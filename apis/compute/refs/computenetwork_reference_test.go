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

package computerefs

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeNetworkRefNormalize(t *testing.T) {
	tests := []struct {
		name           string
		ref            *ComputeNetworkRef
		unstructured   *unstructured.Unstructured
		defaultProject string
		want           *ComputeNetworkRef
	}{
		{
			name: "external only",
			ref: &ComputeNetworkRef{
				External: "projects/p1/global/networks/n1",
			},
			want: &ComputeNetworkRef{
				External: "projects/p1/global/networks/n1",
			},
		},
		{
			name: "external URI prefix",
			ref: &ComputeNetworkRef{
				External: "https://www.googleapis.com/compute/v1/projects/p1/global/networks/n1",
			},
			want: &ComputeNetworkRef{
				External: "projects/p1/global/networks/n1",
			},
		},
		{
			name: "external status selfLink",
			ref: &ComputeNetworkRef{
				Name:      "n1",
				Namespace: "ns1",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"selfLink": "https://www.googleapis.com/compute/v1/projects/p1/global/networks/n1",
					},
				},
			},
			want: &ComputeNetworkRef{
				External: "projects/p1/global/networks/n1",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			var objs []unstructured.Unstructured
			if tc.unstructured != nil {
				tc.unstructured.SetName(tc.ref.Name)
				tc.unstructured.SetNamespace(tc.ref.Namespace)
				tc.unstructured.SetGroupVersionKind(ComputeNetworkGVK)
				objs = append(objs, *tc.unstructured)
			}

			s := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: objs}).Build()

			if err := tc.ref.Normalize(ctx, s, tc.ref.Namespace); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.ref, tc.want); diff != "" {
				t.Errorf("Normalize() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
