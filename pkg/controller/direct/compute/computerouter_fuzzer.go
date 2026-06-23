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
// proto.message: google.cloud.compute.v1.Router
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeRouterFuzzer())
}

func computeRouterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Router{},
		ComputeRouterSpec_v1beta1_FromProto, ComputeRouterSpec_v1beta1_ToProto,
		ComputeRouterStatus_v1beta1_FromProto, ComputeRouterStatus_v1beta1_ToProto,
	)

	// Spec fields mapping:
	// - bgp                                -> .bgp
	//   - advertiseMode                    -> .bgp.advertise_mode
	//   - advertisedGroups                 -> .bgp.advertised_groups
	//   - advertisedIpRanges               -> .bgp.advertised_ip_ranges
	//     - description                    -> .bgp.advertised_ip_ranges[].description
	//     - range                          -> .bgp.advertised_ip_ranges[].range
	//   - asn                              -> .bgp.asn
	//   - keepaliveInterval                -> .bgp.keepalive_interval
	// - description                        -> .description
	// - encryptedInterconnectRouter        -> .encrypted_interconnect_router
	// - networkRef                         -> .network
	// - region                             -> .region
	f.SpecField(".bgp.advertise_mode")
	f.SpecField(".bgp.advertised_groups")
	f.SpecField(".bgp.advertised_ip_ranges")
	f.SpecField(".bgp.advertised_ip_ranges[].description")
	f.SpecField(".bgp.advertised_ip_ranges[].range")
	f.SpecField(".bgp.asn")
	f.SpecField(".bgp.keepalive_interval")
	f.SpecField(".description")
	f.SpecField(".encrypted_interconnect_router")
	f.SpecField(".network")
	f.SpecField(".region")

	// Status fields mapping:
	// - creationTimestamp                 -> .creation_timestamp
	// - selfLink                          -> .self_link
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".bgp.identifier_range")
	f.Unimplemented_NotYetTriaged(".bgp_peers")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".interfaces")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".md5_authentication_keys")
	f.Unimplemented_NotYetTriaged(".nats")
	f.Unimplemented_NotYetTriaged(".params")

	f.FilterSpec = func(in *pb.Router) {
		in.CreationTimestamp = nil
		in.SelfLink = nil
		in.Id = nil
		in.Kind = nil
	}

	f.FilterStatus = func(in *pb.Router) {
		in.Bgp = nil
		in.Description = nil
		in.EncryptedInterconnectRouter = nil
		in.Network = nil
		in.Region = nil
		in.Name = nil
		in.Params = nil
		in.BgpPeers = nil
		in.Interfaces = nil
		in.Md5AuthenticationKeys = nil
		in.Nats = nil
		in.Id = nil
		in.Kind = nil
	}

	return f
}
