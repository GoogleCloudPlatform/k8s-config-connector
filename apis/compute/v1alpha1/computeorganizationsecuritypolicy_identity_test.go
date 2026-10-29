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
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeOrganizationSecurityPolicyIdentity_Conformance(t *testing.T) {
	identity.AssertConformance(t, identity.ConformanceCase[ComputeOrganizationSecurityPolicyIdentity]{
		SampleExternal:        "locations/global/securityPolicies/1234567890",
		WantHost:              "compute.googleapis.com",
		WantParent:            "locations/global",
		ExpectServerGenerated: true,
		WantIdentity: &ComputeOrganizationSecurityPolicyIdentity{
			SecurityPolicy: "1234567890",
		},
		Ref: &ComputeOrganizationSecurityPolicyRef{},
	})
}

func TestComputeOrganizationSecurityPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeOrganizationSecurityPolicyIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "locations/global/securityPolicies/1234567890",
			want: &ComputeOrganizationSecurityPolicyIdentity{
				SecurityPolicy: "1234567890",
			},
		},
		{
			name: "valid full URL with compute v1 prefix",
			ref:  "https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890",
			want: &ComputeOrganizationSecurityPolicyIdentity{
				SecurityPolicy: "1234567890",
			},
		},
		{
			name: "valid full URL without compute prefix",
			ref:  "https://compute.googleapis.com/locations/global/securityPolicies/1234567890",
			want: &ComputeOrganizationSecurityPolicyIdentity{
				SecurityPolicy: "1234567890",
			},
		},
		{
			name:    "invalid format",
			ref:     "invalid/format",
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
			i := &ComputeOrganizationSecurityPolicyIdentity{}
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

func TestComputeOrganizationSecurityPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid relative reference",
			ref:     "locations/global/securityPolicies/1234567890",
			wantErr: false,
		},
		{
			name:    "valid full URL",
			ref:     "https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890",
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "projects/my-project/global/securityPolicies/1234567890",
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
			r := &ComputeOrganizationSecurityPolicyRef{}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeOrganizationSecurityPolicy_GetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = SchemeBuilder.AddToScheme(scheme)
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name         string
		obj          *ComputeOrganizationSecurityPolicy
		wantIdentity *ComputeOrganizationSecurityPolicyIdentity
		wantErr      bool
	}{
		{
			name: "spec resourceID set without status",
			obj: &ComputeOrganizationSecurityPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-policy",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicySpec{
					ResourceID: common.LazyPtr("1234567890"),
					Parent:     "organizations/123456",
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyIdentity{
				SecurityPolicy: "1234567890",
			},
		},
		{
			name: "spec resourceID unset, populated from status.policyId",
			obj: &ComputeOrganizationSecurityPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-policy",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicySpec{
					Parent: "organizations/123456",
				},
				Status: ComputeOrganizationSecurityPolicyStatus{
					PolicyId: common.LazyPtr("9876543210"),
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyIdentity{
				SecurityPolicy: "9876543210",
			},
		},
		{
			name: "spec resourceID matches status.policyId",
			obj: &ComputeOrganizationSecurityPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-policy",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicySpec{
					ResourceID: common.LazyPtr("1234567890"),
					Parent:     "organizations/123456",
				},
				Status: ComputeOrganizationSecurityPolicyStatus{
					PolicyId: common.LazyPtr("1234567890"),
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyIdentity{
				SecurityPolicy: "1234567890",
			},
		},
		{
			name: "spec resourceID conflicts with status.policyId",
			obj: &ComputeOrganizationSecurityPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-policy",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicySpec{
					ResourceID: common.LazyPtr("1234567890"),
					Parent:     "organizations/123456",
				},
				Status: ComputeOrganizationSecurityPolicyStatus{
					PolicyId: common.LazyPtr("9876543210"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.obj.GetIdentity(context.Background(), client)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				policyID, ok := id.(*ComputeOrganizationSecurityPolicyIdentity)
				if !ok {
					t.Fatalf("expected *ComputeOrganizationSecurityPolicyIdentity, got %T", id)
				}
				if diff := cmp.Diff(tt.wantIdentity, policyID); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeOrganizationSecurityPolicyRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = SchemeBuilder.AddToScheme(scheme)

	uReady := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
			"kind":       "ComputeOrganizationSecurityPolicy",
			"metadata": map[string]interface{}{
				"name":      "k8s-policy",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"parent": "organizations/123456",
			},
			"status": map[string]interface{}{
				"policyId": "999888777",
			},
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(uReady).Build()

	tests := []struct {
		name         string
		ref          *ComputeOrganizationSecurityPolicyRef
		wantExternal string
		wantErr      bool
	}{
		{
			name: "external already set with compute prefix",
			ref: &ComputeOrganizationSecurityPolicyRef{
				External: "https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890",
			},
			wantExternal: "locations/global/securityPolicies/1234567890",
		},
		{
			name: "resolve from name and status.policyId fallback",
			ref: &ComputeOrganizationSecurityPolicyRef{
				Name:      "k8s-policy",
				Namespace: "default",
			},
			wantExternal: "locations/global/securityPolicies/999888777",
		},
		{
			name: "referenced resource not found",
			ref: &ComputeOrganizationSecurityPolicyRef{
				Name:      "non-existent",
				Namespace: "default",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), client, "default")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.ref.External != tt.wantExternal {
					t.Errorf("Normalize() external mismatch: got %q, want %q", tt.ref.External, tt.wantExternal)
				}
			}
		})
	}
}
