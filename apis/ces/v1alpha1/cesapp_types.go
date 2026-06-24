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
	dlpv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dlp/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESAppGVK = GroupVersion.WithKind("CESApp")

var _ = apiextensionsv1.JSON{}

// CESGuardrailRef defines the resource reference to CESGuardrail.
type CESGuardrailRef struct {
	// A reference to an externally managed CES Guardrail.
	// Format: `projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}`
	External string `json:"external,omitempty"`

	// The `name` of a `CESGuardrail` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `CESGuardrail` resource.
	Namespace string `json:"namespace,omitempty"`
}

// CESAgentRef defines the resource reference to CESAgent.
type CESAgentRef struct {
	// A reference to an externally managed CES Agent.
	// Format: `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	External string `json:"external,omitempty"`

	// The `name` of a `CESAgent` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `CESAgent` resource.
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1beta.DataStoreSettings
// +kubebuilder:validation:XPreserveUnknownFields
type DataStoreSettings struct {
}

// +kcc:observedstate:proto=google.cloud.ces.v1beta.DataStoreSettings
type DataStoreSettingsObservedState struct {
	// Output only. The engines for the app.
	// +kcc:proto:field=google.cloud.ces.v1beta.DataStoreSettings.engines
	Engines []DataStoreSettings_EngineObservedState `json:"engines,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.ces.v1beta.DataStoreSettings.Engine
type DataStoreSettings_EngineObservedState struct {
	// Output only. The resource name of the engine.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}`
	// +kcc:proto:field=google.cloud.ces.v1beta.DataStoreSettings.Engine.name
	Name *string `json:"name,omitempty"`

	// Output only. The type of the engine.
	// +kcc:proto:field=google.cloud.ces.v1beta.DataStoreSettings.Engine.type
	Type *string `json:"type,omitempty"`
}

// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type Schema struct {
}

