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

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaConfig
type QuotaConfig struct {
	// Required. The preferred value. Must be greater than or equal to -1. If set
	//  to -1, it means the value is "unlimited".
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.preferred_value
	PreferredValue *int64 `json:"preferredValue,omitempty"`

	// Optional. The annotations map for clients to store small amounts of
	//  arbitrary data. Do not put PII or other sensitive information here. See
	//  https://google.aip.dev/128#annotations
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaPreference
type QuotaPreference struct {
	// Required except in the CREATE requests.
	//  The resource name of the quota preference.
	//  The ID component following "locations/" must be "global".
	//  Example:
	//  `projects/123/locations/global/quotaPreferences/my-config-for-us-east1`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.name
	Name *string `json:"name,omitempty"`

	// Immutable. The dimensions that this quota preference applies to. The key of
	//  the map entry is the name of a dimension, such as "region", "zone",
	//  "network_id", and the value of the map entry is the dimension value.
	//
	//  If a dimension is missing from the map of dimensions, the quota preference
	//  applies to all the dimension values except for those that have other quota
	//  preferences configured for the specific value.
	//
	//  NOTE: QuotaPreferences can only be applied across all values of "user" and
	//  "resource" dimension. Do not set values for "user" or "resource" in the
	//  dimension map.
	//
	//  Example: {"provider", "Foo Inc"} where "provider" is a service specific
	//  dimension.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.dimensions
	Dimensions map[string]string `json:"dimensions,omitempty"`

	// Required. Preferred quota configuration.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.quota_config
	QuotaConfig *QuotaConfig `json:"quotaConfig,omitempty"`

	// Optional. The current etag of the quota preference. If an etag is provided
	//  on update and does not match the current server's etag of the quota
	//  preference, the request will be blocked and an ABORTED error will be
	//  returned. See https://google.aip.dev/134#etags for more details on etags.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.etag
	Etag *string `json:"etag,omitempty"`

	// Required. The name of the service to which the quota preference is applied.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.service
	Service *string `json:"service,omitempty"`

	// Required. The id of the quota to which the quota preference is applied. A
	//  quota name is unique in the service. Example: `CpusPerProjectPerRegion`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.quota_id
	QuotaID *string `json:"quotaID,omitempty"`

	// The reason / justification for this quota preference.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.justification
	Justification *string `json:"justification,omitempty"`

	// Input only. An email address that can be used to contact the the user, in
	//  case Google Cloud needs more information to make a decision before
	//  additional quota can be granted.
	//
	//  When requesting a quota increase, the email address is required.
	//  When requesting a quota decrease, the email address is optional.
	//  For example, the email address is optional when the
	//  `QuotaConfig.preferred_value` is smaller than the
	//  `QuotaDetails.reset_value`.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.contact_email
	ContactEmail *string `json:"contactEmail,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaConfig
type QuotaConfigObservedState struct {
	// Output only. Optional details about the state of this quota preference.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.state_detail
	StateDetail *string `json:"stateDetail,omitempty"`

	// Output only. Granted quota value.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.granted_value
	GrantedValue *int64 `json:"grantedValue,omitempty"`

	// Output only. The trace id that the Google Cloud uses to provision the
	//  requested quota. This trace id may be used by the client to contact Cloud
	//  support to track the state of a quota preference request. The trace id is
	//  only produced for increase requests and is unique for each request. The
	//  quota decrease requests do not have a trace id.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.trace_id
	TraceID *string `json:"traceID,omitempty"`

	// Output only. The origin of the quota preference request.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.request_origin
	RequestOrigin *string `json:"requestOrigin,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaPreference
type QuotaPreferenceObservedState struct {
	// Required. Preferred quota configuration.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.quota_config
	QuotaConfig *QuotaConfigObservedState `json:"quotaConfig,omitempty"`

	// Output only. Create time stamp
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Is the quota preference pending Google Cloud approval and
	//  fulfillment.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`
}
