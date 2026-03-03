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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Reference types for CloudBuildTrigger.

// CloudBuildBitbucketServerConfigRef is a reference to a CloudBuildBitbucketServerConfig resource.
type CloudBuildBitbucketServerConfigRef struct {
	// Allowed value: The `name` field of a `CloudBuildBitbucketServerConfig` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildGithubEnterpriseConfigRef is a reference to a CloudBuildGithubEnterpriseConfig resource.
type CloudBuildGithubEnterpriseConfigRef struct {
	// Allowed value: The `name` field of a `CloudBuildGithubEnterpriseConfig` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildV2RepositoryRef is a reference to a CloudBuildV2Repository resource.
type CloudBuildV2RepositoryRef struct {
	// Allowed value: The `name` field of a `CloudBuildV2Repository` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildSecretManagerSecretVersionRef is a reference to a SecretManagerSecretVersion resource.
type CloudBuildSecretManagerSecretVersionRef struct {
	// Allowed value: The `name` field of a `SecretManagerSecretVersion` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildStorageBucketRef is a reference to a StorageBucket resource.
type CloudBuildStorageBucketRef struct {
	// Allowed value: The `name` field of a `StorageBucket` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildKMSCryptoKeyRef is a reference to a KMSCryptoKey resource.
type CloudBuildKMSCryptoKeyRef struct {
	// Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildSourceRepoRepositoryRef is a reference to a SourceRepoRepository resource.
type CloudBuildSourceRepoRepositoryRef struct {
	// Allowed value: The `name` field of a `SourceRepoRepository` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildIAMServiceAccountRef is a reference to an IAMServiceAccount resource.
type CloudBuildIAMServiceAccountRef struct {
	// Allowed value: The `email` field of an `IAMServiceAccount` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildPubSubTopicRef is a reference to a PubSubTopic resource.
type CloudBuildPubSubTopicRef struct {
	// Allowed value: string of the format `projects/{{project}}/topics/{{value}}`, where {{value}} is the `name` field of a `PubSubTopic` resource.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// CloudBuildTriggerApprovalConfig is configuration for manual approval to start a build invocation.
type CloudBuildTriggerApprovalConfig struct {
	// Whether or not approval is needed. If this is set on a build, it will become pending when run,
	// and will need to be explicitly approved to start.
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`
}

// CloudBuildTriggerBitbucketServerTriggerConfig describes the configuration of a trigger that creates
// a build whenever a Bitbucket Server event is received.
type CloudBuildTriggerBitbucketServerTriggerConfig struct {
	// Only `external` field is supported to configure the reference.
	// The full resource name of the bitbucket server config.
	BitbucketServerConfigResourceRef CloudBuildBitbucketServerConfigRef `json:"bitbucketServerConfigResourceRef"`

	// Key of the project that the repo is in.
	ProjectKey string `json:"projectKey"`

	// Filter to match changes in pull requests.
	PullRequest *CloudBuildTriggerBitbucketServerPullRequestFilter `json:"pullRequest,omitempty"`

	// Filter to match changes in refs like branches, tags.
	Push *CloudBuildTriggerBitbucketServerPushFilter `json:"push,omitempty"`

	// Slug of the repository.
	RepoSlug string `json:"repoSlug"`
}

// CloudBuildTriggerBitbucketServerPullRequestFilter is a filter to match changes in pull requests.
type CloudBuildTriggerBitbucketServerPullRequestFilter struct {
	// Regex of branches to match.
	Branch string `json:"branch"`

	// Configure builds to run whether a repository owner or collaborator need to comment /gcbrun.
	CommentControl *string `json:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// CloudBuildTriggerBitbucketServerPushFilter is a filter to match changes in refs like branches, tags.
type CloudBuildTriggerBitbucketServerPushFilter struct {
	// Regex of branches to match. Specify only one of branch or tag.
	Branch *string `json:"branch,omitempty"`

	// When true, only trigger a build if the revision regex does NOT match the gitRef regex.
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Regex of tags to match. Specify only one of branch or tag.
	Tag *string `json:"tag,omitempty"`
}

// CloudBuildTriggerBuild is the contents of the build template.
type CloudBuildTriggerBuild struct {
	// Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
	Artifacts *CloudBuildTriggerBuildArtifacts `json:"artifacts,omitempty"`

	// Secrets and secret environment variables.
	AvailableSecrets *CloudBuildTriggerBuildAvailableSecrets `json:"availableSecrets,omitempty"`

	// A list of images to be pushed upon the successful completion of all build steps.
	Images []string `json:"images,omitempty"`

	// Google Cloud Storage bucket where logs should be written.
	LogsBucketRef *CloudBuildStorageBucketRef `json:"logsBucketRef,omitempty"`

	// Special options for this build.
	Options *CloudBuildTriggerBuildOptions `json:"options,omitempty"`

	// TTL in queue for this build.
	QueueTtl *string `json:"queueTtl,omitempty"`

	// Secrets to decrypt using Cloud Key Management Service.
	Secret []CloudBuildTriggerBuildSecret `json:"secret,omitempty"`

	// The location of the source files to build.
	Source *CloudBuildTriggerBuildSource `json:"source,omitempty"`

	// The operations to be performed on the workspace.
	Step []CloudBuildTriggerBuildStep `json:"step"`

	// Substitutions data for Build resource.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `json:"tags,omitempty"`

	// Amount of time that this build should be allowed to run, to second granularity.
	Timeout *string `json:"timeout,omitempty"`
}

// CloudBuildTriggerBuildArtifacts contains artifacts produced by the build.
type CloudBuildTriggerBuildArtifacts struct {
	// A list of images to be pushed upon the successful completion of all build steps.
	Images []string `json:"images,omitempty"`

	// A list of objects to be uploaded to Cloud Storage upon successful completion of all build steps.
	Objects *CloudBuildTriggerBuildArtifactsObjects `json:"objects,omitempty"`
}

// CloudBuildTriggerBuildArtifactsObjects is a list of objects to be uploaded to Cloud Storage.
type CloudBuildTriggerBuildArtifactsObjects struct {
	// Cloud Storage bucket and optional object path.
	Location *string `json:"location,omitempty"`

	// Path globs used to match files in the build's workspace.
	Paths []string `json:"paths,omitempty"`

	// Output only. Stores timing information for pushing all artifact objects.
	Timing []CloudBuildTriggerBuildArtifactsObjectsTiming `json:"timing,omitempty"`
}

// CloudBuildTriggerBuildArtifactsObjectsTiming stores timing information for pushing artifact objects.
type CloudBuildTriggerBuildArtifactsObjectsTiming struct {
	// End of time span.
	EndTime *string `json:"endTime,omitempty"`

	// Start of time span.
	StartTime *string `json:"startTime,omitempty"`
}

// CloudBuildTriggerBuildAvailableSecrets contains secrets and secret environment variables.
type CloudBuildTriggerBuildAvailableSecrets struct {
	// Pairs a secret environment variable with a SecretVersion in Secret Manager.
	SecretManager []CloudBuildTriggerBuildAvailableSecretsSecretManager `json:"secretManager"`
}

// CloudBuildTriggerBuildAvailableSecretsSecretManager pairs a secret environment variable with a SecretVersion.
type CloudBuildTriggerBuildAvailableSecretsSecretManager struct {
	// Environment variable name to associate with the secret.
	Env string `json:"env"`

	// The SecretVersion in Secret Manager.
	VersionRef CloudBuildSecretManagerSecretVersionRef `json:"versionRef"`
}

// CloudBuildTriggerBuildOptions contains special options for the build.
type CloudBuildTriggerBuildOptions struct {
	// Requested disk size for the VM that runs the build.
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

	// Option to specify whether or not to apply bash style string operations to the substitutions.
	DynamicSubstitutions *bool `json:"dynamicSubstitutions,omitempty"`

	// A list of global environment variable definitions.
	Env []string `json:"env,omitempty"`

	// Option to define build log streaming behavior to Google Cloud Storage.
	LogStreamingOption *string `json:"logStreamingOption,omitempty"`

	// Option to specify the logging mode.
	Logging *string `json:"logging,omitempty"`

	// Compute Engine machine type on which to run the build.
	MachineType *string `json:"machineType,omitempty"`

	// Requested verifiability options.
	RequestedVerifyOption *string `json:"requestedVerifyOption,omitempty"`

	// A list of global environment variables, which are encrypted using a Cloud Key Management Service crypto key.
	SecretEnv []string `json:"secretEnv,omitempty"`

	// Requested hash for SourceProvenance.
	SourceProvenanceHash []string `json:"sourceProvenanceHash,omitempty"`

	// Option to specify behavior when there is an error in the substitution checks.
	SubstitutionOption *string `json:"substitutionOption,omitempty"`

	// Global list of volumes to mount for ALL build steps.
	Volumes []CloudBuildTriggerBuildOptionsVolume `json:"volumes,omitempty"`

	// Option to specify a WorkerPool for the build.
	WorkerPool *string `json:"workerPool,omitempty"`
}

// CloudBuildTriggerBuildOptionsVolume is a volume to mount for build steps.
type CloudBuildTriggerBuildOptionsVolume struct {
	// Name of the volume to mount.
	Name *string `json:"name,omitempty"`

	// Path at which to mount the volume.
	Path *string `json:"path,omitempty"`
}

// CloudBuildTriggerBuildSecret contains a secret to decrypt using Cloud Key Management Service.
type CloudBuildTriggerBuildSecret struct {
	// KMS crypto key to use to decrypt these envs.
	KmsKeyRef CloudBuildKMSCryptoKeyRef `json:"kmsKeyRef"`

	// Map of environment variable name to its encrypted value.
	SecretEnv map[string]string `json:"secretEnv,omitempty"`
}

// CloudBuildTriggerBuildSource is the location of the source files to build.
type CloudBuildTriggerBuildSource struct {
	// Location of the source in a Google Cloud Source Repository.
	RepoSource *CloudBuildTriggerBuildSourceRepoSource `json:"repoSource,omitempty"`

	// Location of the source in an archive file in Google Cloud Storage.
	StorageSource *CloudBuildTriggerBuildSourceStorageSource `json:"storageSource,omitempty"`
}

// CloudBuildTriggerBuildSourceRepoSource is the location of the source in a Google Cloud Source Repository.
type CloudBuildTriggerBuildSourceRepoSource struct {
	// Regex matching branches to build.
	BranchName *string `json:"branchName,omitempty"`

	// Explicit commit SHA to build.
	CommitSha *string `json:"commitSha,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	Dir *string `json:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision regex.
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// ID of the project that owns the Cloud Source Repository.
	ProjectId *string `json:"projectId,omitempty"`

	// The desired Cloud Source Repository.
	RepoRef CloudBuildSourceRepoRepositoryRef `json:"repoRef"`

	// Substitutions to use in a triggered build.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Regex matching tags to build.
	TagName *string `json:"tagName,omitempty"`
}

// CloudBuildTriggerBuildSourceStorageSource is the location of the source in an archive file in Google Cloud Storage.
type CloudBuildTriggerBuildSourceStorageSource struct {
	// Google Cloud Storage bucket containing the source.
	BucketRef CloudBuildStorageBucketRef `json:"bucketRef"`

	// Google Cloud Storage generation for the object.
	Generation *string `json:"generation,omitempty"`

	// Google Cloud Storage object containing the source.
	Object string `json:"object"`
}

// CloudBuildTriggerBuildStep is an operation to be performed on the workspace.
type CloudBuildTriggerBuildStep struct {
	// Allow this build step to fail without failing the entire build if and only if the exit code is one of the specified codes.
	AllowExitCodes []int `json:"allowExitCodes,omitempty"`

	// Allow this build step to fail without failing the entire build.
	AllowFailure *bool `json:"allowFailure,omitempty"`

	// A list of arguments that will be presented to the step when it is started.
	Args []string `json:"args,omitempty"`

	// Working directory to use when running this step's container.
	Dir *string `json:"dir,omitempty"`

	// Entrypoint to be used instead of the build step image's default entrypoint.
	Entrypoint *string `json:"entrypoint,omitempty"`

	// A list of environment variable definitions to be used when running a step.
	Env []string `json:"env,omitempty"`

	// Unique identifier for this build step, used in 'wait_for' to reference this build step as a dependency.
	Id *string `json:"id,omitempty"`

	// The name of the container image that will run this particular build step.
	Name string `json:"name"`

	// A shell script to be executed in the step.
	Script *string `json:"script,omitempty"`

	// A list of environment variables which are encrypted using a Cloud Key Management Service crypto key.
	SecretEnv []string `json:"secretEnv,omitempty"`

	// Time limit for executing this build step.
	Timeout *string `json:"timeout,omitempty"`

	// Output only. Stores timing information for executing this build step.
	Timing *string `json:"timing,omitempty"`

	// List of volumes to mount into the build step.
	Volumes []CloudBuildTriggerBuildStepVolume `json:"volumes,omitempty"`

	// The ID(s) of the step(s) that this build step depends on.
	WaitFor []string `json:"waitFor,omitempty"`
}

// CloudBuildTriggerBuildStepVolume is a volume to mount into a build step.
type CloudBuildTriggerBuildStepVolume struct {
	// Name of the volume to mount.
	Name string `json:"name"`

	// Path at which to mount the volume.
	Path string `json:"path"`
}

// CloudBuildTriggerGitFileSource describes the local or remote Build template.
type CloudBuildTriggerGitFileSource struct {
	// Only `external` field is supported to configure the reference.
	// The full resource name of the bitbucket server config.
	BitbucketServerConfigRef *CloudBuildBitbucketServerConfigRef `json:"bitbucketServerConfigRef,omitempty"`

	// Only `external` field is supported to configure the reference.
	// The full resource name of the github enterprise config.
	GithubEnterpriseConfigRef *CloudBuildGithubEnterpriseConfigRef `json:"githubEnterpriseConfigRef,omitempty"`

	// The path of the file, with the repo root as the root of the path.
	Path string `json:"path"`

	// The type of the repo, since it may not be explicit from the repo field.
	RepoType string `json:"repoType"`

	// Only `external` field is supported to configure the reference.
	// The fully qualified resource name of the Repo API repository.
	RepositoryRef *CloudBuildV2RepositoryRef `json:"repositoryRef,omitempty"`

	// The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the filename.
	Revision *string `json:"revision,omitempty"`

	// The URI of the repo.
	Uri *string `json:"uri,omitempty"`
}

// CloudBuildTriggerGithub describes the configuration of a trigger that creates a build whenever a GitHub event is received.
type CloudBuildTriggerGithub struct {
	// Only `external` field is supported to configure the reference.
	// The full resource name of the github enterprise config.
	EnterpriseConfigResourceNameRef *CloudBuildGithubEnterpriseConfigRef `json:"enterpriseConfigResourceNameRef,omitempty"`

	// Name of the repository.
	Name *string `json:"name,omitempty"`

	// Owner of the repository.
	Owner *string `json:"owner,omitempty"`

	// filter to match changes in pull requests.
	PullRequest *CloudBuildTriggerGithubPullRequest `json:"pullRequest,omitempty"`

	// filter to match changes in refs, like branches or tags.
	Push *CloudBuildTriggerGithubPush `json:"push,omitempty"`
}

// CloudBuildTriggerGithubPullRequest is a filter to match changes in pull requests.
type CloudBuildTriggerGithubPullRequest struct {
	// Regex of branches to match.
	Branch string `json:"branch"`

	// Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
	CommentControl *string `json:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// CloudBuildTriggerGithubPush is a filter to match changes in refs, like branches or tags.
type CloudBuildTriggerGithubPush struct {
	// Regex of branches to match. Specify only one of branch or tag.
	Branch *string `json:"branch,omitempty"`

	// When true, only trigger a build if the revision regex does NOT match the git_ref regex.
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Regex of tags to match. Specify only one of branch or tag.
	Tag *string `json:"tag,omitempty"`
}

// CloudBuildTriggerPubsubConfig describes the configuration of a trigger that creates a build
// whenever a Pub/Sub message is published.
type CloudBuildTriggerPubsubConfig struct {
	// Service account that will make the push request.
	ServiceAccountRef *CloudBuildIAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Potential issues with the underlying Pub/Sub subscription configuration.
	State *string `json:"state,omitempty"`

	// Output only. Name of the subscription.
	Subscription *string `json:"subscription,omitempty"`

	// The name of the topic from which this subscription is receiving messages.
	TopicRef CloudBuildPubSubTopicRef `json:"topicRef"`
}

// CloudBuildTriggerRepositoryEventConfig is the configuration of a trigger that creates a build
// whenever an event from Repo API is received.
type CloudBuildTriggerRepositoryEventConfig struct {
	// Contains filter properties for matching Pull Requests.
	PullRequest *CloudBuildTriggerRepositoryEventConfigPullRequest `json:"pullRequest,omitempty"`

	// Contains filter properties for matching git pushes.
	Push *CloudBuildTriggerRepositoryEventConfigPush `json:"push,omitempty"`

	// The resource name of the Repo API resource.
	Repository *string `json:"repository,omitempty"`
}

// CloudBuildTriggerRepositoryEventConfigPullRequest contains filter properties for matching Pull Requests.
type CloudBuildTriggerRepositoryEventConfigPullRequest struct {
	// Regex of branches to match.
	Branch *string `json:"branch,omitempty"`

	// Configure builds to run whether a repository owner or collaborator need to comment '/gcbrun'.
	CommentControl *string `json:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// CloudBuildTriggerRepositoryEventConfigPush contains filter properties for matching git pushes.
type CloudBuildTriggerRepositoryEventConfigPush struct {
	// Regex of branches to match.
	Branch *string `json:"branch,omitempty"`

	// If true, only trigger a build if the revision regex does NOT match the git_ref regex.
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Regex of tags to match.
	Tag *string `json:"tag,omitempty"`
}

// CloudBuildTriggerSourceToBuild is the repo and ref of the repository from which to build.
type CloudBuildTriggerSourceToBuild struct {
	// Only `external` field is supported to configure the reference.
	// The full resource name of the bitbucket server config.
	BitbucketServerConfigRef *CloudBuildBitbucketServerConfigRef `json:"bitbucketServerConfigRef,omitempty"`

	// Only `external` field is supported to configure the reference.
	// The full resource name of the github enterprise config.
	GithubEnterpriseConfigRef *CloudBuildGithubEnterpriseConfigRef `json:"githubEnterpriseConfigRef,omitempty"`

	// The branch or tag to use. Must start with "refs/" (required).
	Ref string `json:"ref"`

	// The type of the repo, since it may not be explicit from the repo field.
	RepoType string `json:"repoType"`

	// Only `external` field is supported to configure the reference.
	// The qualified resource name of the Repo API repository.
	RepositoryRef *CloudBuildV2RepositoryRef `json:"repositoryRef,omitempty"`

	// The URI of the repo.
	Uri *string `json:"uri,omitempty"`
}

// CloudBuildTriggerTriggerTemplate describes the types of source changes to trigger a build.
type CloudBuildTriggerTriggerTemplate struct {
	// Name of the branch to build.
	BranchName *string `json:"branchName,omitempty"`

	// Explicit commit SHA to build.
	CommitSha *string `json:"commitSha,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	Dir *string `json:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision regex.
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// The Cloud Source Repository to build.
	RepoRef *CloudBuildSourceRepoRepositoryRef `json:"repoRef,omitempty"`

	// Name of the tag to build.
	TagName *string `json:"tagName,omitempty"`
}

// CloudBuildTriggerWebhookConfig describes the configuration of a trigger that creates a build
// whenever a webhook is sent to a trigger's webhook URL.
type CloudBuildTriggerWebhookConfig struct {
	// The secret required.
	SecretRef CloudBuildSecretManagerSecretVersionRef `json:"secretRef"`

	// Potential issues with the underlying Pub/Sub subscription configuration.
	State *string `json:"state,omitempty"`
}

// CloudBuildTriggerSpec defines the desired state of CloudBuildTrigger
type CloudBuildTriggerSpec struct {
	// Configuration for manual approval to start a build invocation of this BuildTrigger.
	ApprovalConfig *CloudBuildTriggerApprovalConfig `json:"approvalConfig,omitempty"`

	// BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build
	// whenever a Bitbucket Server event is received.
	BitbucketServerTriggerConfig *CloudBuildTriggerBitbucketServerTriggerConfig `json:"bitbucketServerTriggerConfig,omitempty"`

	// Contents of the build template. Either a filename or build template must be provided.
	Build *CloudBuildTriggerBuild `json:"build,omitempty"`

	// Human-readable description of the trigger.
	Description *string `json:"description,omitempty"`

	// Whether the trigger is disabled or not.
	Disabled *bool `json:"disabled,omitempty"`

	// Path, from the source root, to a file whose contents is used for the template.
	Filename *string `json:"filename,omitempty"`

	// A Common Expression Language string. Used only with Pub/Sub and Webhook.
	Filter *string `json:"filter,omitempty"`

	// The file source describing the local or remote Build template.
	GitFileSource *CloudBuildTriggerGitFileSource `json:"gitFileSource,omitempty"`

	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	Github *CloudBuildTriggerGithub `json:"github,omitempty"`

	// ignoredFiles and includedFiles are file glob matches.
	IgnoredFiles []string `json:"ignoredFiles,omitempty"`

	// Build logs will be sent back to GitHub as part of the checkrun result.
	IncludeBuildLogs *string `json:"includeBuildLogs,omitempty"`

	// ignoredFiles and includedFiles are file glob matches.
	IncludedFiles []string `json:"includedFiles,omitempty"`

	// Immutable. The location of the Cloud Build trigger.
	// If not specified, "global" is used. More info: cloud.google.com/build/docs/locations.
	Location *string `json:"location,omitempty"`

	// PubsubConfig describes the configuration of a trigger that creates a build
	// whenever a Pub/Sub message is published.
	PubsubConfig *CloudBuildTriggerPubsubConfig `json:"pubsubConfig,omitempty"`

	// The configuration of a trigger that creates a build whenever an event from Repo API is received.
	RepositoryEventConfig *CloudBuildTriggerRepositoryEventConfig `json:"repositoryEventConfig,omitempty"`

	// The service account used for all user-controlled operations.
	ServiceAccountRef *CloudBuildIAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// The repo and ref of the repository from which to build.
	SourceToBuild *CloudBuildTriggerSourceToBuild `json:"sourceToBuild,omitempty"`

	// Substitutions data for Build resource.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Tags for annotation of a BuildTrigger.
	Tags []string `json:"tags,omitempty"`

	// Template describing the types of source changes to trigger a build.
	TriggerTemplate *CloudBuildTriggerTriggerTemplate `json:"triggerTemplate,omitempty"`

	// WebhookConfig describes the configuration of a trigger that creates a build
	// whenever a webhook is sent to a trigger's webhook URL.
	WebhookConfig *CloudBuildTriggerWebhookConfig `json:"webhookConfig,omitempty"`
}

// CloudBuildTriggerStatus defines the observed state of CloudBuildTrigger
type CloudBuildTriggerStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	// Time when the trigger was created.
	CreateTime *string `json:"createTime,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent
	// desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The unique identifier for the trigger.
	TriggerId *string `json:"triggerId,omitempty"`
}

// CloudBuildTrigger is the Schema for the CloudBuildTrigger API
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildtrigger;gcpcloudbuildtriggers
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",description="When 'True', the most recent reconcile of the resource succeeded",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Status",type="string",description="The reason for the value in 'Ready'",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Status Age",type="date",description="The last transition time for the value in 'Status'",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime"
type CloudBuildTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudBuildTriggerSpec   `json:"spec,omitempty"`
	Status CloudBuildTriggerStatus `json:"status,omitempty"`
}

// CloudBuildTriggerList contains a list of CloudBuildTrigger
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CloudBuildTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildTrigger{}, &CloudBuildTriggerList{})
}
