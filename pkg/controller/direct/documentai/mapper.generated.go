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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
)
func DocumentSchema_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EntityTypes = direct.Slice_FromProto(mapCtx, in.EntityTypes, DocumentSchema_EntityType_FromProto)
	out.Metadata = DocumentSchema_Metadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func DocumentSchema_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.EntityTypes = direct.Slice_ToProto(mapCtx, in.EntityTypes, DocumentSchema_EntityType_ToProto)
	out.Metadata = DocumentSchema_Metadata_ToProto(mapCtx, in.Metadata)
	return out
}
func DocumentSchema_EntityType_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType) *krm.DocumentSchema_EntityType {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType{}
	out.EnumValues = DocumentSchema_EntityType_EnumValues_FromProto(mapCtx, in.GetEnumValues())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BaseTypes = in.BaseTypes
	out.Properties = direct.Slice_FromProto(mapCtx, in.Properties, DocumentSchema_EntityType_Property_FromProto)
	out.EntityTypeMetadata = EntityTypeMetadata_FromProto(mapCtx, in.GetEntityTypeMetadata())
	return out
}
func DocumentSchema_EntityType_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType) *pb.DocumentSchema_EntityType {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType{}
	if oneof := DocumentSchema_EntityType_EnumValues_ToProto(mapCtx, in.EnumValues); oneof != nil {
		out.ValueSource = &pb.DocumentSchema_EntityType_EnumValues_{EnumValues: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.BaseTypes = in.BaseTypes
	out.Properties = direct.Slice_ToProto(mapCtx, in.Properties, DocumentSchema_EntityType_Property_ToProto)
	out.EntityTypeMetadata = EntityTypeMetadata_ToProto(mapCtx, in.EntityTypeMetadata)
	return out
}
func DocumentSchema_EntityType_EnumValues_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType_EnumValues) *krm.DocumentSchema_EntityType_EnumValues {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType_EnumValues{}
	out.Values = in.Values
	return out
}
func DocumentSchema_EntityType_EnumValues_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType_EnumValues) *pb.DocumentSchema_EntityType_EnumValues {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType_EnumValues{}
	out.Values = in.Values
	return out
}
func DocumentSchema_EntityType_Property_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType_Property) *krm.DocumentSchema_EntityType_Property {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType_Property{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ValueType = direct.LazyPtr(in.GetValueType())
	out.OccurrenceType = direct.Enum_FromProto(mapCtx, in.GetOccurrenceType())
	out.PropertyMetadata = PropertyMetadata_FromProto(mapCtx, in.GetPropertyMetadata())
	return out
}
func DocumentSchema_EntityType_Property_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType_Property) *pb.DocumentSchema_EntityType_Property {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType_Property{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ValueType = direct.ValueOf(in.ValueType)
	out.OccurrenceType = direct.Enum_ToProto[pb.DocumentSchema_EntityType_Property_OccurrenceType](mapCtx, in.OccurrenceType)
	out.PropertyMetadata = PropertyMetadata_ToProto(mapCtx, in.PropertyMetadata)
	return out
}
func EvaluationReference_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationReference) *krm.EvaluationReference {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationReference{}
	out.Operation = direct.LazyPtr(in.GetOperation())
	out.Evaluation = direct.LazyPtr(in.GetEvaluation())
	out.AggregateMetrics = Evaluation_Metrics_FromProto(mapCtx, in.GetAggregateMetrics())
	out.AggregateMetricsExact = Evaluation_Metrics_FromProto(mapCtx, in.GetAggregateMetricsExact())
	return out
}
func EvaluationReference_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationReference) *pb.EvaluationReference {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationReference{}
	out.Operation = direct.ValueOf(in.Operation)
	out.Evaluation = direct.ValueOf(in.Evaluation)
	out.AggregateMetrics = Evaluation_Metrics_ToProto(mapCtx, in.AggregateMetrics)
	out.AggregateMetricsExact = Evaluation_Metrics_ToProto(mapCtx, in.AggregateMetricsExact)
	return out
}
func ProcessorVersion_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion) *krm.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DocumentSchema = DocumentSchema_FromProto(mapCtx, in.GetDocumentSchema())
	// MISSING: State
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LatestEvaluation = EvaluationReference_FromProto(mapCtx, in.GetLatestEvaluation())
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
func ProcessorVersion_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion) *pb.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DocumentSchema = DocumentSchema_ToProto(mapCtx, in.DocumentSchema)
	// MISSING: State
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LatestEvaluation = EvaluationReference_ToProto(mapCtx, in.LatestEvaluation)
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
func ProcessorVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion) *krm.ProcessorVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DocumentSchema
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: CreateTime
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
func ProcessorVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersionObservedState) *pb.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DocumentSchema
	out.State = direct.Enum_ToProto[pb.ProcessorVersion_State](mapCtx, in.State)
	// MISSING: CreateTime
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
