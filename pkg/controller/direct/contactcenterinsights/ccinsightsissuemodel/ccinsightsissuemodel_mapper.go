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

package ccinsightsissuemodel

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CCInsightsIssueModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsIssueModelSpec) *pb.IssueModel {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.InputDataConfig = InputDataConfig_ToProto(mapCtx, in.InputDataConfig)
	out.ModelType = direct.Enum_ToProto[pb.IssueModel_ModelType](mapCtx, in.ModelType)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}

func CCInsightsIssueModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel) *krm.CCInsightsIssueModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsIssueModelSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.InputDataConfig = InputDataConfig_FromProto(mapCtx, in.GetInputDataConfig())
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}

func CCInsightsIssueModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel) *krm.CCInsightsIssueModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsIssueModelObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.IssueCount = direct.LazyPtr(in.GetIssueCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.InputDataConfig = InputDataConfigObservedState_FromProto(mapCtx, in.GetInputDataConfig())
	out.TrainingStats = IssueModelLabelStats_FromProto(mapCtx, in.GetTrainingStats())
	return out
}

func CCInsightsIssueModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsIssueModelObservedState) *pb.IssueModel {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.IssueCount = direct.ValueOf(in.IssueCount)
	out.State = direct.Enum_ToProto[pb.IssueModel_State](mapCtx, in.State)
	out.InputDataConfig = InputDataConfigObservedState_ToProto(mapCtx, in.InputDataConfig)
	out.TrainingStats = IssueModelLabelStats_ToProto(mapCtx, in.TrainingStats)
	return out
}

func InputDataConfig_ToProto(mapCtx *direct.MapContext, in *krm.IssueModel_InputDataConfig) *pb.IssueModel_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel_InputDataConfig{}
	out.Medium = direct.Enum_ToProto[pb.Conversation_Medium](mapCtx, in.Medium)
	out.Filter = direct.ValueOf(in.Filter)
	return out
}

func InputDataConfig_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel_InputDataConfig) *krm.IssueModel_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &krm.IssueModel_InputDataConfig{}
	out.Medium = direct.Enum_FromProto(mapCtx, in.GetMedium())
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}

func InputDataConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IssueModel_InputDataConfigObservedState) *pb.IssueModel_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel_InputDataConfig{}
	out.TrainingConversationsCount = direct.ValueOf(in.TrainingConversationsCount)
	return out
}

func InputDataConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel_InputDataConfig) *krm.IssueModel_InputDataConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IssueModel_InputDataConfigObservedState{}
	out.TrainingConversationsCount = direct.LazyPtr(in.GetTrainingConversationsCount())
	return out
}

func IssueModelLabelStats_ToProto(mapCtx *direct.MapContext, in *krm.IssueModelLabelStats) *pb.IssueModelLabelStats {
	if in == nil {
		return nil
	}
	out := &pb.IssueModelLabelStats{}
	out.AnalyzedConversationsCount = direct.ValueOf(in.AnalyzedConversationsCount)
	out.UnclassifiedConversationsCount = direct.ValueOf(in.UnclassifiedConversationsCount)
	return out
}

func IssueModelLabelStats_FromProto(mapCtx *direct.MapContext, in *pb.IssueModelLabelStats) *krm.IssueModelLabelStats {
	if in == nil {
		return nil
	}
	out := &krm.IssueModelLabelStats{}
	out.AnalyzedConversationsCount = direct.LazyPtr(in.GetAnalyzedConversationsCount())
	out.UnclassifiedConversationsCount = direct.LazyPtr(in.GetUnclassifiedConversationsCount())
	return out
}
