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

// API sources for DiscoveryEngineWidgetConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/6a1db32dc52fd9421c1743dc8f5faa7a15e518d1/google/cloud/discoveryengine/v1/widget_config.proto
// +kcc:source:service-docs=https://cloud.google.com/generative-ai-app-builder/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1/projects.locations.dataStores.widgetConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineWidgetConfigGVK = GroupVersion.WithKind("DiscoveryEngineWidgetConfig")

// DiscoveryEngineWidgetConfigSpec defines the desired state of DiscoveryEngineWidgetConfig
// +kcc:spec:proto=google.cloud.discoveryengine.v1.WidgetConfig
// +kcc:required-from-proto
type DiscoveryEngineWidgetConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/dataStores/{data_store} this resource belongs to.
	// +kcc:guess
	// DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// The DiscoveryEngineWidgetConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human readable widget config display name. Used in Discovery
	//  UI.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. Specifies the solution type that this WidgetConfig can
	//  be used for.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.solution_type
	// +required
	SolutionType *string `json:"solutionType,omitempty"`

	// Whether to show the result score.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.enable_result_score
	EnableResultScore *bool `json:"enableResultScore,omitempty"`

	// Describes general widget search settings as seen in cloud console
	//  widget configuration page. Replaces top deprecated top level properties.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.ui_settings
	UiSettings *WidgetConfig_UiSettings `json:"uiSettings,omitempty"`

	// Will be used for all widget access settings seen in cloud console
	//  integration page. Replaces top deprecated top level properties.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.access_settings
	AccessSettings *WidgetConfig_AccessSettings `json:"accessSettings,omitempty"`

	// Describes search widget UI branding settings, such as the widget title,
	//  logo, favicons, and colors.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.ui_branding
	UiBranding *WidgetConfig_UiBrandingSettings `json:"uiBranding,omitempty"`

	// Optional. Describes the homepage settings of the widget.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.homepage_setting
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
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.WidgetConfig
type DiscoveryEngineWidgetConfigObservedState struct {
	// Output only. Timestamp the
	//  [WidgetConfig][google.cloud.discoveryengine.v1.WidgetConfig] was created.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp the
	//  [WidgetConfig][google.cloud.discoveryengine.v1.WidgetConfig] was updated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The type of the parent data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// Output only. Unique obfuscated identifier of a
	//  [WidgetConfig][google.cloud.discoveryengine.v1.WidgetConfig].
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.config_id
	ConfigID *string `json:"configID,omitempty"`

	// Output only. The industry vertical that the WidgetConfig registers.
	//  The WidgetConfig industry vertical is based on the associated
	//  [Engine][google.cloud.discoveryengine.v1.Engine].
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.industry_vertical
	IndustryVertical *string `json:"industryVertical,omitempty"`

	// Output only. Whether LLM is enabled in the corresponding data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.llm_enabled
	LlmEnabled *bool `json:"llmEnabled,omitempty"`

	// Output only. Whether the customer accepted data use terms.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.minimum_data_term_accepted
	MinimumDataTermAccepted *bool `json:"minimumDataTermAccepted,omitempty"`

	// Output only. Collection components that lists all collections and child
	//  data stores associated with the widget config, those data sources can be
	//  used for filtering in widget service APIs, users can return results that
	//  from selected data sources.
	//
	//  For SaaS / Business engines, when `LookupWidgetConfig` is called with
	//  `view = WITH_AVAILABLE_CONNECTORS`, this list is additionally augmented
	//  with synthetic placeholder entries for connectors the caller may attach
	//  but has not yet attached (see `CollectionComponent` for the placeholder
	//  contract). The frontend can therefore render a unified list of
	//  already-attached and available-to-attach sources by iterating this single
	//  field. For Enterprise engines and for the default `view`, only
	//  already-attached connectors are returned (today's behavior).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.collection_components
	CollectionComponents []WidgetConfig_CollectionComponentObservedState `json:"collectionComponents,omitempty"`

	// Describes general widget search settings as seen in cloud console
	//  widget configuration page. Replaces top deprecated top level properties.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.ui_settings
	UiSettings *WidgetConfig_UiSettingsObservedState `json:"uiSettings,omitempty"`

	// Optional. Output only. Describes the assistant settings of the widget.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.assistant_settings
	AssistantSettings *WidgetConfig_AssistantSettingsObservedState `json:"assistantSettings,omitempty"`

	// Optional. Output only. Describes the customer related configurations,
	//  currently only used for government customers. This field cannot be modified
	//  after project onboarding.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.customer_provided_config
	CustomerProvidedConfig *WidgetConfig_CustomerProvidedConfigObservedState `json:"customerProvidedConfig,omitempty"`

	// Optional. Output only. Whether to enable private knowledge graph.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.enable_private_knowledge_graph
	EnablePrivateKnowledgeGraph *bool `json:"enablePrivateKnowledgeGraph,omitempty"`

	// Output only. Whether the subscription is gemini bundle or not.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.gemini_bundle
	GeminiBundle *bool `json:"geminiBundle,omitempty"`

	// Output only. The batch authorization statuses for the widget's connectors.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.batch_auth_statuses
	BatchAuthStatuses []WidgetConfig_BatchAuthStatusObservedState `json:"batchAuthStatuses,omitempty"`

	// Output only. The nodes associated with the Widget Config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WidgetConfig.nodes
	Nodes []WidgetConfig_NodeObservedState `json:"nodes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginewidgetconfig;gcpdiscoveryenginewidgetconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
