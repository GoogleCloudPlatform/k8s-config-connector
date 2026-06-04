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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/dns/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(dnsManagedZoneFuzzer())
}

func dnsManagedZoneFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.ManagedZone{},
		DNSManagedZoneSpec_FromAPI, DNSManagedZoneSpec_ToAPI,
		DNSManagedZoneStatus_FromAPI, DNSManagedZoneStatus_ToAPI,
	)

	f.SpecField(".CloudLoggingConfig")
	f.SpecField(".Description")
	f.SpecField(".DnsName")
	f.SpecField(".DnssecConfig")
	f.SpecField(".ForwardingConfig")
	f.SpecField(".PeeringConfig")
	f.SpecField(".PrivateVisibilityConfig")
	f.SpecField(".ReverseLookupConfig")
	f.SpecField(".ServiceDirectoryConfig")
	f.SpecField(".Visibility")

	f.StatusField(".CreationTime")
	f.StatusField(".Id")
	f.StatusField(".NameServers")

	f.IdentityField(".Name")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")
	f.Ignore_JSONBookkeeping(".CloudLoggingConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".CloudLoggingConfig.NullFields")
	f.Ignore_JSONBookkeeping(".DnssecConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".DnssecConfig.NullFields")
	f.Ignore_JSONBookkeeping(".ForwardingConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ForwardingConfig.NullFields")
	f.Ignore_JSONBookkeeping(".PeeringConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".PeeringConfig.NullFields")
	f.Ignore_JSONBookkeeping(".PeeringConfig.TargetNetwork.ForceSendFields")
	f.Ignore_JSONBookkeeping(".PeeringConfig.TargetNetwork.NullFields")
	f.Ignore_JSONBookkeeping(".PrivateVisibilityConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".PrivateVisibilityConfig.NullFields")
	f.Ignore_JSONBookkeeping(".ReverseLookupConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ReverseLookupConfig.NullFields")
	f.Ignore_JSONBookkeeping(".ServiceDirectoryConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ServiceDirectoryConfig.NullFields")
	f.Ignore_JSONBookkeeping(".ServiceDirectoryConfig.Namespace.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ServiceDirectoryConfig.Namespace.NullFields")

	f.Ignore_JSONBookkeeping(".DnssecConfig.DefaultKeySpecs[]")
	f.Ignore_JSONBookkeeping(".ForwardingConfig.TargetNameServers[]")
	f.Ignore_JSONBookkeeping(".PrivateVisibilityConfig.GkeClusters[]")
	f.Ignore_JSONBookkeeping(".PrivateVisibilityConfig.Networks[]")

	f.Unimplemented_NotYetTriaged(".NameServerSet")
	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".Kind")
	f.Unimplemented_NotYetTriaged(".CloudLoggingConfig.Kind")
	f.Unimplemented_NotYetTriaged(".DnssecConfig.Kind")
	f.Unimplemented_NotYetTriaged(".ForwardingConfig.Kind")
	f.Unimplemented_NotYetTriaged(".ForwardingConfig.TargetNameServers[].Kind")
	f.Unimplemented_NotYetTriaged(".DomainName")
	f.Unimplemented_NotYetTriaged(".Ipv6Address")
	f.Unimplemented_NotYetTriaged(".PeeringConfig.Kind")
	f.Unimplemented_NotYetTriaged(".PeeringConfig.TargetNetwork.Kind")
	f.Unimplemented_NotYetTriaged(".PeeringConfig.TargetNetwork.DeactivateTime")
	f.Unimplemented_NotYetTriaged(".PrivateVisibilityConfig.Kind")
	f.Unimplemented_NotYetTriaged(".PrivateVisibilityConfig.GkeClusters[].Kind")
	f.Unimplemented_NotYetTriaged(".PrivateVisibilityConfig.Networks[].Kind")
	f.Unimplemented_NotYetTriaged(".ReverseLookupConfig.Kind")
	f.Unimplemented_NotYetTriaged(".ServiceDirectoryConfig.Kind")
	f.Unimplemented_NotYetTriaged(".ServiceDirectoryConfig.Namespace.Kind")
	f.Unimplemented_NotYetTriaged(".ServiceDirectoryConfig.Namespace.DeletionTime")

	return f
}

func DNSManagedZoneSpec_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZone) *DNSManagedZoneSpec {
	if in == nil {
		return nil
	}
	out := &DNSManagedZoneSpec{}
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

func DNSManagedZoneSpec_ToAPI(mapCtx *direct.MapContext, in *DNSManagedZoneSpec) *api.ManagedZone {
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

func DNSManagedZoneStatus_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZone) *DNSManagedZoneStatus {
	if in == nil {
		return nil
	}
	out := &DNSManagedZoneStatus{}
	out.CreationTime = direct.LazyPtr(in.CreationTime)
	if in.Id != 0 {
		out.ManagedZoneId = direct.LazyPtr(int64(in.Id))
	}
	out.NameServers = in.NameServers
	return out
}

func DNSManagedZoneStatus_ToAPI(mapCtx *direct.MapContext, in *DNSManagedZoneStatus) *api.ManagedZone {
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

func ManagedZoneCloudLoggingConfig_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZoneCloudLoggingConfig) *ManagedZoneCloudLoggingConfig {
	if in == nil {
		return nil
	}
	out := &ManagedZoneCloudLoggingConfig{}
	out.EnableLogging = direct.LazyPtr(in.EnableLogging)
	return out
}

func ManagedZoneCloudLoggingConfig_ToAPI(mapCtx *direct.MapContext, in *ManagedZoneCloudLoggingConfig) *api.ManagedZoneCloudLoggingConfig {
	if in == nil {
		return nil
	}
	out := &api.ManagedZoneCloudLoggingConfig{}
	out.EnableLogging = direct.ValueOf(in.EnableLogging)
	return out
}

func DnsKeySpec_FromAPI(mapCtx *direct.MapContext, in *api.DnsKeySpec) *DnsKeySpec {
	if in == nil {
		return nil
	}
	out := &DnsKeySpec{}
	out.Algorithm = direct.LazyPtr(in.Algorithm)
	out.KeyLength = direct.LazyPtr(in.KeyLength)
	out.KeyType = direct.LazyPtr(in.KeyType)
	out.Kind = direct.LazyPtr(in.Kind)
	return out
}

func DnsKeySpec_ToAPI(mapCtx *direct.MapContext, in *DnsKeySpec) *api.DnsKeySpec {
	if in == nil {
		return nil
	}
	out := &api.DnsKeySpec{}
	out.Algorithm = direct.ValueOf(in.Algorithm)
	out.KeyLength = direct.ValueOf(in.KeyLength)
	out.KeyType = direct.ValueOf(in.KeyType)
	out.Kind = direct.ValueOf(in.Kind)
	return out
}

func ManagedZoneDnsSecConfig_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZoneDnsSecConfig) *ManagedZoneDnsSecConfig {
	if in == nil {
		return nil
	}
	out := &ManagedZoneDnsSecConfig{}
	out.DefaultKeySpecs = direct.Slice_FromProto(mapCtx, in.DefaultKeySpecs, DnsKeySpec_FromAPI)
	out.Kind = direct.LazyPtr(in.Kind)
	out.NonExistence = direct.LazyPtr(in.NonExistence)
	out.State = direct.LazyPtr(in.State)
	return out
}

func ManagedZoneDnsSecConfig_ToAPI(mapCtx *direct.MapContext, in *ManagedZoneDnsSecConfig) *api.ManagedZoneDnsSecConfig {
	if in == nil {
		return nil
	}
	out := &api.ManagedZoneDnsSecConfig{}
	out.DefaultKeySpecs = direct.Slice_ToProto(mapCtx, in.DefaultKeySpecs, DnsKeySpec_ToAPI)
	out.Kind = direct.ValueOf(in.Kind)
	out.NonExistence = direct.ValueOf(in.NonExistence)
	out.State = direct.ValueOf(in.State)
	return out
}

func ManagedZoneForwardingConfigNameServerTarget_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZoneForwardingConfigNameServerTarget) *ManagedZoneForwardingConfigNameServerTarget {
	if in == nil {
		return nil
	}
	out := &ManagedZoneForwardingConfigNameServerTarget{}
	out.ForwardingPath = direct.LazyPtr(in.ForwardingPath)
	out.Ipv4Address = in.Ipv4Address
	return out
}

