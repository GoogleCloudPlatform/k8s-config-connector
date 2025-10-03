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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type ServiceLevelObjective struct{}

func ServiceLevelObjectiveToUnstructured(r *dclService.ServiceLevelObjective) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "monitoring",
			Version: "beta",
			Type:    "ServiceLevelObjective",
		},
		Object: make(map[string]interface{}),
	}
	if r.CalendarPeriod != nil {
		u.Object["calendarPeriod"] = string(*r.CalendarPeriod)
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Goal != nil {
		u.Object["goal"] = *r.Goal
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RollingPeriod != nil {
		u.Object["rollingPeriod"] = *r.RollingPeriod
	}
	if r.Service != nil {
		u.Object["service"] = *r.Service
	}
	if r.ServiceLevelIndicator != nil && r.ServiceLevelIndicator != dclService.EmptyServiceLevelObjectiveServiceLevelIndicator {
		rServiceLevelIndicator := make(map[string]interface{})
		if r.ServiceLevelIndicator.BasicSli != nil && r.ServiceLevelIndicator.BasicSli != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorBasicSli {
			rServiceLevelIndicatorBasicSli := make(map[string]interface{})
			if r.ServiceLevelIndicator.BasicSli.Availability != nil && r.ServiceLevelIndicator.BasicSli.Availability != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
				rServiceLevelIndicatorBasicSliAvailability := make(map[string]interface{})
				rServiceLevelIndicatorBasicSli["availability"] = rServiceLevelIndicatorBasicSliAvailability
			}
			if r.ServiceLevelIndicator.BasicSli.Latency != nil && r.ServiceLevelIndicator.BasicSli.Latency != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
				rServiceLevelIndicatorBasicSliLatency := make(map[string]interface{})
				if r.ServiceLevelIndicator.BasicSli.Latency.Experience != nil {
					rServiceLevelIndicatorBasicSliLatency["experience"] = string(*r.ServiceLevelIndicator.BasicSli.Latency.Experience)
				}
				if r.ServiceLevelIndicator.BasicSli.Latency.Threshold != nil {
					rServiceLevelIndicatorBasicSliLatency["threshold"] = *r.ServiceLevelIndicator.BasicSli.Latency.Threshold
				}
				rServiceLevelIndicatorBasicSli["latency"] = rServiceLevelIndicatorBasicSliLatency
			}
			var rServiceLevelIndicatorBasicSliLocation []interface{}
			for _, rServiceLevelIndicatorBasicSliLocationVal := range r.ServiceLevelIndicator.BasicSli.Location {
				rServiceLevelIndicatorBasicSliLocation = append(rServiceLevelIndicatorBasicSliLocation, rServiceLevelIndicatorBasicSliLocationVal)
			}
			rServiceLevelIndicatorBasicSli["location"] = rServiceLevelIndicatorBasicSliLocation
			var rServiceLevelIndicatorBasicSliMethod []interface{}
			for _, rServiceLevelIndicatorBasicSliMethodVal := range r.ServiceLevelIndicator.BasicSli.Method {
				rServiceLevelIndicatorBasicSliMethod = append(rServiceLevelIndicatorBasicSliMethod, rServiceLevelIndicatorBasicSliMethodVal)
			}
			rServiceLevelIndicatorBasicSli["method"] = rServiceLevelIndicatorBasicSliMethod
			if r.ServiceLevelIndicator.BasicSli.OperationAvailability != nil && r.ServiceLevelIndicator.BasicSli.OperationAvailability != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
				rServiceLevelIndicatorBasicSliOperationAvailability := make(map[string]interface{})
				rServiceLevelIndicatorBasicSli["operationAvailability"] = rServiceLevelIndicatorBasicSliOperationAvailability
			}
			if r.ServiceLevelIndicator.BasicSli.OperationLatency != nil && r.ServiceLevelIndicator.BasicSli.OperationLatency != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
				rServiceLevelIndicatorBasicSliOperationLatency := make(map[string]interface{})
				if r.ServiceLevelIndicator.BasicSli.OperationLatency.Experience != nil {
					rServiceLevelIndicatorBasicSliOperationLatency["experience"] = string(*r.ServiceLevelIndicator.BasicSli.OperationLatency.Experience)
				}
				if r.ServiceLevelIndicator.BasicSli.OperationLatency.Threshold != nil {
					rServiceLevelIndicatorBasicSliOperationLatency["threshold"] = *r.ServiceLevelIndicator.BasicSli.OperationLatency.Threshold
				}
				rServiceLevelIndicatorBasicSli["operationLatency"] = rServiceLevelIndicatorBasicSliOperationLatency
			}
			var rServiceLevelIndicatorBasicSliVersion []interface{}
			for _, rServiceLevelIndicatorBasicSliVersionVal := range r.ServiceLevelIndicator.BasicSli.Version {
				rServiceLevelIndicatorBasicSliVersion = append(rServiceLevelIndicatorBasicSliVersion, rServiceLevelIndicatorBasicSliVersionVal)
			}
			rServiceLevelIndicatorBasicSli["version"] = rServiceLevelIndicatorBasicSliVersion
			rServiceLevelIndicator["basicSli"] = rServiceLevelIndicatorBasicSli
		}
		if r.ServiceLevelIndicator.RequestBased != nil && r.ServiceLevelIndicator.RequestBased != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorRequestBased {
			rServiceLevelIndicatorRequestBased := make(map[string]interface{})
			if r.ServiceLevelIndicator.RequestBased.DistributionCut != nil && r.ServiceLevelIndicator.RequestBased.DistributionCut != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
				rServiceLevelIndicatorRequestBasedDistributionCut := make(map[string]interface{})
				if r.ServiceLevelIndicator.RequestBased.DistributionCut.DistributionFilter != nil {
					rServiceLevelIndicatorRequestBasedDistributionCut["distributionFilter"] = *r.ServiceLevelIndicator.RequestBased.DistributionCut.DistributionFilter
				}
				if r.ServiceLevelIndicator.RequestBased.DistributionCut.Range != nil && r.ServiceLevelIndicator.RequestBased.DistributionCut.Range != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
					rServiceLevelIndicatorRequestBasedDistributionCutRange := make(map[string]interface{})
					if r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Max != nil {
						rServiceLevelIndicatorRequestBasedDistributionCutRange["max"] = *r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Max
					}
					if r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Min != nil {
						rServiceLevelIndicatorRequestBasedDistributionCutRange["min"] = *r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Min
					}
					rServiceLevelIndicatorRequestBasedDistributionCut["range"] = rServiceLevelIndicatorRequestBasedDistributionCutRange
				}
				rServiceLevelIndicatorRequestBased["distributionCut"] = rServiceLevelIndicatorRequestBasedDistributionCut
			}
			if r.ServiceLevelIndicator.RequestBased.GoodTotalRatio != nil && r.ServiceLevelIndicator.RequestBased.GoodTotalRatio != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
				rServiceLevelIndicatorRequestBasedGoodTotalRatio := make(map[string]interface{})
				if r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.BadServiceFilter != nil {
					rServiceLevelIndicatorRequestBasedGoodTotalRatio["badServiceFilter"] = *r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.BadServiceFilter
				}
				if r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.GoodServiceFilter != nil {
					rServiceLevelIndicatorRequestBasedGoodTotalRatio["goodServiceFilter"] = *r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.GoodServiceFilter
				}
				if r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.TotalServiceFilter != nil {
					rServiceLevelIndicatorRequestBasedGoodTotalRatio["totalServiceFilter"] = *r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.TotalServiceFilter
				}
				rServiceLevelIndicatorRequestBased["goodTotalRatio"] = rServiceLevelIndicatorRequestBasedGoodTotalRatio
			}
			rServiceLevelIndicator["requestBased"] = rServiceLevelIndicatorRequestBased
		}
		if r.ServiceLevelIndicator.WindowsBased != nil && r.ServiceLevelIndicator.WindowsBased != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
			rServiceLevelIndicatorWindowsBased := make(map[string]interface{})
			if r.ServiceLevelIndicator.WindowsBased.GoodBadMetricFilter != nil {
				rServiceLevelIndicatorWindowsBased["goodBadMetricFilter"] = *r.ServiceLevelIndicator.WindowsBased.GoodBadMetricFilter
			}
			if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
				rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold := make(map[string]interface{})
				if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance := make(map[string]interface{})
					if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Availability != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Availability != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability := make(map[string]interface{})
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["availability"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability
					}
					if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency := make(map[string]interface{})
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Experience != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency["experience"] = string(*r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Experience)
						}
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Threshold != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency["threshold"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Threshold
						}
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["latency"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency
					}
					var rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLocation []interface{}
					for _, rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLocationVal := range r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Location {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLocation = append(rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLocation, rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLocationVal)
					}
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["location"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLocation
					var rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceMethod []interface{}
					for _, rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceMethodVal := range r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Method {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceMethod = append(rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceMethod, rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceMethodVal)
					}
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["method"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceMethod
					if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationAvailability != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationAvailability != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability := make(map[string]interface{})
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["operationAvailability"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability
					}
					if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency := make(map[string]interface{})
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Experience != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency["experience"] = string(*r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Experience)
						}
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Threshold != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency["threshold"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Threshold
						}
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["operationLatency"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency
					}
					var rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceVersion []interface{}
					for _, rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceVersionVal := range r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Version {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceVersion = append(rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceVersion, rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceVersionVal)
					}
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["version"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceVersion
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["basicSliPerformance"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance
				}
				if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance := make(map[string]interface{})
					if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut := make(map[string]interface{})
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.DistributionFilter != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut["distributionFilter"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.DistributionFilter
						}
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange := make(map[string]interface{})
							if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Max != nil {
								rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange["max"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Max
							}
							if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Min != nil {
								rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange["min"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Min
							}
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut["range"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange
						}
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance["distributionCut"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut
					}
					if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio != nil && r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio := make(map[string]interface{})
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.BadServiceFilter != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["badServiceFilter"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.BadServiceFilter
						}
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.GoodServiceFilter != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["goodServiceFilter"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.GoodServiceFilter
						}
						if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.TotalServiceFilter != nil {
							rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["totalServiceFilter"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.TotalServiceFilter
						}
						rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance["goodTotalRatio"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio
					}
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["performance"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance
				}
				if r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Threshold != nil {
					rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["threshold"] = *r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Threshold
				}
				rServiceLevelIndicatorWindowsBased["goodTotalRatioThreshold"] = rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold
			}
			if r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange != nil && r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
				rServiceLevelIndicatorWindowsBasedMetricMeanInRange := make(map[string]interface{})
				if r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range != nil && r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
					rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange := make(map[string]interface{})
					if r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Max != nil {
						rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange["max"] = *r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Max
					}
					if r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Min != nil {
						rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange["min"] = *r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Min
					}
					rServiceLevelIndicatorWindowsBasedMetricMeanInRange["range"] = rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange
				}
				if r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.TimeSeries != nil {
					rServiceLevelIndicatorWindowsBasedMetricMeanInRange["timeSeries"] = *r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.TimeSeries
				}
				rServiceLevelIndicatorWindowsBased["metricMeanInRange"] = rServiceLevelIndicatorWindowsBasedMetricMeanInRange
			}
			if r.ServiceLevelIndicator.WindowsBased.MetricSumInRange != nil && r.ServiceLevelIndicator.WindowsBased.MetricSumInRange != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
				rServiceLevelIndicatorWindowsBasedMetricSumInRange := make(map[string]interface{})
				if r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range != nil && r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range != dclService.EmptyServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
					rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange := make(map[string]interface{})
					if r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Max != nil {
						rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange["max"] = *r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Max
					}
					if r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Min != nil {
						rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange["min"] = *r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Min
					}
					rServiceLevelIndicatorWindowsBasedMetricSumInRange["range"] = rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange
				}
				if r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.TimeSeries != nil {
					rServiceLevelIndicatorWindowsBasedMetricSumInRange["timeSeries"] = *r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.TimeSeries
				}
				rServiceLevelIndicatorWindowsBased["metricSumInRange"] = rServiceLevelIndicatorWindowsBasedMetricSumInRange
			}
			if r.ServiceLevelIndicator.WindowsBased.WindowPeriod != nil {
				rServiceLevelIndicatorWindowsBased["windowPeriod"] = *r.ServiceLevelIndicator.WindowsBased.WindowPeriod
			}
			rServiceLevelIndicator["windowsBased"] = rServiceLevelIndicatorWindowsBased
		}
		u.Object["serviceLevelIndicator"] = rServiceLevelIndicator
	}
	if r.ServiceManagementOwned != nil {
		u.Object["serviceManagementOwned"] = *r.ServiceManagementOwned
	}
	if r.UserLabels != nil {
		rUserLabels := make(map[string]interface{})
		for k, v := range r.UserLabels {
			rUserLabels[k] = v
		}
		u.Object["userLabels"] = rUserLabels
	}
	return u
}

