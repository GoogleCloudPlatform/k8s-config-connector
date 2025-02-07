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


// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig
type VolumeConfig struct {

	// A transient unique identifier to identify a volume within an
	//  ProvisioningConfig request.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.id
	ID *string `json:"id,omitempty"`

	// Whether snapshots should be enabled.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.snapshots_enabled
	SnapshotsEnabled *bool `json:"snapshotsEnabled,omitempty"`

	// The type of this Volume.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.type
	Type *string `json:"type,omitempty"`

	// Volume protocol.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.protocol
	Protocol *string `json:"protocol,omitempty"`

	// The requested size of this volume, in GB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.size_gb
	SizeGB *int32 `json:"sizeGB,omitempty"`

	// LUN ranges to be configured. Set only when protocol is PROTOCOL_FC.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.lun_ranges
	LunRanges []VolumeConfig_LunRange `json:"lunRanges,omitempty"`

	// Machine ids connected to this volume. Set only when protocol is
	//  PROTOCOL_FC.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.machine_ids
	MachineIds []string `json:"machineIds,omitempty"`

	// NFS exports. Set only when protocol is PROTOCOL_NFS.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.nfs_exports
	NfsExports []VolumeConfig_NfsExport `json:"nfsExports,omitempty"`

	// User note field, it can be used by customers to add additional information
	//  for the BMS Ops team .
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.user_note
	UserNote *string `json:"userNote,omitempty"`

	// The GCP service of the storage volume. Available gcp_service are in
	//  https://cloud.google.com/bare-metal/docs/bms-planning.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.gcp_service
	GcpService *string `json:"gcpService,omitempty"`

	// Performance tier of the Volume.
	//  Default is SHARED.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.performance_tier
	PerformanceTier *string `json:"performanceTier,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig.LunRange
type VolumeConfig_LunRange struct {
	// Number of LUNs to create.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.LunRange.quantity
	Quantity *int32 `json:"quantity,omitempty"`

	// The requested size of each LUN, in GB.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.LunRange.size_gb
	SizeGB *int32 `json:"sizeGB,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport
type VolumeConfig_NfsExport struct {
	// Network to use to publish the export.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.network_id
	NetworkID *string `json:"networkID,omitempty"`

	// Either a single machine, identified by an ID, or a comma-separated
	//  list of machine IDs.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.machine_id
	MachineID *string `json:"machineID,omitempty"`

	// A CIDR range.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.cidr
	Cidr *string `json:"cidr,omitempty"`

	// Export permissions.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.permissions
	Permissions *string `json:"permissions,omitempty"`

	// Disable root squashing, which is a feature of NFS.
	//  Root squash is a special mapping of the remote superuser (root) identity
	//  when using identity authentication.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.no_root_squash
	NoRootSquash *bool `json:"noRootSquash,omitempty"`

	// Allow the setuid flag.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.allow_suid
	AllowSuid *bool `json:"allowSuid,omitempty"`

	// Allow dev flag in NfsShare AllowedClientsRequest.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.NfsExport.allow_dev
	AllowDev *bool `json:"allowDev,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.VolumeConfig
type VolumeConfigObservedState struct {
	// Output only. The name of the volume config.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.VolumeConfig.name
	Name *string `json:"name,omitempty"`
}
