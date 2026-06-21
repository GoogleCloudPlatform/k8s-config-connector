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

func TestComputeAutoscalerIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeAutoscalerIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/zones/us-central1-a/autoscalers/my-autoscaler",
			want: &ComputeAutoscalerIdentity{
				Project:    "my-project",
				Zone:       "us-central1-a",
				Autoscaler: "my-autoscaler",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/zones/us-central1-a/autoscalers/my-autoscaler",
			want: &ComputeAutoscalerIdentity{
				Project:    "my-project",
				Zone:       "us-central1-a",
				Autoscaler: "my-autoscaler",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeAutoscalerIdentity{}
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

func TestComputeAutoscalerRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/zones/us-central1-a/autoscalers/my-autoscaler",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/zones/us-central1-a/autoscalers/my-autoscaler",
			wantErr: true,
		},
		{
			name:    "missing zone",
			ref:     "projects/my-project/autoscalers/my-autoscaler",
			wantErr: true,
		},
		{
			name:    "missing autoscaler",
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
			r := &ComputeAutoscalerRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeAutoscalerRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeAutoscalerRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeAutoscalerRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeAutoscalerRef{
				External: "projects/test-project/zones/us-central1-a/autoscalers/test-autoscaler",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/autoscalers/test-autoscaler",
		},
		{
			name: "external with invalid format",
			ref: &ComputeAutoscalerRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeAutoscaler external="invalid-format" was not known (use projects/{project}/zones/{zone}/autoscalers/{autoscaler})`,
		},
		{
			name: "name specified, with status.selfLink",
			ref: &ComputeAutoscalerRef{
				Name:      "test-autoscaler",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeAutoscaler",
						"metadata": map[string]interface{}{
							"name":      "test-autoscaler",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"selfLink": "https://compute.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/autoscalers/test-autoscaler",
						},
					},
				},
			},
			wantExternal: "projects/test-project/zones/us-central1-a/autoscalers/test-autoscaler",
		},
		{
			name: "name specified, without status.selfLink",
			ref: &ComputeAutoscalerRef{
				Name:      "test-autoscaler",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeAutoscaler",
						"metadata": map[string]interface{}{
							"name":      "test-autoscaler",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeAutoscaler my-namespace/test-autoscaler is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeAutoscalerRef{
				Name:      "test-autoscaler",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeAutoscaler my-namespace/test-autoscaler is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeAutoscaler{})
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
