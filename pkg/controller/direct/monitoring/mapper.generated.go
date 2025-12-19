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
// krm.group: monitoring.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.monitoring.v3

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlertPolicy_AlertStrategy_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_AlertStrategy) *krm.AlertPolicy_AlertStrategy {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_AlertStrategy{}
	out.NotificationRateLimit = AlertPolicy_AlertStrategy_NotificationRateLimit_FromProto(mapCtx, in.GetNotificationRateLimit())
	out.NotificationPrompts = direct.EnumSlice_FromProto(mapCtx, in.NotificationPrompts)
	out.AutoClose = direct.StringDuration_FromProto(mapCtx, in.GetAutoClose())
	out.NotificationChannelStrategy = direct.Slice_FromProto(mapCtx, in.NotificationChannelStrategy, AlertPolicy_AlertStrategy_NotificationChannelStrategy_FromProto)
	return out
}
func AlertPolicy_AlertStrategy_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_AlertStrategy) *pb.AlertPolicy_AlertStrategy {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_AlertStrategy{}
	out.NotificationRateLimit = AlertPolicy_AlertStrategy_NotificationRateLimit_ToProto(mapCtx, in.NotificationRateLimit)
	out.NotificationPrompts = direct.EnumSlice_ToProto[pb.AlertPolicy_AlertStrategy_NotificationPrompt](mapCtx, in.NotificationPrompts)
	out.AutoClose = direct.StringDuration_ToProto(mapCtx, in.AutoClose)
	out.NotificationChannelStrategy = direct.Slice_ToProto(mapCtx, in.NotificationChannelStrategy, AlertPolicy_AlertStrategy_NotificationChannelStrategy_ToProto)
	return out
}
func AlertPolicy_AlertStrategy_NotificationChannelStrategy_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_AlertStrategy_NotificationChannelStrategy) *krm.AlertPolicy_AlertStrategy_NotificationChannelStrategy {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_AlertStrategy_NotificationChannelStrategy{}
	out.NotificationChannelNames = in.NotificationChannelNames
	out.RenotifyInterval = direct.StringDuration_FromProto(mapCtx, in.GetRenotifyInterval())
	return out
}
func AlertPolicy_AlertStrategy_NotificationChannelStrategy_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_AlertStrategy_NotificationChannelStrategy) *pb.AlertPolicy_AlertStrategy_NotificationChannelStrategy {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_AlertStrategy_NotificationChannelStrategy{}
	out.NotificationChannelNames = in.NotificationChannelNames
	out.RenotifyInterval = direct.StringDuration_ToProto(mapCtx, in.RenotifyInterval)
	return out
}
func AlertPolicy_AlertStrategy_NotificationRateLimit_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_AlertStrategy_NotificationRateLimit) *krm.AlertPolicy_AlertStrategy_NotificationRateLimit {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_AlertStrategy_NotificationRateLimit{}
	out.Period = direct.StringDuration_FromProto(mapCtx, in.GetPeriod())
	return out
}
func AlertPolicy_AlertStrategy_NotificationRateLimit_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_AlertStrategy_NotificationRateLimit) *pb.AlertPolicy_AlertStrategy_NotificationRateLimit {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_AlertStrategy_NotificationRateLimit{}
	out.Period = direct.StringDuration_ToProto(mapCtx, in.Period)
	return out
}
func AlertPolicy_Condition_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition) *krm.AlertPolicy_Condition {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ConditionThreshold = AlertPolicy_Condition_MetricThreshold_FromProto(mapCtx, in.GetConditionThreshold())
	out.ConditionAbsent = AlertPolicy_Condition_MetricAbsence_FromProto(mapCtx, in.GetConditionAbsent())
	out.ConditionMatchedLog = AlertPolicy_Condition_LogMatch_FromProto(mapCtx, in.GetConditionMatchedLog())
	out.ConditionMonitoringQueryLanguage = AlertPolicy_Condition_MonitoringQueryLanguageCondition_FromProto(mapCtx, in.GetConditionMonitoringQueryLanguage())
	out.ConditionPrometheusQueryLanguage = AlertPolicy_Condition_PrometheusQueryLanguageCondition_FromProto(mapCtx, in.GetConditionPrometheusQueryLanguage())
	out.ConditionSQL = AlertPolicy_Condition_SQLCondition_FromProto(mapCtx, in.GetConditionSql())
	return out
}
func AlertPolicy_Condition_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition) *pb.AlertPolicy_Condition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := AlertPolicy_Condition_MetricThreshold_ToProto(mapCtx, in.ConditionThreshold); oneof != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionThreshold{ConditionThreshold: oneof}
	}
	if oneof := AlertPolicy_Condition_MetricAbsence_ToProto(mapCtx, in.ConditionAbsent); oneof != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionAbsent{ConditionAbsent: oneof}
	}
	if oneof := AlertPolicy_Condition_LogMatch_ToProto(mapCtx, in.ConditionMatchedLog); oneof != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionMatchedLog{ConditionMatchedLog: oneof}
	}
	if oneof := AlertPolicy_Condition_MonitoringQueryLanguageCondition_ToProto(mapCtx, in.ConditionMonitoringQueryLanguage); oneof != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionMonitoringQueryLanguage{ConditionMonitoringQueryLanguage: oneof}
	}
	if oneof := AlertPolicy_Condition_PrometheusQueryLanguageCondition_ToProto(mapCtx, in.ConditionPrometheusQueryLanguage); oneof != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionPrometheusQueryLanguage{ConditionPrometheusQueryLanguage: oneof}
	}
	if oneof := AlertPolicy_Condition_SQLCondition_ToProto(mapCtx, in.ConditionSQL); oneof != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionSql{ConditionSql: oneof}
	}
	return out
}
func AlertPolicy_Condition_LogMatch_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_LogMatch) *krm.AlertPolicy_Condition_LogMatch {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_LogMatch{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.LabelExtractors = in.LabelExtractors
	return out
}
func AlertPolicy_Condition_LogMatch_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_LogMatch) *pb.AlertPolicy_Condition_LogMatch {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_LogMatch{}
	out.Filter = direct.ValueOf(in.Filter)
	out.LabelExtractors = in.LabelExtractors
	return out
}
func AlertPolicy_Condition_MetricAbsence_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MetricAbsence) *krm.AlertPolicy_Condition_MetricAbsence {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_MetricAbsence{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Aggregations = direct.Slice_FromProto(mapCtx, in.Aggregations, Aggregation_FromProto)
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.Trigger = AlertPolicy_Condition_Trigger_FromProto(mapCtx, in.GetTrigger())
	return out
}
func AlertPolicy_Condition_MetricAbsence_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_MetricAbsence) *pb.AlertPolicy_Condition_MetricAbsence {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MetricAbsence{}
	out.Filter = direct.ValueOf(in.Filter)
	out.Aggregations = direct.Slice_ToProto(mapCtx, in.Aggregations, Aggregation_ToProto)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.Trigger = AlertPolicy_Condition_Trigger_ToProto(mapCtx, in.Trigger)
	return out
}
func AlertPolicy_Condition_MetricThreshold_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MetricThreshold) *krm.AlertPolicy_Condition_MetricThreshold {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_MetricThreshold{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Aggregations = direct.Slice_FromProto(mapCtx, in.Aggregations, Aggregation_FromProto)
	out.DenominatorFilter = direct.LazyPtr(in.GetDenominatorFilter())
	out.DenominatorAggregations = direct.Slice_FromProto(mapCtx, in.DenominatorAggregations, Aggregation_FromProto)
	out.ForecastOptions = AlertPolicy_Condition_MetricThreshold_ForecastOptions_FromProto(mapCtx, in.GetForecastOptions())
	out.Comparison = direct.Enum_FromProto(mapCtx, in.GetComparison())
	out.ThresholdValue = direct.LazyPtr(in.GetThresholdValue())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.Trigger = AlertPolicy_Condition_Trigger_FromProto(mapCtx, in.GetTrigger())
	out.EvaluationMissingData = direct.Enum_FromProto(mapCtx, in.GetEvaluationMissingData())
	return out
}
func AlertPolicy_Condition_MetricThreshold_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_MetricThreshold) *pb.AlertPolicy_Condition_MetricThreshold {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MetricThreshold{}
	out.Filter = direct.ValueOf(in.Filter)
	out.Aggregations = direct.Slice_ToProto(mapCtx, in.Aggregations, Aggregation_ToProto)
	out.DenominatorFilter = direct.ValueOf(in.DenominatorFilter)
	out.DenominatorAggregations = direct.Slice_ToProto(mapCtx, in.DenominatorAggregations, Aggregation_ToProto)
	out.ForecastOptions = AlertPolicy_Condition_MetricThreshold_ForecastOptions_ToProto(mapCtx, in.ForecastOptions)
	out.Comparison = direct.Enum_ToProto[pb.ComparisonType](mapCtx, in.Comparison)
	out.ThresholdValue = direct.ValueOf(in.ThresholdValue)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.Trigger = AlertPolicy_Condition_Trigger_ToProto(mapCtx, in.Trigger)
	out.EvaluationMissingData = direct.Enum_ToProto[pb.AlertPolicy_Condition_EvaluationMissingData](mapCtx, in.EvaluationMissingData)
	return out
}
func AlertPolicy_Condition_MetricThreshold_ForecastOptions_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MetricThreshold_ForecastOptions) *krm.AlertPolicy_Condition_MetricThreshold_ForecastOptions {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_MetricThreshold_ForecastOptions{}
	out.ForecastHorizon = direct.StringDuration_FromProto(mapCtx, in.GetForecastHorizon())
	return out
}
func AlertPolicy_Condition_MetricThreshold_ForecastOptions_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_MetricThreshold_ForecastOptions) *pb.AlertPolicy_Condition_MetricThreshold_ForecastOptions {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MetricThreshold_ForecastOptions{}
	out.ForecastHorizon = direct.StringDuration_ToProto(mapCtx, in.ForecastHorizon)
	return out
}
func AlertPolicy_Condition_MonitoringQueryLanguageCondition_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MonitoringQueryLanguageCondition) *krm.AlertPolicy_Condition_MonitoringQueryLanguageCondition {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_MonitoringQueryLanguageCondition{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.Trigger = AlertPolicy_Condition_Trigger_FromProto(mapCtx, in.GetTrigger())
	out.EvaluationMissingData = direct.Enum_FromProto(mapCtx, in.GetEvaluationMissingData())
	return out
}
func AlertPolicy_Condition_MonitoringQueryLanguageCondition_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_MonitoringQueryLanguageCondition) *pb.AlertPolicy_Condition_MonitoringQueryLanguageCondition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MonitoringQueryLanguageCondition{}
	out.Query = direct.ValueOf(in.Query)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.Trigger = AlertPolicy_Condition_Trigger_ToProto(mapCtx, in.Trigger)
	out.EvaluationMissingData = direct.Enum_ToProto[pb.AlertPolicy_Condition_EvaluationMissingData](mapCtx, in.EvaluationMissingData)
	return out
}
func AlertPolicy_Condition_PrometheusQueryLanguageCondition_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_PrometheusQueryLanguageCondition) *krm.AlertPolicy_Condition_PrometheusQueryLanguageCondition {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_PrometheusQueryLanguageCondition{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.EvaluationInterval = direct.StringDuration_FromProto(mapCtx, in.GetEvaluationInterval())
	out.Labels = in.Labels
	out.RuleGroup = direct.LazyPtr(in.GetRuleGroup())
	out.AlertRule = direct.LazyPtr(in.GetAlertRule())
	out.DisableMetricValidation = direct.LazyPtr(in.GetDisableMetricValidation())
	return out
}
func AlertPolicy_Condition_PrometheusQueryLanguageCondition_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_PrometheusQueryLanguageCondition) *pb.AlertPolicy_Condition_PrometheusQueryLanguageCondition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_PrometheusQueryLanguageCondition{}
	out.Query = direct.ValueOf(in.Query)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.EvaluationInterval = direct.StringDuration_ToProto(mapCtx, in.EvaluationInterval)
	out.Labels = in.Labels
	out.RuleGroup = direct.ValueOf(in.RuleGroup)
	out.AlertRule = direct.ValueOf(in.AlertRule)
	out.DisableMetricValidation = direct.ValueOf(in.DisableMetricValidation)
	return out
}
func AlertPolicy_Condition_SQLCondition_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition) *krm.AlertPolicy_Condition_SQLCondition {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_SQLCondition{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.Minutes = AlertPolicy_Condition_SQLCondition_Minutes_FromProto(mapCtx, in.GetMinutes())
	out.Hourly = AlertPolicy_Condition_SQLCondition_Hourly_FromProto(mapCtx, in.GetHourly())
	out.Daily = AlertPolicy_Condition_SQLCondition_Daily_FromProto(mapCtx, in.GetDaily())
	out.RowCountTest = AlertPolicy_Condition_SQLCondition_RowCountTest_FromProto(mapCtx, in.GetRowCountTest())
	out.BooleanTest = AlertPolicy_Condition_SQLCondition_BooleanTest_FromProto(mapCtx, in.GetBooleanTest())
	return out
}
func AlertPolicy_Condition_SQLCondition_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_SQLCondition) *pb.AlertPolicy_Condition_SqlCondition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition{}
	out.Query = direct.ValueOf(in.Query)
	if oneof := AlertPolicy_Condition_SQLCondition_Minutes_ToProto(mapCtx, in.Minutes); oneof != nil {
		out.Schedule = &pb.AlertPolicy_Condition_SqlCondition_Minutes_{Minutes: oneof}
	}
	if oneof := AlertPolicy_Condition_SQLCondition_Hourly_ToProto(mapCtx, in.Hourly); oneof != nil {
		out.Schedule = &pb.AlertPolicy_Condition_SqlCondition_Hourly_{Hourly: oneof}
	}
	if oneof := AlertPolicy_Condition_SQLCondition_Daily_ToProto(mapCtx, in.Daily); oneof != nil {
		out.Schedule = &pb.AlertPolicy_Condition_SqlCondition_Daily_{Daily: oneof}
	}
	if oneof := AlertPolicy_Condition_SQLCondition_RowCountTest_ToProto(mapCtx, in.RowCountTest); oneof != nil {
		out.Evaluate = &pb.AlertPolicy_Condition_SqlCondition_RowCountTest_{RowCountTest: oneof}
	}
	if oneof := AlertPolicy_Condition_SQLCondition_BooleanTest_ToProto(mapCtx, in.BooleanTest); oneof != nil {
		out.Evaluate = &pb.AlertPolicy_Condition_SqlCondition_BooleanTest_{BooleanTest: oneof}
	}
	return out
}
func AlertPolicy_Condition_SQLCondition_BooleanTest_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_BooleanTest) *krm.AlertPolicy_Condition_SQLCondition_BooleanTest {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_SQLCondition_BooleanTest{}
	out.Column = direct.LazyPtr(in.GetColumn())
	return out
}
func AlertPolicy_Condition_SQLCondition_BooleanTest_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_SQLCondition_BooleanTest) *pb.AlertPolicy_Condition_SqlCondition_BooleanTest {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_BooleanTest{}
	out.Column = direct.ValueOf(in.Column)
	return out
}
func AlertPolicy_Condition_SQLCondition_Daily_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_Daily) *krm.AlertPolicy_Condition_SQLCondition_Daily {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_SQLCondition_Daily{}
	out.Periodicity = direct.LazyPtr(in.GetPeriodicity())
	out.ExecutionTime = TimeOfDay_FromProto(mapCtx, in.GetExecutionTime())
	return out
}
func AlertPolicy_Condition_SQLCondition_Daily_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_SQLCondition_Daily) *pb.AlertPolicy_Condition_SqlCondition_Daily {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_Daily{}
	out.Periodicity = direct.ValueOf(in.Periodicity)
	out.ExecutionTime = TimeOfDay_ToProto(mapCtx, in.ExecutionTime)
	return out
}
func AlertPolicy_Condition_SQLCondition_Hourly_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_Hourly) *krm.AlertPolicy_Condition_SQLCondition_Hourly {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_SQLCondition_Hourly{}
	out.Periodicity = direct.LazyPtr(in.GetPeriodicity())
	out.MinuteOffset = in.MinuteOffset
	return out
}
func AlertPolicy_Condition_SQLCondition_Hourly_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_SQLCondition_Hourly) *pb.AlertPolicy_Condition_SqlCondition_Hourly {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_Hourly{}
	out.Periodicity = direct.ValueOf(in.Periodicity)
	out.MinuteOffset = in.MinuteOffset
	return out
}
func AlertPolicy_Condition_SQLCondition_Minutes_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_Minutes) *krm.AlertPolicy_Condition_SQLCondition_Minutes {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_SQLCondition_Minutes{}
	out.Periodicity = direct.LazyPtr(in.GetPeriodicity())
	return out
}
func AlertPolicy_Condition_SQLCondition_Minutes_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_SQLCondition_Minutes) *pb.AlertPolicy_Condition_SqlCondition_Minutes {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_Minutes{}
	out.Periodicity = direct.ValueOf(in.Periodicity)
	return out
}
func AlertPolicy_Condition_SQLCondition_RowCountTest_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_RowCountTest) *krm.AlertPolicy_Condition_SQLCondition_RowCountTest {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_SQLCondition_RowCountTest{}
	out.Comparison = direct.Enum_FromProto(mapCtx, in.GetComparison())
	out.Threshold = direct.LazyPtr(in.GetThreshold())
	return out
}
func AlertPolicy_Condition_SQLCondition_RowCountTest_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_SQLCondition_RowCountTest) *pb.AlertPolicy_Condition_SqlCondition_RowCountTest {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_RowCountTest{}
	out.Comparison = direct.Enum_ToProto[pb.ComparisonType](mapCtx, in.Comparison)
	out.Threshold = direct.ValueOf(in.Threshold)
	return out
}
func AlertPolicy_Condition_Trigger_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_Trigger) *krm.AlertPolicy_Condition_Trigger {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Condition_Trigger{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.Percent = direct.LazyPtr(in.GetPercent())
	return out
}
func AlertPolicy_Condition_Trigger_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Condition_Trigger) *pb.AlertPolicy_Condition_Trigger {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_Trigger{}
	if oneof := AlertPolicy_Condition_Trigger_Count_ToProto(mapCtx, in.Count); oneof != nil {
		out.Type = oneof
	}
	if oneof := AlertPolicy_Condition_Trigger_Percent_ToProto(mapCtx, in.Percent); oneof != nil {
		out.Type = oneof
	}
	return out
}
func AlertPolicy_Condition_Trigger_Count_ToProto(mapCtx *direct.MapContext, in *int32) *pb.AlertPolicy_Condition_Trigger_Count {
	if in == nil {
		return nil
	}
	return &pb.AlertPolicy_Condition_Trigger_Count{Count: *in}
}
func AlertPolicy_Condition_Trigger_Percent_ToProto(mapCtx *direct.MapContext, in *float64) *pb.AlertPolicy_Condition_Trigger_Percent {
	if in == nil {
		return nil
	}
	return &pb.AlertPolicy_Condition_Trigger_Percent{Percent: *in}
}
func AlertPolicy_Documentation_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Documentation) *krm.AlertPolicy_Documentation {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Documentation{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.Subject = direct.LazyPtr(in.GetSubject())
	out.Links = direct.Slice_FromProto(mapCtx, in.Links, AlertPolicy_Documentation_Link_FromProto)
	return out
}
func AlertPolicy_Documentation_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Documentation) *pb.AlertPolicy_Documentation {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Documentation{}
	out.Content = direct.ValueOf(in.Content)
	out.MimeType = direct.ValueOf(in.MimeType)
	out.Subject = direct.ValueOf(in.Subject)
	out.Links = direct.Slice_ToProto(mapCtx, in.Links, AlertPolicy_Documentation_Link_ToProto)
	return out
}
func AlertPolicy_Documentation_Link_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Documentation_Link) *krm.AlertPolicy_Documentation_Link {
	if in == nil {
		return nil
	}
	out := &krm.AlertPolicy_Documentation_Link{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URL = direct.LazyPtr(in.GetUrl())
	return out
}
func AlertPolicy_Documentation_Link_ToProto(mapCtx *direct.MapContext, in *krm.AlertPolicy_Documentation_Link) *pb.AlertPolicy_Documentation_Link {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Documentation_Link{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Url = direct.ValueOf(in.URL)
	return out
}
func MonitoringAlertPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy) *krm.MonitoringAlertPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringAlertPolicyObservedState{}
	// MISSING: Name
	// MISSING: UserLabels
	// MISSING: Validity
	// MISSING: CreationRecord
	// MISSING: MutationRecord
	// MISSING: Severity
	return out
}
func MonitoringAlertPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringAlertPolicyObservedState) *pb.AlertPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy{}
	// MISSING: Name
	// MISSING: UserLabels
	// MISSING: Validity
	// MISSING: CreationRecord
	// MISSING: MutationRecord
	// MISSING: Severity
	return out
}
func MonitoringAlertPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy) *krm.MonitoringAlertPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringAlertPolicySpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Documentation = AlertPolicy_Documentation_FromProto(mapCtx, in.GetDocumentation())
	// MISSING: UserLabels
	out.Conditions = direct.Slice_FromProto(mapCtx, in.Conditions, AlertPolicy_Condition_FromProto)
	out.Combiner = direct.Enum_FromProto(mapCtx, in.GetCombiner())
	out.Enabled = direct.BoolValue_FromProto(mapCtx, in.GetEnabled())
	// MISSING: Validity
	out.NotificationChannels = in.NotificationChannels
	// MISSING: CreationRecord
	// MISSING: MutationRecord
	out.AlertStrategy = AlertPolicy_AlertStrategy_FromProto(mapCtx, in.GetAlertStrategy())
	// MISSING: Severity
	return out
}
func MonitoringAlertPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringAlertPolicySpec) *pb.AlertPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Documentation = AlertPolicy_Documentation_ToProto(mapCtx, in.Documentation)
	// MISSING: UserLabels
	out.Conditions = direct.Slice_ToProto(mapCtx, in.Conditions, AlertPolicy_Condition_ToProto)
	out.Combiner = direct.Enum_ToProto[pb.AlertPolicy_ConditionCombinerType](mapCtx, in.Combiner)
	out.Enabled = direct.BoolValue_ToProto(mapCtx, in.Enabled)
	// MISSING: Validity
	out.NotificationChannels = in.NotificationChannels
	// MISSING: CreationRecord
	// MISSING: MutationRecord
	out.AlertStrategy = AlertPolicy_AlertStrategy_ToProto(mapCtx, in.AlertStrategy)
	// MISSING: Severity
	return out
}
func MonitoringNotificationChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelSpec{}
	out.Type = direct.LazyPtr(in.GetType())
	// MISSING: Name
	// MISSING: DisplayName
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: UserLabels
	// MISSING: VerificationStatus
	out.Enabled = direct.BoolValue_FromProto(mapCtx, in.GetEnabled())
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelSpec) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	out.Type = direct.ValueOf(in.Type)
	// MISSING: Name
	// MISSING: DisplayName
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	// MISSING: UserLabels
	// MISSING: VerificationStatus
	out.Enabled = direct.BoolValue_ToProto(mapCtx, in.Enabled)
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelStatus_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelStatus{}
	// MISSING: Type
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	out.VerificationStatus = direct.Enum_FromProto(mapCtx, in.GetVerificationStatus())
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelStatus) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	// MISSING: Type
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	out.VerificationStatus = direct.Enum_ToProto[pb.NotificationChannel_VerificationStatus](mapCtx, in.VerificationStatus)
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MutationRecord_FromProto(mapCtx *direct.MapContext, in *pb.MutationRecord) *krm.MutationRecord {
	if in == nil {
		return nil
	}
	out := &krm.MutationRecord{}
	out.MutateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMutateTime())
	out.MutatedBy = direct.LazyPtr(in.GetMutatedBy())
	return out
}
func MutationRecord_ToProto(mapCtx *direct.MapContext, in *krm.MutationRecord) *pb.MutationRecord {
	if in == nil {
		return nil
	}
	out := &pb.MutationRecord{}
	out.MutateTime = direct.StringTimestamp_ToProto(mapCtx, in.MutateTime)
	out.MutatedBy = direct.ValueOf(in.MutatedBy)
	return out
}
