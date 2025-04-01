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

package v1alpha1

import (
	connectorsv1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connector/v1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EventarcChannelGVK = GroupVersion.WithKind("EventarcChannel")

type Parent struct {
	// +required
	Location string `json:"location"`
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// EventarcChannelSpec defines the desired state of EventarcChannel
// +kcc:proto=google.cloud.eventarc.v1.Channel
type EventarcChannelSpec struct {
	Parent `json:",inline"`
	// The EventarcChannel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated
	// with the channel. This provider will be granted permissions to publish
	// events to the channel.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.provider
	Provider *connectorsv1.ProviderRef `json:"provider,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to
	// encrypt/decrypt their event data.
	//
	// It must match the pattern
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.crypto_key_name
	KmsKeyRef *refv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// EventarcChannelStatus defines the config connector machine state of EventarcChannel
type EventarcChannelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EventarcChannel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EventarcChannelObservedState `json:"observedState,omitempty"`
}

// EventarcChannelObservedState is the state of the EventarcChannel resource as most recently observed in GCP.
// +kcc:proto=google.cloud.eventarc.v1.Channel
type EventarcChannelObservedState struct {
	// Output only. Server assigned unique identifier for the channel. The value
	// is a UUID4 string and guaranteed to remain unchanged until the resource is
	// deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The name of the Pub/Sub topic created and managed by
	// Eventarc system as a transport for the event delivery. Format:
	// `projects/{project}/topics/{topic_id}`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.pubsub_topic
	PubsubTopic *string `json:"pubsubTopic,omitempty"`

	// Output only. The state of a Channel.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.state
	State *string `json:"state,omitempty"`

	// Output only. The activation token for the channel. The token must be used
	// by the provider to register the channel for publishing.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.activation_token
	ActivationToken *string `json:"activationToken,omitempty"`

	// Output only. Whether or not this Channel satisfies the requirements of
	// physical zone separation
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpeventarcchannel;gcpeventarcchannels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EventarcChannel is the Schema for the EventarcChannel API
// +k8s:openapi-gen=true
type EventarcChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EventarcChannelSpec   `json:"spec,omitempty"`
	Status EventarcChannelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EventarcChannelList contains a list of EventarcChannel
type EventarcChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventarcChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventarcChannel{}, &EventarcChannelList{})
}
