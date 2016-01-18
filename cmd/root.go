// Copyright © 2016 Kim Eik
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	"github.com/netbrain/gopixabay/pbapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

var cfgFile string

const (
	og = "og"
	lg = "lg"
	md = "md"
	sm = "sm"
	xs = "xs"
)

type Arguments struct {
	request      *pbapi.Request
	size         string
	simultaneous int
}

var arguments = &Arguments{
	request:      &pbapi.Request{},
	simultaneous: 4,
}

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gopixabay /path/to/output",
	Short: "gopixabay image downloader",
	Long:  `gopixabay is an image downloader using the pixabay API`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		validateSizeArgument()

		var result interface{}
		var err error

		switch arguments.request.ResponseGroup {
		case pbapi.ResponseGroupHighResolution:
			result, err = pbapi.QueryHighResolution(arguments.request)
		case pbapi.ResponseGroupImageDetails:
			result, err = pbapi.QueryImageDetails(arguments.request)
		default:
			log.Fatalf("Unknown response group: %s", arguments.request.ResponseGroup)
		}

		if err != nil {
			log.Fatal(err)
		}

		uiprogress.Start()
		wg := &sync.WaitGroup{}
		downloadChan := make(chan string)
		defer close(downloadChan)

		initWorkerThreads(args[0], wg, downloadChan)

		switch realResult := result.(type) {
		case *pbapi.HighResolutionResponse:
			for _, hit := range realResult.Hits {
				switch arguments.size {
				case og:
					downloadChan <- hit.ImageURL
				case lg:
					downloadChan <- hit.FullHDURL
				case md:
					downloadChan <- hit.LargeImageURL
				}
			}
		case *pbapi.ImageDetailsResponse:
			for _, hit := range realResult.Hits {
				switch arguments.size {
				case sm:
					downloadChan <- hit.WebFormatURL
				case xs:
					downloadChan <- hit.PreviewURL
				}
			}
		}

		wg.Wait()
		uiprogress.Stop()
		time.Sleep(time.Second)
	},
}

func validateSizeArgument() {
	switch arguments.size {
	case og:
		fallthrough
	case lg:
		fallthrough
	case md:
		if arguments.request.ResponseGroup != pbapi.ResponseGroupHighResolution {
			log.Fatalf("need %s response group setting to get this image size", pbapi.ResponseGroupHighResolution)
		}
	case sm:
		fallthrough
	case xs:
		if arguments.request.ResponseGroup != pbapi.ResponseGroupImageDetails {
			log.Fatalf("need %s response group setting to get this image size", pbapi.ResponseGroupImageDetails)
		}
	default:
		log.Fatalf("Unknown image size: %s", arguments.size)
	}
}

func initWorkerThreads(outputPath string, wg *sync.WaitGroup, downloadChan <-chan string) {
	for x := 0; x < arguments.simultaneous; x++ {
		go func() {
			for {
				url, more := <-downloadChan
				wg.Add(1)
				if more {
					func() {
						defer wg.Done()
						err := downloadImage(url, outputPath)
						if err != nil && err != io.EOF {
							log.Println(err)
						}
					}()
				} else {
					return
				}

			}
		}()
	}
}

func downloadImage(u, output string) error {
	url, err := url.Parse(u)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(url.Path)
	out, err := os.Create(filepath.Join(output, filename))
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	sum := 0

	bar := uiprogress.AddBar(int(resp.ContentLength)).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return strutil.Resize(humanize.Bytes(uint64(sum))+"/"+humanize.Bytes(uint64(resp.ContentLength))+" "+filename+": ", 42)
	})

	for {
		written, err := io.CopyN(out, resp.Body, 4096)
		sum += int(written)
		bar.Set(sum)
		if err != nil {
			return err
		}
	}
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gopixabay.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().StringVarP(&arguments.request.Key, "key", "k", "", "Your pixabay API key")
	RootCmd.Flags().StringVarP(&arguments.request.ResponseGroup, "response-group", "g", "image_details", "Select the information returned: image details or URLs of high resolution images. High resolution image URLs and details cannot be obtained in a single request. (image_details, high_resolution)")
	RootCmd.Flags().StringVar(&arguments.request.ID, "id", "", "ID, ID hash, or a comma separated list of values for fetching or updating specific images. In a comma separated list, IDs and ID hashs cannot be used together.")
	RootCmd.Flags().StringVarP(&arguments.request.Query, "query", "q", "", "A text query to use when searching for images")
	RootCmd.Flags().StringVarP(&arguments.request.Lang, "lang", "l", "en", "Language code of the language to be searched in. ")
	RootCmd.Flags().StringVarP(&arguments.request.ImageType, "image-type", "t", "all", "A media type to search within. (all, photo, illustration, vector)")
	RootCmd.Flags().StringVarP(&arguments.request.Orientation, "orientation", "o", "all", "Whether an image is wider than it is tall, or taller than it is wide. (all, horizontal, vertical)")
	RootCmd.Flags().StringVarP(&arguments.request.Category, "category", "c", "", "Filter images by category. (fashion, nature, backgrounds, science, education, people, feelings, religion, health, places, animals, industry, food, computer, sports, transportation, travel, buildings, business, music)")
	RootCmd.Flags().IntVar(&arguments.request.MinWidth, "min-width", 0, "Minimum image width. ")
	RootCmd.Flags().IntVar(&arguments.request.MinWidth, "min-height", 0, "Minimum image height. ")
	RootCmd.Flags().BoolVarP(&arguments.request.EditorsChoice, "editors-choice", "e", false, "Select images that have received an Editor's Choice award. ")
	RootCmd.Flags().BoolVarP(&arguments.request.SafeSearch, "safesearch", "s", false, "A flag indicating that only images suitable for all ages should be returned. ")
	RootCmd.Flags().StringVarP(&arguments.request.Order, "order", "r", "popular", "How the results should be ordered. (popular, latest)")
	RootCmd.Flags().IntVarP(&arguments.request.Page, "page", "p", 1, "Returned search results are paginated. Use this parameter to select the page number. ")
	RootCmd.Flags().IntVar(&arguments.request.PerPage, "per-page", 20, "Determine the number of results per page.")
	RootCmd.Flags().StringVar(&arguments.request.Callback, "callback", "", "JSONP callback function name")
	RootCmd.Flags().BoolVar(&arguments.request.Pretty, "pretty", false, "Prettify and indent JSON output. This option is for development puposes only and should not be used in production. ")

	RootCmd.Flags().StringVar(&arguments.size, "size", "og", "Image size to download (og, lg, md, sm, xs) ")
	RootCmd.Flags().IntVar(&arguments.simultaneous, "num", 4, "Number of images to download in paralell")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gopixabay") // name of config file (without extension)
	viper.AddConfigPath("$HOME")       // adding home directory as first search path
	viper.AutomaticEnv()               // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