// CESAppSpec defines the desired state of CESApp
// +kcc:spec:proto=google.cloud.ces.v1beta.App
type CESAppSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CESApp name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Display name of the app.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Human-readable description of the app.
	Description *string `json:"description,omitempty"`

	// Optional. Whether the app is pinned in the app list.
	Pinned *bool `json:"pinned,omitempty"`

	// Optional. The root agent is the entry point of the app.
	//  Format: `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	RootAgentRef *CESAgentRef `json:"rootAgentRef,omitempty"`

	// Optional. Language settings of the app.
	LanguageSettings *LanguageSettings `json:"languageSettings,omitempty"`

	// Optional. TimeZone settings of the app.
	TimeZoneSettings *TimeZoneSettings `json:"timeZoneSettings,omitempty"`

	// Optional. Audio processing configuration of the app.
	AudioProcessingConfig *AudioProcessingConfig `json:"audioProcessingConfig,omitempty"`

	// Optional. Logging settings of the app.
	LoggingSettings *LoggingSettings `json:"loggingSettings,omitempty"`

	// Optional. Error handling settings of the app.
	ErrorHandlingSettings *ErrorHandlingSettings `json:"errorHandlingSettings,omitempty"`

	// Optional. The default LLM model settings for the app.
	//  Individual resources (e.g. agents, guardrails) can override these
	//  configurations as needed.
	ModelSettings *ModelSettings `json:"modelSettings,omitempty"`

	// Optional. The tool execution mode for the app. If not provided, will
	//  default to PARALLEL.
	ToolExecutionMode *string `json:"toolExecutionMode,omitempty"`

	// Optional. The evaluation thresholds for the app.
	EvaluationMetricsThresholds *EvaluationMetricsThresholds `json:"evaluationMetricsThresholds,omitempty"`

	// Optional. The declarations of the variables.
	VariableDeclarations []App_VariableDeclaration `json:"variableDeclarations,omitempty"`

	// Optional. Instructions for all the agents in the app.
	//  You can use this instruction to set up a stable identity or personality
	//  across all the agents.
	GlobalInstruction *string `json:"globalInstruction,omitempty"`

	// Optional. List of guardrails for the app.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}`
	GuardrailRefs []CESGuardrailRef `json:"guardrailRefs,omitempty"`

	// Optional. The data store settings for the app.
	DataStoreSettings *DataStoreSettings `json:"dataStoreSettings,omitempty"`

	// Optional. The default channel profile used by the app.
	DefaultChannelProfile *ChannelProfile `json:"defaultChannelProfile,omitempty"`

	// Optional. Metadata about the app. This field can be used to store
	//  additional information relevant to the app's details or intended usages.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. The default client certificate settings for the app.
	ClientCertificateSettings *ClientCertificateSettings `json:"clientCertificateSettings,omitempty"`

	// Optional. Indicates whether the app is locked for changes. If the app is
	//  locked, modifications to the app resources will be rejected.
	Locked *bool `json:"locked,omitempty"`

	// Optional. The evaluation personas for the app. This field is used to define
	//  the personas that can be used for evaluation. Maximum of 30 personas can be
	//  defined.
	EvaluationPersonas []EvaluationPersona `json:"evaluationPersonas,omitempty"`

	// Optional. The evaluation settings for the app.
	EvaluationSettings *EvaluationSettings `json:"evaluationSettings,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1beta.App.VariableDeclaration
type App_VariableDeclaration struct {
	// Required. The name of the variable. The name must start with a letter or
	//  underscore and contain only letters, numbers, or underscores.
	// +kcc:proto:field=google.cloud.ces.v1beta.App.VariableDeclaration.name
	Name *string `json:"name,omitempty"`

	// Required. The description of the variable.
	// +kcc:proto:field=google.cloud.ces.v1beta.App.VariableDeclaration.description
	Description *string `json:"description,omitempty"`

	// Required. The schema of the variable.
	// +kcc:proto:field=google.cloud.ces.v1beta.App.VariableDeclaration.schema
	Schema *Schema `json:"schema,omitempty"`
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
// +kcc:observedstate:proto=google.cloud.ces.v1beta.App
type CESAppObservedState struct {
	// Output only. The declarations of predefined variables for the app.
	PredefinedVariableDeclarations []App_VariableDeclaration `json:"predefinedVariableDeclarations,omitempty"`

	// Optional. The data store settings for the app.
	DataStoreSettings *DataStoreSettingsObservedState `json:"dataStoreSettings,omitempty"`

	// Output only. Timestamp when the app was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the app was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Etag used to ensure the object hasn't changed during a
	//  read-modify-write operation. If the etag is empty, the update will
	//  overwrite any concurrent changes.
	Etag *string `json:"etag,omitempty"`

	// Output only. Number of deployments in the app.
	DeploymentCount *int32 `json:"deploymentCount,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesapp;gcpcesapps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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

// CESEvaluationPersonaRef defines the resource reference to CESEvaluationPersona.
type CESEvaluationPersonaRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1beta.ClientCertificateSettings
type ClientCertificateSettings struct {
	TLSCertificate *string                                    `json:"tlsCertificate,omitempty"`
	PrivateKeyRef  *refsv1beta1.SecretManagerSecretVersionRef `json:"privateKeyRef,omitempty"`
	PassphraseRef  *refsv1beta1.SecretManagerSecretVersionRef `json:"passphraseRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1beta.EvaluationPersona
type EvaluationPersona struct {
	NameRef     *CESEvaluationPersonaRef `json:"nameRef,omitempty"`
	Description *string                  `json:"description,omitempty"`
	DisplayName *string                  `json:"displayName,omitempty"`
	Instruction *string                  `json:"instruction,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1beta.RedactionConfig
type RedactionConfig struct {
	EnableRedaction       *bool                                 `json:"enableRedaction,omitempty"`
	InspectTemplateRef    *dlpv1alpha1.DLPInspectTemplateRef    `json:"inspectTemplateRef,omitempty"`
	DeidentifyTemplateRef *dlpv1alpha1.DLPDeidentifyTemplateRef `json:"deidentifyTemplateRef,omitempty"`
}
