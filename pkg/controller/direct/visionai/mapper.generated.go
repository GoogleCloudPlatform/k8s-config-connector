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
func Event_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.Event {
	if in == nil {
		return nil
	}
	out := &krm.Event{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.AlignmentClock = direct.Enum_FromProto(mapCtx, in.GetAlignmentClock())
	out.GracePeriod = direct.StringDuration_FromProto(mapCtx, in.GetGracePeriod())
	return out
}
func Event_ToProto(mapCtx *direct.MapContext, in *krm.Event) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.AlignmentClock = direct.Enum_ToProto[pb.Event_Clock](mapCtx, in.AlignmentClock)
	out.GracePeriod = direct.StringDuration_ToProto(mapCtx, in.GracePeriod)
	return out
}
func EventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.EventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: AlignmentClock
	// MISSING: GracePeriod
	return out
}
func EventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventObservedState) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: AlignmentClock
	// MISSING: GracePeriod
	return out
}
func VisionaiEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.VisionaiEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiEventObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: AlignmentClock
	// MISSING: GracePeriod
	return out
}
func VisionaiEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiEventObservedState) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: AlignmentClock
	// MISSING: GracePeriod
	return out
}
func VisionaiEventSpec_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.VisionaiEventSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiEventSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: AlignmentClock
	// MISSING: GracePeriod
	return out
}
func VisionaiEventSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiEventSpec) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: AlignmentClock
	// MISSING: GracePeriod
	return out
}
