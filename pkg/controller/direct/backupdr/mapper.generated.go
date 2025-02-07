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

package backupdr

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BackupApplianceBackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.BackupApplianceBackupConfig) *krm.BackupApplianceBackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackupApplianceBackupConfig{}
	out.BackupApplianceName = direct.LazyPtr(in.GetBackupApplianceName())
	out.BackupApplianceID = direct.LazyPtr(in.GetBackupApplianceId())
	out.SlaID = direct.LazyPtr(in.GetSlaId())
	out.ApplicationName = direct.LazyPtr(in.GetApplicationName())
	out.HostName = direct.LazyPtr(in.GetHostName())
	out.SltName = direct.LazyPtr(in.GetSltName())
	out.SlpName = direct.LazyPtr(in.GetSlpName())
	return out
}
func BackupApplianceBackupConfig_ToProto(mapCtx *direct.MapContext, in *krm.BackupApplianceBackupConfig) *pb.BackupApplianceBackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackupApplianceBackupConfig{}
	out.BackupApplianceName = direct.ValueOf(in.BackupApplianceName)
	out.BackupApplianceId = direct.ValueOf(in.BackupApplianceID)
	out.SlaId = direct.ValueOf(in.SlaID)
	out.ApplicationName = direct.ValueOf(in.ApplicationName)
	out.HostName = direct.ValueOf(in.HostName)
	out.SltName = direct.ValueOf(in.SltName)
	out.SlpName = direct.ValueOf(in.SlpName)
	return out
}
func BackupConfigInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfigInfo) *krm.BackupConfigInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupConfigInfo{}
	// MISSING: LastBackupState
	// MISSING: LastSuccessfulBackupConsistencyTime
	// MISSING: LastBackupError
	out.GcpBackupConfig = GcpBackupConfig_FromProto(mapCtx, in.GetGcpBackupConfig())
	out.BackupApplianceBackupConfig = BackupApplianceBackupConfig_FromProto(mapCtx, in.GetBackupApplianceBackupConfig())
	return out
}
func BackupConfigInfo_ToProto(mapCtx *direct.MapContext, in *krm.BackupConfigInfo) *pb.BackupConfigInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupConfigInfo{}
	// MISSING: LastBackupState
	// MISSING: LastSuccessfulBackupConsistencyTime
	// MISSING: LastBackupError
	if oneof := GcpBackupConfig_ToProto(mapCtx, in.GcpBackupConfig); oneof != nil {
		out.BackupConfig = &pb.BackupConfigInfo_GcpBackupConfig{GcpBackupConfig: oneof}
	}
	if oneof := BackupApplianceBackupConfig_ToProto(mapCtx, in.BackupApplianceBackupConfig); oneof != nil {
		out.BackupConfig = &pb.BackupConfigInfo_BackupApplianceBackupConfig{BackupApplianceBackupConfig: oneof}
	}
	return out
}
func BackupConfigInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfigInfo) *krm.BackupConfigInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupConfigInfoObservedState{}
	out.LastBackupState = direct.Enum_FromProto(mapCtx, in.GetLastBackupState())
	out.LastSuccessfulBackupConsistencyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastSuccessfulBackupConsistencyTime())
	out.LastBackupError = Status_FromProto(mapCtx, in.GetLastBackupError())
	// MISSING: GcpBackupConfig
	// MISSING: BackupApplianceBackupConfig
	return out
}
func BackupConfigInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupConfigInfoObservedState) *pb.BackupConfigInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupConfigInfo{}
	out.LastBackupState = direct.Enum_ToProto[pb.BackupConfigInfo_LastBackupState](mapCtx, in.LastBackupState)
	out.LastSuccessfulBackupConsistencyTime = direct.StringTimestamp_ToProto(mapCtx, in.LastSuccessfulBackupConsistencyTime)
	out.LastBackupError = Status_ToProto(mapCtx, in.LastBackupError)
	// MISSING: GcpBackupConfig
	// MISSING: BackupApplianceBackupConfig
	return out
}
func BackupdrBackupPlanAssociationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupdrBackupPlanAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanAssociationObservedState{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanAssociationObservedState) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupdrBackupPlanAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanAssociationSpec{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanAssociationSpec) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupdrBackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupVaultObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupdrBackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupVaultSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrDataSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.BackupdrDataSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrDataSourceObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrDataSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrDataSourceObservedState) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrDataSourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.BackupdrDataSourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrDataSourceSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrDataSourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrDataSourceSpec) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ComputeInstanceDataSourceProperties_FromProto(mapCtx *direct.MapContext, in *pb.ComputeInstanceDataSourceProperties) *krm.ComputeInstanceDataSourceProperties {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceDataSourceProperties{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.TotalDiskCount = direct.LazyPtr(in.GetTotalDiskCount())
	out.TotalDiskSizeGB = direct.LazyPtr(in.GetTotalDiskSizeGb())
	return out
}
func ComputeInstanceDataSourceProperties_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceDataSourceProperties) *pb.ComputeInstanceDataSourceProperties {
	if in == nil {
		return nil
	}
	out := &pb.ComputeInstanceDataSourceProperties{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.TotalDiskCount = direct.ValueOf(in.TotalDiskCount)
	out.TotalDiskSizeGb = direct.ValueOf(in.TotalDiskSizeGB)
	return out
}
func DataSource_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSource {
	if in == nil {
		return nil
	}
	out := &krm.DataSource{}
	// MISSING: Name
	// MISSING: State
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.BackupCount = in.BackupCount
	out.Etag = in.Etag
	out.TotalStoredBytes = in.TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	out.DataSourceGcpResource = DataSourceGcpResource_FromProto(mapCtx, in.GetDataSourceGcpResource())
	out.DataSourceBackupApplianceApplication = DataSourceBackupApplianceApplication_FromProto(mapCtx, in.GetDataSourceBackupApplianceApplication())
	return out
}
func DataSource_ToProto(mapCtx *direct.MapContext, in *krm.DataSource) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: State
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.BackupCount = in.BackupCount
	out.Etag = in.Etag
	out.TotalStoredBytes = in.TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	if oneof := DataSourceGcpResource_ToProto(mapCtx, in.DataSourceGcpResource); oneof != nil {
		out.SourceResource = &pb.DataSource_DataSourceGcpResource{DataSourceGcpResource: oneof}
	}
	if oneof := DataSourceBackupApplianceApplication_ToProto(mapCtx, in.DataSourceBackupApplianceApplication); oneof != nil {
		out.SourceResource = &pb.DataSource_DataSourceBackupApplianceApplication{DataSourceBackupApplianceApplication: oneof}
	}
	return out
}
func DataSourceBackupApplianceApplication_FromProto(mapCtx *direct.MapContext, in *pb.DataSourceBackupApplianceApplication) *krm.DataSourceBackupApplianceApplication {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceBackupApplianceApplication{}
	out.ApplicationName = direct.LazyPtr(in.GetApplicationName())
	out.BackupAppliance = direct.LazyPtr(in.GetBackupAppliance())
	out.ApplianceID = direct.LazyPtr(in.GetApplianceId())
	out.Type = direct.LazyPtr(in.GetType())
	out.ApplicationID = direct.LazyPtr(in.GetApplicationId())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.HostID = direct.LazyPtr(in.GetHostId())
	return out
}
func DataSourceBackupApplianceApplication_ToProto(mapCtx *direct.MapContext, in *krm.DataSourceBackupApplianceApplication) *pb.DataSourceBackupApplianceApplication {
	if in == nil {
		return nil
	}
	out := &pb.DataSourceBackupApplianceApplication{}
	out.ApplicationName = direct.ValueOf(in.ApplicationName)
	out.BackupAppliance = direct.ValueOf(in.BackupAppliance)
	out.ApplianceId = direct.ValueOf(in.ApplianceID)
	out.Type = direct.ValueOf(in.Type)
	out.ApplicationId = direct.ValueOf(in.ApplicationID)
	out.Hostname = direct.ValueOf(in.Hostname)
	out.HostId = direct.ValueOf(in.HostID)
	return out
}
func DataSourceGcpResource_FromProto(mapCtx *direct.MapContext, in *pb.DataSourceGcpResource) *krm.DataSourceGcpResource {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceGcpResource{}
	// MISSING: GcpResourcename
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Type = direct.LazyPtr(in.GetType())
	out.ComputeInstanceDatasourceProperties = ComputeInstanceDataSourceProperties_FromProto(mapCtx, in.GetComputeInstanceDatasourceProperties())
	return out
}
func DataSourceGcpResource_ToProto(mapCtx *direct.MapContext, in *krm.DataSourceGcpResource) *pb.DataSourceGcpResource {
	if in == nil {
		return nil
	}
	out := &pb.DataSourceGcpResource{}
	// MISSING: GcpResourcename
	out.Location = direct.ValueOf(in.Location)
	out.Type = direct.ValueOf(in.Type)
	if oneof := ComputeInstanceDataSourceProperties_ToProto(mapCtx, in.ComputeInstanceDatasourceProperties); oneof != nil {
		out.GcpResourceProperties = &pb.DataSourceGcpResource_ComputeInstanceDatasourceProperties{ComputeInstanceDatasourceProperties: oneof}
	}
	return out
}
func DataSourceGcpResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSourceGcpResource) *krm.DataSourceGcpResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceGcpResourceObservedState{}
	out.GcpResourcename = direct.LazyPtr(in.GetGcpResourcename())
	// MISSING: Location
	// MISSING: Type
	// MISSING: ComputeInstanceDatasourceProperties
	return out
}
func DataSourceGcpResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataSourceGcpResourceObservedState) *pb.DataSourceGcpResource {
	if in == nil {
		return nil
	}
	out := &pb.DataSourceGcpResource{}
	out.GcpResourcename = direct.ValueOf(in.GcpResourcename)
	// MISSING: Location
	// MISSING: Type
	// MISSING: ComputeInstanceDatasourceProperties
	return out
}
func DataSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	out.ConfigState = direct.Enum_FromProto(mapCtx, in.GetConfigState())
	out.BackupConfigInfo = BackupConfigInfo_FromProto(mapCtx, in.GetBackupConfigInfo())
	out.DataSourceGcpResource = DataSourceGcpResourceObservedState_FromProto(mapCtx, in.GetDataSourceGcpResource())
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func DataSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataSourceObservedState) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.DataSource_State](mapCtx, in.State)
	// MISSING: Labels
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.CreateTime); oneof != nil {
		out.CreateTime = &pb.DataSource_CreateTime{CreateTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime); oneof != nil {
		out.UpdateTime = &pb.DataSource_UpdateTime{UpdateTime: oneof}
	}
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	out.ConfigState = direct.Enum_ToProto[pb.BackupConfigState](mapCtx, in.ConfigState)
	out.BackupConfigInfo = BackupConfigInfo_ToProto(mapCtx, in.BackupConfigInfo)
	if oneof := DataSourceGcpResourceObservedState_ToProto(mapCtx, in.DataSourceGcpResource); oneof != nil {
		out.SourceResource = &pb.DataSource_DataSourceGcpResource{DataSourceGcpResource: oneof}
	}
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func GcpBackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcpBackupConfig) *krm.GcpBackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcpBackupConfig{}
	out.BackupPlan = direct.LazyPtr(in.GetBackupPlan())
	out.BackupPlanDescription = direct.LazyPtr(in.GetBackupPlanDescription())
	out.BackupPlanAssociation = direct.LazyPtr(in.GetBackupPlanAssociation())
	out.BackupPlanRules = in.BackupPlanRules
	return out
}
func GcpBackupConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcpBackupConfig) *pb.GcpBackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcpBackupConfig{}
	out.BackupPlan = direct.ValueOf(in.BackupPlan)
	out.BackupPlanDescription = direct.ValueOf(in.BackupPlanDescription)
	out.BackupPlanAssociation = direct.ValueOf(in.BackupPlanAssociation)
	out.BackupPlanRules = in.BackupPlanRules
	return out
}
