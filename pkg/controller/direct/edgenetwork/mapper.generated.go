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

package edgenetwork

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgenetwork/v1alpha1"
)
func EdgenetworkNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.EdgenetworkNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkNetworkObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Mtu
	return out
}
func EdgenetworkNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkNetworkObservedState) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Mtu
	return out
}
func EdgenetworkNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.EdgenetworkNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkNetworkSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Mtu
	return out
}
func EdgenetworkNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkNetworkSpec) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Mtu
	return out
}
func Network_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.Network {
	if in == nil {
		return nil
	}
	out := &krm.Network{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Mtu = direct.LazyPtr(in.GetMtu())
	return out
}
func Network_ToProto(mapCtx *direct.MapContext, in *krm.Network) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Mtu = direct.ValueOf(in.Mtu)
	return out
}
func NetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.NetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Mtu
	return out
}
func NetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkObservedState) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Mtu
	return out
}
