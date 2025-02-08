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

// +kcc:proto=google.cloud.migrationcenter.v1.Report
type Report struct {

	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Free-text description.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.description
	Description *string `json:"description,omitempty"`

	// Report type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.type
	Type *string `json:"type,omitempty"`

	// Report creation state.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary
type ReportSummary struct {
	// Aggregate statistics for all the assets across all the groups.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.all_assets_stats
	AllAssetsStats *ReportSummary_AssetAggregateStats `json:"allAssetsStats,omitempty"`

	// Findings for each Group included in this report.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.group_findings
	GroupFindings []ReportSummary_GroupFinding `json:"groupFindings,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats
type ReportSummary_AssetAggregateStats struct {
	// Sum of the memory in bytes of all the assets in this collection.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.total_memory_bytes
	TotalMemoryBytes *int64 `json:"totalMemoryBytes,omitempty"`

	// Sum of persistent storage in bytes of all the assets in this collection.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.total_storage_bytes
	TotalStorageBytes *int64 `json:"totalStorageBytes,omitempty"`

	// Sum of the CPU core count of all the assets in this collection.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.total_cores
	TotalCores *int64 `json:"totalCores,omitempty"`

	// Count of the number of unique assets in this collection.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.total_assets
	TotalAssets *int64 `json:"totalAssets,omitempty"`

	// Total memory split into Used/Free buckets.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.memory_utilization_chart
	MemoryUtilizationChart *ReportSummary_UtilizationChartData `json:"memoryUtilizationChart,omitempty"`

	// Total memory split into Used/Free buckets.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.storage_utilization_chart
	StorageUtilizationChart *ReportSummary_UtilizationChartData `json:"storageUtilizationChart,omitempty"`

	// Count of assets grouped by Operating System families.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.operating_system
	OperatingSystem *ReportSummary_ChartData `json:"operatingSystem,omitempty"`

	// Histogram showing a distribution of CPU core counts.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.core_count_histogram
	CoreCountHistogram *ReportSummary_HistogramChartData `json:"coreCountHistogram,omitempty"`

	// Histogram showing a distribution of memory sizes.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.memory_bytes_histogram
	MemoryBytesHistogram *ReportSummary_HistogramChartData `json:"memoryBytesHistogram,omitempty"`

	// Histogram showing a distribution of memory sizes.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.AssetAggregateStats.storage_bytes_histogram
	StorageBytesHistogram *ReportSummary_HistogramChartData `json:"storageBytesHistogram,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.ChartData
type ReportSummary_ChartData struct {
	// Each data point in the chart is represented as a name-value pair
	//  with the name being the x-axis label, and the value being the y-axis
	//  value.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ChartData.data_points
	DataPoints []ReportSummary_ChartData_DataPoint `json:"dataPoints,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.ChartData.DataPoint
type ReportSummary_ChartData_DataPoint struct {
	// The X-axis label for this data point.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ChartData.DataPoint.label
	Label *string `json:"label,omitempty"`

	// The Y-axis value for this data point.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ChartData.DataPoint.value
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.ComputeEngineFinding
type ReportSummary_ComputeEngineFinding struct {
	// Set of regions in which the assets were allocated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ComputeEngineFinding.allocated_regions
	AllocatedRegions []string `json:"allocatedRegions,omitempty"`

	// Count of assets which were allocated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ComputeEngineFinding.allocated_asset_count
	AllocatedAssetCount *int64 `json:"allocatedAssetCount,omitempty"`

	// Distribution of assets based on the Machine Series.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ComputeEngineFinding.machine_series_allocations
	MachineSeriesAllocations []ReportSummary_MachineSeriesAllocation `json:"machineSeriesAllocations,omitempty"`

	// Set of disk types allocated to assets.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.ComputeEngineFinding.allocated_disk_types
	AllocatedDiskTypes []string `json:"allocatedDiskTypes,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.GroupFinding
type ReportSummary_GroupFinding struct {
	// Display Name for the Group.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupFinding.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description for the Group.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupFinding.description
	Description *string `json:"description,omitempty"`

	// Summary statistics for all the assets in this group.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupFinding.asset_aggregate_stats
	AssetAggregateStats *ReportSummary_AssetAggregateStats `json:"assetAggregateStats,omitempty"`

	// This field is deprecated, do not rely on it having a value.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupFinding.overlapping_asset_count
	OverlappingAssetCount *int64 `json:"overlappingAssetCount,omitempty"`

	// Findings for each of the PreferenceSets for this group.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupFinding.preference_set_findings
	PreferenceSetFindings []ReportSummary_GroupPreferenceSetFinding `json:"preferenceSetFindings,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding
type ReportSummary_GroupPreferenceSetFinding struct {
	// Display Name of the Preference Set
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description for the Preference Set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.description
	Description *string `json:"description,omitempty"`

	// A set of preferences that applies to all machines in the context.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.machine_preferences
	MachinePreferences *VirtualMachinePreferences `json:"machinePreferences,omitempty"`

	// Total monthly cost for this preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.monthly_cost_total
	MonthlyCostTotal *Money `json:"monthlyCostTotal,omitempty"`

	// Compute monthly cost for this preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.monthly_cost_compute
	MonthlyCostCompute *Money `json:"monthlyCostCompute,omitempty"`

	// Licensing monthly cost for this preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.monthly_cost_os_license
	MonthlyCostOsLicense *Money `json:"monthlyCostOsLicense,omitempty"`

	// Network Egress monthly cost for this preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.monthly_cost_network_egress
	MonthlyCostNetworkEgress *Money `json:"monthlyCostNetworkEgress,omitempty"`

	// Storage monthly cost for this preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.monthly_cost_storage
	MonthlyCostStorage *Money `json:"monthlyCostStorage,omitempty"`

	// Miscellaneous monthly cost for this preference set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.monthly_cost_other
	MonthlyCostOther *Money `json:"monthlyCostOther,omitempty"`

	// A set of findings that applies to Compute Engine machines in the input.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.compute_engine_finding
	ComputeEngineFinding *ReportSummary_ComputeEngineFinding `json:"computeEngineFinding,omitempty"`

	// A set of findings that applies to VMWare machines in the input.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.vmware_engine_finding
	VmwareEngineFinding *ReportSummary_VmwareEngineFinding `json:"vmwareEngineFinding,omitempty"`

	// A set of findings that applies to Sole-Tenant machines in the input.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.GroupPreferenceSetFinding.sole_tenant_finding
	SoleTenantFinding *ReportSummary_SoleTenantFinding `json:"soleTenantFinding,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.HistogramChartData
type ReportSummary_HistogramChartData struct {
	// Buckets in the histogram.
	//  There will be `n+1` buckets matching `n` lower bounds in the request.
	//  The first bucket will be from -infinity to the first bound.
	//  Subsequent buckets will be between one bound and the next.
	//  The final bucket will be from the final bound to infinity.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.HistogramChartData.buckets
	Buckets []ReportSummary_HistogramChartData_Bucket `json:"buckets,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.HistogramChartData.Bucket
type ReportSummary_HistogramChartData_Bucket struct {
	// Lower bound - inclusive.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.HistogramChartData.Bucket.lower_bound
	LowerBound *int64 `json:"lowerBound,omitempty"`

	// Upper bound - exclusive.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.HistogramChartData.Bucket.upper_bound
	UpperBound *int64 `json:"upperBound,omitempty"`

	// Count of items in the bucket.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.HistogramChartData.Bucket.count
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.MachineSeriesAllocation
type ReportSummary_MachineSeriesAllocation struct {
	// The Machine Series (e.g. "E2", "N2")
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.MachineSeriesAllocation.machine_series
	MachineSeries *MachineSeries `json:"machineSeries,omitempty"`

	// Count of assets allocated to this machine series.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.MachineSeriesAllocation.allocated_asset_count
	AllocatedAssetCount *int64 `json:"allocatedAssetCount,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantFinding
type ReportSummary_SoleTenantFinding struct {
	// Set of regions in which the assets are allocated
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantFinding.allocated_regions
	AllocatedRegions []string `json:"allocatedRegions,omitempty"`

	// Count of assets which are allocated
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantFinding.allocated_asset_count
	AllocatedAssetCount *int64 `json:"allocatedAssetCount,omitempty"`

	// Set of per-nodetype allocation records
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantFinding.node_allocations
	NodeAllocations []ReportSummary_SoleTenantNodeAllocation `json:"nodeAllocations,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantNodeAllocation
type ReportSummary_SoleTenantNodeAllocation struct {
	// Sole Tenant node type, e.g. "m3-node-128-3904"
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantNodeAllocation.node
	Node *SoleTenantNodeType `json:"node,omitempty"`

	// Count of this node type to be provisioned
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantNodeAllocation.node_count
	NodeCount *int64 `json:"nodeCount,omitempty"`

	// Count of assets allocated to these nodes
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.SoleTenantNodeAllocation.allocated_asset_count
	AllocatedAssetCount *int64 `json:"allocatedAssetCount,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.UtilizationChartData
type ReportSummary_UtilizationChartData struct {
	// Aggregate value which falls into the "Used" bucket.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.UtilizationChartData.used
	Used *int64 `json:"used,omitempty"`

	// Aggregate value which falls into the "Free" bucket.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.UtilizationChartData.free
	Free *int64 `json:"free,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.VmwareEngineFinding
type ReportSummary_VmwareEngineFinding struct {
	// Set of regions in which the assets were allocated
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareEngineFinding.allocated_regions
	AllocatedRegions []string `json:"allocatedRegions,omitempty"`

	// Count of assets which are allocated
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareEngineFinding.allocated_asset_count
	AllocatedAssetCount *int64 `json:"allocatedAssetCount,omitempty"`

	// Set of per-nodetype allocation records
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareEngineFinding.node_allocations
	NodeAllocations []ReportSummary_VmwareNodeAllocation `json:"nodeAllocations,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.VmwareNode
type ReportSummary_VmwareNode struct {
	// Code to identify VMware Engine node series, e.g. "ve1-standard-72". Based
	//  on the displayName of
	//  cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.nodeTypes
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareNode.code
	Code *string `json:"code,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportSummary.VmwareNodeAllocation
type ReportSummary_VmwareNodeAllocation struct {
	// VMWare node type, e.g. "ve1-standard-72"
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareNodeAllocation.vmware_node
	VmwareNode *ReportSummary_VmwareNode `json:"vmwareNode,omitempty"`

	// Count of this node type to be provisioned
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareNodeAllocation.node_count
	NodeCount *int64 `json:"nodeCount,omitempty"`

	// Count of assets allocated to these nodes
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportSummary.VmwareNodeAllocation.allocated_asset_count
	AllocatedAssetCount *int64 `json:"allocatedAssetCount,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.SoleTenancyPreferences
type SoleTenancyPreferences struct {
	// CPU overcommit ratio.
	//  Acceptable values are between 1.0 and 2.0 inclusive.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.SoleTenancyPreferences.cpu_overcommit_ratio
	CpuOvercommitRatio *float64 `json:"cpuOvercommitRatio,omitempty"`

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
	CpuOvercommitRatio *float64 `json:"cpuOvercommitRatio,omitempty"`

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

// +kcc:proto=google.type.Money
type Money struct {
	// The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`

	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.Report
type ReportObservedState struct {
	// Output only. Name of resource.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Summary view of the Report.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.summary
	Summary *ReportSummary `json:"summary,omitempty"`
}
