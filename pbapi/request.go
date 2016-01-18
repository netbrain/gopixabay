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

package pbapi

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

//API constant
const APIURL = "https://pixabay.com/api/"

//Constants for ImageType attribute
const (
	ImageTypeAll          = "all"
	ImageTypePhoto        = "photo"
	ImageTypeIllustration = "illustration"
	ImageTypeVector       = "vector"
)

//Constants for ResponseGroup attribute
const (
	ResponseGroupImageDetails   = "image_details"
	ResponseGroupHighResolution = "high_resolution"
)

//Constants for Orientation attribute
const (
	OrientationAll        = "all"
	OrientationHorizontal = "horizontal"
	OrientationVertical   = "vertical"
)

//Constants for Category attribute
const (
	CategoryFashion        = "fashion"
	CategoryNature         = "nature"
	CategoryBackgrounds    = "backgrounds"
	CategoryScience        = "science"
	CategoryEducation      = "education"
	CategoryPeople         = "people"
	CategoryFeelings       = "feelings"
	CategoryReligion       = "religion"
	CategoryHealth         = "health"
	CategoryPlaces         = "places"
	CategoryAnimals        = "animals"
	CategoryIndustry       = "industry"
	CategoryFood           = "food"
	CategoryComputer       = "computer"
	CategorySports         = "sports"
	CategoryTransportation = "transportation"
	CategoryTravel         = "travel"
	CategoryBuildings      = "buildings"
	CategoryBusiness       = "business"
	CategoryMusic          = "music"
)

//Constants for Order attibute
const (
	OrderPopular = "popular"
	OrderLatest  = "latest"
)

//Request payload
type Request struct {
	Key           string `url:"key"`
	ResponseGroup string `url:"response_group"`
	ID            string `url:"id"`
	Query         string `url:"q"`
	Lang          string `url:"lang"`
	ImageType     string `url:"image_type"`
	Orientation   string `url:"orientation"`
	Category      string `url:"category"`
	MinWidth      int    `url:"min_width"`
	MinHeight     int    `url:"min_height"`
	EditorsChoice bool   `url:"editors_choice"`
	SafeSearch    bool   `url:"safesearch"`
	Order         string `url:"order"`
	Page          int    `url:"page"`
	PerPage       int    `url:"per_page"`
	Callback      string `url:"callback"`
	Pretty        bool   `url:"pretty"`
}

//NewRequest creates a new request
func NewRequest(key string) *Request {
	return &Request{
		Key:           key,
		ResponseGroup: ResponseGroupImageDetails,
		Lang:          "en",
		ImageType:     ImageTypeAll,
		Orientation:   OrientationAll,
		MinWidth:      0,
		MinHeight:     0,
		EditorsChoice: false,
		SafeSearch:    false,
		Order:         OrderPopular,
		Page:          1,
		PerPage:       20,
		Pretty:        false,
	}
}

func (r *Request) GetRequestURI() (*url.URL, error) {
	queryParams, err := query.Values(r)
	if err != nil {
		return nil, err
	}
	return url.ParseRequestURI(fmt.Sprintf("%s?%s", APIURL, queryParams.Encode()))
}
