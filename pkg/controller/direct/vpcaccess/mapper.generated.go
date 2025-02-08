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

package vpcaccess

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vpcaccess/apiv1/vpcaccesspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vpcaccess/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Connector_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.Connector {
	if in == nil {
		return nil
	}
	out := &krm.Connector{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.IPCidrRange = direct.LazyPtr(in.GetIpCidrRange())
	// MISSING: State
	out.MinThroughput = direct.LazyPtr(in.GetMinThroughput())
	out.MaxThroughput = direct.LazyPtr(in.GetMaxThroughput())
	// MISSING: ConnectedProjects
	out.Subnet = Connector_Subnet_FromProto(mapCtx, in.GetSubnet())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.MinInstances = direct.LazyPtr(in.GetMinInstances())
	out.MaxInstances = direct.LazyPtr(in.GetMaxInstances())
	return out
}
func Connector_ToProto(mapCtx *direct.MapContext, in *krm.Connector) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
	out.Name = direct.ValueOf(in.Name)
	out.Network = direct.ValueOf(in.Network)
	out.IpCidrRange = direct.ValueOf(in.IPCidrRange)
	// MISSING: State
	out.MinThroughput = direct.ValueOf(in.MinThroughput)
	out.MaxThroughput = direct.ValueOf(in.MaxThroughput)
	// MISSING: ConnectedProjects
	out.Subnet = Connector_Subnet_ToProto(mapCtx, in.Subnet)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.MinInstances = direct.ValueOf(in.MinInstances)
	out.MaxInstances = direct.ValueOf(in.MaxInstances)
	return out
}
func ConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.ConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorObservedState{}
	// MISSING: Name
	// MISSING: Network
	// MISSING: IPCidrRange
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: MinThroughput
	// MISSING: MaxThroughput
	out.ConnectedProjects = in.ConnectedProjects
	// MISSING: Subnet
	// MISSING: MachineType
	// MISSING: MinInstances
	// MISSING: MaxInstances
	return out
}
func ConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorObservedState) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
	// MISSING: Name
	// MISSING: Network
	// MISSING: IPCidrRange
	out.State = direct.Enum_ToProto[pb.Connector_State](mapCtx, in.State)
	// MISSING: MinThroughput
	// MISSING: MaxThroughput
	out.ConnectedProjects = in.ConnectedProjects
	// MISSING: Subnet
	// MISSING: MachineType
	// MISSING: MinInstances
	// MISSING: MaxInstances
	return out
}
func Connector_Subnet_FromProto(mapCtx *direct.MapContext, in *pb.Connector_Subnet) *krm.Connector_Subnet {
	if in == nil {
		return nil
	}
	out := &krm.Connector_Subnet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	return out
}
func Connector_Subnet_ToProto(mapCtx *direct.MapContext, in *krm.Connector_Subnet) *pb.Connector_Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Connector_Subnet{}
	out.Name = direct.ValueOf(in.Name)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	return out
}
func VpcaccessConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.VpcaccessConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VpcaccessConnectorObservedState{}
	// MISSING: Name
	// MISSING: Network
	// MISSING: IPCidrRange
	// MISSING: State
	// MISSING: MinThroughput
	// MISSING: MaxThroughput
	// MISSING: ConnectedProjects
	// MISSING: Subnet
	// MISSING: MachineType
	// MISSING: MinInstances
	// MISSING: MaxInstances
	return out
}
func VpcaccessConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VpcaccessConnectorObservedState) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
	// MISSING: Name
	// MISSING: Network
	// MISSING: IPCidrRange
	// MISSING: State
	// MISSING: MinThroughput
	// MISSING: MaxThroughput
	// MISSING: ConnectedProjects
	// MISSING: Subnet
	// MISSING: MachineType
	// MISSING: MinInstances
	// MISSING: MaxInstances
	return out
}
func VpcaccessConnectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.Connector) *krm.VpcaccessConnectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.VpcaccessConnectorSpec{}
	// MISSING: Name
	// MISSING: Network
	// MISSING: IPCidrRange
	// MISSING: State
	// MISSING: MinThroughput
	// MISSING: MaxThroughput
	// MISSING: ConnectedProjects
	// MISSING: Subnet
	// MISSING: MachineType
	// MISSING: MinInstances
	// MISSING: MaxInstances
	return out
}
func VpcaccessConnectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.VpcaccessConnectorSpec) *pb.Connector {
	if in == nil {
		return nil
	}
	out := &pb.Connector{}
	// MISSING: Name
	// MISSING: Network
	// MISSING: IPCidrRange
	// MISSING: State
	// MISSING: MinThroughput
	// MISSING: MaxThroughput
	// MISSING: ConnectedProjects
	// MISSING: Subnet
	// MISSING: MachineType
	// MISSING: MinInstances
	// MISSING: MaxInstances
	return out
}
