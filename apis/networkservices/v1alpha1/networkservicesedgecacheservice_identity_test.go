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

func TestNetworkServicesEdgeCacheServiceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *NetworkServicesEdgeCacheServiceIdentity
		wantErr bool
	}{
		{
			name: "valid edgecacheservice reference",
			ref:  "projects/my-project/locations/global/edgeCacheServices/my-edgecacheservice",
			want: &NetworkServicesEdgeCacheServiceIdentity{
				Project:          "my-project",
				EdgeCacheService: "my-edgecacheservice",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing edgeCacheServices)",
			ref:     "projects/my-project/locations/global/my-edgecacheservice",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/global/edgeCacheServices/my-edgecacheservice",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkServicesEdgeCacheServiceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesEdgeCacheServiceIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("NetworkServicesEdgeCacheServiceIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNetworkServicesEdgeCacheServiceRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()

	tests := []struct {
		name             string
		ref              *NetworkServicesEdgeCacheServiceRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesEdgeCacheServiceRef{
				External: "projects/my-project/locations/global/edgeCacheServices/my-edgecacheservice",
			},
			wantExternal: "projects/my-project/locations/global/edgeCacheServices/my-edgecacheservice",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesEdgeCacheServiceRef{
				Name:     "my-edgecacheservice",
				External: "projects/my-project/locations/global/edgeCacheServices/my-edgecacheservice",
			},
			wantErr: true,
		},
		{
			name: "resolve from fallback (metadata.name)",
			ref: &NetworkServicesEdgeCacheServiceRef{
				Name: "my-edgecacheservice",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1alpha1",
						"kind":       "NetworkServicesEdgeCacheService",
						"metadata": map[string]interface{}{
							"name":      "my-edgecacheservice",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/global/edgeCacheServices/my-edgecacheservice",
			wantErr:      false,
		},
		{
			name: "resolve from fallback (resourceID)",
			ref: &NetworkServicesEdgeCacheServiceRef{
				Name: "k8s-name",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1alpha1",
						"kind":       "NetworkServicesEdgeCacheService",
						"metadata": map[string]interface{}{
							"name":      "k8s-name",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"resourceID": "my-edgecacheservice",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/global/edgeCacheServices/my-edgecacheservice",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			cl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build()
			err := tt.ref.Normalize(ctx, cl, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesEdgeCacheServiceRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.ref.External != tt.wantExternal {
					t.Errorf("NetworkServicesEdgeCacheServiceRef.Normalize() gotExternal = %v, want %v", tt.ref.External, tt.wantExternal)
				}
			}
		})
	}
}
