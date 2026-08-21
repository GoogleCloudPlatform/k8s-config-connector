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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInstanceGroupManagerFuzzer())
}

func computeInstanceGroupManagerStatus_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManager) *krm.ComputeInstanceGroupManagerStatus {
	out := ComputeInstanceGroupManagerStatus_v1beta1_FromProto(mapCtx, in)
	return &out
}

func computeInstanceGroupManagerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InstanceGroupManager{},
		ComputeInstanceGroupManagerSpec_v1beta1_FromProto, ComputeInstanceGroupManagerSpec_v1beta1_ToProto,
		computeInstanceGroupManagerStatus_FromProto, ComputeInstanceGroupManagerStatus_v1beta1_ToProto,
	)

	// Field comparison: ComputeInstanceGroupManagerSpec vs pb.InstanceGroupManager Proto
	// - Spec.AutoHealingPolicies         maps to proto field .auto_healing_policies
	// - Spec.BaseInstanceName            maps to proto field .base_instance_name
	// - Spec.Description                 maps to proto field .description
	// - Spec.DistributionPolicy          maps to proto field .distribution_policy
	// - Spec.FailoverAction              not represented/mapped in pb.InstanceGroupManager (but check if exists)
	// - Spec.InstanceTemplateRef         maps to proto field .instance_template
	// - Spec.Location                    not represented/mapped in pb.InstanceGroupManager (handled at controller level)
	// - Spec.NamedPorts                  maps to proto field .named_ports
	// - Spec.ProjectRef                  not represented/mapped in pb.InstanceGroupManager
	// - Spec.ResourceID                  maps to proto field .name (via Identity)
	// - Spec.ServiceAccountRef           maps to proto field .service_account (not mapped currently?)
	// - Spec.StatefulPolicy              maps to proto field .stateful_policy
	// - Spec.TargetPools                 maps to proto field .target_pools
	// - Spec.TargetSize                  maps to proto field .target_size
	// - Spec.UpdatePolicy                maps to proto field .update_policy
	// - Spec.Versions                    maps to proto field .versions

	// Field comparison: ComputeInstanceGroupManagerStatus vs pb.InstanceGroupManager Proto
	// - Status.Conditions                not represented in pb.InstanceGroupManager
	// - Status.CreationTimestamp         maps to proto field .creation_timestamp
	// - Status.CurrentActions            maps to proto field .current_actions
	// - Status.Fingerprint               maps to proto field .fingerprint
	// - Status.Id                        maps to proto field .id
	// - Status.InstanceGroup             maps to proto field .instance_group
	// - Status.ObservedGeneration        not represented in pb.InstanceGroupManager
	// - Status.Region                    maps to proto field .region
	// - Status.SelfLink                  maps to proto field .self_link
	// - Status.Status                    maps to proto field .status
	// - Status.UpdatePolicy              not represented/mapped in pb.InstanceGroupManager Status?
	// - Status.Zone                      maps to proto field .zone

	// Spec fields
	f.SpecField(".auto_healing_policies")
	f.SpecField(".auto_healing_policies[].health_check")
	f.SpecField(".auto_healing_policies[].initial_delay_sec")
	f.SpecField(".base_instance_name")
	f.SpecField(".description")
	f.SpecField(".distribution_policy")
	f.SpecField(".distribution_policy.target_shape")
	f.SpecField(".distribution_policy.zones")
	f.SpecField(".distribution_policy.zones[].zone")
	f.SpecField(".instance_template")
	f.SpecField(".named_ports")
	f.SpecField(".named_ports[].name")
	f.SpecField(".named_ports[].port")
	f.SpecField(".stateful_policy")
	f.SpecField(".stateful_policy.preserved_state")
	f.SpecField(".stateful_policy.preserved_state.disks")
	f.SpecField(".stateful_policy.preserved_state.disks[].auto_delete")
	f.SpecField(".stateful_policy.preserved_state.external_ips")
	f.SpecField(".stateful_policy.preserved_state.external_ips[].auto_delete")
	f.SpecField(".stateful_policy.preserved_state.internal_ips")
	f.SpecField(".stateful_policy.preserved_state.internal_ips[].auto_delete")
	f.SpecField(".target_pools")
	f.SpecField(".target_size")
	f.SpecField(".update_policy")
	f.SpecField(".update_policy.instance_redistribution_type")
	f.SpecField(".update_policy.max_surge")
	f.SpecField(".update_policy.max_surge.fixed")
	f.SpecField(".update_policy.max_surge.percent")
	f.SpecField(".update_policy.max_unavailable")
	f.SpecField(".update_policy.max_unavailable.fixed")
	f.SpecField(".update_policy.max_unavailable.percent")
	f.SpecField(".update_policy.min_ready_sec")
	f.SpecField(".update_policy.minimal_action")
	f.SpecField(".update_policy.replacement_method")
	f.SpecField(".update_policy.type")
	f.SpecField(".versions")
	f.SpecField(".versions[].instance_template")
	f.SpecField(".versions[].name")
	f.SpecField(".versions[].target_size")
	f.SpecField(".versions[].target_size.calculated")
	f.SpecField(".versions[].target_size.fixed")
	f.SpecField(".versions[].target_size.percent")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".current_actions")
	f.StatusField(".current_actions.abandoning")
	f.StatusField(".current_actions.creating")
	f.StatusField(".current_actions.creating_without_retries")
	f.StatusField(".current_actions.deleting")
	f.StatusField(".current_actions.none")
	f.StatusField(".current_actions.recreating")
	f.StatusField(".current_actions.refreshing")
	f.StatusField(".current_actions.restarting")
	f.StatusField(".current_actions.verifying")
	f.StatusField(".fingerprint")
	f.StatusField(".id")
	f.StatusField(".instance_group")
	f.StatusField(".region")
	f.StatusField(".self_link")
	f.StatusField(".status")
	f.StatusField(".status.autoscaler")
	f.StatusField(".status.is_stable")
	f.StatusField(".status.stateful")
	f.StatusField(".status.stateful.has_stateful_config")
	f.StatusField(".status.stateful.per_instance_configs")
	f.StatusField(".status.stateful.per_instance_configs.all_effective")
	f.StatusField(".status.version_target")
	f.StatusField(".status.version_target.is_reached")
	f.StatusField(".zone")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".all_instances_config")
	f.Unimplemented_NotYetTriaged(".failover_action")
	f.Unimplemented_NotYetTriaged(".instance_flexibility_policy")
	f.Unimplemented_NotYetTriaged(".instance_lifecycle_policy")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".list_managed_instances_results")
	f.Unimplemented_NotYetTriaged(".resource_policies")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".service_account")
	f.Unimplemented_NotYetTriaged(".standby_policy")
	f.Unimplemented_NotYetTriaged(".target_stopped_size")
	f.Unimplemented_NotYetTriaged(".target_suspended_size")
	f.Unimplemented_NotYetTriaged(".current_actions.resuming")
	f.Unimplemented_NotYetTriaged(".current_actions.starting")
	f.Unimplemented_NotYetTriaged(".current_actions.stopping")
	f.Unimplemented_NotYetTriaged(".current_actions.suspending")
	f.Unimplemented_NotYetTriaged(".status.all_instances_config")
	f.Unimplemented_NotYetTriaged(".update_policy.most_disruptive_allowed_action")
	f.Unimplemented_NotYetTriaged(".update_policy.max_surge.calculated")
	f.Unimplemented_NotYetTriaged(".update_policy.max_unavailable.calculated")
	f.Unimplemented_NotYetTriaged(".target_size_policy")
	f.Unimplemented_NotYetTriaged(".status.current_instance_statuses")
	f.Unimplemented_NotYetTriaged(".status.bulk_instance_operation")
	f.Unimplemented_NotYetTriaged(".status.applied_accelerator_topologies")

	// FilterSpec handles the required non-pointer `target_size` in the KRM Spec
	// (which defaults to 0 on round-trip if originally nil in the proto message).
	f.FilterSpec = func(in *pb.InstanceGroupManager) {
		if in.TargetSize == nil {
			var zero int32 = 0
			in.TargetSize = &zero
		}
	}

	return f
}
