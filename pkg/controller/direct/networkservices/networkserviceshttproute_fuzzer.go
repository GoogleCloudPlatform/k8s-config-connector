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
// proto.message: google.cloud.networkservices.v1.HttpRoute
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(networkServicesHTTPRouteFuzzer())
}

func networkServicesHTTPRouteFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HttpRoute{},
		NetworkServicesHTTPRouteSpec_FromProto, NetworkServicesHTTPRouteSpec_ToProto,
		NetworkServicesHTTPRouteStatus_FromProto, NetworkServicesHTTPRouteStatus_ToProto,
	)

	// Field comparison between KRM NetworkServicesHTTPRouteSpec and Proto HttpRoute:
	// - Description           <=> .description
	// - Hostnames             <=> .hostnames
	// - Gateways              <=> .gateways
	// - Meshes                <=> .meshes
	// - Rules                 <=> .rules
	//
	// Fields ignored because they are not part of KRM Spec:
	// - .name                 (Identity)
	// - .self_link            (Status)
	// - .create_time          (Status)
	// - .update_time          (Status)
	// - .labels               (handled by KCC metadata labels)

	f.SpecField(".description")
	f.SpecField(".hostnames")
	f.SpecField(".gateways")
	f.SpecField(".meshes")
	f.SpecField(".rules")

	f.StatusField(".self_link")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".rules[].action.direct_response")
	f.Unimplemented_NotYetTriaged(".rules[].action.idle_timeout")
	f.Unimplemented_NotYetTriaged(".rules[].action.stateful_session_affinity")
	f.Unimplemented_NotYetTriaged(".rules[].action.destinations[].request_header_modifier")
	f.Unimplemented_NotYetTriaged(".rules[].action.destinations[].response_header_modifier")
	f.Unimplemented_NotYetTriaged(".rules[].action.request_mirror_policy.destination.request_header_modifier")
	f.Unimplemented_NotYetTriaged(".rules[].action.request_mirror_policy.destination.response_header_modifier")
	f.Unimplemented_NotYetTriaged(".rules[].action.request_mirror_policy.mirror_percent")
	f.Unimplemented_NotYetTriaged(".rules[].matches[].query_parameters[].present_match")

	return f
}
