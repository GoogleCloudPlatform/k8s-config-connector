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
	if direct.ValueOf(in.BoolValue) {
		out.Kind = &structpb.Value_BoolValue{
			BoolValue: direct.ValueOf(in.BoolValue),
		}
	}
	if in.ListValue != nil && len(in.ListValue.Values) > 0 {
		out.Kind = &structpb.Value_ListValue{
			ListValue: ListValue_ToProto(mapCtx, in.ListValue),
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
	switch in.GetKind() {
	case &structpb.Value_StringValue{}:
		out.StringValue = direct.LazyPtr(in.GetStringValue())
	case &structpb.Value_NumberValue{}:
		out.NumberValue = direct.LazyPtr(in.GetNumberValue())
	case &structpb.Value_NullValue{}:
		out.NullValue = direct.LazyPtr(in.GetNullValue().String())
	case &structpb.Value_BoolValue{}:
		out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	case &structpb.Value_ListValue{}:
		out.ListValue = ListValue_FromProto(mapCtx, in.GetListValue())
	case &structpb.Value_StructValue{}:
		out.StructValue = StructValue_FromProto(mapCtx, in.GetStructValue())
	}
	return out
}

func ListValue_FromProto(mapCtx *direct.MapContext, in *structpb.ListValue) *krm.ListValue {
	if in == nil {
		return nil
	}
	out := &krm.ListValue{}
	for _, value := range in.Values {
		out.Values = append(out.Values, direct.ValueOf(Value_FromProto(mapCtx, value)))
	}
	return out
}

func ListValue_ToProto(mapCtx *direct.MapContext, in *krm.ListValue) *structpb.ListValue {
	if in == nil {
		return nil
	}
	out := &structpb.ListValue{}
	for _, value := range in.Values {
		out.Values = append(out.Values, Value_ToProto(mapCtx, &value))
	}
	return out
}

func StructValue_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	if in == nil {
		return nil
	}
	var out map[string]string
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
	out.FeatureAttributionsSchemaURI = direct.LazyPtr(in.FeatureAttributionsSchemaUri)
	out.LatentSpaceSource = direct.LazyPtr(in.LatentSpaceSource)
	return out
}

func ExplanationMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationMetadata) *pb.ExplanationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationMetadata{}
	out.FeatureAttributionsSchemaUri = direct.ValueOf(in.FeatureAttributionsSchemaURI)
	out.LatentSpaceSource = direct.ValueOf(in.LatentSpaceSource)
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
