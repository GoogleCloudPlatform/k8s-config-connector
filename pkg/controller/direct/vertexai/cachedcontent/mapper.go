// Copyright 2026 Google LLC
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

package cachedcontent

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	structpb "google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *apiextensionsv1.JSON {
	if in == nil {
		return nil
	}
	b, err := in.MarshalJSON()
	if err != nil {
		mapCtx.Errorf("marshaling structpb.Value to JSON: %v", err)
		return nil
	}
	return &apiextensionsv1.JSON{Raw: b}
}
func Value_ToProto(mapCtx *direct.MapContext, in *apiextensionsv1.JSON) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if err := out.UnmarshalJSON(in.Raw); err != nil {
		mapCtx.Errorf("unmarshaling JSON to structpb.Value: %v", err)
		return nil
	}
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
	if in.Data != nil {
		out.Data = in.Data
	}
	return out
}
func CodeExecutionResult_FromProto(mapCtx *direct.MapContext, in *pb.CodeExecutionResult) *krm.CodeExecutionResult {
	if in == nil {
		return nil
	}
	out := &krm.CodeExecutionResult{}
	out.Outcome = direct.Enum_FromProto(mapCtx, in.GetOutcome())
	out.Output = direct.LazyPtr(in.GetOutput())
	return out
}
func CodeExecutionResult_ToProto(mapCtx *direct.MapContext, in *krm.CodeExecutionResult) *pb.CodeExecutionResult {
	if in == nil {
		return nil
	}
	out := &pb.CodeExecutionResult{}
	out.Outcome = direct.Enum_ToProto[pb.CodeExecutionResult_Outcome](mapCtx, in.Outcome)
	out.Output = direct.ValueOf(in.Output)
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
func EnterpriseWebSearch_FromProto(mapCtx *direct.MapContext, in *pb.EnterpriseWebSearch) *krm.EnterpriseWebSearch {
	if in == nil {
		return nil
	}
	out := &krm.EnterpriseWebSearch{}
	out.ExcludeDomains = in.ExcludeDomains
	return out
}
func EnterpriseWebSearch_ToProto(mapCtx *direct.MapContext, in *krm.EnterpriseWebSearch) *pb.EnterpriseWebSearch {
	if in == nil {
		return nil
	}
	out := &pb.EnterpriseWebSearch{}
	out.ExcludeDomains = in.ExcludeDomains
	return out
}
func ExecutableCode_FromProto(mapCtx *direct.MapContext, in *pb.ExecutableCode) *krm.ExecutableCode {
	if in == nil {
		return nil
	}
	out := &krm.ExecutableCode{}
	out.Language = direct.Enum_FromProto(mapCtx, in.GetLanguage())
	out.Code = direct.LazyPtr(in.GetCode())
	return out
}
func ExecutableCode_ToProto(mapCtx *direct.MapContext, in *krm.ExecutableCode) *pb.ExecutableCode {
	if in == nil {
		return nil
	}
	out := &pb.ExecutableCode{}
	out.Language = direct.Enum_ToProto[pb.ExecutableCode_Language](mapCtx, in.Language)
	out.Code = direct.ValueOf(in.Code)
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
	if val := direct.Struct_FromProto(mapCtx, in.GetArgs()); val != nil {
		out.Args = *val
	}
	return out
}
func FunctionCall_ToProto(mapCtx *direct.MapContext, in *krm.FunctionCall) *pb.FunctionCall {
	if in == nil {
		return nil
	}
	out := &pb.FunctionCall{}
	out.Name = direct.ValueOf(in.Name)
	out.Args = direct.Struct_ToProto(mapCtx, &in.Args)
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
	out.ParametersJsonSchema = Value_FromProto(mapCtx, in.GetParametersJsonSchema())
	out.Response = Schema_FromProto(mapCtx, in.GetResponse())
	out.ResponseJsonSchema = Value_FromProto(mapCtx, in.GetResponseJsonSchema())
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
	out.ParametersJsonSchema = Value_ToProto(mapCtx, in.ParametersJsonSchema)
	out.Response = Schema_ToProto(mapCtx, in.Response)
	out.ResponseJsonSchema = Value_ToProto(mapCtx, in.ResponseJsonSchema)
	return out
}
func GoogleMaps_FromProto(mapCtx *direct.MapContext, in *pb.GoogleMaps) *krm.GoogleMaps {
	if in == nil {
		return nil
	}
	out := &krm.GoogleMaps{}
	return out
}
func GoogleMaps_ToProto(mapCtx *direct.MapContext, in *krm.GoogleMaps) *pb.GoogleMaps {
	if in == nil {
		return nil
	}
	out := &pb.GoogleMaps{}
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
	out.ExecutableCode = ExecutableCode_FromProto(mapCtx, in.GetExecutableCode())
	out.CodeExecutionResult = CodeExecutionResult_FromProto(mapCtx, in.GetCodeExecutionResult())
	out.Thought = direct.LazyPtr(in.GetThought())
	out.ThoughtSignature = in.GetThoughtSignature()
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
	if oneof := ExecutableCode_ToProto(mapCtx, in.ExecutableCode); oneof != nil {
		out.Data = &pb.Part_ExecutableCode{ExecutableCode: oneof}
	}
	if oneof := CodeExecutionResult_ToProto(mapCtx, in.CodeExecutionResult); oneof != nil {
		out.Data = &pb.Part_CodeExecutionResult{CodeExecutionResult: oneof}
	}
	out.Thought = direct.ValueOf(in.Thought)
	if in.ThoughtSignature != nil {
		out.ThoughtSignature = in.ThoughtSignature
	}
	if oneof := VideoMetadata_ToProto(mapCtx, in.VideoMetadata); oneof != nil {
		out.Metadata = &pb.Part_VideoMetadata{VideoMetadata: oneof}
	}
	return out
}
func Part_Text_ToProto(mapCtx *direct.MapContext, in *string) *pb.Part_Text {
	if in == nil {
		return nil
	}
	return &pb.Part_Text{Text: *in}
}
func RagRetrievalConfig_FromProto(mapCtx *direct.MapContext, in *pb.RagRetrievalConfig) *krm.RagRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &krm.RagRetrievalConfig{}
	out.TopK = direct.LazyPtr(in.GetTopK())
	out.Filter = RagRetrievalConfig_Filter_FromProto(mapCtx, in.GetFilter())
	out.Ranking = RagRetrievalConfig_Ranking_FromProto(mapCtx, in.GetRanking())
	return out
}
func RagRetrievalConfig_ToProto(mapCtx *direct.MapContext, in *krm.RagRetrievalConfig) *pb.RagRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &pb.RagRetrievalConfig{}
	out.TopK = direct.ValueOf(in.TopK)
	out.Filter = RagRetrievalConfig_Filter_ToProto(mapCtx, in.Filter)
	out.Ranking = RagRetrievalConfig_Ranking_ToProto(mapCtx, in.Ranking)
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
func RagRetrievalConfig_Filter_VectorDistanceThreshold_ToProto(mapCtx *direct.MapContext, in *float64) *pb.RagRetrievalConfig_Filter_VectorDistanceThreshold {
	if in == nil {
		return nil
	}
	return &pb.RagRetrievalConfig_Filter_VectorDistanceThreshold{VectorDistanceThreshold: *in}
}
func RagRetrievalConfig_Filter_VectorSimilarityThreshold_ToProto(mapCtx *direct.MapContext, in *float64) *pb.RagRetrievalConfig_Filter_VectorSimilarityThreshold {
	if in == nil {
		return nil
	}
	return &pb.RagRetrievalConfig_Filter_VectorSimilarityThreshold{VectorSimilarityThreshold: *in}
}
func RagRetrievalConfig_Ranking_FromProto(mapCtx *direct.MapContext, in *pb.RagRetrievalConfig_Ranking) *krm.RagRetrievalConfig_Ranking {
	if in == nil {
		return nil
	}
	out := &krm.RagRetrievalConfig_Ranking{}
	out.RankService = RagRetrievalConfig_Ranking_RankService_FromProto(mapCtx, in.GetRankService())
	out.LlmRanker = RagRetrievalConfig_Ranking_LlmRanker_FromProto(mapCtx, in.GetLlmRanker())
	return out
}
func RagRetrievalConfig_Ranking_ToProto(mapCtx *direct.MapContext, in *krm.RagRetrievalConfig_Ranking) *pb.RagRetrievalConfig_Ranking {
	if in == nil {
		return nil
	}
	out := &pb.RagRetrievalConfig_Ranking{}
	if oneof := RagRetrievalConfig_Ranking_RankService_ToProto(mapCtx, in.RankService); oneof != nil {
		out.RankingConfig = &pb.RagRetrievalConfig_Ranking_RankService_{RankService: oneof}
	}
	if oneof := RagRetrievalConfig_Ranking_LlmRanker_ToProto(mapCtx, in.LlmRanker); oneof != nil {
		out.RankingConfig = &pb.RagRetrievalConfig_Ranking_LlmRanker_{LlmRanker: oneof}
	}
	return out
}
func RagRetrievalConfig_Ranking_LlmRanker_FromProto(mapCtx *direct.MapContext, in *pb.RagRetrievalConfig_Ranking_LlmRanker) *krm.RagRetrievalConfig_Ranking_LlmRanker {
	if in == nil {
		return nil
	}
	out := &krm.RagRetrievalConfig_Ranking_LlmRanker{}
	out.ModelName = in.ModelName
	return out
}
func RagRetrievalConfig_Ranking_LlmRanker_ToProto(mapCtx *direct.MapContext, in *krm.RagRetrievalConfig_Ranking_LlmRanker) *pb.RagRetrievalConfig_Ranking_LlmRanker {
	if in == nil {
		return nil
	}
	out := &pb.RagRetrievalConfig_Ranking_LlmRanker{}
	out.ModelName = in.ModelName
	return out
}
func RagRetrievalConfig_Ranking_RankService_FromProto(mapCtx *direct.MapContext, in *pb.RagRetrievalConfig_Ranking_RankService) *krm.RagRetrievalConfig_Ranking_RankService {
	if in == nil {
		return nil
	}
	out := &krm.RagRetrievalConfig_Ranking_RankService{}
	out.ModelName = in.ModelName
	return out
}
func RagRetrievalConfig_Ranking_RankService_ToProto(mapCtx *direct.MapContext, in *krm.RagRetrievalConfig_Ranking_RankService) *pb.RagRetrievalConfig_Ranking_RankService {
	if in == nil {
		return nil
	}
	out := &pb.RagRetrievalConfig_Ranking_RankService{}
	out.ModelName = in.ModelName
	return out
}
func Retrieval_FromProto(mapCtx *direct.MapContext, in *pb.Retrieval) *krm.Retrieval {
	if in == nil {
		return nil
	}
	out := &krm.Retrieval{}
	out.VertexAiSearch = VertexAiSearch_FromProto(mapCtx, in.GetVertexAiSearch())
	out.VertexRagStore = VertexRagStore_FromProto(mapCtx, in.GetVertexRagStore())
	out.DisableAttribution = direct.LazyPtr(in.GetDisableAttribution())
	return out
}
func Retrieval_ToProto(mapCtx *direct.MapContext, in *krm.Retrieval) *pb.Retrieval {
	if in == nil {
		return nil
	}
	out := &pb.Retrieval{}
	if oneof := VertexAiSearch_ToProto(mapCtx, in.VertexAiSearch); oneof != nil {
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
	out.LatLng = LatLng_ToProto(mapCtx, in.LatLng)
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
	out.AdditionalProperties = Value_FromProto(mapCtx, in.GetAdditionalProperties())
	out.Ref = direct.LazyPtr(in.GetRef())
	// MISSING: Defs
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
	out.AdditionalProperties = Value_ToProto(mapCtx, in.AdditionalProperties)
	out.Ref = direct.ValueOf(in.Ref)
	// MISSING: Defs
	return out
}
func Tool_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.Tool {
	if in == nil {
		return nil
	}
	out := &krm.Tool{}
	out.FunctionDeclarations = direct.Slice_FromProto(mapCtx, in.FunctionDeclarations, FunctionDeclaration_FromProto)
	out.Retrieval = Retrieval_FromProto(mapCtx, in.GetRetrieval())
	out.GoogleSearch = Tool_GoogleSearch_FromProto(mapCtx, in.GetGoogleSearch())
	out.GoogleSearchRetrieval = GoogleSearchRetrieval_FromProto(mapCtx, in.GetGoogleSearchRetrieval())
	out.GoogleMaps = GoogleMaps_FromProto(mapCtx, in.GetGoogleMaps())
	out.EnterpriseWebSearch = EnterpriseWebSearch_FromProto(mapCtx, in.GetEnterpriseWebSearch())
	out.CodeExecution = Tool_CodeExecution_FromProto(mapCtx, in.GetCodeExecution())
	out.URLContext = URLContext_FromProto(mapCtx, in.GetUrlContext())
	out.ComputerUse = Tool_ComputerUse_FromProto(mapCtx, in.GetComputerUse())
	return out
}
func Tool_ToProto(mapCtx *direct.MapContext, in *krm.Tool) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	out.FunctionDeclarations = direct.Slice_ToProto(mapCtx, in.FunctionDeclarations, FunctionDeclaration_ToProto)
	out.Retrieval = Retrieval_ToProto(mapCtx, in.Retrieval)
	out.GoogleSearch = Tool_GoogleSearch_ToProto(mapCtx, in.GoogleSearch)
	out.GoogleSearchRetrieval = GoogleSearchRetrieval_ToProto(mapCtx, in.GoogleSearchRetrieval)
	out.GoogleMaps = GoogleMaps_ToProto(mapCtx, in.GoogleMaps)
	out.EnterpriseWebSearch = EnterpriseWebSearch_ToProto(mapCtx, in.EnterpriseWebSearch)
	out.CodeExecution = Tool_CodeExecution_ToProto(mapCtx, in.CodeExecution)
	out.UrlContext = URLContext_ToProto(mapCtx, in.URLContext)
	out.ComputerUse = Tool_ComputerUse_ToProto(mapCtx, in.ComputerUse)
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
func Tool_CodeExecution_FromProto(mapCtx *direct.MapContext, in *pb.Tool_CodeExecution) *krm.Tool_CodeExecution {
	if in == nil {
		return nil
	}
	out := &krm.Tool_CodeExecution{}
	return out
}
func Tool_CodeExecution_ToProto(mapCtx *direct.MapContext, in *krm.Tool_CodeExecution) *pb.Tool_CodeExecution {
	if in == nil {
		return nil
	}
	out := &pb.Tool_CodeExecution{}
	return out
}
func Tool_ComputerUse_FromProto(mapCtx *direct.MapContext, in *pb.Tool_ComputerUse) *krm.Tool_ComputerUse {
	if in == nil {
		return nil
	}
	out := &krm.Tool_ComputerUse{}
	out.Environment = direct.Enum_FromProto(mapCtx, in.GetEnvironment())
	return out
}
func Tool_ComputerUse_ToProto(mapCtx *direct.MapContext, in *krm.Tool_ComputerUse) *pb.Tool_ComputerUse {
	if in == nil {
		return nil
	}
	out := &pb.Tool_ComputerUse{}
	out.Environment = direct.Enum_ToProto[pb.Tool_ComputerUse_Environment](mapCtx, in.Environment)
	return out
}
func Tool_GoogleSearch_FromProto(mapCtx *direct.MapContext, in *pb.Tool_GoogleSearch) *krm.Tool_GoogleSearch {
	if in == nil {
		return nil
	}
	out := &krm.Tool_GoogleSearch{}
	out.ExcludeDomains = in.ExcludeDomains
	return out
}
func Tool_GoogleSearch_ToProto(mapCtx *direct.MapContext, in *krm.Tool_GoogleSearch) *pb.Tool_GoogleSearch {
	if in == nil {
		return nil
	}
	out := &pb.Tool_GoogleSearch{}
	out.ExcludeDomains = in.ExcludeDomains
	return out
}
func URLContext_FromProto(mapCtx *direct.MapContext, in *pb.UrlContext) *krm.URLContext {
	if in == nil {
		return nil
	}
	out := &krm.URLContext{}
	return out
}
func URLContext_ToProto(mapCtx *direct.MapContext, in *krm.URLContext) *pb.UrlContext {
	if in == nil {
		return nil
	}
	out := &pb.UrlContext{}
	return out
}
func VertexAiSearch_FromProto(mapCtx *direct.MapContext, in *pb.VertexAISearch) *krm.VertexAiSearch {
	if in == nil {
		return nil
	}
	out := &krm.VertexAiSearch{}
	out.Datastore = direct.LazyPtr(in.GetDatastore())
	out.Engine = direct.LazyPtr(in.GetEngine())
	out.MaxResults = direct.LazyPtr(in.GetMaxResults())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.DataStoreSpecs = direct.Slice_FromProto(mapCtx, in.DataStoreSpecs, VertexAiSearch_DataStoreSpec_FromProto)
	return out
}
func VertexAiSearch_ToProto(mapCtx *direct.MapContext, in *krm.VertexAiSearch) *pb.VertexAISearch {
	if in == nil {
		return nil
	}
	out := &pb.VertexAISearch{}
	out.Datastore = direct.ValueOf(in.Datastore)
	out.Engine = direct.ValueOf(in.Engine)
	out.MaxResults = direct.ValueOf(in.MaxResults)
	out.Filter = direct.ValueOf(in.Filter)
	out.DataStoreSpecs = direct.Slice_ToProto(mapCtx, in.DataStoreSpecs, VertexAiSearch_DataStoreSpec_ToProto)
	return out
}
func VertexAiSearch_DataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.VertexAISearch_DataStoreSpec) *krm.VertexAiSearch_DataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAiSearch_DataStoreSpec{}
	out.DataStore = direct.LazyPtr(in.GetDataStore())
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}
func VertexAiSearch_DataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAiSearch_DataStoreSpec) *pb.VertexAISearch_DataStoreSpec {
	if in == nil {
		return nil
	}
	out := &pb.VertexAISearch_DataStoreSpec{}
	out.DataStore = direct.ValueOf(in.DataStore)
	out.Filter = direct.ValueOf(in.Filter)
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

func FunctionResponse_FromProto(mapCtx *direct.MapContext, in *pb.FunctionResponse) *krm.FunctionResponse {
	if in == nil {
		return nil
	}
	out := &krm.FunctionResponse{}
	out.Name = direct.LazyPtr(in.GetName())
	if val := direct.Struct_FromProto(mapCtx, in.GetResponse()); val != nil {
		out.Response = *val
	}
	return out
}
func FunctionResponse_ToProto(mapCtx *direct.MapContext, in *krm.FunctionResponse) *pb.FunctionResponse {
	if in == nil {
		return nil
	}
	out := &pb.FunctionResponse{}
	out.Name = direct.ValueOf(in.Name)
	out.Response = direct.Struct_ToProto(mapCtx, &in.Response)
	return out
}

func VideoMetadata_FromProto(mapCtx *direct.MapContext, in *pb.VideoMetadata) *krm.VideoMetadata {
	if in == nil {
		return nil
	}
	out := &krm.VideoMetadata{}
	if in.GetStartOffset() != nil {
		out.StartOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartOffset())
	}
	if in.GetEndOffset() != nil {
		out.EndOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndOffset())
	}
	return out
}
func VideoMetadata_ToProto(mapCtx *direct.MapContext, in *krm.VideoMetadata) *pb.VideoMetadata {
	if in == nil {
		return nil
	}
	out := &pb.VideoMetadata{}
	if in.StartOffset != nil {
		out.StartOffset = direct.StringDuration_ToProto(mapCtx, in.StartOffset)
	}
	if in.EndOffset != nil {
		out.EndOffset = direct.StringDuration_ToProto(mapCtx, in.EndOffset)
	}
	return out
}

