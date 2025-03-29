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

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

func EdgeContainerClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.EdgeContainerClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterSpec{}
	out.Labels = in.Labels
	out.Fleet = Fleet_FromProto(mapCtx, in.GetFleet())
	out.Networking = ClusterNetworking_FromProto(mapCtx, in.GetNetworking())
	out.Authorization = Authorization_FromProto(mapCtx, in.GetAuthorization())
	out.DefaultMaxPodsPerNode = direct.LazyPtr(in.GetDefaultMaxPodsPerNode())
	out.MaintenancePolicy = MaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	out.ControlPlane = Cluster_ControlPlane_FromProto(mapCtx, in.GetControlPlane())
	out.SystemAddonsConfig = Cluster_SystemAddonsConfig_FromProto(mapCtx, in.GetSystemAddonsConfig())
	out.ExternalLoadBalancerIPV4AddressPools = in.ExternalLoadBalancerIpv4AddressPools
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryption_FromProto(mapCtx, in.GetControlPlaneEncryption())
	out.TargetVersion = direct.LazyPtr(in.GetTargetVersion())
	out.ReleaseChannel = direct.Enum_FromProto(mapCtx, in.GetReleaseChannel())
	out.SurvivabilityConfig = Cluster_SurvivabilityConfig_FromProto(mapCtx, in.GetSurvivabilityConfig())
	out.ExternalLoadBalancerIPV6AddressPools = in.ExternalLoadBalancerIpv6AddressPools
	return out
}
func EdgeContainerClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Labels = in.Labels
	out.Fleet = Fleet_ToProto(mapCtx, in.Fleet)
	out.Networking = ClusterNetworking_ToProto(mapCtx, in.Networking)
	out.Authorization = Authorization_ToProto(mapCtx, in.Authorization)
	out.DefaultMaxPodsPerNode = direct.ValueOf(in.DefaultMaxPodsPerNode)
	out.MaintenancePolicy = MaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	out.ControlPlane = Cluster_ControlPlane_ToProto(mapCtx, in.ControlPlane)
	out.SystemAddonsConfig = Cluster_SystemAddonsConfig_ToProto(mapCtx, in.SystemAddonsConfig)
	out.ExternalLoadBalancerIpv4AddressPools = in.ExternalLoadBalancerIPV4AddressPools
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryption_ToProto(mapCtx, in.ControlPlaneEncryption)
	out.TargetVersion = direct.ValueOf(in.TargetVersion)
	out.ReleaseChannel = direct.Enum_ToProto[pb.Cluster_ReleaseChannel](mapCtx, in.ReleaseChannel)
	out.SurvivabilityConfig = Cluster_SurvivabilityConfig_ToProto(mapCtx, in.SurvivabilityConfig)
	out.ExternalLoadBalancerIpv6AddressPools = in.ExternalLoadBalancerIPV6AddressPools
	return out
}
func EdgeContainerClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.EdgeContainerClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Fleet = FleetObservedState_FromProto(mapCtx, in.GetFleet())
	out.Endpoint = direct.LazyPtr(in.GetEndpoint())
	out.Port = direct.LazyPtr(in.GetPort())
	out.ClusterCACertificate = direct.LazyPtr(in.GetClusterCaCertificate())
	out.ControlPlaneVersion = direct.LazyPtr(in.GetControlPlaneVersion())
	out.NodeVersion = direct.LazyPtr(in.GetNodeVersion())
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryptionObservedState_FromProto(mapCtx, in.GetControlPlaneEncryption())
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	out.MaintenanceEvents = direct.Slice_FromProto(mapCtx, in.MaintenanceEvents, Cluster_MaintenanceEvent_FromProto)
	out.ConnectionState = Cluster_ConnectionState_FromProto(mapCtx, in.GetConnectionState())
	return out
}
func EdgeContainerClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Fleet = FleetObservedState_ToProto(mapCtx, in.Fleet)
	out.Endpoint = direct.ValueOf(in.Endpoint)
	out.Port = direct.ValueOf(in.Port)
	out.ClusterCaCertificate = direct.ValueOf(in.ClusterCACertificate)
	out.ControlPlaneVersion = direct.ValueOf(in.ControlPlaneVersion)
	out.NodeVersion = direct.ValueOf(in.NodeVersion)
	out.ControlPlaneEncryption = Cluster_ControlPlaneEncryptionObservedState_ToProto(mapCtx, in.ControlPlaneEncryption)
	out.Status = direct.Enum_ToProto[pb.Cluster_Status](mapCtx, in.Status)
	out.MaintenanceEvents = direct.Slice_ToProto(mapCtx, in.MaintenanceEvents, Cluster_MaintenanceEvent_ToProto)
	out.ConnectionState = Cluster_ConnectionState_ToProto(mapCtx, in.ConnectionState)
	return out
}
func Cluster_ControlPlaneEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlaneEncryption) *krm.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_ControlPlaneEncryption{}

	out.KMSKeyRef = &v1beta1.KMSCryptoKeyRef{
		External: in.KmsKey,
	}
	return out
}
func Cluster_ControlPlaneEncryption_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_ControlPlaneEncryption) *pb.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlaneEncryption{}
	if in.KMSKeyRef != nil {
		out.KmsKey = in.KMSKeyRef.External
	}
	return out
}
func Status_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Details = direct.Slice_FromProto(mapCtx, in.Details, Any_FromProto)
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Details = direct.Slice_ToProto(mapCtx, in.Details, Any_ToProto)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func Any_ToProto(mapCtx *direct.MapContext, in *krm.Any) *anypb.Any {
	if in == nil {
		return nil
	}
	out := &anypb.Any{}
	out.TypeUrl = direct.ValueOf(in.TypeURL)
	out.Value = in.Value
	return out
}
func Any_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *krm.Any {
	if in == nil {
		return nil
	}
	out := &krm.Any{}
	out.TypeURL = direct.LazyPtr(in.GetTypeUrl())
	out.Value = in.GetValue()
	return out
}
func Fleet_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) *krm.Fleet {
	if in == nil {
		return nil
	}
	out := &krm.Fleet{}
	out.ProjectRef = &v1beta1.ProjectRef{
		External: in.Project,
	}
	// MISSING: Membership
	return out
}
func Fleet_ToProto(mapCtx *direct.MapContext, in *krm.Fleet) *pb.Fleet {
	if in == nil {
		return nil
	}
	out := &pb.Fleet{}
	if in.ProjectRef != nil {
		out.Project = direct.ValueOf(&in.ProjectRef.External)
	}
	// MISSING: Membership
	return out
}
