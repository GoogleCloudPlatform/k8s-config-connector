// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for IAMAccessPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/iam/v3/access_policy_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/iam/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/iam/docs/reference/rest/v3/organizations.locations.accessPolicies

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAMAccessPolicyGVK = GroupVersion.WithKind("IAMAccessPolicy")

// IAMAccessPolicySpec defines the desired state of IAMAccessPolicy
// +kcc:spec:proto=google.iam.v3.AccessPolicy
// +kcc:required-from-proto
type IAMAccessPolicySpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The IAMAccessPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The etag for the access policy.
	//  If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.iam.v3.AccessPolicy.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The description of the access policy. Must be less than
	//  or equal to 63 characters.
	// +kcc:proto:field=google.iam.v3.AccessPolicy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined annotations. See
	//  https://google.aip.dev/148#annotations for more details such as format and
	//  size limitations
	// +kcc:proto:field=google.iam.v3.AccessPolicy.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. The details for the access policy.
	// +kcc:proto:field=google.iam.v3.AccessPolicy.details
	Details *AccessPolicyDetails `json:"details,omitempty"`
}

// IAMAccessPolicyStatus defines the config connector machine state of IAMAccessPolicy
type IAMAccessPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IAMAccessPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *IAMAccessPolicyObservedState `json:"observedState,omitempty"`
}

// IAMAccessPolicyObservedState is the state of the IAMAccessPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.iam.v3.AccessPolicy
type IAMAccessPolicyObservedState struct {
	// Output only. The globally unique ID of the access policy.
	// +kcc:proto:field=google.iam.v3.AccessPolicy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the access policy was created.
	// +kcc:proto:field=google.iam.v3.AccessPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the access policy was most recently
	//  updated.
	// +kcc:proto:field=google.iam.v3.AccessPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamaccesspolicy;gcpiamaccesspolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMAccessPolicy is the Schema for the IAMAccessPolicy API
// +k8s:openapi-gen=true
type IAMAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAMAccessPolicySpec   `json:"spec,omitempty"`
	Status IAMAccessPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAMAccessPolicyList contains a list of IAMAccessPolicy
type IAMAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMAccessPolicy{}, &IAMAccessPolicyList{})
}
