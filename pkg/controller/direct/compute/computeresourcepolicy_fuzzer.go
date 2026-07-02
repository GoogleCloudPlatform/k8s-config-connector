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
// proto.message: google.cloud.compute.v1.ResourcePolicy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeResourcePolicyFuzzer())
}

func computeResourcePolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ResourcePolicy{},
		ComputeResourcePolicySpec_v1beta1_FromProto, ComputeResourcePolicySpec_v1beta1_ToProto,
		ComputeResourcePolicyObservedState_v1beta1_FromProto, ComputeResourcePolicyObservedState_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".region")
	f.SpecField(".description")
	f.SpecField(".disk_consistency_group_policy")
	f.SpecField(".disk_consistency_group_policy.enabled")
	f.SpecField(".group_placement_policy")
	f.SpecField(".group_placement_policy.availability_domain_count")
	f.SpecField(".group_placement_policy.collocation")
	f.SpecField(".group_placement_policy.vm_count")
	f.SpecField(".instance_schedule_policy")
	f.SpecField(".instance_schedule_policy.expiration_time")
	f.SpecField(".instance_schedule_policy.start_time")
	f.SpecField(".instance_schedule_policy.time_zone")
	f.SpecField(".instance_schedule_policy.vm_start_schedule")
	f.SpecField(".instance_schedule_policy.vm_start_schedule.schedule")
	f.SpecField(".instance_schedule_policy.vm_stop_schedule")
	f.SpecField(".instance_schedule_policy.vm_stop_schedule.schedule")
	f.SpecField(".snapshot_schedule_policy")
	f.SpecField(".snapshot_schedule_policy.retention_policy")
	f.SpecField(".snapshot_schedule_policy.retention_policy.max_retention_days")
	f.SpecField(".snapshot_schedule_policy.retention_policy.on_source_disk_delete")
	f.SpecField(".snapshot_schedule_policy.schedule")
	f.SpecField(".snapshot_schedule_policy.schedule.daily_schedule")
	f.SpecField(".snapshot_schedule_policy.schedule.daily_schedule.days_in_cycle")
	f.SpecField(".snapshot_schedule_policy.schedule.daily_schedule.start_time")
	f.SpecField(".snapshot_schedule_policy.schedule.hourly_schedule")
	f.SpecField(".snapshot_schedule_policy.schedule.hourly_schedule.hours_in_cycle")
	f.SpecField(".snapshot_schedule_policy.schedule.hourly_schedule.start_time")
	f.SpecField(".snapshot_schedule_policy.schedule.weekly_schedule")
	f.SpecField(".snapshot_schedule_policy.schedule.weekly_schedule.day_of_weeks")
	f.SpecField(".snapshot_schedule_policy.schedule.weekly_schedule.day_of_weeks[].day")
	f.SpecField(".snapshot_schedule_policy.schedule.weekly_schedule.day_of_weeks[].start_time")
	f.SpecField(".snapshot_schedule_policy.snapshot_properties")
	f.SpecField(".snapshot_schedule_policy.snapshot_properties.chain_name")
	f.SpecField(".snapshot_schedule_policy.snapshot_properties.guest_flush")
	f.SpecField(".snapshot_schedule_policy.snapshot_properties.labels")
	f.SpecField(".snapshot_schedule_policy.snapshot_properties.storage_locations")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".id")
	f.StatusField(".status")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".self_link")
	f.Unimplemented_NotYetTriaged(".self_link_with_id")
	f.Unimplemented_NotYetTriaged(".resource_status")
	f.Unimplemented_NotYetTriaged(".workload_policy")
	f.Unimplemented_NotYetTriaged(".workload_policy.max_topology_distance")
	f.Unimplemented_NotYetTriaged(".workload_policy.accelerator_topology")
	f.Unimplemented_NotYetTriaged(".workload_policy.type")
	f.Unimplemented_NotYetTriaged(".group_placement_policy.accelerator_topology_mode")
	f.Unimplemented_NotYetTriaged(".group_placement_policy.gpu_topology")
	f.Unimplemented_NotYetTriaged(".group_placement_policy.max_distance")
	f.Unimplemented_NotYetTriaged(".snapshot_schedule_policy.schedule.daily_schedule.duration")
	f.Unimplemented_NotYetTriaged(".snapshot_schedule_policy.schedule.hourly_schedule.duration")
	f.Unimplemented_NotYetTriaged(".snapshot_schedule_policy.schedule.weekly_schedule.day_of_weeks[].duration")

	return f
}
