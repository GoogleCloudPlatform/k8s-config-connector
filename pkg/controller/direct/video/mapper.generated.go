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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
)
func LiveAdTagDetail_FromProto(mapCtx *direct.MapContext, in *pb.LiveAdTagDetail) *krm.LiveAdTagDetail {
	if in == nil {
		return nil
	}
	out := &krm.LiveAdTagDetail{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AdRequests = direct.Slice_FromProto(mapCtx, in.AdRequests, AdRequest_FromProto)
	return out
}
func LiveAdTagDetail_ToProto(mapCtx *direct.MapContext, in *krm.LiveAdTagDetail) *pb.LiveAdTagDetail {
	if in == nil {
		return nil
	}
	out := &pb.LiveAdTagDetail{}
	out.Name = direct.ValueOf(in.Name)
	out.AdRequests = direct.Slice_ToProto(mapCtx, in.AdRequests, AdRequest_ToProto)
	return out
}
func VideoLiveAdTagDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LiveAdTagDetail) *krm.VideoLiveAdTagDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoLiveAdTagDetailObservedState{}
	// MISSING: Name
	// MISSING: AdRequests
	return out
}
func VideoLiveAdTagDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoLiveAdTagDetailObservedState) *pb.LiveAdTagDetail {
	if in == nil {
		return nil
	}
	out := &pb.LiveAdTagDetail{}
	// MISSING: Name
	// MISSING: AdRequests
	return out
}
func VideoLiveAdTagDetailSpec_FromProto(mapCtx *direct.MapContext, in *pb.LiveAdTagDetail) *krm.VideoLiveAdTagDetailSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoLiveAdTagDetailSpec{}
	// MISSING: Name
	// MISSING: AdRequests
	return out
}
func VideoLiveAdTagDetailSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoLiveAdTagDetailSpec) *pb.LiveAdTagDetail {
	if in == nil {
		return nil
	}
	out := &pb.LiveAdTagDetail{}
	// MISSING: Name
	// MISSING: AdRequests
	return out
}
