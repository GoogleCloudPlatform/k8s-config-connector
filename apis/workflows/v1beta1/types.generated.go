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

// +generated:types
// krm.group: workflows.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.workflows.v1
// resource: WorkflowsWorkflow:Workflow

package v1beta1

// +kcc:proto=google.cloud.workflows.v1.Workflow.StateError
type Workflow_StateError struct {
	// Provides specifics about the error.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.StateError.details
	Details *string `json:"details,omitempty"`

	// The type of this state error.
	// +kcc:proto:field=google.cloud.workflows.v1.Workflow.StateError.type
	Type *string `json:"type,omitempty"`
}
