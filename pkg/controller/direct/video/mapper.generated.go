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
func GamVodConfig_FromProto(mapCtx *direct.MapContext, in *pb.GamVodConfig) *krm.GamVodConfig {
	if in == nil {
		return nil
	}
	out := &krm.GamVodConfig{}
	out.NetworkCode = direct.LazyPtr(in.GetNetworkCode())
	return out
}
func GamVodConfig_ToProto(mapCtx *direct.MapContext, in *krm.GamVodConfig) *pb.GamVodConfig {
	if in == nil {
		return nil
	}
	out := &pb.GamVodConfig{}
	out.NetworkCode = direct.ValueOf(in.NetworkCode)
	return out
}
func VideoVodConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VodConfig) *krm.VideoVodConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoVodConfigObservedState{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamVodConfig
	// MISSING: State
	// MISSING: SourceFetchOptions
	return out
}
func VideoVodConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoVodConfigObservedState) *pb.VodConfig {
	if in == nil {
		return nil
	}
	out := &pb.VodConfig{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamVodConfig
	// MISSING: State
	// MISSING: SourceFetchOptions
	return out
}
func VideoVodConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.VodConfig) *krm.VideoVodConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoVodConfigSpec{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamVodConfig
	// MISSING: State
	// MISSING: SourceFetchOptions
	return out
}
func VideoVodConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoVodConfigSpec) *pb.VodConfig {
	if in == nil {
		return nil
	}
	out := &pb.VodConfig{}
	// MISSING: Name
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamVodConfig
	// MISSING: State
	// MISSING: SourceFetchOptions
	return out
}
func VodConfig_FromProto(mapCtx *direct.MapContext, in *pb.VodConfig) *krm.VodConfig {
	if in == nil {
		return nil
	}
	out := &krm.VodConfig{}
	// MISSING: Name
	out.SourceURI = direct.LazyPtr(in.GetSourceUri())
	out.AdTagURI = direct.LazyPtr(in.GetAdTagUri())
	out.GamVodConfig = GamVodConfig_FromProto(mapCtx, in.GetGamVodConfig())
	// MISSING: State
	out.SourceFetchOptions = FetchOptions_FromProto(mapCtx, in.GetSourceFetchOptions())
	return out
}
func VodConfig_ToProto(mapCtx *direct.MapContext, in *krm.VodConfig) *pb.VodConfig {
	if in == nil {
		return nil
	}
	out := &pb.VodConfig{}
	// MISSING: Name
	out.SourceUri = direct.ValueOf(in.SourceURI)
	out.AdTagUri = direct.ValueOf(in.AdTagURI)
	out.GamVodConfig = GamVodConfig_ToProto(mapCtx, in.GamVodConfig)
	// MISSING: State
	out.SourceFetchOptions = FetchOptions_ToProto(mapCtx, in.SourceFetchOptions)
	return out
}
func VodConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VodConfig) *krm.VodConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VodConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamVodConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SourceFetchOptions
	return out
}
func VodConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VodConfigObservedState) *pb.VodConfig {
	if in == nil {
		return nil
	}
	out := &pb.VodConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SourceURI
	// MISSING: AdTagURI
	// MISSING: GamVodConfig
	out.State = direct.Enum_ToProto[pb.VodConfig_State](mapCtx, in.State)
	// MISSING: SourceFetchOptions
	return out
}
