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
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// httpMethodCall holds state for a single method call over HTTP.
type httpMethodCall struct {
	parent *grpcMux
	r      *http.Request
	w      http.ResponseWriter
}

// SendErrorResponse sends an error response for a GRPC method call over HTTP.
func (c *httpMethodCall) SendErrorResponse(err error) {
	ctx := c.r.Context()

	klog.Warningf("sending error response for %T %+v", err, err)

	// Try to map well known errors to HTTP status codes and structured responses
	statusErr, ok := status.FromError(err)
	if ok {
		response := statusErr.Proto()

		httpErrorResponse := &httpErrorResponse{
			Error: &httpError{
				Code:    http.StatusInternalServerError,
				Message: response.Message,
			},
		}

		switch statusErr.Code() {
		case codes.NotFound:
			httpErrorResponse.Error.Code = http.StatusNotFound
			httpErrorResponse.Error.Status = "NOT_FOUND"
		}

		body, err := json.Marshal(httpErrorResponse)
		if err != nil {
			klog.Errorf("failed to marshal error: %v", err)
			http.Error(c.w, "internal error", http.StatusInternalServerError)
			return
		}

		c.w.Header().Set("Content-Type", "application/json")

		c.parent.addGCPHeaders(ctx, c.w, response)

		c.w.WriteHeader(httpErrorResponse.Error.Code)
		if _, err := c.w.Write(body); err != nil {
			klog.Errorf("failed to write error: %v", err)
		}
		klog.Infof("sent response %v with body %v", httpErrorResponse.Error.Code, string(body))
		return
	}
	klog.Warningf("stub-handling error %v", err)
	http.Error(c.w, err.Error(), http.StatusInternalServerError)
}

// ResponseOptions holds options for sending a response.
type ResponseOptions struct {
	Alt []string
}

// populateMarshalOptions populates marshal options based on the options set.
func (o *ResponseOptions) populateMarshalOptions(marshalOptions *protojson.MarshalOptions) {
	if o == nil {
		return
	}
	for _, alt := range o.Alt {
		switch alt {
		case "json;enum-encoding=int":
			marshalOptions.UseEnumNumbers = true
		case "json":
			// Default behavior
		default:
			klog.Fatalf("unhandled alt option: %q", alt)
		}
	}
}

// SendResponse sends a successful response for a GRPC method call over HTTP.
func (c *httpMethodCall) SendResponse(response proto.Message, responseOptions ResponseOptions) {
	ctx := c.r.Context()

	httpCode := http.StatusOK

	c.w.Header().Set("Content-Type", "application/json")

	c.parent.addGCPHeaders(ctx, c.w, response)

	marshalOptions := protojson.MarshalOptions{}
	responseOptions.populateMarshalOptions(&marshalOptions)

	body, err := marshalOptions.Marshal(response)
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
