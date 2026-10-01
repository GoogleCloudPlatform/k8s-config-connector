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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCanonicalizeSubnetworkValue(t *testing.T) {
	cache := projects.NewProjectCache(nil, time.Hour)
	cache.InsertForTest("my-project", 12345)
	projectMapper := projects.NewProjectMapper(cache)

	tests := []struct {
		name            string
		val             string
		parentProjectID string
		parentLocation  string
		want            string
		wantErr         bool
	}{
		{
			name:            "empty",
			val:             "",
			parentProjectID: "my-project",
			parentLocation:  "us-central1",
			want:            "",
		},
		{
			name:            "canonical relative path",
			val:             "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
			parentProjectID: "other-project",
			parentLocation:    "europe-west1",
			want:            "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
		},
		{
			name:            "full HTTPS URI compute.googleapis.com",
			val:             "https://compute.googleapis.com/compute/v1/projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
			parentProjectID: "other-project",
			parentLocation:    "europe-west1",
			want:            "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
		},
		{
			name:            "full HTTPS URI www.googleapis.com",
			val:             "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
			parentProjectID: "other-project",
			parentLocation:    "europe-west1",
			want:            "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
		},
		{
			name:            "short name with regional parent",
			val:             "my-subnetwork",
			parentProjectID: "my-project",
			parentLocation:    "us-central1",
			want:            "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
		},
		{
			name:            "short name with zonal parent location",
			val:             "my-subnetwork",
			parentProjectID: "my-project",
			parentLocation:    "us-central1-a",
			want:            "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
		},
		{
			name:            "project number in relative path resolved to project ID",
			val:             "projects/12345/regions/us-central1/subnetworks/my-subnetwork",
			parentProjectID: "my-project",
			parentLocation:    "us-central1",
			want:            "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
		},
		{
			name:            "short name with missing parent region returns name as is",
			val:             "my-subnetwork",
			parentProjectID: "my-project",
			parentLocation:    "",
			want:            "my-subnetwork",
		},
		{
			name:            "short name with missing parent project returns name as is",
			val:             "my-subnetwork",
			parentProjectID: "",
			parentLocation:    "us-central1",
			want:            "my-subnetwork",
		},
		{
			name:            "invalid external with slashes returns error",
			val:             "projects/invalid/format",
			parentProjectID: "my-project",
			parentLocation:    "us-central1",
			wantErr:         true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			got, err := CanonicalizeSubnetworkValue(ctx, tc.val, tc.parentProjectID, tc.parentLocation, projectMapper)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("CanonicalizeSubnetworkValue(%q) expected error, got nil", tc.val)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("CanonicalizeSubnetworkValue(%q, %q, %q) = %q, want %q", tc.val, tc.parentProjectID, tc.parentLocation, got, tc.want)
			}
		})
	}
}

func TestCanonicalizeAndNormalizeSubnetwork(t *testing.T) {
	cache := projects.NewProjectCache(nil, time.Hour)
	cache.InsertForTest("my-project", 12345)
	projectMapper := projects.NewProjectMapper(cache)

	tests := []struct {
		name            string
		ref             *ComputeSubnetworkRef
		unstructured    *unstructured.Unstructured
		parentProjectID string
		parentLocation    string
		want            *ComputeSubnetworkRef
	}{
		{
			name: "nil ref",
			ref:  nil,
			want: nil,
		},
		{
			name: "short name with parent project and region",
			ref: &ComputeSubnetworkRef{
				External: "default",
			},
			parentProjectID: "my-project",
			parentLocation:    "us-central1",
			want: &ComputeSubnetworkRef{
				External: "projects/my-project/regions/us-central1/subnetworks/default",
			},
		},
		{
			name: "short name with zonal location",
			ref: &ComputeSubnetworkRef{
				External: "default",
			},
			parentProjectID: "my-project",
			parentLocation:    "us-central1-b",
			want: &ComputeSubnetworkRef{
				External: "projects/my-project/regions/us-central1/subnetworks/default",
			},
		},
		{
			name: "full HTTPS URI",
			ref: &ComputeSubnetworkRef{
				External: "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
			},
			parentProjectID: "other-project",
			parentLocation:    "europe-west1",
			want: &ComputeSubnetworkRef{
				External: "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
			},
		},
		{
			name: "project number resolved to project ID",
			ref: &ComputeSubnetworkRef{
				External: "projects/12345/regions/us-central1/subnetworks/my-subnetwork",
			},
			want: &ComputeSubnetworkRef{
				External: "projects/my-project/regions/us-central1/subnetworks/my-subnetwork",
			},
		},
		{
			name: "k8s object reference with selfLink",
			ref: &ComputeSubnetworkRef{
				Name:      "sn1",
				Namespace: "ns1",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"selfLink": "https://www.googleapis.com/compute/v1/projects/p1/regions/us-central1/subnetworks/sn1",
					},
				},
			},
			parentProjectID: "other-project",
			parentLocation:    "europe-west1",
			want: &ComputeSubnetworkRef{
				External: "projects/p1/regions/us-central1/subnetworks/sn1",
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
				tc.unstructured.SetGroupVersionKind(ComputeSubnetworkGVK)
				objs = append(objs, *tc.unstructured)
			}

			s := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: objs}).Build()

			defaultNamespace := ""
			if tc.ref != nil {
				defaultNamespace = tc.ref.Namespace
			}

			if err := tc.ref.CanonicalizeAndNormalize(ctx, s, defaultNamespace, tc.parentProjectID, tc.parentLocation, projectMapper); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.ref, tc.want); diff != "" {
				t.Errorf("CanonicalizeAndNormalize() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
