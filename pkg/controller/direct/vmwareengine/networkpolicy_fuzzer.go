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
// proto.message: google.cloud.vmwareengine.v1.NetworkPolicy
// api.group: vmwareengine.cnrm.cloud.google.com

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vmwareEngineNetworkPolicyFuzzer())
}

func vmwareEngineNetworkPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkPolicy{},
		VMwareEngineNetworkPolicySpec_FromProto, VMwareEngineNetworkPolicySpec_ToProto,
		VMwareEngineNetworkPolicyObservedState_FromProto, VMwareEngineNetworkPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".internet_access")
	f.SpecFields.Insert(".external_ip")
	f.SpecFields.Insert(".edge_services_cidr")
	f.SpecFields.Insert(".vmware_engine_network")
	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".internet_access")
	f.StatusFields.Insert(".external_ip")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".vmware_engine_network_canonical")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
