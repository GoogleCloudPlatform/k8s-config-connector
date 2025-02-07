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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AiplatformAnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationObservedState{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpec{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpec) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AiplatformAnnotationSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpecObservedState) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AiplatformAnnotationSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpecSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpecSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformArtifactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.AiplatformArtifactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformArtifactObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformArtifactObservedState) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactSpec_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.AiplatformArtifactSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformArtifactSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformArtifactSpec) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformBatchPredictionJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.AiplatformBatchPredictionJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformBatchPredictionJobObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformBatchPredictionJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformBatchPredictionJobObservedState) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformBatchPredictionJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.BatchPredictionJob) *krm.AiplatformBatchPredictionJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformBatchPredictionJobSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformBatchPredictionJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformBatchPredictionJobSpec) *pb.BatchPredictionJob {
	if in == nil {
		return nil
	}
	out := &pb.BatchPredictionJob{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: ModelVersionID
	// MISSING: UnmanagedContainerModel
	// MISSING: InputConfig
	// MISSING: InstanceConfig
	// MISSING: ModelParameters
	// MISSING: OutputConfig
	// MISSING: DedicatedResources
	// MISSING: ServiceAccount
	// MISSING: ManualBatchTuningParameters
	// MISSING: GenerateExplanation
	// MISSING: ExplanationSpec
	// MISSING: OutputInfo
	// MISSING: State
	// MISSING: Error
	// MISSING: PartialFailures
	// MISSING: ResourcesConsumed
	// MISSING: CompletionStats
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EncryptionSpec
	// MISSING: DisableContainerLogging
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func AiplatformCachedContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiplatformCachedContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformCachedContentObservedState{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformCachedContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformCachedContentObservedState) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformCachedContentSpec_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiplatformCachedContentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformCachedContentSpec{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiplatformCachedContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformCachedContentSpec) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func Blob_FromProto(mapCtx *direct.MapContext, in *pb.Blob) *krm.Blob {
	if in == nil {
		return nil
	}
	out := &krm.Blob{}
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.Data = in.GetData()
	return out
}
func Blob_ToProto(mapCtx *direct.MapContext, in *krm.Blob) *pb.Blob {
	if in == nil {
		return nil
	}
	out := &pb.Blob{}
	out.MimeType = direct.ValueOf(in.MimeType)
	out.Data = in.Data
	return out
}
func CachedContent_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.CachedContent {
	if in == nil {
		return nil
	}
	out := &krm.CachedContent{}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Model = direct.LazyPtr(in.GetModel())
	out.SystemInstruction = Content_FromProto(mapCtx, in.GetSystemInstruction())
	out.Contents = direct.Slice_FromProto(mapCtx, in.Contents, Content_FromProto)
	out.Tools = direct.Slice_FromProto(mapCtx, in.Tools, Tool_FromProto)
	out.ToolConfig = ToolConfig_FromProto(mapCtx, in.GetToolConfig())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func CachedContent_ToProto(mapCtx *direct.MapContext, in *krm.CachedContent) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.Expiration = &pb.CachedContent_ExpireTime{ExpireTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.Ttl); oneof != nil {
		out.Expiration = &pb.CachedContent_Ttl{Ttl: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Model = direct.ValueOf(in.Model)
	out.SystemInstruction = Content_ToProto(mapCtx, in.SystemInstruction)
	out.Contents = direct.Slice_ToProto(mapCtx, in.Contents, Content_ToProto)
	out.Tools = direct.Slice_ToProto(mapCtx, in.Tools, Tool_ToProto)
	out.ToolConfig = ToolConfig_ToProto(mapCtx, in.ToolConfig)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func CachedContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.CachedContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CachedContentObservedState{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.UsageMetadata = CachedContent_UsageMetadata_FromProto(mapCtx, in.GetUsageMetadata())
	return out
}
func CachedContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CachedContentObservedState) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.UsageMetadata = CachedContent_UsageMetadata_ToProto(mapCtx, in.UsageMetadata)
	return out
}
func Content_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krm.Content {
	if in == nil {
		return nil
	}
	out := &krm.Content{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, Part_FromProto)
	return out
}
func Content_ToProto(mapCtx *direct.MapContext, in *krm.Content) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	out.Role = direct.ValueOf(in.Role)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, Part_ToProto)
	return out
}
func DynamicRetrievalConfig_FromProto(mapCtx *direct.MapContext, in *pb.DynamicRetrievalConfig) *krm.DynamicRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &krm.DynamicRetrievalConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.DynamicThreshold = in.DynamicThreshold
	return out
}
func DynamicRetrievalConfig_ToProto(mapCtx *direct.MapContext, in *krm.DynamicRetrievalConfig) *pb.DynamicRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &pb.DynamicRetrievalConfig{}
	out.Mode = direct.Enum_ToProto[pb.DynamicRetrievalConfig_Mode](mapCtx, in.Mode)
	out.DynamicThreshold = in.DynamicThreshold
	return out
}
func FileData_FromProto(mapCtx *direct.MapContext, in *pb.FileData) *krm.FileData {
	if in == nil {
		return nil
	}
	out := &krm.FileData{}
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.FileURI = direct.LazyPtr(in.GetFileUri())
	return out
}
func FileData_ToProto(mapCtx *direct.MapContext, in *krm.FileData) *pb.FileData {
	if in == nil {
		return nil
	}
	out := &pb.FileData{}
	out.MimeType = direct.ValueOf(in.MimeType)
	out.FileUri = direct.ValueOf(in.FileURI)
	return out
}
func FunctionCall_FromProto(mapCtx *direct.MapContext, in *pb.FunctionCall) *krm.FunctionCall {
	if in == nil {
		return nil
	}
	out := &krm.FunctionCall{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Args = Args_FromProto(mapCtx, in.GetArgs())
	return out
}
func FunctionCall_ToProto(mapCtx *direct.MapContext, in *krm.FunctionCall) *pb.FunctionCall {
	if in == nil {
		return nil
	}
	out := &pb.FunctionCall{}
	out.Name = direct.ValueOf(in.Name)
	out.Args = Args_ToProto(mapCtx, in.Args)
	return out
}
func FunctionCallingConfig_FromProto(mapCtx *direct.MapContext, in *pb.FunctionCallingConfig) *krm.FunctionCallingConfig {
	if in == nil {
		return nil
	}
	out := &krm.FunctionCallingConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.AllowedFunctionNames = in.AllowedFunctionNames
	return out
}
func FunctionCallingConfig_ToProto(mapCtx *direct.MapContext, in *krm.FunctionCallingConfig) *pb.FunctionCallingConfig {
	if in == nil {
		return nil
	}
	out := &pb.FunctionCallingConfig{}
	out.Mode = direct.Enum_ToProto[pb.FunctionCallingConfig_Mode](mapCtx, in.Mode)
	out.AllowedFunctionNames = in.AllowedFunctionNames
	return out
}
func FunctionDeclaration_FromProto(mapCtx *direct.MapContext, in *pb.FunctionDeclaration) *krm.FunctionDeclaration {
	if in == nil {
		return nil
	}
	out := &krm.FunctionDeclaration{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Parameters = Schema_FromProto(mapCtx, in.GetParameters())
	out.Response = Schema_FromProto(mapCtx, in.GetResponse())
	return out
}
func FunctionDeclaration_ToProto(mapCtx *direct.MapContext, in *krm.FunctionDeclaration) *pb.FunctionDeclaration {
	if in == nil {
		return nil
	}
	out := &pb.FunctionDeclaration{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Parameters = Schema_ToProto(mapCtx, in.Parameters)
	out.Response = Schema_ToProto(mapCtx, in.Response)
	return out
}
func GoogleSearchRetrieval_FromProto(mapCtx *direct.MapContext, in *pb.GoogleSearchRetrieval) *krm.GoogleSearchRetrieval {
	if in == nil {
		return nil
	}
	out := &krm.GoogleSearchRetrieval{}
	out.DynamicRetrievalConfig = DynamicRetrievalConfig_FromProto(mapCtx, in.GetDynamicRetrievalConfig())
	return out
}
func GoogleSearchRetrieval_ToProto(mapCtx *direct.MapContext, in *krm.GoogleSearchRetrieval) *pb.GoogleSearchRetrieval {
	if in == nil {
		return nil
	}
	out := &pb.GoogleSearchRetrieval{}
	out.DynamicRetrievalConfig = DynamicRetrievalConfig_ToProto(mapCtx, in.DynamicRetrievalConfig)
	return out
}
func Part_FromProto(mapCtx *direct.MapContext, in *pb.Part) *krm.Part {
	if in == nil {
		return nil
	}
	out := &krm.Part{}
	out.Text = direct.LazyPtr(in.GetText())
	out.InlineData = Blob_FromProto(mapCtx, in.GetInlineData())
	out.FileData = FileData_FromProto(mapCtx, in.GetFileData())
	out.FunctionCall = FunctionCall_FromProto(mapCtx, in.GetFunctionCall())
	out.FunctionResponse = FunctionResponse_FromProto(mapCtx, in.GetFunctionResponse())
	out.VideoMetadata = VideoMetadata_FromProto(mapCtx, in.GetVideoMetadata())
	return out
}
func Part_ToProto(mapCtx *direct.MapContext, in *krm.Part) *pb.Part {
	if in == nil {
		return nil
	}
	out := &pb.Part{}
	if oneof := Part_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Data = oneof
	}
	if oneof := Blob_ToProto(mapCtx, in.InlineData); oneof != nil {
		out.Data = &pb.Part_InlineData{InlineData: oneof}
	}
	if oneof := FileData_ToProto(mapCtx, in.FileData); oneof != nil {
		out.Data = &pb.Part_FileData{FileData: oneof}
	}
	if oneof := FunctionCall_ToProto(mapCtx, in.FunctionCall); oneof != nil {
		out.Data = &pb.Part_FunctionCall{FunctionCall: oneof}
	}
	if oneof := FunctionResponse_ToProto(mapCtx, in.FunctionResponse); oneof != nil {
		out.Data = &pb.Part_FunctionResponse{FunctionResponse: oneof}
	}
	if oneof := VideoMetadata_ToProto(mapCtx, in.VideoMetadata); oneof != nil {
		out.Metadata = &pb.Part_VideoMetadata{VideoMetadata: oneof}
	}
	return out
}
func RagRetrievalConfig_FromProto(mapCtx *direct.MapContext, in *pb.RagRetrievalConfig) *krm.RagRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &krm.RagRetrievalConfig{}
	out.TopK = direct.LazyPtr(in.GetTopK())
	out.Filter = RagRetrievalConfig_Filter_FromProto(mapCtx, in.GetFilter())
	return out
}
func RagRetrievalConfig_ToProto(mapCtx *direct.MapContext, in *krm.RagRetrievalConfig) *pb.RagRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &pb.RagRetrievalConfig{}
	out.TopK = direct.ValueOf(in.TopK)
	out.Filter = RagRetrievalConfig_Filter_ToProto(mapCtx, in.Filter)
	return out
}
func RagRetrievalConfig_Filter_FromProto(mapCtx *direct.MapContext, in *pb.RagRetrievalConfig_Filter) *krm.RagRetrievalConfig_Filter {
	if in == nil {
		return nil
	}
	out := &krm.RagRetrievalConfig_Filter{}
	out.VectorDistanceThreshold = direct.LazyPtr(in.GetVectorDistanceThreshold())
	out.VectorSimilarityThreshold = direct.LazyPtr(in.GetVectorSimilarityThreshold())
	out.MetadataFilter = direct.LazyPtr(in.GetMetadataFilter())
	return out
}
func RagRetrievalConfig_Filter_ToProto(mapCtx *direct.MapContext, in *krm.RagRetrievalConfig_Filter) *pb.RagRetrievalConfig_Filter {
	if in == nil {
		return nil
	}
	out := &pb.RagRetrievalConfig_Filter{}
	if oneof := RagRetrievalConfig_Filter_VectorDistanceThreshold_ToProto(mapCtx, in.VectorDistanceThreshold); oneof != nil {
		out.VectorDbThreshold = oneof
	}
	if oneof := RagRetrievalConfig_Filter_VectorSimilarityThreshold_ToProto(mapCtx, in.VectorSimilarityThreshold); oneof != nil {
		out.VectorDbThreshold = oneof
	}
	out.MetadataFilter = direct.ValueOf(in.MetadataFilter)
	return out
}
func Retrieval_FromProto(mapCtx *direct.MapContext, in *pb.Retrieval) *krm.Retrieval {
	if in == nil {
		return nil
	}
	out := &krm.Retrieval{}
	out.VertexAiSearch = VertexAISearch_FromProto(mapCtx, in.GetVertexAiSearch())
	out.VertexRagStore = VertexRagStore_FromProto(mapCtx, in.GetVertexRagStore())
	out.DisableAttribution = direct.LazyPtr(in.GetDisableAttribution())
	return out
}
func Retrieval_ToProto(mapCtx *direct.MapContext, in *krm.Retrieval) *pb.Retrieval {
	if in == nil {
		return nil
	}
	out := &pb.Retrieval{}
	if oneof := VertexAISearch_ToProto(mapCtx, in.VertexAiSearch); oneof != nil {
		out.Source = &pb.Retrieval_VertexAiSearch{VertexAiSearch: oneof}
	}
	if oneof := VertexRagStore_ToProto(mapCtx, in.VertexRagStore); oneof != nil {
		out.Source = &pb.Retrieval_VertexRagStore{VertexRagStore: oneof}
	}
	out.DisableAttribution = direct.ValueOf(in.DisableAttribution)
	return out
}
func RetrievalConfig_FromProto(mapCtx *direct.MapContext, in *pb.RetrievalConfig) *krm.RetrievalConfig {
	if in == nil {
		return nil
	}
	out := &krm.RetrievalConfig{}
	out.LatLng = LatLng_FromProto(mapCtx, in.GetLatLng())
	out.LanguageCode = in.LanguageCode
	return out
}
func RetrievalConfig_ToProto(mapCtx *direct.MapContext, in *krm.RetrievalConfig) *pb.RetrievalConfig {
	if in == nil {
		return nil
	}
	out := &pb.RetrievalConfig{}
	if oneof := LatLng_ToProto(mapCtx, in.LatLng); oneof != nil {
		out.LatLng = &pb.RetrievalConfig_LatLng{LatLng: oneof}
	}
	out.LanguageCode = in.LanguageCode
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Format = direct.LazyPtr(in.GetFormat())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.Default = Value_FromProto(mapCtx, in.GetDefault())
	out.Items = Schema_FromProto(mapCtx, in.GetItems())
	out.MinItems = direct.LazyPtr(in.GetMinItems())
	out.MaxItems = direct.LazyPtr(in.GetMaxItems())
	out.Enum = in.Enum
	// MISSING: Properties
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.LazyPtr(in.GetMinProperties())
	out.MaxProperties = direct.LazyPtr(in.GetMaxProperties())
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.MinLength = direct.LazyPtr(in.GetMinLength())
	out.MaxLength = direct.LazyPtr(in.GetMaxLength())
	out.Pattern = direct.LazyPtr(in.GetPattern())
	out.Example = Value_FromProto(mapCtx, in.GetExample())
	out.AnyOf = direct.Slice_FromProto(mapCtx, in.AnyOf, Schema_FromProto)
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	out.Type = direct.Enum_ToProto[pb.Type](mapCtx, in.Type)
	out.Format = direct.ValueOf(in.Format)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.Default = Value_ToProto(mapCtx, in.Default)
	out.Items = Schema_ToProto(mapCtx, in.Items)
	out.MinItems = direct.ValueOf(in.MinItems)
	out.MaxItems = direct.ValueOf(in.MaxItems)
	out.Enum = in.Enum
	// MISSING: Properties
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.ValueOf(in.MinProperties)
	out.MaxProperties = direct.ValueOf(in.MaxProperties)
	out.Minimum = direct.ValueOf(in.Minimum)
	out.Maximum = direct.ValueOf(in.Maximum)
	out.MinLength = direct.ValueOf(in.MinLength)
	out.MaxLength = direct.ValueOf(in.MaxLength)
	out.Pattern = direct.ValueOf(in.Pattern)
	out.Example = Value_ToProto(mapCtx, in.Example)
	out.AnyOf = direct.Slice_ToProto(mapCtx, in.AnyOf, Schema_ToProto)
	return out
}
func Tool_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.Tool {
	if in == nil {
		return nil
	}
	out := &krm.Tool{}
	out.FunctionDeclarations = direct.Slice_FromProto(mapCtx, in.FunctionDeclarations, FunctionDeclaration_FromProto)
	out.Retrieval = Retrieval_FromProto(mapCtx, in.GetRetrieval())
	out.GoogleSearchRetrieval = GoogleSearchRetrieval_FromProto(mapCtx, in.GetGoogleSearchRetrieval())
	return out
}
func Tool_ToProto(mapCtx *direct.MapContext, in *krm.Tool) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	out.FunctionDeclarations = direct.Slice_ToProto(mapCtx, in.FunctionDeclarations, FunctionDeclaration_ToProto)
	out.Retrieval = Retrieval_ToProto(mapCtx, in.Retrieval)
	out.GoogleSearchRetrieval = GoogleSearchRetrieval_ToProto(mapCtx, in.GoogleSearchRetrieval)
	return out
}
func ToolConfig_FromProto(mapCtx *direct.MapContext, in *pb.ToolConfig) *krm.ToolConfig {
	if in == nil {
		return nil
	}
	out := &krm.ToolConfig{}
	out.FunctionCallingConfig = FunctionCallingConfig_FromProto(mapCtx, in.GetFunctionCallingConfig())
	out.RetrievalConfig = RetrievalConfig_FromProto(mapCtx, in.GetRetrievalConfig())
	return out
}
func ToolConfig_ToProto(mapCtx *direct.MapContext, in *krm.ToolConfig) *pb.ToolConfig {
	if in == nil {
		return nil
	}
	out := &pb.ToolConfig{}
	out.FunctionCallingConfig = FunctionCallingConfig_ToProto(mapCtx, in.FunctionCallingConfig)
	out.RetrievalConfig = RetrievalConfig_ToProto(mapCtx, in.RetrievalConfig)
	return out
}
func VertexAISearch_FromProto(mapCtx *direct.MapContext, in *pb.VertexAISearch) *krm.VertexAISearch {
	if in == nil {
		return nil
	}
	out := &krm.VertexAISearch{}
	out.Datastore = direct.LazyPtr(in.GetDatastore())
	return out
}
func VertexAISearch_ToProto(mapCtx *direct.MapContext, in *krm.VertexAISearch) *pb.VertexAISearch {
	if in == nil {
		return nil
	}
	out := &pb.VertexAISearch{}
	out.Datastore = direct.ValueOf(in.Datastore)
	return out
}
func VertexRagStore_FromProto(mapCtx *direct.MapContext, in *pb.VertexRagStore) *krm.VertexRagStore {
	if in == nil {
		return nil
	}
	out := &krm.VertexRagStore{}
	out.RagResources = direct.Slice_FromProto(mapCtx, in.RagResources, VertexRagStore_RagResource_FromProto)
	out.SimilarityTopK = in.SimilarityTopK
	out.VectorDistanceThreshold = in.VectorDistanceThreshold
	out.RagRetrievalConfig = RagRetrievalConfig_FromProto(mapCtx, in.GetRagRetrievalConfig())
	return out
}
func VertexRagStore_ToProto(mapCtx *direct.MapContext, in *krm.VertexRagStore) *pb.VertexRagStore {
	if in == nil {
		return nil
	}
	out := &pb.VertexRagStore{}
	out.RagResources = direct.Slice_ToProto(mapCtx, in.RagResources, VertexRagStore_RagResource_ToProto)
	out.SimilarityTopK = in.SimilarityTopK
	out.VectorDistanceThreshold = in.VectorDistanceThreshold
	out.RagRetrievalConfig = RagRetrievalConfig_ToProto(mapCtx, in.RagRetrievalConfig)
	return out
}
func VertexRagStore_RagResource_FromProto(mapCtx *direct.MapContext, in *pb.VertexRagStore_RagResource) *krm.VertexRagStore_RagResource {
	if in == nil {
		return nil
	}
	out := &krm.VertexRagStore_RagResource{}
	out.RagCorpus = direct.LazyPtr(in.GetRagCorpus())
	out.RagFileIds = in.RagFileIds
	return out
}
func VertexRagStore_RagResource_ToProto(mapCtx *direct.MapContext, in *krm.VertexRagStore_RagResource) *pb.VertexRagStore_RagResource {
	if in == nil {
		return nil
	}
	out := &pb.VertexRagStore_RagResource{}
	out.RagCorpus = direct.ValueOf(in.RagCorpus)
	out.RagFileIds = in.RagFileIds
	return out
}
