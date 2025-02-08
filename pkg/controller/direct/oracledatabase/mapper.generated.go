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

package oracledatabase

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
)
func CloudExadataInfrastructure_FromProto(mapCtx *direct.MapContext, in *pb.CloudExadataInfrastructure) *krm.CloudExadataInfrastructure {
	if in == nil {
		return nil
	}
	out := &krm.CloudExadataInfrastructure{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GcpOracleZone = direct.LazyPtr(in.GetGcpOracleZone())
	// MISSING: EntitlementID
	out.Properties = CloudExadataInfrastructureProperties_FromProto(mapCtx, in.GetProperties())
	out.Labels = in.Labels
	// MISSING: CreateTime
	return out
}
func CloudExadataInfrastructure_ToProto(mapCtx *direct.MapContext, in *krm.CloudExadataInfrastructure) *pb.CloudExadataInfrastructure {
	if in == nil {
		return nil
	}
	out := &pb.CloudExadataInfrastructure{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.GcpOracleZone = direct.ValueOf(in.GcpOracleZone)
	// MISSING: EntitlementID
	out.Properties = CloudExadataInfrastructureProperties_ToProto(mapCtx, in.Properties)
	out.Labels = in.Labels
	// MISSING: CreateTime
	return out
}
func CloudExadataInfrastructureObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudExadataInfrastructure) *krm.CloudExadataInfrastructureObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudExadataInfrastructureObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	out.EntitlementID = direct.LazyPtr(in.GetEntitlementId())
	out.Properties = CloudExadataInfrastructurePropertiesObservedState_FromProto(mapCtx, in.GetProperties())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func CloudExadataInfrastructureObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudExadataInfrastructureObservedState) *pb.CloudExadataInfrastructure {
	if in == nil {
		return nil
	}
	out := &pb.CloudExadataInfrastructure{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	out.EntitlementId = direct.ValueOf(in.EntitlementID)
	out.Properties = CloudExadataInfrastructurePropertiesObservedState_ToProto(mapCtx, in.Properties)
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func CloudExadataInfrastructureProperties_FromProto(mapCtx *direct.MapContext, in *pb.CloudExadataInfrastructureProperties) *krm.CloudExadataInfrastructureProperties {
	if in == nil {
		return nil
	}
	out := &krm.CloudExadataInfrastructureProperties{}
	// MISSING: Ocid
	out.ComputeCount = direct.LazyPtr(in.GetComputeCount())
	out.StorageCount = direct.LazyPtr(in.GetStorageCount())
	out.TotalStorageSizeGB = direct.LazyPtr(in.GetTotalStorageSizeGb())
	// MISSING: AvailableStorageSizeGB
	out.MaintenanceWindow = MaintenanceWindow_FromProto(mapCtx, in.GetMaintenanceWindow())
	// MISSING: State
	out.Shape = direct.LazyPtr(in.GetShape())
	// MISSING: OciURL
	// MISSING: CpuCount
	// MISSING: MaxCpuCount
	// MISSING: MemorySizeGB
	// MISSING: MaxMemoryGB
	// MISSING: DbNodeStorageSizeGB
	// MISSING: MaxDbNodeStorageSizeGB
	// MISSING: DataStorageSizeTb
	// MISSING: MaxDataStorageTb
	// MISSING: ActivatedStorageCount
	// MISSING: AdditionalStorageCount
	// MISSING: DbServerVersion
	// MISSING: StorageServerVersion
	// MISSING: NextMaintenanceRunID
	// MISSING: NextMaintenanceRunTime
	// MISSING: NextSecurityMaintenanceRunTime
	out.CustomerContacts = direct.Slice_FromProto(mapCtx, in.CustomerContacts, CustomerContact_FromProto)
	// MISSING: MonthlyStorageServerVersion
	// MISSING: MonthlyDbServerVersion
	return out
}
func CloudExadataInfrastructureProperties_ToProto(mapCtx *direct.MapContext, in *krm.CloudExadataInfrastructureProperties) *pb.CloudExadataInfrastructureProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudExadataInfrastructureProperties{}
	// MISSING: Ocid
	out.ComputeCount = direct.ValueOf(in.ComputeCount)
	out.StorageCount = direct.ValueOf(in.StorageCount)
	out.TotalStorageSizeGb = direct.ValueOf(in.TotalStorageSizeGB)
	// MISSING: AvailableStorageSizeGB
	out.MaintenanceWindow = MaintenanceWindow_ToProto(mapCtx, in.MaintenanceWindow)
	// MISSING: State
	out.Shape = direct.ValueOf(in.Shape)
	// MISSING: OciURL
	// MISSING: CpuCount
	// MISSING: MaxCpuCount
	// MISSING: MemorySizeGB
	// MISSING: MaxMemoryGB
	// MISSING: DbNodeStorageSizeGB
	// MISSING: MaxDbNodeStorageSizeGB
	// MISSING: DataStorageSizeTb
	// MISSING: MaxDataStorageTb
	// MISSING: ActivatedStorageCount
	// MISSING: AdditionalStorageCount
	// MISSING: DbServerVersion
	// MISSING: StorageServerVersion
	// MISSING: NextMaintenanceRunID
	// MISSING: NextMaintenanceRunTime
	// MISSING: NextSecurityMaintenanceRunTime
	out.CustomerContacts = direct.Slice_ToProto(mapCtx, in.CustomerContacts, CustomerContact_ToProto)
	// MISSING: MonthlyStorageServerVersion
	// MISSING: MonthlyDbServerVersion
	return out
}
func CloudExadataInfrastructurePropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudExadataInfrastructureProperties) *krm.CloudExadataInfrastructurePropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudExadataInfrastructurePropertiesObservedState{}
	out.Ocid = direct.LazyPtr(in.GetOcid())
	// MISSING: ComputeCount
	// MISSING: StorageCount
	// MISSING: TotalStorageSizeGB
	out.AvailableStorageSizeGB = direct.LazyPtr(in.GetAvailableStorageSizeGb())
	// MISSING: MaintenanceWindow
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Shape
	out.OciURL = direct.LazyPtr(in.GetOciUrl())
	out.CpuCount = direct.LazyPtr(in.GetCpuCount())
	out.MaxCpuCount = direct.LazyPtr(in.GetMaxCpuCount())
	out.MemorySizeGB = direct.LazyPtr(in.GetMemorySizeGb())
	out.MaxMemoryGB = direct.LazyPtr(in.GetMaxMemoryGb())
	out.DbNodeStorageSizeGB = direct.LazyPtr(in.GetDbNodeStorageSizeGb())
	out.MaxDbNodeStorageSizeGB = direct.LazyPtr(in.GetMaxDbNodeStorageSizeGb())
	out.DataStorageSizeTb = direct.LazyPtr(in.GetDataStorageSizeTb())
	out.MaxDataStorageTb = direct.LazyPtr(in.GetMaxDataStorageTb())
	out.ActivatedStorageCount = direct.LazyPtr(in.GetActivatedStorageCount())
	out.AdditionalStorageCount = direct.LazyPtr(in.GetAdditionalStorageCount())
	out.DbServerVersion = direct.LazyPtr(in.GetDbServerVersion())
	out.StorageServerVersion = direct.LazyPtr(in.GetStorageServerVersion())
	out.NextMaintenanceRunID = direct.LazyPtr(in.GetNextMaintenanceRunId())
	out.NextMaintenanceRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextMaintenanceRunTime())
	out.NextSecurityMaintenanceRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextSecurityMaintenanceRunTime())
	// MISSING: CustomerContacts
	out.MonthlyStorageServerVersion = direct.LazyPtr(in.GetMonthlyStorageServerVersion())
	out.MonthlyDbServerVersion = direct.LazyPtr(in.GetMonthlyDbServerVersion())
	return out
}
func CloudExadataInfrastructurePropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudExadataInfrastructurePropertiesObservedState) *pb.CloudExadataInfrastructureProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudExadataInfrastructureProperties{}
	out.Ocid = direct.ValueOf(in.Ocid)
	// MISSING: ComputeCount
	// MISSING: StorageCount
	// MISSING: TotalStorageSizeGB
	out.AvailableStorageSizeGb = direct.ValueOf(in.AvailableStorageSizeGB)
	// MISSING: MaintenanceWindow
	out.State = direct.Enum_ToProto[pb.CloudExadataInfrastructureProperties_State](mapCtx, in.State)
	// MISSING: Shape
	out.OciUrl = direct.ValueOf(in.OciURL)
	out.CpuCount = direct.ValueOf(in.CpuCount)
	out.MaxCpuCount = direct.ValueOf(in.MaxCpuCount)
	out.MemorySizeGb = direct.ValueOf(in.MemorySizeGB)
	out.MaxMemoryGb = direct.ValueOf(in.MaxMemoryGB)
	out.DbNodeStorageSizeGb = direct.ValueOf(in.DbNodeStorageSizeGB)
	out.MaxDbNodeStorageSizeGb = direct.ValueOf(in.MaxDbNodeStorageSizeGB)
	out.DataStorageSizeTb = direct.ValueOf(in.DataStorageSizeTb)
	out.MaxDataStorageTb = direct.ValueOf(in.MaxDataStorageTb)
	out.ActivatedStorageCount = direct.ValueOf(in.ActivatedStorageCount)
	out.AdditionalStorageCount = direct.ValueOf(in.AdditionalStorageCount)
	out.DbServerVersion = direct.ValueOf(in.DbServerVersion)
	out.StorageServerVersion = direct.ValueOf(in.StorageServerVersion)
	out.NextMaintenanceRunId = direct.ValueOf(in.NextMaintenanceRunID)
	out.NextMaintenanceRunTime = direct.StringTimestamp_ToProto(mapCtx, in.NextMaintenanceRunTime)
	out.NextSecurityMaintenanceRunTime = direct.StringTimestamp_ToProto(mapCtx, in.NextSecurityMaintenanceRunTime)
	// MISSING: CustomerContacts
	out.MonthlyStorageServerVersion = direct.ValueOf(in.MonthlyStorageServerVersion)
	out.MonthlyDbServerVersion = direct.ValueOf(in.MonthlyDbServerVersion)
	return out
}
func CustomerContact_FromProto(mapCtx *direct.MapContext, in *pb.CustomerContact) *krm.CustomerContact {
	if in == nil {
		return nil
	}
	out := &krm.CustomerContact{}
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func CustomerContact_ToProto(mapCtx *direct.MapContext, in *krm.CustomerContact) *pb.CustomerContact {
	if in == nil {
		return nil
	}
	out := &pb.CustomerContact{}
	out.Email = direct.ValueOf(in.Email)
	return out
}
func MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceWindow) *krm.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceWindow{}
	out.Preference = direct.Enum_FromProto(mapCtx, in.GetPreference())
	out.Months = direct.EnumSlice_FromProto(mapCtx, in.Months)
	out.WeeksOfMonth = in.WeeksOfMonth
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	out.HoursOfDay = in.HoursOfDay
	out.LeadTimeWeek = direct.LazyPtr(in.GetLeadTimeWeek())
	out.PatchingMode = direct.Enum_FromProto(mapCtx, in.GetPatchingMode())
	out.CustomActionTimeoutMins = direct.LazyPtr(in.GetCustomActionTimeoutMins())
	out.IsCustomActionTimeoutEnabled = direct.LazyPtr(in.GetIsCustomActionTimeoutEnabled())
	return out
}
func MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceWindow) *pb.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceWindow{}
	out.Preference = direct.Enum_ToProto[pb.MaintenanceWindow_MaintenanceWindowPreference](mapCtx, in.Preference)
	out.Months = direct.EnumSlice_ToProto[pb.Month](mapCtx, in.Months)
	out.WeeksOfMonth = in.WeeksOfMonth
	out.DaysOfWeek = direct.EnumSlice_ToProto[pb.DayOfWeek](mapCtx, in.DaysOfWeek)
	out.HoursOfDay = in.HoursOfDay
	out.LeadTimeWeek = direct.ValueOf(in.LeadTimeWeek)
	out.PatchingMode = direct.Enum_ToProto[pb.MaintenanceWindow_PatchingMode](mapCtx, in.PatchingMode)
	out.CustomActionTimeoutMins = direct.ValueOf(in.CustomActionTimeoutMins)
	out.IsCustomActionTimeoutEnabled = direct.ValueOf(in.IsCustomActionTimeoutEnabled)
	return out
}
func OracledatabaseCloudExadataInfrastructureObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudExadataInfrastructure) *krm.OracledatabaseCloudExadataInfrastructureObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseCloudExadataInfrastructureObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: EntitlementID
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	return out
}
func OracledatabaseCloudExadataInfrastructureObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseCloudExadataInfrastructureObservedState) *pb.CloudExadataInfrastructure {
	if in == nil {
		return nil
	}
	out := &pb.CloudExadataInfrastructure{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: EntitlementID
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	return out
}
func OracledatabaseCloudExadataInfrastructureSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudExadataInfrastructure) *krm.OracledatabaseCloudExadataInfrastructureSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseCloudExadataInfrastructureSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: EntitlementID
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	return out
}
func OracledatabaseCloudExadataInfrastructureSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseCloudExadataInfrastructureSpec) *pb.CloudExadataInfrastructure {
	if in == nil {
		return nil
	}
	out := &pb.CloudExadataInfrastructure{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: EntitlementID
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	return out
}
