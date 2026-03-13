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

package cluster

import (
	"time"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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

func CrossClusterReplicationConfig_FromProto(mapCtx *direct.MapContext, in *pb.CrossClusterReplicationConfig) *krm.CrossClusterReplicationConfig {
	if in == nil {
		return nil
	}
	out := &krm.CrossClusterReplicationConfig{}
	out.ClusterRole = direct.Enum_FromProto(mapCtx, in.GetClusterRole())
	out.PrimaryCluster = CrossClusterReplicationConfig_RemoteCluster_FromProto(mapCtx, in.GetPrimaryCluster())
	out.SecondaryClusters = direct.Slice_FromProto(mapCtx, in.SecondaryClusters, CrossClusterReplicationConfig_RemoteCluster_FromProto)
	return out
}

func CrossClusterReplicationConfig_ToProto(mapCtx *direct.MapContext, in *krm.CrossClusterReplicationConfig) *pb.CrossClusterReplicationConfig {
	if in == nil {
		return nil
	}
	out := &pb.CrossClusterReplicationConfig{}
	out.ClusterRole = direct.Enum_ToProto[pb.CrossClusterReplicationConfig_ClusterRole](mapCtx, in.ClusterRole)
	out.PrimaryCluster = CrossClusterReplicationConfig_RemoteCluster_ToProto(mapCtx, in.PrimaryCluster)
	out.SecondaryClusters = direct.Slice_ToProto(mapCtx, in.SecondaryClusters, CrossClusterReplicationConfig_RemoteCluster_ToProto)
	return out
}

func CrossClusterReplicationConfig_RemoteCluster_FromProto(mapCtx *direct.MapContext, in *pb.CrossClusterReplicationConfig_RemoteCluster) *krm.CrossClusterReplicationConfig_RemoteCluster {
	if in == nil {
		return nil
	}
	out := &krm.CrossClusterReplicationConfig_RemoteCluster{}
	if in.Cluster != "" {
		out.ClusterRef = &refs.RedisClusterRef{External: in.Cluster}
	}
	return out
}

func CrossClusterReplicationConfig_RemoteCluster_ToProto(mapCtx *direct.MapContext, in *krm.CrossClusterReplicationConfig_RemoteCluster) *pb.CrossClusterReplicationConfig_RemoteCluster {
	if in == nil {
		return nil
	}
	out := &pb.CrossClusterReplicationConfig_RemoteCluster{}
	if in.ClusterRef != nil {
		out.Cluster = in.ClusterRef.External
	}
	return out
}

func CrossClusterReplicationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrossClusterReplicationConfig) *krm.CrossClusterReplicationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CrossClusterReplicationConfigObservedState{}
	out.Membership = CrossClusterReplicationConfig_MembershipObservedState_FromProto(mapCtx, in.GetMembership())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func CrossClusterReplicationConfig_MembershipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrossClusterReplicationConfig_Membership) *krm.CrossClusterReplicationConfig_MembershipObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CrossClusterReplicationConfig_MembershipObservedState{}
	out.PrimaryCluster = CrossClusterReplicationConfig_RemoteClusterObservedState_FromProto(mapCtx, in.GetPrimaryCluster())
	out.SecondaryClusters = direct.Slice_FromProto(mapCtx, in.SecondaryClusters, CrossClusterReplicationConfig_RemoteClusterObservedState_FromProto)
	return out
}

func CrossClusterReplicationConfig_RemoteClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CrossClusterReplicationConfig_RemoteCluster) *krm.CrossClusterReplicationConfig_RemoteClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CrossClusterReplicationConfig_RemoteClusterObservedState{}
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}

func CrossClusterReplicationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CrossClusterReplicationConfigObservedState) *pb.CrossClusterReplicationConfig {
	if in == nil {
		return nil
	}
	out := &pb.CrossClusterReplicationConfig{}
	out.Membership = CrossClusterReplicationConfig_MembershipObservedState_ToProto(mapCtx, in.Membership)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func CrossClusterReplicationConfig_MembershipObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CrossClusterReplicationConfig_MembershipObservedState) *pb.CrossClusterReplicationConfig_Membership {
	if in == nil {
		return nil
	}
	out := &pb.CrossClusterReplicationConfig_Membership{}
	out.PrimaryCluster = CrossClusterReplicationConfig_RemoteClusterObservedState_ToProto(mapCtx, in.PrimaryCluster)
	out.SecondaryClusters = direct.Slice_ToProto(mapCtx, in.SecondaryClusters, CrossClusterReplicationConfig_RemoteClusterObservedState_ToProto)
	return out
}

func CrossClusterReplicationConfig_RemoteClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CrossClusterReplicationConfig_RemoteClusterObservedState) *pb.CrossClusterReplicationConfig_RemoteCluster {
	if in == nil {
		return nil
	}
	out := &pb.CrossClusterReplicationConfig_RemoteCluster{}
	out.Cluster = direct.ValueOf(in.Cluster)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
