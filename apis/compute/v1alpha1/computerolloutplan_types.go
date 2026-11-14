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

// API sources for ComputeRolloutPlan, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/7e7f1e9533723bcf3821e681e564b27f92fd46f5/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/compute/docs/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRolloutPlanSpec defines the desired state of ComputeRolloutPlan
// +kcc:spec:proto=google.cloud.compute.v1.RolloutPlan
// +kcc:required-from-proto
type ComputeRolloutPlanSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ComputeRolloutPlan name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// An optional description of this resource. Provide this property when you
	//  create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.description
	Description *string `json:"description,omitempty"`

	// The location scope of the rollout plan. If not specified, the location
	//  scope is considered as ZONAL.
	//  Check the LocationScope enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.location_scope
	LocationScope *string `json:"locationScope,omitempty"`

	// Required. The waves included in this rollout plan.
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.waves
	Waves []RolloutPlanWave `json:"waves,omitempty"`
}

// ComputeRolloutPlanStatus defines the config connector machine state of ComputeRolloutPlan
type ComputeRolloutPlanStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRolloutPlan resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRolloutPlanObservedState `json:"observedState,omitempty"`
}

// ComputeRolloutPlanObservedState is the state of the ComputeRolloutPlan resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.RolloutPlan
type ComputeRolloutPlanObservedState struct {
	// Output only. [Output Only] Creation timestamp inRFC3339
	//  text format.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Output only. [Output Only] The unique identifier for the resource. This identifier is
	//  defined by the server.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.id
	ID *uint64 `json:"id,omitempty"`

	// Output only. [Output Only] Type of the resource. Always compute#rolloutPlan
	//  for rolloutPlans.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. [Output Only] Server-defined fully-qualified URL for this resource.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. [Output Only] Server-defined URL for this resource's resource id.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.RolloutPlan.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputerolloutplan;gcpcomputerolloutplans
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRolloutPlan is the Schema for the ComputeRolloutPlan API
// +k8s:openapi-gen=true
type ComputeRolloutPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRolloutPlanSpec   `json:"spec,omitempty"`
	Status ComputeRolloutPlanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRolloutPlanList contains a list of ComputeRolloutPlan
type ComputeRolloutPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRolloutPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRolloutPlan{}, &ComputeRolloutPlanList{})
}
