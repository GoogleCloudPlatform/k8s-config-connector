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


// +kcc:proto=google.cloud.dataplex.v1.EntryGroup
type EntryGroup struct {

	// Optional. Description of the EntryGroup.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the EntryGroup.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// This checksum is computed by the service, and might be sent on update and
	//  delete requests to ensure the client has an up-to-date value before
	//  proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryGroup
type EntryGroupObservedState struct {
	// Output only. The relative resource name of the EntryGroup, in the format
	//  projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the EntryGroup. If you
	//  delete and recreate the EntryGroup with the same name, this ID will be
	//  different.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the EntryGroup was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the EntryGroup was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Denotes the transfer status of the Entry Group. It is
	//  unspecified for Entry Group created from Dataplex API.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.transfer_status
	TransferStatus *string `json:"transferStatus,omitempty"`
}
