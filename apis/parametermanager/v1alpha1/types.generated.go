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


// +kcc:proto=google.cloud.parametermanager.v1.ParameterVersion
type ParameterVersion struct {
	// Identifier. [Output only] The resource name of the ParameterVersion in the
	//  format `projects/*/locations/*/parameters/*/versions/*`.
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.name
	Name *string `json:"name,omitempty"`

	// Optional. Disabled boolean to determine if a ParameterVersion acts as a
	//  metadata only resource (payload is never returned if disabled is true). If
	//  true any calls will always default to BASIC view even if the user
	//  explicitly passes FULL view as part of the request. A render call on a
	//  disabled resource fails with an error. Default value is False.
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Required. Immutable. Payload content of a ParameterVersion resource.  This
	//  is only returned when the request provides the View value of FULL (default
	//  for GET request).
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.payload
	Payload *ParameterVersionPayload `json:"payload,omitempty"`
}

// +kcc:proto=google.cloud.parametermanager.v1.ParameterVersionPayload
type ParameterVersionPayload struct {
	// Required. bytes data for storing payload.
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersionPayload.data
	Data []byte `json:"data,omitempty"`
}

// +kcc:proto=google.cloud.parametermanager.v1.ParameterVersion
type ParameterVersionObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.ParameterVersion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
