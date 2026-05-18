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
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type ServerMetadata struct {
	HeaderMD  metadata.MD
	TrailerMD metadata.MD
}

type serverMetadataKey struct{}

func NewContextWithServerMetadata(ctx context.Context, md *ServerMetadata) context.Context {
	return context.WithValue(ctx, serverMetadataKey{}, md)
}

func ServerMetadataFromContext(ctx context.Context) (*ServerMetadata, bool) {
	md, ok := ctx.Value(serverMetadataKey{}).(*ServerMetadata)
	return md, ok
}

const MetadataKeyExpires = "x-expires"
const MetadataKeyStatusCode = "x-http-code"

func SetExpiresHeader(ctx context.Context, expiresAt time.Time) {
	expires := expiresAt.UTC().Format(http.TimeFormat)

	if err := grpc.SetHeader(ctx, metadata.Pairs(MetadataKeyExpires, expires)); err != nil {
		// klog.Fatalf("error setting x-expires header: %v", err)
	}
}

func SetStatusCode(ctx context.Context, code int) {
	if err := grpc.SetHeader(ctx, metadata.Pairs(MetadataKeyStatusCode, strconv.Itoa(code))); err != nil {
		// klog.Fatalf("error setting x-http-code header: %v", err)
	}
}

func GetExpiresHeader(ctx context.Context) (string, bool) {
	md, ok := ServerMetadataFromContext(ctx)
	if !ok {
		return "", false
	}

	if vals := md.HeaderMD.Get(MetadataKeyExpires); len(vals) > 0 {
		return vals[0], true
	}
	return "", false
}

func GetStatusCode(ctx context.Context) (int, bool) {
	md, ok := ServerMetadataFromContext(ctx)
	if !ok {
		return 0, false
	}

	// set x-http-code header
	if vals := md.HeaderMD.Get(MetadataKeyStatusCode); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return 0, false
		}
		return code, true
	}
	return 0, false
}

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

	if m.RewriteHeaders != nil {
		m.RewriteHeaders(ctx, w, resp)
	}
	return nil
}

// HTTPErrorResponse is the structure of a GCP error response served over HTTP.
type HTTPErrorResponse struct {
	Error *HTTPError `json:"error,omitempty"`
}

type HTTPError struct {
	Code    int                  `json:"code,omitempty"`
	Message string               `json:"message,omitempty"`
	Status  string               `json:"status,omitempty"`
	Errors  []HTTPErrorDetails `json:"errors,omitempty"`
}

type HTTPErrorDetails struct {
	Domain  string `json:"domain,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}
