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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
)
func CloudVmCluster_FromProto(mapCtx *direct.MapContext, in *pb.CloudVmCluster) *krm.CloudVmCluster {
	if in == nil {
		return nil
	}
	out := &krm.CloudVmCluster{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExadataInfrastructure = direct.LazyPtr(in.GetExadataInfrastructure())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: GcpOracleZone
	out.Properties = CloudVmClusterProperties_FromProto(mapCtx, in.GetProperties())
	out.Labels = in.Labels
	// MISSING: CreateTime
	out.Cidr = direct.LazyPtr(in.GetCidr())
	out.BackupSubnetCidr = direct.LazyPtr(in.GetBackupSubnetCidr())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func CloudVmCluster_ToProto(mapCtx *direct.MapContext, in *krm.CloudVmCluster) *pb.CloudVmCluster {
	if in == nil {
		return nil
	}
	out := &pb.CloudVmCluster{}
	out.Name = direct.ValueOf(in.Name)
	out.ExadataInfrastructure = direct.ValueOf(in.ExadataInfrastructure)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: GcpOracleZone
	out.Properties = CloudVmClusterProperties_ToProto(mapCtx, in.Properties)
	out.Labels = in.Labels
	// MISSING: CreateTime
	out.Cidr = direct.ValueOf(in.Cidr)
	out.BackupSubnetCidr = direct.ValueOf(in.BackupSubnetCidr)
	out.Network = direct.ValueOf(in.Network)
	return out
}
func CloudVmClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudVmCluster) *krm.CloudVmClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudVmClusterObservedState{}
	// MISSING: Name
	// MISSING: ExadataInfrastructure
	// MISSING: DisplayName
	out.GcpOracleZone = direct.LazyPtr(in.GetGcpOracleZone())
	out.Properties = CloudVmClusterPropertiesObservedState_FromProto(mapCtx, in.GetProperties())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Cidr
	// MISSING: BackupSubnetCidr
	// MISSING: Network
	return out
}
func CloudVmClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudVmClusterObservedState) *pb.CloudVmCluster {
	if in == nil {
		return nil
	}
	out := &pb.CloudVmCluster{}
	// MISSING: Name
	// MISSING: ExadataInfrastructure
	// MISSING: DisplayName
	out.GcpOracleZone = direct.ValueOf(in.GcpOracleZone)
	out.Properties = CloudVmClusterPropertiesObservedState_ToProto(mapCtx, in.Properties)
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Cidr
	// MISSING: BackupSubnetCidr
	// MISSING: Network
	return out
}
func CloudVmClusterProperties_FromProto(mapCtx *direct.MapContext, in *pb.CloudVmClusterProperties) *krm.CloudVmClusterProperties {
	if in == nil {
		return nil
	}
	out := &krm.CloudVmClusterProperties{}
	// MISSING: Ocid
	out.LicenseType = direct.Enum_FromProto(mapCtx, in.GetLicenseType())
	out.GiVersion = direct.LazyPtr(in.GetGiVersion())
	out.TimeZone = TimeZone_FromProto(mapCtx, in.GetTimeZone())
	out.SSHPublicKeys = in.SshPublicKeys
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	// MISSING: Shape
	out.OcpuCount = direct.LazyPtr(in.GetOcpuCount())
	out.MemorySizeGB = direct.LazyPtr(in.GetMemorySizeGb())
	out.DbNodeStorageSizeGB = direct.LazyPtr(in.GetDbNodeStorageSizeGb())
	// MISSING: StorageSizeGB
	out.DataStorageSizeTb = direct.LazyPtr(in.GetDataStorageSizeTb())
	out.DiskRedundancy = direct.Enum_FromProto(mapCtx, in.GetDiskRedundancy())
	out.SparseDiskgroupEnabled = direct.LazyPtr(in.GetSparseDiskgroupEnabled())
	out.LocalBackupEnabled = direct.LazyPtr(in.GetLocalBackupEnabled())
	out.HostnamePrefix = direct.LazyPtr(in.GetHostnamePrefix())
	out.DiagnosticsDataCollectionOptions = DataCollectionOptions_FromProto(mapCtx, in.GetDiagnosticsDataCollectionOptions())
	// MISSING: State
	// MISSING: ScanListenerPortTcp
	// MISSING: ScanListenerPortTcpSsl
	// MISSING: Domain
	// MISSING: ScanDns
	// MISSING: Hostname
	out.CpuCoreCount = direct.LazyPtr(in.GetCpuCoreCount())
	out.SystemVersion = direct.LazyPtr(in.GetSystemVersion())
	// MISSING: ScanIPIds
	// MISSING: ScanDnsRecordID
	// MISSING: OciURL
	out.DbServerOcids = in.DbServerOcids
	// MISSING: CompartmentID
	// MISSING: DnsListenerIP
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	return out
}
func CloudVmClusterProperties_ToProto(mapCtx *direct.MapContext, in *krm.CloudVmClusterProperties) *pb.CloudVmClusterProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudVmClusterProperties{}
	// MISSING: Ocid
	out.LicenseType = direct.Enum_ToProto[pb.CloudVmClusterProperties_LicenseType](mapCtx, in.LicenseType)
	out.GiVersion = direct.ValueOf(in.GiVersion)
	out.TimeZone = TimeZone_ToProto(mapCtx, in.TimeZone)
	out.SshPublicKeys = in.SSHPublicKeys
	out.NodeCount = direct.ValueOf(in.NodeCount)
	// MISSING: Shape
	out.OcpuCount = direct.ValueOf(in.OcpuCount)
	out.MemorySizeGb = direct.ValueOf(in.MemorySizeGB)
	out.DbNodeStorageSizeGb = direct.ValueOf(in.DbNodeStorageSizeGB)
	// MISSING: StorageSizeGB
	out.DataStorageSizeTb = direct.ValueOf(in.DataStorageSizeTb)
	out.DiskRedundancy = direct.Enum_ToProto[pb.CloudVmClusterProperties_DiskRedundancy](mapCtx, in.DiskRedundancy)
	out.SparseDiskgroupEnabled = direct.ValueOf(in.SparseDiskgroupEnabled)
	out.LocalBackupEnabled = direct.ValueOf(in.LocalBackupEnabled)
	out.HostnamePrefix = direct.ValueOf(in.HostnamePrefix)
	out.DiagnosticsDataCollectionOptions = DataCollectionOptions_ToProto(mapCtx, in.DiagnosticsDataCollectionOptions)
	// MISSING: State
	// MISSING: ScanListenerPortTcp
	// MISSING: ScanListenerPortTcpSsl
	// MISSING: Domain
	// MISSING: ScanDns
	// MISSING: Hostname
	out.CpuCoreCount = direct.ValueOf(in.CpuCoreCount)
	out.SystemVersion = direct.ValueOf(in.SystemVersion)
	// MISSING: ScanIPIds
	// MISSING: ScanDnsRecordID
	// MISSING: OciURL
	out.DbServerOcids = in.DbServerOcids
	// MISSING: CompartmentID
	// MISSING: DnsListenerIP
	out.ClusterName = direct.ValueOf(in.ClusterName)
	return out
}
func CloudVmClusterPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudVmClusterProperties) *krm.CloudVmClusterPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudVmClusterPropertiesObservedState{}
	out.Ocid = direct.LazyPtr(in.GetOcid())
	// MISSING: LicenseType
	// MISSING: GiVersion
	// MISSING: TimeZone
	// MISSING: SSHPublicKeys
	// MISSING: NodeCount
	out.Shape = direct.LazyPtr(in.GetShape())
	// MISSING: OcpuCount
	// MISSING: MemorySizeGB
	// MISSING: DbNodeStorageSizeGB
	out.StorageSizeGB = direct.LazyPtr(in.GetStorageSizeGb())
	// MISSING: DataStorageSizeTb
	// MISSING: DiskRedundancy
	// MISSING: SparseDiskgroupEnabled
	// MISSING: LocalBackupEnabled
	// MISSING: HostnamePrefix
	// MISSING: DiagnosticsDataCollectionOptions
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ScanListenerPortTcp = direct.LazyPtr(in.GetScanListenerPortTcp())
	out.ScanListenerPortTcpSsl = direct.LazyPtr(in.GetScanListenerPortTcpSsl())
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.ScanDns = direct.LazyPtr(in.GetScanDns())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	// MISSING: CpuCoreCount
	// MISSING: SystemVersion
	out.ScanIPIds = in.ScanIpIds
	out.ScanDnsRecordID = direct.LazyPtr(in.GetScanDnsRecordId())
	out.OciURL = direct.LazyPtr(in.GetOciUrl())
	// MISSING: DbServerOcids
	out.CompartmentID = direct.LazyPtr(in.GetCompartmentId())
	out.DnsListenerIP = direct.LazyPtr(in.GetDnsListenerIp())
	// MISSING: ClusterName
	return out
}
func CloudVmClusterPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudVmClusterPropertiesObservedState) *pb.CloudVmClusterProperties {
	if in == nil {
		return nil
	}
	out := &pb.CloudVmClusterProperties{}
	out.Ocid = direct.ValueOf(in.Ocid)
	// MISSING: LicenseType
	// MISSING: GiVersion
	// MISSING: TimeZone
	// MISSING: SSHPublicKeys
	// MISSING: NodeCount
	out.Shape = direct.ValueOf(in.Shape)
	// MISSING: OcpuCount
	// MISSING: MemorySizeGB
	// MISSING: DbNodeStorageSizeGB
	out.StorageSizeGb = direct.ValueOf(in.StorageSizeGB)
	// MISSING: DataStorageSizeTb
	// MISSING: DiskRedundancy
	// MISSING: SparseDiskgroupEnabled
	// MISSING: LocalBackupEnabled
	// MISSING: HostnamePrefix
	// MISSING: DiagnosticsDataCollectionOptions
	out.State = direct.Enum_ToProto[pb.CloudVmClusterProperties_State](mapCtx, in.State)
	out.ScanListenerPortTcp = direct.ValueOf(in.ScanListenerPortTcp)
	out.ScanListenerPortTcpSsl = direct.ValueOf(in.ScanListenerPortTcpSsl)
	out.Domain = direct.ValueOf(in.Domain)
	out.ScanDns = direct.ValueOf(in.ScanDns)
	out.Hostname = direct.ValueOf(in.Hostname)
	// MISSING: CpuCoreCount
	// MISSING: SystemVersion
	out.ScanIpIds = in.ScanIPIds
	out.ScanDnsRecordId = direct.ValueOf(in.ScanDnsRecordID)
	out.OciUrl = direct.ValueOf(in.OciURL)
	// MISSING: DbServerOcids
	out.CompartmentId = direct.ValueOf(in.CompartmentID)
	out.DnsListenerIp = direct.ValueOf(in.DnsListenerIP)
	// MISSING: ClusterName
	return out
}
func DataCollectionOptions_FromProto(mapCtx *direct.MapContext, in *pb.DataCollectionOptions) *krm.DataCollectionOptions {
	if in == nil {
		return nil
	}
	out := &krm.DataCollectionOptions{}
	out.DiagnosticsEventsEnabled = direct.LazyPtr(in.GetDiagnosticsEventsEnabled())
	out.HealthMonitoringEnabled = direct.LazyPtr(in.GetHealthMonitoringEnabled())
	out.IncidentLogsEnabled = direct.LazyPtr(in.GetIncidentLogsEnabled())
	return out
}
func DataCollectionOptions_ToProto(mapCtx *direct.MapContext, in *krm.DataCollectionOptions) *pb.DataCollectionOptions {
	if in == nil {
		return nil
	}
	out := &pb.DataCollectionOptions{}
	out.DiagnosticsEventsEnabled = direct.ValueOf(in.DiagnosticsEventsEnabled)
	out.HealthMonitoringEnabled = direct.ValueOf(in.HealthMonitoringEnabled)
	out.IncidentLogsEnabled = direct.ValueOf(in.IncidentLogsEnabled)
	return out
}
func OracledatabaseCloudVmClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudVmCluster) *krm.OracledatabaseCloudVmClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseCloudVmClusterObservedState{}
	// MISSING: Name
	// MISSING: ExadataInfrastructure
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: Cidr
	// MISSING: BackupSubnetCidr
	// MISSING: Network
	return out
}
func OracledatabaseCloudVmClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseCloudVmClusterObservedState) *pb.CloudVmCluster {
	if in == nil {
		return nil
	}
	out := &pb.CloudVmCluster{}
	// MISSING: Name
	// MISSING: ExadataInfrastructure
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: Cidr
	// MISSING: BackupSubnetCidr
	// MISSING: Network
	return out
}
func OracledatabaseCloudVmClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudVmCluster) *krm.OracledatabaseCloudVmClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseCloudVmClusterSpec{}
	// MISSING: Name
	// MISSING: ExadataInfrastructure
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: Cidr
	// MISSING: BackupSubnetCidr
	// MISSING: Network
	return out
}
func OracledatabaseCloudVmClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseCloudVmClusterSpec) *pb.CloudVmCluster {
	if in == nil {
		return nil
	}
	out := &pb.CloudVmCluster{}
	// MISSING: Name
	// MISSING: ExadataInfrastructure
	// MISSING: DisplayName
	// MISSING: GcpOracleZone
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: Cidr
	// MISSING: BackupSubnetCidr
	// MISSING: Network
	return out
}
