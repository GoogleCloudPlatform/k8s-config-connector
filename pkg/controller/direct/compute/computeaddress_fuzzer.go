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
// proto.message: google.cloud.compute.v1.Address
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeAddressFuzzer())
}

func computeAddressFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Address{},
		ComputeAddressSpec_v1beta1_FromProto, ComputeAddressSpec_v1beta1_ToProto,
		ComputeAddressStatus_v1beta1_FromProto, ComputeAddressStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".address")
	f.SpecField(".address_type")
	f.SpecField(".description")
	f.SpecField(".ip_version")
	f.SpecField(".ipv6_endpoint_type")
	f.SpecField(".network")
	f.SpecField(".network_tier")
	f.SpecField(".prefix_length")
	f.SpecField(".purpose")
	f.SpecField(".subnetwork")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".label_fingerprint")
	f.StatusField(".self_link")
	f.StatusField(".users")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".ip_collection")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".status")

	return f
}
