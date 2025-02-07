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

package ai

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
)
func AiCachedContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiCachedContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiCachedContentObservedState{}
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
func AiCachedContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiCachedContentObservedState) *pb.CachedContent {
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
func AiCachedContentSpec_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiCachedContentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiCachedContentSpec{}
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
func AiCachedContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiCachedContentSpec) *pb.CachedContent {
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
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Model = in.Model
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
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Model = in.Model
	if oneof := Content_ToProto(mapCtx, in.SystemInstruction); oneof != nil {
		out.SystemInstruction = &pb.CachedContent_SystemInstruction{SystemInstruction: oneof}
	}
	out.Contents = direct.Slice_ToProto(mapCtx, in.Contents, Content_ToProto)
	out.Tools = direct.Slice_ToProto(mapCtx, in.Tools, Tool_ToProto)
	if oneof := ToolConfig_ToProto(mapCtx, in.ToolConfig); oneof != nil {
		out.ToolConfig = &pb.CachedContent_ToolConfig{ToolConfig: oneof}
	}
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
func CodeExecution_FromProto(mapCtx *direct.MapContext, in *pb.CodeExecution) *krm.CodeExecution {
	if in == nil {
		return nil
	}
	out := &krm.CodeExecution{}
	return out
}
func CodeExecution_ToProto(mapCtx *direct.MapContext, in *krm.CodeExecution) *pb.CodeExecution {
	if in == nil {
		return nil
	}
	out := &pb.CodeExecution{}
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
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, Part_FromProto)
	out.Role = direct.LazyPtr(in.GetRole())
	return out
}
func Content_ToProto(mapCtx *direct.MapContext, in *krm.Content) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, Part_ToProto)
	out.Role = direct.ValueOf(in.Role)
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
	out.ID = direct.LazyPtr(in.GetId())
	out.Name = direct.LazyPtr(in.GetName())
	out.Args = Args_FromProto(mapCtx, in.GetArgs())
	return out
}
func FunctionCall_ToProto(mapCtx *direct.MapContext, in *krm.FunctionCall) *pb.FunctionCall {
	if in == nil {
		return nil
	}
	out := &pb.FunctionCall{}
	out.Id = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.Name)
	if oneof := Args_ToProto(mapCtx, in.Args); oneof != nil {
		out.Args = &pb.FunctionCall_Args{Args: oneof}
	}
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
	if oneof := Schema_ToProto(mapCtx, in.Parameters); oneof != nil {
		out.Parameters = &pb.FunctionDeclaration_Parameters{Parameters: oneof}
	}
	if oneof := Schema_ToProto(mapCtx, in.Response); oneof != nil {
		out.Response = &pb.FunctionDeclaration_Response{Response: oneof}
	}
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
	out.FunctionCall = FunctionCall_FromProto(mapCtx, in.GetFunctionCall())
	out.FunctionResponse = FunctionResponse_FromProto(mapCtx, in.GetFunctionResponse())
	out.FileData = FileData_FromProto(mapCtx, in.GetFileData())
	out.ExecutableCode = ExecutableCode_FromProto(mapCtx, in.GetExecutableCode())
	out.CodeExecutionResult = CodeExecutionResult_FromProto(mapCtx, in.GetCodeExecutionResult())
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
	if oneof := FunctionCall_ToProto(mapCtx, in.FunctionCall); oneof != nil {
		out.Data = &pb.Part_FunctionCall{FunctionCall: oneof}
	}
	if oneof := FunctionResponse_ToProto(mapCtx, in.FunctionResponse); oneof != nil {
		out.Data = &pb.Part_FunctionResponse{FunctionResponse: oneof}
	}
	if oneof := FileData_ToProto(mapCtx, in.FileData); oneof != nil {
		out.Data = &pb.Part_FileData{FileData: oneof}
	}
	if oneof := ExecutableCode_ToProto(mapCtx, in.ExecutableCode); oneof != nil {
		out.Data = &pb.Part_ExecutableCode{ExecutableCode: oneof}
	}
	if oneof := CodeExecutionResult_ToProto(mapCtx, in.CodeExecutionResult); oneof != nil {
		out.Data = &pb.Part_CodeExecutionResult{CodeExecutionResult: oneof}
	}
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Format = direct.LazyPtr(in.GetFormat())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.Enum = in.Enum
	out.Items = Schema_FromProto(mapCtx, in.GetItems())
	out.MaxItems = direct.LazyPtr(in.GetMaxItems())
	out.MinItems = direct.LazyPtr(in.GetMinItems())
	// MISSING: Properties
	out.Required = in.Required
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	out.Type = direct.Enum_ToProto[pb.Type](mapCtx, in.Type)
	out.Format = direct.ValueOf(in.Format)
	out.Description = direct.ValueOf(in.Description)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.Enum = in.Enum
	if oneof := Schema_ToProto(mapCtx, in.Items); oneof != nil {
		out.Items = &pb.Schema_Items{Items: oneof}
	}
	out.MaxItems = direct.ValueOf(in.MaxItems)
	out.MinItems = direct.ValueOf(in.MinItems)
	// MISSING: Properties
	out.Required = in.Required
	return out
}
func Tool_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.Tool {
	if in == nil {
		return nil
	}
	out := &krm.Tool{}
	out.FunctionDeclarations = direct.Slice_FromProto(mapCtx, in.FunctionDeclarations, FunctionDeclaration_FromProto)
	out.GoogleSearchRetrieval = GoogleSearchRetrieval_FromProto(mapCtx, in.GetGoogleSearchRetrieval())
	out.CodeExecution = CodeExecution_FromProto(mapCtx, in.GetCodeExecution())
	out.GoogleSearch = Tool_GoogleSearch_FromProto(mapCtx, in.GetGoogleSearch())
	return out
}
func Tool_ToProto(mapCtx *direct.MapContext, in *krm.Tool) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	out.FunctionDeclarations = direct.Slice_ToProto(mapCtx, in.FunctionDeclarations, FunctionDeclaration_ToProto)
	out.GoogleSearchRetrieval = GoogleSearchRetrieval_ToProto(mapCtx, in.GoogleSearchRetrieval)
	out.CodeExecution = CodeExecution_ToProto(mapCtx, in.CodeExecution)
	out.GoogleSearch = Tool_GoogleSearch_ToProto(mapCtx, in.GoogleSearch)
	return out
}
func ToolConfig_FromProto(mapCtx *direct.MapContext, in *pb.ToolConfig) *krm.ToolConfig {
	if in == nil {
		return nil
	}
	out := &krm.ToolConfig{}
	out.FunctionCallingConfig = FunctionCallingConfig_FromProto(mapCtx, in.GetFunctionCallingConfig())
	return out
}
func ToolConfig_ToProto(mapCtx *direct.MapContext, in *krm.ToolConfig) *pb.ToolConfig {
	if in == nil {
		return nil
	}
	out := &pb.ToolConfig{}
	out.FunctionCallingConfig = FunctionCallingConfig_ToProto(mapCtx, in.FunctionCallingConfig)
	return out
}
func Tool_GoogleSearch_FromProto(mapCtx *direct.MapContext, in *pb.Tool_GoogleSearch) *krm.Tool_GoogleSearch {
	if in == nil {
		return nil
	}
	out := &krm.Tool_GoogleSearch{}
	return out
}
func Tool_GoogleSearch_ToProto(mapCtx *direct.MapContext, in *krm.Tool_GoogleSearch) *pb.Tool_GoogleSearch {
	if in == nil {
		return nil
	}
	out := &pb.Tool_GoogleSearch{}
	return out
}
