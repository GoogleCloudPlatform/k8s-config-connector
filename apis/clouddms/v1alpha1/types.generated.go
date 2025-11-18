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

// +generated:types
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1
// resource: CloudDMSConversionWorkspace:ConversionWorkspace
// resource: CloudDMSPrivateConnection:PrivateConnection
// resource: CloudDMSMigrationJob:MigrationJob

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

// +kcc:proto=google.cloud.clouddms.v1.DatabaseEngineInfo
type DatabaseEngineInfo struct {
	// Required. Engine type.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.engine
	Engine *string `json:"engine,omitempty"`

	// Required. Engine named version, for example 12.c.1.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.version
	Version *string `json:"version,omitempty"`
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
type ReverseSSHConnectivity struct {
	// Required. The IP of the virtual machine (Compute Engine) used as the
	//  bastion server for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vm_ip
	VMIP *string `json:"vmIP,omitempty"`

	// Required. The forwarding port of the virtual machine (Compute Engine) used
	//  as the bastion server for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vm_port
	VMPort *int32 `json:"vmPort,omitempty"`

	// The name of the virtual machine (Compute Engine) used as the bastion server
	//  for the SSH tunnel.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vm
	VM *string `json:"vm,omitempty"`

	// The name of the VPC to peer with the Cloud SQL private network.
	// +kcc:proto:field=google.cloud.clouddms.v1.ReverseSshConnectivity.vpc
	VPC *string `json:"vpc,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.StaticIpConnectivity
type StaticIPConnectivity struct {
}

// +kcc:proto=google.cloud.clouddms.v1.VpcPeeringConnectivity
type VPCPeeringConnectivity struct {
	// The name of the VPC network to peer with the Cloud SQL private network.
	// +kcc:proto:field=google.cloud.clouddms.v1.VpcPeeringConnectivity.vpc
	VPC *string `json:"vpc,omitempty"`
}
