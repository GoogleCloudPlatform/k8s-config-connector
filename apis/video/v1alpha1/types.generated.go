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


// +kcc:proto=google.cloud.video.livestream.v1.AudioStream
type AudioStream struct {
	// Specifies whether pass through (transmuxing) is enabled or not.
	//  If set to `true`, the rest of the settings, other than `mapping`, will be
	//  ignored. The default is `false`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.transmux
	Transmux *bool `json:"transmux,omitempty"`

	// The codec for this audio stream. The default is `aac`.
	//
	//  Supported audio codecs:
	//
	//  - `aac`
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.codec
	Codec *string `json:"codec,omitempty"`

	// Required. Audio bitrate in bits per second. Must be between 1 and
	//  10,000,000.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Number of audio channels. Must be between 1 and 6. The default is 2.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.channel_count
	ChannelCount *int32 `json:"channelCount,omitempty"`

	// A list of channel names specifying layout of the audio channels.
	//  This only affects the metadata embedded in the container headers, if
	//  supported by the specified format. The default is `[fl, fr]`.
	//
	//  Supported channel names:
	//
	//  - `fl` - Front left channel
	//  - `fr` - Front right channel
	//  - `sl` - Side left channel
	//  - `sr` - Side right channel
	//  - `fc` - Front center channel
	//  - `lfe` - Low frequency
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.channel_layout
	ChannelLayout []string `json:"channelLayout,omitempty"`

	// The mapping for the input streams and audio channels.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.mapping
	Mapping []AudioStream_AudioMapping `json:"mapping,omitempty"`

	// The audio sample rate in Hertz. The default is 48000 Hertz.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.sample_rate_hertz
	SampleRateHertz *int32 `json:"sampleRateHertz,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.AudioStream.AudioMapping
type AudioStream_AudioMapping struct {
	// Required. The `Channel`
	//  [InputAttachment.key][google.cloud.video.livestream.v1.InputAttachment.key]
	//  that identifies the input that this audio mapping applies to. If an
	//  active input doesn't have an audio mapping, the primary audio track in
	//  the input stream will be selected.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.AudioMapping.input_key
	InputKey *string `json:"inputKey,omitempty"`

	// Required. The zero-based index of the track in the input stream.
	//  All [mapping][google.cloud.video.livestream.v1.AudioStream.mapping]s in
	//  the same [AudioStream][google.cloud.video.livestream.v1.AudioStream] must
	//  have the same input track.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.AudioMapping.input_track
	InputTrack *int32 `json:"inputTrack,omitempty"`

	// Required. The zero-based index of the channel in the input stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.AudioMapping.input_channel
	InputChannel *int32 `json:"inputChannel,omitempty"`

	// Required. The zero-based index of the channel in the output audio stream.
	//  Must be consistent with the
	//  [input_channel][google.cloud.video.livestream.v1.AudioStream.AudioMapping.input_channel].
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.AudioMapping.output_channel
	OutputChannel *int32 `json:"outputChannel,omitempty"`

	// Audio volume control in dB. Negative values decrease volume,
	//  positive values increase. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.AudioStream.AudioMapping.gain_db
	GainDb *float64 `json:"gainDb,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Channel
type Channel struct {
	// The resource name of the channel, in the form of:
	//  `projects/{project}/locations/{location}/channels/{channelId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.name
	Name *string `json:"name,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A list of input attachments that this channel uses.
	//  One channel can have multiple inputs as the input sources. Only one
	//  input can be selected as the input source at one time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.input_attachments
	InputAttachments []InputAttachment `json:"inputAttachments,omitempty"`

	// Required. Information about the output (that is, the Cloud Storage bucket
	//  to store the generated live stream).
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
	//  which is referred to by each MuxStream to indicate which configuration is
	//  used for that output.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.encryptions
	Encryptions []Encryption `json:"encryptions,omitempty"`

	// The configuration for input sources defined in
	//  [input_attachments][google.cloud.video.livestream.v1.Channel.input_attachments].
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.input_config
	InputConfig *InputConfig `json:"inputConfig,omitempty"`

	// Optional. Configuration for retention of output files for this channel.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.retention_config
	RetentionConfig *RetentionConfig `json:"retentionConfig,omitempty"`

	// Optional. List of static overlay images. Those images display over the
	//  output content for the whole duration of the live stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.static_overlays
	StaticOverlays []StaticOverlay `json:"staticOverlays,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Channel.Output
type Channel_Output struct {
	// URI for the output file(s). For example, `gs://my-bucket/outputs/`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.Output.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.ElementaryStream
type ElementaryStream struct {
	// A unique key for this elementary stream. The key must be 1-63
	//  characters in length. The key must begin and end with a letter (regardless
	//  of case) or a number, but can contain dashes or underscores in between.
	// +kcc:proto:field=google.cloud.video.livestream.v1.ElementaryStream.key
	Key *string `json:"key,omitempty"`

	// Encoding of a video stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.ElementaryStream.video_stream
	VideoStream *VideoStream `json:"videoStream,omitempty"`

	// Encoding of an audio stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.ElementaryStream.audio_stream
	AudioStream *AudioStream `json:"audioStream,omitempty"`

	// Encoding of a text stream. For example, closed captions or subtitles.
	// +kcc:proto:field=google.cloud.video.livestream.v1.ElementaryStream.text_stream
	TextStream *TextStream `json:"textStream,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption
type Encryption struct {
	// Required. Identifier for this set of encryption options. The ID must be
	//  1-63 characters in length. The ID must begin and end with a letter
	//  (regardless of case) or a number, but can contain dashes or underscores in
	//  between.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.id
	ID *string `json:"id,omitempty"`

	// For keys stored in Google Secret Manager.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.secret_manager_key_source
	SecretManagerKeySource *Encryption_SecretManagerSource `json:"secretManagerKeySource,omitempty"`

	// Required. Configuration for DRM systems.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.drm_systems
	DrmSystems *Encryption_DrmSystems `json:"drmSystems,omitempty"`

	// Configuration for HLS AES-128 encryption.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.aes128
	Aes128 *Encryption_Aes128Encryption `json:"aes128,omitempty"`

	// Configuration for HLS SAMPLE-AES encryption.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.sample_aes
	SampleAes *Encryption_SampleAesEncryption `json:"sampleAes,omitempty"`

	// Configuration for MPEG-Dash Common Encryption (MPEG-CENC).
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.mpeg_cenc
	MpegCenc *Encryption_MpegCommonEncryption `json:"mpegCenc,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Aes128Encryption
type Encryption_Aes128Encryption struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Clearkey
type Encryption_Clearkey struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.DrmSystems
type Encryption_DrmSystems struct {
	// Widevine configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.DrmSystems.widevine
	Widevine *Encryption_Widevine `json:"widevine,omitempty"`

	// Fairplay configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.DrmSystems.fairplay
	Fairplay *Encryption_Fairplay `json:"fairplay,omitempty"`

	// Playready configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.DrmSystems.playready
	Playready *Encryption_Playready `json:"playready,omitempty"`

	// Clearkey configuration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.DrmSystems.clearkey
	Clearkey *Encryption_Clearkey `json:"clearkey,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Fairplay
type Encryption_Fairplay struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.MpegCommonEncryption
type Encryption_MpegCommonEncryption struct {
	// Required. Specify the encryption scheme, supported schemes:
	//  - `cenc` - AES-CTR subsample
	//  - `cbcs`- AES-CBC subsample pattern
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.MpegCommonEncryption.scheme
	Scheme *string `json:"scheme,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Playready
type Encryption_Playready struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.SampleAesEncryption
type Encryption_SampleAesEncryption struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.SecretManagerSource
type Encryption_SecretManagerSource struct {
	// Required. The name of the Secret Version containing the encryption key.
	//  `projects/{project}/secrets/{secret_id}/versions/{version_number}`
	// +kcc:proto:field=google.cloud.video.livestream.v1.Encryption.SecretManagerSource.secret_version
	SecretVersion *string `json:"secretVersion,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Encryption.Widevine
type Encryption_Widevine struct {
}

// +kcc:proto=google.cloud.video.livestream.v1.InputAttachment
type InputAttachment struct {
	// A unique key for this input attachment. The key must be 1-63
	//  characters in length. The key must begin and end with a letter (regardless
	//  of case) or a number, but can contain dashes or underscores in between.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.key
	Key *string `json:"key,omitempty"`

	// The resource name of an existing input, in the form of:
	//  `projects/{project}/locations/{location}/inputs/{inputId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.input
	Input *string `json:"input,omitempty"`

	// Automatic failover configurations.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.automatic_failover
	AutomaticFailover *InputAttachment_AutomaticFailover `json:"automaticFailover,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.InputAttachment.AutomaticFailover
type InputAttachment_AutomaticFailover struct {
	// The
	//  [InputAttachment.key][google.cloud.video.livestream.v1.InputAttachment.key]s
	//  of inputs to failover to when this input is disconnected. Currently, only
	//  up to one backup input is supported.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputAttachment.AutomaticFailover.input_keys
	InputKeys []string `json:"inputKeys,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.InputConfig
type InputConfig struct {
	// Input switch mode. Default mode is `FAILOVER_PREFER_PRIMARY`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.InputConfig.input_switch_mode
	InputSwitchMode *string `json:"inputSwitchMode,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.LogConfig
type LogConfig struct {
	// The severity level of platform logging for this resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.LogConfig.log_severity
	LogSeverity *string `json:"logSeverity,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Manifest
type Manifest struct {
	// The name of the generated file. The default is `manifest` with the
	//  extension suffix corresponding to the `Manifest`
	//  [type][google.cloud.video.livestream.v1.Manifest.type]. If multiple
	//  manifests are added to the channel, each must have a unique file name.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.file_name
	FileName *string `json:"fileName,omitempty"`

	// Required. Type of the manifest, can be `HLS` or `DASH`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.type
	Type *string `json:"type,omitempty"`

	// Required. List of `MuxStream`
	//  [key][google.cloud.video.livestream.v1.MuxStream.key]s that should appear
	//  in this manifest.
	//
	//  - For HLS, either `fmp4` or `ts` mux streams can be specified but not
	//  mixed.
	//  - For DASH, only `fmp4` mux streams can be specified.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.mux_streams
	MuxStreams []string `json:"muxStreams,omitempty"`

	// Maximum number of segments that this manifest holds. Once the manifest
	//  reaches this maximum number of segments, whenever a new segment is added to
	//  the manifest, the oldest segment will be removed from the manifest.
	//  The minimum value is 3 and the default value is 5.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.max_segment_count
	MaxSegmentCount *int32 `json:"maxSegmentCount,omitempty"`

	// How long to keep a segment on the output Google Cloud Storage bucket after
	//  it is removed from the manifest. This field should be large enough to cover
	//  the manifest propagation delay. Otherwise, a player could receive 404
	//  errors while accessing segments which are listed in the manifest that the
	//  player has, but were already deleted from the output Google Cloud Storage
	//  bucket. Default value is `60s`.
	//
	//  If both segment_keep_duration and
	//  [RetentionConfig.retention_window_duration][google.cloud.video.livestream.v1.RetentionConfig.retention_window_duration]
	//  are set,
	//  [RetentionConfig.retention_window_duration][google.cloud.video.livestream.v1.RetentionConfig.retention_window_duration]
	//  is used and segment_keep_duration is ignored.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.segment_keep_duration
	SegmentKeepDuration *string `json:"segmentKeepDuration,omitempty"`

	// Whether to use the timecode, as specified in timecode config, when setting:
	//
	//  - `availabilityStartTime` attribute in DASH manifests.
	//  - `#EXT-X-PROGRAM-DATE-TIME` tag in HLS manifests.
	//
	//  If false, ignore the input timecode and use the time from system clock
	//  when the manifest is first generated. This is the default behavior.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.use_timecode_as_timeline
	UseTimecodeAsTimeline *bool `json:"useTimecodeAsTimeline,omitempty"`

	// Optional. A unique key for this manifest.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Manifest.key
	Key *string `json:"key,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.MuxStream
type MuxStream struct {
	// A unique key for this multiplexed stream. The key must be 1-63
	//  characters in length. The key must begin and end with a letter (regardless
	//  of case) or a number, but can contain dashes or underscores in between.
	// +kcc:proto:field=google.cloud.video.livestream.v1.MuxStream.key
	Key *string `json:"key,omitempty"`

	// The container format. The default is `fmp4`.
	//
	//  Supported container formats:
	//
	//  - `fmp4` - the corresponding file extension is `.m4s`
	//  - `ts` - the corresponding file extension is `.ts`
	// +kcc:proto:field=google.cloud.video.livestream.v1.MuxStream.container
	Container *string `json:"container,omitempty"`

	// List of `ElementaryStream`
	//  [key][google.cloud.video.livestream.v1.ElementaryStream.key]s multiplexed
	//  in this stream.
	//
	//  - For `fmp4` container, must contain either one video or one audio stream.
	//  - For `ts` container, must contain exactly one audio stream and up to one
	//  video stream.
	// +kcc:proto:field=google.cloud.video.livestream.v1.MuxStream.elementary_streams
	ElementaryStreams []string `json:"elementaryStreams,omitempty"`

	// Segment settings for `fmp4` and `ts`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.MuxStream.segment_settings
	SegmentSettings *SegmentSettings `json:"segmentSettings,omitempty"`

	// Identifier of the encryption configuration to use. If omitted, output
	//  will be unencrypted.
	// +kcc:proto:field=google.cloud.video.livestream.v1.MuxStream.encryption_id
	EncryptionID *string `json:"encryptionID,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.NormalizedCoordinate
type NormalizedCoordinate struct {
	// Optional. Normalized x coordinate. Valid range is [0.0, 1.0]. Default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.NormalizedCoordinate.x
	X *float64 `json:"x,omitempty"`

	// Optional. Normalized y coordinate. Valid range is [0.0, 1.0]. Default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.NormalizedCoordinate.y
	Y *float64 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.NormalizedResolution
type NormalizedResolution struct {
	// Optional. Normalized width. Valid range is [0.0, 1.0]. Default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.NormalizedResolution.w
	W *float64 `json:"w,omitempty"`

	// Optional. Normalized height. Valid range is [0.0, 1.0]. Default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.NormalizedResolution.h
	H *float64 `json:"h,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.RetentionConfig
type RetentionConfig struct {
	// The minimum duration for which the output files from the channel will
	//  remain in the output bucket. After this duration, output files are
	//  deleted asynchronously.
	//
	//  When the channel is deleted, all output files are deleted from the output
	//  bucket asynchronously.
	//
	//  If omitted or set to zero, output files will remain in the output bucket
	//  based on
	//  [Manifest.segment_keep_duration][google.cloud.video.livestream.v1.Manifest.segment_keep_duration],
	//  which defaults to 60s.
	//
	//  If both retention_window_duration and
	//  [Manifest.segment_keep_duration][google.cloud.video.livestream.v1.Manifest.segment_keep_duration]
	//  are set, retention_window_duration is used and
	//  [Manifest.segment_keep_duration][google.cloud.video.livestream.v1.Manifest.segment_keep_duration]
	//  is ignored.
	// +kcc:proto:field=google.cloud.video.livestream.v1.RetentionConfig.retention_window_duration
	RetentionWindowDuration *string `json:"retentionWindowDuration,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.SegmentSettings
type SegmentSettings struct {
	// Duration of the segments in seconds. The default is `6s`. Note that
	//  `segmentDuration` must be greater than or equal to
	//  [gop_duration][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.gop_duration],
	//  and `segmentDuration` must be divisible by
	//  [gop_duration][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.gop_duration].
	//  Valid range is [2s, 20s].
	//
	//  All [mux_streams][google.cloud.video.livestream.v1.Manifest.mux_streams] in
	//  the same manifest must have the same segment duration.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SegmentSettings.segment_duration
	SegmentDuration *string `json:"segmentDuration,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.SpriteSheet
type SpriteSheet struct {
	// Format type. The default is `jpeg`.
	//
	//  Supported formats:
	//
	//  - `jpeg`
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.format
	Format *string `json:"format,omitempty"`

	// Required. File name prefix for the generated sprite sheets. If multiple
	//  sprite sheets are added to the channel, each must have a unique file
	//  prefix.
	//  Each sprite sheet has an incremental 10-digit zero-padded suffix starting
	//  from 0 before the extension, such as `sprite_sheet0000000123.jpeg`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.file_prefix
	FilePrefix *string `json:"filePrefix,omitempty"`

	// Required. The width of the sprite in pixels. Must be an even integer.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.sprite_width_pixels
	SpriteWidthPixels *int32 `json:"spriteWidthPixels,omitempty"`

	// Required. The height of the sprite in pixels. Must be an even integer.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.sprite_height_pixels
	SpriteHeightPixels *int32 `json:"spriteHeightPixels,omitempty"`

	// The maximum number of sprites per row in a sprite sheet. Valid range is
	//  [1, 10] and the default value is 1.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.column_count
	ColumnCount *int32 `json:"columnCount,omitempty"`

	// The maximum number of rows per sprite sheet. When the sprite sheet is full,
	//  a new sprite sheet is created. Valid range is [1, 10] and the default value
	//  is 1.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.row_count
	RowCount *int32 `json:"rowCount,omitempty"`

	// Create sprites at regular intervals. Valid range is [1 second, 1 hour] and
	//  the default value is `10s`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.interval
	Interval *string `json:"interval,omitempty"`

	// The quality of the generated sprite sheet. Enter a value between 1
	//  and 100, where 1 is the lowest quality and 100 is the highest quality.
	//  The default is 100. A high quality value corresponds to a low image data
	//  compression ratio.
	// +kcc:proto:field=google.cloud.video.livestream.v1.SpriteSheet.quality
	Quality *int32 `json:"quality,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.StaticOverlay
type StaticOverlay struct {
	// Required. Asset to use for the overlaid image.
	//  The asset must be represented in the form of:
	//  `projects/{project}/locations/{location}/assets/{assetId}`.
	//  The asset's resource type must be image.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.asset
	Asset *string `json:"asset,omitempty"`

	// Optional. Normalized image resolution, based on output video resolution.
	//  Valid values are [0.0, 1.0]. To respect the original image aspect ratio,
	//  set either `w` or `h` to 0. To use the original image resolution, set both
	//  `w` and `h` to 0. The default is {0, 0}.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.resolution
	Resolution *NormalizedResolution `json:"resolution,omitempty"`

	// Optional. Position of the image in terms of normalized coordinates of the
	//  upper-left corner of the image, based on output video resolution. For
	//  example, use the x and y coordinates {0, 0} to position the top-left corner
	//  of the overlay animation in the top-left corner of the output video.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.position
	Position *NormalizedCoordinate `json:"position,omitempty"`

	// Optional. Target image opacity. Valid values are from `1.0` (solid,
	//  default) to `0.0` (transparent), exclusive. Set this to a value greater
	//  than `0.0`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.StaticOverlay.opacity
	Opacity *float64 `json:"opacity,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.TextStream
type TextStream struct {
	// Required. The codec for this text stream.
	//
	//  Supported text codecs:
	//
	//  - `cea608`
	//  - `cea708`
	// +kcc:proto:field=google.cloud.video.livestream.v1.TextStream.codec
	Codec *string `json:"codec,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.TimecodeConfig
type TimecodeConfig struct {
	// The source of the timecode that will later be used in outputs/manifests.
	//  It determines the initial timecode/timestamp (first frame) of output
	//  streams.
	// +kcc:proto:field=google.cloud.video.livestream.v1.TimecodeConfig.source
	Source *string `json:"source,omitempty"`

	// UTC offset. Must be whole seconds, between -18 hours and +18 hours.
	// +kcc:proto:field=google.cloud.video.livestream.v1.TimecodeConfig.utc_offset
	UtcOffset *string `json:"utcOffset,omitempty"`

	// Time zone e.g. "America/Los_Angeles".
	// +kcc:proto:field=google.cloud.video.livestream.v1.TimecodeConfig.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.VideoStream
type VideoStream struct {
	// H264 codec settings.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.h264
	H264 *VideoStream_H264CodecSettings `json:"h264,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings
type VideoStream_H264CodecSettings struct {
	// Required. The width of the video in pixels. Must be an even integer.
	//  Valid range is [320, 1920].
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.width_pixels
	WidthPixels *int32 `json:"widthPixels,omitempty"`

	// Required. The height of the video in pixels. Must be an even integer.
	//  Valid range is [180, 1080].
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.height_pixels
	HeightPixels *int32 `json:"heightPixels,omitempty"`

	// Required. The target video frame rate in frames per second (FPS). Must be
	//  less than or equal to 60. Will default to the input frame rate if larger
	//  than the input frame rate. The API will generate an output FPS that is
	//  divisible by the input FPS, and smaller or equal to the target FPS. See
	//  [Calculating frame
	//  rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for
	//  more information.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.frame_rate
	FrameRate *float64 `json:"frameRate,omitempty"`

	// Required. The video bitrate in bits per second. Minimum value is 10,000.
	//
	//  - For SD resolution (< 720p), must be <= 3,000,000 (3 Mbps).
	//  - For HD resolution (<= 1080p), must be <= 15,000,000 (15 Mbps).
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Specifies whether an open Group of Pictures (GOP) structure should be
	//  allowed or not. The default is `false`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.allow_open_gop
	AllowOpenGop *bool `json:"allowOpenGop,omitempty"`

	// Select the GOP size based on the specified frame count.
	//  If GOP frame count is set instead of GOP duration, GOP duration will be
	//  calculated by `gopFrameCount`/`frameRate`. The calculated GOP duration
	//  must satisfy the limitations on `gopDuration` as well.
	//  Valid range is [60, 600].
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.gop_frame_count
	GopFrameCount *int32 `json:"gopFrameCount,omitempty"`

	// Select the GOP size based on the specified duration. The default is
	//  `2s`. Note that `gopDuration` must be less than or equal to
	//  [segment_duration][google.cloud.video.livestream.v1.SegmentSettings.segment_duration],
	//  and
	//  [segment_duration][google.cloud.video.livestream.v1.SegmentSettings.segment_duration]
	//  must be divisible by `gopDuration`. Valid range is [2s, 20s].
	//
	//  All video streams in the same channel must have the same GOP size.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.gop_duration
	GopDuration *string `json:"gopDuration,omitempty"`

	// Size of the Video Buffering Verifier (VBV) buffer in bits. Must be
	//  greater than zero. The default is equal to
	//  [bitrate_bps][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.bitrate_bps].
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.vbv_size_bits
	VbvSizeBits *int32 `json:"vbvSizeBits,omitempty"`

	// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.
	//  Must be greater than zero. The default is equal to 90% of
	//  [vbv_size_bits][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.vbv_size_bits].
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.vbv_fullness_bits
	VbvFullnessBits *int32 `json:"vbvFullnessBits,omitempty"`

	// The entropy coder to use. The default is `cabac`.
	//
	//  Supported entropy coders:
	//
	//  - `cavlc`
	//  - `cabac`
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.entropy_coder
	EntropyCoder *string `json:"entropyCoder,omitempty"`

	// Allow B-pyramid for reference frame selection. This may not be supported
	//  on all decoders. The default is `false`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.b_pyramid
	BPyramid *bool `json:"bPyramid,omitempty"`

	// The number of consecutive B-frames. Must be greater than or equal to
	//  zero. Must be less than
	//  [gop_frame_count][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.gop_frame_count]
	//  if set. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.b_frame_count
	BFrameCount *int32 `json:"bFrameCount,omitempty"`

	// Specify the intensity of the adaptive quantizer (AQ). Must be between 0
	//  and 1, where 0 disables the quantizer and 1 maximizes the quantizer. A
	//  higher value equals a lower bitrate but smoother image. The default is 0.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.aq_strength
	AqStrength *float64 `json:"aqStrength,omitempty"`

	// Enforces the specified codec profile. The following profiles are
	//  supported:
	//
	//  *   `baseline`
	//  *   `main` (default)
	//  *   `high`
	//
	//  The available options are [FFmpeg-compatible Profile
	//  Options](https://trac.ffmpeg.org/wiki/Encode/H.264#Profile).
	//  Note that certain values for this field may cause the
	//  transcoder to override other fields you set in the
	//  [H264CodecSettings][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings]
	//  message.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.profile
	Profile *string `json:"profile,omitempty"`

	// Enforces the specified codec tune. The available options are
	//  [FFmpeg-compatible Encode
	//  Options](https://trac.ffmpeg.org/wiki/Encode/H.264#Tune)
	//  Note that certain values for this field may cause the transcoder to
	//  override other fields you set in the
	//  [H264CodecSettings][google.cloud.video.livestream.v1.VideoStream.H264CodecSettings]
	//  message.
	// +kcc:proto:field=google.cloud.video.livestream.v1.VideoStream.H264CodecSettings.tune
	Tune *string `json:"tune,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.TimeZone
type TimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	// +kcc:proto:field=google.type.TimeZone.id
	ID *string `json:"id,omitempty"`

	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	// +kcc:proto:field=google.type.TimeZone.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Channel
type ChannelObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The
	//  [InputAttachment.key][google.cloud.video.livestream.v1.InputAttachment.key]
	//  that serves as the current input source. The first input in the
	//  [input_attachments][google.cloud.video.livestream.v1.Channel.input_attachments]
	//  is the initial input source.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.active_input
	ActiveInput *string `json:"activeInput,omitempty"`

	// Output only. State of the streaming operation.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.streaming_state
	StreamingState *string `json:"streamingState,omitempty"`

	// Output only. A description of the reason for the streaming error. This
	//  property is always present when
	//  [streaming_state][google.cloud.video.livestream.v1.Channel.streaming_state]
	//  is
	//  [STREAMING_ERROR][google.cloud.video.livestream.v1.Channel.StreamingState.STREAMING_ERROR].
	// +kcc:proto:field=google.cloud.video.livestream.v1.Channel.streaming_error
	StreamingError *Status `json:"streamingError,omitempty"`
}
