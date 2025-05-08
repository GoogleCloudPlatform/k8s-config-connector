// Copyright 2025 Google LLC
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

package httptogrpc

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

type httpMethodCall struct {
	parent *grpcMux
	r      *http.Request
	w      http.ResponseWriter
}

func (c *httpMethodCall) SendErrorResponse(err error) {
	ctx := c.r.Context()

	klog.Warningf("sending error response for %T %+v", err, err)

	statusErr, ok := status.FromError(err)
	if ok {
		response := statusErr.Proto()

		httpCode := http.StatusInternalServerError
		switch statusErr.Code() {
		case codes.NotFound:
			httpCode = http.StatusNotFound
		}

		c.w.Header().Set("Content-Type", "application/json")

		c.parent.addGCPHeaders(ctx, c.w, response)

		body, err := protojson.Marshal(response)
		if err != nil {
			klog.Errorf("failed to marshal error: %v", err)
			http.Error(c.w, "internal error", http.StatusInternalServerError)
			return
		}

		c.w.WriteHeader(httpCode)
		if _, err := c.w.Write(body); err != nil {
			klog.Errorf("failed to write error: %v", err)
		}
		klog.Infof("sent response %v with body %v", httpCode, string(body))
		return
	}
	klog.Warningf("stub-handling error %v", err)
	http.Error(c.w, err.Error(), http.StatusInternalServerError)
}

func (c *httpMethodCall) SendResponse(response proto.Message) {
	ctx := c.r.Context()

	httpCode := http.StatusOK

	c.w.Header().Set("Content-Type", "application/json")

	c.parent.addGCPHeaders(ctx, c.w, response)

	body, err := protojson.Marshal(response)
	if err != nil {
		klog.Errorf("failed to marshal response: %v", err)
		http.Error(c.w, "internal error", http.StatusInternalServerError)
		return
	}

	c.w.WriteHeader(httpCode)
	if _, err := c.w.Write(body); err != nil {
		klog.Errorf("failed to write error: %v", err)
	}
	klog.Infof("sent response %v with body %v", httpCode, string(body))
}
