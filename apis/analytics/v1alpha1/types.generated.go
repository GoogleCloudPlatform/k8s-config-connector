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


// +kcc:proto=google.analytics.admin.v1beta.CustomMetric
type CustomMetric struct {

	// Required. Immutable. Tagging name for this custom metric.
	//
	//  If this is an event-scoped metric, then this is the event parameter
	//  name.
	//
	//  May only contain alphanumeric and underscore charactes, starting with a
	//  letter. Max length of 40 characters for event-scoped metrics.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.parameter_name
	ParameterName *string `json:"parameterName,omitempty"`

	// Required. Display name for this custom metric as shown in the Analytics UI.
	//  Max length of 82 characters, alphanumeric plus space and underscore
	//  starting with a letter. Legacy system-generated display names may contain
	//  square brackets, but updates to this field will never permit square
	//  brackets.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for this custom dimension.
	//  Max length of 150 characters.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.description
	Description *string `json:"description,omitempty"`

	// Required. The type for the custom metric's value.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.measurement_unit
	MeasurementUnit *string `json:"measurementUnit,omitempty"`

	// Required. Immutable. The scope of this custom metric.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.scope
	Scope *string `json:"scope,omitempty"`

	// Optional. Types of restricted data that this metric may contain. Required
	//  for metrics with CURRENCY measurement unit. Must be empty for metrics with
	//  a non-CURRENCY measurement unit.
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.restricted_metric_type
	RestrictedMetricType []string `json:"restrictedMetricType,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.CustomMetric
type CustomMetricObservedState struct {
	// Output only. Resource name for this CustomMetric resource.
	//  Format: properties/{property}/customMetrics/{customMetric}
	// +kcc:proto:field=google.analytics.admin.v1beta.CustomMetric.name
	Name *string `json:"name,omitempty"`
}
