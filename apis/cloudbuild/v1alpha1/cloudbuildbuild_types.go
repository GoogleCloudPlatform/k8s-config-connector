// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for CloudBuildBuild, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/devtools/cloudbuild/v1/cloudbuild.proto
// +kcc:source:service-docs=https://cloud.google.com/cloud-build/docs/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/cloud-build/docs/reference/rest/v1/projects.builds

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildBuildGVK = GroupVersion.WithKind("CloudBuildBuild")

// CloudBuildBuildSpec defines the desired state of CloudBuildBuild
// +kcc:spec:proto=google.devtools.cloudbuild.v1.Build
// +kcc:required-from-proto
type CloudBuildBuildSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The CloudBuildBuild name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The location of the source files to build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.source
	Source *Source `json:"source,omitempty"`

	// Required. The operations to be performed on the workspace.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.steps
	Steps []BuildStep `json:"steps,omitempty"`

	// Amount of time that this build should be allowed to run, to second
	//  granularity. If this amount of time elapses, work on the build will cease
	//  and the build status will be `TIMEOUT`.
	//
	//  `timeout` starts ticking from `startTime`.
	//
	//  Default time is 60 minutes.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.timeout
	Timeout *string `json:"timeout,omitempty"`

	// A list of images to be pushed upon the successful completion of all build
	//  steps.
	//
	//  The images are pushed using the builder service account's credentials.
	//
	//  The digests of the pushed images will be stored in the `Build` resource's
	//  results field.
	//
	//  If any of the images fail to be pushed, the build status is marked
	//  `FAILURE`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.images
	Images []string `json:"images,omitempty"`

	// TTL in queue for this build. If provided and the build is enqueued longer
	//  than this value, the build will expire and the build status will be
	//  `EXPIRED`.
	//
	//  The TTL starts ticking from create_time.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.queue_ttl
	QueueTTL *string `json:"queueTTL,omitempty"`

	// Artifacts produced by the build that should be uploaded upon
	//  successful completion of all build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.artifacts
	Artifacts *Artifacts `json:"artifacts,omitempty"`

	// Cloud Storage bucket where logs should be written (see
	//  [Bucket Name
	//  Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	//  Logs file names will be of the format `${logs_bucket}/log-${build_id}.txt`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.logs_bucket
	LogsBucket *string `json:"logsBucket,omitempty"`

	// Special options for this build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.options
	Options *BuildOptions `json:"options,omitempty"`

	// Substitutions data for `Build` resource.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Tags for annotation of a `Build`. These are not docker tags.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.tags
	Tags []string `json:"tags,omitempty"`

	// Secrets to decrypt using Cloud Key Management Service.
	//  Note: Secret Manager is the recommended technique
	//  for managing sensitive data with Cloud Build. Use `available_secrets` to
	//  configure builds to access secrets from Secret Manager. For instructions,
	//  see: https://cloud.google.com/cloud-build/docs/securing-builds/use-secrets
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.secrets
	Secrets []Secret `json:"secrets,omitempty"`

	// IAM service account whose credentials will be used at build runtime.
	//  Must be of the format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`.
	//  ACCOUNT can be email address or uniqueId of the service account.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Secrets and secret environment variables.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.available_secrets
	AvailableSecrets *Secrets `json:"availableSecrets,omitempty"`

	// Optional. Configuration for git operations.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.git_config
	GitConfig *GitConfig `json:"gitConfig,omitempty"`

	// Optional. Dependencies that the Cloud Build worker will fetch before
	//  executing user steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.dependencies
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

// CloudBuildBuildStatus defines the config connector machine state of CloudBuildBuild
type CloudBuildBuildStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudBuildBuild resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudBuildBuildObservedState `json:"observedState,omitempty"`
}

// CloudBuildBuildObservedState is the state of the CloudBuildBuild resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.cloudbuild.v1.Build
type CloudBuildBuildObservedState struct {
	// Output only. Unique identifier of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.id
	ID *string `json:"id,omitempty"`

	// Output only. ID of the project.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Output only. Status of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.status
	Status *string `json:"status,omitempty"`

	// Output only. Customer-readable message about the current status.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.status_detail
	StatusDetail *string `json:"statusDetail,omitempty"`

	// Required. The operations to be performed on the workspace.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.steps
	Steps []BuildStepObservedState `json:"steps,omitempty"`

	// Output only. Results of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.results
	Results *ResultsObservedState `json:"results,omitempty"`

	// Output only. Time at which the request to create the build was received.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which execution of the build was started.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time at which execution of the build was finished.
	//
	//  The difference between finish_time and start_time is the duration of the
	//  build's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.finish_time
	FinishTime *string `json:"finishTime,omitempty"`

	// Artifacts produced by the build that should be uploaded upon
	//  successful completion of all build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.artifacts
	Artifacts *ArtifactsObservedState `json:"artifacts,omitempty"`

	// Output only. A permanent fixed identifier for source.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.source_provenance
	SourceProvenance *SourceProvenanceObservedState `json:"sourceProvenance,omitempty"`

	// Output only. The ID of the `BuildTrigger` that triggered this build, if it
	//  was triggered automatically.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.build_trigger_id
	BuildTriggerID *string `json:"buildTriggerID,omitempty"`

	// Output only. URL to logs for this build in Google Cloud Console.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.log_url
	LogURL *string `json:"logURL,omitempty"`

	// Output only. Stores timing information for phases of the build. Valid keys
	//  are:
	//
	//  * BUILD: time to execute all build steps.
	//  * PUSH: time to push all artifacts including docker images and non docker
	//  artifacts.
	//  * FETCHSOURCE: time to fetch source.
	//  * SETUPBUILD: time to set up build.
	//
	//  If the build does not specify source or images,
	//  these keys will not be included.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.timing
	Timing map[string]TimeSpan `json:"timing,omitempty"`

	// Output only. Describes this build's approval configuration, status,
	//  and result.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.approval
	Approval *BuildApprovalObservedState `json:"approval,omitempty"`

	// Output only. Non-fatal problems encountered during the execution of the
	//  build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.warnings
	Warnings []Build_Warning `json:"warnings,omitempty"`

	// Output only. Contains information about the build when status=FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.failure_info
	FailureInfo *Build_FailureInfo `json:"failureInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildbuild;gcpcloudbuildbuilds
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudBuildBuild is the Schema for the CloudBuildBuild API
// +k8s:openapi-gen=true
type CloudBuildBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudBuildBuildSpec   `json:"spec,omitempty"`
	Status CloudBuildBuildStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBuildBuildList contains a list of CloudBuildBuild
type CloudBuildBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildBuild `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildBuild{}, &CloudBuildBuildList{})
}
