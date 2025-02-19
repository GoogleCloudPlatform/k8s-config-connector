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
func MigrationcenterPreferenceSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PreferenceSet) *krm.MigrationcenterPreferenceSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterPreferenceSetObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VirtualMachinePreferences
	return out
}
func MigrationcenterPreferenceSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterPreferenceSetObservedState) *pb.PreferenceSet {
	if in == nil {
		return nil
	}
	out := &pb.PreferenceSet{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VirtualMachinePreferences
	return out
}
func MigrationcenterPreferenceSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.PreferenceSet) *krm.MigrationcenterPreferenceSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterPreferenceSetSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VirtualMachinePreferences
	return out
}
func MigrationcenterPreferenceSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterPreferenceSetSpec) *pb.PreferenceSet {
	if in == nil {
		return nil
	}
	out := &pb.PreferenceSet{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VirtualMachinePreferences
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
func SoleTenancyPreferences_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenancyPreferences) *krm.SoleTenancyPreferences {
	if in == nil {
		return nil
	}
	out := &krm.SoleTenancyPreferences{}
	out.CPUOvercommitRatio = direct.LazyPtr(in.GetCpuOvercommitRatio())
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
	out.CpuOvercommitRatio = direct.ValueOf(in.CPUOvercommitRatio)
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
	out.CPUOvercommitRatio = direct.LazyPtr(in.GetCpuOvercommitRatio())
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
	out.CpuOvercommitRatio = direct.ValueOf(in.CPUOvercommitRatio)
	out.MemoryOvercommitRatio = direct.ValueOf(in.MemoryOvercommitRatio)
	out.StorageDeduplicationCompressionRatio = direct.ValueOf(in.StorageDeduplicationCompressionRatio)
	out.CommitmentPlan = direct.Enum_ToProto[pb.VmwareEnginePreferences_CommitmentPlan](mapCtx, in.CommitmentPlan)
	return out
}
