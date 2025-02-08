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

package migrationcenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ComputeEnginePreferences_FromProto(mapCtx *direct.MapContext, in *pb.ComputeEnginePreferences) *krm.ComputeEnginePreferences {
	if in == nil {
		return nil
	}
	out := &krm.ComputeEnginePreferences{}
	out.MachinePreferences = MachinePreferences_FromProto(mapCtx, in.GetMachinePreferences())
	out.LicenseType = direct.Enum_FromProto(mapCtx, in.GetLicenseType())
	return out
}
func ComputeEnginePreferences_ToProto(mapCtx *direct.MapContext, in *krm.ComputeEnginePreferences) *pb.ComputeEnginePreferences {
	if in == nil {
		return nil
	}
	out := &pb.ComputeEnginePreferences{}
	out.MachinePreferences = MachinePreferences_ToProto(mapCtx, in.MachinePreferences)
	out.LicenseType = direct.Enum_ToProto[pb.LicenseType](mapCtx, in.LicenseType)
	return out
}
func MachinePreferences_FromProto(mapCtx *direct.MapContext, in *pb.MachinePreferences) *krm.MachinePreferences {
	if in == nil {
		return nil
	}
	out := &krm.MachinePreferences{}
	out.AllowedMachineSeries = direct.Slice_FromProto(mapCtx, in.AllowedMachineSeries, MachineSeries_FromProto)
	return out
}
func MachinePreferences_ToProto(mapCtx *direct.MapContext, in *krm.MachinePreferences) *pb.MachinePreferences {
	if in == nil {
		return nil
	}
	out := &pb.MachinePreferences{}
	out.AllowedMachineSeries = direct.Slice_ToProto(mapCtx, in.AllowedMachineSeries, MachineSeries_ToProto)
	return out
}
func MachineSeries_FromProto(mapCtx *direct.MapContext, in *pb.MachineSeries) *krm.MachineSeries {
	if in == nil {
		return nil
	}
	out := &krm.MachineSeries{}
	out.Code = direct.LazyPtr(in.GetCode())
	return out
}
func MachineSeries_ToProto(mapCtx *direct.MapContext, in *krm.MachineSeries) *pb.MachineSeries {
	if in == nil {
		return nil
	}
	out := &pb.MachineSeries{}
	out.Code = direct.ValueOf(in.Code)
	return out
}
func MigrationcenterReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Report) *krm.MigrationcenterReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterReportObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: State
	// MISSING: Summary
	return out
}
func MigrationcenterReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterReportObservedState) *pb.Report {
	if in == nil {
		return nil
	}
	out := &pb.Report{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: State
	// MISSING: Summary
	return out
}
func MigrationcenterReportSpec_FromProto(mapCtx *direct.MapContext, in *pb.Report) *krm.MigrationcenterReportSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterReportSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: State
	// MISSING: Summary
	return out
}
func MigrationcenterReportSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterReportSpec) *pb.Report {
	if in == nil {
		return nil
	}
	out := &pb.Report{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: State
	// MISSING: Summary
	return out
}
func RegionPreferences_FromProto(mapCtx *direct.MapContext, in *pb.RegionPreferences) *krm.RegionPreferences {
	if in == nil {
		return nil
	}
	out := &krm.RegionPreferences{}
	out.PreferredRegions = in.PreferredRegions
	return out
}
func RegionPreferences_ToProto(mapCtx *direct.MapContext, in *krm.RegionPreferences) *pb.RegionPreferences {
	if in == nil {
		return nil
	}
	out := &pb.RegionPreferences{}
	out.PreferredRegions = in.PreferredRegions
	return out
}
func Report_FromProto(mapCtx *direct.MapContext, in *pb.Report) *krm.Report {
	if in == nil {
		return nil
	}
	out := &krm.Report{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Summary
	return out
}
func Report_ToProto(mapCtx *direct.MapContext, in *krm.Report) *pb.Report {
	if in == nil {
		return nil
	}
	out := &pb.Report{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.Report_Type](mapCtx, in.Type)
	out.State = direct.Enum_ToProto[pb.Report_State](mapCtx, in.State)
	// MISSING: Summary
	return out
}
func ReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Report) *krm.ReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReportObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: State
	out.Summary = ReportSummary_FromProto(mapCtx, in.GetSummary())
	return out
}
func ReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReportObservedState) *pb.Report {
	if in == nil {
		return nil
	}
	out := &pb.Report{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: State
	out.Summary = ReportSummary_ToProto(mapCtx, in.Summary)
	return out
}
func ReportSummary_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary) *krm.ReportSummary {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary{}
	out.AllAssetsStats = ReportSummary_AssetAggregateStats_FromProto(mapCtx, in.GetAllAssetsStats())
	out.GroupFindings = direct.Slice_FromProto(mapCtx, in.GroupFindings, ReportSummary_GroupFinding_FromProto)
	return out
}
func ReportSummary_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary) *pb.ReportSummary {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary{}
	out.AllAssetsStats = ReportSummary_AssetAggregateStats_ToProto(mapCtx, in.AllAssetsStats)
	out.GroupFindings = direct.Slice_ToProto(mapCtx, in.GroupFindings, ReportSummary_GroupFinding_ToProto)
	return out
}
func ReportSummary_AssetAggregateStats_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_AssetAggregateStats) *krm.ReportSummary_AssetAggregateStats {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_AssetAggregateStats{}
	out.TotalMemoryBytes = direct.LazyPtr(in.GetTotalMemoryBytes())
	out.TotalStorageBytes = direct.LazyPtr(in.GetTotalStorageBytes())
	out.TotalCores = direct.LazyPtr(in.GetTotalCores())
	out.TotalAssets = direct.LazyPtr(in.GetTotalAssets())
	out.MemoryUtilizationChart = ReportSummary_UtilizationChartData_FromProto(mapCtx, in.GetMemoryUtilizationChart())
	out.StorageUtilizationChart = ReportSummary_UtilizationChartData_FromProto(mapCtx, in.GetStorageUtilizationChart())
	out.OperatingSystem = ReportSummary_ChartData_FromProto(mapCtx, in.GetOperatingSystem())
	out.CoreCountHistogram = ReportSummary_HistogramChartData_FromProto(mapCtx, in.GetCoreCountHistogram())
	out.MemoryBytesHistogram = ReportSummary_HistogramChartData_FromProto(mapCtx, in.GetMemoryBytesHistogram())
	out.StorageBytesHistogram = ReportSummary_HistogramChartData_FromProto(mapCtx, in.GetStorageBytesHistogram())
	return out
}
func ReportSummary_AssetAggregateStats_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_AssetAggregateStats) *pb.ReportSummary_AssetAggregateStats {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_AssetAggregateStats{}
	out.TotalMemoryBytes = direct.ValueOf(in.TotalMemoryBytes)
	out.TotalStorageBytes = direct.ValueOf(in.TotalStorageBytes)
	out.TotalCores = direct.ValueOf(in.TotalCores)
	out.TotalAssets = direct.ValueOf(in.TotalAssets)
	out.MemoryUtilizationChart = ReportSummary_UtilizationChartData_ToProto(mapCtx, in.MemoryUtilizationChart)
	out.StorageUtilizationChart = ReportSummary_UtilizationChartData_ToProto(mapCtx, in.StorageUtilizationChart)
	out.OperatingSystem = ReportSummary_ChartData_ToProto(mapCtx, in.OperatingSystem)
	out.CoreCountHistogram = ReportSummary_HistogramChartData_ToProto(mapCtx, in.CoreCountHistogram)
	out.MemoryBytesHistogram = ReportSummary_HistogramChartData_ToProto(mapCtx, in.MemoryBytesHistogram)
	out.StorageBytesHistogram = ReportSummary_HistogramChartData_ToProto(mapCtx, in.StorageBytesHistogram)
	return out
}
func ReportSummary_ChartData_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_ChartData) *krm.ReportSummary_ChartData {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_ChartData{}
	out.DataPoints = direct.Slice_FromProto(mapCtx, in.DataPoints, ReportSummary_ChartData_DataPoint_FromProto)
	return out
}
func ReportSummary_ChartData_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_ChartData) *pb.ReportSummary_ChartData {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_ChartData{}
	out.DataPoints = direct.Slice_ToProto(mapCtx, in.DataPoints, ReportSummary_ChartData_DataPoint_ToProto)
	return out
}
func ReportSummary_ChartData_DataPoint_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_ChartData_DataPoint) *krm.ReportSummary_ChartData_DataPoint {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_ChartData_DataPoint{}
	out.Label = direct.LazyPtr(in.GetLabel())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func ReportSummary_ChartData_DataPoint_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_ChartData_DataPoint) *pb.ReportSummary_ChartData_DataPoint {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_ChartData_DataPoint{}
	out.Label = direct.ValueOf(in.Label)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func ReportSummary_ComputeEngineFinding_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_ComputeEngineFinding) *krm.ReportSummary_ComputeEngineFinding {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_ComputeEngineFinding{}
	out.AllocatedRegions = in.AllocatedRegions
	out.AllocatedAssetCount = direct.LazyPtr(in.GetAllocatedAssetCount())
	out.MachineSeriesAllocations = direct.Slice_FromProto(mapCtx, in.MachineSeriesAllocations, ReportSummary_MachineSeriesAllocation_FromProto)
	out.AllocatedDiskTypes = direct.EnumSlice_FromProto(mapCtx, in.AllocatedDiskTypes)
	return out
}
func ReportSummary_ComputeEngineFinding_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_ComputeEngineFinding) *pb.ReportSummary_ComputeEngineFinding {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_ComputeEngineFinding{}
	out.AllocatedRegions = in.AllocatedRegions
	out.AllocatedAssetCount = direct.ValueOf(in.AllocatedAssetCount)
	out.MachineSeriesAllocations = direct.Slice_ToProto(mapCtx, in.MachineSeriesAllocations, ReportSummary_MachineSeriesAllocation_ToProto)
	out.AllocatedDiskTypes = direct.EnumSlice_ToProto[pb.PersistentDiskType](mapCtx, in.AllocatedDiskTypes)
	return out
}
func ReportSummary_GroupFinding_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_GroupFinding) *krm.ReportSummary_GroupFinding {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_GroupFinding{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AssetAggregateStats = ReportSummary_AssetAggregateStats_FromProto(mapCtx, in.GetAssetAggregateStats())
	out.OverlappingAssetCount = direct.LazyPtr(in.GetOverlappingAssetCount())
	out.PreferenceSetFindings = direct.Slice_FromProto(mapCtx, in.PreferenceSetFindings, ReportSummary_GroupPreferenceSetFinding_FromProto)
	return out
}
func ReportSummary_GroupFinding_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_GroupFinding) *pb.ReportSummary_GroupFinding {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_GroupFinding{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.AssetAggregateStats = ReportSummary_AssetAggregateStats_ToProto(mapCtx, in.AssetAggregateStats)
	out.OverlappingAssetCount = direct.ValueOf(in.OverlappingAssetCount)
	out.PreferenceSetFindings = direct.Slice_ToProto(mapCtx, in.PreferenceSetFindings, ReportSummary_GroupPreferenceSetFinding_ToProto)
	return out
}
func ReportSummary_GroupPreferenceSetFinding_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_GroupPreferenceSetFinding) *krm.ReportSummary_GroupPreferenceSetFinding {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_GroupPreferenceSetFinding{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MachinePreferences = VirtualMachinePreferences_FromProto(mapCtx, in.GetMachinePreferences())
	out.MonthlyCostTotal = Money_FromProto(mapCtx, in.GetMonthlyCostTotal())
	out.MonthlyCostCompute = Money_FromProto(mapCtx, in.GetMonthlyCostCompute())
	out.MonthlyCostOsLicense = Money_FromProto(mapCtx, in.GetMonthlyCostOsLicense())
	out.MonthlyCostNetworkEgress = Money_FromProto(mapCtx, in.GetMonthlyCostNetworkEgress())
	out.MonthlyCostStorage = Money_FromProto(mapCtx, in.GetMonthlyCostStorage())
	out.MonthlyCostOther = Money_FromProto(mapCtx, in.GetMonthlyCostOther())
	out.ComputeEngineFinding = ReportSummary_ComputeEngineFinding_FromProto(mapCtx, in.GetComputeEngineFinding())
	out.VmwareEngineFinding = ReportSummary_VmwareEngineFinding_FromProto(mapCtx, in.GetVmwareEngineFinding())
	out.SoleTenantFinding = ReportSummary_SoleTenantFinding_FromProto(mapCtx, in.GetSoleTenantFinding())
	return out
}
func ReportSummary_GroupPreferenceSetFinding_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_GroupPreferenceSetFinding) *pb.ReportSummary_GroupPreferenceSetFinding {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_GroupPreferenceSetFinding{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.MachinePreferences = VirtualMachinePreferences_ToProto(mapCtx, in.MachinePreferences)
	out.MonthlyCostTotal = Money_ToProto(mapCtx, in.MonthlyCostTotal)
	out.MonthlyCostCompute = Money_ToProto(mapCtx, in.MonthlyCostCompute)
	out.MonthlyCostOsLicense = Money_ToProto(mapCtx, in.MonthlyCostOsLicense)
	out.MonthlyCostNetworkEgress = Money_ToProto(mapCtx, in.MonthlyCostNetworkEgress)
	out.MonthlyCostStorage = Money_ToProto(mapCtx, in.MonthlyCostStorage)
	out.MonthlyCostOther = Money_ToProto(mapCtx, in.MonthlyCostOther)
	out.ComputeEngineFinding = ReportSummary_ComputeEngineFinding_ToProto(mapCtx, in.ComputeEngineFinding)
	out.VmwareEngineFinding = ReportSummary_VmwareEngineFinding_ToProto(mapCtx, in.VmwareEngineFinding)
	out.SoleTenantFinding = ReportSummary_SoleTenantFinding_ToProto(mapCtx, in.SoleTenantFinding)
	return out
}
func ReportSummary_HistogramChartData_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_HistogramChartData) *krm.ReportSummary_HistogramChartData {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_HistogramChartData{}
	out.Buckets = direct.Slice_FromProto(mapCtx, in.Buckets, ReportSummary_HistogramChartData_Bucket_FromProto)
	return out
}
func ReportSummary_HistogramChartData_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_HistogramChartData) *pb.ReportSummary_HistogramChartData {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_HistogramChartData{}
	out.Buckets = direct.Slice_ToProto(mapCtx, in.Buckets, ReportSummary_HistogramChartData_Bucket_ToProto)
	return out
}
func ReportSummary_HistogramChartData_Bucket_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_HistogramChartData_Bucket) *krm.ReportSummary_HistogramChartData_Bucket {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_HistogramChartData_Bucket{}
	out.LowerBound = direct.LazyPtr(in.GetLowerBound())
	out.UpperBound = direct.LazyPtr(in.GetUpperBound())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func ReportSummary_HistogramChartData_Bucket_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_HistogramChartData_Bucket) *pb.ReportSummary_HistogramChartData_Bucket {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_HistogramChartData_Bucket{}
	out.LowerBound = direct.ValueOf(in.LowerBound)
	out.UpperBound = direct.ValueOf(in.UpperBound)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func ReportSummary_MachineSeriesAllocation_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_MachineSeriesAllocation) *krm.ReportSummary_MachineSeriesAllocation {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_MachineSeriesAllocation{}
	out.MachineSeries = MachineSeries_FromProto(mapCtx, in.GetMachineSeries())
	out.AllocatedAssetCount = direct.LazyPtr(in.GetAllocatedAssetCount())
	return out
}
func ReportSummary_MachineSeriesAllocation_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_MachineSeriesAllocation) *pb.ReportSummary_MachineSeriesAllocation {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_MachineSeriesAllocation{}
	out.MachineSeries = MachineSeries_ToProto(mapCtx, in.MachineSeries)
	out.AllocatedAssetCount = direct.ValueOf(in.AllocatedAssetCount)
	return out
}
func ReportSummary_SoleTenantFinding_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_SoleTenantFinding) *krm.ReportSummary_SoleTenantFinding {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_SoleTenantFinding{}
	out.AllocatedRegions = in.AllocatedRegions
	out.AllocatedAssetCount = direct.LazyPtr(in.GetAllocatedAssetCount())
	out.NodeAllocations = direct.Slice_FromProto(mapCtx, in.NodeAllocations, ReportSummary_SoleTenantNodeAllocation_FromProto)
	return out
}
func ReportSummary_SoleTenantFinding_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_SoleTenantFinding) *pb.ReportSummary_SoleTenantFinding {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_SoleTenantFinding{}
	out.AllocatedRegions = in.AllocatedRegions
	out.AllocatedAssetCount = direct.ValueOf(in.AllocatedAssetCount)
	out.NodeAllocations = direct.Slice_ToProto(mapCtx, in.NodeAllocations, ReportSummary_SoleTenantNodeAllocation_ToProto)
	return out
}
func ReportSummary_SoleTenantNodeAllocation_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_SoleTenantNodeAllocation) *krm.ReportSummary_SoleTenantNodeAllocation {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_SoleTenantNodeAllocation{}
	out.Node = SoleTenantNodeType_FromProto(mapCtx, in.GetNode())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.AllocatedAssetCount = direct.LazyPtr(in.GetAllocatedAssetCount())
	return out
}
func ReportSummary_SoleTenantNodeAllocation_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_SoleTenantNodeAllocation) *pb.ReportSummary_SoleTenantNodeAllocation {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_SoleTenantNodeAllocation{}
	out.Node = SoleTenantNodeType_ToProto(mapCtx, in.Node)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.AllocatedAssetCount = direct.ValueOf(in.AllocatedAssetCount)
	return out
}
func ReportSummary_UtilizationChartData_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_UtilizationChartData) *krm.ReportSummary_UtilizationChartData {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_UtilizationChartData{}
	out.Used = direct.LazyPtr(in.GetUsed())
	out.Free = direct.LazyPtr(in.GetFree())
	return out
}
func ReportSummary_UtilizationChartData_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_UtilizationChartData) *pb.ReportSummary_UtilizationChartData {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_UtilizationChartData{}
	out.Used = direct.ValueOf(in.Used)
	out.Free = direct.ValueOf(in.Free)
	return out
}
func ReportSummary_VmwareEngineFinding_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_VmwareEngineFinding) *krm.ReportSummary_VmwareEngineFinding {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_VmwareEngineFinding{}
	out.AllocatedRegions = in.AllocatedRegions
	out.AllocatedAssetCount = direct.LazyPtr(in.GetAllocatedAssetCount())
	out.NodeAllocations = direct.Slice_FromProto(mapCtx, in.NodeAllocations, ReportSummary_VmwareNodeAllocation_FromProto)
	return out
}
func ReportSummary_VmwareEngineFinding_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_VmwareEngineFinding) *pb.ReportSummary_VmwareEngineFinding {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_VmwareEngineFinding{}
	out.AllocatedRegions = in.AllocatedRegions
	out.AllocatedAssetCount = direct.ValueOf(in.AllocatedAssetCount)
	out.NodeAllocations = direct.Slice_ToProto(mapCtx, in.NodeAllocations, ReportSummary_VmwareNodeAllocation_ToProto)
	return out
}
func ReportSummary_VmwareNode_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_VmwareNode) *krm.ReportSummary_VmwareNode {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_VmwareNode{}
	out.Code = direct.LazyPtr(in.GetCode())
	return out
}
func ReportSummary_VmwareNode_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_VmwareNode) *pb.ReportSummary_VmwareNode {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_VmwareNode{}
	out.Code = direct.ValueOf(in.Code)
	return out
}
func ReportSummary_VmwareNodeAllocation_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_VmwareNodeAllocation) *krm.ReportSummary_VmwareNodeAllocation {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_VmwareNodeAllocation{}
	out.VmwareNode = ReportSummary_VmwareNode_FromProto(mapCtx, in.GetVmwareNode())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.AllocatedAssetCount = direct.LazyPtr(in.GetAllocatedAssetCount())
	return out
}
func ReportSummary_VmwareNodeAllocation_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_VmwareNodeAllocation) *pb.ReportSummary_VmwareNodeAllocation {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_VmwareNodeAllocation{}
	out.VmwareNode = ReportSummary_VmwareNode_ToProto(mapCtx, in.VmwareNode)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.AllocatedAssetCount = direct.ValueOf(in.AllocatedAssetCount)
	return out
}
func SoleTenancyPreferences_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenancyPreferences) *krm.SoleTenancyPreferences {
	if in == nil {
		return nil
	}
	out := &krm.SoleTenancyPreferences{}
	out.CpuOvercommitRatio = direct.LazyPtr(in.GetCpuOvercommitRatio())
	out.HostMaintenancePolicy = direct.Enum_FromProto(mapCtx, in.GetHostMaintenancePolicy())
	out.CommitmentPlan = direct.Enum_FromProto(mapCtx, in.GetCommitmentPlan())
	out.NodeTypes = direct.Slice_FromProto(mapCtx, in.NodeTypes, SoleTenantNodeType_FromProto)
	return out
}
func SoleTenancyPreferences_ToProto(mapCtx *direct.MapContext, in *krm.SoleTenancyPreferences) *pb.SoleTenancyPreferences {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenancyPreferences{}
	out.CpuOvercommitRatio = direct.ValueOf(in.CpuOvercommitRatio)
	out.HostMaintenancePolicy = direct.Enum_ToProto[pb.SoleTenancyPreferences_HostMaintenancePolicy](mapCtx, in.HostMaintenancePolicy)
	out.CommitmentPlan = direct.Enum_ToProto[pb.SoleTenancyPreferences_CommitmentPlan](mapCtx, in.CommitmentPlan)
	out.NodeTypes = direct.Slice_ToProto(mapCtx, in.NodeTypes, SoleTenantNodeType_ToProto)
	return out
}
func SoleTenantNodeType_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenantNodeType) *krm.SoleTenantNodeType {
	if in == nil {
		return nil
	}
	out := &krm.SoleTenantNodeType{}
	out.NodeName = direct.LazyPtr(in.GetNodeName())
	return out
}
func SoleTenantNodeType_ToProto(mapCtx *direct.MapContext, in *krm.SoleTenantNodeType) *pb.SoleTenantNodeType {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenantNodeType{}
	out.NodeName = direct.ValueOf(in.NodeName)
	return out
}
func VirtualMachinePreferences_FromProto(mapCtx *direct.MapContext, in *pb.VirtualMachinePreferences) *krm.VirtualMachinePreferences {
	if in == nil {
		return nil
	}
	out := &krm.VirtualMachinePreferences{}
	out.TargetProduct = direct.Enum_FromProto(mapCtx, in.GetTargetProduct())
	out.RegionPreferences = RegionPreferences_FromProto(mapCtx, in.GetRegionPreferences())
	out.CommitmentPlan = direct.Enum_FromProto(mapCtx, in.GetCommitmentPlan())
	out.SizingOptimizationStrategy = direct.Enum_FromProto(mapCtx, in.GetSizingOptimizationStrategy())
	out.ComputeEnginePreferences = ComputeEnginePreferences_FromProto(mapCtx, in.GetComputeEnginePreferences())
	out.VmwareEnginePreferences = VmwareEnginePreferences_FromProto(mapCtx, in.GetVmwareEnginePreferences())
	out.SoleTenancyPreferences = SoleTenancyPreferences_FromProto(mapCtx, in.GetSoleTenancyPreferences())
	return out
}
func VirtualMachinePreferences_ToProto(mapCtx *direct.MapContext, in *krm.VirtualMachinePreferences) *pb.VirtualMachinePreferences {
	if in == nil {
		return nil
	}
	out := &pb.VirtualMachinePreferences{}
	out.TargetProduct = direct.Enum_ToProto[pb.ComputeMigrationTargetProduct](mapCtx, in.TargetProduct)
	out.RegionPreferences = RegionPreferences_ToProto(mapCtx, in.RegionPreferences)
	out.CommitmentPlan = direct.Enum_ToProto[pb.CommitmentPlan](mapCtx, in.CommitmentPlan)
	out.SizingOptimizationStrategy = direct.Enum_ToProto[pb.SizingOptimizationStrategy](mapCtx, in.SizingOptimizationStrategy)
	out.ComputeEnginePreferences = ComputeEnginePreferences_ToProto(mapCtx, in.ComputeEnginePreferences)
	out.VmwareEnginePreferences = VmwareEnginePreferences_ToProto(mapCtx, in.VmwareEnginePreferences)
	out.SoleTenancyPreferences = SoleTenancyPreferences_ToProto(mapCtx, in.SoleTenancyPreferences)
	return out
}
func VmwareEnginePreferences_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEnginePreferences) *krm.VmwareEnginePreferences {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEnginePreferences{}
	out.CpuOvercommitRatio = direct.LazyPtr(in.GetCpuOvercommitRatio())
	out.MemoryOvercommitRatio = direct.LazyPtr(in.GetMemoryOvercommitRatio())
	out.StorageDeduplicationCompressionRatio = direct.LazyPtr(in.GetStorageDeduplicationCompressionRatio())
	out.CommitmentPlan = direct.Enum_FromProto(mapCtx, in.GetCommitmentPlan())
	return out
}
func VmwareEnginePreferences_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEnginePreferences) *pb.VmwareEnginePreferences {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEnginePreferences{}
	out.CpuOvercommitRatio = direct.ValueOf(in.CpuOvercommitRatio)
	out.MemoryOvercommitRatio = direct.ValueOf(in.MemoryOvercommitRatio)
	out.StorageDeduplicationCompressionRatio = direct.ValueOf(in.StorageDeduplicationCompressionRatio)
	out.CommitmentPlan = direct.Enum_ToProto[pb.VmwareEnginePreferences_CommitmentPlan](mapCtx, in.CommitmentPlan)
	return out
}
