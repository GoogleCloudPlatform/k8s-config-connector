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

// +kcc:proto=google.cloud.apigeeregistry.v1.ApiVersion
type ApiVersion struct {
	// Resource name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.name
	Name *string `json:"name,omitempty"`

	// Human-meaningful name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A detailed description.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.description
	Description *string `json:"description,omitempty"`

	// A user-definable description of the lifecycle phase of this API version.
	//  Format: free-form, but we expect single words that describe API maturity,
	//  e.g., "CONCEPT", "DESIGN", "DEVELOPMENT", "STAGING", "PRODUCTION",
	//  "DEPRECATED", "RETIRED".
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.state
	State *string `json:"state,omitempty"`

	// Labels attach identifying metadata to resources. Identifying metadata can
	//  be used to filter list operations.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one resource (System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with
	//  `apigeeregistry.googleapis.com/` and cannot be changed.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations attach non-identifying metadata to resources.
	//
	//  Annotation keys and values are less restricted than those of labels, but
	//  should be generally used for small values of broad interest. Larger, topic-
	//  specific metadata should be stored in Artifacts.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.ApiVersion
type ApiVersionObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiVersion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
