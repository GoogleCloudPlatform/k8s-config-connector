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

package networkconnectivity

import (
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/networkconnectivity/v1"
)

func NetworkConnectivityRegionalEndpointSpec_ToAPI(mapCtx *direct.MapContext, in *krm.NetworkConnectivityRegionalEndpointSpec) *api.RegionalEndpoint {
	if in == nil {
		return nil
	}
	out := &api.RegionalEndpoint{}
	out.AccessType = direct.ValueOf(in.AccessType)
	if in.AddressRef != nil {
		out.Address = in.AddressRef.External
	}
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

func NetworkConnectivityRegionalEndpointSpec_FromAPI(mapCtx *direct.MapContext, in *api.RegionalEndpoint) *krm.NetworkConnectivityRegionalEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityRegionalEndpointSpec{}
	out.AccessType = direct.LazyPtr(in.AccessType)
	if in.Address != "" {
		out.AddressRef = &computev1beta1.ComputeAddressRef{External: in.Address}
	}
	out.Description = direct.LazyPtr(in.Description)
	if in.Network != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.Network}
	}
	if in.Subnetwork != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.Subnetwork}
	}
	out.TargetGoogleAPI = direct.LazyPtr(in.TargetGoogleApi)
	return out
}

func NetworkConnectivityRegionalEndpointObservedState_FromAPI(mapCtx *direct.MapContext, in *api.RegionalEndpoint) *krm.NetworkConnectivityRegionalEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityRegionalEndpointObservedState{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.IPAddress = direct.LazyPtr(in.IpAddress)
	out.PSCForwardingRule = direct.LazyPtr(in.PscForwardingRule)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	return out
}

func NetworkConnectivityRegionalEndpointObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.NetworkConnectivityRegionalEndpointObservedState) *api.RegionalEndpoint {
	if in == nil {
		return nil
	}
	out := &api.RegionalEndpoint{}
	out.CreateTime = direct.ValueOf(in.CreateTime)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.PscForwardingRule = direct.ValueOf(in.PSCForwardingRule)
	out.UpdateTime = direct.ValueOf(in.UpdateTime)
	return out
}