func ManagedZoneForwardingConfigNameServerTarget_ToAPI(mapCtx *direct.MapContext, in *ManagedZoneForwardingConfigNameServerTarget) *api.ManagedZoneForwardingConfigNameServerTarget {
	if in == nil {
		return nil
	}
	out := &api.ManagedZoneForwardingConfigNameServerTarget{}
	out.ForwardingPath = direct.ValueOf(in.ForwardingPath)
	out.Ipv4Address = in.Ipv4Address
	return out
}

func ManagedZoneForwardingConfig_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZoneForwardingConfig) *ManagedZoneForwardingConfig {
	if in == nil {
		return nil
	}
	out := &ManagedZoneForwardingConfig{}
	out.TargetNameServers = direct.Slice_FromProto(mapCtx, in.TargetNameServers, ManagedZoneForwardingConfigNameServerTarget_FromAPI)
	return out
}

func ManagedZoneForwardingConfig_ToAPI(mapCtx *direct.MapContext, in *ManagedZoneForwardingConfig) *api.ManagedZoneForwardingConfig {
	if in == nil {
		return nil
	}
	out := &api.ManagedZoneForwardingConfig{}
	out.TargetNameServers = direct.Slice_ToProto(mapCtx, in.TargetNameServers, ManagedZoneForwardingConfigNameServerTarget_ToAPI)
	return out
}

