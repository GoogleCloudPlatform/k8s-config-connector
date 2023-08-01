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
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/klog/v2"
)

type Request interface {
	Run(ctx context.Context, s *MockKubeAPIServer) error
	Init(w http.ResponseWriter, r *http.Request)
}

// baseRequest is the base for our higher-level http requests
type baseRequest struct {
	w http.ResponseWriter
	r *http.Request
}

func (b *baseRequest) Init(w http.ResponseWriter, r *http.Request) {
	b.w = w
	b.r = r
}

func (r *baseRequest) writeResponse(obj interface{}) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("error from json.Marshal on %T: %w", obj, err)
	}
	r.w.Header().Add("Content-Type", "application/json")
	r.w.Header().Add("Cache-Control", "no-cache, private")

	if _, err := r.w.Write(b); err != nil {
		// Too late to send error response
		klog.Warningf("error writing http response: %v", err)
		return nil
	}
	return nil
}

func (r *baseRequest) writeErrorResponse(statusCode int) error {
	klog.Warningf("%d for %s %s", statusCode, r.r.Method, r.r.URL)
	http.Error(r.w, http.StatusText(statusCode), statusCode)

	return nil
}
