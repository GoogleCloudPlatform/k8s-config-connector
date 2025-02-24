// Copyright 2024 Google LLC. All Rights Reserved.
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
package monitoring

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Dashboard struct{}

func DashboardToUnstructured(r *dclService.Dashboard) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "alpha",
			Type:    "Dashboard",
		},
		Object: make(map[string]interface{}),
	}
	if r.ColumnLayout != nil && r.ColumnLayout != dclService.EmptyDashboardColumnLayout {
		rColumnLayout := make(map[string]interface{})
		var rColumnLayoutColumns []interface{}
		for _, rColumnLayoutColumnsVal := range r.ColumnLayout.Columns {
			rColumnLayoutColumnsObject := make(map[string]interface{})
			if rColumnLayoutColumnsVal.Weight != nil {
				rColumnLayoutColumnsObject["weight"] = *rColumnLayoutColumnsVal.Weight
			}
			var rColumnLayoutColumnsValWidgets []interface{}
			for _, rColumnLayoutColumnsValWidgetsVal := range rColumnLayoutColumnsVal.Widgets {
				rColumnLayoutColumnsValWidgetsObject := make(map[string]interface{})
				if rColumnLayoutColumnsValWidgetsVal.Blank != nil && rColumnLayoutColumnsValWidgetsVal.Blank != dclService.EmptyDashboardColumnLayoutColumnsWidgetsBlank {
					rColumnLayoutColumnsValWidgetsValBlank := make(map[string]interface{})
					rColumnLayoutColumnsValWidgetsObject["blank"] = rColumnLayoutColumnsValWidgetsValBlank
				}
				if rColumnLayoutColumnsValWidgetsVal.LogsPanel != nil && rColumnLayoutColumnsValWidgetsVal.LogsPanel != dclService.EmptyDashboardColumnLayoutColumnsWidgetsLogsPanel {
					rColumnLayoutColumnsValWidgetsValLogsPanel := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.LogsPanel.Filter != nil {
						rColumnLayoutColumnsValWidgetsValLogsPanel["filter"] = *rColumnLayoutColumnsValWidgetsVal.LogsPanel.Filter
					}
					var rColumnLayoutColumnsValWidgetsValLogsPanelResourceNames []interface{}
					for _, rColumnLayoutColumnsValWidgetsValLogsPanelResourceNamesVal := range rColumnLayoutColumnsValWidgetsVal.LogsPanel.ResourceNames {
						rColumnLayoutColumnsValWidgetsValLogsPanelResourceNames = append(rColumnLayoutColumnsValWidgetsValLogsPanelResourceNames, rColumnLayoutColumnsValWidgetsValLogsPanelResourceNamesVal)
					}
					rColumnLayoutColumnsValWidgetsValLogsPanel["resourceNames"] = rColumnLayoutColumnsValWidgetsValLogsPanelResourceNames
					rColumnLayoutColumnsValWidgetsObject["logsPanel"] = rColumnLayoutColumnsValWidgetsValLogsPanel
				}
				if rColumnLayoutColumnsValWidgetsVal.Scorecard != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecard {
					rColumnLayoutColumnsValWidgetsValScorecard := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardGaugeView {
						rColumnLayoutColumnsValWidgetsValScorecardGaugeView := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.LowerBound != nil {
							rColumnLayoutColumnsValWidgetsValScorecardGaugeView["lowerBound"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.LowerBound
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.UpperBound != nil {
							rColumnLayoutColumnsValWidgetsValScorecardGaugeView["upperBound"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.GaugeView.UpperBound
						}
						rColumnLayoutColumnsValWidgetsValScorecard["gaugeView"] = rColumnLayoutColumnsValWidgetsValScorecardGaugeView
					}
					if rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardSparkChartView {
						rColumnLayoutColumnsValWidgetsValScorecardSparkChartView := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod != nil {
							rColumnLayoutColumnsValWidgetsValScorecardSparkChartView["minAlignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.SparkChartType != nil {
							rColumnLayoutColumnsValWidgetsValScorecardSparkChartView["sparkChartType"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.SparkChartView.SparkChartType)
						}
						rColumnLayoutColumnsValWidgetsValScorecard["sparkChartView"] = rColumnLayoutColumnsValWidgetsValScorecardSparkChartView
					}
					var rColumnLayoutColumnsValWidgetsValScorecardThresholds []interface{}
					for _, rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.Thresholds {
						rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Color != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["color"] = string(*rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Color)
						}
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Direction != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["direction"] = string(*rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Direction)
						}
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Label != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["label"] = *rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Label
						}
						if rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Value != nil {
							rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject["value"] = *rColumnLayoutColumnsValWidgetsValScorecardThresholdsVal.Value
						}
						rColumnLayoutColumnsValWidgetsValScorecardThresholds = append(rColumnLayoutColumnsValWidgetsValScorecardThresholds, rColumnLayoutColumnsValWidgetsValScorecardThresholdsObject)
					}
					rColumnLayoutColumnsValWidgetsValScorecard["thresholds"] = rColumnLayoutColumnsValWidgetsValScorecardThresholds
					if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery {
						rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
								}
								var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
								for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["filter"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
								}
								var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
								for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
							}
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["timeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
							}
							if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
								}
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
								}
								var rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
								for _, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
								if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
									rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
								}
								rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
							}
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["timeSeriesFilterRatio"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["timeSeriesQueryLanguage"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage
						}
						if rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride != nil {
							rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery["unitOverride"] = *rColumnLayoutColumnsValWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride
						}
						rColumnLayoutColumnsValWidgetsValScorecard["timeSeriesQuery"] = rColumnLayoutColumnsValWidgetsValScorecardTimeSeriesQuery
					}
					rColumnLayoutColumnsValWidgetsObject["scorecard"] = rColumnLayoutColumnsValWidgetsValScorecard
				}
				if rColumnLayoutColumnsValWidgetsVal.Text != nil && rColumnLayoutColumnsValWidgetsVal.Text != dclService.EmptyDashboardColumnLayoutColumnsWidgetsText {
					rColumnLayoutColumnsValWidgetsValText := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.Text.Content != nil {
						rColumnLayoutColumnsValWidgetsValText["content"] = *rColumnLayoutColumnsValWidgetsVal.Text.Content
					}
					if rColumnLayoutColumnsValWidgetsVal.Text.Format != nil {
						rColumnLayoutColumnsValWidgetsValText["format"] = string(*rColumnLayoutColumnsValWidgetsVal.Text.Format)
					}
					rColumnLayoutColumnsValWidgetsObject["text"] = rColumnLayoutColumnsValWidgetsValText
				}
				if rColumnLayoutColumnsValWidgetsVal.Title != nil {
					rColumnLayoutColumnsValWidgetsObject["title"] = *rColumnLayoutColumnsValWidgetsVal.Title
				}
				if rColumnLayoutColumnsValWidgetsVal.XyChart != nil && rColumnLayoutColumnsValWidgetsVal.XyChart != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChart {
					rColumnLayoutColumnsValWidgetsValXyChart := make(map[string]interface{})
					if rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions != nil && rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartChartOptions {
						rColumnLayoutColumnsValWidgetsValXyChartChartOptions := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions.Mode != nil {
							rColumnLayoutColumnsValWidgetsValXyChartChartOptions["mode"] = string(*rColumnLayoutColumnsValWidgetsVal.XyChart.ChartOptions.Mode)
						}
						rColumnLayoutColumnsValWidgetsValXyChart["chartOptions"] = rColumnLayoutColumnsValWidgetsValXyChartChartOptions
					}
					var rColumnLayoutColumnsValWidgetsValXyChartDataSets []interface{}
					for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal := range rColumnLayoutColumnsValWidgetsVal.XyChart.DataSets {
						rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.LegendTemplate != nil {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["legendTemplate"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.LegendTemplate
						}
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.MinAlignmentPeriod != nil {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["minAlignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.MinAlignmentPeriod
						}
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.PlotType != nil {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["plotType"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.PlotType)
						}
						if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery {
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery := make(map[string]interface{})
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["aggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["filter"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
								}
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter
							}
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
										}
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
										}
										var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
										for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
										}
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
										}
										var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
										for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
										if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
											rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
										}
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
								}
								if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
									}
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
									}
									var rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
									for _, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
									if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
										rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
									}
									rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
								}
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilterRatio"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio
							}
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesQueryLanguage"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage
							}
							if rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride != nil {
								rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery["unitOverride"] = *rColumnLayoutColumnsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride
							}
							rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject["timeSeriesQuery"] = rColumnLayoutColumnsValWidgetsValXyChartDataSetsValTimeSeriesQuery
						}
						rColumnLayoutColumnsValWidgetsValXyChartDataSets = append(rColumnLayoutColumnsValWidgetsValXyChartDataSets, rColumnLayoutColumnsValWidgetsValXyChartDataSetsObject)
					}
					rColumnLayoutColumnsValWidgetsValXyChart["dataSets"] = rColumnLayoutColumnsValWidgetsValXyChartDataSets
					var rColumnLayoutColumnsValWidgetsValXyChartThresholds []interface{}
					for _, rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal := range rColumnLayoutColumnsValWidgetsVal.XyChart.Thresholds {
						rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Color != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["color"] = string(*rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Color)
						}
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Direction != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["direction"] = string(*rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Direction)
						}
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Label != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["label"] = *rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Label
						}
						if rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Value != nil {
							rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject["value"] = *rColumnLayoutColumnsValWidgetsValXyChartThresholdsVal.Value
						}
						rColumnLayoutColumnsValWidgetsValXyChartThresholds = append(rColumnLayoutColumnsValWidgetsValXyChartThresholds, rColumnLayoutColumnsValWidgetsValXyChartThresholdsObject)
					}
					rColumnLayoutColumnsValWidgetsValXyChart["thresholds"] = rColumnLayoutColumnsValWidgetsValXyChartThresholds
					if rColumnLayoutColumnsValWidgetsVal.XyChart.TimeshiftDuration != nil {
						rColumnLayoutColumnsValWidgetsValXyChart["timeshiftDuration"] = *rColumnLayoutColumnsValWidgetsVal.XyChart.TimeshiftDuration
					}
					if rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis != nil && rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartXAxis {
						rColumnLayoutColumnsValWidgetsValXyChartXAxis := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Label != nil {
							rColumnLayoutColumnsValWidgetsValXyChartXAxis["label"] = *rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Label
						}
						if rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Scale != nil {
							rColumnLayoutColumnsValWidgetsValXyChartXAxis["scale"] = string(*rColumnLayoutColumnsValWidgetsVal.XyChart.XAxis.Scale)
						}
						rColumnLayoutColumnsValWidgetsValXyChart["xAxis"] = rColumnLayoutColumnsValWidgetsValXyChartXAxis
					}
					if rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis != nil && rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis != dclService.EmptyDashboardColumnLayoutColumnsWidgetsXyChartYAxis {
						rColumnLayoutColumnsValWidgetsValXyChartYAxis := make(map[string]interface{})
						if rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Label != nil {
							rColumnLayoutColumnsValWidgetsValXyChartYAxis["label"] = *rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Label
						}
						if rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Scale != nil {
							rColumnLayoutColumnsValWidgetsValXyChartYAxis["scale"] = string(*rColumnLayoutColumnsValWidgetsVal.XyChart.YAxis.Scale)
						}
						rColumnLayoutColumnsValWidgetsValXyChart["yAxis"] = rColumnLayoutColumnsValWidgetsValXyChartYAxis
					}
					rColumnLayoutColumnsValWidgetsObject["xyChart"] = rColumnLayoutColumnsValWidgetsValXyChart
				}
				rColumnLayoutColumnsValWidgets = append(rColumnLayoutColumnsValWidgets, rColumnLayoutColumnsValWidgetsObject)
			}
			rColumnLayoutColumnsObject["widgets"] = rColumnLayoutColumnsValWidgets
			rColumnLayoutColumns = append(rColumnLayoutColumns, rColumnLayoutColumnsObject)
		}
		rColumnLayout["columns"] = rColumnLayoutColumns
		u.Object["columnLayout"] = rColumnLayout
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.GridLayout != nil && r.GridLayout != dclService.EmptyDashboardGridLayout {
		rGridLayout := make(map[string]interface{})
		if r.GridLayout.Columns != nil {
			rGridLayout["columns"] = *r.GridLayout.Columns
		}
		var rGridLayoutWidgets []interface{}
		for _, rGridLayoutWidgetsVal := range r.GridLayout.Widgets {
			rGridLayoutWidgetsObject := make(map[string]interface{})
			if rGridLayoutWidgetsVal.Blank != nil && rGridLayoutWidgetsVal.Blank != dclService.EmptyDashboardGridLayoutWidgetsBlank {
				rGridLayoutWidgetsValBlank := make(map[string]interface{})
				rGridLayoutWidgetsObject["blank"] = rGridLayoutWidgetsValBlank
			}
			if rGridLayoutWidgetsVal.LogsPanel != nil && rGridLayoutWidgetsVal.LogsPanel != dclService.EmptyDashboardGridLayoutWidgetsLogsPanel {
				rGridLayoutWidgetsValLogsPanel := make(map[string]interface{})
				if rGridLayoutWidgetsVal.LogsPanel.Filter != nil {
					rGridLayoutWidgetsValLogsPanel["filter"] = *rGridLayoutWidgetsVal.LogsPanel.Filter
				}
				var rGridLayoutWidgetsValLogsPanelResourceNames []interface{}
				for _, rGridLayoutWidgetsValLogsPanelResourceNamesVal := range rGridLayoutWidgetsVal.LogsPanel.ResourceNames {
					rGridLayoutWidgetsValLogsPanelResourceNames = append(rGridLayoutWidgetsValLogsPanelResourceNames, rGridLayoutWidgetsValLogsPanelResourceNamesVal)
				}
				rGridLayoutWidgetsValLogsPanel["resourceNames"] = rGridLayoutWidgetsValLogsPanelResourceNames
				rGridLayoutWidgetsObject["logsPanel"] = rGridLayoutWidgetsValLogsPanel
			}
			if rGridLayoutWidgetsVal.Scorecard != nil && rGridLayoutWidgetsVal.Scorecard != dclService.EmptyDashboardGridLayoutWidgetsScorecard {
				rGridLayoutWidgetsValScorecard := make(map[string]interface{})
				if rGridLayoutWidgetsVal.Scorecard.GaugeView != nil && rGridLayoutWidgetsVal.Scorecard.GaugeView != dclService.EmptyDashboardGridLayoutWidgetsScorecardGaugeView {
					rGridLayoutWidgetsValScorecardGaugeView := make(map[string]interface{})
					if rGridLayoutWidgetsVal.Scorecard.GaugeView.LowerBound != nil {
						rGridLayoutWidgetsValScorecardGaugeView["lowerBound"] = *rGridLayoutWidgetsVal.Scorecard.GaugeView.LowerBound
					}
					if rGridLayoutWidgetsVal.Scorecard.GaugeView.UpperBound != nil {
						rGridLayoutWidgetsValScorecardGaugeView["upperBound"] = *rGridLayoutWidgetsVal.Scorecard.GaugeView.UpperBound
					}
					rGridLayoutWidgetsValScorecard["gaugeView"] = rGridLayoutWidgetsValScorecardGaugeView
				}
				if rGridLayoutWidgetsVal.Scorecard.SparkChartView != nil && rGridLayoutWidgetsVal.Scorecard.SparkChartView != dclService.EmptyDashboardGridLayoutWidgetsScorecardSparkChartView {
					rGridLayoutWidgetsValScorecardSparkChartView := make(map[string]interface{})
					if rGridLayoutWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod != nil {
						rGridLayoutWidgetsValScorecardSparkChartView["minAlignmentPeriod"] = *rGridLayoutWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod
					}
					if rGridLayoutWidgetsVal.Scorecard.SparkChartView.SparkChartType != nil {
						rGridLayoutWidgetsValScorecardSparkChartView["sparkChartType"] = string(*rGridLayoutWidgetsVal.Scorecard.SparkChartView.SparkChartType)
					}
					rGridLayoutWidgetsValScorecard["sparkChartView"] = rGridLayoutWidgetsValScorecardSparkChartView
				}
				var rGridLayoutWidgetsValScorecardThresholds []interface{}
				for _, rGridLayoutWidgetsValScorecardThresholdsVal := range rGridLayoutWidgetsVal.Scorecard.Thresholds {
					rGridLayoutWidgetsValScorecardThresholdsObject := make(map[string]interface{})
					if rGridLayoutWidgetsValScorecardThresholdsVal.Color != nil {
						rGridLayoutWidgetsValScorecardThresholdsObject["color"] = string(*rGridLayoutWidgetsValScorecardThresholdsVal.Color)
					}
					if rGridLayoutWidgetsValScorecardThresholdsVal.Direction != nil {
						rGridLayoutWidgetsValScorecardThresholdsObject["direction"] = string(*rGridLayoutWidgetsValScorecardThresholdsVal.Direction)
					}
					if rGridLayoutWidgetsValScorecardThresholdsVal.Label != nil {
						rGridLayoutWidgetsValScorecardThresholdsObject["label"] = *rGridLayoutWidgetsValScorecardThresholdsVal.Label
					}
					if rGridLayoutWidgetsValScorecardThresholdsVal.Value != nil {
						rGridLayoutWidgetsValScorecardThresholdsObject["value"] = *rGridLayoutWidgetsValScorecardThresholdsVal.Value
					}
					rGridLayoutWidgetsValScorecardThresholds = append(rGridLayoutWidgetsValScorecardThresholds, rGridLayoutWidgetsValScorecardThresholdsObject)
				}
				rGridLayoutWidgetsValScorecard["thresholds"] = rGridLayoutWidgetsValScorecardThresholds
				if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQuery {
					rGridLayoutWidgetsValScorecardTimeSeriesQuery := make(map[string]interface{})
					if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter {
						rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
							}
							var rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
							for _, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation
						}
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["filter"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter
						}
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
						}
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
							}
							var rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
							for _, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
						}
						rGridLayoutWidgetsValScorecardTimeSeriesQuery["timeSeriesFilter"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter
					}
					if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio {
						rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
								if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
								}
								if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
								}
								var rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
								for _, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
								}
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
								if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
								}
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
						}
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
								if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
								}
								if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
								}
								var rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
								for _, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
								}
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
								if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
									rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
								}
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
						}
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
						}
						if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
							}
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
							}
							var rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
							for _, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
							if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
								rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
							}
							rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
						}
						rGridLayoutWidgetsValScorecardTimeSeriesQuery["timeSeriesFilterRatio"] = rGridLayoutWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio
					}
					if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
						rGridLayoutWidgetsValScorecardTimeSeriesQuery["timeSeriesQueryLanguage"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage
					}
					if rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride != nil {
						rGridLayoutWidgetsValScorecardTimeSeriesQuery["unitOverride"] = *rGridLayoutWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride
					}
					rGridLayoutWidgetsValScorecard["timeSeriesQuery"] = rGridLayoutWidgetsValScorecardTimeSeriesQuery
				}
				rGridLayoutWidgetsObject["scorecard"] = rGridLayoutWidgetsValScorecard
			}
			if rGridLayoutWidgetsVal.Text != nil && rGridLayoutWidgetsVal.Text != dclService.EmptyDashboardGridLayoutWidgetsText {
				rGridLayoutWidgetsValText := make(map[string]interface{})
				if rGridLayoutWidgetsVal.Text.Content != nil {
					rGridLayoutWidgetsValText["content"] = *rGridLayoutWidgetsVal.Text.Content
				}
				if rGridLayoutWidgetsVal.Text.Format != nil {
					rGridLayoutWidgetsValText["format"] = string(*rGridLayoutWidgetsVal.Text.Format)
				}
				rGridLayoutWidgetsObject["text"] = rGridLayoutWidgetsValText
			}
			if rGridLayoutWidgetsVal.Title != nil {
				rGridLayoutWidgetsObject["title"] = *rGridLayoutWidgetsVal.Title
			}
			if rGridLayoutWidgetsVal.XyChart != nil && rGridLayoutWidgetsVal.XyChart != dclService.EmptyDashboardGridLayoutWidgetsXyChart {
				rGridLayoutWidgetsValXyChart := make(map[string]interface{})
				if rGridLayoutWidgetsVal.XyChart.ChartOptions != nil && rGridLayoutWidgetsVal.XyChart.ChartOptions != dclService.EmptyDashboardGridLayoutWidgetsXyChartChartOptions {
					rGridLayoutWidgetsValXyChartChartOptions := make(map[string]interface{})
					if rGridLayoutWidgetsVal.XyChart.ChartOptions.Mode != nil {
						rGridLayoutWidgetsValXyChartChartOptions["mode"] = string(*rGridLayoutWidgetsVal.XyChart.ChartOptions.Mode)
					}
					rGridLayoutWidgetsValXyChart["chartOptions"] = rGridLayoutWidgetsValXyChartChartOptions
				}
				var rGridLayoutWidgetsValXyChartDataSets []interface{}
				for _, rGridLayoutWidgetsValXyChartDataSetsVal := range rGridLayoutWidgetsVal.XyChart.DataSets {
					rGridLayoutWidgetsValXyChartDataSetsObject := make(map[string]interface{})
					if rGridLayoutWidgetsValXyChartDataSetsVal.LegendTemplate != nil {
						rGridLayoutWidgetsValXyChartDataSetsObject["legendTemplate"] = *rGridLayoutWidgetsValXyChartDataSetsVal.LegendTemplate
					}
					if rGridLayoutWidgetsValXyChartDataSetsVal.MinAlignmentPeriod != nil {
						rGridLayoutWidgetsValXyChartDataSetsObject["minAlignmentPeriod"] = *rGridLayoutWidgetsValXyChartDataSetsVal.MinAlignmentPeriod
					}
					if rGridLayoutWidgetsValXyChartDataSetsVal.PlotType != nil {
						rGridLayoutWidgetsValXyChartDataSetsObject["plotType"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.PlotType)
					}
					if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery {
						rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQuery := make(map[string]interface{})
						if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
							rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
								}
								var rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
								for _, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["aggregation"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation
							}
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["filter"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter
							}
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
							}
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
								}
								var rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
								for _, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
							}
							rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilter"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter
						}
						if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
							rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
									if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
									}
									if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
									}
									var rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
									for _, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
									}
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
									if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
									}
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator
							}
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
									if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
									}
									if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
									}
									var rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
									for _, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
									}
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
									if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
										rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
									}
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator
							}
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
							}
							if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
								}
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
								}
								var rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
								for _, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
								if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
									rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
								}
								rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
							}
							rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilterRatio"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio
						}
						if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
							rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesQueryLanguage"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage
						}
						if rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride != nil {
							rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQuery["unitOverride"] = *rGridLayoutWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride
						}
						rGridLayoutWidgetsValXyChartDataSetsObject["timeSeriesQuery"] = rGridLayoutWidgetsValXyChartDataSetsValTimeSeriesQuery
					}
					rGridLayoutWidgetsValXyChartDataSets = append(rGridLayoutWidgetsValXyChartDataSets, rGridLayoutWidgetsValXyChartDataSetsObject)
				}
				rGridLayoutWidgetsValXyChart["dataSets"] = rGridLayoutWidgetsValXyChartDataSets
				var rGridLayoutWidgetsValXyChartThresholds []interface{}
				for _, rGridLayoutWidgetsValXyChartThresholdsVal := range rGridLayoutWidgetsVal.XyChart.Thresholds {
					rGridLayoutWidgetsValXyChartThresholdsObject := make(map[string]interface{})
					if rGridLayoutWidgetsValXyChartThresholdsVal.Color != nil {
						rGridLayoutWidgetsValXyChartThresholdsObject["color"] = string(*rGridLayoutWidgetsValXyChartThresholdsVal.Color)
					}
					if rGridLayoutWidgetsValXyChartThresholdsVal.Direction != nil {
						rGridLayoutWidgetsValXyChartThresholdsObject["direction"] = string(*rGridLayoutWidgetsValXyChartThresholdsVal.Direction)
					}
					if rGridLayoutWidgetsValXyChartThresholdsVal.Label != nil {
						rGridLayoutWidgetsValXyChartThresholdsObject["label"] = *rGridLayoutWidgetsValXyChartThresholdsVal.Label
					}
					if rGridLayoutWidgetsValXyChartThresholdsVal.Value != nil {
						rGridLayoutWidgetsValXyChartThresholdsObject["value"] = *rGridLayoutWidgetsValXyChartThresholdsVal.Value
					}
					rGridLayoutWidgetsValXyChartThresholds = append(rGridLayoutWidgetsValXyChartThresholds, rGridLayoutWidgetsValXyChartThresholdsObject)
				}
				rGridLayoutWidgetsValXyChart["thresholds"] = rGridLayoutWidgetsValXyChartThresholds
				if rGridLayoutWidgetsVal.XyChart.TimeshiftDuration != nil {
					rGridLayoutWidgetsValXyChart["timeshiftDuration"] = *rGridLayoutWidgetsVal.XyChart.TimeshiftDuration
				}
				if rGridLayoutWidgetsVal.XyChart.XAxis != nil && rGridLayoutWidgetsVal.XyChart.XAxis != dclService.EmptyDashboardGridLayoutWidgetsXyChartXAxis {
					rGridLayoutWidgetsValXyChartXAxis := make(map[string]interface{})
					if rGridLayoutWidgetsVal.XyChart.XAxis.Label != nil {
						rGridLayoutWidgetsValXyChartXAxis["label"] = *rGridLayoutWidgetsVal.XyChart.XAxis.Label
					}
					if rGridLayoutWidgetsVal.XyChart.XAxis.Scale != nil {
						rGridLayoutWidgetsValXyChartXAxis["scale"] = string(*rGridLayoutWidgetsVal.XyChart.XAxis.Scale)
					}
					rGridLayoutWidgetsValXyChart["xAxis"] = rGridLayoutWidgetsValXyChartXAxis
				}
				if rGridLayoutWidgetsVal.XyChart.YAxis != nil && rGridLayoutWidgetsVal.XyChart.YAxis != dclService.EmptyDashboardGridLayoutWidgetsXyChartYAxis {
					rGridLayoutWidgetsValXyChartYAxis := make(map[string]interface{})
					if rGridLayoutWidgetsVal.XyChart.YAxis.Label != nil {
						rGridLayoutWidgetsValXyChartYAxis["label"] = *rGridLayoutWidgetsVal.XyChart.YAxis.Label
					}
					if rGridLayoutWidgetsVal.XyChart.YAxis.Scale != nil {
						rGridLayoutWidgetsValXyChartYAxis["scale"] = string(*rGridLayoutWidgetsVal.XyChart.YAxis.Scale)
					}
					rGridLayoutWidgetsValXyChart["yAxis"] = rGridLayoutWidgetsValXyChartYAxis
				}
				rGridLayoutWidgetsObject["xyChart"] = rGridLayoutWidgetsValXyChart
			}
			rGridLayoutWidgets = append(rGridLayoutWidgets, rGridLayoutWidgetsObject)
		}
		rGridLayout["widgets"] = rGridLayoutWidgets
		u.Object["gridLayout"] = rGridLayout
	}
	if r.MosaicLayout != nil && r.MosaicLayout != dclService.EmptyDashboardMosaicLayout {
		rMosaicLayout := make(map[string]interface{})
		if r.MosaicLayout.Columns != nil {
			rMosaicLayout["columns"] = *r.MosaicLayout.Columns
		}
		var rMosaicLayoutTiles []interface{}
		for _, rMosaicLayoutTilesVal := range r.MosaicLayout.Tiles {
			rMosaicLayoutTilesObject := make(map[string]interface{})
			if rMosaicLayoutTilesVal.Height != nil {
				rMosaicLayoutTilesObject["height"] = *rMosaicLayoutTilesVal.Height
			}
			if rMosaicLayoutTilesVal.Widget != nil && rMosaicLayoutTilesVal.Widget != dclService.EmptyDashboardMosaicLayoutTilesWidget {
				rMosaicLayoutTilesValWidget := make(map[string]interface{})
				if rMosaicLayoutTilesVal.Widget.Blank != nil && rMosaicLayoutTilesVal.Widget.Blank != dclService.EmptyDashboardMosaicLayoutTilesWidgetBlank {
					rMosaicLayoutTilesValWidgetBlank := make(map[string]interface{})
					rMosaicLayoutTilesValWidget["blank"] = rMosaicLayoutTilesValWidgetBlank
				}
				if rMosaicLayoutTilesVal.Widget.LogsPanel != nil && rMosaicLayoutTilesVal.Widget.LogsPanel != dclService.EmptyDashboardMosaicLayoutTilesWidgetLogsPanel {
					rMosaicLayoutTilesValWidgetLogsPanel := make(map[string]interface{})
					if rMosaicLayoutTilesVal.Widget.LogsPanel.Filter != nil {
						rMosaicLayoutTilesValWidgetLogsPanel["filter"] = *rMosaicLayoutTilesVal.Widget.LogsPanel.Filter
					}
					var rMosaicLayoutTilesValWidgetLogsPanelResourceNames []interface{}
					for _, rMosaicLayoutTilesValWidgetLogsPanelResourceNamesVal := range rMosaicLayoutTilesVal.Widget.LogsPanel.ResourceNames {
						rMosaicLayoutTilesValWidgetLogsPanelResourceNames = append(rMosaicLayoutTilesValWidgetLogsPanelResourceNames, rMosaicLayoutTilesValWidgetLogsPanelResourceNamesVal)
					}
					rMosaicLayoutTilesValWidgetLogsPanel["resourceNames"] = rMosaicLayoutTilesValWidgetLogsPanelResourceNames
					rMosaicLayoutTilesValWidget["logsPanel"] = rMosaicLayoutTilesValWidgetLogsPanel
				}
				if rMosaicLayoutTilesVal.Widget.Scorecard != nil && rMosaicLayoutTilesVal.Widget.Scorecard != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecard {
					rMosaicLayoutTilesValWidgetScorecard := make(map[string]interface{})
					if rMosaicLayoutTilesVal.Widget.Scorecard.GaugeView != nil && rMosaicLayoutTilesVal.Widget.Scorecard.GaugeView != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardGaugeView {
						rMosaicLayoutTilesValWidgetScorecardGaugeView := make(map[string]interface{})
						if rMosaicLayoutTilesVal.Widget.Scorecard.GaugeView.LowerBound != nil {
							rMosaicLayoutTilesValWidgetScorecardGaugeView["lowerBound"] = *rMosaicLayoutTilesVal.Widget.Scorecard.GaugeView.LowerBound
						}
						if rMosaicLayoutTilesVal.Widget.Scorecard.GaugeView.UpperBound != nil {
							rMosaicLayoutTilesValWidgetScorecardGaugeView["upperBound"] = *rMosaicLayoutTilesVal.Widget.Scorecard.GaugeView.UpperBound
						}
						rMosaicLayoutTilesValWidgetScorecard["gaugeView"] = rMosaicLayoutTilesValWidgetScorecardGaugeView
					}
					if rMosaicLayoutTilesVal.Widget.Scorecard.SparkChartView != nil && rMosaicLayoutTilesVal.Widget.Scorecard.SparkChartView != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardSparkChartView {
						rMosaicLayoutTilesValWidgetScorecardSparkChartView := make(map[string]interface{})
						if rMosaicLayoutTilesVal.Widget.Scorecard.SparkChartView.MinAlignmentPeriod != nil {
							rMosaicLayoutTilesValWidgetScorecardSparkChartView["minAlignmentPeriod"] = *rMosaicLayoutTilesVal.Widget.Scorecard.SparkChartView.MinAlignmentPeriod
						}
						if rMosaicLayoutTilesVal.Widget.Scorecard.SparkChartView.SparkChartType != nil {
							rMosaicLayoutTilesValWidgetScorecardSparkChartView["sparkChartType"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.SparkChartView.SparkChartType)
						}
						rMosaicLayoutTilesValWidgetScorecard["sparkChartView"] = rMosaicLayoutTilesValWidgetScorecardSparkChartView
					}
					var rMosaicLayoutTilesValWidgetScorecardThresholds []interface{}
					for _, rMosaicLayoutTilesValWidgetScorecardThresholdsVal := range rMosaicLayoutTilesVal.Widget.Scorecard.Thresholds {
						rMosaicLayoutTilesValWidgetScorecardThresholdsObject := make(map[string]interface{})
						if rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Color != nil {
							rMosaicLayoutTilesValWidgetScorecardThresholdsObject["color"] = string(*rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Color)
						}
						if rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Direction != nil {
							rMosaicLayoutTilesValWidgetScorecardThresholdsObject["direction"] = string(*rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Direction)
						}
						if rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Label != nil {
							rMosaicLayoutTilesValWidgetScorecardThresholdsObject["label"] = *rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Label
						}
						if rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Value != nil {
							rMosaicLayoutTilesValWidgetScorecardThresholdsObject["value"] = *rMosaicLayoutTilesValWidgetScorecardThresholdsVal.Value
						}
						rMosaicLayoutTilesValWidgetScorecardThresholds = append(rMosaicLayoutTilesValWidgetScorecardThresholds, rMosaicLayoutTilesValWidgetScorecardThresholdsObject)
					}
					rMosaicLayoutTilesValWidgetScorecard["thresholds"] = rMosaicLayoutTilesValWidgetScorecardThresholds
					if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery {
						rMosaicLayoutTilesValWidgetScorecardTimeSeriesQuery := make(map[string]interface{})
						if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
							rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
								}
								var rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
								for _, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation
							}
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilter["filter"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter
							}
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
							}
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
								}
								var rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
								for _, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
							}
							rMosaicLayoutTilesValWidgetScorecardTimeSeriesQuery["timeSeriesFilter"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilter
						}
						if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
							rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
									if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
									}
									if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
									}
									var rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
									for _, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
									}
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
									if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
									}
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
							}
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
									if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
									}
									if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
									}
									var rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
									for _, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
									}
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
									if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
										rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
									}
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
							}
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
							}
							if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
								}
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
								}
								var rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
								for _, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
								if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
									rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
								}
								rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
							}
							rMosaicLayoutTilesValWidgetScorecardTimeSeriesQuery["timeSeriesFilterRatio"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio
						}
						if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
							rMosaicLayoutTilesValWidgetScorecardTimeSeriesQuery["timeSeriesQueryLanguage"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage
						}
						if rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.UnitOverride != nil {
							rMosaicLayoutTilesValWidgetScorecardTimeSeriesQuery["unitOverride"] = *rMosaicLayoutTilesVal.Widget.Scorecard.TimeSeriesQuery.UnitOverride
						}
						rMosaicLayoutTilesValWidgetScorecard["timeSeriesQuery"] = rMosaicLayoutTilesValWidgetScorecardTimeSeriesQuery
					}
					rMosaicLayoutTilesValWidget["scorecard"] = rMosaicLayoutTilesValWidgetScorecard
				}
				if rMosaicLayoutTilesVal.Widget.Text != nil && rMosaicLayoutTilesVal.Widget.Text != dclService.EmptyDashboardMosaicLayoutTilesWidgetText {
					rMosaicLayoutTilesValWidgetText := make(map[string]interface{})
					if rMosaicLayoutTilesVal.Widget.Text.Content != nil {
						rMosaicLayoutTilesValWidgetText["content"] = *rMosaicLayoutTilesVal.Widget.Text.Content
					}
					if rMosaicLayoutTilesVal.Widget.Text.Format != nil {
						rMosaicLayoutTilesValWidgetText["format"] = string(*rMosaicLayoutTilesVal.Widget.Text.Format)
					}
					rMosaicLayoutTilesValWidget["text"] = rMosaicLayoutTilesValWidgetText
				}
				if rMosaicLayoutTilesVal.Widget.Title != nil {
					rMosaicLayoutTilesValWidget["title"] = *rMosaicLayoutTilesVal.Widget.Title
				}
				if rMosaicLayoutTilesVal.Widget.XyChart != nil && rMosaicLayoutTilesVal.Widget.XyChart != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChart {
					rMosaicLayoutTilesValWidgetXyChart := make(map[string]interface{})
					if rMosaicLayoutTilesVal.Widget.XyChart.ChartOptions != nil && rMosaicLayoutTilesVal.Widget.XyChart.ChartOptions != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartChartOptions {
						rMosaicLayoutTilesValWidgetXyChartChartOptions := make(map[string]interface{})
						if rMosaicLayoutTilesVal.Widget.XyChart.ChartOptions.Mode != nil {
							rMosaicLayoutTilesValWidgetXyChartChartOptions["mode"] = string(*rMosaicLayoutTilesVal.Widget.XyChart.ChartOptions.Mode)
						}
						rMosaicLayoutTilesValWidgetXyChart["chartOptions"] = rMosaicLayoutTilesValWidgetXyChartChartOptions
					}
					var rMosaicLayoutTilesValWidgetXyChartDataSets []interface{}
					for _, rMosaicLayoutTilesValWidgetXyChartDataSetsVal := range rMosaicLayoutTilesVal.Widget.XyChart.DataSets {
						rMosaicLayoutTilesValWidgetXyChartDataSetsObject := make(map[string]interface{})
						if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.LegendTemplate != nil {
							rMosaicLayoutTilesValWidgetXyChartDataSetsObject["legendTemplate"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.LegendTemplate
						}
						if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.MinAlignmentPeriod != nil {
							rMosaicLayoutTilesValWidgetXyChartDataSetsObject["minAlignmentPeriod"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.MinAlignmentPeriod
						}
						if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.PlotType != nil {
							rMosaicLayoutTilesValWidgetXyChartDataSetsObject["plotType"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.PlotType)
						}
						if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery {
							rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQuery := make(map[string]interface{})
							if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
								rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
									}
									var rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
									for _, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["aggregation"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation
								}
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["filter"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter
								}
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
								}
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
									}
									var rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
									for _, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
								}
								rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQuery["timeSeriesFilter"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter
							}
							if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
								rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
										if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
										}
										if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
										}
										var rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
										for _, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
										}
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
										if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
										}
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator
								}
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
										if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
										}
										if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
										}
										var rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
										for _, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
										}
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
										if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
											rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
										}
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator
								}
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
								}
								if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
									}
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
									}
									var rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
									for _, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
									if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
										rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
									}
									rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
								}
								rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQuery["timeSeriesFilterRatio"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio
							}
							if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
								rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQuery["timeSeriesQueryLanguage"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage
							}
							if rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.UnitOverride != nil {
								rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQuery["unitOverride"] = *rMosaicLayoutTilesValWidgetXyChartDataSetsVal.TimeSeriesQuery.UnitOverride
							}
							rMosaicLayoutTilesValWidgetXyChartDataSetsObject["timeSeriesQuery"] = rMosaicLayoutTilesValWidgetXyChartDataSetsValTimeSeriesQuery
						}
						rMosaicLayoutTilesValWidgetXyChartDataSets = append(rMosaicLayoutTilesValWidgetXyChartDataSets, rMosaicLayoutTilesValWidgetXyChartDataSetsObject)
					}
					rMosaicLayoutTilesValWidgetXyChart["dataSets"] = rMosaicLayoutTilesValWidgetXyChartDataSets
					var rMosaicLayoutTilesValWidgetXyChartThresholds []interface{}
					for _, rMosaicLayoutTilesValWidgetXyChartThresholdsVal := range rMosaicLayoutTilesVal.Widget.XyChart.Thresholds {
						rMosaicLayoutTilesValWidgetXyChartThresholdsObject := make(map[string]interface{})
						if rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Color != nil {
							rMosaicLayoutTilesValWidgetXyChartThresholdsObject["color"] = string(*rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Color)
						}
						if rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Direction != nil {
							rMosaicLayoutTilesValWidgetXyChartThresholdsObject["direction"] = string(*rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Direction)
						}
						if rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Label != nil {
							rMosaicLayoutTilesValWidgetXyChartThresholdsObject["label"] = *rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Label
						}
						if rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Value != nil {
							rMosaicLayoutTilesValWidgetXyChartThresholdsObject["value"] = *rMosaicLayoutTilesValWidgetXyChartThresholdsVal.Value
						}
						rMosaicLayoutTilesValWidgetXyChartThresholds = append(rMosaicLayoutTilesValWidgetXyChartThresholds, rMosaicLayoutTilesValWidgetXyChartThresholdsObject)
					}
					rMosaicLayoutTilesValWidgetXyChart["thresholds"] = rMosaicLayoutTilesValWidgetXyChartThresholds
					if rMosaicLayoutTilesVal.Widget.XyChart.TimeshiftDuration != nil {
						rMosaicLayoutTilesValWidgetXyChart["timeshiftDuration"] = *rMosaicLayoutTilesVal.Widget.XyChart.TimeshiftDuration
					}
					if rMosaicLayoutTilesVal.Widget.XyChart.XAxis != nil && rMosaicLayoutTilesVal.Widget.XyChart.XAxis != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartXAxis {
						rMosaicLayoutTilesValWidgetXyChartXAxis := make(map[string]interface{})
						if rMosaicLayoutTilesVal.Widget.XyChart.XAxis.Label != nil {
							rMosaicLayoutTilesValWidgetXyChartXAxis["label"] = *rMosaicLayoutTilesVal.Widget.XyChart.XAxis.Label
						}
						if rMosaicLayoutTilesVal.Widget.XyChart.XAxis.Scale != nil {
							rMosaicLayoutTilesValWidgetXyChartXAxis["scale"] = string(*rMosaicLayoutTilesVal.Widget.XyChart.XAxis.Scale)
						}
						rMosaicLayoutTilesValWidgetXyChart["xAxis"] = rMosaicLayoutTilesValWidgetXyChartXAxis
					}
					if rMosaicLayoutTilesVal.Widget.XyChart.YAxis != nil && rMosaicLayoutTilesVal.Widget.XyChart.YAxis != dclService.EmptyDashboardMosaicLayoutTilesWidgetXyChartYAxis {
						rMosaicLayoutTilesValWidgetXyChartYAxis := make(map[string]interface{})
						if rMosaicLayoutTilesVal.Widget.XyChart.YAxis.Label != nil {
							rMosaicLayoutTilesValWidgetXyChartYAxis["label"] = *rMosaicLayoutTilesVal.Widget.XyChart.YAxis.Label
						}
						if rMosaicLayoutTilesVal.Widget.XyChart.YAxis.Scale != nil {
							rMosaicLayoutTilesValWidgetXyChartYAxis["scale"] = string(*rMosaicLayoutTilesVal.Widget.XyChart.YAxis.Scale)
						}
						rMosaicLayoutTilesValWidgetXyChart["yAxis"] = rMosaicLayoutTilesValWidgetXyChartYAxis
					}
					rMosaicLayoutTilesValWidget["xyChart"] = rMosaicLayoutTilesValWidgetXyChart
				}
				rMosaicLayoutTilesObject["widget"] = rMosaicLayoutTilesValWidget
			}
			if rMosaicLayoutTilesVal.Width != nil {
				rMosaicLayoutTilesObject["width"] = *rMosaicLayoutTilesVal.Width
			}
			if rMosaicLayoutTilesVal.XPos != nil {
				rMosaicLayoutTilesObject["xPos"] = *rMosaicLayoutTilesVal.XPos
			}
			if rMosaicLayoutTilesVal.YPos != nil {
				rMosaicLayoutTilesObject["yPos"] = *rMosaicLayoutTilesVal.YPos
			}
			rMosaicLayoutTiles = append(rMosaicLayoutTiles, rMosaicLayoutTilesObject)
		}
		rMosaicLayout["tiles"] = rMosaicLayoutTiles
		u.Object["mosaicLayout"] = rMosaicLayout
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RowLayout != nil && r.RowLayout != dclService.EmptyDashboardRowLayout {
		rRowLayout := make(map[string]interface{})
		var rRowLayoutRows []interface{}
		for _, rRowLayoutRowsVal := range r.RowLayout.Rows {
			rRowLayoutRowsObject := make(map[string]interface{})
			if rRowLayoutRowsVal.Weight != nil {
				rRowLayoutRowsObject["weight"] = *rRowLayoutRowsVal.Weight
			}
			var rRowLayoutRowsValWidgets []interface{}
			for _, rRowLayoutRowsValWidgetsVal := range rRowLayoutRowsVal.Widgets {
				rRowLayoutRowsValWidgetsObject := make(map[string]interface{})
				if rRowLayoutRowsValWidgetsVal.Blank != nil && rRowLayoutRowsValWidgetsVal.Blank != dclService.EmptyDashboardRowLayoutRowsWidgetsBlank {
					rRowLayoutRowsValWidgetsValBlank := make(map[string]interface{})
					rRowLayoutRowsValWidgetsObject["blank"] = rRowLayoutRowsValWidgetsValBlank
				}
				if rRowLayoutRowsValWidgetsVal.LogsPanel != nil && rRowLayoutRowsValWidgetsVal.LogsPanel != dclService.EmptyDashboardRowLayoutRowsWidgetsLogsPanel {
					rRowLayoutRowsValWidgetsValLogsPanel := make(map[string]interface{})
					if rRowLayoutRowsValWidgetsVal.LogsPanel.Filter != nil {
						rRowLayoutRowsValWidgetsValLogsPanel["filter"] = *rRowLayoutRowsValWidgetsVal.LogsPanel.Filter
					}
					var rRowLayoutRowsValWidgetsValLogsPanelResourceNames []interface{}
					for _, rRowLayoutRowsValWidgetsValLogsPanelResourceNamesVal := range rRowLayoutRowsValWidgetsVal.LogsPanel.ResourceNames {
						rRowLayoutRowsValWidgetsValLogsPanelResourceNames = append(rRowLayoutRowsValWidgetsValLogsPanelResourceNames, rRowLayoutRowsValWidgetsValLogsPanelResourceNamesVal)
					}
					rRowLayoutRowsValWidgetsValLogsPanel["resourceNames"] = rRowLayoutRowsValWidgetsValLogsPanelResourceNames
					rRowLayoutRowsValWidgetsObject["logsPanel"] = rRowLayoutRowsValWidgetsValLogsPanel
				}
				if rRowLayoutRowsValWidgetsVal.Scorecard != nil && rRowLayoutRowsValWidgetsVal.Scorecard != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecard {
					rRowLayoutRowsValWidgetsValScorecard := make(map[string]interface{})
					if rRowLayoutRowsValWidgetsVal.Scorecard.GaugeView != nil && rRowLayoutRowsValWidgetsVal.Scorecard.GaugeView != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardGaugeView {
						rRowLayoutRowsValWidgetsValScorecardGaugeView := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsVal.Scorecard.GaugeView.LowerBound != nil {
							rRowLayoutRowsValWidgetsValScorecardGaugeView["lowerBound"] = *rRowLayoutRowsValWidgetsVal.Scorecard.GaugeView.LowerBound
						}
						if rRowLayoutRowsValWidgetsVal.Scorecard.GaugeView.UpperBound != nil {
							rRowLayoutRowsValWidgetsValScorecardGaugeView["upperBound"] = *rRowLayoutRowsValWidgetsVal.Scorecard.GaugeView.UpperBound
						}
						rRowLayoutRowsValWidgetsValScorecard["gaugeView"] = rRowLayoutRowsValWidgetsValScorecardGaugeView
					}
					if rRowLayoutRowsValWidgetsVal.Scorecard.SparkChartView != nil && rRowLayoutRowsValWidgetsVal.Scorecard.SparkChartView != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardSparkChartView {
						rRowLayoutRowsValWidgetsValScorecardSparkChartView := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod != nil {
							rRowLayoutRowsValWidgetsValScorecardSparkChartView["minAlignmentPeriod"] = *rRowLayoutRowsValWidgetsVal.Scorecard.SparkChartView.MinAlignmentPeriod
						}
						if rRowLayoutRowsValWidgetsVal.Scorecard.SparkChartView.SparkChartType != nil {
							rRowLayoutRowsValWidgetsValScorecardSparkChartView["sparkChartType"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.SparkChartView.SparkChartType)
						}
						rRowLayoutRowsValWidgetsValScorecard["sparkChartView"] = rRowLayoutRowsValWidgetsValScorecardSparkChartView
					}
					var rRowLayoutRowsValWidgetsValScorecardThresholds []interface{}
					for _, rRowLayoutRowsValWidgetsValScorecardThresholdsVal := range rRowLayoutRowsValWidgetsVal.Scorecard.Thresholds {
						rRowLayoutRowsValWidgetsValScorecardThresholdsObject := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Color != nil {
							rRowLayoutRowsValWidgetsValScorecardThresholdsObject["color"] = string(*rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Color)
						}
						if rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Direction != nil {
							rRowLayoutRowsValWidgetsValScorecardThresholdsObject["direction"] = string(*rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Direction)
						}
						if rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Label != nil {
							rRowLayoutRowsValWidgetsValScorecardThresholdsObject["label"] = *rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Label
						}
						if rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Value != nil {
							rRowLayoutRowsValWidgetsValScorecardThresholdsObject["value"] = *rRowLayoutRowsValWidgetsValScorecardThresholdsVal.Value
						}
						rRowLayoutRowsValWidgetsValScorecardThresholds = append(rRowLayoutRowsValWidgetsValScorecardThresholds, rRowLayoutRowsValWidgetsValScorecardThresholdsObject)
					}
					rRowLayoutRowsValWidgetsValScorecard["thresholds"] = rRowLayoutRowsValWidgetsValScorecardThresholds
					if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery {
						rRowLayoutRowsValWidgetsValScorecardTimeSeriesQuery := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter {
							rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
								}
								var rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
								for _, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterAggregation
							}
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["filter"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter
							}
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
							}
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
								}
								var rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
								for _, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
							}
							rRowLayoutRowsValWidgetsValScorecardTimeSeriesQuery["timeSeriesFilter"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilter
						}
						if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio {
							rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
									}
									if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
									}
									var rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
									for _, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
									}
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
									if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
									}
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
							}
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
									}
									if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
									}
									var rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
									for _, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
									}
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
									if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
										rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
									}
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
							}
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
							}
							if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
								}
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
								}
								var rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
								for _, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
								if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
									rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
								}
								rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
							}
							rRowLayoutRowsValWidgetsValScorecardTimeSeriesQuery["timeSeriesFilterRatio"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQueryTimeSeriesFilterRatio
						}
						if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
							rRowLayoutRowsValWidgetsValScorecardTimeSeriesQuery["timeSeriesQueryLanguage"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage
						}
						if rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride != nil {
							rRowLayoutRowsValWidgetsValScorecardTimeSeriesQuery["unitOverride"] = *rRowLayoutRowsValWidgetsVal.Scorecard.TimeSeriesQuery.UnitOverride
						}
						rRowLayoutRowsValWidgetsValScorecard["timeSeriesQuery"] = rRowLayoutRowsValWidgetsValScorecardTimeSeriesQuery
					}
					rRowLayoutRowsValWidgetsObject["scorecard"] = rRowLayoutRowsValWidgetsValScorecard
				}
				if rRowLayoutRowsValWidgetsVal.Text != nil && rRowLayoutRowsValWidgetsVal.Text != dclService.EmptyDashboardRowLayoutRowsWidgetsText {
					rRowLayoutRowsValWidgetsValText := make(map[string]interface{})
					if rRowLayoutRowsValWidgetsVal.Text.Content != nil {
						rRowLayoutRowsValWidgetsValText["content"] = *rRowLayoutRowsValWidgetsVal.Text.Content
					}
					if rRowLayoutRowsValWidgetsVal.Text.Format != nil {
						rRowLayoutRowsValWidgetsValText["format"] = string(*rRowLayoutRowsValWidgetsVal.Text.Format)
					}
					rRowLayoutRowsValWidgetsObject["text"] = rRowLayoutRowsValWidgetsValText
				}
				if rRowLayoutRowsValWidgetsVal.Title != nil {
					rRowLayoutRowsValWidgetsObject["title"] = *rRowLayoutRowsValWidgetsVal.Title
				}
				if rRowLayoutRowsValWidgetsVal.XyChart != nil && rRowLayoutRowsValWidgetsVal.XyChart != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChart {
					rRowLayoutRowsValWidgetsValXyChart := make(map[string]interface{})
					if rRowLayoutRowsValWidgetsVal.XyChart.ChartOptions != nil && rRowLayoutRowsValWidgetsVal.XyChart.ChartOptions != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartChartOptions {
						rRowLayoutRowsValWidgetsValXyChartChartOptions := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsVal.XyChart.ChartOptions.Mode != nil {
							rRowLayoutRowsValWidgetsValXyChartChartOptions["mode"] = string(*rRowLayoutRowsValWidgetsVal.XyChart.ChartOptions.Mode)
						}
						rRowLayoutRowsValWidgetsValXyChart["chartOptions"] = rRowLayoutRowsValWidgetsValXyChartChartOptions
					}
					var rRowLayoutRowsValWidgetsValXyChartDataSets []interface{}
					for _, rRowLayoutRowsValWidgetsValXyChartDataSetsVal := range rRowLayoutRowsValWidgetsVal.XyChart.DataSets {
						rRowLayoutRowsValWidgetsValXyChartDataSetsObject := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.LegendTemplate != nil {
							rRowLayoutRowsValWidgetsValXyChartDataSetsObject["legendTemplate"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.LegendTemplate
						}
						if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.MinAlignmentPeriod != nil {
							rRowLayoutRowsValWidgetsValXyChartDataSetsObject["minAlignmentPeriod"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.MinAlignmentPeriod
						}
						if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.PlotType != nil {
							rRowLayoutRowsValWidgetsValXyChartDataSetsObject["plotType"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.PlotType)
						}
						if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery {
							rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQuery := make(map[string]interface{})
							if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
								rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer)
									}
									var rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields []interface{}
									for _, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFieldsVal)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregationGroupByFields
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["aggregation"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterAggregation
								}
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter != nil {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["filter"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.Filter
								}
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction)
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
								}
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer)
									}
									var rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields []interface{}
									for _, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFieldsVal)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregationGroupByFields
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
								}
								rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilter"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilter
							}
							if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
								rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio := make(map[string]interface{})
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation := make(map[string]interface{})
										if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod != nil {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod
										}
										if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer != nil {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer)
										}
										var rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields []interface{}
										for _, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFieldsVal)
										}
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationGroupByFields
										if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner != nil {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner)
										}
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["denominator"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioDenominator
								}
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation := make(map[string]interface{})
										if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod != nil {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod
										}
										if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer != nil {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer)
										}
										var rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields []interface{}
										for _, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFieldsVal)
										}
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationGroupByFields
										if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner != nil {
											rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner)
										}
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["numerator"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioNumerator
								}
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction)
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
								}
								if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != nil && rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation := make(map[string]interface{})
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod
									}
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer)
									}
									var rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields []interface{}
									for _, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal := range rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields = append(rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields, rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFieldsVal)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationGroupByFields
									if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner != nil {
										rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"] = string(*rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner)
									}
									rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
								}
								rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesFilterRatio"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQueryTimeSeriesFilterRatio
							}
							if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage != nil {
								rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQuery["timeSeriesQueryLanguage"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.TimeSeriesQueryLanguage
							}
							if rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride != nil {
								rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQuery["unitOverride"] = *rRowLayoutRowsValWidgetsValXyChartDataSetsVal.TimeSeriesQuery.UnitOverride
							}
							rRowLayoutRowsValWidgetsValXyChartDataSetsObject["timeSeriesQuery"] = rRowLayoutRowsValWidgetsValXyChartDataSetsValTimeSeriesQuery
						}
						rRowLayoutRowsValWidgetsValXyChartDataSets = append(rRowLayoutRowsValWidgetsValXyChartDataSets, rRowLayoutRowsValWidgetsValXyChartDataSetsObject)
					}
					rRowLayoutRowsValWidgetsValXyChart["dataSets"] = rRowLayoutRowsValWidgetsValXyChartDataSets
					var rRowLayoutRowsValWidgetsValXyChartThresholds []interface{}
					for _, rRowLayoutRowsValWidgetsValXyChartThresholdsVal := range rRowLayoutRowsValWidgetsVal.XyChart.Thresholds {
						rRowLayoutRowsValWidgetsValXyChartThresholdsObject := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Color != nil {
							rRowLayoutRowsValWidgetsValXyChartThresholdsObject["color"] = string(*rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Color)
						}
						if rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Direction != nil {
							rRowLayoutRowsValWidgetsValXyChartThresholdsObject["direction"] = string(*rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Direction)
						}
						if rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Label != nil {
							rRowLayoutRowsValWidgetsValXyChartThresholdsObject["label"] = *rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Label
						}
						if rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Value != nil {
							rRowLayoutRowsValWidgetsValXyChartThresholdsObject["value"] = *rRowLayoutRowsValWidgetsValXyChartThresholdsVal.Value
						}
						rRowLayoutRowsValWidgetsValXyChartThresholds = append(rRowLayoutRowsValWidgetsValXyChartThresholds, rRowLayoutRowsValWidgetsValXyChartThresholdsObject)
					}
					rRowLayoutRowsValWidgetsValXyChart["thresholds"] = rRowLayoutRowsValWidgetsValXyChartThresholds
					if rRowLayoutRowsValWidgetsVal.XyChart.TimeshiftDuration != nil {
						rRowLayoutRowsValWidgetsValXyChart["timeshiftDuration"] = *rRowLayoutRowsValWidgetsVal.XyChart.TimeshiftDuration
					}
					if rRowLayoutRowsValWidgetsVal.XyChart.XAxis != nil && rRowLayoutRowsValWidgetsVal.XyChart.XAxis != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartXAxis {
						rRowLayoutRowsValWidgetsValXyChartXAxis := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsVal.XyChart.XAxis.Label != nil {
							rRowLayoutRowsValWidgetsValXyChartXAxis["label"] = *rRowLayoutRowsValWidgetsVal.XyChart.XAxis.Label
						}
						if rRowLayoutRowsValWidgetsVal.XyChart.XAxis.Scale != nil {
							rRowLayoutRowsValWidgetsValXyChartXAxis["scale"] = string(*rRowLayoutRowsValWidgetsVal.XyChart.XAxis.Scale)
						}
						rRowLayoutRowsValWidgetsValXyChart["xAxis"] = rRowLayoutRowsValWidgetsValXyChartXAxis
					}
					if rRowLayoutRowsValWidgetsVal.XyChart.YAxis != nil && rRowLayoutRowsValWidgetsVal.XyChart.YAxis != dclService.EmptyDashboardRowLayoutRowsWidgetsXyChartYAxis {
						rRowLayoutRowsValWidgetsValXyChartYAxis := make(map[string]interface{})
						if rRowLayoutRowsValWidgetsVal.XyChart.YAxis.Label != nil {
							rRowLayoutRowsValWidgetsValXyChartYAxis["label"] = *rRowLayoutRowsValWidgetsVal.XyChart.YAxis.Label
						}
						if rRowLayoutRowsValWidgetsVal.XyChart.YAxis.Scale != nil {
							rRowLayoutRowsValWidgetsValXyChartYAxis["scale"] = string(*rRowLayoutRowsValWidgetsVal.XyChart.YAxis.Scale)
						}
						rRowLayoutRowsValWidgetsValXyChart["yAxis"] = rRowLayoutRowsValWidgetsValXyChartYAxis
					}
					rRowLayoutRowsValWidgetsObject["xyChart"] = rRowLayoutRowsValWidgetsValXyChart
				}
				rRowLayoutRowsValWidgets = append(rRowLayoutRowsValWidgets, rRowLayoutRowsValWidgetsObject)
			}
			rRowLayoutRowsObject["widgets"] = rRowLayoutRowsValWidgets
			rRowLayoutRows = append(rRowLayoutRows, rRowLayoutRowsObject)
		}
		rRowLayout["rows"] = rRowLayoutRows
		u.Object["rowLayout"] = rRowLayout
	}
	return u
}

