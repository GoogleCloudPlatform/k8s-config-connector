// Copyright 2024 Google LLC
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
// proto.message: google.cloud.edgecontainer.v1.VpnConnection
// api.group: edgecontainer.cnrm.cloud.google.com

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(edgeContainerVpnConnectionFuzzer())
}

func edgeContainerVpnConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.VpnConnection{},
		EdgeContainerVpnConnectionSpec_FromProto, EdgeContainerVpnConnectionSpec_ToProto,
		EdgeContainerVpnConnectionObservedState_FromProto, EdgeContainerVpnConnectionObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".nat_gateway_ip")
	f.SpecFields.Insert(".bgp_routing_mode")
	f.SpecFields.Insert(".cluster")
	f.SpecFields.Insert(".vpc")
	f.SpecFields.Insert(".vpc_project")
	f.SpecFields.Insert(".enable_high_availability")
	f.SpecFields.Insert(".router")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".details")

	f.UnimplementedFields.Insert(".name")
	return f
}
