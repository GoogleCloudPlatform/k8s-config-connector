// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.dialogflow.cx.v3.SecuritySettings
type SecuritySettings struct {
	// Resource name of the settings.
	//  Required for the
	//  [SecuritySettingsService.UpdateSecuritySettings][google.cloud.dialogflow.cx.v3.SecuritySettingsService.UpdateSecuritySettings]
	//  method.
	//  [SecuritySettingsService.CreateSecuritySettings][google.cloud.dialogflow.cx.v3.SecuritySettingsService.CreateSecuritySettings]
	//  populates the name automatically. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/securitySettings/<SecuritySettingsID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the security settings, unique within
	//  the location.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Strategy that defines how we do redaction.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.redaction_strategy
	RedactionStrategy *string `json:"redactionStrategy,omitempty"`

	// Defines the data for which Dialogflow applies redaction. Dialogflow does
	//  not redact data that it does not have access to – for example, Cloud
	//  logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.redaction_scope
	RedactionScope *string `json:"redactionScope,omitempty"`

	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this
	//  template to define inspect base settings.
	//
	//  The `DLP Inspect Templates Reader` role is needed on the Dialogflow
	//  service identity service account (has the form
	//  `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`)
	//  for your agent's project.
	//
	//  If empty, we use the default DLP inspect config.
	//
	//  The template name will have one of the following formats:
	//  `projects/<ProjectID>/locations/<LocationID>/inspectTemplates/<TemplateID>`
	//  OR
	//  `organizations/<OrganizationID>/locations/<LocationID>/inspectTemplates/<TemplateID>`
	//
	//  Note: `inspect_template` must be located in the same region as the
	//  `SecuritySettings`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.inspect_template
	InspectTemplate *string `json:"inspectTemplate,omitempty"`

	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this
	//  template to define de-identification configuration for the content.
	//
	//  The `DLP De-identify Templates Reader` role is needed on the Dialogflow
	//  service identity service account (has the form
	//  `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`)
	//  for your agent's project.
	//
	//  If empty, Dialogflow replaces sensitive info with `[redacted]` text.
	//
	//  The template name will have one of the following formats:
	//  `projects/<ProjectID>/locations/<LocationID>/deidentifyTemplates/<TemplateID>`
	//  OR
	//  `organizations/<OrganizationID>/locations/<LocationID>/deidentifyTemplates/<TemplateID>`
	//
	//  Note: `deidentify_template` must be located in the same region as the
	//  `SecuritySettings`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.deidentify_template
	DeidentifyTemplate *string `json:"deidentifyTemplate,omitempty"`

	// Retains the data for the specified number of days.
	//  User must set a value lower than Dialogflow's default 365d TTL (30 days
	//  for Agent Assist traffic), higher value will be ignored and use default.
	//  Setting a value higher than that has no effect. A missing value or
	//  setting to 0 also means we use default TTL.
	//  When data retention configuration is changed, it only applies to the data
	//  created after the change; the TTL of existing data created before the
	//  change stays intact.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.retention_window_days
	RetentionWindowDays *int32 `json:"retentionWindowDays,omitempty"`

	// Specifies the retention behavior defined by
	//  [SecuritySettings.RetentionStrategy][google.cloud.dialogflow.cx.v3.SecuritySettings.RetentionStrategy].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.retention_strategy
	RetentionStrategy *string `json:"retentionStrategy,omitempty"`

	// List of types of data to remove when retention settings triggers purge.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.purge_data_types
	PurgeDataTypes []string `json:"purgeDataTypes,omitempty"`

	// Controls audio export settings for post-conversation analytics when
	//  ingesting audio to conversations via [Participants.AnalyzeContent][] or
	//  [Participants.StreamingAnalyzeContent][].
	//
	//  If
	//  [retention_strategy][google.cloud.dialogflow.cx.v3.SecuritySettings.retention_strategy]
	//  is set to REMOVE_AFTER_CONVERSATION or [audio_export_settings.gcs_bucket][]
	//  is empty, audio export is disabled.
	//
	//  If audio export is enabled, audio is recorded and saved to
	//  [audio_export_settings.gcs_bucket][], subject to retention policy of
	//  [audio_export_settings.gcs_bucket][].
	//
	//  This setting won't effect audio input for implicit sessions via
	//  [Sessions.DetectIntent][google.cloud.dialogflow.cx.v3.Sessions.DetectIntent]
	//  or
	//  [Sessions.StreamingDetectIntent][google.cloud.dialogflow.cx.v3.Sessions.StreamingDetectIntent].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.audio_export_settings
	AudioExportSettings *SecuritySettings_AudioExportSettings `json:"audioExportSettings,omitempty"`

	// Controls conversation exporting settings to Insights after conversation is
	//  completed.
	//
	//  If
	//  [retention_strategy][google.cloud.dialogflow.cx.v3.SecuritySettings.retention_strategy]
	//  is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter
	//  what you configure here.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.insights_export_settings
	InsightsExportSettings *SecuritySettings_InsightsExportSettings `json:"insightsExportSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings
type SecuritySettings_AudioExportSettings struct {
	// Cloud Storage bucket to export audio record to.
	//  Setting this field would grant the Storage Object Creator role to
	//  the Dialogflow Service Agent.
	//  API caller that tries to modify this field should have the permission of
	//  storage.buckets.setIamPolicy.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings.gcs_bucket
	GcsBucket *string `json:"gcsBucket,omitempty"`

	// Filename pattern for exported audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings.audio_export_pattern
	AudioExportPattern *string `json:"audioExportPattern,omitempty"`

	// Enable audio redaction if it is true.
	//  Note that this only redacts end-user audio data;
	//  Synthesised audio from the virtual agent is not redacted.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings.enable_audio_redaction
	EnableAudioRedaction *bool `json:"enableAudioRedaction,omitempty"`

	// File format for exported audio file. Currently only in telephony
	//  recordings.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings.audio_format
	AudioFormat *string `json:"audioFormat,omitempty"`

	// Whether to store TTS audio. By default, TTS audio from the virtual agent
	//  is not exported.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings.store_tts_audio
	StoreTtsAudio *bool `json:"storeTtsAudio,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.SecuritySettings.InsightsExportSettings
type SecuritySettings_InsightsExportSettings struct {
	// If enabled, we will automatically exports
	//  conversations to Insights and Insights runs its analyzers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SecuritySettings.InsightsExportSettings.enable_insights_export
	EnableInsightsExport *bool `json:"enableInsightsExport,omitempty"`
}
