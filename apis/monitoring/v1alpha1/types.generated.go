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

package v1alpha1


// +kcc:proto=google.monitoring.v3.BasicSli
type BasicSli struct {
	// OPTIONAL: The set of RPCs to which this SLI is relevant. Telemetry from
	//  other methods will not be used to calculate performance for this SLI. If
	//  omitted, this SLI applies to all the Service's methods. For service types
	//  that don't support breaking down by method, setting this field will result
	//  in an error.
	// +kcc:proto:field=google.monitoring.v3.BasicSli.method
	Method []string `json:"method,omitempty"`

	// OPTIONAL: The set of locations to which this SLI is relevant. Telemetry
	//  from other locations will not be used to calculate performance for this
	//  SLI. If omitted, this SLI applies to all locations in which the Service has
	//  activity. For service types that don't support breaking down by location,
	//  setting this field will result in an error.
	// +kcc:proto:field=google.monitoring.v3.BasicSli.location
	Location []string `json:"location,omitempty"`

	// OPTIONAL: The set of API versions to which this SLI is relevant. Telemetry
	//  from other API versions will not be used to calculate performance for this
	//  SLI. If omitted, this SLI applies to all API versions. For service types
	//  that don't support breaking down by version, setting this field will result
	//  in an error.
	// +kcc:proto:field=google.monitoring.v3.BasicSli.version
	Version []string `json:"version,omitempty"`

	// Good service is defined to be the count of requests made to this service
	//  that return successfully.
	// +kcc:proto:field=google.monitoring.v3.BasicSli.availability
	Availability *BasicSli_AvailabilityCriteria `json:"availability,omitempty"`

	// Good service is defined to be the count of requests made to this service
	//  that are fast enough with respect to `latency.threshold`.
	// +kcc:proto:field=google.monitoring.v3.BasicSli.latency
	Latency *BasicSli_LatencyCriteria `json:"latency,omitempty"`
}

// +kcc:proto=google.monitoring.v3.BasicSli.AvailabilityCriteria
type BasicSli_AvailabilityCriteria struct {
}

// +kcc:proto=google.monitoring.v3.BasicSli.LatencyCriteria
type BasicSli_LatencyCriteria struct {
	// Good service is defined to be the count of requests made to this service
	//  that return in no more than `threshold`.
	// +kcc:proto:field=google.monitoring.v3.BasicSli.LatencyCriteria.threshold
	Threshold *string `json:"threshold,omitempty"`
}

// +kcc:proto=google.monitoring.v3.DistributionCut
type DistributionCut struct {
	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	//  specifying a `TimeSeries` aggregating values. Must have `ValueType =
	//  DISTRIBUTION` and `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.
	// +kcc:proto:field=google.monitoring.v3.DistributionCut.distribution_filter
	DistributionFilter *string `json:"distributionFilter,omitempty"`

	// Range of values considered "good." For a one-sided range, set one bound to
	//  an infinite value.
	// +kcc:proto:field=google.monitoring.v3.DistributionCut.range
	Range *Range `json:"range,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Range
type Range struct {
	// Range minimum.
	// +kcc:proto:field=google.monitoring.v3.Range.min
	Min *float64 `json:"min,omitempty"`

	// Range maximum.
	// +kcc:proto:field=google.monitoring.v3.Range.max
	Max *float64 `json:"max,omitempty"`
}

// +kcc:proto=google.monitoring.v3.RequestBasedSli
type RequestBasedSli struct {
	// `good_total_ratio` is used when the ratio of `good_service` to
	//  `total_service` is computed from two `TimeSeries`.
	// +kcc:proto:field=google.monitoring.v3.RequestBasedSli.good_total_ratio
	GoodTotalRatio *TimeSeriesRatio `json:"goodTotalRatio,omitempty"`

	// `distribution_cut` is used when `good_service` is a count of values
	//  aggregated in a `Distribution` that fall into a good range. The
	//  `total_service` is the total count of all values aggregated in the
	//  `Distribution`.
	// +kcc:proto:field=google.monitoring.v3.RequestBasedSli.distribution_cut
	DistributionCut *DistributionCut `json:"distributionCut,omitempty"`
}

// +kcc:proto=google.monitoring.v3.ServiceLevelIndicator
type ServiceLevelIndicator struct {
	// Basic SLI on a well-known service type.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelIndicator.basic_sli
	BasicSli *BasicSli `json:"basicSli,omitempty"`

	// Request-based SLIs
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelIndicator.request_based
	RequestBased *RequestBasedSli `json:"requestBased,omitempty"`

	// Windows-based SLIs
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelIndicator.windows_based
	WindowsBased *WindowsBasedSli `json:"windowsBased,omitempty"`
}

// +kcc:proto=google.monitoring.v3.ServiceLevelObjective
type ServiceLevelObjective struct {
	// Identifier. Resource name for this `ServiceLevelObjective`. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.name
	Name *string `json:"name,omitempty"`

	// Name used for UI elements listing this SLO.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The definition of good service, used to measure and calculate the quality
	//  of the `Service`'s performance with respect to a single aspect of service
	//  quality.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.service_level_indicator
	ServiceLevelIndicator *ServiceLevelIndicator `json:"serviceLevelIndicator,omitempty"`

	// The fraction of service that must be good in order for this objective to be
	//  met. `0 < goal <= 0.9999`.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.goal
	Goal *float64 `json:"goal,omitempty"`

	// A rolling time period, semantically "in the past `<rolling_period>`".
	//  Must be an integer multiple of 1 day no larger than 30 days.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.rolling_period
	RollingPeriod *string `json:"rollingPeriod,omitempty"`

	// A calendar period, semantically "since the start of the current
	//  `<calendar_period>`". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and
	//  `MONTH` are supported.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.calendar_period
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	// Labels which have been used to annotate the service-level objective. Label
	//  keys must start with a letter. Label keys and values may contain lowercase
	//  letters, numbers, underscores, and dashes. Label keys and values have a
	//  maximum length of 63 characters, and must be less than 128 bytes in size.
	//  Up to 64 label entries may be stored. For labels which do not have a
	//  semantic value, the empty string may be supplied for the label value.
	// +kcc:proto:field=google.monitoring.v3.ServiceLevelObjective.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`
}

