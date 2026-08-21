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

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/type/timeofday"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// MonitoringAlertPolicySpec_FromProto maps an AlertPolicy proto to a MonitoringAlertPolicySpec.
func MonitoringAlertPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy) *krm.MonitoringAlertPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringAlertPolicySpec{}
	out.AlertStrategy = AlertpolicyAlertStrategy_FromProto(mapCtx, in.GetAlertStrategy())
	if val := direct.Enum_FromProto(mapCtx, in.GetCombiner()); val != nil {
		out.Combiner = *val
	}
	out.Conditions = direct.Slice_FromProto(mapCtx, in.Conditions, AlertpolicyConditions_FromProto)
	out.DisplayName = in.GetDisplayName()
	out.Documentation = AlertpolicyDocumentation_FromProto(mapCtx, in.GetDocumentation())
	if in.Enabled != nil {
		out.Enabled = &in.Enabled.Value
	}
	if len(in.GetNotificationChannels()) > 0 {
		out.NotificationChannels = make([]krm.MonitoringNotificationChannelRef, len(in.GetNotificationChannels()))
		for i, nc := range in.GetNotificationChannels() {
			out.NotificationChannels[i] = krm.MonitoringNotificationChannelRef{External: nc}
		}
	}
	if in.GetSeverity() != pb.AlertPolicy_SEVERITY_UNSPECIFIED {
		out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	}
	return out
}

// MonitoringAlertPolicySpec_ToProto maps a MonitoringAlertPolicySpec KRM struct to an AlertPolicy proto.
func MonitoringAlertPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringAlertPolicySpec) *pb.AlertPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy{}
	out.AlertStrategy = AlertpolicyAlertStrategy_ToProto(mapCtx, in.AlertStrategy)
	out.Combiner = direct.Enum_ToProto[pb.AlertPolicy_ConditionCombinerType](mapCtx, &in.Combiner)
	out.Conditions = direct.Slice_ToProto(mapCtx, in.Conditions, AlertpolicyConditions_ToProto)
	out.DisplayName = in.DisplayName
	out.Documentation = AlertpolicyDocumentation_ToProto(mapCtx, in.Documentation)
	if in.Enabled != nil {
		out.Enabled = &wrapperspb.BoolValue{Value: *in.Enabled}
	}
	if len(in.NotificationChannels) > 0 {
		out.NotificationChannels = make([]string, len(in.NotificationChannels))
		for i, nc := range in.NotificationChannels {
			out.NotificationChannels[i] = nc.External
		}
	}
	if in.Severity != nil {
		out.Severity = direct.Enum_ToProto[pb.AlertPolicy_Severity](mapCtx, in.Severity)
	}
	return out
}

// MonitoringAlertPolicyStatus_FromProto maps an AlertPolicy proto to a MonitoringAlertPolicyStatus.
func MonitoringAlertPolicyStatus_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy) *krm.MonitoringAlertPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringAlertPolicyStatus{}
	if in.GetName() != "" {
		out.Name = direct.LazyPtr(in.GetName())
	}
	if in.GetCreationRecord() != nil {
		out.CreationRecord = []krm.AlertpolicyCreationRecordStatus{
			*AlertpolicyCreationRecordStatus_FromProto(mapCtx, in.GetCreationRecord()),
		}
	}
	return out
}

// MonitoringAlertPolicyStatus_ToProto maps a MonitoringAlertPolicyStatus KRM struct to an AlertPolicy proto.
func MonitoringAlertPolicyStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringAlertPolicyStatus) *pb.AlertPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy{}
	if in.Name != nil {
		out.Name = *in.Name
	}
	if len(in.CreationRecord) > 0 {
		out.CreationRecord = AlertpolicyCreationRecordStatus_ToProto(mapCtx, &in.CreationRecord[0])
	}
	return out
}

