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
// proto.message: google.cloud.compute.v1.RouterInterface
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeRouterInterfaceFuzzer())
}

func computeRouterInterfaceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.RouterInterface{},
		ComputeRouterInterfaceSpec_v1beta1_FromProto, ComputeRouterInterfaceSpec_v1beta1_ToProto,
		statusFromProto, statusToProto,
	)

	// Spec fields
	f.SpecField(".ip_range")
	f.SpecField(".linked_interconnect_attachment")
	f.SpecField(".linked_vpn_tunnel")
	f.SpecField(".private_ip_address")
	f.SpecField(".redundant_interface")
	f.SpecField(".subnetwork")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".ip_version")
	f.Unimplemented_NotYetTriaged(".management_type")

	return f
}

func statusFromProto(mapCtx *direct.MapContext, in *pb.RouterInterface) *krm.ComputeRouterInterfaceStatus {
	return &krm.ComputeRouterInterfaceStatus{}
}

func statusToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterInterfaceStatus) *pb.RouterInterface {
	if in == nil {
		return nil
	}
	return &pb.RouterInterface{}
}
