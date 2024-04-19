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

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// func ClusterSpec_FromProto(ctx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterSpec {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.RedisClusterSpec{}
// 	out.ReplicaCount = (in.ReplicaCount)
// 	out.AuthorizationMode = direct.Enum_FromProto(ctx, in.AuthorizationMode)
// 	out.TransitEncryptionMode = direct.Enum_FromProto(ctx, in.TransitEncryptionMode)
// 	out.ShardCount = (in.ShardCount)
// 	out.PscConfigs = direct.Slice_FromProto(ctx, in.PscConfigs, PscConfig_FromProto)
// 	return out
// }
// func ClusterSpec_ToProto(ctx *direct.MapContext, in *krm.RedisClusterSpec) *pb.Cluster {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster{}
// 	out.ReplicaCount = (in.ReplicaCount)
// 	out.AuthorizationMode = direct.Enum_ToProto[pb.AuthorizationMode](ctx, in.AuthorizationMode)
// 	out.TransitEncryptionMode = direct.Enum_ToProto[pb.TransitEncryptionMode](ctx, in.TransitEncryptionMode)
// 	out.ShardCount = (in.ShardCount)
// 	out.PscConfigs = direct.Slice_ToProto(ctx, in.PscConfigs, PscConfig_ToProto)
// 	return out
// }

// func ClusterState_FromProto(ctx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterObservedState {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.RedisClusterObservedState{}
// 	// out.CreateTime = Timestamp_FromProto(ctx, in.CreateTime)
// 	// out.State = Enum_FromProto(ctx, &in.State)
// 	// out.Uid = LazyPtr(in.Uid)
// 	out.SizeGb = (in.SizeGb)
// 	out.DiscoveryEndpoints = direct.Slice_FromProto(ctx, in.DiscoveryEndpoints, DiscoveryEndpoint_FromProto)
// 	out.PscConnections = direct.Slice_FromProto(ctx, in.PscConnections, PscConnection_FromProto)

// 	out.StateInfo = Cluster_StateInfo_FromProto(ctx, in.StateInfo)
// 	return out
// }
// func ClusterState_ToProto(ctx *direct.MapContext, in *krm.RedisClusterObservedState) *pb.Cluster {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster{}
// 	// out.CreateTime = Timestamp_ToProto(ctx, in.CreateTime)
// 	// out.State = Enum_FromProto(ctx, &in.State)
// 	// out.Uid = LazyPtr(in.Uid)
// 	out.SizeGb = (in.SizeGb)
// 	out.DiscoveryEndpoints = direct.Slice_ToProto(ctx, in.DiscoveryEndpoints, DiscoveryEndpoint_ToProto)
// 	out.PscConnections = direct.Slice_ToProto(ctx, in.PscConnections, PscConnection_ToProto)
// 	out.StateInfo = Cluster_StateInfo_ToProto(ctx, in.StateInfo)
// 	return out
// }

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

func PscConfig_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krm.PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.PscConfig{}
	if in.Network != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.Network}
	}
	return out
}

func PscConfig_ToProto(mapCtx *direct.MapContext, in *krm.PscConfig) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	return out
}
