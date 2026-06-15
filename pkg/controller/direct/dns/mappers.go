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

package dns

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/dns/v1"
)

func DNSManagedZoneSpec_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZone) *krm.DNSManagedZoneSpec {
	if in == nil {
		return nil
	}
	out := &krm.DNSManagedZoneSpec{}
	out.CloudLoggingConfig = ManagedZoneCloudLoggingConfig_FromAPI(mapCtx, in.CloudLoggingConfig)
	out.Description = direct.LazyPtr(in.Description)
	out.DnsName = in.DnsName
	out.DnssecConfig = ManagedZoneDnsSecConfig_FromAPI(mapCtx, in.DnssecConfig)
	out.ForwardingConfig = ManagedZoneForwardingConfig_FromAPI(mapCtx, in.ForwardingConfig)
	out.PeeringConfig = ManagedZonePeeringConfig_FromAPI(mapCtx, in.PeeringConfig)
	out.PrivateVisibilityConfig = ManagedZonePrivateVisibilityConfig_FromAPI(mapCtx, in.PrivateVisibilityConfig)
	out.ResourceID = direct.LazyPtr(in.Name)
	if in.ReverseLookupConfig != nil {
		out.ReverseLookup = direct.LazyPtr(true)
	}
	out.ServiceDirectoryConfig = ManagedZoneServiceDirectoryConfig_FromAPI(mapCtx, in.ServiceDirectoryConfig)
	out.Visibility = direct.LazyPtr(in.Visibility)
	return out
}

func DNSManagedZoneSpec_ToAPI(mapCtx *direct.MapContext, in *krm.DNSManagedZoneSpec) *api.ManagedZone {
	if in == nil {
		return nil
	}
	out := &api.ManagedZone{}
	out.CloudLoggingConfig = ManagedZoneCloudLoggingConfig_ToAPI(mapCtx, in.CloudLoggingConfig)
	out.Description = direct.ValueOf(in.Description)
	out.DnsName = in.DnsName
	out.DnssecConfig = ManagedZoneDnsSecConfig_ToAPI(mapCtx, in.DnssecConfig)
	out.ForwardingConfig = ManagedZoneForwardingConfig_ToAPI(mapCtx, in.ForwardingConfig)
	out.PeeringConfig = ManagedZonePeeringConfig_ToAPI(mapCtx, in.PeeringConfig)
	out.PrivateVisibilityConfig = ManagedZonePrivateVisibilityConfig_ToAPI(mapCtx, in.PrivateVisibilityConfig)
	out.Name = direct.ValueOf(in.ResourceID)
	if direct.ValueOf(in.ReverseLookup) {
		out.ReverseLookupConfig = &api.ManagedZoneReverseLookupConfig{}
	}
	out.ServiceDirectoryConfig = ManagedZoneServiceDirectoryConfig_ToAPI(mapCtx, in.ServiceDirectoryConfig)
	out.Visibility = direct.ValueOf(in.Visibility)
	return out
}

func DNSManagedZoneStatus_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZone) *krm.DNSManagedZoneStatus {
	if in == nil {
		return nil
	}
	out := &krm.DNSManagedZoneStatus{}
	out.CreationTime = direct.LazyPtr(in.CreationTime)
	if in.Id != 0 {
		out.ManagedZoneId = direct.LazyPtr(int64(in.Id))
	}
	out.NameServers = in.NameServers
	return out
}

func DNSManagedZoneStatus_ToAPI(mapCtx *direct.MapContext, in *krm.DNSManagedZoneStatus) *api.ManagedZone {
	if in == nil {
		return nil
	}
	out := &api.ManagedZone{}
	out.CreationTime = direct.ValueOf(in.CreationTime)
	if in.ManagedZoneId != nil {
		out.Id = uint64(*in.ManagedZoneId)
	}
	out.NameServers = in.NameServers
	return out
}
