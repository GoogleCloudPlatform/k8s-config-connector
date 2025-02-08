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


// +kcc:proto=google.cloud.datacatalog.lineage.v1.Origin
type Origin struct {
	// Type of the source.
	//
	//  Use of a source_type other than `CUSTOM` for process creation
	//  or updating is highly discouraged, and may be restricted in the future
	//  without notice.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Origin.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// If the source_type isn't CUSTOM, the value of this field should be a GCP
	//  resource name of the system, which reports lineage. The project and
	//  location parts of the resource name must match the project and location of
	//  the lineage resource being created. Examples:
	//
	//  - `{source_type: COMPOSER, name:
	//    "projects/foo/locations/us/environments/bar"}`
	//  - `{source_type: BIGQUERY, name: "projects/foo/locations/eu"}`
	//  - `{source_type: CUSTOM,   name: "myCustomIntegration"}`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Origin.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.lineage.v1.Process
type Process struct {
	// Immutable. The resource name of the lineage process. Format:
	//  `projects/{project}/locations/{location}/processes/{process}`.
	//  Can be specified or auto-assigned.
	//  {process} must be not longer than 200 characters and only
	//  contain characters in a set: `a-zA-Z0-9_-:.`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Process.name
	Name *string `json:"name,omitempty"`

	// Optional. A human-readable name you can set to display in a user interface.
	//  Must be not longer than 200 characters and only contain UTF-8 letters
	//  or numbers, spaces or characters like `_-:&.`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Process.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. The origin of this process and its runs and lineage events.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Process.origin
	Origin *Origin `json:"origin,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}
