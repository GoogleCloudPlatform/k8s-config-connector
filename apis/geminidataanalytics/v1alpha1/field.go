// Copyright 2026 Google LLC
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

// +kcc:proto=google.cloud.geminidataanalytics.v1.Field
type Field struct {
	// Optional. The name of the field.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.name
	Name *string `json:"name,omitempty"`

	// Optional. The description of the field.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.description
	Description *string `json:"description,omitempty"`

	// Optional. The type of the field (e.g., STRING, INTEGER).
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.type
	Type *string `json:"type,omitempty"`

	// Optional. The mode of the field (e.g., NULLABLE, REPEATED).
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. A list of alternative names or synonyms that can be used to
	//  refer to this field. For example: ["id", "customerid", "cust_id"]. Currently
	//  only used for BigQuery data sources.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.synonyms
	Synonyms []string `json:"synonyms,omitempty"`

	// Optional. A list of tags or keywords associated with the field, used for
	//  categorization. For example: ["identifier", "customer", "pii"]. Currently
	//  only used for BigQuery data sources.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.tags
	Tags []string `json:"tags,omitempty"`

	// Optional. Field display_name (same as label in
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Recursive property for nested schema structures.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.subfields
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	Subfields []Field `json:"subfields,omitempty"`

	// Optional. Field category, not required, currently only useful for Looker.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.category
	Category *string `json:"category,omitempty"`

	// Optional. Looker only. Value format of the field.
	//  Ref:
	//  https://cloud.google.com/looker/docs/reference/param-field-value-format
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.Field.value_format
	ValueFormat *string `json:"valueFormat,omitempty"`
}
