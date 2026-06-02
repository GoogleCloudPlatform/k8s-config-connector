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

package computesecuritypolicy

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeSecurityPolicyFuzzer())
}

func computeSecurityPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.SecurityPolicy, krm.ComputeSecurityPolicySpec, krm.ComputeSecurityPolicyStatus](&pb.SecurityPolicy{},
		ComputeSecurityPolicySpec_v1beta1_FromProto, ComputeSecurityPolicySpec_v1beta1_ToProto,
		nil, nil,
	)

	// Spec fields
	f.SpecField(".adaptive_protection_config")
	f.SpecField(".advanced_options_config")
	f.SpecField(".description")
	f.SpecField(".recaptcha_options_config")
	f.SpecField(".rules")
	f.SpecField(".type")

	// Status fields
	f.StatusField(".fingerprint")
	f.StatusField(".self_link")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".associations")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".ddos_protection_config")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".parent")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".short_name")
	f.Unimplemented_NotYetTriaged(".user_defined_fields")

	return f
}
