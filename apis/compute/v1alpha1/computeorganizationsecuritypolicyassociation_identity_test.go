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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeOrganizationSecurityPolicyAssociationIdentity_Conformance(t *testing.T) {
	identity.AssertConformance(t, identity.ConformanceCase[ComputeOrganizationSecurityPolicyAssociationIdentity]{
		SampleExternal: "locations/global/securityPolicies/1234567890/associations/my-association",
		WantHost:       "compute.googleapis.com",
		WantParent:     "locations/global/securityPolicies/1234567890",
		WantIdentity: &ComputeOrganizationSecurityPolicyAssociationIdentity{
			SecurityPolicy: "1234567890",
			Association:    "my-association",
		},
		Ref: &ComputeOrganizationSecurityPolicyAssociationRef{},
	})
}

func TestComputeOrganizationSecurityPolicyAssociationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeOrganizationSecurityPolicyAssociationIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "locations/global/securityPolicies/1234567890/associations/my-association",
			want: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "1234567890",
				Association:    "my-association",
			},
		},
		{
			name: "valid full URL with compute v1 prefix",
			ref:  "https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890/associations/my-association",
			want: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "1234567890",
				Association:    "my-association",
			},
		},
		{
			name: "valid full URL without compute prefix",
			ref:  "https://compute.googleapis.com/locations/global/securityPolicies/1234567890/associations/my-association",
			want: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "1234567890",
				Association:    "my-association",
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
			i := &ComputeOrganizationSecurityPolicyAssociationIdentity{}
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

func TestComputeOrganizationSecurityPolicyAssociationRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid relative reference",
			ref:     "locations/global/securityPolicies/1234567890/associations/my-association",
			wantErr: false,
		},
		{
			name:    "valid full URL",
			ref:     "https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890/associations/my-association",
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "locations/global/securityPolicies/1234567890",
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
			r := &ComputeOrganizationSecurityPolicyAssociationRef{}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeOrganizationSecurityPolicyAssociation_GetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = SchemeBuilder.AddToScheme(scheme)
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name         string
		obj          *ComputeOrganizationSecurityPolicyAssociation
		wantIdentity *ComputeOrganizationSecurityPolicyAssociationIdentity
		wantErr      bool
	}{
		{
			name: "spec resourceID set and plain policyID",
			obj: &ComputeOrganizationSecurityPolicyAssociation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "association-k8s-name",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicyAssociationSpec{
					AttachmentID: "organizations/123456789",
					PolicyID:     common.LazyPtr("1234567890"),
					ResourceID:   common.LazyPtr("custom-assoc-id"),
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "1234567890",
				Association:    "custom-assoc-id",
			},
		},
		{
			name: "spec resourceID unset, fallback to metadata.name and full url policyID",
			obj: &ComputeOrganizationSecurityPolicyAssociation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "association-k8s-name",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicyAssociationSpec{
					AttachmentID: "folders/987654321",
					PolicyID:     common.LazyPtr("locations/global/securityPolicies/9876543210"),
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "9876543210",
				Association:    "association-k8s-name",
			},
		},
		{
			name: "spec policyID with compute v1 prefix",
			obj: &ComputeOrganizationSecurityPolicyAssociation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "assoc-1",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicyAssociationSpec{
					AttachmentID: "organizations/123456789",
					PolicyID:     common.LazyPtr("https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/5555555555"),
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "5555555555",
				Association:    "assoc-1",
			},
		},
		{
			name: "spec with securityPolicyRef",
			obj: &ComputeOrganizationSecurityPolicyAssociation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "assoc-ref",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicyAssociationSpec{
					AttachmentID: "projects/123456789",
					SecurityPolicyRef: &ComputeOrganizationSecurityPolicyRef{
						External: "locations/global/securityPolicies/1234567890",
					},
				},
			},
			wantIdentity: &ComputeOrganizationSecurityPolicyAssociationIdentity{
				SecurityPolicy: "1234567890",
				Association:    "assoc-ref",
			},
		},
		{
			name: "empty policyId",
			obj: &ComputeOrganizationSecurityPolicyAssociation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "association-k8s-name",
					Namespace: "default",
				},
				Spec: ComputeOrganizationSecurityPolicyAssociationSpec{
					AttachmentID: "organizations/123456789",
					PolicyID:     common.LazyPtr(""),
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
				assocID, ok := id.(*ComputeOrganizationSecurityPolicyAssociationIdentity)
				if !ok {
					t.Fatalf("expected *ComputeOrganizationSecurityPolicyAssociationIdentity, got %T", id)
				}
				if diff := cmp.Diff(tt.wantIdentity, assocID); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeOrganizationSecurityPolicyAssociationRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = SchemeBuilder.AddToScheme(scheme)

	targetObj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
			"kind":       "ComputeOrganizationSecurityPolicyAssociation",
			"metadata": map[string]interface{}{
				"name":      "assoc-target",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"attachmentId": "organizations/123456",
				"policyId":     "1234567890",
				"resourceID":   "custom-assoc",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
					},
				},
			},
		},
	}

	unreadyObj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
			"kind":       "ComputeOrganizationSecurityPolicyAssociation",
			"metadata": map[string]interface{}{
				"name":      "assoc-unready",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"attachmentId": "organizations/123456",
				"policyId":     "1234567890",
				"resourceID":   "custom-assoc",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "False",
					},
				},
			},
		},
	}

	noConditionsObj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
			"kind":       "ComputeOrganizationSecurityPolicyAssociation",
			"metadata": map[string]interface{}{
				"name":      "assoc-no-conditions",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"attachmentId": "organizations/123456",
				"policyId":     "1234567890",
				"resourceID":   "custom-assoc",
			},
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(targetObj, unreadyObj, noConditionsObj).Build()

	tests := []struct {
		name         string
		ref          *ComputeOrganizationSecurityPolicyAssociationRef
		wantExternal string
		wantErr      bool
		errCheck     func(err error) bool
	}{
		{
			name: "external already set, normalized",
			ref: &ComputeOrganizationSecurityPolicyAssociationRef{
				External: "https://compute.googleapis.com/compute/v1/locations/global/securityPolicies/1234567890/associations/my-association",
			},
			wantExternal: "locations/global/securityPolicies/1234567890/associations/my-association",
		},
		{
			name: "resolve from name and fallback",
			ref: &ComputeOrganizationSecurityPolicyAssociationRef{
				Name:      "assoc-target",
				Namespace: "default",
			},
			wantExternal: "locations/global/securityPolicies/1234567890/associations/custom-assoc",
		},
		{
			name: "resolve from name fails when target has Ready: False condition",
			ref: &ComputeOrganizationSecurityPolicyAssociationRef{
				Name:      "assoc-unready",
				Namespace: "default",
			},
			wantErr: true,
			errCheck: func(err error) bool {
				return k8s.IsReferenceNotReadyError(err)
			},
		},
		{
			name: "resolve from name fails when target has no conditions",
			ref: &ComputeOrganizationSecurityPolicyAssociationRef{
				Name:      "assoc-no-conditions",
				Namespace: "default",
			},
			wantErr: true,
			errCheck: func(err error) bool {
				return k8s.IsReferenceNotReadyError(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), client, "default")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && tt.errCheck != nil {
				if !tt.errCheck(err) {
					t.Errorf("Normalize() unexpected error type: %v", err)
				}
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("Normalize() external = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
