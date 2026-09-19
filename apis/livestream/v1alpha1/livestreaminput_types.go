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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LiveStreamInputGVK = GroupVersion.WithKind("LiveStreamInput")

// LiveStreamInputSpec defines the desired state of LiveStreamInput
// +kcc:spec:proto=google.cloud.video.livestream.v1.Input
type LiveStreamInputSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The LiveStreamInput name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Source type.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.type
	Type *string `json:"type,omitempty"`

	// Tier defines the maximum input specification that is accepted by the
	//  video pipeline. The billing is charged based on the tier specified here.
	//  See [Pricing](https://cloud.google.com/livestream/pricing) for more detail.
	//  The default is `HD`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.tier
	Tier *string `json:"tier,omitempty"`

	// Preprocessing configurations.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.preprocessing_config
	PreprocessingConfig *PreprocessingConfig `json:"preprocessingConfig,omitempty"`

	// Security rule for access control.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.security_rules
	SecurityRules *InputSecurityRule `json:"securityRules,omitempty"`
}

// LiveStreamInputStatus defines the config connector machine state of LiveStreamInput
type LiveStreamInputStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LiveStreamInput resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LiveStreamInputObservedState `json:"observedState,omitempty"`
}

// LiveStreamInputObservedState is the state of the LiveStreamInput resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.livestream.v1.Input
type LiveStreamInputObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. URI to push the input stream to.
	//  Its format depends on the input
	//  [type][google.cloud.video.livestream.v1.Input.type], for example:
	//
	//  *  `RTMP_PUSH`: `rtmp://1.2.3.4/live/{STREAM-ID}`
	//  *  `SRT_PUSH`: `srt://1.2.3.4:4201?streamid={STREAM-ID}`
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.uri
	URI *string `json:"uri,omitempty"`

	// Output only. The information for the input stream. This field will be
	//  present only when this input receives the input stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.input_stream_property
	InputStreamProperty *InputStreamProperty `json:"inputStreamProperty,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplivestreaminput;gcplivestreaminputs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LiveStreamInput is the Schema for the LiveStreamInput API
// +k8s:openapi-gen=true
type LiveStreamInput struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LiveStreamInputSpec   `json:"spec,omitempty"`
	Status LiveStreamInputStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LiveStreamInputList contains a list of LiveStreamInput
type LiveStreamInputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveStreamInput `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LiveStreamInput{}, &LiveStreamInputList{})
}

// +kcc:proto=google.cloud.video.livestream.v1.Input.SecurityRule
type InputSecurityRule struct {
	// At least one ip range must match unless none specified. The IP range is
	//  defined by CIDR block: for example, `192.0.1.0/24` for a range and
	//  `192.0.1.0/32` for a single IP address.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.SecurityRule.ip_ranges
	IPRanges []string `json:"ipRanges,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.PreprocessingConfig
type PreprocessingConfig struct {
	// Audio preprocessing configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.audio
	Audio *PreprocessingConfigAudio `json:"audio,omitempty"`

	// Specify the video cropping configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.crop
	Crop *PreprocessingConfigCrop `json:"crop,omitempty"`

	// Specify the video pad filter configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.pad
	Pad *PreprocessingConfigPad `json:"pad,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.PreprocessingConfig.Audio
type PreprocessingConfigAudio struct {
	// Specify audio loudness normalization in loudness units relative to full
	//  scale (LUFS). Enter a value between -24 and 0 according to the following:
	//
	//  - -24 is the Advanced Television Systems Committee (ATSC A/85)
	//  - -23 is the EU R128 broadcast standard
	//  - -19 is the prior standard for online mono audio
	//  - -18 is the ReplayGain standard
	//  - -16 is the prior standard for stereo audio
	//  - -14 is the new online audio standard recommended by Spotify, as well as
	//  Amazon Echo
	//  - 0 disables normalization. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Audio.lufs
	Lufs *float64 `json:"lufs,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.PreprocessingConfig.Crop
type PreprocessingConfigCrop struct {
	// The number of pixels to crop from the top. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Crop.top_pixels
	TopPixels *int32 `json:"topPixels,omitempty"`

	// The number of pixels to crop from the bottom. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Crop.bottom_pixels
	BottomPixels *int32 `json:"bottomPixels,omitempty"`

	// The number of pixels to crop from the left. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Crop.left_pixels
	LeftPixels *int32 `json:"leftPixels,omitempty"`

	// The number of pixels to crop from the right. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Crop.right_pixels
	RightPixels *int32 `json:"rightPixels,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.PreprocessingConfig.Pad
type PreprocessingConfigPad struct {
	// The number of pixels to add to the top. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Pad.top_pixels
	TopPixels *int32 `json:"topPixels,omitempty"`

	// The number of pixels to add to the bottom. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Pad.bottom_pixels
	BottomPixels *int32 `json:"bottomPixels,omitempty"`

	// The number of pixels to add to the left. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Pad.left_pixels
	LeftPixels *int32 `json:"leftPixels,omitempty"`

	// The number of pixels to add to the right. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.Pad.right_pixels
	RightPixels *int32 `json:"rightPixels,omitempty"`
}