// +kcc:proto=google.monitoring.v3.TimeSeriesRatio
type TimeSeriesRatio struct {
	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	//  specifying a `TimeSeries` quantifying good service provided. Must have
	//  `ValueType = DOUBLE` or `ValueType = INT64` and must have `MetricKind =
	//  DELTA` or `MetricKind = CUMULATIVE`.
	// +kcc:proto:field=google.monitoring.v3.TimeSeriesRatio.good_service_filter
	GoodServiceFilter *string `json:"goodServiceFilter,omitempty"`

	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	//  specifying a `TimeSeries` quantifying bad service, either demanded service
	//  that was not provided or demanded service that was of inadequate quality.
	//  Must have `ValueType = DOUBLE` or `ValueType = INT64` and must have
	//  `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.
	// +kcc:proto:field=google.monitoring.v3.TimeSeriesRatio.bad_service_filter
	BadServiceFilter *string `json:"badServiceFilter,omitempty"`

	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	//  specifying a `TimeSeries` quantifying total demanded service. Must have
	//  `ValueType = DOUBLE` or `ValueType = INT64` and must have `MetricKind =
	//  DELTA` or `MetricKind = CUMULATIVE`.
	// +kcc:proto:field=google.monitoring.v3.TimeSeriesRatio.total_service_filter
	TotalServiceFilter *string `json:"totalServiceFilter,omitempty"`
}

// +kcc:proto=google.monitoring.v3.WindowsBasedSli
type WindowsBasedSli struct {
	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	//  specifying a `TimeSeries` with `ValueType = BOOL`. The window is good if
	//  any `true` values appear in the window.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.good_bad_metric_filter
	GoodBadMetricFilter *string `json:"goodBadMetricFilter,omitempty"`

	// A window is good if its `performance` is high enough.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.good_total_ratio_threshold
	GoodTotalRatioThreshold *WindowsBasedSli_PerformanceThreshold `json:"goodTotalRatioThreshold,omitempty"`

	// A window is good if the metric's value is in a good range, averaged
	//  across returned streams.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.metric_mean_in_range
	MetricMeanInRange *WindowsBasedSli_MetricRange `json:"metricMeanInRange,omitempty"`

	// A window is good if the metric's value is in a good range, summed across
	//  returned streams.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.metric_sum_in_range
	MetricSumInRange *WindowsBasedSli_MetricRange `json:"metricSumInRange,omitempty"`

	// Duration over which window quality is evaluated. Must be an integer
	//  fraction of a day and at least `60s`.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.window_period
	WindowPeriod *string `json:"windowPeriod,omitempty"`
}

// +kcc:proto=google.monitoring.v3.WindowsBasedSli.MetricRange
type WindowsBasedSli_MetricRange struct {
	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	//  specifying the `TimeSeries` to use for evaluating window quality.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.MetricRange.time_series
	TimeSeries *string `json:"timeSeries,omitempty"`

	// Range of values considered "good." For a one-sided range, set one bound
	//  to an infinite value.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.MetricRange.range
	Range *Range `json:"range,omitempty"`
}

// +kcc:proto=google.monitoring.v3.WindowsBasedSli.PerformanceThreshold
type WindowsBasedSli_PerformanceThreshold struct {
	// `RequestBasedSli` to evaluate to judge window quality.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.PerformanceThreshold.performance
	Performance *RequestBasedSli `json:"performance,omitempty"`

	// `BasicSli` to evaluate to judge window quality.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.PerformanceThreshold.basic_sli_performance
	BasicSliPerformance *BasicSli `json:"basicSliPerformance,omitempty"`

	// If window `performance >= threshold`, the window is counted as good.
	// +kcc:proto:field=google.monitoring.v3.WindowsBasedSli.PerformanceThreshold.threshold
	Threshold *float64 `json:"threshold,omitempty"`
}
