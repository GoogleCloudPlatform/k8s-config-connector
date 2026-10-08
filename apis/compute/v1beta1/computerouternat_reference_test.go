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

func TestComputeRouterNATRefNormalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeRouterNATRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeRouterNATRef{
				External: "projects/test-project/regions/us-central1/routers/my-router/test-nat",
			},
			wantExternal: "projects/test-project/regions/us-central1/routers/my-router/test-nat",
		},
		{
			name: "external with invalid format",
			ref: &ComputeRouterNATRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeRouterNAT external="invalid-format" was not known (use projects/{project}/regions/{region}/routers/{router}/{computerouternat})`,
		},
		{
			name: "external is full url",
			ref: &ComputeRouterNATRef{
				External: "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/routers/my-router/test-nat",
			},
			wantExternal: "projects/test-project/regions/us-central1/routers/my-router/test-nat",
		},
		{
			name: "name specified, but resource is not ready (no status conditions)",
			ref: &ComputeRouterNATRef{
				Name:      "test-nat",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeRouterNAT",
						"metadata": map[string]interface{}{
							"name":      "test-nat",
							"namespace": "my-namespace",
							"annotations": map[string]interface{}{
								"cnrm.cloud.google.com/project-id": "test-project",
							},
						},
						"spec": map[string]interface{}{
							"region": "us-central1",
							"routerRef": map[string]interface{}{
								"external": "projects/test-project/regions/us-central1/routers/my-router",
							},
						},
					},
				},
			},
			wantErr: "reference ComputeRouterNAT my-namespace/test-nat is not ready",
		},
		{
			name: "name specified, but resource is not ready (ready condition is false)",
			ref: &ComputeRouterNATRef{
				Name:      "test-nat",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeRouterNAT",
						"metadata": map[string]interface{}{
							"name":      "test-nat",
							"namespace": "my-namespace",
							"annotations": map[string]interface{}{
								"cnrm.cloud.google.com/project-id": "test-project",
							},
						},
						"spec": map[string]interface{}{
							"region": "us-central1",
							"routerRef": map[string]interface{}{
								"external": "projects/test-project/regions/us-central1/routers/my-router",
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
			wantErr: "reference ComputeRouterNAT my-namespace/test-nat is not ready",
		},
		{
			name: "name specified, and resource is ready",
			ref: &ComputeRouterNATRef{
				Name:      "test-nat",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeRouterNAT",
						"metadata": map[string]interface{}{
							"name":      "test-nat",
							"namespace": "my-namespace",
							"annotations": map[string]interface{}{
								"cnrm.cloud.google.com/project-id": "test-project",
							},
						},
						"spec": map[string]interface{}{
							"region": "us-central1",
							"routerRef": map[string]interface{}{
								"external": "projects/test-project/regions/us-central1/routers/my-router",
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
			wantExternal: "projects/test-project/regions/us-central1/routers/my-router/test-nat",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &ComputeRouterNAT{}, &unstructured.Unstructured{})
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
