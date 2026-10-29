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

// +kcc:proto=google.cloud.dataplex.v1.Schema.SchemaField
type Schema_SchemaField struct {
	// Required. The name of the field. Must contain only letters, numbers and
	//  underscores, with a maximum length of 767 characters,
	//  and must begin with a letter or underscore.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.name
	// +required
	Name *string `json:"name,omitempty"`

	// Optional. User friendly field description. Must be less than or equal to
	//  1024 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.description
	Description *string `json:"description,omitempty"`

	// Required. The type of field.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.type
	// +required
	Type *string `json:"type,omitempty"`

	// Required. Additional field semantics.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.mode
	// +required
	Mode *string `json:"mode,omitempty"`

	// Optional. Any nested field for complex types.
	// +kcc:proto:field=google.cloud.dataplex.v1.Schema.SchemaField.fields
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	Fields []Schema_SchemaField `json:"fields,omitempty"`
}
