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

// API sources for ComputeHealthAggregationPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/compute/docs/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeHealthAggregationPolicyGVK = GroupVersion.WithKind("ComputeHealthAggregationPolicy")

// ComputeHealthAggregationPolicySpec defines the desired state of ComputeHealthAggregationPolicy
// +kcc:spec:proto=google.cloud.compute.v1.HealthAggregationPolicy
// +kcc:required-from-proto
type ComputeHealthAggregationPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ComputeHealthAggregationPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.healthy_percent_threshold
	HealthyPercentThreshold *uint32 `json:"healthyPercentThreshold,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.min_healthy_threshold
	MinHealthyThreshold *uint32 `json:"minHealthyThreshold,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.policy_type
	PolicyType *string `json:"policyType,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.region
	Region *string `json:"region,omitempty"`
}

// ComputeHealthAggregationPolicyStatus defines the config connector machine state of ComputeHealthAggregationPolicy
type ComputeHealthAggregationPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeHealthAggregationPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeHealthAggregationPolicyObservedState `json:"observedState,omitempty"`
}

// ComputeHealthAggregationPolicyObservedState is the state of the ComputeHealthAggregationPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.HealthAggregationPolicy
type ComputeHealthAggregationPolicyObservedState struct {
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.id
	ID *uint64 `json:"id,omitempty"`

	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.kind
	Kind *string `json:"kind,omitempty"`

	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputehealthaggregationpolicy;gcpcomputehealthaggregationpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeHealthAggregationPolicy is the Schema for the ComputeHealthAggregationPolicy API
// +k8s:openapi-gen=true
type ComputeHealthAggregationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeHealthAggregationPolicySpec   `json:"spec,omitempty"`
	Status ComputeHealthAggregationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeHealthAggregationPolicyList contains a list of ComputeHealthAggregationPolicy
type ComputeHealthAggregationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeHealthAggregationPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeHealthAggregationPolicy{}, &ComputeHealthAggregationPolicyList{})
}
