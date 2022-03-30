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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/mocktests/mocksecretmanager"
)

type mockRoundTripper struct {
	secretmanager mocksecretmanager.MockService
}

func NewMockRoundTripper() *mockRoundTripper {
	rt := &mockRoundTripper{}
	rt.secretmanager = mocksecretmanager.NewMock(mocksecretmanager.ExpectedHost)
	return rt
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("request: %v %v", req.Method, req.URL)

	// TODO: Make this better ... iterate through a list?

	if req.Host == mocksecretmanager.ExpectedHost {
		return m.secretmanager.RoundTrip(req)
	}

	request := fmt.Sprintf("%s %s", req.Method, req.URL)
	body := make(map[string]interface{})

	response := &http.Response{
		StatusCode: 403,
		Status:     "mockRoundTripper injecting fake response",
	}

	if request == "GET https://openidconnect.googleapis.com/v1/userinfo?alt=json" {
		body["email"] = "test@example.com"

		response.StatusCode = 200
	}

	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			panic("json.Marshal failed")
		}

		log.Printf("response: %d %s", response.StatusCode, string(j))

		response.Body = ioutil.NopCloser(bytes.NewReader(j))
	} else {
		log.Printf("response: %d %s", response.StatusCode, "-")
	}

	return response, nil
}
