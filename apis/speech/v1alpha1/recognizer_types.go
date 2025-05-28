// Copyright 2025 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpeechRecognizerGVK = GroupVersion.WithKind("SpeechRecognizer")

type Parent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// SpeechRecognizerSpec defines the desired state of SpeechRecognizer
// +kcc:spec:proto=google.cloud.speech.v2.Recognizer
type SpeechRecognizerSpec struct {
	// The SpeechRecognizer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// User-settable, human-readable name for the Recognizer. Must be 63
	//  characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.display_name
	DisplayName *string `json:"displayName,omitempty"`

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

// +kcc:proto=google.cloud.speech.v2.RecognitionConfig
type RecognitionConfig struct {
	// Automatically detect decoding parameters.
	//  Preferred for supported formats.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.auto_decoding_config
	/* NOTYET AutoDecodingConfig *AutoDetectDecodingConfig `json:"autoDecodingConfig,omitempty"` */

	// Explicitly specified decoding parameters.
	//  Required if using headerless PCM audio (linear16, mulaw, alaw).
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.explicit_decoding_config
	/* NOTYET ExplicitDecodingConfig *ExplicitDecodingConfig `json:"explicitDecodingConfig,omitempty"`*/

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
	/* NOTYET Features *RecognitionFeatures `json:"features,omitempty"`*/

	// Speech adaptation context that weights recognizer predictions for specific
	//  words and phrases.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.adaptation
	/* NOTYET Adaptation *SpeechAdaptation `json:"adaptation,omitempty"`*/

	// Optional. Use transcription normalization to automatically replace parts of
	//  the transcript with phrases of your choosing. For StreamingRecognize, this
	//  normalization only applies to stable partial transcripts (stability > 0.8)
	//  and final transcripts.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.transcript_normalization
	/* NOTYET TranscriptNormalization *TranscriptNormalization `json:"transcriptNormalization,omitempty"`*/

	// Optional. Optional configuration used to automatically run translation on
	//  the given audio to the desired language for supported models.
	// +kcc:proto:field=google.cloud.speech.v2.RecognitionConfig.translation_config
	/* NOTYET TranslationConfig *TranslationConfig `json:"translationConfig,omitempty"`*/
}

// SpeechRecognizerStatus defines the config connector machine state of SpeechRecognizer
type SpeechRecognizerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpeechRecognizer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpeechRecognizerObservedState `json:"observedState,omitempty"`
}

// SpeechRecognizerObservedState is the state of the SpeechRecognizer resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.speech.v2.Recognizer
type SpeechRecognizerObservedState struct {
	// Output only. Identifier. The resource name of the Recognizer.
	//  Format: `projects/{project}/locations/{location}/recognizers/{recognizer}`.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the Recognizer.
	// +kcc:proto:field=google.cloud.speech.v2.Recognizer.uid
	UID *string `json:"uid,omitempty"`

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspeechrecognizer;gcpspeechrecognizers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpeechRecognizer is the Schema for the SpeechRecognizer API
// +k8s:openapi-gen=true
type SpeechRecognizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpeechRecognizerSpec   `json:"spec,omitempty"`
	Status SpeechRecognizerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpeechRecognizerList contains a list of SpeechRecognizer
type SpeechRecognizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpeechRecognizer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpeechRecognizer{}, &SpeechRecognizerList{})
}

// +kcc:proto=google.cloud.speech.v2.TranslationConfig
type TranslationConfig struct {
	// Required. The language code to translate to.
	// +kcc:proto:field=google.cloud.speech.v2.TranslationConfig.target_language
	// +required
	/* NOTYET TargetLanguage *string `json:"targetLanguage,omitempty"` */
}

