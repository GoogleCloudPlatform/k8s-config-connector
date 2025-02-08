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


// +kcc:proto=google.cloud.securitycenter.v1beta1.Finding
type Finding struct {
	// The relative resource name of this finding. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Example:
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}"
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.name
	Name *string `json:"name,omitempty"`

	// Immutable. The relative resource name of the source the finding belongs to.
	//  See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  This field is immutable after creation time.
	//  For example:
	//  "organizations/{organization_id}/sources/{source_id}"
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.parent
	Parent *string `json:"parent,omitempty"`

	// For findings on Google Cloud resources, the full resource
	//  name of the Google Cloud resource this finding is for. See:
	//  https://cloud.google.com/apis/design/resource_names#full_resource_name
	//  When the finding is for a non-Google Cloud resource, the resourceName can
	//  be a customer or partner defined string. This field is immutable after
	//  creation time.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// The state of the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.state
	State *string `json:"state,omitempty"`

	// The additional taxonomy group within findings from a given source.
	//  This field is immutable after creation time.
	//  Example: "XSS_FLASH_INJECTION"
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.category
	Category *string `json:"category,omitempty"`

	// The URI that, if available, points to a web page outside of Security
	//  Command Center where additional information about the finding can be found.
	//  This field is guaranteed to be either empty or a well formed URL.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.external_uri
	ExternalURI *string `json:"externalURI,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The time at which the event took place, or when an update to the finding
	//  occurred. For example, if the finding represents an open firewall it would
	//  capture the time the detector believes the firewall became open. The
	//  accuracy is determined by the detector. If the finding were to be resolved
	//  afterward, this time would reflect when the finding was resolved.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.event_time
	EventTime *string `json:"eventTime,omitempty"`

	// The time at which the finding was created in Security Command Center.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v1beta1.SecurityMarks
type SecurityMarks struct {
	// The relative resource name of the SecurityMarks. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Examples:
	//  "organizations/{organization_id}/assets/{asset_id}/securityMarks"
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}/securityMarks".
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.SecurityMarks.name
	Name *string `json:"name,omitempty"`

	// Mutable user specified security marks belonging to the parent resource.
	//  Constraints are as follows:
	//
	//    * Keys and values are treated as case insensitive
	//    * Keys must be between 1 - 256 characters (inclusive)
	//    * Keys must be letters, numbers, underscores, or dashes
	//    * Values have leading and trailing whitespace trimmed, remaining
	//      characters must be between 1 - 4096 characters (inclusive)
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.SecurityMarks.marks
	Marks map[string]string `json:"marks,omitempty"`
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

// +kcc:proto=google.cloud.securitycenter.v1beta1.Finding
type FindingObservedState struct {
	// Output only. User specified security marks. These marks are entirely
	//  managed by the user and come from the SecurityMarks resource that belongs
	//  to the finding.
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.Finding.security_marks
	SecurityMarks *SecurityMarks `json:"securityMarks,omitempty"`
}
