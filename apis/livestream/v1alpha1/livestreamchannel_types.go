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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LiveStreamChannelGVK = GroupVersion.WithKind("LiveStreamChannel")

// LiveStreamChannelSpec defines the desired state of LiveStreamChannel
// +kcc:spec:proto=google.cloud.video.livestream.v1.Channel
type LiveStreamChannelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The LiveStreamChannel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A list of input attachments that this channel uses.
	// One channel can have multiple inputs as the input sources. Only one
	// input can be selected as the input source at one time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.input_attachments
	InputAttachments []InputAttachment `json:"inputAttachments,omitempty"`

	// Required. Information about the output (that is, the Cloud Storage bucket
	// to store the generated live stream).
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.output
	Output *Channel_Output `json:"output,omitempty"`

	// List of elementary streams.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.elementary_streams
	ElementaryStreams []ElementaryStream `json:"elementaryStreams,omitempty"`

	// List of multiplexing settings for output streams.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.mux_streams
	MuxStreams []MuxStream `json:"muxStreams,omitempty"`

	// List of output manifests.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.manifests
	Manifests []Manifest `json:"manifests,omitempty"`

	// List of output sprite sheets.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.sprite_sheets
	SpriteSheets []SpriteSheet `json:"spriteSheets,omitempty"`

	// Configuration of platform logs for this channel.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.log_config
	LogConfig *LogConfig `json:"logConfig,omitempty"`

	// Configuration of timecode for this channel.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.timecode_config
	TimecodeConfig *TimecodeConfig `json:"timecodeConfig,omitempty"`

	// Encryption configurations for this channel. Each configuration has an ID
	// which is referred to by each MuxStream to indicate which configuration is
	// used for that output.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.encryptions
	Encryptions []Encryption `json:"encryptions,omitempty"`

	// The configuration for input sources defined in
	// [input_attachments][google.cloud.video.livestream.v1.Channel.input_attachments].
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Optional. Configuration for retention of output files for this channel.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.retention_config
	RetentionConfig *RetentionConfig `json:"retentionConfig,omitempty"`

	// Optional. List of static overlay images. Those images display over the
	// output content for the whole duration of the live stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.static_overlays
	StaticOverlays []StaticOverlay `json:"staticOverlays,omitempty"`
}

// LiveStreamChannelStatus defines the config connector machine state of LiveStreamChannel
type LiveStreamChannelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LiveStreamChannel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LiveStreamChannelObservedState `json:"observedState,omitempty"`
}

// LiveStreamChannelObservedState is the state of the LiveStreamChannel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.livestream.v1.Channel
type LiveStreamChannelObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The
	// [InputAttachment.key][google.cloud.video.livestream.v1.InputAttachment.key]
	// that serves as the current input source. The first input in the
	// [input_attachments][google.cloud.video.livestream.v1.Channel.input_attachments]
	// is the initial input source.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.active_input
	ActiveInput *string `json:"activeInput,omitempty"`

	// Output only. State of the streaming operation.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.streaming_state
	StreamingState *string `json:"streamingState,omitempty"`

	// Output only. A description of the reason for the streaming error. This
	// property is always present when
	// [streaming_state][google.cloud.video.livestream.v1.Channel.streaming_state]
	// is
	// [STREAMING_ERROR][google.cloud.video.livestream.v1.Channel.StreamingState.STREAMING_ERROR].
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.streaming_error
	StreamingError *common.Status `json:"streamingError,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplivestreamchannel;gcplivestreamchannels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LiveStreamChannel is the Schema for the LiveStreamChannel API
// +k8s:openapi-gen=true
type LiveStreamChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LiveStreamChannelSpec   `json:"spec,omitempty"`
	Status LiveStreamChannelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LiveStreamChannelList contains a list of LiveStreamChannel
type LiveStreamChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveStreamChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LiveStreamChannel{}, &LiveStreamChannelList{})
}

// +kcc:proto=google.cloud.video.livestream.v1.InputAttachment
type InputAttachment struct {
	// A unique key for this input attachment. The key must be 1-63
	// characters in length. The key must begin and end with a letter (regardless
	// of case) or a number, but can contain dashes or underscores in between.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.key
	Key *string `json:"key,omitempty"`

	// The resource name of an existing input, in the form of:
	// `projects/{project}/locations/{location}/inputs/{inputId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.input
	InputRef *LiveStreamInputRef `json:"inputRef,omitempty"`

	// Automatic failover configurations.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.automatic_failover
	AutomaticFailover *InputAttachment_AutomaticFailover `json:"automaticFailover,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.StaticOverlay
type StaticOverlay struct {
	// Required. Asset to use for the overlaid image.
	// The asset must be represented in the form of:
	// `projects/{project}/locations/{location}/assets/{assetId}`.
	// The asset's resource type must be image.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.asset
	AssetRef *LiveStreamAssetRef `json:"assetRef,omitempty"`

	// Optional. Normalized image resolution, based on output video resolution.
	// Valid values are [0.0, 1.0]. To respect the original image aspect ratio,
	// set either `w` or `h` to 0. To use the original image resolution, set both
	// `w` and `h` to 0. The default is {0, 0}.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.resolution
	Resolution *NormalizedResolution `json:"resolution,omitempty"`

	// Optional. Position of the image in terms of normalized coordinates of the
	// upper-left corner of the image, based on output video resolution. For
	// example, use the x and y coordinates {0, 0} to position the top-left corner
	// of the overlay animation in the top-left corner of the output video.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.position
	Position *NormalizedCoordinate `json:"position,omitempty"`

	// Optional. Target image opacity. Valid values are from `1.0` (solid,
	// default) to `0.0` (transparent), exclusive. Set this to a value greater
	// than `0.0`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.opacity
	Opacity *float64 `json:"opacity,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.SecretManagerSource
type Encryption_SecretManagerSource struct {
	// Required. The name of the Secret Version containing the encryption key.
	// `projects/{project}/secrets/{secret_id}/versions/{version_number}`
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.SecretManagerSource.secret_version
	SecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionRef,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Aes128Encryption
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Encryption_Aes128Encryption struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Clearkey
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Encryption_Clearkey struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Fairplay
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Encryption_Fairplay struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Playready
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Encryption_Playready struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.SampleAesEncryption
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Encryption_SampleAesEncryption struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Widevine
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Encryption_Widevine struct {
}
