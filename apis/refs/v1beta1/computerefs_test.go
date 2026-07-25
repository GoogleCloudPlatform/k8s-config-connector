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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestResolveComputeBackendBucket(t *testing.T) {
	tests := []struct {
		name         string
		ref          *ComputeBackendBucketRef
		unstructured *unstructured.Unstructured
		want         *ComputeBackendBucketRef
		wantErr      bool
	}{
		{
			name: "external already populated",
			ref: &ComputeBackendBucketRef{
				External: "projects/p1/global/backendBuckets/b1",
			},
			want: &ComputeBackendBucketRef{
				External: "projects/p1/global/backendBuckets/b1",
			},
			wantErr: false,
		},
		{
			name: "resolve from name and namespace with selfLink",
			ref: &ComputeBackendBucketRef{
				Name:      "b1",
				Namespace: "ns1",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"selfLink": "https://www.googleapis.com/compute/v1/projects/p1/global/backendBuckets/b1",
					},
				},
			},
			want: &ComputeBackendBucketRef{
				External:  "projects/p1/global/backendBuckets/b1",
				Name:      "b1",
				Namespace: "ns1",
			},
			wantErr: false,
		},
		{
			name: "resolve empty name and external returns error",
			ref: &ComputeBackendBucketRef{
				Namespace: "ns1",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			var objs []unstructured.Unstructured
			if tc.unstructured != nil {
				tc.unstructured.SetName(tc.ref.Name)
				tc.unstructured.SetNamespace(tc.ref.Namespace)
				tc.unstructured.SetGroupVersionKind(schema.GroupVersionKind{
					Group:   "compute.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "ComputeBackendBucket",
				})
				objs = append(objs, *tc.unstructured)
			}

			s := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: objs}).Build()

			err := ResolveComputeBackendBucket(ctx, s, tc.ref.Namespace, tc.ref)
			if (err != nil) != tc.wantErr {
				t.Errorf("ResolveComputeBackendBucket() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				if diff := cmp.Diff(tc.ref, tc.want); diff != "" {
					t.Errorf("ResolveComputeBackendBucket() mismatch (-got +want):\n%s", diff)
				}
			}
		})
	}
}
