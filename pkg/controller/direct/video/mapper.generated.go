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
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AdBreak_FromProto(mapCtx *direct.MapContext, in *pb.AdBreak) *krm.AdBreak {
	if in == nil {
		return nil
	}
	out := &krm.AdBreak{}
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	return out
}
func AdBreak_ToProto(mapCtx *direct.MapContext, in *krm.AdBreak) *pb.AdBreak {
	if in == nil {
		return nil
	}
	out := &pb.AdBreak{}
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	return out
}
func AudioStream_FromProto(mapCtx *direct.MapContext, in *pb.AudioStream) *krm.AudioStream {
	if in == nil {
		return nil
	}
	out := &krm.AudioStream{}
	out.Codec = direct.LazyPtr(in.GetCodec())
	out.BitrateBps = direct.LazyPtr(in.GetBitrateBps())
	out.ChannelCount = direct.LazyPtr(in.GetChannelCount())
	out.ChannelLayout = in.ChannelLayout
	out.Mapping = direct.Slice_FromProto(mapCtx, in.Mapping, AudioStream_AudioMapping_FromProto)
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func AudioStream_ToProto(mapCtx *direct.MapContext, in *krm.AudioStream) *pb.AudioStream {
	if in == nil {
		return nil
	}
	out := &pb.AudioStream{}
	out.Codec = direct.ValueOf(in.Codec)
	out.BitrateBps = direct.ValueOf(in.BitrateBps)
	out.ChannelCount = direct.ValueOf(in.ChannelCount)
	out.ChannelLayout = in.ChannelLayout
	out.Mapping = direct.Slice_ToProto(mapCtx, in.Mapping, AudioStream_AudioMapping_ToProto)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func AudioStream_AudioMapping_FromProto(mapCtx *direct.MapContext, in *pb.AudioStream_AudioMapping) *krm.AudioStream_AudioMapping {
	if in == nil {
		return nil
	}
	out := &krm.AudioStream_AudioMapping{}
	out.AtomKey = direct.LazyPtr(in.GetAtomKey())
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
	out.AtomKey = direct.ValueOf(in.AtomKey)
	out.InputKey = direct.ValueOf(in.InputKey)
	out.InputTrack = direct.ValueOf(in.InputTrack)
	out.InputChannel = direct.ValueOf(in.InputChannel)
	out.OutputChannel = direct.ValueOf(in.OutputChannel)
	out.GainDb = direct.ValueOf(in.GainDb)
	return out
}
func EditAtom_FromProto(mapCtx *direct.MapContext, in *pb.EditAtom) *krm.EditAtom {
	if in == nil {
		return nil
	}
	out := &krm.EditAtom{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Inputs = in.Inputs
	out.EndTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndTimeOffset())
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	return out
}
func EditAtom_ToProto(mapCtx *direct.MapContext, in *krm.EditAtom) *pb.EditAtom {
	if in == nil {
		return nil
	}
	out := &pb.EditAtom{}
	out.Key = direct.ValueOf(in.Key)
	out.Inputs = in.Inputs
	out.EndTimeOffset = direct.StringDuration_ToProto(mapCtx, in.EndTimeOffset)
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
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
	out.Aes128 = Encryption_Aes128Encryption_FromProto(mapCtx, in.GetAes128())
	out.SampleAes = Encryption_SampleAesEncryption_FromProto(mapCtx, in.GetSampleAes())
	out.MpegCenc = Encryption_MpegCommonEncryption_FromProto(mapCtx, in.GetMpegCenc())
	out.SecretManagerKeySource = Encryption_SecretManagerSource_FromProto(mapCtx, in.GetSecretManagerKeySource())
	out.DrmSystems = Encryption_DrmSystems_FromProto(mapCtx, in.GetDrmSystems())
	return out
}
func Encryption_ToProto(mapCtx *direct.MapContext, in *krm.Encryption) *pb.Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Encryption{}
	out.Id = direct.ValueOf(in.ID)
	if oneof := Encryption_Aes128Encryption_ToProto(mapCtx, in.Aes128); oneof != nil {
		out.EncryptionMode = &pb.Encryption_Aes128{Aes128: oneof}
	}
	if oneof := Encryption_SampleAesEncryption_ToProto(mapCtx, in.SampleAes); oneof != nil {
		out.EncryptionMode = &pb.Encryption_SampleAes{SampleAes: oneof}
	}
	if oneof := Encryption_MpegCommonEncryption_ToProto(mapCtx, in.MpegCenc); oneof != nil {
		out.EncryptionMode = &pb.Encryption_MpegCenc{MpegCenc: oneof}
	}
	if oneof := Encryption_SecretManagerSource_ToProto(mapCtx, in.SecretManagerKeySource); oneof != nil {
		out.SecretSource = &pb.Encryption_SecretManagerKeySource{SecretManagerKeySource: oneof}
	}
	out.DrmSystems = Encryption_DrmSystems_ToProto(mapCtx, in.DrmSystems)
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
func Input_FromProto(mapCtx *direct.MapContext, in *pb.Input) *krm.Input {
	if in == nil {
		return nil
	}
	out := &krm.Input{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.URI = direct.LazyPtr(in.GetUri())
	out.PreprocessingConfig = PreprocessingConfig_FromProto(mapCtx, in.GetPreprocessingConfig())
	return out
}
func Input_ToProto(mapCtx *direct.MapContext, in *krm.Input) *pb.Input {
	if in == nil {
		return nil
	}
	out := &pb.Input{}
	out.Key = direct.ValueOf(in.Key)
	out.Uri = direct.ValueOf(in.URI)
	out.PreprocessingConfig = PreprocessingConfig_ToProto(mapCtx, in.PreprocessingConfig)
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	out.Name = direct.LazyPtr(in.GetName())
	out.InputURI = direct.LazyPtr(in.GetInputUri())
	out.OutputURI = direct.LazyPtr(in.GetOutputUri())
	out.TemplateID = direct.LazyPtr(in.GetTemplateId())
	out.Config = JobConfig_FromProto(mapCtx, in.GetConfig())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	out.TtlAfterCompletionDays = direct.LazyPtr(in.GetTtlAfterCompletionDays())
	out.Labels = in.Labels
	// MISSING: Error
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.BatchModePriority = direct.LazyPtr(in.GetBatchModePriority())
	out.Optimization = direct.Enum_FromProto(mapCtx, in.GetOptimization())
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.InputUri = direct.ValueOf(in.InputURI)
	out.OutputUri = direct.ValueOf(in.OutputURI)
	if oneof := Job_TemplateId_ToProto(mapCtx, in.TemplateID); oneof != nil {
		out.JobConfig = oneof
	}
	if oneof := JobConfig_ToProto(mapCtx, in.Config); oneof != nil {
		out.JobConfig = &pb.Job_Config{Config: oneof}
	}
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	out.TtlAfterCompletionDays = direct.ValueOf(in.TtlAfterCompletionDays)
	out.Labels = in.Labels
	// MISSING: Error
	out.Mode = direct.Enum_ToProto[pb.Job_ProcessingMode](mapCtx, in.Mode)
	out.BatchModePriority = direct.ValueOf(in.BatchModePriority)
	out.Optimization = direct.Enum_ToProto[pb.Job_OptimizationStrategy](mapCtx, in.Optimization)
	return out
}
func JobConfig_FromProto(mapCtx *direct.MapContext, in *pb.JobConfig) *krm.JobConfig {
	if in == nil {
		return nil
	}
	out := &krm.JobConfig{}
	out.Inputs = direct.Slice_FromProto(mapCtx, in.Inputs, Input_FromProto)
	out.EditList = direct.Slice_FromProto(mapCtx, in.EditList, EditAtom_FromProto)
	out.ElementaryStreams = direct.Slice_FromProto(mapCtx, in.ElementaryStreams, ElementaryStream_FromProto)
	out.MuxStreams = direct.Slice_FromProto(mapCtx, in.MuxStreams, MuxStream_FromProto)
	out.Manifests = direct.Slice_FromProto(mapCtx, in.Manifests, Manifest_FromProto)
	out.Output = Output_FromProto(mapCtx, in.GetOutput())
	out.AdBreaks = direct.Slice_FromProto(mapCtx, in.AdBreaks, AdBreak_FromProto)
	out.PubsubDestination = PubsubDestination_FromProto(mapCtx, in.GetPubsubDestination())
	out.SpriteSheets = direct.Slice_FromProto(mapCtx, in.SpriteSheets, SpriteSheet_FromProto)
	out.Overlays = direct.Slice_FromProto(mapCtx, in.Overlays, Overlay_FromProto)
	out.Encryptions = direct.Slice_FromProto(mapCtx, in.Encryptions, Encryption_FromProto)
	return out
}
func JobConfig_ToProto(mapCtx *direct.MapContext, in *krm.JobConfig) *pb.JobConfig {
	if in == nil {
		return nil
	}
	out := &pb.JobConfig{}
	out.Inputs = direct.Slice_ToProto(mapCtx, in.Inputs, Input_ToProto)
	out.EditList = direct.Slice_ToProto(mapCtx, in.EditList, EditAtom_ToProto)
	out.ElementaryStreams = direct.Slice_ToProto(mapCtx, in.ElementaryStreams, ElementaryStream_ToProto)
	out.MuxStreams = direct.Slice_ToProto(mapCtx, in.MuxStreams, MuxStream_ToProto)
	out.Manifests = direct.Slice_ToProto(mapCtx, in.Manifests, Manifest_ToProto)
	out.Output = Output_ToProto(mapCtx, in.Output)
	out.AdBreaks = direct.Slice_ToProto(mapCtx, in.AdBreaks, AdBreak_ToProto)
	out.PubsubDestination = PubsubDestination_ToProto(mapCtx, in.PubsubDestination)
	out.SpriteSheets = direct.Slice_ToProto(mapCtx, in.SpriteSheets, SpriteSheet_ToProto)
	out.Overlays = direct.Slice_ToProto(mapCtx, in.Overlays, Overlay_ToProto)
	out.Encryptions = direct.Slice_ToProto(mapCtx, in.Encryptions, Encryption_ToProto)
	return out
}
func JobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.JobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobObservedState{}
	// MISSING: Name
	// MISSING: InputURI
	// MISSING: OutputURI
	// MISSING: TemplateID
	// MISSING: Config
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: TtlAfterCompletionDays
	// MISSING: Labels
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Mode
	// MISSING: BatchModePriority
	// MISSING: Optimization
	return out
}
func JobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: InputURI
	// MISSING: OutputURI
	// MISSING: TemplateID
	// MISSING: Config
	out.State = direct.Enum_ToProto[pb.Job_ProcessingState](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: TtlAfterCompletionDays
	// MISSING: Labels
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Mode
	// MISSING: BatchModePriority
	// MISSING: Optimization
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
	out.Dash = Manifest_DashConfig_FromProto(mapCtx, in.GetDash())
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
	if oneof := Manifest_DashConfig_ToProto(mapCtx, in.Dash); oneof != nil {
		out.ManifestConfig = &pb.Manifest_Dash{Dash: oneof}
	}
	return out
}
func Manifest_DashConfig_FromProto(mapCtx *direct.MapContext, in *pb.Manifest_DashConfig) *krm.Manifest_DashConfig {
	if in == nil {
		return nil
	}
	out := &krm.Manifest_DashConfig{}
	out.SegmentReferenceScheme = direct.Enum_FromProto(mapCtx, in.GetSegmentReferenceScheme())
	return out
}
func Manifest_DashConfig_ToProto(mapCtx *direct.MapContext, in *krm.Manifest_DashConfig) *pb.Manifest_DashConfig {
	if in == nil {
		return nil
	}
	out := &pb.Manifest_DashConfig{}
	out.SegmentReferenceScheme = direct.Enum_ToProto[pb.Manifest_DashConfig_SegmentReferenceScheme](mapCtx, in.SegmentReferenceScheme)
	return out
}
func MuxStream_FromProto(mapCtx *direct.MapContext, in *pb.MuxStream) *krm.MuxStream {
	if in == nil {
		return nil
	}
	out := &krm.MuxStream{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.FileName = direct.LazyPtr(in.GetFileName())
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
	out.FileName = direct.ValueOf(in.FileName)
	out.Container = direct.ValueOf(in.Container)
	out.ElementaryStreams = in.ElementaryStreams
	out.SegmentSettings = SegmentSettings_ToProto(mapCtx, in.SegmentSettings)
	out.EncryptionId = direct.ValueOf(in.EncryptionID)
	return out
}
func Output_FromProto(mapCtx *direct.MapContext, in *pb.Output) *krm.Output {
	if in == nil {
		return nil
	}
	out := &krm.Output{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Output_ToProto(mapCtx *direct.MapContext, in *krm.Output) *pb.Output {
	if in == nil {
		return nil
	}
	out := &pb.Output{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Overlay_FromProto(mapCtx *direct.MapContext, in *pb.Overlay) *krm.Overlay {
	if in == nil {
		return nil
	}
	out := &krm.Overlay{}
	out.Image = Overlay_Image_FromProto(mapCtx, in.GetImage())
	out.Animations = direct.Slice_FromProto(mapCtx, in.Animations, Overlay_Animation_FromProto)
	return out
}
func Overlay_ToProto(mapCtx *direct.MapContext, in *krm.Overlay) *pb.Overlay {
	if in == nil {
		return nil
	}
	out := &pb.Overlay{}
	out.Image = Overlay_Image_ToProto(mapCtx, in.Image)
	out.Animations = direct.Slice_ToProto(mapCtx, in.Animations, Overlay_Animation_ToProto)
	return out
}
func Overlay_Animation_FromProto(mapCtx *direct.MapContext, in *pb.Overlay_Animation) *krm.Overlay_Animation {
	if in == nil {
		return nil
	}
	out := &krm.Overlay_Animation{}
	out.AnimationStatic = Overlay_AnimationStatic_FromProto(mapCtx, in.GetAnimationStatic())
	out.AnimationFade = Overlay_AnimationFade_FromProto(mapCtx, in.GetAnimationFade())
	out.AnimationEnd = Overlay_AnimationEnd_FromProto(mapCtx, in.GetAnimationEnd())
	return out
}
func Overlay_Animation_ToProto(mapCtx *direct.MapContext, in *krm.Overlay_Animation) *pb.Overlay_Animation {
	if in == nil {
		return nil
	}
	out := &pb.Overlay_Animation{}
	if oneof := Overlay_AnimationStatic_ToProto(mapCtx, in.AnimationStatic); oneof != nil {
		out.AnimationType = &pb.Overlay_Animation_AnimationStatic{AnimationStatic: oneof}
	}
	if oneof := Overlay_AnimationFade_ToProto(mapCtx, in.AnimationFade); oneof != nil {
		out.AnimationType = &pb.Overlay_Animation_AnimationFade{AnimationFade: oneof}
	}
	if oneof := Overlay_AnimationEnd_ToProto(mapCtx, in.AnimationEnd); oneof != nil {
		out.AnimationType = &pb.Overlay_Animation_AnimationEnd{AnimationEnd: oneof}
	}
	return out
}
func Overlay_AnimationEnd_FromProto(mapCtx *direct.MapContext, in *pb.Overlay_AnimationEnd) *krm.Overlay_AnimationEnd {
	if in == nil {
		return nil
	}
	out := &krm.Overlay_AnimationEnd{}
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	return out
}
func Overlay_AnimationEnd_ToProto(mapCtx *direct.MapContext, in *krm.Overlay_AnimationEnd) *pb.Overlay_AnimationEnd {
	if in == nil {
		return nil
	}
	out := &pb.Overlay_AnimationEnd{}
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	return out
}
func Overlay_AnimationFade_FromProto(mapCtx *direct.MapContext, in *pb.Overlay_AnimationFade) *krm.Overlay_AnimationFade {
	if in == nil {
		return nil
	}
	out := &krm.Overlay_AnimationFade{}
	out.FadeType = direct.Enum_FromProto(mapCtx, in.GetFadeType())
	out.Xy = Overlay_NormalizedCoordinate_FromProto(mapCtx, in.GetXy())
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	out.EndTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndTimeOffset())
	return out
}
func Overlay_AnimationFade_ToProto(mapCtx *direct.MapContext, in *krm.Overlay_AnimationFade) *pb.Overlay_AnimationFade {
	if in == nil {
		return nil
	}
	out := &pb.Overlay_AnimationFade{}
	out.FadeType = direct.Enum_ToProto[pb.Overlay_FadeType](mapCtx, in.FadeType)
	out.Xy = Overlay_NormalizedCoordinate_ToProto(mapCtx, in.Xy)
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	out.EndTimeOffset = direct.StringDuration_ToProto(mapCtx, in.EndTimeOffset)
	return out
}
func Overlay_AnimationStatic_FromProto(mapCtx *direct.MapContext, in *pb.Overlay_AnimationStatic) *krm.Overlay_AnimationStatic {
	if in == nil {
		return nil
	}
	out := &krm.Overlay_AnimationStatic{}
	out.Xy = Overlay_NormalizedCoordinate_FromProto(mapCtx, in.GetXy())
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	return out
}
func Overlay_AnimationStatic_ToProto(mapCtx *direct.MapContext, in *krm.Overlay_AnimationStatic) *pb.Overlay_AnimationStatic {
	if in == nil {
		return nil
	}
	out := &pb.Overlay_AnimationStatic{}
	out.Xy = Overlay_NormalizedCoordinate_ToProto(mapCtx, in.Xy)
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	return out
}
func Overlay_Image_FromProto(mapCtx *direct.MapContext, in *pb.Overlay_Image) *krm.Overlay_Image {
	if in == nil {
		return nil
	}
	out := &krm.Overlay_Image{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Resolution = Overlay_NormalizedCoordinate_FromProto(mapCtx, in.GetResolution())
	out.Alpha = direct.LazyPtr(in.GetAlpha())
	return out
}
func Overlay_Image_ToProto(mapCtx *direct.MapContext, in *krm.Overlay_Image) *pb.Overlay_Image {
	if in == nil {
		return nil
	}
	out := &pb.Overlay_Image{}
	out.Uri = direct.ValueOf(in.URI)
	out.Resolution = Overlay_NormalizedCoordinate_ToProto(mapCtx, in.Resolution)
	out.Alpha = direct.ValueOf(in.Alpha)
	return out
}
func Overlay_NormalizedCoordinate_FromProto(mapCtx *direct.MapContext, in *pb.Overlay_NormalizedCoordinate) *krm.Overlay_NormalizedCoordinate {
	if in == nil {
		return nil
	}
	out := &krm.Overlay_NormalizedCoordinate{}
	out.X = direct.LazyPtr(in.GetX())
	out.Y = direct.LazyPtr(in.GetY())
	return out
}
func Overlay_NormalizedCoordinate_ToProto(mapCtx *direct.MapContext, in *krm.Overlay_NormalizedCoordinate) *pb.Overlay_NormalizedCoordinate {
	if in == nil {
		return nil
	}
	out := &pb.Overlay_NormalizedCoordinate{}
	out.X = direct.ValueOf(in.X)
	out.Y = direct.ValueOf(in.Y)
	return out
}
func PreprocessingConfig_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig) *krm.PreprocessingConfig {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig{}
	out.Color = PreprocessingConfig_Color_FromProto(mapCtx, in.GetColor())
	out.Denoise = PreprocessingConfig_Denoise_FromProto(mapCtx, in.GetDenoise())
	out.Deblock = PreprocessingConfig_Deblock_FromProto(mapCtx, in.GetDeblock())
	out.Audio = PreprocessingConfig_Audio_FromProto(mapCtx, in.GetAudio())
	out.Crop = PreprocessingConfig_Crop_FromProto(mapCtx, in.GetCrop())
	out.Pad = PreprocessingConfig_Pad_FromProto(mapCtx, in.GetPad())
	out.Deinterlace = PreprocessingConfig_Deinterlace_FromProto(mapCtx, in.GetDeinterlace())
	return out
}
func PreprocessingConfig_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig) *pb.PreprocessingConfig {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig{}
	out.Color = PreprocessingConfig_Color_ToProto(mapCtx, in.Color)
	out.Denoise = PreprocessingConfig_Denoise_ToProto(mapCtx, in.Denoise)
	out.Deblock = PreprocessingConfig_Deblock_ToProto(mapCtx, in.Deblock)
	out.Audio = PreprocessingConfig_Audio_ToProto(mapCtx, in.Audio)
	out.Crop = PreprocessingConfig_Crop_ToProto(mapCtx, in.Crop)
	out.Pad = PreprocessingConfig_Pad_ToProto(mapCtx, in.Pad)
	out.Deinterlace = PreprocessingConfig_Deinterlace_ToProto(mapCtx, in.Deinterlace)
	return out
}
func PreprocessingConfig_Audio_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Audio) *krm.PreprocessingConfig_Audio {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Audio{}
	out.Lufs = direct.LazyPtr(in.GetLufs())
	out.HighBoost = direct.LazyPtr(in.GetHighBoost())
	out.LowBoost = direct.LazyPtr(in.GetLowBoost())
	return out
}
func PreprocessingConfig_Audio_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Audio) *pb.PreprocessingConfig_Audio {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Audio{}
	out.Lufs = direct.ValueOf(in.Lufs)
	out.HighBoost = direct.ValueOf(in.HighBoost)
	out.LowBoost = direct.ValueOf(in.LowBoost)
	return out
}
func PreprocessingConfig_Color_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Color) *krm.PreprocessingConfig_Color {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Color{}
	out.Saturation = direct.LazyPtr(in.GetSaturation())
	out.Contrast = direct.LazyPtr(in.GetContrast())
	out.Brightness = direct.LazyPtr(in.GetBrightness())
	return out
}
func PreprocessingConfig_Color_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Color) *pb.PreprocessingConfig_Color {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Color{}
	out.Saturation = direct.ValueOf(in.Saturation)
	out.Contrast = direct.ValueOf(in.Contrast)
	out.Brightness = direct.ValueOf(in.Brightness)
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
func PreprocessingConfig_Deblock_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Deblock) *krm.PreprocessingConfig_Deblock {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Deblock{}
	out.Strength = direct.LazyPtr(in.GetStrength())
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func PreprocessingConfig_Deblock_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Deblock) *pb.PreprocessingConfig_Deblock {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Deblock{}
	out.Strength = direct.ValueOf(in.Strength)
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func PreprocessingConfig_Deinterlace_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Deinterlace) *krm.PreprocessingConfig_Deinterlace {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Deinterlace{}
	out.Yadif = PreprocessingConfig_Deinterlace_YadifConfig_FromProto(mapCtx, in.GetYadif())
	out.Bwdif = PreprocessingConfig_Deinterlace_BwdifConfig_FromProto(mapCtx, in.GetBwdif())
	return out
}
func PreprocessingConfig_Deinterlace_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Deinterlace) *pb.PreprocessingConfig_Deinterlace {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Deinterlace{}
	if oneof := PreprocessingConfig_Deinterlace_YadifConfig_ToProto(mapCtx, in.Yadif); oneof != nil {
		out.DeinterlacingFilter = &pb.PreprocessingConfig_Deinterlace_Yadif{Yadif: oneof}
	}
	if oneof := PreprocessingConfig_Deinterlace_BwdifConfig_ToProto(mapCtx, in.Bwdif); oneof != nil {
		out.DeinterlacingFilter = &pb.PreprocessingConfig_Deinterlace_Bwdif{Bwdif: oneof}
	}
	return out
}
func PreprocessingConfig_Deinterlace_BwdifConfig_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Deinterlace_BwdifConfig) *krm.PreprocessingConfig_Deinterlace_BwdifConfig {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Deinterlace_BwdifConfig{}
	out.Mode = direct.LazyPtr(in.GetMode())
	out.Parity = direct.LazyPtr(in.GetParity())
	out.DeinterlaceAllFrames = direct.LazyPtr(in.GetDeinterlaceAllFrames())
	return out
}
func PreprocessingConfig_Deinterlace_BwdifConfig_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Deinterlace_BwdifConfig) *pb.PreprocessingConfig_Deinterlace_BwdifConfig {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Deinterlace_BwdifConfig{}
	out.Mode = direct.ValueOf(in.Mode)
	out.Parity = direct.ValueOf(in.Parity)
	out.DeinterlaceAllFrames = direct.ValueOf(in.DeinterlaceAllFrames)
	return out
}
func PreprocessingConfig_Deinterlace_YadifConfig_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Deinterlace_YadifConfig) *krm.PreprocessingConfig_Deinterlace_YadifConfig {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Deinterlace_YadifConfig{}
	out.Mode = direct.LazyPtr(in.GetMode())
	out.DisableSpatialInterlacing = direct.LazyPtr(in.GetDisableSpatialInterlacing())
	out.Parity = direct.LazyPtr(in.GetParity())
	out.DeinterlaceAllFrames = direct.LazyPtr(in.GetDeinterlaceAllFrames())
	return out
}
func PreprocessingConfig_Deinterlace_YadifConfig_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Deinterlace_YadifConfig) *pb.PreprocessingConfig_Deinterlace_YadifConfig {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Deinterlace_YadifConfig{}
	out.Mode = direct.ValueOf(in.Mode)
	out.DisableSpatialInterlacing = direct.ValueOf(in.DisableSpatialInterlacing)
	out.Parity = direct.ValueOf(in.Parity)
	out.DeinterlaceAllFrames = direct.ValueOf(in.DeinterlaceAllFrames)
	return out
}
func PreprocessingConfig_Denoise_FromProto(mapCtx *direct.MapContext, in *pb.PreprocessingConfig_Denoise) *krm.PreprocessingConfig_Denoise {
	if in == nil {
		return nil
	}
	out := &krm.PreprocessingConfig_Denoise{}
	out.Strength = direct.LazyPtr(in.GetStrength())
	out.Tune = direct.LazyPtr(in.GetTune())
	return out
}
func PreprocessingConfig_Denoise_ToProto(mapCtx *direct.MapContext, in *krm.PreprocessingConfig_Denoise) *pb.PreprocessingConfig_Denoise {
	if in == nil {
		return nil
	}
	out := &pb.PreprocessingConfig_Denoise{}
	out.Strength = direct.ValueOf(in.Strength)
	out.Tune = direct.ValueOf(in.Tune)
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
func PubsubDestination_FromProto(mapCtx *direct.MapContext, in *pb.PubsubDestination) *krm.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &krm.PubsubDestination{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	return out
}
func PubsubDestination_ToProto(mapCtx *direct.MapContext, in *krm.PubsubDestination) *pb.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &pb.PubsubDestination{}
	out.Topic = direct.ValueOf(in.Topic)
	return out
}
func SegmentSettings_FromProto(mapCtx *direct.MapContext, in *pb.SegmentSettings) *krm.SegmentSettings {
	if in == nil {
		return nil
	}
	out := &krm.SegmentSettings{}
	out.SegmentDuration = direct.StringDuration_FromProto(mapCtx, in.GetSegmentDuration())
	out.IndividualSegments = direct.LazyPtr(in.GetIndividualSegments())
	return out
}
func SegmentSettings_ToProto(mapCtx *direct.MapContext, in *krm.SegmentSettings) *pb.SegmentSettings {
	if in == nil {
		return nil
	}
	out := &pb.SegmentSettings{}
	out.SegmentDuration = direct.StringDuration_ToProto(mapCtx, in.SegmentDuration)
	out.IndividualSegments = direct.ValueOf(in.IndividualSegments)
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
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	out.EndTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndTimeOffset())
	out.TotalCount = direct.LazyPtr(in.GetTotalCount())
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
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	out.EndTimeOffset = direct.StringDuration_ToProto(mapCtx, in.EndTimeOffset)
	if oneof := SpriteSheet_TotalCount_ToProto(mapCtx, in.TotalCount); oneof != nil {
		out.ExtractionStrategy = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.Interval); oneof != nil {
		out.ExtractionStrategy = &pb.SpriteSheet_Interval{Interval: oneof}
	}
	out.Quality = direct.ValueOf(in.Quality)
	return out
}
func TextStream_FromProto(mapCtx *direct.MapContext, in *pb.TextStream) *krm.TextStream {
	if in == nil {
		return nil
	}
	out := &krm.TextStream{}
	out.Codec = direct.LazyPtr(in.GetCodec())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.Mapping = direct.Slice_FromProto(mapCtx, in.Mapping, TextStream_TextMapping_FromProto)
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func TextStream_ToProto(mapCtx *direct.MapContext, in *krm.TextStream) *pb.TextStream {
	if in == nil {
		return nil
	}
	out := &pb.TextStream{}
	out.Codec = direct.ValueOf(in.Codec)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.Mapping = direct.Slice_ToProto(mapCtx, in.Mapping, TextStream_TextMapping_ToProto)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func TextStream_TextMapping_FromProto(mapCtx *direct.MapContext, in *pb.TextStream_TextMapping) *krm.TextStream_TextMapping {
	if in == nil {
		return nil
	}
	out := &krm.TextStream_TextMapping{}
	out.AtomKey = direct.LazyPtr(in.GetAtomKey())
	out.InputKey = direct.LazyPtr(in.GetInputKey())
	out.InputTrack = direct.LazyPtr(in.GetInputTrack())
	return out
}
func TextStream_TextMapping_ToProto(mapCtx *direct.MapContext, in *krm.TextStream_TextMapping) *pb.TextStream_TextMapping {
	if in == nil {
		return nil
	}
	out := &pb.TextStream_TextMapping{}
	out.AtomKey = direct.ValueOf(in.AtomKey)
	out.InputKey = direct.ValueOf(in.InputKey)
	out.InputTrack = direct.ValueOf(in.InputTrack)
	return out
}
func VideoJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.VideoJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoJobObservedState{}
	// MISSING: Name
	// MISSING: InputURI
	// MISSING: OutputURI
	// MISSING: TemplateID
	// MISSING: Config
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TtlAfterCompletionDays
	// MISSING: Labels
	// MISSING: Error
	// MISSING: Mode
	// MISSING: BatchModePriority
	// MISSING: Optimization
	return out
}
func VideoJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: InputURI
	// MISSING: OutputURI
	// MISSING: TemplateID
	// MISSING: Config
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TtlAfterCompletionDays
	// MISSING: Labels
	// MISSING: Error
	// MISSING: Mode
	// MISSING: BatchModePriority
	// MISSING: Optimization
	return out
}
func VideoJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.VideoJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoJobSpec{}
	// MISSING: Name
	// MISSING: InputURI
	// MISSING: OutputURI
	// MISSING: TemplateID
	// MISSING: Config
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TtlAfterCompletionDays
	// MISSING: Labels
	// MISSING: Error
	// MISSING: Mode
	// MISSING: BatchModePriority
	// MISSING: Optimization
	return out
}
func VideoJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: InputURI
	// MISSING: OutputURI
	// MISSING: TemplateID
	// MISSING: Config
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: TtlAfterCompletionDays
	// MISSING: Labels
	// MISSING: Error
	// MISSING: Mode
	// MISSING: BatchModePriority
	// MISSING: Optimization
	return out
}
func VideoStream_FromProto(mapCtx *direct.MapContext, in *pb.VideoStream) *krm.VideoStream {
	if in == nil {
		return nil
	}
	out := &krm.VideoStream{}
	out.H264 = VideoStream_H264CodecSettings_FromProto(mapCtx, in.GetH264())
	out.H265 = VideoStream_H265CodecSettings_FromProto(mapCtx, in.GetH265())
	out.Vp9 = VideoStream_Vp9CodecSettings_FromProto(mapCtx, in.GetVp9())
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
	if oneof := VideoStream_H265CodecSettings_ToProto(mapCtx, in.H265); oneof != nil {
		out.CodecSettings = &pb.VideoStream_H265{H265: oneof}
	}
	if oneof := VideoStream_Vp9CodecSettings_ToProto(mapCtx, in.Vp9); oneof != nil {
		out.CodecSettings = &pb.VideoStream_Vp9{Vp9: oneof}
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
	out.PixelFormat = direct.LazyPtr(in.GetPixelFormat())
	out.RateControlMode = direct.LazyPtr(in.GetRateControlMode())
	out.CrfLevel = direct.LazyPtr(in.GetCrfLevel())
	out.AllowOpenGop = direct.LazyPtr(in.GetAllowOpenGop())
	out.GopFrameCount = direct.LazyPtr(in.GetGopFrameCount())
	out.GopDuration = direct.StringDuration_FromProto(mapCtx, in.GetGopDuration())
	out.EnableTwoPass = direct.LazyPtr(in.GetEnableTwoPass())
	out.VbvSizeBits = direct.LazyPtr(in.GetVbvSizeBits())
	out.VbvFullnessBits = direct.LazyPtr(in.GetVbvFullnessBits())
	out.EntropyCoder = direct.LazyPtr(in.GetEntropyCoder())
	out.BPyramid = direct.LazyPtr(in.GetBPyramid())
	out.BFrameCount = direct.LazyPtr(in.GetBFrameCount())
	out.AqStrength = direct.LazyPtr(in.GetAqStrength())
	out.Profile = direct.LazyPtr(in.GetProfile())
	out.Tune = direct.LazyPtr(in.GetTune())
	out.Preset = direct.LazyPtr(in.GetPreset())
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
	out.PixelFormat = direct.ValueOf(in.PixelFormat)
	out.RateControlMode = direct.ValueOf(in.RateControlMode)
	out.CrfLevel = direct.ValueOf(in.CrfLevel)
	out.AllowOpenGop = direct.ValueOf(in.AllowOpenGop)
	if oneof := VideoStream_H264CodecSettings_GopFrameCount_ToProto(mapCtx, in.GopFrameCount); oneof != nil {
		out.GopMode = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.GopDuration); oneof != nil {
		out.GopMode = &pb.VideoStream_H264CodecSettings_GopDuration{GopDuration: oneof}
	}
	out.EnableTwoPass = direct.ValueOf(in.EnableTwoPass)
	out.VbvSizeBits = direct.ValueOf(in.VbvSizeBits)
	out.VbvFullnessBits = direct.ValueOf(in.VbvFullnessBits)
	out.EntropyCoder = direct.ValueOf(in.EntropyCoder)
	out.BPyramid = direct.ValueOf(in.BPyramid)
	out.BFrameCount = direct.ValueOf(in.BFrameCount)
	out.AqStrength = direct.ValueOf(in.AqStrength)
	out.Profile = direct.ValueOf(in.Profile)
	out.Tune = direct.ValueOf(in.Tune)
	out.Preset = direct.ValueOf(in.Preset)
	return out
}
func VideoStream_H265CodecSettings_FromProto(mapCtx *direct.MapContext, in *pb.VideoStream_H265CodecSettings) *krm.VideoStream_H265CodecSettings {
	if in == nil {
		return nil
	}
	out := &krm.VideoStream_H265CodecSettings{}
	out.WidthPixels = direct.LazyPtr(in.GetWidthPixels())
	out.HeightPixels = direct.LazyPtr(in.GetHeightPixels())
	out.FrameRate = direct.LazyPtr(in.GetFrameRate())
	out.BitrateBps = direct.LazyPtr(in.GetBitrateBps())
	out.PixelFormat = direct.LazyPtr(in.GetPixelFormat())
	out.RateControlMode = direct.LazyPtr(in.GetRateControlMode())
	out.CrfLevel = direct.LazyPtr(in.GetCrfLevel())
	out.AllowOpenGop = direct.LazyPtr(in.GetAllowOpenGop())
	out.GopFrameCount = direct.LazyPtr(in.GetGopFrameCount())
	out.GopDuration = direct.StringDuration_FromProto(mapCtx, in.GetGopDuration())
	out.EnableTwoPass = direct.LazyPtr(in.GetEnableTwoPass())
	out.VbvSizeBits = direct.LazyPtr(in.GetVbvSizeBits())
	out.VbvFullnessBits = direct.LazyPtr(in.GetVbvFullnessBits())
	out.BPyramid = direct.LazyPtr(in.GetBPyramid())
	out.BFrameCount = direct.LazyPtr(in.GetBFrameCount())
	out.AqStrength = direct.LazyPtr(in.GetAqStrength())
	out.Profile = direct.LazyPtr(in.GetProfile())
	out.Tune = direct.LazyPtr(in.GetTune())
	out.Preset = direct.LazyPtr(in.GetPreset())
	return out
}
func VideoStream_H265CodecSettings_ToProto(mapCtx *direct.MapContext, in *krm.VideoStream_H265CodecSettings) *pb.VideoStream_H265CodecSettings {
	if in == nil {
		return nil
	}
	out := &pb.VideoStream_H265CodecSettings{}
	out.WidthPixels = direct.ValueOf(in.WidthPixels)
	out.HeightPixels = direct.ValueOf(in.HeightPixels)
	out.FrameRate = direct.ValueOf(in.FrameRate)
	out.BitrateBps = direct.ValueOf(in.BitrateBps)
	out.PixelFormat = direct.ValueOf(in.PixelFormat)
	out.RateControlMode = direct.ValueOf(in.RateControlMode)
	out.CrfLevel = direct.ValueOf(in.CrfLevel)
	out.AllowOpenGop = direct.ValueOf(in.AllowOpenGop)
	if oneof := VideoStream_H265CodecSettings_GopFrameCount_ToProto(mapCtx, in.GopFrameCount); oneof != nil {
		out.GopMode = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.GopDuration); oneof != nil {
		out.GopMode = &pb.VideoStream_H265CodecSettings_GopDuration{GopDuration: oneof}
	}
	out.EnableTwoPass = direct.ValueOf(in.EnableTwoPass)
	out.VbvSizeBits = direct.ValueOf(in.VbvSizeBits)
	out.VbvFullnessBits = direct.ValueOf(in.VbvFullnessBits)
	out.BPyramid = direct.ValueOf(in.BPyramid)
	out.BFrameCount = direct.ValueOf(in.BFrameCount)
	out.AqStrength = direct.ValueOf(in.AqStrength)
	out.Profile = direct.ValueOf(in.Profile)
	out.Tune = direct.ValueOf(in.Tune)
	out.Preset = direct.ValueOf(in.Preset)
	return out
}
func VideoStream_Vp9CodecSettings_FromProto(mapCtx *direct.MapContext, in *pb.VideoStream_Vp9CodecSettings) *krm.VideoStream_Vp9CodecSettings {
	if in == nil {
		return nil
	}
	out := &krm.VideoStream_Vp9CodecSettings{}
	out.WidthPixels = direct.LazyPtr(in.GetWidthPixels())
	out.HeightPixels = direct.LazyPtr(in.GetHeightPixels())
	out.FrameRate = direct.LazyPtr(in.GetFrameRate())
	out.BitrateBps = direct.LazyPtr(in.GetBitrateBps())
	out.PixelFormat = direct.LazyPtr(in.GetPixelFormat())
	out.RateControlMode = direct.LazyPtr(in.GetRateControlMode())
	out.CrfLevel = direct.LazyPtr(in.GetCrfLevel())
	out.GopFrameCount = direct.LazyPtr(in.GetGopFrameCount())
	out.GopDuration = direct.StringDuration_FromProto(mapCtx, in.GetGopDuration())
	out.Profile = direct.LazyPtr(in.GetProfile())
	return out
}
func VideoStream_Vp9CodecSettings_ToProto(mapCtx *direct.MapContext, in *krm.VideoStream_Vp9CodecSettings) *pb.VideoStream_Vp9CodecSettings {
	if in == nil {
		return nil
	}
	out := &pb.VideoStream_Vp9CodecSettings{}
	out.WidthPixels = direct.ValueOf(in.WidthPixels)
	out.HeightPixels = direct.ValueOf(in.HeightPixels)
	out.FrameRate = direct.ValueOf(in.FrameRate)
	out.BitrateBps = direct.ValueOf(in.BitrateBps)
	out.PixelFormat = direct.ValueOf(in.PixelFormat)
	out.RateControlMode = direct.ValueOf(in.RateControlMode)
	out.CrfLevel = direct.ValueOf(in.CrfLevel)
	if oneof := VideoStream_Vp9CodecSettings_GopFrameCount_ToProto(mapCtx, in.GopFrameCount); oneof != nil {
		out.GopMode = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.GopDuration); oneof != nil {
		out.GopMode = &pb.VideoStream_Vp9CodecSettings_GopDuration{GopDuration: oneof}
	}
	out.Profile = direct.ValueOf(in.Profile)
	return out
}
