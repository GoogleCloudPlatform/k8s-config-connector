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

// API sources for CESApp, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1526e545e9d26f23b9c5d0f04af17297def8d045/google/cloud/ces/v1/app.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/reference/rest/v1/projects.locations.apps

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESAppGVK = GroupVersion.WithKind("CESApp")

// +kcc:proto=google.cloud.ces.v1.App.VariableDeclaration
type App_VariableDeclaration struct {
	// Required. The name of the variable. The name must start with a letter or
	//  underscore and contain only letters, numbers, or underscores.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.name
	// +required
	Name *string `json:"name,omitempty"`

	// Required. The description of the variable.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.description
	// +required
	Description *string `json:"description,omitempty"`

	// Required. The schema of the variable.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.schema
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	// +required
	Schema *SchemaWithRequired `json:"schema,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.ces.v1.App.VariableDeclaration
type App_VariableDeclarationObservedState struct {
	// Required. The name of the variable. The name must start with a letter or
	//  underscore and contain only letters, numbers, or underscores.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.name
	Name *string `json:"name,omitempty"`

	// Required. The description of the variable.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.description
	Description *string `json:"description,omitempty"`

	// Required. The schema of the variable.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.schema
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	Schema *SchemaObservedState `json:"schema,omitempty"`
}

// CESAppSpec defines the desired state of CESApp
// +kcc:spec:proto=google.cloud.ces.v1.App
// +kcc:required-from-proto
type CESAppSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CESApp name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Human-readable description of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.description
	Description *string `json:"description,omitempty"`

	// Optional. Whether the app is pinned in the app list.
	// +kcc:proto:field=google.cloud.ces.v1.App.pinned
	Pinned *bool `json:"pinned,omitempty"`

	// Optional. The root agent is the entry point of the app.
	//  Format: `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.App.root_agent
	RootAgent *string `json:"rootAgent,omitempty"`

	// Optional. Language settings of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.language_settings
	LanguageSettings *LanguageSettings `json:"languageSettings,omitempty"`

	// Optional. TimeZone settings of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.time_zone_settings
	TimeZoneSettings *TimeZoneSettings `json:"timeZoneSettings,omitempty"`

	// Optional. Audio processing configuration of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.audio_processing_config
	AudioProcessingConfig *AudioProcessingConfig `json:"audioProcessingConfig,omitempty"`

	// Optional. Logging settings of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.logging_settings
	LoggingSettings *LoggingSettings `json:"loggingSettings,omitempty"`

	// Optional. Error handling settings of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.error_handling_settings
	ErrorHandlingSettings *ErrorHandlingSettings `json:"errorHandlingSettings,omitempty"`

	// Optional. The default LLM model settings for the app.
	//  Individual resources (e.g. agents, guardrails) can override these
	//  configurations as needed.
	// +kcc:proto:field=google.cloud.ces.v1.App.model_settings
	ModelSettings *ModelSettings `json:"modelSettings,omitempty"`

	// Optional. The tool execution mode for the app. If not provided, will
	//  default to PARALLEL.
	// +kcc:proto:field=google.cloud.ces.v1.App.tool_execution_mode
	ToolExecutionMode *string `json:"toolExecutionMode,omitempty"`

	// Optional. The evaluation thresholds for the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.evaluation_metrics_thresholds
	EvaluationMetricsThresholds *EvaluationMetricsThresholds `json:"evaluationMetricsThresholds,omitempty"`

	// Optional. The declarations of the variables.
	// +kcc:proto:field=google.cloud.ces.v1.App.variable_declarations
	VariableDeclarations []App_VariableDeclaration `json:"variableDeclarations,omitempty"`

	// Optional. Instructions for all the agents in the app.
	//  You can use this instruction to set up a stable identity or personality
	//  across all the agents.
	// +kcc:proto:field=google.cloud.ces.v1.App.global_instruction
	GlobalInstruction *string `json:"globalInstruction,omitempty"`

	// Optional. List of guardrails for the app.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}`
	// +kcc:proto:field=google.cloud.ces.v1.App.guardrails
	Guardrails []string `json:"guardrails,omitempty"`

	// Optional. The default channel profile used by the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.default_channel_profile
	DefaultChannelProfile *ChannelProfile `json:"defaultChannelProfile,omitempty"`

	// Optional. Metadata about the app. This field can be used to store
	//  additional information relevant to the app's details or intended usages.
	// +kcc:proto:field=google.cloud.ces.v1.App.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. The default client certificate settings for the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.client_certificate_settings
	ClientCertificateSettings *ClientCertificateSettings `json:"clientCertificateSettings,omitempty"`

	// Optional. Indicates whether the app is locked for changes. If the app is
	//  locked, modifications to the app resources will be rejected.
	// +kcc:proto:field=google.cloud.ces.v1.App.locked
	Locked *bool `json:"locked,omitempty"`
}

// CESAppStatus defines the config connector machine state of CESApp
type CESAppStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESApp resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESAppObservedState `json:"observedState,omitempty"`
}

// CESAppObservedState is the state of the CESApp resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1.App
type CESAppObservedState struct {
	// Output only. The declarations of predefined variables for the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.predefined_variable_declarations
	PredefinedVariableDeclarations []App_VariableDeclarationObservedState `json:"predefinedVariableDeclarations,omitempty"`

	// Optional. The data store settings for the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.data_store_settings
	DataStoreSettings *DataStoreSettingsObservedState `json:"dataStoreSettings,omitempty"`

	// Output only. Timestamp when the app was created.
	// +kcc:proto:field=google.cloud.ces.v1.App.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the app was last updated.
	// +kcc:proto:field=google.cloud.ces.v1.App.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Etag used to ensure the object hasn't changed during a
	//  read-modify-write operation. If the etag is empty, the update will
	//  overwrite any concurrent changes.
	// +kcc:proto:field=google.cloud.ces.v1.App.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Number of deployments in the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.deployment_count
	DeploymentCount *int32 `json:"deploymentCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesapp;gcpcesapps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESApp is the Schema for the CESApp API
// +k8s:openapi-gen=true
type CESApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESAppSpec   `json:"spec,omitempty"`
	Status CESAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESAppList contains a list of CESApp
type CESAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESApp{}, &CESAppList{})
}
