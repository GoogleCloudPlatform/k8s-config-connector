// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.video.livestream.v1.Channel
// api.group: livestream.cnrm.cloud.google.com

package livestream

import (
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(livestreamChannelFuzzer())
}

func livestreamChannelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Channel{},
		LiveStreamChannelSpec_FromProto, LiveStreamChannelSpec_ToProto,
		LiveStreamChannelObservedState_FromProto, LiveStreamChannelObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz
	f.SpecField(".labels")

	f.SpecField(".input_attachments")
	f.SpecField(".input_attachments.key")
	f.SpecField(".input_attachments.input")
	f.SpecField(".input_attachments.automatic_failover")
	f.SpecField(".input_attachments.automatic_failover.input_keys")

	f.SpecField(".output")
	f.SpecField(".output.uri")

	f.SpecField(".elementary_streams")
	f.SpecField(".elementary_streams.key")
	f.SpecField(".elementary_streams.video_stream")
	f.SpecField(".elementary_streams.video_stream.h264")
	f.SpecField(".elementary_streams.video_stream.h264.width_pixels")
	f.SpecField(".elementary_streams.video_stream.h264.height_pixels")
	f.SpecField(".elementary_streams.video_stream.h264.frame_rate")
	f.SpecField(".elementary_streams.video_stream.h264.bitrate_bps")
	f.SpecField(".elementary_streams.video_stream.h264.allow_open_gop")
	f.SpecField(".elementary_streams.video_stream.h264.gop_frame_count")
	f.SpecField(".elementary_streams.video_stream.h264.gop_duration")
	f.SpecField(".elementary_streams.video_stream.h264.entropy_coder")
	f.SpecField(".elementary_streams.video_stream.h264.b_frame_count")
	f.SpecField(".elementary_streams.video_stream.h264.aq_strength")
	f.SpecField(".elementary_streams.video_stream.h264.profile")
	f.SpecField(".elementary_streams.video_stream.h264.tune")
	f.SpecField(".elementary_streams.video_stream.h264.b_pyramid")
	f.SpecField(".elementary_streams.video_stream.h264.vbv_size_bits")
	f.SpecField(".elementary_streams.video_stream.h264.vbv_fullness_bits")
	f.Unimplemented_NotYetTriaged(".elementary_streams[].video_stream.h265")

	f.SpecField(".elementary_streams.audio_stream")
	f.SpecField(".elementary_streams.audio_stream.transmux")
	f.SpecField(".elementary_streams.audio_stream.codec")
	f.SpecField(".elementary_streams.audio_stream.bitrate_bps")
	f.SpecField(".elementary_streams.audio_stream.channel_count")
	f.SpecField(".elementary_streams.audio_stream.channel_layout")
	f.SpecField(".elementary_streams.audio_stream.mapping")
	f.SpecField(".elementary_streams.audio_stream.mapping.input_key")
	f.SpecField(".elementary_streams.audio_stream.mapping.input_track")
	f.SpecField(".elementary_streams.audio_stream.mapping.input_channel")
	f.SpecField(".elementary_streams.audio_stream.mapping.output_channel")
	f.SpecField(".elementary_streams.audio_stream.mapping.gain_db")
	f.SpecField(".elementary_streams.audio_stream.sample_rate_hertz")

	f.SpecField(".elementary_streams.text_stream")
	f.SpecField(".elementary_streams.text_stream.codec")
	f.Unimplemented_NotYetTriaged(".elementary_streams[].text_stream.language_code")
	f.Unimplemented_NotYetTriaged(".elementary_streams[].text_stream.display_name")
	f.Unimplemented_NotYetTriaged(".elementary_streams[].text_stream.output_cea_channel")
	f.Unimplemented_NotYetTriaged(".elementary_streams[].text_stream.mapping")

	f.SpecField(".mux_streams")
	f.SpecField(".mux_streams.key")
	f.SpecField(".mux_streams.container")
	f.SpecField(".mux_streams.elementary_streams")
	f.SpecField(".mux_streams.segment_settings")
	f.SpecField(".mux_streams.segment_settings.segment_duration")
	f.SpecField(".mux_streams.encryption_id")

	f.SpecField(".manifests")
	f.SpecField(".manifests.key")
	f.SpecField(".manifests.file_name")
	f.SpecField(".manifests.type")
	f.SpecField(".manifests.mux_streams")
	f.SpecField(".manifests.max_segment_count")
	f.SpecField(".manifests.segment_keep_duration")
	f.SpecField(".manifests.use_timecode_as_timeline")

	f.SpecField(".sprite_sheets")
	f.SpecField(".sprite_sheets.format")
	f.SpecField(".sprite_sheets.file_prefix")
	f.SpecField(".sprite_sheets.total_count")
	f.SpecField(".sprite_sheets.density")
	f.SpecField(".sprite_sheets.sprite_width_pixels")
	f.SpecField(".sprite_sheets.sprite_height_pixels")
	f.SpecField(".sprite_sheets.column_count")
	f.SpecField(".sprite_sheets.row_count")
	f.SpecField(".sprite_sheets.interval")
	f.SpecField(".sprite_sheets.quality")

	f.SpecField(".log_config")
	f.SpecField(".log_config.log_severity")

	f.SpecField(".timecode_config")
	f.SpecField(".timecode_config.source")
	f.SpecField(".timecode_config.utc_offset")
	f.SpecField(".timecode_config.time_zone")
	f.SpecField(".timecode_config.time_zone.id")
	f.SpecField(".timecode_config.time_zone.version")

	f.SpecField(".encryptions")
	f.SpecField(".encryptions.id")
	f.SpecField(".encryptions.secret_manager_key_source")
	f.SpecField(".encryptions.secret_manager_key_source.secret_version")
	f.SpecField(".encryptions.drm_systems")
	f.SpecField(".encryptions.drm_systems.widevine")
	f.SpecField(".encryptions.drm_systems.fairplay")
	f.SpecField(".encryptions.drm_systems.playready")
	f.SpecField(".encryptions.drm_systems.clearkey")
	f.SpecField(".encryptions.aes128")
	f.SpecField(".encryptions.sample_aes")
	f.SpecField(".encryptions.mpeg_cenc")
	f.SpecField(".encryptions.mpeg_cenc.scheme")

	f.SpecField(".input_config")
	f.SpecField(".input_config.input_switch_mode")

	f.SpecField(".retention_config")
	f.SpecField(".retention_config.retention_window_duration")

	f.SpecField(".static_overlays")
	f.SpecField(".static_overlays.asset")
	f.SpecField(".static_overlays.resolution")
	f.SpecField(".static_overlays.resolution.w")
	f.SpecField(".static_overlays.resolution.h")
	f.SpecField(".static_overlays.position")
	f.SpecField(".static_overlays.position.x")
	f.SpecField(".static_overlays.position.y")
	f.SpecField(".static_overlays.opacity")

	// Status / Observed state fields to fuzz
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".active_input")
	f.StatusField(".streaming_state")
	f.StatusField(".streaming_error")
	f.StatusField(".streaming_error.code")
	f.StatusField(".streaming_error.message")
	f.Unimplemented_NotYetTriaged(".streaming_error.details")

	// Proto fields not exposed in KRM
	f.Unimplemented_NotYetTriaged(".distribution_streams")
	f.Unimplemented_NotYetTriaged(".distributions")
	f.Unimplemented_NotYetTriaged(".auto_transcription_config")

	return f
}
