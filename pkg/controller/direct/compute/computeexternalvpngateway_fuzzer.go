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
// proto.message: google.cloud.compute.v1.ExternalVpnGateway
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeExternalVPNGatewayFuzzer())
}

func computeExternalVPNGatewayFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ExternalVpnGateway{},
		ComputeExternalVPNGatewaySpec_v1beta1_FromProto, ComputeExternalVPNGatewaySpec_v1beta1_ToProto,
		ComputeExternalVPNGatewayStatus_v1beta1_FromProto, ComputeExternalVPNGatewayStatus_v1beta1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".interfaces")
	f.SpecField(".interfaces[].id")
	f.SpecField(".interfaces[].ip_address")
	f.SpecField(".redundancy_type")

	f.StatusField(".label_fingerprint")
	f.StatusField(".self_link")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_Internal(".creation_timestamp")
	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".interfaces[].ipv6_address")
	f.Unimplemented_NotYetTriaged(".params")

	return f
}
