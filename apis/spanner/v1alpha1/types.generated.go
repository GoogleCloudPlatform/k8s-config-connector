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


// +kcc:proto=google.spanner.admin.instance.v1.InstancePartition
type InstancePartition struct {
	// Required. A unique identifier for the instance partition. Values are of the
	//  form
	//  `projects/<project>/instances/<instance>/instancePartitions/[a-z][-a-z0-9]*[a-z0-9]`.
	//  The final segment of the name must be between 2 and 64 characters in
	//  length. An instance partition's name cannot be changed after the instance
	//  partition is created.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the instance partition's configuration. Values are of
	//  the form `projects/<project>/instanceConfigs/<configuration>`. See also
	//  [InstanceConfig][google.spanner.admin.instance.v1.InstanceConfig] and
	//  [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs].
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.config
	Config *string `json:"config,omitempty"`

	// Required. The descriptive name for this instance partition as it appears in
	//  UIs. Must be unique per project and between 4 and 30 characters in length.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The number of nodes allocated to this instance partition.
	//
	//  Users can set the `node_count` field to specify the target number of
	//  nodes allocated to the instance partition.
	//
	//  This may be zero in API responses for instance partitions that are not
	//  yet in state `READY`.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// The number of processing units allocated to this instance partition.
	//
	//  Users can set the `processing_units` field to specify the target number
	//  of processing units allocated to the instance partition.
	//
	//  This might be zero in API responses for instance partitions that are not
	//  yet in the `READY` state.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.processing_units
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`

	// Used for optimistic concurrency control as a way
	//  to help prevent simultaneous updates of a instance partition from
	//  overwriting each other. It is strongly suggested that systems make use of
	//  the etag in the read-modify-write cycle to perform instance partition
	//  updates in order to avoid race conditions: An etag is returned in the
	//  response which contains instance partitions, and systems are expected to
	//  put that etag in the request to update instance partitions to ensure that
	//  their change will be applied to the same version of the instance partition.
	//  If no etag is provided in the call to update instance partition, then the
	//  existing instance partition is overwritten blindly.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.InstancePartition
type InstancePartitionObservedState struct {
	// Output only. The current instance partition state.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.state
	State *string `json:"state,omitempty"`

	// Output only. The time at which the instance partition was created.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the instance partition was most recently
	//  updated.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The names of the databases that reference this
	//  instance partition. Referencing databases should share the parent instance.
	//  The existence of any referencing database prevents the instance partition
	//  from being deleted.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.referencing_databases
	ReferencingDatabases []string `json:"referencingDatabases,omitempty"`

	// Output only. Deprecated: This field is not populated.
	//  Output only. The names of the backups that reference this instance
	//  partition. Referencing backups should share the parent instance. The
	//  existence of any referencing backup prevents the instance partition from
	//  being deleted.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.referencing_backups
	ReferencingBackups []string `json:"referencingBackups,omitempty"`
}
