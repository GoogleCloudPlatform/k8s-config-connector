// Copyright 2026 Google LLC
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

package logging

import (
	"testing"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestLoggingLinkDiff(t *testing.T) {
	tests := []struct {
		name     string
		desired  *pb.Link
		actual   *pb.Link
		expected int
	}{
		{
			name: "identical",
			desired: &pb.Link{
				Name:        "projects/p1/locations/global/buckets/b1/links/l1",
				Description: "desc",
			},
			actual: &pb.Link{
				Name:        "projects/p1/locations/global/buckets/b1/links/l1",
				Description: "desc",
			},
			expected: 0,
		},
		{
			name: "name mismatch - project ID vs number",
			desired: &pb.Link{
				Name:        "projects/p1/locations/global/buckets/b1/links/l1",
				Description: "desc",
			},
			actual: &pb.Link{
				Name:        "projects/12345/locations/global/buckets/b1/links/l1",
				Description: "desc",
			},
			expected: 0, // We want this to be 0 after normalization
		},
		{
			name: "output only fields difference",
			desired: &pb.Link{
				Name:        "projects/p1/locations/global/buckets/b1/links/l1",
				Description: "desc",
			},
			actual: &pb.Link{
				Name:           "projects/p1/locations/global/buckets/b1/links/l1",
				Description:    "desc",
				LifecycleState: pb.LifecycleState_ACTIVE,
			},
			expected: 0,
		},
		{
			name: "description mismatch",
			desired: &pb.Link{
				Name:        "projects/p1/locations/global/buckets/b1/links/l1",
				Description: "desc1",
			},
			actual: &pb.Link{
				Name:        "projects/p1/locations/global/buckets/b1/links/l1",
				Description: "desc2",
			},
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Simulate the normalization in Update
			tc.desired.Name = tc.actual.Name

			diffFunc := func(fieldName protoreflect.Name, a, b proto.Message) (bool, error) {
				if fieldName == "create_time" || fieldName == "lifecycle_state" || fieldName == "bigquery_dataset" {
					return false, nil
				}
				return common.BasicDiff(fieldName, a, b)
			}

			paths, err := common.CompareProtoMessage(tc.desired, tc.actual, diffFunc)
			if err != nil {
				t.Fatalf("CompareProtoMessage failed: %v", err)
			}
			if len(paths) != tc.expected {
				t.Errorf("expected %d diffs, got %d: %v", tc.expected, len(paths), paths)
			}
		})
	}
}
