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
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"k8s.io/klog/v2"
)

type VCRRequest struct {
	// Body of request
	Body string `yaml:"body"`

	// Request header
	Header http.Header `yaml:"header"`

	// Request URL
	URL string `yaml:"url"`

	// Request method
	Method string `yaml:"method"`
}

func NewVCRRequest(request *http.Request) *VCRRequest {
	requestBody := &bytes.Buffer{}
	if request.Body != nil {
		requestBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			klog.Fatalf("[VCR] Failed to read request body.")
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	}

	return &VCRRequest{
		Body:   requestBody.String(),
		Header: request.Header,
		URL:    request.URL.String(),
		Method: request.Method,
	}
}

func (r *VCRRequest) GetHTTPRequest() *http.Request {
	httpURL, err := url.Parse(r.URL)
	if err != nil {
		return nil
	}
	return &http.Request{
		Body:   io.NopCloser(strings.NewReader(r.Body)),
		Header: r.Header,
		URL:    httpURL,
		Method: r.Method,
	}
}
