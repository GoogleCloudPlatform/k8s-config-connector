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
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkPolicy_NetworkService_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krm.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPolicy_NetworkService{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	// MISSING: State
	return out
}
func NetworkPolicy_NetworkService_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPolicy_NetworkService) *pb.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy_NetworkService{}
	out.Enabled = direct.ValueOf(in.Enabled)
	// MISSING: State
	return out
}
func NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krm.NetworkPolicy_NetworkServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPolicy_NetworkServiceObservedState{}
	// MISSING: Enabled
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPolicy_NetworkServiceObservedState) *pb.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy_NetworkService{}
	// MISSING: Enabled
	out.State = direct.Enum_ToProto[pb.NetworkPolicy_NetworkService_State](mapCtx, in.State)
	return out
}
func VMwareEngineNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VMwareEngineNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VpcNetworks
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Uid
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func VMwareEngineNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkSpec) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: VpcNetworks
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_Type](mapCtx, in.Type)
	// MISSING: Uid
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func VmwareEngineNetwork_VpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krm.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork_VpcNetwork) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VpcNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krm.VmwareEngineNetwork_VpcNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork_VpcNetworkObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func VmwareEngineNetwork_VpcNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork_VpcNetworkObservedState) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_VpcNetwork_Type](mapCtx, in.Type)
	out.Network = direct.ValueOf(in.Network)
	return out
}
