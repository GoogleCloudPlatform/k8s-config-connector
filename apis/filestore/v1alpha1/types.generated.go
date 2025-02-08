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


// +kcc:proto=google.cloud.filestore.v1beta1.NfsExportOptions
type NfsExportOptions struct {
	// List of either an IPv4 addresses in the format
	//  `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format
	//  `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the
	//  file share.
	//  Overlapping IP ranges are not allowed, both within and across
	//  NfsExportOptions. An error will be returned.
	//  The limit is 64 IP ranges/addresses for each FileShareConfig among all
	//  NfsExportOptions.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.NfsExportOptions.ip_ranges
	IPRanges []string `json:"ipRanges,omitempty"`

	// Either READ_ONLY, for allowing only read requests on the exported
	//  directory, or READ_WRITE, for allowing both read and write requests.
	//  The default is READ_WRITE.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.NfsExportOptions.access_mode
	AccessMode *string `json:"accessMode,omitempty"`

	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory,
	//  or ROOT_SQUASH, for not allowing root access. The default is
	//  NO_ROOT_SQUASH.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.NfsExportOptions.squash_mode
	SquashMode *string `json:"squashMode,omitempty"`

	// An integer representing the anonymous user id with a default value of
	//  65534.
	//  Anon_uid may only be set with squash_mode of ROOT_SQUASH.  An error will be
	//  returned if this field is specified for other squash_mode settings.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.NfsExportOptions.anon_uid
	AnonUid *int64 `json:"anonUid,omitempty"`

	// An integer representing the anonymous group id with a default value of
	//  65534.
	//  Anon_gid may only be set with squash_mode of ROOT_SQUASH.  An error will be
	//  returned if this field is specified for other squash_mode settings.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.NfsExportOptions.anon_gid
	AnonGid *int64 `json:"anonGid,omitempty"`

	// The security flavors allowed for mount operations.
	//  The default is AUTH_SYS.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.NfsExportOptions.security_flavors
	SecurityFlavors []string `json:"securityFlavors,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1beta1.Share
type Share struct {

	// The mount name of the share. Must be 63 characters or less and consist of
	//  uppercase or lowercase letters, numbers, and underscores.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.mount_name
	MountName *string `json:"mountName,omitempty"`

	// A description of the share with 2048 characters or less. Requests with
	//  longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.description
	Description *string `json:"description,omitempty"`

	// File share capacity in gigabytes (GB). Filestore defines 1 GB as
	//  1024^3 bytes. Must be greater than 0.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.capacity_gb
	CapacityGB *int64 `json:"capacityGB,omitempty"`

	// Nfs Export Options.
	//  There is a limit of 10 export options per file share.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.nfs_export_options
	NfsExportOptions []NfsExportOptions `json:"nfsExportOptions,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Full name of the Cloud Filestore Backup resource that this
	//  Share is restored from, in the format of
	//  projects/{project_id}/locations/{location_id}/backups/{backup_id}.
	//  Empty, if the Share is created from scratch and not restored from a
	//  backup.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.backup
	Backup *string `json:"backup,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1beta1.Share
type ShareObservedState struct {
	// Output only. The resource name of the share, in the format
	//  `projects/{project_id}/locations/{location_id}/instances/{instance_id}/shares/{share_id}`.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.name
	Name *string `json:"name,omitempty"`

	// Output only. The share state.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the share was created.
	// +kcc:proto:field=google.cloud.filestore.v1beta1.Share.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
