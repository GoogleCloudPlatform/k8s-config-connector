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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNetworkServicesServiceLBPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *NetworkServicesServiceLBPolicyIdentity
		wantErr bool
	}{
		{
			name: "valid serviceLbPolicy reference",
			ref:  "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			want: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "my-project",
				Location:        "us-central1",
				ServiceLbPolicy: "my-policy",
			},
			wantErr: false,
		},
		{
			name: "valid serviceLbPolicy reference with host",
			ref:  "networkservices.googleapis.com/projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			want: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "my-project",
				Location:        "us-central1",
				ServiceLbPolicy: "my-policy",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing serviceLbPolicies)",
			ref:     "projects/my-project/locations/us-central1/my-policy",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/us-central1/serviceLbPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "empty reference",
			ref:     "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkServicesServiceLBPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesServiceLBPolicyIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("NetworkServicesServiceLBPolicyIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNetworkServicesServiceLBPolicy_GetIdentity(t *testing.T) {
	ctx := context.Background()

	projectObj := &unstructured.Unstructured{}
	projectObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	projectObj.SetName("referenced-project")
	projectObj.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(projectObj.Object, "resolved-project-id", "spec", "resourceID"); err != nil {
		t.Fatalf("failed to set nested field: %v", err)
	}

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	scheme.AddKnownTypeWithName(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ProjectList",
	}, &unstructured.UnstructuredList{})

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(projectObj).Build()

	tests := []struct {
		name         string
		obj          *NetworkServicesServiceLBPolicy
		wantIdentity *NetworkServicesServiceLBPolicyIdentity
		wantErr      bool
	}{
		{
			name: "resolution with explicit resourceID",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-policy-name",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location:   common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("custom-policy-id"),
				},
			},
			wantIdentity: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "my-project",
				Location:        "us-central1",
				ServiceLbPolicy: "custom-policy-id",
			},
			wantErr: false,
		},
		{
			name: "resolution defaulting resourceID to metadata.name",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-policy-name",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location: common.LazyPtr("global"),
				},
			},
			wantIdentity: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "my-project",
				Location:        "global",
				ServiceLbPolicy: "default-policy-name",
			},
			wantErr: false,
		},
		{
			name: "resolution with projectRef via reader",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						Name: "referenced-project",
					},
					Location: common.LazyPtr("us-central1"),
				},
			},
			wantIdentity: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "resolved-project-id",
				Location:        "us-central1",
				ServiceLbPolicy: "my-policy",
			},
			wantErr: false,
		},
		{
			name: "matching status.externalRef",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location: common.LazyPtr("us-central1"),
				},
				Status: NetworkServicesServiceLBPolicyStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/serviceLbPolicies/my-policy"),
				},
			},
			wantIdentity: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "my-project",
				Location:        "us-central1",
				ServiceLbPolicy: "my-policy",
			},
			wantErr: false,
		},
		{
			name: "mismatched status.externalRef returns error",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location: common.LazyPtr("us-central1"),
				},
				Status: NetworkServicesServiceLBPolicyStatus{
					ExternalRef: common.LazyPtr("projects/other-project/locations/us-central1/serviceLbPolicies/my-policy"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing location returns error",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing project returns error",
			obj: &NetworkServicesServiceLBPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "test-ns",
				},
				Spec: NetworkServicesServiceLBPolicySpec{
					Location: common.LazyPtr("us-central1"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				gotIdentity, ok := got.(*NetworkServicesServiceLBPolicyIdentity)
				if !ok {
					t.Fatalf("GetIdentity() did not return *NetworkServicesServiceLBPolicyIdentity")
				}
				if diff := cmp.Diff(tt.wantIdentity, gotIdentity); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNewNetworkServicesServiceLBPolicyIdentity(t *testing.T) {
	ctx := context.Background()

	t.Run("valid obj", func(t *testing.T) {
		obj := &NetworkServicesServiceLBPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-policy",
				Namespace: "test-ns",
			},
			Spec: NetworkServicesServiceLBPolicySpec{
				ProjectRef: &refsv1beta1.ProjectRef{
					External: "my-project",
				},
				Location: common.LazyPtr("us-central1"),
			},
		}

		got, err := NewNetworkServicesServiceLBPolicyIdentity(ctx, nil, obj)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := &NetworkServicesServiceLBPolicyIdentity{
			Project:         "my-project",
			Location:        "us-central1",
			ServiceLbPolicy: "my-policy",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("NewNetworkServicesServiceLBPolicyIdentity() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("mismatched status", func(t *testing.T) {
		obj := &NetworkServicesServiceLBPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-policy",
				Namespace: "test-ns",
			},
			Spec: NetworkServicesServiceLBPolicySpec{
				ProjectRef: &refsv1beta1.ProjectRef{
					External: "my-project",
				},
				Location: common.LazyPtr("us-central1"),
			},
			Status: NetworkServicesServiceLBPolicyStatus{
				ExternalRef: common.LazyPtr("projects/other-project/locations/us-central1/serviceLbPolicies/my-policy"),
			},
		}

		_, err := NewNetworkServicesServiceLBPolicyIdentity(ctx, nil, obj)
		if err == nil {
			t.Fatal("expected error for mismatched status, got nil")
		}
	})
}

func TestNetworkServicesServiceLBPolicyRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()

	tests := []struct {
		name             string
		ref              *NetworkServicesServiceLBPolicyRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesServiceLBPolicyRef{
				External: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			},
			wantExternal: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesServiceLBPolicyRef{
				Name:     "my-policy",
				External: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			},
			wantErr: true,
		},
		{
			name: "resolve from status.externalRef",
			ref: &NetworkServicesServiceLBPolicyRef{
				Name: "my-policy",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1alpha1",
						"kind":       "NetworkServicesServiceLBPolicy",
						"metadata": map[string]interface{}{
							"name":      "my-policy",
							"namespace": "test-ns",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			wantErr:      false,
		},
		{
			name: "missing reference",
			ref: &NetworkServicesServiceLBPolicyRef{
				Name: "non-existent",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build()
			err := tt.ref.Normalize(context.Background(), client, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesServiceLBPolicyRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("NetworkServicesServiceLBPolicyRef.Normalize() got = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}

func TestNetworkServicesServiceLBPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/serviceLbPolicies/my-policy",
			wantErr: true,
		},
		{
			name:    "missing policy id",
			ref:     "projects/my-project/locations/us-central1/serviceLbPolicies",
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
			r := &NetworkServicesServiceLBPolicyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesServiceLBPolicyRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetworkServicesServiceLBPolicyRef_ParseExternalToIdentity(t *testing.T) {
	r := &NetworkServicesServiceLBPolicyRef{
		External: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
	}
	id, err := r.ParseExternalToIdentity()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	policyID, ok := id.(*NetworkServicesServiceLBPolicyIdentity)
	if !ok {
		t.Fatalf("expected *NetworkServicesServiceLBPolicyIdentity, got %T", id)
	}
	want := &NetworkServicesServiceLBPolicyIdentity{
		Project:         "my-project",
		Location:        "us-central1",
		ServiceLbPolicy: "my-policy",
	}
	if diff := cmp.Diff(want, policyID); diff != "" {
		t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
	}
}

func TestNetworkServicesServiceLBPolicyIdentity_Interfaces(t *testing.T) {
	var _ identity.IdentityV2 = &NetworkServicesServiceLBPolicyIdentity{}
	var _ identity.Resource = &NetworkServicesServiceLBPolicy{}
}
