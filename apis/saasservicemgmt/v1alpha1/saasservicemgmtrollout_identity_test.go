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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSaaSServiceMgmtRolloutIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *SaaSServiceMgmtRolloutIdentity
	}{
		{
			name: "valid relative resource name",
			ref:  "projects/my-project/locations/us-central1/rollouts/my-rollout",
			want: &SaaSServiceMgmtRolloutIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Rollout:  "my-rollout",
			},
		},
		{
			name: "valid URI with scheme and host",
			ref:  "//saasservicemgmt.googleapis.com/projects/my-project/locations/us-central1/rollouts/my-rollout",
			want: &SaaSServiceMgmtRolloutIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Rollout:  "my-rollout",
			},
		},
		{
			name: "valid HTTP URL",
			ref:  "https://saasservicemgmt.googleapis.com/projects/my-project/locations/us-central1/rollouts/my-rollout",
			want: &SaaSServiceMgmtRolloutIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Rollout:  "my-rollout",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "invalid wrong resource type",
			ref:     "projects/my-project/locations/us-central1/releases/my-release",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &SaaSServiceMgmtRolloutIdentity{}
			err := got.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if gotStr := got.String(); gotStr != "projects/my-project/locations/us-central1/rollouts/my-rollout" {
					t.Errorf("String() = %v, want %v", gotStr, "projects/my-project/locations/us-central1/rollouts/my-rollout")
				}
				if gotParent := got.ParentString(); gotParent != "projects/my-project/locations/us-central1" {
					t.Errorf("ParentString() = %v, want %v", gotParent, "projects/my-project/locations/us-central1")
				}
			}
		})
	}
}

func TestSaaSServiceMgmtRollout_GetIdentity(t *testing.T) {
	tests := []struct {
		name    string
		obj     *SaaSServiceMgmtRollout
		want    *SaaSServiceMgmtRolloutIdentity
		wantErr bool
	}{
		{
			name: "spec identity with ResourceID",
			obj: &SaaSServiceMgmtRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-rollout-krm",
				},
				Spec: SaaSServiceMgmtRolloutSpec{
					Location:   common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-rollout"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &SaaSServiceMgmtRolloutIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Rollout:  "my-rollout",
			},
		},
		{
			name: "spec identity with metadata.name fallback",
			obj: &SaaSServiceMgmtRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-rollout",
				},
				Spec: SaaSServiceMgmtRolloutSpec{
					Location: common.LazyPtr("us-central1"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &SaaSServiceMgmtRolloutIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Rollout:  "my-rollout",
			},
		},
		{
			name: "status externalRef matches spec identity",
			obj: &SaaSServiceMgmtRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-rollout",
				},
				Spec: SaaSServiceMgmtRolloutSpec{
					Location: common.LazyPtr("us-central1"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
				Status: SaaSServiceMgmtRolloutStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/rollouts/my-rollout"),
				},
			},
			want: &SaaSServiceMgmtRolloutIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Rollout:  "my-rollout",
			},
		},
		{
			name: "status externalRef mismatch returns error",
			obj: &SaaSServiceMgmtRollout{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-rollout",
				},
				Spec: SaaSServiceMgmtRolloutSpec{
					Location: common.LazyPtr("us-central1"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
				Status: SaaSServiceMgmtRolloutStatus{
					ExternalRef: common.LazyPtr("projects/other-project/locations/us-central1/rollouts/other-rollout"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(context.Background(), nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
