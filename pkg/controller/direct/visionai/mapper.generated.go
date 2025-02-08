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
func Channel_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.Channel {
	if in == nil {
		return nil
	}
	out := &krm.Channel{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.Stream = direct.LazyPtr(in.GetStream())
	out.Event = direct.LazyPtr(in.GetEvent())
	return out
}
func Channel_ToProto(mapCtx *direct.MapContext, in *krm.Channel) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.Stream = direct.ValueOf(in.Stream)
	out.Event = direct.ValueOf(in.Event)
	return out
}
func ChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.ChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func ChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.VisionaiChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiChannelObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.VisionaiChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiChannelSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
func VisionaiChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: Stream
	// MISSING: Event
	return out
}
