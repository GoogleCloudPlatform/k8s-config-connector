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


// +kcc:proto=google.analytics.data.v1beta.ComparisonMetadata
type ComparisonMetadata struct {
	// This comparison's resource name. Useable in [Comparison](#Comparison)'s
	//  `comparison` field. For example, 'comparisons/1234'.
	// +kcc:proto:field=google.analytics.data.v1beta.ComparisonMetadata.api_name
	ApiName *string `json:"apiName,omitempty"`

	// This comparison's name within the Google Analytics user interface.
	// +kcc:proto:field=google.analytics.data.v1beta.ComparisonMetadata.ui_name
	UiName *string `json:"uiName,omitempty"`

	// This comparison's description.
	// +kcc:proto:field=google.analytics.data.v1beta.ComparisonMetadata.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.analytics.data.v1beta.DimensionMetadata
type DimensionMetadata struct {
	// This dimension's name. Useable in [Dimension](#Dimension)'s `name`. For
	//  example, `eventName`.
	// +kcc:proto:field=google.analytics.data.v1beta.DimensionMetadata.api_name
	ApiName *string `json:"apiName,omitempty"`

	// This dimension's name within the Google Analytics user interface. For
	//  example, `Event name`.
	// +kcc:proto:field=google.analytics.data.v1beta.DimensionMetadata.ui_name
	UiName *string `json:"uiName,omitempty"`

	// Description of how this dimension is used and calculated.
	// +kcc:proto:field=google.analytics.data.v1beta.DimensionMetadata.description
	Description *string `json:"description,omitempty"`

	// Still usable but deprecated names for this dimension. If populated, this
	//  dimension is available by either `apiName` or one of `deprecatedApiNames`
	//  for a period of time. After the deprecation period, the dimension will be
	//  available only by `apiName`.
	// +kcc:proto:field=google.analytics.data.v1beta.DimensionMetadata.deprecated_api_names
	DeprecatedApiNames []string `json:"deprecatedApiNames,omitempty"`

	// True if the dimension is custom to this property. This includes user,
	//  event, & item scoped custom dimensions; to learn more about custom
	//  dimensions, see https://support.google.com/analytics/answer/14240153. This
	//  also include custom channel groups; to learn more about custom channel
	//  groups, see https://support.google.com/analytics/answer/13051316.
	// +kcc:proto:field=google.analytics.data.v1beta.DimensionMetadata.custom_definition
	CustomDefinition *bool `json:"customDefinition,omitempty"`

	// The display name of the category that this dimension belongs to. Similar
	//  dimensions and metrics are categorized together.
	// +kcc:proto:field=google.analytics.data.v1beta.DimensionMetadata.category
	Category *string `json:"category,omitempty"`
}

// +kcc:proto=google.analytics.data.v1beta.Metadata
type Metadata struct {
	// Resource name of this metadata.
	// +kcc:proto:field=google.analytics.data.v1beta.Metadata.name
	Name *string `json:"name,omitempty"`

	// The dimension descriptions.
	// +kcc:proto:field=google.analytics.data.v1beta.Metadata.dimensions
	Dimensions []DimensionMetadata `json:"dimensions,omitempty"`

	// The metric descriptions.
	// +kcc:proto:field=google.analytics.data.v1beta.Metadata.metrics
	Metrics []MetricMetadata `json:"metrics,omitempty"`

	// The comparison descriptions.
	// +kcc:proto:field=google.analytics.data.v1beta.Metadata.comparisons
	Comparisons []ComparisonMetadata `json:"comparisons,omitempty"`
}

// +kcc:proto=google.analytics.data.v1beta.MetricMetadata
type MetricMetadata struct {
	// A metric name. Useable in [Metric](#Metric)'s `name`. For example,
	//  `eventCount`.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.api_name
	ApiName *string `json:"apiName,omitempty"`

	// This metric's name within the Google Analytics user interface. For example,
	//  `Event count`.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.ui_name
	UiName *string `json:"uiName,omitempty"`

	// Description of how this metric is used and calculated.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.description
	Description *string `json:"description,omitempty"`

	// Still usable but deprecated names for this metric. If populated, this
	//  metric is available by either `apiName` or one of `deprecatedApiNames`
	//  for a period of time. After the deprecation period, the metric will be
	//  available only by `apiName`.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.deprecated_api_names
	DeprecatedApiNames []string `json:"deprecatedApiNames,omitempty"`

	// The type of this metric.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.type
	Type *string `json:"type,omitempty"`

	// The mathematical expression for this derived metric. Can be used in
	//  [Metric](#Metric)'s `expression` field for equivalent reports. Most metrics
	//  are not expressions, and for non-expressions, this field is empty.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.expression
	Expression *string `json:"expression,omitempty"`

	// True if the metric is a custom metric for this property.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.custom_definition
	CustomDefinition *bool `json:"customDefinition,omitempty"`

	// If reasons are specified, your access is blocked to this metric for this
	//  property. API requests from you to this property for this metric will
	//  succeed; however, the report will contain only zeros for this metric. API
	//  requests with metric filters on blocked metrics will fail. If reasons are
	//  empty, you have access to this metric.
	//
	//  To learn more, see [Access and data-restriction
	//  management](https://support.google.com/analytics/answer/10851388).
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.blocked_reasons
	BlockedReasons []string `json:"blockedReasons,omitempty"`

	// The display name of the category that this metrics belongs to. Similar
	//  dimensions and metrics are categorized together.
	// +kcc:proto:field=google.analytics.data.v1beta.MetricMetadata.category
	Category *string `json:"category,omitempty"`
}
