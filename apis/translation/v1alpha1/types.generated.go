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


// +kcc:proto=google.cloud.translation.v3.AdaptiveMtFile
type AdaptiveMtFile struct {
	// Required. The resource name of the file, in form of
	//  `projects/{project-number-or-id}/locations/{location_id}/adaptiveMtDatasets/{dataset}/adaptiveMtFiles/{file}`
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtFile.name
	Name *string `json:"name,omitempty"`

	// The file's display name.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtFile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The number of entries that the file contains.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtFile.entry_count
	EntryCount *int32 `json:"entryCount,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.AdaptiveMtFile
type AdaptiveMtFileObservedState struct {
	// Output only. Timestamp when this file was created.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtFile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this file was last updated.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtFile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
