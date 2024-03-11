// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vcr

import (
	"io"
	"net/http"
	"strings"
)

type VCRResponse struct {
	// Body of response
	Body string `yaml:"body"`

	// Response headers
	Header http.Header `yaml:"header"`

	// Response status message
	Status string `yaml:"status"`

	// Response status code
	StatusCode int `yaml:"code"`
}

func NewVCRResponse(response *http.Response) *VCRResponse {
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	return &VCRResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       string(respBody),
	}
}

func (r *VCRResponse) GetHTTPResponse() *http.Response {
	return &http.Response{
		Status:     r.Status,
		StatusCode: r.StatusCode,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     r.Header,
		Body:       io.NopCloser(strings.NewReader(r.Body)),
	}
}
