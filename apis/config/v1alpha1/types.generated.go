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


// +kcc:proto=google.cloud.config.v1.GitSource
type GitSource struct {
	// Optional. Repository URL.
	//  Example: 'https://github.com/kubernetes/examples.git'
	// +kcc:proto:field=google.cloud.config.v1.GitSource.repo
	Repo *string `json:"repo,omitempty"`

	// Optional. Subdirectory inside the repository.
	//  Example: 'staging/my-package'
	// +kcc:proto:field=google.cloud.config.v1.GitSource.directory
	Directory *string `json:"directory,omitempty"`

	// Optional. Git reference (e.g. branch or tag).
	// +kcc:proto:field=google.cloud.config.v1.GitSource.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.PreviewArtifacts
type PreviewArtifacts struct {
}

// +kcc:proto=google.cloud.config.v1.TerraformBlueprint
type TerraformBlueprint struct {
	// URI of an object in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`
	//
	//  URI may also specify an object version for zipped objects.
	//  Format: `gs://{bucket}/{object}#{version}`
	// +kcc:proto:field=google.cloud.config.v1.TerraformBlueprint.gcs_source
	GcsSource *string `json:"gcsSource,omitempty"`

	// URI of a public Git repo.
	// +kcc:proto:field=google.cloud.config.v1.TerraformBlueprint.git_source
	GitSource *GitSource `json:"gitSource,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.config.v1.TerraformError
type TerraformError struct {
	// Address of the resource associated with the error,
	//  e.g. `google_compute_network.vpc_network`.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.resource_address
	ResourceAddress *string `json:"resourceAddress,omitempty"`

	// HTTP response code returned from Google Cloud Platform APIs when Terraform
	//  fails to provision the resource. If unset or 0, no HTTP response code was
	//  returned by Terraform.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.http_response_code
	HTTPResponseCode *int32 `json:"httpResponseCode,omitempty"`

	// A human-readable error description.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.error_description
	ErrorDescription *string `json:"errorDescription,omitempty"`

	// Original error response from underlying Google API, if available.
	// +kcc:proto:field=google.cloud.config.v1.TerraformError.error
	Error *Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.TerraformVariable
type TerraformVariable struct {
	// Input variable value.
	// +kcc:proto:field=google.cloud.config.v1.TerraformVariable.input_value
	InputValue *Value `json:"inputValue,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
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

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.PreviewArtifacts
type PreviewArtifactsObservedState struct {
	// Output only. Location of a blueprint copy and other content in Google Cloud
	//  Storage. Format: `gs://{bucket}/{object}`
	// +kcc:proto:field=google.cloud.config.v1.PreviewArtifacts.content
	Content *string `json:"content,omitempty"`

	// Output only. Location of artifacts in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`
	// +kcc:proto:field=google.cloud.config.v1.PreviewArtifacts.artifacts
	Artifacts *string `json:"artifacts,omitempty"`
}
