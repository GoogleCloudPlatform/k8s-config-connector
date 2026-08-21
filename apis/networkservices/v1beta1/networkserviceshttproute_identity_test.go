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

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNetworkServicesHTTPRouteIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *NetworkServicesHTTPRouteIdentity
		wantErr bool
	}{
		{
			name: "valid httpRoute reference",
			ref:  "projects/my-project/locations/us-central1/httpRoutes/my-route",
			want: &NetworkServicesHTTPRouteIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				HttpRoute: "my-route",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing httpRoutes)",
			ref:     "projects/my-project/locations/us-central1/my-route",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/us-central1/httpRoutes/my-route",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkServicesHTTPRouteIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesHTTPRouteIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("NetworkServicesHTTPRouteIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNetworkServicesHTTPRouteRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()

	tests := []struct {
		name             string
		ref              *NetworkServicesHTTPRouteRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesHTTPRouteRef{
				External: "projects/my-project/locations/us-central1/httpRoutes/my-route",
			},
			wantExternal: "projects/my-project/locations/us-central1/httpRoutes/my-route",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesHTTPRouteRef{
				Name:     "my-route",
				External: "projects/my-project/locations/us-central1/httpRoutes/my-route",
			},
			wantErr: true,
		},
		{
			name: "resolve from fallback (metadata.name)",
			ref: &NetworkServicesHTTPRouteRef{
				Name: "my-route",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesHTTPRoute",
						"metadata": map[string]interface{}{
							"name":      "my-route",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"location": "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": "True",
								},
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/httpRoutes/my-route",
			wantErr:      false,
		},
		{
			name: "resolve from fallback fails when not ready",
			ref: &NetworkServicesHTTPRouteRef{
				Name: "my-route",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesHTTPRoute",
						"metadata": map[string]interface{}{
							"name":      "my-route",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"location": "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": "False",
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "resolve from fallback (resourceID)",
			ref: &NetworkServicesHTTPRouteRef{
				Name: "k8s-name",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesHTTPRoute",
						"metadata": map[string]interface{}{
							"name":      "k8s-name",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"resourceID": "my-route",
							"location":   "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": "True",
								},
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/httpRoutes/my-route",
			wantErr:      false,
		},
		{
			name: "missing reference",
			ref: &NetworkServicesHTTPRouteRef{
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
				t.Errorf("NetworkServicesHTTPRouteRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("NetworkServicesHTTPRouteRef.Normalize() got = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
