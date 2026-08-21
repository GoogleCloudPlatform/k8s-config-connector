// Copyright 2026 Google LLC
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

package livestream

import (
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/livestream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func LiveStreamAssetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.LiveStreamAssetSpec {
	if in == nil {
		return nil
	}
	out := &krm.LiveStreamAssetSpec{}
	out.Labels = in.Labels
	out.Video = AssetVideoAsset_FromProto(mapCtx, in.GetVideo())
	out.Image = AssetImageAsset_FromProto(mapCtx, in.GetImage())
	out.Crc32c = direct.LazyPtr(in.GetCrc32C())
	return out
}

func LiveStreamAssetSpec_ToProto(mapCtx *direct.MapContext, in *krm.LiveStreamAssetSpec) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.Labels = in.Labels
	if oneof := AssetVideoAsset_ToProto(mapCtx, in.Video); oneof != nil {
		out.Resource = &pb.Asset_Video{Video: oneof}
	}
	if oneof := AssetImageAsset_ToProto(mapCtx, in.Image); oneof != nil {
		out.Resource = &pb.Asset_Image{Image: oneof}
	}
	out.Crc32C = direct.ValueOf(in.Crc32c)
	return out
}
