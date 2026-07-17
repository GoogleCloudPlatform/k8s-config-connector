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
// proto.message: google.cloud.compute.v1.Snapshot
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeSnapshotFuzzer())
}

func computeSnapshotFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Snapshot{},
		ComputeSnapshotSpec_v1beta1_FromProto, ComputeSnapshotSpec_v1beta1_ToProto,
		ComputeSnapshotStatus_v1beta1_FromProto, ComputeSnapshotStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".chain_name")
	f.SpecField(".description")
	f.SpecField(".snapshot_encryption_key")
	f.SpecField(".snapshot_encryption_key.kms_key_name")
	f.SpecField(".snapshot_encryption_key.kms_key_service_account")
	f.SpecField(".snapshot_encryption_key.raw_key")
	f.SpecField(".snapshot_encryption_key.sha256")
	f.SpecField(".source_disk_encryption_key")
	f.SpecField(".source_disk_encryption_key.kms_key_service_account")
	f.SpecField(".source_disk_encryption_key.raw_key")
	f.SpecField(".source_disk")
	f.SpecField(".storage_locations")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".label_fingerprint")
	f.StatusField(".licenses")
	f.StatusField(".self_link")
	f.StatusField(".storage_bytes")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".architecture")
	f.Unimplemented_NotYetTriaged(".auto_created")
	f.Unimplemented_NotYetTriaged(".creation_size_bytes")
	f.Unimplemented_NotYetTriaged(".download_bytes")
	f.Unimplemented_NotYetTriaged(".enable_confidential_compute")
	f.Unimplemented_NotYetTriaged(".guest_flush")
	f.Unimplemented_NotYetTriaged(".guest_os_features")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".license_codes")
	f.Unimplemented_NotYetTriaged(".location_hint")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".snapshot_type")
	f.Unimplemented_NotYetTriaged(".source_disk_for_recovery_checkpoint")
	f.Unimplemented_NotYetTriaged(".source_disk_id")
	f.Unimplemented_NotYetTriaged(".source_instant_snapshot")
	f.Unimplemented_NotYetTriaged(".source_instant_snapshot_encryption_key")
	f.Unimplemented_NotYetTriaged(".source_instant_snapshot_id")
	f.Unimplemented_NotYetTriaged(".source_snapshot_schedule_policy")
	f.Unimplemented_NotYetTriaged(".source_snapshot_schedule_policy_id")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".storage_bytes_status")

	f.Unimplemented_NotYetTriaged(".zone")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".params.resource_manager_tags")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".disk_size_gb")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".snapshot_group_id")
	f.Unimplemented_NotYetTriaged(".snapshot_group_name")
	f.Unimplemented_NotYetTriaged(".snapshot_encryption_key.rsa_encrypted_key")
	f.Unimplemented_NotYetTriaged(".source_disk_encryption_key.rsa_encrypted_key")
	f.Unimplemented_NotYetTriaged(".source_disk_encryption_key.sha256")
	f.Unimplemented_NotYetTriaged(".source_disk_encryption_key.kms_key_name")

	return f
}