func UnstructuredToServiceLevelObjective(u *unstructured.Resource) (*dclService.ServiceLevelObjective, error) {
	r := &dclService.ServiceLevelObjective{}
	if _, ok := u.Object["calendarPeriod"]; ok {
		if s, ok := u.Object["calendarPeriod"].(string); ok {
			r.CalendarPeriod = dclService.ServiceLevelObjectiveCalendarPeriodEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.CalendarPeriod: expected string")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deleteTime"]; ok {
		if s, ok := u.Object["deleteTime"].(string); ok {
			r.DeleteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DeleteTime: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["goal"]; ok {
		if f, ok := u.Object["goal"].(float64); ok {
			r.Goal = dcl.Float64(f)
		} else {
			return nil, fmt.Errorf("r.Goal: expected float64")
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
	if _, ok := u.Object["rollingPeriod"]; ok {
		if s, ok := u.Object["rollingPeriod"].(string); ok {
			r.RollingPeriod = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RollingPeriod: expected string")
		}
	}
	if _, ok := u.Object["service"]; ok {
		if s, ok := u.Object["service"].(string); ok {
			r.Service = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Service: expected string")
		}
	}
	if _, ok := u.Object["serviceLevelIndicator"]; ok {
		if rServiceLevelIndicator, ok := u.Object["serviceLevelIndicator"].(map[string]interface{}); ok {
			r.ServiceLevelIndicator = &dclService.ServiceLevelObjectiveServiceLevelIndicator{}
			if _, ok := rServiceLevelIndicator["basicSli"]; ok {
				if rServiceLevelIndicatorBasicSli, ok := rServiceLevelIndicator["basicSli"].(map[string]interface{}); ok {
					r.ServiceLevelIndicator.BasicSli = &dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSli{}
					if _, ok := rServiceLevelIndicatorBasicSli["availability"]; ok {
						if _, ok := rServiceLevelIndicatorBasicSli["availability"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.BasicSli.Availability = &dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Availability: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorBasicSli["latency"]; ok {
						if rServiceLevelIndicatorBasicSliLatency, ok := rServiceLevelIndicatorBasicSli["latency"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.BasicSli.Latency = &dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{}
							if _, ok := rServiceLevelIndicatorBasicSliLatency["experience"]; ok {
								if s, ok := rServiceLevelIndicatorBasicSliLatency["experience"].(string); ok {
									r.ServiceLevelIndicator.BasicSli.Latency.Experience = dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Latency.Experience: expected string")
								}
							}
							if _, ok := rServiceLevelIndicatorBasicSliLatency["threshold"]; ok {
								if s, ok := rServiceLevelIndicatorBasicSliLatency["threshold"].(string); ok {
									r.ServiceLevelIndicator.BasicSli.Latency.Threshold = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Latency.Threshold: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Latency: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorBasicSli["location"]; ok {
						if s, ok := rServiceLevelIndicatorBasicSli["location"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.ServiceLevelIndicator.BasicSli.Location = append(r.ServiceLevelIndicator.BasicSli.Location, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Location: expected []interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorBasicSli["method"]; ok {
						if s, ok := rServiceLevelIndicatorBasicSli["method"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.ServiceLevelIndicator.BasicSli.Method = append(r.ServiceLevelIndicator.BasicSli.Method, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Method: expected []interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorBasicSli["operationAvailability"]; ok {
						if _, ok := rServiceLevelIndicatorBasicSli["operationAvailability"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.BasicSli.OperationAvailability = &dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.OperationAvailability: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorBasicSli["operationLatency"]; ok {
						if rServiceLevelIndicatorBasicSliOperationLatency, ok := rServiceLevelIndicatorBasicSli["operationLatency"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.BasicSli.OperationLatency = &dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{}
							if _, ok := rServiceLevelIndicatorBasicSliOperationLatency["experience"]; ok {
								if s, ok := rServiceLevelIndicatorBasicSliOperationLatency["experience"].(string); ok {
									r.ServiceLevelIndicator.BasicSli.OperationLatency.Experience = dclService.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.OperationLatency.Experience: expected string")
								}
							}
							if _, ok := rServiceLevelIndicatorBasicSliOperationLatency["threshold"]; ok {
								if s, ok := rServiceLevelIndicatorBasicSliOperationLatency["threshold"].(string); ok {
									r.ServiceLevelIndicator.BasicSli.OperationLatency.Threshold = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.OperationLatency.Threshold: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.OperationLatency: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorBasicSli["version"]; ok {
						if s, ok := rServiceLevelIndicatorBasicSli["version"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.ServiceLevelIndicator.BasicSli.Version = append(r.ServiceLevelIndicator.BasicSli.Version, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli.Version: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ServiceLevelIndicator.BasicSli: expected map[string]interface{}")
				}
			}
			if _, ok := rServiceLevelIndicator["requestBased"]; ok {
				if rServiceLevelIndicatorRequestBased, ok := rServiceLevelIndicator["requestBased"].(map[string]interface{}); ok {
					r.ServiceLevelIndicator.RequestBased = &dclService.ServiceLevelObjectiveServiceLevelIndicatorRequestBased{}
					if _, ok := rServiceLevelIndicatorRequestBased["distributionCut"]; ok {
						if rServiceLevelIndicatorRequestBasedDistributionCut, ok := rServiceLevelIndicatorRequestBased["distributionCut"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.RequestBased.DistributionCut = &dclService.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{}
							if _, ok := rServiceLevelIndicatorRequestBasedDistributionCut["distributionFilter"]; ok {
								if s, ok := rServiceLevelIndicatorRequestBasedDistributionCut["distributionFilter"].(string); ok {
									r.ServiceLevelIndicator.RequestBased.DistributionCut.DistributionFilter = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.DistributionCut.DistributionFilter: expected string")
								}
							}
							if _, ok := rServiceLevelIndicatorRequestBasedDistributionCut["range"]; ok {
								if rServiceLevelIndicatorRequestBasedDistributionCutRange, ok := rServiceLevelIndicatorRequestBasedDistributionCut["range"].(map[string]interface{}); ok {
									r.ServiceLevelIndicator.RequestBased.DistributionCut.Range = &dclService.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{}
									if _, ok := rServiceLevelIndicatorRequestBasedDistributionCutRange["max"]; ok {
										if f, ok := rServiceLevelIndicatorRequestBasedDistributionCutRange["max"].(float64); ok {
											r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Max = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Max: expected float64")
										}
									}
									if _, ok := rServiceLevelIndicatorRequestBasedDistributionCutRange["min"]; ok {
										if f, ok := rServiceLevelIndicatorRequestBasedDistributionCutRange["min"].(float64); ok {
											r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Min = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.DistributionCut.Range.Min: expected float64")
										}
									}
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.DistributionCut.Range: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.DistributionCut: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorRequestBased["goodTotalRatio"]; ok {
						if rServiceLevelIndicatorRequestBasedGoodTotalRatio, ok := rServiceLevelIndicatorRequestBased["goodTotalRatio"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.RequestBased.GoodTotalRatio = &dclService.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{}
							if _, ok := rServiceLevelIndicatorRequestBasedGoodTotalRatio["badServiceFilter"]; ok {
								if s, ok := rServiceLevelIndicatorRequestBasedGoodTotalRatio["badServiceFilter"].(string); ok {
									r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.BadServiceFilter = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.BadServiceFilter: expected string")
								}
							}
							if _, ok := rServiceLevelIndicatorRequestBasedGoodTotalRatio["goodServiceFilter"]; ok {
								if s, ok := rServiceLevelIndicatorRequestBasedGoodTotalRatio["goodServiceFilter"].(string); ok {
									r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.GoodServiceFilter = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.GoodServiceFilter: expected string")
								}
							}
							if _, ok := rServiceLevelIndicatorRequestBasedGoodTotalRatio["totalServiceFilter"]; ok {
								if s, ok := rServiceLevelIndicatorRequestBasedGoodTotalRatio["totalServiceFilter"].(string); ok {
									r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.TotalServiceFilter = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.GoodTotalRatio.TotalServiceFilter: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased.GoodTotalRatio: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ServiceLevelIndicator.RequestBased: expected map[string]interface{}")
				}
			}
			if _, ok := rServiceLevelIndicator["windowsBased"]; ok {
				if rServiceLevelIndicatorWindowsBased, ok := rServiceLevelIndicator["windowsBased"].(map[string]interface{}); ok {
					r.ServiceLevelIndicator.WindowsBased = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased{}
					if _, ok := rServiceLevelIndicatorWindowsBased["goodBadMetricFilter"]; ok {
						if s, ok := rServiceLevelIndicatorWindowsBased["goodBadMetricFilter"].(string); ok {
							r.ServiceLevelIndicator.WindowsBased.GoodBadMetricFilter = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodBadMetricFilter: expected string")
						}
					}
					if _, ok := rServiceLevelIndicatorWindowsBased["goodTotalRatioThreshold"]; ok {
						if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold, ok := rServiceLevelIndicatorWindowsBased["goodTotalRatioThreshold"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{}
							if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["basicSliPerformance"]; ok {
								if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["basicSliPerformance"].(map[string]interface{}); ok {
									r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["availability"]; ok {
										if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["availability"].(map[string]interface{}); ok {
											r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Availability = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Availability: expected map[string]interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["latency"]; ok {
										if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["latency"].(map[string]interface{}); ok {
											r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency["experience"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency["experience"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Experience = dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Experience: expected string")
												}
											}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency["threshold"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency["threshold"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Threshold = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency.Threshold: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Latency: expected map[string]interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["location"]; ok {
										if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["location"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Location = append(r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Location, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Location: expected []interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["method"]; ok {
										if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["method"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Method = append(r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Method, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Method: expected []interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["operationAvailability"]; ok {
										if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["operationAvailability"].(map[string]interface{}); ok {
											r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationAvailability = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationAvailability: expected map[string]interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["operationLatency"]; ok {
										if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["operationLatency"].(map[string]interface{}); ok {
											r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency["experience"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency["experience"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Experience = dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumRef(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Experience: expected string")
												}
											}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency["threshold"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency["threshold"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Threshold = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency.Threshold: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.OperationLatency: expected map[string]interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["version"]; ok {
										if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance["version"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Version = append(r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Version, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance.Version: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.BasicSliPerformance: expected map[string]interface{}")
								}
							}
							if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["performance"]; ok {
								if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["performance"].(map[string]interface{}); ok {
									r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance["distributionCut"]; ok {
										if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance["distributionCut"].(map[string]interface{}); ok {
											r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut["distributionFilter"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut["distributionFilter"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.DistributionFilter = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.DistributionFilter: expected string")
												}
											}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut["range"]; ok {
												if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut["range"].(map[string]interface{}); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{}
													if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange["max"]; ok {
														if f, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange["max"].(float64); ok {
															r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Max = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Max: expected float64")
														}
													}
													if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange["min"]; ok {
														if f, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange["min"].(float64); ok {
															r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Min = dcl.Float64(f)
														} else {
															return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range.Min: expected float64")
														}
													}
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut.Range: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.DistributionCut: expected map[string]interface{}")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance["goodTotalRatio"]; ok {
										if rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance["goodTotalRatio"].(map[string]interface{}); ok {
											r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["badServiceFilter"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["badServiceFilter"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.BadServiceFilter = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.BadServiceFilter: expected string")
												}
											}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["goodServiceFilter"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["goodServiceFilter"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.GoodServiceFilter = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.GoodServiceFilter: expected string")
												}
											}
											if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["totalServiceFilter"]; ok {
												if s, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio["totalServiceFilter"].(string); ok {
													r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.TotalServiceFilter = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio.TotalServiceFilter: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance.GoodTotalRatio: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Performance: expected map[string]interface{}")
								}
							}
							if _, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["threshold"]; ok {
								if f, ok := rServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold["threshold"].(float64); ok {
									r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Threshold = dcl.Float64(f)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold.Threshold: expected float64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.GoodTotalRatioThreshold: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorWindowsBased["metricMeanInRange"]; ok {
						if rServiceLevelIndicatorWindowsBasedMetricMeanInRange, ok := rServiceLevelIndicatorWindowsBased["metricMeanInRange"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{}
							if _, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRange["range"]; ok {
								if rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRange["range"].(map[string]interface{}); ok {
									r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{}
									if _, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange["max"]; ok {
										if f, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange["max"].(float64); ok {
											r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Max = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Max: expected float64")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange["min"]; ok {
										if f, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange["min"].(float64); ok {
											r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Min = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range.Min: expected float64")
										}
									}
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.Range: expected map[string]interface{}")
								}
							}
							if _, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRange["timeSeries"]; ok {
								if s, ok := rServiceLevelIndicatorWindowsBasedMetricMeanInRange["timeSeries"].(string); ok {
									r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.TimeSeries = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange.TimeSeries: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricMeanInRange: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorWindowsBased["metricSumInRange"]; ok {
						if rServiceLevelIndicatorWindowsBasedMetricSumInRange, ok := rServiceLevelIndicatorWindowsBased["metricSumInRange"].(map[string]interface{}); ok {
							r.ServiceLevelIndicator.WindowsBased.MetricSumInRange = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{}
							if _, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRange["range"]; ok {
								if rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRange["range"].(map[string]interface{}); ok {
									r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range = &dclService.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{}
									if _, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange["max"]; ok {
										if f, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange["max"].(float64); ok {
											r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Max = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Max: expected float64")
										}
									}
									if _, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange["min"]; ok {
										if f, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRangeRange["min"].(float64); ok {
											r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Min = dcl.Float64(f)
										} else {
											return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range.Min: expected float64")
										}
									}
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.Range: expected map[string]interface{}")
								}
							}
							if _, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRange["timeSeries"]; ok {
								if s, ok := rServiceLevelIndicatorWindowsBasedMetricSumInRange["timeSeries"].(string); ok {
									r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.TimeSeries = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricSumInRange.TimeSeries: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.MetricSumInRange: expected map[string]interface{}")
						}
					}
					if _, ok := rServiceLevelIndicatorWindowsBased["windowPeriod"]; ok {
						if s, ok := rServiceLevelIndicatorWindowsBased["windowPeriod"].(string); ok {
							r.ServiceLevelIndicator.WindowsBased.WindowPeriod = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased.WindowPeriod: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ServiceLevelIndicator.WindowsBased: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ServiceLevelIndicator: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["serviceManagementOwned"]; ok {
		if b, ok := u.Object["serviceManagementOwned"].(bool); ok {
			r.ServiceManagementOwned = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.ServiceManagementOwned: expected bool")
		}
	}
	if _, ok := u.Object["userLabels"]; ok {
		if rUserLabels, ok := u.Object["userLabels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rUserLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.UserLabels = m
		} else {
			return nil, fmt.Errorf("r.UserLabels: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetServiceLevelObjective(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceLevelObjective(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetServiceLevelObjective(ctx, r)
	if err != nil {
		return nil, err
	}
	return ServiceLevelObjectiveToUnstructured(r), nil
}

func ListServiceLevelObjective(ctx context.Context, config *dcl.Config, project string, service string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListServiceLevelObjective(ctx, project, service)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ServiceLevelObjectiveToUnstructured(r))
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

func ApplyServiceLevelObjective(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceLevelObjective(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServiceLevelObjective(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyServiceLevelObjective(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ServiceLevelObjectiveToUnstructured(r), nil
}

func ServiceLevelObjectiveHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceLevelObjective(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServiceLevelObjective(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyServiceLevelObjective(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteServiceLevelObjective(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceLevelObjective(u)
	if err != nil {
		return err
	}
	return c.DeleteServiceLevelObjective(ctx, r)
}

func ServiceLevelObjectiveID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToServiceLevelObjective(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ServiceLevelObjective) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"monitoring",
		"ServiceLevelObjective",
		"beta",
	}
}

func (r *ServiceLevelObjective) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceLevelObjective) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceLevelObjective) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *ServiceLevelObjective) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceLevelObjective) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceLevelObjective) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *ServiceLevelObjective) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetServiceLevelObjective(ctx, config, resource)
}

func (r *ServiceLevelObjective) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyServiceLevelObjective(ctx, config, resource, opts...)
}

func (r *ServiceLevelObjective) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ServiceLevelObjectiveHasDiff(ctx, config, resource, opts...)
}

func (r *ServiceLevelObjective) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteServiceLevelObjective(ctx, config, resource)
}

func (r *ServiceLevelObjective) ID(resource *unstructured.Resource) (string, error) {
	return ServiceLevelObjectiveID(resource)
}

func init() {
	unstructured.Register(&ServiceLevelObjective{})
}
