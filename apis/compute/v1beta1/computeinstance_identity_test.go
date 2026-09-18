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

func TestComputeInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *ComputeInstanceIdentity
		wantString string
		wantParent string
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/zones/us-central1-a/instances/my-instance",
			want: &ComputeInstanceIdentity{
				Project:  "my-project",
				Zone:     "us-central1-a",
				Instance: "my-instance",
			},
			wantString: "projects/my-project/zones/us-central1-a/instances/my-instance",
			wantParent: "projects/my-project/zones/us-central1-a",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url with compute.googleapis.com",
			ref:  "https://www.googleapis.com/compute/v1/projects/my-project/zones/us-central1-a/instances/my-instance",
			want: &ComputeInstanceIdentity{
				Project:  "my-project",
				Zone:     "us-central1-a",
				Instance: "my-instance",
			},
			wantString: "projects/my-project/zones/us-central1-a/instances/my-instance",
			wantParent: "projects/my-project/zones/us-central1-a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeInstanceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != tt.wantString {
					t.Errorf("String() = %q, want %q", i.String(), tt.wantString)
				}
				if i.ParentString() != tt.wantParent {
					t.Errorf("ParentString() = %q, want %q", i.ParentString(), tt.wantParent)
				}
			}
		})
	}
}

func TestInstanceRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *InstanceRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &InstanceRef{
				External: "projects/test-project/zones/us-central1-a/instances/test-instance",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/instances/test-instance",
		},
		{
			name: "external with invalid format",
			ref: &InstanceRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeInstance external="invalid-format" was not known (use projects/{project}/zones/{zone}/instances/{instance})`,
		},
		{
			name: "name specified, with status.selfLink and ready condition",
			ref: &InstanceRef{
				Name:      "test-instance",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeInstance",
						"metadata": map[string]interface{}{
							"name":      "test-instance",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"selfLink": "https://compute.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/instances/test-instance",
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
			wantExternal: "projects/test-project/zones/us-central1-a/instances/test-instance",
		},
		{
			name: "name specified, with status.selfLink but not ready",
			ref: &InstanceRef{
				Name:      "test-instance",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeInstance",
						"metadata": map[string]interface{}{
							"name":      "test-instance",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"selfLink": "https://compute.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/instances/test-instance",
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
			wantErr: `reference ComputeInstance my-namespace/test-instance is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &InstanceRef{
				Name:      "test-instance",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeInstance my-namespace/test-instance is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeInstance{})
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
