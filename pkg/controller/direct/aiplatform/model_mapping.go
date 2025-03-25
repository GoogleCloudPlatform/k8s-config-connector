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

package aiplatform

import (
	"strconv"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func AIPlatformModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AIPlatformModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformModelObservedState{}
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	out.VersionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionCreateTime())
	out.VersionUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionUpdateTime())
	out.SupportedExportFormats = direct.Slice_FromProto(mapCtx, in.SupportedExportFormats, Model_ExportFormat_FromProto)
	out.TrainingPipeline = direct.LazyPtr(in.GetTrainingPipeline())
	out.SupportedDeploymentResourcesTypes = direct.EnumSlice_FromProto(mapCtx, in.SupportedDeploymentResourcesTypes)
	out.SupportedInputStorageFormats = in.SupportedInputStorageFormats
	out.SupportedOutputStorageFormats = in.SupportedOutputStorageFormats
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeployedModels = direct.Slice_FromProto(mapCtx, in.DeployedModels, DeployedModelRef_FromProto)
	out.ModelSourceInfo = ModelSourceInfo_FromProto(mapCtx, in.GetModelSourceInfo())
	out.OriginalModelInfo = Model_OriginalModelInfo_FromProto(mapCtx, in.GetOriginalModelInfo())
	out.MetadataArtifact = direct.LazyPtr(in.GetMetadataArtifact())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())

	return out
}
func AIPlatformModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.VersionId = direct.ValueOf(in.VersionID)
	out.VersionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.VersionCreateTime)
	out.VersionUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.VersionUpdateTime)
	out.SupportedExportFormats = direct.Slice_ToProto(mapCtx, in.SupportedExportFormats, Model_ExportFormat_ToProto)
	out.TrainingPipeline = direct.ValueOf(in.TrainingPipeline)
	out.SupportedDeploymentResourcesTypes = direct.EnumSlice_ToProto[pb.Model_DeploymentResourcesType](mapCtx, in.SupportedDeploymentResourcesTypes)
	out.SupportedInputStorageFormats = in.SupportedInputStorageFormats
	out.SupportedOutputStorageFormats = in.SupportedOutputStorageFormats
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeployedModels = direct.Slice_ToProto(mapCtx, in.DeployedModels, DeployedModelRef_ToProto)
	out.ModelSourceInfo = ModelSourceInfo_ToProto(mapCtx, in.ModelSourceInfo)
	out.OriginalModelInfo = Model_OriginalModelInfo_ToProto(mapCtx, in.OriginalModelInfo)
	out.MetadataArtifact = direct.ValueOf(in.MetadataArtifact)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func AIPlatformModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AIPlatformModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformModelSpec{}
	out.VersionAliases = in.VersionAliases
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.VersionDescription = direct.LazyPtr(in.GetVersionDescription())
	out.PredictSchemata = PredictSchemata_FromProto(mapCtx, in.GetPredictSchemata())
	out.MetadataSchemaURI = direct.LazyPtr(in.GetMetadataSchemaUri())
	out.Metadata = Value_FromProto(mapCtx, in.GetMetadata())
	out.PipelineJob = direct.LazyPtr(in.GetPipelineJob())
	out.ContainerSpec = ModelContainerSpec_FromProto(mapCtx, in.GetContainerSpec())
	out.ArtifactURI = direct.LazyPtr(in.GetArtifactUri())
	out.ExplanationSpec = ExplanationSpec_FromProto(mapCtx, in.GetExplanationSpec())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.DataStats = Model_DataStats_FromProto(mapCtx, in.GetDataStats())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.BaseModelSource = Model_BaseModelSource_FromProto(mapCtx, in.GetBaseModelSource())
	return out
}
func AIPlatformModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformModelSpec) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.VersionAliases = in.VersionAliases
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.VersionDescription = direct.ValueOf(in.VersionDescription)
	out.PredictSchemata = PredictSchemata_ToProto(mapCtx, in.PredictSchemata)
	out.MetadataSchemaUri = direct.ValueOf(in.MetadataSchemaURI)
	out.Metadata = Value_ToProto(mapCtx, in.Metadata)
	out.PipelineJob = direct.ValueOf(in.PipelineJob)
	out.ContainerSpec = ModelContainerSpec_ToProto(mapCtx, in.ContainerSpec)
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	out.ExplanationSpec = ExplanationSpec_ToProto(mapCtx, in.ExplanationSpec)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.DataStats = Model_DataStats_ToProto(mapCtx, in.DataStats)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.BaseModelSource = Model_BaseModelSource_ToProto(mapCtx, in.BaseModelSource)
	return out
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{
			BoolValue: direct.ValueOf(in.BoolValue),
		}
	}
	if in.NullValue != nil {
		value, err := strconv.Atoi(direct.ValueOf(in.NullValue))
		if err != nil {
			mapCtx.Errorf("error converting value %s from string to int", direct.ValueOf(in.NullValue))
		}
		out.Kind = &structpb.Value_NullValue{
			NullValue: structpb.NullValue(value),
		}
	}
	if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{
			NumberValue: direct.ValueOf(in.NumberValue),
		}
	}
	if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{
			StringValue: direct.ValueOf(in.StringValue),
		}
	}
	if in.StructValue != nil {
		out.Kind = &structpb.Value_StructValue{
			StructValue: StructValue_ToProto(mapCtx, in.StructValue),
		}
	}
	return out
}

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch in.GetKind().(type) {
	case *structpb.Value_StringValue:
		value := in.GetStringValue()
		out.StringValue = &value
	case *structpb.Value_NumberValue:
		value := in.GetNumberValue()
		out.NumberValue = &value
	case *structpb.Value_NullValue:
		value := in.GetNullValue().String()
		out.NullValue = &value
	case *structpb.Value_BoolValue:
		value := in.GetBoolValue()
		out.BoolValue = &value
	case *structpb.Value_StructValue:
		out.StructValue = StructValue_FromProto(mapCtx, in.GetStructValue())
	}
	return out
}

