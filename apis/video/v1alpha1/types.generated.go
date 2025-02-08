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


// +kcc:proto=google.cloud.video.transcoder.v1.AdBreak
type AdBreak struct {
	// Start time in seconds for the ad break, relative to the output file
	//  timeline. The default is `0s`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AdBreak.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.AudioStream
type AudioStream struct {
	// The codec for this audio stream. The default is `aac`.
	//
	//  Supported audio codecs:
	//
	//  - `aac`
	//  - `aac-he`
	//  - `aac-he-v2`
	//  - `mp3`
	//  - `ac3`
	//  - `eac3`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.codec
	Codec *string `json:"codec,omitempty"`

	// Required. Audio bitrate in bits per second. Must be between 1 and
	//  10,000,000.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Number of audio channels. Must be between 1 and 6. The default is 2.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.channel_count
	ChannelCount *int32 `json:"channelCount,omitempty"`

	// A list of channel names specifying layout of the audio channels.
	//  This only affects the metadata embedded in the container headers, if
	//  supported by the specified format. The default is `["fl", "fr"]`.
	//
	//  Supported channel names:
	//
	//  - `fl` - Front left channel
	//  - `fr` - Front right channel
	//  - `sl` - Side left channel
	//  - `sr` - Side right channel
	//  - `fc` - Front center channel
	//  - `lfe` - Low frequency
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.channel_layout
	ChannelLayout []string `json:"channelLayout,omitempty"`

	// The mapping for the `Job.edit_list` atoms with audio `EditAtom.inputs`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.mapping
	Mapping []AudioStream_AudioMapping `json:"mapping,omitempty"`

	// The audio sample rate in Hertz. The default is 48000 Hertz.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`

	// The BCP-47 language code, such as `en-US` or `sr-Latn`. For more
	//  information, see
	//  https://www.unicode.org/reports/tr35/#Unicode_locale_identifier. Not
	//  supported in MP4 files.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The name for this particular audio stream that
	//  will be added to the HLS/DASH manifest. Not supported in MP4 files.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.AudioStream.AudioMapping
type AudioStream_AudioMapping struct {
	// Required. The `EditAtom.key` that references the atom with audio inputs
	//  in the `Job.edit_list`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.AudioMapping.atom_key
	AtomKey *string `json:"atomKey,omitempty"`

	// Required. The `Input.key` that identifies the input file.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.AudioMapping.input_key
	InputKey *string `json:"inputKey,omitempty"`

	// Required. The zero-based index of the track in the input file.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.AudioMapping.input_track
	InputTrack *int32 `json:"inputTrack,omitempty"`

	// Required. The zero-based index of the channel in the input audio stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.AudioMapping.input_channel
	InputChannel *int32 `json:"inputChannel,omitempty"`

	// Required. The zero-based index of the channel in the output audio stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.AudioMapping.output_channel
	OutputChannel *int32 `json:"outputChannel,omitempty"`

	// Audio volume control in dB. Negative values decrease volume,
	//  positive values increase. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.AudioStream.AudioMapping.gain_db
	GainDb *float64 `json:"gainDb,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.EditAtom
type EditAtom struct {
	// A unique key for this atom. Must be specified when using advanced
	//  mapping.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.EditAtom.key
	Key *string `json:"key,omitempty"`

	// List of `Input.key`s identifying files that should be used in this atom.
	//  The listed `inputs` must have the same timeline.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.EditAtom.inputs
	Inputs []string `json:"inputs,omitempty"`

	// End time in seconds for the atom, relative to the input file timeline.
	//  When `end_time_offset` is not specified, the `inputs` are used until
	//  the end of the atom.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.EditAtom.end_time_offset
	EndTimeOffset *string `json:"endTimeOffset,omitempty"`

	// Start time in seconds for the atom, relative to the input file timeline.
	//  The default is `0s`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.EditAtom.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.ElementaryStream
type ElementaryStream struct {
	// A unique key for this elementary stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.ElementaryStream.key
	Key *string `json:"key,omitempty"`

	// Encoding of a video stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.ElementaryStream.video_stream
	VideoStream *VideoStream `json:"videoStream,omitempty"`

	// Encoding of an audio stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.ElementaryStream.audio_stream
	AudioStream *AudioStream `json:"audioStream,omitempty"`

	// Encoding of a text stream. For example, closed captions or subtitles.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.ElementaryStream.text_stream
	TextStream *TextStream `json:"textStream,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption
type Encryption struct {
	// Required. Identifier for this set of encryption options.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.id
	ID *string `json:"id,omitempty"`

	// Configuration for AES-128 encryption.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.aes_128
	Aes128 *Encryption_Aes128Encryption `json:"aes128,omitempty"`

	// Configuration for SAMPLE-AES encryption.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.sample_aes
	SampleAes *Encryption_SampleAesEncryption `json:"sampleAes,omitempty"`

	// Configuration for MPEG Common Encryption (MPEG-CENC).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.mpeg_cenc
	MpegCenc *Encryption_MpegCommonEncryption `json:"mpegCenc,omitempty"`

	// Keys are stored in Google Secret Manager.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.secret_manager_key_source
	SecretManagerKeySource *Encryption_SecretManagerSource `json:"secretManagerKeySource,omitempty"`

	// Required. DRM system(s) to use; at least one must be specified. If a
	//  DRM system is omitted, it is considered disabled.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.drm_systems
	DrmSystems *Encryption_DrmSystems `json:"drmSystems,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.Aes128Encryption
type Encryption_Aes128Encryption struct {
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.Clearkey
type Encryption_Clearkey struct {
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.DrmSystems
type Encryption_DrmSystems struct {
	// Widevine configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.DrmSystems.widevine
	Widevine *Encryption_Widevine `json:"widevine,omitempty"`

	// Fairplay configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.DrmSystems.fairplay
	Fairplay *Encryption_Fairplay `json:"fairplay,omitempty"`

	// Playready configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.DrmSystems.playready
	Playready *Encryption_Playready `json:"playready,omitempty"`

	// Clearkey configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.DrmSystems.clearkey
	Clearkey *Encryption_Clearkey `json:"clearkey,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.Fairplay
type Encryption_Fairplay struct {
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.MpegCommonEncryption
type Encryption_MpegCommonEncryption struct {
	// Required. Specify the encryption scheme.
	//
	//  Supported encryption schemes:
	//
	//  - `cenc`
	//  - `cbcs`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.MpegCommonEncryption.scheme
	Scheme *string `json:"scheme,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.Playready
type Encryption_Playready struct {
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.SampleAesEncryption
type Encryption_SampleAesEncryption struct {
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.SecretManagerSource
type Encryption_SecretManagerSource struct {
	// Required. The name of the Secret Version containing the encryption key in
	//  the following format:
	//  `projects/{project}/secrets/{secret_id}/versions/{version_number}`
	//
	//  Note that only numbered versions are supported. Aliases like "latest" are
	//  not supported.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.SecretManagerSource.secret_version
	SecretVersion *string `json:"secretVersion,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.Widevine
type Encryption_Widevine struct {
}

// +kcc:proto=google.cloud.video.transcoder.v1.Input
type Input struct {
	// A unique key for this input. Must be specified when using advanced
	//  mapping and edit lists.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Input.key
	Key *string `json:"key,omitempty"`

	// URI of the media. Input files must be at least 5 seconds in duration and
	//  stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
	//  If empty, the value is populated from `Job.input_uri`. See
	//  [Supported input and output
	//  formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Input.uri
	URI *string `json:"uri,omitempty"`

	// Preprocessing configurations.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Input.preprocessing_config
	PreprocessingConfig *PreprocessingConfig `json:"preprocessingConfig,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.JobConfig
type JobConfig struct {
	// List of input assets stored in Cloud Storage.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.inputs
	Inputs []Input `json:"inputs,omitempty"`

	// List of `Edit atom`s. Defines the ultimate timeline of the resulting
	//  file or manifest.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.edit_list
	EditList []EditAtom `json:"editList,omitempty"`

	// List of elementary streams.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.elementary_streams
	ElementaryStreams []ElementaryStream `json:"elementaryStreams,omitempty"`

	// List of multiplexing settings for output streams.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.mux_streams
	MuxStreams []MuxStream `json:"muxStreams,omitempty"`

	// List of output manifests.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.manifests
	Manifests []Manifest `json:"manifests,omitempty"`

	// Output configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.output
	Output *Output `json:"output,omitempty"`

	// List of ad breaks. Specifies where to insert ad break tags in the output
	//  manifests.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.ad_breaks
	AdBreaks []AdBreak `json:"adBreaks,omitempty"`

	// Destination on Pub/Sub.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.pubsub_destination
	PubsubDestination *PubsubDestination `json:"pubsubDestination,omitempty"`

	// List of output sprite sheets.
	//  Spritesheets require at least one VideoStream in the Jobconfig.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.sprite_sheets
	SpriteSheets []SpriteSheet `json:"spriteSheets,omitempty"`

	// List of overlays on the output video, in descending Z-order.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.overlays
	Overlays []Overlay `json:"overlays,omitempty"`

	// List of encryption configurations for the content.
	//  Each configuration has an ID. Specify this ID in the
	//  [MuxStream.encryption_id][google.cloud.video.transcoder.v1.MuxStream.encryption_id]
	//  field to indicate the configuration to use for that `MuxStream` output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.encryptions
	Encryptions []Encryption `json:"encryptions,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.JobTemplate
type JobTemplate struct {
	// The resource name of the job template.
	//  Format:
	//  `projects/{project_number}/locations/{location}/jobTemplates/{job_template}`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobTemplate.name
	Name *string `json:"name,omitempty"`

	// The configuration for this template.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobTemplate.config
	Config *JobConfig `json:"config,omitempty"`

	// The labels associated with this job template. You can use these to organize
	//  and group your job templates.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Manifest
type Manifest struct {
	// The name of the generated file. The default is `manifest` with the
	//  extension suffix corresponding to the `Manifest.type`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Manifest.file_name
	FileName *string `json:"fileName,omitempty"`

	// Required. Type of the manifest.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Manifest.type
	Type *string `json:"type,omitempty"`

	// Required. List of user given `MuxStream.key`s that should appear in this
	//  manifest.
	//
	//  When `Manifest.type` is `HLS`, a media manifest with name `MuxStream.key`
	//  and `.m3u8` extension is generated for each element of the
	//  `Manifest.mux_streams`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Manifest.mux_streams
	MuxStreams []string `json:"muxStreams,omitempty"`

	// `DASH` manifest configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Manifest.dash
	Dash *Manifest_DashConfig `json:"dash,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Manifest.DashConfig
type Manifest_DashConfig struct {
	// The segment reference scheme for a `DASH` manifest. The default is
	//  `SEGMENT_LIST`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Manifest.DashConfig.segment_reference_scheme
	SegmentReferenceScheme *string `json:"segmentReferenceScheme,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.MuxStream
type MuxStream struct {
	// A unique key for this multiplexed stream. HLS media manifests will be
	//  named `MuxStream.key` with the `.m3u8` extension suffix.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.MuxStream.key
	Key *string `json:"key,omitempty"`

	// The name of the generated file. The default is `MuxStream.key` with the
	//  extension suffix corresponding to the `MuxStream.container`.
	//
	//  Individual segments also have an incremental 10-digit zero-padded suffix
	//  starting from 0 before the extension, such as `mux_stream0000000123.ts`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.MuxStream.file_name
	FileName *string `json:"fileName,omitempty"`

	// The container format. The default is `mp4`
	//
	//  Supported container formats:
	//
	//  - `ts`
	//  - `fmp4`- the corresponding file extension is `.m4s`
	//  - `mp4`
	//  - `vtt`
	//
	//  See also:
	//  [Supported input and output
	//  formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats)
	// +kcc:proto:field=google.cloud.video.transcoder.v1.MuxStream.container
	Container *string `json:"container,omitempty"`

	// List of `ElementaryStream.key`s multiplexed in this stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.MuxStream.elementary_streams
	ElementaryStreams []string `json:"elementaryStreams,omitempty"`

	// Segment settings for `ts`, `fmp4` and `vtt`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.MuxStream.segment_settings
	SegmentSettings *SegmentSettings `json:"segmentSettings,omitempty"`

	// Identifier of the encryption configuration to use. If omitted, output will
	//  be unencrypted.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.MuxStream.encryption_id
	EncryptionID *string `json:"encryptionID,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Output
type Output struct {
	// URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	//  If empty, the value is populated from `Job.output_uri`. See
	//  [Supported input and output
	//  formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Output.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay
type Overlay struct {
	// Image overlay.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.image
	Image *Overlay_Image `json:"image,omitempty"`

	// List of Animations. The list should be chronological, without any time
	//  overlap.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.animations
	Animations []Overlay_Animation `json:"animations,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay.Animation
type Overlay_Animation struct {
	// Display static overlay object.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.Animation.animation_static
	AnimationStatic *Overlay_AnimationStatic `json:"animationStatic,omitempty"`

	// Display overlay object with fade animation.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.Animation.animation_fade
	AnimationFade *Overlay_AnimationFade `json:"animationFade,omitempty"`

	// End previous animation.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.Animation.animation_end
	AnimationEnd *Overlay_AnimationEnd `json:"animationEnd,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay.AnimationEnd
type Overlay_AnimationEnd struct {
	// The time to end overlay object, in seconds. Default: 0
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationEnd.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay.AnimationFade
type Overlay_AnimationFade struct {
	// Required. Type of fade animation: `FADE_IN` or `FADE_OUT`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationFade.fade_type
	FadeType *string `json:"fadeType,omitempty"`

	// Normalized coordinates based on output video resolution. Valid
	//  values: `0.0`–`1.0`. `xy` is the upper-left coordinate of the overlay
	//  object. For example, use the x and y coordinates {0,0} to position the
	//  top-left corner of the overlay animation in the top-left corner of the
	//  output video.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationFade.xy
	Xy *Overlay_NormalizedCoordinate `json:"xy,omitempty"`

	// The time to start the fade animation, in seconds. Default: 0
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationFade.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`

	// The time to end the fade animation, in seconds. Default:
	//  `start_time_offset` + 1s
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationFade.end_time_offset
	EndTimeOffset *string `json:"endTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay.AnimationStatic
type Overlay_AnimationStatic struct {
	// Normalized coordinates based on output video resolution. Valid
	//  values: `0.0`–`1.0`. `xy` is the upper-left coordinate of the overlay
	//  object. For example, use the x and y coordinates {0,0} to position the
	//  top-left corner of the overlay animation in the top-left corner of the
	//  output video.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationStatic.xy
	Xy *Overlay_NormalizedCoordinate `json:"xy,omitempty"`

	// The time to start displaying the overlay object, in seconds. Default: 0
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.AnimationStatic.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay.Image
type Overlay_Image struct {
	// Required. URI of the image in Cloud Storage. For example,
	//  `gs://bucket/inputs/image.png`. Only PNG and JPEG images are supported.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.Image.uri
	URI *string `json:"uri,omitempty"`

	// Normalized image resolution, based on output video resolution. Valid
	//  values: `0.0`–`1.0`. To respect the original image aspect ratio, set
	//  either `x` or `y` to `0.0`. To use the original image resolution, set
	//  both `x` and `y` to `0.0`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.Image.resolution
	Resolution *Overlay_NormalizedCoordinate `json:"resolution,omitempty"`

	// Target image opacity. Valid values are from  `1.0` (solid, default) to
	//  `0.0` (transparent), exclusive. Set this to a value greater than `0.0`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.Image.alpha
	Alpha *float64 `json:"alpha,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Overlay.NormalizedCoordinate
type Overlay_NormalizedCoordinate struct {
	// Normalized x coordinate.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.NormalizedCoordinate.x
	X *float64 `json:"x,omitempty"`

	// Normalized y coordinate.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Overlay.NormalizedCoordinate.y
	Y *float64 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig
type PreprocessingConfig struct {
	// Color preprocessing configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.color
	Color *PreprocessingConfig_Color `json:"color,omitempty"`

	// Denoise preprocessing configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.denoise
	Denoise *PreprocessingConfig_Denoise `json:"denoise,omitempty"`

	// Deblock preprocessing configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.deblock
	Deblock *PreprocessingConfig_Deblock `json:"deblock,omitempty"`

	// Audio preprocessing configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.audio
	Audio *PreprocessingConfig_Audio `json:"audio,omitempty"`

	// Specify the video cropping configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.crop
	Crop *PreprocessingConfig_Crop `json:"crop,omitempty"`

	// Specify the video pad filter configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.pad
	Pad *PreprocessingConfig_Pad `json:"pad,omitempty"`

	// Specify the video deinterlace configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.deinterlace
	Deinterlace *PreprocessingConfig_Deinterlace `json:"deinterlace,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Audio
type PreprocessingConfig_Audio struct {
	// Specify audio loudness normalization in loudness units relative to full
	//  scale (LUFS). Enter a value between -24 and 0 (the default), where:
	//
	//  *   -24 is the Advanced Television Systems Committee (ATSC A/85) standard
	//  *   -23 is the EU R128 broadcast standard
	//  *   -19 is the prior standard for online mono audio
	//  *   -18 is the ReplayGain standard
	//  *   -16 is the prior standard for stereo audio
	//  *   -14 is the new online audio standard recommended by Spotify, as well
	//      as Amazon Echo
	//  *   0 disables normalization
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Audio.lufs
	Lufs *float64 `json:"lufs,omitempty"`

	// Enable boosting high frequency components. The default is `false`.
	//
	//  **Note:** This field is not supported.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Audio.high_boost
	HighBoost *bool `json:"highBoost,omitempty"`

	// Enable boosting low frequency components. The default is `false`.
	//
	//  **Note:** This field is not supported.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Audio.low_boost
	LowBoost *bool `json:"lowBoost,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Color
type PreprocessingConfig_Color struct {
	// Control color saturation of the video. Enter a value between -1 and 1,
	//  where -1 is fully desaturated and 1 is maximum saturation. 0 is no
	//  change. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Color.saturation
	Saturation *float64 `json:"saturation,omitempty"`

	// Control black and white contrast of the video. Enter a value between -1
	//  and 1, where -1 is minimum contrast and 1 is maximum contrast. 0 is no
	//  change. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Color.contrast
	Contrast *float64 `json:"contrast,omitempty"`

	// Control brightness of the video. Enter a value between -1 and 1, where -1
	//  is minimum brightness and 1 is maximum brightness. 0 is no change. The
	//  default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Color.brightness
	Brightness *float64 `json:"brightness,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Crop
type PreprocessingConfig_Crop struct {
	// The number of pixels to crop from the top. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Crop.top_pixels
	TopPixels *int32 `json:"topPixels,omitempty"`

	// The number of pixels to crop from the bottom. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Crop.bottom_pixels
	BottomPixels *int32 `json:"bottomPixels,omitempty"`

	// The number of pixels to crop from the left. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Crop.left_pixels
	LeftPixels *int32 `json:"leftPixels,omitempty"`

	// The number of pixels to crop from the right. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Crop.right_pixels
	RightPixels *int32 `json:"rightPixels,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Deblock
type PreprocessingConfig_Deblock struct {
	// Set strength of the deblocker. Enter a value between 0 and 1. The higher
	//  the value, the stronger the block removal. 0 is no deblocking. The
	//  default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deblock.strength
	Strength *float64 `json:"strength,omitempty"`

	// Enable deblocker. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deblock.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace
type PreprocessingConfig_Deinterlace struct {
	// Specifies the Yet Another Deinterlacing Filter Configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.yadif
	Yadif *PreprocessingConfig_Deinterlace_YadifConfig `json:"yadif,omitempty"`

	// Specifies the Bob Weaver Deinterlacing Filter Configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.bwdif
	Bwdif *PreprocessingConfig_Deinterlace_BwdifConfig `json:"bwdif,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.BwdifConfig
type PreprocessingConfig_Deinterlace_BwdifConfig struct {
	// Specifies the deinterlacing mode to adopt.
	//  The default is `send_frame`.
	//  Supported values:
	//
	//  - `send_frame`: Output one frame for each frame
	//  - `send_field`: Output one frame for each field
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.BwdifConfig.mode
	Mode *string `json:"mode,omitempty"`

	// The picture field parity assumed for the input interlaced video.
	//  The default is `auto`.
	//  Supported values:
	//
	//  - `tff`: Assume the top field is first
	//  - `bff`: Assume the bottom field is first
	//  - `auto`: Enable automatic detection of field parity
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.BwdifConfig.parity
	Parity *string `json:"parity,omitempty"`

	// Deinterlace all frames rather than just the frames identified as
	//  interlaced. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.BwdifConfig.deinterlace_all_frames
	DeinterlaceAllFrames *bool `json:"deinterlaceAllFrames,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.YadifConfig
type PreprocessingConfig_Deinterlace_YadifConfig struct {
	// Specifies the deinterlacing mode to adopt.
	//  The default is `send_frame`.
	//  Supported values:
	//
	//  - `send_frame`: Output one frame for each frame
	//  - `send_field`: Output one frame for each field
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.YadifConfig.mode
	Mode *string `json:"mode,omitempty"`

	// Disable spacial interlacing.
	//  The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.YadifConfig.disable_spatial_interlacing
	DisableSpatialInterlacing *bool `json:"disableSpatialInterlacing,omitempty"`

	// The picture field parity assumed for the input interlaced video.
	//  The default is `auto`.
	//  Supported values:
	//
	//  - `tff`: Assume the top field is first
	//  - `bff`: Assume the bottom field is first
	//  - `auto`: Enable automatic detection of field parity
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.YadifConfig.parity
	Parity *string `json:"parity,omitempty"`

	// Deinterlace all frames rather than just the frames identified as
	//  interlaced. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Deinterlace.YadifConfig.deinterlace_all_frames
	DeinterlaceAllFrames *bool `json:"deinterlaceAllFrames,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Denoise
type PreprocessingConfig_Denoise struct {
	// Set strength of the denoise. Enter a value between 0 and 1. The higher
	//  the value, the smoother the image. 0 is no denoising. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Denoise.strength
	Strength *float64 `json:"strength,omitempty"`

	// Set the denoiser mode. The default is `standard`.
	//
	//  Supported denoiser modes:
	//
	//  - `standard`
	//  - `grain`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Denoise.tune
	Tune *string `json:"tune,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PreprocessingConfig.Pad
type PreprocessingConfig_Pad struct {
	// The number of pixels to add to the top. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Pad.top_pixels
	TopPixels *int32 `json:"topPixels,omitempty"`

	// The number of pixels to add to the bottom. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Pad.bottom_pixels
	BottomPixels *int32 `json:"bottomPixels,omitempty"`

	// The number of pixels to add to the left. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Pad.left_pixels
	LeftPixels *int32 `json:"leftPixels,omitempty"`

	// The number of pixels to add to the right. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PreprocessingConfig.Pad.right_pixels
	RightPixels *int32 `json:"rightPixels,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PubsubDestination
type PubsubDestination struct {
	// The name of the Pub/Sub topic to publish job completion notification
	//  to. For example: `projects/{project}/topics/{topic}`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PubsubDestination.topic
	Topic *string `json:"topic,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.SegmentSettings
type SegmentSettings struct {
	// Duration of the segments in seconds. The default is `6.0s`. Note that
	//  `segmentDuration` must be greater than or equal to
	//  [`gopDuration`](#videostream), and `segmentDuration` must be divisible by
	//  [`gopDuration`](#videostream).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SegmentSettings.segment_duration
	SegmentDuration *string `json:"segmentDuration,omitempty"`

	// Required. Create an individual segment file. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SegmentSettings.individual_segments
	IndividualSegments *bool `json:"individualSegments,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.SpriteSheet
type SpriteSheet struct {
	// Format type. The default is `jpeg`.
	//
	//  Supported formats:
	//
	//  - `jpeg`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.format
	Format *string `json:"format,omitempty"`

	// Required. File name prefix for the generated sprite sheets.
	//
	//  Each sprite sheet has an incremental 10-digit zero-padded suffix starting
	//  from 0 before the extension, such as `sprite_sheet0000000123.jpeg`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.file_prefix
	FilePrefix *string `json:"filePrefix,omitempty"`

	// Required. The width of sprite in pixels. Must be an even integer. To
	//  preserve the source aspect ratio, set the
	//  [SpriteSheet.sprite_width_pixels][google.cloud.video.transcoder.v1.SpriteSheet.sprite_width_pixels]
	//  field or the
	//  [SpriteSheet.sprite_height_pixels][google.cloud.video.transcoder.v1.SpriteSheet.sprite_height_pixels]
	//  field, but not both (the API will automatically calculate the missing
	//  field).
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the width, in pixels, per the horizontal ASR. The API calculates
	//  the height per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.sprite_width_pixels
	SpriteWidthPixels *int32 `json:"spriteWidthPixels,omitempty"`

	// Required. The height of sprite in pixels. Must be an even integer. To
	//  preserve the source aspect ratio, set the
	//  [SpriteSheet.sprite_height_pixels][google.cloud.video.transcoder.v1.SpriteSheet.sprite_height_pixels]
	//  field or the
	//  [SpriteSheet.sprite_width_pixels][google.cloud.video.transcoder.v1.SpriteSheet.sprite_width_pixels]
	//  field, but not both (the API will automatically calculate the missing
	//  field).
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the height, in pixels, per the horizontal ASR. The API calculates
	//  the width per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.sprite_height_pixels
	SpriteHeightPixels *int32 `json:"spriteHeightPixels,omitempty"`

	// The maximum number of sprites per row in a sprite sheet. The default is 0,
	//  which indicates no maximum limit.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.column_count
	ColumnCount *int32 `json:"columnCount,omitempty"`

	// The maximum number of rows per sprite sheet. When the sprite sheet is full,
	//  a new sprite sheet is created. The default is 0, which indicates no maximum
	//  limit.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.row_count
	RowCount *int32 `json:"rowCount,omitempty"`

	// Start time in seconds, relative to the output file timeline. Determines the
	//  first sprite to pick. The default is `0s`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`

	// End time in seconds, relative to the output file timeline. When
	//  `end_time_offset` is not specified, the sprites are generated until the end
	//  of the output file.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.end_time_offset
	EndTimeOffset *string `json:"endTimeOffset,omitempty"`

	// Total number of sprites. Create the specified number of sprites
	//  distributed evenly across the timeline of the output media. The default
	//  is 100.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.total_count
	TotalCount *int32 `json:"totalCount,omitempty"`

	// Starting from `0s`, create sprites at regular intervals. Specify the
	//  interval value in seconds.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.interval
	Interval *string `json:"interval,omitempty"`

	// The quality of the generated sprite sheet. Enter a value between 1
	//  and 100, where 1 is the lowest quality and 100 is the highest quality.
	//  The default is 100. A high quality value corresponds to a low image data
	//  compression ratio.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.SpriteSheet.quality
	Quality *int32 `json:"quality,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.TextStream
type TextStream struct {
	// The codec for this text stream. The default is `webvtt`.
	//
	//  Supported text codecs:
	//
	//  - `srt`
	//  - `ttml`
	//  - `cea608`
	//  - `cea708`
	//  - `webvtt`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.codec
	Codec *string `json:"codec,omitempty"`

	// The BCP-47 language code, such as `en-US` or `sr-Latn`. For more
	//  information, see
	//  https://www.unicode.org/reports/tr35/#Unicode_locale_identifier. Not
	//  supported in MP4 files.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The mapping for the `Job.edit_list` atoms with text `EditAtom.inputs`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.mapping
	Mapping []TextStream_TextMapping `json:"mapping,omitempty"`

	// The name for this particular text stream that
	//  will be added to the HLS/DASH manifest. Not supported in MP4 files.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.TextStream.TextMapping
type TextStream_TextMapping struct {
	// Required. The `EditAtom.key` that references atom with text inputs in the
	//  `Job.edit_list`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.TextMapping.atom_key
	AtomKey *string `json:"atomKey,omitempty"`

	// Required. The `Input.key` that identifies the input file.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.TextMapping.input_key
	InputKey *string `json:"inputKey,omitempty"`

	// Required. The zero-based index of the track in the input file.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.TextStream.TextMapping.input_track
	InputTrack *int32 `json:"inputTrack,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.VideoStream
type VideoStream struct {
	// H264 codec settings.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.h264
	H264 *VideoStream_H264CodecSettings `json:"h264,omitempty"`

	// H265 codec settings.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.h265
	H265 *VideoStream_H265CodecSettings `json:"h265,omitempty"`

	// VP9 codec settings.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.vp9
	Vp9 *VideoStream_Vp9CodecSettings `json:"vp9,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings
type VideoStream_H264CodecSettings struct {
	// The width of the video in pixels. Must be an even integer.
	//  When not specified, the width is adjusted to match the specified height
	//  and input aspect ratio. If both are omitted, the input width is used.
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the width, in pixels, per the horizontal ASR. The API calculates
	//  the height per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.width_pixels
	WidthPixels *int32 `json:"widthPixels,omitempty"`

	// The height of the video in pixels. Must be an even integer.
	//  When not specified, the height is adjusted to match the specified width
	//  and input aspect ratio. If both are omitted, the input height is used.
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the height, in pixels, per the horizontal ASR. The API calculates
	//  the width per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.height_pixels
	HeightPixels *int32 `json:"heightPixels,omitempty"`

	// Required. The target video frame rate in frames per second (FPS). Must be
	//  less than or equal to 120. Will default to the input frame rate if larger
	//  than the input frame rate. The API will generate an output FPS that is
	//  divisible by the input FPS, and smaller or equal to the target FPS. See
	//  [Calculating frame
	//  rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for
	//  more information.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.frame_rate
	FrameRate *float64 `json:"frameRate,omitempty"`

	// Required. The video bitrate in bits per second. The minimum value is
	//  1,000. The maximum value is 800,000,000.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Pixel format to use. The default is `yuv420p`.
	//
	//  Supported pixel formats:
	//
	//  - `yuv420p` pixel format
	//  - `yuv422p` pixel format
	//  - `yuv444p` pixel format
	//  - `yuv420p10` 10-bit HDR pixel format
	//  - `yuv422p10` 10-bit HDR pixel format
	//  - `yuv444p10` 10-bit HDR pixel format
	//  - `yuv420p12` 12-bit HDR pixel format
	//  - `yuv422p12` 12-bit HDR pixel format
	//  - `yuv444p12` 12-bit HDR pixel format
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.pixel_format
	PixelFormat *string `json:"pixelFormat,omitempty"`

	// Specify the `rate_control_mode`. The default is `vbr`.
	//
	//  Supported rate control modes:
	//
	//  - `vbr` - variable bitrate
	//  - `crf` - constant rate factor
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.rate_control_mode
	RateControlMode *string `json:"rateControlMode,omitempty"`

	// Target CRF level. Must be between 10 and 36, where 10 is the highest
	//  quality and 36 is the most efficient compression. The default is 21.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.crf_level
	CrfLevel *int32 `json:"crfLevel,omitempty"`

	// Specifies whether an open Group of Pictures (GOP) structure should be
	//  allowed or not. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.allow_open_gop
	AllowOpenGop *bool `json:"allowOpenGop,omitempty"`

	// Select the GOP size based on the specified frame count. Must be greater
	//  than zero.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.gop_frame_count
	GopFrameCount *int32 `json:"gopFrameCount,omitempty"`

	// Select the GOP size based on the specified duration. The default is
	//  `3s`. Note that `gopDuration` must be less than or equal to
	//  [`segmentDuration`](#SegmentSettings), and
	//  [`segmentDuration`](#SegmentSettings) must be divisible by
	//  `gopDuration`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.gop_duration
	GopDuration *string `json:"gopDuration,omitempty"`

	// Use two-pass encoding strategy to achieve better video quality.
	//  `VideoStream.rate_control_mode` must be `vbr`. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.enable_two_pass
	EnableTwoPass *bool `json:"enableTwoPass,omitempty"`

	// Size of the Video Buffering Verifier (VBV) buffer in bits. Must be
	//  greater than zero. The default is equal to `VideoStream.bitrate_bps`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.vbv_size_bits
	VbvSizeBits *int32 `json:"vbvSizeBits,omitempty"`

	// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.
	//  Must be greater than zero. The default is equal to 90% of
	//  `VideoStream.vbv_size_bits`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.vbv_fullness_bits
	VbvFullnessBits *int32 `json:"vbvFullnessBits,omitempty"`

	// The entropy coder to use. The default is `cabac`.
	//
	//  Supported entropy coders:
	//
	//  - `cavlc`
	//  - `cabac`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.entropy_coder
	EntropyCoder *string `json:"entropyCoder,omitempty"`

	// Allow B-pyramid for reference frame selection. This may not be supported
	//  on all decoders. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.b_pyramid
	BPyramid *bool `json:"bPyramid,omitempty"`

	// The number of consecutive B-frames. Must be greater than or equal to
	//  zero. Must be less than `VideoStream.gop_frame_count` if set. The default
	//  is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.b_frame_count
	BFrameCount *int32 `json:"bFrameCount,omitempty"`

	// Specify the intensity of the adaptive quantizer (AQ). Must be between 0
	//  and 1, where 0 disables the quantizer and 1 maximizes the quantizer. A
	//  higher value equals a lower bitrate but smoother image. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.aq_strength
	AqStrength *float64 `json:"aqStrength,omitempty"`

	// Enforces the specified codec profile. The following profiles are
	//  supported:
	//
	//  *   `baseline`
	//  *   `main`
	//  *   `high` (default)
	//
	//  The available options are
	//  [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.264#Tune).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `H264CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.profile
	Profile *string `json:"profile,omitempty"`

	// Enforces the specified codec tune. The available options are
	//  [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.264#Tune).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `H264CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.tune
	Tune *string `json:"tune,omitempty"`

	// Enforces the specified codec preset. The default is `veryfast`. The
	//  available options are
	//  [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.264#Preset).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `H264CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H264CodecSettings.preset
	Preset *string `json:"preset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings
type VideoStream_H265CodecSettings struct {
	// The width of the video in pixels. Must be an even integer.
	//  When not specified, the width is adjusted to match the specified height
	//  and input aspect ratio. If both are omitted, the input width is used.
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the width, in pixels, per the horizontal ASR. The API calculates
	//  the height per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.width_pixels
	WidthPixels *int32 `json:"widthPixels,omitempty"`

	// The height of the video in pixels. Must be an even integer.
	//  When not specified, the height is adjusted to match the specified width
	//  and input aspect ratio. If both are omitted, the input height is used.
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the height, in pixels, per the horizontal ASR. The API calculates
	//  the width per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.height_pixels
	HeightPixels *int32 `json:"heightPixels,omitempty"`

	// Required. The target video frame rate in frames per second (FPS). Must be
	//  less than or equal to 120. Will default to the input frame rate if larger
	//  than the input frame rate. The API will generate an output FPS that is
	//  divisible by the input FPS, and smaller or equal to the target FPS. See
	//  [Calculating frame
	//  rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for
	//  more information.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.frame_rate
	FrameRate *float64 `json:"frameRate,omitempty"`

	// Required. The video bitrate in bits per second. The minimum value is
	//  1,000. The maximum value is 800,000,000.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Pixel format to use. The default is `yuv420p`.
	//
	//  Supported pixel formats:
	//
	//  - `yuv420p` pixel format
	//  - `yuv422p` pixel format
	//  - `yuv444p` pixel format
	//  - `yuv420p10` 10-bit HDR pixel format
	//  - `yuv422p10` 10-bit HDR pixel format
	//  - `yuv444p10` 10-bit HDR pixel format
	//  - `yuv420p12` 12-bit HDR pixel format
	//  - `yuv422p12` 12-bit HDR pixel format
	//  - `yuv444p12` 12-bit HDR pixel format
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.pixel_format
	PixelFormat *string `json:"pixelFormat,omitempty"`

	// Specify the `rate_control_mode`. The default is `vbr`.
	//
	//  Supported rate control modes:
	//
	//  - `vbr` - variable bitrate
	//  - `crf` - constant rate factor
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.rate_control_mode
	RateControlMode *string `json:"rateControlMode,omitempty"`

	// Target CRF level. Must be between 10 and 36, where 10 is the highest
	//  quality and 36 is the most efficient compression. The default is 21.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.crf_level
	CrfLevel *int32 `json:"crfLevel,omitempty"`

	// Specifies whether an open Group of Pictures (GOP) structure should be
	//  allowed or not. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.allow_open_gop
	AllowOpenGop *bool `json:"allowOpenGop,omitempty"`

	// Select the GOP size based on the specified frame count. Must be greater
	//  than zero.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.gop_frame_count
	GopFrameCount *int32 `json:"gopFrameCount,omitempty"`

	// Select the GOP size based on the specified duration. The default is
	//  `3s`. Note that `gopDuration` must be less than or equal to
	//  [`segmentDuration`](#SegmentSettings), and
	//  [`segmentDuration`](#SegmentSettings) must be divisible by
	//  `gopDuration`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.gop_duration
	GopDuration *string `json:"gopDuration,omitempty"`

	// Use two-pass encoding strategy to achieve better video quality.
	//  `VideoStream.rate_control_mode` must be `vbr`. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.enable_two_pass
	EnableTwoPass *bool `json:"enableTwoPass,omitempty"`

	// Size of the Video Buffering Verifier (VBV) buffer in bits. Must be
	//  greater than zero. The default is equal to `VideoStream.bitrate_bps`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.vbv_size_bits
	VbvSizeBits *int32 `json:"vbvSizeBits,omitempty"`

	// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.
	//  Must be greater than zero. The default is equal to 90% of
	//  `VideoStream.vbv_size_bits`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.vbv_fullness_bits
	VbvFullnessBits *int32 `json:"vbvFullnessBits,omitempty"`

	// Allow B-pyramid for reference frame selection. This may not be supported
	//  on all decoders. The default is `false`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.b_pyramid
	BPyramid *bool `json:"bPyramid,omitempty"`

	// The number of consecutive B-frames. Must be greater than or equal to
	//  zero. Must be less than `VideoStream.gop_frame_count` if set. The default
	//  is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.b_frame_count
	BFrameCount *int32 `json:"bFrameCount,omitempty"`

	// Specify the intensity of the adaptive quantizer (AQ). Must be between 0
	//  and 1, where 0 disables the quantizer and 1 maximizes the quantizer. A
	//  higher value equals a lower bitrate but smoother image. The default is 0.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.aq_strength
	AqStrength *float64 `json:"aqStrength,omitempty"`

	// Enforces the specified codec profile. The following profiles are
	//  supported:
	//
	//  *   8-bit profiles
	//      *   `main` (default)
	//      *   `main-intra`
	//      *   `mainstillpicture`
	//  *   10-bit profiles
	//      *   `main10` (default)
	//      *   `main10-intra`
	//      *   `main422-10`
	//      *   `main422-10-intra`
	//      *   `main444-10`
	//      *   `main444-10-intra`
	//  *   12-bit profiles
	//      *   `main12` (default)
	//      *   `main12-intra`
	//      *   `main422-12`
	//      *   `main422-12-intra`
	//      *   `main444-12`
	//      *   `main444-12-intra`
	//
	//  The available options are
	//  [FFmpeg-compatible](https://x265.readthedocs.io/).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `H265CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.profile
	Profile *string `json:"profile,omitempty"`

	// Enforces the specified codec tune. The available options are
	//  [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.265).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `H265CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.tune
	Tune *string `json:"tune,omitempty"`

	// Enforces the specified codec preset. The default is `veryfast`. The
	//  available options are
	//  [FFmpeg-compatible](https://trac.ffmpeg.org/wiki/Encode/H.265).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `H265CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.H265CodecSettings.preset
	Preset *string `json:"preset,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings
type VideoStream_Vp9CodecSettings struct {
	// The width of the video in pixels. Must be an even integer.
	//  When not specified, the width is adjusted to match the specified height
	//  and input aspect ratio. If both are omitted, the input width is used.
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the width, in pixels, per the horizontal ASR. The API calculates
	//  the height per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.width_pixels
	WidthPixels *int32 `json:"widthPixels,omitempty"`

	// The height of the video in pixels. Must be an even integer.
	//  When not specified, the height is adjusted to match the specified width
	//  and input aspect ratio. If both are omitted, the input height is used.
	//
	//  For portrait videos that contain horizontal ASR and rotation metadata,
	//  provide the height, in pixels, per the horizontal ASR. The API calculates
	//  the width per the horizontal ASR. The API detects any rotation metadata
	//  and swaps the requested height and width for the output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.height_pixels
	HeightPixels *int32 `json:"heightPixels,omitempty"`

	// Required. The target video frame rate in frames per second (FPS). Must be
	//  less than or equal to 120. Will default to the input frame rate if larger
	//  than the input frame rate. The API will generate an output FPS that is
	//  divisible by the input FPS, and smaller or equal to the target FPS. See
	//  [Calculating frame
	//  rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for
	//  more information.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.frame_rate
	FrameRate *float64 `json:"frameRate,omitempty"`

	// Required. The video bitrate in bits per second. The minimum value is
	//  1,000. The maximum value is 480,000,000.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Pixel format to use. The default is `yuv420p`.
	//
	//  Supported pixel formats:
	//
	//  - `yuv420p` pixel format
	//  - `yuv422p` pixel format
	//  - `yuv444p` pixel format
	//  - `yuv420p10` 10-bit HDR pixel format
	//  - `yuv422p10` 10-bit HDR pixel format
	//  - `yuv444p10` 10-bit HDR pixel format
	//  - `yuv420p12` 12-bit HDR pixel format
	//  - `yuv422p12` 12-bit HDR pixel format
	//  - `yuv444p12` 12-bit HDR pixel format
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.pixel_format
	PixelFormat *string `json:"pixelFormat,omitempty"`

	// Specify the `rate_control_mode`. The default is `vbr`.
	//
	//  Supported rate control modes:
	//
	//  - `vbr` - variable bitrate
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.rate_control_mode
	RateControlMode *string `json:"rateControlMode,omitempty"`

	// Target CRF level. Must be between 10 and 36, where 10 is the highest
	//  quality and 36 is the most efficient compression. The default is 21.
	//
	//  **Note:** This field is not supported.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.crf_level
	CrfLevel *int32 `json:"crfLevel,omitempty"`

	// Select the GOP size based on the specified frame count. Must be greater
	//  than zero.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.gop_frame_count
	GopFrameCount *int32 `json:"gopFrameCount,omitempty"`

	// Select the GOP size based on the specified duration. The default is
	//  `3s`. Note that `gopDuration` must be less than or equal to
	//  [`segmentDuration`](#SegmentSettings), and
	//  [`segmentDuration`](#SegmentSettings) must be divisible by
	//  `gopDuration`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.gop_duration
	GopDuration *string `json:"gopDuration,omitempty"`

	// Enforces the specified codec profile. The following profiles are
	//  supported:
	//
	//  *   `profile0` (default)
	//  *   `profile1`
	//  *   `profile2`
	//  *   `profile3`
	//
	//  The available options are
	//  [WebM-compatible](https://www.webmproject.org/vp9/profiles/).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the `Vp9CodecSettings`
	//  message.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.VideoStream.Vp9CodecSettings.profile
	Profile *string `json:"profile,omitempty"`
}
