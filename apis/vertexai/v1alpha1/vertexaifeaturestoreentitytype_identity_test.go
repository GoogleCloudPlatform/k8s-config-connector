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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestVertexAIFeaturestoreEntityTypeIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAIFeaturestoreEntityTypeIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
			want: &VertexAIFeaturestoreEntityTypeIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
			},
		},
		{
			name: "full url with scheme",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
			want: &VertexAIFeaturestoreEntityTypeIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
			},
		},
		{
			name: "full url without scheme",
			ref:  "//aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
			want: &VertexAIFeaturestoreEntityTypeIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty reference",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAIFeaturestoreEntityTypeIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et")
				}
			}
		})
	}
}

func TestVertexAIFeaturestoreEntityTypeIdentity_ParentString(t *testing.T) {
	i := &VertexAIFeaturestoreEntityTypeIdentity{
		Project:      "my-project",
		Location:     "us-central1",
		Featurestore: "my-fs",
		EntityType:   "my-et",
	}
	want := "projects/my-project/locations/us-central1/featurestores/my-fs"
	if got := i.ParentString(); got != want {
		t.Errorf("ParentString() = %q, want %q", got, want)
	}
}

func TestGetIdentityFromVertexAIFeaturestoreEntityTypeSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *VertexAIFeaturestoreEntityType
		want    *VertexAIFeaturestoreEntityTypeIdentity
		wantErr bool
	}{
		{
			name: "valid spec with resourceID",
			obj: &VertexAIFeaturestoreEntityType{
				Spec: VertexAIFeaturestoreEntityTypeSpec{
					Featurestore: "projects/my-project/locations/us-central1/featurestores/my-fs",
					ResourceID:   common.LazyPtr("my-et"),
				},
			},
			want: &VertexAIFeaturestoreEntityTypeIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
			},
		},
		{
			name: "valid spec defaulting to metadata.name",
			obj: &VertexAIFeaturestoreEntityType{
				ObjectMeta: metav1.ObjectMeta{
					Name: "default-et",
				},
				Spec: VertexAIFeaturestoreEntityTypeSpec{
					Featurestore: "projects/my-project/locations/us-central1/featurestores/my-fs",
				},
			},
			want: &VertexAIFeaturestoreEntityTypeIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "default-et",
			},
		},
		{
			name: "valid spec with fully qualified featurestore parent URL",
			obj: &VertexAIFeaturestoreEntityType{
				ObjectMeta: metav1.ObjectMeta{
					Name: "default-et",
				},
				Spec: VertexAIFeaturestoreEntityTypeSpec{
					Featurestore: "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs",
				},
			},
			want: &VertexAIFeaturestoreEntityTypeIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "default-et",
			},
		},
		{
			name: "missing featurestore",
			obj: &VertexAIFeaturestoreEntityType{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-et",
				},
				Spec: VertexAIFeaturestoreEntityTypeSpec{
					Featurestore: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid featurestore format",
			obj: &VertexAIFeaturestoreEntityType{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-et",
				},
				Spec: VertexAIFeaturestoreEntityTypeSpec{
					Featurestore: "invalid/featurestore/format",
				},
			},
			wantErr: true,
		},
		{
			name: "missing resourceID and name",
			obj: &VertexAIFeaturestoreEntityType{
				Spec: VertexAIFeaturestoreEntityTypeSpec{
					Featurestore: "projects/my-project/locations/us-central1/featurestores/my-fs",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromVertexAIFeaturestoreEntityTypeSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromVertexAIFeaturestoreEntityTypeSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromVertexAIFeaturestoreEntityTypeSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestVertexAIFeaturestoreEntityTypeRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     *VertexAIFeaturestoreEntityTypeRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &VertexAIFeaturestoreEntityTypeRef{
				External: "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
			},
			wantErr: false,
		},
		{
			name: "invalid external reference",
			ref: &VertexAIFeaturestoreEntityTypeRef{
				External: "invalid/external/ref",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.ValidateExternal(tt.ref.External)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
