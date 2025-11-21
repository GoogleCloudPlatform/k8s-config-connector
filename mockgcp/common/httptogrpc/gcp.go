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
	"context"
	"net/http"

	"google.golang.org/protobuf/proto"
)

// addGCPHeaders adds standard GCP headers to the HTTP response.
// If we made this GRPC gateway general-purpose, we could make this configurable per-service.
func (m *grpcMux) addGCPHeaders(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if w.Header().Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
	w.Header().Set("Cache-Control", "private")
	w.Header().Set("Server", "ESF")
	w.Header()["Vary"] = []string{"Origin", "X-Origin", "Referer"}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Xss-Protection", "0")

	return nil
}

// httpErrorResponse is the structure of a GCP error response served over HTTP.
type httpErrorResponse struct {
	Error *httpError `json:"error,omitempty"`
}

type httpError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}
