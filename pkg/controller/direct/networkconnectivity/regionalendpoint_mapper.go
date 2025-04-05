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

package networkconnectivity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkConnectivityRegionalEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RegionalEndpoint) *krm.NetworkConnectivityRegionalEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityRegionalEndpointObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.PSCForwardingRule = direct.LazyPtr(in.GetPscForwardingRule())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NetworkConnectivityRegionalEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityRegionalEndpointObservedState) *pb.RegionalEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.RegionalEndpoint{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.PscForwardingRule = direct.ValueOf(in.PSCForwardingRule)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func NetworkConnectivityRegionalEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.RegionalEndpoint) *krm.NetworkConnectivityRegionalEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityRegionalEndpointSpec{}
	out.AccessType = direct.LazyPtr(in.GetAccessType())
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: Name
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef =  &refsv1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.TargetGoogleAPI = direct.LazyPtr(in.GetTargetGoogleApi())
	return out
}
func NetworkConnectivityRegionalEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityRegionalEndpointSpec) *pb.RegionalEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.RegionalEndpoint{}
	out.AccessType = direct.ValueOf(in.AccessType)
	out.Address = direct.ValueOf(in.Address)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.TargetGoogleApi = direct.ValueOf(in.TargetGoogleAPI)
	return out
}