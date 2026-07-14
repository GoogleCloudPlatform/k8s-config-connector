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
// proto.message: google.cloud.networksecurity.v1.FirewallEndpoint
// api.group: networksecurity.cnrm.cloud.google.com

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(firewallEndpointFuzzer())
}

func firewallEndpointFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallEndpoint{},
		NetworkSecurityFirewallEndpointSpec_v1alpha1_FromProto, NetworkSecurityFirewallEndpointSpec_v1alpha1_ToProto,
		NetworkSecurityFirewallEndpointObservedState_v1alpha1_FromProto, NetworkSecurityFirewallEndpointObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".labels")
	f.SpecField(".billing_project_id")
	f.SpecField(".endpoint_settings")

	f.StatusField(".state")
	f.StatusField(".reconciling")
	f.StatusField(".associations")

	f.IdentityField(".name")

	return f
}
