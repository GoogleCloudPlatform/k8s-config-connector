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

// API sources for TPUQueuedResource, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/tpu/v2/cloud_tpu.proto
// +kcc:source:service-docs=https://cloud.google.com/tpu/
// +kcc:source:resource-docs=https://docs.cloud.google.com/tpu/docs/reference/rest/v2/projects.locations.queuedResources

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TPUQueuedResourceGVK = GroupVersion.WithKind("TPUQueuedResource")

// TPUQueuedResourceSpec defines the desired state of TPUQueuedResource
// +kcc:spec:proto=google.cloud.tpu.v2.QueuedResource
// +kcc:required-from-proto
type TPUQueuedResourceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The TPUQueuedResource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Defines a TPU resource.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.tpu
	Tpu *QueuedResource_Tpu `json:"tpu,omitempty"`

	// Optional. The Spot tier.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.spot
	Spot *QueuedResource_Spot `json:"spot,omitempty"`

	// Optional. The Guaranteed tier
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.guaranteed
	Guaranteed *QueuedResource_Guaranteed `json:"guaranteed,omitempty"`

	// Optional. The queueing policy of the QueuedRequest.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.queueing_policy
	QueueingPolicy *QueuedResource_QueueingPolicy `json:"queueingPolicy,omitempty"`

	// Optional. Name of the reservation in which the resource should be
	//  provisioned. Format:
	//  projects/{project}/locations/{zone}/reservations/{reservation}
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.reservation_name
	ReservationName *string `json:"reservationName,omitempty"`
}

// TPUQueuedResourceStatus defines the config connector machine state of TPUQueuedResource
type TPUQueuedResourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TPUQueuedResource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TPUQueuedResourceObservedState `json:"observedState,omitempty"`
}

// TPUQueuedResourceObservedState is the state of the TPUQueuedResource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.tpu.v2.QueuedResource
type TPUQueuedResourceObservedState struct {
	// Output only. The time when the QueuedResource was created.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Defines a TPU resource.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.tpu
	Tpu *QueuedResource_TpuObservedState `json:"tpu,omitempty"`

	// Output only. State of the QueuedResource request.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.state
	State *QueuedResourceStateObservedState `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptpuqueuedresource;gcptpuqueuedresources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TPUQueuedResource is the Schema for the TPUQueuedResource API
// +k8s:openapi-gen=true
type TPUQueuedResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TPUQueuedResourceSpec   `json:"spec,omitempty"`
	Status TPUQueuedResourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TPUQueuedResourceList contains a list of TPUQueuedResource
type TPUQueuedResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TPUQueuedResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TPUQueuedResource{}, &TPUQueuedResourceList{})
}
