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

// +tool:fuzz-gen
// proto.message: google.cloud.networkservices.v1.Gateway
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(networkServicesGatewayFuzzer())
}

func networkServicesGatewayFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Gateway{},
		NetworkServicesGatewaySpec_FromProto, NetworkServicesGatewaySpec_ToProto,
		NetworkServicesGatewayStatus_FromProto, NetworkServicesGatewayStatus_ToProto,
	)

	// Field comparison between KRM NetworkServicesGatewaySpec and Proto Gateway:
	// - Addresses             <=> .addresses
	// - Description           <=> .description
	// - Ports                 <=> .ports
	// - Scope                 <=> .scope
	// - ServerTlsPolicyRef    <=> .server_tls_policy
	// - Type                  <=> .type
	//
	// Fields ignored because they are not part of KRM Spec:
	// - .name                 (Identity)
	// - .self_link            (Status)
	// - .create_time          (Status)
	// - .update_time          (Status)
	// - .labels               (handled by KCC metadata labels)
	// - .certificate_urls     (unimplemented/not yet triaged)
	// - .gateway_security_policy (unimplemented/not yet triaged)
	// - .network              (unimplemented/not yet triaged)
	// - .subnetwork           (unimplemented/not yet triaged)
	// - .ip_version           (unimplemented/not yet triaged)
	// - .envoy_headers        (unimplemented/not yet triaged)
	// - .routing_mode         (unimplemented/not yet triaged)

	f.SpecField(".addresses")
	f.SpecField(".description")
	f.SpecField(".ports")
	f.SpecField(".scope")
	f.SpecField(".server_tls_policy")
	f.SpecField(".type")

	f.StatusField(".self_link")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".certificate_urls")
	f.Unimplemented_NotYetTriaged(".gateway_security_policy")
	f.Unimplemented_NotYetTriaged(".network")
	f.Unimplemented_NotYetTriaged(".subnetwork")
	f.Unimplemented_NotYetTriaged(".ip_version")
	f.Unimplemented_NotYetTriaged(".envoy_headers")
	f.Unimplemented_NotYetTriaged(".routing_mode")
	f.Unimplemented_NotYetTriaged(".all_ports")
	f.Unimplemented_NotYetTriaged(".allow_global_access")

	return f
}
