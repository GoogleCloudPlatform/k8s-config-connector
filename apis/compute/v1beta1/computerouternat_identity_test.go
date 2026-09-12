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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeRouterNATIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeRouterNATIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/routers/my-router/my-nat",
			want: &ComputeRouterNATIdentity{
				Project:          "my-project",
				Region:           "us-central1",
				Router:           "my-router",
				ComputeRouterNAT: "my-nat",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/routers/my-router/my-nat",
			want: &ComputeRouterNATIdentity{
				Project:          "my-project",
				Region:           "us-central1",
				Router:           "my-router",
				ComputeRouterNAT: "my-nat",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeRouterNATIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestGetIdentity(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(GroupVersion, &ComputeRouterNAT{}, &unstructured.Unstructured{})

	tests := []struct {
		name       string
		obj        *ComputeRouterNAT
		objs       []runtime.Object
		want       *ComputeRouterNATIdentity
		wantErrSub string
	}{
		{
			name: "1. Short external router name with project-id annotation and spec.region",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "proj-1",
					},
				},
				Spec: ComputeRouterNATSpec{
					Region: "region-1",
					RouterRef: ComputeRouterRef{
						External: "router-1",
					},
				},
			},
			want: &ComputeRouterNATIdentity{
				Project:          "proj-1",
				Region:           "region-1",
				Router:           "router-1",
				ComputeRouterNAT: "nat-1",
			},
		},
		{
			name: "2. Short external router name without project-id annotation",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: ComputeRouterNATSpec{
					Region: "region-1",
					RouterRef: ComputeRouterRef{
						External: "router-1",
					},
				},
			},
			wantErrSub: "cannot resolve project: please set the 'cnrm.cloud.google.com/project-id' annotation",
		},
		{
			name: "3. Short external router name without spec.region",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "proj-1",
					},
				},
				Spec: ComputeRouterNATSpec{
					RouterRef: ComputeRouterRef{
						External: "router-1",
					},
				},
			},
			wantErrSub: "spec.region is required when routerRef is a short name",
		},
		{
			name: "4. Canonical path",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: ComputeRouterNATSpec{
					RouterRef: ComputeRouterRef{
						External: "projects/p1/regions/r1/routers/example-router",
					},
				},
			},
			want: &ComputeRouterNATIdentity{
				Project:          "p1",
				Region:           "r1",
				Router:           "example-router",
				ComputeRouterNAT: "nat-1",
			},
		},
		{
			name: "5. Canonical path with conflicting spec.region",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: ComputeRouterNATSpec{
					Region: "r2",
					RouterRef: ComputeRouterRef{
						External: "projects/p1/regions/r1/routers/example-router",
					},
				},
			},
			wantErrSub: "spec.region (\"r2\") does not match routerRef region (\"r1\")",
		},
		{
			name: "6. Canonical path with conflicting project-id annotation",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "p2",
					},
				},
				Spec: ComputeRouterNATSpec{
					RouterRef: ComputeRouterRef{
						External: "projects/p1/regions/r1/routers/example-router",
					},
				},
			},
			wantErrSub: "project-id annotation (\"p2\") does not match routerRef project (\"p1\")",
		},
		{
			name: "7. Full URI",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: ComputeRouterNATSpec{
					RouterRef: ComputeRouterRef{
						External: "https://www.googleapis.com/compute/v1/projects/p1/regions/r1/routers/example-router",
					},
				},
			},
			want: &ComputeRouterNATIdentity{
				Project:          "p1",
				Region:           "r1",
				Router:           "example-router",
				ComputeRouterNAT: "nat-1",
			},
		},
		{
			name: "8a. KRM referent object with status.externalRef",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: ComputeRouterNATSpec{
					RouterRef: ComputeRouterRef{
						Name: "router-1",
					},
				},
			},
			objs: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeRouter",
						"metadata": map[string]interface{}{
							"name":      "router-1",
							"namespace": "ns-1",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/p1/regions/r1/routers/example-router",
						},
					},
				},
			},
			want: &ComputeRouterNATIdentity{
				Project:          "p1",
				Region:           "r1",
				Router:           "example-router",
				ComputeRouterNAT: "nat-1",
			},
		},
		{
			name: "8b. KRM referent object with status.selfLink fallback",
			obj: &ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: ComputeRouterNATSpec{
					RouterRef: ComputeRouterRef{
						Name: "router-1",
					},
				},
			},
			objs: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeRouter",
						"metadata": map[string]interface{}{
							"name":      "router-1",
							"namespace": "ns-1",
						},
						"status": map[string]interface{}{
							"selfLink": "https://www.googleapis.com/compute/v1/projects/p1/regions/r1/routers/example-router",
						},
					},
				},
			},
			want: &ComputeRouterNATIdentity{
				Project:          "p1",
				Region:           "r1",
				Router:           "example-router",
				ComputeRouterNAT: "nat-1",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			fakeClient := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objs...).Build()

			got, err := tc.obj.GetIdentity(ctx, fakeClient)
			if tc.wantErrSub != "" {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tc.wantErrSub) {
					t.Fatalf("expected error containing %q, got: %v", tc.wantErrSub, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.want, got.(*ComputeRouterNATIdentity)); diff != "" {
				t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
