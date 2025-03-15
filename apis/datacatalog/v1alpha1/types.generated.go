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

// +kcc:proto=google.cloud.datacatalog.v1.EntryGroup
type EntryGroup struct {
	// Identifier. The resource name of the entry group in URL format.
	//
	//  Note: The entry group itself and its child resources might not be
	//  stored in the location specified in its name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.name
	Name *string `json:"name,omitempty"`

	// A short name to identify the entry group, for example,
	//  "analytics data - jan 2011". Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Entry group description. Can consist of several sentences or
	//  paragraphs that describe the entry group contents.
	//  Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. When set to [true], it means DataCatalog EntryGroup was
	//  transferred to Dataplex Catalog Service. It makes EntryGroup and its
	//  Entries to be read-only in DataCatalog. However, new Tags on EntryGroup and
	//  its Entries can be created. After setting the flag to [true] it cannot be
	//  unset.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.transferred_to_dataplex
	TransferredToDataplex *bool `json:"transferredToDataplex,omitempty"`
}

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

// +kcc:proto=google.cloud.datacatalog.v1.EntryGroup
type EntryGroupObservedState struct {
	// Output only. Timestamps of the entry group. Default value is empty.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.data_catalog_timestamps
	DataCatalogTimestamps *SystemTimestamps `json:"dataCatalogTimestamps,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestampsObservedState struct {
	// Output only. Expiration timestamp of the resource within the given system.
	//
	//  Currently only applicable to BigQuery resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}
