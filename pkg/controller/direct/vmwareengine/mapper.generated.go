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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VMwareEngineNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VMwareEngineNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: VpcNetworks
	// MISSING: State
	// MISSING: Type
	// MISSING: Uid
	// MISSING: Etag
	return out
}
func VMwareEngineNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkObservedState) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: VpcNetworks
	// MISSING: State
	// MISSING: Type
	// MISSING: Uid
	// MISSING: Etag
	return out
}
func VMwareEngineNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VMwareEngineNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: VpcNetworks
	// MISSING: State
	// MISSING: Type
	// MISSING: Uid
	// MISSING: Etag
	return out
}
func VMwareEngineNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkSpec) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: VpcNetworks
	// MISSING: State
	// MISSING: Type
	// MISSING: Uid
	// MISSING: Etag
	return out
}
func VmwareEngineNetwork_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VpcNetworks
	// MISSING: State
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Uid
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func VmwareEngineNetwork_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	// MISSING: VpcNetworks
	// MISSING: State
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_Type](mapCtx, in.Type)
	// MISSING: Uid
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func VmwareEngineNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VmwareEngineNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetworkObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	out.VpcNetworks = direct.Slice_FromProto(mapCtx, in.VpcNetworks, VmwareEngineNetwork_VpcNetwork_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Type
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Etag
	return out
}
func VmwareEngineNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetworkObservedState) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	out.VpcNetworks = direct.Slice_ToProto(mapCtx, in.VpcNetworks, VmwareEngineNetwork_VpcNetwork_ToProto)
	out.State = direct.Enum_ToProto[pb.VmwareEngineNetwork_State](mapCtx, in.State)
	// MISSING: Type
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Etag
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
