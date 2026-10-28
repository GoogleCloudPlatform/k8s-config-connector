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
// proto.message: google.cloud.beyondcorp.appgateways.v1.AppGateway
// api.group: beyondcorp.cnrm.cloud.google.com

package beyondcorpappgateway

import (
	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(beyondCorpAppGatewayFuzzer())
}

func beyondCorpAppGatewayFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AppGateway{},
		BeyondCorpAppGatewaySpec_FromProto, BeyondCorpAppGatewaySpec_ToProto,
		BeyondCorpAppGatewayStatus_FromProto, BeyondCorpAppGatewayStatus_ToProto,
	)

	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | KRM Spec Field                                            | Proto Field                       | Fuzzer/Mapping Status                                       |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | displayName                                               | display_name                      | SpecField (".display_name")                                 |
	// | hostType                                                  | host_type                         | SpecField (".host_type")                                    |
	// | projectRef                                                | N/A                               | KRM spec-only GCP project reference                         |
	// | region                                                    | N/A                               | KRM spec-only GCP location URL segment                      |
	// | resourceID                                                | N/A                               | KRM spec-only resource acquisition ID                       |
	// | type                                                      | type                              | SpecField (".type")                                         |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | KRM Status Field                                          | Proto Field                       | Fuzzer/Mapping Status                                       |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | allocatedConnections                                      | allocated_connections             | StatusField (".allocated_connections")                      |
	// | allocatedConnections[].ingressPort                        | allocated_connections[].ingress_port | StatusField (via .allocated_connections)                 |
	// | allocatedConnections[].pscUri                             | allocated_connections[].psc_uri   | StatusField (via .allocated_connections)                    |
	// | conditions                                                | N/A                               | KRM status condition list                                   |
	// | externalRef                                               | name                              | Identity (URL)                                              |
	// | observedGeneration                                        | N/A                               | KRM status observed generation                              |
	// | state                                                     | state                             | StatusField (".state")                                      |
	// | uri                                                       | uri                               | StatusField (".uri")                                        |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".type")
	f.SpecField(".host_type")

	// Status fields
	f.StatusField(".state")
	f.StatusField(".uri")
	f.StatusField(".allocated_connections")

	// Identity fields
	f.Unimplemented_Identity(".name")

	// Labels and Annotations
	f.Unimplemented_LabelsAnnotations(".labels")

	// Output-only / unmapped fields on proto
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".uid")

	return f
}
