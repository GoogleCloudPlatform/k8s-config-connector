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
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	krmredisv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func RedisClusterEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterEndpointObservedState{}
	out.ClusterEndpoints = direct.Slice_FromProto(mapCtx, in.ClusterEndpoints, ClusterEndpoint_ClusterEndpointObservedState_FromProto)
	return out
}

func RedisClusterEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterEndpointObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.ClusterEndpoints = direct.Slice_ToProto(mapCtx, in.ClusterEndpoints, ClusterEndpoint_ClusterEndpointObservedState_ToProto)
	return out
}

func RedisClusterEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.RedisClusterEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterEndpointSpec{}
	out.ClusterEndpoints = direct.Slice_FromProto(mapCtx, in.ClusterEndpoints, ClusterEndpoint_ClusterEndpoint_FromProto)
	return out
}

func RedisClusterEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterEndpointSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.ClusterEndpoints = direct.Slice_ToProto(mapCtx, in.ClusterEndpoints, ClusterEndpoint_ClusterEndpoint_ToProto)
	return out
}
func ClusterEndpoint_ClusterEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.ClusterEndpoint) *krmredisv1alpha1.ClusterEndpoint_ClusterEndpoint {
	if in == nil {
		return nil
	}
	out := &krmredisv1alpha1.ClusterEndpoint_ClusterEndpoint{}
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, ClusterEndpoint_ConnectionDetail_FromProto)
	return out
}
func ClusterEndpoint_ClusterEndpoint_ToProto(mapCtx *direct.MapContext, in *krmredisv1alpha1.ClusterEndpoint_ClusterEndpoint) *pb.ClusterEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.ClusterEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, ClusterEndpoint_ConnectionDetail_ToProto)
	return out
}
func ClusterEndpoint_ClusterEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClusterEndpoint) *krmredisv1alpha1.ClusterEndpoint_ClusterEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krmredisv1alpha1.ClusterEndpoint_ClusterEndpointObservedState{}
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, ClusterEndpoint_ConnectionDetailObservedState_FromProto)
	return out
}
func ClusterEndpoint_ClusterEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krmredisv1alpha1.ClusterEndpoint_ClusterEndpointObservedState) *pb.ClusterEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.ClusterEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, ClusterEndpoint_ConnectionDetailObservedState_ToProto)
	return out
}
func ClusterEndpoint_ConnectionDetail_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionDetail) *krmredisv1alpha1.ClusterEndpoint_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &krmredisv1alpha1.ClusterEndpoint_ConnectionDetail{}
	// MISSING: PSCAutoConnection
	out.PSCConnection = ClusterEndpoint_PSCConnection_FromProto(mapCtx, in.GetPscConnection())
	return out
}
func ClusterEndpoint_ConnectionDetail_ToProto(mapCtx *direct.MapContext, in *krmredisv1alpha1.ClusterEndpoint_ConnectionDetail) *pb.ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionDetail{}
	// MISSING: PSCAutoConnection
	if oneof := ClusterEndpoint_PSCConnection_ToProto(mapCtx, in.PSCConnection); oneof != nil {
		out.Connection = &pb.ConnectionDetail_PscConnection{PscConnection: oneof}
	}
	return out
}
func ClusterEndpoint_ConnectionDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionDetail) *krmredisv1alpha1.ClusterEndpoint_ConnectionDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krmredisv1alpha1.ClusterEndpoint_ConnectionDetailObservedState{}
	// MISSING: PSCAutoConnection
	out.PSCConnection = ClusterEndpoint_PSCConnectionObservedState_FromProto(mapCtx, in.GetPscConnection())
	return out
}
func ClusterEndpoint_ConnectionDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krmredisv1alpha1.ClusterEndpoint_ConnectionDetailObservedState) *pb.ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionDetail{}
	// MISSING: PSCAutoConnection
	if oneof := ClusterEndpoint_PSCConnectionObservedState_ToProto(mapCtx, in.PSCConnection); oneof != nil {
		out.Connection = &pb.ConnectionDetail_PscConnection{PscConnection: oneof}
	}
	return out
}
func ClusterEndpoint_PSCConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmredisv1alpha1.ClusterEndpoint_PSCConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmredisv1alpha1.ClusterEndpoint_PSCConnectionObservedState{}
	// MISSING: PSCConnectionID
	// MISSING: Address
	// MISSING: ForwardingRule
	// MISSING: ProjectID
	// MISSING: Network
	// MISSING: ServiceAttachment
	out.PSCConnectionStatus = direct.Enum_FromProto(mapCtx, in.GetPscConnectionStatus())
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func ClusterEndpoint_PSCConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmredisv1alpha1.ClusterEndpoint_PSCConnectionObservedState) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	// MISSING: PSCConnectionID
	// MISSING: Address
	// MISSING: ForwardingRule
	// MISSING: ProjectID
	// MISSING: Network
	// MISSING: ServiceAttachment
	out.PscConnectionStatus = direct.Enum_ToProto[pb.PscConnectionStatus](mapCtx, in.PSCConnectionStatus)
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func ClusterEndpoint_PSCConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmredisv1alpha1.ClusterEndpoint_PSCConnection {
	if in == nil {
		return nil
	}
	out := &krmredisv1alpha1.ClusterEndpoint_PSCConnection{}
	out.PSCConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	if in.GetAddress() != "" {
		out.AddressRef = &krmcomputev1beta1.ComputeAddressRef{External: in.GetAddress()}
	}
	if in.GetForwardingRule() != "" {
		out.ForwardingRuleRef = &krmcomputev1beta1.ForwardingRuleRef{External: in.GetForwardingRule()}
	}
	if in.GetProjectId() != "" {
		out.ProjectRef = &refsv1beta1.ProjectRef{External: in.GetProjectId()}
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetServiceAttachment() != "" {
		out.ServiceAttachmentRef = &refsv1beta1.ComputeServiceAttachmentRef{External: in.GetServiceAttachment()}
	}
	// MISSING: PSCConnectionStatus
	// MISSING: ConnectionType
	return out
}
func ClusterEndpoint_PSCConnection_ToProto(mapCtx *direct.MapContext, in *krmredisv1alpha1.ClusterEndpoint_PSCConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.PscConnectionId = direct.ValueOf(in.PSCConnectionID)
	if in.AddressRef != nil {
		out.Address = in.AddressRef.External
	}
	if in.ForwardingRuleRef != nil {
		out.ForwardingRule = in.ForwardingRuleRef.External
	}
	if in.ProjectRef != nil {
		out.ProjectId = in.ProjectRef.External
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.ServiceAttachmentRef != nil {
		out.ServiceAttachment = in.ServiceAttachmentRef.External
	}
	// MISSING: PSCConnectionStatus
	// MISSING: ConnectionType
	return out
}
