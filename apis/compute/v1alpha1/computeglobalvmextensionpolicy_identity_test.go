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

func TestComputeGlobalVMExtensionPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeGlobalVMExtensionPolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/global/vmExtensionPolicies/my-policy",
			want: &ComputeGlobalVMExtensionPolicyIdentity{
				Project:                 "my-project",
				Globalvmextensionpolicy: "my-policy",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "reference with host prefix",
			ref:  "//compute.googleapis.com/projects/my-project/global/vmExtensionPolicies/my-policy",
			want: &ComputeGlobalVMExtensionPolicyIdentity{
				Project:                 "my-project",
				Globalvmextensionpolicy: "my-policy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeGlobalVMExtensionPolicyIdentity{}
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

func TestComputeGlobalVMExtensionPolicyIdentity_ParentString(t *testing.T) {
	id := &ComputeGlobalVMExtensionPolicyIdentity{
		Project:                 "my-project",
		Globalvmextensionpolicy: "my-policy",
	}
	want := "projects/my-project/global"
	got := id.ParentString()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ParentString() mismatch (-want +got):\n%s", diff)
	}
}

func TestComputeGlobalVMExtensionPolicyIdentity_GetIdentity(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(GroupVersion, &ComputeGlobalVMExtensionPolicy{})
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	tests := []struct {
		name    string
		obj     *ComputeGlobalVMExtensionPolicy
		want    *ComputeGlobalVMExtensionPolicyIdentity
		wantErr bool
	}{
		{
			name: "from spec with resourceID",
			obj: &ComputeGlobalVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeGlobalVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					ResourceID: ptrTo("custom-policy"),
				},
			},
			want: &ComputeGlobalVMExtensionPolicyIdentity{
				Project:                 "my-project",
				Globalvmextensionpolicy: "custom-policy",
			},
		},
		{
			name: "from spec defaulting to metadata.name",
			obj: &ComputeGlobalVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeGlobalVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
				},
			},
			want: &ComputeGlobalVMExtensionPolicyIdentity{
				Project:                 "my-project",
				Globalvmextensionpolicy: "k8s-policy",
			},
		},
		{
			name: "matching status.externalRef",
			obj: &ComputeGlobalVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeGlobalVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
				},
				Status: ComputeGlobalVMExtensionPolicyStatus{
					ExternalRef: ptrTo("projects/my-project/global/vmExtensionPolicies/k8s-policy"),
				},
			},
			want: &ComputeGlobalVMExtensionPolicyIdentity{
				Project:                 "my-project",
				Globalvmextensionpolicy: "k8s-policy",
			},
		},
		{
			name: "mismatched status.externalRef",
			obj: &ComputeGlobalVMExtensionPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy",
					Namespace: "default",
				},
				Spec: ComputeGlobalVMExtensionPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
				},
				Status: ComputeGlobalVMExtensionPolicyStatus{
					ExternalRef: ptrTo("projects/other-project/global/vmExtensionPolicies/k8s-policy"),
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

func TestComputeGlobalVMExtensionPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/global/vmExtensionPolicies/my-policy",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/global/vmExtensionPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "missing global segment",
			ref:     "projects/my-project/vmExtensionPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "missing policy id",
			ref:     "projects/my-project/global/vmExtensionPolicies",
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
			r := &ComputeGlobalVMExtensionPolicyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeGlobalVMExtensionPolicyRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeGlobalVMExtensionPolicyRef_Normalize(t *testing.T) {
	testCases := []struct {
		name         string
		ref          *ComputeGlobalVMExtensionPolicyRef
		objects      []runtime.Object
		wantExternal string
		wantErr      string
	}{
		{
			name: "external with valid format",
			ref: &ComputeGlobalVMExtensionPolicyRef{
				External: "projects/test-project/global/vmExtensionPolicies/test-policy",
			},
			wantExternal: "projects/test-project/global/vmExtensionPolicies/test-policy",
		},
		{
			name: "external with invalid format",
			ref: &ComputeGlobalVMExtensionPolicyRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeGlobalVMExtensionPolicy external="invalid-format" was not known (use projects/{project}/global/vmExtensionPolicies/{globalvmextensionpolicy})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeGlobalVMExtensionPolicyRef{
				Name:      "test-policy",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeGlobalVMExtensionPolicy",
						"metadata": map[string]interface{}{
							"name":      "test-policy",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/global/vmExtensionPolicies/test-policy",
						},
					},
				},
			},
			wantExternal: "projects/test-project/global/vmExtensionPolicies/test-policy",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeGlobalVMExtensionPolicyRef{
				Name:      "test-policy",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeGlobalVMExtensionPolicy",
						"metadata": map[string]interface{}{
							"name":      "test-policy",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeGlobalVMExtensionPolicy my-namespace/test-policy is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeGlobalVMExtensionPolicyRef{
				Name:      "test-policy",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeGlobalVMExtensionPolicy my-namespace/test-policy is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeGlobalVMExtensionPolicy{})
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
