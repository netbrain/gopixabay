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

type BaseResponse struct {
	Total     int `json:"total"`
	TotalHits int `json:"totalHits"`
}

type ImageDetailsResponse struct {
	BaseResponse
	Hits []*ImageDetailsHitResponse `json:"hits"`
}

type HighResolutionResponse struct {
	BaseResponse
	Hits []*HighResolutionHitResponse `json:"hits"`
}

type baseHitResponse struct {
	Type            string `json:"type"`
	User            string `json:"user"`
	UserID          int    `json:"user_id"`
	PreviewURL      string `json:"previewURL"`
	PreviewWidth    int    `json:"previewWidth"`
	PreviewHeight   int    `json:"previewHeight"`
	ImageWidth      int    `json:"imageWidth"`
	ImageHeight     int    `json:"imageHeight"`
	WebFormatURL    string `json:"webformatURL"`
	WebFormatHeight int    `json:"webformatHeight"`
	WebFormatWidth  int    `json:"webformatWidth"`
}

type HighResolutionHitResponse struct {
	baseHitResponse
	ID            string `json:"id_hash"`
	LargeImageURL string `json:"largeImageURL"`
	ImageURL      string `json:"imageURL"`
	FullHDURL     string `json:"fullHDURL"`
	UserImageURL  string `json:"userImageURL"`
}

type ImageDetailsHitResponse struct {
	baseHitResponse
	ID        int    `json:"id"`
	Favorites int    `json:"favorites"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
	Views     int    `json:"views"`
	Tags      string `json:"tags"`
	Downloads int    `json:"downloads"`
	PageURL   string `json:"pageURL"`
}
