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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Stream_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.Stream {
	if in == nil {
		return nil
	}
	out := &krm.Stream{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.EnableHlsPlayback = direct.LazyPtr(in.GetEnableHlsPlayback())
	out.MediaWarehouseAsset = direct.LazyPtr(in.GetMediaWarehouseAsset())
	return out
}
func Stream_ToProto(mapCtx *direct.MapContext, in *krm.Stream) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.EnableHlsPlayback = direct.ValueOf(in.EnableHlsPlayback)
	out.MediaWarehouseAsset = direct.ValueOf(in.MediaWarehouseAsset)
	return out
}
func StreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.StreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StreamObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: EnableHlsPlayback
	// MISSING: MediaWarehouseAsset
	return out
}
func StreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: EnableHlsPlayback
	// MISSING: MediaWarehouseAsset
	return out
}
func VisionaiStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.VisionaiStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiStreamObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: EnableHlsPlayback
	// MISSING: MediaWarehouseAsset
	return out
}
func VisionaiStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiStreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: EnableHlsPlayback
	// MISSING: MediaWarehouseAsset
	return out
}
func VisionaiStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.VisionaiStreamSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiStreamSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: EnableHlsPlayback
	// MISSING: MediaWarehouseAsset
	return out
}
func VisionaiStreamSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiStreamSpec) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: EnableHlsPlayback
	// MISSING: MediaWarehouseAsset
	return out
}
