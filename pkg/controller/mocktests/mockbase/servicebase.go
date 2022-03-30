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

package mockbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"k8s.io/klog/v2"
)

type ServiceBase struct {
	ExpectedHost string

	handler Handler

	operationsMutex sync.Mutex
	operations      map[string]*Operation
}

type Handler interface {
	ProcessRequest(request *http.Request) (*http.Response, error)
}

func (s *ServiceBase) Init(expectedHost string, handler Handler) {
	s.ExpectedHost = expectedHost
	s.handler = handler
}

func (s *ServiceBase) RoundTrip(request *http.Request) (*http.Response, error) {
	host := request.Host
	if host != s.ExpectedHost {
		return nil, fmt.Errorf("unexpected host in request, got %q, want %q", host, s.ExpectedHost)
	}

	response, err := s.handler.ProcessRequest(request)
	if err != nil {
		klog.Warningf("error serving request %s %s: %v", request.Method, request.URL, err)
		httpResponse := &http.Response{
			Status:     http.StatusText(http.StatusInternalServerError),
			StatusCode: http.StatusInternalServerError,
		}
		var b bytes.Buffer
		b.WriteString(fmt.Sprintf("error: %v", err))
		httpResponse.Body = io.NopCloser(&b)
		return httpResponse, nil
	}
	return response, nil
}

func (s *ServiceBase) ErrorMethodNotAllowed(request *http.Request) (*http.Response, error) {
	klog.Warningf("unhandled method: %s %s %#v", request.Method, request.URL, request)
	httpResponse := &http.Response{
		Status:     http.StatusText(http.StatusMethodNotAllowed),
		StatusCode: http.StatusMethodNotAllowed,
		Body:       s.EmptyBody(),
	}
	return httpResponse, nil
}

func (s *ServiceBase) ErrorNotFound(request *http.Request) (*http.Response, error) {
	httpResponse := &http.Response{
		Status:     http.StatusText(http.StatusNotFound),
		StatusCode: http.StatusNotFound,
		Body:       s.EmptyBody(),
	}
	return httpResponse, nil
}

func (s *ServiceBase) ErrorForbidden(request *http.Request) (*http.Response, error) {
	httpResponse := &http.Response{
		Status:     http.StatusText(http.StatusForbidden),
		StatusCode: http.StatusForbidden,
		Body:       s.EmptyBody(),
	}
	return httpResponse, nil
}

func (s *ServiceBase) ErrorBadRequest(request *http.Request, message string) (*http.Response, error) {
	klog.Warningf("bad request: %s %s %s", request.Method, request.URL, message)
	httpResponse := &http.Response{
		Status:     http.StatusText(http.StatusBadRequest),
		StatusCode: http.StatusBadRequest,
		Body:       s.EmptyBody(),
	}
	return httpResponse, nil
}

func (s *ServiceBase) ReplyJSON(data interface{}) (*http.Response, error) {
	httpResponse := &http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error building json: %w", err)
	}
	klog.Infof("sending response %s", string(b))
	httpResponse.Header = make(http.Header)
	httpResponse.Header.Add("Content-Type", "application/json; charset=UTF-8")
	httpResponse.Header.Add("Cache-Control", "private")
	httpResponse.Body = io.NopCloser(bytes.NewReader(b))
	return httpResponse, nil
}

func (s *ServiceBase) ParseRequest(r *http.Request, request interface{}) *http.Response {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return &http.Response{
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest,
			Body:       s.EmptyBody(),
		}
	}
	if json.Unmarshal(b, request); err != nil {
		return &http.Response{
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest,
			Body:       s.EmptyBody(),
		}
	}
	return nil
}

func (s *ServiceBase) EmptyBody() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte{}))
}
