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
// proto.message: google.cloud.aiplatform.v1.Model
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
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
	f.SpecFields.Insert(".etag")
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

	return f
}
