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
// proto.message: google.cloud.aiplatform.v1.Schedule
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIScheduleFuzzer())
}

func vertexAIScheduleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Schedule{},
		VertexAIScheduleSpec_FromProto, VertexAIScheduleSpec_ToProto,
		VertexAIScheduleObservedState_FromProto, VertexAIScheduleObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".start_time")
	f.SpecField(".end_time")
	f.SpecField(".max_run_count")
	f.SpecField(".max_concurrent_run_count")
	f.SpecField(".allow_queueing")

	// time_specification oneof
	f.SpecField(".cron")

	// request oneof - pipeline job request is fully supported and triaged
	f.SpecField(".create_pipeline_job_request")

	// Status fields (ObservedState)
	f.StatusField(".started_run_count")
	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".next_run_time")
	f.StatusField(".last_pause_time")
	f.StatusField(".last_resume_time")
	f.StatusField(".catch_up")
	f.StatusField(".last_scheduled_run_response")

	// Parent fields of request are resolved via KRM projectRef/location, so they are not mapped in Spec/ObservedState
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.parent")

	// Spec-only fields that are not yet triaged or mapped in spec
	f.Unimplemented_NotYetTriaged(".max_concurrent_active_run_count")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job_id")

	// Ignore fields within the nested PipelineJob inside CreatePipelineJobRequest that are not fully mapped/triaged
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.name")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.runtime_config.input_artifacts")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.runtime_config.parameter_values")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.runtime_config.parameters")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.error")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.job_detail")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.create_time")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.start_time")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.end_time")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.update_time")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.state")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.template_metadata")
	f.Unimplemented_NotYetTriaged(".create_pipeline_job_request.pipeline_job.schedule_name")

	// CreateNotebookExecutionJobRequest is a highly complex nested structure whose details are skipped in KRM fuzzer to avoid endless roundtrip checks
	f.Unimplemented_NotYetTriaged(".create_notebook_execution_job_request")

	return f
}
