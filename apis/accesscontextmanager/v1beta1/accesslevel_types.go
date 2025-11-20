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

var AccessContextManagerAccessLevelGVK = GroupVersion.WithKind("AccessContextManagerAccessLevel")

// AccessContextManagerAccessLevelSpec defines the desired state of AccessContextManagerAccessLevel
// +kcc:spec:proto=google.identity.accesscontextmanager.v1.AccessLevel
type AccessContextManagerAccessLevelSpec struct {
	// The AccessPolicy that this resource belongs to.
	// +required
	AccessPolicyRef *AccessPolicyRef `json:"accessPolicyRef,omitempty"`

	// The AccessContextManagerAccessLevel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Human readable title. Must be unique within the Policy.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.title
	Title *string `json:"title,omitempty"`

	// Description of the `AccessLevel` and its use. Does not affect behavior.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.description
	Description *string `json:"description,omitempty"`

	// A `BasicLevel` composed of `Conditions`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.basic
	Basic *BasicLevel `json:"basic,omitempty"`

	// A `CustomLevel` written in the Common Expression Language.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.custom
	Custom *CustomLevel `json:"custom,omitempty"`
}

// AccessContextManagerAccessLevelStatus defines the config connector machine state of AccessContextManagerAccessLevel
type AccessContextManagerAccessLevelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AccessContextManagerAccessLevel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AccessContextManagerAccessLevelObservedState `json:"observedState,omitempty"`

	// Output only. Time the `AccessLevel` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `AccessLevel` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessLevel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// AccessContextManagerAccessLevelObservedState is the state of the AccessContextManagerAccessLevel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.identity.accesscontextmanager.v1.AccessLevel
type AccessContextManagerAccessLevelObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanageraccesslevel;gcpaccesscontextmanageraccesslevels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AccessContextManagerAccessLevel is the Schema for the AccessContextManagerAccessLevel API
// +k8s:openapi-gen=true
type AccessContextManagerAccessLevel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AccessContextManagerAccessLevelSpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessLevelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AccessContextManagerAccessLevelList contains a list of AccessContextManagerAccessLevel
type AccessContextManagerAccessLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerAccessLevel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessLevel{}, &AccessContextManagerAccessLevelList{})
}
