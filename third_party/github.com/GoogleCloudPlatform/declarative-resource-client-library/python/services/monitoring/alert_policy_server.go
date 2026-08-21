// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// Server implements the gRPC interface for AlertPolicy.
type AlertPolicyServer struct{}

// ProtoToAlertPolicyConditionsResourceStateFilterEnum converts a AlertPolicyConditionsResourceStateFilterEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsResourceStateFilterEnum(e monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum) *monitoring.AlertPolicyConditionsResourceStateFilterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsResourceStateFilterEnum(n[len("MonitoringAlertPolicyConditionsResourceStateFilterEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum converts a AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(n[len("MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum converts a AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(n[len("MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(n[len("MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(n[len("MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionThresholdComparisonEnum converts a AlertPolicyConditionsConditionThresholdComparisonEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdComparisonEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum) *monitoring.AlertPolicyConditionsConditionThresholdComparisonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionThresholdComparisonEnum(n[len("MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum converts a AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(n[len("MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum converts a AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(n[len("MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum converts a AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum) *monitoring.AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(n[len("MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum converts a AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum) *monitoring.AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(n[len("MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionRateComparisonEnum converts a AlertPolicyConditionsConditionRateComparisonEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateComparisonEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum) *monitoring.AlertPolicyConditionsConditionRateComparisonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionRateComparisonEnum(n[len("MonitoringAlertPolicyConditionsConditionRateComparisonEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyConditionsConditionProcessCountComparisonEnum converts a AlertPolicyConditionsConditionProcessCountComparisonEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum(e monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum) *monitoring.AlertPolicyConditionsConditionProcessCountComparisonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyConditionsConditionProcessCountComparisonEnum(n[len("MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyCombinerEnum converts a AlertPolicyCombinerEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyCombinerEnum(e monitoringpb.MonitoringAlertPolicyCombinerEnum) *monitoring.AlertPolicyCombinerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyCombinerEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyCombinerEnum(n[len("MonitoringAlertPolicyCombinerEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyIncidentStrategyTypeEnum converts a AlertPolicyIncidentStrategyTypeEnum enum from its proto representation.
func ProtoToMonitoringAlertPolicyIncidentStrategyTypeEnum(e monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum) *monitoring.AlertPolicyIncidentStrategyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum_name[int32(e)]; ok {
		e := monitoring.AlertPolicyIncidentStrategyTypeEnum(n[len("MonitoringAlertPolicyIncidentStrategyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAlertPolicyDocumentation converts a AlertPolicyDocumentation resource from its proto representation.
func ProtoToMonitoringAlertPolicyDocumentation(p *monitoringpb.MonitoringAlertPolicyDocumentation) *monitoring.AlertPolicyDocumentation {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyDocumentation{
		Content:  dcl.StringOrNil(p.Content),
		MimeType: dcl.StringOrNil(p.MimeType),
	}
	return obj
}

// ProtoToAlertPolicyConditions converts a AlertPolicyConditions resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditions(p *monitoringpb.MonitoringAlertPolicyConditions) *monitoring.AlertPolicyConditions {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditions{
		Name:                             dcl.StringOrNil(p.Name),
		DisplayName:                      dcl.StringOrNil(p.DisplayName),
		ResourceStateFilter:              ProtoToMonitoringAlertPolicyConditionsResourceStateFilterEnum(p.GetResourceStateFilter()),
		ConditionThreshold:               ProtoToMonitoringAlertPolicyConditionsConditionThreshold(p.GetConditionThreshold()),
		ConditionAbsent:                  ProtoToMonitoringAlertPolicyConditionsConditionAbsent(p.GetConditionAbsent()),
		ConditionMatchedLog:              ProtoToMonitoringAlertPolicyConditionsConditionMatchedLog(p.GetConditionMatchedLog()),
		ConditionClusterOutlier:          ProtoToMonitoringAlertPolicyConditionsConditionClusterOutlier(p.GetConditionClusterOutlier()),
		ConditionRate:                    ProtoToMonitoringAlertPolicyConditionsConditionRate(p.GetConditionRate()),
		ConditionUpMon:                   ProtoToMonitoringAlertPolicyConditionsConditionUpMon(p.GetConditionUpMon()),
		ConditionProcessCount:            ProtoToMonitoringAlertPolicyConditionsConditionProcessCount(p.GetConditionProcessCount()),
		ConditionTimeSeriesQueryLanguage: ProtoToMonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguage(p.GetConditionTimeSeriesQueryLanguage()),
		ConditionMonitoringQueryLanguage: ProtoToMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage(p.GetConditionMonitoringQueryLanguage()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThreshold converts a AlertPolicyConditionsConditionThreshold resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThreshold(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThreshold) *monitoring.AlertPolicyConditionsConditionThreshold {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThreshold{
		Filter:            dcl.StringOrNil(p.Filter),
		DenominatorFilter: dcl.StringOrNil(p.DenominatorFilter),
		Comparison:        ProtoToMonitoringAlertPolicyConditionsConditionThresholdComparisonEnum(p.GetComparison()),
		ThresholdValue:    dcl.Float64OrNil(p.ThresholdValue),
		Duration:          dcl.StringOrNil(p.Duration),
		Trigger:           ProtoToMonitoringAlertPolicyConditionsConditionThresholdTrigger(p.GetTrigger()),
	}
	for _, r := range p.GetAggregations() {
		obj.Aggregations = append(obj.Aggregations, *ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregations(r))
	}
	for _, r := range p.GetDenominatorAggregations() {
		obj.DenominatorAggregations = append(obj.DenominatorAggregations, *ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations(r))
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregations converts a AlertPolicyConditionsConditionThresholdAggregations resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregations(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregations) *monitoring.AlertPolicyConditionsConditionThresholdAggregations {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregations{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams converts a AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregations converts a AlertPolicyConditionsConditionThresholdDenominatorAggregations resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregations{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionThresholdTrigger converts a AlertPolicyConditionsConditionThresholdTrigger resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionThresholdTrigger(p *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdTrigger) *monitoring.AlertPolicyConditionsConditionThresholdTrigger {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionThresholdTrigger{
		Count:   dcl.Int64OrNil(p.Count),
		Percent: dcl.Float64OrNil(p.Percent),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsent converts a AlertPolicyConditionsConditionAbsent resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsent(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsent) *monitoring.AlertPolicyConditionsConditionAbsent {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsent{
		Filter:   dcl.StringOrNil(p.Filter),
		Duration: ProtoToMonitoringAlertPolicyConditionsConditionAbsentDuration(p.GetDuration()),
		Trigger:  ProtoToMonitoringAlertPolicyConditionsConditionAbsentTrigger(p.GetTrigger()),
	}
	for _, r := range p.GetAggregations() {
		obj.Aggregations = append(obj.Aggregations, *ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregations(r))
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregations converts a AlertPolicyConditionsConditionAbsentAggregations resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregations(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregations) *monitoring.AlertPolicyConditionsConditionAbsentAggregations {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregations{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams converts a AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentDuration converts a AlertPolicyConditionsConditionAbsentDuration resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentDuration(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentDuration) *monitoring.AlertPolicyConditionsConditionAbsentDuration {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionAbsentTrigger converts a AlertPolicyConditionsConditionAbsentTrigger resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionAbsentTrigger(p *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentTrigger) *monitoring.AlertPolicyConditionsConditionAbsentTrigger {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionAbsentTrigger{
		Count:   dcl.Int64OrNil(p.Count),
		Percent: dcl.Float64OrNil(p.Percent),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionMatchedLog converts a AlertPolicyConditionsConditionMatchedLog resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionMatchedLog(p *monitoringpb.MonitoringAlertPolicyConditionsConditionMatchedLog) *monitoring.AlertPolicyConditionsConditionMatchedLog {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionMatchedLog{
		Filter: dcl.StringOrNil(p.Filter),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionClusterOutlier converts a AlertPolicyConditionsConditionClusterOutlier resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionClusterOutlier(p *monitoringpb.MonitoringAlertPolicyConditionsConditionClusterOutlier) *monitoring.AlertPolicyConditionsConditionClusterOutlier {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionClusterOutlier{
		Filter: dcl.StringOrNil(p.Filter),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRate converts a AlertPolicyConditionsConditionRate resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRate(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRate) *monitoring.AlertPolicyConditionsConditionRate {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRate{
		Filter:         dcl.StringOrNil(p.Filter),
		Comparison:     ProtoToMonitoringAlertPolicyConditionsConditionRateComparisonEnum(p.GetComparison()),
		ThresholdValue: dcl.Float64OrNil(p.ThresholdValue),
		TimeWindow:     ProtoToMonitoringAlertPolicyConditionsConditionRateTimeWindow(p.GetTimeWindow()),
		Trigger:        ProtoToMonitoringAlertPolicyConditionsConditionRateTrigger(p.GetTrigger()),
	}
	for _, r := range p.GetAggregations() {
		obj.Aggregations = append(obj.Aggregations, *ProtoToMonitoringAlertPolicyConditionsConditionRateAggregations(r))
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregations converts a AlertPolicyConditionsConditionRateAggregations resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregations(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregations) *monitoring.AlertPolicyConditionsConditionRateAggregations {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregations{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams converts a AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateTimeWindow converts a AlertPolicyConditionsConditionRateTimeWindow resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateTimeWindow(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateTimeWindow) *monitoring.AlertPolicyConditionsConditionRateTimeWindow {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateTimeWindow{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionRateTrigger converts a AlertPolicyConditionsConditionRateTrigger resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionRateTrigger(p *monitoringpb.MonitoringAlertPolicyConditionsConditionRateTrigger) *monitoring.AlertPolicyConditionsConditionRateTrigger {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionRateTrigger{
		Count:   dcl.Int64OrNil(p.Count),
		Percent: dcl.Float64OrNil(p.Percent),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionUpMon converts a AlertPolicyConditionsConditionUpMon resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionUpMon(p *monitoringpb.MonitoringAlertPolicyConditionsConditionUpMon) *monitoring.AlertPolicyConditionsConditionUpMon {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionUpMon{
		Filter:     dcl.StringOrNil(p.Filter),
		EndpointId: dcl.StringOrNil(p.EndpointId),
		CheckId:    dcl.StringOrNil(p.CheckId),
		Duration:   ProtoToMonitoringAlertPolicyConditionsConditionUpMonDuration(p.GetDuration()),
		Trigger:    ProtoToMonitoringAlertPolicyConditionsConditionUpMonTrigger(p.GetTrigger()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionUpMonDuration converts a AlertPolicyConditionsConditionUpMonDuration resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionUpMonDuration(p *monitoringpb.MonitoringAlertPolicyConditionsConditionUpMonDuration) *monitoring.AlertPolicyConditionsConditionUpMonDuration {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionUpMonDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionUpMonTrigger converts a AlertPolicyConditionsConditionUpMonTrigger resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionUpMonTrigger(p *monitoringpb.MonitoringAlertPolicyConditionsConditionUpMonTrigger) *monitoring.AlertPolicyConditionsConditionUpMonTrigger {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionUpMonTrigger{
		Count:   dcl.Int64OrNil(p.Count),
		Percent: dcl.Float64OrNil(p.Percent),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionProcessCount converts a AlertPolicyConditionsConditionProcessCount resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionProcessCount(p *monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCount) *monitoring.AlertPolicyConditionsConditionProcessCount {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionProcessCount{
		Process:               dcl.StringOrNil(p.Process),
		User:                  dcl.StringOrNil(p.User),
		Filter:                dcl.StringOrNil(p.Filter),
		Comparison:            ProtoToMonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum(p.GetComparison()),
		ProcessCountThreshold: dcl.Int64OrNil(p.ProcessCountThreshold),
		Trigger:               ProtoToMonitoringAlertPolicyConditionsConditionProcessCountTrigger(p.GetTrigger()),
		Duration:              ProtoToMonitoringAlertPolicyConditionsConditionProcessCountDuration(p.GetDuration()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionProcessCountTrigger converts a AlertPolicyConditionsConditionProcessCountTrigger resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionProcessCountTrigger(p *monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountTrigger) *monitoring.AlertPolicyConditionsConditionProcessCountTrigger {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionProcessCountTrigger{
		Count:   dcl.Int64OrNil(p.Count),
		Percent: dcl.Float64OrNil(p.Percent),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionProcessCountDuration converts a AlertPolicyConditionsConditionProcessCountDuration resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionProcessCountDuration(p *monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountDuration) *monitoring.AlertPolicyConditionsConditionProcessCountDuration {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionProcessCountDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionTimeSeriesQueryLanguage converts a AlertPolicyConditionsConditionTimeSeriesQueryLanguage resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguage(p *monitoringpb.MonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguage) *monitoring.AlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionTimeSeriesQueryLanguage{
		Query:   dcl.StringOrNil(p.Query),
		Summary: dcl.StringOrNil(p.Summary),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionMonitoringQueryLanguage converts a AlertPolicyConditionsConditionMonitoringQueryLanguage resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage(p *monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage) *monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguage {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguage{
		Query:    dcl.StringOrNil(p.Query),
		Duration: ProtoToMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(p.GetDuration()),
		Trigger:  ProtoToMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(p.GetTrigger()),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionMonitoringQueryLanguageDuration converts a AlertPolicyConditionsConditionMonitoringQueryLanguageDuration resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(p *monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration) *monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger converts a AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger resource from its proto representation.
func ProtoToMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(p *monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) *monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{
		Count:   dcl.Int64OrNil(p.Count),
		Percent: dcl.Float64OrNil(p.Percent),
	}
	return obj
}

// ProtoToAlertPolicyEnabled converts a AlertPolicyEnabled resource from its proto representation.
func ProtoToMonitoringAlertPolicyEnabled(p *monitoringpb.MonitoringAlertPolicyEnabled) *monitoring.AlertPolicyEnabled {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyEnabled{
		Value: dcl.Bool(p.Value),
	}
	return obj
}

// ProtoToAlertPolicyValidity converts a AlertPolicyValidity resource from its proto representation.
func ProtoToMonitoringAlertPolicyValidity(p *monitoringpb.MonitoringAlertPolicyValidity) *monitoring.AlertPolicyValidity {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyValidity{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToMonitoringAlertPolicyValidityDetails(r))
	}
	return obj
}

// ProtoToAlertPolicyValidityDetails converts a AlertPolicyValidityDetails resource from its proto representation.
func ProtoToMonitoringAlertPolicyValidityDetails(p *monitoringpb.MonitoringAlertPolicyValidityDetails) *monitoring.AlertPolicyValidityDetails {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyValidityDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToAlertPolicyCreationRecord converts a AlertPolicyCreationRecord resource from its proto representation.
func ProtoToMonitoringAlertPolicyCreationRecord(p *monitoringpb.MonitoringAlertPolicyCreationRecord) *monitoring.AlertPolicyCreationRecord {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyCreationRecord{
		MutateTime: ProtoToMonitoringAlertPolicyCreationRecordMutateTime(p.GetMutateTime()),
		MutatedBy:  dcl.StringOrNil(p.MutatedBy),
	}
	return obj
}

// ProtoToAlertPolicyCreationRecordMutateTime converts a AlertPolicyCreationRecordMutateTime resource from its proto representation.
func ProtoToMonitoringAlertPolicyCreationRecordMutateTime(p *monitoringpb.MonitoringAlertPolicyCreationRecordMutateTime) *monitoring.AlertPolicyCreationRecordMutateTime {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyCreationRecordMutateTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyMutationRecord converts a AlertPolicyMutationRecord resource from its proto representation.
func ProtoToMonitoringAlertPolicyMutationRecord(p *monitoringpb.MonitoringAlertPolicyMutationRecord) *monitoring.AlertPolicyMutationRecord {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyMutationRecord{
		MutateTime: ProtoToMonitoringAlertPolicyMutationRecordMutateTime(p.GetMutateTime()),
		MutatedBy:  dcl.StringOrNil(p.MutatedBy),
	}
	return obj
}

// ProtoToAlertPolicyMutationRecordMutateTime converts a AlertPolicyMutationRecordMutateTime resource from its proto representation.
func ProtoToMonitoringAlertPolicyMutationRecordMutateTime(p *monitoringpb.MonitoringAlertPolicyMutationRecordMutateTime) *monitoring.AlertPolicyMutationRecordMutateTime {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyMutationRecordMutateTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToAlertPolicyIncidentStrategy converts a AlertPolicyIncidentStrategy resource from its proto representation.
func ProtoToMonitoringAlertPolicyIncidentStrategy(p *monitoringpb.MonitoringAlertPolicyIncidentStrategy) *monitoring.AlertPolicyIncidentStrategy {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyIncidentStrategy{
		Type: ProtoToMonitoringAlertPolicyIncidentStrategyTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToAlertPolicyMetadata converts a AlertPolicyMetadata resource from its proto representation.
func ProtoToMonitoringAlertPolicyMetadata(p *monitoringpb.MonitoringAlertPolicyMetadata) *monitoring.AlertPolicyMetadata {
	if p == nil {
		return nil
	}
	obj := &monitoring.AlertPolicyMetadata{}
	for _, r := range p.GetSloNames() {
		obj.SloNames = append(obj.SloNames, r)
	}
	return obj
}

// ProtoToAlertPolicy converts a AlertPolicy resource from its proto representation.
func ProtoToAlertPolicy(p *monitoringpb.MonitoringAlertPolicy) *monitoring.AlertPolicy {
	obj := &monitoring.AlertPolicy{
		Name:             dcl.StringOrNil(p.Name),
		DisplayName:      dcl.StringOrNil(p.DisplayName),
		Documentation:    ProtoToMonitoringAlertPolicyDocumentation(p.GetDocumentation()),
		Combiner:         ProtoToMonitoringAlertPolicyCombinerEnum(p.GetCombiner()),
		Disabled:         dcl.Bool(p.Disabled),
		Enabled:          ProtoToMonitoringAlertPolicyEnabled(p.GetEnabled()),
		Validity:         ProtoToMonitoringAlertPolicyValidity(p.GetValidity()),
		CreationRecord:   ProtoToMonitoringAlertPolicyCreationRecord(p.GetCreationRecord()),
		MutationRecord:   ProtoToMonitoringAlertPolicyMutationRecord(p.GetMutationRecord()),
		IncidentStrategy: ProtoToMonitoringAlertPolicyIncidentStrategy(p.GetIncidentStrategy()),
		Metadata:         ProtoToMonitoringAlertPolicyMetadata(p.GetMetadata()),
		Project:          dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToMonitoringAlertPolicyConditions(r))
	}
	for _, r := range p.GetNotificationChannels() {
		obj.NotificationChannels = append(obj.NotificationChannels, r)
	}
	return obj
}

// AlertPolicyConditionsResourceStateFilterEnumToProto converts a AlertPolicyConditionsResourceStateFilterEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsResourceStateFilterEnumToProto(e *monitoring.AlertPolicyConditionsResourceStateFilterEnum) monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum_value["AlertPolicyConditionsResourceStateFilterEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsResourceStateFilterEnum(0)
}

// AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumToProto converts a AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumToProto(e *monitoring.AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum_value["AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(0)
}

// AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumToProto converts a AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumToProto(e *monitoring.AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum_value["AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(0)
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumToProto(e *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum_value["AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(0)
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumToProto(e *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum_value["AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(0)
}

// AlertPolicyConditionsConditionThresholdComparisonEnumToProto converts a AlertPolicyConditionsConditionThresholdComparisonEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdComparisonEnumToProto(e *monitoring.AlertPolicyConditionsConditionThresholdComparisonEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum_value["AlertPolicyConditionsConditionThresholdComparisonEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum(0)
}

// AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumToProto converts a AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumToProto(e *monitoring.AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum_value["AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(0)
}

// AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumToProto converts a AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumToProto(e *monitoring.AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum_value["AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(0)
}

// AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumToProto converts a AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumToProto(e *monitoring.AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum_value["AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(0)
}

// AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumToProto converts a AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumToProto(e *monitoring.AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum_value["AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(0)
}

// AlertPolicyConditionsConditionRateComparisonEnumToProto converts a AlertPolicyConditionsConditionRateComparisonEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateComparisonEnumToProto(e *monitoring.AlertPolicyConditionsConditionRateComparisonEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum_value["AlertPolicyConditionsConditionRateComparisonEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionRateComparisonEnum(0)
}

// AlertPolicyConditionsConditionProcessCountComparisonEnumToProto converts a AlertPolicyConditionsConditionProcessCountComparisonEnum enum to its proto representation.
func MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnumToProto(e *monitoring.AlertPolicyConditionsConditionProcessCountComparisonEnum) monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum_value["AlertPolicyConditionsConditionProcessCountComparisonEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum(0)
}

// AlertPolicyCombinerEnumToProto converts a AlertPolicyCombinerEnum enum to its proto representation.
func MonitoringAlertPolicyCombinerEnumToProto(e *monitoring.AlertPolicyCombinerEnum) monitoringpb.MonitoringAlertPolicyCombinerEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyCombinerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyCombinerEnum_value["AlertPolicyCombinerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyCombinerEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyCombinerEnum(0)
}

// AlertPolicyIncidentStrategyTypeEnumToProto converts a AlertPolicyIncidentStrategyTypeEnum enum to its proto representation.
func MonitoringAlertPolicyIncidentStrategyTypeEnumToProto(e *monitoring.AlertPolicyIncidentStrategyTypeEnum) monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum_value["AlertPolicyIncidentStrategyTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum(v)
	}
	return monitoringpb.MonitoringAlertPolicyIncidentStrategyTypeEnum(0)
}

// AlertPolicyDocumentationToProto converts a AlertPolicyDocumentation resource to its proto representation.
func MonitoringAlertPolicyDocumentationToProto(o *monitoring.AlertPolicyDocumentation) *monitoringpb.MonitoringAlertPolicyDocumentation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyDocumentation{
		Content:  dcl.ValueOrEmptyString(o.Content),
		MimeType: dcl.ValueOrEmptyString(o.MimeType),
	}
	return p
}

// AlertPolicyConditionsToProto converts a AlertPolicyConditions resource to its proto representation.
func MonitoringAlertPolicyConditionsToProto(o *monitoring.AlertPolicyConditions) *monitoringpb.MonitoringAlertPolicyConditions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditions{
		Name:                             dcl.ValueOrEmptyString(o.Name),
		DisplayName:                      dcl.ValueOrEmptyString(o.DisplayName),
		ResourceStateFilter:              MonitoringAlertPolicyConditionsResourceStateFilterEnumToProto(o.ResourceStateFilter),
		ConditionThreshold:               MonitoringAlertPolicyConditionsConditionThresholdToProto(o.ConditionThreshold),
		ConditionAbsent:                  MonitoringAlertPolicyConditionsConditionAbsentToProto(o.ConditionAbsent),
		ConditionMatchedLog:              MonitoringAlertPolicyConditionsConditionMatchedLogToProto(o.ConditionMatchedLog),
		ConditionClusterOutlier:          MonitoringAlertPolicyConditionsConditionClusterOutlierToProto(o.ConditionClusterOutlier),
		ConditionRate:                    MonitoringAlertPolicyConditionsConditionRateToProto(o.ConditionRate),
		ConditionUpMon:                   MonitoringAlertPolicyConditionsConditionUpMonToProto(o.ConditionUpMon),
		ConditionProcessCount:            MonitoringAlertPolicyConditionsConditionProcessCountToProto(o.ConditionProcessCount),
		ConditionTimeSeriesQueryLanguage: MonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguageToProto(o.ConditionTimeSeriesQueryLanguage),
		ConditionMonitoringQueryLanguage: MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageToProto(o.ConditionMonitoringQueryLanguage),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdToProto converts a AlertPolicyConditionsConditionThreshold resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdToProto(o *monitoring.AlertPolicyConditionsConditionThreshold) *monitoringpb.MonitoringAlertPolicyConditionsConditionThreshold {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThreshold{
		Filter:            dcl.ValueOrEmptyString(o.Filter),
		DenominatorFilter: dcl.ValueOrEmptyString(o.DenominatorFilter),
		Comparison:        MonitoringAlertPolicyConditionsConditionThresholdComparisonEnumToProto(o.Comparison),
		ThresholdValue:    dcl.ValueOrEmptyDouble(o.ThresholdValue),
		Duration:          dcl.ValueOrEmptyString(o.Duration),
		Trigger:           MonitoringAlertPolicyConditionsConditionThresholdTriggerToProto(o.Trigger),
	}
	for _, r := range o.Aggregations {
		p.Aggregations = append(p.Aggregations, MonitoringAlertPolicyConditionsConditionThresholdAggregationsToProto(&r))
	}
	for _, r := range o.DenominatorAggregations {
		p.DenominatorAggregations = append(p.DenominatorAggregations, MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsToProto(&r))
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsToProto converts a AlertPolicyConditionsConditionThresholdAggregations resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregations) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregations {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregations{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{
		BucketOptions:    MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingToProto converts a AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregations resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregations) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{
		BucketOptions:    MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingToProto converts a AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// AlertPolicyConditionsConditionThresholdTriggerToProto converts a AlertPolicyConditionsConditionThresholdTrigger resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionThresholdTriggerToProto(o *monitoring.AlertPolicyConditionsConditionThresholdTrigger) *monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdTrigger {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionThresholdTrigger{
		Count:   dcl.ValueOrEmptyInt64(o.Count),
		Percent: dcl.ValueOrEmptyDouble(o.Percent),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentToProto converts a AlertPolicyConditionsConditionAbsent resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentToProto(o *monitoring.AlertPolicyConditionsConditionAbsent) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsent {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsent{
		Filter:   dcl.ValueOrEmptyString(o.Filter),
		Duration: MonitoringAlertPolicyConditionsConditionAbsentDurationToProto(o.Duration),
		Trigger:  MonitoringAlertPolicyConditionsConditionAbsentTriggerToProto(o.Trigger),
	}
	for _, r := range o.Aggregations {
		p.Aggregations = append(p.Aggregations, MonitoringAlertPolicyConditionsConditionAbsentAggregationsToProto(&r))
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsToProto converts a AlertPolicyConditionsConditionAbsentAggregations resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregations) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregations {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregations{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{
		BucketOptions:    MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingToProto converts a AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentDurationToProto converts a AlertPolicyConditionsConditionAbsentDuration resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentDurationToProto(o *monitoring.AlertPolicyConditionsConditionAbsentDuration) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentDuration {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyConditionsConditionAbsentTriggerToProto converts a AlertPolicyConditionsConditionAbsentTrigger resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionAbsentTriggerToProto(o *monitoring.AlertPolicyConditionsConditionAbsentTrigger) *monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentTrigger {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionAbsentTrigger{
		Count:   dcl.ValueOrEmptyInt64(o.Count),
		Percent: dcl.ValueOrEmptyDouble(o.Percent),
	}
	return p
}

// AlertPolicyConditionsConditionMatchedLogToProto converts a AlertPolicyConditionsConditionMatchedLog resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionMatchedLogToProto(o *monitoring.AlertPolicyConditionsConditionMatchedLog) *monitoringpb.MonitoringAlertPolicyConditionsConditionMatchedLog {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionMatchedLog{
		Filter: dcl.ValueOrEmptyString(o.Filter),
	}
	p.LabelExtractors = make(map[string]string)
	for k, r := range o.LabelExtractors {
		p.LabelExtractors[k] = r
	}
	return p
}

// AlertPolicyConditionsConditionClusterOutlierToProto converts a AlertPolicyConditionsConditionClusterOutlier resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionClusterOutlierToProto(o *monitoring.AlertPolicyConditionsConditionClusterOutlier) *monitoringpb.MonitoringAlertPolicyConditionsConditionClusterOutlier {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionClusterOutlier{
		Filter: dcl.ValueOrEmptyString(o.Filter),
	}
	return p
}

// AlertPolicyConditionsConditionRateToProto converts a AlertPolicyConditionsConditionRate resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateToProto(o *monitoring.AlertPolicyConditionsConditionRate) *monitoringpb.MonitoringAlertPolicyConditionsConditionRate {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRate{
		Filter:         dcl.ValueOrEmptyString(o.Filter),
		Comparison:     MonitoringAlertPolicyConditionsConditionRateComparisonEnumToProto(o.Comparison),
		ThresholdValue: dcl.ValueOrEmptyDouble(o.ThresholdValue),
		TimeWindow:     MonitoringAlertPolicyConditionsConditionRateTimeWindowToProto(o.TimeWindow),
		Trigger:        MonitoringAlertPolicyConditionsConditionRateTriggerToProto(o.Trigger),
	}
	for _, r := range o.Aggregations {
		p.Aggregations = append(p.Aggregations, MonitoringAlertPolicyConditionsConditionRateAggregationsToProto(&r))
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsToProto converts a AlertPolicyConditionsConditionRateAggregations resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregations) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregations {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregations{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{
		BucketOptions:    MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingToProto converts a AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// AlertPolicyConditionsConditionRateTimeWindowToProto converts a AlertPolicyConditionsConditionRateTimeWindow resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateTimeWindowToProto(o *monitoring.AlertPolicyConditionsConditionRateTimeWindow) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateTimeWindow {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateTimeWindow{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyConditionsConditionRateTriggerToProto converts a AlertPolicyConditionsConditionRateTrigger resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionRateTriggerToProto(o *monitoring.AlertPolicyConditionsConditionRateTrigger) *monitoringpb.MonitoringAlertPolicyConditionsConditionRateTrigger {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionRateTrigger{
		Count:   dcl.ValueOrEmptyInt64(o.Count),
		Percent: dcl.ValueOrEmptyDouble(o.Percent),
	}
	return p
}

// AlertPolicyConditionsConditionUpMonToProto converts a AlertPolicyConditionsConditionUpMon resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionUpMonToProto(o *monitoring.AlertPolicyConditionsConditionUpMon) *monitoringpb.MonitoringAlertPolicyConditionsConditionUpMon {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionUpMon{
		Filter:     dcl.ValueOrEmptyString(o.Filter),
		EndpointId: dcl.ValueOrEmptyString(o.EndpointId),
		CheckId:    dcl.ValueOrEmptyString(o.CheckId),
		Duration:   MonitoringAlertPolicyConditionsConditionUpMonDurationToProto(o.Duration),
		Trigger:    MonitoringAlertPolicyConditionsConditionUpMonTriggerToProto(o.Trigger),
	}
	return p
}

// AlertPolicyConditionsConditionUpMonDurationToProto converts a AlertPolicyConditionsConditionUpMonDuration resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionUpMonDurationToProto(o *monitoring.AlertPolicyConditionsConditionUpMonDuration) *monitoringpb.MonitoringAlertPolicyConditionsConditionUpMonDuration {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionUpMonDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyConditionsConditionUpMonTriggerToProto converts a AlertPolicyConditionsConditionUpMonTrigger resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionUpMonTriggerToProto(o *monitoring.AlertPolicyConditionsConditionUpMonTrigger) *monitoringpb.MonitoringAlertPolicyConditionsConditionUpMonTrigger {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionUpMonTrigger{
		Count:   dcl.ValueOrEmptyInt64(o.Count),
		Percent: dcl.ValueOrEmptyDouble(o.Percent),
	}
	return p
}

// AlertPolicyConditionsConditionProcessCountToProto converts a AlertPolicyConditionsConditionProcessCount resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionProcessCountToProto(o *monitoring.AlertPolicyConditionsConditionProcessCount) *monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCount {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCount{
		Process:               dcl.ValueOrEmptyString(o.Process),
		User:                  dcl.ValueOrEmptyString(o.User),
		Filter:                dcl.ValueOrEmptyString(o.Filter),
		Comparison:            MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnumToProto(o.Comparison),
		ProcessCountThreshold: dcl.ValueOrEmptyInt64(o.ProcessCountThreshold),
		Trigger:               MonitoringAlertPolicyConditionsConditionProcessCountTriggerToProto(o.Trigger),
		Duration:              MonitoringAlertPolicyConditionsConditionProcessCountDurationToProto(o.Duration),
	}
	return p
}

// AlertPolicyConditionsConditionProcessCountTriggerToProto converts a AlertPolicyConditionsConditionProcessCountTrigger resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionProcessCountTriggerToProto(o *monitoring.AlertPolicyConditionsConditionProcessCountTrigger) *monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountTrigger {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountTrigger{
		Count:   dcl.ValueOrEmptyInt64(o.Count),
		Percent: dcl.ValueOrEmptyDouble(o.Percent),
	}
	return p
}

// AlertPolicyConditionsConditionProcessCountDurationToProto converts a AlertPolicyConditionsConditionProcessCountDuration resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionProcessCountDurationToProto(o *monitoring.AlertPolicyConditionsConditionProcessCountDuration) *monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountDuration {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionProcessCountDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyConditionsConditionTimeSeriesQueryLanguageToProto converts a AlertPolicyConditionsConditionTimeSeriesQueryLanguage resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguageToProto(o *monitoring.AlertPolicyConditionsConditionTimeSeriesQueryLanguage) *monitoringpb.MonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguage {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguage{
		Query:   dcl.ValueOrEmptyString(o.Query),
		Summary: dcl.ValueOrEmptyString(o.Summary),
	}
	return p
}

// AlertPolicyConditionsConditionMonitoringQueryLanguageToProto converts a AlertPolicyConditionsConditionMonitoringQueryLanguage resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageToProto(o *monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguage) *monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage{
		Query:    dcl.ValueOrEmptyString(o.Query),
		Duration: MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDurationToProto(o.Duration),
		Trigger:  MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerToProto(o.Trigger),
	}
	return p
}

// AlertPolicyConditionsConditionMonitoringQueryLanguageDurationToProto converts a AlertPolicyConditionsConditionMonitoringQueryLanguageDuration resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDurationToProto(o *monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) *monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyConditionsConditionMonitoringQueryLanguageTriggerToProto converts a AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger resource to its proto representation.
func MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerToProto(o *monitoring.AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) *monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{
		Count:   dcl.ValueOrEmptyInt64(o.Count),
		Percent: dcl.ValueOrEmptyDouble(o.Percent),
	}
	return p
}

// AlertPolicyEnabledToProto converts a AlertPolicyEnabled resource to its proto representation.
func MonitoringAlertPolicyEnabledToProto(o *monitoring.AlertPolicyEnabled) *monitoringpb.MonitoringAlertPolicyEnabled {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyEnabled{
		Value: dcl.ValueOrEmptyBool(o.Value),
	}
	return p
}

// AlertPolicyValidityToProto converts a AlertPolicyValidity resource to its proto representation.
func MonitoringAlertPolicyValidityToProto(o *monitoring.AlertPolicyValidity) *monitoringpb.MonitoringAlertPolicyValidity {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyValidity{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, MonitoringAlertPolicyValidityDetailsToProto(&r))
	}
	return p
}

// AlertPolicyValidityDetailsToProto converts a AlertPolicyValidityDetails resource to its proto representation.
func MonitoringAlertPolicyValidityDetailsToProto(o *monitoring.AlertPolicyValidityDetails) *monitoringpb.MonitoringAlertPolicyValidityDetails {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyValidityDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// AlertPolicyCreationRecordToProto converts a AlertPolicyCreationRecord resource to its proto representation.
func MonitoringAlertPolicyCreationRecordToProto(o *monitoring.AlertPolicyCreationRecord) *monitoringpb.MonitoringAlertPolicyCreationRecord {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyCreationRecord{
		MutateTime: MonitoringAlertPolicyCreationRecordMutateTimeToProto(o.MutateTime),
		MutatedBy:  dcl.ValueOrEmptyString(o.MutatedBy),
	}
	return p
}

// AlertPolicyCreationRecordMutateTimeToProto converts a AlertPolicyCreationRecordMutateTime resource to its proto representation.
func MonitoringAlertPolicyCreationRecordMutateTimeToProto(o *monitoring.AlertPolicyCreationRecordMutateTime) *monitoringpb.MonitoringAlertPolicyCreationRecordMutateTime {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyCreationRecordMutateTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyMutationRecordToProto converts a AlertPolicyMutationRecord resource to its proto representation.
func MonitoringAlertPolicyMutationRecordToProto(o *monitoring.AlertPolicyMutationRecord) *monitoringpb.MonitoringAlertPolicyMutationRecord {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyMutationRecord{
		MutateTime: MonitoringAlertPolicyMutationRecordMutateTimeToProto(o.MutateTime),
		MutatedBy:  dcl.ValueOrEmptyString(o.MutatedBy),
	}
	return p
}

// AlertPolicyMutationRecordMutateTimeToProto converts a AlertPolicyMutationRecordMutateTime resource to its proto representation.
func MonitoringAlertPolicyMutationRecordMutateTimeToProto(o *monitoring.AlertPolicyMutationRecordMutateTime) *monitoringpb.MonitoringAlertPolicyMutationRecordMutateTime {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyMutationRecordMutateTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// AlertPolicyIncidentStrategyToProto converts a AlertPolicyIncidentStrategy resource to its proto representation.
func MonitoringAlertPolicyIncidentStrategyToProto(o *monitoring.AlertPolicyIncidentStrategy) *monitoringpb.MonitoringAlertPolicyIncidentStrategy {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyIncidentStrategy{
		Type: MonitoringAlertPolicyIncidentStrategyTypeEnumToProto(o.Type),
	}
	return p
}

// AlertPolicyMetadataToProto converts a AlertPolicyMetadata resource to its proto representation.
func MonitoringAlertPolicyMetadataToProto(o *monitoring.AlertPolicyMetadata) *monitoringpb.MonitoringAlertPolicyMetadata {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringAlertPolicyMetadata{}
	for _, r := range o.SloNames {
		p.SloNames = append(p.SloNames, r)
	}
	return p
}

// AlertPolicyToProto converts a AlertPolicy resource to its proto representation.
func AlertPolicyToProto(resource *monitoring.AlertPolicy) *monitoringpb.MonitoringAlertPolicy {
	p := &monitoringpb.MonitoringAlertPolicy{
		Name:             dcl.ValueOrEmptyString(resource.Name),
		DisplayName:      dcl.ValueOrEmptyString(resource.DisplayName),
		Documentation:    MonitoringAlertPolicyDocumentationToProto(resource.Documentation),
		Combiner:         MonitoringAlertPolicyCombinerEnumToProto(resource.Combiner),
		Disabled:         dcl.ValueOrEmptyBool(resource.Disabled),
		Enabled:          MonitoringAlertPolicyEnabledToProto(resource.Enabled),
		Validity:         MonitoringAlertPolicyValidityToProto(resource.Validity),
		CreationRecord:   MonitoringAlertPolicyCreationRecordToProto(resource.CreationRecord),
		MutationRecord:   MonitoringAlertPolicyMutationRecordToProto(resource.MutationRecord),
		IncidentStrategy: MonitoringAlertPolicyIncidentStrategyToProto(resource.IncidentStrategy),
		Metadata:         MonitoringAlertPolicyMetadataToProto(resource.Metadata),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, MonitoringAlertPolicyConditionsToProto(&r))
	}
	for _, r := range resource.NotificationChannels {
		p.NotificationChannels = append(p.NotificationChannels, r)
	}

	return p
}

// ApplyAlertPolicy handles the gRPC request by passing it to the underlying AlertPolicy Apply() method.
func (s *AlertPolicyServer) applyAlertPolicy(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringAlertPolicyRequest) (*monitoringpb.MonitoringAlertPolicy, error) {
	p := ProtoToAlertPolicy(request.GetResource())
	res, err := c.ApplyAlertPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AlertPolicyToProto(res)
	return r, nil
}

// ApplyAlertPolicy handles the gRPC request by passing it to the underlying AlertPolicy Apply() method.
func (s *AlertPolicyServer) ApplyMonitoringAlertPolicy(ctx context.Context, request *monitoringpb.ApplyMonitoringAlertPolicyRequest) (*monitoringpb.MonitoringAlertPolicy, error) {
	cl, err := createConfigAlertPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAlertPolicy(ctx, cl, request)
}

// DeleteAlertPolicy handles the gRPC request by passing it to the underlying AlertPolicy Delete() method.
func (s *AlertPolicyServer) DeleteMonitoringAlertPolicy(ctx context.Context, request *monitoringpb.DeleteMonitoringAlertPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAlertPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAlertPolicy(ctx, ProtoToAlertPolicy(request.GetResource()))

}

// ListMonitoringAlertPolicy handles the gRPC request by passing it to the underlying AlertPolicyList() method.
func (s *AlertPolicyServer) ListMonitoringAlertPolicy(ctx context.Context, request *monitoringpb.ListMonitoringAlertPolicyRequest) (*monitoringpb.ListMonitoringAlertPolicyResponse, error) {
	cl, err := createConfigAlertPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAlertPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringAlertPolicy
	for _, r := range resources.Items {
		rp := AlertPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &monitoringpb.ListMonitoringAlertPolicyResponse{Items: protos}, nil
}

func createConfigAlertPolicy(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
