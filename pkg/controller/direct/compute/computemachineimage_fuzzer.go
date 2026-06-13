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
// proto.message: google.cloud.compute.v1.MachineImage
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeMachineImageFuzzer())
}

func computeMachineImageFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MachineImage{},
		ComputeMachineImageSpec_v1alpha1_FromProto, ComputeMachineImageSpec_v1alpha1_ToProto,
		ComputeMachineImageStatus_v1alpha1_FromProto, ComputeMachineImageStatus_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".guest_flush")
	f.SpecField(".machine_image_encryption_key")
	f.SpecField(".source_instance")

	f.StatusField(".self_link")
	f.StatusField(".storage_locations")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_Internal(".kind")

	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".instance_properties")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_NotYetTriaged(".machine_image_encryption_key.rsa_encrypted_key")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".saved_disks")
	f.Unimplemented_NotYetTriaged(".source_disk_encryption_keys")
	f.Unimplemented_NotYetTriaged(".source_instance_properties")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".total_storage_bytes")

	return f
}