func UnstructuredToDashboard(u *unstructured.Resource) (*dclService.Dashboard, error) {
	r := &dclService.Dashboard{}
	if _, ok := u.Object["columnLayout"]; ok {
		if rColumnLayout, ok := u.Object["columnLayout"].(map[string]interface{}); ok {
			r.ColumnLayout = &dclService.DashboardColumnLayout{}
			if _, ok := rColumnLayout["columns"]; ok {
				if s, ok := rColumnLayout["columns"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rColumnLayoutColumns dclService.DashboardColumnLayoutColumns
							if _, ok := objval["weight"]; ok {
								if i, ok := objval["weight"].(int64); ok {
									rColumnLayoutColumns.Weight = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rColumnLayoutColumns.Weight: expected int64")
								}
							}
							if _, ok := objval["widgets"]; ok {
								if s, ok := objval["widgets"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rColumnLayoutColumnsWidgets dclService.DashboardColumnLayoutColumnsWidgets
											if _, ok := objval["blank"]; ok {
												if _, ok := objval["blank"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.Blank = &dclService.DashboardColumnLayoutColumnsWidgetsBlank{}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Blank: expected map[string]interface{}")
												}
											}
											if _, ok := objval["logsPanel"]; ok {
												if rColumnLayoutColumnsWidgetsLogsPanel, ok := objval["logsPanel"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.LogsPanel = &dclService.DashboardColumnLayoutColumnsWidgetsLogsPanel{}
													if _, ok := rColumnLayoutColumnsWidgetsLogsPanel["filter"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsLogsPanel["filter"].(string); ok {
															rColumnLayoutColumnsWidgets.LogsPanel.Filter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.LogsPanel.Filter: expected string")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsLogsPanel["resourceNames"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsLogsPanel["resourceNames"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rColumnLayoutColumnsWidgets.LogsPanel.ResourceNames = append(rColumnLayoutColumnsWidgets.LogsPanel.ResourceNames, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.LogsPanel.ResourceNames: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.LogsPanel: expected map[string]interface{}")
												}
											}
											if _, ok := objval["scorecard"]; ok {
												if rColumnLayoutColumnsWidgetsScorecard, ok := objval["scorecard"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.Scorecard = &dclService.DashboardColumnLayoutColumnsWidgetsScorecard{}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["gaugeView"]; ok {
														if rColumnLayoutColumnsWidgetsScorecardGaugeView, ok := rColumnLayoutColumnsWidgetsScorecard["gaugeView"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.Scorecard.GaugeView = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardGaugeView{}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["lowerBound"]; ok {
																if f, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["lowerBound"].(float64); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.GaugeView.LowerBound = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.GaugeView.LowerBound: expected float64")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["upperBound"]; ok {
																if f, ok := rColumnLayoutColumnsWidgetsScorecardGaugeView["upperBound"].(float64); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.GaugeView.UpperBound = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.GaugeView.UpperBound: expected float64")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.GaugeView: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["sparkChartView"]; ok {
														if rColumnLayoutColumnsWidgetsScorecardSparkChartView, ok := rColumnLayoutColumnsWidgetsScorecard["sparkChartView"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.Scorecard.SparkChartView = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardSparkChartView{}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["minAlignmentPeriod"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["minAlignmentPeriod"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.MinAlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.MinAlignmentPeriod: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["sparkChartType"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardSparkChartView["sparkChartType"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.SparkChartType = dclService.DashboardColumnLayoutColumnsWidgetsScorecardSparkChartViewSparkChartTypeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.SparkChartView.SparkChartType: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.SparkChartView: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["thresholds"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsScorecard["thresholds"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rColumnLayoutColumnsWidgetsScorecardThresholds dclService.DashboardColumnLayoutColumnsWidgetsScorecardThresholds
																	if _, ok := objval["color"]; ok {
																		if s, ok := objval["color"].(string); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Color = dclService.DashboardColumnLayoutColumnsWidgetsScorecardThresholdsColorEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Color: expected string")
																		}
																	}
																	if _, ok := objval["direction"]; ok {
																		if s, ok := objval["direction"].(string); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Direction = dclService.DashboardColumnLayoutColumnsWidgetsScorecardThresholdsDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Direction: expected string")
																		}
																	}
																	if _, ok := objval["label"]; ok {
																		if s, ok := objval["label"].(string); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Label = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Label: expected string")
																		}
																	}
																	if _, ok := objval["value"]; ok {
																		if f, ok := objval["value"].(float64); ok {
																			rColumnLayoutColumnsWidgetsScorecardThresholds.Value = dcl.Float64(f)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsScorecardThresholds.Value: expected float64")
																		}
																	}
																	rColumnLayoutColumnsWidgets.Scorecard.Thresholds = append(rColumnLayoutColumnsWidgets.Scorecard.Thresholds, rColumnLayoutColumnsWidgetsScorecardThresholds)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.Thresholds: expected []interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsScorecard["timeSeriesQuery"]; ok {
														if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery, ok := rColumnLayoutColumnsWidgetsScorecard["timeSeriesQuery"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery{}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"]; ok {
																if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter{}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																		if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																				if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																				if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																		if rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["unitOverride"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsScorecardTimeSeriesQuery["unitOverride"].(string); ok {
																	rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.UnitOverride = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery.UnitOverride: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard.TimeSeriesQuery: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Scorecard: expected map[string]interface{}")
												}
											}
											if _, ok := objval["text"]; ok {
												if rColumnLayoutColumnsWidgetsText, ok := objval["text"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.Text = &dclService.DashboardColumnLayoutColumnsWidgetsText{}
													if _, ok := rColumnLayoutColumnsWidgetsText["content"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsText["content"].(string); ok {
															rColumnLayoutColumnsWidgets.Text.Content = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Text.Content: expected string")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsText["format"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsText["format"].(string); ok {
															rColumnLayoutColumnsWidgets.Text.Format = dclService.DashboardColumnLayoutColumnsWidgetsTextFormatEnumRef(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Text.Format: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Text: expected map[string]interface{}")
												}
											}
											if _, ok := objval["title"]; ok {
												if s, ok := objval["title"].(string); ok {
													rColumnLayoutColumnsWidgets.Title = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.Title: expected string")
												}
											}
											if _, ok := objval["xyChart"]; ok {
												if rColumnLayoutColumnsWidgetsXyChart, ok := objval["xyChart"].(map[string]interface{}); ok {
													rColumnLayoutColumnsWidgets.XyChart = &dclService.DashboardColumnLayoutColumnsWidgetsXyChart{}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["chartOptions"]; ok {
														if rColumnLayoutColumnsWidgetsXyChartChartOptions, ok := rColumnLayoutColumnsWidgetsXyChart["chartOptions"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.XyChart.ChartOptions = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartChartOptions{}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartChartOptions["mode"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartChartOptions["mode"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.ChartOptions.Mode = dclService.DashboardColumnLayoutColumnsWidgetsXyChartChartOptionsModeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.ChartOptions.Mode: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.ChartOptions: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["dataSets"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsXyChart["dataSets"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rColumnLayoutColumnsWidgetsXyChartDataSets dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSets
																	if _, ok := objval["legendTemplate"]; ok {
																		if s, ok := objval["legendTemplate"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.LegendTemplate = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.LegendTemplate: expected string")
																		}
																	}
																	if _, ok := objval["minAlignmentPeriod"]; ok {
																		if s, ok := objval["minAlignmentPeriod"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.MinAlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.MinAlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := objval["plotType"]; ok {
																		if s, ok := objval["plotType"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.PlotType = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsPlotTypeEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.PlotType: expected string")
																		}
																	}
																	if _, ok := objval["timeSeriesQuery"]; ok {
																		if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery, ok := objval["timeSeriesQuery"].(map[string]interface{}); ok {
																			rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery{}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"]; ok {
																				if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																						if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																								if i, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																				if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																								if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																											for _, ss := range s {
																												if strval, ok := ss.(string); ok {
																													rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																								if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																											for _, ss := range s {
																												if strval, ok := ss.(string); ok {
																													rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																										}
																									}
																									if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																										if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																								if i, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																						if rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																							rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																								if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																									rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																				}
																			}
																			if _, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"]; ok {
																				if s, ok := rColumnLayoutColumnsWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"].(string); ok {
																					rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartDataSets.TimeSeriesQuery: expected map[string]interface{}")
																		}
																	}
																	rColumnLayoutColumnsWidgets.XyChart.DataSets = append(rColumnLayoutColumnsWidgets.XyChart.DataSets, rColumnLayoutColumnsWidgetsXyChartDataSets)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.DataSets: expected []interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["thresholds"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsXyChart["thresholds"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rColumnLayoutColumnsWidgetsXyChartThresholds dclService.DashboardColumnLayoutColumnsWidgetsXyChartThresholds
																	if _, ok := objval["color"]; ok {
																		if s, ok := objval["color"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Color = dclService.DashboardColumnLayoutColumnsWidgetsXyChartThresholdsColorEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Color: expected string")
																		}
																	}
																	if _, ok := objval["direction"]; ok {
																		if s, ok := objval["direction"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Direction = dclService.DashboardColumnLayoutColumnsWidgetsXyChartThresholdsDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Direction: expected string")
																		}
																	}
																	if _, ok := objval["label"]; ok {
																		if s, ok := objval["label"].(string); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Label = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Label: expected string")
																		}
																	}
																	if _, ok := objval["value"]; ok {
																		if f, ok := objval["value"].(float64); ok {
																			rColumnLayoutColumnsWidgetsXyChartThresholds.Value = dcl.Float64(f)
																		} else {
																			return nil, fmt.Errorf("rColumnLayoutColumnsWidgetsXyChartThresholds.Value: expected float64")
																		}
																	}
																	rColumnLayoutColumnsWidgets.XyChart.Thresholds = append(rColumnLayoutColumnsWidgets.XyChart.Thresholds, rColumnLayoutColumnsWidgetsXyChartThresholds)
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.Thresholds: expected []interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["timeshiftDuration"]; ok {
														if s, ok := rColumnLayoutColumnsWidgetsXyChart["timeshiftDuration"].(string); ok {
															rColumnLayoutColumnsWidgets.XyChart.TimeshiftDuration = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.TimeshiftDuration: expected string")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["xAxis"]; ok {
														if rColumnLayoutColumnsWidgetsXyChartXAxis, ok := rColumnLayoutColumnsWidgetsXyChart["xAxis"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.XyChart.XAxis = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartXAxis{}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["label"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["label"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.XAxis.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.XAxis.Label: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["scale"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartXAxis["scale"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.XAxis.Scale = dclService.DashboardColumnLayoutColumnsWidgetsXyChartXAxisScaleEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.XAxis.Scale: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.XAxis: expected map[string]interface{}")
														}
													}
													if _, ok := rColumnLayoutColumnsWidgetsXyChart["yAxis"]; ok {
														if rColumnLayoutColumnsWidgetsXyChartYAxis, ok := rColumnLayoutColumnsWidgetsXyChart["yAxis"].(map[string]interface{}); ok {
															rColumnLayoutColumnsWidgets.XyChart.YAxis = &dclService.DashboardColumnLayoutColumnsWidgetsXyChartYAxis{}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["label"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["label"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.YAxis.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.YAxis.Label: expected string")
																}
															}
															if _, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["scale"]; ok {
																if s, ok := rColumnLayoutColumnsWidgetsXyChartYAxis["scale"].(string); ok {
																	rColumnLayoutColumnsWidgets.XyChart.YAxis.Scale = dclService.DashboardColumnLayoutColumnsWidgetsXyChartYAxisScaleEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.YAxis.Scale: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart.YAxis: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rColumnLayoutColumnsWidgets.XyChart: expected map[string]interface{}")
												}
											}
											rColumnLayoutColumns.Widgets = append(rColumnLayoutColumns.Widgets, rColumnLayoutColumnsWidgets)
										}
									}
								} else {
									return nil, fmt.Errorf("rColumnLayoutColumns.Widgets: expected []interface{}")
								}
							}
							r.ColumnLayout.Columns = append(r.ColumnLayout.Columns, rColumnLayoutColumns)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ColumnLayout.Columns: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ColumnLayout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["gridLayout"]; ok {
		if rGridLayout, ok := u.Object["gridLayout"].(map[string]interface{}); ok {
			r.GridLayout = &dclService.DashboardGridLayout{}
			if _, ok := rGridLayout["columns"]; ok {
				if i, ok := rGridLayout["columns"].(int64); ok {
					r.GridLayout.Columns = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.GridLayout.Columns: expected int64")
				}
			}
			if _, ok := rGridLayout["widgets"]; ok {
				if s, ok := rGridLayout["widgets"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rGridLayoutWidgets dclService.DashboardGridLayoutWidgets
							if _, ok := objval["blank"]; ok {
								if _, ok := objval["blank"].(map[string]interface{}); ok {
									rGridLayoutWidgets.Blank = &dclService.DashboardGridLayoutWidgetsBlank{}
								} else {
									return nil, fmt.Errorf("rGridLayoutWidgets.Blank: expected map[string]interface{}")
								}
							}
							if _, ok := objval["logsPanel"]; ok {
								if rGridLayoutWidgetsLogsPanel, ok := objval["logsPanel"].(map[string]interface{}); ok {
									rGridLayoutWidgets.LogsPanel = &dclService.DashboardGridLayoutWidgetsLogsPanel{}
									if _, ok := rGridLayoutWidgetsLogsPanel["filter"]; ok {
										if s, ok := rGridLayoutWidgetsLogsPanel["filter"].(string); ok {
											rGridLayoutWidgets.LogsPanel.Filter = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.LogsPanel.Filter: expected string")
										}
									}
									if _, ok := rGridLayoutWidgetsLogsPanel["resourceNames"]; ok {
										if s, ok := rGridLayoutWidgetsLogsPanel["resourceNames"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rGridLayoutWidgets.LogsPanel.ResourceNames = append(rGridLayoutWidgets.LogsPanel.ResourceNames, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.LogsPanel.ResourceNames: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rGridLayoutWidgets.LogsPanel: expected map[string]interface{}")
								}
							}
							if _, ok := objval["scorecard"]; ok {
								if rGridLayoutWidgetsScorecard, ok := objval["scorecard"].(map[string]interface{}); ok {
									rGridLayoutWidgets.Scorecard = &dclService.DashboardGridLayoutWidgetsScorecard{}
									if _, ok := rGridLayoutWidgetsScorecard["gaugeView"]; ok {
										if rGridLayoutWidgetsScorecardGaugeView, ok := rGridLayoutWidgetsScorecard["gaugeView"].(map[string]interface{}); ok {
											rGridLayoutWidgets.Scorecard.GaugeView = &dclService.DashboardGridLayoutWidgetsScorecardGaugeView{}
											if _, ok := rGridLayoutWidgetsScorecardGaugeView["lowerBound"]; ok {
												if f, ok := rGridLayoutWidgetsScorecardGaugeView["lowerBound"].(float64); ok {
													rGridLayoutWidgets.Scorecard.GaugeView.LowerBound = dcl.Float64(f)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.GaugeView.LowerBound: expected float64")
												}
											}
											if _, ok := rGridLayoutWidgetsScorecardGaugeView["upperBound"]; ok {
												if f, ok := rGridLayoutWidgetsScorecardGaugeView["upperBound"].(float64); ok {
													rGridLayoutWidgets.Scorecard.GaugeView.UpperBound = dcl.Float64(f)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.GaugeView.UpperBound: expected float64")
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.GaugeView: expected map[string]interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsScorecard["sparkChartView"]; ok {
										if rGridLayoutWidgetsScorecardSparkChartView, ok := rGridLayoutWidgetsScorecard["sparkChartView"].(map[string]interface{}); ok {
											rGridLayoutWidgets.Scorecard.SparkChartView = &dclService.DashboardGridLayoutWidgetsScorecardSparkChartView{}
											if _, ok := rGridLayoutWidgetsScorecardSparkChartView["minAlignmentPeriod"]; ok {
												if s, ok := rGridLayoutWidgetsScorecardSparkChartView["minAlignmentPeriod"].(string); ok {
													rGridLayoutWidgets.Scorecard.SparkChartView.MinAlignmentPeriod = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.SparkChartView.MinAlignmentPeriod: expected string")
												}
											}
											if _, ok := rGridLayoutWidgetsScorecardSparkChartView["sparkChartType"]; ok {
												if s, ok := rGridLayoutWidgetsScorecardSparkChartView["sparkChartType"].(string); ok {
													rGridLayoutWidgets.Scorecard.SparkChartView.SparkChartType = dclService.DashboardGridLayoutWidgetsScorecardSparkChartViewSparkChartTypeEnumRef(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.SparkChartView.SparkChartType: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.SparkChartView: expected map[string]interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsScorecard["thresholds"]; ok {
										if s, ok := rGridLayoutWidgetsScorecard["thresholds"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rGridLayoutWidgetsScorecardThresholds dclService.DashboardGridLayoutWidgetsScorecardThresholds
													if _, ok := objval["color"]; ok {
														if s, ok := objval["color"].(string); ok {
															rGridLayoutWidgetsScorecardThresholds.Color = dclService.DashboardGridLayoutWidgetsScorecardThresholdsColorEnumRef(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsScorecardThresholds.Color: expected string")
														}
													}
													if _, ok := objval["direction"]; ok {
														if s, ok := objval["direction"].(string); ok {
															rGridLayoutWidgetsScorecardThresholds.Direction = dclService.DashboardGridLayoutWidgetsScorecardThresholdsDirectionEnumRef(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsScorecardThresholds.Direction: expected string")
														}
													}
													if _, ok := objval["label"]; ok {
														if s, ok := objval["label"].(string); ok {
															rGridLayoutWidgetsScorecardThresholds.Label = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsScorecardThresholds.Label: expected string")
														}
													}
													if _, ok := objval["value"]; ok {
														if f, ok := objval["value"].(float64); ok {
															rGridLayoutWidgetsScorecardThresholds.Value = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsScorecardThresholds.Value: expected float64")
														}
													}
													rGridLayoutWidgets.Scorecard.Thresholds = append(rGridLayoutWidgets.Scorecard.Thresholds, rGridLayoutWidgetsScorecardThresholds)
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.Thresholds: expected []interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsScorecard["timeSeriesQuery"]; ok {
										if rGridLayoutWidgetsScorecardTimeSeriesQuery, ok := rGridLayoutWidgetsScorecard["timeSeriesQuery"].(map[string]interface{}); ok {
											rGridLayoutWidgets.Scorecard.TimeSeriesQuery = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQuery{}
											if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"]; ok {
												if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
													rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter{}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
														}
													}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
														if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
														}
													}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																if i, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
														}
													}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
												}
											}
											if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
												if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
													rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
														}
													}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																		if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
														}
													}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																if i, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
														}
													}
													if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
														if rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
															rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																	rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
												}
											}
											if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
												if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
													rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
												}
											}
											if _, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["unitOverride"]; ok {
												if s, ok := rGridLayoutWidgetsScorecardTimeSeriesQuery["unitOverride"].(string); ok {
													rGridLayoutWidgets.Scorecard.TimeSeriesQuery.UnitOverride = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery.UnitOverride: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard.TimeSeriesQuery: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rGridLayoutWidgets.Scorecard: expected map[string]interface{}")
								}
							}
							if _, ok := objval["text"]; ok {
								if rGridLayoutWidgetsText, ok := objval["text"].(map[string]interface{}); ok {
									rGridLayoutWidgets.Text = &dclService.DashboardGridLayoutWidgetsText{}
									if _, ok := rGridLayoutWidgetsText["content"]; ok {
										if s, ok := rGridLayoutWidgetsText["content"].(string); ok {
											rGridLayoutWidgets.Text.Content = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.Text.Content: expected string")
										}
									}
									if _, ok := rGridLayoutWidgetsText["format"]; ok {
										if s, ok := rGridLayoutWidgetsText["format"].(string); ok {
											rGridLayoutWidgets.Text.Format = dclService.DashboardGridLayoutWidgetsTextFormatEnumRef(s)
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.Text.Format: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rGridLayoutWidgets.Text: expected map[string]interface{}")
								}
							}
							if _, ok := objval["title"]; ok {
								if s, ok := objval["title"].(string); ok {
									rGridLayoutWidgets.Title = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rGridLayoutWidgets.Title: expected string")
								}
							}
							if _, ok := objval["xyChart"]; ok {
								if rGridLayoutWidgetsXyChart, ok := objval["xyChart"].(map[string]interface{}); ok {
									rGridLayoutWidgets.XyChart = &dclService.DashboardGridLayoutWidgetsXyChart{}
									if _, ok := rGridLayoutWidgetsXyChart["chartOptions"]; ok {
										if rGridLayoutWidgetsXyChartChartOptions, ok := rGridLayoutWidgetsXyChart["chartOptions"].(map[string]interface{}); ok {
											rGridLayoutWidgets.XyChart.ChartOptions = &dclService.DashboardGridLayoutWidgetsXyChartChartOptions{}
											if _, ok := rGridLayoutWidgetsXyChartChartOptions["mode"]; ok {
												if s, ok := rGridLayoutWidgetsXyChartChartOptions["mode"].(string); ok {
													rGridLayoutWidgets.XyChart.ChartOptions.Mode = dclService.DashboardGridLayoutWidgetsXyChartChartOptionsModeEnumRef(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.ChartOptions.Mode: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.ChartOptions: expected map[string]interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsXyChart["dataSets"]; ok {
										if s, ok := rGridLayoutWidgetsXyChart["dataSets"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rGridLayoutWidgetsXyChartDataSets dclService.DashboardGridLayoutWidgetsXyChartDataSets
													if _, ok := objval["legendTemplate"]; ok {
														if s, ok := objval["legendTemplate"].(string); ok {
															rGridLayoutWidgetsXyChartDataSets.LegendTemplate = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.LegendTemplate: expected string")
														}
													}
													if _, ok := objval["minAlignmentPeriod"]; ok {
														if s, ok := objval["minAlignmentPeriod"].(string); ok {
															rGridLayoutWidgetsXyChartDataSets.MinAlignmentPeriod = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.MinAlignmentPeriod: expected string")
														}
													}
													if _, ok := objval["plotType"]; ok {
														if s, ok := objval["plotType"].(string); ok {
															rGridLayoutWidgetsXyChartDataSets.PlotType = dclService.DashboardGridLayoutWidgetsXyChartDataSetsPlotTypeEnumRef(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.PlotType: expected string")
														}
													}
													if _, ok := objval["timeSeriesQuery"]; ok {
														if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery, ok := objval["timeSeriesQuery"].(map[string]interface{}); ok {
															rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery{}
															if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"]; ok {
																if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																	rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																		if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																	rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																				if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																				if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																		if rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																			rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardGridLayoutWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																}
															}
															if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																	rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																}
															}
															if _, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"]; ok {
																if s, ok := rGridLayoutWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"].(string); ok {
																	rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartDataSets.TimeSeriesQuery: expected map[string]interface{}")
														}
													}
													rGridLayoutWidgets.XyChart.DataSets = append(rGridLayoutWidgets.XyChart.DataSets, rGridLayoutWidgetsXyChartDataSets)
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.DataSets: expected []interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsXyChart["thresholds"]; ok {
										if s, ok := rGridLayoutWidgetsXyChart["thresholds"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rGridLayoutWidgetsXyChartThresholds dclService.DashboardGridLayoutWidgetsXyChartThresholds
													if _, ok := objval["color"]; ok {
														if s, ok := objval["color"].(string); ok {
															rGridLayoutWidgetsXyChartThresholds.Color = dclService.DashboardGridLayoutWidgetsXyChartThresholdsColorEnumRef(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartThresholds.Color: expected string")
														}
													}
													if _, ok := objval["direction"]; ok {
														if s, ok := objval["direction"].(string); ok {
															rGridLayoutWidgetsXyChartThresholds.Direction = dclService.DashboardGridLayoutWidgetsXyChartThresholdsDirectionEnumRef(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartThresholds.Direction: expected string")
														}
													}
													if _, ok := objval["label"]; ok {
														if s, ok := objval["label"].(string); ok {
															rGridLayoutWidgetsXyChartThresholds.Label = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartThresholds.Label: expected string")
														}
													}
													if _, ok := objval["value"]; ok {
														if f, ok := objval["value"].(float64); ok {
															rGridLayoutWidgetsXyChartThresholds.Value = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("rGridLayoutWidgetsXyChartThresholds.Value: expected float64")
														}
													}
													rGridLayoutWidgets.XyChart.Thresholds = append(rGridLayoutWidgets.XyChart.Thresholds, rGridLayoutWidgetsXyChartThresholds)
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.Thresholds: expected []interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsXyChart["timeshiftDuration"]; ok {
										if s, ok := rGridLayoutWidgetsXyChart["timeshiftDuration"].(string); ok {
											rGridLayoutWidgets.XyChart.TimeshiftDuration = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.TimeshiftDuration: expected string")
										}
									}
									if _, ok := rGridLayoutWidgetsXyChart["xAxis"]; ok {
										if rGridLayoutWidgetsXyChartXAxis, ok := rGridLayoutWidgetsXyChart["xAxis"].(map[string]interface{}); ok {
											rGridLayoutWidgets.XyChart.XAxis = &dclService.DashboardGridLayoutWidgetsXyChartXAxis{}
											if _, ok := rGridLayoutWidgetsXyChartXAxis["label"]; ok {
												if s, ok := rGridLayoutWidgetsXyChartXAxis["label"].(string); ok {
													rGridLayoutWidgets.XyChart.XAxis.Label = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.XAxis.Label: expected string")
												}
											}
											if _, ok := rGridLayoutWidgetsXyChartXAxis["scale"]; ok {
												if s, ok := rGridLayoutWidgetsXyChartXAxis["scale"].(string); ok {
													rGridLayoutWidgets.XyChart.XAxis.Scale = dclService.DashboardGridLayoutWidgetsXyChartXAxisScaleEnumRef(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.XAxis.Scale: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.XAxis: expected map[string]interface{}")
										}
									}
									if _, ok := rGridLayoutWidgetsXyChart["yAxis"]; ok {
										if rGridLayoutWidgetsXyChartYAxis, ok := rGridLayoutWidgetsXyChart["yAxis"].(map[string]interface{}); ok {
											rGridLayoutWidgets.XyChart.YAxis = &dclService.DashboardGridLayoutWidgetsXyChartYAxis{}
											if _, ok := rGridLayoutWidgetsXyChartYAxis["label"]; ok {
												if s, ok := rGridLayoutWidgetsXyChartYAxis["label"].(string); ok {
													rGridLayoutWidgets.XyChart.YAxis.Label = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.YAxis.Label: expected string")
												}
											}
											if _, ok := rGridLayoutWidgetsXyChartYAxis["scale"]; ok {
												if s, ok := rGridLayoutWidgetsXyChartYAxis["scale"].(string); ok {
													rGridLayoutWidgets.XyChart.YAxis.Scale = dclService.DashboardGridLayoutWidgetsXyChartYAxisScaleEnumRef(s)
												} else {
													return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.YAxis.Scale: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rGridLayoutWidgets.XyChart.YAxis: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rGridLayoutWidgets.XyChart: expected map[string]interface{}")
								}
							}
							r.GridLayout.Widgets = append(r.GridLayout.Widgets, rGridLayoutWidgets)
						}
					}
				} else {
					return nil, fmt.Errorf("r.GridLayout.Widgets: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.GridLayout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["mosaicLayout"]; ok {
		if rMosaicLayout, ok := u.Object["mosaicLayout"].(map[string]interface{}); ok {
			r.MosaicLayout = &dclService.DashboardMosaicLayout{}
			if _, ok := rMosaicLayout["columns"]; ok {
				if i, ok := rMosaicLayout["columns"].(int64); ok {
					r.MosaicLayout.Columns = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.MosaicLayout.Columns: expected int64")
				}
			}
			if _, ok := rMosaicLayout["tiles"]; ok {
				if s, ok := rMosaicLayout["tiles"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMosaicLayoutTiles dclService.DashboardMosaicLayoutTiles
							if _, ok := objval["height"]; ok {
								if i, ok := objval["height"].(int64); ok {
									rMosaicLayoutTiles.Height = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.Height: expected int64")
								}
							}
							if _, ok := objval["widget"]; ok {
								if rMosaicLayoutTilesWidget, ok := objval["widget"].(map[string]interface{}); ok {
									rMosaicLayoutTiles.Widget = &dclService.DashboardMosaicLayoutTilesWidget{}
									if _, ok := rMosaicLayoutTilesWidget["blank"]; ok {
										if _, ok := rMosaicLayoutTilesWidget["blank"].(map[string]interface{}); ok {
											rMosaicLayoutTiles.Widget.Blank = &dclService.DashboardMosaicLayoutTilesWidgetBlank{}
										} else {
											return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Blank: expected map[string]interface{}")
										}
									}
									if _, ok := rMosaicLayoutTilesWidget["logsPanel"]; ok {
										if rMosaicLayoutTilesWidgetLogsPanel, ok := rMosaicLayoutTilesWidget["logsPanel"].(map[string]interface{}); ok {
											rMosaicLayoutTiles.Widget.LogsPanel = &dclService.DashboardMosaicLayoutTilesWidgetLogsPanel{}
											if _, ok := rMosaicLayoutTilesWidgetLogsPanel["filter"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetLogsPanel["filter"].(string); ok {
													rMosaicLayoutTiles.Widget.LogsPanel.Filter = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.LogsPanel.Filter: expected string")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetLogsPanel["resourceNames"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetLogsPanel["resourceNames"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rMosaicLayoutTiles.Widget.LogsPanel.ResourceNames = append(rMosaicLayoutTiles.Widget.LogsPanel.ResourceNames, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.LogsPanel.ResourceNames: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.LogsPanel: expected map[string]interface{}")
										}
									}
									if _, ok := rMosaicLayoutTilesWidget["scorecard"]; ok {
										if rMosaicLayoutTilesWidgetScorecard, ok := rMosaicLayoutTilesWidget["scorecard"].(map[string]interface{}); ok {
											rMosaicLayoutTiles.Widget.Scorecard = &dclService.DashboardMosaicLayoutTilesWidgetScorecard{}
											if _, ok := rMosaicLayoutTilesWidgetScorecard["gaugeView"]; ok {
												if rMosaicLayoutTilesWidgetScorecardGaugeView, ok := rMosaicLayoutTilesWidgetScorecard["gaugeView"].(map[string]interface{}); ok {
													rMosaicLayoutTiles.Widget.Scorecard.GaugeView = &dclService.DashboardMosaicLayoutTilesWidgetScorecardGaugeView{}
													if _, ok := rMosaicLayoutTilesWidgetScorecardGaugeView["lowerBound"]; ok {
														if f, ok := rMosaicLayoutTilesWidgetScorecardGaugeView["lowerBound"].(float64); ok {
															rMosaicLayoutTiles.Widget.Scorecard.GaugeView.LowerBound = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.GaugeView.LowerBound: expected float64")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetScorecardGaugeView["upperBound"]; ok {
														if f, ok := rMosaicLayoutTilesWidgetScorecardGaugeView["upperBound"].(float64); ok {
															rMosaicLayoutTiles.Widget.Scorecard.GaugeView.UpperBound = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.GaugeView.UpperBound: expected float64")
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.GaugeView: expected map[string]interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetScorecard["sparkChartView"]; ok {
												if rMosaicLayoutTilesWidgetScorecardSparkChartView, ok := rMosaicLayoutTilesWidgetScorecard["sparkChartView"].(map[string]interface{}); ok {
													rMosaicLayoutTiles.Widget.Scorecard.SparkChartView = &dclService.DashboardMosaicLayoutTilesWidgetScorecardSparkChartView{}
													if _, ok := rMosaicLayoutTilesWidgetScorecardSparkChartView["minAlignmentPeriod"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetScorecardSparkChartView["minAlignmentPeriod"].(string); ok {
															rMosaicLayoutTiles.Widget.Scorecard.SparkChartView.MinAlignmentPeriod = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.SparkChartView.MinAlignmentPeriod: expected string")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetScorecardSparkChartView["sparkChartType"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetScorecardSparkChartView["sparkChartType"].(string); ok {
															rMosaicLayoutTiles.Widget.Scorecard.SparkChartView.SparkChartType = dclService.DashboardMosaicLayoutTilesWidgetScorecardSparkChartViewSparkChartTypeEnumRef(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.SparkChartView.SparkChartType: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.SparkChartView: expected map[string]interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetScorecard["thresholds"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetScorecard["thresholds"].([]interface{}); ok {
													for _, o := range s {
														if objval, ok := o.(map[string]interface{}); ok {
															var rMosaicLayoutTilesWidgetScorecardThresholds dclService.DashboardMosaicLayoutTilesWidgetScorecardThresholds
															if _, ok := objval["color"]; ok {
																if s, ok := objval["color"].(string); ok {
																	rMosaicLayoutTilesWidgetScorecardThresholds.Color = dclService.DashboardMosaicLayoutTilesWidgetScorecardThresholdsColorEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetScorecardThresholds.Color: expected string")
																}
															}
															if _, ok := objval["direction"]; ok {
																if s, ok := objval["direction"].(string); ok {
																	rMosaicLayoutTilesWidgetScorecardThresholds.Direction = dclService.DashboardMosaicLayoutTilesWidgetScorecardThresholdsDirectionEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetScorecardThresholds.Direction: expected string")
																}
															}
															if _, ok := objval["label"]; ok {
																if s, ok := objval["label"].(string); ok {
																	rMosaicLayoutTilesWidgetScorecardThresholds.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetScorecardThresholds.Label: expected string")
																}
															}
															if _, ok := objval["value"]; ok {
																if f, ok := objval["value"].(float64); ok {
																	rMosaicLayoutTilesWidgetScorecardThresholds.Value = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetScorecardThresholds.Value: expected float64")
																}
															}
															rMosaicLayoutTiles.Widget.Scorecard.Thresholds = append(rMosaicLayoutTiles.Widget.Scorecard.Thresholds, rMosaicLayoutTilesWidgetScorecardThresholds)
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.Thresholds: expected []interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetScorecard["timeSeriesQuery"]; ok {
												if rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery, ok := rMosaicLayoutTilesWidgetScorecard["timeSeriesQuery"].(map[string]interface{}); ok {
													rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQuery{}
													if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["timeSeriesFilter"]; ok {
														if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
															rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter{}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																}
															}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																}
															}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																		if i, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																}
															}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
														if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
															rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																		if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																}
															}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																		if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																}
															}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																		if i, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																}
															}
															if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																if rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																	rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																			rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
															rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["unitOverride"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetScorecardTimeSeriesQuery["unitOverride"].(string); ok {
															rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.UnitOverride = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery.UnitOverride: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard.TimeSeriesQuery: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Scorecard: expected map[string]interface{}")
										}
									}
									if _, ok := rMosaicLayoutTilesWidget["text"]; ok {
										if rMosaicLayoutTilesWidgetText, ok := rMosaicLayoutTilesWidget["text"].(map[string]interface{}); ok {
											rMosaicLayoutTiles.Widget.Text = &dclService.DashboardMosaicLayoutTilesWidgetText{}
											if _, ok := rMosaicLayoutTilesWidgetText["content"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetText["content"].(string); ok {
													rMosaicLayoutTiles.Widget.Text.Content = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Text.Content: expected string")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetText["format"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetText["format"].(string); ok {
													rMosaicLayoutTiles.Widget.Text.Format = dclService.DashboardMosaicLayoutTilesWidgetTextFormatEnumRef(s)
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Text.Format: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Text: expected map[string]interface{}")
										}
									}
									if _, ok := rMosaicLayoutTilesWidget["title"]; ok {
										if s, ok := rMosaicLayoutTilesWidget["title"].(string); ok {
											rMosaicLayoutTiles.Widget.Title = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.Title: expected string")
										}
									}
									if _, ok := rMosaicLayoutTilesWidget["xyChart"]; ok {
										if rMosaicLayoutTilesWidgetXyChart, ok := rMosaicLayoutTilesWidget["xyChart"].(map[string]interface{}); ok {
											rMosaicLayoutTiles.Widget.XyChart = &dclService.DashboardMosaicLayoutTilesWidgetXyChart{}
											if _, ok := rMosaicLayoutTilesWidgetXyChart["chartOptions"]; ok {
												if rMosaicLayoutTilesWidgetXyChartChartOptions, ok := rMosaicLayoutTilesWidgetXyChart["chartOptions"].(map[string]interface{}); ok {
													rMosaicLayoutTiles.Widget.XyChart.ChartOptions = &dclService.DashboardMosaicLayoutTilesWidgetXyChartChartOptions{}
													if _, ok := rMosaicLayoutTilesWidgetXyChartChartOptions["mode"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetXyChartChartOptions["mode"].(string); ok {
															rMosaicLayoutTiles.Widget.XyChart.ChartOptions.Mode = dclService.DashboardMosaicLayoutTilesWidgetXyChartChartOptionsModeEnumRef(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.ChartOptions.Mode: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.ChartOptions: expected map[string]interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetXyChart["dataSets"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetXyChart["dataSets"].([]interface{}); ok {
													for _, o := range s {
														if objval, ok := o.(map[string]interface{}); ok {
															var rMosaicLayoutTilesWidgetXyChartDataSets dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSets
															if _, ok := objval["legendTemplate"]; ok {
																if s, ok := objval["legendTemplate"].(string); ok {
																	rMosaicLayoutTilesWidgetXyChartDataSets.LegendTemplate = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.LegendTemplate: expected string")
																}
															}
															if _, ok := objval["minAlignmentPeriod"]; ok {
																if s, ok := objval["minAlignmentPeriod"].(string); ok {
																	rMosaicLayoutTilesWidgetXyChartDataSets.MinAlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.MinAlignmentPeriod: expected string")
																}
															}
															if _, ok := objval["plotType"]; ok {
																if s, ok := objval["plotType"].(string); ok {
																	rMosaicLayoutTilesWidgetXyChartDataSets.PlotType = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsPlotTypeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.PlotType: expected string")
																}
															}
															if _, ok := objval["timeSeriesQuery"]; ok {
																if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery, ok := objval["timeSeriesQuery"].(map[string]interface{}); ok {
																	rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery{}
																	if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"]; ok {
																		if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																			rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																				if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																						if i, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																		if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																			rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																						if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																						if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																								if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																						if i, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																				if rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																					rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																						if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																							rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																			rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																		}
																	}
																	if _, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["unitOverride"]; ok {
																		if s, ok := rMosaicLayoutTilesWidgetXyChartDataSetsTimeSeriesQuery["unitOverride"].(string); ok {
																			rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.UnitOverride = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery.UnitOverride: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartDataSets.TimeSeriesQuery: expected map[string]interface{}")
																}
															}
															rMosaicLayoutTiles.Widget.XyChart.DataSets = append(rMosaicLayoutTiles.Widget.XyChart.DataSets, rMosaicLayoutTilesWidgetXyChartDataSets)
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.DataSets: expected []interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetXyChart["thresholds"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetXyChart["thresholds"].([]interface{}); ok {
													for _, o := range s {
														if objval, ok := o.(map[string]interface{}); ok {
															var rMosaicLayoutTilesWidgetXyChartThresholds dclService.DashboardMosaicLayoutTilesWidgetXyChartThresholds
															if _, ok := objval["color"]; ok {
																if s, ok := objval["color"].(string); ok {
																	rMosaicLayoutTilesWidgetXyChartThresholds.Color = dclService.DashboardMosaicLayoutTilesWidgetXyChartThresholdsColorEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartThresholds.Color: expected string")
																}
															}
															if _, ok := objval["direction"]; ok {
																if s, ok := objval["direction"].(string); ok {
																	rMosaicLayoutTilesWidgetXyChartThresholds.Direction = dclService.DashboardMosaicLayoutTilesWidgetXyChartThresholdsDirectionEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartThresholds.Direction: expected string")
																}
															}
															if _, ok := objval["label"]; ok {
																if s, ok := objval["label"].(string); ok {
																	rMosaicLayoutTilesWidgetXyChartThresholds.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartThresholds.Label: expected string")
																}
															}
															if _, ok := objval["value"]; ok {
																if f, ok := objval["value"].(float64); ok {
																	rMosaicLayoutTilesWidgetXyChartThresholds.Value = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rMosaicLayoutTilesWidgetXyChartThresholds.Value: expected float64")
																}
															}
															rMosaicLayoutTiles.Widget.XyChart.Thresholds = append(rMosaicLayoutTiles.Widget.XyChart.Thresholds, rMosaicLayoutTilesWidgetXyChartThresholds)
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.Thresholds: expected []interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetXyChart["timeshiftDuration"]; ok {
												if s, ok := rMosaicLayoutTilesWidgetXyChart["timeshiftDuration"].(string); ok {
													rMosaicLayoutTiles.Widget.XyChart.TimeshiftDuration = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.TimeshiftDuration: expected string")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetXyChart["xAxis"]; ok {
												if rMosaicLayoutTilesWidgetXyChartXAxis, ok := rMosaicLayoutTilesWidgetXyChart["xAxis"].(map[string]interface{}); ok {
													rMosaicLayoutTiles.Widget.XyChart.XAxis = &dclService.DashboardMosaicLayoutTilesWidgetXyChartXAxis{}
													if _, ok := rMosaicLayoutTilesWidgetXyChartXAxis["label"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetXyChartXAxis["label"].(string); ok {
															rMosaicLayoutTiles.Widget.XyChart.XAxis.Label = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.XAxis.Label: expected string")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetXyChartXAxis["scale"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetXyChartXAxis["scale"].(string); ok {
															rMosaicLayoutTiles.Widget.XyChart.XAxis.Scale = dclService.DashboardMosaicLayoutTilesWidgetXyChartXAxisScaleEnumRef(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.XAxis.Scale: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.XAxis: expected map[string]interface{}")
												}
											}
											if _, ok := rMosaicLayoutTilesWidgetXyChart["yAxis"]; ok {
												if rMosaicLayoutTilesWidgetXyChartYAxis, ok := rMosaicLayoutTilesWidgetXyChart["yAxis"].(map[string]interface{}); ok {
													rMosaicLayoutTiles.Widget.XyChart.YAxis = &dclService.DashboardMosaicLayoutTilesWidgetXyChartYAxis{}
													if _, ok := rMosaicLayoutTilesWidgetXyChartYAxis["label"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetXyChartYAxis["label"].(string); ok {
															rMosaicLayoutTiles.Widget.XyChart.YAxis.Label = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.YAxis.Label: expected string")
														}
													}
													if _, ok := rMosaicLayoutTilesWidgetXyChartYAxis["scale"]; ok {
														if s, ok := rMosaicLayoutTilesWidgetXyChartYAxis["scale"].(string); ok {
															rMosaicLayoutTiles.Widget.XyChart.YAxis.Scale = dclService.DashboardMosaicLayoutTilesWidgetXyChartYAxisScaleEnumRef(s)
														} else {
															return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.YAxis.Scale: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart.YAxis: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rMosaicLayoutTiles.Widget.XyChart: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.Widget: expected map[string]interface{}")
								}
							}
							if _, ok := objval["width"]; ok {
								if i, ok := objval["width"].(int64); ok {
									rMosaicLayoutTiles.Width = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.Width: expected int64")
								}
							}
							if _, ok := objval["xPos"]; ok {
								if i, ok := objval["xPos"].(int64); ok {
									rMosaicLayoutTiles.XPos = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.XPos: expected int64")
								}
							}
							if _, ok := objval["yPos"]; ok {
								if i, ok := objval["yPos"].(int64); ok {
									rMosaicLayoutTiles.YPos = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rMosaicLayoutTiles.YPos: expected int64")
								}
							}
							r.MosaicLayout.Tiles = append(r.MosaicLayout.Tiles, rMosaicLayoutTiles)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MosaicLayout.Tiles: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MosaicLayout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["rowLayout"]; ok {
		if rRowLayout, ok := u.Object["rowLayout"].(map[string]interface{}); ok {
			r.RowLayout = &dclService.DashboardRowLayout{}
			if _, ok := rRowLayout["rows"]; ok {
				if s, ok := rRowLayout["rows"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rRowLayoutRows dclService.DashboardRowLayoutRows
							if _, ok := objval["weight"]; ok {
								if i, ok := objval["weight"].(int64); ok {
									rRowLayoutRows.Weight = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rRowLayoutRows.Weight: expected int64")
								}
							}
							if _, ok := objval["widgets"]; ok {
								if s, ok := objval["widgets"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rRowLayoutRowsWidgets dclService.DashboardRowLayoutRowsWidgets
											if _, ok := objval["blank"]; ok {
												if _, ok := objval["blank"].(map[string]interface{}); ok {
													rRowLayoutRowsWidgets.Blank = &dclService.DashboardRowLayoutRowsWidgetsBlank{}
												} else {
													return nil, fmt.Errorf("rRowLayoutRowsWidgets.Blank: expected map[string]interface{}")
												}
											}
											if _, ok := objval["logsPanel"]; ok {
												if rRowLayoutRowsWidgetsLogsPanel, ok := objval["logsPanel"].(map[string]interface{}); ok {
													rRowLayoutRowsWidgets.LogsPanel = &dclService.DashboardRowLayoutRowsWidgetsLogsPanel{}
													if _, ok := rRowLayoutRowsWidgetsLogsPanel["filter"]; ok {
														if s, ok := rRowLayoutRowsWidgetsLogsPanel["filter"].(string); ok {
															rRowLayoutRowsWidgets.LogsPanel.Filter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.LogsPanel.Filter: expected string")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsLogsPanel["resourceNames"]; ok {
														if s, ok := rRowLayoutRowsWidgetsLogsPanel["resourceNames"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rRowLayoutRowsWidgets.LogsPanel.ResourceNames = append(rRowLayoutRowsWidgets.LogsPanel.ResourceNames, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.LogsPanel.ResourceNames: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rRowLayoutRowsWidgets.LogsPanel: expected map[string]interface{}")
												}
											}
											if _, ok := objval["scorecard"]; ok {
												if rRowLayoutRowsWidgetsScorecard, ok := objval["scorecard"].(map[string]interface{}); ok {
													rRowLayoutRowsWidgets.Scorecard = &dclService.DashboardRowLayoutRowsWidgetsScorecard{}
													if _, ok := rRowLayoutRowsWidgetsScorecard["gaugeView"]; ok {
														if rRowLayoutRowsWidgetsScorecardGaugeView, ok := rRowLayoutRowsWidgetsScorecard["gaugeView"].(map[string]interface{}); ok {
															rRowLayoutRowsWidgets.Scorecard.GaugeView = &dclService.DashboardRowLayoutRowsWidgetsScorecardGaugeView{}
															if _, ok := rRowLayoutRowsWidgetsScorecardGaugeView["lowerBound"]; ok {
																if f, ok := rRowLayoutRowsWidgetsScorecardGaugeView["lowerBound"].(float64); ok {
																	rRowLayoutRowsWidgets.Scorecard.GaugeView.LowerBound = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.GaugeView.LowerBound: expected float64")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsScorecardGaugeView["upperBound"]; ok {
																if f, ok := rRowLayoutRowsWidgetsScorecardGaugeView["upperBound"].(float64); ok {
																	rRowLayoutRowsWidgets.Scorecard.GaugeView.UpperBound = dcl.Float64(f)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.GaugeView.UpperBound: expected float64")
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.GaugeView: expected map[string]interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsScorecard["sparkChartView"]; ok {
														if rRowLayoutRowsWidgetsScorecardSparkChartView, ok := rRowLayoutRowsWidgetsScorecard["sparkChartView"].(map[string]interface{}); ok {
															rRowLayoutRowsWidgets.Scorecard.SparkChartView = &dclService.DashboardRowLayoutRowsWidgetsScorecardSparkChartView{}
															if _, ok := rRowLayoutRowsWidgetsScorecardSparkChartView["minAlignmentPeriod"]; ok {
																if s, ok := rRowLayoutRowsWidgetsScorecardSparkChartView["minAlignmentPeriod"].(string); ok {
																	rRowLayoutRowsWidgets.Scorecard.SparkChartView.MinAlignmentPeriod = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.SparkChartView.MinAlignmentPeriod: expected string")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsScorecardSparkChartView["sparkChartType"]; ok {
																if s, ok := rRowLayoutRowsWidgetsScorecardSparkChartView["sparkChartType"].(string); ok {
																	rRowLayoutRowsWidgets.Scorecard.SparkChartView.SparkChartType = dclService.DashboardRowLayoutRowsWidgetsScorecardSparkChartViewSparkChartTypeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.SparkChartView.SparkChartType: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.SparkChartView: expected map[string]interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsScorecard["thresholds"]; ok {
														if s, ok := rRowLayoutRowsWidgetsScorecard["thresholds"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rRowLayoutRowsWidgetsScorecardThresholds dclService.DashboardRowLayoutRowsWidgetsScorecardThresholds
																	if _, ok := objval["color"]; ok {
																		if s, ok := objval["color"].(string); ok {
																			rRowLayoutRowsWidgetsScorecardThresholds.Color = dclService.DashboardRowLayoutRowsWidgetsScorecardThresholdsColorEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsScorecardThresholds.Color: expected string")
																		}
																	}
																	if _, ok := objval["direction"]; ok {
																		if s, ok := objval["direction"].(string); ok {
																			rRowLayoutRowsWidgetsScorecardThresholds.Direction = dclService.DashboardRowLayoutRowsWidgetsScorecardThresholdsDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsScorecardThresholds.Direction: expected string")
																		}
																	}
																	if _, ok := objval["label"]; ok {
																		if s, ok := objval["label"].(string); ok {
																			rRowLayoutRowsWidgetsScorecardThresholds.Label = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsScorecardThresholds.Label: expected string")
																		}
																	}
																	if _, ok := objval["value"]; ok {
																		if f, ok := objval["value"].(float64); ok {
																			rRowLayoutRowsWidgetsScorecardThresholds.Value = dcl.Float64(f)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsScorecardThresholds.Value: expected float64")
																		}
																	}
																	rRowLayoutRowsWidgets.Scorecard.Thresholds = append(rRowLayoutRowsWidgets.Scorecard.Thresholds, rRowLayoutRowsWidgetsScorecardThresholds)
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.Thresholds: expected []interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsScorecard["timeSeriesQuery"]; ok {
														if rRowLayoutRowsWidgetsScorecardTimeSeriesQuery, ok := rRowLayoutRowsWidgetsScorecard["timeSeriesQuery"].(map[string]interface{}); ok {
															rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQuery{}
															if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"]; ok {
																if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																	rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter{}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																		if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																		}
																	}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																	rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio{}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																				if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																				if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																							for _, ss := range s {
																								if strval, ok := ss.(string); ok {
																									rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																				if i, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																		if rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																					rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																	rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["unitOverride"]; ok {
																if s, ok := rRowLayoutRowsWidgetsScorecardTimeSeriesQuery["unitOverride"].(string); ok {
																	rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.UnitOverride = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery.UnitOverride: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard.TimeSeriesQuery: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rRowLayoutRowsWidgets.Scorecard: expected map[string]interface{}")
												}
											}
											if _, ok := objval["text"]; ok {
												if rRowLayoutRowsWidgetsText, ok := objval["text"].(map[string]interface{}); ok {
													rRowLayoutRowsWidgets.Text = &dclService.DashboardRowLayoutRowsWidgetsText{}
													if _, ok := rRowLayoutRowsWidgetsText["content"]; ok {
														if s, ok := rRowLayoutRowsWidgetsText["content"].(string); ok {
															rRowLayoutRowsWidgets.Text.Content = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.Text.Content: expected string")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsText["format"]; ok {
														if s, ok := rRowLayoutRowsWidgetsText["format"].(string); ok {
															rRowLayoutRowsWidgets.Text.Format = dclService.DashboardRowLayoutRowsWidgetsTextFormatEnumRef(s)
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.Text.Format: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rRowLayoutRowsWidgets.Text: expected map[string]interface{}")
												}
											}
											if _, ok := objval["title"]; ok {
												if s, ok := objval["title"].(string); ok {
													rRowLayoutRowsWidgets.Title = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRowLayoutRowsWidgets.Title: expected string")
												}
											}
											if _, ok := objval["xyChart"]; ok {
												if rRowLayoutRowsWidgetsXyChart, ok := objval["xyChart"].(map[string]interface{}); ok {
													rRowLayoutRowsWidgets.XyChart = &dclService.DashboardRowLayoutRowsWidgetsXyChart{}
													if _, ok := rRowLayoutRowsWidgetsXyChart["chartOptions"]; ok {
														if rRowLayoutRowsWidgetsXyChartChartOptions, ok := rRowLayoutRowsWidgetsXyChart["chartOptions"].(map[string]interface{}); ok {
															rRowLayoutRowsWidgets.XyChart.ChartOptions = &dclService.DashboardRowLayoutRowsWidgetsXyChartChartOptions{}
															if _, ok := rRowLayoutRowsWidgetsXyChartChartOptions["mode"]; ok {
																if s, ok := rRowLayoutRowsWidgetsXyChartChartOptions["mode"].(string); ok {
																	rRowLayoutRowsWidgets.XyChart.ChartOptions.Mode = dclService.DashboardRowLayoutRowsWidgetsXyChartChartOptionsModeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.ChartOptions.Mode: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.ChartOptions: expected map[string]interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsXyChart["dataSets"]; ok {
														if s, ok := rRowLayoutRowsWidgetsXyChart["dataSets"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rRowLayoutRowsWidgetsXyChartDataSets dclService.DashboardRowLayoutRowsWidgetsXyChartDataSets
																	if _, ok := objval["legendTemplate"]; ok {
																		if s, ok := objval["legendTemplate"].(string); ok {
																			rRowLayoutRowsWidgetsXyChartDataSets.LegendTemplate = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.LegendTemplate: expected string")
																		}
																	}
																	if _, ok := objval["minAlignmentPeriod"]; ok {
																		if s, ok := objval["minAlignmentPeriod"].(string); ok {
																			rRowLayoutRowsWidgetsXyChartDataSets.MinAlignmentPeriod = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.MinAlignmentPeriod: expected string")
																		}
																	}
																	if _, ok := objval["plotType"]; ok {
																		if s, ok := objval["plotType"].(string); ok {
																			rRowLayoutRowsWidgetsXyChartDataSets.PlotType = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsPlotTypeEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.PlotType: expected string")
																		}
																	}
																	if _, ok := objval["timeSeriesQuery"]; ok {
																		if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery, ok := objval["timeSeriesQuery"].(map[string]interface{}); ok {
																			rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery{}
																			if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"]; ok {
																				if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilter"].(map[string]interface{}); ok {
																					rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["aggregation"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["alignmentPeriod"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["crossSeriesReducer"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields = append(rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation["perSeriesAligner"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Aggregation: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"]; ok {
																						if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["filter"].(string); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.Filter: expected string")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["direction"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.Direction: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"]; ok {
																								if i, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter["rankingMethod"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter.RankingMethod: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.PickTimeSeriesFilter: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilter["secondaryAggregation"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["alignmentPeriod"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["crossSeriesReducer"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields = append(rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation["perSeriesAligner"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter.SecondaryAggregation: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilter: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"]; ok {
																				if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesFilterRatio"].(map[string]interface{}); ok {
																					rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["denominator"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"]; ok {
																								if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["aggregation"].(map[string]interface{}); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["alignmentPeriod"].(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.AlignmentPeriod: expected string")
																										}
																									}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["crossSeriesReducer"].(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.CrossSeriesReducer: expected string")
																										}
																									}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["groupByFields"].([]interface{}); ok {
																											for _, ss := range s {
																												if strval, ok := ss.(string); ok {
																													rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields = append(rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields, strval)
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.GroupByFields: expected []interface{}")
																										}
																									}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation["perSeriesAligner"].(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation.PerSeriesAligner: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Aggregation: expected map[string]interface{}")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator["filter"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator.Filter: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Denominator: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["numerator"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"]; ok {
																								if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["aggregation"].(map[string]interface{}); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["alignmentPeriod"].(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod = dcl.String(s)
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.AlignmentPeriod: expected string")
																										}
																									}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["crossSeriesReducer"].(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.CrossSeriesReducer: expected string")
																										}
																									}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["groupByFields"].([]interface{}); ok {
																											for _, ss := range s {
																												if strval, ok := ss.(string); ok {
																													rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields = append(rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields, strval)
																												}
																											}
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.GroupByFields: expected []interface{}")
																										}
																									}
																									if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"]; ok {
																										if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation["perSeriesAligner"].(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s)
																										} else {
																											return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation.PerSeriesAligner: expected string")
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Aggregation: expected map[string]interface{}")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator["filter"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator.Filter: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.Numerator: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["pickTimeSeriesFilter"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["direction"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.Direction: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"]; ok {
																								if i, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["numTimeSeries"].(int64); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries = dcl.Int64(i)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.NumTimeSeries: expected int64")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter["rankingMethod"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter.RankingMethod: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.PickTimeSeriesFilter: expected map[string]interface{}")
																						}
																					}
																					if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"]; ok {
																						if rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio["secondaryAggregation"].(map[string]interface{}); ok {
																							rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation = &dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["alignmentPeriod"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod = dcl.String(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.AlignmentPeriod: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["crossSeriesReducer"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.CrossSeriesReducer: expected string")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["groupByFields"].([]interface{}); ok {
																									for _, ss := range s {
																										if strval, ok := ss.(string); ok {
																											rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields = append(rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields, strval)
																										}
																									}
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.GroupByFields: expected []interface{}")
																								}
																							}
																							if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"]; ok {
																								if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation["perSeriesAligner"].(string); ok {
																									rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner = dclService.DashboardRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s)
																								} else {
																									return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation.PerSeriesAligner: expected string")
																								}
																							}
																						} else {
																							return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio.SecondaryAggregation: expected map[string]interface{}")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesFilterRatio: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["timeSeriesQueryLanguage"].(string); ok {
																					rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.TimeSeriesQueryLanguage: expected string")
																				}
																			}
																			if _, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"]; ok {
																				if s, ok := rRowLayoutRowsWidgetsXyChartDataSetsTimeSeriesQuery["unitOverride"].(string); ok {
																					rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery.UnitOverride: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartDataSets.TimeSeriesQuery: expected map[string]interface{}")
																		}
																	}
																	rRowLayoutRowsWidgets.XyChart.DataSets = append(rRowLayoutRowsWidgets.XyChart.DataSets, rRowLayoutRowsWidgetsXyChartDataSets)
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.DataSets: expected []interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsXyChart["thresholds"]; ok {
														if s, ok := rRowLayoutRowsWidgetsXyChart["thresholds"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rRowLayoutRowsWidgetsXyChartThresholds dclService.DashboardRowLayoutRowsWidgetsXyChartThresholds
																	if _, ok := objval["color"]; ok {
																		if s, ok := objval["color"].(string); ok {
																			rRowLayoutRowsWidgetsXyChartThresholds.Color = dclService.DashboardRowLayoutRowsWidgetsXyChartThresholdsColorEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartThresholds.Color: expected string")
																		}
																	}
																	if _, ok := objval["direction"]; ok {
																		if s, ok := objval["direction"].(string); ok {
																			rRowLayoutRowsWidgetsXyChartThresholds.Direction = dclService.DashboardRowLayoutRowsWidgetsXyChartThresholdsDirectionEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartThresholds.Direction: expected string")
																		}
																	}
																	if _, ok := objval["label"]; ok {
																		if s, ok := objval["label"].(string); ok {
																			rRowLayoutRowsWidgetsXyChartThresholds.Label = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartThresholds.Label: expected string")
																		}
																	}
																	if _, ok := objval["value"]; ok {
																		if f, ok := objval["value"].(float64); ok {
																			rRowLayoutRowsWidgetsXyChartThresholds.Value = dcl.Float64(f)
																		} else {
																			return nil, fmt.Errorf("rRowLayoutRowsWidgetsXyChartThresholds.Value: expected float64")
																		}
																	}
																	rRowLayoutRowsWidgets.XyChart.Thresholds = append(rRowLayoutRowsWidgets.XyChart.Thresholds, rRowLayoutRowsWidgetsXyChartThresholds)
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.Thresholds: expected []interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsXyChart["timeshiftDuration"]; ok {
														if s, ok := rRowLayoutRowsWidgetsXyChart["timeshiftDuration"].(string); ok {
															rRowLayoutRowsWidgets.XyChart.TimeshiftDuration = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.TimeshiftDuration: expected string")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsXyChart["xAxis"]; ok {
														if rRowLayoutRowsWidgetsXyChartXAxis, ok := rRowLayoutRowsWidgetsXyChart["xAxis"].(map[string]interface{}); ok {
															rRowLayoutRowsWidgets.XyChart.XAxis = &dclService.DashboardRowLayoutRowsWidgetsXyChartXAxis{}
															if _, ok := rRowLayoutRowsWidgetsXyChartXAxis["label"]; ok {
																if s, ok := rRowLayoutRowsWidgetsXyChartXAxis["label"].(string); ok {
																	rRowLayoutRowsWidgets.XyChart.XAxis.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.XAxis.Label: expected string")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsXyChartXAxis["scale"]; ok {
																if s, ok := rRowLayoutRowsWidgetsXyChartXAxis["scale"].(string); ok {
																	rRowLayoutRowsWidgets.XyChart.XAxis.Scale = dclService.DashboardRowLayoutRowsWidgetsXyChartXAxisScaleEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.XAxis.Scale: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.XAxis: expected map[string]interface{}")
														}
													}
													if _, ok := rRowLayoutRowsWidgetsXyChart["yAxis"]; ok {
														if rRowLayoutRowsWidgetsXyChartYAxis, ok := rRowLayoutRowsWidgetsXyChart["yAxis"].(map[string]interface{}); ok {
															rRowLayoutRowsWidgets.XyChart.YAxis = &dclService.DashboardRowLayoutRowsWidgetsXyChartYAxis{}
															if _, ok := rRowLayoutRowsWidgetsXyChartYAxis["label"]; ok {
																if s, ok := rRowLayoutRowsWidgetsXyChartYAxis["label"].(string); ok {
																	rRowLayoutRowsWidgets.XyChart.YAxis.Label = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.YAxis.Label: expected string")
																}
															}
															if _, ok := rRowLayoutRowsWidgetsXyChartYAxis["scale"]; ok {
																if s, ok := rRowLayoutRowsWidgetsXyChartYAxis["scale"].(string); ok {
																	rRowLayoutRowsWidgets.XyChart.YAxis.Scale = dclService.DashboardRowLayoutRowsWidgetsXyChartYAxisScaleEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.YAxis.Scale: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart.YAxis: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rRowLayoutRowsWidgets.XyChart: expected map[string]interface{}")
												}
											}
											rRowLayoutRows.Widgets = append(rRowLayoutRows.Widgets, rRowLayoutRowsWidgets)
										}
									}
								} else {
									return nil, fmt.Errorf("rRowLayoutRows.Widgets: expected []interface{}")
								}
							}
							r.RowLayout.Rows = append(r.RowLayout.Rows, rRowLayoutRows)
						}
					}
				} else {
					return nil, fmt.Errorf("r.RowLayout.Rows: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RowLayout: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetDashboard(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDashboard(ctx, r)
	if err != nil {
		return nil, err
	}
	return DashboardToUnstructured(r), nil
}

func ListDashboard(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDashboard(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DashboardToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyDashboard(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDashboard(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDashboard(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DashboardToUnstructured(r), nil
}

func DashboardHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDashboard(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDashboard(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDashboard(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return err
	}
	return c.DeleteDashboard(ctx, r)
}

func DashboardID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDashboard(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Dashboard) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"Dashboard",
		"alpha",
	}
}

func (r *Dashboard) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Dashboard) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Dashboard) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDashboard(ctx, config, resource)
}

func (r *Dashboard) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDashboard(ctx, config, resource, opts...)
}

func (r *Dashboard) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DashboardHasDiff(ctx, config, resource, opts...)
}

func (r *Dashboard) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDashboard(ctx, config, resource)
}

func (r *Dashboard) ID(resource *unstructured.Resource) (string, error) {
	return DashboardID(resource)
}

func init() {
	unstructured.Register(&Dashboard{})
}
