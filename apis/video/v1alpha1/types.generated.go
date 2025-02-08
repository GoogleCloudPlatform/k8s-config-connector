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


// +kcc:proto=google.cloud.video.livestream.v1.AudioFormat
type AudioFormat struct {
	// Audio codec used in this audio stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioFormat.codec
	Codec *string `json:"codec,omitempty"`

	// The number of audio channels.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioFormat.channel_count
	ChannelCount *int32 `json:"channelCount,omitempty"`

	// A list of channel names specifying the layout of the audio channels.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioFormat.channel_layout
	ChannelLayout []string `json:"channelLayout,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.AudioStreamProperty
type AudioStreamProperty struct {
	// Index of this audio stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStreamProperty.index
	Index *int32 `json:"index,omitempty"`

	// Properties of the audio format.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStreamProperty.audio_format
	AudioFormat *AudioFormat `json:"audioFormat,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Input
type Input struct {
	// The resource name of the input, in the form of:
	//  `projects/{project}/locations/{location}/inputs/{inputId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.name
	Name *string `json:"name,omitempty"`

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
	SecurityRules *Input_SecurityRule `json:"securityRules,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Input.SecurityRule
type Input_SecurityRule struct {
	// At least one ip range must match unless none specified. The IP range is
	//  defined by CIDR block: for example, `192.0.1.0/24` for a range and
	//  `192.0.1.0/32` for a single IP address.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Input.SecurityRule.ip_ranges
	IPRanges []string `json:"ipRanges,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.InputStreamProperty
type InputStreamProperty struct {
	// The time that the current input stream is accepted and the connection is
	//  established.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputStreamProperty.last_establish_time
	LastEstablishTime *string `json:"lastEstablishTime,omitempty"`

	// Properties of the video streams.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputStreamProperty.video_streams
	VideoStreams []VideoStreamProperty `json:"videoStreams,omitempty"`

	// Properties of the audio streams.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputStreamProperty.audio_streams
	AudioStreams []AudioStreamProperty `json:"audioStreams,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.PreprocessingConfig
type PreprocessingConfig struct {
	// Audio preprocessing configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.audio
	Audio *PreprocessingConfig_Audio `json:"audio,omitempty"`

	// Specify the video cropping configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.crop
	Crop *PreprocessingConfig_Crop `json:"crop,omitempty"`

	// Specify the video pad filter configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.PreprocessingConfig.pad
	Pad *PreprocessingConfig_Pad `json:"pad,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.PreprocessingConfig.Audio
type PreprocessingConfig_Audio struct {
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
type PreprocessingConfig_Crop struct {
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
type PreprocessingConfig_Pad struct {
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

// +kcc:proto=google.cloud.video.livestream.v1.VideoFormat
type VideoFormat struct {
	// Video codec used in this video stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoFormat.codec
	Codec *string `json:"codec,omitempty"`

	// The width of the video stream in pixels.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoFormat.width_pixels
	WidthPixels *int32 `json:"widthPixels,omitempty"`

	// The height of the video stream in pixels.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoFormat.height_pixels
	HeightPixels *int32 `json:"heightPixels,omitempty"`

	// The frame rate of the input video stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoFormat.frame_rate
	FrameRate *float64 `json:"frameRate,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.VideoStreamProperty
type VideoStreamProperty struct {
	// Index of this video stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStreamProperty.index
	Index *int32 `json:"index,omitempty"`

	// Properties of the video format.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStreamProperty.video_format
	VideoFormat *VideoFormat `json:"videoFormat,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Input
type InputObservedState struct {
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
