// Copyright 2026 Google LLC
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
// proto.message: google.cloud.notebooks.v1.Schedule
// api.group: notebooks.cnrm.cloud.google.com

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(notebooksScheduleFuzzer())
}

func notebooksScheduleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Schedule{},
		NotebooksScheduleSpec_v1alpha1_FromProto, NotebooksScheduleSpec_v1alpha1_ToProto,
		NotebooksScheduleObservedState_v1alpha1_FromProto, NotebooksScheduleObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".description")
	f.SpecField(".state")
	f.SpecField(".cron_schedule")
	f.SpecField(".time_zone")
	f.SpecField(".execution_template")

	f.StatusField(".display_name")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".recent_executions")

	f.Unimplemented_NotYetTriaged(".recent_executions[].name")
	f.Unimplemented_NotYetTriaged(".recent_executions[].description")
	f.Unimplemented_NotYetTriaged(".recent_executions[].output_notebook_file")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.scale_tier")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.master_type")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.accelerator_config")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.accelerator_config.type")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.accelerator_config.core_count")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.labels")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.input_notebook_file")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.container_image_uri")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.output_notebook_folder")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.params_yaml_file")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.parameters")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.service_account")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.job_type")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.dataproc_parameters")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.dataproc_parameters.cluster")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.vertex_ai_parameters")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.vertex_ai_parameters.network")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.vertex_ai_parameters.env")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.kernel_spec")
	f.Unimplemented_NotYetTriaged(".recent_executions[].execution_template.tensorboard")

	return f
}
