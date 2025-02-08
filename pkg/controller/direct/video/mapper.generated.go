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
func AdStitchDetail_FromProto(mapCtx *direct.MapContext, in *pb.AdStitchDetail) *krm.AdStitchDetail {
	if in == nil {
		return nil
	}
	out := &krm.AdStitchDetail{}
	out.AdBreakID = direct.LazyPtr(in.GetAdBreakId())
	out.AdID = direct.LazyPtr(in.GetAdId())
	out.AdTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetAdTimeOffset())
	out.SkipReason = direct.LazyPtr(in.GetSkipReason())
	// MISSING: Media
	return out
}
func AdStitchDetail_ToProto(mapCtx *direct.MapContext, in *krm.AdStitchDetail) *pb.AdStitchDetail {
	if in == nil {
		return nil
	}
	out := &pb.AdStitchDetail{}
	out.AdBreakId = direct.ValueOf(in.AdBreakID)
	out.AdId = direct.ValueOf(in.AdID)
	out.AdTimeOffset = direct.StringDuration_ToProto(mapCtx, in.AdTimeOffset)
	out.SkipReason = direct.ValueOf(in.SkipReason)
	// MISSING: Media
	return out
}
func VideoVodStitchDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VodStitchDetail) *krm.VideoVodStitchDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoVodStitchDetailObservedState{}
	// MISSING: Name
	// MISSING: AdStitchDetails
	return out
}
func VideoVodStitchDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoVodStitchDetailObservedState) *pb.VodStitchDetail {
	if in == nil {
		return nil
	}
	out := &pb.VodStitchDetail{}
	// MISSING: Name
	// MISSING: AdStitchDetails
	return out
}
func VideoVodStitchDetailSpec_FromProto(mapCtx *direct.MapContext, in *pb.VodStitchDetail) *krm.VideoVodStitchDetailSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoVodStitchDetailSpec{}
	// MISSING: Name
	// MISSING: AdStitchDetails
	return out
}
func VideoVodStitchDetailSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoVodStitchDetailSpec) *pb.VodStitchDetail {
	if in == nil {
		return nil
	}
	out := &pb.VodStitchDetail{}
	// MISSING: Name
	// MISSING: AdStitchDetails
	return out
}
func VodStitchDetail_FromProto(mapCtx *direct.MapContext, in *pb.VodStitchDetail) *krm.VodStitchDetail {
	if in == nil {
		return nil
	}
	out := &krm.VodStitchDetail{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AdStitchDetails = direct.Slice_FromProto(mapCtx, in.AdStitchDetails, AdStitchDetail_FromProto)
	return out
}
func VodStitchDetail_ToProto(mapCtx *direct.MapContext, in *krm.VodStitchDetail) *pb.VodStitchDetail {
	if in == nil {
		return nil
	}
	out := &pb.VodStitchDetail{}
	out.Name = direct.ValueOf(in.Name)
	out.AdStitchDetails = direct.Slice_ToProto(mapCtx, in.AdStitchDetails, AdStitchDetail_ToProto)
	return out
}
