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

package bigquerydataset

import (
	"testing"

	pb "cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// TestBigQueryDatasetStatus_FromProto_selfLink checks the self-link built from
// FullID, which is where a project ID containing a colon previously caused the
// link to be dropped entirely.
func TestBigQueryDatasetStatus_FromProto_selfLink(t *testing.T) {
	tests := []struct {
		name     string
		fullID   string
		expected string
	}{
		{
			name:     "plainProjectID",
			fullID:   "my-project:my_dataset",
			expected: "https://bigquery.googleapis.com/bigquery/v2/projects/my-project/datasets/my_dataset",
		},
		{
			name:     "universeQualifiedProjectID",
			fullID:   "my-universe:my-project:my_dataset",
			expected: "https://bigquery.googleapis.com/bigquery/v2/projects/my-universe:my-project/datasets/my_dataset",
		},
		{
			name:     "domainScopedProjectID",
			fullID:   "example.com:my-project:my_dataset",
			expected: "https://bigquery.googleapis.com/bigquery/v2/projects/example.com:my-project/datasets/my_dataset",
		},
		{
			name:     "malformedFullID_noSelfLink",
			fullID:   "my_dataset",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			status := BigQueryDatasetStatus_FromProto(mapCtx, &pb.DatasetMetadata{FullID: tc.fullID})
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("BigQueryDatasetStatus_FromProto returned error: %v", err)
			}

			got := ""
			if status.SelfLink != nil {
				got = *status.SelfLink
			}
			if got != tc.expected {
				t.Errorf("SelfLink = %q, want %q", got, tc.expected)
			}
		})
	}
}

// TestBigQueryDatasetSpec_FromProto_resourceID checks the resourceID derived
// from FullID. Before the fix, a project ID containing a colon left resourceID
// unset, and the create path failed outright.
func TestBigQueryDatasetSpec_FromProto_resourceID(t *testing.T) {
	tests := []struct {
		name     string
		fullID   string
		expected string
	}{
		{
			name:     "plainProjectID",
			fullID:   "my-project:my_dataset",
			expected: "my_dataset",
		},
		{
			name:     "universeQualifiedProjectID",
			fullID:   "my-universe:my-project:my_dataset",
			expected: "my_dataset",
		},
		{
			name:     "domainScopedProjectID",
			fullID:   "example.com:my-project:my_dataset",
			expected: "my_dataset",
		},
		{
			name:     "malformedFullID_noResourceID",
			fullID:   "my_dataset",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			spec := BigQueryDatasetSpec_FromProto(mapCtx, &pb.DatasetMetadata{FullID: tc.fullID})
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("BigQueryDatasetSpec_FromProto returned error: %v", err)
			}

			got := ""
			if spec.ResourceID != nil {
				got = *spec.ResourceID
			}
			if got != tc.expected {
				t.Errorf("ResourceID = %q, want %q", got, tc.expected)
			}
		})
	}
}
