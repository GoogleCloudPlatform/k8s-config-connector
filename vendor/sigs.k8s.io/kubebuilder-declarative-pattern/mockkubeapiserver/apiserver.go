/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockkubeapiserver

import (
	"net"
	"net/http"
	"strings"

	"k8s.io/klog/v2"
)

func NewMockKubeAPIServer(addr string) (*MockKubeAPIServer, error) {
	s := &MockKubeAPIServer{}
	if addr == "" {
		addr = ":http"
	}

	s.httpServer = &http.Server{Addr: addr, Handler: s}

	s.storage = NewMemoryStorage()

	return s, nil
}

type MockKubeAPIServer struct {
	httpServer *http.Server
	listener   net.Listener

	storage *MemoryStorage
}

func (s *MockKubeAPIServer) StartServing() (net.Addr, error) {
	listener, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return nil, err
	}
	s.listener = listener
	addr := listener.Addr()
	go func() {
		if err := s.httpServer.Serve(s.listener); err != nil {
			if err != http.ErrServerClosed {
				klog.Errorf("error serving: %v", err)
			}
		}
	}()
	return addr, nil
}

func (s *MockKubeAPIServer) Stop() error {
	return s.httpServer.Close()
}

func (s *MockKubeAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	klog.Infof("kubeapiserver request: %s %s", r.Method, r.URL)

	path := r.URL.Path
	tokens := strings.Split(strings.Trim(path, "/"), "/")

	var req Request

	// matchedPath is bool if we recognized the path, but if we didn't build a req we should send StatusMethodNotAllowed instead of NotFound
	var matchedPath bool

	if len(tokens) == 2 {
		if tokens[0] == "api" && tokens[1] == "v1" {
			matchedPath = true

			switch r.Method {
			case http.MethodGet:
				req = &apiResourceList{
					Group:   "",
					Version: "v1",
				}
			}
		}

		if tokens[0] == "openapi" && tokens[1] == "v2" {
			matchedPath = true

			switch r.Method {
			case http.MethodGet:
				req = &openapiRequest{}
			}
		}
	}

	if len(tokens) == 1 {
		if tokens[0] == "api" {
			matchedPath = true

			switch r.Method {
			case http.MethodGet:
				req = &apiVersionsRequest{}
			}
		}

		if tokens[0] == "apis" {
			matchedPath = true
			switch r.Method {
			case http.MethodGet:
				req = &apiGroupList{}
			}
		}
	}

	if len(tokens) == 3 {
		if tokens[0] == "apis" {
			matchedPath = true
			switch r.Method {
			case http.MethodGet:
				req = &apiResourceList{
					Group:   tokens[1],
					Version: tokens[2],
				}
			}
		}

		if tokens[0] == "api" && tokens[1] == "v1" {
			matchedPath = true
			switch r.Method {
			case http.MethodPost:
				req = &postResource{
					Group:     "",
					Version:   tokens[1],
					Resource:  tokens[2],
					Namespace: "",
				}
			case http.MethodGet:
				req = &listResource{
					Group:     "",
					Version:   tokens[1],
					Resource:  tokens[2],
					Namespace: "",
				}
			}
		}
	}

	buildObjectRequest := func(common resourceRequestBase) {
		switch r.Method {
		case http.MethodGet:
			req = &getResource{
				resourceRequestBase: common,
			}
		case http.MethodPatch:
			req = &patchResource{
				resourceRequestBase: common,
			}
		case http.MethodPut:
			req = &putResource{
				resourceRequestBase: common,
			}
		case http.MethodDelete:
			req = &deleteResource{
				resourceRequestBase: common,
			}
		}
	}

	if len(tokens) == 4 {
		if tokens[0] == "api" {
			buildObjectRequest(resourceRequestBase{
				Group:    "",
				Version:  tokens[1],
				Resource: tokens[2],
				Name:     tokens[3],
			})
			matchedPath = true
		}

		if tokens[0] == "apis" {
			matchedPath = true
			switch r.Method {
			case http.MethodPost:
				req = &postResource{
					Group:     tokens[1],
					Version:   tokens[2],
					Resource:  tokens[3],
					Namespace: "",
				}
			case http.MethodGet:
				req = &listResource{
					Group:     tokens[1],
					Version:   tokens[2],
					Resource:  tokens[3],
					Namespace: "",
				}
			}
		}
	}

	if len(tokens) == 5 {
		if tokens[0] == "api" && tokens[1] == "v1" && tokens[2] == "namespaces" {
			matchedPath = true
			switch r.Method {
			case http.MethodPost:
				req = &postResource{
					Group:     "",
					Version:   tokens[1],
					Resource:  tokens[4],
					Namespace: tokens[3],
				}
			case http.MethodGet:
				req = &listResource{
					Group:     "",
					Version:   tokens[1],
					Resource:  tokens[4],
					Namespace: tokens[3],
				}
			}
		}
	}

	if len(tokens) == 6 {
		if tokens[0] == "api" && tokens[2] == "namespaces" {
			buildObjectRequest(resourceRequestBase{
				Group:     "",
				Version:   tokens[1],
				Resource:  tokens[4],
				Namespace: tokens[3],
				Name:      tokens[5],
			})
			matchedPath = true
		}

		if tokens[0] == "apis" && tokens[3] == "namespaces" {
			matchedPath = true
			switch r.Method {
			case http.MethodPost:
				req = &postResource{
					Group:     tokens[1],
					Version:   tokens[2],
					Resource:  tokens[5],
					Namespace: tokens[4],
				}
			case http.MethodGet:
				req = &listResource{
					Group:     tokens[1],
					Version:   tokens[2],
					Resource:  tokens[5],
					Namespace: tokens[4],
				}
			}
		}
	}
	if len(tokens) == 7 {
		if tokens[0] == "apis" && tokens[3] == "namespaces" {
			buildObjectRequest(resourceRequestBase{
				Group:     tokens[1],
				Version:   tokens[2],
				Namespace: tokens[4],
				Resource:  tokens[5],
				Name:      tokens[6],
			})
			matchedPath = true
		}
	}
	if len(tokens) == 8 {
		if tokens[0] == "apis" && tokens[3] == "namespaces" {
			buildObjectRequest(resourceRequestBase{
				Group:       tokens[1],
				Version:     tokens[2],
				Namespace:   tokens[4],
				Resource:    tokens[5],
				Name:        tokens[6],
				SubResource: tokens[7],
			})
			matchedPath = true
		}
	}

	if req == nil {
		if matchedPath {
			klog.Warningf("method not allowed for %s %s", r.Method, r.URL)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		} else {
			klog.Warningf("404 for %s %s tokens=%#v", r.Method, r.URL, tokens)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		return
	}

	req.Init(w, r)

	err := req.Run(ctx, s)
	if err != nil {
		klog.Warningf("internal error for %s %s: %v", r.Method, r.URL, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
