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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineSessionGVK = GroupVersion.WithKind("DiscoveryEngineSession")

// DiscoveryEngineSessionSpec defines the desired state of DiscoveryEngineSession
// +kcc:spec:proto=google.cloud.discoveryengine.v1.Session
type DiscoveryEngineSessionSpec struct {
	// The DataStore this session should be part of.
	// +required
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef"`

	// The DiscoveryEngineSession name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The display name of the session.
	// This field is used to identify the session in the UI.
	// By default, the display name is the first turn query text in the session.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// The state of the session.
	// +optional
	State *string `json:"state,omitempty"`

	// A unique identifier for tracking users.
	// +optional
	UserPseudoID *string `json:"userPseudoID,omitempty"`

	// Turns.
	// +optional
	Turns []Session_Turn `json:"turns,omitempty"`

	// Optional. Whether the session is pinned, pinned session will be displayed
	// on the top of the session list.
	// +optional
	IsPinned *bool `json:"isPinned,omitempty"`
}

// DiscoveryEngineSessionStatus defines the config connector machine state of DiscoveryEngineSession
type DiscoveryEngineSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineSessionObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineSessionObservedState is the state of the DiscoveryEngineSession resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.Session
type DiscoveryEngineSessionObservedState struct {
	// Turns.
	// +optional
	Turns []Session_TurnObservedState `json:"turns,omitempty"`

	// Output only. The time the session started.
	// +optional
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time the session finished.
	// +optional
	EndTime *string `json:"endTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginesession;gcpdiscoveryenginesessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineSession is the Schema for the DiscoveryEngineSession API
// +k8s:openapi-gen=true
type DiscoveryEngineSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineSessionSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineSessionList contains a list of DiscoveryEngineSession
type DiscoveryEngineSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineSession{}, &DiscoveryEngineSessionList{})
}