func LatLng_FromProto(mapCtx *direct.MapContext, in *latlng.LatLng) *krm.LatLng { return nil }
func LatLng_ToProto(mapCtx *direct.MapContext, in *krm.LatLng) *latlng.LatLng   { return nil }
func DynamicRetrievalConfig_FromProto(mapCtx *direct.MapContext, in *pb.DynamicRetrievalConfig) *krm.DynamicRetrievalConfig {
	if in == nil {
		return nil
	}
	out := &krm.DynamicRetrievalConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.DynamicThreshold = direct.LazyPtr(in.GetDynamicThreshold())
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

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	out.KMSKeyRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	return out
}

func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func VertexAICachedContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAICachedContentSpec) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Model = direct.ValueOf(in.Model)
	out.SystemInstruction = Content_ToProto(mapCtx, in.SystemInstruction)
	if in.Contents != nil {
		out.Contents = make([]*pb.Content, len(in.Contents))
		for i, v := range in.Contents {
			out.Contents[i] = Content_ToProto(mapCtx, &v)
		}
	}
	if in.Tools != nil {
		out.Tools = make([]*pb.Tool, len(in.Tools))
		for i, v := range in.Tools {
			out.Tools[i] = Tool_ToProto(mapCtx, &v)
		}
	}
	out.ToolConfig = ToolConfig_ToProto(mapCtx, in.ToolConfig)

	if in.ExpireTime != nil {
		out.Expiration = &pb.CachedContent_ExpireTime{ExpireTime: direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)}
	} else if in.TTL != nil {
		out.Expiration = &pb.CachedContent_Ttl{Ttl: direct.StringDuration_ToProto(mapCtx, in.TTL)}
	}

	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)

	return out
}

func CachedContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.VertexAICachedContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAICachedContentObservedState{}

	if in.GetCreateTime() != nil {
		out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	}
	if in.GetUpdateTime() != nil {
		out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	}

	// Handle Expiration
	switch expiration := in.Expiration.(type) {
	case *pb.CachedContent_ExpireTime:
		out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, expiration.ExpireTime)
	case *pb.CachedContent_Ttl:
		// TTL is a duration, usually not mapped back to a string in ObservedState unless explicitly asked.
		_ = expiration
	}

	// Fallback to in.GetExpireTime() if present
	if in.GetExpireTime() != nil {
		out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	}

	out.UsageMetadata = CachedContent_UsageMetadata_FromProto(mapCtx, in.GetUsageMetadata())
	return out
}

func CachedContent_UsageMetadata_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent_UsageMetadata) *krm.CachedContent_UsageMetadata {
	if in == nil {
		return nil
	}
	out := &krm.CachedContent_UsageMetadata{}
	out.TotalTokenCount = direct.LazyPtr(in.GetTotalTokenCount())
	out.AudioDurationSeconds = direct.LazyPtr(in.GetAudioDurationSeconds())
	out.VideoDurationSeconds = direct.LazyPtr(in.GetVideoDurationSeconds())
	out.TextCount = direct.LazyPtr(in.GetTextCount())
	out.ImageCount = direct.LazyPtr(in.GetImageCount())
	return out
}

func CachedContent_UsageMetadata_ToProto(mapCtx *direct.MapContext, in *krm.CachedContent_UsageMetadata) *pb.CachedContent_UsageMetadata {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent_UsageMetadata{}
	out.TotalTokenCount = direct.ValueOf(in.TotalTokenCount)
	out.AudioDurationSeconds = direct.ValueOf(in.AudioDurationSeconds)
	out.VideoDurationSeconds = direct.ValueOf(in.VideoDurationSeconds)
	out.TextCount = direct.ValueOf(in.TextCount)
	out.ImageCount = direct.ValueOf(in.ImageCount)
	return out
}
