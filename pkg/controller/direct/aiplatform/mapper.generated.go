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
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BlurBaselineConfig_FromProto(mapCtx *direct.MapContext, in *pb.BlurBaselineConfig) *krm.BlurBaselineConfig {
	if in == nil {
		return nil
	}
	out := &krm.BlurBaselineConfig{}
	out.MaxBlurSigma = direct.LazyPtr(in.GetMaxBlurSigma())
	return out
}
func BlurBaselineConfig_ToProto(mapCtx *direct.MapContext, in *krm.BlurBaselineConfig) *pb.BlurBaselineConfig {
	if in == nil {
		return nil
	}
	out := &pb.BlurBaselineConfig{}
	out.MaxBlurSigma = direct.ValueOf(in.MaxBlurSigma)
	return out
}
func DeployedModelRef_FromProto(mapCtx *direct.MapContext, in *pb.DeployedModelRef) *krm.DeployedModelRef {
	if in == nil {
		return nil
	}
	out := &krm.DeployedModelRef{}
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	out.DeployedModelID = direct.LazyPtr(in.GetDeployedModelId())
	return out
}
func DeployedModelRef_ToProto(mapCtx *direct.MapContext, in *krm.DeployedModelRef) *pb.DeployedModelRef {
	if in == nil {
		return nil
	}
	out := &pb.DeployedModelRef{}
	out.Endpoint = direct.ValueOf(in.Endpoint)
	out.DeployedModelId = direct.ValueOf(in.DeployedModelID)
	return out
}
func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func EnvVar_FromProto(mapCtx *direct.MapContext, in *pb.EnvVar) *krm.EnvVar {
	if in == nil {
		return nil
	}
	out := &krm.EnvVar{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func EnvVar_ToProto(mapCtx *direct.MapContext, in *krm.EnvVar) *pb.EnvVar {
	if in == nil {
		return nil
	}
	out := &pb.EnvVar{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func Examples_FromProto(mapCtx *direct.MapContext, in *pb.Examples) *krm.Examples {
	if in == nil {
		return nil
	}
	out := &krm.Examples{}
	out.ExampleGcsSource = Examples_ExampleGcsSource_FromProto(mapCtx, in.GetExampleGcsSource())
	out.NearestNeighborSearchConfig = Value_FromProto(mapCtx, in.GetNearestNeighborSearchConfig())
	out.Presets = Presets_FromProto(mapCtx, in.GetPresets())
	out.NeighborCount = direct.LazyPtr(in.GetNeighborCount())
	return out
}
func Examples_ToProto(mapCtx *direct.MapContext, in *krm.Examples) *pb.Examples {
	if in == nil {
		return nil
	}
	out := &pb.Examples{}
	if oneof := Examples_ExampleGcsSource_ToProto(mapCtx, in.ExampleGcsSource); oneof != nil {
		out.Source = &pb.Examples_ExampleGcsSource_{ExampleGcsSource: oneof}
	}
	if oneof := Value_ToProto(mapCtx, in.NearestNeighborSearchConfig); oneof != nil {
		out.Config = &pb.Examples_NearestNeighborSearchConfig{NearestNeighborSearchConfig: oneof}
	}
	if oneof := Presets_ToProto(mapCtx, in.Presets); oneof != nil {
		out.Config = &pb.Examples_Presets{Presets: oneof}
	}
	out.NeighborCount = direct.ValueOf(in.NeighborCount)
	return out
}
func Examples_ExampleGcsSource_FromProto(mapCtx *direct.MapContext, in *pb.Examples_ExampleGcsSource) *krm.Examples_ExampleGcsSource {
	if in == nil {
		return nil
	}
	out := &krm.Examples_ExampleGcsSource{}
	out.DataFormat = direct.Enum_FromProto(mapCtx, in.GetDataFormat())
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	return out
}
func Examples_ExampleGcsSource_ToProto(mapCtx *direct.MapContext, in *krm.Examples_ExampleGcsSource) *pb.Examples_ExampleGcsSource {
	if in == nil {
		return nil
	}
	out := &pb.Examples_ExampleGcsSource{}
	out.DataFormat = direct.Enum_ToProto[pb.Examples_ExampleGcsSource_DataFormat](mapCtx, in.DataFormat)
	out.GcsSource = GcsSource_ToProto(mapCtx, in.GcsSource)
	return out
}
func ExplanationParameters_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationParameters) *krm.ExplanationParameters {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationParameters{}
	out.SampledShapleyAttribution = SampledShapleyAttribution_FromProto(mapCtx, in.GetSampledShapleyAttribution())
	out.IntegratedGradientsAttribution = IntegratedGradientsAttribution_FromProto(mapCtx, in.GetIntegratedGradientsAttribution())
	out.XraiAttribution = XraiAttribution_FromProto(mapCtx, in.GetXraiAttribution())
	out.Examples = Examples_FromProto(mapCtx, in.GetExamples())
	out.TopK = direct.LazyPtr(in.GetTopK())
	out.OutputIndices = ListValue_FromProto(mapCtx, in.GetOutputIndices())
	return out
}
func ExplanationParameters_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationParameters) *pb.ExplanationParameters {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationParameters{}
	if oneof := SampledShapleyAttribution_ToProto(mapCtx, in.SampledShapleyAttribution); oneof != nil {
		out.Method = &pb.ExplanationParameters_SampledShapleyAttribution{SampledShapleyAttribution: oneof}
	}
	if oneof := IntegratedGradientsAttribution_ToProto(mapCtx, in.IntegratedGradientsAttribution); oneof != nil {
		out.Method = &pb.ExplanationParameters_IntegratedGradientsAttribution{IntegratedGradientsAttribution: oneof}
	}
	if oneof := XraiAttribution_ToProto(mapCtx, in.XraiAttribution); oneof != nil {
		out.Method = &pb.ExplanationParameters_XraiAttribution{XraiAttribution: oneof}
	}
	if oneof := Examples_ToProto(mapCtx, in.Examples); oneof != nil {
		out.Method = &pb.ExplanationParameters_Examples{Examples: oneof}
	}
	out.TopK = direct.ValueOf(in.TopK)
	out.OutputIndices = ListValue_ToProto(mapCtx, in.OutputIndices)
	return out
}
func ExplanationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationSpec) *krm.ExplanationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationSpec{}
	out.Parameters = ExplanationParameters_FromProto(mapCtx, in.GetParameters())
	out.Metadata = ExplanationMetadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func ExplanationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationSpec) *pb.ExplanationSpec {
	if in == nil {
		return nil
	}
	out := &pb.ExplanationSpec{}
	out.Parameters = ExplanationParameters_ToProto(mapCtx, in.Parameters)
	out.Metadata = ExplanationMetadata_ToProto(mapCtx, in.Metadata)
	return out
}
func FeatureNoiseSigma_FromProto(mapCtx *direct.MapContext, in *pb.FeatureNoiseSigma) *krm.FeatureNoiseSigma {
	if in == nil {
		return nil
	}
	out := &krm.FeatureNoiseSigma{}
	out.NoiseSigma = direct.Slice_FromProto(mapCtx, in.NoiseSigma, FeatureNoiseSigma_NoiseSigmaForFeature_FromProto)
	return out
}
func FeatureNoiseSigma_ToProto(mapCtx *direct.MapContext, in *krm.FeatureNoiseSigma) *pb.FeatureNoiseSigma {
	if in == nil {
		return nil
	}
	out := &pb.FeatureNoiseSigma{}
	out.NoiseSigma = direct.Slice_ToProto(mapCtx, in.NoiseSigma, FeatureNoiseSigma_NoiseSigmaForFeature_ToProto)
	return out
}
func FeatureNoiseSigma_NoiseSigmaForFeature_FromProto(mapCtx *direct.MapContext, in *pb.FeatureNoiseSigma_NoiseSigmaForFeature) *krm.FeatureNoiseSigma_NoiseSigmaForFeature {
	if in == nil {
		return nil
	}
	out := &krm.FeatureNoiseSigma_NoiseSigmaForFeature{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Sigma = direct.LazyPtr(in.GetSigma())
	return out
}
func FeatureNoiseSigma_NoiseSigmaForFeature_ToProto(mapCtx *direct.MapContext, in *krm.FeatureNoiseSigma_NoiseSigmaForFeature) *pb.FeatureNoiseSigma_NoiseSigmaForFeature {
	if in == nil {
		return nil
	}
	out := &pb.FeatureNoiseSigma_NoiseSigmaForFeature{}
	out.Name = direct.ValueOf(in.Name)
	out.Sigma = direct.ValueOf(in.Sigma)
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
func GenieSource_FromProto(mapCtx *direct.MapContext, in *pb.GenieSource) *krm.GenieSource {
	if in == nil {
		return nil
	}
	out := &krm.GenieSource{}
	out.BaseModelURI = direct.LazyPtr(in.GetBaseModelUri())
	return out
}
func GenieSource_ToProto(mapCtx *direct.MapContext, in *krm.GenieSource) *pb.GenieSource {
	if in == nil {
		return nil
	}
	out := &pb.GenieSource{}
	out.BaseModelUri = direct.ValueOf(in.BaseModelURI)
	return out
}
func IntegratedGradientsAttribution_FromProto(mapCtx *direct.MapContext, in *pb.IntegratedGradientsAttribution) *krm.IntegratedGradientsAttribution {
	if in == nil {
		return nil
	}
	out := &krm.IntegratedGradientsAttribution{}
	out.StepCount = direct.LazyPtr(in.GetStepCount())
	out.SmoothGradConfig = SmoothGradConfig_FromProto(mapCtx, in.GetSmoothGradConfig())
	out.BlurBaselineConfig = BlurBaselineConfig_FromProto(mapCtx, in.GetBlurBaselineConfig())
	return out
}
func IntegratedGradientsAttribution_ToProto(mapCtx *direct.MapContext, in *krm.IntegratedGradientsAttribution) *pb.IntegratedGradientsAttribution {
	if in == nil {
		return nil
	}
	out := &pb.IntegratedGradientsAttribution{}
	out.StepCount = direct.ValueOf(in.StepCount)
	out.SmoothGradConfig = SmoothGradConfig_ToProto(mapCtx, in.SmoothGradConfig)
	out.BlurBaselineConfig = BlurBaselineConfig_ToProto(mapCtx, in.BlurBaselineConfig)
	return out
}
func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: VersionID
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
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: VersionID
	out.VersionAliases = in.VersionAliases
	// MISSING: VersionCreateTime
	// MISSING: VersionUpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.VersionDescription = direct.ValueOf(in.VersionDescription)
	out.PredictSchemata = PredictSchemata_ToProto(mapCtx, in.PredictSchemata)
	out.MetadataSchemaUri = direct.ValueOf(in.MetadataSchemaURI)
	out.Metadata = Value_ToProto(mapCtx, in.Metadata)
	// MISSING: SupportedExportFormats
	// MISSING: TrainingPipeline
	out.PipelineJob = direct.ValueOf(in.PipelineJob)
	out.ContainerSpec = ModelContainerSpec_ToProto(mapCtx, in.ContainerSpec)
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	// MISSING: SupportedDeploymentResourcesTypes
	// MISSING: SupportedInputStorageFormats
	// MISSING: SupportedOutputStorageFormats
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeployedModels
	out.ExplanationSpec = ExplanationSpec_ToProto(mapCtx, in.ExplanationSpec)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.DataStats = Model_DataStats_ToProto(mapCtx, in.DataStats)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: ModelSourceInfo
	// MISSING: OriginalModelInfo
	// MISSING: MetadataArtifact
	out.BaseModelSource = Model_BaseModelSource_ToProto(mapCtx, in.BaseModelSource)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ModelContainerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelContainerSpec) *krm.ModelContainerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ModelContainerSpec{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, EnvVar_FromProto)
	out.Ports = direct.Slice_FromProto(mapCtx, in.Ports, Port_FromProto)
	out.PredictRoute = direct.LazyPtr(in.GetPredictRoute())
	out.HealthRoute = direct.LazyPtr(in.GetHealthRoute())
	out.GrpcPorts = direct.Slice_FromProto(mapCtx, in.GrpcPorts, Port_FromProto)
	out.DeploymentTimeout = direct.StringDuration_FromProto(mapCtx, in.GetDeploymentTimeout())
	out.SharedMemorySizeMb = direct.LazyPtr(in.GetSharedMemorySizeMb())
	out.StartupProbe = Probe_FromProto(mapCtx, in.GetStartupProbe())
	out.HealthProbe = Probe_FromProto(mapCtx, in.GetHealthProbe())
	return out
}
func ModelContainerSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelContainerSpec) *pb.ModelContainerSpec {
	if in == nil {
		return nil
	}
	out := &pb.ModelContainerSpec{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_ToProto)
	out.Ports = direct.Slice_ToProto(mapCtx, in.Ports, Port_ToProto)
	out.PredictRoute = direct.ValueOf(in.PredictRoute)
	out.HealthRoute = direct.ValueOf(in.HealthRoute)
	out.GrpcPorts = direct.Slice_ToProto(mapCtx, in.GrpcPorts, Port_ToProto)
	out.DeploymentTimeout = direct.StringDuration_ToProto(mapCtx, in.DeploymentTimeout)
	out.SharedMemorySizeMb = direct.ValueOf(in.SharedMemorySizeMb)
	out.StartupProbe = Probe_ToProto(mapCtx, in.StartupProbe)
	out.HealthProbe = Probe_ToProto(mapCtx, in.HealthProbe)
	return out
}
func ModelGardenSource_FromProto(mapCtx *direct.MapContext, in *pb.ModelGardenSource) *krm.ModelGardenSource {
	if in == nil {
		return nil
	}
	out := &krm.ModelGardenSource{}
	out.PublicModelName = direct.LazyPtr(in.GetPublicModelName())
	return out
}
func ModelGardenSource_ToProto(mapCtx *direct.MapContext, in *krm.ModelGardenSource) *pb.ModelGardenSource {
	if in == nil {
		return nil
	}
	out := &pb.ModelGardenSource{}
	out.PublicModelName = direct.ValueOf(in.PublicModelName)
	return out
}
func ModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.ModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ModelObservedState{}
	// MISSING: Name
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	// MISSING: VersionAliases
	out.VersionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionCreateTime())
	out.VersionUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionUpdateTime())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionDescription
	// MISSING: PredictSchemata
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	out.SupportedExportFormats = direct.Slice_FromProto(mapCtx, in.SupportedExportFormats, Model_ExportFormat_FromProto)
	out.TrainingPipeline = direct.LazyPtr(in.GetTrainingPipeline())
	// MISSING: PipelineJob
	// MISSING: ContainerSpec
	// MISSING: ArtifactURI
	out.SupportedDeploymentResourcesTypes = direct.EnumSlice_FromProto(mapCtx, in.SupportedDeploymentResourcesTypes)
	out.SupportedInputStorageFormats = in.SupportedInputStorageFormats
	out.SupportedOutputStorageFormats = in.SupportedOutputStorageFormats
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeployedModels = direct.Slice_FromProto(mapCtx, in.DeployedModels, DeployedModelRef_FromProto)
	// MISSING: ExplanationSpec
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: DataStats
	// MISSING: EncryptionSpec
	out.ModelSourceInfo = ModelSourceInfo_FromProto(mapCtx, in.GetModelSourceInfo())
	out.OriginalModelInfo = Model_OriginalModelInfo_FromProto(mapCtx, in.GetOriginalModelInfo())
	out.MetadataArtifact = direct.LazyPtr(in.GetMetadataArtifact())
	// MISSING: BaseModelSource
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func ModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	out.VersionId = direct.ValueOf(in.VersionID)
	// MISSING: VersionAliases
	out.VersionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.VersionCreateTime)
	out.VersionUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.VersionUpdateTime)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionDescription
	// MISSING: PredictSchemata
	// MISSING: MetadataSchemaURI
	// MISSING: Metadata
	out.SupportedExportFormats = direct.Slice_ToProto(mapCtx, in.SupportedExportFormats, Model_ExportFormat_ToProto)
	out.TrainingPipeline = direct.ValueOf(in.TrainingPipeline)
	// MISSING: PipelineJob
	// MISSING: ContainerSpec
	// MISSING: ArtifactURI
	out.SupportedDeploymentResourcesTypes = direct.EnumSlice_ToProto[pb.Model_DeploymentResourcesType](mapCtx, in.SupportedDeploymentResourcesTypes)
	out.SupportedInputStorageFormats = in.SupportedInputStorageFormats
	out.SupportedOutputStorageFormats = in.SupportedOutputStorageFormats
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeployedModels = direct.Slice_ToProto(mapCtx, in.DeployedModels, DeployedModelRef_ToProto)
	// MISSING: ExplanationSpec
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: DataStats
	// MISSING: EncryptionSpec
	out.ModelSourceInfo = ModelSourceInfo_ToProto(mapCtx, in.ModelSourceInfo)
	out.OriginalModelInfo = Model_OriginalModelInfo_ToProto(mapCtx, in.OriginalModelInfo)
	out.MetadataArtifact = direct.ValueOf(in.MetadataArtifact)
	// MISSING: BaseModelSource
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func ModelSourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.ModelSourceInfo) *krm.ModelSourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.ModelSourceInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.Copy = direct.LazyPtr(in.GetCopy())
	return out
}
func ModelSourceInfo_ToProto(mapCtx *direct.MapContext, in *krm.ModelSourceInfo) *pb.ModelSourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.ModelSourceInfo{}
	out.SourceType = direct.Enum_ToProto[pb.ModelSourceInfo_ModelSourceType](mapCtx, in.SourceType)
	out.Copy = direct.ValueOf(in.Copy)
	return out
}
func Model_BaseModelSource_FromProto(mapCtx *direct.MapContext, in *pb.Model_BaseModelSource) *krm.Model_BaseModelSource {
	if in == nil {
		return nil
	}
	out := &krm.Model_BaseModelSource{}
	out.ModelGardenSource = ModelGardenSource_FromProto(mapCtx, in.GetModelGardenSource())
	out.GenieSource = GenieSource_FromProto(mapCtx, in.GetGenieSource())
	return out
}
func Model_BaseModelSource_ToProto(mapCtx *direct.MapContext, in *krm.Model_BaseModelSource) *pb.Model_BaseModelSource {
	if in == nil {
		return nil
	}
	out := &pb.Model_BaseModelSource{}
	if oneof := ModelGardenSource_ToProto(mapCtx, in.ModelGardenSource); oneof != nil {
		out.Source = &pb.Model_BaseModelSource_ModelGardenSource{ModelGardenSource: oneof}
	}
	if oneof := GenieSource_ToProto(mapCtx, in.GenieSource); oneof != nil {
		out.Source = &pb.Model_BaseModelSource_GenieSource{GenieSource: oneof}
	}
	return out
}
func Model_DataStats_FromProto(mapCtx *direct.MapContext, in *pb.Model_DataStats) *krm.Model_DataStats {
	if in == nil {
		return nil
	}
	out := &krm.Model_DataStats{}
	out.TrainingDataItemsCount = direct.LazyPtr(in.GetTrainingDataItemsCount())
	out.ValidationDataItemsCount = direct.LazyPtr(in.GetValidationDataItemsCount())
	out.TestDataItemsCount = direct.LazyPtr(in.GetTestDataItemsCount())
	out.TrainingAnnotationsCount = direct.LazyPtr(in.GetTrainingAnnotationsCount())
	out.ValidationAnnotationsCount = direct.LazyPtr(in.GetValidationAnnotationsCount())
	out.TestAnnotationsCount = direct.LazyPtr(in.GetTestAnnotationsCount())
	return out
}
func Model_DataStats_ToProto(mapCtx *direct.MapContext, in *krm.Model_DataStats) *pb.Model_DataStats {
	if in == nil {
		return nil
	}
	out := &pb.Model_DataStats{}
	out.TrainingDataItemsCount = direct.ValueOf(in.TrainingDataItemsCount)
	out.ValidationDataItemsCount = direct.ValueOf(in.ValidationDataItemsCount)
	out.TestDataItemsCount = direct.ValueOf(in.TestDataItemsCount)
	out.TrainingAnnotationsCount = direct.ValueOf(in.TrainingAnnotationsCount)
	out.ValidationAnnotationsCount = direct.ValueOf(in.ValidationAnnotationsCount)
	out.TestAnnotationsCount = direct.ValueOf(in.TestAnnotationsCount)
	return out
}
func Model_ExportFormat_FromProto(mapCtx *direct.MapContext, in *pb.Model_ExportFormat) *krm.Model_ExportFormat {
	if in == nil {
		return nil
	}
	out := &krm.Model_ExportFormat{}
	// MISSING: ID
	// MISSING: ExportableContents
	return out
}
func Model_ExportFormat_ToProto(mapCtx *direct.MapContext, in *krm.Model_ExportFormat) *pb.Model_ExportFormat {
	if in == nil {
		return nil
	}
	out := &pb.Model_ExportFormat{}
	// MISSING: ID
	// MISSING: ExportableContents
	return out
}
func Model_ExportFormatObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model_ExportFormat) *krm.Model_ExportFormatObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Model_ExportFormatObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.ExportableContents = direct.EnumSlice_FromProto(mapCtx, in.ExportableContents)
	return out
}
func Model_ExportFormatObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Model_ExportFormatObservedState) *pb.Model_ExportFormat {
	if in == nil {
		return nil
	}
	out := &pb.Model_ExportFormat{}
	out.Id = direct.ValueOf(in.ID)
	out.ExportableContents = direct.EnumSlice_ToProto[pb.Model_ExportFormat_ExportableContent](mapCtx, in.ExportableContents)
	return out
}
func Model_OriginalModelInfo_FromProto(mapCtx *direct.MapContext, in *pb.Model_OriginalModelInfo) *krm.Model_OriginalModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.Model_OriginalModelInfo{}
	// MISSING: Model
	return out
}
func Model_OriginalModelInfo_ToProto(mapCtx *direct.MapContext, in *krm.Model_OriginalModelInfo) *pb.Model_OriginalModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.Model_OriginalModelInfo{}
	// MISSING: Model
	return out
}
func Model_OriginalModelInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model_OriginalModelInfo) *krm.Model_OriginalModelInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Model_OriginalModelInfoObservedState{}
	out.Model = direct.LazyPtr(in.GetModel())
	return out
}
func Model_OriginalModelInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Model_OriginalModelInfoObservedState) *pb.Model_OriginalModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.Model_OriginalModelInfo{}
	out.Model = direct.ValueOf(in.Model)
	return out
}
func Port_FromProto(mapCtx *direct.MapContext, in *pb.Port) *krm.Port {
	if in == nil {
		return nil
	}
	out := &krm.Port{}
	out.ContainerPort = direct.LazyPtr(in.GetContainerPort())
	return out
}
func Port_ToProto(mapCtx *direct.MapContext, in *krm.Port) *pb.Port {
	if in == nil {
		return nil
	}
	out := &pb.Port{}
	out.ContainerPort = direct.ValueOf(in.ContainerPort)
	return out
}
func PredictSchemata_FromProto(mapCtx *direct.MapContext, in *pb.PredictSchemata) *krm.PredictSchemata {
	if in == nil {
		return nil
	}
	out := &krm.PredictSchemata{}
	out.InstanceSchemaURI = direct.LazyPtr(in.GetInstanceSchemaUri())
	out.ParametersSchemaURI = direct.LazyPtr(in.GetParametersSchemaUri())
	out.PredictionSchemaURI = direct.LazyPtr(in.GetPredictionSchemaUri())
	return out
}
func PredictSchemata_ToProto(mapCtx *direct.MapContext, in *krm.PredictSchemata) *pb.PredictSchemata {
	if in == nil {
		return nil
	}
	out := &pb.PredictSchemata{}
	out.InstanceSchemaUri = direct.ValueOf(in.InstanceSchemaURI)
	out.ParametersSchemaUri = direct.ValueOf(in.ParametersSchemaURI)
	out.PredictionSchemaUri = direct.ValueOf(in.PredictionSchemaURI)
	return out
}
func Presets_FromProto(mapCtx *direct.MapContext, in *pb.Presets) *krm.Presets {
	if in == nil {
		return nil
	}
	out := &krm.Presets{}
	out.Query = direct.Enum_FromProto(mapCtx, in.GetQuery())
	out.Modality = direct.Enum_FromProto(mapCtx, in.GetModality())
	return out
}
func Presets_ToProto(mapCtx *direct.MapContext, in *krm.Presets) *pb.Presets {
	if in == nil {
		return nil
	}
	out := &pb.Presets{}
	if oneof := Presets_Query_ToProto(mapCtx, in.Query); oneof != nil {
		out.Query = oneof
	}
	out.Modality = direct.Enum_ToProto[pb.Presets_Modality](mapCtx, in.Modality)
	return out
}
func Probe_FromProto(mapCtx *direct.MapContext, in *pb.Probe) *krm.Probe {
	if in == nil {
		return nil
	}
	out := &krm.Probe{}
	out.Exec = Probe_ExecAction_FromProto(mapCtx, in.GetExec())
	out.PeriodSeconds = direct.LazyPtr(in.GetPeriodSeconds())
	out.TimeoutSeconds = direct.LazyPtr(in.GetTimeoutSeconds())
	return out
}
func Probe_ToProto(mapCtx *direct.MapContext, in *krm.Probe) *pb.Probe {
	if in == nil {
		return nil
	}
	out := &pb.Probe{}
	if oneof := Probe_ExecAction_ToProto(mapCtx, in.Exec); oneof != nil {
		out.ProbeType = &pb.Probe_Exec{Exec: oneof}
	}
	out.PeriodSeconds = direct.ValueOf(in.PeriodSeconds)
	out.TimeoutSeconds = direct.ValueOf(in.TimeoutSeconds)
	return out
}
func Probe_ExecAction_FromProto(mapCtx *direct.MapContext, in *pb.Probe_ExecAction) *krm.Probe_ExecAction {
	if in == nil {
		return nil
	}
	out := &krm.Probe_ExecAction{}
	out.Command = in.Command
	return out
}
func Probe_ExecAction_ToProto(mapCtx *direct.MapContext, in *krm.Probe_ExecAction) *pb.Probe_ExecAction {
	if in == nil {
		return nil
	}
	out := &pb.Probe_ExecAction{}
	out.Command = in.Command
	return out
}
func SampledShapleyAttribution_FromProto(mapCtx *direct.MapContext, in *pb.SampledShapleyAttribution) *krm.SampledShapleyAttribution {
	if in == nil {
		return nil
	}
	out := &krm.SampledShapleyAttribution{}
	out.PathCount = direct.LazyPtr(in.GetPathCount())
	return out
}
func SampledShapleyAttribution_ToProto(mapCtx *direct.MapContext, in *krm.SampledShapleyAttribution) *pb.SampledShapleyAttribution {
	if in == nil {
		return nil
	}
	out := &pb.SampledShapleyAttribution{}
	out.PathCount = direct.ValueOf(in.PathCount)
	return out
}
func SmoothGradConfig_FromProto(mapCtx *direct.MapContext, in *pb.SmoothGradConfig) *krm.SmoothGradConfig {
	if in == nil {
		return nil
	}
	out := &krm.SmoothGradConfig{}
	out.NoiseSigma = direct.LazyPtr(in.GetNoiseSigma())
	out.FeatureNoiseSigma = FeatureNoiseSigma_FromProto(mapCtx, in.GetFeatureNoiseSigma())
	out.NoisySampleCount = direct.LazyPtr(in.GetNoisySampleCount())
	return out
}
func SmoothGradConfig_ToProto(mapCtx *direct.MapContext, in *krm.SmoothGradConfig) *pb.SmoothGradConfig {
	if in == nil {
		return nil
	}
	out := &pb.SmoothGradConfig{}
	if oneof := SmoothGradConfig_NoiseSigma_ToProto(mapCtx, in.NoiseSigma); oneof != nil {
		out.GradientNoiseSigma = oneof
	}
	if oneof := FeatureNoiseSigma_ToProto(mapCtx, in.FeatureNoiseSigma); oneof != nil {
		out.GradientNoiseSigma = &pb.SmoothGradConfig_FeatureNoiseSigma{FeatureNoiseSigma: oneof}
	}
	out.NoisySampleCount = direct.ValueOf(in.NoisySampleCount)
	return out
}
func XraiAttribution_FromProto(mapCtx *direct.MapContext, in *pb.XraiAttribution) *krm.XraiAttribution {
	if in == nil {
		return nil
	}
	out := &krm.XraiAttribution{}
	out.StepCount = direct.LazyPtr(in.GetStepCount())
	out.SmoothGradConfig = SmoothGradConfig_FromProto(mapCtx, in.GetSmoothGradConfig())
	out.BlurBaselineConfig = BlurBaselineConfig_FromProto(mapCtx, in.GetBlurBaselineConfig())
	return out
}
func XraiAttribution_ToProto(mapCtx *direct.MapContext, in *krm.XraiAttribution) *pb.XraiAttribution {
	if in == nil {
		return nil
	}
	out := &pb.XraiAttribution{}
	out.StepCount = direct.ValueOf(in.StepCount)
	out.SmoothGradConfig = SmoothGradConfig_ToProto(mapCtx, in.SmoothGradConfig)
	out.BlurBaselineConfig = BlurBaselineConfig_ToProto(mapCtx, in.BlurBaselineConfig)
	return out
}
