// Copyright 2022 Google LLC
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

package webhook

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestValidateIAMPolicy(t *testing.T) {
	tests := []struct {
		name                 string
		policy               *v1beta1.IAMPolicy
		refResourceRCs       []*v1alpha1.ResourceConfig
		isDCLResource        bool
		expectedAllowedValue bool
	}{
		{
			name:                 "IAMPolicy without conditions, with reference that doesn't support conditions",
			policy:               newTestPolicy(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicy without conditions, with reference that supports conditions",
			policy:               newTestPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicy with conditions, with reference that doesn't support conditions",
			policy:               newTestPolicyWithIAMConditions(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: false,
		},
		{
			name:                 "IAMPolicy with conditions, with reference that supports conditions",
			policy:               newTestPolicyWithIAMConditions(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicy without audit configs, with reference that doesn't support audit configs",
			policy:               newTestPolicy(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicy without audit configs, with reference that supports audit configs",
			policy:               newTestPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportAuditConfigs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicy with audit configs, with reference that doesn't support audit configs",
			policy:               newTestPolicyWithAuditConfigs(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: false,
		},
		{
			name:                 "IAMPolicy with audit configs, with reference that supports audit configs",
			policy:               newTestPolicyWithAuditConfigs(),
			refResourceRCs:       newTestResourceRCsThatSupportAuditConfigs(),
			expectedAllowedValue: true,
		},
		// TODO(kcc-eng): Remove the four tests below when we drop support for headless IAM.
		{
			name:                 "Headless Project IAMPolicy without conditions",
			policy:               newTestHeadlessProjectPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "Headless Project IAMPolicy with conditions",
			policy:               newTestHeadlessProjectPolicyWithIAMConditions(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "Headless Project IAMPolicy without audit configs",
			policy:               newTestHeadlessProjectPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportAuditConfigs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "Headless Project IAMPolicy with audit configs",
			policy:               newTestHeadlessProjectPolicyWithAuditConfigs(),
			refResourceRCs:       newTestResourceRCsThatSupportAuditConfigs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicy with external-only reference",
			policy:               newTestExternalOnlyRefPolicy(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name: "DCL-based IAMPolicy",
			policy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1beta1.IAMPolicyGVK.Kind,
					APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind:       "DataprocCluster",
						Namespace:  "my-namespace",
						Name:       "my-resource-name",
						APIVersion: "dataproc/v1beta1",
					},
				},
			},
			isDCLResource:        true,
			expectedAllowedValue: true,
		},
		{
			name:                 "DCL-based IAMPolicy with audit configs",
			policy:               newTestHeadlessProjectPolicyWithAuditConfigs(),
			isDCLResource:        true,
			expectedAllowedValue: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclSchemaLoader, err := dclschemaloader.New()
			if err != nil {
				t.Fatalf("error creating a new DCL schema loader: %v", err)
			}
			a := iamValidatorHandler{
				smLoader:              testservicemappingloader.New(t),
				serviceMetadataLoader: dclmetadata.New(),
				schemaLoader:          dclSchemaLoader,
			}
			var response admission.Response
			if tc.isDCLResource {
				response = a.dclValidateIAMPolicy(tc.policy)
			} else {
				response = a.tfValidateIAMPolicy(tc.policy, tc.refResourceRCs)
			}
			if response.Allowed != tc.expectedAllowedValue {
				t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, tc.expectedAllowedValue)
			}
		})
	}
}

func TestValidateIAMPartialPolicy(t *testing.T) {
	tests := []struct {
		name                 string
		partialPolicy        *v1beta1.IAMPartialPolicy
		refResourceRCs       []*v1alpha1.ResourceConfig
		isDCLResource        bool
		expectedAllowedValue bool
	}{
		{
			name:                 "IAMPartialPolicy without conditions, with reference that doesn't support conditions",
			partialPolicy:        newTestPartialPolicy(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPartialPolicy without conditions, with reference that supports conditions",
			partialPolicy:        newTestPartialPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPartialPolicy with conditions, with reference that doesn't support conditions",
			partialPolicy:        newTestPartialPolicyWithIAMConditions(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: false,
		},
		{
			name:                 "IAMPartialPolicy with conditions, with reference that supports conditions",
			partialPolicy:        newTestPartialPolicyWithIAMConditions(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPartialPolicy without audit configs, with reference that supports audit configs",
			partialPolicy:        newTestPartialPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportAuditConfigs(),
			expectedAllowedValue: true,
		},
		// TODO(kcc-eng): Remove the four tests below when we drop support for headless IAM.
		{
			name:                 "Headless Project IAMPartialPolicy without conditions",
			partialPolicy:        newTestHeadlessProjectPartialPolicy(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "Headless Project IAMPartialPolicy with conditions",
			partialPolicy:        newTestHeadlessProjectPartialPolicyWithIAMConditions(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPartialPolicy with external-only reference",
			partialPolicy:        newTestExternalOnlyRefPartialPolicy(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name: "DCL-based IAMPartialPolicy",
			partialPolicy: &v1beta1.IAMPartialPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1beta1.IAMPolicyGVK.Kind,
					APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
				},
				Spec: v1beta1.IAMPartialPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind:       "DataprocCluster",
						Namespace:  "my-namespace",
						Name:       "my-resource-name",
						APIVersion: "dataproc/v1beta1",
					},
				},
			},
			isDCLResource:        true,
			expectedAllowedValue: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclSchemaLoader, err := dclschemaloader.New()
			if err != nil {
				t.Fatalf("error creating a new DCL schema loader: %v", err)
			}
			a := iamValidatorHandler{
				smLoader:              testservicemappingloader.New(t),
				serviceMetadataLoader: dclmetadata.New(),
				schemaLoader:          dclSchemaLoader,
			}
			var response admission.Response
			if tc.isDCLResource {
				response = a.dclValidateIAMPartialPolicy(tc.partialPolicy)
			} else {
				response = a.tfValidateIAMPartialPolicy(tc.partialPolicy, tc.refResourceRCs)
			}
			if response.Allowed != tc.expectedAllowedValue {
				t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, tc.expectedAllowedValue)
			}
		})
	}
}

func TestValidateIAMPolicyMember(t *testing.T) {
	tests := []struct {
		name                 string
		policyMember         *v1beta1.IAMPolicyMember
		refResourceRCs       []*v1alpha1.ResourceConfig
		isDCLResource        bool
		expectedAllowedValue bool
	}{
		{
			name:                 "IAMPolicyMember without conditions, with reference that doesn't support conditions",
			policyMember:         newTestPolicyMember(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicyMember without conditions, with reference that supports conditions",
			policyMember:         newTestPolicyMember(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicyMember with conditions, with reference that doesn't support conditions",
			policyMember:         newTestPolicyMemberWithIAMConditions(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: false,
		},
		{
			name:                 "IAMPolicyMember with conditions, with reference that supports conditions",
			policyMember:         newTestPolicyMemberWithIAMConditions(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		// TODO(kcc-eng): Remove the two tests below when we drop support for headless IAM.
		{
			name:                 "Headless Project IAMPolicyMember without conditions",
			policyMember:         newTestHeadlessProjectPolicyMember(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "Headless Project IAMPolicyMember with conditions",
			policyMember:         newTestHeadlessProjectPolicyMemberWithIAMConditions(),
			refResourceRCs:       newTestResourceRCsThatSupportConditions(),
			expectedAllowedValue: true,
		},
		{
			name:                 "IAMPolicyMember with external-only reference",
			policyMember:         newTestExternalOnlyRefPolicyMember(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: true,
		},
		{
			name: "DCL-based IAMPolicyMember",
			policyMember: &v1beta1.IAMPolicyMember{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1beta1.IAMPolicyGVK.Kind,
					APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind:       "DataprocCluster",
						Namespace:  "my-namespace",
						Name:       "my-resource-name",
						APIVersion: "dataproc/v1beta1",
					},
					Role: "role",
				},
			},
			isDCLResource:        true,
			expectedAllowedValue: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclSchemaLoader, err := dclschemaloader.New()
			if err != nil {
				t.Fatalf("error creating a new DCL schema loader: %v", err)
			}
			a := iamValidatorHandler{
				smLoader:              testservicemappingloader.New(t),
				serviceMetadataLoader: dclmetadata.New(),
				schemaLoader:          dclSchemaLoader,
			}
			response := a.tfValidateIAMPolicyMember(tc.policyMember, tc.refResourceRCs)
			if response.Allowed != tc.expectedAllowedValue {
				t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, tc.expectedAllowedValue)
			}
		})
	}
}

func TestValidateIAMAuditConfig(t *testing.T) {
	tests := []struct {
		name                 string
		auditConfig          *v1beta1.IAMAuditConfig
		refResourceRCs       []*v1alpha1.ResourceConfig
		expectedAllowedValue bool
	}{
		{
			name:                 "IAMAuditConfig with reference that doesn't support audit configs",
			auditConfig:          newTestAuditConfig(),
			refResourceRCs:       newTestResourceRCs(),
			expectedAllowedValue: false,
		},
		{
			name:                 "IAMAuditConfig with reference that supports audit configs",
			auditConfig:          newTestAuditConfig(),
			refResourceRCs:       newTestResourceRCsThatSupportAuditConfigs(),
			expectedAllowedValue: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			response := validateIAMAuditConfig(tc.auditConfig, tc.refResourceRCs)
			if response.Allowed != tc.expectedAllowedValue {
				t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, tc.expectedAllowedValue)
			}
		})
	}
}

func newTestPolicy() *v1beta1.IAMPolicy {
	return &v1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMPolicyGVK.Kind,
			APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMPolicySpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-resource-name",
				APIVersion: "my-api-version",
			},
			Bindings: []v1beta1.IAMPolicyBinding{
				{
					Members: []v1beta1.Member{"member"},
					Role:    "role",
				},
			},
		},
	}
}

