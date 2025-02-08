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
func Analysis_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.Analysis {
	if in == nil {
		return nil
	}
	out := &krm.Analysis{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.AnalysisDefinition = AnalysisDefinition_FromProto(mapCtx, in.GetAnalysisDefinition())
	out.InputStreamsMapping = in.InputStreamsMapping
	out.OutputStreamsMapping = in.OutputStreamsMapping
	out.DisableEventWatch = direct.LazyPtr(in.GetDisableEventWatch())
	return out
}
func Analysis_ToProto(mapCtx *direct.MapContext, in *krm.Analysis) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.AnalysisDefinition = AnalysisDefinition_ToProto(mapCtx, in.AnalysisDefinition)
	out.InputStreamsMapping = in.InputStreamsMapping
	out.OutputStreamsMapping = in.OutputStreamsMapping
	out.DisableEventWatch = direct.ValueOf(in.DisableEventWatch)
	return out
}
func AnalysisDefinition_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisDefinition) *krm.AnalysisDefinition {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisDefinition{}
	out.Analyzers = direct.Slice_FromProto(mapCtx, in.Analyzers, AnalyzerDefinition_FromProto)
	return out
}
func AnalysisDefinition_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisDefinition) *pb.AnalysisDefinition {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisDefinition{}
	out.Analyzers = direct.Slice_ToProto(mapCtx, in.Analyzers, AnalyzerDefinition_ToProto)
	return out
}
func AnalysisObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.AnalysisObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: AnalysisDefinition
	// MISSING: InputStreamsMapping
	// MISSING: OutputStreamsMapping
	// MISSING: DisableEventWatch
	return out
}
func AnalysisObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisObservedState) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: AnalysisDefinition
	// MISSING: InputStreamsMapping
	// MISSING: OutputStreamsMapping
	// MISSING: DisableEventWatch
	return out
}
func AnalyzerDefinition_FromProto(mapCtx *direct.MapContext, in *pb.AnalyzerDefinition) *krm.AnalyzerDefinition {
	if in == nil {
		return nil
	}
	out := &krm.AnalyzerDefinition{}
	out.Analyzer = direct.LazyPtr(in.GetAnalyzer())
	out.Operator = direct.LazyPtr(in.GetOperator())
	out.Inputs = direct.Slice_FromProto(mapCtx, in.Inputs, AnalyzerDefinition_StreamInput_FromProto)
	// MISSING: Attrs
	out.DebugOptions = AnalyzerDefinition_DebugOptions_FromProto(mapCtx, in.GetDebugOptions())
	out.OperatorOption = AnalyzerDefinition_OperatorOption_FromProto(mapCtx, in.GetOperatorOption())
	return out
}
func AnalyzerDefinition_ToProto(mapCtx *direct.MapContext, in *krm.AnalyzerDefinition) *pb.AnalyzerDefinition {
	if in == nil {
		return nil
	}
	out := &pb.AnalyzerDefinition{}
	out.Analyzer = direct.ValueOf(in.Analyzer)
	out.Operator = direct.ValueOf(in.Operator)
	out.Inputs = direct.Slice_ToProto(mapCtx, in.Inputs, AnalyzerDefinition_StreamInput_ToProto)
	// MISSING: Attrs
	out.DebugOptions = AnalyzerDefinition_DebugOptions_ToProto(mapCtx, in.DebugOptions)
	out.OperatorOption = AnalyzerDefinition_OperatorOption_ToProto(mapCtx, in.OperatorOption)
	return out
}
func AnalyzerDefinition_DebugOptions_FromProto(mapCtx *direct.MapContext, in *pb.AnalyzerDefinition_DebugOptions) *krm.AnalyzerDefinition_DebugOptions {
	if in == nil {
		return nil
	}
	out := &krm.AnalyzerDefinition_DebugOptions{}
	out.EnvironmentVariables = in.EnvironmentVariables
	return out
}
func AnalyzerDefinition_DebugOptions_ToProto(mapCtx *direct.MapContext, in *krm.AnalyzerDefinition_DebugOptions) *pb.AnalyzerDefinition_DebugOptions {
	if in == nil {
		return nil
	}
	out := &pb.AnalyzerDefinition_DebugOptions{}
	out.EnvironmentVariables = in.EnvironmentVariables
	return out
}
func AnalyzerDefinition_OperatorOption_FromProto(mapCtx *direct.MapContext, in *pb.AnalyzerDefinition_OperatorOption) *krm.AnalyzerDefinition_OperatorOption {
	if in == nil {
		return nil
	}
	out := &krm.AnalyzerDefinition_OperatorOption{}
	out.Tag = direct.LazyPtr(in.GetTag())
	out.Registry = direct.LazyPtr(in.GetRegistry())
	return out
}
func AnalyzerDefinition_OperatorOption_ToProto(mapCtx *direct.MapContext, in *krm.AnalyzerDefinition_OperatorOption) *pb.AnalyzerDefinition_OperatorOption {
	if in == nil {
		return nil
	}
	out := &pb.AnalyzerDefinition_OperatorOption{}
	out.Tag = direct.ValueOf(in.Tag)
	out.Registry = direct.ValueOf(in.Registry)
	return out
}
func AnalyzerDefinition_StreamInput_FromProto(mapCtx *direct.MapContext, in *pb.AnalyzerDefinition_StreamInput) *krm.AnalyzerDefinition_StreamInput {
	if in == nil {
		return nil
	}
	out := &krm.AnalyzerDefinition_StreamInput{}
	out.Input = direct.LazyPtr(in.GetInput())
	return out
}
func AnalyzerDefinition_StreamInput_ToProto(mapCtx *direct.MapContext, in *krm.AnalyzerDefinition_StreamInput) *pb.AnalyzerDefinition_StreamInput {
	if in == nil {
		return nil
	}
	out := &pb.AnalyzerDefinition_StreamInput{}
	out.Input = direct.ValueOf(in.Input)
	return out
}
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
func VisionaiAnalysisObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.VisionaiAnalysisObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiAnalysisObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: AnalysisDefinition
	// MISSING: InputStreamsMapping
	// MISSING: OutputStreamsMapping
	// MISSING: DisableEventWatch
	return out
}
func VisionaiAnalysisObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiAnalysisObservedState) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: AnalysisDefinition
	// MISSING: InputStreamsMapping
	// MISSING: OutputStreamsMapping
	// MISSING: DisableEventWatch
	return out
}
func VisionaiAnalysisSpec_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.VisionaiAnalysisSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiAnalysisSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: AnalysisDefinition
	// MISSING: InputStreamsMapping
	// MISSING: OutputStreamsMapping
	// MISSING: DisableEventWatch
	return out
}
func VisionaiAnalysisSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiAnalysisSpec) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: AnalysisDefinition
	// MISSING: InputStreamsMapping
	// MISSING: OutputStreamsMapping
	// MISSING: DisableEventWatch
	return out
}
