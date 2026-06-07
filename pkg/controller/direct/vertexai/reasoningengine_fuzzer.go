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
// proto.message: google.cloud.aiplatform.v1beta1.ReasoningEngine
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIReasoningEngineFuzzer())
}

func vertexAIReasoningEngineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ReasoningEngine{},
		VertexAIReasoningEngineSpec_FromProto, VertexAIReasoningEngineSpec_ToProto,
		VertexAIReasoningEngineObservedState_FromProto, VertexAIReasoningEngineObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".spec")
	f.SpecFields.Insert(".context_spec")
	f.SpecFields.Insert(".encryption_spec")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.granular_ttl_config.create_ttl")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.granular_ttl_config.create_ttl.seconds")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.granular_ttl_config.generate_updated_ttl")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.granular_ttl_config.generate_updated_ttl.seconds")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.granular_ttl_config.generate_created_ttl")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.granular_ttl_config.generate_created_ttl.seconds")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.default_ttl")
	f.Unimplemented_NotYetTriaged(".context_spec.memory_bank_config.ttl_config.default_ttl.seconds")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".name")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.developer_connect_source")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.developer_connect_source.config")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.developer_connect_source.config.git_repository_link")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.developer_connect_source.config.revision")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.inline_source")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.inline_source.source_archive")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.python_spec")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.python_spec.entrypoint_module")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.python_spec.entrypoint_object")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec.python_spec.requirements_file")
	f.Unimplemented_NotYetTriaged(".spec.container_spec")

	return f
}
