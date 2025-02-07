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


// +kcc:proto=google.analytics.admin.v1beta.CustomDimension
type CustomDimension struct {

	// Required. Immutable. Tagging parameter name for this custom dimension.
	//
	//  If this is a user-scoped dimension, then this is the user property name.
	//  If this is an event-scoped dimension, then this is the event parameter
	//  name.
	//
	//  If this is an item-scoped dimension, then this is the parameter
	//  name found in the eCommerce items array.
	//
	//  May only contain alphanumeric and underscore characters, starting with a
	//  letter. Max length of 24 characters for user-scoped dimensions, 40
	//  characters for event-scoped dimensions.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomDimension.parameter_name
	ParameterName *string `json:"parameterName,omitempty"`

	// Required. Display name for this custom dimension as shown in the Analytics
	//  UI. Max length of 82 characters, alphanumeric plus space and underscore
	//  starting with a letter. Legacy system-generated display names may contain
	//  square brackets, but updates to this field will never permit square
	//  brackets.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomDimension.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for this custom dimension. Max length of 150
	//  characters.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomDimension.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The scope of this dimension.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomDimension.scope
	Scope *string `json:"scope,omitempty"`

	// Optional. If set to true, sets this dimension as NPA and excludes it from
	//  ads personalization.
	//
	//  This is currently only supported by user-scoped custom dimensions.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomDimension.disallow_ads_personalization
	DisallowAdsPersonalization *bool `json:"disallowAdsPersonalization,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.CustomDimension
type CustomDimensionObservedState struct {
	// Output only. Resource name for this CustomDimension resource.
	//  Format: properties/{property}/customDimensions/{customDimension}
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomDimension.name
	Name *string `json:"name,omitempty"`
}
