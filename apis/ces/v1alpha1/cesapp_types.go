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
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	dlpv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dlp/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESAppGVK = GroupVersion.WithKind("CESApp")

// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type Schema struct {
}

// +kcc:proto=google.cloud.ces.v1.App.VariableDeclaration
type App_VariableDeclaration struct {
	// Required. The name of the variable. The name must start with a letter or
	//  underscore and contain only letters, numbers, or underscores.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.name
	Name *string `json:"name,omitempty"`

	// Required. The description of the variable.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.description
	Description *string `json:"description,omitempty"`

	// Required. The schema of the variable.
	// +kcc:proto:field=google.cloud.ces.v1.App.VariableDeclaration.schema
	Schema *Schema `json:"schema,omitempty"`
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
	Schema *Schema `json:"schema,omitempty"`
}

// CESAppSpec defines the desired state of CESApp
// +kcc:spec:proto=google.cloud.ces.v1.App
type CESAppSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The CESApp name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the app.
	// +kcc:proto:field=google.cloud.ces.v1.App.display_name
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
	RootAgentRef *CesAgentRef `json:"rootAgentRef,omitempty"`

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
	GuardrailRefs []CesGuardrailRef `json:"guardrailRefs,omitempty"`

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

	// Optional. The data store settings for the app.
	DataStoreSettings *DataStoreSettings `json:"dataStoreSettings,omitempty"`

	// Optional. The evaluation personas for the app.
	EvaluationPersonas []EvaluationPersona `json:"evaluationPersonas,omitempty"`

	// Optional. The evaluation settings for the app.
	EvaluationSettings *EvaluationSettings `json:"evaluationSettings,omitempty"`
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

