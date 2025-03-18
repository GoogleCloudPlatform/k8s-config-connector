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
// proto.message: google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate
// api.group: colab.cnrm.cloud.google.com

package colab

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(colabRuntimeTemplateFuzzer())
}

func colabRuntimeTemplateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NotebookRuntimeTemplate{},
		ColabRuntimeTemplateSpec_FromProto, ColabRuntimeTemplateSpec_ToProto,
		ColabRuntimeTemplateObservedState_FromProto, ColabRuntimeTemplateObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")       // special field
	f.UnimplementedFields.Insert(".is_default") // deprecated field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".encryption_spec")
	f.SpecFields.Insert(".machine_spec")
	f.SpecFields.Insert(".data_persistent_disk_spec")
	f.SpecFields.Insert(".network_spec")
	f.SpecFields.Insert(".service_account")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".idle_shutdown_config")
	f.SpecFields.Insert(".euc_config")
	f.SpecFields.Insert(".notebook_runtime_type")
	f.SpecFields.Insert(".shielded_vm_config")
	f.SpecFields.Insert(".network_tags")

	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".euc_config")

	return f
}
