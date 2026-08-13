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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
)

func TestVertexAIDatasetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAIDatasetIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &VertexAIDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &VertexAIDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &VertexAIDatasetIdentity{}
			err := got.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FromExternal mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestVertexAIDataset_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *VertexAIDataset
		want    *VertexAIDatasetIdentity
		wantErr bool
	}{
		{
			name: "GetIdentity with only spec",
			obj: &VertexAIDataset{
				Spec: VertexAIDatasetSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Region:     common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-dataset"),
				},
			},
			want: &VertexAIDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name: "GetIdentity overrides spec with numeric project and dataset ID from status",
			obj: &VertexAIDataset{
				Spec: VertexAIDatasetSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Region:     common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-dataset"),
				},
				Status: VertexAIDatasetStatus{
					ObservedState: &VertexAIDatasetObservedState{
						Name: common.LazyPtr("projects/1234567890/locations/us-central1/datasets/987654321"),
					},
				},
			},
			want: &VertexAIDatasetIdentity{
				Project:  "1234567890",
				Location: "us-central1",
				Dataset:  "987654321",
			},
		},
		{
			name: "GetIdentity allows non-numeric status dataset matching",
			obj: &VertexAIDataset{
				Spec: VertexAIDatasetSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Region:     common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-dataset"),
				},
				Status: VertexAIDatasetStatus{
					ObservedState: &VertexAIDatasetObservedState{
						Name: common.LazyPtr("projects/my-project/locations/us-central1/datasets/my-dataset"),
					},
				},
			},
			want: &VertexAIDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name: "GetIdentity fails on mismatched location",
			obj: &VertexAIDataset{
				Spec: VertexAIDatasetSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Region:     common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-dataset"),
				},
				Status: VertexAIDatasetStatus{
					ObservedState: &VertexAIDatasetObservedState{
						Name: common.LazyPtr("projects/my-project/locations/us-east1/datasets/my-dataset"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity fails on mismatched dataset (non-numeric)",
			obj: &VertexAIDataset{
				Spec: VertexAIDatasetSpec{
					ProjectRef: refs.ProjectRef{
						External: "my-project",
					},
					Region:     common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-dataset"),
				},
				Status: VertexAIDatasetStatus{
					ObservedState: &VertexAIDatasetObservedState{
						Name: common.LazyPtr("projects/my-project/locations/us-central1/datasets/other-dataset"),
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				gotIdentity := got.(*VertexAIDatasetIdentity)
				if diff := cmp.Diff(tt.want, gotIdentity); diff != "" {
					t.Errorf("GetIdentity mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
