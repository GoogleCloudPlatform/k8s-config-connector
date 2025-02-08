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


// +kcc:proto=google.cloud.filestore.v1beta1.Snapshot
type Snapshot struct {

	// A description of the snapshot with 2048 characters or less.
	//  Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Snapshot.description
	Description *string `json:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Snapshot.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1beta1.Snapshot
type SnapshotObservedState struct {
	// Output only. The resource name of the snapshot, in the format
	//  `projects/{project_id}/locations/{location_id}/instances/{instance_id}/snapshots/{snapshot_id}`.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Snapshot.name
	Name *string `json:"name,omitempty"`

	// Output only. The snapshot state.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Snapshot.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the snapshot was created.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Snapshot.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The amount of bytes needed to allocate a full copy of the
	//  snapshot content
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Snapshot.filesystem_used_bytes
	FilesystemUsedBytes *int64 `json:"filesystemUsedBytes,omitempty"`
}
