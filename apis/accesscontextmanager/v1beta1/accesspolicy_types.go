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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AccessContextManagerAccessPolicyGVK = GroupVersion.WithKind("AccessContextManagerAccessPolicy")

// AccessContextManagerAccessPolicySpec defines the desired state of AccessContextManagerAccessPolicy
// +kcc:spec:proto=google.identity.accesscontextmanager.v1.AccessPolicy
type AccessContextManagerAccessPolicySpec struct {
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Human readable title. Does not affect behavior.
	// +required
	Title *string `json:"title"`
}

// AccessContextManagerAccessPolicyStatus defines the config connector machine state of AccessContextManagerAccessPolicy
type AccessContextManagerAccessPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Output only. Time the AccessPolicy was created in UTC.
	CreateTime string `json:"createTime,omitempty"`

	// Resource name of the AccessPolicy. Format: {policy_id}.
	Name string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. Time the AccessPolicy was updated in UTC.
	UpdateTime string `json:"updateTime,omitempty"`
}

// AccessContextManagerAccessPolicyObservedState is the state of the AccessContextManagerAccessPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.identity.accesscontextmanager.v1.AccessPolicy
type AccessContextManagerAccessPolicyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanageraccesspolicy;gcpaccesscontextmanageraccesspolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AccessContextManagerAccessPolicy is the Schema for the AccessContextManagerAccessPolicy API
// As per https://cloud.google.com/config-connector/docs/reference/resource-docs/accesscontextmanager/accesscontextmanageraccesspolicy#annotations
// the parent is organization which is stored in the cnrm.cloud.google.com/organization-id annotation.
// +k8s:openapi-gen=true
type AccessContextManagerAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AccessContextManagerAccessPolicySpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AccessContextManagerAccessPolicyList contains a list of AccessContextManagerAccessPolicy
type AccessContextManagerAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessPolicy{}, &AccessContextManagerAccessPolicyList{})
}
