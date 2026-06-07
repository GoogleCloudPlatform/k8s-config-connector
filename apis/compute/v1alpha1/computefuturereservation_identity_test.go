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

package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeFutureReservationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeFutureReservationIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/zones/us-central1-a/futureReservations/my-fr",
			want: &ComputeFutureReservationIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				FutureReservation: "my-fr",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/zones/us-central1-a/futureReservations/my-fr",
			want: &ComputeFutureReservationIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				FutureReservation: "my-fr",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeFutureReservationIdentity{}
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

func TestComputeFutureReservationRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/zones/us-central1-a/futureReservations/my-fr",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/zones/us-central1-a/futureReservations/my-fr",
			wantErr: true,
		},
		{
			name:    "missing zone",
			ref:     "projects/my-project/futureReservations/my-fr",
			wantErr: true,
		},
		{
			name:    "missing futureReservation",
			ref:     "projects/my-project/zones/us-central1-a",
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
			r := &ComputeFutureReservationRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeFutureReservationRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeFutureReservationRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeFutureReservationRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeFutureReservationRef{
				External: "projects/test-project/zones/us-central1-a/futureReservations/test-fr",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/futureReservations/test-fr",
		},
		{
			name: "external with invalid format",
			ref: &ComputeFutureReservationRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeFutureReservation external="invalid-format" was not known (use projects/{project}/zones/{zone}/futureReservations/{futureReservation})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeFutureReservationRef{
				Name:      "test-fr",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeFutureReservation",
						"metadata": map[string]interface{}{
							"name":      "test-fr",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/zones/us-central1-a/futureReservations/test-fr",
						},
					},
				},
			},
			wantExternal: "projects/test-project/zones/us-central1-a/futureReservations/test-fr",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeFutureReservationRef{
				Name:      "test-fr",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeFutureReservation",
						"metadata": map[string]interface{}{
							"name":      "test-fr",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeFutureReservation my-namespace/test-fr is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeFutureReservationRef{
				Name:      "test-fr",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeFutureReservation my-namespace/test-fr is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeFutureReservation{})
			cl := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()

			err := tc.ref.Normalize(context.TODO(), cl, "default")
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("Normalize() expected error %q, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Errorf("Normalize() error = %q, want %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Normalize() unexpected error: %v", err)
			}
			if tc.ref.External != tc.wantExternal {
				t.Errorf("Normalize() external = %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}
