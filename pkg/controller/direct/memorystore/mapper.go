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

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PscAutoConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnection) *krm.PscAutoConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PscAutoConnectionSpec{}
	if in.Network != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.Network}
	}
	if in.ProjectId != "" {
		out.ProjectRef = &refs.ProjectRef{External: in.ProjectId}
	}
	return out
}
func PscAutoConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PscAutoConnectionSpec) *pb.PscAutoConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnection{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.ProjectRef != nil {
		out.ProjectId = in.ProjectRef.External
	}
	return out
}
func MemorystoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceSpec{}
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
	for _, pscConfig := range in.GetPscAutoConnections() {
		out.PscAutoConnectionsSpec = append(out.PscAutoConnectionsSpec, *PscAutoConnectionSpec_FromProto(mapCtx, pscConfig))
	}
	return out
}
func MemorystoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
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
	for _, pscConfig := range in.PscAutoConnectionsSpec {
		out.PscAutoConnections = append(out.PscAutoConnections, PscAutoConnectionSpec_ToProto(mapCtx, &pscConfig))
	}
	return out
}
func MemorystoreInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func MemorystoreInstanceStatusObservedState_FromProto(mapCtx *direct.MapContext, updated *pb.Instance) *krm.MemorystoreInstanceObservedState {
	if updated == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceObservedState{}
	out.State = direct.Enum_FromProto[pb.Instance_State](mapCtx, updated.State)
	return out
}
