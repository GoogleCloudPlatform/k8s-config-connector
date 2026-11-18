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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCloudDeployRolloutIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *CloudDeployRolloutIdentity
		wantParent string
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release/rollouts/my-rollout",
			want: &CloudDeployRolloutIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
				Rollout:          "my-rollout",
			},
			wantParent: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
		},
		{
			name: "valid reference with host prefix",
			ref:  "//clouddeploy.googleapis.com/projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release/rollouts/my-rollout",
			want: &CloudDeployRolloutIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
				Rollout:          "my-rollout",
			},
			wantParent: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "invalid reference segment (wrong collection)",
			ref:     "projects/my-project/locations/us-central1/pipelines/my-pipeline/releases/my-release/rollouts/my-rollout",
			wantErr: true,
		},
		{
			name:    "invalid reference segment (wrong rollouts collection)",
			ref:     "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release/deployRollouts/my-rollout",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudDeployRolloutIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("(-want +got):\n%s", diff)
				}
				if gotParent := i.ParentString(); gotParent != tt.wantParent {
					t.Errorf("ParentString() = %v, want %v", gotParent, tt.wantParent)
				}
				if i.Host() != "clouddeploy.googleapis.com" {
					t.Errorf("Host() = %v, want %v", i.Host(), "clouddeploy.googleapis.com")
				}
			}
		})
	}
}

func TestCloudDeployRollout_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *CloudDeployRollout
		want    *CloudDeployRolloutIdentity
		wantErr bool
	}{
		{
			name: "spec with parentRef and resourceID",
			obj: &CloudDeployRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-rollout",
					Namespace: "test-ns",
				},
				Spec: CloudDeployRolloutSpec{
					ParentRef: &CloudDeployReleaseRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
					},
					ResourceID: direct.LazyPtr("custom-rollout-id"),
				},
			},
			want: &CloudDeployRolloutIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
				Rollout:          "custom-rollout-id",
			},
		},
		{
			name: "spec with parentRef and fallback to metadata.name",
			obj: &CloudDeployRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-rollout-name",
					Namespace: "test-ns",
				},
				Spec: CloudDeployRolloutSpec{
					ParentRef: &CloudDeployReleaseRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
					},
				},
			},
			want: &CloudDeployRolloutIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
				Rollout:          "k8s-rollout-name",
			},
		},
		{
			name: "status.externalRef matches spec",
			obj: &CloudDeployRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rollout",
					Namespace: "test-ns",
				},
				Spec: CloudDeployRolloutSpec{
					ParentRef: &CloudDeployReleaseRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
					},
				},
				Status: CloudDeployRolloutStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release/rollouts/my-rollout"),
				},
			},
			want: &CloudDeployRolloutIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
				Rollout:          "my-rollout",
			},
		},
		{
			name: "status.externalRef drift error",
			obj: &CloudDeployRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rollout",
					Namespace: "test-ns",
				},
				Spec: CloudDeployRolloutSpec{
					ParentRef: &CloudDeployReleaseRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
					},
				},
				Status: CloudDeployRolloutStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/other-release/rollouts/my-rollout"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing parentRef",
			obj: &CloudDeployRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rollout",
					Namespace: "test-ns",
				},
				Spec: CloudDeployRolloutSpec{},
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
				gotIdentity, ok := got.(*CloudDeployRolloutIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned unexpected type %T", got)
				}
				if diff := cmp.Diff(tt.want, gotIdentity); diff != "" {
					t.Errorf("(-want +got):\n%s", diff)
				}
			}
		})
	}
}
