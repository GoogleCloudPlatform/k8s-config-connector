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

package metastore

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/metastore/apiv1beta/metastorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
)
func AuxiliaryVersionConfig_FromProto(mapCtx *direct.MapContext, in *pb.AuxiliaryVersionConfig) *krm.AuxiliaryVersionConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuxiliaryVersionConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.ConfigOverrides = in.ConfigOverrides
	// MISSING: NetworkConfig
	return out
}
func AuxiliaryVersionConfig_ToProto(mapCtx *direct.MapContext, in *krm.AuxiliaryVersionConfig) *pb.AuxiliaryVersionConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuxiliaryVersionConfig{}
	out.Version = direct.ValueOf(in.Version)
	out.ConfigOverrides = in.ConfigOverrides
	// MISSING: NetworkConfig
	return out
}
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: ServiceRevision
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: RestoringServices
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: ServiceRevision
	out.Description = direct.ValueOf(in.Description)
	// MISSING: RestoringServices
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceRevision = Service_FromProto(mapCtx, in.GetServiceRevision())
	// MISSING: Description
	out.RestoringServices = in.RestoringServices
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceRevision = Service_ToProto(mapCtx, in.ServiceRevision)
	// MISSING: Description
	out.RestoringServices = in.RestoringServices
	return out
}
func DataCatalogConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataCatalogConfig) *krm.DataCatalogConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func DataCatalogConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogConfig) *pb.DataCatalogConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataCatalogConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func DataplexConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataplexConfig) *krm.DataplexConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataplexConfig{}
	// MISSING: LakeResources
	return out
}
func DataplexConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataplexConfig) *pb.DataplexConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataplexConfig{}
	// MISSING: LakeResources
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
func HiveMetastoreConfig_FromProto(mapCtx *direct.MapContext, in *pb.HiveMetastoreConfig) *krm.HiveMetastoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.HiveMetastoreConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.ConfigOverrides = in.ConfigOverrides
	out.KerberosConfig = KerberosConfig_FromProto(mapCtx, in.GetKerberosConfig())
	out.EndpointProtocol = direct.Enum_FromProto(mapCtx, in.GetEndpointProtocol())
	// MISSING: AuxiliaryVersions
	return out
}
func HiveMetastoreConfig_ToProto(mapCtx *direct.MapContext, in *krm.HiveMetastoreConfig) *pb.HiveMetastoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.HiveMetastoreConfig{}
	out.Version = direct.ValueOf(in.Version)
	out.ConfigOverrides = in.ConfigOverrides
	out.KerberosConfig = KerberosConfig_ToProto(mapCtx, in.KerberosConfig)
	out.EndpointProtocol = direct.Enum_ToProto[pb.HiveMetastoreConfig_EndpointProtocol](mapCtx, in.EndpointProtocol)
	// MISSING: AuxiliaryVersions
	return out
}
func KerberosConfig_FromProto(mapCtx *direct.MapContext, in *pb.KerberosConfig) *krm.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &krm.KerberosConfig{}
	out.Keytab = Secret_FromProto(mapCtx, in.GetKeytab())
	out.Principal = direct.LazyPtr(in.GetPrincipal())
	out.Krb5ConfigGcsURI = direct.LazyPtr(in.GetKrb5ConfigGcsUri())
	return out
}
func KerberosConfig_ToProto(mapCtx *direct.MapContext, in *krm.KerberosConfig) *pb.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &pb.KerberosConfig{}
	out.Keytab = Secret_ToProto(mapCtx, in.Keytab)
	out.Principal = direct.ValueOf(in.Principal)
	out.Krb5ConfigGcsUri = direct.ValueOf(in.Krb5ConfigGcsURI)
	return out
}
func Lake_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krm.Lake {
	if in == nil {
		return nil
	}
	out := &krm.Lake{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Lake_ToProto(mapCtx *direct.MapContext, in *krm.Lake) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceWindow) *krm.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceWindow{}
	out.HourOfDay = Int32Value_FromProto(mapCtx, in.GetHourOfDay())
	out.DayOfWeek = direct.Enum_FromProto(mapCtx, in.GetDayOfWeek())
	return out
}
func MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceWindow) *pb.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceWindow{}
	out.HourOfDay = Int32Value_ToProto(mapCtx, in.HourOfDay)
	out.DayOfWeek = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.DayOfWeek)
	return out
}
func MetadataExport_FromProto(mapCtx *direct.MapContext, in *pb.MetadataExport) *krm.MetadataExport {
	if in == nil {
		return nil
	}
	out := &krm.MetadataExport{}
	// MISSING: DestinationGcsURI
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DatabaseDumpType
	return out
}
func MetadataExport_ToProto(mapCtx *direct.MapContext, in *krm.MetadataExport) *pb.MetadataExport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataExport{}
	// MISSING: DestinationGcsURI
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DatabaseDumpType
	return out
}
func MetadataExportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataExport) *krm.MetadataExportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataExportObservedState{}
	out.DestinationGcsURI = direct.LazyPtr(in.GetDestinationGcsUri())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DatabaseDumpType = direct.Enum_FromProto(mapCtx, in.GetDatabaseDumpType())
	return out
}
func MetadataExportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataExportObservedState) *pb.MetadataExport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataExport{}
	if oneof := MetadataExportObservedState_DestinationGcsUri_ToProto(mapCtx, in.DestinationGcsURI); oneof != nil {
		out.Destination = oneof
	}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.MetadataExport_State](mapCtx, in.State)
	out.DatabaseDumpType = direct.Enum_ToProto[pb.DatabaseDumpSpec_Type](mapCtx, in.DatabaseDumpType)
	return out
}
func MetadataIntegration_FromProto(mapCtx *direct.MapContext, in *pb.MetadataIntegration) *krm.MetadataIntegration {
	if in == nil {
		return nil
	}
	out := &krm.MetadataIntegration{}
	out.DataCatalogConfig = DataCatalogConfig_FromProto(mapCtx, in.GetDataCatalogConfig())
	out.DataplexConfig = DataplexConfig_FromProto(mapCtx, in.GetDataplexConfig())
	return out
}
func MetadataIntegration_ToProto(mapCtx *direct.MapContext, in *krm.MetadataIntegration) *pb.MetadataIntegration {
	if in == nil {
		return nil
	}
	out := &pb.MetadataIntegration{}
	out.DataCatalogConfig = DataCatalogConfig_ToProto(mapCtx, in.DataCatalogConfig)
	out.DataplexConfig = DataplexConfig_ToProto(mapCtx, in.DataplexConfig)
	return out
}
func MetadataManagementActivity_FromProto(mapCtx *direct.MapContext, in *pb.MetadataManagementActivity) *krm.MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := &krm.MetadataManagementActivity{}
	// MISSING: MetadataExports
	// MISSING: Restores
	return out
}
func MetadataManagementActivity_ToProto(mapCtx *direct.MapContext, in *krm.MetadataManagementActivity) *pb.MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := &pb.MetadataManagementActivity{}
	// MISSING: MetadataExports
	// MISSING: Restores
	return out
}
func MetadataManagementActivityObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataManagementActivity) *krm.MetadataManagementActivityObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataManagementActivityObservedState{}
	out.MetadataExports = direct.Slice_FromProto(mapCtx, in.MetadataExports, MetadataExport_FromProto)
	out.Restores = direct.Slice_FromProto(mapCtx, in.Restores, Restore_FromProto)
	return out
}
func MetadataManagementActivityObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataManagementActivityObservedState) *pb.MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := &pb.MetadataManagementActivity{}
	out.MetadataExports = direct.Slice_ToProto(mapCtx, in.MetadataExports, MetadataExport_ToProto)
	out.Restores = direct.Slice_ToProto(mapCtx, in.Restores, Restore_ToProto)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.Consumers = direct.Slice_FromProto(mapCtx, in.Consumers, NetworkConfig_Consumer_FromProto)
	out.CustomRoutesEnabled = direct.LazyPtr(in.GetCustomRoutesEnabled())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Consumers = direct.Slice_ToProto(mapCtx, in.Consumers, NetworkConfig_Consumer_ToProto)
	out.CustomRoutesEnabled = direct.ValueOf(in.CustomRoutesEnabled)
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	out.Consumers = direct.Slice_FromProto(mapCtx, in.Consumers, NetworkConfig_ConsumerObservedState_FromProto)
	// MISSING: CustomRoutesEnabled
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Consumers = direct.Slice_ToProto(mapCtx, in.Consumers, NetworkConfig_ConsumerObservedState_ToProto)
	// MISSING: CustomRoutesEnabled
	return out
}
func NetworkConfig_Consumer_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig_Consumer) *krm.NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig_Consumer{}
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	// MISSING: EndpointURI
	// MISSING: EndpointLocation
	return out
}
func NetworkConfig_Consumer_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig_Consumer) *pb.NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig_Consumer{}
	if oneof := NetworkConfig_Consumer_Subnetwork_ToProto(mapCtx, in.Subnetwork); oneof != nil {
		out.VpcResource = oneof
	}
	// MISSING: EndpointURI
	// MISSING: EndpointLocation
	return out
}
func NetworkConfig_ConsumerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig_Consumer) *krm.NetworkConfig_ConsumerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig_ConsumerObservedState{}
	// MISSING: Subnetwork
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	out.EndpointLocation = direct.LazyPtr(in.GetEndpointLocation())
	return out
}
func NetworkConfig_ConsumerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig_ConsumerObservedState) *pb.NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig_Consumer{}
	// MISSING: Subnetwork
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	out.EndpointLocation = direct.ValueOf(in.EndpointLocation)
	return out
}
func Restore_FromProto(mapCtx *direct.MapContext, in *pb.Restore) *krm.Restore {
	if in == nil {
		return nil
	}
	out := &krm.Restore{}
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: Backup
	// MISSING: Type
	// MISSING: Details
	return out
}
func Restore_ToProto(mapCtx *direct.MapContext, in *krm.Restore) *pb.Restore {
	if in == nil {
		return nil
	}
	out := &pb.Restore{}
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: Backup
	// MISSING: Type
	// MISSING: Details
	return out
}
func RestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Restore) *krm.RestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RestoreObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Details = direct.LazyPtr(in.GetDetails())
	return out
}
func RestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RestoreObservedState) *pb.Restore {
	if in == nil {
		return nil
	}
	out := &pb.Restore{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Restore_State](mapCtx, in.State)
	out.Backup = direct.ValueOf(in.Backup)
	out.Type = direct.Enum_ToProto[pb.Restore_RestoreType](mapCtx, in.Type)
	out.Details = direct.ValueOf(in.Details)
	return out
}
func ScalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.ScalingConfig) *krm.ScalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ScalingConfig{}
	out.InstanceSize = direct.Enum_FromProto(mapCtx, in.GetInstanceSize())
	out.ScalingFactor = direct.LazyPtr(in.GetScalingFactor())
	return out
}
func ScalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.ScalingConfig) *pb.ScalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.ScalingConfig{}
	if oneof := ScalingConfig_InstanceSize_ToProto(mapCtx, in.InstanceSize); oneof != nil {
		out.ScalingModel = oneof
	}
	if oneof := ScalingConfig_ScalingFactor_ToProto(mapCtx, in.ScalingFactor); oneof != nil {
		out.ScalingModel = oneof
	}
	return out
}
func Secret_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.Secret {
	if in == nil {
		return nil
	}
	out := &krm.Secret{}
	out.CloudSecret = direct.LazyPtr(in.GetCloudSecret())
	return out
}
func Secret_ToProto(mapCtx *direct.MapContext, in *krm.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	if oneof := Secret_CloudSecret_ToProto(mapCtx, in.CloudSecret); oneof != nil {
		out.Value = oneof
	}
	return out
}
func Service_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.Service {
	if in == nil {
		return nil
	}
	out := &krm.Service{}
	out.HiveMetastoreConfig = HiveMetastoreConfig_FromProto(mapCtx, in.GetHiveMetastoreConfig())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: EndpointURI
	out.Port = direct.LazyPtr(in.GetPort())
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ArtifactGcsURI
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.MetadataIntegration = MetadataIntegration_FromProto(mapCtx, in.GetMetadataIntegration())
	out.MaintenanceWindow = MaintenanceWindow_FromProto(mapCtx, in.GetMaintenanceWindow())
	// MISSING: Uid
	// MISSING: MetadataManagementActivity
	out.ReleaseChannel = direct.Enum_FromProto(mapCtx, in.GetReleaseChannel())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.DatabaseType = direct.Enum_FromProto(mapCtx, in.GetDatabaseType())
	out.TelemetryConfig = TelemetryConfig_FromProto(mapCtx, in.GetTelemetryConfig())
	out.ScalingConfig = ScalingConfig_FromProto(mapCtx, in.GetScalingConfig())
	return out
}
func Service_ToProto(mapCtx *direct.MapContext, in *krm.Service) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	if oneof := HiveMetastoreConfig_ToProto(mapCtx, in.HiveMetastoreConfig); oneof != nil {
		out.MetastoreConfig = &pb.Service_HiveMetastoreConfig{HiveMetastoreConfig: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Network = direct.ValueOf(in.Network)
	// MISSING: EndpointURI
	out.Port = direct.ValueOf(in.Port)
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ArtifactGcsURI
	out.Tier = direct.Enum_ToProto[pb.Service_Tier](mapCtx, in.Tier)
	out.MetadataIntegration = MetadataIntegration_ToProto(mapCtx, in.MetadataIntegration)
	out.MaintenanceWindow = MaintenanceWindow_ToProto(mapCtx, in.MaintenanceWindow)
	// MISSING: Uid
	// MISSING: MetadataManagementActivity
	out.ReleaseChannel = direct.Enum_ToProto[pb.Service_ReleaseChannel](mapCtx, in.ReleaseChannel)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.DatabaseType = direct.Enum_ToProto[pb.Service_DatabaseType](mapCtx, in.DatabaseType)
	out.TelemetryConfig = TelemetryConfig_ToProto(mapCtx, in.TelemetryConfig)
	out.ScalingConfig = ScalingConfig_ToProto(mapCtx, in.ScalingConfig)
	return out
}
func ServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceObservedState{}
	// MISSING: HiveMetastoreConfig
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Network
	out.EndpointURI = direct.LazyPtr(in.GetEndpointUri())
	// MISSING: Port
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.ArtifactGcsURI = direct.LazyPtr(in.GetArtifactGcsUri())
	// MISSING: Tier
	// MISSING: MetadataIntegration
	// MISSING: MaintenanceWindow
	out.Uid = direct.LazyPtr(in.GetUid())
	out.MetadataManagementActivity = MetadataManagementActivity_FromProto(mapCtx, in.GetMetadataManagementActivity())
	// MISSING: ReleaseChannel
	// MISSING: EncryptionConfig
	out.NetworkConfig = NetworkConfigObservedState_FromProto(mapCtx, in.GetNetworkConfig())
	// MISSING: DatabaseType
	// MISSING: TelemetryConfig
	// MISSING: ScalingConfig
	return out
}
func ServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: HiveMetastoreConfig
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Network
	out.EndpointUri = direct.ValueOf(in.EndpointURI)
	// MISSING: Port
	out.State = direct.Enum_ToProto[pb.Service_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.ArtifactGcsUri = direct.ValueOf(in.ArtifactGcsURI)
	// MISSING: Tier
	// MISSING: MetadataIntegration
	// MISSING: MaintenanceWindow
	out.Uid = direct.ValueOf(in.Uid)
	out.MetadataManagementActivity = MetadataManagementActivity_ToProto(mapCtx, in.MetadataManagementActivity)
	// MISSING: ReleaseChannel
	// MISSING: EncryptionConfig
	out.NetworkConfig = NetworkConfigObservedState_ToProto(mapCtx, in.NetworkConfig)
	// MISSING: DatabaseType
	// MISSING: TelemetryConfig
	// MISSING: ScalingConfig
	return out
}
func TelemetryConfig_FromProto(mapCtx *direct.MapContext, in *pb.TelemetryConfig) *krm.TelemetryConfig {
	if in == nil {
		return nil
	}
	out := &krm.TelemetryConfig{}
	out.LogFormat = direct.Enum_FromProto(mapCtx, in.GetLogFormat())
	return out
}
func TelemetryConfig_ToProto(mapCtx *direct.MapContext, in *krm.TelemetryConfig) *pb.TelemetryConfig {
	if in == nil {
		return nil
	}
	out := &pb.TelemetryConfig{}
	out.LogFormat = direct.Enum_ToProto[pb.TelemetryConfig_LogFormat](mapCtx, in.LogFormat)
	return out
}
