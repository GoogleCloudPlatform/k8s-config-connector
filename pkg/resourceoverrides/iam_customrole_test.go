// Copyright 2024 Google LLC
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

package resourceoverrides_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestIAMCustomRoleConfigValidate(t *testing.T) {
	tests := []struct {
		name        string
		resourceID  string
		expectError bool
	}{
		{
			name:        "valid resourceID with underscores and dots",
			resourceID:  "my_custom.role_123",
			expectError: false,
		},
		{
			name:        "unset resourceID is allowed in admission (defaulted later)",
			resourceID:  "",
			expectError: false,
		},
		{
			name:        "invalid resourceID containing hyphens",
			resourceID:  "my-custom-role",
			expectError: true,
		},
		{
			name:        "invalid resourceID too short",
			resourceID:  "ab",
			expectError: true,
		},
		{
			name:        "invalid resourceID containing special symbols",
			resourceID:  "my@custom#role",
			expectError: true,
		},
	}

	handler := resourceoverrides.NewResourceOverridesHandler()
	handler.Register(resourceoverrides.GetIAMCustomRoleResourceOverrides())

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
					"kind":       "IAMCustomRole",
					"metadata": map[string]interface{}{
						"name": "sample-role",
					},
					"spec": map[string]interface{}{},
				},
			}
			if tc.resourceID != "" {
				_ = unstructured.SetNestedField(u.Object, tc.resourceID, "spec", "resourceID")
			}

			err := handler.ConfigValidate(u)
			if tc.expectError && err == nil {
				t.Fatalf("expected error for resourceID %q, got nil", tc.resourceID)
			}
			if !tc.expectError && err != nil {
				t.Fatalf("unexpected error for resourceID %q: %v", tc.resourceID, err)
			}
		})
	}
}

func TestIAMCustomRolePreActuationTransform(t *testing.T) {
	tests := []struct {
		name               string
		resourceName       string
		initialResourceID  string
		expectedResourceID string
		expectError        bool
	}{
		{
			name:               "kebab-case name defaults to snake_case resourceID",
			resourceName:       "audit-iamcustomrole",
			initialResourceID:  "",
			expectedResourceID: "audit_iamcustomrole",
			expectError:        false,
		},
		{
			name:               "already snake_case name preserved",
			resourceName:       "audit_role_1",
			initialResourceID:  "",
			expectedResourceID: "audit_role_1",
			expectError:        false,
		},
		{
			name:               "explicit spec.resourceID preserved",
			resourceName:       "some-other-name",
			initialResourceID:  "explicit_role_id",
			expectedResourceID: "explicit_role_id",
			expectError:        false,
		},
		{
			name:               "name too short to be valid role ID",
			resourceName:       "ab",
			initialResourceID:  "",
			expectedResourceID: "",
			expectError:        true,
		},
	}

	handler := resourceoverrides.NewResourceOverridesHandler()
	handler.Register(resourceoverrides.GetIAMCustomRoleResourceOverrides())

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "iam.cnrm.cloud.google.com/v1beta1",
					Kind:       "IAMCustomRole",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      tc.resourceName,
					Namespace: "default",
				},
				Spec: map[string]interface{}{},
			}
			if tc.initialResourceID != "" {
				res.Spec["resourceID"] = tc.initialResourceID
			}

			err := handler.PreActuationTransform(res)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			actualID, found := res.Spec["resourceID"]
			if !found {
				t.Fatalf("expected spec.resourceID to be set, but was missing")
			}
			if actualID != tc.expectedResourceID {
				t.Fatalf("expected spec.resourceID %q, got %q", tc.expectedResourceID, actualID)
			}
		})
	}
}

func TestIAMCustomRolePreTerraformExport(t *testing.T) {
	handler := resourceoverrides.NewResourceOverridesHandler()
	handler.Register(resourceoverrides.GetIAMCustomRoleResourceOverrides())

	// Test project scoped
	opProject := &operations.TerraformExport{
		TerraformState: &terraform.InstanceState{
			Attributes: map[string]string{
				"project": "sample-project",
			},
		},
		TerraformInfo: &terraform.InstanceInfo{},
	}
	err := handler.PreTerraformExport(context.Background(), schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMCustomRole"}, opProject)
	if err != nil {
		t.Fatalf("unexpected error for project scoped: %v", err)
	}
	if opProject.TerraformInfo.Type != "google_project_iam_custom_role" {
		t.Fatalf("expected TerraformInfo.Type 'google_project_iam_custom_role', got %q", opProject.TerraformInfo.Type)
	}

	// Test org scoped
	opOrg := &operations.TerraformExport{
		TerraformState: &terraform.InstanceState{
			Attributes: map[string]string{
				"org_id": "1234567890",
			},
		},
		TerraformInfo: &terraform.InstanceInfo{},
	}
	err = handler.PreTerraformExport(context.Background(), schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMCustomRole"}, opOrg)
	if err != nil {
		t.Fatalf("unexpected error for org scoped: %v", err)
	}
	if opOrg.TerraformInfo.Type != "google_organization_iam_custom_role" {
		t.Fatalf("expected TerraformInfo.Type 'google_organization_iam_custom_role', got %q", opOrg.TerraformInfo.Type)
	}
}
