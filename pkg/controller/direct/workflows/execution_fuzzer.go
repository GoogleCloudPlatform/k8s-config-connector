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
// proto.message: google.cloud.workflows.executions.v1.Execution
// api.group: workflows.cnrm.cloud.google.com

package workflows

import (
	pb "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(workflowsExecutionFuzzer())
}

func workflowsExecutionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Execution{},
		WorkflowsExecutionSpec_FromProto, WorkflowsExecutionSpec_ToProto,
		WorkflowsExecutionObservedState_FromProto, WorkflowsExecutionObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".argument")
	f.SpecFields.Insert(".call_log_level")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".start_time")
	f.StatusFields.Insert(".end_time")
	f.StatusFields.Insert(".duration")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".result")
	f.StatusFields.Insert(".error")
	f.StatusFields.Insert(".workflow_revision_id")
	f.StatusFields.Insert(".status")
	f.StatusFields.Insert(".state_error")

	return f
}