// AlertpolicyConditionAbsent is hand-coded because Duration is string in KRM but *durationpb.Duration in proto.
func AlertpolicyConditionAbsent_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MetricAbsence) *krm.AlertpolicyConditionAbsent {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyConditionAbsent{}
	out.Aggregations = direct.Slice_FromProto(mapCtx, in.Aggregations, AlertpolicyAggregations_FromProto)
	if d := direct.StringDuration_FromProto(mapCtx, in.GetDuration()); d != nil {
		out.Duration = *d
	}
	if in.GetFilter() != "" {
		out.Filter = direct.LazyPtr(in.GetFilter())
	}
	out.Trigger = AlertpolicyTrigger_FromProto(mapCtx, in.GetTrigger())
	return out
}

func AlertpolicyConditionAbsent_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyConditionAbsent) *pb.AlertPolicy_Condition_MetricAbsence {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MetricAbsence{}
	out.Aggregations = direct.Slice_ToProto(mapCtx, in.Aggregations, AlertpolicyAggregations_ToProto)
	if in.Duration != "" {
		out.Duration = direct.StringDuration_ToProto(mapCtx, &in.Duration)
	}
	if in.Filter != nil {
		out.Filter = *in.Filter
	}
	out.Trigger = AlertpolicyTrigger_ToProto(mapCtx, in.Trigger)
	return out
}

// AlertpolicyConditionMatchedLog is hand-coded because Filter is non-pointer string.
func AlertpolicyConditionMatchedLog_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_LogMatch) *krm.AlertpolicyConditionMatchedLog {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyConditionMatchedLog{}
	out.Filter = in.GetFilter()
	out.LabelExtractors = in.GetLabelExtractors()
	return out
}

func AlertpolicyConditionMatchedLog_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyConditionMatchedLog) *pb.AlertPolicy_Condition_LogMatch {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_LogMatch{}
	out.Filter = in.Filter
	out.LabelExtractors = in.LabelExtractors
	return out
}

// AlertpolicyConditionMonitoringQueryLanguage is hand-coded because Duration and Query are non-pointer strings.
func AlertpolicyConditionMonitoringQueryLanguage_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MonitoringQueryLanguageCondition) *krm.AlertpolicyConditionMonitoringQueryLanguage {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyConditionMonitoringQueryLanguage{}
	if d := direct.StringDuration_FromProto(mapCtx, in.GetDuration()); d != nil {
		out.Duration = *d
	}
	out.EvaluationMissingData = direct.Enum_FromProto(mapCtx, in.GetEvaluationMissingData())
	out.Query = in.GetQuery()
	out.Trigger = AlertpolicyTrigger_FromProto(mapCtx, in.GetTrigger())
	return out
}

func AlertpolicyConditionMonitoringQueryLanguage_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyConditionMonitoringQueryLanguage) *pb.AlertPolicy_Condition_MonitoringQueryLanguageCondition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MonitoringQueryLanguageCondition{}
	if in.Duration != "" {
		out.Duration = direct.StringDuration_ToProto(mapCtx, &in.Duration)
	}
	out.EvaluationMissingData = direct.Enum_ToProto[pb.AlertPolicy_Condition_EvaluationMissingData](mapCtx, in.EvaluationMissingData)
	out.Query = in.Query
	out.Trigger = AlertpolicyTrigger_ToProto(mapCtx, in.Trigger)
	return out
}

// AlertpolicyConditionPrometheusQueryLanguage is hand-coded because Query is non-pointer string.
func AlertpolicyConditionPrometheusQueryLanguage_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_PrometheusQueryLanguageCondition) *krm.AlertpolicyConditionPrometheusQueryLanguage {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyConditionPrometheusQueryLanguage{}
	if in.GetAlertRule() != "" {
		out.AlertRule = direct.LazyPtr(in.GetAlertRule())
	}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.EvaluationInterval = direct.StringDuration_FromProto(mapCtx, in.GetEvaluationInterval())
	out.Labels = in.GetLabels()
	out.Query = in.GetQuery()
	if in.GetRuleGroup() != "" {
		out.RuleGroup = direct.LazyPtr(in.GetRuleGroup())
	}
	return out
}

