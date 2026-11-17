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

// API sources for ComputeHealthSource, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/regionHealthSources

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeHealthSourceGVK = GroupVersion.WithKind("ComputeHealthSource")

// ComputeHealthSourceSpec defines the desired state of ComputeHealthSource
// +kcc:spec:proto=google.cloud.compute.v1.HealthSource
// +kcc:required-from-proto
type ComputeHealthSourceSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Location represents the geographical location of the ComputeHealthSource. Specify a region name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// The ComputeHealthSource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.health_aggregation_policy
	HealthAggregationPolicyRef *ComputeHealthAggregationPolicyRef `json:"healthAggregationPolicyRef,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.sources
	Sources []computev1beta1.ComputeBackendServiceRef `json:"sources,omitempty"`
}

// ComputeHealthSourceStatus defines the config connector machine state of ComputeHealthSource
type ComputeHealthSourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeHealthSource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeHealthSourceObservedState `json:"observedState,omitempty"`
}

// ComputeHealthSourceObservedState is the state of the ComputeHealthSource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.HealthSource
type ComputeHealthSourceObservedState struct {
	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.id
	ID *uint64 `json:"id,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.kind
	Kind *string `json:"kind,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.region
	Region *string `json:"region,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthSource.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputehealthsource;gcpcomputehealthsources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeHealthSource is the Schema for the ComputeHealthSource API
// +k8s:openapi-gen=true
type ComputeHealthSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeHealthSourceSpec   `json:"spec,omitempty"`
	Status ComputeHealthSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeHealthSourceList contains a list of ComputeHealthSource
type ComputeHealthSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeHealthSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeHealthSource{}, &ComputeHealthSourceList{})
}
