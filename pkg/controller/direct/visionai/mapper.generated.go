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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
)
func Series_FromProto(mapCtx *direct.MapContext, in *pb.Series) *krm.Series {
	if in == nil {
		return nil
	}
	out := &krm.Series{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.Stream = direct.LazyPtr(in.GetStream())
	out.Event = direct.LazyPtr(in.GetEvent())
	return out
}
func Series_ToProto(mapCtx *direct.MapContext, in *krm.Series) *pb.Series {
	if in == nil {
		return nil
	}
	out := &pb.Series{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.Stream = direct.ValueOf(in.Stream)
	out.Event = direct.ValueOf(in.Event)
	return out
}
func SeriesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Series) *krm.SeriesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SeriesObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func SeriesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SeriesObservedState) *pb.Series {
	if in == nil {
		return nil
	}
	out := &pb.Series{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiSeriesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Series) *krm.VisionaiSeriesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiSeriesObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiSeriesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiSeriesObservedState) *pb.Series {
	if in == nil {
		return nil
	}
	out := &pb.Series{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiSeriesSpec_FromProto(mapCtx *direct.MapContext, in *pb.Series) *krm.VisionaiSeriesSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiSeriesSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiSeriesSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiSeriesSpec) *pb.Series {
	if in == nil {
		return nil
	}
	out := &pb.Series{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
