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

// +generated:mapper
// krmv1beta1.group: memorystore.cnrm.cloud.google.com
// krmv1beta1.version: v1beta1
// proto.service: google.cloud.memorystore.v1

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	dayofweekpb "google.golang.org/genproto/googleapis/type/dayofweek"
	timeofdaypb "google.golang.org/genproto/googleapis/type/timeofday"
)

func AutomatedBackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupConfig) *krmv1beta1.AutomatedBackupConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AutomatedBackupConfig{}
	out.FixedFrequencySchedule = AutomatedBackupConfig_FixedFrequencySchedule_FromProto(mapCtx, in.GetFixedFrequencySchedule())
	out.AutomatedBackupMode = direct.Enum_FromProto(mapCtx, in.GetAutomatedBackupMode())
	out.Retention = direct.StringDuration_FromProto(mapCtx, in.GetRetention())
	return out
}
func AutomatedBackupConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AutomatedBackupConfig) *pb.AutomatedBackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupConfig{}
	if oneof := AutomatedBackupConfig_FixedFrequencySchedule_ToProto(mapCtx, in.FixedFrequencySchedule); oneof != nil {
		out.Schedule = &pb.AutomatedBackupConfig_FixedFrequencySchedule_{FixedFrequencySchedule: oneof}
	}
	out.AutomatedBackupMode = direct.Enum_ToProto[pb.AutomatedBackupConfig_AutomatedBackupMode](mapCtx, in.AutomatedBackupMode)
	out.Retention = direct.StringDuration_ToProto(mapCtx, in.Retention)
	return out
}
func AutomatedBackupConfig_FixedFrequencySchedule_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupConfig_FixedFrequencySchedule) *krmv1beta1.AutomatedBackupConfig_FixedFrequencySchedule {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AutomatedBackupConfig_FixedFrequencySchedule{}
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	return out
}
func AutomatedBackupConfig_FixedFrequencySchedule_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AutomatedBackupConfig_FixedFrequencySchedule) *pb.AutomatedBackupConfig_FixedFrequencySchedule {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupConfig_FixedFrequencySchedule{}
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	return out
}
func CrossInstanceReplicationConfig_FromProto(mapCtx *direct.MapContext, in *pb.CrossInstanceReplicationConfig) *krmv1beta1.CrossInstanceReplicationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CrossInstanceReplicationConfig{}
	out.InstanceRole = direct.Enum_FromProto(mapCtx, in.GetInstanceRole())
	out.PrimaryInstance = CrossInstanceReplicationConfig_RemoteInstance_FromProto(mapCtx, in.GetPrimaryInstance())
	out.SecondaryInstances = direct.Slice_FromProto(mapCtx, in.SecondaryInstances, CrossInstanceReplicationConfig_RemoteInstance_FromProto)
	return out
}
func CrossInstanceReplicationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CrossInstanceReplicationConfig) *pb.CrossInstanceReplicationConfig {
	if in == nil {
		return nil
	}
	out := &pb.CrossInstanceReplicationConfig{}
	out.InstanceRole = direct.Enum_ToProto[pb.CrossInstanceReplicationConfig_InstanceRole](mapCtx, in.InstanceRole)
	out.PrimaryInstance = CrossInstanceReplicationConfig_RemoteInstance_ToProto(mapCtx, in.PrimaryInstance)
	out.SecondaryInstances = direct.Slice_ToProto(mapCtx, in.SecondaryInstances, CrossInstanceReplicationConfig_RemoteInstance_ToProto)
	return out
}
func CrossInstanceReplicationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrossInstanceReplicationConfig) *krmv1beta1.CrossInstanceReplicationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CrossInstanceReplicationConfigObservedState{}
	out.PrimaryInstance = CrossInstanceReplicationConfig_RemoteInstanceObservedState_FromProto(mapCtx, in.GetPrimaryInstance())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Membership = CrossInstanceReplicationConfig_MembershipObservedState_FromProto(mapCtx, in.GetMembership())
	return out
}
func CrossInstanceReplicationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CrossInstanceReplicationConfigObservedState) *pb.CrossInstanceReplicationConfig {
	if in == nil {
		return nil
	}
	out := &pb.CrossInstanceReplicationConfig{}
	out.PrimaryInstance = CrossInstanceReplicationConfig_RemoteInstanceObservedState_ToProto(mapCtx, in.PrimaryInstance)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Membership = CrossInstanceReplicationConfig_MembershipObservedState_ToProto(mapCtx, in.Membership)
	return out
}
func CrossInstanceReplicationConfig_MembershipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrossInstanceReplicationConfig_Membership) *krmv1beta1.CrossInstanceReplicationConfig_MembershipObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CrossInstanceReplicationConfig_MembershipObservedState{}
	out.PrimaryInstance = CrossInstanceReplicationConfig_RemoteInstanceObservedState_FromProto(mapCtx, in.GetPrimaryInstance())
	out.SecondaryInstances = direct.Slice_FromProto(mapCtx, in.SecondaryInstances, CrossInstanceReplicationConfig_RemoteInstanceObservedState_FromProto)
	return out
}
func CrossInstanceReplicationConfig_MembershipObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CrossInstanceReplicationConfig_MembershipObservedState) *pb.CrossInstanceReplicationConfig_Membership {
	if in == nil {
		return nil
	}
	out := &pb.CrossInstanceReplicationConfig_Membership{}
	out.PrimaryInstance = CrossInstanceReplicationConfig_RemoteInstanceObservedState_ToProto(mapCtx, in.PrimaryInstance)
	out.SecondaryInstances = direct.Slice_ToProto(mapCtx, in.SecondaryInstances, CrossInstanceReplicationConfig_RemoteInstanceObservedState_ToProto)
	return out
}
func CrossInstanceReplicationConfig_RemoteInstance_FromProto(mapCtx *direct.MapContext, in *pb.CrossInstanceReplicationConfig_RemoteInstance) *krmv1beta1.CrossInstanceReplicationConfig_RemoteInstance {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CrossInstanceReplicationConfig_RemoteInstance{}
	out.Instance = direct.LazyPtr(in.GetInstance())
	return out
}
func CrossInstanceReplicationConfig_RemoteInstance_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CrossInstanceReplicationConfig_RemoteInstance) *pb.CrossInstanceReplicationConfig_RemoteInstance {
	if in == nil {
		return nil
	}
	out := &pb.CrossInstanceReplicationConfig_RemoteInstance{}
	out.Instance = direct.ValueOf(in.Instance)
	return out
}
func CrossInstanceReplicationConfig_RemoteInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrossInstanceReplicationConfig_RemoteInstance) *krmv1beta1.CrossInstanceReplicationConfig_RemoteInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CrossInstanceReplicationConfig_RemoteInstanceObservedState{}
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func CrossInstanceReplicationConfig_RemoteInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CrossInstanceReplicationConfig_RemoteInstanceObservedState) *pb.CrossInstanceReplicationConfig_RemoteInstance {
	if in == nil {
		return nil
	}
	out := &pb.CrossInstanceReplicationConfig_RemoteInstance{}
	out.Instance = direct.ValueOf(in.Instance)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func Instance_ConnectionDetail_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionDetail) *krmv1beta1.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_ConnectionDetail{}
	out.PscAutoConnection = PscAutoConnection_FromProto(mapCtx, in.GetPscAutoConnection())
	out.PscConnection = PscConnection_FromProto(mapCtx, in.GetPscConnection())
	return out
}
func Instance_ConnectionDetail_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_ConnectionDetail) *pb.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionDetail{}
	if oneof := PscAutoConnection_ToProto(mapCtx, in.PscAutoConnection); oneof != nil {
		out.Connection = &pb.Instance_ConnectionDetail_PscAutoConnection{PscAutoConnection: oneof}
	}
	if oneof := PscConnection_ToProto(mapCtx, in.PscConnection); oneof != nil {
		out.Connection = &pb.Instance_ConnectionDetail_PscConnection{PscConnection: oneof}
	}
	return out
}
func Instance_ConnectionDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionDetail) *krmv1beta1.Instance_ConnectionDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_ConnectionDetailObservedState{}
	out.PscAutoConnection = PscAutoConnectionObservedState_FromProto(mapCtx, in.GetPscAutoConnection())
	out.PscConnection = PscConnectionObservedState_FromProto(mapCtx, in.GetPscConnection())
	return out
}
func Instance_ConnectionDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_ConnectionDetailObservedState) *pb.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionDetail{}
	if oneof := PscAutoConnectionObservedState_ToProto(mapCtx, in.PscAutoConnection); oneof != nil {
		out.Connection = &pb.Instance_ConnectionDetail_PscAutoConnection{PscAutoConnection: oneof}
	}
	if oneof := PscConnectionObservedState_ToProto(mapCtx, in.PscConnection); oneof != nil {
		out.Connection = &pb.Instance_ConnectionDetail_PscConnection{PscConnection: oneof}
	}
	return out
}
func Instance_InstanceEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceEndpoint) *krmv1beta1.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, Instance_ConnectionDetail_FromProto)
	return out
}
func Instance_InstanceEndpoint_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_InstanceEndpoint) *pb.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Instance_ConnectionDetail_ToProto)
	return out
}
func Instance_InstanceEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceEndpoint) *krmv1beta1.Instance_InstanceEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_InstanceEndpointObservedState{}
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, Instance_ConnectionDetailObservedState_FromProto)
	return out
}
func Instance_InstanceEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_InstanceEndpointObservedState) *pb.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Instance_ConnectionDetailObservedState_ToProto)
	return out
}
func Instance_GCSBackupSource_FromProto(mapCtx *direct.MapContext, in *pb.Instance_GcsBackupSource) *krmv1beta1.Instance_GCSBackupSource {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_GCSBackupSource{}
	out.Uris = in.Uris
	return out
}
func Instance_GCSBackupSource_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_GCSBackupSource) *pb.Instance_GcsBackupSource {
	if in == nil {
		return nil
	}
	out := &pb.Instance_GcsBackupSource{}
	out.Uris = in.Uris
	return out
}
func Instance_ManagedBackupSource_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ManagedBackupSource) *krmv1beta1.Instance_ManagedBackupSource {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_ManagedBackupSource{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	return out
}
func Instance_ManagedBackupSource_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_ManagedBackupSource) *pb.Instance_ManagedBackupSource {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ManagedBackupSource{}
	out.Backup = direct.ValueOf(in.Backup)
	return out
}
func Instance_StateInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo) *krmv1beta1.Instance_StateInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_StateInfoObservedState{}
	out.UpdateInfo = Instance_StateInfo_UpdateInfoObservedState_FromProto(mapCtx, in.GetUpdateInfo())
	return out
}
func Instance_StateInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_StateInfoObservedState) *pb.Instance_StateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo{}
	if oneof := Instance_StateInfo_UpdateInfoObservedState_ToProto(mapCtx, in.UpdateInfo); oneof != nil {
		out.Info = &pb.Instance_StateInfo_UpdateInfo_{UpdateInfo: oneof}
	}
	return out
}
func Instance_StateInfo_UpdateInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_StateInfo_UpdateInfo) *krmv1beta1.Instance_StateInfo_UpdateInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Instance_StateInfo_UpdateInfoObservedState{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func Instance_StateInfo_UpdateInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Instance_StateInfo_UpdateInfoObservedState) *pb.Instance_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Instance_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func MaintenancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krmv1beta1.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MaintenancePolicy{}
	out.WeeklyMaintenanceWindow = direct.Slice_FromProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindow_FromProto)
	return out
}
func MaintenancePolicy_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MaintenancePolicy) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	out.WeeklyMaintenanceWindow = direct.Slice_ToProto(mapCtx, in.WeeklyMaintenanceWindow, WeeklyMaintenanceWindow_ToProto)
	return out
}
func MaintenancePolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krmv1beta1.MaintenancePolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MaintenancePolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func MaintenancePolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MaintenancePolicyObservedState) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func MaintenanceScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krmv1beta1.MaintenanceScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MaintenanceScheduleObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func MaintenanceScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MaintenanceScheduleObservedState) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func MemorystoreInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1beta1.MemorystoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MemorystoreInstanceObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateInfo = Instance_StateInfoObservedState_FromProto(mapCtx, in.GetStateInfo())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.NodeConfig = NodeConfigObservedState_FromProto(mapCtx, in.GetNodeConfig())
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Instance_InstanceEndpointObservedState_FromProto)
	out.MaintenancePolicy = MaintenancePolicyObservedState_FromProto(mapCtx, in.MaintenancePolicy)
	out.MaintenanceSchedule = MaintenanceScheduleObservedState_FromProto(mapCtx, in.MaintenanceSchedule)
	out.CrossInstanceReplicationConfig = CrossInstanceReplicationConfigObservedState_FromProto(mapCtx, in.CrossInstanceReplicationConfig)
	out.BackupCollection = direct.LazyPtr(in.GetBackupCollection())
	return out
}
func MemorystoreInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MemorystoreInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateInfo = Instance_StateInfoObservedState_ToProto(mapCtx, in.StateInfo)
	out.Uid = direct.ValueOf(in.Uid)
	out.NodeConfig = NodeConfigObservedState_ToProto(mapCtx, in.NodeConfig)
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Instance_InstanceEndpointObservedState_ToProto)
	return out
}
func MemorystoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmv1beta1.MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MemorystoreInstanceSpec{}
	out.Labels = in.Labels
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.EngineConfigs = in.EngineConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.GCSSource = Instance_GCSBackupSource_FromProto(mapCtx, in.GetGcsSource())
	out.ManagedBackupSource = Instance_ManagedBackupSource_FromProto(mapCtx, in.GetManagedBackupSource())
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Instance_InstanceEndpoint_FromProto)
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.MaintenancePolicy = MaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	out.CrossInstanceReplicationConfig = CrossInstanceReplicationConfig_FromProto(mapCtx, in.GetCrossInstanceReplicationConfig())
	out.AutomatedBackupConfig = AutomatedBackupConfig_FromProto(mapCtx, in.GetAutomatedBackupConfig())
	return out
}
func MemorystoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MemorystoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Labels = in.Labels
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.Instance_AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.EngineConfigs = in.EngineConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	if oneof := Instance_GCSBackupSource_ToProto(mapCtx, in.GCSSource); oneof != nil {
		out.ImportSources = &pb.Instance_GcsSource{GcsSource: oneof}
	}
	if oneof := Instance_ManagedBackupSource_ToProto(mapCtx, in.ManagedBackupSource); oneof != nil {
		out.ImportSources = &pb.Instance_ManagedBackupSource_{ManagedBackupSource: oneof}
	}
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Instance_InstanceEndpoint_ToProto)
	out.Mode = direct.Enum_ToProto[pb.Instance_Mode](mapCtx, in.Mode)
	out.MaintenancePolicy = MaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	out.CrossInstanceReplicationConfig = CrossInstanceReplicationConfig_ToProto(mapCtx, in.CrossInstanceReplicationConfig)
	out.AutomatedBackupConfig = AutomatedBackupConfig_ToProto(mapCtx, in.AutomatedBackupConfig)
	return out
}
func NodeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krmv1beta1.NodeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.NodeConfigObservedState{}
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	return out
}
func NodeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.NodeConfigObservedState) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.SizeGb = direct.ValueOf(in.SizeGB)
	return out
}
func PersistenceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig) *krmv1beta1.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PersistenceConfig{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.RdbConfig = PersistenceConfig_RdbConfig_FromProto(mapCtx, in.GetRdbConfig())
	out.AofConfig = PersistenceConfig_AofConfig_FromProto(mapCtx, in.GetAofConfig())
	return out
}
func PersistenceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PersistenceConfig) *pb.PersistenceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig{}
	out.Mode = direct.Enum_ToProto[pb.PersistenceConfig_PersistenceMode](mapCtx, in.Mode)
	out.RdbConfig = PersistenceConfig_RdbConfig_ToProto(mapCtx, in.RdbConfig)
	out.AofConfig = PersistenceConfig_AofConfig_ToProto(mapCtx, in.AofConfig)
	return out
}
func PersistenceConfig_AofConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig_AOFConfig) *krmv1beta1.PersistenceConfig_AofConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PersistenceConfig_AofConfig{}
	out.AppendFsync = direct.Enum_FromProto(mapCtx, in.GetAppendFsync())
	return out
}
func PersistenceConfig_AofConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PersistenceConfig_AofConfig) *pb.PersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig_AOFConfig{}
	out.AppendFsync = direct.Enum_ToProto[pb.PersistenceConfig_AOFConfig_AppendFsync](mapCtx, in.AppendFsync)
	return out
}
func PersistenceConfig_RdbConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersistenceConfig_RDBConfig) *krmv1beta1.PersistenceConfig_RdbConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PersistenceConfig_RdbConfig{}
	out.RdbSnapshotPeriod = direct.Enum_FromProto(mapCtx, in.GetRdbSnapshotPeriod())
	out.RdbSnapshotStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRdbSnapshotStartTime())
	return out
}
func PersistenceConfig_RdbConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PersistenceConfig_RdbConfig) *pb.PersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersistenceConfig_RDBConfig{}
	out.RdbSnapshotPeriod = direct.Enum_ToProto[pb.PersistenceConfig_RDBConfig_SnapshotPeriod](mapCtx, in.RdbSnapshotPeriod)
	out.RdbSnapshotStartTime = direct.StringTimestamp_ToProto(mapCtx, in.RdbSnapshotStartTime)
	return out
}
func PscAutoConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krmv1beta1.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PscAutoConnection{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetProjectId() != "" {
		out.ProjectRef = &refs.ProjectRef{External: in.GetProjectId()}
	}
	return out
}
func PscAutoConnection_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PscAutoConnection) *pb.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnection{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.ProjectRef != nil {
		project := refs.Project{}
		if err := project.FromExternal(in.ProjectRef.External); err != nil {
			mapCtx.Errorf("unable to get reference for the project: %v", err)
		}
		out.ProjectId = project.ProjectID
	}
	return out
}
func PscAutoConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krmv1beta1.PscAutoConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PscAutoConnectionObservedState{}
	out.Port = direct.LazyPtr(in.GetPort())
	out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	out.ForwardingRule = direct.LazyPtr(in.GetForwardingRule())
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	out.PscConnectionStatus = direct.Enum_FromProto(mapCtx, in.GetPscConnectionStatus())
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func PscAutoConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PscAutoConnectionObservedState) *pb.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnection{}
	if oneof := PscAutoConnectionObservedState_Port_ToProto(mapCtx, in.Port); oneof != nil {
		out.Ports = oneof
	}
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.IpAddress = direct.ValueOf(in.IpAddress)
	out.ForwardingRule = direct.ValueOf(in.ForwardingRule)
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	out.PscConnectionStatus = direct.Enum_ToProto[pb.PscConnectionStatus](mapCtx, in.PscConnectionStatus)
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}

