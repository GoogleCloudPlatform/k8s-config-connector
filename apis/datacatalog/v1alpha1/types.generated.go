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

// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestamps struct {
	// Creation timestamp of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Timestamp of the last modification of the resource or its metadata within
	//  a given system.
	//
	//  Note: Depending on the source system, not every modification updates this
	//  timestamp.
	//  For example, BigQuery timestamps every metadata modification but not data
	//  or permission changes.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestampsObservedState struct {
	// Output only. Expiration timestamp of the resource within the given system.
	//
	//  Currently only applicable to BigQuery resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
