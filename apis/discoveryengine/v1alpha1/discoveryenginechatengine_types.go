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

var DiscoveryEngineChatEngineGVK = GroupVersion.WithKind("DiscoveryEngineChatEngine")

// DiscoveryEngineChatEngineSpec defines the desired state of DiscoveryEngineChatEngine
// +kcc:spec:proto=google.cloud.discoveryengine.v1.Engine
type DiscoveryEngineChatEngineSpec struct {
	// Required. The display name of the engine. Should be human readable. UTF-8
	// encoded string with limit of 1024 characters.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The data stores associated with this engine.
	// Multiple DataStores in the same Collection can be associated here.
	// Note that when used in CreateEngineRequest, one DataStore must be
	// provided as the system will use it for necessary initializations.
	DataStoreRefs []*DiscoveryEngineDataStoreRef `json:"dataStoreRefs,omitempty"`

	// The industry vertical that the engine registers.
	// The restriction of the Engine industry vertical is based on
	// DataStore: If unspecified, default to `GENERIC`. Vertical on Engine
	// has to match vertical of the DataStore linked to the engine.
	IndustryVertical *string `json:"industryVertical,omitempty"`

	// Common config spec that specifies the metadata of the engine.
	CommonConfig *Engine_CommonConfig `json:"commonConfig,omitempty"`

	// Required. Configurations for the Chat Engine.
	// +required
	ChatEngineConfig *Engine_ChatEngineConfig `json:"chatEngineConfig,omitempty"`

	// Optional. Whether to disable analytics for searches performed on this
	// engine.
	DisableAnalytics *bool `json:"disableAnalytics,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the resource. */
	// +required
	Location string `json:"location"`

	// Immutable. The collection for the Engine.
	// +required
	Collection string `json:"collection"`

	// Immutable.
	// The DiscoveryEngineChatEngine name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// DiscoveryEngineChatEngineStatus defines the config connector machine state of DiscoveryEngineChatEngine
type DiscoveryEngineChatEngineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineChatEngine resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineChatEngineObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineChatEngineObservedState is the state of the DiscoveryEngineChatEngine resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.Engine
type DiscoveryEngineChatEngineObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginechatengine;gcpdiscoveryenginechatengines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineChatEngine is the Schema for the DiscoveryEngineChatEngine API
// +k8s:openapi-gen=true
type DiscoveryEngineChatEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineChatEngineSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineChatEngineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineChatEngineList contains a list of DiscoveryEngineChatEngine
type DiscoveryEngineChatEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineChatEngine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineChatEngine{}, &DiscoveryEngineChatEngineList{})
}
