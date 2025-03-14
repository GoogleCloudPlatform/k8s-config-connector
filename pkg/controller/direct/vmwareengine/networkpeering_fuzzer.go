// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.vmwareengine.v1.NetworkPeering
// api.group: vmwareengine.cnrm.cloud.google.com

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vmwareEngineNetworkPeeringFuzzer())
}

func vmwareEngineNetworkPeeringFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkPeering{},
		VMwareEngineNetworkPeeringSpec_FromProto, VMwareEngineNetworkPeeringSpec_ToProto,
		VMwareEngineNetworkPeeringObservedState_FromProto, VMwareEngineNetworkPeeringObservedState_ToProto,
	)

	f.SpecFields.Insert(".peer_network")
	f.SpecFields.Insert(".export_custom_routes")
	f.SpecFields.Insert(".import_custom_routes")
	f.SpecFields.Insert(".exchange_subnet_routes")
	f.SpecFields.Insert(".export_custom_routes_with_public_ip")
	f.SpecFields.Insert(".import_custom_routes_with_public_ip")
	f.SpecFields.Insert(".peer_mtu")
	f.SpecFields.Insert(".peer_network_type")
	f.SpecFields.Insert(".vmware_engine_network")
	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_details")
	f.StatusFields.Insert(".uid")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