func StructValue_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for key, val := range in.Fields {
		out[key] = val.GetStringValue()
	}
	return out
}

func StructValue_ToProto(mapCtx *direct.MapContext, in map[string]string) *structpb.Struct {
	if in == nil {
		return nil
	}
	out := &structpb.Struct{}
	if len(in) > 0 {
		out.Fields = make(map[string]*structpb.Value)
	}
	for key, val := range in {
		value := &structpb.Value_StringValue{
			StringValue: val,
		}
		out.Fields[key] = &structpb.Value{
			Kind: value,
		}
	}
	return out
}

func ExplanationMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata) *krm.ExplanationMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationMetadata{}
	out.Inputs = make(map[string]*krm.ExplanationMetadata_InputMetadata)
	for k, v := range in.Inputs {
		out.Inputs[k] = ExplanationMetadata_InputMetadata_FromProto(mapCtx, v)
	}
	out.Outputs = make(map[string]*krm.ExplanationMetadata_OutputMetadata)
	for k, v := range in.Outputs {
		out.Outputs[k] = ExplanationMetadata_OutputMetadata_FromProto(mapCtx, v)
	}
	out.FeatureAttributionsSchemaURI = direct.LazyPtr(in.FeatureAttributionsSchemaUri)
	out.LatentSpaceSource = direct.LazyPtr(in.LatentSpaceSource)
	return out
}

func ExplanationMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationMetadata) *pb.ExplanationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata{}
	out.Inputs = make(map[string]*pb.ExplanationMetadata_InputMetadata)
	for k, v := range in.Inputs {
		out.Inputs[k] = ExplanationMetadata_InputMetadata_ToProto(mapCtx, v)
	}
	out.Outputs = make(map[string]*pb.ExplanationMetadata_OutputMetadata)
	for k, v := range in.Outputs {
		out.Outputs[k] = ExplanationMetadata_OutputMetadata_ToProto(mapCtx, v)
	}
	out.FeatureAttributionsSchemaUri = direct.ValueOf(in.FeatureAttributionsSchemaURI)
	out.LatentSpaceSource = direct.ValueOf(in.LatentSpaceSource)
	return out
}

