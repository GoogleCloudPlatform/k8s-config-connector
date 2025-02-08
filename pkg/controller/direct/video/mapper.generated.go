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

package video

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AudioFormat_FromProto(mapCtx *direct.MapContext, in *pb.AudioFormat) *krm.AudioFormat {
	if in == nil {
		return nil
	}
	out := &krm.AudioFormat{}
	out.Codec = direct.LazyPtr(in.GetCodec())
	out.ChannelCount = direct.LazyPtr(in.GetChannelCount())
	out.ChannelLayout = in.ChannelLayout
	return out
}
func AudioFormat_ToProto(mapCtx *direct.MapContext, in *krm.AudioFormat) *pb.AudioFormat {
	if in == nil {
		return nil
	}
	out := &pb.AudioFormat{}
	out.Codec = direct.ValueOf(in.Codec)
	out.ChannelCount = direct.ValueOf(in.ChannelCount)
	out.ChannelLayout = in.ChannelLayout
	return out
}
func AudioStreamProperty_FromProto(mapCtx *direct.MapContext, in *pb.AudioStreamProperty) *krm.AudioStreamProperty {
	if in == nil {
		return nil
	}
	out := &krm.AudioStreamProperty{}
	out.Index = direct.LazyPtr(in.GetIndex())
	out.AudioFormat = AudioFormat_FromProto(mapCtx, in.GetAudioFormat())
	return out
}
func AudioStreamProperty_ToProto(mapCtx *direct.MapContext, in *krm.AudioStreamProperty) *pb.AudioStreamProperty {
	if in == nil {
		return nil
	}
	out := &pb.AudioStreamProperty{}
	out.Index = direct.ValueOf(in.Index)
	out.AudioFormat = AudioFormat_ToProto(mapCtx, in.AudioFormat)
	return out
}
func Input_FromProto(mapCtx *direct.MapContext, in *pb.Input) *krm.Input {
	if in == nil {
		return nil
	}
	out := &krm.Input{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	// MISSING: URI
	out.PreprocessingConfig = PreprocessingConfig_FromProto(mapCtx, in.GetPreprocessingConfig())
	out.SecurityRules = Input_SecurityRule_FromProto(mapCtx, in.GetSecurityRules())
	// MISSING: InputStreamProperty
	return out
}
func Input_ToProto(mapCtx *direct.MapContext, in *krm.Input) *pb.Input {
	if in == nil {
		return nil
	}
	out := &pb.Input{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.Input_Type](mapCtx, in.Type)
	out.Tier = direct.Enum_ToProto[pb.Input_Tier](mapCtx, in.Tier)
	// MISSING: URI
	out.PreprocessingConfig = PreprocessingConfig_ToProto(mapCtx, in.PreprocessingConfig)
	out.SecurityRules = Input_SecurityRule_ToProto(mapCtx, in.SecurityRules)
	// MISSING: InputStreamProperty
	return out
}
func InputObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Input) *krm.InputObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InputObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Type
	// MISSING: Tier
	out.URI = direct.LazyPtr(in.GetUri())
	// MISSING: PreprocessingConfig
	// MISSING: SecurityRules
	out.InputStreamProperty = InputStreamProperty_FromProto(mapCtx, in.GetInputStreamProperty())
	return out
}
func InputObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InputObservedState) *pb.Input {
	if in == nil {
		return nil
	}
	out := &pb.Input{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Type
	// MISSING: Tier
	out.Uri = direct.ValueOf(in.URI)
	// MISSING: PreprocessingConfig
	// MISSING: SecurityRules
	out.InputStreamProperty = InputStreamProperty_ToProto(mapCtx, in.InputStreamProperty)
	return out
}
func InputStreamProperty_FromProto(mapCtx *direct.MapContext, in *pb.InputStreamProperty) *krm.InputStreamProperty {
	if in == nil {
		return nil
	}
	out := &krm.InputStreamProperty{}
	out.LastEstablishTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastEstablishTime())
	out.VideoStreams = direct.Slice_FromProto(mapCtx, in.VideoStreams, VideoStreamProperty_FromProto)
	out.AudioStreams = direct.Slice_FromProto(mapCtx, in.AudioStreams, AudioStreamProperty_FromProto)
	return out
}
func InputStreamProperty_ToProto(mapCtx *direct.MapContext, in *krm.InputStreamProperty) *pb.InputStreamProperty {
	if in == nil {
		return nil
	}
	out := &pb.InputStreamProperty{}
	out.LastEstablishTime = direct.StringTimestamp_ToProto(mapCtx, in.LastEstablishTime)
	out.VideoStreams = direct.Slice_ToProto(mapCtx, in.VideoStreams, VideoStreamProperty_ToProto)
	out.AudioStreams = direct.Slice_ToProto(mapCtx, in.AudioStreams, AudioStreamProperty_ToProto)
	return out
}
func Input_SecurityRule_FromProto(mapCtx *direct.MapContext, in *pb.Input_SecurityRule) *krm.Input_SecurityRule {
	if in == nil {
		return nil
	}
	out := &krm.Input_SecurityRule{}
	out.IPRanges = in.IpRanges
	return out
}
func Input_SecurityRule_ToProto(mapCtx *direct.MapContext, in *krm.Input_SecurityRule) *pb.Input_SecurityRule {
	if in == nil {
		return nil
	}
	out := &pb.Input_SecurityRule{}
	out.IpRanges = in.IPRanges
	return out
}
func PreprocessingConfig_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig) *krm.PreprocessingConfig {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig{}
	out.Audio = PreprocessingConfig_Audio_FromProto(mapCtx, in.GetAudio())
	out.Crop = PreprocessingConfig_Crop_FromProto(mapCtx, in.GetCrop())
	out.Pad = PreprocessingConfig_Pad_FromProto(mapCtx, in.GetPad())
	return out
}
func PreprocessingConfig_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig) *pb.PreprocessingConfig {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig{}
	out.Audio = PreprocessingConfig_Audio_ToProto(mapCtx, in.Audio)
	out.Crop = PreprocessingConfig_Crop_ToProto(mapCtx, in.Crop)
	out.Pad = PreprocessingConfig_Pad_ToProto(mapCtx, in.Pad)
	return out
}
func PreprocessingConfig_Audio_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Audio) *krm.PreprocessingConfig_Audio {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Audio{}
	out.Lufs = direct.LazyPtr(in.GetLufs())
	return out
}
func PreprocessingConfig_Audio_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Audio) *pb.PreprocessingConfig_Audio {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Audio{}
	out.Lufs = direct.ValueOf(in.Lufs)
	return out
}
func PreprocessingConfig_Crop_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Crop) *krm.PreprocessingConfig_Crop {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Crop{}
	out.TopPixels = direct.LazyPtr(in.GetTopPixels())
	out.BottomPixels = direct.LazyPtr(in.GetBottomPixels())
	out.LeftPixels = direct.LazyPtr(in.GetLeftPixels())
	out.RightPixels = direct.LazyPtr(in.GetRightPixels())
	return out
}
func PreprocessingConfig_Crop_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Crop) *pb.PreprocessingConfig_Crop {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Crop{}
	out.TopPixels = direct.ValueOf(in.TopPixels)
	out.BottomPixels = direct.ValueOf(in.BottomPixels)
	out.LeftPixels = direct.ValueOf(in.LeftPixels)
	out.RightPixels = direct.ValueOf(in.RightPixels)
	return out
}
func PreprocessingConfig_Pad_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Pad) *krm.PreprocessingConfig_Pad {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Pad{}
	out.TopPixels = direct.LazyPtr(in.GetTopPixels())
	out.BottomPixels = direct.LazyPtr(in.GetBottomPixels())
	out.LeftPixels = direct.LazyPtr(in.GetLeftPixels())
	out.RightPixels = direct.LazyPtr(in.GetRightPixels())
	return out
}
func PreprocessingConfig_Pad_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Pad) *pb.PreprocessingConfig_Pad {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Pad{}
	out.TopPixels = direct.ValueOf(in.TopPixels)
	out.BottomPixels = direct.ValueOf(in.BottomPixels)
	out.LeftPixels = direct.ValueOf(in.LeftPixels)
	out.RightPixels = direct.ValueOf(in.RightPixels)
	return out
}
func VideoFormat_FromProto(mapCtx *direct.MapContext, in *pb.VideoFormat) *krm.VideoFormat {
	if in == nil {
		return nil
	}
	out := &krm.VideoFormat{}
	out.Codec = direct.LazyPtr(in.GetCodec())
	out.WidthPixels = direct.LazyPtr(in.GetWidthPixels())
	out.HeightPixels = direct.LazyPtr(in.GetHeightPixels())
	out.FrameRate = direct.LazyPtr(in.GetFrameRate())
	return out
}
func VideoFormat_ToProto(mapCtx *direct.MapContext, in *krm.VideoFormat) *pb.VideoFormat {
	if in == nil {
		return nil
	}
	out := &pb.VideoFormat{}
	out.Codec = direct.ValueOf(in.Codec)
	out.WidthPixels = direct.ValueOf(in.WidthPixels)
	out.HeightPixels = direct.ValueOf(in.HeightPixels)
	out.FrameRate = direct.ValueOf(in.FrameRate)
	return out
}
func VideoInputObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Input) *krm.VideoInputObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoInputObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: Tier
	// MISSING: URI
	// MISSING: PreprocessingConfig
	// MISSING: SecurityRules
	// MISSING: InputStreamProperty
	return out
}
func VideoInputObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoInputObservedState) *pb.Input {
	if in == nil {
		return nil
	}
	out := &pb.Input{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: Tier
	// MISSING: URI
	// MISSING: PreprocessingConfig
	// MISSING: SecurityRules
	// MISSING: InputStreamProperty
	return out
}
func VideoInputSpec_FromProto(mapCtx *direct.MapContext, in *pb.Input) *krm.VideoInputSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoInputSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: Tier
	// MISSING: URI
	// MISSING: PreprocessingConfig
	// MISSING: SecurityRules
	// MISSING: InputStreamProperty
	return out
}
func VideoInputSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoInputSpec) *pb.Input {
	if in == nil {
		return nil
	}
	out := &pb.Input{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: Tier
	// MISSING: URI
	// MISSING: PreprocessingConfig
	// MISSING: SecurityRules
	// MISSING: InputStreamProperty
	return out
}
func VideoStreamProperty_FromProto(mapCtx *direct.MapContext, in *pb.VideoStreamProperty) *krm.VideoStreamProperty {
	if in == nil {
		return nil
	}
	out := &krm.VideoStreamProperty{}
	out.Index = direct.LazyPtr(in.GetIndex())
	out.VideoFormat = VideoFormat_FromProto(mapCtx, in.GetVideoFormat())
	return out
}
func VideoStreamProperty_ToProto(mapCtx *direct.MapContext, in *krm.VideoStreamProperty) *pb.VideoStreamProperty {
	if in == nil {
		return nil
	}
	out := &pb.VideoStreamProperty{}
	out.Index = direct.ValueOf(in.Index)
	out.VideoFormat = VideoFormat_ToProto(mapCtx, in.VideoFormat)
	return out
}
