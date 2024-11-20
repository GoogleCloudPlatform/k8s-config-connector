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

package httpmux

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

type wrappedStatus struct {
	Error *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int                    `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`
	Status  string                 `json:"status,omitempty"`
	Errors  []ErrorResponseDetails `json:"errors,omitempty"`
}

type ErrorResponseDetails struct {
	Domain  string `json:"domain,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// customErrorHandler wraps errors in an error block
func (m *ServeMux) customErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s := status.Convert(err)
	// pb := s.Proto()

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	httpStatusCode := runtime.HTTPStatusFromCode(s.Code())
	wrapped := &wrappedStatus{
		Error: &ErrorResponse{
			Code:    httpStatusCode,
			Message: s.Message(),
		},
	}

	switch s.Code() {
	case codes.PermissionDenied:
		wrapped.Error.Status = "PERMISSION_DENIED"
	case codes.AlreadyExists:
		wrapped.Error.Status = "ALREADY_EXISTS"
	case codes.NotFound:
		wrapped.Error.Status = "NOT_FOUND"
		wrapped.Error.Errors = append(wrapped.Error.Errors, ErrorResponseDetails{
			Domain:  "global",
			Message: wrapped.Error.Message,
			Reason:  "notFound",
		})
	}

	if m.RewriteError != nil {
		// wrapped.Error is changed in-place
		m.RewriteError(ctx, wrapped.Error)
	}

	buf, merr := json.Marshal(wrapped)
	if merr != nil {
		klog.Warningf("Failed to marshal error message %q: %v", s, merr)
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		return
	}

	if err := m.addGCPHeaders(ctx, w, nil); err != nil {
		klog.Warningf("unexpected error from header filter: %v", err)
	}

	w.WriteHeader(httpStatusCode)
	if _, err := w.Write(buf); err != nil {
		klog.Warningf("Failed to write response: %v", err)
	}

}
