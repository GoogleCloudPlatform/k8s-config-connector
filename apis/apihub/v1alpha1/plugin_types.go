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

var APIHubPluginGVK = GroupVersion.WithKind("APIHubPlugin")

// APIHubPluginSpec defines the desired state of APIHubPlugin
// +kcc:spec:proto=google.cloud.apihub.v1.Plugin
type APIHubPluginSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The APIHubPlugin name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the plugin. Max length is 50 characters (Unicode code points).
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of the API.
	// +kubebuilder:validation:Required
	Type *AttributeValues `json:"type,omitempty"`

	// Optional. The plugin description. Max length is 2000 characters (Unicode code points).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`
}

// APIHubPluginStatus defines the config connector machine state of APIHubPlugin
type APIHubPluginStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubPlugin resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *APIHubPluginObservedState `json:"observedState,omitempty"`
}

// APIHubPluginObservedState is the state of the APIHubPlugin resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Plugin
type APIHubPluginObservedState struct {
	// Output only. Represents the state of the plugin.
	// +kcc:proto:field=google.cloud.apihub.v1.Plugin.state
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Required. The type of the API.
	// +kcc:proto:field=google.cloud.apihub.v1.Plugin.type
	// +kubebuilder:validation:Optional
	Type *AttributeValuesObservedState `json:"type,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubplugin;gcpapihubplugins
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubPlugin is the Schema for the APIHubPlugin API
// +k8s:openapi-gen=true
type APIHubPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubPluginSpec   `json:"spec,omitempty"`
	Status APIHubPluginStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubPluginList contains a list of APIHubPlugin
type APIHubPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubPlugin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubPlugin{}, &APIHubPluginList{})
}
