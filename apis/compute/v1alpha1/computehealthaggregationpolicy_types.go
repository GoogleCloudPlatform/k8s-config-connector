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
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/regionHealthAggregationPolicies

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
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Location represents the geographical location of the ComputeHealthAggregationPolicy. Specify a region name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// The ComputeHealthAggregationPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.description
	Description *string `json:"description,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthAggregationPolicy. An up-to-date fingerprint must be provided in order to patch/update the HealthAggregationPolicy; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthAggregationPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.healthy_percent_threshold
	HealthyPercentThreshold *uint32 `json:"healthyPercentThreshold,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.min_healthy_threshold
	MinHealthyThreshold *uint32 `json:"minHealthyThreshold,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.policy_type
	PolicyType *string `json:"policyType,omitempty"`
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
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.id
	ID *uint64 `json:"id,omitempty"`

	// [Output only] Type of the resource. Always compute#healthAggregationPolicy for health aggregation policies.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the health aggregation policy resides. This field is not applicable to global health aggregation policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.HealthAggregationPolicy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource with the resource id.
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
