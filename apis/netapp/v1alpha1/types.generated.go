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


// +kcc:proto=google.cloud.netapp.v1.DestinationVolumeParameters
type DestinationVolumeParameters struct {
	// Required. Existing destination StoragePool name.
	// +kcc:proto:field=google.cloud.netapp.v1.DestinationVolumeParameters.storage_pool
	StoragePool *string `json:"storagePool,omitempty"`

	// Desired destination volume resource id. If not specified, source volume's
	//  resource id will be used.
	//  This value must start with a lowercase letter followed by up to 62
	//  lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
	// +kcc:proto:field=google.cloud.netapp.v1.DestinationVolumeParameters.volume_id
	VolumeID *string `json:"volumeID,omitempty"`

	// Destination volume's share name. If not specified, source volume's share
	//  name will be used.
	// +kcc:proto:field=google.cloud.netapp.v1.DestinationVolumeParameters.share_name
	ShareName *string `json:"shareName,omitempty"`

	// Description for the destination volume.
	// +kcc:proto:field=google.cloud.netapp.v1.DestinationVolumeParameters.description
	Description *string `json:"description,omitempty"`

	// Optional. Tiering policy for the volume.
	// +kcc:proto:field=google.cloud.netapp.v1.DestinationVolumeParameters.tiering_policy
	TieringPolicy *TieringPolicy `json:"tieringPolicy,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.HybridPeeringDetails
type HybridPeeringDetails struct {
	// Optional. IP address of the subnet.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridPeeringDetails.subnet_ip
	SubnetIP *string `json:"subnetIP,omitempty"`

	// Optional. Copy-paste-able commands to be used on user's ONTAP to accept
	//  peering requests.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridPeeringDetails.command
	Command *string `json:"command,omitempty"`

	// Optional. Expiration time for the peering command to be executed on user's
	//  ONTAP.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridPeeringDetails.command_expiry_time
	CommandExpiryTime *string `json:"commandExpiryTime,omitempty"`

	// Optional. Temporary passphrase generated to accept cluster peering command.
	// +kcc:proto:field=google.cloud.netapp.v1.HybridPeeringDetails.passphrase
	Passphrase *string `json:"passphrase,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.Replication
type Replication struct {
	// Identifier. The resource name of the Replication.
	//  Format:
	//  `projects/{project_id}/locations/{location}/volumes/{volume_id}/replications/{replication_id}`.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.name
	Name *string `json:"name,omitempty"`

	// Required. Indicates the schedule for replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.replication_schedule
	ReplicationSchedule *string `json:"replicationSchedule,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A description about this replication relationship.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.description
	Description *string `json:"description,omitempty"`

	// Required. Input only. Destination volume parameters
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.destination_volume_parameters
	DestinationVolumeParameters *DestinationVolumeParameters `json:"destinationVolumeParameters,omitempty"`

	// Optional. Location of the user cluster.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.cluster_location
	ClusterLocation *string `json:"clusterLocation,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.TieringPolicy
type TieringPolicy struct {
	// Optional. Flag indicating if the volume has tiering policy enable/pause.
	//  Default is PAUSED.
	// +kcc:proto:field=google.cloud.netapp.v1.TieringPolicy.tier_action
	TierAction *string `json:"tierAction,omitempty"`

	// Optional. Time in days to mark the volume's data block as cold and make it
	//  eligible for tiering, can be range from 7-183. Default is 31.
	// +kcc:proto:field=google.cloud.netapp.v1.TieringPolicy.cooling_threshold_days
	CoolingThresholdDays *int32 `json:"coolingThresholdDays,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.TransferStats
type TransferStats struct {
	// Cumulative bytes trasferred so far for the replication relatinonship.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.transfer_bytes
	TransferBytes *int64 `json:"transferBytes,omitempty"`

	// Cumulative time taken across all transfers for the replication
	//  relationship.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.total_transfer_duration
	TotalTransferDuration *string `json:"totalTransferDuration,omitempty"`

	// Last transfer size in bytes.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.last_transfer_bytes
	LastTransferBytes *int64 `json:"lastTransferBytes,omitempty"`

	// Time taken during last transfer.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.last_transfer_duration
	LastTransferDuration *string `json:"lastTransferDuration,omitempty"`

	// Lag duration indicates the duration by which Destination region volume
	//  content lags behind the primary region volume content.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.lag_duration
	LagDuration *string `json:"lagDuration,omitempty"`

	// Time when progress was updated last.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Time when last transfer completed.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.last_transfer_end_time
	LastTransferEndTime *string `json:"lastTransferEndTime,omitempty"`

	// A message describing the cause of the last transfer failure.
	// +kcc:proto:field=google.cloud.netapp.v1.TransferStats.last_transfer_error
	LastTransferError *string `json:"lastTransferError,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.Replication
type ReplicationObservedState struct {
	// Output only. State of the replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Indicates whether this points to source or destination.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.role
	Role *string `json:"role,omitempty"`

	// Output only. Indicates the state of mirroring.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.mirror_state
	MirrorState *string `json:"mirrorState,omitempty"`

	// Output only. Condition of the relationship. Can be one of the following:
	//  - true: The replication relationship is healthy. It has not missed the most
	//  recent scheduled transfer.
	//  - false: The replication relationship is not healthy. It has missed the
	//  most recent scheduled transfer.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.healthy
	Healthy *bool `json:"healthy,omitempty"`

	// Output only. Replication create time.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Full name of destination volume resource.
	//  Example : "projects/{project}/locations/{location}/volumes/{volume_id}"
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.destination_volume
	DestinationVolume *string `json:"destinationVolume,omitempty"`

	// Output only. Replication transfer statistics.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.transfer_stats
	TransferStats *TransferStats `json:"transferStats,omitempty"`

	// Output only. Full name of source volume resource.
	//  Example : "projects/{project}/locations/{location}/volumes/{volume_id}"
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.source_volume
	SourceVolume *string `json:"sourceVolume,omitempty"`

	// Output only. Hybrid peering details.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.hybrid_peering_details
	HybridPeeringDetails *HybridPeeringDetails `json:"hybridPeeringDetails,omitempty"`

	// Output only. Type of the hybrid replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.hybrid_replication_type
	HybridReplicationType *string `json:"hybridReplicationType,omitempty"`
}
