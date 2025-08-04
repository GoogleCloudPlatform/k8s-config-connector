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

package servicenetworking

// We don't have a proto for servicenetworking, so we write the mappers manually (TIP: do start with the generated mappers to make this easy)

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicenetworking/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicenetworking/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/servicenetworking/v1"
)

func ServiceNetworkingPeeredDNSDomainObservedState_FromProto(mapCtx *direct.MapContext, in *api.PeeredDnsDomain) *krmv1alpha1.ServiceNetworkingPeeredDNSDomainObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ServiceNetworkingPeeredDNSDomainObservedState{}
	// MISSING: DNSSuffix
	// MISSING: Name
	return out
}
func ServiceNetworkingPeeredDNSDomainObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ServiceNetworkingPeeredDNSDomainObservedState) *api.PeeredDnsDomain {
	if in == nil {
		return nil
	}
	out := &api.PeeredDnsDomain{}
	// MISSING: DNSSuffix
	// MISSING: Name
	return out
}
func ServiceNetworkingPeeredDNSDomainSpec_FromProto(mapCtx *direct.MapContext, in *api.PeeredDnsDomain) *krmv1alpha1.ServiceNetworkingPeeredDNSDomainSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ServiceNetworkingPeeredDNSDomainSpec{}
	out.DNSSuffix = direct.LazyPtr(in.DnsSuffix)
	// MISSING: Name
	return out
}
func ServiceNetworkingPeeredDNSDomainSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ServiceNetworkingPeeredDNSDomainSpec) *api.PeeredDnsDomain {
	if in == nil {
		return nil
	}
	out := &api.PeeredDnsDomain{}
	out.DnsSuffix = direct.ValueOf(in.DNSSuffix)
	// MISSING: Name
	return out
}

func ServiceNetworkingConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *api.Connection) *krmv1beta1.ServiceNetworkingConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ServiceNetworkingConnectionObservedState{}
	// MISSING: Peering
	return out
}
func ServiceNetworkingConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ServiceNetworkingConnectionObservedState) *api.Connection {
	if in == nil {
		return nil
	}
	out := &api.Connection{}
	// MISSING: Peering
	return out
}
func ServiceNetworkingConnectionSpec_FromProto(mapCtx *direct.MapContext, in *api.Connection) *krmv1beta1.ServiceNetworkingConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ServiceNetworkingConnectionSpec{}
	if in.Network != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.Network}
	}
	// MISSING: Peering
	out.ReservedPeeringRanges = ServiceNetworkingConnectionSpec_ReservedPeeringRanges_FromProto(mapCtx, in.ReservedPeeringRanges)
	return out
}
func ServiceNetworkingConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ServiceNetworkingConnectionSpec) *api.Connection {
	if in == nil {
		return nil
	}
	out := &api.Connection{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	// MISSING: Peering
	out.ReservedPeeringRanges = ServiceNetworkingConnectionSpec_ReservedPeeringRanges_ToProto(mapCtx, in.ReservedPeeringRanges)
	return out
}

func ServiceNetworkingConnectionSpec_ReservedPeeringRanges_FromProto(mapCtx *direct.MapContext, in []string) []*refs.ComputeAddressRef {
	if in == nil {
		return nil
	}
	out := make([]*refs.ComputeAddressRef, len(in))
	for i, item := range in {
		out[i] = &refs.ComputeAddressRef{External: item}
	}
	return out
}
func ServiceNetworkingConnectionSpec_ReservedPeeringRanges_ToProto(mapCtx *direct.MapContext, in []*refs.ComputeAddressRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, item := range in {
		if item == nil {
			continue
		}
		out[i] = item.External
	}
	return out
}
