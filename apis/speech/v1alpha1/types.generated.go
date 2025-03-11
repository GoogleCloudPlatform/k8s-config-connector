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

// +kcc:proto=google.cloud.speech.v2.AutoDetectDecodingConfig
type AutoDetectDecodingConfig struct {
}

// +kcc:proto=google.cloud.speech.v2.CustomClass
type CustomClass struct {

	// Optional. User-settable, human-readable name for the CustomClass. Must be
	//  63 characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A collection of class items.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.items
	Items []CustomClass_ClassItem `json:"items,omitempty"`

	// Optional. Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.CustomClass.ClassItem
type CustomClass_ClassItem struct {
	// The class item's value.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.ClassItem.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.ExplicitDecodingConfig
type ExplicitDecodingConfig struct {
	// Required. Encoding of the audio data sent for recognition.
	// +kcc:proto:field=google.cloud.speech.v2.ExplicitDecodingConfig.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Sample rate in Hertz of the audio data sent for recognition.
	//  Valid values are: 8000-48000. 16000 is optimal. For best results, set the
	//  sampling rate of the audio source to 16000 Hz. If that's not possible, use
	//  the native sample rate of the audio source (instead of re-sampling).
	//  Note that this field is marked as OPTIONAL for backward compatibility
	//  reasons. It is (and has always been) effectively REQUIRED.
	// +kcc:proto:field=google.cloud.speech.v2.ExplicitDecodingConfig.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`

	// Optional. Number of channels present in the audio data sent for
	//  recognition. Note that this field is marked as OPTIONAL for backward
	//  compatibility reasons. It is (and has always been) effectively REQUIRED.
	//
	//  The maximum allowed value is 8.
	// +kcc:proto:field=google.cloud.speech.v2.ExplicitDecodingConfig.audio_channel_count
	AudioChannelCount *int32 `json:"audioChannelCount,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.PhraseSet
type PhraseSet struct {

	// A list of word and phrases.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.phrases
	Phrases []PhraseSet_Phrase `json:"phrases,omitempty"`

	// Hint Boost. Positive value will increase the probability that a specific
	//  phrase will be recognized over other similar sounding phrases. The higher
	//  the boost, the higher the chance of false positive recognition as well.
	//  Valid `boost` values are between 0 (exclusive) and 20. We recommend using a
	//  binary search approach to finding the optimal value for your use case as
	//  well as adding phrases both with and without boost to your requests.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.boost
	Boost *float32 `json:"boost,omitempty"`

	// User-settable, human-readable name for the PhraseSet. Must be 63
	//  characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.PhraseSet.Phrase
type PhraseSet_Phrase struct {
	// The phrase itself.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.Phrase.value
	Value *string `json:"value,omitempty"`

	// Hint Boost. Overrides the boost set at the phrase set level.
	//  Positive value will increase the probability that a specific phrase will
	//  be recognized over other similar sounding phrases. The higher the boost,
	//  the higher the chance of false positive recognition as well. Negative
	//  boost values would correspond to anti-biasing. Anti-biasing is not
	//  enabled, so negative boost values will return an error. Boost values must
	//  be between 0 and 20. Any values outside that range will return an error.
	//  We recommend using a binary search approach to finding the optimal value
	//  for your use case as well as adding phrases both with and without boost
	//  to your requests.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.Phrase.boost
	Boost *float32 `json:"boost,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.RecognitionConfig
type RecognitionConfig struct {
	// Automatically detect decoding parameters.
	//  Preferred for supported formats.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.auto_decoding_config
	AutoDecodingConfig *AutoDetectDecodingConfig `json:"autoDecodingConfig,omitempty"`

	// Explicitly specified decoding parameters.
	//  Required if using headerless PCM audio (linear16, mulaw, alaw).
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.explicit_decoding_config
	ExplicitDecodingConfig *ExplicitDecodingConfig `json:"explicitDecodingConfig,omitempty"`

	// Optional. Which model to use for recognition requests. Select the model
	//  best suited to your domain to get best results.
	//
	//  Guidance for choosing which model to use can be found in the [Transcription
	//  Models
	//  Documentation](https://cloud.google.com/speech-to-text/v2/docs/transcription-model)
	//  and the models supported in each region can be found in the [Table Of
	//  Supported
	//  Models](https://cloud.google.com/speech-to-text/v2/docs/speech-to-text-supported-languages).
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.model
	Model *string `json:"model,omitempty"`

	// Optional. The language of the supplied audio as a
	//  [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag.
	//  Language tags are normalized to BCP-47 before they are used eg "en-us"
	//  becomes "en-US".
	//
	//  Supported languages for each model are listed in the [Table of Supported
	//  Models](https://cloud.google.com/speech-to-text/v2/docs/speech-to-text-supported-languages).
	//
	//  If additional languages are provided, recognition result will contain
	//  recognition in the most likely language detected. The recognition result
	//  will include the language tag of the language detected in the audio.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.language_codes
	LanguageCodes []string `json:"languageCodes,omitempty"`

	// Speech recognition features to enable.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.features
	Features *RecognitionFeatures `json:"features,omitempty"`

	// Speech adaptation context that weights recognizer predictions for specific
	//  words and phrases.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.adaptation
	Adaptation *SpeechAdaptation `json:"adaptation,omitempty"`

	// Optional. Use transcription normalization to automatically replace parts of
	//  the transcript with phrases of your choosing. For StreamingRecognize, this
	//  normalization only applies to stable partial transcripts (stability > 0.8)
	//  and final transcripts.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.transcript_normalization
	TranscriptNormalization *TranscriptNormalization `json:"transcriptNormalization,omitempty"`

	// Optional. Optional configuration used to automatically run translation on
	//  the given audio to the desired language for supported models.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.translation_config
	TranslationConfig *TranslationConfig `json:"translationConfig,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.RecognitionFeatures
type RecognitionFeatures struct {
	// If set to `true`, the server will attempt to filter out profanities,
	//  replacing all but the initial character in each filtered word with
	//  asterisks, for instance, "f***". If set to `false` or omitted, profanities
	//  won't be filtered out.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.profanity_filter
	ProfanityFilter *bool `json:"profanityFilter,omitempty"`

	// If `true`, the top result includes a list of words and the start and end
	//  time offsets (timestamps) for those words. If `false`, no word-level time
	//  offset information is returned. The default is `false`.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_word_time_offsets
	EnableWordTimeOffsets *bool `json:"enableWordTimeOffsets,omitempty"`

	// If `true`, the top result includes a list of words and the confidence for
	//  those words. If `false`, no word-level confidence information is returned.
	//  The default is `false`.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_word_confidence
	EnableWordConfidence *bool `json:"enableWordConfidence,omitempty"`

	// If `true`, adds punctuation to recognition result hypotheses. This feature
	//  is only available in select languages. The default `false` value does not
	//  add punctuation to result hypotheses.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_automatic_punctuation
	EnableAutomaticPunctuation *bool `json:"enableAutomaticPunctuation,omitempty"`

	// The spoken punctuation behavior for the call. If `true`, replaces spoken
	//  punctuation with the corresponding symbols in the request. For example,
	//  "how are you question mark" becomes "how are you?". See
	//  https://cloud.google.com/speech-to-text/docs/spoken-punctuation for
	//  support. If `false`, spoken punctuation is not replaced.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_spoken_punctuation
	EnableSpokenPunctuation *bool `json:"enableSpokenPunctuation,omitempty"`

	// The spoken emoji behavior for the call. If `true`, adds spoken emoji
	//  formatting for the request. This will replace spoken emojis with the
	//  corresponding Unicode symbols in the final transcript. If `false`, spoken
	//  emojis are not replaced.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_spoken_emojis
	EnableSpokenEmojis *bool `json:"enableSpokenEmojis,omitempty"`

	// Mode for recognizing multi-channel audio.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.multi_channel_mode
	MultiChannelMode *string `json:"multiChannelMode,omitempty"`

	// Configuration to enable speaker diarization and set additional
	//  parameters to make diarization better suited for your application.
	//  When this is enabled, we send all the words from the beginning of the
	//  audio for the top alternative in every consecutive STREAMING responses.
	//  This is done in order to improve our speaker tags as our models learn to
	//  identify the speakers in the conversation over time.
	//  For non-streaming requests, the diarization results will be provided only
	//  in the top alternative of the FINAL SpeechRecognitionResult.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.diarization_config
	DiarizationConfig *SpeakerDiarizationConfig `json:"diarizationConfig,omitempty"`

	// Maximum number of recognition hypotheses to be returned.
	//  The server may return fewer than `max_alternatives`.
	//  Valid values are `0`-`30`. A value of `0` or `1` will return a maximum of
	//  one. If omitted, will return a maximum of one.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.max_alternatives
	MaxAlternatives *int32 `json:"maxAlternatives,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.Recognizer
type Recognizer struct {

	// User-settable, human-readable name for the Recognizer. Must be 63
	//  characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. This field is now deprecated. Prefer the
	//  [`model`][google.cloud.speech.v2.RecognitionConfig.model] field in the
	//  [`RecognitionConfig`][google.cloud.speech.v2.RecognitionConfig] message.
	//
	//  Which model to use for recognition requests. Select the model best suited
	//  to your domain to get best results.
	//
	//  Guidance for choosing which model to use can be found in the [Transcription
	//  Models
	//  Documentation](https://cloud.google.com/speech-to-text/v2/docs/transcription-model)
	//  and the models supported in each region can be found in the [Table Of
	//  Supported
	//  Models](https://cloud.google.com/speech-to-text/v2/docs/speech-to-text-supported-languages).
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.model
	Model *string `json:"model,omitempty"`

	// Optional. This field is now deprecated. Prefer the
	//  [`language_codes`][google.cloud.speech.v2.RecognitionConfig.language_codes]
	//  field in the
	//  [`RecognitionConfig`][google.cloud.speech.v2.RecognitionConfig] message.
	//
	//  The language of the supplied audio as a
	//  [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag.
	//
	//  Supported languages for each model are listed in the [Table of Supported
	//  Models](https://cloud.google.com/speech-to-text/v2/docs/speech-to-text-supported-languages).
	//
	//  If additional languages are provided, recognition result will contain
	//  recognition in the most likely language detected. The recognition result
	//  will include the language tag of the language detected in the audio.
	//  When you create or update a Recognizer, these values are
	//  stored in normalized BCP-47 form. For example, "en-us" is stored as
	//  "en-US".
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.language_codes
	LanguageCodes []string `json:"languageCodes,omitempty"`

	// Default configuration to use for requests with this Recognizer.
	//  This can be overwritten by inline configuration in the
	//  [RecognizeRequest.config][google.cloud.speech.v2.RecognizeRequest.config]
	//  field.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.default_recognition_config
	DefaultRecognitionConfig *RecognitionConfig `json:"defaultRecognitionConfig,omitempty"`

	// Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.SpeakerDiarizationConfig
type SpeakerDiarizationConfig struct {
	// Required. Minimum number of speakers in the conversation. This range gives
	//  you more flexibility by allowing the system to automatically determine the
	//  correct number of speakers.
	//
	//  To fix the number of speakers detected in the audio, set
	//  `min_speaker_count` = `max_speaker_count`.
	// +kcc:proto:field=google.cloud.speech.v2.SpeakerDiarizationConfig.min_speaker_count
	MinSpeakerCount *int32 `json:"minSpeakerCount,omitempty"`

	// Required. Maximum number of speakers in the conversation. Valid values are:
	//  1-6. Must be >= `min_speaker_count`. This range gives you more flexibility
	//  by allowing the system to automatically determine the correct number of
	//  speakers.
	// +kcc:proto:field=google.cloud.speech.v2.SpeakerDiarizationConfig.max_speaker_count
	MaxSpeakerCount *int32 `json:"maxSpeakerCount,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation
type SpeechAdaptation struct {
	// A list of inline or referenced PhraseSets.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.phrase_sets
	PhraseSets []SpeechAdaptation_AdaptationPhraseSet `json:"phraseSets,omitempty"`

	// A list of inline CustomClasses. Existing CustomClass resources can be
	//  referenced directly in a PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.custom_classes
	CustomClasses []CustomClass `json:"customClasses,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet
type SpeechAdaptation_AdaptationPhraseSet struct {
	// The name of an existing PhraseSet resource. The user must have read
	//  access to the resource and it must not be deleted.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet.phrase_set
	PhraseSet *string `json:"phraseSet,omitempty"`

	// An inline defined PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet.inline_phrase_set
	InlinePhraseSet *PhraseSet `json:"inlinePhraseSet,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.TranscriptNormalization
type TranscriptNormalization struct {
	// A list of replacement entries. We will perform replacement with one entry
	//  at a time. For example, the second entry in ["cat" => "dog", "mountain cat"
	//  => "mountain dog"] will never be applied because we will always process the
	//  first entry before it. At most 100 entries.
	// +kcc:proto:field=google.cloud.speech.v2.TranscriptNormalization.entries
	Entries []TranscriptNormalization_Entry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.TranscriptNormalization.Entry
type TranscriptNormalization_Entry struct {
	// What to replace. Max length is 100 characters.
	// +kcc:proto:field=google.cloud.speech.v2.TranscriptNormalization.Entry.search
	Search *string `json:"search,omitempty"`

	// What to replace with. Max length is 100 characters.
	// +kcc:proto:field=google.cloud.speech.v2.TranscriptNormalization.Entry.replace
	Replace *string `json:"replace,omitempty"`

	// Whether the search is case sensitive.
	// +kcc:proto:field=google.cloud.speech.v2.TranscriptNormalization.Entry.case_sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.TranslationConfig
type TranslationConfig struct {
	// Required. The language code to translate to.
	// +kcc:proto:field=google.cloud.speech.v2.TranslationConfig.target_language
	TargetLanguage *string `json:"targetLanguage,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.CustomClass
type CustomClassObservedState struct {
	// Output only. Identifier. The resource name of the CustomClass.
	//  Format:
	//  `projects/{project}/locations/{location}/customClasses/{custom_class}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.name
	Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the CustomClass.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The CustomClass lifecycle state.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this resource was requested for deletion.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this resource will be purged.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields. This may be sent on update, undelete, and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Whether or not this CustomClass is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	//  the CustomClass is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The [KMS key version
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	//  with which the CustomClass is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.PhraseSet
type PhraseSetObservedState struct {
	// Output only. Identifier. The resource name of the PhraseSet.
	//  Format: `projects/{project}/locations/{location}/phraseSets/{phrase_set}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.name
	Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The PhraseSet lifecycle state.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this resource was requested for deletion.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this resource will be purged.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields. This may be sent on update, undelete, and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Whether or not this PhraseSet is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	//  the PhraseSet is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The [KMS key version
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	//  with which the PhraseSet is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.RecognitionConfig
type RecognitionConfigObservedState struct {
	// Speech adaptation context that weights recognizer predictions for specific
	//  words and phrases.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.adaptation
	Adaptation *SpeechAdaptationObservedState `json:"adaptation,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.Recognizer
type RecognizerObservedState struct {
	// Output only. Identifier. The resource name of the Recognizer.
	//  Format: `projects/{project}/locations/{location}/recognizers/{recognizer}`.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.name
	Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the Recognizer.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.uid
	Uid *string `json:"uid,omitempty"`

	// Default configuration to use for requests with this Recognizer.
	//  This can be overwritten by inline configuration in the
	//  [RecognizeRequest.config][google.cloud.speech.v2.RecognizeRequest.config]
	//  field.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.default_recognition_config
	DefaultRecognitionConfig *RecognitionConfigObservedState `json:"defaultRecognitionConfig,omitempty"`

	// Output only. The Recognizer lifecycle state.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this Recognizer was modified.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this Recognizer was requested for deletion.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this Recognizer will be purged.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields. This may be sent on update, undelete, and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Whether or not this Recognizer is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	//  the Recognizer is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The [KMS key version
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	//  with which the Recognizer is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation
type SpeechAdaptationObservedState struct {
	// A list of inline or referenced PhraseSets.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.phrase_sets
	PhraseSets []SpeechAdaptation_AdaptationPhraseSetObservedState `json:"phraseSets,omitempty"`

	// A list of inline CustomClasses. Existing CustomClass resources can be
	//  referenced directly in a PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.custom_classes
	CustomClasses []CustomClassObservedState `json:"customClasses,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet
type SpeechAdaptation_AdaptationPhraseSetObservedState struct {
	// An inline defined PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet.inline_phrase_set
	InlinePhraseSet *PhraseSetObservedState `json:"inlinePhraseSet,omitempty"`
}
