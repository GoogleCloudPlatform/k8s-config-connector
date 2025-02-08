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

package contactcenterinsights

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ContactcenterinsightsIssueModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel) *krm.ContactcenterinsightsIssueModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsIssueModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IssueCount
	// MISSING: State
	// MISSING: InputDataConfig
	// MISSING: TrainingStats
	// MISSING: ModelType
	// MISSING: LanguageCode
	return out
}
func ContactcenterinsightsIssueModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsIssueModelObservedState) *pb.IssueModel {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IssueCount
	// MISSING: State
	// MISSING: InputDataConfig
	// MISSING: TrainingStats
	// MISSING: ModelType
	// MISSING: LanguageCode
	return out
}
func ContactcenterinsightsIssueModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel) *krm.ContactcenterinsightsIssueModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsIssueModelSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IssueCount
	// MISSING: State
	// MISSING: InputDataConfig
	// MISSING: TrainingStats
	// MISSING: ModelType
	// MISSING: LanguageCode
	return out
}
func ContactcenterinsightsIssueModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsIssueModelSpec) *pb.IssueModel {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IssueCount
	// MISSING: State
	// MISSING: InputDataConfig
	// MISSING: TrainingStats
	// MISSING: ModelType
	// MISSING: LanguageCode
	return out
}
func IssueModel_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel) *krm.IssueModel {
	if in == nil {
		return nil
	}
	out := &krm.IssueModel{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IssueCount
	// MISSING: State
	out.InputDataConfig = IssueModel_InputDataConfig_FromProto(mapCtx, in.GetInputDataConfig())
	// MISSING: TrainingStats
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func IssueModel_ToProto(mapCtx *direct.MapContext, in *krm.IssueModel) *pb.IssueModel {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IssueCount
	// MISSING: State
	out.InputDataConfig = IssueModel_InputDataConfig_ToProto(mapCtx, in.InputDataConfig)
	// MISSING: TrainingStats
	out.ModelType = direct.Enum_ToProto[pb.IssueModel_ModelType](mapCtx, in.ModelType)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
func IssueModelLabelStats_FromProto(mapCtx *direct.MapContext, in *pb.IssueModelLabelStats) *krm.IssueModelLabelStats {
	if in == nil {
		return nil
	}
	out := &krm.IssueModelLabelStats{}
	out.AnalyzedConversationsCount = direct.LazyPtr(in.GetAnalyzedConversationsCount())
	out.UnclassifiedConversationsCount = direct.LazyPtr(in.GetUnclassifiedConversationsCount())
	// MISSING: IssueStats
	return out
}
func IssueModelLabelStats_ToProto(mapCtx *direct.MapContext, in *krm.IssueModelLabelStats) *pb.IssueModelLabelStats {
	if in == nil {
		return nil
	}
	out := &pb.IssueModelLabelStats{}
	out.AnalyzedConversationsCount = direct.ValueOf(in.AnalyzedConversationsCount)
	out.UnclassifiedConversationsCount = direct.ValueOf(in.UnclassifiedConversationsCount)
	// MISSING: IssueStats
	return out
}
func IssueModelLabelStats_IssueStats_FromProto(mapCtx *direct.MapContext, in *pb.IssueModelLabelStats_IssueStats) *krm.IssueModelLabelStats_IssueStats {
	if in == nil {
		return nil
	}
	out := &krm.IssueModelLabelStats_IssueStats{}
	out.Issue = direct.LazyPtr(in.GetIssue())
	out.LabeledConversationsCount = direct.LazyPtr(in.GetLabeledConversationsCount())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func IssueModelLabelStats_IssueStats_ToProto(mapCtx *direct.MapContext, in *krm.IssueModelLabelStats_IssueStats) *pb.IssueModelLabelStats_IssueStats {
	if in == nil {
		return nil
	}
	out := &pb.IssueModelLabelStats_IssueStats{}
	out.Issue = direct.ValueOf(in.Issue)
	out.LabeledConversationsCount = direct.ValueOf(in.LabeledConversationsCount)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func IssueModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel) *krm.IssueModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IssueModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.IssueCount = direct.LazyPtr(in.GetIssueCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.InputDataConfig = IssueModel_InputDataConfigObservedState_FromProto(mapCtx, in.GetInputDataConfig())
	out.TrainingStats = IssueModelLabelStats_FromProto(mapCtx, in.GetTrainingStats())
	// MISSING: ModelType
	// MISSING: LanguageCode
	return out
}
func IssueModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IssueModelObservedState) *pb.IssueModel {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.IssueCount = direct.ValueOf(in.IssueCount)
	out.State = direct.Enum_ToProto[pb.IssueModel_State](mapCtx, in.State)
	out.InputDataConfig = IssueModel_InputDataConfigObservedState_ToProto(mapCtx, in.InputDataConfig)
	out.TrainingStats = IssueModelLabelStats_ToProto(mapCtx, in.TrainingStats)
	// MISSING: ModelType
	// MISSING: LanguageCode
	return out
}
func IssueModel_InputDataConfig_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel_InputDataConfig) *krm.IssueModel_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &krm.IssueModel_InputDataConfig{}
	out.Medium = direct.Enum_FromProto(mapCtx, in.GetMedium())
	// MISSING: TrainingConversationsCount
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}
func IssueModel_InputDataConfig_ToProto(mapCtx *direct.MapContext, in *krm.IssueModel_InputDataConfig) *pb.IssueModel_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel_InputDataConfig{}
	out.Medium = direct.Enum_ToProto[pb.Conversation_Medium](mapCtx, in.Medium)
	// MISSING: TrainingConversationsCount
	out.Filter = direct.ValueOf(in.Filter)
	return out
}
func IssueModel_InputDataConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IssueModel_InputDataConfig) *krm.IssueModel_InputDataConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IssueModel_InputDataConfigObservedState{}
	// MISSING: Medium
	out.TrainingConversationsCount = direct.LazyPtr(in.GetTrainingConversationsCount())
	// MISSING: Filter
	return out
}
func IssueModel_InputDataConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IssueModel_InputDataConfigObservedState) *pb.IssueModel_InputDataConfig {
	if in == nil {
		return nil
	}
	out := &pb.IssueModel_InputDataConfig{}
	// MISSING: Medium
	out.TrainingConversationsCount = direct.ValueOf(in.TrainingConversationsCount)
	// MISSING: Filter
	return out
}
