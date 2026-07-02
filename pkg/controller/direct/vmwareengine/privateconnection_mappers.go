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

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VMwareEnginePrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.VMwareEnginePrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEnginePrivateConnectionObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.VmwareEngineNetworkCanonical = direct.LazyPtr(in.GetVmwareEngineNetworkCanonical())
	out.PeeringID = direct.LazyPtr(in.GetPeeringId())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.PeeringState = direct.Enum_FromProto(mapCtx, in.GetPeeringState())
	return out
}

func VMwareEnginePrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEnginePrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.VmwareEngineNetworkCanonical = direct.ValueOf(in.VmwareEngineNetworkCanonical)
	out.PeeringId = direct.ValueOf(in.PeeringID)
	out.Uid = direct.ValueOf(in.Uid)
	out.PeeringState = direct.Enum_ToProto[pb.PrivateConnection_PeeringState](mapCtx, in.PeeringState)
	return out
}

func VMwareEnginePrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.VMwareEnginePrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEnginePrivateConnectionSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetVmwareEngineNetwork() != "" {
		out.VMwareEngineNetworkRef = &krm.VmwareEngineNetworkRef{
			External: in.GetVmwareEngineNetwork(),
		}
	}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.RoutingMode = direct.Enum_FromProto(mapCtx, in.GetRoutingMode())
	if in.GetServiceNetwork() != "" {
		out.ServiceNetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetServiceNetwork()}
	}
	return out
}

func VMwareEnginePrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEnginePrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	out.Description = direct.ValueOf(in.Description)
	if in.VMwareEngineNetworkRef != nil {
		out.VmwareEngineNetwork = in.VMwareEngineNetworkRef.External
	}
	out.Type = direct.Enum_ToProto[pb.PrivateConnection_Type](mapCtx, in.Type)
	out.RoutingMode = direct.Enum_ToProto[pb.PrivateConnection_RoutingMode](mapCtx, in.RoutingMode)
	if in.ServiceNetworkRef != nil {
		out.ServiceNetwork = in.ServiceNetworkRef.External
	}
	return out
}
