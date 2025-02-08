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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
)
func LiveSession_FromProto(mapCtx *direct.MapContext, in *pb.LiveSession) *krm.LiveSession {
	if in == nil {
		return nil
	}
	out := &krm.LiveSession{}
	// MISSING: Name
	// MISSING: PlayURI
	out.AdTagMacros = in.AdTagMacros
	out.ManifestOptions = ManifestOptions_FromProto(mapCtx, in.GetManifestOptions())
	out.GamSettings = LiveSession_GamSettings_FromProto(mapCtx, in.GetGamSettings())
	out.LiveConfig = direct.LazyPtr(in.GetLiveConfig())
	out.AdTracking = direct.Enum_FromProto(mapCtx, in.GetAdTracking())
	return out
}
func LiveSession_ToProto(mapCtx *direct.MapContext, in *krm.LiveSession) *pb.LiveSession {
	if in == nil {
		return nil
	}
	out := &pb.LiveSession{}
	// MISSING: Name
	// MISSING: PlayURI
	out.AdTagMacros = in.AdTagMacros
	out.ManifestOptions = ManifestOptions_ToProto(mapCtx, in.ManifestOptions)
	out.GamSettings = LiveSession_GamSettings_ToProto(mapCtx, in.GamSettings)
	out.LiveConfig = direct.ValueOf(in.LiveConfig)
	out.AdTracking = direct.Enum_ToProto[pb.AdTracking](mapCtx, in.AdTracking)
	return out
}
func LiveSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LiveSession) *krm.LiveSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LiveSessionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PlayURI = direct.LazyPtr(in.GetPlayUri())
	// MISSING: AdTagMacros
	// MISSING: ManifestOptions
	// MISSING: GamSettings
	// MISSING: LiveConfig
	// MISSING: AdTracking
	return out
}
func LiveSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LiveSessionObservedState) *pb.LiveSession {
	if in == nil {
		return nil
	}
	out := &pb.LiveSession{}
	out.Name = direct.ValueOf(in.Name)
	out.PlayUri = direct.ValueOf(in.PlayURI)
	// MISSING: AdTagMacros
	// MISSING: ManifestOptions
	// MISSING: GamSettings
	// MISSING: LiveConfig
	// MISSING: AdTracking
	return out
}
func LiveSession_GamSettings_FromProto(mapCtx *direct.MapContext, in *pb.LiveSession_GamSettings) *krm.LiveSession_GamSettings {
	if in == nil {
		return nil
	}
	out := &krm.LiveSession_GamSettings{}
	out.StreamID = direct.LazyPtr(in.GetStreamId())
	out.TargetingParameters = in.TargetingParameters
	return out
}
func LiveSession_GamSettings_ToProto(mapCtx *direct.MapContext, in *krm.LiveSession_GamSettings) *pb.LiveSession_GamSettings {
	if in == nil {
		return nil
	}
	out := &pb.LiveSession_GamSettings{}
	out.StreamId = direct.ValueOf(in.StreamID)
	out.TargetingParameters = in.TargetingParameters
	return out
}
func ManifestOptions_FromProto(mapCtx *direct.MapContext, in *pb.ManifestOptions) *krm.ManifestOptions {
	if in == nil {
		return nil
	}
	out := &krm.ManifestOptions{}
	out.IncludeRenditions = direct.Slice_FromProto(mapCtx, in.IncludeRenditions, RenditionFilter_FromProto)
	out.BitrateOrder = direct.Enum_FromProto(mapCtx, in.GetBitrateOrder())
	return out
}
func ManifestOptions_ToProto(mapCtx *direct.MapContext, in *krm.ManifestOptions) *pb.ManifestOptions {
	if in == nil {
		return nil
	}
	out := &pb.ManifestOptions{}
	out.IncludeRenditions = direct.Slice_ToProto(mapCtx, in.IncludeRenditions, RenditionFilter_ToProto)
	out.BitrateOrder = direct.Enum_ToProto[pb.ManifestOptions_OrderPolicy](mapCtx, in.BitrateOrder)
	return out
}
func RenditionFilter_FromProto(mapCtx *direct.MapContext, in *pb.RenditionFilter) *krm.RenditionFilter {
	if in == nil {
		return nil
	}
	out := &krm.RenditionFilter{}
	out.BitrateBps = direct.LazyPtr(in.GetBitrateBps())
	out.Codecs = direct.LazyPtr(in.GetCodecs())
	return out
}
func RenditionFilter_ToProto(mapCtx *direct.MapContext, in *krm.RenditionFilter) *pb.RenditionFilter {
	if in == nil {
		return nil
	}
	out := &pb.RenditionFilter{}
	out.BitrateBps = direct.ValueOf(in.BitrateBps)
	out.Codecs = direct.ValueOf(in.Codecs)
	return out
}
func VideoLiveSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LiveSession) *krm.VideoLiveSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoLiveSessionObservedState{}
	// MISSING: Name
	// MISSING: PlayURI
	// MISSING: AdTagMacros
	// MISSING: ManifestOptions
	// MISSING: GamSettings
	// MISSING: LiveConfig
	// MISSING: AdTracking
	return out
}
func VideoLiveSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoLiveSessionObservedState) *pb.LiveSession {
	if in == nil {
		return nil
	}
	out := &pb.LiveSession{}
	// MISSING: Name
	// MISSING: PlayURI
	// MISSING: AdTagMacros
	// MISSING: ManifestOptions
	// MISSING: GamSettings
	// MISSING: LiveConfig
	// MISSING: AdTracking
	return out
}
func VideoLiveSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.LiveSession) *krm.VideoLiveSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoLiveSessionSpec{}
	// MISSING: Name
	// MISSING: PlayURI
	// MISSING: AdTagMacros
	// MISSING: ManifestOptions
	// MISSING: GamSettings
	// MISSING: LiveConfig
	// MISSING: AdTracking
	return out
}
func VideoLiveSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoLiveSessionSpec) *pb.LiveSession {
	if in == nil {
		return nil
	}
	out := &pb.LiveSession{}
	// MISSING: Name
	// MISSING: PlayURI
	// MISSING: AdTagMacros
	// MISSING: ManifestOptions
	// MISSING: GamSettings
	// MISSING: LiveConfig
	// MISSING: AdTracking
	return out
}
