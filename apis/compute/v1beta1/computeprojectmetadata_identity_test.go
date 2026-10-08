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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestComputeProjectMetadataIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeProjectMetadataIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project",
			want: &ComputeProjectMetadataIdentity{
				Project: "my-project",
			},
		},
		{
			name: "short reference (just project ID)",
			ref:  "my-project",
			want: &ComputeProjectMetadataIdentity{
				Project: "my-project",
			},
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/compute/v1/projects/my-project",
			want: &ComputeProjectMetadataIdentity{
				Project: "my-project",
			},
		},
		{
			name: "www googleapis url",
			ref:  "https://www.googleapis.com/compute/v1/projects/my-project",
			want: &ComputeProjectMetadataIdentity{
				Project: "my-project",
			},
		},
		{
			name: "cai url",
			ref:  "//compute.googleapis.com/projects/my-project",
			want: &ComputeProjectMetadataIdentity{
				Project: "my-project",
			},
		},
		{
			name:    "invalid format with extra segments",
			ref:     "projects/my-project/zones/us-central1-a",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeProjectMetadataIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeProjectMetadataIdentity_String(t *testing.T) {
	id := &ComputeProjectMetadataIdentity{
		Project: "my-project",
	}
	expected := "projects/my-project"
	if actual := id.String(); actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestComputeProjectMetadataIdentity_ParentString(t *testing.T) {
	id := &ComputeProjectMetadataIdentity{
		Project: "my-project",
	}
	expected := "projects/my-project"
	if actual := id.ParentString(); actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestComputeProjectMetadataIdentity_Host(t *testing.T) {
	id := &ComputeProjectMetadataIdentity{
		Project: "my-project",
	}
	expected := "compute.googleapis.com"
	if actual := id.Host(); actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestComputeProjectMetadataRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref:  "projects/my-project",
		},
		{
			name: "valid full compute url",
			ref:  "https://compute.googleapis.com/compute/v1/projects/my-project",
		},
		{
			name:    "invalid external reference",
			ref:     "invalid/format/too/many/parts",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeProjectMetadataRef{
				External: tt.ref,
			}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeProjectMetadataRef_ParseExternalToIdentity(t *testing.T) {
	r := &ComputeProjectMetadataRef{
		External: "projects/my-project",
	}
	id, err := r.ParseExternalToIdentity()
	if err != nil {
		t.Fatalf("ParseExternalToIdentity() error = %v", err)
	}
	expected := &ComputeProjectMetadataIdentity{
		Project: "my-project",
	}
	if diff := cmp.Diff(expected, id); diff != "" {
		t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
	}
}
