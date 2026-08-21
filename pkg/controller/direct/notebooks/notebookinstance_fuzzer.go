// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.notebooks.v1.Instance
// api.group: notebooks.cnrm.cloud.google.com

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(notebookInstanceFuzzer())
}

func notebookInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		NotebookInstanceSpec_v1beta1_FromProto, NotebookInstanceSpec_v1beta1_ToProto,
		NotebookInstanceObservedState_v1beta1_FromProto, NotebookInstanceObservedState_v1beta1_ToProto,
	)

	f.Unimplemented_Identity(".name") // special field

	// Field comparison between KRM Spec (NotebookInstanceSpec) and GCP Proto (google.cloud.notebooks.v1.Instance):
	// - Zone (KRM Spec, Parent) maps to location segment in GCP parent URL path.
	// - ProjectRef (KRM Spec, Parent) maps to project segment in GCP parent URL path.
	// - ResourceID (KRM Spec) maps to the last segment of the resource .name.
	// - VMImage (KRM Spec) maps to `.vm_image` (GCP Proto).
	// - ContainerImage (KRM Spec) maps to `.container_image` (GCP Proto).
	// - PostStartupScript (KRM Spec) maps to `.post_startup_script` (GCP Proto).
	// - InstanceOwners (KRM Spec) maps to `.instance_owners` (GCP Proto).
	// - ServiceAccountRef (KRM Spec) maps to `.service_account` (GCP Proto).
	// - ServiceAccountScopes (KRM Spec) maps to `.service_account_scopes` (GCP Proto).
	// - MachineType (KRM Spec) maps to `.machine_type` (GCP Proto).
	// - AcceleratorConfig (KRM Spec) maps to `.accelerator_config` (GCP Proto).
	// - InstallGpuDriver (KRM Spec) maps to `.install_gpu_driver` (GCP Proto).
	// - CustomGpuDriverPath (KRM Spec) maps to `.custom_gpu_driver_path` (GCP Proto).
	// - BootDiskType (KRM Spec) maps to `.boot_disk_type` (GCP Proto).
	// - BootDiskSizeGB (KRM Spec) maps to `.boot_disk_size_gb` (GCP Proto).
	// - DataDiskType (KRM Spec) maps to `.data_disk_type` (GCP Proto).
	// - DataDiskSizeGB (KRM Spec) maps to `.data_disk_size_gb` (GCP Proto).
	// - NoRemoveDataDisk (KRM Spec) maps to `.no_remove_data_disk` (GCP Proto).
	// - DiskEncryption (KRM Spec) maps to `.disk_encryption` (GCP Proto).
	// - KMSKeyRef (KRM Spec) maps to `.kms_key` (GCP Proto).
	// - ShieldedInstanceConfig (KRM Spec) maps to `.shielded_instance_config` (GCP Proto).
	// - NoPublicIP (KRM Spec) maps to `.no_public_ip` (GCP Proto).
	// - NoProxyAccess (KRM Spec) maps to `.no_proxy_access` (GCP Proto).
	// - NetworkRef (KRM Spec) maps to `.network` (GCP Proto).
	// - SubnetRef (KRM Spec) maps to `.subnet` (GCP Proto).
	// - Labels (KRM Spec) maps to `.labels` (GCP Proto).
	// - Metadata (KRM Spec) maps to `.metadata` (GCP Proto).
	// - Tags (KRM Spec) maps to `.tags` (GCP Proto).
	// - UpgradeHistory (KRM Spec) maps to `.upgrade_history` (GCP Proto).
	// - NicType (KRM Spec) maps to `.nic_type` (GCP Proto).
	// - ReservationAffinity (KRM Spec) maps to `.reservation_affinity` (GCP Proto).
	// - CanIPForward (KRM Spec) maps to `.can_ip_forward` (GCP Proto).
	//
	// Field comparison between KRM Status (NotebookInstanceObservedState) and GCP Proto (google.cloud.notebooks.v1.Instance):
	// - ProxyURI (KRM Status) maps to `.proxy_uri` (GCP Proto).
	// - State (KRM Status) maps to `.state` (GCP Proto).
	// - Disks (KRM Status) maps to `.disks` (GCP Proto).
	// - Creator (KRM Status) maps to `.creator` (GCP Proto).
	// - CreateTime (KRM Status) maps to `.create_time` (GCP Proto).
	// - UpdateTime (KRM Status) maps to `.update_time` (GCP Proto).

	f.SpecField(".vm_image")
	f.SpecField(".container_image")
	f.SpecField(".post_startup_script")
	f.SpecField(".instance_owners")
	f.SpecField(".service_account")
	f.SpecField(".service_account_scopes")
	f.SpecField(".machine_type")
	f.SpecField(".accelerator_config")
	f.SpecField(".install_gpu_driver")
	f.SpecField(".custom_gpu_driver_path")
	f.SpecField(".boot_disk_type")
	f.SpecField(".boot_disk_size_gb")
	f.SpecField(".data_disk_type")
	f.SpecField(".data_disk_size_gb")
	f.SpecField(".no_remove_data_disk")
	f.SpecField(".disk_encryption")
	f.SpecField(".kms_key")
	f.SpecField(".shielded_instance_config")
	f.SpecField(".no_public_ip")
	f.SpecField(".no_proxy_access")
	f.SpecField(".network")
	f.SpecField(".subnet")
	f.SpecField(".labels")
	f.SpecField(".metadata")
	f.SpecField(".tags")
	f.SpecField(".upgrade_history")
	f.SpecField(".nic_type")
	f.SpecField(".reservation_affinity")
	f.SpecField(".can_ip_forward")

	f.StatusField(".proxy_uri")
	f.StatusField(".state")
	f.StatusField(".disks")
	f.StatusField(".creator")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
