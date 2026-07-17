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
// proto.message: google.cloud.compute.v1.SecurityPolicy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeOrganizationSecurityPolicyFuzzer())
}

func computeOrganizationSecurityPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SecurityPolicy{},
		ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto, ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto,
		ComputeOrganizationSecurityPolicyStatus_v1alpha1_FromProto, ComputeOrganizationSecurityPolicyStatus_v1alpha1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".short_name")
	f.SpecField(".parent")
	f.SpecField(".type")

	// Status fields
	f.StatusField(".fingerprint")
	f.StatusField(".id")

	// Unimplemented / Identity fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".region")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".ddos_protection_config")
	f.Unimplemented_NotYetTriaged(".ddos_protection_config.ddos_protection")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".self_link")
	f.Unimplemented_NotYetTriaged(".user_defined_fields")
	f.Unimplemented_NotYetTriaged(".associations")
	f.Unimplemented_NotYetTriaged(".adaptive_protection_config")
	f.Unimplemented_NotYetTriaged(".advanced_options_config")
	f.Unimplemented_NotYetTriaged(".recaptcha_options_config")
	f.Unimplemented_NotYetTriaged(".rules")

	return f
}
