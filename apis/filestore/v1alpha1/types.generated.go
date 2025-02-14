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

// +kcc:proto=google.cloud.filestore.v1.NetworkConfig
type NetworkConfig struct {
	// The name of the Google Compute Engine
	//  [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the
	//  instance is connected.
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.network
	Network *string `json:"network,omitempty"`

	// Internet protocol versions for which the instance has IP addresses
	//  assigned. For this version, only MODE_IPV4 is supported.
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.modes
	Modes []string `json:"modes,omitempty"`

	// Optional, reserved_ip_range can have one of the following two types of
	//  values.
	//
	//  * CIDR range value when using DIRECT_PEERING connect mode.
	//  * [Allocated IP address
	//  range](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
	//  when using PRIVATE_SERVICE_ACCESS connect mode.
	//
	//  When the name of an allocated IP address range is specified, it must be one
	//  of the ranges associated with the private service access connection.
	//  When specified as a direct CIDR value, it must be a /29 CIDR block for
	//  Basic tier, a /24 CIDR block for High Scale tier, or a /26 CIDR block for
	//  Enterprise tier in one of the [internal IP address
	//  ranges](https://www.arin.net/reference/research/statistics/address_filters/)
	//  that identifies the range of IP addresses reserved for this instance. For
	//  example, 10.0.0.0/29, 192.168.0.0/24 or 192.168.0.0/26, respectively. The
	//  range you specify can't overlap with either existing subnets or assigned IP
	//  address ranges for other Filestore instances in the selected VPC
	//  network.
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`

	// The network connect mode of the Filestore instance.
	//  If not provided, the connect mode defaults to DIRECT_PEERING.
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.connect_mode
	ConnectMode *string `json:"connectMode,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1.NfsExportOptions
type NfsExportOptions struct {
	// List of either an IPv4 addresses in the format
	//  `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format
	//  `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the
	//  file share.
	//  Overlapping IP ranges are not allowed, both within and across
	//  NfsExportOptions. An error will be returned.
	//  The limit is 64 IP ranges/addresses for each FileShareConfig among all
	//  NfsExportOptions.
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.ip_ranges
	IPRanges []string `json:"ipRanges,omitempty"`

	// Either READ_ONLY, for allowing only read requests on the exported
	//  directory, or READ_WRITE, for allowing both read and write requests.
	//  The default is READ_WRITE.
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.access_mode
	AccessMode *string `json:"accessMode,omitempty"`

	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory,
	//  or ROOT_SQUASH, for not allowing root access. The default is
	//  NO_ROOT_SQUASH.
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.squash_mode
	SquashMode *string `json:"squashMode,omitempty"`

	// An integer representing the anonymous user id with a default value of
	//  65534.
	//  Anon_uid may only be set with squash_mode of ROOT_SQUASH.  An error will be
	//  returned if this field is specified for other squash_mode settings.
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.anon_uid
	AnonUid *int64 `json:"anonUid,omitempty"`

	// An integer representing the anonymous group id with a default value of
	//  65534.
	//  Anon_gid may only be set with squash_mode of ROOT_SQUASH.  An error will be
	//  returned if this field is specified for other squash_mode settings.
	// +kcc:proto:field=google.cloud.filestore.v1.NfsExportOptions.anon_gid
	AnonGid *int64 `json:"anonGid,omitempty"`
}

// +kcc:proto=google.cloud.filestore.v1.NetworkConfig
type NetworkConfigObservedState struct {
	// Output only. IPv4 addresses in the format
	//  `{octet1}.{octet2}.{octet3}.{octet4}` or IPv6 addresses in the format
	//  `{block1}:{block2}:{block3}:{block4}:{block5}:{block6}:{block7}:{block8}`.
	// +kcc:proto:field=google.cloud.filestore.v1.NetworkConfig.ip_addresses
	IPAddresses []string `json:"ipAddresses,omitempty"`
}
