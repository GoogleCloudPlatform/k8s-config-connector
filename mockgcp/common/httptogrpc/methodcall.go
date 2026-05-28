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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// httpMethodCall holds state for a single method call over HTTP.
type httpMethodCall struct {
	parent     *grpcMux
	grpcMethod *grpcMethod
	r          *http.Request
	w          http.ResponseWriter

	headerMD  metadata.MD
	trailerMD metadata.MD
}

// SendErrorResponse sends an error response for a GRPC method call over HTTP.
func (c *httpMethodCall) SendErrorResponse(err error) {
	ctx := c.r.Context()
	if c.headerMD != nil || c.trailerMD != nil {
		md := &ServerMetadata{
			HeaderMD:  c.headerMD,
			TrailerMD: c.trailerMD,
		}
		ctx = NewContextWithServerMetadata(ctx, md)
	}

	klog.Warningf("sending error response for %T %+v", err, err)

	// Try to map well known errors to HTTP status codes and structured responses
	statusErr, ok := status.FromError(err)
	if ok {
		response := statusErr.Proto()

		HTTPErrorResponse := &HTTPErrorResponse{
			Error: &HTTPError{
				Code:    http.StatusInternalServerError,
				Message: response.Message,
			},
		}

		switch statusErr.Code() {
		case codes.NotFound:
			HTTPErrorResponse.Error.Code = http.StatusNotFound
			HTTPErrorResponse.Error.Status = "NOT_FOUND"
		case codes.AlreadyExists:
			HTTPErrorResponse.Error.Code = http.StatusConflict
			HTTPErrorResponse.Error.Status = "ALREADY_EXISTS"
		case codes.PermissionDenied:
			HTTPErrorResponse.Error.Code = http.StatusForbidden
			HTTPErrorResponse.Error.Status = "PERMISSION_DENIED"
		case codes.InvalidArgument:
			HTTPErrorResponse.Error.Code = http.StatusBadRequest
			HTTPErrorResponse.Error.Status = "INVALID_ARGUMENT"
		}

		if c.parent.RewriteError != nil {
			c.parent.RewriteError(ctx, HTTPErrorResponse.Error)
		}

		body, err := json.Marshal(HTTPErrorResponse)
		if err != nil {
			klog.Errorf("failed to marshal error: %v", err)
			http.Error(c.w, "internal error", http.StatusInternalServerError)
			return
		}

		c.w.Header().Set("Content-Type", "application/json")

		c.parent.addGCPHeaders(ctx, c.w, response)

		c.w.WriteHeader(HTTPErrorResponse.Error.Code)
		if _, err := c.w.Write(body); err != nil {
			klog.Errorf("failed to write error: %v", err)
		}
		klog.Infof("sent response %v with body %v", HTTPErrorResponse.Error.Code, string(body))
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
	if c.headerMD != nil || c.trailerMD != nil {
		md := &ServerMetadata{
			HeaderMD:  c.headerMD,
			TrailerMD: c.trailerMD,
		}
		ctx = NewContextWithServerMetadata(ctx, md)
	}

	httpCode := http.StatusOK
	if code, found := GetStatusCode(ctx); found {
		httpCode = code
	}

	c.w.Header().Set("Content-Type", "application/json")

	c.parent.addGCPHeaders(ctx, c.w, response)

	marshalOptions := protojson.MarshalOptions{}
	responseOptions.populateMarshalOptions(&marshalOptions)

	if c.grpcMethod != nil {
		if c.grpcMethod.parentService.options.EmitUnpopulated {
			marshalOptions.EmitUnpopulated = true
		}
	}

	if httpCode == http.StatusNoContent {
		c.w.WriteHeader(httpCode)
		return
	}

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
