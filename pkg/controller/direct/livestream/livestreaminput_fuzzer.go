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
// proto.message: google.cloud.video.livestream.v1.Input
// api.group: livestream.cnrm.cloud.google.com

package livestream

import (
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(livestreamInputFuzzer())
}

func livestreamInputFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Input{},
		LiveStreamInputSpec_FromProto, LiveStreamInputSpec_ToProto,
		LiveStreamInputObservedState_FromProto, LiveStreamInputObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields to fuzz
	f.SpecField(".labels")
	f.SpecField(".type")
	f.SpecField(".tier")

	f.SpecField(".preprocessing_config")
	f.SpecField(".preprocessing_config.audio")
	f.SpecField(".preprocessing_config.audio.lufs")
	f.SpecField(".preprocessing_config.crop")
	f.SpecField(".preprocessing_config.crop.top_pixels")
	f.SpecField(".preprocessing_config.crop.bottom_pixels")
	f.SpecField(".preprocessing_config.crop.left_pixels")
	f.SpecField(".preprocessing_config.crop.right_pixels")
	f.SpecField(".preprocessing_config.pad")
	f.SpecField(".preprocessing_config.pad.top_pixels")
	f.SpecField(".preprocessing_config.pad.bottom_pixels")
	f.SpecField(".preprocessing_config.pad.left_pixels")
	f.SpecField(".preprocessing_config.pad.right_pixels")

	f.SpecField(".security_rules")
	f.SpecField(".security_rules.ip_ranges")

	// Observed state fields to fuzz
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".uri")

	f.StatusField(".input_stream_property")
	f.StatusField(".input_stream_property.last_establish_time")
	f.StatusField(".input_stream_property.video_streams")
	f.StatusField(".input_stream_property.video_streams.index")
	f.StatusField(".input_stream_property.video_streams.video_format")
	f.StatusField(".input_stream_property.video_streams.video_format.codec")
	f.StatusField(".input_stream_property.video_streams.video_format.width_pixels")
	f.StatusField(".input_stream_property.video_streams.video_format.height_pixels")
	f.StatusField(".input_stream_property.video_streams.video_format.frame_rate")
	f.StatusField(".input_stream_property.audio_streams")
	f.StatusField(".input_stream_property.audio_streams.index")
	f.StatusField(".input_stream_property.audio_streams.audio_format")
	f.StatusField(".input_stream_property.audio_streams.audio_format.codec")
	f.StatusField(".input_stream_property.audio_streams.audio_format.channel_count")
	f.StatusField(".input_stream_property.audio_streams.audio_format.channel_layout")

	return f
}
