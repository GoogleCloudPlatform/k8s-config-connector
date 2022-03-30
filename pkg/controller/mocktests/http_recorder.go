// Copyright 2022 Google LLC
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

package mocktests

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Method string      `json:"method,omitempty"`
	URL    string      `json:"url,omitempty"`
	Header http.Header `json:"header,omitempty"`
	Body   string      `json:"body,omitempty"`
}

type HTTPRecorder struct {
	inner    http.RoundTripper
	Requests []Request
}

func NewRecorder(inner http.RoundTripper) *HTTPRecorder {
	rt := &HTTPRecorder{inner: inner}
	return rt
}

func (m *HTTPRecorder) RoundTrip(req *http.Request) (*http.Response, error) {

	c := Request{
		Method: req.Method,
		URL:    req.URL.String(),
		Header: req.Header,
	}

	if req.Body != nil {
		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic("failed to read request body")
		}
		c.Body = string(requestBody)
		req.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	}
	m.Requests = append(m.Requests, c)

	return m.inner.RoundTrip(req)
}
