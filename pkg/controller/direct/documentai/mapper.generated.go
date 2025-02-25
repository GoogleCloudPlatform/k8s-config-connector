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

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DocumentAIProcessorVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion) *krm.DocumentAIProcessorVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DocumentAIProcessorVersionSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())

	// MISSING: State

	//out.LatestEvaluation = EvaluationReference_FromProto(mapCtx, in.GetLatestEvaluation())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	// MISSING: GoogleManaged
	out.DeprecationInfo = ProcessorVersion_DeprecationInfo_FromProto(mapCtx, in.GetDeprecationInfo())
	// MISSING: ModelType
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: GenAiModelInfo
	return out
}
func DocumentAIProcessorVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DocumentAIProcessorVersionSpec) *pb.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	//out.DocumentSchema = DocumentSchema_ToProto(mapCtx, in.DocumentSchema)
	// MISSING: State
	//out.LatestEvaluation = EvaluationReference_ToProto(mapCtx, in.LatestEvaluation)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	// MISSING: GoogleManaged
	out.DeprecationInfo = ProcessorVersion_DeprecationInfo_ToProto(mapCtx, in.DeprecationInfo)
	// MISSING: ModelType
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: GenAiModelInfo
	return out
}
func DocumentAIProcessorVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion) *krm.DocumentAIProcessorVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentAIProcessorVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	//out.DocumentSchema = DocumentSchema_FromProto(mapCtx, in.GetDocumentSchema())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: LatestEvaluation
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	out.GoogleManaged = direct.LazyPtr(in.GetGoogleManaged())
	// MISSING: DeprecationInfo
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	out.GenAiModelInfo = ProcessorVersion_GenAiModelInfo_FromProto(mapCtx, in.GetGenAiModelInfo())
	return out
}
func DocumentAIProcessorVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentAIProcessorVersionObservedState) *pb.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DocumentSchema
	out.State = direct.Enum_ToProto[pb.ProcessorVersion_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: LatestEvaluation
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	out.GoogleManaged = direct.ValueOf(in.GoogleManaged)
	// MISSING: DeprecationInfo
	out.ModelType = direct.Enum_ToProto[pb.ProcessorVersion_ModelType](mapCtx, in.ModelType)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	out.GenAiModelInfo = ProcessorVersion_GenAiModelInfo_ToProto(mapCtx, in.GenAiModelInfo)
	return out
}
func ProcessorVersion_DeprecationInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_DeprecationInfo) *krm.ProcessorVersion_DeprecationInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_DeprecationInfo{}
	out.DeprecationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeprecationTime())
	out.ReplacementProcessorVersion = direct.LazyPtr(in.GetReplacementProcessorVersion())
	return out
}
func ProcessorVersion_DeprecationInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_DeprecationInfo) *pb.ProcessorVersion_DeprecationInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_DeprecationInfo{}
	out.DeprecationTime = direct.StringTimestamp_ToProto(mapCtx, in.DeprecationTime)
	out.ReplacementProcessorVersion = direct.ValueOf(in.ReplacementProcessorVersion)
	return out
}
func ProcessorVersion_GenAiModelInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_GenAiModelInfo) *krm.ProcessorVersion_GenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_GenAiModelInfo{}
	out.FoundationGenAiModelInfo = ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_FromProto(mapCtx, in.GetFoundationGenAiModelInfo())
	out.CustomGenAiModelInfo = ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_FromProto(mapCtx, in.GetCustomGenAiModelInfo())
	return out
}
func ProcessorVersion_GenAiModelInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_GenAiModelInfo) *pb.ProcessorVersion_GenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_GenAiModelInfo{}
	if oneof := ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_ToProto(mapCtx, in.FoundationGenAiModelInfo); oneof != nil {
		out.ModelInfo = &pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_{FoundationGenAiModelInfo: oneof}
	}
	if oneof := ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_ToProto(mapCtx, in.CustomGenAiModelInfo); oneof != nil {
		out.ModelInfo = &pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_{CustomGenAiModelInfo: oneof}
	}
	return out
}
func ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo) *krm.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo{}
	out.CustomModelType = direct.Enum_FromProto(mapCtx, in.GetCustomModelType())
	out.BaseProcessorVersionID = direct.LazyPtr(in.GetBaseProcessorVersionId())
	return out
}
func ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo) *pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo{}
	out.CustomModelType = direct.Enum_ToProto[pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_CustomModelType](mapCtx, in.CustomModelType)
	out.BaseProcessorVersionId = direct.ValueOf(in.BaseProcessorVersionID)
	return out
}
func ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo) *krm.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo{}
	out.FinetuningAllowed = direct.LazyPtr(in.GetFinetuningAllowed())
	out.MinTrainLabeledDocuments = direct.LazyPtr(in.GetMinTrainLabeledDocuments())
	return out
}
func ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo) *pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo{}
	out.FinetuningAllowed = direct.ValueOf(in.FinetuningAllowed)
	out.MinTrainLabeledDocuments = direct.ValueOf(in.MinTrainLabeledDocuments)
	return out
}
func SummaryOptions_FromProto(mapCtx *direct.MapContext, in *pb.SummaryOptions) *krm.SummaryOptions {
	if in == nil {
		return nil
	}
	out := &krm.SummaryOptions{}
	out.Length = direct.Enum_FromProto(mapCtx, in.GetLength())
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	return out
}
func SummaryOptions_ToProto(mapCtx *direct.MapContext, in *krm.SummaryOptions) *pb.SummaryOptions {
	if in == nil {
		return nil
	}
	out := &pb.SummaryOptions{}
	out.Length = direct.Enum_ToProto[pb.SummaryOptions_Length](mapCtx, in.Length)
	out.Format = direct.Enum_ToProto[pb.SummaryOptions_Format](mapCtx, in.Format)
	return out
}
