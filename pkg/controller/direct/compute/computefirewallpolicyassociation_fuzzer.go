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
// proto.message: google.cloud.compute.v1.FirewallPolicyAssociation
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeFirewallPolicyAssociationFuzzer())
}

func computeFirewallPolicyAssociationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicyAssociation{},
		ComputeFirewallPolicyAssociationSpec_v1beta1_FromProto, ComputeFirewallPolicyAssociationSpec_v1beta1_ToProto,
		ComputeFirewallPolicyAssociationStatus_v1beta1_FromProto, ComputeFirewallPolicyAssociationStatus_v1beta1_ToProto,
	)

	f.SpecField(".attachment_target")
	f.SpecField(".firewall_policy_id")
	f.SpecField(".name")

	f.StatusField(".short_name")

	f.Unimplemented_Identity(".display_name")

	return f
}
