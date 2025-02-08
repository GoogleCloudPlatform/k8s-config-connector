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
func CustomProcessorSourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.CustomProcessorSourceInfo) *krm.CustomProcessorSourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.CustomProcessorSourceInfo{}
	out.VertexModel = direct.LazyPtr(in.GetVertexModel())
	out.ProductRecognizerArtifact = CustomProcessorSourceInfo_ProductRecognizerArtifact_FromProto(mapCtx, in.GetProductRecognizerArtifact())
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	// MISSING: AdditionalInfo
	out.ModelSchema = CustomProcessorSourceInfo_ModelSchema_FromProto(mapCtx, in.GetModelSchema())
	return out
}
func CustomProcessorSourceInfo_ToProto(mapCtx *direct.MapContext, in *krm.CustomProcessorSourceInfo) *pb.CustomProcessorSourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.CustomProcessorSourceInfo{}
	if oneof := CustomProcessorSourceInfo_VertexModel_ToProto(mapCtx, in.VertexModel); oneof != nil {
		out.ArtifactPath = oneof
	}
	if oneof := CustomProcessorSourceInfo_ProductRecognizerArtifact_ToProto(mapCtx, in.ProductRecognizerArtifact); oneof != nil {
		out.ArtifactPath = &pb.CustomProcessorSourceInfo_ProductRecognizerArtifact_{ProductRecognizerArtifact: oneof}
	}
	out.SourceType = direct.Enum_ToProto[pb.CustomProcessorSourceInfo_SourceType](mapCtx, in.SourceType)
	// MISSING: AdditionalInfo
	out.ModelSchema = CustomProcessorSourceInfo_ModelSchema_ToProto(mapCtx, in.ModelSchema)
	return out
}
func CustomProcessorSourceInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomProcessorSourceInfo) *krm.CustomProcessorSourceInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomProcessorSourceInfoObservedState{}
	// MISSING: VertexModel
	// MISSING: ProductRecognizerArtifact
	// MISSING: SourceType
	out.AdditionalInfo = in.AdditionalInfo
	// MISSING: ModelSchema
	return out
}
func CustomProcessorSourceInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomProcessorSourceInfoObservedState) *pb.CustomProcessorSourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.CustomProcessorSourceInfo{}
	// MISSING: VertexModel
	// MISSING: ProductRecognizerArtifact
	// MISSING: SourceType
	out.AdditionalInfo = in.AdditionalInfo
	// MISSING: ModelSchema
	return out
}
func CustomProcessorSourceInfo_ModelSchema_FromProto(mapCtx *direct.MapContext, in *pb.CustomProcessorSourceInfo_ModelSchema) *krm.CustomProcessorSourceInfo_ModelSchema {
	if in == nil {
		return nil
	}
	out := &krm.CustomProcessorSourceInfo_ModelSchema{}
	out.InstancesSchema = GcsSource_FromProto(mapCtx, in.GetInstancesSchema())
	out.ParametersSchema = GcsSource_FromProto(mapCtx, in.GetParametersSchema())
	out.PredictionsSchema = GcsSource_FromProto(mapCtx, in.GetPredictionsSchema())
	return out
}
func CustomProcessorSourceInfo_ModelSchema_ToProto(mapCtx *direct.MapContext, in *krm.CustomProcessorSourceInfo_ModelSchema) *pb.CustomProcessorSourceInfo_ModelSchema {
	if in == nil {
		return nil
	}
	out := &pb.CustomProcessorSourceInfo_ModelSchema{}
	out.InstancesSchema = GcsSource_ToProto(mapCtx, in.InstancesSchema)
	out.ParametersSchema = GcsSource_ToProto(mapCtx, in.ParametersSchema)
	out.PredictionsSchema = GcsSource_ToProto(mapCtx, in.PredictionsSchema)
	return out
}
func CustomProcessorSourceInfo_ProductRecognizerArtifact_FromProto(mapCtx *direct.MapContext, in *pb.CustomProcessorSourceInfo_ProductRecognizerArtifact) *krm.CustomProcessorSourceInfo_ProductRecognizerArtifact {
	if in == nil {
		return nil
	}
	out := &krm.CustomProcessorSourceInfo_ProductRecognizerArtifact{}
	out.RetailProductRecognitionIndex = direct.LazyPtr(in.GetRetailProductRecognitionIndex())
	out.VertexModel = direct.LazyPtr(in.GetVertexModel())
	return out
}
func CustomProcessorSourceInfo_ProductRecognizerArtifact_ToProto(mapCtx *direct.MapContext, in *krm.CustomProcessorSourceInfo_ProductRecognizerArtifact) *pb.CustomProcessorSourceInfo_ProductRecognizerArtifact {
	if in == nil {
		return nil
	}
	out := &pb.CustomProcessorSourceInfo_ProductRecognizerArtifact{}
	out.RetailProductRecognitionIndex = direct.ValueOf(in.RetailProductRecognitionIndex)
	out.VertexModel = direct.ValueOf(in.VertexModel)
	return out
}
func GcsSource_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSource {
	if in == nil {
		return nil
	}
	out := &krm.GcsSource{}
	out.Uris = in.Uris
	return out
}
func GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GcsSource) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	out.Uris = in.Uris
	return out
}
func Processor_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.Processor {
	if in == nil {
		return nil
	}
	out := &krm.Processor{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: ProcessorType
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.CustomProcessorSourceInfo = CustomProcessorSourceInfo_FromProto(mapCtx, in.GetCustomProcessorSourceInfo())
	// MISSING: State
	// MISSING: ProcessorIoSpec
	// MISSING: ConfigurationTypeurl
	// MISSING: SupportedAnnotationTypes
	out.SupportsPostProcessing = direct.LazyPtr(in.GetSupportsPostProcessing())
	out.SupportedInstanceTypes = direct.EnumSlice_FromProto(mapCtx, in.SupportedInstanceTypes)
	return out
}
func Processor_ToProto(mapCtx *direct.MapContext, in *krm.Processor) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: ProcessorType
	out.ModelType = direct.Enum_ToProto[pb.ModelType](mapCtx, in.ModelType)
	out.CustomProcessorSourceInfo = CustomProcessorSourceInfo_ToProto(mapCtx, in.CustomProcessorSourceInfo)
	// MISSING: State
	// MISSING: ProcessorIoSpec
	// MISSING: ConfigurationTypeurl
	// MISSING: SupportedAnnotationTypes
	out.SupportsPostProcessing = direct.ValueOf(in.SupportsPostProcessing)
	out.SupportedInstanceTypes = direct.EnumSlice_ToProto[pb.Instance_InstanceType](mapCtx, in.SupportedInstanceTypes)
	return out
}
func ProcessorIOSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorIOSpec) *krm.ProcessorIOSpec {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorIOSpec{}
	out.GraphInputChannelSpecs = direct.Slice_FromProto(mapCtx, in.GraphInputChannelSpecs, ProcessorIOSpec_GraphInputChannelSpec_FromProto)
	out.GraphOutputChannelSpecs = direct.Slice_FromProto(mapCtx, in.GraphOutputChannelSpecs, ProcessorIOSpec_GraphOutputChannelSpec_FromProto)
	out.InstanceResourceInputBindingSpecs = direct.Slice_FromProto(mapCtx, in.InstanceResourceInputBindingSpecs, ProcessorIOSpec_InstanceResourceInputBindingSpec_FromProto)
	out.InstanceResourceOutputBindingSpecs = direct.Slice_FromProto(mapCtx, in.InstanceResourceOutputBindingSpecs, ProcessorIOSpec_InstanceResourceOutputBindingSpec_FromProto)
	return out
}
func ProcessorIOSpec_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorIOSpec) *pb.ProcessorIOSpec {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorIOSpec{}
	out.GraphInputChannelSpecs = direct.Slice_ToProto(mapCtx, in.GraphInputChannelSpecs, ProcessorIOSpec_GraphInputChannelSpec_ToProto)
	out.GraphOutputChannelSpecs = direct.Slice_ToProto(mapCtx, in.GraphOutputChannelSpecs, ProcessorIOSpec_GraphOutputChannelSpec_ToProto)
	out.InstanceResourceInputBindingSpecs = direct.Slice_ToProto(mapCtx, in.InstanceResourceInputBindingSpecs, ProcessorIOSpec_InstanceResourceInputBindingSpec_ToProto)
	out.InstanceResourceOutputBindingSpecs = direct.Slice_ToProto(mapCtx, in.InstanceResourceOutputBindingSpecs, ProcessorIOSpec_InstanceResourceOutputBindingSpec_ToProto)
	return out
}
func ProcessorIOSpec_GraphInputChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorIOSpec_GraphInputChannelSpec) *krm.ProcessorIOSpec_GraphInputChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorIOSpec_GraphInputChannelSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	out.AcceptedDataTypeUris = in.AcceptedDataTypeUris
	out.Required = direct.LazyPtr(in.GetRequired())
	out.MaxConnectionAllowed = direct.LazyPtr(in.GetMaxConnectionAllowed())
	return out
}
func ProcessorIOSpec_GraphInputChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorIOSpec_GraphInputChannelSpec) *pb.ProcessorIOSpec_GraphInputChannelSpec {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorIOSpec_GraphInputChannelSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.DataType = direct.Enum_ToProto[pb.DataType](mapCtx, in.DataType)
	out.AcceptedDataTypeUris = in.AcceptedDataTypeUris
	out.Required = direct.ValueOf(in.Required)
	out.MaxConnectionAllowed = direct.ValueOf(in.MaxConnectionAllowed)
	return out
}
func ProcessorIOSpec_GraphOutputChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorIOSpec_GraphOutputChannelSpec) *krm.ProcessorIOSpec_GraphOutputChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorIOSpec_GraphOutputChannelSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	out.DataTypeURI = direct.LazyPtr(in.GetDataTypeUri())
	return out
}
func ProcessorIOSpec_GraphOutputChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorIOSpec_GraphOutputChannelSpec) *pb.ProcessorIOSpec_GraphOutputChannelSpec {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorIOSpec_GraphOutputChannelSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.DataType = direct.Enum_ToProto[pb.DataType](mapCtx, in.DataType)
	out.DataTypeUri = direct.ValueOf(in.DataTypeURI)
	return out
}
func ProcessorIOSpec_InstanceResourceInputBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorIOSpec_InstanceResourceInputBindingSpec) *krm.ProcessorIOSpec_InstanceResourceInputBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorIOSpec_InstanceResourceInputBindingSpec{}
	out.ConfigTypeURI = direct.LazyPtr(in.GetConfigTypeUri())
	out.ResourceTypeURI = direct.LazyPtr(in.GetResourceTypeUri())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func ProcessorIOSpec_InstanceResourceInputBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorIOSpec_InstanceResourceInputBindingSpec) *pb.ProcessorIOSpec_InstanceResourceInputBindingSpec {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorIOSpec_InstanceResourceInputBindingSpec{}
	if oneof := ProcessorIOSpec_InstanceResourceInputBindingSpec_ConfigTypeUri_ToProto(mapCtx, in.ConfigTypeURI); oneof != nil {
		out.ResourceType = oneof
	}
	if oneof := ProcessorIOSpec_InstanceResourceInputBindingSpec_ResourceTypeUri_ToProto(mapCtx, in.ResourceTypeURI); oneof != nil {
		out.ResourceType = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func ProcessorIOSpec_InstanceResourceOutputBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorIOSpec_InstanceResourceOutputBindingSpec) *krm.ProcessorIOSpec_InstanceResourceOutputBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorIOSpec_InstanceResourceOutputBindingSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ResourceTypeURI = direct.LazyPtr(in.GetResourceTypeUri())
	out.Explicit = direct.LazyPtr(in.GetExplicit())
	return out
}
func ProcessorIOSpec_InstanceResourceOutputBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorIOSpec_InstanceResourceOutputBindingSpec) *pb.ProcessorIOSpec_InstanceResourceOutputBindingSpec {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorIOSpec_InstanceResourceOutputBindingSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.ResourceTypeUri = direct.ValueOf(in.ResourceTypeURI)
	out.Explicit = direct.ValueOf(in.Explicit)
	return out
}
func ProcessorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.ProcessorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	out.ProcessorType = direct.Enum_FromProto(mapCtx, in.GetProcessorType())
	// MISSING: ModelType
	out.CustomProcessorSourceInfo = CustomProcessorSourceInfoObservedState_FromProto(mapCtx, in.GetCustomProcessorSourceInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ProcessorIoSpec = ProcessorIOSpec_FromProto(mapCtx, in.GetProcessorIoSpec())
	out.ConfigurationTypeurl = direct.LazyPtr(in.GetConfigurationTypeurl())
	out.SupportedAnnotationTypes = direct.EnumSlice_FromProto(mapCtx, in.SupportedAnnotationTypes)
	// MISSING: SupportsPostProcessing
	// MISSING: SupportedInstanceTypes
	return out
}
func ProcessorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorObservedState) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	out.ProcessorType = direct.Enum_ToProto[pb.Processor_ProcessorType](mapCtx, in.ProcessorType)
	// MISSING: ModelType
	out.CustomProcessorSourceInfo = CustomProcessorSourceInfoObservedState_ToProto(mapCtx, in.CustomProcessorSourceInfo)
	out.State = direct.Enum_ToProto[pb.Processor_ProcessorState](mapCtx, in.State)
	out.ProcessorIoSpec = ProcessorIOSpec_ToProto(mapCtx, in.ProcessorIoSpec)
	out.ConfigurationTypeurl = direct.ValueOf(in.ConfigurationTypeurl)
	out.SupportedAnnotationTypes = direct.EnumSlice_ToProto[pb.StreamAnnotationType](mapCtx, in.SupportedAnnotationTypes)
	// MISSING: SupportsPostProcessing
	// MISSING: SupportedInstanceTypes
	return out
}
func VisionaiProcessorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.VisionaiProcessorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiProcessorObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ProcessorType
	// MISSING: ModelType
	// MISSING: CustomProcessorSourceInfo
	// MISSING: State
	// MISSING: ProcessorIoSpec
	// MISSING: ConfigurationTypeurl
	// MISSING: SupportedAnnotationTypes
	// MISSING: SupportsPostProcessing
	// MISSING: SupportedInstanceTypes
	return out
}
func VisionaiProcessorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiProcessorObservedState) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ProcessorType
	// MISSING: ModelType
	// MISSING: CustomProcessorSourceInfo
	// MISSING: State
	// MISSING: ProcessorIoSpec
	// MISSING: ConfigurationTypeurl
	// MISSING: SupportedAnnotationTypes
	// MISSING: SupportsPostProcessing
	// MISSING: SupportedInstanceTypes
	return out
}
func VisionaiProcessorSpec_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.VisionaiProcessorSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiProcessorSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ProcessorType
	// MISSING: ModelType
	// MISSING: CustomProcessorSourceInfo
	// MISSING: State
	// MISSING: ProcessorIoSpec
	// MISSING: ConfigurationTypeurl
	// MISSING: SupportedAnnotationTypes
	// MISSING: SupportsPostProcessing
	// MISSING: SupportedInstanceTypes
	return out
}
func VisionaiProcessorSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiProcessorSpec) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ProcessorType
	// MISSING: ModelType
	// MISSING: CustomProcessorSourceInfo
	// MISSING: State
	// MISSING: ProcessorIoSpec
	// MISSING: ConfigurationTypeurl
	// MISSING: SupportedAnnotationTypes
	// MISSING: SupportsPostProcessing
	// MISSING: SupportedInstanceTypes
	return out
}
