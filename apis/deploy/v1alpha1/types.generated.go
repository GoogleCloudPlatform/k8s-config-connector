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


// +kcc:proto=google.cloud.deploy.v1.AdvanceChildRolloutJobRun
type AdvanceChildRolloutJobRun struct {
}

// +kcc:proto=google.cloud.deploy.v1.CloudRunMetadata
type CloudRunMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.CreateChildRolloutJobRun
type CreateChildRolloutJobRun struct {
}

// +kcc:proto=google.cloud.deploy.v1.CustomMetadata
type CustomMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.CustomTargetDeployMetadata
type CustomTargetDeployMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.DeployArtifact
type DeployArtifact struct {
}

// +kcc:proto=google.cloud.deploy.v1.DeployJobRun
type DeployJobRun struct {
}

// +kcc:proto=google.cloud.deploy.v1.DeployJobRunMetadata
type DeployJobRunMetadata struct {
}

// +kcc:proto=google.cloud.deploy.v1.JobRun
type JobRun struct {
	// Optional. Name of the `JobRun`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{releases}/rollouts/{rollouts}/jobRuns/{uuid}`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PostdeployJobRun
type PostdeployJobRun struct {
}

// +kcc:proto=google.cloud.deploy.v1.PredeployJobRun
type PredeployJobRun struct {
}

// +kcc:proto=google.cloud.deploy.v1.VerifyJobRun
type VerifyJobRun struct {
}

// +kcc:proto=google.cloud.deploy.v1.AdvanceChildRolloutJobRun
type AdvanceChildRolloutJobRunObservedState struct {
	// Output only. Name of the `ChildRollout`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/{rollout}`.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceChildRolloutJobRun.rollout
	Rollout *string `json:"rollout,omitempty"`

	// Output only. the ID of the ChildRollout's Phase.
	// +kcc:proto:field=google.cloud.deploy.v1.AdvanceChildRolloutJobRun.rollout_phase_id
	RolloutPhaseID *string `json:"rolloutPhaseID,omitempty"`
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

// +kcc:proto=google.cloud.deploy.v1.CreateChildRolloutJobRun
type CreateChildRolloutJobRunObservedState struct {
	// Output only. Name of the `ChildRollout`. Format is
	//  `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/{rollout}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CreateChildRolloutJobRun.rollout
	Rollout *string `json:"rollout,omitempty"`

	// Output only. The ID of the childRollout Phase initiated by this JobRun.
	// +kcc:proto:field=google.cloud.deploy.v1.CreateChildRolloutJobRun.rollout_phase_id
	RolloutPhaseID *string `json:"rolloutPhaseID,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomMetadata
type CustomMetadataObservedState struct {
	// Output only. Key-value pairs provided by the user-defined operation.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomMetadata.values
	Values map[string]string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTargetDeployMetadata
type CustomTargetDeployMetadataObservedState struct {
	// Output only. Skip message provided in the results of a custom deploy
	//  operation.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetDeployMetadata.skip_message
	SkipMessage *string `json:"skipMessage,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployArtifact
type DeployArtifactObservedState struct {
	// Output only. URI of a directory containing the artifacts. All paths are
	//  relative to this location.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployArtifact.artifact_uri
	ArtifactURI *string `json:"artifactURI,omitempty"`

	// Output only. File paths of the manifests applied during the deploy
	//  operation relative to the URI.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployArtifact.manifest_paths
	ManifestPaths []string `json:"manifestPaths,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployJobRun
type DeployJobRunObservedState struct {
	// Output only. The resource name of the Cloud Build `Build` object that is
	//  used to deploy. Format is
	//  `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRun.build
	Build *string `json:"build,omitempty"`

	// Output only. The reason the deploy failed. This will always be unspecified
	//  while the deploy is in progress or if it succeeded.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRun.failure_cause
	FailureCause *string `json:"failureCause,omitempty"`

	// Output only. Additional information about the deploy failure, if available.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRun.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`

	// Output only. Metadata containing information about the deploy job run.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRun.metadata
	Metadata *DeployJobRunMetadata `json:"metadata,omitempty"`

	// Output only. The artifact of a deploy job run, if available.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRun.artifact
	Artifact *DeployArtifact `json:"artifact,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployJobRunMetadata
type DeployJobRunMetadataObservedState struct {
	// Output only. The name of the Cloud Run Service that is associated with a
	//  `DeployJobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRunMetadata.cloud_run
	CloudRun *CloudRunMetadata `json:"cloudRun,omitempty"`

	// Output only. Custom Target metadata associated with a `DeployJobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRunMetadata.custom_target
	CustomTarget *CustomTargetDeployMetadata `json:"customTarget,omitempty"`

	// Output only. Custom metadata provided by user-defined deploy operation.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployJobRunMetadata.custom
	Custom *CustomMetadata `json:"custom,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.JobRun
type JobRunObservedState struct {
	// Output only. Unique identifier of the `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. ID of the `Rollout` phase this `JobRun` belongs in.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.phase_id
	PhaseID *string `json:"phaseID,omitempty"`

	// Output only. ID of the `Rollout` job this `JobRun` corresponds to.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.job_id
	JobID *string `json:"jobID,omitempty"`

	// Output only. Time at which the `JobRun` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the `JobRun` was started.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time at which the `JobRun` ended.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.state
	State *string `json:"state,omitempty"`

	// Output only. Information specific to a deploy `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.deploy_job_run
	DeployJobRun *DeployJobRun `json:"deployJobRun,omitempty"`

	// Output only. Information specific to a verify `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.verify_job_run
	VerifyJobRun *VerifyJobRun `json:"verifyJobRun,omitempty"`

	// Output only. Information specific to a predeploy `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.predeploy_job_run
	PredeployJobRun *PredeployJobRun `json:"predeployJobRun,omitempty"`

	// Output only. Information specific to a postdeploy `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.postdeploy_job_run
	PostdeployJobRun *PostdeployJobRun `json:"postdeployJobRun,omitempty"`

	// Output only. Information specific to a createChildRollout `JobRun`.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.create_child_rollout_job_run
	CreateChildRolloutJobRun *CreateChildRolloutJobRun `json:"createChildRolloutJobRun,omitempty"`

	// Output only. Information specific to an advanceChildRollout `JobRun`
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.advance_child_rollout_job_run
	AdvanceChildRolloutJobRun *AdvanceChildRolloutJobRun `json:"advanceChildRolloutJobRun,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.JobRun.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PostdeployJobRun
type PostdeployJobRunObservedState struct {
	// Output only. The resource name of the Cloud Build `Build` object that is
	//  used to execute the custom actions associated with the postdeploy Job.
	//  Format is `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.PostdeployJobRun.build
	Build *string `json:"build,omitempty"`

	// Output only. The reason the postdeploy failed. This will always be
	//  unspecified while the postdeploy is in progress or if it succeeded.
	// +kcc:proto:field=google.cloud.deploy.v1.PostdeployJobRun.failure_cause
	FailureCause *string `json:"failureCause,omitempty"`

	// Output only. Additional information about the postdeploy failure, if
	//  available.
	// +kcc:proto:field=google.cloud.deploy.v1.PostdeployJobRun.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PredeployJobRun
type PredeployJobRunObservedState struct {
	// Output only. The resource name of the Cloud Build `Build` object that is
	//  used to execute the custom actions associated with the predeploy Job.
	//  Format is `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.PredeployJobRun.build
	Build *string `json:"build,omitempty"`

	// Output only. The reason the predeploy failed. This will always be
	//  unspecified while the predeploy is in progress or if it succeeded.
	// +kcc:proto:field=google.cloud.deploy.v1.PredeployJobRun.failure_cause
	FailureCause *string `json:"failureCause,omitempty"`

	// Output only. Additional information about the predeploy failure, if
	//  available.
	// +kcc:proto:field=google.cloud.deploy.v1.PredeployJobRun.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.VerifyJobRun
type VerifyJobRunObservedState struct {
	// Output only. The resource name of the Cloud Build `Build` object that is
	//  used to verify. Format is
	//  `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.VerifyJobRun.build
	Build *string `json:"build,omitempty"`

	// Output only. URI of a directory containing the verify artifacts. This
	//  contains the Skaffold event log.
	// +kcc:proto:field=google.cloud.deploy.v1.VerifyJobRun.artifact_uri
	ArtifactURI *string `json:"artifactURI,omitempty"`

	// Output only. File path of the Skaffold event log relative to the artifact
	//  URI.
	// +kcc:proto:field=google.cloud.deploy.v1.VerifyJobRun.event_log_path
	EventLogPath *string `json:"eventLogPath,omitempty"`

	// Output only. The reason the verify failed. This will always be unspecified
	//  while the verify is in progress or if it succeeded.
	// +kcc:proto:field=google.cloud.deploy.v1.VerifyJobRun.failure_cause
	FailureCause *string `json:"failureCause,omitempty"`

	// Output only. Additional information about the verify failure, if available.
	// +kcc:proto:field=google.cloud.deploy.v1.VerifyJobRun.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`
}
