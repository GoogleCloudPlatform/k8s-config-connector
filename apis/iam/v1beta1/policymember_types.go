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

package v1beta1

import (
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMPolicyMemberSpec defines the desired state of IAMPolicyMember
type IAMPolicyMemberSpec struct {
	// Immutable. Required. The GCP resource to set the IAM policy on (e.g.
	// organization, project...)
	ResourceReference ResourceReference `json:"resourceRef"`

	// Immutable. The IAM identity to be bound to the role. Exactly one of
	// 'member' or 'memberFrom' must be used.
	Member Member `json:"member,omitempty"`

	// Immutable. The IAM identity to be bound to the role. Exactly one of
	// 'member' or 'memberFrom' must be used, and only one subfield within
	// 'memberFrom' can be used.
	MemberFrom *MemberSource `json:"memberFrom,omitempty"`

	// Immutable. Required. The role for which the Member will be bound.
	// +kubebuilder:validation:Pattern=^((projects|organizations)/[^/]+/)?roles/[\w_\.]+$
	Role string `json:"role"`
	// Immutable. Optional. The condition under which the binding applies.
	Condition *IAMCondition `json:"condition,omitempty"`
}

// IAMPolicyMemberStatus defines the observed state of IAMPolicyMember
type IAMPolicyMemberStatus struct {
	// Conditions represent the latest available observations of the IAM
	// policy's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPolicyMember is the Schema for the iampolicies API
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].status",description="When 'True' the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].reason",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",type="date",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime"
// +kubebuilder:subresource:status
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:resource:categories=gcp,shortName=gcpiampolicymember;gcpiampolicymembers
type IAMPolicyMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPolicyMemberSpec   `json:"spec,omitempty"`
	Status IAMPolicyMemberStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPolicyMemberList contains a list of IAMPolicyMember
type IAMPolicyMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMPolicyMember `json:"items"`
}

const IAMPolicyMemberReconcileInterval = 10 * time.Minute

func init() {
	SchemeBuilder.Register(&IAMPolicyMember{}, &IAMPolicyMemberList{})
}
