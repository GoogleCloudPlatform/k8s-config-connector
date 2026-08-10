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
)

func TestVertexAITensorboardExperimentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAITensorboardExperimentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/tensorboards/my-tensorboard/experiments/my-experiment",
			want: &VertexAITensorboardExperimentIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
				Experiment:  "my-experiment",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/tensorboards/my-tensorboard/experiments/my-experiment",
			want: &VertexAITensorboardExperimentIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
				Experiment:  "my-experiment",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAITensorboardExperimentIdentity{}
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

func TestGetIdentityFromVertexAITensorboardExperimentSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *VertexAITensorboardExperiment
		want    *VertexAITensorboardExperimentIdentity
		wantErr bool
	}{
		{
			name: "valid external references",
			obj: &VertexAITensorboardExperiment{
				Spec: VertexAITensorboardExperimentSpec{
					Location: common.LazyPtr("us-central1"),
					TensorboardRef: &VertexAITensorboardRef{
						External: "projects/my-project/locations/us-central1/tensorboards/my-tensorboard",
					},
					ResourceID: common.LazyPtr("my-experiment"),
				},
			},
			want: &VertexAITensorboardExperimentIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				Tensorboard: "my-tensorboard",
				Experiment:  "my-experiment",
			},
		},
		{
			name: "missing location",
			obj: &VertexAITensorboardExperiment{
				Spec: VertexAITensorboardExperimentSpec{
					TensorboardRef: &VertexAITensorboardRef{
						External: "projects/my-project/locations/us-central1/tensorboards/my-tensorboard",
					},
					ResourceID: common.LazyPtr("my-experiment"),
				},
			},
			wantErr: true,
		},
		{
			name: "mismatched location",
			obj: &VertexAITensorboardExperiment{
				Spec: VertexAITensorboardExperimentSpec{
					Location: common.LazyPtr("us-east1"),
					TensorboardRef: &VertexAITensorboardRef{
						External: "projects/my-project/locations/us-central1/tensorboards/my-tensorboard",
					},
					ResourceID: common.LazyPtr("my-experiment"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromVertexAITensorboardExperimentSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromVertexAITensorboardExperimentSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromVertexAITensorboardExperimentSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
