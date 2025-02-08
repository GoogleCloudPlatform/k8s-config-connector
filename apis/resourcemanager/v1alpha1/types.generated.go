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


// +kcc:proto=google.cloud.resourcemanager.v3.Folder
type Folder struct {

	// Required. The folder's parent's resource name.
	//  Updates to the folder's parent must be performed using
	//  [MoveFolder][google.cloud.resourcemanager.v3.Folders.MoveFolder].
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.parent
	Parent *string `json:"parent,omitempty"`

	// The folder's display name.
	//  A folder's display name must be unique amongst its siblings. For example,
	//  no two folders with the same parent can share the same display name.
	//  The display name must start and end with a letter or digit, may contain
	//  letters, digits, spaces, hyphens and underscores and can be no longer
	//  than 30 characters. This is captured by the regular expression:
	//  `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v3.Folder
type FolderObservedState struct {
	// Output only. The resource name of the folder.
	//  Its format is `folders/{folder_id}`, for example: "folders/1234".
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.name
	Name *string `json:"name,omitempty"`

	// Output only. The lifecycle state of the folder.
	//  Updates to the state must be performed using
	//  [DeleteFolder][google.cloud.resourcemanager.v3.Folders.DeleteFolder] and
	//  [UndeleteFolder][google.cloud.resourcemanager.v3.Folders.UndeleteFolder].
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.state
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when the folder was created.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the folder was last modified.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when the folder was requested to be deleted.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. A checksum computed by the server based on the current value
	//  of the folder resource. This may be sent on update and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Folder.etag
	Etag *string `json:"etag,omitempty"`
}
