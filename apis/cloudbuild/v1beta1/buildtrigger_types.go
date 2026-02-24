// Copyright 2025 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildTriggerGVK = GroupVersion.WithKind("CloudBuildTrigger")

// CloudBuildTriggerSpec defines the desired state of CloudBuildTrigger
// +kcc:spec:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CloudBuildTrigger name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Human-readable description of this trigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.description
	Description *string `json:"description,omitempty"`

	// User-assigned name of the trigger. Must be unique within the project.
	//  Trigger names must meet the following requirements:
	//
	//  + They must contain only alphanumeric characters and dashes.
	//  + They can be 1-64 characters long.
	//  + They must begin and end with an alphanumeric character.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.name
	Name *string `json:"name,omitempty"`

	// Tags for annotation of a `BuildTrigger`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.tags
	Tags []string `json:"tags,omitempty"`

	// Template describing the types of source changes to trigger a build.
	//
	//  Branch and tag names in trigger templates are interpreted as regular
	//  expressions. Any branch or tag change that matches that regular expression
	//  will trigger a build.
	//
	//  Mutually exclusive with `github`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.trigger_template
	TriggerTemplate *RepoSource `json:"triggerTemplate,omitempty"`

	// GitHubEventsConfig describes the configuration of a trigger that creates
	//  a build whenever a GitHub event is received.
	//
	//  Mutually exclusive with `trigger_template`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.github
	Github *GitHubEventsConfig `json:"github,omitempty"`

	// PubsubConfig describes the configuration of a trigger that
	//  creates a build whenever a Pub/Sub message is published.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.pubsub_config
	PubsubConfig *PubsubConfig `json:"pubsubConfig,omitempty"`

	// WebhookConfig describes the configuration of a trigger that
	//  creates a build whenever a webhook is sent to a trigger's webhook URL.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.webhook_config
	WebhookConfig *WebhookConfig `json:"webhookConfig,omitempty"`

	// Autodetect build configuration.  The following precedence is used (case
	//  insensitive):
	//
	//  1. cloudbuild.yaml
	//  2. cloudbuild.yml
	//  3. cloudbuild.json
	//  4. Dockerfile
	//
	//  Currently only available for GitHub App Triggers.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.autodetect
	Autodetect *bool `json:"autodetect,omitempty"`

	// Contents of the build template.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.build
	Build *Build `json:"build,omitempty"`

	// Path, from the source root, to the build configuration file
	//  (i.e. cloudbuild.yaml).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.filename
	Filename *string `json:"filename,omitempty"`

	// The file source describing the local or remote Build template.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.git_file_source
	GitFileSource *GitFileSource `json:"gitFileSource,omitempty"`

	// If true, the trigger will never automatically execute a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Substitutions for Build resource. The keys must match the following
	//  regular expression: `^_[A-Z0-9_]+$`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// ignored_files and included_files are file glob matches using
	//  https://golang.org/pkg/path/filepath/#Match extended with support for "**".
	//
	//  If ignored_files and changed files are both empty, then they are
	//  not used to determine whether or not to trigger a build.
	//
	//  If ignored_files is not empty, then we ignore any files that match
	//  any of the ignored_file globs. If the change has no files that are
	//  outside of the ignored_files globs, then we do not trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.ignored_files
	IgnoredFiles []string `json:"ignoredFiles,omitempty"`

	// If any of the files altered in the commit pass the ignored_files
	//  filter and included_files is empty, then as far as this filter is
	//  concerned, we should trigger the build.
	//
	//  If any of the files altered in the commit pass the ignored_files
	//  filter and included_files is not empty, then we make sure that at
	//  least one of those files matches a included_files glob. If not,
	//  then we do not trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.included_files
	IncludedFiles []string `json:"includedFiles,omitempty"`

	// Optional. A Common Expression Language string.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.filter
	Filter *string `json:"filter,omitempty"`

	// The repo and ref of the repository from which to build. This field
	//  is used only for those triggers that do not respond to SCM events.
	//  Triggers that respond to such events build source at whatever commit
	//  caused the event.
	//  This field is currently only used by Webhook, Pub/Sub, Manual, and Cron
	//  triggers.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.source_to_build
	SourceToBuild *GitRepoSource `json:"sourceToBuild,omitempty"`

	// The service account used for all user-controlled operations including
	//  UpdateBuildTrigger, RunBuildTrigger, CreateBuild, and CancelBuild.
	//  If no service account is set and the legacy Cloud Build service account
	//  (`[PROJECT_NUM]@cloudbuild.gserviceaccount.com`) is the default for the
	//  project then it will be used instead.
	//  Format: `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The configuration of a trigger that creates a build whenever an event from
	//  Repo API is received.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.repository_event_config
	RepositoryEventConfig *RepositoryEventConfig `json:"repositoryEventConfig,omitempty"`
}

// CloudBuildTriggerStatus defines the config connector machine state of CloudBuildTrigger
type CloudBuildTriggerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudBuildTrigger resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudBuildTriggerObservedState `json:"observedState,omitempty"`
}

// CloudBuildTriggerObservedState is the state of the CloudBuildTrigger resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerObservedState struct {
	// Output only. Unique identifier of the trigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.id
	ID *string `json:"id,omitempty"`

	// PubsubConfig describes the configuration of a trigger that
	//  creates a build whenever a Pub/Sub message is published.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.pubsub_config
	PubsubConfig *PubsubConfigObservedState `json:"pubsubConfig,omitempty"`

	// Contents of the build template.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.build
	Build *BuildObservedState `json:"build,omitempty"`

	// Output only. Time when the trigger was created.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The configuration of a trigger that creates a build whenever an event from
	//  Repo API is received.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.repository_event_config
	RepositoryEventConfig *RepositoryEventConfigObservedState `json:"repositoryEventConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildtrigger;gcpcloudbuildtriggers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
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
