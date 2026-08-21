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
	api "google.golang.org/api/dns/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
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
