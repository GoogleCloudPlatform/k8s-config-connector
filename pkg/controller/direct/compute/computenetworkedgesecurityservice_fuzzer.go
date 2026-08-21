// Copyright 2024 Google LLC
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
// proto.message: google.cloud.compute.v1.NetworkEdgeSecurityService
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeNetworkEdgeSecurityServiceFuzzer())
}

func computeNetworkEdgeSecurityServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkEdgeSecurityService{},
		ComputeNetworkEdgeSecurityServiceSpec_v1alpha1_FromProto, ComputeNetworkEdgeSecurityServiceSpec_v1alpha1_ToProto,
		ComputeNetworkEdgeSecurityServiceObservedState_v1alpha1_FromProto, ComputeNetworkEdgeSecurityServiceObservedState_v1alpha1_ToProto,
	)

	// Spec fields mapping:
	// - description       -> .description
	// - fingerprint       -> .fingerprint
	// - securityPolicyRef -> .security_policy
	f.SpecField(".description")
	f.SpecField(".fingerprint")
	f.SpecField(".security_policy")

	// Status fields mapping:
	// - creationTimestamp -> .creation_timestamp
	// - id                -> .id
	// - kind              -> .kind
	// - region            -> .region
	// - selfLink          -> .self_link
	// - selfLinkWithID    -> .self_link_with_id
	f.StatusField(".creation_timestamp")
	f.StatusField(".id")
	f.StatusField(".kind")
	f.StatusField(".region")
	f.StatusField(".self_link")
	f.StatusField(".self_link_with_id")

	// Identity/Special fields
	f.Unimplemented_Identity(".name")

	return f
}
