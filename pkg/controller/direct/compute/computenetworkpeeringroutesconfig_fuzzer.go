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
// proto.message: google.cloud.compute.v1.NetworkPeering
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeNetworkPeeringRoutesConfigFuzzer())
}

func computeNetworkPeeringRoutesConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.NetworkPeering{},
		ComputeNetworkPeeringRoutesConfigSpec_v1alpha1_FromProto, ComputeNetworkPeeringRoutesConfigSpec_v1alpha1_ToProto,
	)

	f.FilterSpec = func(in *pb.NetworkPeering) {
		if in.ExportCustomRoutes == nil {
			in.ExportCustomRoutes = direct.PtrTo(false)
		}
		if in.ImportCustomRoutes == nil {
			in.ImportCustomRoutes = direct.PtrTo(false)
		}
	}

	f.SpecField(".export_custom_routes")
	f.SpecField(".import_custom_routes")
	f.SpecField(".network")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".auto_create_routes")
	f.Unimplemented_NotYetTriaged(".connection_status")
	f.Unimplemented_NotYetTriaged(".exchange_subnet_routes")
	f.Unimplemented_NotYetTriaged(".export_subnet_routes_with_public_ip")
	f.Unimplemented_NotYetTriaged(".import_subnet_routes_with_public_ip")
	f.Unimplemented_NotYetTriaged(".peer_mtu")
	f.Unimplemented_NotYetTriaged(".stack_type")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".state_details")
	f.Unimplemented_NotYetTriaged(".update_strategy")

	return f
}