// +kcc:proto=google.cloud.speech.v2.ExplicitDecodingConfig
type ExplicitDecodingConfig struct {
	// Required. Encoding of the audio data sent for recognition.
	// +kcc:proto:field=google.cloud.speech.v2.ExplicitDecodingConfig.encoding
	// +required
	/* NOTYET Encoding *string `json:"encoding,omitempty"` */

	// Optional. Sample rate in Hertz of the audio data sent for recognition.
	//  Valid values are: 8000-48000. 16000 is optimal. For best results, set the
	//  sampling rate of the audio source to 16000 Hz. If that's not possible, use
	//  the native sample rate of the audio source (instead of re-sampling).
	//  Note that this field is marked as OPTIONAL for backward compatibility
	//  reasons. It is (and has always been) effectively REQUIRED.
	// +kcc:proto:field=google.cloud.speech.v2.ExplicitDecodingConfig.sample_rate_hertz
	/* NOTYET SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"` */

	// Optional. Number of channels present in the audio data sent for
	//  recognition. Note that this field is marked as OPTIONAL for backward
	//  compatibility reasons. It is (and has always been) effectively REQUIRED.
	//
	//  The maximum allowed value is 8.
	// +kcc:proto:field=google.cloud.speech.v2.ExplicitDecodingConfig.audio_channel_count
	/* NOTYET AudioChannelCount *int32 `json:"audioChannelCount,omitempty"` */
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
	// +required
	/* NOTYET MinSpeakerCount *int32 `json:"minSpeakerCount,omitempty"` */

	// Required. Maximum number of speakers in the conversation. Valid values are:
	//  1-6. Must be >= `min_speaker_count`. This range gives you more flexibility
	//  by allowing the system to automatically determine the correct number of
	//  speakers.
	// +kcc:proto:field=google.cloud.speech.v2.SpeakerDiarizationConfig.max_speaker_count
	// +required
	/* NOTYET MaxSpeakerCount *int32 `json:"maxSpeakerCount,omitempty"` */
}

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation
type SpeechAdaptation struct {
	// A list of inline or referenced PhraseSets.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.phrase_sets
	PhraseSets []SpeechAdaptation_AdaptationPhraseSet `json:"phraseSets,omitempty"`

	// A list of inline CustomClasses. Existing CustomClass resources can be
	//  referenced directly in a PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.custom_classes
	CustomClasses []InlineCustomClass `json:"customClasses,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.speech.v2.SpeechAdaptation
type SpeechAdaptationObservedState struct {
	// A list of inline or referenced PhraseSets.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.phrase_sets
	PhraseSets []SpeechAdaptation_AdaptationPhraseSetObservedState `json:"phraseSets,omitempty"`

	// A list of inline CustomClasses. Existing CustomClass resources can be
	//  referenced directly in a PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.custom_classes
	CustomClasses []InlineCustomClassObservedState `json:"customClasses,omitempty"`
}

// We need InlineCustomClass and InlineCustomClassObservedState because they are inline fields in SpeechAdaptation
// and we need to distinguish them from Config Connector resource SpeechCustomClass
// SpeechCustomClass is a Config Connector resource which contains Config Connector specific fields

type InlineCustomClass struct {
	// Optional. User-settable, human-readable name for the CustomClass. Must be
	//  63 characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.display_name
	/* NOTYET
	DisplayName *string `json:"displayName,omitempty"`
	*/

	// A collection of class items.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.items
	/* NOTYET
	Items []CustomClass_ClassItem `json:"items,omitempty"`
	*/

	// Optional. Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

type InlineCustomClassObservedState struct {
	// Output only. Identifier. The resource name of the CustomClass.
	//  Format:
	//  `projects/{project}/locations/{location}/customClasses/{custom_class}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.name
	Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the CustomClass.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.uid
	UID *string `json:"uid,omitempty"`

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

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet
type SpeechAdaptation_AdaptationPhraseSet struct {
	// The name of an existing PhraseSet resource. The user must have read
	//  access to the resource and it must not be deleted.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet.phrase_set
	/* NOTYET PhraseSetRef *PhraseSetRef `json:"phraseSetRef,omitempty"` */

	// An inline defined PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet.inline_phrase_set
	InlinePhraseSet *InlinePhraseSet `json:"inlinePhraseSet,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet
type SpeechAdaptation_AdaptationPhraseSetObservedState struct {
	// An inline defined PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.SpeechAdaptation.AdaptationPhraseSet.inline_phrase_set
	InlinePhraseSet *InlinePhraseSetObservedState `json:"inlinePhraseSet,omitempty"`
}

// We need InlinePhraseSet and InlinePhraseSetObservedState because they are inline fields in SpeechAdaptation
// and we need to distinguish them from Config Connector resource SpeechPhraseSet
// SpeechPhraseSet is a Config Connector resource which contains Config Connector specific fields

// +kcc:proto=google.cloud.speech.v2.PhraseSet
type InlinePhraseSet struct {
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
	/* NOTYET
	Boost *string `json:"boost,omitempty"`
	*/

	// User-settable, human-readable name for the PhraseSet. Must be 63
	//  characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.display_name
	/* NOTYET
	DisplayName *string `json:"displayName,omitempty"`
	*/

	// Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.PhraseSet
type InlinePhraseSetObservedState struct {
	// Output only. Identifier. The resource name of the PhraseSet.
	//  Format: `projects/{project}/locations/{location}/phraseSets/{phrase_set}`.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.name
	Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the PhraseSet.
	// +kcc:proto:field=google.cloud.speech.v2.PhraseSet.uid
	UID *string `json:"uid,omitempty"`

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
