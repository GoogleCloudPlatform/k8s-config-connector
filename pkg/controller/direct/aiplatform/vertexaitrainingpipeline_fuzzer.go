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
// proto.message: google.cloud.aiplatform.v1.TrainingPipeline
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAITrainingPipelineFuzzer())
}

func vertexAITrainingPipelineSpec_FromProto(mapCtx *direct.MapContext, in *pb.TrainingPipeline) *krm.VertexAITrainingPipelineSpec {
	out := VertexAITrainingPipelineSpec_FromProto(mapCtx, in)
	if out != nil && in.GetParentModel() != "" {
		out.ModelRef = &krm.AIPlatformModelRef{External: in.GetParentModel()}
	}
	return out
}

func vertexAITrainingPipelineSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAITrainingPipelineSpec) *pb.TrainingPipeline {
	out := VertexAITrainingPipelineSpec_ToProto(mapCtx, in)
	if out != nil && in.ModelRef != nil {
		out.ParentModel = in.ModelRef.External
	}
	return out
}

func vertexAITrainingPipelineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TrainingPipeline{},
		vertexAITrainingPipelineSpec_FromProto, vertexAITrainingPipelineSpec_ToProto,
		VertexAITrainingPipelineObservedState_FromProto, VertexAITrainingPipelineObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".input_data_config")
	f.SpecField(".training_task_definition")
	f.SpecField(".training_task_inputs")
	f.SpecField(".model_to_upload")
	f.SpecField(".model_id")
	f.SpecField(".parent_model")
	f.SpecField(".labels")
	f.SpecField(".encryption_spec")

	// Status fields (ObservedState)
	f.StatusField(".state")
	f.StatusField(".error")
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".update_time")
	f.StatusField(".training_task_metadata")

	// Model fields (nested model_to_upload)
	f.SpecField(".model_to_upload.version_aliases")
	f.SpecField(".model_to_upload.display_name")
	f.SpecField(".model_to_upload.description")
	f.SpecField(".model_to_upload.version_description")
	f.SpecField(".model_to_upload.predict_schemata")
	f.SpecField(".model_to_upload.metadata_schema_uri")
	f.SpecField(".model_to_upload.metadata")
	f.SpecField(".model_to_upload.pipeline_job")
	f.SpecField(".model_to_upload.container_spec")
	f.SpecField(".model_to_upload.artifact_uri")
	f.SpecField(".model_to_upload.explanation_spec")
	f.SpecField(".model_to_upload.labels")
	f.SpecField(".model_to_upload.data_stats")
	f.SpecField(".model_to_upload.encryption_spec")
	f.SpecField(".model_to_upload.base_model_source")

	f.StatusField(".model_to_upload.version_id")
	f.StatusField(".model_to_upload.version_create_time")
	f.StatusField(".model_to_upload.version_update_time")
	f.StatusField(".model_to_upload.training_pipeline")
	f.StatusField(".model_to_upload.supported_deployment_resources_types")
	f.StatusField(".model_to_upload.supported_input_storage_formats")
	f.StatusField(".model_to_upload.supported_output_storage_formats")
	f.StatusField(".model_to_upload.create_time")
	f.StatusField(".model_to_upload.update_time")
	f.StatusField(".model_to_upload.deployed_models")
	f.StatusField(".model_to_upload.model_source_info")
	f.StatusField(".model_to_upload.metadata_artifact")
	f.StatusField(".model_to_upload.satisfies_pzs")
	f.StatusField(".model_to_upload.satisfies_pzi")

	// Unmapped model fields / unimplemented fields
	f.Unimplemented_NotYetTriaged(".model_to_upload.etag")
	f.Unimplemented_NotYetTriaged(".model_to_upload.name")
	f.Unimplemented_NotYetTriaged(".model_to_upload.original_model_info")
	f.Unimplemented_NotYetTriaged(".model_to_upload.supported_export_formats")
	f.Unimplemented_NotYetTriaged(".model_to_upload.default_checkpoint_id")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.health_probe.success_threshold")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.health_probe.initial_delay_seconds")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.startup_probe.http_get")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.startup_probe.success_threshold")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.liveness_probe")
	f.Unimplemented_NotYetTriaged(".model_to_upload.explanation_spec.explanation_parameters.output_indices")
	f.Unimplemented_NotYetTriaged(".model_to_upload.metadata.list_value")

	f.Unimplemented_NotYetTriaged(".model_to_upload.checkpoints")
	f.Unimplemented_NotYetTriaged(".model_to_upload.checkpoints[].epoch")
	f.Unimplemented_NotYetTriaged(".model_to_upload.base_model_source.model_garden_source.skip_hf_model_cache")
	f.Unimplemented_NotYetTriaged(".model_to_upload.base_model_source.model_garden_source.version_id")
	f.Unimplemented_NotYetTriaged(".model_to_upload.explanation_spec.parameters.output_indices")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.invoke_route_prefix")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.startup_probe.tcp_socket")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.startup_probe.failure_threshold")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.health_probe.failure_threshold")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.health_probe.grpc")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.health_probe.tcp_socket")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.health_probe.http_get")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.startup_probe.grpc")
	f.Unimplemented_NotYetTriaged(".model_to_upload.container_spec.startup_probe.initial_delay_seconds")
	f.Unimplemented_NotYetTriaged(".model_to_upload.explanation_spec.parameters.examples.example_gcs_source.gcs_source")
	f.Unimplemented_NotYetTriaged(".model_to_upload.explanation_spec.parameters.examples.presets.query")

	// Unmapped root/pipeline fields
	f.Unimplemented_NotYetTriaged(".training_task_inputs.list_value")
	f.Unimplemented_NotYetTriaged(".training_task_metadata.list_value")
	f.Unimplemented_NotYetTriaged(".error.details")

	f.FilterSpec = func(in *pb.TrainingPipeline) {
		if in.TrainingTaskInputs != nil {
			clearUnsupportedValueFields(in.TrainingTaskInputs)
		}
		if in.ModelToUpload != nil {
			if in.ModelToUpload.Metadata != nil {
				clearUnsupportedValueFields(in.ModelToUpload.Metadata)
			}
			if in.ModelToUpload.ExplanationSpec != nil {
				if in.ModelToUpload.ExplanationSpec.Metadata != nil {
					for _, input := range in.ModelToUpload.ExplanationSpec.Metadata.Inputs {
						for _, b := range input.InputBaselines {
							clearUnsupportedValueFields(b)
						}
						for _, b := range input.EncodedBaselines {
							clearUnsupportedValueFields(b)
						}
						if input.Visualization != nil {
							input.Visualization.Type = 0
							input.Visualization.Polarity = 0
							input.Visualization.OverlayType = 0
						}
					}
					for _, output := range in.ModelToUpload.ExplanationSpec.Metadata.Outputs {
						clearUnsupportedValueFields(output.GetIndexDisplayNameMapping())
					}
				}
				if in.ModelToUpload.ExplanationSpec.Parameters != nil {
					if in.ModelToUpload.ExplanationSpec.Parameters.GetExamples() != nil {
						clearUnsupportedValueFields(in.ModelToUpload.ExplanationSpec.Parameters.GetExamples().GetNearestNeighborSearchConfig())
					}
				}
			}
		}
	}

	return f
}
