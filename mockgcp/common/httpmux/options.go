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
	"net/http"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"k8s.io/klog/v2"
)

const MetadataKeyExpires = "x-expires"
const MetadataKeyStatusCode = "x-http-code"

func SetExpiresHeader(ctx context.Context, expiresAt time.Time) {
	expires := expiresAt.UTC().Format(http.TimeFormat)

	if err := grpc.SetHeader(ctx, metadata.Pairs(MetadataKeyExpires, expires)); err != nil {
		klog.Fatalf("error setting x-expires header: %v", err)
	}
}

func GetExpiresHeader(ctx context.Context) (string, bool) {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		klog.Fatalf("Failed to get server metadata from context")
	}

	if vals := md.HeaderMD.Get(MetadataKeyExpires); len(vals) > 0 {
		return vals[0], true
	}
	return "", false
}

func SetStatusCode(ctx context.Context, code int) {
	if err := grpc.SetHeader(ctx, metadata.Pairs(MetadataKeyStatusCode, strconv.Itoa(code))); err != nil {
		klog.Fatalf("error setting x-http-code header: %v", err)
	}
}

func GetStatusCode(ctx context.Context) (int, bool) {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		klog.Fatalf("Failed to get server metadata from context")
	}

	// set x-http-code header
	if vals := md.HeaderMD.Get(MetadataKeyStatusCode); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			klog.Fatalf("error parsing x-http-code %q", vals[0])
		}
		return code, true
	}
	return 0, false
}
