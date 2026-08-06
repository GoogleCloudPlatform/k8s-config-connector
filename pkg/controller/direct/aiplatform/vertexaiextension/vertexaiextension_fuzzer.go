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
// proto.message: google.cloud.aiplatform.v1beta1.Extension
// api.group: aiplatform.cnrm.cloud.google.com

package vertexaiextension

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIExtensionFuzzer())
}

func vertexAIExtensionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Extension{},
		aiplatform.VertexAIExtensionSpec_FromProto, aiplatform.VertexAIExtensionSpec_ToProto,
		aiplatform.VertexAIExtensionObservedState_FromProto, aiplatform.VertexAIExtensionObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".manifest")
	f.SpecField(".manifest.name")
	f.SpecField(".manifest.description")
	f.SpecField(".manifest.api_spec")
	f.SpecField(".manifest.api_spec.open_api_yaml")
	f.SpecField(".manifest.api_spec.open_api_gcs_uri")
	f.SpecField(".manifest.auth_config")
	f.SpecField(".manifest.auth_config.api_key_config")
	f.SpecField(".manifest.auth_config.api_key_config.name")
	f.SpecField(".manifest.auth_config.api_key_config.api_key_secret")
	f.SpecField(".manifest.auth_config.api_key_config.http_element_location")
	f.SpecField(".manifest.auth_config.google_service_account_config")
	f.SpecField(".manifest.auth_config.google_service_account_config.service_account")
	f.SpecField(".manifest.auth_config.http_basic_auth_config")
	f.SpecField(".manifest.auth_config.http_basic_auth_config.credential_secret")
	f.SpecField(".manifest.auth_config.oauth_config")
	f.SpecField(".manifest.auth_config.oauth_config.access_token")
	f.SpecField(".manifest.auth_config.oauth_config.service_account")
	f.SpecField(".manifest.auth_config.oidc_config")
	f.SpecField(".manifest.auth_config.oidc_config.id_token")
	f.SpecField(".manifest.auth_config.oidc_config.service_account")
	f.SpecField(".manifest.auth_config.auth_type")
	f.SpecField(".runtime_config")
	f.SpecField(".runtime_config.code_interpreter_runtime_config")
	f.SpecField(".runtime_config.code_interpreter_runtime_config.file_input_gcs_bucket")
	f.SpecField(".runtime_config.code_interpreter_runtime_config.file_output_gcs_bucket")
	f.SpecField(".runtime_config.vertex_ai_search_runtime_config")
	f.SpecField(".runtime_config.vertex_ai_search_runtime_config.serving_config")
	f.SpecField(".runtime_config.vertex_ai_search_runtime_config.engine")
	f.SpecField(".runtime_config.default_params")
	f.SpecField(".tool_use_examples")
	f.SpecField(".private_service_connect_config")
	f.SpecField(".private_service_connect_config.service_directory")

	// Status fields (ObservedState)
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".extension_operations")

	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.parameters")
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.parameters_json_schema")
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.response")
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.response_json_schema")

	f.Unimplemented_Etag()

	return f
}
