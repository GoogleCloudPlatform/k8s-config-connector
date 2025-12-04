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

func TestComputeNetworkRefNormalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeNetworkRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeNetworkRef{
				External: "projects/test-project/global/networks/test-network",
			},
			wantExternal: "projects/test-project/global/networks/test-network",
		},
		{
			name: "external with invalid format",
			ref: &ComputeNetworkRef{
				External: "invalid-format",
			},
			wantErr: `format of computenetwork external="invalid-format" was not known (use https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/global/networks/{{networkId}} or projects/{{projectId}}/global/networks/{{networkId}})`,
		},
		{
			name: "external is full url(v1)",
			ref: &ComputeNetworkRef{
				External: "https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network",
			},
			wantExternal: "projects/test-project/global/networks/test-network",
		},
		{
			name: "external is full url(v1beta1)",
			ref: &ComputeNetworkRef{
				External: "https://www.googleapis.com/compute/v1beta1/projects/test-project/global/networks/test-network",
			},
			wantExternal: "projects/test-project/global/networks/test-network",
		},
		{
			name: "external is full url(beta)",
			ref: &ComputeNetworkRef{
				External: "https://www.googleapis.com/compute/beta/projects/test-project/global/networks/test-network",
			},
			wantExternal: "projects/test-project/global/networks/test-network",
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeNetworkRef{
				Name:      "test-network",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeNetwork",
						"metadata": map[string]interface{}{
							"name":      "test-network",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/global/networks/test-network",
						},
					},
				},
			},
			wantExternal: "projects/test-project/global/networks/test-network",
		},
		{
			name: "name specified, with status.selfLink",
			ref: &ComputeNetworkRef{
				Name:      "test-network",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeNetwork",
						"metadata": map[string]interface{}{
							"name":      "test-network",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"selfLink": "https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network",
						},
					},
				},
			},
			wantExternal: "projects/test-project/global/networks/test-network",
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeNetworkRef{
				Name:      "test-network",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeNetwork my-namespace/test-network is not found`,
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
