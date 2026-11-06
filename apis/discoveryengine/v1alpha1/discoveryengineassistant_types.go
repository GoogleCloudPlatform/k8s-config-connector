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

// API sources for DiscoveryEngineAssistant, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/6a1db32dc52fd9421c1743dc8f5faa7a15e518d1/google/cloud/discoveryengine/v1/assistant.proto
// +kcc:source:service-docs=https://cloud.google.com/generative-ai-app-builder/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1/projects.locations.collections.engines.assistants

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineAssistantGVK = GroupVersion.WithKind("DiscoveryEngineAssistant")

// DiscoveryEngineAssistantSpec defines the desired state of DiscoveryEngineAssistant
// +kcc:spec:proto=google.cloud.discoveryengine.v1.Assistant
// +kcc:required-from-proto
type DiscoveryEngineAssistantSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/collections/{collection}/engines/{engine} this resource belongs to.
	// +kcc:guess
	// EngineRef *DiscoveryEngineEngineRef `json:"engineRef,omitempty"`

	// The DiscoveryEngineAssistant name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The assistant display name.
	//
	//  It must be a UTF-8 encoded string with a length limit of 128 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for additional information. Expected to be shown on
	//  the configuration UI, not to the users of the assistant.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.description
	Description *string `json:"description,omitempty"`

	// Optional. Configuration for the generation of the assistant response.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.generation_config
	GenerationConfig *Assistant_GenerationConfig `json:"generationConfig,omitempty"`

	// Optional. The type of web grounding to use.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.web_grounding_type
	WebGroundingType *string `json:"webGroundingType,omitempty"`

	// Optional. This field controls the default web grounding toggle for end
	//  users if `web_grounding_type` is set to `WEB_GROUNDING_TYPE_GOOGLE_SEARCH`
	//  or `WEB_GROUNDING_TYPE_ENTERPRISE_WEB_SEARCH`. By default, this field is
	//  set to false. If `web_grounding_type` is `WEB_GROUNDING_TYPE_GOOGLE_SEARCH`
	//  or `WEB_GROUNDING_TYPE_ENTERPRISE_WEB_SEARCH`, end users will have web
	//  grounding enabled by default on UI. If true, grounding toggle will be
	//  disabled by default on UI. End users can still enable web grounding in
	//  the UI if web grounding is enabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.default_web_grounding_toggle_off
	DefaultWebGroundingToggleOff *bool `json:"defaultWebGroundingToggleOff,omitempty"`

	// Optional. Note: not implemented yet. Use
	//  [enabled_actions][google.cloud.discoveryengine.v1.Assistant.enabled_actions]
	//  instead. The enabled tools on this assistant. The keys are connector name,
	//  for example
	//  "projects/{projectId}/locations/{locationId}/collections/{collectionId}/dataconnector
	//  The values consist of admin enabled tools towards the connector
	//  instance. Admin can selectively enable multiple tools on any of the
	//  connector instances that they created in the project. For example
	//  {"jira1ConnectorName": [(toolId1, "createTicket"), (toolId2,
	//  "transferTicket")],
	//   "gmail1ConnectorName": [(toolId3, "sendEmail"),..] }
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.enabled_tools
	EnabledTools map[string]Assistant_ToolList `json:"enabledTools,omitempty"`

	// Optional. Customer policy for the assistant.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.customer_policy
	CustomerPolicy *Assistant_CustomerPolicy `json:"customerPolicy,omitempty"`
}

// DiscoveryEngineAssistantStatus defines the config connector machine state of DiscoveryEngineAssistant
type DiscoveryEngineAssistantStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineAssistant resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineAssistantObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineAssistantObservedState is the state of the DiscoveryEngineAssistant resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.Assistant
type DiscoveryEngineAssistantObservedState struct {
	// Output only. Represents the time when this Assistant was created.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Represents the time when this Assistant was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Assistant.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineassistant;gcpdiscoveryengineassistants
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineAssistant is the Schema for the DiscoveryEngineAssistant API
// +k8s:openapi-gen=true
type DiscoveryEngineAssistant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineAssistantSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineAssistantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineAssistantList contains a list of DiscoveryEngineAssistant
type DiscoveryEngineAssistantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineAssistant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineAssistant{}, &DiscoveryEngineAssistantList{})
}
