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


// +kcc:proto=google.cloud.deploy.v1.AdvanceChildRolloutJob
type AdvanceChildRolloutJob struct {
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRolloutMetadata
type AutomationRolloutMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.ChildRolloutJobs
type ChildRolloutJobs struct {
}

// +kcc:proto=google.cloud.deploy.v1.CloudRunMetadata
type CloudRunMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.CreateChildRolloutJob
type CreateChildRolloutJob struct {
}

// +kcc:proto=google.cloud.deploy.v1.CustomMetadata
type CustomMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.DeployJob
type DeployJob struct {
}

// +kcc:proto=google.cloud.deploy.v1.DeploymentJobs
type DeploymentJobs struct {
}

// +kcc:proto=google.cloud.deploy.v1.Job
type Job struct {
}

// +kcc:proto=google.cloud.deploy.v1.Metadata
type Metadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.Phase
type Phase struct {
}

// +kcc:proto=google.cloud.deploy.v1.PostdeployJob
type PostdeployJob struct {
}

// +kcc:proto=google.cloud.deploy.v1.PredeployJob
type PredeployJob struct {
}

// +kcc:proto=google.cloud.deploy.v1.Rollout
type Rollout struct {
	// Optional. Name of the `Rollout`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/{rollout}`.
	//  The `rollout` component must match `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.name
	Name *string `json:"name,omitempty"`

	// Description of the `Rollout` for user purposes. Max length is 255
	//  characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.description
	Description *string `json:"description,omitempty"`

	// User annotations. These attributes can only be set and used by the
	//  user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations
	//  for more details such as format and size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels are attributes that can be set and used by both the
	//  user and by Cloud Deploy. Labels must meet the following constraints:
	//
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//  underscores, and dashes.
	//  * All characters must use UTF-8 encoding, and international characters are
	//  allowed.
	//  * Keys must start with a lowercase letter or international character.
	//  * Each resource is limited to a maximum of 64 labels.
	//
	//  Both keys and values are additionally constrained to be <= 128 bytes.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The ID of Target to which this `Rollout` is deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.target_id
	TargetID *string `json:"targetID,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.VerifyJob
type VerifyJob struct {
}

// +kcc:proto=google.cloud.deploy.v1.AutomationRolloutMetadata
type AutomationRolloutMetadataObservedState struct {
	// Output only. The name of the AutomationRun initiated by a promote release
	//  rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRolloutMetadata.promote_automation_run
	PromoteAutomationRun *string `json:"promoteAutomationRun,omitempty"`

	// Output only. The names of the AutomationRuns initiated by an advance
	//  rollout rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRolloutMetadata.advance_automation_runs
	AdvanceAutomationRuns []string `json:"advanceAutomationRuns,omitempty"`

	// Output only. The names of the AutomationRuns initiated by a repair rollout
	//  rule.
	// +kcc:proto:field=google.cloud.deploy.v1.AutomationRolloutMetadata.repair_automation_runs
	RepairAutomationRuns []string `json:"repairAutomationRuns,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.ChildRolloutJobs
type ChildRolloutJobsObservedState struct {
	// Output only. List of CreateChildRolloutJobs
	// +kcc:proto:field=google.cloud.deploy.v1.ChildRolloutJobs.create_rollout_jobs
	CreateRolloutJobs []Job `json:"createRolloutJobs,omitempty"`

	// Output only. List of AdvanceChildRolloutJobs
	// +kcc:proto:field=google.cloud.deploy.v1.ChildRolloutJobs.advance_rollout_jobs
	AdvanceRolloutJobs []Job `json:"advanceRolloutJobs,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CloudRunMetadata
type CloudRunMetadataObservedState struct {
	// Output only. The name of the Cloud Run Service that is associated with a
	//  `Rollout`. Format is
	//  `projects/{project}/locations/{location}/services/{service}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunMetadata.service
	Service *string `json:"service,omitempty"`

	// Output only. The Cloud Run Service urls that are associated with a
	//  `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunMetadata.service_urls
	ServiceUrls []string `json:"serviceUrls,omitempty"`

	// Output only. The Cloud Run Revision id associated with a `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunMetadata.revision
	Revision *string `json:"revision,omitempty"`

	// Output only. The name of the Cloud Run job that is associated with a
	//  `Rollout`. Format is
	//  `projects/{project}/locations/{location}/jobs/{job_name}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunMetadata.job
	Job *string `json:"job,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomMetadata
type CustomMetadataObservedState struct {
	// Output only. Key-value pairs provided by the user-defined operation.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomMetadata.values
	Values map[string]string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeploymentJobs
type DeploymentJobsObservedState struct {
	// Output only. The deploy Job. This is the deploy job in the phase.
	// +kcc:proto:field=google.cloud.deploy.v1.DeploymentJobs.deploy_job
	DeployJob *Job `json:"deployJob,omitempty"`

	// Output only. The verify Job. Runs after a deploy if the deploy succeeds.
	// +kcc:proto:field=google.cloud.deploy.v1.DeploymentJobs.verify_job
	VerifyJob *Job `json:"verifyJob,omitempty"`

	// Output only. The predeploy Job, which is the first job on the phase.
	// +kcc:proto:field=google.cloud.deploy.v1.DeploymentJobs.predeploy_job
	PredeployJob *Job `json:"predeployJob,omitempty"`

	// Output only. The postdeploy Job, which is the last job on the phase.
	// +kcc:proto:field=google.cloud.deploy.v1.DeploymentJobs.postdeploy_job
	PostdeployJob *Job `json:"postdeployJob,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Job
type JobObservedState struct {
	// Output only. The ID of the Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.id
	ID *string `json:"id,omitempty"`

	// Output only. The current state of the Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information on why the Job was skipped, if
	//  available.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.skip_message
	SkipMessage *string `json:"skipMessage,omitempty"`

	// Output only. The name of the `JobRun` responsible for the most recent
	//  invocation of this Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.job_run
	JobRun *string `json:"jobRun,omitempty"`

	// Output only. A deploy Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.deploy_job
	DeployJob *DeployJob `json:"deployJob,omitempty"`

	// Output only. A verify Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.verify_job
	VerifyJob *VerifyJob `json:"verifyJob,omitempty"`

	// Output only. A predeploy Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.predeploy_job
	PredeployJob *PredeployJob `json:"predeployJob,omitempty"`

	// Output only. A postdeploy Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.postdeploy_job
	PostdeployJob *PostdeployJob `json:"postdeployJob,omitempty"`

	// Output only. A createChildRollout Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.create_child_rollout_job
	CreateChildRolloutJob *CreateChildRolloutJob `json:"createChildRolloutJob,omitempty"`

	// Output only. An advanceChildRollout Job.
	// +kcc:proto:field=google.cloud.deploy.v1.Job.advance_child_rollout_job
	AdvanceChildRolloutJob *AdvanceChildRolloutJob `json:"advanceChildRolloutJob,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Metadata
type MetadataObservedState struct {
	// Output only. The name of the Cloud Run Service that is associated with a
	//  `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Metadata.cloud_run
	CloudRun *CloudRunMetadata `json:"cloudRun,omitempty"`

	// Output only. AutomationRolloutMetadata contains the information about the
	//  interactions between Automation service and this rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.Metadata.automation
	Automation *AutomationRolloutMetadata `json:"automation,omitempty"`

	// Output only. Custom metadata provided by user-defined `Rollout` operations.
	// +kcc:proto:field=google.cloud.deploy.v1.Metadata.custom
	Custom *CustomMetadata `json:"custom,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Phase
type PhaseObservedState struct {
	// Output only. The ID of the Phase.
	// +kcc:proto:field=google.cloud.deploy.v1.Phase.id
	ID *string `json:"id,omitempty"`

	// Output only. Current state of the Phase.
	// +kcc:proto:field=google.cloud.deploy.v1.Phase.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information on why the Phase was skipped, if
	//  available.
	// +kcc:proto:field=google.cloud.deploy.v1.Phase.skip_message
	SkipMessage *string `json:"skipMessage,omitempty"`

	// Output only. Deployment job composition.
	// +kcc:proto:field=google.cloud.deploy.v1.Phase.deployment_jobs
	DeploymentJobs *DeploymentJobs `json:"deploymentJobs,omitempty"`

	// Output only. ChildRollout job composition.
	// +kcc:proto:field=google.cloud.deploy.v1.Phase.child_rollout_jobs
	ChildRolloutJobs *ChildRolloutJobs `json:"childRolloutJobs,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PostdeployJob
type PostdeployJobObservedState struct {
	// Output only. The custom actions that the postdeploy Job executes.
	// +kcc:proto:field=google.cloud.deploy.v1.PostdeployJob.actions
	Actions []string `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PredeployJob
type PredeployJobObservedState struct {
	// Output only. The custom actions that the predeploy Job executes.
	// +kcc:proto:field=google.cloud.deploy.v1.PredeployJob.actions
	Actions []string `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.Rollout
type RolloutObservedState struct {
	// Output only. Unique identifier of the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `Rollout` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the `Rollout` was approved.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.approve_time
	ApproveTime *string `json:"approveTime,omitempty"`

	// Output only. Time at which the `Rollout` was enqueued.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.enqueue_time
	EnqueueTime *string `json:"enqueueTime,omitempty"`

	// Output only. Time at which the `Rollout` started deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploy_start_time
	DeployStartTime *string `json:"deployStartTime,omitempty"`

	// Output only. Time at which the `Rollout` finished deploying.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploy_end_time
	DeployEndTime *string `json:"deployEndTime,omitempty"`

	// Output only. Approval state of the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.approval_state
	ApprovalState *string `json:"approvalState,omitempty"`

	// Output only. Current state of the `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the rollout failure, if
	//  available.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.failure_reason
	FailureReason *string `json:"failureReason,omitempty"`

	// Output only. The resource name of the Cloud Build `Build` object that is
	//  used to deploy the Rollout. Format is
	//  `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploying_build
	DeployingBuild *string `json:"deployingBuild,omitempty"`

	// Output only. The reason this rollout failed. This will always be
	//  unspecified while the rollout is in progress.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.deploy_failure_cause
	DeployFailureCause *string `json:"deployFailureCause,omitempty"`

	// Output only. The phases that represent the workflows of this `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.phases
	Phases []Phase `json:"phases,omitempty"`

	// Output only. Metadata contains information about the rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.metadata
	Metadata *Metadata `json:"metadata,omitempty"`

	// Output only. Name of the `ControllerRollout`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/{rollout}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.controller_rollout
	ControllerRollout *string `json:"controllerRollout,omitempty"`

	// Output only. Name of the `Rollout` that is rolled back by this `Rollout`.
	//  Empty if this `Rollout` wasn't created as a rollback.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.rollback_of_rollout
	RollbackOfRollout *string `json:"rollbackOfRollout,omitempty"`

	// Output only. Names of `Rollouts` that rolled back this `Rollout`.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.rolled_back_by_rollouts
	RolledBackByRollouts []string `json:"rolledBackByRollouts,omitempty"`

	// Output only. The AutomationRun actively repairing the rollout.
	// +kcc:proto:field=google.cloud.deploy.v1.Rollout.active_repair_automation_run
	ActiveRepairAutomationRun *string `json:"activeRepairAutomationRun,omitempty"`
}
