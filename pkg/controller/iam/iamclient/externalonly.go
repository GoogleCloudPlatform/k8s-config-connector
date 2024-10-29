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

package iamclient

import (
	"fmt"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/externalonlygvks"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ExternalOnlyType is a KCC resource type that KCC does not support as a
// core resource, but does support referencing externally in IAM.
type ExternalOnlyType struct {
	// UnstructHandler is a function that fills in the external field information
	// from the given reference into the given unstructured object.
	UnstructHandler func(ref iamv1beta1.ResourceReference, u *unstructured.Unstructured) *unstructured.Unstructured

	// ResourceConfig is a skeleton resource config that includes IAM configuration
	// needed to map to the proper Terraform resource.
	ResourceConfig *corekccv1alpha1.ResourceConfig

	// ExternalFormat is the format the external field is expected to match. This
	// is used for documentation only. Ex. "{{org_id}}"
	ExternalFormat string
}

var ExternalOnlyTypes = map[schema.GroupVersionKind]ExternalOnlyType{
	externalonlygvks.OrganizationGVK: {
		UnstructHandler: func(ref iamv1beta1.ResourceReference, u *unstructured.Unstructured) *unstructured.Unstructured {
			u.Object["spec"] = map[string]interface{}{
				"org_id": ref.External,
			}
			return u
		},
		ResourceConfig: &corekccv1alpha1.ResourceConfig{
			IAMConfig: corekccv1alpha1.IAMConfig{
				PolicyName:       "google_organization_iam_policy",
				PolicyMemberName: "google_organization_iam_member",
				AuditConfigName:  "google_organization_iam_audit_config",
				ReferenceField: corekccv1alpha1.IAMReferenceField{
					Name: "org_id",
					Type: "id",
				},
				SupportsConditions: true,
			},
		},
		ExternalFormat: "{{org_id}}",
	},
	externalonlygvks.BillingAccountGVK: {
		UnstructHandler: func(ref iamv1beta1.ResourceReference, u *unstructured.Unstructured) *unstructured.Unstructured {
			u.Object["spec"] = map[string]interface{}{
				"billing_account_id": ref.External,
			}
			return u
		},
		ResourceConfig: &corekccv1alpha1.ResourceConfig{
			IAMConfig: corekccv1alpha1.IAMConfig{
				PolicyName:       "google_billing_account_iam_policy",
				PolicyMemberName: "google_billing_account_iam_member",
				ReferenceField: corekccv1alpha1.IAMReferenceField{
					Name: "billing_account_id",
					Type: "id",
				},
				SupportsConditions: true,
			},
		},
		ExternalFormat: "{{billing_account_id}}",
	},
}

func GetResourceConfigForExternalOnlyGVK(gvk schema.GroupVersionKind) (*corekccv1alpha1.ResourceConfig, error) {
	ext, ok := ExternalOnlyTypes[gvk]
	if !ok {
		return nil, fmt.Errorf("unsupported external-only reference of type %v", gvk)
	}
	return ext.ResourceConfig, nil
}

func unstructuredIAMSkeletonForExternalOnlyRef(resourceRef iamv1beta1.ResourceReference, u *unstructured.Unstructured) (
	*unstructured.Unstructured, error) {
	gvk := resourceRef.GroupVersionKind()
	ext, ok := ExternalOnlyTypes[gvk]
	if !ok {
		return nil, fmt.Errorf("unsupported external-only reference of type %v", gvk)
	}
	return ext.UnstructHandler(resourceRef, u), nil
}
