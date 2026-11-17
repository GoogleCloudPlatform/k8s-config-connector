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
// +kcc:source:resource-docs=https://cloud.google.com/build/docs/api/reference/rest/v1/projects.locations.builds

package v1alpha1

import (
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	developerconnectv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildBuildGVK = GroupVersion.WithKind("CloudBuildBuild")

// CloudBuildBuildSpec defines the desired state of CloudBuildBuild
// +kcc:spec:proto=google.devtools.cloudbuild.v1.Build
// +kcc:required-from-proto
type CloudBuildBuildSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// Immutable. Optional. The server-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// The location of the source files to build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.source
	Source *Build_Source `json:"source,omitempty"`

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
	LogsBucketRef *storagev1beta1.StorageBucketRef `json:"logsBucketRef,omitempty"`

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
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

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

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.GoModule
type Artifacts_GoModule struct {
	// Optional. Location of the Artifact Registry repository. i.e. us-east1
	//  Defaults to the build’s location.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.repository_location
	RepositoryLocation *string `json:"repositoryLocation,omitempty"`

	// Optional. Name of the Artifact Registry repository.
	//  e.g. my-go-repo
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.repository_name
	RepositoryName *string `json:"repositoryName,omitempty"`

	// Optional. Project ID of the Artifact Registry repository.
	//  Defaults to the build project.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.repository_project_id
	RepositoryProjectRef *refsv1beta1.ProjectRef `json:"repositoryProjectRef,omitempty"`

	// Optional. Source path of the go.mod file in the build's workspace. If not
	//  specified, this will default to the current directory.
	//  e.g. ~/code/go/mypackage
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.source_path
	SourcePath *string `json:"sourcePath,omitempty"`

	// Optional. The Go module's "module path".
	//  e.g. example.com/foo/v2
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.module_path
	ModulePath *string `json:"modulePath,omitempty"`

	// Optional. The Go module's semantic version in the form vX.Y.Z. e.g.
	//  v0.1.1 Pre-release identifiers can also be added by appending a dash and
	//  dot separated ASCII alphanumeric characters and hyphens.
	//  e.g. v0.2.3-alpha.x.12m.5
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.module_version
	ModuleVersion *string `json:"moduleVersion,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildOptions
type BuildOptions struct {
	// Requested hash for SourceProvenance.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.source_provenance_hash
	SourceProvenanceHash []string `json:"sourceProvenanceHash,omitempty"`

	// Requested verifiability options.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.requested_verify_option
	RequestedVerifyOption *string `json:"requestedVerifyOption,omitempty"`

	// Compute Engine machine type on which to run the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Requested disk size for the VM that runs the build. Note that this is *NOT*
	//  "disk free"; some of the space will be used by the operating system and
	//  build utilities. Also note that this is the minimum disk size that will be
	//  allocated for the build (see until_waste).
	//  Disk size is specified in GB.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Option to specify behavior when there is an error in the substitution checks.
	//
	//  NOTE: this is always set to ALLOW_LOOSE for builds.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.substitution_option
	SubstitutionOption *string `json:"substitutionOption,omitempty"`

	// Option to specify whether default substitutions should be replaced with
	//  a regex-based dynamic substitutions, and whether dynamic substitutions
	//  should be evaluated.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.dynamic_substitutions
	DynamicSubstitutions *bool `json:"dynamicSubstitutions,omitempty"`

	// Option to specify whether reserved substitutions should be automapped.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.automap_substitutions
	AutomapSubstitutions *bool `json:"automapSubstitutions,omitempty"`

	// Option to define build log streaming behavior to Google Cloud
	//  Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.log_streaming_option
	LogStreamingOption *string `json:"logStreamingOption,omitempty"`

	// Optional. Specification for execution on a `WorkerPool`.
	//
	//  See [running builds in a private
	//  pool](https://cloud.google.com/build/docs/private-pools/run-builds-in-private-pool)
	//  for more information.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.pool
	Pool *BuildOptions_PoolOption `json:"pool,omitempty"`

	// Option to specify the logging mode, which determines if and where build
	//  logs are stored.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.logging
	Logging *string `json:"logging,omitempty"`

	// A list of global environment variable definitions that will exist for all
	//  build steps in this build. If a variable is defined in both this field and
	//  in a `BuildStep`, the value in `BuildStep` will be used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.env
	Env []string `json:"env,omitempty"`

	// A list of global environment variables, which are encrypted using a Cloud
	//  Key Management Service crypto key. These values must be specified in the
	//  build's `Secret`. These variables will be available to all build steps
	//  in this build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.secret_env
	SecretEnv []string `json:"secretEnv,omitempty"`

	// Global list of volumes to mount for ALL build steps
	//
	//  Each volume is created as an empty volume prior to starting the build
	//  process. Upon completion of the build, volumes and their contents are
	//  discarded. Global volume names was be unique among all volumes.
	//
	//  In addition, a volume will also be created for each normal build step
	//  under `/workspace` to serve as the common workspace for all steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// Optional. Option to specify how default logs buckets are localized.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.default_logs_bucket_behavior
	DefaultLogsBucketBehavior *string `json:"defaultLogsBucketBehavior,omitempty"`

	// Optional. Option to specify whether structured logging is enabled.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.enable_structured_logging
	EnableStructuredLogging *bool `json:"enableStructuredLogging,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildOptions.PoolOption
type BuildOptions_PoolOption struct {
	// The WorkerPool resource to execute the build on.
	//  You must have `cloudbuild.workerpools.use` on the project hosting the
	//  WorkerPool.
	//
	//  Format projects/{project}/locations/{location}/workerPools/{workerPool}
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.PoolOption.name
	NameRef *cloudbuildv1beta1.CloudBuildWorkerPoolRef `json:"nameRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Dependency.GitSourceRepository
type Dependency_GitSourceRepository struct {
	// Location of the Git repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Dependency.GitSourceRepository.url
	URL *string `json:"url,omitempty"`

	// The Developer Connect Git repository link or the url that matches a
	//  repository link in the current project, formatted as
	//  `projects/*/locations/*/connections/*/gitRepositoryLink/*`
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Dependency.GitSourceRepository.developer_connect
	DeveloperConnectRef *developerconnectv1alpha1.DeveloperConnectGitRepositoryLinkRef `json:"developerConnectRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitConfig.HttpConfig
type GitConfig_HTTPConfig struct {
	// SecretVersion resource of the HTTP proxy URL. The expected format is
	//  `projects/{project}/secrets/{secret}/versions/{version}`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitConfig.HttpConfig.proxy_secret_version_name
	ProxySecretVersionRef *secretmanagerv1beta1.SecretManagerSecretVersionRef `json:"proxySecretVersionRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.InlineSecret
type InlineSecret struct {
	// Resource name of Cloud KMS crypto key to decrypt the encrypted value.
	//  In format: projects/*/locations/*/keyRings/*/cryptoKeys/*
	// +kcc:proto:field=google.devtools.cloudbuild.v1.InlineSecret.kms_key_name
	KMSKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Source
type Build_Source struct {
	// If provided, get the source from this location in Cloud Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.storage_source
	StorageSource *Build_StorageSource `json:"storageSource,omitempty"`

	// If provided, get the source from this location in a Cloud Source
	//  Repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.repo_source
	RepoSource *Build_RepoSource `json:"repoSource,omitempty"`

	// If provided, get the source from this location in a Cloud Storage
	//  manifest.
	//
	//  This feature is in Preview.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.storage_source_manifest
	StorageSourceManifest *Build_StorageSourceManifest `json:"storageSourceManifest,omitempty"`

	// If provided, get the source from this Git repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.git_source
	GitSource *GitSource `json:"gitSource,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.RepoSource
type Build_RepoSource struct {
	// Optional. ID of the project that owns the Cloud Source Repository. If
	//  omitted, the project ID requesting the build is assumed.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.project_id
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Required. Name of the Cloud Source Repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.repo_name
	RepoName *string `json:"repoName,omitempty"`

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

	// Optional. Directory, relative to the source root, in which to run the build.
	//
	//  This must be a relative path. If a step's `dir` is specified and is an
	//  absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.dir
	Dir *string `json:"dir,omitempty"`

	// Optional. Only trigger a build if the revision regex does NOT match the
	//  revision regex.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`

	// Optional. Substitutions to use in a triggered build.
	//  Should only be used with RunBuildTrigger
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Secret
type Secret struct {
	// Cloud KMS key name to use to decrypt these envs.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Secret.kms_key_name
	KMSKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.SecretManagerSecret
type SecretManagerSecret struct {
	// Resource name of the SecretVersion. In format:
	//  projects/*/secrets/*/versions/*
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SecretManagerSecret.version_name
	VersionRef *secretmanagerv1beta1.SecretManagerSecretVersionRef `json:"versionRef,omitempty"`

	// Environment variable name to associate with the secret. Secret environment
	//  variables must be unique across all of a build's secrets, and must be used
	//  by at least one build step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SecretManagerSecret.env
	Env *string `json:"env,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.StorageSource
type Build_StorageSource struct {
	// Cloud Storage bucket.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Cloud Storage object containing the source.
	//
	//  This object must be a gzipped archive file (`.tar.gz`) containing source to
	//  build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.object
	Object *string `json:"object,omitempty"`

	// Cloud Storage generation for the object. If the generation is
	//  omitted, the latest generation will be used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.generation
	Generation *int64 `json:"generation,omitempty"`

	// Optional. Option to specify the source code source,
	//  currently used for Source-to-Image Builds.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.source_fetcher
	SourceFetcher *string `json:"sourceFetcher,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.StorageSourceManifest
type Build_StorageSourceManifest struct {
	// Cloud Storage bucket.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSourceManifest.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Cloud Storage object containing the source.
	//
	//  This object must be a gzipped archive file (`.tar.gz`) containing source to
	//  build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSourceManifest.object
	Object *string `json:"object,omitempty"`

	// Cloud Storage generation for the object. If the generation is
	//  omitted, the latest generation will be used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSourceManifest.generation
	Generation *int64 `json:"generation,omitempty"`
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
