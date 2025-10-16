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

// +generated:mapper
// krm.group: documentai.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.documentai.v1

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krmdocumentaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1beta1"
	krmkmsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DocumentAIProcessorObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krmdocumentaiv1alpha1.DocumentAIProcessorObservedState {
	if in == nil {
		return nil
	}
	out := &krmdocumentaiv1alpha1.DocumentAIProcessorObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DefaultProcessorVersion = direct.LazyPtr(in.GetDefaultProcessorVersion())
	out.ProcessorVersionAliases = direct.Slice_FromProto(mapCtx, in.ProcessorVersionAliases, ProcessorVersionAlias_v1alpha1_FromProto)
	out.ProcessEndpoint = direct.LazyPtr(in.GetProcessEndpoint())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: KMSKeyName
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentAIProcessorObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdocumentaiv1alpha1.DocumentAIProcessorObservedState) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Processor_State](mapCtx, in.State)
	out.DefaultProcessorVersion = direct.ValueOf(in.DefaultProcessorVersion)
	out.ProcessorVersionAliases = direct.Slice_ToProto(mapCtx, in.ProcessorVersionAliases, ProcessorVersionAlias_v1alpha1_ToProto)
	out.ProcessEndpoint = direct.ValueOf(in.ProcessEndpoint)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: KMSKeyName
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentAIProcessorSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krmdocumentaiv1alpha1.DocumentAIProcessorSpec {
	if in == nil {
		return nil
	}
	out := &krmdocumentaiv1alpha1.DocumentAIProcessorSpec{}
	out.Type = direct.LazyPtr(in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: KMSKeyName
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentAIProcessorSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdocumentaiv1alpha1.DocumentAIProcessorSpec) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: KMSKeyName
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentAIProcessorVersionObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion) *krm.DocumentAIProcessorVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentAIProcessorVersionObservedState{}
	// MISSING: Name
	out.DocumentSchema = DocumentSchema_v1beta1_FromProto(mapCtx, in.GetDocumentSchema())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LatestEvaluation = EvaluationReference_v1beta1_FromProto(mapCtx, in.GetLatestEvaluation())
	out.GoogleManaged = direct.LazyPtr(in.GetGoogleManaged())
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	out.GenAiModelInfo = ProcessorVersion_GenAiModelInfo_v1beta1_FromProto(mapCtx, in.GetGenAiModelInfo())
	return out
}
func DocumentAIProcessorVersionObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DocumentAIProcessorVersionObservedState) *pb.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion{}
	// MISSING: Name
	out.DocumentSchema = DocumentSchema_v1beta1_ToProto(mapCtx, in.DocumentSchema)
	out.State = direct.Enum_ToProto[pb.ProcessorVersion_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LatestEvaluation = EvaluationReference_v1beta1_ToProto(mapCtx, in.LatestEvaluation)
	out.GoogleManaged = direct.ValueOf(in.GoogleManaged)
	out.ModelType = direct.Enum_ToProto[pb.ProcessorVersion_ModelType](mapCtx, in.ModelType)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	out.GenAiModelInfo = ProcessorVersion_GenAiModelInfo_v1beta1_ToProto(mapCtx, in.GenAiModelInfo)
	return out
}
func DocumentAIProcessorVersionSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion) *krm.DocumentAIProcessorVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DocumentAIProcessorVersionSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	if in.GetKmsKeyName() != "" {
		out.KMSKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	if in.GetKmsKeyVersionName() != "" {
		out.KMSKeyVersionNameRef = &krmkmsv1alpha1.KMSCryptoKeyVersionRef{External: in.GetKmsKeyVersionName()}
	}
	out.DeprecationInfo = ProcessorVersion_DeprecationInfo_v1beta1_FromProto(mapCtx, in.GetDeprecationInfo())
	return out
}
func DocumentAIProcessorVersionSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DocumentAIProcessorVersionSpec) *pb.ProcessorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.KMSKeyNameRef != nil {
		out.KmsKeyName = in.KMSKeyNameRef.External
	}
	if in.KMSKeyVersionNameRef != nil {
		out.KmsKeyVersionName = in.KMSKeyVersionNameRef.External
	}
	out.DeprecationInfo = ProcessorVersion_DeprecationInfo_v1beta1_ToProto(mapCtx, in.DeprecationInfo)
	return out
}
func DocumentSchema_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EntityTypes = direct.Slice_FromProto(mapCtx, in.EntityTypes, DocumentSchema_EntityType_v1beta1_FromProto)
	out.Metadata = DocumentSchema_Metadata_v1beta1_FromProto(mapCtx, in.GetMetadata())
	return out
}
func DocumentSchema_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.EntityTypes = direct.Slice_ToProto(mapCtx, in.EntityTypes, DocumentSchema_EntityType_v1beta1_ToProto)
	out.Metadata = DocumentSchema_Metadata_v1beta1_ToProto(mapCtx, in.Metadata)
	return out
}
func DocumentSchema_EntityType_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType) *krm.DocumentSchema_EntityType {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType{}
	out.EnumValues = DocumentSchema_EntityType_EnumValues_v1beta1_FromProto(mapCtx, in.GetEnumValues())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Name = direct.LazyPtr(in.GetName())
	out.BaseTypes = in.BaseTypes
	out.Properties = direct.Slice_FromProto(mapCtx, in.Properties, DocumentSchema_EntityType_Property_v1beta1_FromProto)
	return out
}
func DocumentSchema_EntityType_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType) *pb.DocumentSchema_EntityType {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType{}
	if oneof := DocumentSchema_EntityType_EnumValues_v1beta1_ToProto(mapCtx, in.EnumValues); oneof != nil {
		out.ValueSource = &pb.DocumentSchema_EntityType_EnumValues_{EnumValues: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Name = direct.ValueOf(in.Name)
	out.BaseTypes = in.BaseTypes
	out.Properties = direct.Slice_ToProto(mapCtx, in.Properties, DocumentSchema_EntityType_Property_v1beta1_ToProto)
	return out
}
func DocumentSchema_EntityType_EnumValues_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType_EnumValues) *krm.DocumentSchema_EntityType_EnumValues {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType_EnumValues{}
	out.Values = in.Values
	return out
}
func DocumentSchema_EntityType_EnumValues_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType_EnumValues) *pb.DocumentSchema_EntityType_EnumValues {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType_EnumValues{}
	out.Values = in.Values
	return out
}
func DocumentSchema_EntityType_Property_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType_Property) *krm.DocumentSchema_EntityType_Property {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType_Property{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ValueType = direct.LazyPtr(in.GetValueType())
	out.OccurrenceType = direct.Enum_FromProto(mapCtx, in.GetOccurrenceType())
	out.Method = direct.Enum_FromProto(mapCtx, in.GetMethod())
	return out
}
func DocumentSchema_EntityType_Property_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType_Property) *pb.DocumentSchema_EntityType_Property {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType_Property{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ValueType = direct.ValueOf(in.ValueType)
	out.OccurrenceType = direct.Enum_ToProto[pb.DocumentSchema_EntityType_Property_OccurrenceType](mapCtx, in.OccurrenceType)
	out.Method = direct.Enum_ToProto[pb.DocumentSchema_EntityType_Property_Method](mapCtx, in.Method)
	return out
}
func EvaluationReference_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationReference) *krm.EvaluationReference {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationReference{}
	out.Operation = direct.LazyPtr(in.GetOperation())
	out.Evaluation = direct.LazyPtr(in.GetEvaluation())
	out.AggregateMetrics = Evaluation_Metrics_v1beta1_FromProto(mapCtx, in.GetAggregateMetrics())
	out.AggregateMetricsExact = Evaluation_Metrics_v1beta1_FromProto(mapCtx, in.GetAggregateMetricsExact())
	return out
}
func EvaluationReference_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationReference) *pb.EvaluationReference {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationReference{}
	out.Operation = direct.ValueOf(in.Operation)
	out.Evaluation = direct.ValueOf(in.Evaluation)
	out.AggregateMetrics = Evaluation_Metrics_v1beta1_ToProto(mapCtx, in.AggregateMetrics)
	out.AggregateMetricsExact = Evaluation_Metrics_v1beta1_ToProto(mapCtx, in.AggregateMetricsExact)
	return out
}
func ProcessorVersionAlias_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersionAlias) *krmdocumentaiv1alpha1.ProcessorVersionAlias {
	if in == nil {
		return nil
	}
	out := &krmdocumentaiv1alpha1.ProcessorVersionAlias{}
	out.Alias = direct.LazyPtr(in.GetAlias())
	out.ProcessorVersion = direct.LazyPtr(in.GetProcessorVersion())
	return out
}
func ProcessorVersionAlias_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdocumentaiv1alpha1.ProcessorVersionAlias) *pb.ProcessorVersionAlias {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersionAlias{}
	out.Alias = direct.ValueOf(in.Alias)
	out.ProcessorVersion = direct.ValueOf(in.ProcessorVersion)
	return out
}
func ProcessorVersion_DeprecationInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_DeprecationInfo) *krm.ProcessorVersion_DeprecationInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_DeprecationInfo{}
	out.DeprecationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeprecationTime())
	out.ReplacementProcessorVersion = direct.LazyPtr(in.GetReplacementProcessorVersion())
	return out
}
func ProcessorVersion_DeprecationInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_DeprecationInfo) *pb.ProcessorVersion_DeprecationInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_DeprecationInfo{}
	out.DeprecationTime = direct.StringTimestamp_ToProto(mapCtx, in.DeprecationTime)
	out.ReplacementProcessorVersion = direct.ValueOf(in.ReplacementProcessorVersion)
	return out
}
func ProcessorVersion_GenAiModelInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_GenAiModelInfo) *krm.ProcessorVersion_GenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_GenAiModelInfo{}
	out.FoundationGenAiModelInfo = ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_v1beta1_FromProto(mapCtx, in.GetFoundationGenAiModelInfo())
	out.CustomGenAiModelInfo = ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_v1beta1_FromProto(mapCtx, in.GetCustomGenAiModelInfo())
	return out
}
func ProcessorVersion_GenAiModelInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_GenAiModelInfo) *pb.ProcessorVersion_GenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_GenAiModelInfo{}
	if oneof := ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_v1beta1_ToProto(mapCtx, in.FoundationGenAiModelInfo); oneof != nil {
		out.ModelInfo = &pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_{FoundationGenAiModelInfo: oneof}
	}
	if oneof := ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_v1beta1_ToProto(mapCtx, in.CustomGenAiModelInfo); oneof != nil {
		out.ModelInfo = &pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_{CustomGenAiModelInfo: oneof}
	}
	return out
}
func ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo) *krm.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo{}
	out.CustomModelType = direct.Enum_FromProto(mapCtx, in.GetCustomModelType())
	out.BaseProcessorVersionID = direct.LazyPtr(in.GetBaseProcessorVersionId())
	return out
}
func ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo) *pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo{}
	out.CustomModelType = direct.Enum_ToProto[pb.ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo_CustomModelType](mapCtx, in.CustomModelType)
	out.BaseProcessorVersionId = direct.ValueOf(in.BaseProcessorVersionID)
	return out
}
func ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo) *krm.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo{}
	out.FinetuningAllowed = direct.LazyPtr(in.GetFinetuningAllowed())
	out.MinTrainLabeledDocuments = direct.LazyPtr(in.GetMinTrainLabeledDocuments())
	return out
}
func ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo) *pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo{}
	out.FinetuningAllowed = direct.ValueOf(in.FinetuningAllowed)
	out.MinTrainLabeledDocuments = direct.ValueOf(in.MinTrainLabeledDocuments)
	return out
}
