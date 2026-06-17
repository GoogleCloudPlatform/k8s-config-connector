// Copyright 2024 Google LLC
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
	"time"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	timeofdaypb "google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Cluster_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func Cluster_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func RDBConfig_RdbSnapshotStartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func RDBConfig_RdbSnapshotStartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Timestamp_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	if in == nil {
		return nil
	}
	t := in.AsTime()
	s := t.Format(time.RFC3339Nano)
	return &s
}
func Timestamp_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	if in == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *in)
	if err != nil {
		mapCtx.Errorf("invalid timestamp %q", *in)
	}
	ts := timestamppb.New(t)
	return ts
}

func PscConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krm.PscConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.PscConfigSpec{}
	if in.Network != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.Network}
	}
	return out
}

func PscConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.PscConfigSpec) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	return out
}

func TimeOfDay_FromProto(mapCtx *direct.MapContext, in *timeofdaypb.TimeOfDay) *krm.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &krm.TimeOfDay{}
	out.Hours = direct.LazyPtr(in.GetHours())
	out.Minutes = direct.LazyPtr(in.GetMinutes())
	out.Seconds = direct.LazyPtr(in.GetSeconds())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func TimeOfDay_ToProto(mapCtx *direct.MapContext, in *krm.TimeOfDay) *timeofdaypb.TimeOfDay {
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

func RedisClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterSpec{}
	out.GCSSource = Cluster_GCSBackupSource_FromProto(mapCtx, in.GetGcsSource())
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = in.ShardCount
	out.PSCConfigs = direct.Slice_FromProto(mapCtx, in.PscConfigs, PscConfigSpec_FromProto)
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = ClusterPersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.RedisConfigs = in.RedisConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.MaintenancePolicy = ClusterMaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	if in.GetKmsKey() != "" {
		out.KMSKeyRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	out.AutomatedBackupConfig = AutomatedBackupConfig_FromProto(mapCtx, in.GetAutomatedBackupConfig())
	return out
}

func RedisClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	if oneof := Cluster_GCSBackupSource_ToProto(mapCtx, in.GCSSource); oneof != nil {
		out.ImportSources = &pb.Cluster_GcsSource{GcsSource: oneof}
	}
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = in.ShardCount
	out.PscConfigs = direct.Slice_ToProto(mapCtx, in.PSCConfigs, PscConfigSpec_ToProto)
	out.NodeType = direct.Enum_ToProto[pb.NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = ClusterPersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.RedisConfigs = in.RedisConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.MaintenancePolicy = ClusterMaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	if in.KMSKeyRef != nil {
		out.KmsKey = &in.KMSKeyRef.External
	}
	out.AutomatedBackupConfig = AutomatedBackupConfig_ToProto(mapCtx, in.AutomatedBackupConfig)
	return out
}
