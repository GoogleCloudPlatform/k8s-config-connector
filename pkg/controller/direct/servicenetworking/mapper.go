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
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicenetworking/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/servicenetworking/v1"
)

func ServiceNetworkingPeeredDnsDomainObservedState_FromProto(mapCtx *direct.MapContext, in *api.PeeredDnsDomain) *krmv1alpha1.ServiceNetworkingPeeredDnsDomainObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ServiceNetworkingPeeredDnsDomainObservedState{}
	// MISSING: DNSSuffix
	// MISSING: Name
	return out
}
func ServiceNetworkingPeeredDnsDomainObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ServiceNetworkingPeeredDnsDomainObservedState) *api.PeeredDnsDomain {
	if in == nil {
		return nil
	}
	out := &api.PeeredDnsDomain{}
	// MISSING: DNSSuffix
	// MISSING: Name
	return out
}
func ServiceNetworkingPeeredDnsDomainSpec_FromProto(mapCtx *direct.MapContext, in *api.PeeredDnsDomain) *krmv1alpha1.ServiceNetworkingPeeredDnsDomainSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ServiceNetworkingPeeredDnsDomainSpec{}
	out.DNSSuffix = direct.LazyPtr(in.DnsSuffix)
	// MISSING: Name
	return out
}
func ServiceNetworkingPeeredDnsDomainSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ServiceNetworkingPeeredDnsDomainSpec) *api.PeeredDnsDomain {
	if in == nil {
		return nil
	}
	out := &api.PeeredDnsDomain{}
	out.DnsSuffix = direct.ValueOf(in.DNSSuffix)
	// MISSING: Name
	return out
}
