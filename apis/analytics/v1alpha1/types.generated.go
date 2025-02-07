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


// +kcc:proto=google.analytics.data.v1beta.AudienceDimension
type AudienceDimension struct {
	// Optional. The API name of the dimension. See the [API
	//  Dimensions](https://developers.google.com/analytics/devguides/reporting/data/v1/audience-list-api-schema#dimensions)
	//  for the list of dimension names.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceDimension.dimension_name
	DimensionName *string `json:"dimensionName,omitempty"`
}

// +kcc:proto=google.analytics.data.v1beta.AudienceExport
type AudienceExport struct {

	// Required. The audience resource name. This resource name identifies the
	//  audience being listed and is shared between the Analytics Data & Admin
	//  APIs.
	//
	//  Format: `properties/{property}/audiences/{audience}`
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.audience
	Audience *string `json:"audience,omitempty"`

	// Required. The dimensions requested and displayed in the query response.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.dimensions
	Dimensions []AudienceDimension `json:"dimensions,omitempty"`
}

// +kcc:proto=google.analytics.data.v1beta.AudienceExport
type AudienceExportObservedState struct {
	// Output only. Identifier. The audience export resource name assigned during
	//  creation. This resource name identifies this `AudienceExport`.
	//
	//  Format: `properties/{property}/audienceExports/{audience_export}`
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.name
	Name *string `json:"name,omitempty"`

	// Output only. The descriptive display name for this audience. For example,
	//  "Purchasers".
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.audience_display_name
	AudienceDisplayName *string `json:"audienceDisplayName,omitempty"`

	// Output only. The current state for this AudienceExport.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.state
	State *string `json:"state,omitempty"`

	// Output only. The time when CreateAudienceExport was called and the
	//  AudienceExport began the `CREATING` state.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.begin_creating_time
	BeginCreatingTime *string `json:"beginCreatingTime,omitempty"`

	// Output only. The total quota tokens charged during creation of the
	//  AudienceExport. Because this token count is based on activity from the
	//  `CREATING` state, this tokens charged will be fixed once an AudienceExport
	//  enters the `ACTIVE` or `FAILED` states.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.creation_quota_tokens_charged
	CreationQuotaTokensCharged *int32 `json:"creationQuotaTokensCharged,omitempty"`

	// Output only. The total number of rows in the AudienceExport result.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.row_count
	RowCount *int32 `json:"rowCount,omitempty"`

	// Output only. Error message is populated when an audience export fails
	//  during creation. A common reason for such a failure is quota exhaustion.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Output only. The percentage completed for this audience export ranging
	//  between 0 to 100.
	// +kcc:proto:field=google.analytics.data.v1beta.AudienceExport.percentage_completed
	PercentageCompleted *float64 `json:"percentageCompleted,omitempty"`
}
