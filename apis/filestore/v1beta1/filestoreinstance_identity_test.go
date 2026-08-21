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

package v1beta1

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
)

func TestFilestoreInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *FilestoreInstanceIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/instances/my-instance",
			want: &FilestoreInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://file.googleapis.com/projects/my-project/locations/us-central1/instances/my-instance",
			want: &FilestoreInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &FilestoreInstanceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if gotString := i.String(); gotString != "projects/my-project/locations/us-central1/instances/my-instance" {
					t.Errorf("String() = %v, want %v", gotString, "projects/my-project/locations/us-central1/instances/my-instance")
				}
				if gotParent := i.ParentString(); gotParent != "projects/my-project/locations/us-central1" {
					t.Errorf("ParentString() = %v, want %v", gotParent, "projects/my-project/locations/us-central1")
				}
			}
		})
	}
}

func TestFilestoreInstance_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		obj         *FilestoreInstance
		expected    string
		expectError bool
	}{
		{
			name: "resolved identity without externalRef",
			obj: &FilestoreInstance{
				Spec: FilestoreInstanceSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Location:   "us-central1",
					ResourceID: func() *string { s := "my-instance"; return &s }(),
				},
			},
			expected: "projects/my-project/locations/us-central1/instances/my-instance",
		},
		{
			name: "resolved identity matching externalRef",
			obj: &FilestoreInstance{
				Spec: FilestoreInstanceSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Location:   "us-central1",
					ResourceID: func() *string { s := "my-instance"; return &s }(),
				},
				Status: FilestoreInstanceStatus{
					ExternalRef: func() *string { s := "projects/my-project/locations/us-central1/instances/my-instance"; return &s }(),
				},
			},
			expected: "projects/my-project/locations/us-central1/instances/my-instance",
		},
		{
			name: "resolved identity mismatching externalRef",
			obj: &FilestoreInstance{
				Spec: FilestoreInstanceSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Location:   "us-central1",
					ResourceID: func() *string { s := "my-instance"; return &s }(),
				},
				Status: FilestoreInstanceStatus{
					ExternalRef: func() *string { s := "projects/my-project/locations/us-central1/instances/other-instance"; return &s }(),
				},
			},
			expectError: true,
		},
		{
			name: "invalid externalRef format",
			obj: &FilestoreInstance{
				Spec: FilestoreInstanceSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Location:   "us-central1",
					ResourceID: func() *string { s := "my-instance"; return &s }(),
				},
				Status: FilestoreInstanceStatus{
					ExternalRef: func() *string { s := "invalid-ref"; return &s }(),
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, nil)
			if (err != nil) != tt.expectError {
				t.Errorf("GetIdentity() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError {
				if got.String() != tt.expected {
					t.Errorf("GetIdentity() = %q, expected %q", got.String(), tt.expected)
				}
			}
		})
	}
}
