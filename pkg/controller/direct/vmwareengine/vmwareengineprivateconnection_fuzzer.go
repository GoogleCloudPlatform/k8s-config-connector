// Copyright 2026 Google LLC
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
// proto.message: google.cloud.vmwareengine.v1.PrivateConnection
// api.group: vmwareengine.cnrm.cloud.google.com

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vmwareEnginePrivateConnectionFuzzer())
}

func vmwareEnginePrivateConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PrivateConnection{},
		VMwareEnginePrivateConnectionSpec_FromProto, VMwareEnginePrivateConnectionSpec_ToProto,
		VMwareEnginePrivateConnectionObservedState_FromProto, VMwareEnginePrivateConnectionObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".vmware_engine_network")
	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".routing_mode")
	f.SpecFields.Insert(".service_network")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".vmware_engine_network_canonical")
	f.StatusFields.Insert(".peering_id")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".peering_state")

	return f
}
