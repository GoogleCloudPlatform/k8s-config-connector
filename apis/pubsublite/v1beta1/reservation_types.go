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

// +tool:krm-type-terraform
// proto.message: google.cloud.pubsublite.v1.Reservation
// crd.kind: PubSubLiteReservation
// crd.version: v1beta1
// terraform.src: github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsublite/resource_pubsub_lite_reservation.go

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PubSubLiteReservationGVK = GroupVersion.WithKind("PubSubLiteReservation")

// PubSubLiteReservationSpec defines the desired state of PubSubLiteReservation
// +kcc:spec:proto=google.cloud.pubsublite.v1.Reservation
type PubSubLiteReservationSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *parent.ProjectRef `json:"projectRef,omitempty"`

	// The region of the pubsub lite reservation.
	// +required
	Region *string `json:"region,omitempty"`

	// The PubSubLiteReservation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The reserved throughput capacity. Every unit of throughput capacity is
	//  equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	//  messages.
	//
	//  Any topics which are declared as using capacity from a Reservation will
	//  consume resources from this reservation instead of being charged
	//  individually.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Reservation.throughput_capacity
	// +required
	ThroughputCapacity *int64 `json:"throughputCapacity,omitempty"`
}

// PubSubLiteReservationStatus defines the config connector machine state of PubSubLiteReservation
type PubSubLiteReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the PubSubLiteReservation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET(terraform)
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *PubSubLiteReservationObservedState `json:"observedState,omitempty"`
}

// PubSubLiteReservationObservedState is the state of the PubSubLiteReservation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.pubsublite.v1.Reservation
type PubSubLiteReservationObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsublitereservation;gcppubsublitereservations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=beta"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubLiteReservation is the Schema for the PubSubLiteReservation API
// +k8s:openapi-gen=true
type PubSubLiteReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PubSubLiteReservationSpec   `json:"spec,omitempty"`
	Status PubSubLiteReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PubSubLiteReservationList contains a list of PubSubLiteReservation
type PubSubLiteReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubLiteReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubLiteReservation{}, &PubSubLiteReservationList{})
}
