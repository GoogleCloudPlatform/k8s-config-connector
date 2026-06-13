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
// proto.message: google.cloud.compute.v1.NetworkEndpoint
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeGlobalNetworkEndpointFuzzer())
}

func computeGlobalNetworkEndpointFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.NetworkEndpoint, krm.ComputeGlobalNetworkEndpointSpec](&pb.NetworkEndpoint{},
		ComputeGlobalNetworkEndpointSpec_v1alpha1_FromProto, ComputeGlobalNetworkEndpointSpec_v1alpha1_ToProto,
	)

	f.SpecField(".fqdn")
	f.SpecField(".ip_address")
	f.SpecField(".port")

	f.Unimplemented_NotYetTriaged(".annotations")
	f.Unimplemented_NotYetTriaged(".client_destination_port")
	f.Unimplemented_NotYetTriaged(".instance")
	f.Unimplemented_NotYetTriaged(".ipv6_address")

	return f
}
