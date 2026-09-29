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

package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVertexAIExtensionIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAIExtensionIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/extensions/my-extension",
			want: &VertexAIExtensionIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Extension: "my-extension",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/extensions/my-extension",
			want: &VertexAIExtensionIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Extension: "my-extension",
			},
		},
		{
			name: "cai format",
			ref:  "//aiplatform.googleapis.com/projects/my-project/locations/us-central1/extensions/my-extension",
			want: &VertexAIExtensionIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Extension: "my-extension",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &VertexAIExtensionIdentity{}
			err := got.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got.String() != "projects/my-project/locations/us-central1/extensions/my-extension" {
					t.Errorf("String() = %v, want %v", got.String(), "projects/my-project/locations/us-central1/extensions/my-extension")
				}
				if got.Host() != "aiplatform.googleapis.com" {
					t.Errorf("Host() = %v, want %v", got.Host(), "aiplatform.googleapis.com")
				}
			}
		})
	}
}

func TestVertexAIExtensionRef_Normalize(t *testing.T) {
	tests := []struct {
		name    string
		ref     *VertexAIExtensionRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &VertexAIExtensionRef{
				External: "projects/my-project/locations/us-central1/extensions/my-extension",
			},
			wantErr: false,
		},
		{
			name: "empty external reference",
			ref: &VertexAIExtensionRef{
				External: "",
			},
			wantErr: true,
		},
		{
			name: "invalid external reference format",
			ref: &VertexAIExtensionRef{
				External: "invalid/format",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), nil, "default")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
