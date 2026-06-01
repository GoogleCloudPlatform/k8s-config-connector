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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AccessContextManagerAccessLevelConditionGVK = GroupVersion.WithKind("AccessContextManagerAccessLevelCondition")

// AccessContextManagerAccessLevelConditionSpec defines the desired state of AccessContextManagerAccessLevelCondition
// +kcc:spec:proto=google.identity.accesscontextmanager.v1.Condition
type AccessContextManagerAccessLevelConditionSpec struct {
	// The AccessLevel that this resource belongs to.
	// +required
	AccessLevelRef *AccessLevelRef `json:"accessLevelRef,omitempty"`

	// The AccessContextManagerAccessLevelCondition name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.device_policy
	DevicePolicy *DevicePolicy `json:"devicePolicy,omitempty"`

	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.ip_subnetworks
	IPSubnetworks []string `json:"ipSubnetworks,omitempty"`

	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.members
	Members []Member `json:"members,omitempty"`

	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.negate
	Negate *bool `json:"negate,omitempty"`

	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.regions
	Regions []string `json:"regions,omitempty"`

	// +kcc:proto:field=google.identity.accesscontextmanager.v1.Condition.required_access_levels
	RequiredAccessLevels []AccessLevelRef `json:"requiredAccessLevels,omitempty"`
}

// AccessContextManagerAccessLevelConditionStatus defines the config connector machine state of AccessContextManagerAccessLevelCondition
type AccessContextManagerAccessLevelConditionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AccessContextManagerAccessLevelCondition resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AccessContextManagerAccessLevelConditionObservedState `json:"observedState,omitempty"`
}

// AccessContextManagerAccessLevelConditionObservedState is the state of the AccessContextManagerAccessLevelCondition resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.identity.accesscontextmanager.v1.Condition
type AccessContextManagerAccessLevelConditionObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanageraccesslevelcondition;gcpaccesscontextmanageraccesslevelconditions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AccessContextManagerAccessLevelCondition is the Schema for the AccessContextManagerAccessLevelCondition API
// +k8s:openapi-gen=true
type AccessContextManagerAccessLevelCondition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AccessContextManagerAccessLevelConditionSpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessLevelConditionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AccessContextManagerAccessLevelConditionList contains a list of AccessContextManagerAccessLevelCondition
type AccessContextManagerAccessLevelConditionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerAccessLevelCondition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessLevelCondition{}, &AccessContextManagerAccessLevelConditionList{})
}
