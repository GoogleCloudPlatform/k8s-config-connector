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

	f.Unimplemented_Identity(".name")

	f.SpecField(".description")
	f.SpecField(".vmware_engine_network")
	f.SpecField(".type")
	f.SpecField(".routing_mode")
	f.SpecField(".service_network")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".vmware_engine_network_canonical")
	f.StatusField(".peering_id")
	f.StatusField(".uid")
	f.StatusField(".peering_state")

	return f
}