func newTestPartialPolicy() *v1beta1.IAMPartialPolicy {
	return &v1beta1.IAMPartialPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMPolicyGVK.Kind,
			APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMPartialPolicySpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-resource-name",
				APIVersion: "my-api-version",
			},
			Bindings: []v1beta1.IAMPartialPolicyBinding{
				{
					Members: []v1beta1.IAMPartialPolicyMember{
						{
							Member: "member",
						},
					},
					Role: "role",
				},
			},
		},
	}
}

func newTestPolicyMember() *v1beta1.IAMPolicyMember {
	return &v1beta1.IAMPolicyMember{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMPolicyMemberGVK.Kind,
			APIVersion: v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMPolicyMemberSpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-resource-name",
				APIVersion: "my-api-version",
			},
			Member: "member",
			Role:   "role",
		},
	}
}

func newTestAuditConfig() *v1beta1.IAMAuditConfig {
	return &v1beta1.IAMAuditConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMAuditConfigGVK.Kind,
			APIVersion: v1beta1.IAMAuditConfigGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMAuditConfigSpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-resource-name",
				APIVersion: "my-api-version",
			},
			Service: "service",
			AuditLogConfigs: []v1beta1.AuditLogConfig{
				{
					LogType:         "log-type",
					ExemptedMembers: []v1beta1.Member{"member"},
				},
			},
		},
	}
}

