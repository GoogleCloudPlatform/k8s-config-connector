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
func FetchOptions_FromProto(mapCtx *direct.MapContext, in *pb.FetchOptions) *krm.FetchOptions {
	if in == nil {
		return nil
	}
	out := &krm.FetchOptions{}
	out.Headers = in.Headers
	return out
}
func FetchOptions_ToProto(mapCtx *direct.MapContext, in *krm.FetchOptions) *pb.FetchOptions {
	if in == nil {
		return nil
	}
	out := &pb.FetchOptions{}
	out.Headers = in.Headers
	return out
}
func GamLiveConfig_FromProto(mapCtx *direct.MapContext, in *pb.GamLiveConfig) *krm.GamLiveConfig {
	if in == nil {
		return nil
	}
	out := &krm.GamLiveConfig{}
	out.NetworkCode = direct.LazyPtr(in.GetNetworkCode())
	// MISSING: AssetKey
	// MISSING: CustomAssetKey
	return out
}
func GamLiveConfig_ToProto(mapCtx *direct.MapContext, in *krm.GamLiveConfig) *pb.GamLiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.GamLiveConfig{}
	out.NetworkCode = direct.ValueOf(in.NetworkCode)
	// MISSING: AssetKey
	// MISSING: CustomAssetKey
	return out
}
func GamLiveConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GamLiveConfig) *krm.GamLiveConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GamLiveConfigObservedState{}
	// MISSING: NetworkCode
	out.AssetKey = direct.LazyPtr(in.GetAssetKey())
	out.CustomAssetKey = direct.LazyPtr(in.GetCustomAssetKey())
	return out
}
func GamLiveConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GamLiveConfigObservedState) *pb.GamLiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.GamLiveConfig{}
	// MISSING: NetworkCode
	out.AssetKey = direct.ValueOf(in.AssetKey)
	out.CustomAssetKey = direct.ValueOf(in.CustomAssetKey)
	return out
}
func LiveConfig_FromProto(mapCtx *direct.MapContext, in *pb.LiveConfig) *krm.LiveConfig {
	if in == nil {
		return nil
	}
	out := &krm.LiveConfig{}
	// MISSING: Name
	out.SourceURI = direct.LazyPtr(in.GetSourceUri())
	out.AdTagURI = direct.LazyPtr(in.GetAdTagUri())
	out.GamLiveConfig = GamLiveConfig_FromProto(mapCtx, in.GetGamLiveConfig())
	// MISSING: State
	out.AdTracking = direct.Enum_FromProto(mapCtx, in.GetAdTracking())
	out.DefaultSlate = direct.LazyPtr(in.GetDefaultSlate())
	out.StitchingPolicy = direct.Enum_FromProto(mapCtx, in.GetStitchingPolicy())
	out.PrefetchConfig = PrefetchConfig_FromProto(mapCtx, in.GetPrefetchConfig())
	out.SourceFetchOptions = FetchOptions_FromProto(mapCtx, in.GetSourceFetchOptions())
	return out
}
func LiveConfig_ToProto(mapCtx *direct.MapContext, in *krm.LiveConfig) *pb.LiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.LiveConfig{}
	// MISSING: Name
	out.SourceUri = direct.ValueOf(in.SourceURI)
	out.AdTagUri = direct.ValueOf(in.AdTagURI)
	out.GamLiveConfig = GamLiveConfig_ToProto(mapCtx, in.GamLiveConfig)
	// MISSING: State
	out.AdTracking = direct.Enum_ToProto[pb.AdTracking](mapCtx, in.AdTracking)
	out.DefaultSlate = direct.ValueOf(in.DefaultSlate)
	out.StitchingPolicy = direct.Enum_ToProto[pb.LiveConfig_StitchingPolicy](mapCtx, in.StitchingPolicy)
	out.PrefetchConfig = PrefetchConfig_ToProto(mapCtx, in.PrefetchConfig)
	out.SourceFetchOptions = FetchOptions_ToProto(mapCtx, in.SourceFetchOptions)
	return out
}
func LiveConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LiveConfig) *krm.LiveConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LiveConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SourceURI
	// MISSING: AdTagURI
	out.GamLiveConfig = GamLiveConfigObservedState_FromProto(mapCtx, in.GetGamLiveConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: AdTracking
	// MISSING: DefaultSlate
	// MISSING: StitchingPolicy
	// MISSING: PrefetchConfig
	// MISSING: SourceFetchOptions
	return out
}
func LiveConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LiveConfigObservedState) *pb.LiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.LiveConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SourceURI
	// MISSING: AdTagURI
	out.GamLiveConfig = GamLiveConfigObservedState_ToProto(mapCtx, in.GamLiveConfig)
	out.State = direct.Enum_ToProto[pb.LiveConfig_State](mapCtx, in.State)
	// MISSING: AdTracking
	// MISSING: DefaultSlate
	// MISSING: StitchingPolicy
	// MISSING: PrefetchConfig
	// MISSING: SourceFetchOptions
	return out
}
func PrefetchConfig_FromProto(mapCtx *direct.MapContext, in *pb.PrefetchConfig) *krm.PrefetchConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrefetchConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.InitialAdRequestDuration = direct.StringDuration_FromProto(mapCtx, in.GetInitialAdRequestDuration())
	return out
}
func PrefetchConfig_ToProto(mapCtx *direct.MapContext, in *krm.PrefetchConfig) *pb.PrefetchConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrefetchConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.InitialAdRequestDuration = direct.StringDuration_ToProto(mapCtx, in.InitialAdRequestDuration)
	return out
}
func VideoLiveConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LiveConfig) *krm.VideoLiveConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoLiveConfigObservedState{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamLiveConfig
	// MISSING: State
	// MISSING: AdTracking
	// MISSING: DefaultSlate
	// MISSING: StitchingPolicy
	// MISSING: PrefetchConfig
	// MISSING: SourceFetchOptions
	return out
}
func VideoLiveConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoLiveConfigObservedState) *pb.LiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.LiveConfig{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamLiveConfig
	// MISSING: State
	// MISSING: AdTracking
	// MISSING: DefaultSlate
	// MISSING: StitchingPolicy
	// MISSING: PrefetchConfig
	// MISSING: SourceFetchOptions
	return out
}
func VideoLiveConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.LiveConfig) *krm.VideoLiveConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoLiveConfigSpec{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamLiveConfig
	// MISSING: State
	// MISSING: AdTracking
	// MISSING: DefaultSlate
	// MISSING: StitchingPolicy
	// MISSING: PrefetchConfig
	// MISSING: SourceFetchOptions
	return out
}
func VideoLiveConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoLiveConfigSpec) *pb.LiveConfig {
	if in == nil {
		return nil
	}
	out := &pb.LiveConfig{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamLiveConfig
	// MISSING: State
	// MISSING: AdTracking
	// MISSING: DefaultSlate
	// MISSING: StitchingPolicy
	// MISSING: PrefetchConfig
	// MISSING: SourceFetchOptions
	return out
}
