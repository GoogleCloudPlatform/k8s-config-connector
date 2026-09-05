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

func TestComputeGlobalNetworkEndpointGroupIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeGlobalNetworkEndpointGroupIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/global/networkEndpointGroups/my-neg",
			want: &ComputeGlobalNetworkEndpointGroupIdentity{
				Project:                    "my-project",
				GlobalNetworkEndpointGroup: "my-neg",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/global/networkEndpointGroups/my-neg",
			want: &ComputeGlobalNetworkEndpointGroupIdentity{
				Project:                    "my-project",
				GlobalNetworkEndpointGroup: "my-neg",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeGlobalNetworkEndpointGroupIdentity{}
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

func TestComputeGlobalNetworkEndpointGroupRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/global/networkEndpointGroups/my-neg",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/global/networkEndpointGroups/my-neg",
			wantErr: true,
		},
		{
			name:    "missing global/networkEndpointGroups",
			ref:     "projects/my-project/my-neg",
			wantErr: true,
		},
		{
			name:    "missing neg",
			ref:     "projects/my-project/global/networkEndpointGroups",
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
			r := &ComputeGlobalNetworkEndpointGroupRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeGlobalNetworkEndpointGroupRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeGlobalNetworkEndpointGroupRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeGlobalNetworkEndpointGroupRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeGlobalNetworkEndpointGroupRef{
				External: "projects/test-project/global/networkEndpointGroups/test-neg",
			},
			wantExternal: "projects/test-project/global/networkEndpointGroups/test-neg",
		},
		{
			name: "external with invalid format",
			ref: &ComputeGlobalNetworkEndpointGroupRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeGlobalNetworkEndpointGroup external="invalid-format" was not known (use projects/{project}/global/networkEndpointGroups/{globalNetworkEndpointGroup})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeGlobalNetworkEndpointGroupRef{
				Name:      "test-neg",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeGlobalNetworkEndpointGroup",
						"metadata": map[string]interface{}{
							"name":      "test-neg",
							"namespace": "my-namespace",
						},
						"spec": map[string]interface{}{
							"projectRef": map[string]interface{}{
								"name": "project-name",
							},
						},
						"status": map[string]interface{}{
							"externalRef": "projects/project-id/global/networkEndpointGroups/test-neg",
						},
					},
				},
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
						"kind":       "Project",
						"metadata": map[string]interface{}{
							"name":      "project-name",
							"namespace": "my-namespace",
						},
						"spec": map[string]interface{}{
							"resourceID": "project-id",
						},
					},
				},
			},
			wantExternal: "projects/project-id/global/networkEndpointGroups/test-neg",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			err := AddToScheme(s)
			if err != nil {
				t.Fatalf("failed to add scheme: %v", err)
			}
			cl := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()
			err = tc.ref.Normalize(context.Background(), cl, tc.otherNamespace)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error: %s, got none", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Fatalf("expected error: %s, got: %s", tc.wantErr, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if tc.ref.External != tc.wantExternal {
					t.Errorf("Normalize() mismatch got %q, want %q", tc.ref.External, tc.wantExternal)
				}
			}
		})
	}
}
