// Copyright 2025 Google LLC
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
// proto.message: google.cloud.run.v2.Job
// api.group: run.cnrm.cloud.google.com

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(runJobFuzzer())
}

func runJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Job{},
		RunJobSpec_FromProto, RunJobSpec_ToProto,
		RunJobObservedState_FromProto, RunJobObservedState_ToProto,
	)

	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".binary_authorization")
	f.SpecFields.Insert(".client")
	f.SpecFields.Insert(".client_version")
	f.SpecFields.Insert(".launch_stage")
	f.SpecFields.Insert(".template")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".creator")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".execution_count")
	f.StatusFields.Insert(".expire_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".last_modifier")
	f.StatusFields.Insert(".latest_created_execution")
	f.StatusFields.Insert(".reconciling")
	f.StatusFields.Insert(".terminal_condition")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".update_time")

	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".binary_authorization.policy")
	f.Unimplemented_NotYetTriaged(".binary_authorization.use_default")
	f.Unimplemented_NotYetTriaged(".conditions")
	f.Unimplemented_NotYetTriaged(".generation")
	f.Unimplemented_NotYetTriaged(".observed_generation")
	f.Unimplemented_NotYetTriaged(".run_execution_token")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".template.labels")
	f.Unimplemented_NotYetTriaged(".template.template.containers.depends_on")
	f.Unimplemented_NotYetTriaged(".template.template.containers.base_image_uri")
	f.Unimplemented_NotYetTriaged(".template.template.node_selector")
	f.Unimplemented_NotYetTriaged(".template.template.volumes")
	f.Unimplemented_NotYetTriaged(".terminal_condition.execution_reason")

	f.UnimplementedFields.Insert(".name") // Read-only, special handling

	return f
}
