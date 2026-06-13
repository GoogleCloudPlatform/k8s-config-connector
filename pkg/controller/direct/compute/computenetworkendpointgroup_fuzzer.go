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
// proto.message: google.cloud.compute.v1.NetworkEndpointGroup
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeNetworkEndpointGroupFuzzer())
}

func computeNetworkEndpointGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkEndpointGroup{},
		ComputeNetworkEndpointGroupSpec_v1beta1_FromProto, ComputeNetworkEndpointGroupSpec_v1beta1_ToProto,
		ComputeNetworkEndpointGroupStatus_v1beta1_FromProto, ComputeNetworkEndpointGroupStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".default_port")
	f.SpecField(".description")
	f.SpecField(".network")
	f.SpecField(".network_endpoint_type")
	f.SpecField(".subnetwork")

	// Status fields
	f.StatusField(".self_link")
	f.StatusField(".size")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".annotations")
	f.Unimplemented_NotYetTriaged(".app_engine")
	f.Unimplemented_NotYetTriaged(".cloud_function")
	f.Unimplemented_NotYetTriaged(".cloud_run")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".psc_data")
	f.Unimplemented_NotYetTriaged(".psc_target_service")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".zone")

	return f
}
