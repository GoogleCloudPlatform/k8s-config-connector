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

package mockcontainer

import (
	"context"
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

func PtrTo[T any](t T) *T {
	return &t
}

func ValueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}

// AsZonalLink will convert a "location" link to a "zonal" link, if the location is a zone.
// For example, projects/${projectNumber}/locations/us-central1-a/operations/${operationID}
// will be converted to projects/${projectNumber}/zones/us-central1-a/operations/${operationID}
func AsZonalLink(link string) string {
	tokens := strings.Split(link, "/")

	for i := 0; i+1 < len(tokens); i++ {
		switch tokens[i] {
		case "locations":
			location := tokens[i+1]
			if isZone(location) {
				tokens[i] = "zones"
			}
		}
	}

	return strings.Join(tokens, "/")
}

// isZone returns true if the location appears to be a GCP zone (as oppposed to a region)
// The logic is pretty simple right now, based on the number of hyphens.
func isZone(location string) bool {
	tokens := strings.Split(location, "-")
	return len(tokens) == 3
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

// buildSelfLink constructs a full self link (including https://container.googleapis.com/<version>/)
func buildSelfLink(ctx context.Context, fqn string) string {
	version := getAPIVersion(ctx)
	return "https://container.googleapis.com/" + version + "/" + fqn
}

// checkInvalidOSVersion checks the request context metadata for a query parameter
// indicating that the client sent "OS_2022" which is rejected with 400 Bad Request by real GKE.
func checkInvalidOSVersion(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)
	if md != nil {
		values := md.Get(httpmux.MetadataKeyHttpRequestQuery)
		if len(values) > 0 {
			q, err := url.ParseQuery(values[0])
			if err == nil {
				if q.Get("mockgcp-invalid-os-version") == "OS_2022" {
					return status.Errorf(codes.InvalidArgument, "Invalid value at 'node_pool.config.windows_node_config.os_version' (type.googleapis.com/google.container.v1beta1.WindowsNodeConfig.OSVersion), \"OS_2022\"")
				}
			}
		}
	}
	return nil
}
