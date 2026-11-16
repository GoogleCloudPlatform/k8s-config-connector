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

package compute

import (
	"context"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestInstantSnapshotAdapter_Export(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		id           *krm.ComputeInstantSnapshotIdentity
		actual       *computepb.InstantSnapshot
		wantProject  string
		wantLocation string
		wantName     string
		wantDesc     string
	}{
		{
			name: "zonal instant snapshot",
			id: &krm.ComputeInstantSnapshotIdentity{
				Project:         "my-project",
				Zone:            "us-central1-a",
				InstantSnapshot: "my-snapshot",
			},
			actual: &computepb.InstantSnapshot{
				Name:        proto.String("my-snapshot"),
				Description: proto.String("zonal snapshot description"),
			},
			wantProject:  "my-project",
			wantLocation: "us-central1-a",
			wantName:     "my-snapshot",
			wantDesc:     "zonal snapshot description",
		},
		{
			name: "regional instant snapshot",
			id: &krm.ComputeInstantSnapshotIdentity{
				Project:         "my-project",
				Region:          "us-central1",
				InstantSnapshot: "my-regional-snapshot",
			},
			actual: &computepb.InstantSnapshot{
				Name:        proto.String("my-regional-snapshot"),
				Description: proto.String("regional snapshot description"),
			},
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantName:     "my-regional-snapshot",
			wantDesc:     "regional snapshot description",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &InstantSnapshotAdapter{
				id:     tc.id,
				actual: tc.actual,
			}

			u, err := adapter.Export(ctx)
			if err != nil {
				t.Fatalf("Export() returned unexpected error: %v", err)
			}

			if got := u.GetName(); got != tc.wantName {
				t.Errorf("u.GetName() = %q, want %q", got, tc.wantName)
			}

			if got := u.GetKind(); got != "ComputeInstantSnapshot" {
				t.Errorf("u.GetKind() = %q, want ComputeInstantSnapshot", got)
			}

			projectRefExternal, found, err := unstructured.NestedString(u.Object, "spec", "projectRef", "external")
			if err != nil || !found {
				t.Fatalf("spec.projectRef.external not found: %v", err)
			}
			if projectRefExternal != tc.wantProject {
				t.Errorf("spec.projectRef.external = %q, want %q", projectRefExternal, tc.wantProject)
			}

			location, found, err := unstructured.NestedString(u.Object, "spec", "location")
			if err != nil || !found {
				t.Fatalf("spec.location not found: %v", err)
			}
			if location != tc.wantLocation {
				t.Errorf("spec.location = %q, want %q", location, tc.wantLocation)
			}

			desc, found, err := unstructured.NestedString(u.Object, "spec", "description")
			if err != nil || !found {
				t.Fatalf("spec.description not found: %v", err)
			}
			if desc != tc.wantDesc {
				t.Errorf("spec.description = %q, want %q", desc, tc.wantDesc)
			}
		})
	}
}
