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
func NetworkPolicy_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy) *krm.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPolicy{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.InternetAccess = NetworkPolicy_NetworkService_FromProto(mapCtx, in.GetInternetAccess())
	out.ExternalIP = NetworkPolicy_NetworkService_FromProto(mapCtx, in.GetExternalIp())
	out.EdgeServicesCidr = direct.LazyPtr(in.GetEdgeServicesCidr())
	// MISSING: Uid
	out.VmwareEngineNetwork = direct.LazyPtr(in.GetVmwareEngineNetwork())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VmwareEngineNetworkCanonical
	return out
}
func NetworkPolicy_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPolicy) *pb.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.InternetAccess = NetworkPolicy_NetworkService_ToProto(mapCtx, in.InternetAccess)
	out.ExternalIp = NetworkPolicy_NetworkService_ToProto(mapCtx, in.ExternalIP)
	out.EdgeServicesCidr = direct.ValueOf(in.EdgeServicesCidr)
	// MISSING: Uid
	out.VmwareEngineNetwork = direct.ValueOf(in.VmwareEngineNetwork)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: VmwareEngineNetworkCanonical
	return out
}
func NetworkPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy) *krm.NetworkPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPolicyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.InternetAccess = NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx, in.GetInternetAccess())
	// MISSING: ExternalIP
	// MISSING: EdgeServicesCidr
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	out.VmwareEngineNetworkCanonical = direct.LazyPtr(in.GetVmwareEngineNetworkCanonical())
	return out
}
func NetworkPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPolicyObservedState) *pb.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.InternetAccess = NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx, in.InternetAccess)
	// MISSING: ExternalIP
	// MISSING: EdgeServicesCidr
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	out.VmwareEngineNetworkCanonical = direct.ValueOf(in.VmwareEngineNetworkCanonical)
	return out
}
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
func VmwareengineNetworkPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy) *krm.VmwareengineNetworkPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNetworkPolicyObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternetAccess
	// MISSING: ExternalIP
	// MISSING: EdgeServicesCidr
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	// MISSING: VmwareEngineNetworkCanonical
	return out
}
func VmwareengineNetworkPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNetworkPolicyObservedState) *pb.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternetAccess
	// MISSING: ExternalIP
	// MISSING: EdgeServicesCidr
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	// MISSING: VmwareEngineNetworkCanonical
	return out
}
func VmwareengineNetworkPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy) *krm.VmwareengineNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNetworkPolicySpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternetAccess
	// MISSING: ExternalIP
	// MISSING: EdgeServicesCidr
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	// MISSING: VmwareEngineNetworkCanonical
	return out
}
func VmwareengineNetworkPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNetworkPolicySpec) *pb.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternetAccess
	// MISSING: ExternalIP
	// MISSING: EdgeServicesCidr
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	// MISSING: VmwareEngineNetworkCanonical
	return out
}
