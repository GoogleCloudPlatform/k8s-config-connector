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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildTriggerGVK = GroupVersion.WithKind("CloudBuildTrigger")

// +kcc:spec:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerSpec_ApprovalConfig struct {
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`
}
type CloudBuildTriggerSpec_BitbucketServerTriggerConfig_BitbucketServerConfigResourceRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_BitbucketServerTriggerConfig_PullRequest struct {
	Branch         *string `json:"branch,omitempty"`
	CommentControl *string `json:"commentControl,omitempty"`
	InvertRegex    *bool   `json:"invertRegex,omitempty"`
}
type CloudBuildTriggerSpec_BitbucketServerTriggerConfig_Push struct {
	Branch      *string `json:"branch,omitempty"`
	InvertRegex *bool   `json:"invertRegex,omitempty"`
	Tag         *string `json:"tag,omitempty"`
}
type CloudBuildTriggerSpec_BitbucketServerTriggerConfig struct {
	BitbucketServerConfigResourceRef *CloudBuildTriggerSpec_BitbucketServerTriggerConfig_BitbucketServerConfigResourceRef `json:"bitbucketServerConfigResourceRef,omitempty"`
	ProjectKey                       *string                                                                              `json:"projectKey,omitempty"`
	PullRequest                      *CloudBuildTriggerSpec_BitbucketServerTriggerConfig_PullRequest                      `json:"pullRequest,omitempty"`
	Push                             *CloudBuildTriggerSpec_BitbucketServerTriggerConfig_Push                             `json:"push,omitempty"`
	RepoSlug                         *string                                                                              `json:"repoSlug,omitempty"`
}
type CloudBuildTriggerSpec_Build_Artifacts_Objects_TimingItem struct {
	EndTime   *string `json:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
}
type CloudBuildTriggerSpec_Build_Artifacts_Objects struct {
	Location *string                                                    `json:"location,omitempty"`
	Paths    []string                                                   `json:"paths,omitempty"`
	Timing   []CloudBuildTriggerSpec_Build_Artifacts_Objects_TimingItem `json:"timing,omitempty"`
}
type CloudBuildTriggerSpec_Build_Artifacts struct {
	Images  []string                                       `json:"images,omitempty"`
	Objects *CloudBuildTriggerSpec_Build_Artifacts_Objects `json:"objects,omitempty"`
}
type CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem_VersionRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem struct {
	Env        *string                                                                    `json:"env,omitempty"`
	VersionRef *CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem_VersionRef `json:"versionRef,omitempty"`
}
type CloudBuildTriggerSpec_Build_AvailableSecrets struct {
	SecretManager []CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem `json:"secretManager,omitempty"`
}
type CloudBuildTriggerSpec_Build_LogsBucketRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_Build_Options_VolumesItem struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}
type CloudBuildTriggerSpec_Build_Options struct {
	DiskSizeGb            *int                                              `json:"diskSizeGb,omitempty"`
	DynamicSubstitutions  *bool                                             `json:"dynamicSubstitutions,omitempty"`
	Env                   []string                                          `json:"env,omitempty"`
	LogStreamingOption    *string                                           `json:"logStreamingOption,omitempty"`
	Logging               *string                                           `json:"logging,omitempty"`
	MachineType           *string                                           `json:"machineType,omitempty"`
	RequestedVerifyOption *string                                           `json:"requestedVerifyOption,omitempty"`
	SecretEnv             []string                                          `json:"secretEnv,omitempty"`
	SourceProvenanceHash  []string                                          `json:"sourceProvenanceHash,omitempty"`
	SubstitutionOption    *string                                           `json:"substitutionOption,omitempty"`
	Volumes               []CloudBuildTriggerSpec_Build_Options_VolumesItem `json:"volumes,omitempty"`
	// Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}
	//
	// This field is experimental.
	WorkerPool *string `json:"workerPool,omitempty"`
}
type CloudBuildTriggerSpec_Build_SecretItem_KmsKeyRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_Build_SecretItem struct {
	KmsKeyRef *CloudBuildTriggerSpec_Build_SecretItem_KmsKeyRef `json:"kmsKeyRef,omitempty"`
	SecretEnv map[string]string                                 `json:"secretEnv,omitempty"`
}
type CloudBuildTriggerSpec_Build_Source_RepoSource_RepoRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_Build_Source_RepoSource struct {
	BranchName    *string                                                `json:"branchName,omitempty"`
	CommitSha     *string                                                `json:"commitSha,omitempty"`
	Dir           *string                                                `json:"dir,omitempty"`
	InvertRegex   *bool                                                  `json:"invertRegex,omitempty"`
	ProjectId     *string                                                `json:"projectId,omitempty"`
	RepoRef       *CloudBuildTriggerSpec_Build_Source_RepoSource_RepoRef `json:"repoRef,omitempty"`
	Substitutions map[string]string                                      `json:"substitutions,omitempty"`
	TagName       *string                                                `json:"tagName,omitempty"`
}
type CloudBuildTriggerSpec_Build_Source_StorageSource_BucketRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_Build_Source_StorageSource struct {
	BucketRef  *CloudBuildTriggerSpec_Build_Source_StorageSource_BucketRef `json:"bucketRef,omitempty"`
	Generation *string                                                     `json:"generation,omitempty"`
	Object     *string                                                     `json:"object,omitempty"`
}
type CloudBuildTriggerSpec_Build_Source struct {
	RepoSource    *CloudBuildTriggerSpec_Build_Source_RepoSource    `json:"repoSource,omitempty"`
	StorageSource *CloudBuildTriggerSpec_Build_Source_StorageSource `json:"storageSource,omitempty"`
}
type CloudBuildTriggerSpec_Build_StepItem_VolumesItem struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}
type CloudBuildTriggerSpec_Build_StepItem struct {
	AllowExitCodes []int                                              `json:"allowExitCodes,omitempty"`
	AllowFailure   *bool                                              `json:"allowFailure,omitempty"`
	Args           []string                                           `json:"args,omitempty"`
	Dir            *string                                            `json:"dir,omitempty"`
	Entrypoint     *string                                            `json:"entrypoint,omitempty"`
	Env            []string                                           `json:"env,omitempty"`
	Id             *string                                            `json:"id,omitempty"`
	Name           *string                                            `json:"name,omitempty"`
	Script         *string                                            `json:"script,omitempty"`
	SecretEnv      []string                                           `json:"secretEnv,omitempty"`
	Timeout        *string                                            `json:"timeout,omitempty"`
	Timing         *string                                            `json:"timing,omitempty"`
	Volumes        []CloudBuildTriggerSpec_Build_StepItem_VolumesItem `json:"volumes,omitempty"`
	WaitFor        []string                                           `json:"waitFor,omitempty"`
}
type CloudBuildTriggerSpec_Build struct {
	Artifacts        *CloudBuildTriggerSpec_Build_Artifacts        `json:"artifacts,omitempty"`
	AvailableSecrets *CloudBuildTriggerSpec_Build_AvailableSecrets `json:"availableSecrets,omitempty"`
	Images           []string                                      `json:"images,omitempty"`
	LogsBucketRef    *CloudBuildTriggerSpec_Build_LogsBucketRef    `json:"logsBucketRef,omitempty"`
	Options          *CloudBuildTriggerSpec_Build_Options          `json:"options,omitempty"`
	QueueTtl         *string                                       `json:"queueTtl,omitempty"`
	Secret           []CloudBuildTriggerSpec_Build_SecretItem      `json:"secret,omitempty"`
	Source           *CloudBuildTriggerSpec_Build_Source           `json:"source,omitempty"`
	Step             []CloudBuildTriggerSpec_Build_StepItem        `json:"step,omitempty"`
	Substitutions    map[string]string                             `json:"substitutions,omitempty"`
	Tags             []string                                      `json:"tags,omitempty"`
	Timeout          *string                                       `json:"timeout,omitempty"`
}
type CloudBuildTriggerSpec_GitFileSource_BitbucketServerConfigRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_GitFileSource_GithubEnterpriseConfigRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_GitFileSource_RepositoryRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_GitFileSource struct {
	BitbucketServerConfigRef  *CloudBuildTriggerSpec_GitFileSource_BitbucketServerConfigRef  `json:"bitbucketServerConfigRef,omitempty"`
	GithubEnterpriseConfigRef *CloudBuildTriggerSpec_GitFileSource_GithubEnterpriseConfigRef `json:"githubEnterpriseConfigRef,omitempty"`
	Path                      *string                                                        `json:"path,omitempty"`
	RepoType                  *string                                                        `json:"repoType,omitempty"`
	RepositoryRef             *CloudBuildTriggerSpec_GitFileSource_RepositoryRef             `json:"repositoryRef,omitempty"`
	Revision                  *string                                                        `json:"revision,omitempty"`
	Uri                       *string                                                        `json:"uri,omitempty"`
}
type CloudBuildTriggerSpec_Github_EnterpriseConfigResourceNameRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_Github_PullRequest struct {
	Branch         *string `json:"branch,omitempty"`
	CommentControl *string `json:"commentControl,omitempty"`
	InvertRegex    *bool   `json:"invertRegex,omitempty"`
}
type CloudBuildTriggerSpec_Github_Push struct {
	Branch      *string `json:"branch,omitempty"`
	InvertRegex *bool   `json:"invertRegex,omitempty"`
	Tag         *string `json:"tag,omitempty"`
}
type CloudBuildTriggerSpec_Github struct {
	EnterpriseConfigResourceNameRef *CloudBuildTriggerSpec_Github_EnterpriseConfigResourceNameRef `json:"enterpriseConfigResourceNameRef,omitempty"`
	Name                            *string                                                       `json:"name,omitempty"`
	Owner                           *string                                                       `json:"owner,omitempty"`
	PullRequest                     *CloudBuildTriggerSpec_Github_PullRequest                     `json:"pullRequest,omitempty"`
	Push                            *CloudBuildTriggerSpec_Github_Push                            `json:"push,omitempty"`
}
type CloudBuildTriggerSpec_PubsubConfig_ServiceAccountRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_PubsubConfig_TopicRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_PubsubConfig struct {
	ServiceAccountRef *CloudBuildTriggerSpec_PubsubConfig_ServiceAccountRef `json:"serviceAccountRef,omitempty"`
	State             *string                                               `json:"state,omitempty"`
	Subscription      *string                                               `json:"subscription,omitempty"`
	TopicRef          *CloudBuildTriggerSpec_PubsubConfig_TopicRef          `json:"topicRef,omitempty"`
}
type CloudBuildTriggerSpec_RepositoryEventConfig_PullRequest struct {
	Branch         *string `json:"branch,omitempty"`
	CommentControl *string `json:"commentControl,omitempty"`
	InvertRegex    *bool   `json:"invertRegex,omitempty"`
}
type CloudBuildTriggerSpec_RepositoryEventConfig_Push struct {
	Branch      *string `json:"branch,omitempty"`
	InvertRegex *bool   `json:"invertRegex,omitempty"`
	Tag         *string `json:"tag,omitempty"`
}
type CloudBuildTriggerSpec_RepositoryEventConfig struct {
	PullRequest *CloudBuildTriggerSpec_RepositoryEventConfig_PullRequest `json:"pullRequest,omitempty"`
	Push        *CloudBuildTriggerSpec_RepositoryEventConfig_Push        `json:"push,omitempty"`
	Repository  *string                                                  `json:"repository,omitempty"`
}
type CloudBuildTriggerSpec_ServiceAccountRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_SourceToBuild_BitbucketServerConfigRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_SourceToBuild_GithubEnterpriseConfigRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_SourceToBuild_RepositoryRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_SourceToBuild struct {
	BitbucketServerConfigRef  *CloudBuildTriggerSpec_SourceToBuild_BitbucketServerConfigRef  `json:"bitbucketServerConfigRef,omitempty"`
	GithubEnterpriseConfigRef *CloudBuildTriggerSpec_SourceToBuild_GithubEnterpriseConfigRef `json:"githubEnterpriseConfigRef,omitempty"`
	Ref                       *string                                                        `json:"ref,omitempty"`
	RepoType                  *string                                                        `json:"repoType,omitempty"`
	RepositoryRef             *CloudBuildTriggerSpec_SourceToBuild_RepositoryRef             `json:"repositoryRef,omitempty"`
	Uri                       *string                                                        `json:"uri,omitempty"`
}
type CloudBuildTriggerSpec_TriggerTemplate_RepoRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_TriggerTemplate struct {
	BranchName  *string                                        `json:"branchName,omitempty"`
	CommitSha   *string                                        `json:"commitSha,omitempty"`
	Dir         *string                                        `json:"dir,omitempty"`
	InvertRegex *bool                                          `json:"invertRegex,omitempty"`
	RepoRef     *CloudBuildTriggerSpec_TriggerTemplate_RepoRef `json:"repoRef,omitempty"`
	TagName     *string                                        `json:"tagName,omitempty"`
}
type CloudBuildTriggerSpec_WebhookConfig_SecretRef struct {
	External  *string `json:"external,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}
type CloudBuildTriggerSpec_WebhookConfig struct {
	SecretRef *CloudBuildTriggerSpec_WebhookConfig_SecretRef `json:"secretRef,omitempty"`
	State     *string                                        `json:"state,omitempty"`
}
type CloudBuildTriggerSpec struct {
	ApprovalConfig               *CloudBuildTriggerSpec_ApprovalConfig               `json:"approvalConfig,omitempty"`
	BitbucketServerTriggerConfig *CloudBuildTriggerSpec_BitbucketServerTriggerConfig `json:"bitbucketServerTriggerConfig,omitempty"`
	Build                        *CloudBuildTriggerSpec_Build                        `json:"build,omitempty"`
	Description                  *string                                             `json:"description,omitempty"`
	Disabled                     *bool                                               `json:"disabled,omitempty"`
	Filename                     *string                                             `json:"filename,omitempty"`
	Filter                       *string                                             `json:"filter,omitempty"`
	GitFileSource                *CloudBuildTriggerSpec_GitFileSource                `json:"gitFileSource,omitempty"`
	Github                       *CloudBuildTriggerSpec_Github                       `json:"github,omitempty"`
	IgnoredFiles                 []string                                            `json:"ignoredFiles,omitempty"`
	IncludeBuildLogs             *string                                             `json:"includeBuildLogs,omitempty"`
	IncludedFiles                []string                                            `json:"includedFiles,omitempty"`
	Location                     *string                                             `json:"location"`
	PubsubConfig                 *CloudBuildTriggerSpec_PubsubConfig                 `json:"pubsubConfig,omitempty"`
	RepositoryEventConfig        *CloudBuildTriggerSpec_RepositoryEventConfig        `json:"repositoryEventConfig,omitempty"`
	ServiceAccountRef            *CloudBuildTriggerSpec_ServiceAccountRef            `json:"serviceAccountRef,omitempty"`
	SourceToBuild                *CloudBuildTriggerSpec_SourceToBuild                `json:"sourceToBuild,omitempty"`
	Substitutions                map[string]string                                   `json:"substitutions,omitempty"`
	Tags                         []string                                            `json:"tags,omitempty"`
	TriggerTemplate              *CloudBuildTriggerSpec_TriggerTemplate              `json:"triggerTemplate,omitempty"`
	WebhookConfig                *CloudBuildTriggerSpec_WebhookConfig                `json:"webhookConfig,omitempty"`
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

	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.create_time
	CreateTime *string `json:"createTime,omitempty"`

	TriggerId *string `json:"triggerId,omitempty"`
}

// CloudBuildTriggerObservedState is the state of the CloudBuildTrigger resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildtrigger;gcpcloudbuildtriggers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
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

// +kcc:proto=google.devtools.cloudbuild.v1.Hash
type Hash struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Hash.type
	Type *string `json:"type,omitempty"`
	// +kcc:proto:field=-
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Results
type Results struct {
	// +kcc:proto:field=-
	BuildStepImages []string `json:"buildStepImages,omitempty"`
	// +kcc:proto:field=-
	BuildStepOutputs [][]byte `json:"buildStepOutputs,omitempty"`
}
