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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EdgecontainerVpnConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection) *krm.EdgecontainerVpnConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgecontainerVpnConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NatGatewayIP
	// MISSING: BgpRoutingMode
	// MISSING: Cluster
	// MISSING: Vpc
	// MISSING: VpcProject
	// MISSING: EnableHighAvailability
	// MISSING: Router
	// MISSING: Details
	return out
}
func EdgecontainerVpnConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgecontainerVpnConnectionObservedState) *pb.VpnConnection {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NatGatewayIP
	// MISSING: BgpRoutingMode
	// MISSING: Cluster
	// MISSING: Vpc
	// MISSING: VpcProject
	// MISSING: EnableHighAvailability
	// MISSING: Router
	// MISSING: Details
	return out
}
func EdgecontainerVpnConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection) *krm.EdgecontainerVpnConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgecontainerVpnConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NatGatewayIP
	// MISSING: BgpRoutingMode
	// MISSING: Cluster
	// MISSING: Vpc
	// MISSING: VpcProject
	// MISSING: EnableHighAvailability
	// MISSING: Router
	// MISSING: Details
	return out
}
func EdgecontainerVpnConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgecontainerVpnConnectionSpec) *pb.VpnConnection {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NatGatewayIP
	// MISSING: BgpRoutingMode
	// MISSING: Cluster
	// MISSING: Vpc
	// MISSING: VpcProject
	// MISSING: EnableHighAvailability
	// MISSING: Router
	// MISSING: Details
	return out
}
func VpnConnection_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection) *krm.VpnConnection {
	if in == nil {
		return nil
	}
	out := &krm.VpnConnection{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.NatGatewayIP = direct.LazyPtr(in.GetNatGatewayIp())
	out.BgpRoutingMode = direct.Enum_FromProto(mapCtx, in.GetBgpRoutingMode())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.Vpc = direct.LazyPtr(in.GetVpc())
	out.VpcProject = VpnConnection_VpcProject_FromProto(mapCtx, in.GetVpcProject())
	out.EnableHighAvailability = direct.LazyPtr(in.GetEnableHighAvailability())
	out.Router = direct.LazyPtr(in.GetRouter())
	// MISSING: Details
	return out
}
func VpnConnection_ToProto(mapCtx *direct.MapContext, in *krm.VpnConnection) *pb.VpnConnection {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.NatGatewayIp = direct.ValueOf(in.NatGatewayIP)
	out.BgpRoutingMode = direct.Enum_ToProto[pb.VpnConnection_BgpRoutingMode](mapCtx, in.BgpRoutingMode)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.Vpc = direct.ValueOf(in.Vpc)
	out.VpcProject = VpnConnection_VpcProject_ToProto(mapCtx, in.VpcProject)
	out.EnableHighAvailability = direct.ValueOf(in.EnableHighAvailability)
	out.Router = direct.ValueOf(in.Router)
	// MISSING: Details
	return out
}
func VpnConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection) *krm.VpnConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VpnConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: NatGatewayIP
	// MISSING: BgpRoutingMode
	// MISSING: Cluster
	// MISSING: Vpc
	// MISSING: VpcProject
	// MISSING: EnableHighAvailability
	// MISSING: Router
	out.Details = VpnConnection_Details_FromProto(mapCtx, in.GetDetails())
	return out
}
func VpnConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VpnConnectionObservedState) *pb.VpnConnection {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: NatGatewayIP
	// MISSING: BgpRoutingMode
	// MISSING: Cluster
	// MISSING: Vpc
	// MISSING: VpcProject
	// MISSING: EnableHighAvailability
	// MISSING: Router
	out.Details = VpnConnection_Details_ToProto(mapCtx, in.Details)
	return out
}
func VpnConnection_Details_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection_Details) *krm.VpnConnection_Details {
	if in == nil {
		return nil
	}
	out := &krm.VpnConnection_Details{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = direct.LazyPtr(in.GetError())
	out.CloudRouter = VpnConnection_Details_CloudRouter_FromProto(mapCtx, in.GetCloudRouter())
	out.CloudVpns = direct.Slice_FromProto(mapCtx, in.CloudVpns, VpnConnection_Details_CloudVpn_FromProto)
	return out
}
func VpnConnection_Details_ToProto(mapCtx *direct.MapContext, in *krm.VpnConnection_Details) *pb.VpnConnection_Details {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection_Details{}
	out.State = direct.Enum_ToProto[pb.VpnConnection_Details_State](mapCtx, in.State)
	out.Error = direct.ValueOf(in.Error)
	out.CloudRouter = VpnConnection_Details_CloudRouter_ToProto(mapCtx, in.CloudRouter)
	out.CloudVpns = direct.Slice_ToProto(mapCtx, in.CloudVpns, VpnConnection_Details_CloudVpn_ToProto)
	return out
}
func VpnConnection_Details_CloudRouter_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection_Details_CloudRouter) *krm.VpnConnection_Details_CloudRouter {
	if in == nil {
		return nil
	}
	out := &krm.VpnConnection_Details_CloudRouter{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func VpnConnection_Details_CloudRouter_ToProto(mapCtx *direct.MapContext, in *krm.VpnConnection_Details_CloudRouter) *pb.VpnConnection_Details_CloudRouter {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection_Details_CloudRouter{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func VpnConnection_Details_CloudVpn_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection_Details_CloudVpn) *krm.VpnConnection_Details_CloudVpn {
	if in == nil {
		return nil
	}
	out := &krm.VpnConnection_Details_CloudVpn{}
	out.Gateway = direct.LazyPtr(in.GetGateway())
	return out
}
func VpnConnection_Details_CloudVpn_ToProto(mapCtx *direct.MapContext, in *krm.VpnConnection_Details_CloudVpn) *pb.VpnConnection_Details_CloudVpn {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection_Details_CloudVpn{}
	out.Gateway = direct.ValueOf(in.Gateway)
	return out
}
func VpnConnection_VpcProject_FromProto(mapCtx *direct.MapContext, in *pb.VpnConnection_VpcProject) *krm.VpnConnection_VpcProject {
	if in == nil {
		return nil
	}
	out := &krm.VpnConnection_VpcProject{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func VpnConnection_VpcProject_ToProto(mapCtx *direct.MapContext, in *krm.VpnConnection_VpcProject) *pb.VpnConnection_VpcProject {
	if in == nil {
		return nil
	}
	out := &pb.VpnConnection_VpcProject{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	return out
}
