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

package aiplatform

import (
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/types/known/structpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIExtensionFuzzer())
}

func vertexAIExtensionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&aiplatformpb.Extension{},
		VertexAIExtensionSpec_FromProto, VertexAIExtensionSpec_ToProto,
		VertexAIExtensionObservedState_FromProto, VertexAIExtensionObservedState_ToProto,
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
	f.SpecField(".manifest.auth_config.http_basic_auth_config")
	f.SpecField(".manifest.auth_config.http_basic_auth_config.credential_secret")
	f.SpecField(".manifest.auth_config.google_service_account_config")
	f.SpecField(".manifest.auth_config.google_service_account_config.service_account")
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
	f.SpecField(".runtime_config.vertex_ai_search_runtime_config.serving_config_name")
	f.SpecField(".runtime_config.vertex_ai_search_runtime_config.engine_id")
	f.SpecField(".runtime_config.default_params")
	f.SpecField(".tool_use_examples")
	f.SpecField(".tool_use_examples[].display_name")
	f.SpecField(".tool_use_examples[].query")
	f.SpecField(".tool_use_examples[].extension_operation")
	f.SpecField(".tool_use_examples[].extension_operation.extension")
	f.SpecField(".tool_use_examples[].extension_operation.operation_id")
	f.SpecField(".tool_use_examples[].function_name")
	f.SpecField(".tool_use_examples[].request_params")
	f.SpecField(".tool_use_examples[].response_params")
	f.SpecField(".tool_use_examples[].response_summary")
	f.SpecField(".private_service_connect_config")
	f.SpecField(".private_service_connect_config.service_directory")

	// Status fields (ObservedState)
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".etag")
	f.StatusField(".extension_operations")
	f.StatusField(".extension_operations[].operation_id")
	f.StatusField(".extension_operations[].function_declaration")
	f.StatusField(".extension_operations[].function_declaration.name")
	f.StatusField(".extension_operations[].function_declaration.description")

	// Unimplemented fields during fuzz roundtrip
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.parameters")
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.parameters_json_schema")
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.response")
	f.Unimplemented_NotYetTriaged(".extension_operations[].function_declaration.response_json_schema")

	f.FilterSpec = func(in *aiplatformpb.Extension) {
		if in.RuntimeConfig != nil && in.RuntimeConfig.DefaultParams != nil {
			clearUnsupportedStructFields(in.RuntimeConfig.DefaultParams)
		}
		if in.ToolUseExamples != nil {
			for _, example := range in.ToolUseExamples {
				if example.RequestParams != nil {
					clearUnsupportedStructFields(example.RequestParams)
				}
				if example.ResponseParams != nil {
					clearUnsupportedStructFields(example.ResponseParams)
				}
			}
		}
	}

	return f
}

func clearUnsupportedStructFields(s *structpb.Struct) {
	if s == nil {
		return
	}
	for _, val := range s.Fields {
		clearUnsupportedValueFields(val)
	}
}
