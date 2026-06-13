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
// See the_identity.go specific language governing permissions and
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

func TestComputeGlobalNetworkEndpointIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeGlobalNetworkEndpointIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/global/networkEndpointGroups/my-neg/networkEndpoints/80",
			want: &ComputeGlobalNetworkEndpointIdentity{
				Project:                    "my-project",
				GlobalNetworkEndpointGroup: "my-neg",
				Port:                       "80",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/global/networkEndpointGroups/my-neg/networkEndpoints/8080",
			want: &ComputeGlobalNetworkEndpointIdentity{
				Project:                    "my-project",
				GlobalNetworkEndpointGroup: "my-neg",
				Port:                       "8080",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeGlobalNetworkEndpointIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeGlobalNetworkEndpointRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/global/networkEndpointGroups/my-neg/networkEndpoints/80",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/global/networkEndpointGroups/my-neg/networkEndpoints/80",
			wantErr: true,
		},
		{
			name:    "missing port",
			ref:     "projects/my-project/global/networkEndpointGroups/my-neg/networkEndpoints",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeGlobalNetworkEndpointRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeGlobalNetworkEndpointRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeGlobalNetworkEndpointRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeGlobalNetworkEndpointRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeGlobalNetworkEndpointRef{
				External: "projects/test-project/global/networkEndpointGroups/test-neg/networkEndpoints/80",
			},
			wantExternal: "projects/test-project/global/networkEndpointGroups/test-neg/networkEndpoints/80",
		},
		{
			name: "external with invalid format",
			ref: &ComputeGlobalNetworkEndpointRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeGlobalNetworkEndpoint external="invalid-format" was not known (use projects/{project}/global/networkEndpointGroups/{globalNetworkEndpointGroup}/networkEndpoints/{port})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeGlobalNetworkEndpointRef{
				Name:      "test-endpoint",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeGlobalNetworkEndpoint",
						"metadata": map[string]interface{}{
							"name":      "test-endpoint",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/global/networkEndpointGroups/test-neg/networkEndpoints/80",
						},
					},
				},
			},
			wantExternal: "projects/test-project/global/networkEndpointGroups/test-neg/networkEndpoints/80",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeGlobalNetworkEndpointRef{
				Name:      "test-endpoint",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeGlobalNetworkEndpoint",
						"metadata": map[string]interface{}{
							"name":      "test-endpoint",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeGlobalNetworkEndpoint my-namespace/test-endpoint is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeGlobalNetworkEndpointRef{
				Name:      "test-endpoint",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeGlobalNetworkEndpoint my-namespace/test-endpoint is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeGlobalNetworkEndpoint{})
			cl := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()

			err := tc.ref.Normalize(context.TODO(), cl, "default")
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("Normalize() expected error %q, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Errorf("Normalize() error = %q, want %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Normalize() unexpected error: %v", err)
			}
			if tc.ref.External != tc.wantExternal {
				t.Errorf("Normalize() external = %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}