func ExplanationMetadata_InputMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata_InputMetadata) *krm.ExplanationMetadata_InputMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationMetadata_InputMetadata{}
	out.GroupName = direct.LazyPtr(in.GetGroupName())
	out.IndicesTensorName = direct.LazyPtr(in.GetIndicesTensorName())
	out.InputTensorName = direct.LazyPtr(in.GetInputTensorName())
	out.EncodedTensorName = direct.LazyPtr(in.GetEncodedTensorName())
	out.DenseShapeTensorName = direct.LazyPtr(in.GetDenseShapeTensorName())
	out.Encoding = direct.LazyPtr(in.GetEncoding().String())
	out.Modality = direct.LazyPtr(in.GetModality())
	out.IndexFeatureMapping = in.GetIndexFeatureMapping()
	out.FeatureValueDomain = ExplanationMetadata_InputMetadata_FeatureValueDomain_FromProto(mapCtx, in.GetFeatureValueDomain())
	out.Visualization = ExplanationMetadata_InputMetadata_Visualization_FromProto(mapCtx, in.GetVisualization())
	out.InputBaselines = direct.Slice_FromProto(mapCtx, in.InputBaselines, Value_FromProto)
	out.EncodedBaselines = direct.Slice_FromProto(mapCtx, in.EncodedBaselines, Value_FromProto)
	return out
}

func ExplanationMetadata_InputMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationMetadata_InputMetadata) *pb.ExplanationMetadata_InputMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata_InputMetadata{}
	out.GroupName = direct.ValueOf(in.GroupName)
	out.IndicesTensorName = direct.ValueOf(in.IndicesTensorName)
	out.InputTensorName = direct.ValueOf(in.InputTensorName)
	out.EncodedTensorName = direct.ValueOf(in.EncodedTensorName)
	out.DenseShapeTensorName = direct.ValueOf(in.DenseShapeTensorName)
	out.Encoding = pb.ExplanationMetadata_InputMetadata_Encoding(pb.ExplanationMetadata_InputMetadata_Encoding_value[direct.ValueOf(in.Encoding)])
	out.Modality = direct.ValueOf(in.Modality)
	out.IndexFeatureMapping = in.IndexFeatureMapping
	out.EncodedBaselines = direct.Slice_ToProto(mapCtx, in.EncodedBaselines, Value_ToProto)
	out.FeatureValueDomain = ExplanationMetadata_InputMetadata_FeatureValueDomain_ToProto(mapCtx, in.FeatureValueDomain)
	out.Visualization = ExplanationMetadata_InputMetadata_Visualization_ToProto(mapCtx, in.Visualization)
	out.InputBaselines = direct.Slice_ToProto(mapCtx, in.InputBaselines, Value_ToProto)
	out.EncodedBaselines = direct.Slice_ToProto(mapCtx, in.EncodedBaselines, Value_ToProto)
	return out
}

func ExplanationMetadata_InputMetadata_FeatureValueDomain_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata_InputMetadata_FeatureValueDomain) *krm.ExplanationMetadata_InputMetadata_FeatureValueDomain {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationMetadata_InputMetadata_FeatureValueDomain{}
	out.MaxValue = direct.LazyPtr(in.GetMaxValue())
	out.MinValue = direct.LazyPtr(in.GetMinValue())
	out.OriginalMean = direct.LazyPtr(in.GetOriginalMean())
	out.OriginalStddev = direct.LazyPtr(in.GetOriginalStddev())
	return out
}