func ManagedZonePeeringConfig_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZonePeeringConfig) *ManagedZonePeeringConfig {
	if in == nil {
		return nil
	}
	out := &ManagedZonePeeringConfig{}
	if in.TargetNetwork != nil {
		out.TargetNetwork = &ManagedZonePeeringConfigTargetNetwork{}
		out.TargetNetwork.NetworkRef.External = in.TargetNetwork.NetworkUrl
	}
	return out
}

func ManagedZonePeeringConfig_ToAPI(mapCtx *direct.MapContext, in *ManagedZonePeeringConfig) *api.ManagedZonePeeringConfig {
	if in == nil {
		return nil
	}
	out := &api.ManagedZonePeeringConfig{}
	if in.TargetNetwork != nil && in.TargetNetwork.NetworkRef.External != "" {
		out.TargetNetwork = &api.ManagedZonePeeringConfigTargetNetwork{
			NetworkUrl: in.TargetNetwork.NetworkRef.External,
		}
	}
	return out
}

func ManagedZonePrivateVisibilityConfigGKECluster_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZonePrivateVisibilityConfigGKECluster) *ManagedZonePrivateVisibilityConfigGKECluster {
	if in == nil {
		return nil
	}
	out := &ManagedZonePrivateVisibilityConfigGKECluster{}
	out.GkeClusterNameRef.External = in.GkeClusterName
	return out
}

func ManagedZonePrivateVisibilityConfigGKECluster_ToAPI(mapCtx *direct.MapContext, in *ManagedZonePrivateVisibilityConfigGKECluster) *api.ManagedZonePrivateVisibilityConfigGKECluster {
	if in == nil {
		return nil
	}
	out := &api.ManagedZonePrivateVisibilityConfigGKECluster{}
	out.GkeClusterName = in.GkeClusterNameRef.External
	return out
}

func ManagedZonePrivateVisibilityConfigNetwork_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZonePrivateVisibilityConfigNetwork) *ManagedZonePrivateVisibilityConfigNetwork {
	if in == nil {
		return nil
	}
	out := &ManagedZonePrivateVisibilityConfigNetwork{}
	out.NetworkRef.External = in.NetworkUrl
	return out
}

func ManagedZonePrivateVisibilityConfigNetwork_ToAPI(mapCtx *direct.MapContext, in *ManagedZonePrivateVisibilityConfigNetwork) *api.ManagedZonePrivateVisibilityConfigNetwork {
	if in == nil {
		return nil
	}
	out := &api.ManagedZonePrivateVisibilityConfigNetwork{}
	out.NetworkUrl = in.NetworkRef.External
	return out
}

func ManagedZonePrivateVisibilityConfig_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZonePrivateVisibilityConfig) *ManagedZonePrivateVisibilityConfig {
	if in == nil {
		return nil
	}
	out := &ManagedZonePrivateVisibilityConfig{}
	out.GKEClusters = direct.Slice_FromProto(mapCtx, in.GkeClusters, ManagedZonePrivateVisibilityConfigGKECluster_FromAPI)
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, ManagedZonePrivateVisibilityConfigNetwork_FromAPI)
	return out
}

func ManagedZonePrivateVisibilityConfig_ToAPI(mapCtx *direct.MapContext, in *ManagedZonePrivateVisibilityConfig) *api.ManagedZonePrivateVisibilityConfig {
	if in == nil {
		return nil
	}
	out := &api.ManagedZonePrivateVisibilityConfig{}
	out.GkeClusters = direct.Slice_ToProto(mapCtx, in.GKEClusters, ManagedZonePrivateVisibilityConfigGKECluster_ToAPI)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, ManagedZonePrivateVisibilityConfigNetwork_ToAPI)
	return out
}

func ManagedZoneServiceDirectoryConfig_FromAPI(mapCtx *direct.MapContext, in *api.ManagedZoneServiceDirectoryConfig) *ManagedZoneServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &ManagedZoneServiceDirectoryConfig{}
	if in.Namespace != nil {
		out.Namespace = &ManagedZoneServiceDirectoryConfigNamespace{}
		out.Namespace.NamespaceUrl = in.Namespace.NamespaceUrl
	}
	return out
}

func ManagedZoneServiceDirectoryConfig_ToAPI(mapCtx *direct.MapContext, in *ManagedZoneServiceDirectoryConfig) *api.ManagedZoneServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &api.ManagedZoneServiceDirectoryConfig{}
	if in.Namespace != nil && in.Namespace.NamespaceUrl != "" {
		out.Namespace = &api.ManagedZoneServiceDirectoryConfigNamespace{
			NamespaceUrl: in.Namespace.NamespaceUrl,
		}
	}
	return out
}
