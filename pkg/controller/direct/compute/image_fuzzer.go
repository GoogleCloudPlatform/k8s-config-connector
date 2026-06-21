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
// proto.message: google.cloud.compute.v1.Image
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeImageFuzzer())
}

func computeImageFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Image{},
		ComputeImageSpec_v1beta1_FromProto, ComputeImageSpec_v1beta1_ToProto,
		ComputeImageStatus_v1beta1_FromProto, ComputeImageStatus_v1beta1_ToProto,
	)

	// Field comparison of KRM ComputeImageSpec to proto fields:
	// - description => .description (SpecField)
	// - diskRef => .source_disk (SpecField)
	// - diskSizeGb => .disk_size_gb (SpecField)
	// - family => .family (SpecField)
	// - guestOsFeatures => .guest_os_features (SpecField)
	// - imageEncryptionKey => .image_encryption_key (SpecField)
	// - licenses => .licenses (SpecField)
	// - rawDisk => .raw_disk (SpecField)
	// - resourceID => (ignored, used for name/identity)
	// - sourceImageRef => .source_image (SpecField)
	// - sourceSnapshotRef => .source_snapshot (SpecField)
	// - storageLocations => .storage_locations (SpecField)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".disk_size_gb")
	f.SpecField(".family")
	f.SpecField(".licenses")
	f.SpecField(".storage_locations")

	// Spec Reference fields (mapped to flat strings in proto)
	f.SpecField(".source_disk")
	f.SpecField(".source_image")
	f.SpecField(".source_snapshot")

	// Spec Sub-structures
	f.SpecField(".guest_os_features")
	f.SpecField(".guest_os_features[].type")

	f.SpecField(".raw_disk")
	f.SpecField(".raw_disk.container_type")
	f.SpecField(".raw_disk.sha1_checksum")
	f.SpecField(".raw_disk.source")

	f.SpecField(".image_encryption_key")
	f.SpecField(".image_encryption_key.kms_key_name")
	f.SpecField(".image_encryption_key.kms_key_service_account")

	// Field comparison of KRM ComputeImageStatus to proto fields:
	// - archiveSizeBytes => .archive_size_bytes (StatusField)
	// - creationTimestamp => .creation_timestamp (StatusField)
	// - labelFingerprint => .label_fingerprint (StatusField)
	// - selfLink => .self_link (StatusField)

	// Status fields
	f.StatusField(".archive_size_bytes")
	f.StatusField(".creation_timestamp")
	f.StatusField(".label_fingerprint")
	f.StatusField(".self_link")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Labels / Annotations
	f.Unimplemented_LabelsAnnotations(".labels")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".architecture")
	f.Unimplemented_NotYetTriaged(".deprecated")
	f.Unimplemented_NotYetTriaged(".enable_confidential_compute")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".license_codes")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".shielded_instance_initial_state")
	f.Unimplemented_NotYetTriaged(".source_disk_encryption_key")
	f.Unimplemented_NotYetTriaged(".source_disk_id")
	f.Unimplemented_NotYetTriaged(".source_image_encryption_key")
	f.Unimplemented_NotYetTriaged(".source_image_id")
	f.Unimplemented_NotYetTriaged(".source_snapshot_encryption_key")
	f.Unimplemented_NotYetTriaged(".source_snapshot_id")
	f.Unimplemented_NotYetTriaged(".source_type")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".params")

	// Sub-structure leaf fields not mapped
	f.Unimplemented_NotYetTriaged(".image_encryption_key.raw_key")
	f.Unimplemented_NotYetTriaged(".image_encryption_key.rsa_encrypted_key")
	f.Unimplemented_NotYetTriaged(".image_encryption_key.sha256")

	f.FilterSpec = func(in *pb.Image) {
		if in.RawDisk != nil && in.RawDisk.ContainerType == nil && in.RawDisk.Sha1Checksum == nil && in.RawDisk.Source == nil {
			in.RawDisk = nil
		}
		if in.ImageEncryptionKey != nil && in.ImageEncryptionKey.KmsKeyName == nil && in.ImageEncryptionKey.KmsKeyServiceAccount == nil {
			in.ImageEncryptionKey = nil
		}
	}

	f.FilterStatus = func(in *pb.Image) {
		in.Description = nil
		in.DiskSizeGb = nil
		in.Family = nil
		in.GuestOsFeatures = nil
		in.ImageEncryptionKey = nil
		in.Licenses = nil
		in.RawDisk = nil
		in.SourceDisk = nil
		in.SourceImage = nil
		in.SourceSnapshot = nil
		in.StorageLocations = nil
	}

	return f
}
