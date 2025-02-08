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


// +kcc:proto=google.monitoring.v3.Aggregation
type Aggregation struct {
	// The `alignment_period` specifies a time interval, in seconds, that is used
	//  to divide the data in all the
	//  [time series][google.monitoring.v3.TimeSeries] into consistent blocks of
	//  time. This will be done before the per-series aligner can be applied to
	//  the data.
	//
	//  The value must be at least 60 seconds. If a per-series
	//  aligner other than `ALIGN_NONE` is specified, this field is required or an
	//  error is returned. If no per-series aligner is specified, or the aligner
	//  `ALIGN_NONE` is specified, then this field is ignored.
	//
	//  The maximum value of the `alignment_period` is 104 weeks (2 years) for
	//  charts, and 90,000 seconds (25 hours) for alerting policies.
	// +kcc:proto:field=google.monitoring.v3.Aggregation.alignment_period
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	// An `Aligner` describes how to bring the data points in a single
	//  time series into temporal alignment. Except for `ALIGN_NONE`, all
	//  alignments cause all the data points in an `alignment_period` to be
	//  mathematically grouped together, resulting in a single data point for
	//  each `alignment_period` with end timestamp at the end of the period.
	//
	//  Not all alignment operations may be applied to all time series. The valid
	//  choices depend on the `metric_kind` and `value_type` of the original time
	//  series. Alignment can change the `metric_kind` or the `value_type` of
	//  the time series.
	//
	//  Time series data must be aligned in order to perform cross-time
	//  series reduction. If `cross_series_reducer` is specified, then
	//  `per_series_aligner` must be specified and not equal to `ALIGN_NONE`
	//  and `alignment_period` must be specified; otherwise, an error is
	//  returned.
	// +kcc:proto:field=google.monitoring.v3.Aggregation.per_series_aligner
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`

	// The reduction operation to be used to combine time series into a single
	//  time series, where the value of each data point in the resulting series is
	//  a function of all the already aligned values in the input time series.
	//
	//  Not all reducer operations can be applied to all time series. The valid
	//  choices depend on the `metric_kind` and the `value_type` of the original
	//  time series. Reduction can yield a time series with a different
	//  `metric_kind` or `value_type` than the input time series.
	//
	//  Time series data must first be aligned (see `per_series_aligner`) in order
	//  to perform cross-time series reduction. If `cross_series_reducer` is
	//  specified, then `per_series_aligner` must be specified, and must not be
	//  `ALIGN_NONE`. An `alignment_period` must also be specified; otherwise, an
	//  error is returned.
	// +kcc:proto:field=google.monitoring.v3.Aggregation.cross_series_reducer
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	// The set of fields to preserve when `cross_series_reducer` is
	//  specified. The `group_by_fields` determine how the time series are
	//  partitioned into subsets prior to applying the aggregation
	//  operation. Each subset contains time series that have the same
	//  value for each of the grouping fields. Each individual time
	//  series is a member of exactly one subset. The
	//  `cross_series_reducer` is applied to each subset of time series.
	//  It is not possible to reduce across different resource types, so
	//  this field implicitly contains `resource.type`.  Fields not
	//  specified in `group_by_fields` are aggregated away.  If
	//  `group_by_fields` is not specified and all the time series have
	//  the same resource type, then the time series are aggregated into
	//  a single output time series. If `cross_series_reducer` is not
	//  defined, this field is ignored.
	// +kcc:proto:field=google.monitoring.v3.Aggregation.group_by_fields
	GroupByFields []string `json:"groupByFields,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy
type AlertPolicy struct {
	// Identifier. Required if the policy exists. The resource name for this
	//  policy. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID]
	//
	//  `[ALERT_POLICY_ID]` is assigned by Cloud Monitoring when the policy
	//  is created. When calling the
	//  [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	//  method, do not include the `name` field in the alerting policy passed as
	//  part of the request.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.name
	Name *string `json:"name,omitempty"`

	// A short name or phrase used to identify the policy in dashboards,
	//  notifications, and incidents. To avoid confusion, don't use the same
	//  display name for multiple policies in the same project. The name is
	//  limited to 512 Unicode characters.
	//
	//  The convention for the display_name of a PrometheusQueryLanguageCondition
	//  is "{rule group name}/{alert name}", where the {rule group name} and
	//  {alert name} should be taken from the corresponding Prometheus
	//  configuration file. This convention is not enforced.
	//  In any case the display_name is not a unique key of the AlertPolicy.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Documentation that is included with notifications and incidents related to
	//  this policy. Best practice is for the documentation to include information
	//  to help responders understand, mitigate, escalate, and correct the
	//  underlying problems detected by the alerting policy. Notification channels
	//  that have limited capacity might not show this documentation.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.documentation
	Documentation *AlertPolicy_Documentation `json:"documentation,omitempty"`

	// User-supplied key/value data to be used for organizing and
	//  identifying the `AlertPolicy` objects.
	//
	//  The field can contain up to 64 entries. Each key and value is limited to
	//  63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	//  values can contain only lowercase letters, numerals, underscores, and
	//  dashes. Keys must begin with a letter.
	//
	//  Note that Prometheus {alert name} is a
	//  [valid Prometheus label
	//  names](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels),
	//  whereas Prometheus {rule group} is an unrestricted UTF-8 string.
	//  This means that they cannot be stored as-is in user labels, because
	//  they may contain characters that are not allowed in user-label values.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// A list of conditions for the policy. The conditions are combined by AND or
	//  OR according to the `combiner` field. If the combined conditions evaluate
	//  to true, then an incident is created. A policy can have from one to six
	//  conditions.
	//  If `condition_time_series_query_language` is present, it must be the only
	//  `condition`.
	//  If `condition_monitoring_query_language` is present, it must be the only
	//  `condition`.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.conditions
	Conditions []AlertPolicy_Condition `json:"conditions,omitempty"`

	// How to combine the results of multiple conditions to determine if an
	//  incident should be opened.
	//  If `condition_time_series_query_language` is present, this must be
	//  `COMBINE_UNSPECIFIED`.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.combiner
	Combiner *string `json:"combiner,omitempty"`

	// Whether or not the policy is enabled. On write, the default interpretation
	//  if unset is that the policy is enabled. On read, clients should not make
	//  any assumption about the state if it has not been populated. The
	//  field should always be populated on List and Get operations, unless
	//  a field projection has been specified that strips it out.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Read-only description of how the alerting policy is invalid. This field is
	//  only set when the alerting policy is invalid. An invalid alerting policy
	//  will not generate incidents.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.validity
	Validity *Status `json:"validity,omitempty"`

	// Identifies the notification channels to which notifications should be sent
	//  when incidents are opened or closed or when new violations occur on
	//  an already opened incident. Each element of this array corresponds to
	//  the `name` field in each of the
	//  [`NotificationChannel`][google.monitoring.v3.NotificationChannel]
	//  objects that are returned from the [`ListNotificationChannels`]
	//  [google.monitoring.v3.NotificationChannelService.ListNotificationChannels]
	//  method. The format of the entries in this field is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.notification_channels
	NotificationChannels []string `json:"notificationChannels,omitempty"`

	// A read-only record of the creation of the alerting policy. If provided
	//  in a call to create or update, this field will be ignored.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.creation_record
	CreationRecord *MutationRecord `json:"creationRecord,omitempty"`

	// A read-only record of the most recent change to the alerting policy. If
	//  provided in a call to create or update, this field will be ignored.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.mutation_record
	MutationRecord *MutationRecord `json:"mutationRecord,omitempty"`

	// Control over how this alerting policy's notification channels are notified.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.alert_strategy
	AlertStrategy *AlertPolicy_AlertStrategy `json:"alertStrategy,omitempty"`

	// Optional. The severity of an alerting policy indicates how important
	//  incidents generated by that policy are. The severity level will be
	//  displayed on the Incident detail page and in notifications.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.severity
	Severity *string `json:"severity,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.AlertStrategy
type AlertPolicy_AlertStrategy struct {
	// Required for log-based alerting policies, i.e. policies with a `LogMatch`
	//  condition.
	//
	//  This limit is not implemented for alerting policies that do not have
	//  a LogMatch condition.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.notification_rate_limit
	NotificationRateLimit *AlertPolicy_AlertStrategy_NotificationRateLimit `json:"notificationRateLimit,omitempty"`

	// For log-based alert policies, the notification prompts is always
	//  [OPENED]. For non log-based alert policies, the notification prompts can
	//  be [OPENED] or [OPENED, CLOSED].
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.notification_prompts
	NotificationPrompts []string `json:"notificationPrompts,omitempty"`

	// If an alerting policy that was active has no data for this long, any open
	//  incidents will close
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.auto_close
	AutoClose *string `json:"autoClose,omitempty"`

	// Control how notifications will be sent out, on a per-channel basis.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.notification_channel_strategy
	NotificationChannelStrategy []AlertPolicy_AlertStrategy_NotificationChannelStrategy `json:"notificationChannelStrategy,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.AlertStrategy.NotificationChannelStrategy
type AlertPolicy_AlertStrategy_NotificationChannelStrategy struct {
	// The full REST resource name for the notification channels that these
	//  settings apply to. Each of these correspond to the name field in one
	//  of the NotificationChannel objects referenced in the
	//  notification_channels field of this AlertPolicy.
	//  The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.NotificationChannelStrategy.notification_channel_names
	NotificationChannelNames []string `json:"notificationChannelNames,omitempty"`

	// The frequency at which to send reminder notifications for open
	//  incidents.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.NotificationChannelStrategy.renotify_interval
	RenotifyInterval *string `json:"renotifyInterval,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.AlertStrategy.NotificationRateLimit
type AlertPolicy_AlertStrategy_NotificationRateLimit struct {
	// Not more than one notification per `period`.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.NotificationRateLimit.period
	Period *string `json:"period,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition
type AlertPolicy_Condition struct {
	// Required if the condition exists. The unique resource name for this
	//  condition. Its format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	//
	//  `[CONDITION_ID]` is assigned by Cloud Monitoring when the
	//  condition is created as part of a new or updated alerting policy.
	//
	//  When calling the
	//  [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	//  method, do not include the `name` field in the conditions of the
	//  requested alerting policy. Cloud Monitoring creates the
	//  condition identifiers and includes them in the new policy.
	//
	//  When calling the
	//  [alertPolicies.update][google.monitoring.v3.AlertPolicyService.UpdateAlertPolicy]
	//  method to update a policy, including a condition `name` causes the
	//  existing condition to be updated. Conditions without names are added to
	//  the updated policy. Existing conditions are deleted if they are not
	//  updated.
	//
	//  Best practice is to preserve `[CONDITION_ID]` if you make only small
	//  changes, such as those to condition thresholds, durations, or trigger
	//  values.  Otherwise, treat the change as a new condition and let the
	//  existing condition be deleted.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.name
	Name *string `json:"name,omitempty"`

	// A short name or phrase used to identify the condition in dashboards,
	//  notifications, and incidents. To avoid confusion, don't use the same
	//  display name for multiple conditions in the same policy.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A condition that compares a time series against a threshold.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_threshold
	ConditionThreshold *AlertPolicy_Condition_MetricThreshold `json:"conditionThreshold,omitempty"`

	// A condition that checks that a time series continues to
	//  receive new data points.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_absent
	ConditionAbsent *AlertPolicy_Condition_MetricAbsence `json:"conditionAbsent,omitempty"`

	// A condition that checks for log messages matching given constraints. If
	//  set, no other conditions can be present.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_matched_log
	ConditionMatchedLog *AlertPolicy_Condition_LogMatch `json:"conditionMatchedLog,omitempty"`

	// A condition that uses the Monitoring Query Language to define
	//  alerts.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_monitoring_query_language
	ConditionMonitoringQueryLanguage *AlertPolicy_Condition_MonitoringQueryLanguageCondition `json:"conditionMonitoringQueryLanguage,omitempty"`

	// A condition that uses the Prometheus query language to define alerts.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_prometheus_query_language
	ConditionPrometheusQueryLanguage *AlertPolicy_Condition_PrometheusQueryLanguageCondition `json:"conditionPrometheusQueryLanguage,omitempty"`

	// A condition that periodically evaluates a SQL query result.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_sql
	ConditionSql *AlertPolicy_Condition_SqlCondition `json:"conditionSql,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.LogMatch
type AlertPolicy_Condition_LogMatch struct {
	// Required. A logs-based filter. See [Advanced Logs
	//  Queries](https://cloud.google.com/logging/docs/view/advanced-queries)
	//  for how this filter should be constructed.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.LogMatch.filter
	Filter *string `json:"filter,omitempty"`

	// Optional. A map from a label key to an extractor expression, which is
	//  used to extract the value for this label key. Each entry in this map is
	//  a specification for how data should be extracted from log entries that
	//  match `filter`. Each combination of extracted values is treated as a
	//  separate rule for the purposes of triggering notifications. Label keys
	//  and corresponding values can be used in notifications generated by this
	//  condition.
	//
	//  Please see [the documentation on logs-based metric
	//  `valueExtractor`s](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric.FIELDS.value_extractor)
	//  for syntax and examples.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.LogMatch.label_extractors
	LabelExtractors map[string]string `json:"labelExtractors,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.MetricAbsence
type AlertPolicy_Condition_MetricAbsence struct {
	// Required. A
	//  [filter](https://cloud.google.com/monitoring/api/v3/filters) that
	//  identifies which time series should be compared with the threshold.
	//
	//  The filter is similar to the one that is specified in the
	//  [`ListTimeSeries`
	//  request](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list)
	//  (that call is useful to verify the time series that will be retrieved /
	//  processed). The filter must specify the metric type and the resource
	//  type. Optionally, it can specify resource labels and metric labels.
	//  This field must not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricAbsence.filter
	Filter *string `json:"filter,omitempty"`

	// Specifies the alignment of data points in individual time series as
	//  well as how to combine the retrieved time series together (such as
	//  when aggregating multiple streams on each resource to a single
	//  stream for each resource or when aggregating streams across all
	//  members of a group of resources). Multiple aggregations
	//  are applied in the order specified.
	//
	//  This field is similar to the one in the [`ListTimeSeries`
	//  request](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
	//  It is advisable to use the `ListTimeSeries` method when debugging this
	//  field.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricAbsence.aggregations
	Aggregations []Aggregation `json:"aggregations,omitempty"`

	// The amount of time that a time series must fail to report new
	//  data to be considered failing. The minimum value of this field
	//  is 120 seconds. Larger values that are a multiple of a
	//  minute--for example, 240 or 300 seconds--are supported.
	//  If an invalid value is given, an
	//  error will be returned. The `Duration.nanos` field is
	//  ignored.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricAbsence.duration
	Duration *string `json:"duration,omitempty"`

	// The number/percent of time series for which the comparison must hold
	//  in order for the condition to trigger. If unspecified, then the
	//  condition will trigger if the comparison is true for any of the
	//  time series that have been identified by `filter` and `aggregations`.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricAbsence.trigger
	Trigger *AlertPolicy_Condition_Trigger `json:"trigger,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold
type AlertPolicy_Condition_MetricThreshold struct {
	// Required. A
	//  [filter](https://cloud.google.com/monitoring/api/v3/filters) that
	//  identifies which time series should be compared with the threshold.
	//
	//  The filter is similar to the one that is specified in the
	//  [`ListTimeSeries`
	//  request](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list)
	//  (that call is useful to verify the time series that will be retrieved /
	//  processed). The filter must specify the metric type and the resource
	//  type. Optionally, it can specify resource labels and metric labels.
	//  This field must not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.filter
	Filter *string `json:"filter,omitempty"`

	// Specifies the alignment of data points in individual time series as
	//  well as how to combine the retrieved time series together (such as
	//  when aggregating multiple streams on each resource to a single
	//  stream for each resource or when aggregating streams across all
	//  members of a group of resources). Multiple aggregations
	//  are applied in the order specified.
	//
	//  This field is similar to the one in the [`ListTimeSeries`
	//  request](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
	//  It is advisable to use the `ListTimeSeries` method when debugging this
	//  field.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.aggregations
	Aggregations []Aggregation `json:"aggregations,omitempty"`

	// A [filter](https://cloud.google.com/monitoring/api/v3/filters) that
	//  identifies a time series that should be used as the denominator of a
	//  ratio that will be compared with the threshold. If a
	//  `denominator_filter` is specified, the time series specified by the
	//  `filter` field will be used as the numerator.
	//
	//  The filter must specify the metric type and optionally may contain
	//  restrictions on resource type, resource labels, and metric labels.
	//  This field may not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.denominator_filter
	DenominatorFilter *string `json:"denominatorFilter,omitempty"`

	// Specifies the alignment of data points in individual time series
	//  selected by `denominatorFilter` as
	//  well as how to combine the retrieved time series together (such as
	//  when aggregating multiple streams on each resource to a single
	//  stream for each resource or when aggregating streams across all
	//  members of a group of resources).
	//
	//  When computing ratios, the `aggregations` and
	//  `denominator_aggregations` fields must use the same alignment period
	//  and produce time series that have the same periodicity and labels.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.denominator_aggregations
	DenominatorAggregations []Aggregation `json:"denominatorAggregations,omitempty"`

	// When this field is present, the `MetricThreshold` condition forecasts
	//  whether the time series is predicted to violate the threshold within
	//  the `forecast_horizon`. When this field is not set, the
	//  `MetricThreshold` tests the current value of the timeseries against the
	//  threshold.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.forecast_options
	ForecastOptions *AlertPolicy_Condition_MetricThreshold_ForecastOptions `json:"forecastOptions,omitempty"`

	// The comparison to apply between the time series (indicated by `filter`
	//  and `aggregation`) and the threshold (indicated by `threshold_value`).
	//  The comparison is applied on each time series, with the time series
	//  on the left-hand side and the threshold on the right-hand side.
	//
	//  Only `COMPARISON_LT` and `COMPARISON_GT` are supported currently.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.comparison
	Comparison *string `json:"comparison,omitempty"`

	// A value against which to compare the time series.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.threshold_value
	ThresholdValue *float64 `json:"thresholdValue,omitempty"`

	// The amount of time that a time series must violate the
	//  threshold to be considered failing. Currently, only values
	//  that are a multiple of a minute--e.g., 0, 60, 120, or 300
	//  seconds--are supported. If an invalid value is given, an
	//  error will be returned. When choosing a duration, it is useful to
	//  keep in mind the frequency of the underlying time series data
	//  (which may also be affected by any alignments specified in the
	//  `aggregations` field); a good duration is long enough so that a single
	//  outlier does not generate spurious alerts, but short enough that
	//  unhealthy states are detected and alerted on quickly.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.duration
	Duration *string `json:"duration,omitempty"`

	// The number/percent of time series for which the comparison must hold
	//  in order for the condition to trigger. If unspecified, then the
	//  condition will trigger if the comparison is true for any of the
	//  time series that have been identified by `filter` and `aggregations`,
	//  or by the ratio, if `denominator_filter` and `denominator_aggregations`
	//  are specified.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.trigger
	Trigger *AlertPolicy_Condition_Trigger `json:"trigger,omitempty"`

	// A condition control that determines how metric-threshold conditions
	//  are evaluated when data stops arriving. To use this control, the value
	//  of the `duration` field must be greater than or equal to 60 seconds.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.evaluation_missing_data
	EvaluationMissingData *string `json:"evaluationMissingData,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.ForecastOptions
type AlertPolicy_Condition_MetricThreshold_ForecastOptions struct {
	// Required. The length of time into the future to forecast whether a
	//  time series will violate the threshold. If the predicted value is
	//  found to violate the threshold, and the violation is observed in all
	//  forecasts made for the configured `duration`, then the time series is
	//  considered to be failing.
	//  The forecast horizon can range from 1 hour to 60 hours.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MetricThreshold.ForecastOptions.forecast_horizon
	ForecastHorizon *string `json:"forecastHorizon,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.MonitoringQueryLanguageCondition
type AlertPolicy_Condition_MonitoringQueryLanguageCondition struct {
	// [Monitoring Query Language](https://cloud.google.com/monitoring/mql)
	//  query that outputs a boolean stream.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MonitoringQueryLanguageCondition.query
	Query *string `json:"query,omitempty"`

	// The amount of time that a time series must violate the
	//  threshold to be considered failing. Currently, only values
	//  that are a multiple of a minute--e.g., 0, 60, 120, or 300
	//  seconds--are supported. If an invalid value is given, an
	//  error will be returned. When choosing a duration, it is useful to
	//  keep in mind the frequency of the underlying time series data
	//  (which may also be affected by any alignments specified in the
	//  `aggregations` field); a good duration is long enough so that a single
	//  outlier does not generate spurious alerts, but short enough that
	//  unhealthy states are detected and alerted on quickly.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MonitoringQueryLanguageCondition.duration
	Duration *string `json:"duration,omitempty"`

	// The number/percent of time series for which the comparison must hold
	//  in order for the condition to trigger. If unspecified, then the
	//  condition will trigger if the comparison is true for any of the
	//  time series that have been identified by `filter` and `aggregations`,
	//  or by the ratio, if `denominator_filter` and `denominator_aggregations`
	//  are specified.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MonitoringQueryLanguageCondition.trigger
	Trigger *AlertPolicy_Condition_Trigger `json:"trigger,omitempty"`

	// A condition control that determines how metric-threshold conditions
	//  are evaluated when data stops arriving.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.MonitoringQueryLanguageCondition.evaluation_missing_data
	EvaluationMissingData *string `json:"evaluationMissingData,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition
type AlertPolicy_Condition_PrometheusQueryLanguageCondition struct {
	// Required. The PromQL expression to evaluate. Every evaluation cycle
	//  this expression is evaluated at the current time, and all resultant
	//  time series become pending/firing alerts. This field must not be empty.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.query
	Query *string `json:"query,omitempty"`

	// Optional. Alerts are considered firing once their PromQL expression was
	//  evaluated to be "true" for this long.
	//  Alerts whose PromQL expression was not evaluated to be "true" for
	//  long enough are considered pending.
	//  Must be a non-negative duration or missing.
	//  This field is optional. Its default value is zero.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.duration
	Duration *string `json:"duration,omitempty"`

	// Optional. How often this rule should be evaluated.
	//  Must be a positive multiple of 30 seconds or missing.
	//  This field is optional. Its default value is 30 seconds.
	//  If this PrometheusQueryLanguageCondition was generated from a
	//  Prometheus alerting rule, then this value should be taken from the
	//  enclosing rule group.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.evaluation_interval
	EvaluationInterval *string `json:"evaluationInterval,omitempty"`

	// Optional. Labels to add to or overwrite in the PromQL query result.
	//  Label names [must be
	//  valid](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels).
	//  Label values can be [templatized by using
	//  variables](https://cloud.google.com/monitoring/alerts/doc-variables#doc-vars).
	//  The only available variable names are the names of the labels in the
	//  PromQL result, including "__name__" and "value". "labels" may be empty.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The rule group name of this alert in the corresponding
	//  Prometheus configuration file.
	//
	//  Some external tools may require this field to be populated correctly
	//  in order to refer to the original Prometheus configuration file.
	//  The rule group name and the alert name are necessary to update the
	//  relevant AlertPolicies in case the definition of the rule group changes
	//  in the future.
	//
	//  This field is optional. If this field is not empty, then it must
	//  contain a valid UTF-8 string.
	//  This field may not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.rule_group
	RuleGroup *string `json:"ruleGroup,omitempty"`

	// Optional. The alerting rule name of this alert in the corresponding
	//  Prometheus configuration file.
	//
	//  Some external tools may require this field to be populated correctly
	//  in order to refer to the original Prometheus configuration file.
	//  The rule group name and the alert name are necessary to update the
	//  relevant AlertPolicies in case the definition of the rule group changes
	//  in the future.
	//
	//  This field is optional. If this field is not empty, then it must be a
	//  [valid Prometheus label
	//  name](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels).
	//  This field may not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.alert_rule
	AlertRule *string `json:"alertRule,omitempty"`

	// Optional. Whether to disable metric existence validation for this
	//  condition.
	//
	//  This allows alerting policies to be defined on metrics that do not yet
	//  exist, improving advanced customer workflows such as configuring
	//  alerting policies using Terraform.
	//
	//  Users with the `monitoring.alertPolicyViewer` role are able to see the
	//  name of the non-existent metric in the alerting policy condition.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.disable_metric_validation
	DisableMetricValidation *bool `json:"disableMetricValidation,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.SqlCondition
type AlertPolicy_Condition_SqlCondition struct {
	// Required. The Log Analytics SQL query to run, as a string.  The query
	//  must conform to the required shape. Specifically, the query must not
	//  try to filter the input by time.  A filter will automatically be
	//  applied to filter the input so that the query receives all rows
	//  received since the last time the query was run.
	//
	//  For example, the following query extracts all log entries containing an
	//  HTTP request:
	//
	//      SELECT
	//        timestamp, log_name, severity, http_request, resource, labels
	//      FROM
	//        my-project.global._Default._AllLogs
	//      WHERE
	//        http_request IS NOT NULL
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.query
	Query *string `json:"query,omitempty"`

	// Schedule the query to execute every so many minutes.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.minutes
	Minutes *AlertPolicy_Condition_SqlCondition_Minutes `json:"minutes,omitempty"`

	// Schedule the query to execute every so many hours.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.hourly
	Hourly *AlertPolicy_Condition_SqlCondition_Hourly `json:"hourly,omitempty"`

	// Schedule the query to execute every so many days.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.daily
	Daily *AlertPolicy_Condition_SqlCondition_Daily `json:"daily,omitempty"`

	// Test the row count against a threshold.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.row_count_test
	RowCountTest *AlertPolicy_Condition_SqlCondition_RowCountTest `json:"rowCountTest,omitempty"`

	// Test the boolean value in the indicated column.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.boolean_test
	BooleanTest *AlertPolicy_Condition_SqlCondition_BooleanTest `json:"booleanTest,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.BooleanTest
type AlertPolicy_Condition_SqlCondition_BooleanTest struct {
	// Required. The name of the column containing the boolean value. If the
	//  value in a row is NULL, that row is ignored.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.BooleanTest.column
	Column *string `json:"column,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Daily
type AlertPolicy_Condition_SqlCondition_Daily struct {
	// Required. The number of days between runs. Must be greater than or
	//  equal to 1 day and less than or equal to 31 days.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Daily.periodicity
	Periodicity *int32 `json:"periodicity,omitempty"`

	// Optional. The time of day (in UTC) at which the query should run. If
	//  left unspecified, the server picks an arbitrary time of day and runs
	//  the query at the same time each day.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Daily.execution_time
	ExecutionTime *TimeOfDay `json:"executionTime,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Hourly
type AlertPolicy_Condition_SqlCondition_Hourly struct {
	// Required. The number of hours between runs. Must be greater than or
	//  equal to 1 hour and less than or equal to 48 hours.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Hourly.periodicity
	Periodicity *int32 `json:"periodicity,omitempty"`

	// Optional. The number of minutes after the hour (in UTC) to run the
	//  query. Must be greater than or equal to 0 minutes and less than or
	//  equal to 59 minutes.  If left unspecified, then an arbitrary offset
	//  is used.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Hourly.minute_offset
	MinuteOffset *int32 `json:"minuteOffset,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Minutes
type AlertPolicy_Condition_SqlCondition_Minutes struct {
	// Required. Number of minutes between runs. The interval must be
	//  greater than or equal to 5 minutes and less than or equal to 1440
	//  minutes.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.Minutes.periodicity
	Periodicity *int32 `json:"periodicity,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.RowCountTest
type AlertPolicy_Condition_SqlCondition_RowCountTest struct {
	// Required. The comparison to apply between the number of rows returned
	//  by the query and the threshold.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.RowCountTest.comparison
	Comparison *string `json:"comparison,omitempty"`

	// Required. The value against which to compare the row count.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.SqlCondition.RowCountTest.threshold
	Threshold *int64 `json:"threshold,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.Trigger
type AlertPolicy_Condition_Trigger struct {
	// The absolute number of time series that must fail
	//  the predicate for the condition to be triggered.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.Trigger.count
	Count *int32 `json:"count,omitempty"`

	// The percentage of time series that must fail the
	//  predicate for the condition to be triggered.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.Trigger.percent
	Percent *float64 `json:"percent,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Documentation
type AlertPolicy_Documentation struct {
	// The body of the documentation, interpreted according to `mime_type`.
	//  The content may not exceed 8,192 Unicode characters and may not exceed
	//  more than 10,240 bytes when encoded in UTF-8 format, whichever is
	//  smaller. This text can be [templatized by using
	//  variables](https://cloud.google.com/monitoring/alerts/doc-variables#doc-vars).
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.content
	Content *string `json:"content,omitempty"`

	// The format of the `content` field. Presently, only the value
	//  `"text/markdown"` is supported. See
	//  [Markdown](https://en.wikipedia.org/wiki/Markdown) for more information.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Optional. The subject line of the notification. The subject line may not
	//  exceed 10,240 bytes. In notifications generated by this policy, the
	//  contents of the subject line after variable expansion will be truncated
	//  to 255 bytes or shorter at the latest UTF-8 character boundary. The
	//  255-byte limit is recommended by [this
	//  thread](https://stackoverflow.com/questions/1592291/what-is-the-email-subject-length-limit).
	//  It is both the limit imposed by some third-party ticketing products and
	//  it is common to define textual fields in databases as VARCHAR(255).
	//
	//  The contents of the subject line can be [templatized by using
	//  variables](https://cloud.google.com/monitoring/alerts/doc-variables#doc-vars).
	//  If this field is missing or empty, a default subject line will be
	//  generated.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.subject
	Subject *string `json:"subject,omitempty"`

	// Optional. Links to content such as playbooks, repositories, and other
	//  resources. This field can contain up to 3 entries.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.links
	Links []AlertPolicy_Documentation_Link `json:"links,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Documentation.Link
type AlertPolicy_Documentation_Link struct {
	// A short display name for the link. The display name must not be empty
	//  or exceed 63 characters. Example: "playbook".
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.Link.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The url of a webpage.
	//  A url can be templatized by using variables
	//  in the path or the query parameters. The total length of a URL should
	//  not exceed 2083 characters before and after variable expansion.
	//  Example: "https://my_domain.com/playbook?name=${resource.name}"
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.Link.url
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.monitoring.v3.MutationRecord
type MutationRecord struct {
	// When the change occurred.
	// +kcc:proto:field=google.monitoring.v3.MutationRecord.mutate_time
	MutateTime *string `json:"mutateTime,omitempty"`

	// The email address of the user making the change.
	// +kcc:proto:field=google.monitoring.v3.MutationRecord.mutated_by
	MutatedBy *string `json:"mutatedBy,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.TimeOfDay
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	//  to allow the value "24:00:00" for scenarios like business closing time.
	// +kcc:proto:field=google.type.TimeOfDay.hours
	Hours *int32 `json:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.TimeOfDay.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	//  allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.TimeOfDay.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kcc:proto:field=google.type.TimeOfDay.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}
