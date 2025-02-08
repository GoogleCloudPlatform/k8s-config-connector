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


// +kcc:proto=google.cloud.netapp.v1.Snapshot
type Snapshot struct {
	// Identifier. The resource name of the snapshot.
	//  Format:
	//  `projects/{project_id}/locations/{location}/volumes/{volume_id}/snapshots/{snapshot_id}`.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.name
	Name *string `json:"name,omitempty"`

	// A description of the snapshot with 2048 characters or less.
	//  Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.description
	Description *string `json:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.Snapshot
type SnapshotObservedState struct {
	// Output only. The snapshot state.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Current storage usage for the snapshot in bytes.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.used_bytes
	UsedBytes *float64 `json:"usedBytes,omitempty"`

	// Output only. The time when the snapshot was created.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
