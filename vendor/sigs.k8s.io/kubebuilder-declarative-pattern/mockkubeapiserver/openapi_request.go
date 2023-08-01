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

	"k8s.io/klog/v2"
)

// openapiRequest is a request to patch a single resource
type openapiRequest struct {
	resourceRequestBase
}

// Run serves the http request
func (req *openapiRequest) Run(ctx context.Context, s *MockKubeAPIServer) error {
	klog.Warningf("returning empty information for openapi/v2 request")

	b := []byte{}
	w := req.baseRequest.w
	w.Header().Add("Content-Type", "application/com.github.proto-openapi.spec.v2")
	w.Header().Add("Cache-Control", "no-cache, private")

	if _, err := w.Write(b); err != nil {
		// Too late to send error response
		klog.Warningf("error writing http response: %v", err)
		return nil
	}
	return nil
}
