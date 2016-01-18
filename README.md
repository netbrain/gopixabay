# gopixabay


<a href="https://pixabay.com/">
  <img src="https://pixabay.com/static/img/public/medium_rectangle_a.png" alt="Pixabay">
</a>


Have you ever needed a bunch of images with specific content, size or resolution? 
This, and http://pixabay.com can help with that!

gopixabay be used as an api library or as a standalone cli application.

## CLI

```
$ gopixabay -k <insert api key> -t photo --per-page 10 --size og -g high_resolution --min-width 1920 --min-height 1080 /tmp

2016/01/18 08:22:45 Getting: https://pixabay.com/api/?...
2.6 MB/2.6 MB e834b0062bf4003ecd0b440de...   --- [====================================================================] 100%
2.6 MB/2.6 MB e834b0062bf4003ecd0b440de...   --- [====================================================================] 100%
2.3 MB/2.3 MB e834b2062df4093ecd0b440de...   --- [====================================================================] 100%
968 kB/968 kB e834b20929f0063ecd0b440de...   --- [====================================================================] 100%
1.0 MB/1.0 MB e834b2062bf1003ecd0b440de...   --- [====================================================================] 100%
6.6 MB/6.6 MB e834b20f2bfc083ecd0b440de...   --- [====================================================================] 100%
4.2 MB/4.2 MB e834b20928fd023ecd0b440de...   --- [====================================================================] 100%
5.3 MB/5.3 MB e834b20b2dfd093ecd0b440de...   --- [====================================================================] 100%
505 kB/505 kB e834b20929f7003ecd0b440de...   --- [====================================================================] 100%
4.5 MB/4.5 MB e834b20b2ef4043ecd0b440de...   --- [====================================================================] 100%
4.5 MB/4.5 MB e834b2062bfd093ecd0b440de...   --- [====================================================================] 100%
```

### Usage and flags

```
gopixabay is an image downloader using the pixabay API

Usage:
  gopixabay /path/to/output [flags]

Flags:
      --callback string         JSONP callback function name
  -c, --category string         Filter images by category. (fashion, nature, backgrounds, science, education, people, feelings, religion, health, places, animals, industry, food, computer, sports, transportation, travel, buildings, business, music)
      --config string           config file (default is $HOME/.gopixabay.yaml)
  -e, --editors-choice          Select images that have received an Editor's Choice award. 
  -h, --help                    help for gopixabay
      --id string               ID, ID hash, or a comma separated list of values for fetching or updating specific images. In a comma separated list, IDs and ID hashs cannot be used together.
  -t, --image-type string       A media type to search within. (all, photo, illustration, vector) (default "all")
  -k, --key string              Your pixabay API key
  -l, --lang string             Language code of the language to be searched in.  (default "en")
      --min-height int          Minimum image height. 
      --min-width int           Minimum image width. 
      --num int                 Number of images to download in paralell (default 4)
  -r, --order string            How the results should be ordered. (popular, latest) (default "popular")
  -o, --orientation string      Whether an image is wider than it is tall, or taller than it is wide. (all, horizontal, vertical) (default "all")
  -p, --page int                Returned search results are paginated. Use this parameter to select the page number.  (default 1)
      --per-page int            Determine the number of results per page. (default 20)
      --pretty                  Prettify and indent JSON output. This option is for development puposes only and should not be used in production. 
  -q, --query string            A text query to use when searching for images
  -g, --response-group string   Select the information returned: image details or URLs of high resolution images. High resolution image URLs and details cannot be obtained in a single request. (image_details, high_resolution) (default "image_details")
  -s, --safesearch              A flag indicating that only images suitable for all ages should be returned. 
      --size string             Image size to download (og, lg, md, sm, xs)  (default "og")
```
