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
			name: "regional URL with API version",
			ref:  "https://us-central1-aiplatform.googleapis.com/v1beta1/projects/my-project/locations/us-central1/extensions/my-extension",
			want: &VertexAIExtensionIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Extension: "my-extension",
			},
		},
		{
			name: "regional URL without API version",
			ref:  "https://us-central1-aiplatform.googleapis.com/projects/my-project/locations/us-central1/extensions/my-extension",
			want: &VertexAIExtensionIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Extension: "my-extension",
			},
		},
		{
			name: "regional URL with v1",
			ref:  "https://us-east4-aiplatform.googleapis.com/v1/projects/my-project/locations/us-central1/extensions/my-extension",
			want: &VertexAIExtensionIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Extension: "my-extension",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAIExtensionIdentity{}
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

func TestGetServiceGeneratedResourceID(t *testing.T) {
	tests := []struct {
		name string
		spec *VertexAIExtensionSpec
		want string
	}{
		{
			name: "resourceID set",
			spec: &VertexAIExtensionSpec{
				ResourceID: ptrTo("my-id"),
			},
			want: "my-id",
		},
		{
			name: "resourceID set with prefix",
			spec: &VertexAIExtensionSpec{
				ResourceID: ptrTo("extensions/my-id"),
			},
			want: "my-id",
		},
		{
			name: "resourceID unset",
			spec: &VertexAIExtensionSpec{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &VertexAIExtension{
				Spec: *tt.spec,
			}
			got := GetServiceGeneratedResourceID(obj)
			if got != tt.want {
				t.Errorf("GetServiceGeneratedResourceID() = %q, want %q", got, tt.want)
			}
		})
	}
}

func ptrTo[T any](v T) *T {
	return &v
}
