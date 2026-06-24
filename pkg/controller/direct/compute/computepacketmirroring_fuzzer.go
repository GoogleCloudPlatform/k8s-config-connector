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
// proto.message: google.cloud.compute.v1.PacketMirroring
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computePacketMirroringFuzzer())
}

func computePacketMirroringFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PacketMirroring{},
		ComputePacketMirroringSpec_v1beta1_FromProto, ComputePacketMirroringSpec_v1beta1_ToProto,
		ComputePacketMirroringStatus_v1beta1_FromProto, ComputePacketMirroringStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".collector_ilb")
	f.SpecField(".collector_ilb.url")
	f.SpecField(".description")
	f.SpecField(".enable")
	f.SpecField(".filter")
	f.SpecField(".filter.cidr_ranges")
	f.SpecField(".filter.direction")
	f.SpecField(".filter.I_p_protocols")
	f.SpecField(".mirrored_resources")
	f.SpecField(".mirrored_resources.instances")
	f.SpecField(".mirrored_resources.instances[].canonical_url")
	f.SpecField(".mirrored_resources.instances[].url")
	f.SpecField(".mirrored_resources.subnetworks")
	f.SpecField(".mirrored_resources.subnetworks[].canonical_url")
	f.SpecField(".mirrored_resources.subnetworks[].url")
	f.SpecField(".mirrored_resources.tags")
	f.SpecField(".network")
	f.SpecField(".network.url")
	f.SpecField(".priority")

	// Status fields
	f.StatusField(".collector_ilb.canonical_url")
	f.StatusField(".id")
	f.StatusField(".network.canonical_url")
	f.StatusField(".region")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".kind")

	f.FilterSpec = func(in *pb.PacketMirroring) {
		if in.CollectorIlb != nil {
			if in.CollectorIlb.Url == nil || *in.CollectorIlb.Url == "" {
				in.CollectorIlb = nil
			}
		}
		if in.Network != nil {
			if in.Network.Url == nil || *in.Network.Url == "" {
				in.Network = nil
			}
		}
		if in.MirroredResources != nil {
			var instances []*pb.PacketMirroringMirroredResourceInfoInstanceInfo
			for _, inst := range in.MirroredResources.Instances {
				if inst != nil && (inst.CanonicalUrl != nil || (inst.Url != nil && *inst.Url != "")) {
					instances = append(instances, inst)
				}
			}
			in.MirroredResources.Instances = instances

			var subnetworks []*pb.PacketMirroringMirroredResourceInfoSubnetInfo
			for _, subnet := range in.MirroredResources.Subnetworks {
				if subnet != nil && (subnet.CanonicalUrl != nil || (subnet.Url != nil && *subnet.Url != "")) {
					subnetworks = append(subnetworks, subnet)
				}
			}
			in.MirroredResources.Subnetworks = subnetworks

			if len(in.MirroredResources.Instances) == 0 && len(in.MirroredResources.Subnetworks) == 0 && len(in.MirroredResources.Tags) == 0 {
				in.MirroredResources = nil
			}
		}
		if in.Filter != nil {
			if len(in.Filter.CidrRanges) == 0 && in.Filter.Direction == nil && len(in.Filter.IPProtocols) == 0 {
				in.Filter = nil
			}
		}
	}

	f.FilterStatus = func(in *pb.PacketMirroring) {
		if in.CollectorIlb != nil {
			if in.CollectorIlb.CanonicalUrl == nil {
				in.CollectorIlb = nil
			}
		}
		if in.Network != nil {
			if in.Network.CanonicalUrl == nil {
				in.Network = nil
			}
		}
	}

	return f
}
