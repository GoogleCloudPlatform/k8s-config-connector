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
// proto.message: google.cloud.compute.v1.InstanceGroup
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInstanceGroupFuzzer())
}

func computeInstanceGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InstanceGroup{},
		ComputeInstanceGroupSpec_v1beta1_FromProto, ComputeInstanceGroupSpec_v1beta1_ToProto,
		ComputeInstanceGroupStatus_v1beta1_FromProto, ComputeInstanceGroupStatus_v1beta1_ToProto,
	)

	// Field comparison: ComputeInstanceGroupSpec vs pb.InstanceGroup Proto
	// - Spec.Description                 maps to proto field .description
	// - Spec.Instances                   not represented in pb.InstanceGroup (managed separately)
	// - Spec.NamedPorts                  maps to proto field .named_ports
	// - Spec.NetworkRef                  maps to proto field .network
	// - Spec.ResourceID                  maps to proto field .name (via Identity)
	// - Spec.Zone                        maps to proto field .zone

	// Field comparison: ComputeInstanceGroupStatus vs pb.InstanceGroup Proto
	// - Status.Conditions                not represented in pb.InstanceGroup
	// - Status.ObservedGeneration        not represented in pb.InstanceGroup
	// - Status.SelfLink                  maps to proto field .self_link
	// - Status.Size                      maps to proto field .size

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".named_ports")
	f.SpecField(".network")
	f.SpecField(".zone")

	// Status fields
	f.StatusField(".self_link")
	f.StatusField(".size")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".fingerprint")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".subnetwork")

	return f
}
