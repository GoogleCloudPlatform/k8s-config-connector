// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain
// crd.kind: ServiceNetworkingPeeredDNSDomain

package servicenetworking

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/servicenetworking/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(peeredDnsDomainFuzzer())
}

func peeredDnsDomainFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.PeeredDnsDomain{},
		ServiceNetworkingPeeredDNSDomainSpec_FromProto, ServiceNetworkingPeeredDNSDomainSpec_ToProto,
		ServiceNetworkingPeeredDNSDomainObservedState_FromProto, ServiceNetworkingPeeredDNSDomainObservedState_ToProto,
	)

	// Spec fields
	f.SpecFields.Insert(".DnsSuffix")

	// System / naming fields
	f.UnimplementedFields.Insert(".Name")
	f.UnimplementedFields.Insert(".ForceSendFields")
	f.UnimplementedFields.Insert(".NullFields")

	return f
}
