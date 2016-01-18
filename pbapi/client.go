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
	"encoding/json"
	"log"
	"net/http"
)

func doQuery(req *Request, result interface{}) error {
	uri, err := req.GetRequestURI()
	if err != nil {
		return err
	}

	log.Println("Getting: " + uri.String())
	resp, err := http.Get(uri.String())
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(result)
	return nil
}

func QueryImageDetails(req *Request) (*ImageDetailsResponse, error) {
	req.ResponseGroup = ResponseGroupImageDetails

	result := &ImageDetailsResponse{}
	if err := doQuery(req, result); err != nil {
		return nil, err
	}
	return result, nil
}

func QueryHighResolution(req *Request) (*HighResolutionResponse, error) {
	req.ResponseGroup = ResponseGroupHighResolution

	result := &HighResolutionResponse{}
	if err := doQuery(req, result); err != nil {
		return nil, err
	}
	return result, nil
}
