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
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AttributeValue_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValue) *krm.AttributeValue {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValue{}
	out.I = direct.LazyPtr(in.GetI())
	out.F = direct.LazyPtr(in.GetF())
	out.B = direct.LazyPtr(in.GetB())
	out.S = in.GetS()
	return out
}
func AttributeValue_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValue) *pb.AttributeValue {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValue{}
	if oneof := AttributeValue_I_ToProto(mapCtx, in.I); oneof != nil {
		out.Value = oneof
	}
	if oneof := AttributeValue_F_ToProto(mapCtx, in.F); oneof != nil {
		out.Value = oneof
	}
	if oneof := AttributeValue_B_ToProto(mapCtx, in.B); oneof != nil {
		out.Value = oneof
	}
	if oneof := AttributeValue_S_ToProto(mapCtx, in.S); oneof != nil {
		out.Value = oneof
	}
	return out
}
func Operator_FromProto(mapCtx *direct.MapContext, in *pb.Operator) *krm.Operator {
	if in == nil {
		return nil
	}
	out := &krm.Operator{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.OperatorDefinition = OperatorDefinition_FromProto(mapCtx, in.GetOperatorDefinition())
	out.DockerImage = direct.LazyPtr(in.GetDockerImage())
	return out
}
func Operator_ToProto(mapCtx *direct.MapContext, in *krm.Operator) *pb.Operator {
	if in == nil {
		return nil
	}
	out := &pb.Operator{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.OperatorDefinition = OperatorDefinition_ToProto(mapCtx, in.OperatorDefinition)
	out.DockerImage = direct.ValueOf(in.DockerImage)
	return out
}
func OperatorDefinition_FromProto(mapCtx *direct.MapContext, in *pb.OperatorDefinition) *krm.OperatorDefinition {
	if in == nil {
		return nil
	}
	out := &krm.OperatorDefinition{}
	out.Operator = direct.LazyPtr(in.GetOperator())
	out.InputArgs = direct.Slice_FromProto(mapCtx, in.InputArgs, OperatorDefinition_ArgumentDefinition_FromProto)
	out.OutputArgs = direct.Slice_FromProto(mapCtx, in.OutputArgs, OperatorDefinition_ArgumentDefinition_FromProto)
	out.Attributes = direct.Slice_FromProto(mapCtx, in.Attributes, OperatorDefinition_AttributeDefinition_FromProto)
	out.Resources = ResourceSpecification_FromProto(mapCtx, in.GetResources())
	out.ShortDescription = direct.LazyPtr(in.GetShortDescription())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func OperatorDefinition_ToProto(mapCtx *direct.MapContext, in *krm.OperatorDefinition) *pb.OperatorDefinition {
	if in == nil {
		return nil
	}
	out := &pb.OperatorDefinition{}
	out.Operator = direct.ValueOf(in.Operator)
	out.InputArgs = direct.Slice_ToProto(mapCtx, in.InputArgs, OperatorDefinition_ArgumentDefinition_ToProto)
	out.OutputArgs = direct.Slice_ToProto(mapCtx, in.OutputArgs, OperatorDefinition_ArgumentDefinition_ToProto)
	out.Attributes = direct.Slice_ToProto(mapCtx, in.Attributes, OperatorDefinition_AttributeDefinition_ToProto)
	out.Resources = ResourceSpecification_ToProto(mapCtx, in.Resources)
	out.ShortDescription = direct.ValueOf(in.ShortDescription)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func OperatorDefinition_ArgumentDefinition_FromProto(mapCtx *direct.MapContext, in *pb.OperatorDefinition_ArgumentDefinition) *krm.OperatorDefinition_ArgumentDefinition {
	if in == nil {
		return nil
	}
	out := &krm.OperatorDefinition_ArgumentDefinition{}
	out.Argument = direct.LazyPtr(in.GetArgument())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func OperatorDefinition_ArgumentDefinition_ToProto(mapCtx *direct.MapContext, in *krm.OperatorDefinition_ArgumentDefinition) *pb.OperatorDefinition_ArgumentDefinition {
	if in == nil {
		return nil
	}
	out := &pb.OperatorDefinition_ArgumentDefinition{}
	out.Argument = direct.ValueOf(in.Argument)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func OperatorDefinition_AttributeDefinition_FromProto(mapCtx *direct.MapContext, in *pb.OperatorDefinition_AttributeDefinition) *krm.OperatorDefinition_AttributeDefinition {
	if in == nil {
		return nil
	}
	out := &krm.OperatorDefinition_AttributeDefinition{}
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	out.Type = direct.LazyPtr(in.GetType())
	out.DefaultValue = AttributeValue_FromProto(mapCtx, in.GetDefaultValue())
	return out
}
func OperatorDefinition_AttributeDefinition_ToProto(mapCtx *direct.MapContext, in *krm.OperatorDefinition_AttributeDefinition) *pb.OperatorDefinition_AttributeDefinition {
	if in == nil {
		return nil
	}
	out := &pb.OperatorDefinition_AttributeDefinition{}
	out.Attribute = direct.ValueOf(in.Attribute)
	out.Type = direct.ValueOf(in.Type)
	out.DefaultValue = AttributeValue_ToProto(mapCtx, in.DefaultValue)
	return out
}
func OperatorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Operator) *krm.OperatorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OperatorObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: OperatorDefinition
	// MISSING: DockerImage
	return out
}
func OperatorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OperatorObservedState) *pb.Operator {
	if in == nil {
		return nil
	}
	out := &pb.Operator{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: OperatorDefinition
	// MISSING: DockerImage
	return out
}
func ResourceSpecification_FromProto(mapCtx *direct.MapContext, in *pb.ResourceSpecification) *krm.ResourceSpecification {
	if in == nil {
		return nil
	}
	out := &krm.ResourceSpecification{}
	out.Cpu = direct.LazyPtr(in.GetCpu())
	out.CpuLimits = direct.LazyPtr(in.GetCpuLimits())
	out.Memory = direct.LazyPtr(in.GetMemory())
	out.MemoryLimits = direct.LazyPtr(in.GetMemoryLimits())
	out.Gpus = direct.LazyPtr(in.GetGpus())
	out.LatencyBudgetMs = direct.LazyPtr(in.GetLatencyBudgetMs())
	return out
}
func ResourceSpecification_ToProto(mapCtx *direct.MapContext, in *krm.ResourceSpecification) *pb.ResourceSpecification {
	if in == nil {
		return nil
	}
	out := &pb.ResourceSpecification{}
	out.Cpu = direct.ValueOf(in.Cpu)
	out.CpuLimits = direct.ValueOf(in.CpuLimits)
	out.Memory = direct.ValueOf(in.Memory)
	out.MemoryLimits = direct.ValueOf(in.MemoryLimits)
	out.Gpus = direct.ValueOf(in.Gpus)
	out.LatencyBudgetMs = direct.ValueOf(in.LatencyBudgetMs)
	return out
}
func VisionaiOperatorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Operator) *krm.VisionaiOperatorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiOperatorObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: OperatorDefinition
	// MISSING: DockerImage
	return out
}
func VisionaiOperatorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiOperatorObservedState) *pb.Operator {
	if in == nil {
		return nil
	}
	out := &pb.Operator{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: OperatorDefinition
	// MISSING: DockerImage
	return out
}
func VisionaiOperatorSpec_FromProto(mapCtx *direct.MapContext, in *pb.Operator) *krm.VisionaiOperatorSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiOperatorSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: OperatorDefinition
	// MISSING: DockerImage
	return out
}
func VisionaiOperatorSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiOperatorSpec) *pb.Operator {
	if in == nil {
		return nil
	}
	out := &pb.Operator{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: OperatorDefinition
	// MISSING: DockerImage
	return out
}