// +kcc:proto=google.cloud.ces.v1.AmbientSoundConfig
type AmbientSoundConfig struct {
	// Optional. Deprecated: `prebuilt_ambient_noise` is deprecated in favor of
	//  `prebuilt_ambient_sound`.
	// +kcc:proto:field=google.cloud.ces.v1.AmbientSoundConfig.prebuilt_ambient_noise
	PrebuiltAmbientNoise *string `json:"prebuiltAmbientNoise,omitempty"`

	// Optional. Ambient noise as a mono-channel, 16kHz WAV file stored in
	//  [Cloud Storage](https://cloud.google.com/storage).
	//
	//  Note: Please make sure the CES service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com` has
	//  `storage.objects.get` permission to the Cloud Storage object.
	// +kcc:proto:field=google.cloud.ces.v1.AmbientSoundConfig.gcs_uri
	GCSURI *string `json:"gcsURI,omitempty"`

	// Optional. Name of the prebuilt ambient sound.
	//  Valid values are:
	//  - "coffee_shop"
	//  - "keyboard"
	//  - "keypad"
	//  - "hum"
	//  - "office_1"
	//  - "office_2"
	//  - "office_3"
	//  - "room_1"
	//  - "room_2"
	//  - "room_3"
	//  - "room_4"
	//  - "room_5"
	//  - "air_conditioner"
	// +kcc:proto:field=google.cloud.ces.v1.AmbientSoundConfig.prebuilt_ambient_sound
	PrebuiltAmbientSound *string `json:"prebuiltAmbientSound,omitempty"`

	// Optional. Volume gain (in dB) of the normal native volume supported by
	//  ambient noise, in the range [-96.0, 16.0]. If unset, or set to a value of
	//  0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	//  will play at approximately half the amplitude of the normal native signal
	//  amplitude. A value of +6.0 (dB) will play at approximately twice the
	//  amplitude of the normal native signal amplitude. We strongly recommend not
	//  to exceed +10 (dB) as there's usually no effective increase in loudness for
	//  any value greater than that.
	// +kcc:proto:field=google.cloud.ces.v1.AmbientSoundConfig.volume_gain_db
	VolumeGainDb *float64 `json:"volumeGainDb,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.BargeInConfig
type BargeInConfig struct {
	// Optional. Deprecated: `disable_barge_in` is deprecated in favor of
	//  `barge_in_awareness`.
	//  Whether to disable user barge-in control in the conversation.
	//  - **true**: User interruptions are disabled while the agent is speaking.
	//  - **false**: User interruptions are allowed while the agent is speaking.
	// +kcc:proto:field=google.cloud.ces.v1.BargeInConfig.disable_barge_in
	DisableBargeIn *bool `json:"disableBargeIn,omitempty"`

	// Optional. Whether to disable user barge-in control in the conversation.
	// +kcc:proto:field=google.cloud.ces.v1.BargeInConfig.barge_in_awareness
	BargeInAwareness *bool `json:"bargeInAwareness,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ClientCertificateSettings
type ClientCertificateSettings struct {
	// Required. The TLS certificate encoded in PEM format. This string must
	//  include the begin header and end footer lines.
	// +kcc:proto:field=google.cloud.ces.v1.ClientCertificateSettings.tls_certificate
	TLSCertificate *string `json:"tlsCertificate,omitempty"`

	// Required. The name of the SecretManager secret version resource storing the
	//  private key encoded in PEM format.
	//  Format: `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.ces.v1.ClientCertificateSettings.private_key
	PrivateKeyRef *refsv1beta1.SecretManagerSecretVersionRef `json:"privateKeyRef,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing the
	//  passphrase to decrypt the private key. Should be left unset if the private
	//  key is not encrypted.
	//  Format: `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.ces.v1.ClientCertificateSettings.passphrase
	PassphraseRef *refsv1beta1.SecretManagerSecretVersionRef `json:"passphraseRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.EvaluationMetricsThresholds
type EvaluationMetricsThresholds struct {
	// Optional. The golden evaluation metrics thresholds.
	// +kcc:proto:field=google.cloud.ces.v1.EvaluationMetricsThresholds.golden_evaluation_metrics_thresholds
	GoldenEvaluationMetricsThresholds *EvaluationMetricsThresholds_GoldenEvaluationMetricsThresholds `json:"goldenEvaluationMetricsThresholds,omitempty"`

	// Optional. Deprecated: `hallucination_metric_behavior` is deprecated in
	//  favor of `golden_hallucination_metric_behavior` and
	//  `scenario_hallucination_metric_behavior`.
	//  Behavior if hallucination metric threshold is not met.
	// +kcc:proto:field=google.cloud.ces.v1.EvaluationMetricsThresholds.hallucination_metric_behavior
	HallucinationMetricBehavior *string `json:"hallucinationMetricBehavior,omitempty"`

	// Optional. The hallucination metric behavior for golden evaluations.
	// +kcc:proto:field=google.cloud.ces.v1.EvaluationMetricsThresholds.golden_hallucination_metric_behavior
	GoldenHallucinationMetricBehavior *string `json:"goldenHallucinationMetricBehavior,omitempty"`

	// Optional. The hallucination metric behavior for scenario evaluations.
	// +kcc:proto:field=google.cloud.ces.v1.EvaluationMetricsThresholds.scenario_hallucination_metric_behavior
	ScenarioHallucinationMetricBehavior *string `json:"scenarioHallucinationMetricBehavior,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.LanguageSettings
type LanguageSettings struct {
	// Required. The default language code of the app.
	// +kcc:proto:field=google.cloud.ces.v1.LanguageSettings.default_language_code
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`

	// Optional. The list of languages supported by the app.
	// +kcc:proto:field=google.cloud.ces.v1.LanguageSettings.supported_language_codes
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// Optional. Whether to enable multilingual support for the app.
	// +kcc:proto:field=google.cloud.ces.v1.LanguageSettings.enable_multilingual_support
	EnableMultilingualSupport *bool `json:"enableMultilingualSupport,omitempty"`

	// Optional. Deprecated: `fallback_action` is deprecated.
	//  Action to take when a user speaks in a language that is not supported.
	// +kcc:proto:field=google.cloud.ces.v1.LanguageSettings.fallback_action
	FallbackAction *string `json:"fallbackAction,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.AudioRecordingConfig
type AudioRecordingConfig struct {
	// Optional. The [Cloud Storage](https://cloud.google.com/storage) bucket to
	//  store the session audio recordings. The URI must start with "gs://".
	//
	//  Please choose a bucket location that meets your data residency
	//  requirements.
	//
	//  Note: If the Cloud Storage bucket is in a different project from the app,
	//  you should grant `storage.objects.create` permission to the CES service
	//  agent `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.AudioRecordingConfig.gcs_bucket
	GCSBucket *string `json:"gcsBucket,omitempty"`

	// Reference to the Cloud Storage bucket to store session audio recordings.
	GCSBucketRef *storagev1beta1.StorageBucketRef `json:"gcsBucketRef,omitempty"`

	// Optional. The Cloud Storage path prefix for audio recordings.
	//
	//  This prefix can include the following placeholders, which will be
	//  dynamically substituted at serving time:
	//    - $project:   project ID
	//    - $location:  app location
	//    - $app:       app ID
	//    - $date:      session date in YYYY-MM-DD format
	// +kcc:proto:field=google.cloud.ces.v1.AudioRecordingConfig.gcs_path_prefix
	GCSPathPrefix *string `json:"gcsPathPrefix,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.BigQueryExportSettings
type BigQueryExportSettings struct {
	// Optional. Indicates whether the BigQuery export is enabled.
	// +kcc:proto:field=google.cloud.ces.v1.BigQueryExportSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. The **project ID** of the BigQuery dataset to export the data to.
	//
	//  Note: If the BigQuery dataset is in a different project from the app,
	//  you should grant `roles/bigquery.admin` role to the CES service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.BigQueryExportSettings.project
	Project *string `json:"project,omitempty"`

	// Reference to the Project of the BigQuery dataset.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Optional. The BigQuery **dataset ID** to export the data to.
	// +kcc:proto:field=google.cloud.ces.v1.BigQueryExportSettings.dataset
	Dataset *string `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.RedactionConfig
type RedactionConfig struct {
	// Optional. Whether redaction is enabled.
	// +kcc:proto:field=google.cloud.ces.v1.RedactionConfig.enable_redaction
	EnableRedaction *bool `json:"enableRedaction,omitempty"`

	// Optional. [DLP](https://cloud.google.com/dlp/docs) inspect template name to
	//  configure detection of sensitive data types.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/inspectTemplates/{inspect_template}`
	// +kcc:proto:field=google.cloud.ces.v1.RedactionConfig.inspect_template
	InspectTemplateRef *dlpv1alpha1.DLPInspectTemplateRef `json:"inspectTemplateRef,omitempty"`

	// Optional. [DLP](https://cloud.google.com/dlp/docs) deidentify template name
	//  to instruct on how to de-identify content.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}`
	// +kcc:proto:field=google.cloud.ces.v1.RedactionConfig.deidentify_template
	DeidentifyTemplateRef *dlpv1alpha1.DLPDeidentifyTemplateRef `json:"deidentifyTemplateRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ErrorHandlingSettings
type ErrorHandlingSettings struct {
	// Optional. The strategy to use for error handling.
	// +kcc:proto:field=google.cloud.ces.v1.ErrorHandlingSettings.error_handling_strategy
	ErrorHandlingStrategy *string `json:"errorHandlingStrategy,omitempty"`

	EndSessionConfig       *ErrorHandlingSettings_EndSessionConfig       `json:"endSessionConfig,omitempty"`
	FallbackResponseConfig *ErrorHandlingSettings_FallbackResponseConfig `json:"fallbackResponseConfig,omitempty"`
}

type ErrorHandlingSettings_EndSessionConfig struct {
	EscalateSession *bool `json:"escalateSession,omitempty"`
}

type ErrorHandlingSettings_FallbackResponseConfig struct {
	CustomFallbackMessages map[string]string `json:"customFallbackMessages,omitempty"`
	MaxFallbackAttempts    *int32            `json:"maxFallbackAttempts,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ConversationLoggingSettings
type ConversationLoggingSettings struct {
	// Optional. Whether to disable conversation logging.
	// +kcc:proto:field=google.cloud.ces.v1.ConversationLoggingSettings.disable_conversation_logging
	DisableConversationLogging *bool `json:"disableConversationLogging,omitempty"`

	RetentionWindow *string `json:"retentionWindow,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.LoggingSettings
type LoggingSettings struct {
	// Optional. Configuration of how the audio should be recorded.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.audio_recording_config
	AudioRecordingConfig *AudioRecordingConfig `json:"audioRecordingConfig,omitempty"`

	// Optional. Configuration of how the data should be exported to BigQuery.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.bigquery_export_settings
	BigqueryExportSettings *BigQueryExportSettings `json:"bigqueryExportSettings,omitempty"`

	// Optional. Configuration of how the data should be logged to Cloud Logging.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.cloud_logging_settings
	CloudLoggingSettings *CloudLoggingSettings `json:"cloudLoggingSettings,omitempty"`

	// Optional. Configuration of how the conversation should be logged.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.conversation_logging_settings
	ConversationLoggingSettings *ConversationLoggingSettings `json:"conversationLoggingSettings,omitempty"`

	// Optional. Configuration of how the evaluation audio should be recorded.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.evaluation_audio_recording_config
	EvaluationAudioRecordingConfig *AudioRecordingConfig `json:"evaluationAudioRecordingConfig,omitempty"`

	// Optional. Configuration of how the metric analysis should be performed.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.metric_analysis_settings
	MetricAnalysisSettings *MetricAnalysisSettings `json:"metricAnalysisSettings,omitempty"`

	// Optional. Configuration of how the sensitive data should be redacted.
	// +kcc:proto:field=google.cloud.ces.v1.LoggingSettings.redaction_config
	RedactionConfig *RedactionConfig `json:"redactionConfig,omitempty"`

	// Optional. Configuration of unredacted audio recording.
	UnredactedAudioRecordingConfig *AudioRecordingConfig `json:"unredactedAudioRecordingConfig,omitempty"`
}

// +kubebuilder:validation:Schemaless
// +kubebuilder:validation:XPreserveUnknownFields
// +kubebuilder:validation:Type=object
type DataStoreSettings struct {
}

type CESEvaluationPersonaRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type EvaluationPersona struct {
	NameRef     *CESEvaluationPersonaRef `json:"nameRef,omitempty"`
	Description *string                  `json:"description,omitempty"`
	DisplayName *string                  `json:"displayName,omitempty"`
	Instruction *string                  `json:"instruction,omitempty"`
}

type EvaluationSettings struct {
	ScenarioConversationInitiator       *string                  `json:"scenarioConversationInitiator,omitempty"`
	GoldenRunMethod                     *string                  `json:"goldenRunMethod,omitempty"`
	GoldenEvaluationToolCallBehaviour   *string                  `json:"goldenEvaluationToolCallBehaviour,omitempty"`
	ScenarioEvaluationToolCallBehaviour *string                  `json:"scenarioEvaluationToolCallBehaviour,omitempty"`
	MetricsConfig                       *EvaluationMetricsConfig `json:"metricsConfig,omitempty"`
	ScenarioExecutionMode               *string                  `json:"scenarioExecutionMode,omitempty"`
}

type EvaluationMetricsConfig struct {
	GoldenMetricsConfig   *EvaluationMetricsConfig_GoldenMetricsConfig   `json:"goldenMetricsConfig,omitempty"`
	ScenarioMetricsConfig *EvaluationMetricsConfig_ScenarioMetricsConfig `json:"scenarioMetricsConfig,omitempty"`
}

type EvaluationMetricsConfig_GoldenMetricsConfig struct {
	SemanticSimilarityMetricsConfig  *EvaluationMetricsConfig_SemanticSimilarityMetricsConfig `json:"semanticSimilarityMetricsConfig,omitempty"`
	ToolCorrectnessMetricsConfig     *EvaluationMetricsConfig_ToolCorrectnessMetricsConfig    `json:"toolCorrectnessMetricsConfig,omitempty"`
	StepToolCorrectnessMetricsConfig *EvaluationMetricsConfig_ToolCorrectnessMetricsConfig    `json:"stepToolCorrectnessMetricsConfig,omitempty"`
}

type EvaluationMetricsConfig_ScenarioMetricsConfig struct {
	UserGoalMetMetricsConfig     *EvaluationMetricsConfig_UserGoalMetMetricsConfig     `json:"userGoalMetMetricsConfig,omitempty"`
	ExpectationsMetMetricsConfig *EvaluationMetricsConfig_ExpectationsMetMetricsConfig `json:"expectationsMetMetricsConfig,omitempty"`
}

type EvaluationMetricsConfig_SemanticSimilarityMetricsConfig struct {
	EnableSemanticSimilarityMetrics *bool `json:"enableSemanticSimilarityMetrics,omitempty"`
}

type EvaluationMetricsConfig_ToolCorrectnessMetricsConfig struct {
	EnableToolCorrectnessMetrics *bool `json:"enableToolCorrectnessMetrics,omitempty"`
}

type EvaluationMetricsConfig_UserGoalMetMetricsConfig struct {
	EnableUserGoalMetMetrics *bool `json:"enableUserGoalMetMetrics,omitempty"`
}

type EvaluationMetricsConfig_ExpectationsMetMetricsConfig struct {
	EnableExpectationsMetMetrics *bool `json:"enableExpectationsMetMetrics,omitempty"`
}
