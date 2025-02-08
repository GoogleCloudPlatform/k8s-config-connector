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


// +kcc:proto=google.cloud.gkebackup.v1.NamespacedName
type NamespacedName struct {
	// Optional. The Namespace of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.namespace
	Namespace *string `json:"namespace,omitempty"`

	// Optional. The name of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.VolumeRestore
type VolumeRestore struct {
}

// +kcc:proto=google.cloud.gkebackup.v1.VolumeRestore
type VolumeRestoreObservedState struct {
	// Output only. Full name of the VolumeRestore resource.
	//  Format: `projects/*/locations/*/restorePlans/*/restores/*/volumeRestores/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.name
	Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp when this VolumeRestore resource was
	//  created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this VolumeRestore resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The full name of the VolumeBackup from which the volume will
	//  be restored. Format:
	//  `projects/*/locations/*/backupPlans/*/backups/*/volumeBackups/*`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.volume_backup
	VolumeBackup *string `json:"volumeBackup,omitempty"`

	// Output only. The reference to the target Kubernetes PVC to be restored.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.target_pvc
	TargetPvc *NamespacedName `json:"targetPvc,omitempty"`

	// Output only. A storage system-specific opaque handler to the underlying
	//  volume created for the target PVC from the volume backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.volume_handle
	VolumeHandle *string `json:"volumeHandle,omitempty"`

	// Output only. The type of volume provisioned
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.volume_type
	VolumeType *string `json:"volumeType,omitempty"`

	// Output only. The timestamp when the associated underlying volume
	//  restoration completed.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. The current state of this VolumeRestore.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.state
	State *string `json:"state,omitempty"`

	// Output only. A human readable message explaining why the VolumeRestore is
	//  in its current state.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a volume restore from overwriting each
	//  other. It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform volume restore updates in order to avoid
	//  race conditions.
	// +kcc:proto:field=google.cloud.gkebackup.v1.VolumeRestore.etag
	Etag *string `json:"etag,omitempty"`
}
