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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Experiment_FromProto(mapCtx *direct.MapContext, in *pb.Experiment) *krm.Experiment {
	if in == nil {
		return nil
	}
	out := &krm.Experiment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Definition = Experiment_Definition_FromProto(mapCtx, in.GetDefinition())
	out.RolloutConfig = RolloutConfig_FromProto(mapCtx, in.GetRolloutConfig())
	out.RolloutState = RolloutState_FromProto(mapCtx, in.GetRolloutState())
	out.RolloutFailureReason = direct.LazyPtr(in.GetRolloutFailureReason())
	out.Result = Experiment_Result_FromProto(mapCtx, in.GetResult())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.ExperimentLength = direct.StringDuration_FromProto(mapCtx, in.GetExperimentLength())
	out.VariantsHistory = direct.Slice_FromProto(mapCtx, in.VariantsHistory, VariantsHistory_FromProto)
	return out
}
func Experiment_ToProto(mapCtx *direct.MapContext, in *krm.Experiment) *pb.Experiment {
	if in == nil {
		return nil
	}
	out := &pb.Experiment{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.State = direct.Enum_ToProto[pb.Experiment_State](mapCtx, in.State)
	out.Definition = Experiment_Definition_ToProto(mapCtx, in.Definition)
	out.RolloutConfig = RolloutConfig_ToProto(mapCtx, in.RolloutConfig)
	out.RolloutState = RolloutState_ToProto(mapCtx, in.RolloutState)
	out.RolloutFailureReason = direct.ValueOf(in.RolloutFailureReason)
	out.Result = Experiment_Result_ToProto(mapCtx, in.Result)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	out.ExperimentLength = direct.StringDuration_ToProto(mapCtx, in.ExperimentLength)
	out.VariantsHistory = direct.Slice_ToProto(mapCtx, in.VariantsHistory, VariantsHistory_ToProto)
	return out
}
func Experiment_Definition_FromProto(mapCtx *direct.MapContext, in *pb.Experiment_Definition) *krm.Experiment_Definition {
	if in == nil {
		return nil
	}
	out := &krm.Experiment_Definition{}
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.VersionVariants = VersionVariants_FromProto(mapCtx, in.GetVersionVariants())
	return out
}
func Experiment_Definition_ToProto(mapCtx *direct.MapContext, in *krm.Experiment_Definition) *pb.Experiment_Definition {
	if in == nil {
		return nil
	}
	out := &pb.Experiment_Definition{}
	out.Condition = direct.ValueOf(in.Condition)
	if oneof := VersionVariants_ToProto(mapCtx, in.VersionVariants); oneof != nil {
		out.Variants = &pb.Experiment_Definition_VersionVariants{VersionVariants: oneof}
	}
	return out
}
func Experiment_Result_FromProto(mapCtx *direct.MapContext, in *pb.Experiment_Result) *krm.Experiment_Result {
	if in == nil {
		return nil
	}
	out := &krm.Experiment_Result{}
	out.VersionMetrics = direct.Slice_FromProto(mapCtx, in.VersionMetrics, Experiment_Result_VersionMetrics_FromProto)
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	return out
}
func Experiment_Result_ToProto(mapCtx *direct.MapContext, in *krm.Experiment_Result) *pb.Experiment_Result {
	if in == nil {
		return nil
	}
	out := &pb.Experiment_Result{}
	out.VersionMetrics = direct.Slice_ToProto(mapCtx, in.VersionMetrics, Experiment_Result_VersionMetrics_ToProto)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	return out
}
func Experiment_Result_ConfidenceInterval_FromProto(mapCtx *direct.MapContext, in *pb.Experiment_Result_ConfidenceInterval) *krm.Experiment_Result_ConfidenceInterval {
	if in == nil {
		return nil
	}
	out := &krm.Experiment_Result_ConfidenceInterval{}
	out.ConfidenceLevel = direct.LazyPtr(in.GetConfidenceLevel())
	out.Ratio = direct.LazyPtr(in.GetRatio())
	out.LowerBound = direct.LazyPtr(in.GetLowerBound())
	out.UpperBound = direct.LazyPtr(in.GetUpperBound())
	return out
}
func Experiment_Result_ConfidenceInterval_ToProto(mapCtx *direct.MapContext, in *krm.Experiment_Result_ConfidenceInterval) *pb.Experiment_Result_ConfidenceInterval {
	if in == nil {
		return nil
	}
	out := &pb.Experiment_Result_ConfidenceInterval{}
	out.ConfidenceLevel = direct.ValueOf(in.ConfidenceLevel)
	out.Ratio = direct.ValueOf(in.Ratio)
	out.LowerBound = direct.ValueOf(in.LowerBound)
	out.UpperBound = direct.ValueOf(in.UpperBound)
	return out
}
func Experiment_Result_Metric_FromProto(mapCtx *direct.MapContext, in *pb.Experiment_Result_Metric) *krm.Experiment_Result_Metric {
	if in == nil {
		return nil
	}
	out := &krm.Experiment_Result_Metric{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.CountType = direct.Enum_FromProto(mapCtx, in.GetCountType())
	out.Ratio = direct.LazyPtr(in.GetRatio())
	out.Count = direct.LazyPtr(in.GetCount())
	out.ConfidenceInterval = Experiment_Result_ConfidenceInterval_FromProto(mapCtx, in.GetConfidenceInterval())
	return out
}
func Experiment_Result_Metric_ToProto(mapCtx *direct.MapContext, in *krm.Experiment_Result_Metric) *pb.Experiment_Result_Metric {
	if in == nil {
		return nil
	}
	out := &pb.Experiment_Result_Metric{}
	out.Type = direct.Enum_ToProto[pb.Experiment_Result_MetricType](mapCtx, in.Type)
	out.CountType = direct.Enum_ToProto[pb.Experiment_Result_CountType](mapCtx, in.CountType)
	if oneof := Experiment_Result_Metric_Ratio_ToProto(mapCtx, in.Ratio); oneof != nil {
		out.Value = oneof
	}
	if oneof := Experiment_Result_Metric_Count_ToProto(mapCtx, in.Count); oneof != nil {
		out.Value = oneof
	}
	out.ConfidenceInterval = Experiment_Result_ConfidenceInterval_ToProto(mapCtx, in.ConfidenceInterval)
	return out
}
func Experiment_Result_VersionMetrics_FromProto(mapCtx *direct.MapContext, in *pb.Experiment_Result_VersionMetrics) *krm.Experiment_Result_VersionMetrics {
	if in == nil {
		return nil
	}
	out := &krm.Experiment_Result_VersionMetrics{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, Experiment_Result_Metric_FromProto)
	out.SessionCount = direct.LazyPtr(in.GetSessionCount())
	return out
}
func Experiment_Result_VersionMetrics_ToProto(mapCtx *direct.MapContext, in *krm.Experiment_Result_VersionMetrics) *pb.Experiment_Result_VersionMetrics {
	if in == nil {
		return nil
	}
	out := &pb.Experiment_Result_VersionMetrics{}
	out.Version = direct.ValueOf(in.Version)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, Experiment_Result_Metric_ToProto)
	out.SessionCount = direct.ValueOf(in.SessionCount)
	return out
}
func RolloutConfig_FromProto(mapCtx *direct.MapContext, in *pb.RolloutConfig) *krm.RolloutConfig {
	if in == nil {
		return nil
	}
	out := &krm.RolloutConfig{}
	out.RolloutSteps = direct.Slice_FromProto(mapCtx, in.RolloutSteps, RolloutConfig_RolloutStep_FromProto)
	out.RolloutCondition = direct.LazyPtr(in.GetRolloutCondition())
	out.FailureCondition = direct.LazyPtr(in.GetFailureCondition())
	return out
}
func RolloutConfig_ToProto(mapCtx *direct.MapContext, in *krm.RolloutConfig) *pb.RolloutConfig {
	if in == nil {
		return nil
	}
	out := &pb.RolloutConfig{}
	out.RolloutSteps = direct.Slice_ToProto(mapCtx, in.RolloutSteps, RolloutConfig_RolloutStep_ToProto)
	out.RolloutCondition = direct.ValueOf(in.RolloutCondition)
	out.FailureCondition = direct.ValueOf(in.FailureCondition)
	return out
}
func RolloutConfig_RolloutStep_FromProto(mapCtx *direct.MapContext, in *pb.RolloutConfig_RolloutStep) *krm.RolloutConfig_RolloutStep {
	if in == nil {
		return nil
	}
	out := &krm.RolloutConfig_RolloutStep{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.TrafficPercent = direct.LazyPtr(in.GetTrafficPercent())
	out.MinDuration = direct.StringDuration_FromProto(mapCtx, in.GetMinDuration())
	return out
}
func RolloutConfig_RolloutStep_ToProto(mapCtx *direct.MapContext, in *krm.RolloutConfig_RolloutStep) *pb.RolloutConfig_RolloutStep {
	if in == nil {
		return nil
	}
	out := &pb.RolloutConfig_RolloutStep{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.TrafficPercent = direct.ValueOf(in.TrafficPercent)
	out.MinDuration = direct.StringDuration_ToProto(mapCtx, in.MinDuration)
	return out
}
func RolloutState_FromProto(mapCtx *direct.MapContext, in *pb.RolloutState) *krm.RolloutState {
	if in == nil {
		return nil
	}
	out := &krm.RolloutState{}
	out.Step = direct.LazyPtr(in.GetStep())
	out.StepIndex = direct.LazyPtr(in.GetStepIndex())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	return out
}
func RolloutState_ToProto(mapCtx *direct.MapContext, in *krm.RolloutState) *pb.RolloutState {
	if in == nil {
		return nil
	}
	out := &pb.RolloutState{}
	out.Step = direct.ValueOf(in.Step)
	out.StepIndex = direct.ValueOf(in.StepIndex)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	return out
}
func VariantsHistory_FromProto(mapCtx *direct.MapContext, in *pb.VariantsHistory) *krm.VariantsHistory {
	if in == nil {
		return nil
	}
	out := &krm.VariantsHistory{}
	out.VersionVariants = VersionVariants_FromProto(mapCtx, in.GetVersionVariants())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func VariantsHistory_ToProto(mapCtx *direct.MapContext, in *krm.VariantsHistory) *pb.VariantsHistory {
	if in == nil {
		return nil
	}
	out := &pb.VariantsHistory{}
	if oneof := VersionVariants_ToProto(mapCtx, in.VersionVariants); oneof != nil {
		out.Variants = &pb.VariantsHistory_VersionVariants{VersionVariants: oneof}
	}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func VersionVariants_FromProto(mapCtx *direct.MapContext, in *pb.VersionVariants) *krm.VersionVariants {
	if in == nil {
		return nil
	}
	out := &krm.VersionVariants{}
	out.Variants = direct.Slice_FromProto(mapCtx, in.Variants, VersionVariants_Variant_FromProto)
	return out
}
func VersionVariants_ToProto(mapCtx *direct.MapContext, in *krm.VersionVariants) *pb.VersionVariants {
	if in == nil {
		return nil
	}
	out := &pb.VersionVariants{}
	out.Variants = direct.Slice_ToProto(mapCtx, in.Variants, VersionVariants_Variant_ToProto)
	return out
}
func VersionVariants_Variant_FromProto(mapCtx *direct.MapContext, in *pb.VersionVariants_Variant) *krm.VersionVariants_Variant {
	if in == nil {
		return nil
	}
	out := &krm.VersionVariants_Variant{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.TrafficAllocation = direct.LazyPtr(in.GetTrafficAllocation())
	out.IsControlGroup = direct.LazyPtr(in.GetIsControlGroup())
	return out
}
func VersionVariants_Variant_ToProto(mapCtx *direct.MapContext, in *krm.VersionVariants_Variant) *pb.VersionVariants_Variant {
	if in == nil {
		return nil
	}
	out := &pb.VersionVariants_Variant{}
	out.Version = direct.ValueOf(in.Version)
	out.TrafficAllocation = direct.ValueOf(in.TrafficAllocation)
	out.IsControlGroup = direct.ValueOf(in.IsControlGroup)
	return out
}
