// Copyright 2025 Google LLC
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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAMDenyPolicyGVK = GroupVersion.WithKind("IAMDenyPolicy")

// IAMDenyPolicySpec defines the desired state of IAMDenyPolicy
// +kcc:spec:proto=google.iam.v2.Policy
type IAMDenyPolicySpec struct {
	// The IAMDenyPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// // Immutable. The resource name of the `Policy`, which must be unique. Format:
	// //  `policies/{attachment_point}/denypolicies/{policy_id}`
	// //
	// //
	// //  The attachment point is identified by its URL-encoded full resource name,
	// //  which means that the forward-slash character, `/`, must be written as
	// //  `%2F`. For example,
	// //  `policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-deny-policy`.
	// //
	// //  For organizations and folders, use the numeric ID in the full resource
	// //  name. For projects, requests can use the alphanumeric or the numeric ID.
	// //  Responses always contain the numeric ID.
	// // +kcc:proto:field=google.iam.v2.Policy.name
	// Name *string `json:"name,omitempty"`

	// A user-specified description of the `Policy`. This value can be up to 63
	//  characters.
	// +kcc:proto:field=google.iam.v2.Policy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// // NOTYET: not yet supporting annotations
	// // A key-value map to store arbitrary metadata for the `Policy`. Keys
	// //  can be up to 63 characters. Values can be up to 255 characters.
	// // +kcc:proto:field=google.iam.v2.Policy.annotations
	// Annotations map[string]string `json:"annotations,omitempty"`

	// A list of rules that specify the behavior of the `Policy`. All of the rules
	//  should be of the `kind` specified in the `Policy`.
	// +kcc:proto:field=google.iam.v2.Policy.rules
	Rules []PolicyRule `json:"rules,omitempty"`

	// NOTYET: not really documented?
	// // Immutable. Specifies that this policy is managed by an authority and can only be
	// //  modified by that authority. Usage is restricted.
	// // +kcc:proto:field=google.iam.v2.Policy.managing_authority
	// ManagingAuthority *string `json:"managingAuthority,omitempty"`
}

// IAMDenyPolicyStatus defines the config connector machine state of IAMDenyPolicy
type IAMDenyPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IAMDenyPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *IAMDenyPolicyObservedState `json:"observedState,omitempty"`
}

// IAMDenyPolicyObservedState is the state of the IAMDenyPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.iam.v2.Policy
type IAMDenyPolicyObservedState struct {
	// NOTYET: no clear use case?
	// // Immutable. The globally unique ID of the `Policy`. Assigned automatically when the
	// //  `Policy` is created.
	// // +kcc:proto:field=google.iam.v2.Policy.uid
	// Uid *string `json:"uid,omitempty"`

	// NOTYET: no clear use case?
	// 	// An opaque tag that identifies the current version of the `Policy`. IAM uses
	// //  this value to help manage concurrent updates, so they do not cause one
	// //  update to be overwritten by another.
	// //
	// //  If this field is present in a [CreatePolicy][] request, the value is
	// //  ignored.
	// // +kcc:proto:field=google.iam.v2.Policy.etag
	// Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamdenypolicy;gcpiamdenypolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMDenyPolicy is the Schema for the IAMDenyPolicy API
// +k8s:openapi-gen=true
type IAMDenyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAMDenyPolicySpec   `json:"spec,omitempty"`
	Status IAMDenyPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAMDenyPolicyList contains a list of IAMDenyPolicy
type IAMDenyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMDenyPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMDenyPolicy{}, &IAMDenyPolicyList{})
}
