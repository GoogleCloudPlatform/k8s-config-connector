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


// +kcc:proto=google.cloud.visionai.v1.FacetProperty
type FacetProperty struct {
	// Fixed range facet bucket config.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.fixed_range_bucket_spec
	FixedRangeBucketSpec *FacetProperty_FixedRangeBucketSpec `json:"fixedRangeBucketSpec,omitempty"`

	// Custom range facet bucket config.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.custom_range_bucket_spec
	CustomRangeBucketSpec *FacetProperty_CustomRangeBucketSpec `json:"customRangeBucketSpec,omitempty"`

	// Datetime range facet bucket config.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.datetime_bucket_spec
	DatetimeBucketSpec *FacetProperty_DateTimeBucketSpec `json:"datetimeBucketSpec,omitempty"`

	// Name of the facets, which are the dimensions users want to use to refine
	//  search results. `mapped_fields` will match UserSpecifiedDataSchema keys.
	//
	//  For example, user can add a bunch of UGAs with the same key, such as
	//  player:adam, player:bob, player:charles. When multiple mapped_fields are
	//  specified, will merge their value together as final facet value. E.g.
	//  home_team: a, home_team:b, away_team:a, away_team:c, when facet_field =
	//  [home_team, away_team], facet_value will be [a, b, c].
	//
	//  UNLESS this is a 1:1 facet dimension (mapped_fields.size() == 1) AND the
	//  mapped_field equals the parent SearchConfig.name, the parent must
	//  also contain a SearchCriteriaProperty that maps to the same fields.
	//  mapped_fields must not be empty.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.mapped_fields
	MappedFields []string `json:"mappedFields,omitempty"`

	// Display name of the facet. To be used by UI for facet rendering.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Maximum number of unique bucket to return for one facet. Bucket number can
	//  be large for high-cardinality facet such as "player". We only return top-n
	//  most related ones to user. If it's <= 0, the server will decide the
	//  appropriate result_size.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.result_size
	ResultSize *int64 `json:"resultSize,omitempty"`

	// Facet bucket type e.g. value, range.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.bucket_type
	BucketType *string `json:"bucketType,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.FacetProperty.CustomRangeBucketSpec
type FacetProperty_CustomRangeBucketSpec struct {
	// Currently, only integer type is supported for this field.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.CustomRangeBucketSpec.endpoints
	Endpoints []FacetValue `json:"endpoints,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.FacetProperty.DateTimeBucketSpec
type FacetProperty_DateTimeBucketSpec struct {
	// Granularity of date type facet.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.DateTimeBucketSpec.granularity
	Granularity *string `json:"granularity,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.FacetProperty.FixedRangeBucketSpec
type FacetProperty_FixedRangeBucketSpec struct {
	// Lower bound of the bucket. NOTE: Only integer type is currently supported
	//  for this field.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.FixedRangeBucketSpec.bucket_start
	BucketStart *FacetValue `json:"bucketStart,omitempty"`

	// Bucket granularity. NOTE: Only integer type is currently supported for
	//  this field.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.FixedRangeBucketSpec.bucket_granularity
	BucketGranularity *FacetValue `json:"bucketGranularity,omitempty"`

	// Total number of buckets.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetProperty.FixedRangeBucketSpec.bucket_count
	BucketCount *int32 `json:"bucketCount,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.FacetValue
type FacetValue struct {
	// String type value.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetValue.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Integer type value.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetValue.integer_value
	IntegerValue *int64 `json:"integerValue,omitempty"`

	// Datetime type value.
	// +kcc:proto:field=google.cloud.visionai.v1.FacetValue.datetime_value
	DatetimeValue *DateTime `json:"datetimeValue,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.SearchConfig
type SearchConfig struct {
	// Resource name of the search configuration.
	//  For CustomSearchCriteria, search_config would be the search
	//  operator name. For Facets, search_config would be the facet
	//  dimension name.
	//  Format:
	//  `projects/{project_number}/locations/{location}/corpora/{corpus}/searchConfigs/{search_config}`
	// +kcc:proto:field=google.cloud.visionai.v1.SearchConfig.name
	Name *string `json:"name,omitempty"`

	// Establishes a FacetDimension and associated specifications.
	// +kcc:proto:field=google.cloud.visionai.v1.SearchConfig.facet_property
	FacetProperty *FacetProperty `json:"facetProperty,omitempty"`

	// Creates a mapping between a custom SearchCriteria and one or more UGA keys.
	// +kcc:proto:field=google.cloud.visionai.v1.SearchConfig.search_criteria_property
	SearchCriteriaProperty *SearchCriteriaProperty `json:"searchCriteriaProperty,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.SearchCriteriaProperty
type SearchCriteriaProperty struct {
	// Each mapped_field corresponds to a UGA key. To understand how this property
	//  works, take the following example. In the SearchConfig table, the
	//  user adds this entry:
	//    search_config {
	//      name: "person"
	//      search_criteria_property {
	//        mapped_fields: "player"
	//        mapped_fields: "coach"
	//      }
	//    }
	//
	//  Now, when a user issues a query like:
	//    criteria {
	//      field: "person"
	//      text_array {
	//        txt_values: "Tom Brady"
	//        txt_values: "Bill Belichick"
	//      }
	//    }
	//
	//  MWH search will return search documents where (player=Tom Brady ||
	//  coach=Tom Brady || player=Bill Belichick || coach=Bill Belichick).
	// +kcc:proto:field=google.cloud.visionai.v1.SearchCriteriaProperty.mapped_fields
	MappedFields []string `json:"mappedFields,omitempty"`
}

// +kcc:proto=google.type.DateTime
type DateTime struct {
	// Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
	//  datetime without a year.
	// +kcc:proto:field=google.type.DateTime.year
	Year *int32 `json:"year,omitempty"`

	// Required. Month of year. Must be from 1 to 12.
	// +kcc:proto:field=google.type.DateTime.month
	Month *int32 `json:"month,omitempty"`

	// Required. Day of month. Must be from 1 to 31 and valid for the year and
	//  month.
	// +kcc:proto:field=google.type.DateTime.day
	Day *int32 `json:"day,omitempty"`

	// Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
	//  may choose to allow the value "24:00:00" for scenarios like business
	//  closing time.
	// +kcc:proto:field=google.type.DateTime.hours
	Hours *int32 `json:"hours,omitempty"`

	// Required. Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.DateTime.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
	//  API may allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.DateTime.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Required. Fractions of seconds in nanoseconds. Must be from 0 to
	//  999,999,999.
	// +kcc:proto:field=google.type.DateTime.nanos
	Nanos *int32 `json:"nanos,omitempty"`

	// UTC offset. Must be whole seconds, between -18 hours and +18 hours.
	//  For example, a UTC offset of -4:00 would be represented as
	//  { seconds: -14400 }.
	// +kcc:proto:field=google.type.DateTime.utc_offset
	UtcOffset *string `json:"utcOffset,omitempty"`

	// Time zone.
	// +kcc:proto:field=google.type.DateTime.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

// +kcc:proto=google.type.TimeZone
type TimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	// +kcc:proto:field=google.type.TimeZone.id
	ID *string `json:"id,omitempty"`

	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	// +kcc:proto:field=google.type.TimeZone.version
	Version *string `json:"version,omitempty"`
}
