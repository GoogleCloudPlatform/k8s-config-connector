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
func Slate_FromProto(mapCtx *direct.MapContext, in *pb.Slate) *krm.Slate {
	if in == nil {
		return nil
	}
	out := &krm.Slate{}
	// MISSING: Name
	out.URI = direct.LazyPtr(in.GetUri())
	out.GamSlate = Slate_GamSlate_FromProto(mapCtx, in.GetGamSlate())
	return out
}
func Slate_ToProto(mapCtx *direct.MapContext, in *krm.Slate) *pb.Slate {
	if in == nil {
		return nil
	}
	out := &pb.Slate{}
	// MISSING: Name
	out.Uri = direct.ValueOf(in.URI)
	out.GamSlate = Slate_GamSlate_ToProto(mapCtx, in.GamSlate)
	return out
}
func SlateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Slate) *krm.SlateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SlateObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: URI
	out.GamSlate = Slate_GamSlateObservedState_FromProto(mapCtx, in.GetGamSlate())
	return out
}
func SlateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SlateObservedState) *pb.Slate {
	if in == nil {
		return nil
	}
	out := &pb.Slate{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: URI
	out.GamSlate = Slate_GamSlateObservedState_ToProto(mapCtx, in.GamSlate)
	return out
}
func Slate_GamSlate_FromProto(mapCtx *direct.MapContext, in *pb.Slate_GamSlate) *krm.Slate_GamSlate {
	if in == nil {
		return nil
	}
	out := &krm.Slate_GamSlate{}
	out.NetworkCode = direct.LazyPtr(in.GetNetworkCode())
	// MISSING: GamSlateID
	return out
}
func Slate_GamSlate_ToProto(mapCtx *direct.MapContext, in *krm.Slate_GamSlate) *pb.Slate_GamSlate {
	if in == nil {
		return nil
	}
	out := &pb.Slate_GamSlate{}
	out.NetworkCode = direct.ValueOf(in.NetworkCode)
	// MISSING: GamSlateID
	return out
}
func Slate_GamSlateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Slate_GamSlate) *krm.Slate_GamSlateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Slate_GamSlateObservedState{}
	// MISSING: NetworkCode
	out.GamSlateID = direct.LazyPtr(in.GetGamSlateId())
	return out
}
func Slate_GamSlateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Slate_GamSlateObservedState) *pb.Slate_GamSlate {
	if in == nil {
		return nil
	}
	out := &pb.Slate_GamSlate{}
	// MISSING: NetworkCode
	out.GamSlateId = direct.ValueOf(in.GamSlateID)
	return out
}
func VideoSlateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Slate) *krm.VideoSlateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoSlateObservedState{}
	// MISSING: Name
	// MISSING: URI
	// MISSING: GamSlate
	return out
}
func VideoSlateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoSlateObservedState) *pb.Slate {
	if in == nil {
		return nil
	}
	out := &pb.Slate{}
	// MISSING: Name
	// MISSING: URI
	// MISSING: GamSlate
	return out
}
func VideoSlateSpec_FromProto(mapCtx *direct.MapContext, in *pb.Slate) *krm.VideoSlateSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoSlateSpec{}
	// MISSING: Name
	// MISSING: URI
	// MISSING: GamSlate
	return out
}
func VideoSlateSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoSlateSpec) *pb.Slate {
	if in == nil {
		return nil
	}
	out := &pb.Slate{}
	// MISSING: Name
	// MISSING: URI
	// MISSING: GamSlate
	return out
}
