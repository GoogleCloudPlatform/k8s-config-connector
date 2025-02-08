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
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Clip_FromProto(mapCtx *direct.MapContext, in *pb.Clip) *krm.Clip {
	if in == nil {
		return nil
	}
	out := &krm.Clip{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	out.OutputURI = direct.LazyPtr(in.GetOutputUri())
	// MISSING: Error
	out.Slices = direct.Slice_FromProto(mapCtx, in.Slices, Clip_Slice_FromProto)
	out.ClipManifests = direct.Slice_FromProto(mapCtx, in.ClipManifests, Clip_ClipManifest_FromProto)
	return out
}
func Clip_ToProto(mapCtx *direct.MapContext, in *krm.Clip) *pb.Clip {
	if in == nil {
		return nil
	}
	out := &pb.Clip{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	out.OutputUri = direct.ValueOf(in.OutputURI)
	// MISSING: Error
	out.Slices = direct.Slice_ToProto(mapCtx, in.Slices, Clip_Slice_ToProto)
	out.ClipManifests = direct.Slice_ToProto(mapCtx, in.ClipManifests, Clip_ClipManifest_ToProto)
	return out
}
func ClipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Clip) *krm.ClipObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClipObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: OutputURI
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Slices
	out.ClipManifests = direct.Slice_FromProto(mapCtx, in.ClipManifests, Clip_ClipManifestObservedState_FromProto)
	return out
}
func ClipObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClipObservedState) *pb.Clip {
	if in == nil {
		return nil
	}
	out := &pb.Clip{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Clip_State](mapCtx, in.State)
	// MISSING: OutputURI
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Slices
	out.ClipManifests = direct.Slice_ToProto(mapCtx, in.ClipManifests, Clip_ClipManifestObservedState_ToProto)
	return out
}
func Clip_ClipManifest_FromProto(mapCtx *direct.MapContext, in *pb.Clip_ClipManifest) *krm.Clip_ClipManifest {
	if in == nil {
		return nil
	}
	out := &krm.Clip_ClipManifest{}
	out.ManifestKey = direct.LazyPtr(in.GetManifestKey())
	// MISSING: OutputURI
	return out
}
func Clip_ClipManifest_ToProto(mapCtx *direct.MapContext, in *krm.Clip_ClipManifest) *pb.Clip_ClipManifest {
	if in == nil {
		return nil
	}
	out := &pb.Clip_ClipManifest{}
	out.ManifestKey = direct.ValueOf(in.ManifestKey)
	// MISSING: OutputURI
	return out
}
func Clip_ClipManifestObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Clip_ClipManifest) *krm.Clip_ClipManifestObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Clip_ClipManifestObservedState{}
	// MISSING: ManifestKey
	out.OutputURI = direct.LazyPtr(in.GetOutputUri())
	return out
}
func Clip_ClipManifestObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Clip_ClipManifestObservedState) *pb.Clip_ClipManifest {
	if in == nil {
		return nil
	}
	out := &pb.Clip_ClipManifest{}
	// MISSING: ManifestKey
	out.OutputUri = direct.ValueOf(in.OutputURI)
	return out
}
func Clip_Slice_FromProto(mapCtx *direct.MapContext, in *pb.Clip_Slice) *krm.Clip_Slice {
	if in == nil {
		return nil
	}
	out := &krm.Clip_Slice{}
	out.TimeSlice = Clip_TimeSlice_FromProto(mapCtx, in.GetTimeSlice())
	return out
}
func Clip_Slice_ToProto(mapCtx *direct.MapContext, in *krm.Clip_Slice) *pb.Clip_Slice {
	if in == nil {
		return nil
	}
	out := &pb.Clip_Slice{}
	if oneof := Clip_TimeSlice_ToProto(mapCtx, in.TimeSlice); oneof != nil {
		out.Kind = &pb.Clip_Slice_TimeSlice{TimeSlice: oneof}
	}
	return out
}
func Clip_TimeSlice_FromProto(mapCtx *direct.MapContext, in *pb.Clip_TimeSlice) *krm.Clip_TimeSlice {
	if in == nil {
		return nil
	}
	out := &krm.Clip_TimeSlice{}
	out.MarkinTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMarkinTime())
	out.MarkoutTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMarkoutTime())
	return out
}
func Clip_TimeSlice_ToProto(mapCtx *direct.MapContext, in *krm.Clip_TimeSlice) *pb.Clip_TimeSlice {
	if in == nil {
		return nil
	}
	out := &pb.Clip_TimeSlice{}
	out.MarkinTime = direct.StringTimestamp_ToProto(mapCtx, in.MarkinTime)
	out.MarkoutTime = direct.StringTimestamp_ToProto(mapCtx, in.MarkoutTime)
	return out
}
func VideoClipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Clip) *krm.VideoClipObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoClipObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: OutputURI
	// MISSING: Error
	// MISSING: Slices
	// MISSING: ClipManifests
	return out
}
func VideoClipObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoClipObservedState) *pb.Clip {
	if in == nil {
		return nil
	}
	out := &pb.Clip{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: OutputURI
	// MISSING: Error
	// MISSING: Slices
	// MISSING: ClipManifests
	return out
}
func VideoClipSpec_FromProto(mapCtx *direct.MapContext, in *pb.Clip) *krm.VideoClipSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoClipSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: OutputURI
	// MISSING: Error
	// MISSING: Slices
	// MISSING: ClipManifests
	return out
}
func VideoClipSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoClipSpec) *pb.Clip {
	if in == nil {
		return nil
	}
	out := &pb.Clip{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: OutputURI
	// MISSING: Error
	// MISSING: Slices
	// MISSING: ClipManifests
	return out
}