func AlertpolicyConditionPrometheusQueryLanguage_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyConditionPrometheusQueryLanguage) *pb.AlertPolicy_Condition_PrometheusQueryLanguageCondition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_PrometheusQueryLanguageCondition{}
	if in.AlertRule != nil {
		out.AlertRule = *in.AlertRule
	}
	if in.Duration != nil && *in.Duration != "" {
		out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	}
	if in.EvaluationInterval != nil && *in.EvaluationInterval != "" {
		out.EvaluationInterval = direct.StringDuration_ToProto(mapCtx, in.EvaluationInterval)
	}
	out.Labels = in.Labels
	out.Query = in.Query
	if in.RuleGroup != nil {
		out.RuleGroup = *in.RuleGroup
	}
	return out
}

// AlertpolicyConditionThreshold is hand-coded because Duration and Comparison are non-pointer types.
func AlertpolicyConditionThreshold_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MetricThreshold) *krm.AlertpolicyConditionThreshold {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyConditionThreshold{}
	out.Aggregations = direct.Slice_FromProto(mapCtx, in.Aggregations, AlertpolicyAggregations_FromProto)
	if val := direct.Enum_FromProto(mapCtx, in.GetComparison()); val != nil {
		out.Comparison = *val
	}
	out.DenominatorAggregations = direct.Slice_FromProto(mapCtx, in.DenominatorAggregations, AlertpolicyDenominatorAggregations_FromProto)
	if in.GetDenominatorFilter() != "" {
		out.DenominatorFilter = direct.LazyPtr(in.GetDenominatorFilter())
	}
	if d := direct.StringDuration_FromProto(mapCtx, in.GetDuration()); d != nil {
		out.Duration = *d
	}
	out.EvaluationMissingData = direct.Enum_FromProto(mapCtx, in.GetEvaluationMissingData())
	if in.GetFilter() != "" {
		out.Filter = direct.LazyPtr(in.GetFilter())
	}
	out.ForecastOptions = AlertpolicyForecastOptions_FromProto(mapCtx, in.GetForecastOptions())
	if in.ThresholdValue != 0 {
		out.ThresholdValue = direct.LazyPtr(in.GetThresholdValue())
	}
	out.Trigger = AlertpolicyTrigger_FromProto(mapCtx, in.GetTrigger())
	return out
}

func AlertpolicyConditionThreshold_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyConditionThreshold) *pb.AlertPolicy_Condition_MetricThreshold {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MetricThreshold{}
	out.Aggregations = direct.Slice_ToProto(mapCtx, in.Aggregations, AlertpolicyAggregations_ToProto)
	out.Comparison = direct.Enum_ToProto[pb.ComparisonType](mapCtx, &in.Comparison)
	out.DenominatorAggregations = direct.Slice_ToProto(mapCtx, in.DenominatorAggregations, AlertpolicyDenominatorAggregations_ToProto)
	if in.DenominatorFilter != nil {
		out.DenominatorFilter = *in.DenominatorFilter
	}
	if in.Duration != "" {
		out.Duration = direct.StringDuration_ToProto(mapCtx, &in.Duration)
	}
	out.EvaluationMissingData = direct.Enum_ToProto[pb.AlertPolicy_Condition_EvaluationMissingData](mapCtx, in.EvaluationMissingData)
	if in.Filter != nil {
		out.Filter = *in.Filter
	}
	out.ForecastOptions = AlertpolicyForecastOptions_ToProto(mapCtx, in.ForecastOptions)
	if in.ThresholdValue != nil {
		out.ThresholdValue = *in.ThresholdValue
	}
	out.Trigger = AlertpolicyTrigger_ToProto(mapCtx, in.Trigger)
	return out
}

