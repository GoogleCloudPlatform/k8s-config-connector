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
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestContainerClusterIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ContainerClusterIdentity
	}{
		{
			name: "valid regional reference",
			ref:  "projects/my-project/locations/us-central1/clusters/my-cluster",
			want: &ContainerClusterIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
			},
		},
		{
			name: "valid zonal reference",
			ref:  "projects/my-project/zones/us-central1-a/clusters/my-cluster",
			want: &ContainerClusterIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
			},
		},
		{
			name: "valid full regional url",
			ref:  "https://container.googleapis.com/v1/projects/my-project/locations/us-central1/clusters/my-cluster",
			want: &ContainerClusterIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
			},
		},
		{
			name: "valid full zonal url",
			ref:  "https://container.googleapis.com/v1/projects/my-project/zones/us-central1-a/clusters/my-cluster",
			want: &ContainerClusterIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ContainerClusterIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.Cluster != tt.want.Cluster {
					t.Errorf("Cluster = %v, want %v", i.Cluster, tt.want.Cluster)
				}
			}
		})
	}
}

func TestContainerClusterIdentity_String(t *testing.T) {
	identity := &ContainerClusterIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Cluster:  "my-cluster",
	}
	expected := "projects/my-project/locations/us-central1/clusters/my-cluster"
	if actual := identity.String(); actual != expected {
		t.Errorf("String() = %v, want %v", actual, expected)
	}
}

func TestContainerCluster_GetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() unexpected error: %v", err)
	}
	reader := fake.NewClientBuilder().WithScheme(scheme).Build()

	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name    string
		obj     *ContainerCluster
		want    *ContainerClusterIdentity
		wantStr string
		wantErr bool
	}{
		{
			name: "valid_spec_and_matching_v1beta1_self_link",
			obj: &ContainerCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-cluster",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ContainerClusterSpec{
					Location: strPtr("us-central1"),
				},
				Status: ContainerClusterStatus{
					SelfLink: strPtr("https://container.googleapis.com/v1beta1/projects/my-project/locations/us-central1/clusters/my-cluster"),
				},
			},
			want: &ContainerClusterIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
			},
			wantStr: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name: "valid_zonal_spec_and_matching_zonal_self_link",
			obj: &ContainerCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-cluster",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ContainerClusterSpec{
					Location: strPtr("us-central1-a"),
				},
				Status: ContainerClusterStatus{
					SelfLink: strPtr("https://container.googleapis.com/v1/projects/my-project/zones/us-central1-a/clusters/my-cluster"),
				},
			},
			want: &ContainerClusterIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
			},
			wantStr: "projects/my-project/locations/us-central1-a/clusters/my-cluster",
		},
		{
			name: "resource_id_path_traversal_rejected",
			obj: &ContainerCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-cluster",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ContainerClusterSpec{
					Location:   strPtr("us-central1"),
					ResourceID: strPtr("real-cluster/nodePools/np-1"),
				},
			},
			wantErr: true,
		},
		{
			name: "project_id_path_traversal_rejected",
			obj: &ContainerCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-cluster",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project/locations/us",
					},
				},
				Spec: ContainerClusterSpec{
					Location: strPtr("us-central1"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.obj.GetIdentity(t.Context(), reader)

			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if got.String() != tc.wantStr {
					t.Errorf("GetIdentity().String() = %q, want %q", got.String(), tc.wantStr)
				}
			}
		})
	}
}
