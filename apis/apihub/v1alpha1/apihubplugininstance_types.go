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

// API sources for APIHubPluginInstance, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/apihub/v1/plugin_service.proto
// +kcc:source:service-docs=https://cloud.google.com/apigee/docs/api-hub/what-is-api-hub
// +kcc:source:resource-docs=https://docs.cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.plugins.instances

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIHubPluginInstanceGVK = GroupVersion.WithKind("APIHubPluginInstance")

// APIHubPluginInstanceSpec defines the desired state of APIHubPluginInstance
// +kcc:spec:proto=google.cloud.apihub.v1.PluginInstance
// +kcc:required-from-proto
type APIHubPluginInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The APIHubPluginInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name for this plugin instance. Max length is 255
	//  characters.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The authentication information for this plugin instance.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.auth_config
	AuthConfig *AuthConfig `json:"authConfig,omitempty"`

	// Optional. The additional information for this plugin instance corresponding
	//  to the additional config template of the plugin. This information will be
	//  sent to plugin hosting service on each call to plugin hosted service. The
	//  key will be the config_variable_template.display_name to uniquely identify
	//  the config variable.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.additional_config
	AdditionalConfig map[string]ConfigVariable `json:"additionalConfig,omitempty"`

	// Required. The action status for the plugin instance.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.actions
	// +required
	Actions []PluginInstanceAction `json:"actions,omitempty"`

	// Optional. The source project id of the plugin instance. This will be the id
	//  of runtime project in case of gcp based plugins and org id in case of non
	//  gcp based plugins. This field will be a required field for Google provided
	//  on-ramp plugins.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.source_project_id
	SourceProjectID *string `json:"sourceProjectID,omitempty"`
}

// APIHubPluginInstanceStatus defines the config connector machine state of APIHubPluginInstance
type APIHubPluginInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubPluginInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubPluginInstanceObservedState `json:"observedState,omitempty"`
}

// APIHubPluginInstanceObservedState is the state of the APIHubPluginInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.PluginInstance
type APIHubPluginInstanceObservedState struct {
	// Output only. The current state of the plugin instance (e.g., enabled,
	//  disabled, provisioning).
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.state
	State *string `json:"state,omitempty"`

	// Output only. Error message describing the failure, if any, during Create,
	//  Delete or ApplyConfig operation corresponding to the plugin instance.This
	//  field will only be populated if the plugin instance is in the ERROR or
	//  FAILED state.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Required. The action status for the plugin instance.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.actions
	Actions []PluginInstanceActionObservedState `json:"actions,omitempty"`

	// Output only. Timestamp indicating when the plugin instance was created.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp indicating when the plugin instance was last
	//  updated.
	// +kcc:proto:field=google.cloud.apihub.v1.PluginInstance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubplugininstance;gcpapihubplugininstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubPluginInstance is the Schema for the APIHubPluginInstance API
// +k8s:openapi-gen=true
type APIHubPluginInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubPluginInstanceSpec   `json:"spec,omitempty"`
	Status APIHubPluginInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubPluginInstanceList contains a list of APIHubPluginInstance
type APIHubPluginInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubPluginInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubPluginInstance{}, &APIHubPluginInstanceList{})
}
