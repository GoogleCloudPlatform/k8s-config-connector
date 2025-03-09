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

package mockcompute

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func computeFingerprint(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func labelsFingerprint(labels map[string]string) string {
	b, err := json.Marshal(labels)
	if err != nil {
		panic(fmt.Sprintf("converting to json: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func ValueOf[T any](t *T) T {
	if t == nil {
		var defaultT T
		return defaultT
	}
	return *t
}

// getAPIVersion returns the version of the compute API the caller is using.
// It defaults to v1
func getAPIVersion(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	path := ""
	if md != nil {
		for _, v := range md.Get("path") {
			path = v
		}
	}
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimPrefix(path, "compute/")
	version, _, _ := strings.Cut(path, "/")
	if version == "" {
		// Default to v1
		version = "v1"
	}
	return version
}

// buildComputeSelfLink constructs a full self link (including https://www.googleapis.com/compute/)
func buildComputeSelfLink(ctx context.Context, fqn string) string {
	version := getAPIVersion(ctx)
	return "https://www.googleapis.com/compute/" + version + "/" + fqn
}

// makeFullyQualifiedRegion will convert a short-form region name to a fully-qualified name
func makeFullyQualifiedRegion(ctx context.Context, projectID string, region string) string {
	s := region
	tokens := strings.Split(s, "/")
	if len(tokens) == 1 {
		s = buildComputeSelfLink(ctx, "projects/"+projectID+"/regions/"+region)
	}
	return s
}
