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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"

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

func TestCanonicalizeNetworkValue(t *testing.T) {
	tests := []struct {
		name            string
		val             string
		parentProjectID string
		want            string
	}{
		{
			name: "empty string",
			val:  "",
			want: "",
		},
		{
			name:            "short network name with parent project",
			val:             "default",
			parentProjectID: "my-project",
			want:            "projects/my-project/global/networks/default",
		},
		{
			name:            "short network name without parent project",
			val:             "default",
			parentProjectID: "",
			want:            "default",
		},
		{
			name:            "relative resource path",
			val:             "projects/my-project/global/networks/my-vpc",
			parentProjectID: "other-project",
			want:            "projects/my-project/global/networks/my-vpc",
		},
		{
			name:            "full HTTPS URI",
			val:             "https://www.googleapis.com/compute/v1/projects/my-project/global/networks/my-vpc",
			parentProjectID: "other-project",
			want:            "projects/my-project/global/networks/my-vpc",
		},
		{
			name:            "invalid external format",
			val:             "projects/my-project/invalid/path",
			parentProjectID: "my-project",
			want:            "projects/my-project/invalid/path",
		},
		{
			name:            "project number present in mapper cache",
			val:             "projects/12345/global/networks/my-vpc",
			parentProjectID: "",
			want:            "projects/my-project/global/networks/my-vpc",
		},
		{
			name:            "project number not in mapper cache",
			val:             "projects/99999/global/networks/my-vpc",
			parentProjectID: "",
			want:            "projects/99999/global/networks/my-vpc",
		},
	}

	cache := projects.NewProjectCache(nil, time.Hour)
	cache.InsertForTest("my-project", 12345)
	projectMapper := projects.NewProjectMapper(cache)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			got := canonicalizeNetworkValue(ctx, tc.val, tc.parentProjectID, projectMapper)
			if got != tc.want {
				t.Errorf("canonicalizeNetworkValue(%q, %q) = %q, want %q", tc.val, tc.parentProjectID, got, tc.want)
			}
		})
	}
}

func TestCanonicalizeAndNormalize(t *testing.T) {
	cache := projects.NewProjectCache(nil, time.Hour)
	cache.InsertForTest("my-project", 12345)
	projectMapper := projects.NewProjectMapper(cache)

	tests := []struct {
		name            string
		ref             *ComputeNetworkRef
		unstructured    *unstructured.Unstructured
		parentProjectID string
		want            *ComputeNetworkRef
	}{
		{
			name: "nil ref",
			ref:  nil,
			want: nil,
		},
		{
			name: "short name with parent project",
			ref: &ComputeNetworkRef{
				External: "default",
			},
			parentProjectID: "my-project",
			want: &ComputeNetworkRef{
				External: "projects/my-project/global/networks/default",
			},
		},
		{
			name: "full HTTPS URI",
			ref: &ComputeNetworkRef{
				External: "https://www.googleapis.com/compute/v1/projects/my-project/global/networks/my-vpc",
			},
			parentProjectID: "other-project",
			want: &ComputeNetworkRef{
				External: "projects/my-project/global/networks/my-vpc",
			},
		},
		{
			name: "project number in external resolved to project ID",
			ref: &ComputeNetworkRef{
				External: "projects/12345/global/networks/my-vpc",
			},
			want: &ComputeNetworkRef{
				External: "projects/my-project/global/networks/my-vpc",
			},
		},
		{
			name: "k8s object reference",
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
			parentProjectID: "other-project",
			want: &ComputeNetworkRef{
				External: "projects/p1/global/networks/n1",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			var objs []unstructured.Unstructured
			if tc.unstructured != nil && tc.ref != nil {
				tc.unstructured.SetName(tc.ref.Name)
				tc.unstructured.SetNamespace(tc.ref.Namespace)
				tc.unstructured.SetGroupVersionKind(ComputeNetworkGVK)
				objs = append(objs, *tc.unstructured)
			}

			s := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: objs}).Build()

			defaultNamespace := ""
			if tc.ref != nil {
				defaultNamespace = tc.ref.Namespace
			}

			if err := tc.ref.CanonicalizeAndNormalize(ctx, s, defaultNamespace, tc.parentProjectID, projectMapper); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.ref, tc.want); diff != "" {
				t.Errorf("CanonicalizeAndNormalize() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

