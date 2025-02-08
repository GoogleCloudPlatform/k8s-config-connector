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


// +kcc:proto=google.cloud.dialogflow.v2.Version
type Version struct {

	// Optional. The developer-provided description of this version.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Version.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Version
type VersionObservedState struct {
	// Output only. The unique identifier of this agent version.
	//  Supported formats:
	//
	//  - `projects/<Project ID>/agent/versions/<Version ID>`
	//  - `projects/<Project ID>/locations/<Location ID>/agent/versions/<Version
	//    ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2.Version.name
	Name *string `json:"name,omitempty"`

	// Output only. The sequential number of this version. This field is read-only
	//  which means it cannot be set by create and update methods.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Version.version_number
	VersionNumber *int32 `json:"versionNumber,omitempty"`

	// Output only. The creation time of this version. This field is read-only,
	//  i.e., it cannot be set by create and update methods.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Version.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The status of this version. This field is read-only and cannot
	//  be set by create and update methods.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Version.status
	Status *string `json:"status,omitempty"`
}
