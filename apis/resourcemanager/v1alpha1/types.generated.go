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


// +kcc:proto=google.cloud.resourcemanager.v2.Folder
type Folder struct {

	// Required. The Folder’s parent's resource name.
	//  Updates to the folder's parent must be performed via
	//  [MoveFolder][google.cloud.resourcemanager.v2.Folders.MoveFolder].
	// +kcc:proto:field=google.cloud.resourcemanager.v2.Folder.parent
	Parent *string `json:"parent,omitempty"`

	// The folder’s display name.
	//  A folder’s display name must be unique amongst its siblings, e.g.
	//  no two folders with the same parent can share the same display name.
	//  The display name must start and end with a letter or digit, may contain
	//  letters, digits, spaces, hyphens and underscores and can be no longer
	//  than 30 characters. This is captured by the regular expression:
	//  [\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?.
	// +kcc:proto:field=google.cloud.resourcemanager.v2.Folder.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v2.Folder
type FolderObservedState struct {
	// Output only. The resource name of the Folder.
	//  Its format is `folders/{folder_id}`, for example: "folders/1234".
	// +kcc:proto:field=google.cloud.resourcemanager.v2.Folder.name
	Name *string `json:"name,omitempty"`

	// Output only. The lifecycle state of the folder.
	//  Updates to the lifecycle_state must be performed via
	//  [DeleteFolder][google.cloud.resourcemanager.v2.Folders.DeleteFolder] and
	//  [UndeleteFolder][google.cloud.resourcemanager.v2.Folders.UndeleteFolder].
	// +kcc:proto:field=google.cloud.resourcemanager.v2.Folder.lifecycle_state
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// Output only. Timestamp when the Folder was created. Assigned by the server.
	// +kcc:proto:field=google.cloud.resourcemanager.v2.Folder.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the Folder was last modified.
	// +kcc:proto:field=google.cloud.resourcemanager.v2.Folder.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
