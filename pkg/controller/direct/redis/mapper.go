// Copyright 2026 Google LLC
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
	redispb "cloud.google.com/go/redis/apiv1/redispb"
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	timeofdaypb "google.golang.org/genproto/googleapis/type/timeofday"
)

func TimeOfDay_v1beta1_FromProto(mapCtx *direct.MapContext, in *timeofdaypb.TimeOfDay) *krm.TimeOfDay {
	return TimeOfDay_FromProto(mapCtx, in)
}

func TimeOfDay_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.TimeOfDay) *timeofdaypb.TimeOfDay {
	return TimeOfDay_ToProto(mapCtx, in)
}

func PscConfigSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krm.PscConfigSpec {
	return PscConfigSpec_FromProto(mapCtx, in)
}

func PscConfigSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PscConfigSpec) *pb.PscConfig {
	return PscConfigSpec_ToProto(mapCtx, in)
}

func InstanceWeeklyMaintenanceWindow_v1beta1_FromProto(mapCtx *direct.MapContext, in *redispb.WeeklyMaintenanceWindow) *krm.InstanceWeeklyMaintenanceWindow {
	return InstanceWeeklyMaintenanceWindow_FromProto(mapCtx, in)
}

func InstanceWeeklyMaintenanceWindow_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceWeeklyMaintenanceWindow) *redispb.WeeklyMaintenanceWindow {
	return InstanceWeeklyMaintenanceWindow_ToProto(mapCtx, in)
}

func RedisInstanceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *redispb.Instance) *krm.RedisInstanceSpec {
	return RedisInstanceSpec_FromProto(mapCtx, in)
}

func RedisInstanceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RedisInstanceSpec) *redispb.Instance {
	return RedisInstanceSpec_ToProto(mapCtx, in)
}

func InstanceObservedStateStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *redispb.Instance) *krm.InstanceObservedStateStatus {
	return InstanceObservedStateStatus_FromProto(mapCtx, in)
}

func InstanceObservedStateStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedStateStatus) *redispb.Instance {
	return InstanceObservedStateStatus_ToProto(mapCtx, in)
}

func RedisClusterSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterSpec {
	return RedisClusterSpec_FromProto(mapCtx, in)
}

func RedisClusterSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterSpec) *pb.Cluster {
	return RedisClusterSpec_ToProto(mapCtx, in)
}

func Cluster_ManagedBackupSource_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ManagedBackupSource) *krm.Cluster_ManagedBackupSource {
	return Cluster_ManagedBackupSource_FromProto(mapCtx, in)
}

func Cluster_ManagedBackupSource_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ManagedBackupSource) *pb.Cluster_ManagedBackupSource {
	return Cluster_ManagedBackupSource_ToProto(mapCtx, in)
}

func PSCAutoConnection_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krm.PSCAutoConnection {
	return PSCAutoConnection_FromProto(mapCtx, in)
}

func PSCAutoConnection_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PSCAutoConnection) *pb.PscAutoConnection {
	return PSCAutoConnection_ToProto(mapCtx, in)
}

func PSCConnectionObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krm.PSCConnectionObservedState {
	return PSCConnectionObservedState_FromProto(mapCtx, in)
}

func PSCConnectionObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PSCConnectionObservedState) *pb.PscConnection {
	return PSCConnectionObservedState_ToProto(mapCtx, in)
}

var (
	RedisClusterObservedState_FromProto = RedisClusterObservedState_v1beta1_FromProto
	RedisClusterObservedState_ToProto   = RedisClusterObservedState_v1beta1_ToProto

	AutomatedBackupConfig_FromProto = AutomatedBackupConfig_v1beta1_FromProto
	AutomatedBackupConfig_ToProto   = AutomatedBackupConfig_v1beta1_ToProto

	ClusterMaintenancePolicy_FromProto = ClusterMaintenancePolicy_v1beta1_FromProto
	ClusterMaintenancePolicy_ToProto   = ClusterMaintenancePolicy_v1beta1_ToProto

	ClusterPersistenceConfig_FromProto = ClusterPersistenceConfig_v1beta1_FromProto
	ClusterPersistenceConfig_ToProto   = ClusterPersistenceConfig_v1beta1_ToProto

	CrossClusterReplicationConfig_FromProto = CrossClusterReplicationConfig_v1beta1_FromProto
	CrossClusterReplicationConfig_ToProto   = CrossClusterReplicationConfig_v1beta1_ToProto

	ZoneDistributionConfig_FromProto = ZoneDistributionConfig_v1beta1_FromProto
	ZoneDistributionConfig_ToProto   = ZoneDistributionConfig_v1beta1_ToProto

	ZoneDistributionConfigObservedState_FromProto = ZoneDistributionConfigObservedState_v1beta1_FromProto
	ZoneDistributionConfigObservedState_ToProto   = ZoneDistributionConfigObservedState_v1beta1_ToProto

	Cluster_GCSBackupSource_FromProto = Cluster_GCSBackupSource_v1beta1_FromProto
	Cluster_GCSBackupSource_ToProto   = Cluster_GCSBackupSource_v1beta1_ToProto

	ClusterEndpoint_FromProto = ClusterEndpoint_v1beta1_FromProto
	ClusterEndpoint_ToProto   = ClusterEndpoint_v1beta1_ToProto

	ConnectionDetail_FromProto = ConnectionDetail_v1beta1_FromProto
	ConnectionDetail_ToProto   = ConnectionDetail_v1beta1_ToProto

	InstanceMaintenancePolicy_FromProto = InstanceMaintenancePolicy_v1beta1_FromProto
	InstanceMaintenancePolicy_ToProto   = InstanceMaintenancePolicy_v1beta1_ToProto

	InstanceMaintenanceSchedule_FromProto = InstanceMaintenanceSchedule_v1beta1_FromProto
	InstanceMaintenanceSchedule_ToProto   = InstanceMaintenanceSchedule_v1beta1_ToProto

	InstancePersistenceConfig_FromProto = InstancePersistenceConfig_v1beta1_FromProto
	InstancePersistenceConfig_ToProto   = InstancePersistenceConfig_v1beta1_ToProto
)
