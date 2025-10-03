// Copyright 2025 Google LLC
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
// proto.message: google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate
// api.group: colab.cnrm.cloud.google.com

package colab

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(colabRuntimeFuzzer())
}

func colabRuntimeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NotebookRuntime{},
		ColabRuntimeSpec_FromProto, ColabRuntimeSpec_ToProto,
		ColabRuntimeObservedState_FromProto, ColabRuntimeObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")          // special field
	f.UnimplementedFields.Insert(".satisfies_pzs") // field for future use
	f.UnimplementedFields.Insert(".satisfies_pzi") // field for future use
	f.UnimplementedFields.Insert(".data_persistent_disk_spec")
	f.UnimplementedFields.Insert(".machine_spec")
	f.UnimplementedFields.Insert(".shielded_vm_config")
	f.UnimplementedFields.Insert(".software_config")
	f.UnimplementedFields.Insert(".network_spec")
	f.UnimplementedFields.Insert(".euc_config")

	f.SpecFields.Insert(".notebook_runtime_template_ref")
	f.SpecFields.Insert(".runtime_user")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".network_tags")

	f.StatusFields.Insert(".proxy_uri")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".health_state")
	f.StatusFields.Insert(".service_account")
	f.StatusFields.Insert(".runtime_state")
	f.StatusFields.Insert(".is_upgradable")
	f.StatusFields.Insert(".expiration_time")
	f.StatusFields.Insert(".version")
	f.StatusFields.Insert(".notebook_runtime_type")
	f.StatusFields.Insert(".idle_shutdown_config")
	f.StatusFields.Insert(".encryption_spec")

	return f
}
