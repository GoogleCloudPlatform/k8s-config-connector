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
// proto.message: google.cloud.aiplatform.v1.PipelineJob
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIPipelineJobFuzzer())
}

func vertexAIPipelineJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PipelineJob{},
		VertexAIPipelineJobSpec_FromProto, VertexAIPipelineJobSpec_ToProto,
		VertexAIPipelineJobObservedState_FromProto, VertexAIPipelineJobObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".pipeline_spec")
	f.SpecField(".labels")
	f.SpecField(".runtime_config")
	f.SpecField(".encryption_spec")
	f.SpecField(".service_account")
	f.SpecField(".network")
	f.SpecField(".reserved_ip_ranges")
	f.SpecField(".psc_interface_config")
	f.SpecField(".template_uri")
	f.SpecField(".preflight_validations")

	// Status fields (ObservedState)
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".job_detail")
	f.StatusField(".error")
	f.StatusField(".template_metadata")
	f.StatusField(".schedule_name")

	// Unmapped fields to ignore in the fuzzer roundtrip
	f.Unimplemented_NotYetTriaged(".runtime_config.input_artifacts")
	f.Unimplemented_NotYetTriaged(".runtime_config.parameter_values")
	f.Unimplemented_NotYetTriaged(".runtime_config.parameters")
	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".job_detail.task_details[].inputs")
	f.Unimplemented_NotYetTriaged(".job_detail.task_details[].outputs")
	f.Unimplemented_NotYetTriaged(".job_detail.task_details[].error.details")
	f.Unimplemented_NotYetTriaged(".job_detail.task_details[].pipeline_task_status[].error.details")
	f.Unimplemented_NotYetTriaged(".job_detail.task_details[].pipeline_task_status[].error.details[].type_url")
	f.Unimplemented_NotYetTriaged(".job_detail.task_details[].pipeline_task_status[].error.details[].value")

	return f
}
