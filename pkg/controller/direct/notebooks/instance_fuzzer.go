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
		NotebookInstanceSpec_FromProto, NotebookInstanceSpec_ToProto,
		NotebookInstanceObservedState_FromProto, NotebookInstanceObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".vm_image")
	f.SpecFields.Insert(".container_image")
	f.SpecFields.Insert(".post_startup_script")
	f.SpecFields.Insert(".instance_owners")
	f.SpecFields.Insert(".service_account")
	f.SpecFields.Insert(".service_account_scopes")
	f.SpecFields.Insert(".machine_type")
	f.SpecFields.Insert(".accelerator_config")
	f.SpecFields.Insert(".install_gpu_driver")
	f.SpecFields.Insert(".custom_gpu_driver_path")
	f.SpecFields.Insert(".boot_disk_type")
	f.SpecFields.Insert(".boot_disk_size_gb")
	f.SpecFields.Insert(".data_disk_type")
	f.SpecFields.Insert(".data_disk_size_gb")
	f.SpecFields.Insert(".no_remove_data_disk")
	f.SpecFields.Insert(".disk_encryption")
	f.SpecFields.Insert(".kms_key")
	f.SpecFields.Insert(".shielded_instance_config")
	f.SpecFields.Insert(".no_public_ip")
	f.SpecFields.Insert(".no_proxy_access")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".subnet")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".metadata")
	f.SpecFields.Insert(".tags")
	f.SpecFields.Insert(".upgrade_history")
	f.SpecFields.Insert(".nic_type")
	f.SpecFields.Insert(".reservation_affinity")
	f.SpecFields.Insert(".can_ip_forward")

	f.StatusFields.Insert(".proxy_uri")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".disks")
	f.StatusFields.Insert(".creator")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	return f
}
