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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EventarcMessageBusGVK = GroupVersion.WithKind("EventarcMessageBus")

// EventarcMessageBusSpec defines the desired state of EventarcMessageBus
// +kcc:spec:proto=google.cloud.eventarc.v1.MessageBus
type EventarcMessageBusSpec struct {
	Parent `json:",inline"`

	// The EventarcMessageBus name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Resource display name.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.display_name
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt their event data.
	//
	//  It must match the pattern
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.crypto_key_name
	// +kubebuilder:validation:Optional
	CryptoKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"cryptoKeyRef,omitempty"`

	// Optional. Config to control Platform logging for the Message Bus. This log
	//  configuration is applied to the Message Bus itself, and all the Enrollments
	//  attached to it.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.logging_config
	// +kubebuilder:validation:Optional
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// EventarcMessageBusStatus defines the config connector machine state of EventarcMessageBus
type EventarcMessageBusStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EventarcMessageBus resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EventarcMessageBusObservedState `json:"observedState,omitempty"`
}

// EventarcMessageBusObservedState is the state of the EventarcMessageBus resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.eventarc.v1.MessageBus
type EventarcMessageBusObservedState struct {
	// Output only. Server assigned unique identifier for the channel. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.uid
	// +kubebuilder:validation:Optional
	Uid *string `json:"uid,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and might be sent only on update and delete requests to
	//  ensure that the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.etag
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.create_time
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.MessageBus.update_time
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpeventarcmessagebus;gcpeventarcmessagebuses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EventarcMessageBus is the Schema for the EventarcMessageBus API
// +k8s:openapi-gen=true
type EventarcMessageBus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EventarcMessageBusSpec   `json:"spec,omitempty"`
	Status EventarcMessageBusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EventarcMessageBusList contains a list of EventarcMessageBus
type EventarcMessageBusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventarcMessageBus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventarcMessageBus{}, &EventarcMessageBusList{})
}
