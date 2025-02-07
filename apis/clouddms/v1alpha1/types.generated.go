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


// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspaceInfo
type ConversionWorkspaceInfo struct {
	// The resource name (URI) of the conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspaceInfo.name
	Name *string `json:"name,omitempty"`

	// The commit ID of the conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspaceInfo.commit_id
	CommitID *string `json:"commitID,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DatabaseType
type DatabaseType struct {
	// The database provider.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseType.provider
	Provider *string `json:"provider,omitempty"`

	// The database engine.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseType.engine
	Engine *string `json:"engine,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob
type MigrationJob struct {
	// The name (URI) of this migration job resource, in the form of:
	//  projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.name
	Name *string `json:"name,omitempty"`

	// The resource labels for migration job to use to annotate any related
	//  underlying resources such as Compute Engine VMs. An object containing a
	//  list of "key": "value" pairs.
	//
	//  Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The migration job display name.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The current migration job state.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.state
	State *string `json:"state,omitempty"`

	// Required. The migration job type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.type
	Type *string `json:"type,omitempty"`

	// The path to the dump file in Google Cloud Storage,
	//  in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
	//  This field and the "dump_flags" field are mutually exclusive.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.dump_path
	DumpPath *string `json:"dumpPath,omitempty"`

	// The initial dump flags.
	//  This field and the "dump_path" field are mutually exclusive.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.dump_flags
	DumpFlags *MigrationJob_DumpFlags `json:"dumpFlags,omitempty"`

	// Required. The resource name (URI) of the source connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.source
	Source *string `json:"source,omitempty"`

	// Required. The resource name (URI) of the destination connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.destination
	Destination *string `json:"destination,omitempty"`

	// The details needed to communicate to the source over Reverse SSH
	//  tunnel connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.reverse_ssh_connectivity
	ReverseSSHConnectivity *ReverseSshConnectivity `json:"reverseSSHConnectivity,omitempty"`

	// The details of the VPC network that the source database is located in.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.vpc_peering_connectivity
	VpcPeeringConnectivity *VpcPeeringConnectivity `json:"vpcPeeringConnectivity,omitempty"`

	// static ip connectivity data (default, no additional details needed).
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.static_ip_connectivity
	StaticIPConnectivity *StaticIpConnectivity `json:"staticIPConnectivity,omitempty"`

	// The database engine type and provider of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.source_database
	SourceDatabase *DatabaseType `json:"sourceDatabase,omitempty"`

	// The database engine type and provider of the destination.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.destination_database
	DestinationDatabase *DatabaseType `json:"destinationDatabase,omitempty"`

	// The conversion workspace used by the migration.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.conversion_workspace
	ConversionWorkspace *ConversionWorkspaceInfo `json:"conversionWorkspace,omitempty"`

	// This field can be used to select the entities to migrate as part of
	//  the migration job. It uses AIP-160 notation to select a subset of the
	//  entities configured on the associated conversion-workspace. This field
	//  should not be set on migration-jobs that are not associated with a
	//  conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.filter
	Filter *string `json:"filter,omitempty"`

	// The CMEK (customer-managed encryption key) fully qualified key name used
	//  for the migration job.
	//  This field supports all migration jobs types except for:
	//  * Mysql to Mysql (use the cmek field in the cloudsql connection profile
	//  instead).
	//  * PostrgeSQL to PostgreSQL (use the cmek field in the cloudsql
	//  connection profile instead).
	//  * PostgreSQL to AlloyDB (use the kms_key_name field in the alloydb
	//  connection profile instead).
	//  Each Cloud CMEK key has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.cmek_key_name
	CmekKeyName *string `json:"cmekKeyName,omitempty"`

	// Optional. Data dump parallelism settings used by the migration.
	//  Currently applicable only for MySQL to Cloud SQL for MySQL migrations only.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.performance_config
	PerformanceConfig *MigrationJob_PerformanceConfig `json:"performanceConfig,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.DumpFlag
type MigrationJob_DumpFlag struct {
	// The name of the flag
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlag.name
	Name *string `json:"name,omitempty"`

	// The value of the flag.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlag.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.DumpFlags
type MigrationJob_DumpFlags struct {
	// The flags for the initial dump.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.DumpFlags.dump_flags
	DumpFlags []MigrationJob_DumpFlag `json:"dumpFlags,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob.PerformanceConfig
type MigrationJob_PerformanceConfig struct {
	// Initial dump parallelism level.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.PerformanceConfig.dump_parallel_level
	DumpParallelLevel *string `json:"dumpParallelLevel,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ReverseSshConnectivity
type ReverseSshConnectivity struct {
	// Required. The IP of the virtual machine (Compute Engine) used as the
	//  bastion server for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vm_ip
	VmIP *string `json:"vmIP,omitempty"`

	// Required. The forwarding port of the virtual machine (Compute Engine) used
	//  as the bastion server for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vm_port
	VmPort *int32 `json:"vmPort,omitempty"`

	// The name of the virtual machine (Compute Engine) used as the bastion server
	//  for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vm
	Vm *string `json:"vm,omitempty"`

	// The name of the VPC to peer with the Cloud SQL private network.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vpc
	Vpc *string `json:"vpc,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.StaticIpConnectivity
type StaticIpConnectivity struct {
}

// +kcc:proto=google.cloud.clouddms.v1.VpcPeeringConnectivity
type VpcPeeringConnectivity struct {
	// The name of the VPC network to peer with the Cloud SQL private network.
	// +kcc:proto:field=google.cloud.clouddms.v1.VpcPeeringConnectivity.vpc
	Vpc *string `json:"vpc,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MigrationJob
type MigrationJobObservedState struct {
	// Output only. The timestamp when the migration job resource was created.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the migration job resource was last
	//  updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current migration job phase.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.phase
	Phase *string `json:"phase,omitempty"`

	// Output only. The duration of the migration job (in seconds). A duration in
	//  seconds with up to nine fractional digits, terminated by 's'. Example:
	//  "3.5s".
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.duration
	Duration *string `json:"duration,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. If the migration job is completed, the time when it was
	//  completed.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.end_time
	EndTime *string `json:"endTime,omitempty"`
}
