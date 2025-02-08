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


// +kcc:proto=google.cloud.dialogflow.v2beta1.Environment
type Environment struct {

	// Optional. The developer-provided description for this environment.
	//  The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.description
	Description *string `json:"description,omitempty"`

	// Optional. The agent version loaded into this environment.
	//  Supported formats:
	//  - `projects/<Project ID>/agent/versions/<Version ID>`
	//  - `projects/<Project ID>/locations/<Location ID>/agent/versions/<Version
	//  ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.agent_version
	AgentVersion *string `json:"agentVersion,omitempty"`

	// Optional. Text to speech settings for this environment.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.text_to_speech_settings
	TextToSpeechSettings *TextToSpeechSettings `json:"textToSpeechSettings,omitempty"`

	// Optional. The fulfillment settings to use for this environment.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.fulfillment
	Fulfillment *Fulfillment `json:"fulfillment,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Fulfillment
type Fulfillment struct {
	// Required. The unique identifier of the fulfillment.
	//  Supported formats:
	//
	//  - `projects/<Project ID>/agent/fulfillment`
	//  - `projects/<Project ID>/locations/<Location ID>/agent/fulfillment`
	//
	//  This field is not used for Fulfillment in an Environment.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.name
	Name *string `json:"name,omitempty"`

	// The human-readable name of the fulfillment, unique within the agent.
	//
	//  This field is not used for Fulfillment in an Environment.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for a generic web service.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.generic_web_service
	GenericWebService *Fulfillment_GenericWebService `json:"genericWebService,omitempty"`

	// Whether fulfillment is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The field defines whether the fulfillment is enabled for certain features.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.features
	Features []Fulfillment_Feature `json:"features,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Fulfillment.Feature
type Fulfillment_Feature struct {
	// The type of the feature that enabled for fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.Feature.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Fulfillment.GenericWebService
type Fulfillment_GenericWebService struct {
	// Required. The fulfillment URI for receiving POST requests.
	//  It must use https protocol.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.GenericWebService.uri
	URI *string `json:"uri,omitempty"`

	// The user name for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.GenericWebService.username
	Username *string `json:"username,omitempty"`

	// The password for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.GenericWebService.password
	Password *string `json:"password,omitempty"`

	// The HTTP request headers to send together with fulfillment requests.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.GenericWebService.request_headers
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// Optional. Indicates if generic web service is created through Cloud
	//  Functions integration. Defaults to false.
	//
	//  is_cloud_function is deprecated. Cloud functions can be configured by
	//  its uri as a regular web service now.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Fulfillment.GenericWebService.is_cloud_function
	IsCloudFunction *bool `json:"isCloudFunction,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig
type SynthesizeSpeechConfig struct {
	// Optional. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal
	//  native speed supported by the specific voice. 2.0 is twice as fast, and 0.5
	//  is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any other
	//  values < 0.25 or > 4.0 will return an error.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.speaking_rate
	SpeakingRate *float64 `json:"speakingRate,omitempty"`

	// Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20
	//  semitones from the original pitch. -20 means decrease 20 semitones from the
	//  original pitch.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.pitch
	Pitch *float64 `json:"pitch,omitempty"`

	// Optional. Volume gain (in dB) of the normal native volume supported by the
	//  specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of
	//  0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	//  will play at approximately half the amplitude of the normal native signal
	//  amplitude. A value of +6.0 (dB) will play at approximately twice the
	//  amplitude of the normal native signal amplitude. We strongly recommend not
	//  to exceed +10 (dB) as there's usually no effective increase in loudness for
	//  any value greater than that.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.volume_gain_db
	VolumeGainDb *float64 `json:"volumeGainDb,omitempty"`

	// Optional. An identifier which selects 'audio effects' profiles that are
	//  applied on (post synthesized) text to speech. Effects are applied on top of
	//  each other in the order they are given.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.effects_profile_id
	EffectsProfileID []string `json:"effectsProfileID,omitempty"`

	// Optional. The desired voice of the synthesized audio.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SynthesizeSpeechConfig.voice
	Voice *VoiceSelectionParams `json:"voice,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.TextToSpeechSettings
type TextToSpeechSettings struct {
	// Optional. Indicates whether text to speech is enabled. Even when this field
	//  is false, other settings in this proto are still retained.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.TextToSpeechSettings.enable_text_to_speech
	EnableTextToSpeech *bool `json:"enableTextToSpeech,omitempty"`

	// Required. Audio encoding of the synthesized audio content.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.TextToSpeechSettings.output_audio_encoding
	OutputAudioEncoding *string `json:"outputAudioEncoding,omitempty"`

	// Optional. The synthesis sample rate (in hertz) for this audio. If not
	//  provided, then the synthesizer will use the default sample rate based on
	//  the audio encoding. If this is different from the voice's natural sample
	//  rate, then the synthesizer will honor this request by converting to the
	//  desired sample rate (which might result in worse audio quality).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.TextToSpeechSettings.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.dialogflow.v2beta1.VoiceSelectionParams
type VoiceSelectionParams struct {
	// Optional. The name of the voice. If not set, the service will choose a
	//  voice based on the other parameters such as language_code and
	//  [ssml_gender][google.cloud.dialogflow.v2beta1.VoiceSelectionParams.ssml_gender].
	//
	//  For the list of available voices, please refer to [Supported voices and
	//  languages](https://cloud.google.com/text-to-speech/docs/voices).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.VoiceSelectionParams.name
	Name *string `json:"name,omitempty"`

	// Optional. The preferred gender of the voice. If not set, the service will
	//  choose a voice based on the other parameters such as language_code and
	//  [name][google.cloud.dialogflow.v2beta1.VoiceSelectionParams.name]. Note
	//  that this is only a preference, not requirement. If a voice of the
	//  appropriate gender is not available, the synthesizer should substitute a
	//  voice with a different gender rather than failing the request.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.VoiceSelectionParams.ssml_gender
	SsmlGender *string `json:"ssmlGender,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Environment
type EnvironmentObservedState struct {
	// Output only. The unique identifier of this agent environment.
	//  Supported formats:
	//  - `projects/<Project ID>/agent/environments/<Environment ID>`
	//  - `projects/<Project ID>/locations/<Location
	//  ID>/agent/environments/<Environment ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of this environment. This field is read-only, i.e.,
	//  it cannot be set by create and update methods.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.state
	State *string `json:"state,omitempty"`

	// Output only. The last update time of this environment. This field is
	//  read-only, i.e., it cannot be set by create and update methods.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Environment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