// AlertpolicyConditions is hand-coded to map the Condition oneof and DisplayName non-pointer string.
func AlertpolicyConditions_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition) *krm.AlertpolicyConditions {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyConditions{}
	out.DisplayName = in.GetDisplayName()
	if in.GetName() != "" {
		out.Name = direct.LazyPtr(in.GetName())
	}
	switch c := in.Condition.(type) {
	case *pb.AlertPolicy_Condition_ConditionThreshold:
		out.ConditionThreshold = AlertpolicyConditionThreshold_FromProto(mapCtx, c.ConditionThreshold)
	case *pb.AlertPolicy_Condition_ConditionAbsent:
		out.ConditionAbsent = AlertpolicyConditionAbsent_FromProto(mapCtx, c.ConditionAbsent)
	case *pb.AlertPolicy_Condition_ConditionMatchedLog:
		out.ConditionMatchedLog = AlertpolicyConditionMatchedLog_FromProto(mapCtx, c.ConditionMatchedLog)
	case *pb.AlertPolicy_Condition_ConditionMonitoringQueryLanguage:
		out.ConditionMonitoringQueryLanguage = AlertpolicyConditionMonitoringQueryLanguage_FromProto(mapCtx, c.ConditionMonitoringQueryLanguage)
	case *pb.AlertPolicy_Condition_ConditionPrometheusQueryLanguage:
		out.ConditionPrometheusQueryLanguage = AlertpolicyConditionPrometheusQueryLanguage_FromProto(mapCtx, c.ConditionPrometheusQueryLanguage)
	case *pb.AlertPolicy_Condition_ConditionSql:
		out.ConditionSql = AlertpolicyConditionSql_FromProto(mapCtx, c.ConditionSql)
	}
	return out
}

func AlertpolicyConditions_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyConditions) *pb.AlertPolicy_Condition {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition{}
	out.DisplayName = in.DisplayName
	if in.Name != nil {
		out.Name = *in.Name
	}
	if in.ConditionThreshold != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionThreshold{
			ConditionThreshold: AlertpolicyConditionThreshold_ToProto(mapCtx, in.ConditionThreshold),
		}
	} else if in.ConditionAbsent != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionAbsent{
			ConditionAbsent: AlertpolicyConditionAbsent_ToProto(mapCtx, in.ConditionAbsent),
		}
	} else if in.ConditionMatchedLog != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionMatchedLog{
			ConditionMatchedLog: AlertpolicyConditionMatchedLog_ToProto(mapCtx, in.ConditionMatchedLog),
		}
	} else if in.ConditionMonitoringQueryLanguage != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionMonitoringQueryLanguage{
			ConditionMonitoringQueryLanguage: AlertpolicyConditionMonitoringQueryLanguage_ToProto(mapCtx, in.ConditionMonitoringQueryLanguage),
		}
	} else if in.ConditionPrometheusQueryLanguage != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionPrometheusQueryLanguage{
			ConditionPrometheusQueryLanguage: AlertpolicyConditionPrometheusQueryLanguage_ToProto(mapCtx, in.ConditionPrometheusQueryLanguage),
		}
	} else if in.ConditionSql != nil {
		out.Condition = &pb.AlertPolicy_Condition_ConditionSql{
			ConditionSql: AlertpolicyConditionSql_ToProto(mapCtx, in.ConditionSql),
		}
	}
	return out
}

// AlertpolicyForecastOptions is hand-coded because ForecastHorizon is non-pointer string in KRM.
func AlertpolicyForecastOptions_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_MetricThreshold_ForecastOptions) *krm.AlertpolicyForecastOptions {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyForecastOptions{}
	if d := direct.StringDuration_FromProto(mapCtx, in.GetForecastHorizon()); d != nil {
		out.ForecastHorizon = *d
	}
	return out
}

func AlertpolicyForecastOptions_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyForecastOptions) *pb.AlertPolicy_Condition_MetricThreshold_ForecastOptions {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_MetricThreshold_ForecastOptions{}
	if in.ForecastHorizon != "" {
		out.ForecastHorizon = direct.StringDuration_ToProto(mapCtx, &in.ForecastHorizon)
	}
	return out
}

// AlertpolicyTrigger is hand-coded because Count is int64 in KRM but int32 in proto.
func AlertpolicyTrigger_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_Trigger) *krm.AlertpolicyTrigger {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyTrigger{}
	if in.GetCount() != 0 {
		out.Count = direct.LazyPtr(int64(in.GetCount()))
	}
	if in.GetPercent() != 0 {
		out.Percent = direct.LazyPtr(in.GetPercent())
	}
	return out
}

