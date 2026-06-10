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
// proto.message: google.cloud.aiplatform.v1.Model
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/types/known/structpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aiplatformModelFuzzer())
}

func aiplatformModelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Model{},
		AIPlatformModelSpec_FromProto, AIPlatformModelSpec_ToProto,
		AIPlatformModelObservedState_FromProto, AIPlatformModelObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")                     // special field
	f.UnimplementedFields.Insert(".original_model_info")      // NOTYET
	f.UnimplementedFields.Insert(".satisfies_pzi")            // NOTYET
	f.UnimplementedFields.Insert(".satisfies_pzs")            // NOTYET
	f.UnimplementedFields.Insert(".supported_export_formats") // NOTYET
	f.UnimplementedFields.Insert(".default_checkpoint_id")
	f.UnimplementedFields.Insert(".container_spec.health_probe.success_threshold")
	f.UnimplementedFields.Insert(".container_spec.health_probe.initial_delay_seconds")
	f.UnimplementedFields.Insert(".container_spec.startup_probe.http_get")
	f.UnimplementedFields.Insert(".container_spec.startup_probe.success_threshold")
	f.UnimplementedFields.Insert(".container_spec.startup_probe.grpc")
	f.UnimplementedFields.Insert(".container_spec.startup_probe.tcp_socket")
	f.UnimplementedFields.Insert(".container_spec.startup_probe.failure_threshold")
	f.UnimplementedFields.Insert(".container_spec.health_probe.tcp_socket")
	f.UnimplementedFields.Insert(".container_spec.health_probe.failure_threshold")
	f.UnimplementedFields.Insert(".container_spec.liveness_probe")
	f.UnimplementedFields.Insert(".container_spec.invoke_route_prefix")
	f.UnimplementedFields.Insert(".base_model_source.model_garden_source.skip_hf_model_cache")
	f.UnimplementedFields.Insert(".base_model_source.model_garden_source.version_id")
	f.UnimplementedFields.Insert(".container_spec.health_probe.grpc")
	f.UnimplementedFields.Insert(".container_spec.health_probe.http_get")
	f.UnimplementedFields.Insert(".container_spec.startup_probe.initial_delay_seconds")
	f.UnimplementedFields.Insert(".explanation_spec.parameters.examples.example_gcs_source")
	f.UnimplementedFields.Insert(".explanation_spec.explanation_parameters.output_indices")
	f.UnimplementedFields.Insert(".explanation_spec.parameters.output_indices")
	f.UnimplementedFields.Insert(".explanation_spec.metadata.inputs.visualization.overlay_type")
	f.UnimplementedFields.Insert(".explanation_spec.metadata.inputs.visualization.polarity")
	f.UnimplementedFields.Insert(".explanation_spec.metadata.inputs.visualization.type")
	f.UnimplementedFields.Insert(".explanation_spec.metadata.inputs.visualization.clear_point")
	f.UnimplementedFields.Insert(".metadata.list_value")
	f.UnimplementedFields.Insert(".explanation_spec.parameters.examples.nearest_neighbor_search_config.list_value")
	f.UnimplementedFields.Insert(".checkpoints")

	f.SpecFields.Insert(".version_aliases")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".version_description")
	f.SpecFields.Insert(".predict_schemata")
	f.SpecFields.Insert(".metadata_schema_uri")
	f.SpecFields.Insert(".metadata")
	f.SpecFields.Insert(".pipeline_job")
	f.SpecFields.Insert(".container_spec")
	f.SpecFields.Insert(".artifact_uri")
	f.SpecFields.Insert(".explanation_spec")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".data_stats")
	f.SpecFields.Insert(".encryption_spec")
	f.SpecFields.Insert(".base_model_source")

	f.StatusFields.Insert(".version_id")
	f.StatusFields.Insert(".version_create_time")
	f.StatusFields.Insert(".version_update_time")
	f.StatusFields.Insert(".training_pipeline")
	f.StatusFields.Insert(".supported_deployment_resources_types")
	f.StatusFields.Insert(".supported_input_storage_formats")
	f.StatusFields.Insert(".supported_output_storage_formats")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".deployed_models")
	f.StatusFields.Insert(".model_source_info")
	f.StatusFields.Insert(".metadata_artifact")

	f.Unimplemented_Etag()

	f.FilterSpec = func(in *pb.Model) {
		if in.ExplanationSpec != nil {
			if in.ExplanationSpec.Metadata != nil {
				for _, input := range in.ExplanationSpec.Metadata.Inputs {
					if input.Visualization != nil {
						input.Visualization.Type = 0
						input.Visualization.Polarity = 0
						input.Visualization.OverlayType = 0
					}
					for _, baseline := range input.InputBaselines {
						if _, ok := baseline.Kind.(*structpb.Value_ListValue); ok {
							baseline.Kind = nil
						}
						if _, ok := baseline.Kind.(*structpb.Value_StructValue); ok {
							baseline.Kind = nil
						}
					}
					for _, baseline := range input.EncodedBaselines {
						if _, ok := baseline.Kind.(*structpb.Value_ListValue); ok {
							baseline.Kind = nil
						}
						if _, ok := baseline.Kind.(*structpb.Value_StructValue); ok {
							baseline.Kind = nil
						}
					}
				}
				for _, output := range in.ExplanationSpec.Metadata.Outputs {
					if output.GetIndexDisplayNameMapping() != nil {
						output.DisplayNameMapping = nil
					}
				}
			}
			if in.ExplanationSpec.Parameters != nil && in.ExplanationSpec.Parameters.GetExamples() != nil && in.ExplanationSpec.Parameters.GetExamples().GetPresets() != nil {
				in.ExplanationSpec.Parameters.GetExamples().GetPresets().Query = nil
			}
		}
	}
	return f
}
