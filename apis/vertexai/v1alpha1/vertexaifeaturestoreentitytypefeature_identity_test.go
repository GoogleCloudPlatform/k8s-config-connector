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

func TestVertexAIFeaturestoreEntityTypeFeatureIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAIFeaturestoreEntityTypeFeatureIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feature",
			want: &VertexAIFeaturestoreEntityTypeFeatureIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
				Feature:      "my-feature",
			},
		},
		{
			name: "full url with scheme",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feature",
			want: &VertexAIFeaturestoreEntityTypeFeatureIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
				Feature:      "my-feature",
			},
		},
		{
			name: "full url without scheme",
			ref:  "//aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feature",
			want: &VertexAIFeaturestoreEntityTypeFeatureIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
				Feature:      "my-feature",
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
			i := &VertexAIFeaturestoreEntityTypeFeatureIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feature" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feature")
				}
			}
		})
	}
}

func TestVertexAIFeaturestoreEntityTypeFeatureIdentity_ParentString(t *testing.T) {
	i := &VertexAIFeaturestoreEntityTypeFeatureIdentity{
		Project:      "my-project",
		Location:     "us-central1",
		Featurestore: "my-fs",
		EntityType:   "my-et",
		Feature:      "my-feature",
	}
	want := "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et"
	if got := i.ParentString(); got != want {
		t.Errorf("ParentString() = %q, want %q", got, want)
	}
}

func TestGetIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *VertexAIFeaturestoreEntityTypeFeature
		want    *VertexAIFeaturestoreEntityTypeFeatureIdentity
		wantErr bool
	}{
		{
			name: "valid spec with resourceID",
			obj: &VertexAIFeaturestoreEntityTypeFeature{
				Spec: VertexAIFeaturestoreEntityTypeFeatureSpec{
					Entitytype: "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
					ResourceID: common.LazyPtr("my-feature"),
				},
			},
			want: &VertexAIFeaturestoreEntityTypeFeatureIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
				Feature:      "my-feature",
			},
		},
		{
			name: "valid spec defaulting to metadata.name",
			obj: &VertexAIFeaturestoreEntityTypeFeature{
				ObjectMeta: metav1.ObjectMeta{
					Name: "default-feature",
				},
				Spec: VertexAIFeaturestoreEntityTypeFeatureSpec{
					Entitytype: "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
				},
			},
			want: &VertexAIFeaturestoreEntityTypeFeatureIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
				Feature:      "default-feature",
			},
		},
		{
			name: "valid spec with fully qualified entitytype parent URL",
			obj: &VertexAIFeaturestoreEntityTypeFeature{
				ObjectMeta: metav1.ObjectMeta{
					Name: "default-feature",
				},
				Spec: VertexAIFeaturestoreEntityTypeFeatureSpec{
					Entitytype: "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
				},
			},
			want: &VertexAIFeaturestoreEntityTypeFeatureIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				Featurestore: "my-fs",
				EntityType:   "my-et",
				Feature:      "default-feature",
			},
		},
		{
			name: "missing entitytype",
			obj: &VertexAIFeaturestoreEntityTypeFeature{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-feature",
				},
				Spec: VertexAIFeaturestoreEntityTypeFeatureSpec{
					Entitytype: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid entitytype format",
			obj: &VertexAIFeaturestoreEntityTypeFeature{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-feature",
				},
				Spec: VertexAIFeaturestoreEntityTypeFeatureSpec{
					Entitytype: "invalid/entitytype/format",
				},
			},
			wantErr: true,
		},
		{
			name: "missing resourceID and name",
			obj: &VertexAIFeaturestoreEntityTypeFeature{
				Spec: VertexAIFeaturestoreEntityTypeFeatureSpec{
					Entitytype: "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestVertexAIFeaturestoreEntityTypeFeatureRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     *VertexAIFeaturestoreEntityTypeFeatureRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &VertexAIFeaturestoreEntityTypeFeatureRef{
				External: "projects/my-project/locations/us-central1/featurestores/my-fs/entityTypes/my-et/features/my-feature",
			},
			wantErr: false,
		},
		{
			name: "invalid external reference",
			ref: &VertexAIFeaturestoreEntityTypeFeatureRef{
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
