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


// +kcc:proto=google.api.cloudquotas.v1beta.DimensionsInfo
type DimensionsInfo struct {
	// The map of dimensions for this dimensions info. The key of a map entry
	//  is "region", "zone" or the name of a service specific dimension, and the
	//  value of a map entry is the value of the dimension.  If a dimension does
	//  not appear in the map of dimensions, the dimensions info applies to all
	//  the dimension values except for those that have another DimenisonInfo
	//  instance configured for the specific value.
	//  Example: {"provider" : "Foo Inc"} where "provider" is a service specific
	//  dimension of a quota.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.DimensionsInfo.dimensions
	Dimensions map[string]string `json:"dimensions,omitempty"`

	// Quota details for the specified dimensions.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.DimensionsInfo.details
	Details *QuotaDetails `json:"details,omitempty"`

	// The applicable regions or zones of this dimensions info. The field will be
	//  set to ['global'] for quotas that are not per region or per zone.
	//  Otherwise, it will be set to the list of locations this dimension info is
	//  applicable to.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.DimensionsInfo.applicable_locations
	ApplicableLocations []string `json:"applicableLocations,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaDetails
type QuotaDetails struct {
	// The value currently in effect and being enforced.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaDetails.value
	Value *int64 `json:"value,omitempty"`

	// Rollout information of this quota.
	//  This field is present only if the effective limit will change due to the
	//  ongoing rollout of the service config.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaDetails.rollout_info
	RolloutInfo *RolloutInfo `json:"rolloutInfo,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaIncreaseEligibility
type QuotaIncreaseEligibility struct {
	// Whether a higher quota value can be requested for the quota.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaIncreaseEligibility.is_eligible
	IsEligible *bool `json:"isEligible,omitempty"`

	// The reason of why it is ineligible to request increased value of the quota.
	//  If the is_eligible field is true, it defaults to
	//  INELIGIBILITY_REASON_UNSPECIFIED.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaIncreaseEligibility.ineligibility_reason
	IneligibilityReason *string `json:"ineligibilityReason,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaInfo
type QuotaInfo struct {
	// Resource name of this QuotaInfo.
	//  The ID component following "locations/" must be "global".
	//  Example:
	//  `projects/123/locations/global/services/compute.googleapis.com/quotaInfos/CpusPerProjectPerRegion`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.name
	Name *string `json:"name,omitempty"`

	// The id of the quota, which is unquie within the service.
	//  Example: `CpusPerProjectPerRegion`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.quota_id
	QuotaID *string `json:"quotaID,omitempty"`

	// The metric of the quota. It specifies the resources consumption the quota
	//  is defined for.
	//  Example: `compute.googleapis.com/cpus`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.metric
	Metric *string `json:"metric,omitempty"`

	// The name of the service in which the quota is defined.
	//  Example: `compute.googleapis.com`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.service
	Service *string `json:"service,omitempty"`

	// Whether this is a precise quota. A precise quota is tracked with absolute
	//  precision. In contrast, an imprecise quota is not tracked with precision.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.is_precise
	IsPrecise *bool `json:"isPrecise,omitempty"`

	// The reset time interval for the quota. Refresh interval applies to rate
	//  quota only.
	//  Example: "minute" for per minute, "day" for per day, or "10 seconds" for
	//  every 10 seconds.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.refresh_interval
	RefreshInterval *string `json:"refreshInterval,omitempty"`

	// The container type of the QuotaInfo.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.container_type
	ContainerType *string `json:"containerType,omitempty"`

	// The dimensions the quota is defined on.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.dimensions
	Dimensions []string `json:"dimensions,omitempty"`

	// The display name of the quota metric
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.metric_display_name
	MetricDisplayName *string `json:"metricDisplayName,omitempty"`

	// The display name of the quota.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.quota_display_name
	QuotaDisplayName *string `json:"quotaDisplayName,omitempty"`

	// The unit in which the metric value is reported, e.g., "MByte".
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.metric_unit
	MetricUnit *string `json:"metricUnit,omitempty"`

	// Whether it is eligible to request a higher quota value for this quota.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.quota_increase_eligibility
	QuotaIncreaseEligibility *QuotaIncreaseEligibility `json:"quotaIncreaseEligibility,omitempty"`

	// Whether the quota value is fixed or adjustable
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.is_fixed
	IsFixed *bool `json:"isFixed,omitempty"`

	// The collection of dimensions info ordered by their dimensions from more
	//  specific ones to less specific ones.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.dimensions_infos
	DimensionsInfos []DimensionsInfo `json:"dimensionsInfos,omitempty"`

	// Whether the quota is a concurrent quota. Concurrent quotas are enforced
	//  on the total number of concurrent operations in flight at any given time.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.is_concurrent
	IsConcurrent *bool `json:"isConcurrent,omitempty"`

	// URI to the page where users can request more quota for the cloud
	//  service—for example,
	//  https://console.cloud.google.com/iam-admin/quotas.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaInfo.service_request_quota_uri
	ServiceRequestQuotaURI *string `json:"serviceRequestQuotaURI,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.RolloutInfo
type RolloutInfo struct {
	// Whether there is an ongoing rollout for a quota or not.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.RolloutInfo.ongoing_rollout
	OngoingRollout *bool `json:"ongoingRollout,omitempty"`
}
