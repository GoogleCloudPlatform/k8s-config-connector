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


// +kcc:proto=google.cloud.datacatalog.v1beta1.EntryGroup
type EntryGroup struct {
	// Identifier. The resource name of the entry group in URL format. Example:
	//
	//  * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}
	//
	//  Note that this EntryGroup and its child resources may not actually be
	//  stored in the location in this name.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.EntryGroup.name
	Name *string `json:"name,omitempty"`

	// A short name to identify the entry group, for example,
	//  "analytics data - jan 2011". Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.EntryGroup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Entry group description, which can consist of several sentences or
	//  paragraphs that describe entry group contents. Default value is an empty
	//  string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.EntryGroup.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.SystemTimestamps
type SystemTimestamps struct {
	// The creation time of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.SystemTimestamps.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The last-modified time of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.SystemTimestamps.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.EntryGroup
type EntryGroupObservedState struct {
	// Output only. Timestamps about this EntryGroup. Default value is empty
	//  timestamps.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.EntryGroup.data_catalog_timestamps
	DataCatalogTimestamps *SystemTimestamps `json:"dataCatalogTimestamps,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.SystemTimestamps
type SystemTimestampsObservedState struct {
	// Output only. The expiration time of the resource within the given system.
	//  Currently only apllicable to BigQuery resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.SystemTimestamps.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
