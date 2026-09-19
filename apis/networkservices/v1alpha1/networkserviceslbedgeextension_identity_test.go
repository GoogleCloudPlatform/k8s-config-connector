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

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNetworkServicesLBEdgeExtensionIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *LBEdgeExtensionIdentity
		wantErr bool
	}{
		{
			name: "valid lbEdgeExtension reference",
			ref:  "projects/my-project/locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			want: &LBEdgeExtensionIdentity{
				Project:         "my-project",
				Location:        "us-central1",
				LbEdgeExtension: "my-lbedgeextension",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing lbEdgeExtensions)",
			ref:     "projects/my-project/locations/us-central1/my-lbedgeextension",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &LBEdgeExtensionIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("LBEdgeExtensionIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("LBEdgeExtensionIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNetworkServicesLBEdgeExtensionRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()

	tests := []struct {
		name             string
		ref              *NetworkServicesLBEdgeExtensionRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesLBEdgeExtensionRef{
				External: "projects/my-project/locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			},
			wantExternal: "projects/my-project/locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesLBEdgeExtensionRef{
				Name:     "my-lbedgeextension",
				External: "projects/my-project/locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			},
			wantErr: true,
		},
		{
			name: "resolve from fallback (metadata.name)",
			ref: &NetworkServicesLBEdgeExtensionRef{
				Name: "my-lbedgeextension",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1alpha1",
						"kind":       "NetworkServicesLBEdgeExtension",
						"metadata": map[string]interface{}{
							"name":      "my-lbedgeextension",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"location": "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			wantErr:      false,
		},
		{
			name: "resolve from fallback (resourceID)",
			ref: &NetworkServicesLBEdgeExtensionRef{
				Name: "k8s-name",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1alpha1",
						"kind":       "NetworkServicesLBEdgeExtension",
						"metadata": map[string]interface{}{
							"name":      "k8s-name",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"resourceID": "my-lbedgeextension",
							"location":   "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/lbEdgeExtensions/my-lbedgeextension",
			wantErr:      false,
		},
		{
			name: "missing reference",
			ref: &NetworkServicesLBEdgeExtensionRef{
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
				t.Errorf("NetworkServicesLBEdgeExtensionRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("NetworkServicesLBEdgeExtensionRef.Normalize() got = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