func PscAutoConnectionObservedState_Port_ToProto(mapCtx *direct.MapContext, in *int32) *pb.PscAutoConnection_Port {
	out := &pb.PscAutoConnection_Port{}
	out.Port = direct.ValueOf(in)
	return out
}

func PscConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmv1beta1.PscConnection {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PscConnection{}
	if in.GetPscConnectionId() != "" {
		out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	}
	if in.GetIpAddress() != "" {
		out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetServiceAttachment() != "" {
		out.ServiceAttachmentRef = &refs.ComputeServiceAttachmentRef{External: in.GetServiceAttachment()}
	}
	return out
}
func PscConnection_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.IpAddress = direct.ValueOf(in.IpAddress)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.ServiceAttachmentRef != nil {
		out.ServiceAttachment = in.ServiceAttachmentRef.External
	}
	return out
}
func PscConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmv1beta1.PscConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.PscConnectionObservedState{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.PscConnectionStatus = direct.Enum_FromProto(mapCtx, in.GetPscConnectionStatus())
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func PscConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.PscConnectionObservedState) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.PscConnectionStatus = direct.Enum_ToProto[pb.PscConnectionStatus](mapCtx, in.PscConnectionStatus)
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func ZoneDistributionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ZoneDistributionConfig) *krmv1beta1.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ZoneDistributionConfig{}
	out.Zone = direct.LazyPtr(in.GetZone())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func ZoneDistributionConfig_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ZoneDistributionConfig) *pb.ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ZoneDistributionConfig{}
	out.Zone = direct.ValueOf(in.Zone)
	out.Mode = direct.Enum_ToProto[pb.ZoneDistributionConfig_ZoneDistributionMode](mapCtx, in.Mode)
	return out
}
func TimeOfDay_FromProto(mapCtx *direct.MapContext, in *timeofdaypb.TimeOfDay) *krmv1beta1.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.TimeOfDay{}
	out.Hours = direct.PtrTo(in.GetHours())
	out.Minutes = direct.PtrTo(in.GetMinutes())
	out.Seconds = direct.PtrTo(in.GetSeconds())
	out.Nanos = direct.PtrTo(in.GetNanos())
	return out
}
func TimeOfDay_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.TimeOfDay) *timeofdaypb.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &timeofdaypb.TimeOfDay{}
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}
func WeeklyMaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyMaintenanceWindow) *krmv1beta1.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	return out
}
func WeeklyMaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.WeeklyMaintenanceWindow) *pb.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_ToProto[dayofweekpb.DayOfWeek](mapCtx, in.Day)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	return out
}
