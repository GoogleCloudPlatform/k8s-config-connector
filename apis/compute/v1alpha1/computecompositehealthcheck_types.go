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

// API sources for ComputeCompositeHealthCheck, recorded by generate-types:
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

var ComputeCompositeHealthCheckGVK = GroupVersion.WithKind("ComputeCompositeHealthCheck")

// ComputeCompositeHealthCheckSpec defines the desired state of ComputeCompositeHealthCheck
// +kcc:spec:proto=google.cloud.compute.v1.CompositeHealthCheck
// +kcc:required-from-proto
type ComputeCompositeHealthCheckSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ComputeCompositeHealthCheck name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// An optional description of this resource. Provide this property when you
	//  create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.description
	Description *string `json:"description,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object.
	//  This field is used in optimistic locking. This field will be ignored when
	//  inserting a CompositeHealthCheck. An up-to-date fingerprint
	//  must be provided in order to patch the CompositeHealthCheck; Otherwise,
	//  the request will fail with error 412 conditionNotMet. To see
	//  the latest fingerprint, make a get() request to retrieve the
	//  CompositeHealthCheck.
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// URL to the destination resource. Must be set. Must be aForwardingRule. The ForwardingRule must have
	//  load balancing scheme INTERNAL orINTERNAL_MANAGED and must be regional and in the same region
	//  as the CompositeHealthCheck (cross-region deployment forINTERNAL_MANAGED is not supported). Can be mutated.
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.health_destination
	HealthDestination *string `json:"healthDestination,omitempty"`

	// URLs to the HealthSource resources whose results are AND'ed.
	//  I.e. he aggregated result is is HEALTHY only if all sources
	//  are HEALTHY. Must have at least 1. Must not have more than 10.
	//  Must be regional and in the same region as theCompositeHealthCheck. Can be mutated.
	// +kcc:guess=possible-reference target=ComputeHealthSource
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.health_sources
	HealthSources []string `json:"healthSources,omitempty"`

	// Output only. [Output Only] URL of the region where the composite health check resides.
	//  This field applies only to the regional resource. You must specify this
	//  field as part of the HTTP request URL. It is not settable as a field in
	//  the request body.
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.region
	Region *string `json:"region,omitempty"`
}

// ComputeCompositeHealthCheckStatus defines the config connector machine state of ComputeCompositeHealthCheck
type ComputeCompositeHealthCheckStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeCompositeHealthCheck resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeCompositeHealthCheckObservedState `json:"observedState,omitempty"`
}

// ComputeCompositeHealthCheckObservedState is the state of the ComputeCompositeHealthCheck resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.CompositeHealthCheck
type ComputeCompositeHealthCheckObservedState struct {
	// Output only. [Output Only] Creation timestamp inRFC3339
	//  text format.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Output only. [Output Only] A unique identifier for this resource type. The server
	//  generates this identifier.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.id
	ID *uint64 `json:"id,omitempty"`

	// Output only. [Output Only] Type of the resource. Alwayscompute#compositeHealthCheck for composite health checks.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. [Output Only] Server-defined URL for the resource.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. [Output Only] Server-defined URL with id for the resource.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.CompositeHealthCheck.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputecompositehealthcheck;gcpcomputecompositehealthchecks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeCompositeHealthCheck is the Schema for the ComputeCompositeHealthCheck API
// +k8s:openapi-gen=true
type ComputeCompositeHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeCompositeHealthCheckSpec   `json:"spec,omitempty"`
	Status ComputeCompositeHealthCheckStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeCompositeHealthCheckList contains a list of ComputeCompositeHealthCheck
type ComputeCompositeHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeCompositeHealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeCompositeHealthCheck{}, &ComputeCompositeHealthCheckList{})
}
