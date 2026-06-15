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
