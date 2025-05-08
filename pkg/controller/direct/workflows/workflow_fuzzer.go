// Copyright 2025 Google LLC
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

package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzWorkflowsWorkflow())
}

func fuzzWorkflowsWorkflow() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Workflow{},
		WorkflowsWorkflowSpec_FromProto, WorkflowsWorkflowSpec_ToProto,
		WorkflowsWorkflowObservedState_FromProto, WorkflowsWorkflowObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".service_account")
	f.SpecFields.Insert(".source_contents")
	f.SpecFields.Insert(".crypto_key_name")
	f.SpecFields.Insert(".call_log_level")
	f.SpecFields.Insert(".user_env_vars")
	f.SpecFields.Insert(".execution_history_level")
	f.SpecFields.Insert(".tags")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".revision_id")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".revision_create_time")
	f.StatusFields.Insert(".state_error")
	f.StatusFields.Insert(".state_error.details")
	f.StatusFields.Insert(".state_error.type")
	f.StatusFields.Insert(".all_kms_keys")
	f.StatusFields.Insert(".all_kms_keys_versions")
	f.StatusFields.Insert(".crypto_key_version")
	return f
}
