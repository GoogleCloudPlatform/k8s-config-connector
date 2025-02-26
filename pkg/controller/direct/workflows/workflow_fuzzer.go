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

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".serviceAccount")
	f.SpecFields.Insert(".sourceContents")
	f.SpecFields.Insert(".cryptoKeyName")
	f.SpecFields.Insert(".callLogLevel")
	f.SpecFields.Insert(".userEnvVars")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".revisionId")
	f.StatusFields.Insert(".createTime")
	f.StatusFields.Insert(".updateTime")
	f.StatusFields.Insert(".revisionCreateTime")
	f.StatusFields.Insert(".stateError.details")
	f.StatusFields.Insert(".stateError.type")
	return f
}
