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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCloudDeployReleaseIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *CloudDeployReleaseIdentity
		wantParent string
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
			want: &CloudDeployReleaseIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
			},
			wantParent: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
		},
		{
			name: "valid reference with host prefix",
			ref:  "//clouddeploy.googleapis.com/projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release",
			want: &CloudDeployReleaseIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
			},
			wantParent: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "invalid reference segment (wrong collection)",
			ref:     "projects/my-project/locations/us-central1/pipelines/my-pipeline/releases/my-release",
			wantErr: true,
		},
		{
			name:    "invalid reference segment (wrong releases collection)",
			ref:     "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/deployReleases/my-release",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudDeployReleaseIdentity{}
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

func TestCloudDeployRelease_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *CloudDeployRelease
		want    *CloudDeployReleaseIdentity
		wantErr bool
	}{
		{
			name: "spec with parentRef and resourceID",
			obj: &CloudDeployRelease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-release",
					Namespace: "test-ns",
				},
				Spec: CloudDeployReleaseSpec{
					ParentRef: &v1beta1.DeliveryPipelineRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
					},
					ResourceID: direct.LazyPtr("custom-release-id"),
				},
			},
			want: &CloudDeployReleaseIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "custom-release-id",
			},
		},
		{
			name: "spec with parentRef and fallback to metadata.name",
			obj: &CloudDeployRelease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-release-name",
					Namespace: "test-ns",
				},
				Spec: CloudDeployReleaseSpec{
					ParentRef: &v1beta1.DeliveryPipelineRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
					},
				},
			},
			want: &CloudDeployReleaseIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "k8s-release-name",
			},
		},
		{
			name: "status.externalRef matches spec",
			obj: &CloudDeployRelease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-release",
					Namespace: "test-ns",
				},
				Spec: CloudDeployReleaseSpec{
					ParentRef: &v1beta1.DeliveryPipelineRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
					},
				},
				Status: CloudDeployReleaseStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/releases/my-release"),
				},
			},
			want: &CloudDeployReleaseIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DeliveryPipeline: "my-pipeline",
				Release:          "my-release",
			},
		},
		{
			name: "status.externalRef drift error",
			obj: &CloudDeployRelease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-release",
					Namespace: "test-ns",
				},
				Spec: CloudDeployReleaseSpec{
					ParentRef: &v1beta1.DeliveryPipelineRef{
						External: "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
					},
				},
				Status: CloudDeployReleaseStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/deliveryPipelines/other-pipeline/releases/my-release"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing parentRef",
			obj: &CloudDeployRelease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-release",
					Namespace: "test-ns",
				},
				Spec: CloudDeployReleaseSpec{},
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
				gotIdentity, ok := got.(*CloudDeployReleaseIdentity)
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
