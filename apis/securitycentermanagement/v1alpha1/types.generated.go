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


// +kcc:proto=google.cloud.securitycentermanagement.v1.SimulatedFinding
type SimulatedFinding struct {
	// Identifier. The [relative resource name](https://google.aip.dev/122) of the
	//  finding, in one of the following formats:
	//
	//  * `organizations/{organization_id}/sources/{source_id}/findings/{finding_id}`
	//  * `folders/{folder_id}/sources/{source_id}/findings/{finding_id}`
	//  * `projects/{project_id}/sources/{source_id}/findings/{finding_id}`
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.name
	Name *string `json:"name,omitempty"`

	// The [relative resource name](https://google.aip.dev/122) of the source the
	//  finding belongs to. For example,
	//  `organizations/{organization_id}/sources/{source_id}`. This field is
	//  immutable after creation time.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.parent
	Parent *string `json:"parent,omitempty"`

	// For findings on Google Cloud resources, the
	//  [full resource name](https://google.aip.dev/122#full-resource-names) of the
	//  Google Cloud resource this finding is for. When the finding is for a
	//  non-Google Cloud resource, the value can be a customer or partner defined
	//  string. This field is immutable after creation time.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// The additional taxonomy group within findings from a given source. For
	//  example, `XSS_FLASH_INJECTION`. This field is immutable after creation
	//  time.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.category
	Category *string `json:"category,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The time the finding was first detected. If an existing finding is updated,
	//  then this is the time the update occurred. If the finding is later
	//  resolved, then this time reflects when the finding was resolved.
	//
	//  For example, if the finding represents an open firewall, this property
	//  captures the time the detector believes the firewall became open. The
	//  accuracy is determined by the detector.
	//
	//  The event time must not be set to a value greater than the current
	//  timestamp.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.event_time
	EventTime *string `json:"eventTime,omitempty"`

	// The severity of the finding. This field is managed by the source that
	//  writes the finding.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.severity
	Severity *string `json:"severity,omitempty"`

	// The class of the finding.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.finding_class
	FindingClass *string `json:"findingClass,omitempty"`
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

// +kcc:proto=google.cloud.securitycentermanagement.v1.SimulatedFinding
type SimulatedFindingObservedState struct {
	// Output only. The state of the finding.
	// +kcc:proto:field=google.cloud.securitycentermanagement.v1.SimulatedFinding.state
	State *string `json:"state,omitempty"`
}
