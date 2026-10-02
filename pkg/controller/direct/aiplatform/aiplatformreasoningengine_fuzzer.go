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
// proto.message: google.cloud.aiplatform.v1.ReasoningEngine
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aiplatformReasoningEngineFuzzer())
}

func aiplatformReasoningEngineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ReasoningEngine{},
		AIPlatformReasoningEngineSpec_FromProto, AIPlatformReasoningEngineSpec_ToProto,
		AIPlatformReasoningEngineObservedState_FromProto, AIPlatformReasoningEngineObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".spec.class_methods")
	f.Unimplemented_NotYetTriaged(".spec.source_code_spec")
	f.Unimplemented_NotYetTriaged(".spec.container_spec")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".spec")
	f.SpecField(".spec.service_account")
	f.SpecField(".spec.package_spec")
	f.SpecField(".spec.package_spec.pickle_object_gcs_uri")
	f.SpecField(".spec.package_spec.dependency_files_gcs_uri")
	f.SpecField(".spec.package_spec.requirements_gcs_uri")
	f.SpecField(".spec.package_spec.python_version")
	f.SpecField(".spec.deployment_spec")
	f.SpecField(".spec.deployment_spec.env")
	f.SpecField(".spec.deployment_spec.secret_env")
	f.SpecField(".spec.deployment_spec.psc_interface_config")
	f.SpecField(".spec.deployment_spec.min_instances")
	f.SpecField(".spec.deployment_spec.max_instances")
	f.SpecField(".spec.deployment_spec.resource_limits")
	f.SpecField(".spec.deployment_spec.container_concurrency")
	f.SpecField(".spec.agent_framework")
	f.SpecField(".encryption_spec")
	f.SpecField(".encryption_spec.kms_key_name")
	f.SpecField(".labels")

	// Status fields (ObservedState)
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
