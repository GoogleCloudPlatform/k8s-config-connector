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


// +kcc:proto=google.cloud.datacatalog.lineage.v1.Run
type Run struct {
	// Immutable. The resource name of the run. Format:
	//  `projects/{project}/locations/{location}/processes/{process}/runs/{run}`.
	//  Can be specified or auto-assigned.
	//  {run} must be not longer than 200 characters and only
	//  contain characters in a set: `a-zA-Z0-9_-:.`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.name
	Name *string `json:"name,omitempty"`

	// Optional. A human-readable name you can set to display in a user interface.
	//  Must be not longer than 1024 characters and only contain UTF-8 letters
	//  or numbers, spaces or characters like `_-:&.`
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Required. The timestamp of the start of the run.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. The timestamp of the end of the run.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Required. The state of the run.
	// +kcc:proto:field=google.cloud.datacatalog.lineage.v1.Run.state
	State *string `json:"state,omitempty"`
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
