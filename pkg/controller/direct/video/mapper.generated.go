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
	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Companion_FromProto(mapCtx *direct.MapContext, in *pb.Companion) *krm.Companion {
	if in == nil {
		return nil
	}
	out := &krm.Companion{}
	out.IframeAdResource = IframeAdResource_FromProto(mapCtx, in.GetIframeAdResource())
	out.StaticAdResource = StaticAdResource_FromProto(mapCtx, in.GetStaticAdResource())
	out.HTMLAdResource = HtmlAdResource_FromProto(mapCtx, in.GetHtmlAdResource())
	out.ApiFramework = direct.LazyPtr(in.GetApiFramework())
	out.HeightPx = direct.LazyPtr(in.GetHeightPx())
	out.WidthPx = direct.LazyPtr(in.GetWidthPx())
	out.AssetHeightPx = direct.LazyPtr(in.GetAssetHeightPx())
	out.ExpandedHeightPx = direct.LazyPtr(in.GetExpandedHeightPx())
	out.AssetWidthPx = direct.LazyPtr(in.GetAssetWidthPx())
	out.ExpandedWidthPx = direct.LazyPtr(in.GetExpandedWidthPx())
	out.AdSlotID = direct.LazyPtr(in.GetAdSlotId())
	out.Events = direct.Slice_FromProto(mapCtx, in.Events, Event_FromProto)
	return out
}
func Companion_ToProto(mapCtx *direct.MapContext, in *krm.Companion) *pb.Companion {
	if in == nil {
		return nil
	}
	out := &pb.Companion{}
	if oneof := IframeAdResource_ToProto(mapCtx, in.IframeAdResource); oneof != nil {
		out.AdResource = &pb.Companion_IframeAdResource{IframeAdResource: oneof}
	}
	if oneof := StaticAdResource_ToProto(mapCtx, in.StaticAdResource); oneof != nil {
		out.AdResource = &pb.Companion_StaticAdResource{StaticAdResource: oneof}
	}
	if oneof := HtmlAdResource_ToProto(mapCtx, in.HTMLAdResource); oneof != nil {
		out.AdResource = &pb.Companion_HtmlAdResource{HtmlAdResource: oneof}
	}
	out.ApiFramework = direct.ValueOf(in.ApiFramework)
	out.HeightPx = direct.ValueOf(in.HeightPx)
	out.WidthPx = direct.ValueOf(in.WidthPx)
	out.AssetHeightPx = direct.ValueOf(in.AssetHeightPx)
	out.ExpandedHeightPx = direct.ValueOf(in.ExpandedHeightPx)
	out.AssetWidthPx = direct.ValueOf(in.AssetWidthPx)
	out.ExpandedWidthPx = direct.ValueOf(in.ExpandedWidthPx)
	out.AdSlotId = direct.ValueOf(in.AdSlotID)
	out.Events = direct.Slice_ToProto(mapCtx, in.Events, Event_ToProto)
	return out
}
func CompanionAds_FromProto(mapCtx *direct.MapContext, in *pb.CompanionAds) *krm.CompanionAds {
	if in == nil {
		return nil
	}
	out := &krm.CompanionAds{}
	out.DisplayRequirement = direct.Enum_FromProto(mapCtx, in.GetDisplayRequirement())
	out.Companions = direct.Slice_FromProto(mapCtx, in.Companions, Companion_FromProto)
	return out
}
func CompanionAds_ToProto(mapCtx *direct.MapContext, in *krm.CompanionAds) *pb.CompanionAds {
	if in == nil {
		return nil
	}
	out := &pb.CompanionAds{}
	out.DisplayRequirement = direct.Enum_ToProto[pb.CompanionAds_DisplayRequirement](mapCtx, in.DisplayRequirement)
	out.Companions = direct.Slice_ToProto(mapCtx, in.Companions, Companion_ToProto)
	return out
}
func Event_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.Event {
	if in == nil {
		return nil
	}
	out := &krm.Event{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.URI = direct.LazyPtr(in.GetUri())
	out.ID = direct.LazyPtr(in.GetId())
	out.Offset = direct.StringDuration_FromProto(mapCtx, in.GetOffset())
	return out
}
func Event_ToProto(mapCtx *direct.MapContext, in *krm.Event) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	out.Type = direct.Enum_ToProto[pb.Event_EventType](mapCtx, in.Type)
	out.Uri = direct.ValueOf(in.URI)
	out.Id = direct.ValueOf(in.ID)
	out.Offset = direct.StringDuration_ToProto(mapCtx, in.Offset)
	return out
}
func HtmlAdResource_FromProto(mapCtx *direct.MapContext, in *pb.HtmlAdResource) *krm.HtmlAdResource {
	if in == nil {
		return nil
	}
	out := &krm.HtmlAdResource{}
	out.HTMLSource = direct.LazyPtr(in.GetHtmlSource())
	return out
}
func HtmlAdResource_ToProto(mapCtx *direct.MapContext, in *krm.HtmlAdResource) *pb.HtmlAdResource {
	if in == nil {
		return nil
	}
	out := &pb.HtmlAdResource{}
	out.HtmlSource = direct.ValueOf(in.HTMLSource)
	return out
}
func IframeAdResource_FromProto(mapCtx *direct.MapContext, in *pb.IframeAdResource) *krm.IframeAdResource {
	if in == nil {
		return nil
	}
	out := &krm.IframeAdResource{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func IframeAdResource_ToProto(mapCtx *direct.MapContext, in *krm.IframeAdResource) *pb.IframeAdResource {
	if in == nil {
		return nil
	}
	out := &pb.IframeAdResource{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Interstitials_FromProto(mapCtx *direct.MapContext, in *pb.Interstitials) *krm.Interstitials {
	if in == nil {
		return nil
	}
	out := &krm.Interstitials{}
	out.AdBreaks = direct.Slice_FromProto(mapCtx, in.AdBreaks, VodSessionAdBreak_FromProto)
	out.SessionContent = VodSessionContent_FromProto(mapCtx, in.GetSessionContent())
	return out
}
func Interstitials_ToProto(mapCtx *direct.MapContext, in *krm.Interstitials) *pb.Interstitials {
	if in == nil {
		return nil
	}
	out := &pb.Interstitials{}
	out.AdBreaks = direct.Slice_ToProto(mapCtx, in.AdBreaks, VodSessionAdBreak_ToProto)
	out.SessionContent = VodSessionContent_ToProto(mapCtx, in.SessionContent)
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
func ProgressEvent_FromProto(mapCtx *direct.MapContext, in *pb.ProgressEvent) *krm.ProgressEvent {
	if in == nil {
		return nil
	}
	out := &krm.ProgressEvent{}
	out.TimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetTimeOffset())
	out.Events = direct.Slice_FromProto(mapCtx, in.Events, Event_FromProto)
	return out
}
func ProgressEvent_ToProto(mapCtx *direct.MapContext, in *krm.ProgressEvent) *pb.ProgressEvent {
	if in == nil {
		return nil
	}
	out := &pb.ProgressEvent{}
	out.TimeOffset = direct.StringDuration_ToProto(mapCtx, in.TimeOffset)
	out.Events = direct.Slice_ToProto(mapCtx, in.Events, Event_ToProto)
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
func StaticAdResource_FromProto(mapCtx *direct.MapContext, in *pb.StaticAdResource) *krm.StaticAdResource {
	if in == nil {
		return nil
	}
	out := &krm.StaticAdResource{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.CreativeType = direct.LazyPtr(in.GetCreativeType())
	return out
}
func StaticAdResource_ToProto(mapCtx *direct.MapContext, in *krm.StaticAdResource) *pb.StaticAdResource {
	if in == nil {
		return nil
	}
	out := &pb.StaticAdResource{}
	out.Uri = direct.ValueOf(in.URI)
	out.CreativeType = direct.ValueOf(in.CreativeType)
	return out
}
func VideoVodSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VodSession) *krm.VideoVodSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoVodSessionObservedState{}
	// MISSING: Name
	// MISSING: Interstitials
	// MISSING: PlayURI
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: AdTagMacroMap
	// MISSING: ManifestOptions
	// MISSING: AssetID
	// MISSING: AdTracking
	// MISSING: GamSettings
	// MISSING: VodConfig
	return out
}
func VideoVodSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoVodSessionObservedState) *pb.VodSession {
	if in == nil {
		return nil
	}
	out := &pb.VodSession{}
	// MISSING: Name
	// MISSING: Interstitials
	// MISSING: PlayURI
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: AdTagMacroMap
	// MISSING: ManifestOptions
	// MISSING: AssetID
	// MISSING: AdTracking
	// MISSING: GamSettings
	// MISSING: VodConfig
	return out
}
func VideoVodSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.VodSession) *krm.VideoVodSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoVodSessionSpec{}
	// MISSING: Name
	// MISSING: Interstitials
	// MISSING: PlayURI
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: AdTagMacroMap
	// MISSING: ManifestOptions
	// MISSING: AssetID
	// MISSING: AdTracking
	// MISSING: GamSettings
	// MISSING: VodConfig
	return out
}
func VideoVodSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoVodSessionSpec) *pb.VodSession {
	if in == nil {
		return nil
	}
	out := &pb.VodSession{}
	// MISSING: Name
	// MISSING: Interstitials
	// MISSING: PlayURI
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: AdTagMacroMap
	// MISSING: ManifestOptions
	// MISSING: AssetID
	// MISSING: AdTracking
	// MISSING: GamSettings
	// MISSING: VodConfig
	return out
}
func VodSession_FromProto(mapCtx *direct.MapContext, in *pb.VodSession) *krm.VodSession {
	if in == nil {
		return nil
	}
	out := &krm.VodSession{}
	// MISSING: Name
	// MISSING: Interstitials
	// MISSING: PlayURI
	out.SourceURI = direct.LazyPtr(in.GetSourceUri())
	out.AdTagURI = direct.LazyPtr(in.GetAdTagUri())
	out.AdTagMacroMap = in.AdTagMacroMap
	out.ManifestOptions = ManifestOptions_FromProto(mapCtx, in.GetManifestOptions())
	// MISSING: AssetID
	out.AdTracking = direct.Enum_FromProto(mapCtx, in.GetAdTracking())
	out.GamSettings = VodSession_GamSettings_FromProto(mapCtx, in.GetGamSettings())
	out.VodConfig = direct.LazyPtr(in.GetVodConfig())
	return out
}
func VodSession_ToProto(mapCtx *direct.MapContext, in *krm.VodSession) *pb.VodSession {
	if in == nil {
		return nil
	}
	out := &pb.VodSession{}
	// MISSING: Name
	// MISSING: Interstitials
	// MISSING: PlayURI
	out.SourceUri = direct.ValueOf(in.SourceURI)
	out.AdTagUri = direct.ValueOf(in.AdTagURI)
	out.AdTagMacroMap = in.AdTagMacroMap
	out.ManifestOptions = ManifestOptions_ToProto(mapCtx, in.ManifestOptions)
	// MISSING: AssetID
	out.AdTracking = direct.Enum_ToProto[pb.AdTracking](mapCtx, in.AdTracking)
	out.GamSettings = VodSession_GamSettings_ToProto(mapCtx, in.GamSettings)
	out.VodConfig = direct.ValueOf(in.VodConfig)
	return out
}
func VodSessionAd_FromProto(mapCtx *direct.MapContext, in *pb.VodSessionAd) *krm.VodSessionAd {
	if in == nil {
		return nil
	}
	out := &krm.VodSessionAd{}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.CompanionAds = CompanionAds_FromProto(mapCtx, in.GetCompanionAds())
	out.ActivityEvents = direct.Slice_FromProto(mapCtx, in.ActivityEvents, Event_FromProto)
	return out
}
func VodSessionAd_ToProto(mapCtx *direct.MapContext, in *krm.VodSessionAd) *pb.VodSessionAd {
	if in == nil {
		return nil
	}
	out := &pb.VodSessionAd{}
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.CompanionAds = CompanionAds_ToProto(mapCtx, in.CompanionAds)
	out.ActivityEvents = direct.Slice_ToProto(mapCtx, in.ActivityEvents, Event_ToProto)
	return out
}
func VodSessionAdBreak_FromProto(mapCtx *direct.MapContext, in *pb.VodSessionAdBreak) *krm.VodSessionAdBreak {
	if in == nil {
		return nil
	}
	out := &krm.VodSessionAdBreak{}
	out.ProgressEvents = direct.Slice_FromProto(mapCtx, in.ProgressEvents, ProgressEvent_FromProto)
	out.Ads = direct.Slice_FromProto(mapCtx, in.Ads, VodSessionAd_FromProto)
	out.EndTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndTimeOffset())
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	return out
}
func VodSessionAdBreak_ToProto(mapCtx *direct.MapContext, in *krm.VodSessionAdBreak) *pb.VodSessionAdBreak {
	if in == nil {
		return nil
	}
	out := &pb.VodSessionAdBreak{}
	out.ProgressEvents = direct.Slice_ToProto(mapCtx, in.ProgressEvents, ProgressEvent_ToProto)
	out.Ads = direct.Slice_ToProto(mapCtx, in.Ads, VodSessionAd_ToProto)
	out.EndTimeOffset = direct.StringDuration_ToProto(mapCtx, in.EndTimeOffset)
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	return out
}
func VodSessionContent_FromProto(mapCtx *direct.MapContext, in *pb.VodSessionContent) *krm.VodSessionContent {
	if in == nil {
		return nil
	}
	out := &krm.VodSessionContent{}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func VodSessionContent_ToProto(mapCtx *direct.MapContext, in *krm.VodSessionContent) *pb.VodSessionContent {
	if in == nil {
		return nil
	}
	out := &pb.VodSessionContent{}
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
func VodSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VodSession) *krm.VodSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VodSessionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Interstitials = Interstitials_FromProto(mapCtx, in.GetInterstitials())
	out.PlayURI = direct.LazyPtr(in.GetPlayUri())
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: AdTagMacroMap
	// MISSING: ManifestOptions
	out.AssetID = direct.LazyPtr(in.GetAssetId())
	// MISSING: AdTracking
	// MISSING: GamSettings
	// MISSING: VodConfig
	return out
}
func VodSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VodSessionObservedState) *pb.VodSession {
	if in == nil {
		return nil
	}
	out := &pb.VodSession{}
	out.Name = direct.ValueOf(in.Name)
	out.Interstitials = Interstitials_ToProto(mapCtx, in.Interstitials)
	out.PlayUri = direct.ValueOf(in.PlayURI)
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: AdTagMacroMap
	// MISSING: ManifestOptions
	out.AssetId = direct.ValueOf(in.AssetID)
	// MISSING: AdTracking
	// MISSING: GamSettings
	// MISSING: VodConfig
	return out
}
func VodSession_GamSettings_FromProto(mapCtx *direct.MapContext, in *pb.VodSession_GamSettings) *krm.VodSession_GamSettings {
	if in == nil {
		return nil
	}
	out := &krm.VodSession_GamSettings{}
	out.NetworkCode = direct.LazyPtr(in.GetNetworkCode())
	out.StreamID = direct.LazyPtr(in.GetStreamId())
	return out
}
func VodSession_GamSettings_ToProto(mapCtx *direct.MapContext, in *krm.VodSession_GamSettings) *pb.VodSession_GamSettings {
	if in == nil {
		return nil
	}
	out := &pb.VodSession_GamSettings{}
	out.NetworkCode = direct.ValueOf(in.NetworkCode)
	out.StreamId = direct.ValueOf(in.StreamID)
	return out
}
