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
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildTriggerGVK = GroupVersion.WithKind("CloudBuildTrigger")

// +kcc:proto=google.devtools.cloudbuild.v1.BuildTrigger
type CloudBuildTriggerSpec struct {
	// Configuration for manual approval to start a build invocation of this BuildTrigger.
	//  Builds created by this trigger will require approval before they execute.
	//  Any user with a Cloud Build Approver role for the project can approve a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.approval_config
	ApprovalConfig *ApprovalConfig `json:"approvalConfig,omitempty"`

	// BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build whenever a Bitbucket Server event is received.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.bitbucket_server_trigger_config
	BitbucketServerTriggerConfig *BitbucketServerTriggerConfig `json:"bitbucketServerTriggerConfig,omitempty"`

	// Contents of the build template. Either a filename or build template must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.build
	Build *Build `json:"build,omitempty"`

	// Human-readable description of the trigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.description
	Description *string `json:"description,omitempty"`

	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided. Set this only when using trigger_template or github. When using Pub/Sub, Webhook or Manual set the file name using git_file_source instead.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.filename
	Filename *string `json:"filename,omitempty"`

	// A Common Expression Language string. Used only with Pub/Sub and Webhook.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.filter
	Filter *string `json:"filter,omitempty"`

	// The file source describing the local or remote Build template.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.git_file_source
	GitFileSource *GitFileSource `json:"gitFileSource,omitempty"`

	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	//  One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.github
	Github *GitHubEventsConfig `json:"github,omitempty"`

	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for '**'. If ignoredFiles and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignoredFiles is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignoredFiles globs, then we do not trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.ignored_files
	IgnoredFiles []string `json:"ignoredFiles,omitempty"`

	// Build logs will be sent back to GitHub as part of the checkrun result. Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or INCLUDE_BUILD_LOGS_WITH_STATUS Possible values: ["INCLUDE_BUILD_LOGS_UNSPECIFIED", "INCLUDE_BUILD_LOGS_WITH_STATUS"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.include_build_logs
	IncludeBuildLogs *string `json:"includeBuildLogs,omitempty"`

	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for '**'. If any of the files altered in the commit pass the ignoredFiles filter and includedFiles is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignoredFiles filter and includedFiles is not empty, then we make sure that at least one of those files matches a includedFiles glob. If not, then we do not trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.included_files
	IncludedFiles []string `json:"includedFiles,omitempty"`

	// Immutable. The location of the Cloud Build trigger. If not specified, "global" is used. More info: cloud.google.com/build/docs/locations.
	Location *string `json:"location,omitempty"`

	// PubsubConfig describes the configuration of a trigger that creates a build whenever a Pub/Sub message is published.
	//  One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.pubsub_config
	PubsubConfig *PubsubConfig `json:"pubsubConfig,omitempty"`

	// The configuration of a trigger that creates a build whenever an event from Repo API is received.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.repository_event_config
	RepositoryEventConfig *RepositoryEventConfig `json:"repositoryEventConfig,omitempty"`

	// The service account used for all user-controlled operations including triggers.patch, triggers.run, builds.create, and builds.cancel. If no service account is set, then the standard Cloud Build service account ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead. When populating via the external field, the following format is supported: projects/{PROJECT_ID}/serviceAccounts/{SERVICE_ACCOUNT_EMAIL}
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// The repo and ref of the repository from which to build. This field is used only for those triggers that do not respond to SCM events. Triggers that respond to such events build source at whatever commit caused the event. This field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers. One of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.source_to_build
	SourceToBuild *SourceToBuild `json:"sourceToBuild,omitempty"`

	// Substitutions data for Build resource.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Tags for annotation of a BuildTrigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.tags
	Tags []string `json:"tags,omitempty"`

	// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. One of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.trigger_template
	TriggerTemplate *RepoSource `json:"triggerTemplate,omitempty"`

	// WebhookConfig describes the configuration of a trigger that creates a build whenever a webhook is sent to a trigger's webhook URL. One of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.webhook_config
	WebhookConfig *WebhookConfig `json:"webhookConfig,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.ApprovalConfig
type ApprovalConfig struct {
	// Whether or not approval is needed. If this is set on a build, it will become pending when run, and will need to be explicitly approved to start.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalConfig.approval_required
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BitbucketServerTriggerConfig
type BitbucketServerTriggerConfig struct {
	// Only `external` field is supported to configure the reference. The full resource name of the bitbucket server config. Format: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BitbucketServerTriggerConfig.bitbucket_server_config_resource
	BitbucketServerConfigResourceRef *refsv1beta1.BitbucketServerConfigRef `json:"bitbucketServerConfigResourceRef"`

	// Key of the project that the repo is in. For example: The key for https://mybitbucket.server/projects/TEST/repos/test-repo is "TEST".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BitbucketServerTriggerConfig.project_key
	ProjectKey string `json:"projectKey"`

	// Filter to match changes in pull requests.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BitbucketServerTriggerConfig.pull_request
	PullRequest *BitbucketServerTriggerConfig_PullRequest `json:"pullRequest,omitempty"`

	// Filter to match changes in refs like branches, tags.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BitbucketServerTriggerConfig.push
	Push *BitbucketServerTriggerConfig_Push `json:"push,omitempty"`

	// Slug of the repository. A repository slug is a URL-friendly version of a repository name, automatically generated by Bitbucket for use in the URL. For example, if the repository name is 'test repo', in the URL it would become 'test-repo' as in https://mybitbucket.server/projects/TEST/repos/test-repo.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BitbucketServerTriggerConfig.repo_slug
	RepoSlug string `json:"repoSlug"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PullRequestFilter
type BitbucketServerTriggerConfig_PullRequest struct {
	// Regex of branches to match. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.branch
	Branch *string `json:"branch"`

	// Configure builds to run whether a repository owner or collaborator need to comment /gcbrun. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.comment_control
	CommentControl *string `json:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PushFilter
type BitbucketServerTriggerConfig_Push struct {
	// Regex of branches to match. Specify only one of branch or tag.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.branch
	Branch *string `json:"branch,omitempty"`

	// When true, only trigger a build if the revision regex does NOT match the gitRef regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Regex of tags to match. Specify only one of branch or tag.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build
type Build struct {
	// Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.artifacts
	Artifacts *Artifacts `json:"artifacts,omitempty"`

	// Secrets and secret environment variables.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.available_secrets
	AvailableSecrets *Secrets `json:"availableSecrets,omitempty"`

	// A list of images to be pushed upon the successful completion of all build steps. The images are pushed using the builder service account's credentials. The digests of the pushed images will be stored in the Build resource's results field. If any of the images fail to be pushed, the build status is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.images
	Images []string `json:"images,omitempty"`

	// Google Cloud Storage bucket where logs should be written. Logs file names will be of the format ${logsBucket}/log-${build_id}.txt.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.logs_bucket
	LogsBucketRef *storagev1beta1.StorageBucketRef `json:"logsBucketRef,omitempty"`

	// Special options for this build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.options
	Options *BuildOptions `json:"options,omitempty"`

	// TTL in queue for this build. If provided and the build is enqueued longer than this value, the build will expire and the build status will be EXPIRED. The TTL starts ticking from createTime. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.queue_ttl
	QueueTTL *string `json:"queueTtl,omitempty"`

	// Secrets to decrypt using Cloud Key Management Service.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.secrets
	Secret []Secret `json:"secret,omitempty"`

	// The location of the source files to build. One of 'storageSource' or 'repoSource' must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.source
	Source *Source `json:"source,omitempty"`

	// The operations to be performed on the workspace.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.steps
	Step []BuildStep `json:"step,omitempty"`

	// Substitutions data for Build resource.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Tags for annotation of a Build. These are not docker tags.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.tags
	Tags []string `json:"tags,omitempty"`

	// Amount of time that this build should be allowed to run, to second granularity. If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT. This timeout must be equal to or greater than the sum of the timeouts for build steps within the build. The expected format is the number of seconds followed by s. Default time is ten minutes (600s).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.timeout
	Timeout *string `json:"timeout,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts
type Artifacts struct {
	// A list of images to be pushed upon the successful completion of all build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.images
	Images []string `json:"images,omitempty"`

	// A list of objects to be uploaded to Cloud Storage upon successful completion of all build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.objects
	Objects *Artifacts_ArtifactObjects `json:"objects,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects
type Artifacts_ArtifactObjects struct {
	// Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/". Files in the workspace matching any path pattern will be uploaded to Cloud Storage with this location as a prefix.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects.location
	Location *string `json:"location,omitempty"`

	// Path globs used to match files in the build's workspace.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects.paths
	Paths []string `json:"paths,omitempty"`

	// Output only. Stores timing information for pushing all artifacts.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects.timing
	Timing []*Artifacts_Timing `json:"timing,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.TimeSpan
type Artifacts_Timing struct {
	// End of time span.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.TimeSpan.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Start of time span.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.TimeSpan.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.TimeSpan
type BuildStep_Timing struct {
	// End of time span.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.TimeSpan.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Start of time span.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.TimeSpan.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Secrets
type Secrets struct {
	// Secret Manager secrets.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Secrets.secret_manager
	SecretManager []Secrets_SecretManager `json:"secretManager,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.SecretManagerSecret
type Secrets_SecretManager struct {
	// Environment variable name to associate with the secret. Secret environment variables must be unique across all of a build's secrets, and must be used by at least one build step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SecretManagerSecret.env
	Env *string `json:"env"`

	// The name of the SecretManagerSecretVersion resource.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SecretManagerSecret.version_name
	VersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"versionRef"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildOptions
type BuildOptions struct {
	// Requested disk size for the VM that runs the build. Note that this is *NOT* "disk free"; some of the space will be used by the operating system and build utilities. Also note that this is the minimum disk size that will be allocated for the build -- the build may run with a larger disk than requested. At present, the maximum disk size is 4000GB; builds that request more than the maximum are rejected with an error.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.disk_size_gb
	DiskSizeGB *int `json:"diskSizeGb,omitempty"`

	// Option to specify whether or not to apply bash style string operations to the substitutions. NOTE: this is always enabled for triggered builds and cannot be overridden in the build configuration file.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.dynamic_substitutions
	DynamicSubstitutions *bool `json:"dynamicSubstitutions,omitempty"`

	// A list of global environment variable definitions that will exist for all build steps in this build. If a variable is defined in both globally and in a build step, the variable will use the build step value. The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.env
	Env []string `json:"env,omitempty"`

	// Option to define build log streaming behavior to Google Cloud Storage. Possible values: ["STREAM_DEFAULT", "STREAM_ON", "STREAM_OFF"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.log_streaming_option
	LogStreamingOption *string `json:"logStreamingOption,omitempty"`

	// Option to specify the logging mode, which determines if and where build logs are stored. Possible values: ["LOGGING_UNSPECIFIED", "LEGACY", "GCS_ONLY", "STACKDRIVER_ONLY", "CLOUD_LOGGING_ONLY", "NONE"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.logging
	Logging *string `json:"logging,omitempty"`

	// Compute Engine machine type on which to run the build. Possible values: ["UNSPECIFIED", "N1_HIGHCPU_8", "N1_HIGHCPU_32", "E2_HIGHCPU_8", "E2_HIGHCPU_32"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Requested verifiability options. Possible values: ["NOT_VERIFIED", "VERIFIED"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.requested_verify_option
	RequestedVerifyOption *string `json:"requestedVerifyOption,omitempty"`

	// A list of global environment variables, which are encrypted using a Cloud Key Management Service crypto key. These values must be specified in the build's Secret. These variables will be available to all build steps in this build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.secret_env
	SecretEnv []string `json:"secretEnv,omitempty"`

	// Requested hash for SourceProvenance. Possible values: ["NONE", "SHA256", "MD5"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.source_provenance_hash
	SourceProvenanceHash []string `json:"sourceProvenanceHash,omitempty"`

	// Option to specify behavior when there is an error in the substitution checks. NOTE: this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden in the build configuration file. Possible values: ["MUST_MATCH", "ALLOW_LOOSE"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.substitution_option
	SubstitutionOption *string `json:"substitutionOption,omitempty"`

	// Global list of volumes to mount for ALL build steps Each volume is created as an empty volume prior to starting the build process. Upon completion of the build, volumes and their contents are discarded. Global volume names and paths cannot conflict with the volumes defined a build step. Using a global volume in a build with only one step is not valid as it is indicative of a build request with an incorrect configuration.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.volumes
	Volumes []BuildOptions_Volume `json:"volumes,omitempty"`

	// This field deprecated; please use pool.name instead.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Volume
type BuildOptions_Volume struct {
	// Name of the volume to mount.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Volume.name
	Name *string `json:"name,omitempty"`

	// Path at which to mount the volume.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Volume.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Secret
type Secret struct {
	// Resource name of Cloud KMS crypto key to decrypt the encrypted value. In format: projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Secret.kms_key_name
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef"`

	// Map of environment variable name to its encrypted value. Secret environment variables must be unique across all of a build's secrets, and must be used by at least one build step.
	SecretEnv map[string]string `json:"secretEnv,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Source
type Source struct {
	// RepoSource describes the location of the source in a Google Cloud Source Repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.repo_source
	RepoSource *Source_RepoSource `json:"repoSource,omitempty"`

	// StorageSource describes the location of the source in an archive file in Google Cloud Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.storage_source
	StorageSource *Source_StorageSource `json:"storageSource,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.RepoSource
type Source_RepoSource struct {
	// Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided. This field is a regular expression.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.branch_name
	BranchName *string `json:"branchName,omitempty"`

	// Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`

	// Directory, relative to the source root, in which to run the build. This must be a relative path. If a step's dir is specified and is an absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.dir
	Dir *string `json:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.project_id
	ProjectID *string `json:"projectId,omitempty"`

	// The Cloud Source Repository to build. If omitted, the repo with name "default" is assumed.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.repo_name
	RepoReference refsv1beta1.SourceRepoRepositoryRef `json:"repoRef"`

	// Substitutions to use in a triggered build. Should only be used with RunBuildTrigger
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided. This field is a regular expression.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.tag_name
	TagName *string `json:"tagName,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.StorageSource
type Source_StorageSource struct {
	// Google Cloud Storage bucket containing the source.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef"`

	// Google Cloud Storage generation for the object. If the generation is omitted, the latest generation will be used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.generation
	Generation *string `json:"generation,omitempty"`

	// Google Cloud Storage object containing the source. This object must be a gzipped archive file (.tar.gz) containing source to build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.object
	Object *string `json:"object"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildStep
type BuildStep struct {
	// Allow this build step to fail without failing the entire build if and only if the exit code is one of the specified codes. If allowFailure is also specified, this field will take precedence.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.allow_exit_codes
	AllowExitCodes []int `json:"allowExitCodes,omitempty"`

	// Allow this build step to fail without failing the entire build. If false, the entire build will fail if this step fails. Otherwise, the build will succeed, but this step will still have a failure status. Error information will be reported in the failureDetail field.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.allow_failure
	AllowFailure *bool `json:"allowFailure,omitempty"`

	// A list of arguments that will be presented to the step when it is started. If the image used to run the step's container has an entrypoint, the args are used as arguments to that entrypoint. If the image does not define an entrypoint, the first element in args is used as the entrypoint, and the remainder will be used as arguments.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.args
	Args []string `json:"args,omitempty"`

	// Working directory to use when running this step's container. If this value is a relative path, it is relative to the build's working directory. If this value is absolute, it may be outside the build's working directory, in which case the contents of the path may not be persisted across build step executions, unless a volume for that path is specified. If the build specifies a RepoSource with dir and a step with a dir, which specifies an absolute path, the RepoSource dir is ignored for the step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.dir
	Dir *string `json:"dir,omitempty"`

	// Entrypoint to be used instead of the build step image's default entrypoint. If unset, the image's default entrypoint is used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.entrypoint
	Entrypoint *string `json:"entrypoint,omitempty"`

	// A list of environment variable definitions to be used when running a step. The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.env
	Env []string `json:"env,omitempty"`

	// Unique identifier for this build step, used in waitFor to reference this build step as a dependency.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.id
	ID *string `json:"id,omitempty"`

	// The name of the container image that will run this particular build step. If the image is available in the host's Docker daemon's cache, it will be run directly. If not, the host will attempt to pull the image first, using the builder service account's credentials if necessary. The Docker daemon's cache will already have the latest versions of all of the officially supported build steps (https://github.com/GoogleCloudPlatform/cloud-builders). The Docker daemon will also have cached many of the layers for some popular images, like "ubuntu", "debian", but they will be refreshed at the time you attempt to use them. If you built an image in a previous build step, it will be stored in the host's Docker daemon's cache and is available to use as the name for a later build step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.name
	Name *string `json:"name"`

	// A shell script to be executed in the step. When script is provided, the user cannot specify the entrypoint or args.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.script
	Script *string `json:"script,omitempty"`

	// A list of environment variables which are encrypted using a Cloud Key Management Service crypto key. These values must be specified in the build's Secret.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.secret_env
	SecretEnv []string `json:"secretEnv,omitempty"`

	// Time limit for executing this build step. If not defined, the step has no time limit and will be allowed to continue to run until either it completes or the build itself times out.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Output only. Stores timing information for executing this build step.
	// +kcc:proto:field=MISSING
	StepTiming *string `json:"timing,omitempty"`

	// List of volumes to mount into the build step. Each volume is created as an empty volume prior to execution of the build step. Upon completion of the build, volumes and their contents are discarded. Using a named volume in only one step is not valid as it is indicative of a build request with an incorrect configuration.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.volumes
	Volumes []BuildStep_Volume `json:"volumes,omitempty"`

	// The ID(s) of the step(s) that this build step depends on. This build step will not start until all the build steps in waitFor have completed successfully. If waitFor is empty, this build step will start when all previous build steps in the Build.Steps list have completed successfully.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.wait_for
	WaitFor []string `json:"waitFor,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Volume
type BuildStep_Volume struct {
	// Name of the volume to mount.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Volume.name
	Name *string `json:"name"`

	// Path at which to mount the volume.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Volume.path
	Path *string `json:"path"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitFileSource
type GitFileSource struct {
	// Only `external` field is supported to configure the reference. The full resource name of the bitbucket server config. Format: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.
	// +kcc:proto:field=MISSING
	BitbucketServerConfigReference *refsv1beta1.BitbucketServerConfigRef `json:"bitbucketServerConfigRef,omitempty"`

	// Only `external` field is supported to configure the reference. The full resource name of the github enterprise config. Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.
	// +kcc:proto:field=MISSING
	GithubEnterpriseConfigReference *refsv1beta1.GithubEnterpriseConfigRef `json:"githubEnterpriseConfigRef,omitempty"`

	// The path of the file, with the repo root as the root of the path.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.path
	Path *string `json:"path"`

	// The type of the repo, since it may not be explicit from the repo field (e.g from a URL). Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.repo_type
	RepoType *string `json:"repoType"`

	// Only `external` field is supported to configure the reference. The qualified resource name of the Repo API repository. Either uri or repository can be specified and is required.
	// +kcc:proto:field=MISSING
	RepositoryReference *refsv1beta1.CloudBuildV2RepositoryRef `json:"repositoryRef,omitempty"`

	// The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the filename (optional). This field respects the same syntax/resolution as described here: https://git-scm.com/docs/gitrevisions If unspecified, the revision from which the trigger invocation originated is assumed to be the revision from which to read the specified path.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.revision
	Revision *string `json:"revision,omitempty"`

	// The URI of the repo (optional). If unspecified, the repo from which the trigger invocation originated is assumed to be the repo from which to read the specified path.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitFileSource.uri
	Uri *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitHubEventsConfig
type GitHubEventsConfig struct {
	// Only `external` field is supported to configure the reference. The full resource name of the github enterprise config. Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}.
	// +kcc:proto:field=MISSING
	EnterpriseConfigResourceNameReference *refsv1beta1.GithubEnterpriseConfigRef `json:"enterpriseConfigResourceNameRef,omitempty"`

	// Name of the repository. For example: The name for https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEventsConfig.name
	Name *string `json:"name,omitempty"`

	// Owner of the repository. For example: The owner for https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEventsConfig.owner
	Owner *string `json:"owner,omitempty"`

	// filter to match changes in pull requests. Specify only one of 'pull_request' or 'push'.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEventsConfig.pull_request
	PullRequest *GitHubEventsConfig_PullRequest `json:"pullRequest,omitempty"`

	// filter to match changes in refs, like branches or tags. Specify only one of 'pull_request' or 'push'.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEventsConfig.push
	Push *GitHubEventsConfig_Push `json:"push,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PullRequestFilter
type GitHubEventsConfig_PullRequest struct {
	// Regex of branches to match.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.branch
	Branch *string `json:"branch"`

	// Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.comment_control
	CommentControl *string `json:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PushFilter
type GitHubEventsConfig_Push struct {
	// Regex of branches to match. Specify only one of branch or tag.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.branch
	Branch *string `json:"branch,omitempty"`

	// When true, only trigger a build if the revision regex does NOT match the git_ref regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Regex of tags to match. Specify only one of branch or tag.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PubsubConfig
type PubsubConfig struct {
	// Service account that will make the push request.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.service_account_email
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Potential issues with the underlying Pub/Sub subscription configuration. Only populated on get requests.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.state
	State *string `json:"state,omitempty"`

	// Output only. Name of the subscription.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.subscription
	Subscription *string `json:"subscription,omitempty"`

	// The name of the topic from which this subscription is receiving messages.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PubsubConfig.topic
	TopicRef *pubsubv1beta1.PubSubTopicRef `json:"topicRef"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.RepositoryEventConfig
type RepositoryEventConfig struct {
	// Contains filter properties for matching Pull Requests.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepositoryEventConfig.pull_request
	PullRequest *RepositoryEventConfig_PullRequest `json:"pullRequest,omitempty"`

	// Contains filter properties for matching git pushes.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepositoryEventConfig.push
	Push *RepositoryEventConfig_Push `json:"push,omitempty"`

	// The resource name of the Repo API resource.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepositoryEventConfig.repository
	Repository *string `json:"repository,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PullRequestFilter
type RepositoryEventConfig_PullRequest struct {
	// Regex of branches to match. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.branch
	Branch *string `json:"branch,omitempty"`

	// Configure builds to run whether a repository owner or collaborator need to comment '/gcbrun'. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.comment_control
	CommentControl *string `json:"commentControl,omitempty"`

	// If true, branches that do NOT match the git_ref will trigger a build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PullRequestFilter.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PushFilter
type RepositoryEventConfig_Push struct {
	// Regex of branches to match. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.branch
	Branch *string `json:"branch,omitempty"`

	// If true, only trigger a build if the revision regex does NOT match the git_ref regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Regex of tags to match. The syntax of the regular expressions accepted is the syntax accepted by RE2 and described at https://github.com/google/re2/wiki/Syntax.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.PushFilter.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitRepoSource
type SourceToBuild struct {
	// Only `external` field is supported to configure the reference. The full resource name of the bitbucket server config. Format: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.
	// +kcc:proto:field=MISSING
	BitbucketServerConfigReference *refsv1beta1.BitbucketServerConfigRef `json:"bitbucketServerConfigRef,omitempty"`

	// Only `external` field is supported to configure the reference. The full resource name of the github enterprise config. Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}.
	// +kcc:proto:field=MISSING
	GithubEnterpriseConfigReference *refsv1beta1.GithubEnterpriseConfigRef `json:"githubEnterpriseConfigRef,omitempty"`

	// The branch or tag to use. Must start with "refs/" (required).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitRepoSource.ref
	Ref *string `json:"ref"`

	// The type of the repo, since it may not be explicit from the repo field (e.g from a URL). Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitRepoSource.repo_type
	RepoType *string `json:"repoType"`

	// Only `external` field is supported to configure the reference. The qualified resource name of the Repo API repository. Either uri or repository can be specified and is required.
	// +kcc:proto:field=MISSING
	RepositoryReference *refsv1beta1.CloudBuildV2RepositoryRef `json:"repositoryRef,omitempty"`

	// The URI of the repo.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitRepoSource.uri
	Uri *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.RepoSource
type RepoSource struct {
	// Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided. This field is a regular expression.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.branch_name
	BranchName *string `json:"branchName,omitempty"`

	// Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`

	// Directory, relative to the source root, in which to run the build. This must be a relative path. If a step's dir is specified and is an absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.dir
	Dir *string `json:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// The Cloud Source Repository to build. If omitted, the repo with name "default" is assumed.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.repo_name
	RepoReference *refsv1beta1.SourceRepoRepositoryRef `json:"repoRef,omitempty"`

	// Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided. This field is a regular expression.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.tag_name
	TagName *string `json:"tagName,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.WebhookConfig
type WebhookConfig struct {
	// The secret required
	// +kcc:proto:field=MISSING
	SecretReference *refsv1beta1.SecretManagerSecretVersionRef `json:"secretRef"`

	// Potential issues with the underlying Pub/Sub subscription configuration. Only populated on get requests.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.WebhookConfig.state
	State *string `json:"state,omitempty"`
}

type CloudBuildTriggerStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Time when the trigger was created.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The unique identifier for the trigger.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildTrigger.id
	TriggerID *string `json:"triggerId,omitempty"`

	// A unique specifier for the CloudBuildTrigger resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudBuildTriggerObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildTrigger
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

// +kcc:proto=google.devtools.cloudbuild.v1.BuildTrigger
// +k8s:openapi-gen=true
type CloudBuildTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

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

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts
type ArtifactsObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects
type Artifacts_ArtifactObjectsObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildTrigger
type BuildTriggerObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.PubsubConfig
type PubsubConfigObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.RepositoryEventConfig
type RepositoryEventConfigObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.ApprovalResult
type ApprovalResult struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.decision
	Decision *string `json:"decision,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.comment
	Comment *string `json:"comment,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.url
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.ApprovalResult
type ApprovalResultObservedState struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.approver_account
	ApproverAccount *string `json:"approverAccount,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.approval_time
	ApprovalTime *string `json:"approvalTime,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.GoModule

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.NpmPackage

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.PythonPackage

// +kcc:proto=google.devtools.cloudbuild.v1.FileHashes

// +kcc:proto=google.devtools.cloudbuild.v1.Hash

// +kcc:proto=google.devtools.cloudbuild.v1.BuiltImage
type BuiltImage struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.name
	Name *string `json:"name,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.digest
	Digest *string `json:"digest,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.push_timing
	PushTiming *Artifacts_Timing `json:"pushTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuiltImage
type BuiltImageObservedState struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.push_timing
	PushTiming *Artifacts_Timing `json:"pushTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Results
type Results struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.images
	Images []BuiltImage `json:"images,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.build_step_images
	BuildStepImages []string `json:"buildStepImages,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.artifact_manifest
	ArtifactManifest *string `json:"artifactManifest,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.num_artifacts
	NumArtifacts *int64 `json:"numArtifacts,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.artifact_timing
	ArtifactTiming *Artifacts_Timing `json:"artifactTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build
type BuildObservedState struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.id
	ID *string `json:"id,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.status
	Status *string `json:"status,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildApproval
type BuildApproval struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.state
	State *string `json:"state,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.config
	Config *ApprovalConfig `json:"config,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.result
	Result *ApprovalResult `json:"result,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildApproval
type BuildApprovalObservedState struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.state
	State *string `json:"state,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.config
	Config *ApprovalConfig `json:"config,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.result
	Result *ApprovalResult `json:"result,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildOptions.PoolOption
type BuildOptions_PoolOption struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build.FailureInfo
type Build_FailureInfo struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.FailureInfo.type
	Type *string `json:"type,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.FailureInfo.detail
	Detail *string `json:"detail,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build.Warning
type Build_Warning struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.Warning.text
	Text *string `json:"text,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.Warning.priority
	Priority *string `json:"priority,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildStep
type BuildStepObservedState struct {
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.status
	Status *string `json:"status,omitempty"`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.exit_code
	ExitCode *int32 `json:"exitCode,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Dependency
type Dependency struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.Dependency.GitSourceDependency
type Dependency_GitSourceDependency struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.Dependency.GitSourceRepository
type Dependency_GitSourceRepository struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitConfig
type GitConfig struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitConfig.HttpConfig
type GitConfig_HTTPConfig struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitSource
type GitSource struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.InlineSecret
type InlineSecret struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.Results
type ResultsObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.SourceProvenance
type SourceProvenance struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.SourceProvenance
type SourceProvenanceObservedState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v1.StorageSourceManifest

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedGoModule

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedGoModule

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedMavenArtifact

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedMavenArtifact

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedNpmPackage

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedNpmPackage

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedPythonPackage

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedPythonPackage
