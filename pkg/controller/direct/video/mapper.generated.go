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
func Asset_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.Asset {
	if in == nil {
		return nil
	}
	out := &krm.Asset{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Video = Asset_VideoAsset_FromProto(mapCtx, in.GetVideo())
	out.Image = Asset_ImageAsset_FromProto(mapCtx, in.GetImage())
	out.Crc32c = direct.LazyPtr(in.GetCrc32c())
	// MISSING: State
	// MISSING: Error
	return out
}
func Asset_ToProto(mapCtx *direct.MapContext, in *krm.Asset) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	if oneof := Asset_VideoAsset_ToProto(mapCtx, in.Video); oneof != nil {
		out.Resource = &pb.Asset_Video{Video: oneof}
	}
	if oneof := Asset_ImageAsset_ToProto(mapCtx, in.Image); oneof != nil {
		out.Resource = &pb.Asset_Image{Image: oneof}
	}
	out.Crc32c = direct.ValueOf(in.Crc32c)
	// MISSING: State
	// MISSING: Error
	return out
}
func AssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.AssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Video
	// MISSING: Image
	// MISSING: Crc32c
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func AssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Video
	// MISSING: Image
	// MISSING: Crc32c
	out.State = direct.Enum_ToProto[pb.Asset_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func Asset_ImageAsset_FromProto(mapCtx *direct.MapContext, in *pb.Asset_ImageAsset) *krm.Asset_ImageAsset {
	if in == nil {
		return nil
	}
	out := &krm.Asset_ImageAsset{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Asset_ImageAsset_ToProto(mapCtx *direct.MapContext, in *krm.Asset_ImageAsset) *pb.Asset_ImageAsset {
	if in == nil {
		return nil
	}
	out := &pb.Asset_ImageAsset{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Asset_VideoAsset_FromProto(mapCtx *direct.MapContext, in *pb.Asset_VideoAsset) *krm.Asset_VideoAsset {
	if in == nil {
		return nil
	}
	out := &krm.Asset_VideoAsset{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Asset_VideoAsset_ToProto(mapCtx *direct.MapContext, in *krm.Asset_VideoAsset) *pb.Asset_VideoAsset {
	if in == nil {
		return nil
	}
	out := &pb.Asset_VideoAsset{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func VideoAssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.VideoAssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoAssetObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Video
	// MISSING: Image
	// MISSING: Crc32c
	// MISSING: State
	// MISSING: Error
	return out
}
func VideoAssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoAssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Video
	// MISSING: Image
	// MISSING: Crc32c
	// MISSING: State
	// MISSING: Error
	return out
}
func VideoAssetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.VideoAssetSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoAssetSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Video
	// MISSING: Image
	// MISSING: Crc32c
	// MISSING: State
	// MISSING: Error
	return out
}
func VideoAssetSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoAssetSpec) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Video
	// MISSING: Image
	// MISSING: Crc32c
	// MISSING: State
	// MISSING: Error
	return out
}
