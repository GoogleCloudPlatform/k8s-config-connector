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


// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings
type AdvancedSettings struct {
	// If present, incoming audio is exported by Dialogflow to the configured
	//  Google Cloud Storage destination.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.audio_export_gcs_destination
	AudioExportGcsDestination *GcsDestination `json:"audioExportGcsDestination,omitempty"`

	// Settings for speech to text detection.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	//  - Page level
	//  - Parameter level
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.speech_settings
	SpeechSettings *AdvancedSettings_SpeechSettings `json:"speechSettings,omitempty"`

	// Settings for DTMF.
	//  Exposed at the following levels:
	//  - Agent level
	//  - Flow level
	//  - Page level
	//  - Parameter level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.dtmf_settings
	DtmfSettings *AdvancedSettings_DtmfSettings `json:"dtmfSettings,omitempty"`

	// Settings for logging.
	//  Settings for Dialogflow History, Contact Center messages, StackDriver logs,
	//  and speech logging.
	//  Exposed at the following levels:
	//  - Agent level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.logging_settings
	LoggingSettings *AdvancedSettings_LoggingSettings `json:"loggingSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings
type AdvancedSettings_DtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequency)
	//  events. For example, if the caller presses a button on their telephone
	//  keypad and DTMF processing is enabled, Dialogflow will detect the
	//  event (e.g. a "3" was pressed) in the incoming audio and pass the event
	//  to the bot to drive business logic (e.g. when 3 is pressed, return the
	//  account balance).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Max length of DTMF digits.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.max_digits
	MaxDigits *int32 `json:"maxDigits,omitempty"`

	// The digit that terminates a DTMF digit sequence.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.finish_digit
	FinishDigit *string `json:"finishDigit,omitempty"`

	// Interdigit timeout setting for matching dtmf input to regex.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.interdigit_timeout_duration
	InterdigitTimeoutDuration *string `json:"interdigitTimeoutDuration,omitempty"`

	// Endpoint timeout setting for matching dtmf input to regex.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.DtmfSettings.endpointing_timeout_duration
	EndpointingTimeoutDuration *string `json:"endpointingTimeoutDuration,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings
type AdvancedSettings_LoggingSettings struct {
	// Enables Google Cloud Logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Enables DF Interaction logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings.enable_interaction_logging
	EnableInteractionLogging *bool `json:"enableInteractionLogging,omitempty"`

	// Enables consent-based end-user input redaction, if true, a pre-defined
	//  session parameter `$session.params.conversation-redaction` will be
	//  used to determine if the utterance should be redacted.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings.enable_consent_based_redaction
	EnableConsentBasedRedaction *bool `json:"enableConsentBasedRedaction,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings
type AdvancedSettings_SpeechSettings struct {
	// Sensitivity of the speech model that detects the end of speech.
	//  Scale from 0 to 100.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.endpointer_sensitivity
	EndpointerSensitivity *int32 `json:"endpointerSensitivity,omitempty"`

	// Timeout before detecting no speech.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.no_speech_timeout
	NoSpeechTimeout *string `json:"noSpeechTimeout,omitempty"`

	// Use timeout based endpointing, interpreting endpointer sensitivy as
	//  seconds of timeout value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.use_timeout_based_endpointing
	UseTimeoutBasedEndpointing *bool `json:"useTimeoutBasedEndpointing,omitempty"`

	// Mapping from language to Speech-to-Text model. The mapped Speech-to-Text
	//  model will be selected for requests from its corresponding language.
	//  For more information, see
	//  [Speech
	//  models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AdvancedSettings.SpeechSettings.models
	Models map[string]string `json:"models,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent
type Agent struct {
	// The unique identifier of the agent.
	//  Required for the
	//  [Agents.UpdateAgent][google.cloud.dialogflow.cx.v3.Agents.UpdateAgent]
	//  method.
	//  [Agents.CreateAgent][google.cloud.dialogflow.cx.v3.Agents.CreateAgent]
	//  populates the name automatically.
	//  Format: `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the agent, unique within the location.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The default language of the agent as a language tag.
	//  See [Language
	//  Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	//  for a list of the currently supported language codes.
	//  This field cannot be set by the
	//  [Agents.UpdateAgent][google.cloud.dialogflow.cx.v3.Agents.UpdateAgent]
	//  method.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.default_language_code
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`

	// The list of all languages supported by the agent (except for the
	//  `default_language_code`).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.supported_language_codes
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// Required. The time zone of the agent from the [time zone
	//  database](https://www.iana.org/time-zones), e.g., America/New_York,
	//  Europe/Paris.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// The description of the agent. The maximum length is 500 characters. If
	//  exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.description
	Description *string `json:"description,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow
	//  console and in the self-hosted [Web
	//  Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo)
	//  integration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.avatar_uri
	AvatarURI *string `json:"avatarURI,omitempty"`

	// Speech recognition related settings.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.speech_to_text_settings
	SpeechToTextSettings *SpeechToTextSettings `json:"speechToTextSettings,omitempty"`

	// Immutable. Name of the start flow in this agent. A start flow will be
	//  automatically created when the agent is created, and can only be deleted by
	//  deleting the agent. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.start_flow
	StartFlow *string `json:"startFlow,omitempty"`

	// Name of the
	//  [SecuritySettings][google.cloud.dialogflow.cx.v3.SecuritySettings]
	//  reference for the agent. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/securitySettings/<SecuritySettingsID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.security_settings
	SecuritySettings *string `json:"securitySettings,omitempty"`

	// Indicates if stackdriver logging is enabled for the agent.
	//  Please use
	//  [agent.advanced_settings][google.cloud.dialogflow.cx.v3.AdvancedSettings.LoggingSettings]
	//  instead.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.enable_stackdriver_logging
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent
	//  requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.enable_spell_correction
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty"`

	// Optional. Enable training multi-lingual models for this agent. These models
	//  will be trained on all the languages supported by the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.enable_multi_language_training
	EnableMultiLanguageTraining *bool `json:"enableMultiLanguageTraining,omitempty"`

	// Indicates whether the agent is locked for changes. If the agent is locked,
	//  modifications to the agent will be rejected except for [RestoreAgent][].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.locked
	Locked *bool `json:"locked,omitempty"`

	// Hierarchical advanced settings for this agent. The settings exposed at the
	//  lower level overrides the settings exposed at the higher level.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.advanced_settings
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`

	// Git integration settings for this agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.git_integration_settings
	GitIntegrationSettings *Agent_GitIntegrationSettings `json:"gitIntegrationSettings,omitempty"`

	// Settings on instructing the speech synthesizer on how to generate the
	//  output audio content.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.text_to_speech_settings
	TextToSpeechSettings *TextToSpeechSettings `json:"textToSpeechSettings,omitempty"`

	// Gen App Builder-related agent-level settings.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.gen_app_builder_settings
	GenAppBuilderSettings *Agent_GenAppBuilderSettings `json:"genAppBuilderSettings,omitempty"`

	// Optional. Answer feedback collection settings.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.answer_feedback_settings
	AnswerFeedbackSettings *Agent_AnswerFeedbackSettings `json:"answerFeedbackSettings,omitempty"`

	// Optional. Settings for end user personalization.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.personalization_settings
	PersonalizationSettings *Agent_PersonalizationSettings `json:"personalizationSettings,omitempty"`

	// Optional. Settings for custom client certificates.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.client_certificate_settings
	ClientCertificateSettings *Agent_ClientCertificateSettings `json:"clientCertificateSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent.AnswerFeedbackSettings
type Agent_AnswerFeedbackSettings struct {
	// Optional. If enabled, end users will be able to provide
	//  [answer feedback][google.cloud.dialogflow.cx.v3.AnswerFeedback] to
	//  Dialogflow responses. Feature works only if interaction logging is
	//  enabled in the Dialogflow agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.AnswerFeedbackSettings.enable_answer_feedback
	EnableAnswerFeedback *bool `json:"enableAnswerFeedback,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent.ClientCertificateSettings
type Agent_ClientCertificateSettings struct {
	// Required. The ssl certificate encoded in PEM format. This string must
	//  include the begin header and end footer lines.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.ClientCertificateSettings.ssl_certificate
	SslCertificate *string `json:"sslCertificate,omitempty"`

	// Required. The name of the SecretManager secret version resource storing
	//  the private key encoded in PEM format. Format:
	//  `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.ClientCertificateSettings.private_key
	PrivateKey *string `json:"privateKey,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the passphrase. 'passphrase' should be left unset if the private key is
	//  not encrypted.
	//  Format: `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.ClientCertificateSettings.passphrase
	Passphrase *string `json:"passphrase,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent.GenAppBuilderSettings
type Agent_GenAppBuilderSettings struct {
	// Required. The full name of the Gen App Builder engine related to this
	//  agent if there is one. Format: `projects/{Project ID}/locations/{Location
	//  ID}/collections/{Collection ID}/engines/{Engine ID}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GenAppBuilderSettings.engine
	Engine *string `json:"engine,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings
type Agent_GitIntegrationSettings struct {
	// GitHub settings.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.github_settings
	GithubSettings *Agent_GitIntegrationSettings_GithubSettings `json:"githubSettings,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.GithubSettings
type Agent_GitIntegrationSettings_GithubSettings struct {
	// The unique repository display name for the GitHub repository.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.GithubSettings.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The GitHub repository URI related to the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.GithubSettings.repository_uri
	RepositoryURI *string `json:"repositoryURI,omitempty"`

	// The branch of the GitHub repository tracked for this agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.GithubSettings.tracking_branch
	TrackingBranch *string `json:"trackingBranch,omitempty"`

	// The access token used to authenticate the access to the GitHub
	//  repository.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.GithubSettings.access_token
	AccessToken *string `json:"accessToken,omitempty"`

	// A list of branches configured to be used from Dialogflow.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.GitIntegrationSettings.GithubSettings.branches
	Branches []string `json:"branches,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Agent.PersonalizationSettings
type Agent_PersonalizationSettings struct {
	// Optional. Default end user metadata, used when processing DetectIntent
	//  requests. Recommended to be filled as a template instead of hard-coded
	//  value, for example { "age": "$session.params.age" }. The data will be
	//  merged with the
	//  [QueryParameters.end_user_metadata][google.cloud.dialogflow.cx.v3.QueryParameters.end_user_metadata]
	//  in
	//  [DetectIntentRequest.query_params][google.cloud.dialogflow.cx.v3.DetectIntentRequest.query_params]
	//  during query processing.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Agent.PersonalizationSettings.default_end_user_metadata
	DefaultEndUserMetadata map[string]string `json:"defaultEndUserMetadata,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.GcsDestination
type GcsDestination struct {
	// Required. The Google Cloud Storage URI for the exported objects. A URI is
	//  of the form: `gs://bucket/object-name-or-prefix` Whether a full object
	//  name, or just a prefix, its usage depends on the Dialogflow operation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.GcsDestination.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.SpeechToTextSettings
type SpeechToTextSettings struct {
	// Whether to use speech adaptation for speech recognition.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SpeechToTextSettings.enable_speech_adaptation
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.SynthesizeSpeechConfig
type SynthesizeSpeechConfig struct {
	// Optional. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal
	//  native speed supported by the specific voice. 2.0 is twice as fast, and
	//  0.5 is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any
	//  other values < 0.25 or > 4.0 will return an error.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SynthesizeSpeechConfig.speaking_rate
	SpeakingRate *float64 `json:"speakingRate,omitempty"`

	// Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20
	//  semitones from the original pitch. -20 means decrease 20 semitones from the
	//  original pitch.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SynthesizeSpeechConfig.pitch
	Pitch *float64 `json:"pitch,omitempty"`

	// Optional. Volume gain (in dB) of the normal native volume supported by the
	//  specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of
	//  0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	//  will play at approximately half the amplitude of the normal native signal
	//  amplitude. A value of +6.0 (dB) will play at approximately twice the
	//  amplitude of the normal native signal amplitude. We strongly recommend not
	//  to exceed +10 (dB) as there's usually no effective increase in loudness for
	//  any value greater than that.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SynthesizeSpeechConfig.volume_gain_db
	VolumeGainDb *float64 `json:"volumeGainDb,omitempty"`

	// Optional. An identifier which selects 'audio effects' profiles that are
	//  applied on (post synthesized) text to speech. Effects are applied on top of
	//  each other in the order they are given.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SynthesizeSpeechConfig.effects_profile_id
	EffectsProfileID []string `json:"effectsProfileID,omitempty"`

	// Optional. The desired voice of the synthesized audio.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SynthesizeSpeechConfig.voice
	Voice *VoiceSelectionParams `json:"voice,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TextToSpeechSettings
type TextToSpeechSettings struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.dialogflow.cx.v3.VoiceSelectionParams
type VoiceSelectionParams struct {
	// Optional. The name of the voice. If not set, the service will choose a
	//  voice based on the other parameters such as language_code and
	//  [ssml_gender][google.cloud.dialogflow.cx.v3.VoiceSelectionParams.ssml_gender].
	//
	//  For the list of available voices, please refer to [Supported voices and
	//  languages](https://cloud.google.com/text-to-speech/docs/voices).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.VoiceSelectionParams.name
	Name *string `json:"name,omitempty"`

	// Optional. The preferred gender of the voice. If not set, the service will
	//  choose a voice based on the other parameters such as language_code and
	//  [name][google.cloud.dialogflow.cx.v3.VoiceSelectionParams.name]. Note that
	//  this is only a preference, not requirement. If a voice of the appropriate
	//  gender is not available, the synthesizer substitutes a voice with a
	//  different gender rather than failing the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.VoiceSelectionParams.ssml_gender
	SsmlGender *string `json:"ssmlGender,omitempty"`
}
