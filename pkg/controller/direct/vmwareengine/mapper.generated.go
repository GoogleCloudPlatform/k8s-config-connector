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

package vmwareengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func PrivateConnection_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.VmwareEngineNetwork = direct.LazyPtr(in.GetVmwareEngineNetwork())
	// MISSING: VmwareEngineNetworkCanonical
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: PeeringID
	out.RoutingMode = direct.Enum_FromProto(mapCtx, in.GetRoutingMode())
	// MISSING: Uid
	out.ServiceNetwork = direct.LazyPtr(in.GetServiceNetwork())
	// MISSING: PeeringState
	return out
}
func PrivateConnection_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnection) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.VmwareEngineNetwork = direct.ValueOf(in.VmwareEngineNetwork)
	// MISSING: VmwareEngineNetworkCanonical
	out.Type = direct.Enum_ToProto[pb.PrivateConnection_Type](mapCtx, in.Type)
	// MISSING: PeeringID
	out.RoutingMode = direct.Enum_ToProto[pb.PrivateConnection_RoutingMode](mapCtx, in.RoutingMode)
	// MISSING: Uid
	out.ServiceNetwork = direct.ValueOf(in.ServiceNetwork)
	// MISSING: PeeringState
	return out
}
func PrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.PrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: VmwareEngineNetwork
	out.VmwareEngineNetworkCanonical = direct.LazyPtr(in.GetVmwareEngineNetworkCanonical())
	// MISSING: Type
	out.PeeringID = direct.LazyPtr(in.GetPeeringId())
	// MISSING: RoutingMode
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: ServiceNetwork
	out.PeeringState = direct.Enum_FromProto(mapCtx, in.GetPeeringState())
	return out
}
func PrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	// MISSING: VmwareEngineNetwork
	out.VmwareEngineNetworkCanonical = direct.ValueOf(in.VmwareEngineNetworkCanonical)
	// MISSING: Type
	out.PeeringId = direct.ValueOf(in.PeeringID)
	// MISSING: RoutingMode
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: ServiceNetwork
	out.PeeringState = direct.Enum_ToProto[pb.PrivateConnection_PeeringState](mapCtx, in.PeeringState)
	return out
}
func VmwareenginePrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.VmwareenginePrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareenginePrivateConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: State
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: Type
	// MISSING: PeeringID
	// MISSING: RoutingMode
	// MISSING: Uid
	// MISSING: ServiceNetwork
	// MISSING: PeeringState
	return out
}
func VmwareenginePrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareenginePrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: State
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: Type
	// MISSING: PeeringID
	// MISSING: RoutingMode
	// MISSING: Uid
	// MISSING: ServiceNetwork
	// MISSING: PeeringState
	return out
}
func VmwareenginePrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.VmwareenginePrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareenginePrivateConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: State
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: Type
	// MISSING: PeeringID
	// MISSING: RoutingMode
	// MISSING: Uid
	// MISSING: ServiceNetwork
	// MISSING: PeeringState
	return out
}
func VmwareenginePrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareenginePrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: State
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: Type
	// MISSING: PeeringID
	// MISSING: RoutingMode
	// MISSING: Uid
	// MISSING: ServiceNetwork
	// MISSING: PeeringState
	return out
}
