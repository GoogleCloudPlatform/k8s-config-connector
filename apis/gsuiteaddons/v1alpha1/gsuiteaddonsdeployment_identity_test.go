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
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGSuiteAddonsDeploymentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name           string
		ref            string
		wantProject    string
		wantDeployment string
		wantErr        bool
	}{
		{
			name:           "valid external ref",
			ref:            "projects/my-project/deployments/my-deployment",
			wantProject:    "my-project",
			wantDeployment: "my-deployment",
			wantErr:        false,
		},
		{
			name:           "full url",
			ref:            "https://gsuiteaddons.googleapis.com/projects/my-project/deployments/my-deployment",
			wantProject:    "my-project",
			wantDeployment: "my-deployment",
			wantErr:        false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/other/my-deployment",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &GSuiteAddonsDeploymentIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("GSuiteAddonsDeploymentIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("GSuiteAddonsDeploymentIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Deployment != tt.wantDeployment {
					t.Errorf("GSuiteAddonsDeploymentIdentity.FromExternal() Deployment = %v, want %v", i.Deployment, tt.wantDeployment)
				}
				if got := i.String(); got != "projects/"+tt.wantProject+"/deployments/"+tt.wantDeployment {
					t.Errorf("GSuiteAddonsDeploymentIdentity.String() = %v", got)
				}
			}
		})
	}
}

func TestGSuiteAddonsDeployment_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *GSuiteAddonsDeployment
		wantErr bool
		want    *GSuiteAddonsDeploymentIdentity
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &GSuiteAddonsDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-name",
				},
				Spec: GSuiteAddonsDeploymentSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					ResourceID: common.LazyPtr("my-deployment"),
				},
			},
			want: &GSuiteAddonsDeploymentIdentity{
				Project:    "my-project",
				Deployment: "my-deployment",
			},
		},
		{
			name: "GetIdentity fallback to metadata.name",
			obj: &GSuiteAddonsDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-name",
				},
				Spec: GSuiteAddonsDeploymentSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &GSuiteAddonsDeploymentIdentity{
				Project:    "my-project",
				Deployment: "k8s-name",
			},
		},
		{
			name: "GetIdentity mismatch with status externalRef",
			obj: &GSuiteAddonsDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-name",
				},
				Spec: GSuiteAddonsDeploymentSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					ResourceID: common.LazyPtr("my-deployment"),
				},
				Status: GSuiteAddonsDeploymentStatus{
					ExternalRef: common.LazyPtr("projects/my-project/deployments/other-deployment"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
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
