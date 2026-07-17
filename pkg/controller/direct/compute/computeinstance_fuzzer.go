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
// proto.message: google.cloud.compute.v1.Instance
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInstanceFuzzer())
}

func computeInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		ComputeInstanceSpec_v1beta1_FromProto, ComputeInstanceSpec_v1beta1_ToProto,
		ComputeInstanceStatus_v1beta1_FromProto, ComputeInstanceStatus_v1beta1_ToProto,
	)

	// Field comparison: ComputeInstance Spec vs pb.Instance Proto
	// - Spec.DeletionProtection          maps to proto field .deletion_protection
	// - Spec.Description                 maps to proto field .description
	// - Spec.Hostname                    maps to proto field .hostname
	// - Spec.MachineType                 maps to proto field .machine_type
	// - Spec.Metadata                    maps to proto field .metadata
	// - Spec.ResourcePolicies            maps to proto field .resource_policies
	// - Spec.Tags                        maps to proto field .tags
	// - Spec.Zone                        maps to proto field .zone

	// Spec fields
	f.SpecField(".deletion_protection")
	f.SpecField(".description")
	f.SpecField(".hostname")
	f.SpecField(".machine_type")
	f.SpecField(".metadata")
	f.SpecField(".resource_policies")
	f.SpecField(".tags")
	f.SpecField(".zone")

	// Status fields
	f.StatusField(".cpu_platform")
	f.StatusField(".status")
	f.StatusField(".id")
	f.StatusField(".label_fingerprint")
	f.StatusField(".metadata.fingerprint")
	f.StatusField(".self_link")
	f.StatusField(".tags.fingerprint")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".advanced_machine_features")
	f.Unimplemented_NotYetTriaged(".can_ip_forward")
	f.Unimplemented_NotYetTriaged(".confidential_instance_config")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".disks")
	f.Unimplemented_NotYetTriaged(".display_device")
	f.Unimplemented_NotYetTriaged(".fingerprint")
	f.Unimplemented_NotYetTriaged(".guest_accelerators")
	f.Unimplemented_NotYetTriaged(".instance_encryption_key")
	f.Unimplemented_NotYetTriaged(".key_revocation_action_type")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".last_start_timestamp")
	f.Unimplemented_NotYetTriaged(".last_stop_timestamp")
	f.Unimplemented_NotYetTriaged(".last_suspended_timestamp")
	f.Unimplemented_NotYetTriaged(".metadata.kind")
	f.Unimplemented_NotYetTriaged(".min_cpu_platform")
	f.Unimplemented_NotYetTriaged(".network_interfaces")
	f.Unimplemented_NotYetTriaged(".network_performance_config")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".private_ipv6_google_access")
	f.Unimplemented_NotYetTriaged(".reservation_affinity")
	f.Unimplemented_NotYetTriaged(".resource_status")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".scheduling")
	f.Unimplemented_NotYetTriaged(".service_accounts")
	f.Unimplemented_NotYetTriaged(".shielded_instance_config")
	f.Unimplemented_NotYetTriaged(".shielded_instance_integrity_policy")
	f.Unimplemented_NotYetTriaged(".source_machine_image")
	f.Unimplemented_NotYetTriaged(".source_machine_image_encryption_key")
	f.Unimplemented_NotYetTriaged(".start_restricted")
	f.Unimplemented_NotYetTriaged(".status_message")
	f.Unimplemented_NotYetTriaged(".workload_identity_config")

	return f
}