func newTestHeadlessProjectPolicy() *v1beta1.IAMPolicy {
	p := newTestPolicy()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestHeadlessProjectPolicyWithIAMConditions() *v1beta1.IAMPolicy {
	p := newTestPolicyWithIAMConditions()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestHeadlessProjectPartialPolicy() *v1beta1.IAMPartialPolicy {
	p := newTestPartialPolicy()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestHeadlessProjectPartialPolicyWithIAMConditions() *v1beta1.IAMPartialPolicy {
	p := newTestPartialPolicyWithIAMConditions()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestHeadlessProjectPolicyMember() *v1beta1.IAMPolicyMember {
	p := newTestPolicyMember()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestHeadlessProjectPolicyMemberWithIAMConditions() *v1beta1.IAMPolicyMember {
	p := newTestPolicyMemberWithIAMConditions()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestPolicyWithIAMConditions() *v1beta1.IAMPolicy {
	policy := newTestPolicy()
	bindings := make([]v1beta1.IAMPolicyBinding, 0)
	for _, binding := range policy.Spec.Bindings {
		binding.Condition = newTestIAMCondition()
		bindings = append(bindings, binding)
	}
	policy.Spec.Bindings = bindings
	return policy
}

func newTestPartialPolicyWithIAMConditions() *v1beta1.IAMPartialPolicy {
	pPolicy := newTestPartialPolicy()
	bindings := make([]v1beta1.IAMPartialPolicyBinding, 0)
	for _, binding := range pPolicy.Spec.Bindings {
		binding.Condition = newTestIAMCondition()
		bindings = append(bindings, binding)
	}
	pPolicy.Spec.Bindings = bindings
	return pPolicy
}

func newTestPolicyMemberWithIAMConditions() *v1beta1.IAMPolicyMember {
	policyMember := newTestPolicyMember()
	policyMember.Spec.Condition = newTestIAMCondition()
	return policyMember
}

func newTestIAMCondition() *v1beta1.IAMCondition {
	return &v1beta1.IAMCondition{
		Title:       "test-iam-condition",
		Description: "Test IAM Condition",
		Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
	}
}

func newTestPolicyWithAuditConfigs() *v1beta1.IAMPolicy {
	policy := newTestPolicy()
	policy.Spec.AuditConfigs = []v1beta1.IAMPolicyAuditConfig{
		{
			Service: "service",
			AuditLogConfigs: []v1beta1.AuditLogConfig{
				{
					LogType:         "log-type",
					ExemptedMembers: []v1beta1.Member{"member"},
				},
			},
		},
	}
	return policy
}

func newTestHeadlessProjectPolicyWithAuditConfigs() *v1beta1.IAMPolicy {
	p := newTestPolicyWithAuditConfigs()
	p.Spec.ResourceReference = v1beta1.ResourceReference{
		Kind: kcciamclient.ProjectKind,
	}
	return p
}

func newTestExternalOnlyRefPolicy() *v1beta1.IAMPolicy {
	p := newTestPolicy()
	resourceRef := v1beta1.ResourceReference{}
	resourceRef.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Organization",
	})
	p.Spec.ResourceReference = resourceRef
	return p
}

func newTestExternalOnlyRefPartialPolicy() *v1beta1.IAMPartialPolicy {
	p := newTestPartialPolicy()
	resourceRef := v1beta1.ResourceReference{}
	resourceRef.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Organization",
	})
	p.Spec.ResourceReference = resourceRef
	return p
}

func newTestExternalOnlyRefPolicyMember() *v1beta1.IAMPolicyMember {
	policyMember := newTestPolicyMember()
	resourceRef := v1beta1.ResourceReference{}
	resourceRef.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Organization",
	})
	policyMember.Spec.ResourceReference = resourceRef
	return policyMember
}

func newTestResourceRCs() []*v1alpha1.ResourceConfig {
	return []*v1alpha1.ResourceConfig{{
		Name: "google_resource_type",
		Kind: "ResourceType",
		IAMConfig: v1alpha1.IAMConfig{
			PolicyName:       "google_resource_type_iam_policy",
			PolicyMemberName: "google_resource_type_iam_policy_member",
			ReferenceField: v1alpha1.IAMReferenceField{
				Name: "name",
				Type: v1alpha1.IAMReferenceTypeName,
			},
		},
	},
	}
}

func newTestResourceRCsThatSupportConditions() []*v1alpha1.ResourceConfig {
	rcs := newTestResourceRCs()
	for _, rc := range rcs {
		rc.IAMConfig.SupportsConditions = true
	}
	return rcs
}

func newTestResourceRCsThatSupportAuditConfigs() []*v1alpha1.ResourceConfig {
	rcs := newTestResourceRCs()
	for _, rc := range rcs {
		rc.IAMConfig.AuditConfigName = "google_resource_type_iam_audit_config"
	}
	return rcs
}
