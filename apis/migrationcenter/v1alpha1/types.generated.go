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

// +kcc:proto=google.cloud.migrationcenter.v1.ComputeEnginePreferences
type ComputeEnginePreferences struct {
	// Preferences concerning the machine types to consider on Compute Engine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ComputeEnginePreferences.machine_preferences
	MachinePreferences *MachinePreferences `json:"machinePreferences,omitempty"`

	// License type to consider when calculating costs for virtual machine
	//  insights and recommendations. If unspecified, costs are calculated
	//  based on the default licensing plan.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ComputeEnginePreferences.license_type
	LicenseType *string `json:"licenseType,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MachinePreferences
type MachinePreferences struct {
	// Compute Engine machine series to consider for insights and recommendations.
	//  If empty, no restriction is applied on the machine series.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachinePreferences.allowed_machine_series
	AllowedMachineSeries []MachineSeries `json:"allowedMachineSeries,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.MachineSeries
type MachineSeries struct {
	// Code to identify a Compute Engine machine series. Consult
	//  https://cloud.google.com/compute/docs/machine-resource#machine_type_comparison
	//  for more details on the available series.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.MachineSeries.code
	Code *string `json:"code,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.PreferenceSet
type PreferenceSet struct {

	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PreferenceSet.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description of the preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PreferenceSet.description
	Description *string `json:"description,omitempty"`

	// A set of preferences that applies to all virtual machines in the context.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PreferenceSet.virtual_machine_preferences
	VirtualMachinePreferences *VirtualMachinePreferences `json:"virtualMachinePreferences,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.RegionPreferences
type RegionPreferences struct {
	// A list of preferred regions,
	//  ordered by the most preferred region first.
	//  Set only valid Google Cloud region names.
	//  See https://cloud.google.com/compute/docs/regions-zones
	//  for available regions.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.RegionPreferences.preferred_regions
	PreferredRegions []string `json:"preferredRegions,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.SoleTenancyPreferences
type SoleTenancyPreferences struct {
	// CPU overcommit ratio.
	//  Acceptable values are between 1.0 and 2.0 inclusive.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.SoleTenancyPreferences.cpu_overcommit_ratio
	CPUOvercommitRatio *float64 `json:"cpuOvercommitRatio,omitempty"`

	// Sole Tenancy nodes maintenance policy.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.SoleTenancyPreferences.host_maintenance_policy
	HostMaintenancePolicy *string `json:"hostMaintenancePolicy,omitempty"`

	// Commitment plan to consider when calculating costs for virtual machine
	//  insights and recommendations.
	//  If you are unsure which value to set, a 3 year commitment plan is often a
	//  good value to start with.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.SoleTenancyPreferences.commitment_plan
	CommitmentPlan *string `json:"commitmentPlan,omitempty"`

	// A list of sole tenant node types.
	//  An empty list means that all possible node types will be considered.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.SoleTenancyPreferences.node_types
	NodeTypes []SoleTenantNodeType `json:"nodeTypes,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.SoleTenantNodeType
type SoleTenantNodeType struct {
	// Name of the Sole Tenant node. Consult
	//  https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes
	// +kcc:proto:field=google.cloud.migrationcenter.v1.SoleTenantNodeType.node_name
	NodeName *string `json:"nodeName,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.VirtualMachinePreferences
type VirtualMachinePreferences struct {
	// Target product for assets using this preference set.
	//  Specify either target product or business goal, but
	//  not both.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.target_product
	TargetProduct *string `json:"targetProduct,omitempty"`

	// Region preferences for assets using this preference set.
	//  If you are unsure which value to set, the migration service API region is
	//  often a good value to start with.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.region_preferences
	RegionPreferences *RegionPreferences `json:"regionPreferences,omitempty"`

	// Commitment plan to consider when calculating costs for virtual machine
	//  insights and recommendations.
	//  If you are unsure which value to set, a 3 year commitment plan is often a
	//  good value to start with.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.commitment_plan
	CommitmentPlan *string `json:"commitmentPlan,omitempty"`

	// Sizing optimization strategy specifies the preferred strategy used when
	//  extrapolating usage data to calculate insights and recommendations for a
	//  virtual machine.
	//  If you are unsure which value to set, a moderate sizing optimization
	//  strategy is often a good value to start with.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.sizing_optimization_strategy
	SizingOptimizationStrategy *string `json:"sizingOptimizationStrategy,omitempty"`

	// Compute Engine preferences concern insights and recommendations for Compute
	//  Engine target.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.compute_engine_preferences
	ComputeEnginePreferences *ComputeEnginePreferences `json:"computeEnginePreferences,omitempty"`

	// Preferences concerning insights and recommendations for
	//  Google Cloud VMware Engine.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.vmware_engine_preferences
	VmwareEnginePreferences *VmwareEnginePreferences `json:"vmwareEnginePreferences,omitempty"`

	// Preferences concerning Sole Tenant nodes and virtual machines.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VirtualMachinePreferences.sole_tenancy_preferences
	SoleTenancyPreferences *SoleTenancyPreferences `json:"soleTenancyPreferences,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.VmwareEnginePreferences
type VmwareEnginePreferences struct {
	// CPU overcommit ratio.
	//  Acceptable values are between 1.0 and 8.0, with 0.1 increment.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareEnginePreferences.cpu_overcommit_ratio
	CPUOvercommitRatio *float64 `json:"cpuOvercommitRatio,omitempty"`

	// Memory overcommit ratio.
	//  Acceptable values are 1.0, 1.25, 1.5, 1.75 and 2.0.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareEnginePreferences.memory_overcommit_ratio
	MemoryOvercommitRatio *float64 `json:"memoryOvercommitRatio,omitempty"`

	// The Deduplication and Compression ratio is based on the logical (Used
	//  Before) space required to store data before applying deduplication and
	//  compression, in relation to the physical (Used After) space required after
	//  applying deduplication and compression. Specifically, the ratio is the Used
	//  Before space divided by the Used After space. For example, if the Used
	//  Before space is 3 GB, but the physical Used After space is 1 GB, the
	//  deduplication and compression ratio is 3x. Acceptable values are
	//  between 1.0 and 4.0.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareEnginePreferences.storage_deduplication_compression_ratio
	StorageDeduplicationCompressionRatio *float64 `json:"storageDeduplicationCompressionRatio,omitempty"`

	// Commitment plan to consider when calculating costs for virtual machine
	//  insights and recommendations.
	//  If you are unsure which value to set, a 3 year commitment plan is often a
	//  good value to start with.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.VmwareEnginePreferences.commitment_plan
	CommitmentPlan *string `json:"commitmentPlan,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.PreferenceSet
type PreferenceSetObservedState struct {
	// Output only. Name of the preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PreferenceSet.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the preference set was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PreferenceSet.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the preference set was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.PreferenceSet.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
