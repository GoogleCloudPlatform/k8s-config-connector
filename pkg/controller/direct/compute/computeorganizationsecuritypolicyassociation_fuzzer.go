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
	fuzztesting.RegisterKRMFuzzer(computeOrganizationSecurityPolicyAssociationFuzzer())
}

func computeOrganizationSecurityPolicyAssociationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicyAssociation{},
		ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto, ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto,
		ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_FromProto, ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_ToProto,
	)

	// Spec fields
	f.SpecField(".attachment_target")
	f.SpecField(".firewall_policy_id")

	// Status fields
	f.StatusField(".display_name")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".short_name")

	return f
}
