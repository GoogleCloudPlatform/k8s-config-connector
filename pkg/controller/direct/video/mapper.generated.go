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
func AudioStream_FromProto(mapCtx *direct.MapContext, in *pb.AudioStream) *krm.AudioStream {
	if in == nil {
		return nil
	}
	out := &krm.AudioStream{}
	out.Transmux = direct.LazyPtr(in.GetTransmux())
	out.Codec = direct.LazyPtr(in.GetCodec())
	out.BitrateBps = direct.LazyPtr(in.GetBitrateBps())
	out.ChannelCount = direct.LazyPtr(in.GetChannelCount())
	out.ChannelLayout = in.ChannelLayout
	out.Mapping = direct.Slice_FromProto(mapCtx, in.Mapping, AudioStream_AudioMapping_FromProto)
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	return out
}
func AudioStream_ToProto(mapCtx *direct.MapContext, in *krm.AudioStream) *pb.AudioStream {
	if in == nil {
		return nil
	}
	out := &pb.AudioStream{}
	out.Transmux = direct.ValueOf(in.Transmux)
	out.Codec = direct.ValueOf(in.Codec)
	out.BitrateBps = direct.ValueOf(in.BitrateBps)
	out.ChannelCount = direct.ValueOf(in.ChannelCount)
	out.ChannelLayout = in.ChannelLayout
	out.Mapping = direct.Slice_ToProto(mapCtx, in.Mapping, AudioStream_AudioMapping_ToProto)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	return out
}
func AudioStream_AudioMapping_FromProto(mapCtx *direct.MapContext, in *pb.AudioStream_AudioMapping) *krm.AudioStream_AudioMapping {
	if in == nil {
		return nil
	}
	out := &krm.AudioStream_AudioMapping{}
	out.InputKey = direct.LazyPtr(in.GetInputKey())
	out.InputTrack = direct.LazyPtr(in.GetInputTrack())
	out.InputChannel = direct.LazyPtr(in.GetInputChannel())
	out.OutputChannel = direct.LazyPtr(in.GetOutputChannel())
	out.GainDb = direct.LazyPtr(in.GetGainDb())
	return out
}
func AudioStream_AudioMapping_ToProto(mapCtx *direct.MapContext, in *krm.AudioStream_AudioMapping) *pb.AudioStream_AudioMapping {
	if in == nil {
		return nil
	}
	out := &pb.AudioStream_AudioMapping{}
	out.InputKey = direct.ValueOf(in.InputKey)
	out.InputTrack = direct.ValueOf(in.InputTrack)
	out.InputChannel = direct.ValueOf(in.InputChannel)
	out.OutputChannel = direct.ValueOf(in.OutputChannel)
	out.GainDb = direct.ValueOf(in.GainDb)
	return out
}
func Channel_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.Channel {
	if in == nil {
		return nil
	}
	out := &krm.Channel{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.InputAttachments = direct.Slice_FromProto(mapCtx, in.InputAttachments, InputAttachment_FromProto)
	// MISSING: ActiveInput
	out.Output = Channel_Output_FromProto(mapCtx, in.GetOutput())
	out.ElementaryStreams = direct.Slice_FromProto(mapCtx, in.ElementaryStreams, ElementaryStream_FromProto)
	out.MuxStreams = direct.Slice_FromProto(mapCtx, in.MuxStreams, MuxStream_FromProto)
	out.Manifests = direct.Slice_FromProto(mapCtx, in.Manifests, Manifest_FromProto)
	out.SpriteSheets = direct.Slice_FromProto(mapCtx, in.SpriteSheets, SpriteSheet_FromProto)
	// MISSING: StreamingState
	// MISSING: StreamingError
	out.LogConfig = LogConfig_FromProto(mapCtx, in.GetLogConfig())
	out.TimecodeConfig = TimecodeConfig_FromProto(mapCtx, in.GetTimecodeConfig())
	out.Encryptions = direct.Slice_FromProto(mapCtx, in.Encryptions, Encryption_FromProto)
	out.InputConfig = InputConfig_FromProto(mapCtx, in.GetInputConfig())
	out.RetentionConfig = RetentionConfig_FromProto(mapCtx, in.GetRetentionConfig())
	out.StaticOverlays = direct.Slice_FromProto(mapCtx, in.StaticOverlays, StaticOverlay_FromProto)
	return out
}
func Channel_ToProto(mapCtx *direct.MapContext, in *krm.Channel) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.InputAttachments = direct.Slice_ToProto(mapCtx, in.InputAttachments, InputAttachment_ToProto)
	// MISSING: ActiveInput
	out.Output = Channel_Output_ToProto(mapCtx, in.Output)
	out.ElementaryStreams = direct.Slice_ToProto(mapCtx, in.ElementaryStreams, ElementaryStream_ToProto)
	out.MuxStreams = direct.Slice_ToProto(mapCtx, in.MuxStreams, MuxStream_ToProto)
	out.Manifests = direct.Slice_ToProto(mapCtx, in.Manifests, Manifest_ToProto)
	out.SpriteSheets = direct.Slice_ToProto(mapCtx, in.SpriteSheets, SpriteSheet_ToProto)
	// MISSING: StreamingState
	// MISSING: StreamingError
	out.LogConfig = LogConfig_ToProto(mapCtx, in.LogConfig)
	out.TimecodeConfig = TimecodeConfig_ToProto(mapCtx, in.TimecodeConfig)
	out.Encryptions = direct.Slice_ToProto(mapCtx, in.Encryptions, Encryption_ToProto)
	out.InputConfig = InputConfig_ToProto(mapCtx, in.InputConfig)
	out.RetentionConfig = RetentionConfig_ToProto(mapCtx, in.RetentionConfig)
	out.StaticOverlays = direct.Slice_ToProto(mapCtx, in.StaticOverlays, StaticOverlay_ToProto)
	return out
}
func ChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.ChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: InputAttachments
	out.ActiveInput = direct.LazyPtr(in.GetActiveInput())
	// MISSING: Output
	// MISSING: ElementaryStreams
	// MISSING: MuxStreams
	// MISSING: Manifests
	// MISSING: SpriteSheets
	out.StreamingState = direct.Enum_FromProto(mapCtx, in.GetStreamingState())
	out.StreamingError = Status_FromProto(mapCtx, in.GetStreamingError())
	// MISSING: LogConfig
	// MISSING: TimecodeConfig
	// MISSING: Encryptions
	// MISSING: InputConfig
	// MISSING: RetentionConfig
	// MISSING: StaticOverlays
	return out
}
func ChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: InputAttachments
	out.ActiveInput = direct.ValueOf(in.ActiveInput)
	// MISSING: Output
	// MISSING: ElementaryStreams
	// MISSING: MuxStreams
	// MISSING: Manifests
	// MISSING: SpriteSheets
	out.StreamingState = direct.Enum_ToProto[pb.Channel_StreamingState](mapCtx, in.StreamingState)
	out.StreamingError = Status_ToProto(mapCtx, in.StreamingError)
	// MISSING: LogConfig
	// MISSING: TimecodeConfig
	// MISSING: Encryptions
	// MISSING: InputConfig
	// MISSING: RetentionConfig
	// MISSING: StaticOverlays
	return out
}
func Channel_Output_FromProto(mapCtx *direct.MapContext, in *pb.Channel_Output) *krm.Channel_Output {
	if in == nil {
		return nil
	}
	out := &krm.Channel_Output{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Channel_Output_ToProto(mapCtx *direct.MapContext, in *krm.Channel_Output) *pb.Channel_Output {
	if in == nil {
		return nil
	}
	out := &pb.Channel_Output{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func ElementaryStream_FromProto(mapCtx *direct.MapContext, in *pb.ElementaryStream) *krm.ElementaryStream {
	if in == nil {
		return nil
	}
	out := &krm.ElementaryStream{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.VideoStream = VideoStream_FromProto(mapCtx, in.GetVideoStream())
	out.AudioStream = AudioStream_FromProto(mapCtx, in.GetAudioStream())
	out.TextStream = TextStream_FromProto(mapCtx, in.GetTextStream())
	return out
}
func ElementaryStream_ToProto(mapCtx *direct.MapContext, in *krm.ElementaryStream) *pb.ElementaryStream {
	if in == nil {
		return nil
	}
	out := &pb.ElementaryStream{}
	out.Key = direct.ValueOf(in.Key)
	if oneof := VideoStream_ToProto(mapCtx, in.VideoStream); oneof != nil {
		out.ElementaryStream = &pb.ElementaryStream_VideoStream{VideoStream: oneof}
	}
	if oneof := AudioStream_ToProto(mapCtx, in.AudioStream); oneof != nil {
		out.ElementaryStream = &pb.ElementaryStream_AudioStream{AudioStream: oneof}
	}
	if oneof := TextStream_ToProto(mapCtx, in.TextStream); oneof != nil {
		out.ElementaryStream = &pb.ElementaryStream_TextStream{TextStream: oneof}
	}
	return out
}
func Encryption_FromProto(mapCtx *direct.MapContext, in *pb.Encryption) *krm.Encryption {
	if in == nil {
		return nil
	}
	out := &krm.Encryption{}
	out.ID = direct.LazyPtr(in.GetId())
	out.SecretManagerKeySource = Encryption_SecretManagerSource_FromProto(mapCtx, in.GetSecretManagerKeySource())
	out.DrmSystems = Encryption_DrmSystems_FromProto(mapCtx, in.GetDrmSystems())
	out.Aes128 = Encryption_Aes128Encryption_FromProto(mapCtx, in.GetAes128())
	out.SampleAes = Encryption_SampleAesEncryption_FromProto(mapCtx, in.GetSampleAes())
	out.MpegCenc = Encryption_MpegCommonEncryption_FromProto(mapCtx, in.GetMpegCenc())
	return out
}
func Encryption_ToProto(mapCtx *direct.MapContext, in *krm.Encryption) *pb.Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Encryption{}
	out.Id = direct.ValueOf(in.ID)
	if oneof := Encryption_SecretManagerSource_ToProto(mapCtx, in.SecretManagerKeySource); oneof != nil {
		out.SecretSource = &pb.Encryption_SecretManagerKeySource{SecretManagerKeySource: oneof}
	}
	out.DrmSystems = Encryption_DrmSystems_ToProto(mapCtx, in.DrmSystems)
	if oneof := Encryption_Aes128Encryption_ToProto(mapCtx, in.Aes128); oneof != nil {
		out.EncryptionMode = &pb.Encryption_Aes128{Aes128: oneof}
	}
	if oneof := Encryption_SampleAesEncryption_ToProto(mapCtx, in.SampleAes); oneof != nil {
		out.EncryptionMode = &pb.Encryption_SampleAes{SampleAes: oneof}
	}
	if oneof := Encryption_MpegCommonEncryption_ToProto(mapCtx, in.MpegCenc); oneof != nil {
		out.EncryptionMode = &pb.Encryption_MpegCenc{MpegCenc: oneof}
	}
	return out
}
func Encryption_Aes128Encryption_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_Aes128Encryption) *krm.Encryption_Aes128Encryption {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_Aes128Encryption{}
	return out
}
func Encryption_Aes128Encryption_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_Aes128Encryption) *pb.Encryption_Aes128Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_Aes128Encryption{}
	return out
}
func Encryption_Clearkey_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_Clearkey) *krm.Encryption_Clearkey {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_Clearkey{}
	return out
}
func Encryption_Clearkey_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_Clearkey) *pb.Encryption_Clearkey {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_Clearkey{}
	return out
}
func Encryption_DrmSystems_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_DrmSystems) *krm.Encryption_DrmSystems {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_DrmSystems{}
	out.Widevine = Encryption_Widevine_FromProto(mapCtx, in.GetWidevine())
	out.Fairplay = Encryption_Fairplay_FromProto(mapCtx, in.GetFairplay())
	out.Playready = Encryption_Playready_FromProto(mapCtx, in.GetPlayready())
	out.Clearkey = Encryption_Clearkey_FromProto(mapCtx, in.GetClearkey())
	return out
}
func Encryption_DrmSystems_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_DrmSystems) *pb.Encryption_DrmSystems {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_DrmSystems{}
	out.Widevine = Encryption_Widevine_ToProto(mapCtx, in.Widevine)
	out.Fairplay = Encryption_Fairplay_ToProto(mapCtx, in.Fairplay)
	out.Playready = Encryption_Playready_ToProto(mapCtx, in.Playready)
	out.Clearkey = Encryption_Clearkey_ToProto(mapCtx, in.Clearkey)
	return out
}
func Encryption_Fairplay_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_Fairplay) *krm.Encryption_Fairplay {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_Fairplay{}
	return out
}
func Encryption_Fairplay_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_Fairplay) *pb.Encryption_Fairplay {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_Fairplay{}
	return out
}
func Encryption_MpegCommonEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_MpegCommonEncryption) *krm.Encryption_MpegCommonEncryption {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_MpegCommonEncryption{}
	out.Scheme = direct.LazyPtr(in.GetScheme())
	return out
}
func Encryption_MpegCommonEncryption_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_MpegCommonEncryption) *pb.Encryption_MpegCommonEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_MpegCommonEncryption{}
	out.Scheme = direct.ValueOf(in.Scheme)
	return out
}
func Encryption_Playready_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_Playready) *krm.Encryption_Playready {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_Playready{}
	return out
}
func Encryption_Playready_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_Playready) *pb.Encryption_Playready {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_Playready{}
	return out
}
func Encryption_SampleAesEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_SampleAesEncryption) *krm.Encryption_SampleAesEncryption {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_SampleAesEncryption{}
	return out
}
func Encryption_SampleAesEncryption_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_SampleAesEncryption) *pb.Encryption_SampleAesEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_SampleAesEncryption{}
	return out
}
func Encryption_SecretManagerSource_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_SecretManagerSource) *krm.Encryption_SecretManagerSource {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_SecretManagerSource{}
	out.SecretVersion = direct.LazyPtr(in.GetSecretVersion())
	return out
}
func Encryption_SecretManagerSource_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_SecretManagerSource) *pb.Encryption_SecretManagerSource {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_SecretManagerSource{}
	out.SecretVersion = direct.ValueOf(in.SecretVersion)
	return out
}
func Encryption_Widevine_FromProto(mapCtx *direct.MapContext, in *pb.Encryption_Widevine) *krm.Encryption_Widevine {
	if in == nil {
		return nil
	}
	out := &krm.Encryption_Widevine{}
	return out
}
func Encryption_Widevine_ToProto(mapCtx *direct.MapContext, in *krm.Encryption_Widevine) *pb.Encryption_Widevine {
	if in == nil {
		return nil
	}
	out := &pb.Encryption_Widevine{}
	return out
}
func InputAttachment_FromProto(mapCtx *direct.MapContext, in *pb.InputAttachment) *krm.InputAttachment {
	if in == nil {
		return nil
	}
	out := &krm.InputAttachment{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Input = direct.LazyPtr(in.GetInput())
	out.AutomaticFailover = InputAttachment_AutomaticFailover_FromProto(mapCtx, in.GetAutomaticFailover())
	return out
}
func InputAttachment_ToProto(mapCtx *direct.MapContext, in *krm.InputAttachment) *pb.InputAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InputAttachment{}
	out.Key = direct.ValueOf(in.Key)
	out.Input = direct.ValueOf(in.Input)
	out.AutomaticFailover = InputAttachment_AutomaticFailover_ToProto(mapCtx, in.AutomaticFailover)
	return out
}
func InputAttachment_AutomaticFailover_FromProto(mapCtx *direct.MapContext, in *pb.InputAttachment_AutomaticFailover) *krm.InputAttachment_AutomaticFailover {
	if in == nil {
		return nil
	}
	out := &krm.InputAttachment_AutomaticFailover{}
	out.InputKeys = in.InputKeys
	return out
}
func InputAttachment_AutomaticFailover_ToProto(mapCtx *direct.MapContext, in *krm.InputAttachment_AutomaticFailover) *pb.InputAttachment_AutomaticFailover {
	if in == nil {
		return nil
	}
	out := &pb.InputAttachment_AutomaticFailover{}
	out.InputKeys = in.InputKeys
	return out
}
func InputConfig_FromProto(mapCtx *direct.MapContext, in *pb.InputConfig) *krm.InputConfig {
	if in == nil {
		return nil
	}
	out := &krm.InputConfig{}
	out.InputSwitchMode = direct.Enum_FromProto(mapCtx, in.GetInputSwitchMode())
	return out
}
func InputConfig_ToProto(mapCtx *direct.MapContext, in *krm.InputConfig) *pb.InputConfig {
	if in == nil {
		return nil
	}
	out := &pb.InputConfig{}
	out.InputSwitchMode = direct.Enum_ToProto[pb.InputConfig_InputSwitchMode](mapCtx, in.InputSwitchMode)
	return out
}
func LogConfig_FromProto(mapCtx *direct.MapContext, in *pb.LogConfig) *krm.LogConfig {
	if in == nil {
		return nil
	}
	out := &krm.LogConfig{}
	out.LogSeverity = direct.Enum_FromProto(mapCtx, in.GetLogSeverity())
	return out
}
func LogConfig_ToProto(mapCtx *direct.MapContext, in *krm.LogConfig) *pb.LogConfig {
	if in == nil {
		return nil
	}
	out := &pb.LogConfig{}
	out.LogSeverity = direct.Enum_ToProto[pb.LogConfig_LogSeverity](mapCtx, in.LogSeverity)
	return out
}
func Manifest_FromProto(mapCtx *direct.MapContext, in *pb.Manifest) *krm.Manifest {
	if in == nil {
		return nil
	}
	out := &krm.Manifest{}
	out.FileName = direct.LazyPtr(in.GetFileName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.MuxStreams = in.MuxStreams
	out.MaxSegmentCount = direct.LazyPtr(in.GetMaxSegmentCount())
	out.SegmentKeepDuration = direct.StringDuration_FromProto(mapCtx, in.GetSegmentKeepDuration())
	out.UseTimecodeAsTimeline = direct.LazyPtr(in.GetUseTimecodeAsTimeline())
	out.Key = direct.LazyPtr(in.GetKey())
	return out
}
func Manifest_ToProto(mapCtx *direct.MapContext, in *krm.Manifest) *pb.Manifest {
	if in == nil {
		return nil
	}
	out := &pb.Manifest{}
	out.FileName = direct.ValueOf(in.FileName)
	out.Type = direct.Enum_ToProto[pb.Manifest_ManifestType](mapCtx, in.Type)
	out.MuxStreams = in.MuxStreams
	out.MaxSegmentCount = direct.ValueOf(in.MaxSegmentCount)
	out.SegmentKeepDuration = direct.StringDuration_ToProto(mapCtx, in.SegmentKeepDuration)
	out.UseTimecodeAsTimeline = direct.ValueOf(in.UseTimecodeAsTimeline)
	out.Key = direct.ValueOf(in.Key)
	return out
}
func MuxStream_FromProto(mapCtx *direct.MapContext, in *pb.MuxStream) *krm.MuxStream {
	if in == nil {
		return nil
	}
	out := &krm.MuxStream{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Container = direct.LazyPtr(in.GetContainer())
	out.ElementaryStreams = in.ElementaryStreams
	out.SegmentSettings = SegmentSettings_FromProto(mapCtx, in.GetSegmentSettings())
	out.EncryptionID = direct.LazyPtr(in.GetEncryptionId())
	return out
}
func MuxStream_ToProto(mapCtx *direct.MapContext, in *krm.MuxStream) *pb.MuxStream {
	if in == nil {
		return nil
	}
	out := &pb.MuxStream{}
	out.Key = direct.ValueOf(in.Key)
	out.Container = direct.ValueOf(in.Container)
	out.ElementaryStreams = in.ElementaryStreams
	out.SegmentSettings = SegmentSettings_ToProto(mapCtx, in.SegmentSettings)
	out.EncryptionId = direct.ValueOf(in.EncryptionID)
	return out
}
func NormalizedCoordinate_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedCoordinate) *krm.NormalizedCoordinate {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedCoordinate{}
	out.X = direct.LazyPtr(in.GetX())
	out.Y = direct.LazyPtr(in.GetY())
	return out
}
func NormalizedCoordinate_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedCoordinate) *pb.NormalizedCoordinate {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedCoordinate{}
	out.X = direct.ValueOf(in.X)
	out.Y = direct.ValueOf(in.Y)
	return out
}
func NormalizedResolution_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedResolution) *krm.NormalizedResolution {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedResolution{}
	out.W = direct.LazyPtr(in.GetW())
	out.H = direct.LazyPtr(in.GetH())
	return out
}
func NormalizedResolution_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedResolution) *pb.NormalizedResolution {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedResolution{}
	out.W = direct.ValueOf(in.W)
	out.H = direct.ValueOf(in.H)
	return out
}
func RetentionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RetentionConfig) *krm.RetentionConfig {
	if in == nil {
		return nil
	}
	out := &krm.RetentionConfig{}
	out.RetentionWindowDuration = direct.StringDuration_FromProto(mapCtx, in.GetRetentionWindowDuration())
	return out
}
func RetentionConfig_ToProto(mapCtx *direct.MapContext, in *krm.RetentionConfig) *pb.RetentionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RetentionConfig{}
	out.RetentionWindowDuration = direct.StringDuration_ToProto(mapCtx, in.RetentionWindowDuration)
	return out
}
func SegmentSettings_FromProto(mapCtx *direct.MapContext, in *pb.SegmentSettings) *krm.SegmentSettings {
	if in == nil {
		return nil
	}
	out := &krm.SegmentSettings{}
	out.SegmentDuration = direct.StringDuration_FromProto(mapCtx, in.GetSegmentDuration())
	return out
}
func SegmentSettings_ToProto(mapCtx *direct.MapContext, in *krm.SegmentSettings) *pb.SegmentSettings {
	if in == nil {
		return nil
	}
	out := &pb.SegmentSettings{}
	out.SegmentDuration = direct.StringDuration_ToProto(mapCtx, in.SegmentDuration)
	return out
}
func SpriteSheet_FromProto(mapCtx *direct.MapContext, in *pb.SpriteSheet) *krm.SpriteSheet {
	if in == nil {
		return nil
	}
	out := &krm.SpriteSheet{}
	out.Format = direct.LazyPtr(in.GetFormat())
	out.FilePrefix = direct.LazyPtr(in.GetFilePrefix())
	out.SpriteWidthPixels = direct.LazyPtr(in.GetSpriteWidthPixels())
	out.SpriteHeightPixels = direct.LazyPtr(in.GetSpriteHeightPixels())
	out.ColumnCount = direct.LazyPtr(in.GetColumnCount())
	out.RowCount = direct.LazyPtr(in.GetRowCount())
	out.Interval = direct.StringDuration_FromProto(mapCtx, in.GetInterval())
	out.Quality = direct.LazyPtr(in.GetQuality())
	return out
}
func SpriteSheet_ToProto(mapCtx *direct.MapContext, in *krm.SpriteSheet) *pb.SpriteSheet {
	if in == nil {
		return nil
	}
	out := &pb.SpriteSheet{}
	out.Format = direct.ValueOf(in.Format)
	out.FilePrefix = direct.ValueOf(in.FilePrefix)
	out.SpriteWidthPixels = direct.ValueOf(in.SpriteWidthPixels)
	out.SpriteHeightPixels = direct.ValueOf(in.SpriteHeightPixels)
	out.ColumnCount = direct.ValueOf(in.ColumnCount)
	out.RowCount = direct.ValueOf(in.RowCount)
	out.Interval = direct.StringDuration_ToProto(mapCtx, in.Interval)
	out.Quality = direct.ValueOf(in.Quality)
	return out
}
func StaticOverlay_FromProto(mapCtx *direct.MapContext, in *pb.StaticOverlay) *krm.StaticOverlay {
	if in == nil {
		return nil
	}
	out := &krm.StaticOverlay{}
	out.Asset = direct.LazyPtr(in.GetAsset())
	out.Resolution = NormalizedResolution_FromProto(mapCtx, in.GetResolution())
	out.Position = NormalizedCoordinate_FromProto(mapCtx, in.GetPosition())
	out.Opacity = direct.LazyPtr(in.GetOpacity())
	return out
}
func StaticOverlay_ToProto(mapCtx *direct.MapContext, in *krm.StaticOverlay) *pb.StaticOverlay {
	if in == nil {
		return nil
	}
	out := &pb.StaticOverlay{}
	out.Asset = direct.ValueOf(in.Asset)
	out.Resolution = NormalizedResolution_ToProto(mapCtx, in.Resolution)
	out.Position = NormalizedCoordinate_ToProto(mapCtx, in.Position)
	out.Opacity = direct.ValueOf(in.Opacity)
	return out
}
func TextStream_FromProto(mapCtx *direct.MapContext, in *pb.TextStream) *krm.TextStream {
	if in == nil {
		return nil
	}
	out := &krm.TextStream{}
	out.Codec = direct.LazyPtr(in.GetCodec())
	return out
}
func TextStream_ToProto(mapCtx *direct.MapContext, in *krm.TextStream) *pb.TextStream {
	if in == nil {
		return nil
	}
	out := &pb.TextStream{}
	out.Codec = direct.ValueOf(in.Codec)
	return out
}
func TimecodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.TimecodeConfig) *krm.TimecodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.TimecodeConfig{}
	out.Source = direct.Enum_FromProto(mapCtx, in.GetSource())
	out.UtcOffset = direct.StringDuration_FromProto(mapCtx, in.GetUtcOffset())
	out.TimeZone = TimeZone_FromProto(mapCtx, in.GetTimeZone())
	return out
}
func TimecodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.TimecodeConfig) *pb.TimecodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.TimecodeConfig{}
	out.Source = direct.Enum_ToProto[pb.TimecodeConfig_TimecodeSource](mapCtx, in.Source)
	if oneof := direct.StringDuration_ToProto(mapCtx, in.UtcOffset); oneof != nil {
		out.TimeOffset = &pb.TimecodeConfig_UtcOffset{UtcOffset: oneof}
	}
	if oneof := TimeZone_ToProto(mapCtx, in.TimeZone); oneof != nil {
		out.TimeOffset = &pb.TimecodeConfig_TimeZone{TimeZone: oneof}
	}
	return out
}
func VideoChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.VideoChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoChannelObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputAttachments
	// MISSING: ActiveInput
	// MISSING: Output
	// MISSING: ElementaryStreams
	// MISSING: MuxStreams
	// MISSING: Manifests
	// MISSING: SpriteSheets
	// MISSING: StreamingState
	// MISSING: StreamingError
	// MISSING: LogConfig
	// MISSING: TimecodeConfig
	// MISSING: Encryptions
	// MISSING: InputConfig
	// MISSING: RetentionConfig
	// MISSING: StaticOverlays
	return out
}
func VideoChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputAttachments
	// MISSING: ActiveInput
	// MISSING: Output
	// MISSING: ElementaryStreams
	// MISSING: MuxStreams
	// MISSING: Manifests
	// MISSING: SpriteSheets
	// MISSING: StreamingState
	// MISSING: StreamingError
	// MISSING: LogConfig
	// MISSING: TimecodeConfig
	// MISSING: Encryptions
	// MISSING: InputConfig
	// MISSING: RetentionConfig
	// MISSING: StaticOverlays
	return out
}
func VideoChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.VideoChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoChannelSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputAttachments
	// MISSING: ActiveInput
	// MISSING: Output
	// MISSING: ElementaryStreams
	// MISSING: MuxStreams
	// MISSING: Manifests
	// MISSING: SpriteSheets
	// MISSING: StreamingState
	// MISSING: StreamingError
	// MISSING: LogConfig
	// MISSING: TimecodeConfig
	// MISSING: Encryptions
	// MISSING: InputConfig
	// MISSING: RetentionConfig
	// MISSING: StaticOverlays
	return out
}
func VideoChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: InputAttachments
	// MISSING: ActiveInput
	// MISSING: Output
	// MISSING: ElementaryStreams
	// MISSING: MuxStreams
	// MISSING: Manifests
	// MISSING: SpriteSheets
	// MISSING: StreamingState
	// MISSING: StreamingError
	// MISSING: LogConfig
	// MISSING: TimecodeConfig
	// MISSING: Encryptions
	// MISSING: InputConfig
	// MISSING: RetentionConfig
	// MISSING: StaticOverlays
	return out
}
func VideoStream_FromProto(mapCtx *direct.MapContext, in *pb.VideoStream) *krm.VideoStream {
	if in == nil {
		return nil
	}
	out := &krm.VideoStream{}
	out.H264 = VideoStream_H264CodecSettings_FromProto(mapCtx, in.GetH264())
	return out
}
func VideoStream_ToProto(mapCtx *direct.MapContext, in *krm.VideoStream) *pb.VideoStream {
	if in == nil {
		return nil
	}
	out := &pb.VideoStream{}
	if oneof := VideoStream_H264CodecSettings_ToProto(mapCtx, in.H264); oneof != nil {
		out.CodecSettings = &pb.VideoStream_H264{H264: oneof}
	}
	return out
}
func VideoStream_H264CodecSettings_FromProto(mapCtx *direct.MapContext, in *pb.VideoStream_H264CodecSettings) *krm.VideoStream_H264CodecSettings {
	if in == nil {
		return nil
	}
	out := &krm.VideoStream_H264CodecSettings{}
	out.WidthPixels = direct.LazyPtr(in.GetWidthPixels())
	out.HeightPixels = direct.LazyPtr(in.GetHeightPixels())
	out.FrameRate = direct.LazyPtr(in.GetFrameRate())
	out.BitrateBps = direct.LazyPtr(in.GetBitrateBps())
	out.AllowOpenGop = direct.LazyPtr(in.GetAllowOpenGop())
	out.GopFrameCount = direct.LazyPtr(in.GetGopFrameCount())
	out.GopDuration = direct.StringDuration_FromProto(mapCtx, in.GetGopDuration())
	out.VbvSizeBits = direct.LazyPtr(in.GetVbvSizeBits())
	out.VbvFullnessBits = direct.LazyPtr(in.GetVbvFullnessBits())
	out.EntropyCoder = direct.LazyPtr(in.GetEntropyCoder())
	out.BPyramid = direct.LazyPtr(in.GetBPyramid())
	out.BFrameCount = direct.LazyPtr(in.GetBFrameCount())
	out.AqStrength = direct.LazyPtr(in.GetAqStrength())
	out.Profile = direct.LazyPtr(in.GetProfile())
	out.Tune = direct.LazyPtr(in.GetTune())
	return out
}
func VideoStream_H264CodecSettings_ToProto(mapCtx *direct.MapContext, in *krm.VideoStream_H264CodecSettings) *pb.VideoStream_H264CodecSettings {
	if in == nil {
		return nil
	}
	out := &pb.VideoStream_H264CodecSettings{}
	out.WidthPixels = direct.ValueOf(in.WidthPixels)
	out.HeightPixels = direct.ValueOf(in.HeightPixels)
	out.FrameRate = direct.ValueOf(in.FrameRate)
	out.BitrateBps = direct.ValueOf(in.BitrateBps)
	out.AllowOpenGop = direct.ValueOf(in.AllowOpenGop)
	if oneof := VideoStream_H264CodecSettings_GopFrameCount_ToProto(mapCtx, in.GopFrameCount); oneof != nil {
		out.GopMode = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.GopDuration); oneof != nil {
		out.GopMode = &pb.VideoStream_H264CodecSettings_GopDuration{GopDuration: oneof}
	}
	out.VbvSizeBits = direct.ValueOf(in.VbvSizeBits)
	out.VbvFullnessBits = direct.ValueOf(in.VbvFullnessBits)
	out.EntropyCoder = direct.ValueOf(in.EntropyCoder)
	out.BPyramid = direct.ValueOf(in.BPyramid)
	out.BFrameCount = direct.ValueOf(in.BFrameCount)
	out.AqStrength = direct.ValueOf(in.AqStrength)
	out.Profile = direct.ValueOf(in.Profile)
	out.Tune = direct.ValueOf(in.Tune)
	return out
}
