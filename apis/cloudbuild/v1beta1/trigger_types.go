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

package v1beta1

import (
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildTriggerGVK = GroupVersion.WithKind("CloudBuildTrigger")

// CloudBuildTriggerSpec defines the desired state of CloudBuildTrigger
// +kcc:spec:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// Immutable. The location of this resource.
	// +required
	Location string `json:"location"`

	// Human-readable description of the trigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.description
	Description *string `json:"description,omitempty"`

	// Tags for annotation of a `BuildTrigger`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.tags
	Tags []string `json:"tags,omitempty"`

	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// Mutually exclusive with `github`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.trigger_template
	TriggerTemplate *RepoSource `json:"triggerTemplate,omitempty"`

	// GitHubEventsConfig describes the configuration of a trigger that creates
	// a build whenever a GitHub event is received.
	// Mutually exclusive with `trigger_template`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.github
	Github *GitHubEventsConfig `json:"github,omitempty"`

	// PubsubConfig describes the configuration of a subscription that
	// creates a build whenever a Pub/Sub message is published.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.pubsub_config
	PubsubConfig *PubsubConfig `json:"pubsubConfig,omitempty"`

	// WebhookConfig describes the configuration of a trigger that
	// creates a build whenever a webhook is sent to a trigger's webhook URL.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.webhook_config
	WebhookConfig *WebhookConfig `json:"webhookConfig,omitempty"`

	// The git file source definition for the trigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.git_file_source
	GitFileSource *GitFileSource `json:"gitFileSource,omitempty"`

	// The repo source to build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.source_to_build
	SourceToBuild *GitRepoSource `json:"sourceToBuild,omitempty"`

	// Contents of the build template.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.build
	Build *Build `json:"build,omitempty"`

	// Path, from the source root, to the build configuration file
	// (i.e. cloudbuild.yaml).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.filename
	Filename *string `json:"filename,omitempty"`

	// Substitutions for Build resource. The keys must match the following
	// regular expression: `^_[A-Z0-9_]+$`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// ignored_files and included_files are file glob matches using
	// https://golang.org/pkg/path/filepath/#Match extended with support for "**".
	//
	// If ignored_files and changed files are both empty, then they are
	// not used to determine whether or not to trigger a build.
	//
	// If ignored_files is not empty, then we ignore any files that match
	// any of the ignored_file globs. If the change has no files that are
	// outside of the ignored_files globs, then we do not trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.ignored_files
	IgnoredFiles []string `json:"ignoredFiles,omitempty"`

	// If any of the files altered in the commit pass the ignored_files
	// filter and included_files is empty, then as far as this filter is
	// concerned, we should trigger the build.
	//
	// If any of the files altered in the commit pass the ignored_files
	// filter and included_files is not empty, then we make sure that at
	// least one of those files matches a included_files glob. If not,
	// then we do not trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.included_files
	IncludedFiles []string `json:"includedFiles,omitempty"`

	// Optional. A Common Expression Language string.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.filter
	Filter *string `json:"filter,omitempty"`

	// The IAM service account email to use as the build service account.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// If true, the trigger will never automatically execute a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// The configuration of a trigger that creates a build whenever an event from
	// Repo API is received.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.repository_event_config
	RepositoryEventConfig *RepositoryEventConfig `json:"repositoryEventConfig,omitempty"`

	// Configuration for manual approval to start a build invocation of this BuildTrigger.
	// Builds created by this trigger will require approval before they execute.
	// Any user with a Cloud Build Approver role for the project can approve a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.approval_config
	ApprovalConfig *ApprovalConfig `json:"approvalConfig,omitempty"`

	// Build logs will be sent back to GitHub as part of the checkrun
	// result.  Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or
	// INCLUDE_BUILD_LOGS_WITH_STATUS Possible values: ["INCLUDE_BUILD_LOGS_UNSPECIFIED", "INCLUDE_BUILD_LOGS_WITH_STATUS"].
	IncludeBuildLogs *string `json:"includeBuildLogs,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.RepoSource
type RepoSource struct {
	// ID of the project that owns the Cloud Source Repository. If omitted, the
	//  project ID requesting the build is assumed.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Name of the Cloud Source Repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.repo_name
	RepoName *string `json:"repoName,omitempty"`

	// The Cloud Source Repository to build. If omitted, the repo with
	// name "default" is assumed.
	RepoRef *SourceRepoRepositoryRef `json:"repoRef,omitempty"`

	// Regex matching branches to build.
	//
	//  The syntax of the regular expressions accepted is the syntax accepted by
	//  RE2 and described at https://github.com/google/re2/wiki/Syntax
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.branch_name
	BranchName *string `json:"branchName,omitempty"`

	// Regex matching tags to build.
	//
	//  The syntax of the regular expressions accepted is the syntax accepted by
	//  RE2 and described at https://github.com/google/re2/wiki/Syntax
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.tag_name
	TagName *string `json:"tagName,omitempty"`

	// Explicit commit SHA to build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	//
	//  This must be a relative path. If a step's `dir` is specified and is an
	//  absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.dir
	Dir *string `json:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision
	//  regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Substitutions to use in a triggered build. Should only be used with
	//  RunBuildTrigger
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`
}

type SourceRepoRepositoryRef struct {
	/* The `name` field of a `SourceRepoRepository` resource. */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.WebhookConfig
type WebhookConfig struct {
	// Required. Resource name for the secret required as a URL parameter.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.WebhookConfig.secret
	Secret *string `json:"secret,omitempty"`

	// The secret required
	SecretRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretRef,omitempty"`

	// Potential issues with the underlying Pub/Sub subscription configuration.
	//  Only populated on get requests.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.WebhookConfig.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PubsubConfig
type PubsubConfig struct {
	// Output only. Name of the subscription.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.subscription
	Subscription *string `json:"subscription,omitempty"`

	// The name of the topic from which this subscription is receiving messages.
	//  Format is `projects/{project}/topics/{topic}`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.topic
	Topic *string `json:"topic,omitempty"`

	// The name of the topic from which this subscription
	// is receiving messages.
	TopicRef *pubsubv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`

	// Service account that will make the push request.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Potential issues with the underlying Pub/Sub subscription configuration.
	//  Only populated on get requests.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitFileSource
type GitFileSource struct {
	// The URI of the repo. Either uri or repository can be specified. If
	//  unspecified, the repo from which the trigger invocation originated is
	//  assumed to be the repo from which to read the specified path.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.uri
	URI *string `json:"uri,omitempty"`

	// The fully qualified resource name of the Repos API repository. Either URI or
	//  repository can be specified. If unspecified, the repo from which the
	//  trigger invocation originated is assumed to be the repo from which to read
	//  the specified path.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.repository
	Repository *string `json:"repository,omitempty"`

	RepositoryRef *CloudBuildV2RepositoryRef `json:"repositoryRef,omitempty"`

	// The branch, tag, arbitrary ref, or SHA version of the repo to use when
	//  resolving the filename (optional). This field respects the same
	//  syntax/resolution as described here:
	//  https://git-scm.com/docs/gitrevisions If unspecified, the revision from
	//  which the trigger invocation originated is assumed to be the revision from
	//  which to read the specified path.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.revision
	Revision *string `json:"revision,omitempty"`

	// See RepoType above.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.repo_type
	RepoType *string `json:"repoType,omitempty"`

	// The full resource name of the github enterprise config.
	//  Format: `projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}`.
	//  `projects/{project}/githubEnterpriseConfigs/{id}`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.github_enterprise_config
	GithubEnterpriseConfig *string `json:"githubEnterpriseConfig,omitempty"`
}

type CloudBuildV2RepositoryRef struct {
	/* The `name` field of a `CloudBuildV2Repository` resource. */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildTriggerStatus defines the config connector machine state of CloudBuildTrigger
type CloudBuildTriggerStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudBuildTriggerObservedState `json:"observedState,omitempty"`

	// The unique identifier for the trigger.
	TriggerId *string `json:"triggerId,omitempty"`
}

// CloudBuildTriggerObservedState is the state of the CloudBuildTrigger resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerObservedState struct {
	// Output only. Time when the trigger was created.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.create_time
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildtrigger;gcpcloudbuildtriggers
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=beta"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudBuildTrigger is the Schema for the CloudBuildTrigger API
// +k8s:openapi-gen=true
type CloudBuildTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudBuildTriggerSpec   `json:"spec,omitempty"`
	Status CloudBuildTriggerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBuildTriggerList contains a list of CloudBuildTrigger
type CloudBuildTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildTrigger{}, &CloudBuildTriggerList{})
}
