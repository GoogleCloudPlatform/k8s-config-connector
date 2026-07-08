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
// proto.message: google.cloud.compute.v1.SslPolicy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeSSLPolicyFuzzer())
}

func computeSSLPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SslPolicy{},
		ComputeSSLPolicySpec_v1beta1_FromProto, ComputeSSLPolicySpec_v1beta1_ToProto,
		ComputeSSLPolicyStatus_v1beta1_FromProto, ComputeSSLPolicyStatus_v1beta1_ToProto,
	)

	// Field mapping comparison:
	// KRM Spec Fields:
	// - customFeatures -> .custom_features
	// - description    -> .description
	// - minTlsVersion  -> .min_tls_version
	// - profile        -> .profile
	// - resourceID     -> .name (Identity)
	//
	// KRM Status Fields:
	// - creationTimestamp -> .creation_timestamp
	// - enabledFeatures   -> .enabled_features
	// - fingerprint       -> .fingerprint
	// - selfLink          -> .self_link
	// - observedGeneration (not mapped directly to proto)

	// Spec fields
	f.SpecField(".custom_features")
	f.SpecField(".description")
	f.SpecField(".min_tls_version")
	f.SpecField(".profile")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".enabled_features")
	f.StatusField(".fingerprint")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".warnings")
	f.Unimplemented_NotYetTriaged(".post_quantum_key_exchange")

	return f
}
