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

package dialogflow

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
)
func ConversationModelEvaluation_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModelEvaluation) *krm.ConversationModelEvaluation {
	if in == nil {
		return nil
	}
	out := &krm.ConversationModelEvaluation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.EvaluationConfig = EvaluationConfig_FromProto(mapCtx, in.GetEvaluationConfig())
	// MISSING: CreateTime
	// MISSING: SmartReplyMetrics
	// MISSING: RawHumanEvalTemplateCsv
	return out
}
func ConversationModelEvaluation_ToProto(mapCtx *direct.MapContext, in *krm.ConversationModelEvaluation) *pb.ConversationModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModelEvaluation{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.EvaluationConfig = EvaluationConfig_ToProto(mapCtx, in.EvaluationConfig)
	// MISSING: CreateTime
	// MISSING: SmartReplyMetrics
	// MISSING: RawHumanEvalTemplateCsv
	return out
}
func ConversationModelEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModelEvaluation) *krm.ConversationModelEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationModelEvaluationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EvaluationConfig
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SmartReplyMetrics = SmartReplyMetrics_FromProto(mapCtx, in.GetSmartReplyMetrics())
	out.RawHumanEvalTemplateCsv = direct.LazyPtr(in.GetRawHumanEvalTemplateCsv())
	return out
}
func ConversationModelEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationModelEvaluationObservedState) *pb.ConversationModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModelEvaluation{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EvaluationConfig
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	if oneof := SmartReplyMetrics_ToProto(mapCtx, in.SmartReplyMetrics); oneof != nil {
		out.Metrics = &pb.ConversationModelEvaluation_SmartReplyMetrics{SmartReplyMetrics: oneof}
	}
	out.RawHumanEvalTemplateCsv = direct.ValueOf(in.RawHumanEvalTemplateCsv)
	return out
}
func DialogflowConversationModelEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModelEvaluation) *krm.DialogflowConversationModelEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelEvaluationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EvaluationConfig
	// MISSING: CreateTime
	// MISSING: SmartReplyMetrics
	// MISSING: RawHumanEvalTemplateCsv
	return out
}
func DialogflowConversationModelEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelEvaluationObservedState) *pb.ConversationModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModelEvaluation{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EvaluationConfig
	// MISSING: CreateTime
	// MISSING: SmartReplyMetrics
	// MISSING: RawHumanEvalTemplateCsv
	return out
}
func DialogflowConversationModelEvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModelEvaluation) *krm.DialogflowConversationModelEvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelEvaluationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EvaluationConfig
	// MISSING: CreateTime
	// MISSING: SmartReplyMetrics
	// MISSING: RawHumanEvalTemplateCsv
	return out
}
func DialogflowConversationModelEvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelEvaluationSpec) *pb.ConversationModelEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModelEvaluation{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EvaluationConfig
	// MISSING: CreateTime
	// MISSING: SmartReplyMetrics
	// MISSING: RawHumanEvalTemplateCsv
	return out
}
func EvaluationConfig_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationConfig) *krm.EvaluationConfig {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationConfig{}
	out.Datasets = direct.Slice_FromProto(mapCtx, in.Datasets, InputDataset_FromProto)
	out.SmartReplyConfig = EvaluationConfig_SmartReplyConfig_FromProto(mapCtx, in.GetSmartReplyConfig())
	out.SmartComposeConfig = EvaluationConfig_SmartComposeConfig_FromProto(mapCtx, in.GetSmartComposeConfig())
	return out
}
func EvaluationConfig_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationConfig) *pb.EvaluationConfig {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationConfig{}
	out.Datasets = direct.Slice_ToProto(mapCtx, in.Datasets, InputDataset_ToProto)
	if oneof := EvaluationConfig_SmartReplyConfig_ToProto(mapCtx, in.SmartReplyConfig); oneof != nil {
		out.ModelSpecificConfig = &pb.EvaluationConfig_SmartReplyConfig_{SmartReplyConfig: oneof}
	}
	if oneof := EvaluationConfig_SmartComposeConfig_ToProto(mapCtx, in.SmartComposeConfig); oneof != nil {
		out.ModelSpecificConfig = &pb.EvaluationConfig_SmartComposeConfig_{SmartComposeConfig: oneof}
	}
	return out
}
func EvaluationConfig_SmartComposeConfig_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationConfig_SmartComposeConfig) *krm.EvaluationConfig_SmartComposeConfig {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationConfig_SmartComposeConfig{}
	out.AllowlistDocument = direct.LazyPtr(in.GetAllowlistDocument())
	out.MaxResultCount = direct.LazyPtr(in.GetMaxResultCount())
	return out
}
func EvaluationConfig_SmartComposeConfig_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationConfig_SmartComposeConfig) *pb.EvaluationConfig_SmartComposeConfig {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationConfig_SmartComposeConfig{}
	out.AllowlistDocument = direct.ValueOf(in.AllowlistDocument)
	out.MaxResultCount = direct.ValueOf(in.MaxResultCount)
	return out
}
func EvaluationConfig_SmartReplyConfig_FromProto(mapCtx *direct.MapContext, in *pb.EvaluationConfig_SmartReplyConfig) *krm.EvaluationConfig_SmartReplyConfig {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationConfig_SmartReplyConfig{}
	out.AllowlistDocument = direct.LazyPtr(in.GetAllowlistDocument())
	out.MaxResultCount = direct.LazyPtr(in.GetMaxResultCount())
	return out
}
func EvaluationConfig_SmartReplyConfig_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationConfig_SmartReplyConfig) *pb.EvaluationConfig_SmartReplyConfig {
	if in == nil {
		return nil
	}
	out := &pb.EvaluationConfig_SmartReplyConfig{}
	out.AllowlistDocument = direct.ValueOf(in.AllowlistDocument)
	out.MaxResultCount = direct.ValueOf(in.MaxResultCount)
	return out
}
func InputDataset_FromProto(mapCtx *direct.MapContext, in *pb.InputDataset) *krm.InputDataset {
	if in == nil {
		return nil
	}
	out := &krm.InputDataset{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	return out
}
func InputDataset_ToProto(mapCtx *direct.MapContext, in *krm.InputDataset) *pb.InputDataset {
	if in == nil {
		return nil
	}
	out := &pb.InputDataset{}
	out.Dataset = direct.ValueOf(in.Dataset)
	return out
}
func SmartReplyMetrics_FromProto(mapCtx *direct.MapContext, in *pb.SmartReplyMetrics) *krm.SmartReplyMetrics {
	if in == nil {
		return nil
	}
	out := &krm.SmartReplyMetrics{}
	out.AllowlistCoverage = direct.LazyPtr(in.GetAllowlistCoverage())
	out.TopNMetrics = direct.Slice_FromProto(mapCtx, in.TopNMetrics, SmartReplyMetrics_TopNMetrics_FromProto)
	out.ConversationCount = direct.LazyPtr(in.GetConversationCount())
	return out
}
func SmartReplyMetrics_ToProto(mapCtx *direct.MapContext, in *krm.SmartReplyMetrics) *pb.SmartReplyMetrics {
	if in == nil {
		return nil
	}
	out := &pb.SmartReplyMetrics{}
	out.AllowlistCoverage = direct.ValueOf(in.AllowlistCoverage)
	out.TopNMetrics = direct.Slice_ToProto(mapCtx, in.TopNMetrics, SmartReplyMetrics_TopNMetrics_ToProto)
	out.ConversationCount = direct.ValueOf(in.ConversationCount)
	return out
}
func SmartReplyMetrics_TopNMetrics_FromProto(mapCtx *direct.MapContext, in *pb.SmartReplyMetrics_TopNMetrics) *krm.SmartReplyMetrics_TopNMetrics {
	if in == nil {
		return nil
	}
	out := &krm.SmartReplyMetrics_TopNMetrics{}
	out.N = direct.LazyPtr(in.GetN())
	out.Recall = direct.LazyPtr(in.GetRecall())
	return out
}
func SmartReplyMetrics_TopNMetrics_ToProto(mapCtx *direct.MapContext, in *krm.SmartReplyMetrics_TopNMetrics) *pb.SmartReplyMetrics_TopNMetrics {
	if in == nil {
		return nil
	}
	out := &pb.SmartReplyMetrics_TopNMetrics{}
	out.N = direct.ValueOf(in.N)
	out.Recall = direct.ValueOf(in.Recall)
	return out
}
