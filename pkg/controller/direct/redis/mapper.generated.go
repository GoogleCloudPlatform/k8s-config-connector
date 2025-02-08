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

package redis

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/redis/apiv1/redispb"
)
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.LocationID = direct.LazyPtr(in.GetLocationId())
	out.AlternativeLocationID = direct.LazyPtr(in.GetAlternativeLocationId())
	out.RedisVersion = direct.LazyPtr(in.GetRedisVersion())
	out.ReservedIPRange = direct.LazyPtr(in.GetReservedIpRange())
	out.SecondaryIPRange = direct.LazyPtr(in.GetSecondaryIpRange())
	// MISSING: Host
	// MISSING: Port
	// MISSING: CurrentLocationID
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: StatusMessage
	out.RedisConfigs = in.RedisConfigs
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.MemorySizeGB = direct.LazyPtr(in.GetMemorySizeGb())
	out.AuthorizedNetwork = direct.LazyPtr(in.GetAuthorizedNetwork())
	// MISSING: PersistenceIamIdentity
	out.ConnectMode = direct.Enum_FromProto(mapCtx, in.GetConnectMode())
	out.AuthEnabled = direct.LazyPtr(in.GetAuthEnabled())
	// MISSING: ServerCaCerts
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.MaintenancePolicy = MaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	// MISSING: MaintenanceSchedule
	out.ReplicaCount = direct.LazyPtr(in.GetReplicaCount())
	// MISSING: Nodes
	// MISSING: ReadEndpoint
	// MISSING: ReadEndpointPort
	out.ReadReplicasMode = direct.Enum_FromProto(mapCtx, in.GetReadReplicasMode())
	out.CustomerManagedKey = direct.LazyPtr(in.GetCustomerManagedKey())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.SuspensionReasons = direct.EnumSlice_FromProto(mapCtx, in.SuspensionReasons)
	out.MaintenanceVersion = direct.LazyPtr(in.GetMaintenanceVersion())
	out.AvailableMaintenanceVersions = in.AvailableMaintenanceVersions
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.LocationId = direct.ValueOf(in.LocationID)
	out.AlternativeLocationId = direct.ValueOf(in.AlternativeLocationID)
	out.RedisVersion = direct.ValueOf(in.RedisVersion)
	out.ReservedIpRange = direct.ValueOf(in.ReservedIPRange)
	out.SecondaryIpRange = direct.ValueOf(in.SecondaryIPRange)
	// MISSING: Host
	// MISSING: Port
	// MISSING: CurrentLocationID
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: StatusMessage
	out.RedisConfigs = in.RedisConfigs
	out.Tier = direct.Enum_ToProto[pb.Instance_Tier](mapCtx, in.Tier)
	out.MemorySizeGb = direct.ValueOf(in.MemorySizeGB)
	out.AuthorizedNetwork = direct.ValueOf(in.AuthorizedNetwork)
	// MISSING: PersistenceIamIdentity
	out.ConnectMode = direct.Enum_ToProto[pb.Instance_ConnectMode](mapCtx, in.ConnectMode)
	out.AuthEnabled = direct.ValueOf(in.AuthEnabled)
	// MISSING: ServerCaCerts
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.MaintenancePolicy = MaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	// MISSING: MaintenanceSchedule
	out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	// MISSING: Nodes
	// MISSING: ReadEndpoint
	// MISSING: ReadEndpointPort
	out.ReadReplicasMode = direct.Enum_ToProto[pb.Instance_ReadReplicasMode](mapCtx, in.ReadReplicasMode)
	out.CustomerManagedKey = direct.ValueOf(in.CustomerManagedKey)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.SuspensionReasons = direct.EnumSlice_ToProto[pb.Instance_SuspensionReason](mapCtx, in.SuspensionReasons)
	out.MaintenanceVersion = direct.ValueOf(in.MaintenanceVersion)
	out.AvailableMaintenanceVersions = in.AvailableMaintenanceVersions
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: LocationID
	// MISSING: AlternativeLocationID
	// MISSING: RedisVersion
	// MISSING: ReservedIPRange
	// MISSING: SecondaryIPRange
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.CurrentLocationID = direct.LazyPtr(in.GetCurrentLocationId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	// MISSING: RedisConfigs
	// MISSING: Tier
	// MISSING: MemorySizeGB
	// MISSING: AuthorizedNetwork
	out.PersistenceIamIdentity = direct.LazyPtr(in.GetPersistenceIamIdentity())
	// MISSING: ConnectMode
	// MISSING: AuthEnabled
	out.ServerCaCerts = direct.Slice_FromProto(mapCtx, in.ServerCaCerts, TlsCertificate_FromProto)
	// MISSING: TransitEncryptionMode
	out.MaintenancePolicy = MaintenancePolicyObservedState_FromProto(mapCtx, in.GetMaintenancePolicy())
	out.MaintenanceSchedule = MaintenanceSchedule_FromProto(mapCtx, in.GetMaintenanceSchedule())
	// MISSING: ReplicaCount
	out.Nodes = direct.Slice_FromProto(mapCtx, in.Nodes, NodeInfo_FromProto)
	out.ReadEndpoint = direct.LazyPtr(in.GetReadEndpoint())
	out.ReadEndpointPort = direct.LazyPtr(in.GetReadEndpointPort())
	// MISSING: ReadReplicasMode
	// MISSING: CustomerManagedKey
	out.PersistenceConfig = PersistenceConfigObservedState_FromProto(mapCtx, in.GetPersistenceConfig())
	// MISSING: SuspensionReasons
	// MISSING: MaintenanceVersion
	// MISSING: AvailableMaintenanceVersions
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: LocationID
	// MISSING: AlternativeLocationID
	// MISSING: RedisVersion
	// MISSING: ReservedIPRange
	// MISSING: SecondaryIPRange
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.CurrentLocationId = direct.ValueOf(in.CurrentLocationID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	// MISSING: RedisConfigs
	// MISSING: Tier
	// MISSING: MemorySizeGB
	// MISSING: AuthorizedNetwork
	out.PersistenceIamIdentity = direct.ValueOf(in.PersistenceIamIdentity)
	// MISSING: ConnectMode
	// MISSING: AuthEnabled
	out.ServerCaCerts = direct.Slice_ToProto(mapCtx, in.ServerCaCerts, TlsCertificate_ToProto)
	// MISSING: TransitEncryptionMode
	out.MaintenancePolicy = MaintenancePolicyObservedState_ToProto(mapCtx, in.MaintenancePolicy)
	out.MaintenanceSchedule = MaintenanceSchedule_ToProto(mapCtx, in.MaintenanceSchedule)
	// MISSING: ReplicaCount
	out.Nodes = direct.Slice_ToProto(mapCtx, in.Nodes, NodeInfo_ToProto)
	out.ReadEndpoint = direct.ValueOf(in.ReadEndpoint)
	out.ReadEndpointPort = direct.ValueOf(in.ReadEndpointPort)
	// MISSING: ReadReplicasMode
	// MISSING: CustomerManagedKey
	out.PersistenceConfig = PersistenceConfigObservedState_ToProto(mapCtx, in.PersistenceConfig)
	// MISSING: SuspensionReasons
	// MISSING: MaintenanceVersion
	// MISSING: AvailableMaintenanceVersions
	return out
}
func MaintenancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krm.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MaintenancePolicy{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.WeeklyMaintenanceWindow = direct.Slice_FromProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindow_FromProto)
	return out
}
func MaintenancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenancePolicy) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.WeeklyMaintenanceWindow = direct.Slice_ToProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindow_ToProto)
	return out
}
func MaintenancePolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krm.MaintenancePolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MaintenancePolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	out.WeeklyMaintenanceWindow = direct.Slice_FromProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindowObservedState_FromProto)
	return out
}
func MaintenancePolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MaintenancePolicyObservedState) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	out.WeeklyMaintenanceWindow = direct.Slice_ToProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindowObservedState_ToProto)
	return out
}
func MaintenanceSchedule_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceSchedule{}
	// MISSING: StartTime
	// MISSING: EndTime
	out.CanReschedule = direct.LazyPtr(in.GetCanReschedule())
	// MISSING: ScheduleDeadlineTime
	return out
}
func MaintenanceSchedule_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceSchedule) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	// MISSING: StartTime
	// MISSING: EndTime
	out.CanReschedule = direct.ValueOf(in.CanReschedule)
	// MISSING: ScheduleDeadlineTime
	return out
}
func MaintenanceScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceScheduleObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: CanReschedule
	out.ScheduleDeadlineTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleDeadlineTime())
	return out
}
func MaintenanceScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceScheduleObservedState) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: CanReschedule
	out.ScheduleDeadlineTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleDeadlineTime)
	return out
}
func NodeInfo_FromProto(mapCtx *direct.MapContext, in *pb.NodeInfo) *krm.NodeInfo {
	if in == nil {
		return nil
	}
	out := &krm.NodeInfo{}
	// MISSING: ID
	// MISSING: Zone
	return out
}
func NodeInfo_ToProto(mapCtx *direct.MapContext, in *krm.NodeInfo) *pb.NodeInfo {
	if in == nil {
		return nil
	}
	out := &pb.NodeInfo{}
	// MISSING: ID
	// MISSING: Zone
	return out
}
func NodeInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeInfo) *krm.NodeInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NodeInfoObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func NodeInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NodeInfoObservedState) *pb.NodeInfo {
	if in == nil {
		return nil
	}
	out := &pb.NodeInfo{}
	out.Id = direct.ValueOf(in.ID)
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
func PersistenceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig) *krm.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersistenceConfig{}
	out.PersistenceMode = direct.Enum_FromProto(mapCtx, in.GetPersistenceMode())
	out.RdbSnapshotPeriod = direct.Enum_FromProto(mapCtx, in.GetRdbSnapshotPeriod())
	// MISSING: RdbNextSnapshotTime
	out.RdbSnapshotStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRdbSnapshotStartTime())
	return out
}
func PersistenceConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersistenceConfig) *pb.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig{}
	out.PersistenceMode = direct.Enum_ToProto[pb.PersistenceConfig_PersistenceMode](mapCtx, in.PersistenceMode)
	out.RdbSnapshotPeriod = direct.Enum_ToProto[pb.PersistenceConfig_SnapshotPeriod](mapCtx, in.RdbSnapshotPeriod)
	// MISSING: RdbNextSnapshotTime
	out.RdbSnapshotStartTime = direct.StringTimestamp_ToProto(mapCtx, in.RdbSnapshotStartTime)
	return out
}
func PersistenceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig) *krm.PersistenceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PersistenceConfigObservedState{}
	// MISSING: PersistenceMode
	// MISSING: RdbSnapshotPeriod
	out.RdbNextSnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRdbNextSnapshotTime())
	// MISSING: RdbSnapshotStartTime
	return out
}
func PersistenceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PersistenceConfigObservedState) *pb.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig{}
	// MISSING: PersistenceMode
	// MISSING: RdbSnapshotPeriod
	out.RdbNextSnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.RdbNextSnapshotTime)
	// MISSING: RdbSnapshotStartTime
	return out
}
func RedisInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.RedisInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RedisInstanceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: LocationID
	// MISSING: AlternativeLocationID
	// MISSING: RedisVersion
	// MISSING: ReservedIPRange
	// MISSING: SecondaryIPRange
	// MISSING: Host
	// MISSING: Port
	// MISSING: CurrentLocationID
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: RedisConfigs
	// MISSING: Tier
	// MISSING: MemorySizeGB
	// MISSING: AuthorizedNetwork
	// MISSING: PersistenceIamIdentity
	// MISSING: ConnectMode
	// MISSING: AuthEnabled
	// MISSING: ServerCaCerts
	// MISSING: TransitEncryptionMode
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: ReplicaCount
	// MISSING: Nodes
	// MISSING: ReadEndpoint
	// MISSING: ReadEndpointPort
	// MISSING: ReadReplicasMode
	// MISSING: CustomerManagedKey
	// MISSING: PersistenceConfig
	// MISSING: SuspensionReasons
	// MISSING: MaintenanceVersion
	// MISSING: AvailableMaintenanceVersions
	return out
}
func RedisInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RedisInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: LocationID
	// MISSING: AlternativeLocationID
	// MISSING: RedisVersion
	// MISSING: ReservedIPRange
	// MISSING: SecondaryIPRange
	// MISSING: Host
	// MISSING: Port
	// MISSING: CurrentLocationID
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: RedisConfigs
	// MISSING: Tier
	// MISSING: MemorySizeGB
	// MISSING: AuthorizedNetwork
	// MISSING: PersistenceIamIdentity
	// MISSING: ConnectMode
	// MISSING: AuthEnabled
	// MISSING: ServerCaCerts
	// MISSING: TransitEncryptionMode
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: ReplicaCount
	// MISSING: Nodes
	// MISSING: ReadEndpoint
	// MISSING: ReadEndpointPort
	// MISSING: ReadReplicasMode
	// MISSING: CustomerManagedKey
	// MISSING: PersistenceConfig
	// MISSING: SuspensionReasons
	// MISSING: MaintenanceVersion
	// MISSING: AvailableMaintenanceVersions
	return out
}
func RedisInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.RedisInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisInstanceSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: LocationID
	// MISSING: AlternativeLocationID
	// MISSING: RedisVersion
	// MISSING: ReservedIPRange
	// MISSING: SecondaryIPRange
	// MISSING: Host
	// MISSING: Port
	// MISSING: CurrentLocationID
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: RedisConfigs
	// MISSING: Tier
	// MISSING: MemorySizeGB
	// MISSING: AuthorizedNetwork
	// MISSING: PersistenceIamIdentity
	// MISSING: ConnectMode
	// MISSING: AuthEnabled
	// MISSING: ServerCaCerts
	// MISSING: TransitEncryptionMode
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: ReplicaCount
	// MISSING: Nodes
	// MISSING: ReadEndpoint
	// MISSING: ReadEndpointPort
	// MISSING: ReadReplicasMode
	// MISSING: CustomerManagedKey
	// MISSING: PersistenceConfig
	// MISSING: SuspensionReasons
	// MISSING: MaintenanceVersion
	// MISSING: AvailableMaintenanceVersions
	return out
}
func RedisInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: LocationID
	// MISSING: AlternativeLocationID
	// MISSING: RedisVersion
	// MISSING: ReservedIPRange
	// MISSING: SecondaryIPRange
	// MISSING: Host
	// MISSING: Port
	// MISSING: CurrentLocationID
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: RedisConfigs
	// MISSING: Tier
	// MISSING: MemorySizeGB
	// MISSING: AuthorizedNetwork
	// MISSING: PersistenceIamIdentity
	// MISSING: ConnectMode
	// MISSING: AuthEnabled
	// MISSING: ServerCaCerts
	// MISSING: TransitEncryptionMode
	// MISSING: MaintenancePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: ReplicaCount
	// MISSING: Nodes
	// MISSING: ReadEndpoint
	// MISSING: ReadEndpointPort
	// MISSING: ReadReplicasMode
	// MISSING: CustomerManagedKey
	// MISSING: PersistenceConfig
	// MISSING: SuspensionReasons
	// MISSING: MaintenanceVersion
	// MISSING: AvailableMaintenanceVersions
	return out
}
func TlsCertificate_FromProto(mapCtx *direct.MapContext, in *pb.TlsCertificate) *krm.TlsCertificate {
	if in == nil {
		return nil
	}
	out := &krm.TlsCertificate{}
	out.SerialNumber = direct.LazyPtr(in.GetSerialNumber())
	out.Cert = direct.LazyPtr(in.GetCert())
	// MISSING: CreateTime
	// MISSING: ExpireTime
	out.Sha1Fingerprint = direct.LazyPtr(in.GetSha1Fingerprint())
	return out
}
func TlsCertificate_ToProto(mapCtx *direct.MapContext, in *krm.TlsCertificate) *pb.TlsCertificate {
	if in == nil {
		return nil
	}
	out := &pb.TlsCertificate{}
	out.SerialNumber = direct.ValueOf(in.SerialNumber)
	out.Cert = direct.ValueOf(in.Cert)
	// MISSING: CreateTime
	// MISSING: ExpireTime
	out.Sha1Fingerprint = direct.ValueOf(in.Sha1Fingerprint)
	return out
}
func TlsCertificateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TlsCertificate) *krm.TlsCertificateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TlsCertificateObservedState{}
	// MISSING: SerialNumber
	// MISSING: Cert
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Sha1Fingerprint
	return out
}
func TlsCertificateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TlsCertificateObservedState) *pb.TlsCertificate {
	if in == nil {
		return nil
	}
	out := &pb.TlsCertificate{}
	// MISSING: SerialNumber
	// MISSING: Cert
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Sha1Fingerprint
	return out
}
func WeeklyMaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyMaintenanceWindow) *krm.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	// MISSING: Duration
	return out
}
func WeeklyMaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.WeeklyMaintenanceWindow) *pb.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.Day)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	// MISSING: Duration
	return out
}
func WeeklyMaintenanceWindowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyMaintenanceWindow) *krm.WeeklyMaintenanceWindowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WeeklyMaintenanceWindowObservedState{}
	// MISSING: Day
	// MISSING: StartTime
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}
func WeeklyMaintenanceWindowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WeeklyMaintenanceWindowObservedState) *pb.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyMaintenanceWindow{}
	// MISSING: Day
	// MISSING: StartTime
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}
