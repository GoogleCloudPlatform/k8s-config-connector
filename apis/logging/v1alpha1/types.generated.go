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

// +kcc:proto=google.api.Distribution.BucketOptions
type Distribution_BucketOptions struct {
	// The linear bucket.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.linear_buckets
	LinearBuckets *Distribution_BucketOptions_Linear `json:"linearBuckets,omitempty"`

	// The exponential buckets.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.exponential_buckets
	ExponentialBuckets *Distribution_BucketOptions_Exponential `json:"exponentialBuckets,omitempty"`

	// The explicit buckets.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.explicit_buckets
	ExplicitBuckets *Distribution_BucketOptions_Explicit `json:"explicitBuckets,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Explicit
type Distribution_BucketOptions_Explicit struct {
	// The values must be monotonically increasing.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Explicit.bounds
	Bounds []float64 `json:"bounds,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Exponential
type Distribution_BucketOptions_Exponential struct {
	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.num_finite_buckets
	NumFiniteBuckets *int32 `json:"numFiniteBuckets,omitempty"`

	// Must be greater than 1.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.growth_factor
	GrowthFactor *float64 `json:"growthFactor,omitempty"`

	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Exponential.scale
	Scale *float64 `json:"scale,omitempty"`
}

// +kcc:proto=google.api.Distribution.BucketOptions.Linear
type Distribution_BucketOptions_Linear struct {
	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.num_finite_buckets
	NumFiniteBuckets *int32 `json:"numFiniteBuckets,omitempty"`

	// Must be greater than 0.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.width
	Width *float64 `json:"width,omitempty"`

	// Lower bound of the first bucket.
	// +kcc:proto:field=google.api.Distribution.BucketOptions.Linear.offset
	Offset *float64 `json:"offset,omitempty"`
}

// +kcc:proto=google.api.LabelDescriptor
type LabelDescriptor struct {
	// The label key.
	// +kcc:proto:field=google.api.LabelDescriptor.key
	Key *string `json:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// +kcc:proto:field=google.api.LabelDescriptor.value_type
	ValueType *string `json:"valueType,omitempty"`

	// A human-readable description for the label.
	// +kcc:proto:field=google.api.LabelDescriptor.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.api.MetricDescriptor
type MetricDescriptor struct {
	// The resource name of the metric descriptor.
	// +kcc:proto:field=google.api.MetricDescriptor.name
	Name *string `json:"name,omitempty"`

	// The metric type, including its DNS name prefix. The type is not
	//  URL-encoded. All user-defined metric types have the DNS name
	//  `custom.googleapis.com` or `external.googleapis.com`. Metric types should
	//  use a natural hierarchical grouping. For example:
	//
	//      "custom.googleapis.com/invoice/paid/amount"
	//      "external.googleapis.com/prometheus/up"
	//      "appengine.googleapis.com/http/server/response_latencies"
	// +kcc:proto:field=google.api.MetricDescriptor.type
	Type *string `json:"type,omitempty"`

	// The set of labels that can be used to describe a specific
	//  instance of this metric type. For example, the
	//  `appengine.googleapis.com/http/server/response_latencies` metric
	//  type has a label for the HTTP response code, `response_code`, so
	//  you can look at latencies for successful responses or just
	//  for responses that failed.
	// +kcc:proto:field=google.api.MetricDescriptor.labels
	Labels []LabelDescriptor `json:"labels,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc.
	//  Some combinations of `metric_kind` and `value_type` might not be supported.
	// +kcc:proto:field=google.api.MetricDescriptor.metric_kind
	MetricKind *string `json:"metricKind,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	//  Some combinations of `metric_kind` and `value_type` might not be supported.
	// +kcc:proto:field=google.api.MetricDescriptor.value_type
	ValueType *string `json:"valueType,omitempty"`

	// The units in which the metric value is reported. It is only applicable
	//  if the `value_type` is `INT64`, `DOUBLE`, or `DISTRIBUTION`. The `unit`
	//  defines the representation of the stored metric values.
	//
	//  Different systems might scale the values to be more easily displayed (so a
	//  value of `0.02kBy` _might_ be displayed as `20By`, and a value of
	//  `3523kBy` _might_ be displayed as `3.5MBy`). However, if the `unit` is
	//  `kBy`, then the value of the metric is always in thousands of bytes, no
	//  matter how it might be displayed.
	//
	//  If you want a custom metric to record the exact number of CPU-seconds used
	//  by a job, you can create an `INT64 CUMULATIVE` metric whose `unit` is
	//  `s{CPU}` (or equivalently `1s{CPU}` or just `s`). If the job uses 12,005
	//  CPU-seconds, then the value is written as `12005`.
	//
	//  Alternatively, if you want a custom metric to record data in a more
	//  granular way, you can create a `DOUBLE CUMULATIVE` metric whose `unit` is
	//  `ks{CPU}`, and then write the value `12.005` (which is `12005/1000`),
	//  or use `Kis{CPU}` and write `11.723` (which is `12005/1024`).
	//
	//  The supported units are a subset of [The Unified Code for Units of
	//  Measure](https://unitsofmeasure.org/ucum.html) standard:
	//
	//  **Basic units (UNIT)**
	//
	//  * `bit`   bit
	//  * `By`    byte
	//  * `s`     second
	//  * `min`   minute
	//  * `h`     hour
	//  * `d`     day
	//  * `1`     dimensionless
	//
	//  **Prefixes (PREFIX)**
	//
	//  * `k`     kilo    (10^3)
	//  * `M`     mega    (10^6)
	//  * `G`     giga    (10^9)
	//  * `T`     tera    (10^12)
	//  * `P`     peta    (10^15)
	//  * `E`     exa     (10^18)
	//  * `Z`     zetta   (10^21)
	//  * `Y`     yotta   (10^24)
	//
	//  * `m`     milli   (10^-3)
	//  * `u`     micro   (10^-6)
	//  * `n`     nano    (10^-9)
	//  * `p`     pico    (10^-12)
	//  * `f`     femto   (10^-15)
	//  * `a`     atto    (10^-18)
	//  * `z`     zepto   (10^-21)
	//  * `y`     yocto   (10^-24)
	//
	//  * `Ki`    kibi    (2^10)
	//  * `Mi`    mebi    (2^20)
	//  * `Gi`    gibi    (2^30)
	//  * `Ti`    tebi    (2^40)
	//  * `Pi`    pebi    (2^50)
	//
	//  **Grammar**
	//
	//  The grammar also includes these connectors:
	//
	//  * `/`    division or ratio (as an infix operator). For examples,
	//           `kBy/{email}` or `MiBy/10ms` (although you should almost never
	//           have `/s` in a metric `unit`; rates should always be computed at
	//           query time from the underlying cumulative or delta value).
	//  * `.`    multiplication or composition (as an infix operator). For
	//           examples, `GBy.d` or `k{watt}.h`.
	//
	//  The grammar for a unit is as follows:
	//
	//      Expression = Component { "." Component } { "/" Component } ;
	//
	//      Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ]
	//                | Annotation
	//                | "1"
	//                ;
	//
	//      Annotation = "{" NAME "}" ;
	//
	//  Notes:
	//
	//  * `Annotation` is just a comment if it follows a `UNIT`. If the annotation
	//     is used alone, then the unit is equivalent to `1`. For examples,
	//     `{request}/s == 1/s`, `By{transmitted}/s == By/s`.
	//  * `NAME` is a sequence of non-blank printable ASCII characters not
	//     containing `{` or `}`.
	//  * `1` represents a unitary [dimensionless
	//     unit](https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such
	//     as in `1/s`. It is typically used when none of the basic units are
	//     appropriate. For example, "new users per day" can be represented as
	//     `1/d` or `{new-users}/d` (and a metric value `5` would mean "5 new
	//     users). Alternatively, "thousands of page views per day" would be
	//     represented as `1000/d` or `k1/d` or `k{page_views}/d` (and a metric
	//     value of `5.3` would mean "5300 page views per day").
	//  * `%` represents dimensionless value of 1/100, and annotates values giving
	//     a percentage (so the metric values are typically in the range of 0..100,
	//     and a metric value `3` means "3 percent").
	//  * `10^2.%` indicates a metric contains a ratio, typically in the range
	//     0..1, that will be multiplied by 100 and displayed as a percentage
	//     (so a metric value `0.03` means "3 percent").
	// +kcc:proto:field=google.api.MetricDescriptor.unit
	Unit *string `json:"unit,omitempty"`

	// A detailed description of the metric, which can be used in documentation.
	// +kcc:proto:field=google.api.MetricDescriptor.description
	Description *string `json:"description,omitempty"`

	// A concise name for the metric, which can be displayed in user interfaces.
	//  Use sentence case without an ending period, for example "Request count".
	//  This field is optional but it is recommended to be set for any metrics
	//  associated with user-visible concepts, such as Quota.
	// +kcc:proto:field=google.api.MetricDescriptor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Metadata which can be used to guide usage of the metric.
	// +kcc:proto:field=google.api.MetricDescriptor.metadata
	Metadata *MetricDescriptor_MetricDescriptorMetadata `json:"metadata,omitempty"`

	// Optional. The launch stage of the metric definition.
	// +kcc:proto:field=google.api.MetricDescriptor.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// Read-only. If present, then a [time
	//  series][google.monitoring.v3.TimeSeries], which is identified partially by
	//  a metric type and a
	//  [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor], that
	//  is associated with this metric type can only be associated with one of the
	//  monitored resource types listed here.
	// +kcc:proto:field=google.api.MetricDescriptor.monitored_resource_types
	MonitoredResourceTypes []string `json:"monitoredResourceTypes,omitempty"`
}

// +kcc:proto=google.api.MetricDescriptor.MetricDescriptorMetadata
type MetricDescriptor_MetricDescriptorMetadata struct {
	// Deprecated. Must use the
	//  [MetricDescriptor.launch_stage][google.api.MetricDescriptor.launch_stage]
	//  instead.
	// +kcc:proto:field=google.api.MetricDescriptor.MetricDescriptorMetadata.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// The sampling period of metric data points. For metrics which are written
	//  periodically, consecutive data points are stored at this time interval,
	//  excluding data loss due to errors. Metrics with a higher granularity have
	//  a smaller sampling period.
	// +kcc:proto:field=google.api.MetricDescriptor.MetricDescriptorMetadata.sample_period
	SamplePeriod *string `json:"samplePeriod,omitempty"`

	// The delay of data points caused by ingestion. Data points older than this
	//  age are guaranteed to be ingested and available to be read, excluding
	//  data loss due to errors.
	// +kcc:proto:field=google.api.MetricDescriptor.MetricDescriptorMetadata.ingest_delay
	IngestDelay *string `json:"ingestDelay,omitempty"`

	// The scope of the timeseries data of the metric.
	// +kcc:proto:field=google.api.MetricDescriptor.MetricDescriptorMetadata.time_series_resource_hierarchy_level
	TimeSeriesResourceHierarchyLevel []string `json:"timeSeriesResourceHierarchyLevel,omitempty"`
}