func AlertpolicyTrigger_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyTrigger) *pb.AlertPolicy_Condition_Trigger {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_Trigger{}
	if in.Count != nil {
		out.Type = &pb.AlertPolicy_Condition_Trigger_Count{Count: int32(*in.Count)}
	} else if in.Percent != nil {
		out.Type = &pb.AlertPolicy_Condition_Trigger_Percent{Percent: *in.Percent}
	}
	return out
}

func AlertpolicyTrigger_Count_ToProto(mapCtx *direct.MapContext, in *int64) *pb.AlertPolicy_Condition_Trigger_Count {
	if in == nil {
		return nil
	}
	return &pb.AlertPolicy_Condition_Trigger_Count{Count: int32(*in)}
}

func AlertpolicyTrigger_Percent_ToProto(mapCtx *direct.MapContext, in *float64) *pb.AlertPolicy_Condition_Trigger_Percent {
	if in == nil {
		return nil
	}
	return &pb.AlertPolicy_Condition_Trigger_Percent{Percent: *in}
}

func AlertpolicyBooleanTest_Column_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func AlertpolicyConditionSql_Query_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func AlertpolicyExecutionTime_FromProto(mapCtx *direct.MapContext, in *timeofday.TimeOfDay) *krm.AlertpolicyExecutionTime {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyExecutionTime{}
	out.Hours = direct.LazyPtr(in.GetHours())
	out.Minutes = direct.LazyPtr(in.GetMinutes())
	out.Seconds = direct.LazyPtr(in.GetSeconds())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func AlertpolicyExecutionTime_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyExecutionTime) *timeofday.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &timeofday.TimeOfDay{}
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func AlertpolicyDaily_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_Daily) *krm.AlertpolicyDaily {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyDaily{}
	out.Periodicity = in.GetPeriodicity()
	out.ExecutionTime = AlertpolicyExecutionTime_FromProto(mapCtx, in.GetExecutionTime())
	return out
}

func AlertpolicyDaily_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyDaily) *pb.AlertPolicy_Condition_SqlCondition_Daily {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_Daily{}
	out.Periodicity = in.Periodicity
	out.ExecutionTime = AlertpolicyExecutionTime_ToProto(mapCtx, in.ExecutionTime)
	return out
}

func AlertpolicyHourly_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_Hourly) *krm.AlertpolicyHourly {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyHourly{}
	out.Periodicity = in.GetPeriodicity()
	out.MinuteOffset = direct.LazyPtr(in.GetMinuteOffset())
	return out
}

func AlertpolicyHourly_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyHourly) *pb.AlertPolicy_Condition_SqlCondition_Hourly {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_Hourly{}
	out.Periodicity = in.Periodicity
	out.MinuteOffset = in.MinuteOffset
	return out
}

func AlertpolicyMinutes_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_Minutes) *krm.AlertpolicyMinutes {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyMinutes{}
	out.Periodicity = in.GetPeriodicity()
	return out
}

func AlertpolicyMinutes_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyMinutes) *pb.AlertPolicy_Condition_SqlCondition_Minutes {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_Minutes{}
	out.Periodicity = in.Periodicity
	return out
}

func AlertpolicyRowCountTest_FromProto(mapCtx *direct.MapContext, in *pb.AlertPolicy_Condition_SqlCondition_RowCountTest) *krm.AlertpolicyRowCountTest {
	if in == nil {
		return nil
	}
	out := &krm.AlertpolicyRowCountTest{}
	out.Comparison = in.GetComparison().String()
	out.Threshold = in.GetThreshold()
	return out
}

func AlertpolicyRowCountTest_ToProto(mapCtx *direct.MapContext, in *krm.AlertpolicyRowCountTest) *pb.AlertPolicy_Condition_SqlCondition_RowCountTest {
	if in == nil {
		return nil
	}
	out := &pb.AlertPolicy_Condition_SqlCondition_RowCountTest{}
	out.Comparison = direct.Enum_ToProto[pb.ComparisonType](mapCtx, &in.Comparison)
	out.Threshold = in.Threshold
	return out
}
