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

func TestComputeReservationRefNormalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeReservationRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeReservationRef{
				External: "projects/test-project/zones/us-central1-a/reservations/test-reservation",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/reservations/test-reservation",
		},
		{
			name: "external with invalid format",
			ref: &ComputeReservationRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeReservation external="invalid-format" was not known (use projects/{project}/zones/{zone}/reservations/{reservation})`,
		},
		{
			name: "external is full url",
			ref: &ComputeReservationRef{
				External: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/reservations/test-reservation",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/reservations/test-reservation",
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeReservationRef{
				Name:      "test-reservation",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeReservation",
						"metadata": map[string]interface{}{
							"name":      "test-reservation",
							"namespace": "my-namespace",
						},
						"spec": map[string]interface{}{
							"zone": "us-central1-a",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/zones/us-central1-a/reservations/test-reservation",
						},
					},
				},
			},
			wantExternal: "projects/test-project/zones/us-central1-a/reservations/test-reservation",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &ComputeReservation{}, &unstructured.Unstructured{})
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
