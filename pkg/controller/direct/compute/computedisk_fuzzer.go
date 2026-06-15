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
// proto.message: google.cloud.compute.v1.Disk
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeDiskFuzzer())
}

func computeDiskFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Disk{},
		ComputeDiskSpec_v1beta1_FromProto, ComputeDiskSpec_v1beta1_ToProto,
		ComputeDiskStatus_v1beta1_FromProto, ComputeDiskStatus_v1beta1_ToProto,
	)

	// Filter empty elements from GuestOsFeatures list to avoid roundtrip mismatches
	f.FilterSpec = func(in *pb.Disk) {
		var filtered []*pb.GuestOsFeature
		for _, feat := range in.GuestOsFeatures {
			if feat.GetType() != "" {
				filtered = append(filtered, feat)
			}
		}
		in.GuestOsFeatures = filtered
	}

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".enable_confidential_compute")
	f.SpecField(".guest_os_features")
	f.SpecField(".guest_os_features[].type")
	f.SpecField(".interface")
	f.SpecField(".licenses")
	f.SpecField(".multi_writer")
	f.SpecField(".physical_block_size_bytes")
	f.SpecField(".provisioned_iops")
	f.SpecField(".provisioned_throughput")
	f.SpecField(".replica_zones")
	f.SpecField(".resource_policies")
	f.SpecField(".size_gb")
	f.SpecField(".source_disk")
	f.SpecField(".source_image")
	f.SpecField(".source_snapshot")
	f.SpecField(".type")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".label_fingerprint")
	f.StatusField(".last_attach_timestamp")
	f.StatusField(".last_detach_timestamp")
	f.StatusField(".self_link")
	f.StatusField(".source_disk_id")
	f.StatusField(".source_image_id")
	f.StatusField(".source_snapshot_id")
	f.StatusField(".users")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".architecture")
	f.Unimplemented_NotYetTriaged(".async_secondary_disks")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".license_codes")
	f.Unimplemented_NotYetTriaged(".location_hint")
	f.Unimplemented_NotYetTriaged(".options")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".resource_status")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".source_storage_object")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".storage_pool")
	f.Unimplemented_NotYetTriaged(".zone")
	f.Unimplemented_NotYetTriaged(".access_mode")

	// Nested structs with potential empty-struct/nil discrepancies under proto generation
	f.Unimplemented_NotYetTriaged(".async_primary_disk")
	f.Unimplemented_NotYetTriaged(".disk_encryption_key")
	f.Unimplemented_NotYetTriaged(".source_image_encryption_key")
	f.Unimplemented_NotYetTriaged(".source_snapshot_encryption_key")

	// Additional nested unimplemented fields identified by fuzz runner
	f.Unimplemented_NotYetTriaged(".async_primary_disk.consistency_group_policy")
	f.Unimplemented_NotYetTriaged(".async_primary_disk.consistency_group_policy_id")
	f.Unimplemented_NotYetTriaged(".async_primary_disk.disk_id")
	f.Unimplemented_NotYetTriaged(".source_consistency_group_policy")
	f.Unimplemented_NotYetTriaged(".source_consistency_group_policy_id")
	f.Unimplemented_NotYetTriaged(".source_instant_snapshot")
	f.Unimplemented_NotYetTriaged(".source_instant_snapshot_id")
	f.Unimplemented_NotYetTriaged(".source_image_encryption_key.rsa_encrypted_key")
	f.Unimplemented_NotYetTriaged(".source_snapshot_encryption_key.rsa_encrypted_key")

	return f
}
