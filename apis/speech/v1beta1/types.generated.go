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

// +generated:types
// krm.group: speech.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.speech.v2
// resource: SpeechRecognizer:Recognizer
// resource: SpeechCustomClass:CustomClass
// resource: SpeechPhraseSet:PhraseSet

package v1beta1

// +kcc:proto=google.cloud.speech.v2.AutoDetectDecodingConfig
type AutoDetectDecodingConfig struct {
}

// +kcc:proto=google.cloud.speech.v2.CustomClass.ClassItem
type CustomClass_ClassItem struct {
	// The class item's value.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.ClassItem.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.RecognitionFeatures
type RecognitionFeatures struct {
	// If set to `true`, the server will attempt to filter out profanities,
	//  replacing all but the initial character in each filtered word with
	//  asterisks, for instance, "f***". If set to `false` or omitted, profanities
	//  won't be filtered out.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.profanity_filter
	/* NOTYET ProfanityFilter *bool `json:"profanityFilter,omitempty"` */

	// If `true`, the top result includes a list of words and the start and end
	//  time offsets (timestamps) for those words. If `false`, no word-level time
	//  offset information is returned. The default is `false`.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_word_time_offsets
	/* NOTYET EnableWordTimeOffsets *bool `json:"enableWordTimeOffsets,omitempty"` */

	// If `true`, the top result includes a list of words and the confidence for
	//  those words. If `false`, no word-level confidence information is returned.
	//  The default is `false`.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_word_confidence
	/* NOTYET EnableWordConfidence *bool `json:"enableWordConfidence,omitempty"` */

	// If `true`, adds punctuation to recognition result hypotheses. This feature
	//  is only available in select languages. The default `false` value does not
	//  add punctuation to result hypotheses.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_automatic_punctuation
	/* NOTYET EnableAutomaticPunctuation *bool `json:"enableAutomaticPunctuation,omitempty"` */

	// The spoken punctuation behavior for the call. If `true`, replaces spoken
	//  punctuation with the corresponding symbols in the request. For example,
	//  "how are you question mark" becomes "how are you?". See
	//  https://cloud.google.com/speech-to-text/docs/spoken-punctuation for
	//  support. If `false`, spoken punctuation is not replaced.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_spoken_punctuation
	/* NOTYET EnableSpokenPunctuation *bool `json:"enableSpokenPunctuation,omitempty"` */

	// The spoken emoji behavior for the call. If `true`, adds spoken emoji
	//  formatting for the request. This will replace spoken emojis with the
	//  corresponding Unicode symbols in the final transcript. If `false`, spoken
	//  emojis are not replaced.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.enable_spoken_emojis
	/* NOTYET EnableSpokenEmojis *bool `json:"enableSpokenEmojis,omitempty"` */

	// Mode for recognizing multi-channel audio.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionFeatures.multi_channel_mode
	/* NOTYET MultiChannelMode *string `json:"multiChannelMode,omitempty"` */

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
	/* NOTYET MaxAlternatives *int32 `json:"maxAlternatives,omitempty"` */
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
	/* NOTYET Search *string `json:"search,omitempty"` */

	// What to replace with. Max length is 100 characters.
	// +kcc:proto:field=google.cloud.speech.v2.TranscriptNormalization.Entry.replace
	/* NOTYET Replace *string `json:"replace,omitempty"` */

	// Whether the search is case sensitive.
	// +kcc:proto:field=google.cloud.speech.v2.TranscriptNormalization.Entry.case_sensitive
	/* NOTYET CaseSensitive *bool `json:"caseSensitive,omitempty"` */
}

// +kcc:proto=google.cloud.speech.v2.RecognitionConfig
type RecognitionConfigObservedState struct {
	// Speech adaptation context that weights recognizer predictions for specific
	//  words and phrases.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.adaptation
	Adaptation *SpeechAdaptationObservedState `json:"adaptation,omitempty"`
}
