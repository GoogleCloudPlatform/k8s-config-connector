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


// +kcc:proto=google.devtools.cloudbuild.v1.ApprovalConfig
type ApprovalConfig struct {
	// Whether or not approval is needed. If this is set on a build, it will
	//  become pending when created, and will need to be explicitly approved
	//  to start.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalConfig.approval_required
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.ApprovalResult
type ApprovalResult struct {

	// Required. The decision of this manual approval.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.decision
	Decision *string `json:"decision,omitempty"`

	// Optional. An optional comment for this manual approval result.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.comment
	Comment *string `json:"comment,omitempty"`

	// Optional. An optional URL tied to this manual approval result. This field
	//  is essentially the same as comment, except that it will be rendered by the
	//  UI differently. An example use case is a link to an external job that
	//  approved this Build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.url
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts
type Artifacts struct {
	// A list of images to be pushed upon the successful completion of all build
	//  steps.
	//
	//  The images will be pushed using the builder service account's credentials.
	//
	//  The digests of the pushed images will be stored in the Build resource's
	//  results field.
	//
	//  If any of the images fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.images
	Images []string `json:"images,omitempty"`

	// A list of objects to be uploaded to Cloud Storage upon successful
	//  completion of all build steps.
	//
	//  Files in the workspace matching specified paths globs will be uploaded to
	//  the specified Cloud Storage location using the builder service account's
	//  credentials.
	//
	//  The location and generation of the uploaded objects will be stored in the
	//  Build resource's results field.
	//
	//  If any objects fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.objects
	Objects *Artifacts_ArtifactObjects `json:"objects,omitempty"`

	// A list of Maven artifacts to be uploaded to Artifact Registry upon
	//  successful completion of all build steps.
	//
	//  Artifacts in the workspace matching specified paths globs will be uploaded
	//  to the specified Artifact Registry repository using the builder service
	//  account's credentials.
	//
	//  If any artifacts fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.maven_artifacts
	MavenArtifacts []Artifacts_MavenArtifact `json:"mavenArtifacts,omitempty"`

	// Optional. A list of Go modules to be uploaded to Artifact Registry upon
	//  successful completion of all build steps.
	//
	//  If any objects fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.go_modules
	GoModules []Artifacts_GoModule `json:"goModules,omitempty"`

	// A list of Python packages to be uploaded to Artifact Registry upon
	//  successful completion of all build steps.
	//
	//  The build service account credentials will be used to perform the upload.
	//
	//  If any objects fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.python_packages
	PythonPackages []Artifacts_PythonPackage `json:"pythonPackages,omitempty"`

	// A list of npm packages to be uploaded to Artifact Registry upon
	//  successful completion of all build steps.
	//
	//  Npm packages in the specified paths will be uploaded
	//  to the specified Artifact Registry repository using the builder service
	//  account's credentials.
	//
	//  If any packages fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.npm_packages
	NpmPackages []Artifacts_NpmPackage `json:"npmPackages,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects
type Artifacts_ArtifactObjects struct {
	// Cloud Storage bucket and optional object path, in the form
	//  "gs://bucket/path/to/somewhere/". (see [Bucket Name
	//  Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	//
	//  Files in the workspace matching any path pattern will be uploaded to
	//  Cloud Storage with this location as a prefix.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects.location
	Location *string `json:"location,omitempty"`

	// Path globs used to match files in the build's workspace.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects.paths
	Paths []string `json:"paths,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.GoModule
type Artifacts_GoModule struct {
	// Optional. Artifact Registry repository name.
	//
	//  Specified Go modules will be zipped and uploaded to Artifact Registry
	//  with this location as a prefix.
	//  e.g. my-go-repo
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.repository_name
	RepositoryName *string `json:"repositoryName,omitempty"`

	// Optional. Location of the Artifact Registry repository. i.e. us-east1
	//  Defaults to the build’s location.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.repository_location
	RepositoryLocation *string `json:"repositoryLocation,omitempty"`

	// Optional. Project ID of the Artifact Registry repository.
	//  Defaults to the build project.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.GoModule.repository_project_id
	RepositoryProjectID *string `json:"repositoryProjectID,omitempty"`

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

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact
type Artifacts_MavenArtifact struct {
	// Artifact Registry repository, in the form
	//  "https://$REGION-maven.pkg.dev/$PROJECT/$REPOSITORY"
	//
	//  Artifact in the workspace specified by path will be uploaded to
	//  Artifact Registry with this location as a prefix.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact.repository
	Repository *string `json:"repository,omitempty"`

	// Path to an artifact in the build's workspace to be uploaded to
	//  Artifact Registry.
	//  This can be either an absolute path,
	//  e.g. /workspace/my-app/target/my-app-1.0.SNAPSHOT.jar
	//  or a relative path from /workspace,
	//  e.g. my-app/target/my-app-1.0.SNAPSHOT.jar.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact.path
	Path *string `json:"path,omitempty"`

	// Maven `artifactId` value used when uploading the artifact to Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact.artifact_id
	ArtifactID *string `json:"artifactID,omitempty"`

	// Maven `groupId` value used when uploading the artifact to Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact.group_id
	GroupID *string `json:"groupID,omitempty"`

	// Maven `version` value used when uploading the artifact to Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.MavenArtifact.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.NpmPackage
type Artifacts_NpmPackage struct {
	// Artifact Registry repository, in the form
	//  "https://$REGION-npm.pkg.dev/$PROJECT/$REPOSITORY"
	//
	//  Npm package in the workspace specified by path will be zipped and
	//  uploaded to Artifact Registry with this location as a prefix.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.NpmPackage.repository
	Repository *string `json:"repository,omitempty"`

	// Path to the package.json.
	//  e.g. workspace/path/to/package
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.NpmPackage.package_path
	PackagePath *string `json:"packagePath,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.PythonPackage
type Artifacts_PythonPackage struct {
	// Artifact Registry repository, in the form
	//  "https://$REGION-python.pkg.dev/$PROJECT/$REPOSITORY"
	//
	//  Files in the workspace matching any path pattern will be uploaded to
	//  Artifact Registry with this location as a prefix.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.PythonPackage.repository
	Repository *string `json:"repository,omitempty"`

	// Path globs used to match files in the build's workspace. For Python/
	//  Twine, this is usually `dist/*`, and sometimes additionally an `.asc`
	//  file.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.PythonPackage.paths
	Paths []string `json:"paths,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build
type Build struct {

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
	QueueTtl *string `json:"queueTtl,omitempty"`

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
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build.FailureInfo
type Build_FailureInfo struct {
	// The name of the failure.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.FailureInfo.type
	Type *string `json:"type,omitempty"`

	// Explains the failure issue in more detail using hard-coded text.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.FailureInfo.detail
	Detail *string `json:"detail,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build.Warning
type Build_Warning struct {
	// Explanation of the warning generated.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.Warning.text
	Text *string `json:"text,omitempty"`

	// The priority for this warning.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.Warning.priority
	Priority *string `json:"priority,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildApproval
type BuildApproval struct {
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
	//  allocated for the build -- the build may run with a larger disk than
	//  requested. At present, the maximum disk size is 2000GB; builds that request
	//  more than the maximum are rejected with an error.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`

	// Option to specify behavior when there is an error in the substitution
	//  checks.
	//
	//  NOTE: this is always set to ALLOW_LOOSE for triggered builds and cannot
	//  be overridden in the build configuration file.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.substitution_option
	SubstitutionOption *string `json:"substitutionOption,omitempty"`

	// Option to specify whether or not to apply bash style string
	//  operations to the substitutions.
	//
	//  NOTE: this is always enabled for triggered builds and cannot be
	//  overridden in the build configuration file.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.dynamic_substitutions
	DynamicSubstitutions *bool `json:"dynamicSubstitutions,omitempty"`

	// Option to include built-in and custom substitutions as env variables
	//  for all build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.automap_substitutions
	AutomapSubstitutions *bool `json:"automapSubstitutions,omitempty"`

	// Option to define build log streaming behavior to Cloud
	//  Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.log_streaming_option
	LogStreamingOption *string `json:"logStreamingOption,omitempty"`

	// This field deprecated; please use `pool.name` instead.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

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
	//  build steps in this build. If a variable is defined in both globally and in
	//  a build step, the variable will use the build step value.
	//
	//  The elements are of the form "KEY=VALUE" for the environment variable "KEY"
	//  being given the value "VALUE".
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
	//  discarded. Global volume names and paths cannot conflict with the volumes
	//  defined a build step.
	//
	//  Using a global volume in a build with only one step is not valid as
	//  it is indicative of a build request with an incorrect configuration.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// Optional. Option to specify how default logs buckets are setup.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.default_logs_bucket_behavior
	DefaultLogsBucketBehavior *string `json:"defaultLogsBucketBehavior,omitempty"`

	// Optional. Option to specify whether structured logging is enabled.
	//
	//  If true, JSON-formatted logs are parsed as structured logs.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.enable_structured_logging
	EnableStructuredLogging *bool `json:"enableStructuredLogging,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildOptions.PoolOption
type BuildOptions_PoolOption struct {
	// The `WorkerPool` resource to execute the build on.
	//  You must have `cloudbuild.workerpools.use` on the project hosting the
	//  WorkerPool.
	//
	//  Format projects/{project}/locations/{location}/workerPools/{workerPoolId}
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildOptions.PoolOption.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildStep
type BuildStep struct {
	// Required. The name of the container image that will run this particular
	//  build step.
	//
	//  If the image is available in the host's Docker daemon's cache, it
	//  will be run directly. If not, the host will attempt to pull the image
	//  first, using the builder service account's credentials if necessary.
	//
	//  The Docker daemon's cache will already have the latest versions of all of
	//  the officially supported build steps
	//  ([https://github.com/GoogleCloudPlatform/cloud-builders](https://github.com/GoogleCloudPlatform/cloud-builders)).
	//  The Docker daemon will also have cached many of the layers for some popular
	//  images, like "ubuntu", "debian", but they will be refreshed at the time you
	//  attempt to use them.
	//
	//  If you built an image in a previous build step, it will be stored in the
	//  host's Docker daemon's cache and is available to use as the name for a
	//  later build step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.name
	Name *string `json:"name,omitempty"`

	// A list of environment variable definitions to be used when running a step.
	//
	//  The elements are of the form "KEY=VALUE" for the environment variable "KEY"
	//  being given the value "VALUE".
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.env
	Env []string `json:"env,omitempty"`

	// A list of arguments that will be presented to the step when it is started.
	//
	//  If the image used to run the step's container has an entrypoint, the `args`
	//  are used as arguments to that entrypoint. If the image does not define
	//  an entrypoint, the first element in args is used as the entrypoint,
	//  and the remainder will be used as arguments.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.args
	Args []string `json:"args,omitempty"`

	// Working directory to use when running this step's container.
	//
	//  If this value is a relative path, it is relative to the build's working
	//  directory. If this value is absolute, it may be outside the build's working
	//  directory, in which case the contents of the path may not be persisted
	//  across build step executions, unless a `volume` for that path is specified.
	//
	//  If the build specifies a `RepoSource` with `dir` and a step with a `dir`,
	//  which specifies an absolute path, the `RepoSource` `dir` is ignored for
	//  the step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.dir
	Dir *string `json:"dir,omitempty"`

	// Unique identifier for this build step, used in `wait_for` to
	//  reference this build step as a dependency.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.id
	ID *string `json:"id,omitempty"`

	// The ID(s) of the step(s) that this build step depends on.
	//  This build step will not start until all the build steps in `wait_for`
	//  have completed successfully. If `wait_for` is empty, this build step will
	//  start when all previous build steps in the `Build.Steps` list have
	//  completed successfully.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.wait_for
	WaitFor []string `json:"waitFor,omitempty"`

	// Entrypoint to be used instead of the build step image's default entrypoint.
	//  If unset, the image's default entrypoint is used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.entrypoint
	Entrypoint *string `json:"entrypoint,omitempty"`

	// A list of environment variables which are encrypted using a Cloud Key
	//  Management Service crypto key. These values must be specified in the
	//  build's `Secret`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.secret_env
	SecretEnv []string `json:"secretEnv,omitempty"`

	// List of volumes to mount into the build step.
	//
	//  Each volume is created as an empty volume prior to execution of the
	//  build step. Upon completion of the build, volumes and their contents are
	//  discarded.
	//
	//  Using a named volume in only one step is not valid as it is indicative
	//  of a build request with an incorrect configuration.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// Time limit for executing this build step. If not defined, the step has no
	//  time limit and will be allowed to continue to run until either it completes
	//  or the build itself times out.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Allow this build step to fail without failing the entire build.
	//
	//  If false, the entire build will fail if this step fails. Otherwise, the
	//  build will succeed, but this step will still have a failure status.
	//  Error information will be reported in the failure_detail field.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.allow_failure
	AllowFailure *bool `json:"allowFailure,omitempty"`

	// Allow this build step to fail without failing the entire build if and
	//  only if the exit code is one of the specified codes. If allow_failure
	//  is also specified, this field will take precedence.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.allow_exit_codes
	AllowExitCodes []int32 `json:"allowExitCodes,omitempty"`

	// A shell script to be executed in the step.
	//
	//  When script is provided, the user cannot specify the entrypoint or args.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.script
	Script *string `json:"script,omitempty"`

	// Option to include built-in and custom substitutions as env variables
	//  for this build step. This option will override the global option
	//  in BuildOption.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.automap_substitutions
	AutomapSubstitutions *bool `json:"automapSubstitutions,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuiltImage
type BuiltImage struct {
	// Name used to push the container image to Google Container Registry, as
	//  presented to `docker push`.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.name
	Name *string `json:"name,omitempty"`

	// Docker Registry 2.0 digest.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.digest
	Digest *string `json:"digest,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.FileHashes
type FileHashes struct {
	// Collection of file hashes.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.FileHashes.file_hash
	FileHash []Hash `json:"fileHash,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitSource
type GitSource struct {
	// Location of the Git repo to build.
	//
	//  This will be used as a `git remote`, see
	//  https://git-scm.com/docs/git-remote.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitSource.url
	URL *string `json:"url,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	//
	//  This must be a relative path. If a step's `dir` is specified and is an
	//  absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitSource.dir
	Dir *string `json:"dir,omitempty"`

	// The revision to fetch from the Git repository such as a branch, a tag, a
	//  commit SHA, or any Git ref.
	//
	//  Cloud Build uses `git fetch` to fetch the revision from the Git
	//  repository; therefore make sure that the string you provide for `revision`
	//  is parsable  by the command. For information on string values accepted by
	//  `git fetch`, see
	//  https://git-scm.com/docs/gitrevisions#_specifying_revisions. For
	//  information on `git fetch`, see https://git-scm.com/docs/git-fetch.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitSource.revision
	Revision *string `json:"revision,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Hash
type Hash struct {
	// The type of hash that was performed.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Hash.type
	Type *string `json:"type,omitempty"`

	// The hash value.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Hash.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.InlineSecret
type InlineSecret struct {
	// Resource name of Cloud KMS crypto key to decrypt the encrypted value.
	//  In format: projects/*/locations/*/keyRings/*/cryptoKeys/*
	// +kcc:proto:field=google.devtools.cloudbuild.v1.InlineSecret.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// TODO: unsupported map type with key string and value bytes

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

	// Substitutions to use in a triggered build.
	//  Should only be used with RunBuildTrigger
	// +kcc:proto:field=google.devtools.cloudbuild.v1.RepoSource.substitutions
	Substitutions map[string]string `json:"substitutions,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Results
type Results struct {
	// Container images that were built as a part of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.images
	Images []BuiltImage `json:"images,omitempty"`

	// List of build step digests, in the order corresponding to build step
	//  indices.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.build_step_images
	BuildStepImages []string `json:"buildStepImages,omitempty"`

	// Path to the artifact manifest for non-container artifacts uploaded to Cloud
	//  Storage. Only populated when artifacts are uploaded to Cloud Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.artifact_manifest
	ArtifactManifest *string `json:"artifactManifest,omitempty"`

	// Number of non-container artifacts uploaded to Cloud Storage. Only populated
	//  when artifacts are uploaded to Cloud Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.num_artifacts
	NumArtifacts *int64 `json:"numArtifacts,omitempty"`

	// List of build step outputs, produced by builder images, in the order
	//  corresponding to build step indices.
	//
	//  [Cloud Builders](https://cloud.google.com/cloud-build/docs/cloud-builders)
	//  can produce this output by writing to `$BUILDER_OUTPUT/output`.
	//  Only the first 4KB of data is stored.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.build_step_outputs
	BuildStepOutputs [][]byte `json:"buildStepOutputs,omitempty"`

	// Time to push all non-container artifacts to Cloud Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.artifact_timing
	ArtifactTiming *TimeSpan `json:"artifactTiming,omitempty"`

	// Python artifacts uploaded to Artifact Registry at the end of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.python_packages
	PythonPackages []UploadedPythonPackage `json:"pythonPackages,omitempty"`

	// Maven artifacts uploaded to Artifact Registry at the end of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.maven_artifacts
	MavenArtifacts []UploadedMavenArtifact `json:"mavenArtifacts,omitempty"`

	// Optional. Go module artifacts uploaded to Artifact Registry at the end of
	//  the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.go_modules
	GoModules []UploadedGoModule `json:"goModules,omitempty"`

	// Npm packages uploaded to Artifact Registry at the end of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.npm_packages
	NpmPackages []UploadedNpmPackage `json:"npmPackages,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Secret
type Secret struct {
	// Cloud KMS key name to use to decrypt these envs.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Secret.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// TODO: unsupported map type with key string and value bytes

}

// +kcc:proto=google.devtools.cloudbuild.v1.SecretManagerSecret
type SecretManagerSecret struct {
	// Resource name of the SecretVersion. In format:
	//  projects/*/secrets/*/versions/*
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SecretManagerSecret.version_name
	VersionName *string `json:"versionName,omitempty"`

	// Environment variable name to associate with the secret.
	//  Secret environment variables must be unique across all of a build's
	//  secrets, and must be used by at least one build step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SecretManagerSecret.env
	Env *string `json:"env,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Secrets
type Secrets struct {
	// Secrets in Secret Manager and associated secret environment variable.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Secrets.secret_manager
	SecretManager []SecretManagerSecret `json:"secretManager,omitempty"`

	// Secrets encrypted with KMS key and the associated secret environment
	//  variable.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Secrets.inline
	Inline []InlineSecret `json:"inline,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Source
type Source struct {
	// If provided, get the source from this location in Cloud Storage.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.storage_source
	StorageSource *StorageSource `json:"storageSource,omitempty"`

	// If provided, get the source from this location in a Cloud Source
	//  Repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.repo_source
	RepoSource *RepoSource `json:"repoSource,omitempty"`

	// If provided, get the source from this Git repository.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.git_source
	GitSource *GitSource `json:"gitSource,omitempty"`

	// If provided, get the source from this manifest in Cloud Storage.
	//  This feature is in Preview; see description
	//  [here](https://github.com/GoogleCloudPlatform/cloud-builders/tree/master/gcs-fetcher).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Source.storage_source_manifest
	StorageSourceManifest *StorageSourceManifest `json:"storageSourceManifest,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.SourceProvenance
type SourceProvenance struct {
	// A copy of the build's `source.storage_source`, if exists, with any
	//  generations resolved.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SourceProvenance.resolved_storage_source
	ResolvedStorageSource *StorageSource `json:"resolvedStorageSource,omitempty"`

	// A copy of the build's `source.repo_source`, if exists, with any
	//  revisions resolved.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SourceProvenance.resolved_repo_source
	ResolvedRepoSource *RepoSource `json:"resolvedRepoSource,omitempty"`

	// A copy of the build's `source.storage_source_manifest`, if exists, with any
	//  revisions resolved.
	//  This feature is in Preview.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.SourceProvenance.resolved_storage_source_manifest
	ResolvedStorageSourceManifest *StorageSourceManifest `json:"resolvedStorageSourceManifest,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.StorageSource
type StorageSource struct {
	// Cloud Storage bucket containing the source (see
	//  [Bucket Name
	//  Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Cloud Storage object containing the source.
	//
	//  This object must be a zipped (`.zip`) or gzipped archive file (`.tar.gz`)
	//  containing source to build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.object
	Object *string `json:"object,omitempty"`

	// Cloud Storage generation for the object. If the generation is
	//  omitted, the latest generation will be used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.generation
	Generation *int64 `json:"generation,omitempty"`

	// Option to specify the tool to fetch the source file for the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSource.source_fetcher
	SourceFetcher *string `json:"sourceFetcher,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.StorageSourceManifest
type StorageSourceManifest struct {
	// Cloud Storage bucket containing the source manifest (see [Bucket
	//  Name
	//  Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSourceManifest.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Cloud Storage object containing the source manifest.
	//
	//  This object must be a JSON file.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSourceManifest.object
	Object *string `json:"object,omitempty"`

	// Cloud Storage generation for the object. If the generation is
	//  omitted, the latest generation will be used.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.StorageSourceManifest.generation
	Generation *int64 `json:"generation,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.TimeSpan
type TimeSpan struct {
	// Start of time span.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.TimeSpan.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End of time span.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.TimeSpan.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedGoModule
type UploadedGoModule struct {
	// URI of the uploaded artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedGoModule.uri
	URI *string `json:"uri,omitempty"`

	// Hash types and values of the Go Module Artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedGoModule.file_hashes
	FileHashes *FileHashes `json:"fileHashes,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedMavenArtifact
type UploadedMavenArtifact struct {
	// URI of the uploaded artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedMavenArtifact.uri
	URI *string `json:"uri,omitempty"`

	// Hash types and values of the Maven Artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedMavenArtifact.file_hashes
	FileHashes *FileHashes `json:"fileHashes,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedNpmPackage
type UploadedNpmPackage struct {
	// URI of the uploaded npm package.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedNpmPackage.uri
	URI *string `json:"uri,omitempty"`

	// Hash types and values of the npm package.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedNpmPackage.file_hashes
	FileHashes *FileHashes `json:"fileHashes,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedPythonPackage
type UploadedPythonPackage struct {
	// URI of the uploaded artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedPythonPackage.uri
	URI *string `json:"uri,omitempty"`

	// Hash types and values of the Python Artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedPythonPackage.file_hashes
	FileHashes *FileHashes `json:"fileHashes,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Volume
type Volume struct {
	// Name of the volume to mount.
	//
	//  Volume names must be unique per build step and must be valid names for
	//  Docker volumes. Each named volume must be used by at least two build steps.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Volume.name
	Name *string `json:"name,omitempty"`

	// Path at which to mount the volume.
	//
	//  Paths must be absolute and cannot conflict with other volume paths on the
	//  same build step or with certain reserved volume paths.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Volume.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.ApprovalResult
type ApprovalResultObservedState struct {
	// Output only. Email of the user that called the ApproveBuild API to
	//  approve or reject a build at the time that the API was called.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.approver_account
	ApproverAccount *string `json:"approverAccount,omitempty"`

	// Output only. The time when the approval decision was made.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.ApprovalResult.approval_time
	ApprovalTime *string `json:"approvalTime,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts
type ArtifactsObservedState struct {
	// A list of objects to be uploaded to Cloud Storage upon successful
	//  completion of all build steps.
	//
	//  Files in the workspace matching specified paths globs will be uploaded to
	//  the specified Cloud Storage location using the builder service account's
	//  credentials.
	//
	//  The location and generation of the uploaded objects will be stored in the
	//  Build resource's results field.
	//
	//  If any objects fail to be pushed, the build is marked FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.objects
	Objects *Artifacts_ArtifactObjectsObservedState `json:"objects,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects
type Artifacts_ArtifactObjectsObservedState struct {
	// Output only. Stores timing information for pushing all artifact objects.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Artifacts.ArtifactObjects.timing
	Timing *TimeSpan `json:"timing,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Build
type BuildObservedState struct {
	// Output only. The 'Build' name with format:
	//  `projects/{project}/locations/{location}/builds/{build}`, where {build}
	//  is a unique identifier generated by the service.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.name
	Name *string `json:"name,omitempty"`

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
	Results *Results `json:"results,omitempty"`

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
	SourceProvenance *SourceProvenance `json:"sourceProvenance,omitempty"`

	// Output only. The ID of the `BuildTrigger` that triggered this build, if it
	//  was triggered automatically.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.build_trigger_id
	BuildTriggerID *string `json:"buildTriggerID,omitempty"`

	// Output only. URL to logs for this build in Google Cloud Console.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.log_url
	LogURL *string `json:"logURL,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. Describes this build's approval configuration, status,
	//  and result.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.approval
	Approval *BuildApproval `json:"approval,omitempty"`

	// Output only. Non-fatal problems encountered during the execution of the
	//  build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.warnings
	Warnings []Build_Warning `json:"warnings,omitempty"`

	// Output only. Contains information about the build when status=FAILURE.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Build.failure_info
	FailureInfo *Build_FailureInfo `json:"failureInfo,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildApproval
type BuildApprovalObservedState struct {
	// Output only. The state of this build's approval.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.state
	State *string `json:"state,omitempty"`

	// Output only. Configuration for manual approval of this build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.config
	Config *ApprovalConfig `json:"config,omitempty"`

	// Output only. Result of manual approval for this Build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildApproval.result
	Result *ApprovalResult `json:"result,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuildStep
type BuildStepObservedState struct {
	// Output only. Stores timing information for executing this build step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.timing
	Timing *TimeSpan `json:"timing,omitempty"`

	// Output only. Stores timing information for pulling this build step's
	//  builder image only.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.pull_timing
	PullTiming *TimeSpan `json:"pullTiming,omitempty"`

	// Output only. Status of the build step. At this time, build step status is
	//  only updated on build completion; step status is not updated in real-time
	//  as the build progresses.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.status
	Status *string `json:"status,omitempty"`

	// Output only. Return code from running the step.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuildStep.exit_code
	ExitCode *int32 `json:"exitCode,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.BuiltImage
type BuiltImageObservedState struct {
	// Output only. Stores timing information for pushing the specified image.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.BuiltImage.push_timing
	PushTiming *TimeSpan `json:"pushTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.Results
type ResultsObservedState struct {
	// Container images that were built as a part of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.images
	Images []BuiltImageObservedState `json:"images,omitempty"`

	// Python artifacts uploaded to Artifact Registry at the end of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.python_packages
	PythonPackages []UploadedPythonPackageObservedState `json:"pythonPackages,omitempty"`

	// Maven artifacts uploaded to Artifact Registry at the end of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.maven_artifacts
	MavenArtifacts []UploadedMavenArtifactObservedState `json:"mavenArtifacts,omitempty"`

	// Optional. Go module artifacts uploaded to Artifact Registry at the end of
	//  the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.go_modules
	GoModules []UploadedGoModuleObservedState `json:"goModules,omitempty"`

	// Npm packages uploaded to Artifact Registry at the end of the build.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.Results.npm_packages
	NpmPackages []UploadedNpmPackageObservedState `json:"npmPackages,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.SourceProvenance
type SourceProvenanceObservedState struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedGoModule
type UploadedGoModuleObservedState struct {
	// Output only. Stores timing information for pushing the specified artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedGoModule.push_timing
	PushTiming *TimeSpan `json:"pushTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedMavenArtifact
type UploadedMavenArtifactObservedState struct {
	// Output only. Stores timing information for pushing the specified artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedMavenArtifact.push_timing
	PushTiming *TimeSpan `json:"pushTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedNpmPackage
type UploadedNpmPackageObservedState struct {
	// Output only. Stores timing information for pushing the specified artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedNpmPackage.push_timing
	PushTiming *TimeSpan `json:"pushTiming,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.UploadedPythonPackage
type UploadedPythonPackageObservedState struct {
	// Output only. Stores timing information for pushing the specified artifact.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.UploadedPythonPackage.push_timing
	PushTiming *TimeSpan `json:"pushTiming,omitempty"`
}
