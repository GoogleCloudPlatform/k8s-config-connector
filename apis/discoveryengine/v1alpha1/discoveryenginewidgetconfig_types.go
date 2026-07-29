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

var DiscoveryEngineWidgetConfigGVK = GroupVersion.WithKind("DiscoveryEngineWidgetConfig")

// DiscoveryEngineWidgetConfigSpec defines the desired state of DiscoveryEngineWidgetConfig
// +kcc:spec:proto=google.cloud.discoveryengine.v1alpha.WidgetConfig
type DiscoveryEngineWidgetConfigSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	Location string `json:"location"`

	// Immutable. The collection ID. If not specified, it defaults to "default_collection".
	// +optional
	CollectionID *string `json:"collectionID,omitempty"`

	// Immutable. The Engine that this WidgetConfig belongs to.
	// +required
	EngineRef *DiscoveryEngineEngineRef `json:"engineRef"`

	// The DiscoveryEngineWidgetConfig name (or ID). If not specified, metadata.name will be used.
	// Currently, GCP only accepts "default_search_widget_config".
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Access settings.
	// +optional
	AccessSettings *WidgetConfig_AccessSettings `json:"accessSettings,omitempty"`

	// UI settings.
	// +optional
	UiSettings *WidgetConfig_UiSettings `json:"uiSettings,omitempty"`

	// UI branding.
	// +optional
	UiBranding *WidgetConfig_UiBranding `json:"uiBranding,omitempty"`

	// Homepage setting.
	// +optional
	HomepageSetting *WidgetConfig_HomepageSetting `json:"homepageSetting,omitempty"`
}

// DiscoveryEngineWidgetConfigStatus defines the config connector machine state of DiscoveryEngineWidgetConfig
type DiscoveryEngineWidgetConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineWidgetConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineWidgetConfigObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineWidgetConfigObservedState is the state of the DiscoveryEngineWidgetConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1alpha.WidgetConfig
type DiscoveryEngineWidgetConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginewidgetconfig;gcpdiscoveryenginewidgetconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineWidgetConfig is the Schema for the DiscoveryEngineWidgetConfig API
// +k8s:openapi-gen=true
type DiscoveryEngineWidgetConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineWidgetConfigSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineWidgetConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineWidgetConfigList contains a list of DiscoveryEngineWidgetConfig
type DiscoveryEngineWidgetConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineWidgetConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineWidgetConfig{}, &DiscoveryEngineWidgetConfigList{})
}
