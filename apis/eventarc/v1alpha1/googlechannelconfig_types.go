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

// +generated:types
// krm.group: eventarc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.eventarc.v1
// resource: EventarcGoogleChannelConfig:GoogleChannelConfig

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EventarcGoogleChannelConfigGVK = GroupVersion.WithKind("EventarcGoogleChannelConfig")

// EventarcGoogleChannelConfigSpec defines the desired state of EventarcGoogleChannelConfig
// +kcc:spec:proto=google.cloud.eventarc.v1.GoogleChannelConfig
type EventarcGoogleChannelConfigSpec struct {
	Parent `json:",inline"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt their event data.
	//
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleChannelConfig.crypto_key_name
	CryptoKeyRef *refv1beta1.KMSCryptoKeyRef `json:"cryptoKeyRef,omitempty"`

	// The user-provided name of the EventarcGoogleChannelConfig. If not specified, the name of the KRM resource will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// EventarcGoogleChannelConfigStatus defines the config connector machine state of EventarcGoogleChannelConfig
type EventarcGoogleChannelConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EventarcGoogleChannelConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EventarcGoogleChannelConfigObservedState `json:"observedState,omitempty"`
}

// EventarcGoogleChannelConfigObservedState is the state of the EventarcGoogleChannelConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.eventarc.v1.GoogleChannelConfig
type EventarcGoogleChannelConfigObservedState struct {
	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleChannelConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpeventarcgooglechannelconfig;gcpeventarcgooglechannelconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EventarcGoogleChannelConfig is the Schema for the EventarcGoogleChannelConfig API
// +k8s:openapi-gen=true
type EventarcGoogleChannelConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EventarcGoogleChannelConfigSpec   `json:"spec,omitempty"`
	Status EventarcGoogleChannelConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EventarcGoogleChannelConfigList contains a list of EventarcGoogleChannelConfig
type EventarcGoogleChannelConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventarcGoogleChannelConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventarcGoogleChannelConfig{}, &EventarcGoogleChannelConfigList{})
}
