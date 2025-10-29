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

package mocknetworkservices

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// ProtoClone is a type-safe wrapper around proto.Clone.
func ProtoClone[T proto.Message](obj T) T {
	return proto.Clone(obj).(T)
}

// getAPIVersion returns the version of the API the caller is using.
// It defaults to v1beta1
func getAPIVersion(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	path := ""
	if md != nil {
		for _, v := range md.Get("path") {
			path = v
		}
	}
	path = strings.TrimPrefix(path, "/")
	version, _, _ := strings.Cut(path, "/")
	if version == "" {
		// We could default to v1beta1, but because this is a test we panic instead
		klog.Fatalf("could not determine API version from path %q", path)
	}
	if version != "v1beta1" && version != "v1" {
		// This does not look like an api version
		klog.Fatalf("unexpected API version %q", version)
	}
	return version
}

// buildSelfLink constructs a full self link (including https://networkservices.googleapis.com/<version>/)
func buildSelfLink(ctx context.Context, fqn string) string {

	// Oddly the self link version appears to always be v1alpha1
	// version := getAPIVersion(ctx)
	version := "v1alpha1"

	return "https://networkservices.googleapis.com/" + version + "/" + fqn
}
