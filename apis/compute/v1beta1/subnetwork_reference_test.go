// Copyright 2025 Google LLC
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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeSubnetworkRefNormalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeSubnetworkRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeSubnetworkRef{
				External: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
			},
			wantExternal: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
		},
		{
			name: "external with invalid format",
			ref: &ComputeSubnetworkRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeSubnetwork external="invalid-format" was not known (use https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/regions/{{region}}/subnetworks/{{subnetworkId}} or projects/{{projectId}}/regions/{{region}}/subnetworks/{{subnetworkId}})`,
		},
		{
			name: "external is full url(v1)",
			ref: &ComputeSubnetworkRef{
				External: "https://www.googleapis.com/compute/v1/projects/test-project/regions/test-region/subnetworks/test-subnetwork",
			},
			wantExternal: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
		},
		{
			name: "external is full url(v1beta1)",
			ref: &ComputeSubnetworkRef{
				External: "https://www.googleapis.com/compute/v1beta1/projects/test-project/regions/test-region/subnetworks/test-subnetwork",
			},
			wantExternal: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
		},
		{
			name: "external is full url(beta)",
			ref: &ComputeSubnetworkRef{
				External: "https://www.googleapis.com/compute/beta/projects/test-project/regions/test-region/subnetworks/test-subnetwork",
			},
			wantExternal: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeSubnetworkRef{
				Name:      "test-subnetwork",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeSubnetwork",
						"metadata": map[string]interface{}{
							"name":      "test-subnetwork",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
						},
					},
				},
			},
			wantExternal: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
		},
		{
			name: "name specified, with status.selfLink",
			ref: &ComputeSubnetworkRef{
				Name:      "test-subnetwork",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeSubnetwork",
						"metadata": map[string]interface{}{
							"name":      "test-subnetwork",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"selfLink": "https://www.googleapis.com/compute/v1/projects/test-project/regions/test-region/subnetworks/test-subnetwork",
						},
					},
				},
			},
			wantExternal: "projects/test-project/regions/test-region/subnetworks/test-subnetwork",
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeSubnetworkRef{
				Name:      "test-subnetwork",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeSubnetwork my-namespace/test-subnetwork is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(schema.GroupVersion{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1"}, &unstructured.Unstructured{})
			fakeClient := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()

			err := tc.ref.Normalize(context.TODO(), fakeClient, tc.otherNamespace)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("got nil error, want %q", tc.wantErr)
				}
				if !cmp.Equal(err.Error(), tc.wantErr) {
					t.Errorf("got error %q, want %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			gotExternal := tc.ref.External
			if gotExternal != tc.wantExternal {
				t.Errorf("got external %q, want %q", gotExternal, tc.wantExternal)
			}
		})
	}
}
