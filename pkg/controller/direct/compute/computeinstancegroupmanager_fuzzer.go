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
// proto.message: google.cloud.compute.v1.InstanceGroupManager
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInstanceGroupManagerFuzzer())
}

func computeInstanceGroupManagerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InstanceGroupManager{},
		ComputeInstanceGroupManagerSpec_v1beta1_FromProto, ComputeInstanceGroupManagerSpec_v1beta1_ToProto,
		ComputeInstanceGroupManagerStatus_v1beta1_FromProto, ComputeInstanceGroupManagerStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".auto_healing_policies")
	f.SpecField(".base_instance_name")
	f.SpecField(".description")
	f.SpecField(".distribution_policy")
	f.SpecField(".instance_template")
	f.SpecField(".named_ports")
	f.SpecField(".stateful_policy")
	f.SpecField(".target_pools")
	f.SpecField(".target_size")
	f.SpecField(".update_policy")
	f.SpecField(".versions")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".current_actions")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".instance_group")
	f.StatusField(".region")
	f.StatusField(".self_link")
	f.StatusField(".status")
	f.StatusField(".update_policy")
	f.StatusField(".zone")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".all_instances_config")
	f.Unimplemented_NotYetTriaged(".instance_flexibility_policy")
	f.Unimplemented_NotYetTriaged(".instance_lifecycle_policy")
	f.Unimplemented_NotYetTriaged(".list_managed_instances_results")
	f.Unimplemented_NotYetTriaged(".resource_policies")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".standby_policy")
	f.Unimplemented_NotYetTriaged(".target_stopped_size")
	f.Unimplemented_NotYetTriaged(".target_suspended_size")

	// Inner nested current_actions fields that are not in KRM
	f.Unimplemented_NotYetTriaged(".current_actions.starting")
	f.Unimplemented_NotYetTriaged(".current_actions.resuming")
	f.Unimplemented_NotYetTriaged(".current_actions.stopping")
	f.Unimplemented_NotYetTriaged(".current_actions.suspending")

	// Inner nested status fields that are not in KRM
	f.Unimplemented_NotYetTriaged(".status.all_instances_config")
	f.Unimplemented_NotYetTriaged(".status.standby_policy")

	f.FilterSpec = func(in *pb.InstanceGroupManager) {
		if in.TargetSize == nil {
			in.TargetSize = direct.PtrTo(int32(0))
		}
	}

	return f
}
