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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
)

func TestVertexAITensorboardIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAITensorboardIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/tensorboards/my-tensorboard",
			want: &VertexAITensorboardIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/tensorboards/my-tensorboard",
			want: &VertexAITensorboardIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAITensorboardIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestGetIdentityFromVertexAITensorboardSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *VertexAITensorboard
		want    *VertexAITensorboardIdentity
		wantErr bool
	}{
		{
			name: "resourceID is simple name",
			obj: &VertexAITensorboard{
				Spec: VertexAITensorboardSpec{
					Region:     "us-central1",
					ResourceID: common.LazyPtr("my-tensorboard"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &VertexAITensorboardIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
			},
		},
		{
			name: "resourceID is full GCP URI",
			obj: &VertexAITensorboard{
				Spec: VertexAITensorboardSpec{
					Region:     "us-central1",
					ResourceID: common.LazyPtr("projects/my-project/locations/us-central1/tensorboards/my-tensorboard"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &VertexAITensorboardIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromVertexAITensorboardSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromVertexAITensorboardSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromVertexAITensorboardSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
