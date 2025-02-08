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

package v1alpha1


// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignmentReport
type OSPolicyAssignmentReport struct {
	// The `OSPolicyAssignmentReport` API resource name.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location}/instances/{instance_id}/osPolicyAssignments/{os_policy_assignment_id}/report`
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.name
	Name *string `json:"name,omitempty"`

	// The Compute Engine VM instance name.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.instance
	Instance *string `json:"instance,omitempty"`

	// Reference to the `OSPolicyAssignment` API resource that the `OSPolicy`
	//  belongs to.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id@revision_id}`
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.os_policy_assignment
	OsPolicyAssignment *string `json:"osPolicyAssignment,omitempty"`

	// Compliance data for each `OSPolicy` that is applied to the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.os_policy_compliances
	OsPolicyCompliances []OSPolicyAssignmentReport_OSPolicyCompliance `json:"osPolicyCompliances,omitempty"`

	// Timestamp for when the report was last generated.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Unique identifier of the last attempted run to apply the OS policies
	//  associated with this assignment on the VM.
	//
	//  This ID is logged by the OS Config agent while applying the OS
	//  policies associated with this assignment on the VM.
	//  NOTE: If the service is unable to successfully connect to the agent for
	//  this run, then this id will not be available in the agent logs.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.last_run_id
	LastRunID *string `json:"lastRunID,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance
type OSPolicyAssignmentReport_OSPolicyCompliance struct {
	// The OS policy id
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.os_policy_id
	OsPolicyID *string `json:"osPolicyID,omitempty"`

	// The compliance state of the OS policy.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.compliance_state
	ComplianceState *string `json:"complianceState,omitempty"`

	// The reason for the OS policy to be in an unknown compliance state.
	//  This field is always populated when `compliance_state` is `UNKNOWN`.
	//
	//  If populated, the field can contain one of the following values:
	//
	//  * `vm-not-running`: The VM was not running.
	//  * `os-policies-not-supported-by-agent`: The version of the OS Config
	//  agent running on the VM does not support running OS policies.
	//  * `no-agent-detected`: The OS Config agent is not detected for the VM.
	//  * `resource-execution-errors`: The OS Config agent encountered errors
	//  while executing one or more resources in the policy. See
	//  `os_policy_resource_compliances` for details.
	//  * `task-timeout`: The task sent to the agent to apply the policy timed
	//  out.
	//  * `unexpected-agent-state`: The OS Config agent did not report the final
	//  status of the task that attempted to apply the policy. Instead, the agent
	//  unexpectedly started working on a different task. This mostly happens
	//  when the agent or VM unexpectedly restarts while applying OS policies.
	//  * `internal-service-errors`: Internal service errors were encountered
	//  while attempting to apply the policy.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.compliance_state_reason
	ComplianceStateReason *string `json:"complianceStateReason,omitempty"`

	// Compliance data for each resource within the policy that is applied to
	//  the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.os_policy_resource_compliances
	OsPolicyResourceCompliances []OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance `json:"osPolicyResourceCompliances,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance
type OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance struct {
	// The ID of the OS policy resource.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.os_policy_resource_id
	OsPolicyResourceID *string `json:"osPolicyResourceID,omitempty"`

	// Ordered list of configuration completed by the agent for the OS policy
	//  resource.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.config_steps
	ConfigSteps []OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep `json:"configSteps,omitempty"`

	// The compliance state of the resource.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.compliance_state
	ComplianceState *string `json:"complianceState,omitempty"`

	// A reason for the resource to be in the given compliance state.
	//  This field is always populated when `compliance_state` is `UNKNOWN`.
	//
	//  The following values are supported when `compliance_state == UNKNOWN`
	//
	//  * `execution-errors`: Errors were encountered by the agent while
	//  executing the resource and the compliance state couldn't be
	//  determined.
	//  * `execution-skipped-by-agent`: Resource execution was skipped by the
	//  agent because errors were encountered while executing prior resources
	//  in the OS policy.
	//  * `os-policy-execution-attempt-failed`: The execution of the OS policy
	//  containing this resource failed and the compliance state couldn't be
	//  determined.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.compliance_state_reason
	ComplianceStateReason *string `json:"complianceStateReason,omitempty"`

	// ExecResource specific output.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.exec_resource_output
	ExecResourceOutput *OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput `json:"execResourceOutput,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.ExecResourceOutput
type OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput struct {
	// Output from enforcement phase output file (if run).
	//  Output size is limited to 100K bytes.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.ExecResourceOutput.enforcement_output
	EnforcementOutput []byte `json:"enforcementOutput,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.OSPolicyResourceConfigStep
type OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep struct {
	// Configuration step type.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.OSPolicyResourceConfigStep.type
	Type *string `json:"type,omitempty"`

	// An error message recorded during the execution of this step.
	//  Only populated if errors were encountered during this step execution.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignmentReport.OSPolicyCompliance.OSPolicyResourceCompliance.OSPolicyResourceConfigStep.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}