func ExplanationMetadata_InputMetadata_FeatureValueDomain_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationMetadata_InputMetadata_FeatureValueDomain) *pb.ExplanationMetadata_InputMetadata_FeatureValueDomain {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata_InputMetadata_FeatureValueDomain{}
	out.MaxValue = direct.ValueOf(in.MaxValue)
	out.MinValue = direct.ValueOf(in.MinValue)
	out.OriginalMean = direct.ValueOf(in.OriginalMean)
	out.OriginalStddev = direct.ValueOf(in.OriginalStddev)
	return out
}

func ExplanationMetadata_InputMetadata_Visualization_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata_InputMetadata_Visualization) *krm.ExplanationMetadata_InputMetadata_Visualization {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationMetadata_InputMetadata_Visualization{}
	out.ClipPercentLowerbound = direct.LazyPtr(in.GetClipPercentLowerbound())
	out.ClipPercentUpperbound = direct.LazyPtr(in.GetClipPercentUpperbound())
	out.ColorMap = direct.LazyPtr(in.GetColorMap().String())
	return out
}

func ExplanationMetadata_InputMetadata_Visualization_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationMetadata_InputMetadata_Visualization) *pb.ExplanationMetadata_InputMetadata_Visualization {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata_InputMetadata_Visualization{}
	out.ClipPercentLowerbound = direct.ValueOf(in.ClipPercentLowerbound)
	out.ClipPercentUpperbound = direct.ValueOf(in.ClipPercentUpperbound)
	out.ColorMap = pb.ExplanationMetadata_InputMetadata_Visualization_ColorMap(pb.ExplanationMetadata_InputMetadata_Visualization_ColorMap_value[direct.ValueOf(in.ColorMap)])
	return out
}

func ExplanationMetadata_OutputMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata_OutputMetadata) *krm.ExplanationMetadata_OutputMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationMetadata_OutputMetadata{}
	out.DisplayNameMappingKey = DisplayNameMappingKey_FromProto(mapCtx, in.GetDisplayNameMappingKey())
	out.IndexDisplayNameMapping = IndexDisplayNameMapping_FromProto(mapCtx, in.GetIndexDisplayNameMapping())
	out.OutputTensorName = direct.LazyPtr(in.GetOutputTensorName())
	return out
}

func ExplanationMetadata_OutputMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationMetadata_OutputMetadata) *pb.ExplanationMetadata_OutputMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata_OutputMetadata{}
	out.OutputTensorName = direct.ValueOf(in.OutputTensorName)
	if oneof := DisplayNameMappingKey_ToProto(mapCtx, in.DisplayNameMappingKey); oneof != nil {
		out.DisplayNameMapping = oneof
	}
	if oneof := IndexDisplayNameMapping_ToProto(mapCtx, in.IndexDisplayNameMapping); oneof != nil {
		out.DisplayNameMapping = oneof
	}
	return out
}

func DisplayNameMappingKey_FromProto(mapCtx *direct.MapContext, in string) *string {
	if in == "" {
		return nil
	}
	out := direct.LazyPtr(in)
	return out
}

func DisplayNameMappingKey_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExplanationMetadata_OutputMetadata_DisplayNameMappingKey {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata_OutputMetadata_DisplayNameMappingKey{
		DisplayNameMappingKey: direct.ValueOf(in),
	}
	return out
}

func IndexDisplayNameMapping_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := Value_FromProto(mapCtx, in)
	return out
}

func IndexDisplayNameMapping_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping{
		IndexDisplayNameMapping: Value_ToProto(mapCtx, in),
	}
	return out
}

func Presets_Query_ToProto(mapCtx *direct.MapContext, in *string) *pb.Presets_Query {
	if in == nil {
		return nil
	}
	query := pb.Presets_Query_value[direct.ValueOf(in)]
	out := direct.LazyPtr(pb.Presets_Query(query))
	return out
}

func SmoothGradConfig_NoiseSigma_ToProto(mapCtx *direct.MapContext, in *float32) *pb.SmoothGradConfig_NoiseSigma {
	if in == nil {
		return nil
	}
	out := &pb.SmoothGradConfig_NoiseSigma{}
	out.NoiseSigma = direct.ValueOf(in)
	return out
}
