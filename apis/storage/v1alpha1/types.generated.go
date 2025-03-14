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


// +kcc:proto=google.storage.control.v2.Folder
type Folder struct {
	// Identifier. The name of this folder.
	//  Format: `projects/{project}/buckets/{bucket}/folders/{folder}`
	// +kcc:proto:field=google.storage.control.v2.Folder.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.storage.control.v2.PendingRenameInfo
type PendingRenameInfo struct {
}

// +kcc:proto=google.storage.control.v2.Folder
type FolderObservedState struct {
	// Output only. The version of the metadata for this folder. Used for
	//  preconditions and for detecting changes in metadata.
	// +kcc:proto:field=google.storage.control.v2.Folder.metageneration
	Metageneration *int64 `json:"metageneration,omitempty"`

	// Output only. The creation time of the folder.
	// +kcc:proto:field=google.storage.control.v2.Folder.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The modification time of the folder.
	// +kcc:proto:field=google.storage.control.v2.Folder.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only present if the folder is part of an ongoing RenameFolder
	//  operation. Contains information which can be used to query the operation
	//  status. The presence of this field also indicates all write operations are
	//  blocked for this folder, including folder, managed folder, and object
	//  operations.
	// +kcc:proto:field=google.storage.control.v2.Folder.pending_rename_info
	PendingRenameInfo *PendingRenameInfo `json:"pendingRenameInfo,omitempty"`
}

// +kcc:proto=google.storage.control.v2.PendingRenameInfo
type PendingRenameInfoObservedState struct {
	// Output only. The name of the rename operation.
	// +kcc:proto:field=google.storage.control.v2.PendingRenameInfo.operation
	Operation *string `json:"operation,omitempty"`
}
