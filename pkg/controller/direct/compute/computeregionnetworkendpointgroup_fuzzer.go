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
	fuzztesting.RegisterKRMFuzzer(computeRegionNetworkEndpointGroupFuzzer())
}

func computeRegionNetworkEndpointGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkEndpointGroup{},
		ComputeRegionNetworkEndpointGroupSpec_v1beta1_FromProto, ComputeRegionNetworkEndpointGroupSpec_v1beta1_ToProto,
		ComputeRegionNetworkEndpointGroupStatus_v1beta1_FromProto, ComputeRegionNetworkEndpointGroupStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".cloud_function")
	f.SpecField(".cloud_function.function")
	f.SpecField(".cloud_function.url_mask")
	f.SpecField(".cloud_run")
	f.SpecField(".cloud_run.service")
	f.SpecField(".cloud_run.tag")
	f.SpecField(".cloud_run.url_mask")
	f.SpecField(".description")
	f.SpecField(".network")
	f.SpecField(".network_endpoint_type")
	f.SpecField(".psc_target_service")
	f.SpecField(".region")
	f.SpecField(".subnetwork")

	// Status fields
	f.StatusField(".self_link")

	// Unimplemented / Identity fields
	f.Unimplemented_Identity(".id")
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".kind")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".annotations")
	f.Unimplemented_NotYetTriaged(".app_engine")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".default_port")
	f.Unimplemented_NotYetTriaged(".psc_data")
	f.Unimplemented_NotYetTriaged(".size")
	f.Unimplemented_NotYetTriaged(".zone")

	return f
}
