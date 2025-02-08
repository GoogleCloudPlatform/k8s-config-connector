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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
)
func ApplicationNodeAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.ApplicationNodeAnnotation) *krm.ApplicationNodeAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.ApplicationNodeAnnotation{}
	out.Node = direct.LazyPtr(in.GetNode())
	out.Annotations = direct.Slice_FromProto(mapCtx, in.Annotations, StreamAnnotation_FromProto)
	return out
}
func ApplicationNodeAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.ApplicationNodeAnnotation) *pb.ApplicationNodeAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.ApplicationNodeAnnotation{}
	out.Node = direct.ValueOf(in.Node)
	out.Annotations = direct.Slice_ToProto(mapCtx, in.Annotations, StreamAnnotation_ToProto)
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.InstanceType = direct.Enum_FromProto(mapCtx, in.GetInstanceType())
	out.InputResources = direct.Slice_FromProto(mapCtx, in.InputResources, Instance_InputResource_FromProto)
	out.OutputResources = direct.Slice_FromProto(mapCtx, in.OutputResources, Instance_OutputResource_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.InstanceType = direct.Enum_ToProto[pb.Instance_InstanceType](mapCtx, in.InstanceType)
	out.InputResources = direct.Slice_ToProto(mapCtx, in.InputResources, Instance_InputResource_ToProto)
	out.OutputResources = direct.Slice_ToProto(mapCtx, in.OutputResources, Instance_OutputResource_ToProto)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InstanceType
	// MISSING: InputResources
	out.OutputResources = direct.Slice_FromProto(mapCtx, in.OutputResources, Instance_OutputResourceObservedState_FromProto)
	// MISSING: State
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InstanceType
	// MISSING: InputResources
	out.OutputResources = direct.Slice_ToProto(mapCtx, in.OutputResources, Instance_OutputResourceObservedState_ToProto)
	// MISSING: State
	return out
}
func Instance_InputResource_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InputResource) *krm.Instance_InputResource {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InputResource{}
	out.InputResource = direct.LazyPtr(in.GetInputResource())
	out.AnnotatedStream = StreamWithAnnotation_FromProto(mapCtx, in.GetAnnotatedStream())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	out.ConsumerNode = direct.LazyPtr(in.GetConsumerNode())
	out.InputResourceBinding = direct.LazyPtr(in.GetInputResourceBinding())
	out.Annotations = ResourceAnnotations_FromProto(mapCtx, in.GetAnnotations())
	return out
}
func Instance_InputResource_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InputResource) *pb.Instance_InputResource {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InputResource{}
	if oneof := Instance_InputResource_InputResource_ToProto(mapCtx, in.InputResource); oneof != nil {
		out.InputResourceInformation = oneof
	}
	if oneof := StreamWithAnnotation_ToProto(mapCtx, in.AnnotatedStream); oneof != nil {
		out.InputResourceInformation = &pb.Instance_InputResource_AnnotatedStream{AnnotatedStream: oneof}
	}
	out.DataType = direct.Enum_ToProto[pb.DataType](mapCtx, in.DataType)
	out.ConsumerNode = direct.ValueOf(in.ConsumerNode)
	out.InputResourceBinding = direct.ValueOf(in.InputResourceBinding)
	out.Annotations = ResourceAnnotations_ToProto(mapCtx, in.Annotations)
	return out
}
func Instance_OutputResource_FromProto(mapCtx *direct.MapContext, in *pb.Instance_OutputResource) *krm.Instance_OutputResource {
	if in == nil {
		return nil
	}
	out := &krm.Instance_OutputResource{}
	out.OutputResource = direct.LazyPtr(in.GetOutputResource())
	out.ProducerNode = direct.LazyPtr(in.GetProducerNode())
	out.OutputResourceBinding = direct.LazyPtr(in.GetOutputResourceBinding())
	// MISSING: IsTemporary
	// MISSING: Autogen
	return out
}
func Instance_OutputResource_ToProto(mapCtx *direct.MapContext, in *krm.Instance_OutputResource) *pb.Instance_OutputResource {
	if in == nil {
		return nil
	}
	out := &pb.Instance_OutputResource{}
	out.OutputResource = direct.ValueOf(in.OutputResource)
	out.ProducerNode = direct.ValueOf(in.ProducerNode)
	out.OutputResourceBinding = direct.ValueOf(in.OutputResourceBinding)
	// MISSING: IsTemporary
	// MISSING: Autogen
	return out
}
func Instance_OutputResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_OutputResource) *krm.Instance_OutputResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_OutputResourceObservedState{}
	// MISSING: OutputResource
	// MISSING: ProducerNode
	// MISSING: OutputResourceBinding
	out.IsTemporary = direct.LazyPtr(in.GetIsTemporary())
	out.Autogen = direct.LazyPtr(in.GetAutogen())
	return out
}
func Instance_OutputResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_OutputResourceObservedState) *pb.Instance_OutputResource {
	if in == nil {
		return nil
	}
	out := &pb.Instance_OutputResource{}
	// MISSING: OutputResource
	// MISSING: ProducerNode
	// MISSING: OutputResourceBinding
	out.IsTemporary = direct.ValueOf(in.IsTemporary)
	out.Autogen = direct.ValueOf(in.Autogen)
	return out
}
func NormalizedPolygon_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedPolygon) *krm.NormalizedPolygon {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedPolygon{}
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func NormalizedPolygon_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedPolygon) *pb.NormalizedPolygon {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedPolygon{}
	out.NormalizedVertices = direct.Slice_ToProto(mapCtx, in.NormalizedVertices, NormalizedVertex_ToProto)
	return out
}
func NormalizedPolyline_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedPolyline) *krm.NormalizedPolyline {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedPolyline{}
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func NormalizedPolyline_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedPolyline) *pb.NormalizedPolyline {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedPolyline{}
	out.NormalizedVertices = direct.Slice_ToProto(mapCtx, in.NormalizedVertices, NormalizedVertex_ToProto)
	return out
}
func NormalizedVertex_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedVertex) *krm.NormalizedVertex {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedVertex{}
	out.X = direct.LazyPtr(in.GetX())
	out.Y = direct.LazyPtr(in.GetY())
	return out
}
func NormalizedVertex_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedVertex) *pb.NormalizedVertex {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedVertex{}
	out.X = direct.ValueOf(in.X)
	out.Y = direct.ValueOf(in.Y)
	return out
}
func ResourceAnnotations_FromProto(mapCtx *direct.MapContext, in *pb.ResourceAnnotations) *krm.ResourceAnnotations {
	if in == nil {
		return nil
	}
	out := &krm.ResourceAnnotations{}
	out.ApplicationAnnotations = direct.Slice_FromProto(mapCtx, in.ApplicationAnnotations, StreamAnnotation_FromProto)
	out.NodeAnnotations = direct.Slice_FromProto(mapCtx, in.NodeAnnotations, ApplicationNodeAnnotation_FromProto)
	return out
}
func ResourceAnnotations_ToProto(mapCtx *direct.MapContext, in *krm.ResourceAnnotations) *pb.ResourceAnnotations {
	if in == nil {
		return nil
	}
	out := &pb.ResourceAnnotations{}
	out.ApplicationAnnotations = direct.Slice_ToProto(mapCtx, in.ApplicationAnnotations, StreamAnnotation_ToProto)
	out.NodeAnnotations = direct.Slice_ToProto(mapCtx, in.NodeAnnotations, ApplicationNodeAnnotation_ToProto)
	return out
}
func StreamAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.StreamAnnotation) *krm.StreamAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.StreamAnnotation{}
	out.ActiveZone = NormalizedPolygon_FromProto(mapCtx, in.GetActiveZone())
	out.CrossingLine = NormalizedPolyline_FromProto(mapCtx, in.GetCrossingLine())
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SourceStream = direct.LazyPtr(in.GetSourceStream())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func StreamAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.StreamAnnotation) *pb.StreamAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.StreamAnnotation{}
	if oneof := NormalizedPolygon_ToProto(mapCtx, in.ActiveZone); oneof != nil {
		out.AnnotationPayload = &pb.StreamAnnotation_ActiveZone{ActiveZone: oneof}
	}
	if oneof := NormalizedPolyline_ToProto(mapCtx, in.CrossingLine); oneof != nil {
		out.AnnotationPayload = &pb.StreamAnnotation_CrossingLine{CrossingLine: oneof}
	}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SourceStream = direct.ValueOf(in.SourceStream)
	out.Type = direct.Enum_ToProto[pb.StreamAnnotationType](mapCtx, in.Type)
	return out
}
func StreamWithAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.StreamWithAnnotation) *krm.StreamWithAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.StreamWithAnnotation{}
	out.Stream = direct.LazyPtr(in.GetStream())
	out.ApplicationAnnotations = direct.Slice_FromProto(mapCtx, in.ApplicationAnnotations, StreamAnnotation_FromProto)
	out.NodeAnnotations = direct.Slice_FromProto(mapCtx, in.NodeAnnotations, StreamWithAnnotation_NodeAnnotation_FromProto)
	return out
}
func StreamWithAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.StreamWithAnnotation) *pb.StreamWithAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.StreamWithAnnotation{}
	out.Stream = direct.ValueOf(in.Stream)
	out.ApplicationAnnotations = direct.Slice_ToProto(mapCtx, in.ApplicationAnnotations, StreamAnnotation_ToProto)
	out.NodeAnnotations = direct.Slice_ToProto(mapCtx, in.NodeAnnotations, StreamWithAnnotation_NodeAnnotation_ToProto)
	return out
}
func StreamWithAnnotation_NodeAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.StreamWithAnnotation_NodeAnnotation) *krm.StreamWithAnnotation_NodeAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.StreamWithAnnotation_NodeAnnotation{}
	out.Node = direct.LazyPtr(in.GetNode())
	out.Annotations = direct.Slice_FromProto(mapCtx, in.Annotations, StreamAnnotation_FromProto)
	return out
}
func StreamWithAnnotation_NodeAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.StreamWithAnnotation_NodeAnnotation) *pb.StreamWithAnnotation_NodeAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.StreamWithAnnotation_NodeAnnotation{}
	out.Node = direct.ValueOf(in.Node)
	out.Annotations = direct.Slice_ToProto(mapCtx, in.Annotations, StreamAnnotation_ToProto)
	return out
}
func VisionaiInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.VisionaiInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InstanceType
	// MISSING: InputResources
	// MISSING: OutputResources
	// MISSING: State
	return out
}
func VisionaiInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InstanceType
	// MISSING: InputResources
	// MISSING: OutputResources
	// MISSING: State
	return out
}
func VisionaiInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.VisionaiInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InstanceType
	// MISSING: InputResources
	// MISSING: OutputResources
	// MISSING: State
	return out
}
func VisionaiInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InstanceType
	// MISSING: InputResources
	// MISSING: OutputResources
	// MISSING: State
	return out
}
