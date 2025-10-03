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

func VMwareEngineNetworkPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy) *krm.VMwareEngineNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkPolicySpec{}
	// MISSING: Name
	out.InternetAccess = NetworkPolicy_NetworkService_FromProto(mapCtx, in.GetInternetAccess())
	out.ExternalIP = NetworkPolicy_NetworkService_FromProto(mapCtx, in.GetExternalIp())
	out.EdgeServicesCIDR = direct.LazyPtr(in.GetEdgeServicesCidr())
	if in.GetVmwareEngineNetwork() != "" {
		out.VMwareEngineNetworkRef = &krm.VmwareEngineNetworkRef{
			External: in.GetVmwareEngineNetwork(),
		}
	}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func VMwareEngineNetworkPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkPolicySpec) *pb.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy{}
	// MISSING: Name
	out.InternetAccess = NetworkPolicy_NetworkService_ToProto(mapCtx, in.InternetAccess)
	out.ExternalIp = NetworkPolicy_NetworkService_ToProto(mapCtx, in.ExternalIP)
	out.EdgeServicesCidr = direct.ValueOf(in.EdgeServicesCIDR)
	if in.VMwareEngineNetworkRef != nil {
		out.VmwareEngineNetwork = in.VMwareEngineNetworkRef.External
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func VMwareEngineNetworkPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy) *krm.VMwareEngineNetworkPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkPolicyObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.InternetAccess = NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx, in.GetInternetAccess())
	out.ExternalIP = NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx, in.GetExternalIp())
	out.UID = direct.LazyPtr(in.GetUid())
	out.VMwareEngineNetworkCanonical = direct.LazyPtr(in.GetVmwareEngineNetworkCanonical())
	return out
}
func VMwareEngineNetworkPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkPolicyObservedState) *pb.NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.InternetAccess = NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx, in.InternetAccess)
	out.ExternalIp = NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx, in.ExternalIP)
	out.Uid = direct.ValueOf(in.UID)
	out.VmwareEngineNetworkCanonical = direct.ValueOf(in.VMwareEngineNetworkCanonical)
	return out
}
