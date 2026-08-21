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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNetworkServicesMeshIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *NetworkServicesMeshIdentity
		wantErr bool
	}{
		{
			name: "valid global mesh reference",
			ref:  "projects/my-project/locations/global/meshes/my-mesh",
			want: &NetworkServicesMeshIdentity{
				ProjectID: "my-project",
				Location:  "global",
				MeshID:    "my-mesh",
			},
			wantErr: false,
		},
		{
			name: "valid regional mesh reference",
			ref:  "projects/my-project/locations/us-central1/meshes/my-mesh",
			want: &NetworkServicesMeshIdentity{
				ProjectID: "my-project",
				Location:  "us-central1",
				MeshID:    "my-mesh",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing meshes)",
			ref:     "projects/my-project/locations/global/my-mesh",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/global/meshes/my-mesh",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkServicesMeshIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesMeshIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if *i != *tt.want {
					t.Errorf("NetworkServicesMeshIdentity.FromExternal() = %v, want %v", i, tt.want)
				}
			}
		})
	}
}

func TestNetworkServicesMeshRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()
	// No need to register actual types if we use Unstructured

	tests := []struct {
		name             string
		ref              *NetworkServicesMeshRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesMeshRef{
				External: "projects/my-project/locations/global/meshes/my-mesh",
			},
			wantExternal: "projects/my-project/locations/global/meshes/my-mesh",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesMeshRef{
				Name:     "my-mesh",
				External: "projects/my-project/locations/global/meshes/my-mesh",
			},
			wantErr: true,
		},
		{
			name: "resolve from status.externalRef",
			ref: &NetworkServicesMeshRef{
				Name: "my-mesh",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesMesh",
						"metadata": map[string]interface{}{
							"name":      "my-mesh",
							"namespace": "test-ns",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/my-project/locations/global/meshes/my-mesh",
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/global/meshes/my-mesh",
			wantErr:      false,
		},
		{
			name: "resolve from fallback (metadata.name)",
			ref: &NetworkServicesMeshRef{
				Name: "my-mesh",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesMesh",
						"metadata": map[string]interface{}{
							"name":      "my-mesh",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"location": "global",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/global/meshes/my-mesh",
			wantErr:      false,
		},
		{
			name: "resolve from fallback (resourceID)",
			ref: &NetworkServicesMeshRef{
				Name: "k8s-name",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesMesh",
						"metadata": map[string]interface{}{
							"name":      "k8s-name",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"resourceID": "my-mesh",
							"location":   "global",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/global/meshes/my-mesh",
			wantErr:      false,
		},
		{
			name: "missing reference",
			ref: &NetworkServicesMeshRef{
				Name: "non-existent",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build()
			err := tt.ref.Normalize(context.Background(), client, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesMeshRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("NetworkServicesMeshRef.Normalize() got = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
