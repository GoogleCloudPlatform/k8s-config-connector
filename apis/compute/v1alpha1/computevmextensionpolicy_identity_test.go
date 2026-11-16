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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeVMExtensionPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeVMExtensionPolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/zones/us-central1-a/vmExtensionPolicies/my-policy",
			want: &ComputeVMExtensionPolicyIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				Vmextensionpolicy: "my-policy",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "reference with host prefix",
			ref:  "//compute.googleapis.com/projects/my-project/zones/us-central1-a/vmExtensionPolicies/my-policy",
			want: &ComputeVMExtensionPolicyIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				Vmextensionpolicy: "my-policy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeVMExtensionPolicyIdentity{}
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

func TestComputeVMExtensionPolicyIdentity_ParentString(t *testing.T) {
	id := &ComputeVMExtensionPolicyIdentity{
		Project:           "my-project",
		Zone:              "us-central1-a",
		Vmextensionpolicy: "my-policy",
	}
	want := "projects/my-project/zones/us-central1-a"
	got := id.ParentString()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ParentString() mismatch (-want +got):\n%s", diff)
	}
}

func TestComputeVMExtensionPolicyIdentity_GetIdentity(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(GroupVersion, &ComputeVMExtensionPolicy{})
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	tests := []struct {
		name    string
		obj     *ComputeVMExtensionPolicy
		want    *ComputeVMExtensionPolicyIdentity
		wantErr bool
	}{
		{
			name: "from spec with resourceID",
			obj: &ComputeVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location:   "us-central1-a",
					ResourceID: ptrTo("custom-policy"),
				},
			},
			want: &ComputeVMExtensionPolicyIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				Vmextensionpolicy: "custom-policy",
			},
		},
		{
			name: "from spec defaulting to metadata.name",
			obj: &ComputeVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1-a",
				},
			},
			want: &ComputeVMExtensionPolicyIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				Vmextensionpolicy: "k8s-policy",
			},
		},
		{
			name: "matching status.externalRef",
			obj: &ComputeVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1-a",
				},
				Status: ComputeVMExtensionPolicyStatus{
					ExternalRef: ptrTo("projects/my-project/zones/us-central1-a/vmExtensionPolicies/k8s-policy"),
				},
			},
			want: &ComputeVMExtensionPolicyIdentity{
				Project:           "my-project",
				Zone:              "us-central1-a",
				Vmextensionpolicy: "k8s-policy",
			},
		},
		{
			name: "mismatched status.externalRef",
			obj: &ComputeVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1-a",
				},
				Status: ComputeVMExtensionPolicyStatus{
					ExternalRef: ptrTo("projects/other-project/zones/us-central1-a/vmExtensionPolicies/k8s-policy"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(context.Background(), cl)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeVMExtensionPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/zones/us-central1-a/vmExtensionPolicies/my-policy",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/zones/us-central1-a/vmExtensionPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "missing zone",
			ref:     "projects/my-project/vmExtensionPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "missing policy id",
			ref:     "projects/my-project/zones/us-central1-a/vmExtensionPolicies",
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
			r := &ComputeVMExtensionPolicyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeVMExtensionPolicyRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeVMExtensionPolicyRef_Normalize(t *testing.T) {
	testCases := []struct {
		name         string
		ref          *ComputeVMExtensionPolicyRef
		objects      []runtime.Object
		wantExternal string
		wantErr      string
	}{
		{
			name: "external with valid format",
			ref: &ComputeVMExtensionPolicyRef{
				External: "projects/test-project/zones/us-central1-a/vmExtensionPolicies/test-policy",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/vmExtensionPolicies/test-policy",
		},
		{
			name: "external with invalid format",
			ref: &ComputeVMExtensionPolicyRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeVMExtensionPolicy external="invalid-format" was not known (use projects/{project}/zones/{zone}/vmExtensionPolicies/{vmextensionpolicy})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeVMExtensionPolicyRef{
				Name:      "test-policy",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeVMExtensionPolicy",
						"metadata": map[string]interface{}{
							"name":      "test-policy",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/zones/us-central1-a/vmExtensionPolicies/test-policy",
						},
					},
				},
			},
			wantExternal: "projects/test-project/zones/us-central1-a/vmExtensionPolicies/test-policy",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeVMExtensionPolicyRef{
				Name:      "test-policy",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeVMExtensionPolicy",
						"metadata": map[string]interface{}{
							"name":      "test-policy",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeVMExtensionPolicy my-namespace/test-policy is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeVMExtensionPolicyRef{
				Name:      "test-policy",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeVMExtensionPolicy my-namespace/test-policy is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeVMExtensionPolicy{})
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

func ptrTo[T any](v T) *T {
	return &v
}
