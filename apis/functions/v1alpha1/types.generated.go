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


// +kcc:proto=google.cloud.functions.v2.AutomaticUpdatePolicy
type AutomaticUpdatePolicy struct {
}

// +kcc:proto=google.cloud.functions.v2.BuildConfig
type BuildConfig struct {
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.automatic_update_policy
	AutomaticUpdatePolicy *AutomaticUpdatePolicy `json:"automaticUpdatePolicy,omitempty"`

	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.on_deploy_update_policy
	OnDeployUpdatePolicy *OnDeployUpdatePolicy `json:"onDeployUpdatePolicy,omitempty"`

	// The runtime in which to run the function. Required when deploying a new
	//  function, optional when updating an existing function. For a complete
	//  list of possible choices, see the
	//  [`gcloud` command
	//  reference](https://cloud.google.com/sdk/gcloud/reference/functions/deploy#--runtime).
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.runtime
	Runtime *string `json:"runtime,omitempty"`

	// The name of the function (as defined in source code) that will be
	//  executed. Defaults to the resource name suffix, if not specified. For
	//  backward compatibility, if function with given name is not found, then the
	//  system will try to use function named "function".
	//  For Node.js this is name of a function exported by the module specified
	//  in `source_location`.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.entry_point
	EntryPoint *string `json:"entryPoint,omitempty"`

	// The location of the function source code.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.source
	Source *Source `json:"source,omitempty"`

	// Name of the Cloud Build Custom Worker Pool that should be used to build the
	//  function. The format of this field is
	//  `projects/{project}/locations/{region}/workerPools/{workerPool}` where
	//  {project} and {region} are the project id and region respectively where the
	//  worker pool is defined and {workerPool} is the short name of the worker
	//  pool.
	//
	//  If the project id is not the same as the function, then the Cloud
	//  Functions Service Agent
	//  (service-<project_number>@gcf-admin-robot.iam.gserviceaccount.com) must be
	//  granted the role Cloud Build Custom Workers Builder
	//  (roles/cloudbuild.customworkers.builder) in the project.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// User-provided build-time environment variables for the function
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.environment_variables
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	// Docker Registry to use for this deployment. This configuration is only
	//  applicable to 1st Gen functions, 2nd Gen functions can only use Artifact
	//  Registry.
	//
	//  If unspecified, it defaults to `ARTIFACT_REGISTRY`.
	//  If `docker_repository` field is specified, this field should either be left
	//  unspecified or set to `ARTIFACT_REGISTRY`.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.docker_registry
	DockerRegistry *string `json:"dockerRegistry,omitempty"`

	// Repository in Artifact Registry to which the function docker image will be
	//  pushed after it is built by Cloud Build. If specified by user, it is
	//  created and managed by user with a customer managed encryption key.
	//  Otherwise, GCF will create and use a repository named 'gcf-artifacts'
	//  for every deployed region.
	//
	//  It must match the pattern
	//  `projects/{project}/locations/{location}/repositories/{repository}`.
	//
	//  Cross-project repositories are not supported.
	//  Cross-location repositories are not supported.
	//  Repository format must be 'DOCKER'.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.docker_repository
	DockerRepository *string `json:"dockerRepository,omitempty"`

	// Service account to be used for building the container. The format of this
	//  field is `projects/{projectId}/serviceAccounts/{serviceAccountEmail}`.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.EventFilter
type EventFilter struct {
	// Required. The name of a CloudEvents attribute.
	// +kcc:proto:field=google.cloud.functions.v2.EventFilter.attribute
	Attribute *string `json:"attribute,omitempty"`

	// Required. The value for the attribute.
	// +kcc:proto:field=google.cloud.functions.v2.EventFilter.value
	Value *string `json:"value,omitempty"`

	// Optional. The operator used for matching the events with the value of the
	//  filter. If not specified, only events that have an exact key-value pair
	//  specified in the filter are matched. The only allowed value is
	//  `match-path-pattern`.
	// +kcc:proto:field=google.cloud.functions.v2.EventFilter.operator
	Operator *string `json:"operator,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.EventTrigger
type EventTrigger struct {

	// The region that the trigger will be in. The trigger will only receive
	//  events originating in this region. It can be the same
	//  region as the function, a different region or multi-region, or the global
	//  region. If not provided, defaults to the same region as the function.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.trigger_region
	TriggerRegion *string `json:"triggerRegion,omitempty"`

	// Required. The type of event to observe. For example:
	//  `google.cloud.audit.log.v1.written` or
	//  `google.cloud.pubsub.topic.v1.messagePublished`.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.event_type
	EventType *string `json:"eventType,omitempty"`

	// Criteria used to filter events.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.event_filters
	EventFilters []EventFilter `json:"eventFilters,omitempty"`

	// Optional. The name of a Pub/Sub topic in the same project that will be used
	//  as the transport topic for the event delivery. Format:
	//  `projects/{project}/topics/{topic}`.
	//
	//  This is only valid for events of type
	//  `google.cloud.pubsub.topic.v1.messagePublished`. The topic provided here
	//  will not be deleted at function deletion.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.pubsub_topic
	PubsubTopic *string `json:"pubsubTopic,omitempty"`

	// Optional. The email of the trigger's service account. The service account
	//  must have permission to invoke Cloud Run services, the permission is
	//  `run.routes.invoke`.
	//  If empty, defaults to the Compute Engine default service account:
	//  `{project_number}-compute@developer.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Optional. If unset, then defaults to ignoring failures (i.e. not retrying
	//  them).
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.retry_policy
	RetryPolicy *string `json:"retryPolicy,omitempty"`

	// Optional. The name of the channel associated with the trigger in
	//  `projects/{project}/locations/{location}/channels/{channel}` format.
	//  You must provide a channel to receive events from Eventarc SaaS partners.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.channel
	Channel *string `json:"channel,omitempty"`

	// Optional. The hostname of the service that 1st Gen function should be
	//  observed.
	//
	//  If no string is provided, the default service implementing the API will
	//  be used. For example, `storage.googleapis.com` is the default for all
	//  event types in the `google.storage` namespace.
	//
	//  The field is only applicable to 1st Gen functions.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.service
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.Function
type Function struct {
	// A user-defined name of the function. Function names must be unique
	//  globally and match pattern `projects/*/locations/*/functions/*`
	// +kcc:proto:field=google.cloud.functions.v2.Function.name
	Name *string `json:"name,omitempty"`

	// User-provided description of a function.
	// +kcc:proto:field=google.cloud.functions.v2.Function.description
	Description *string `json:"description,omitempty"`

	// Describes the Build step of the function that builds a container from the
	//  given source.
	// +kcc:proto:field=google.cloud.functions.v2.Function.build_config
	BuildConfig *BuildConfig `json:"buildConfig,omitempty"`

	// Describes the Service being deployed. Currently deploys services to Cloud
	//  Run (fully managed).
	// +kcc:proto:field=google.cloud.functions.v2.Function.service_config
	ServiceConfig *ServiceConfig `json:"serviceConfig,omitempty"`

	// An Eventarc trigger managed by Google Cloud Functions that fires events in
	//  response to a condition in another service.
	// +kcc:proto:field=google.cloud.functions.v2.Function.event_trigger
	EventTrigger *EventTrigger `json:"eventTrigger,omitempty"`

	// Labels associated with this Cloud Function.
	// +kcc:proto:field=google.cloud.functions.v2.Function.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Describe whether the function is 1st Gen or 2nd Gen.
	// +kcc:proto:field=google.cloud.functions.v2.Function.environment
	Environment *string `json:"environment,omitempty"`

	// [Preview] Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt function resources.
	//
	//  It must match the pattern
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.functions.v2.Function.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.OnDeployUpdatePolicy
type OnDeployUpdatePolicy struct {
}

// +kcc:proto=google.cloud.functions.v2.RepoSource
type RepoSource struct {
	// Regex matching branches to build.
	//
	//  The syntax of the regular expressions accepted is the syntax accepted by
	//  RE2 and described at https://github.com/google/re2/wiki/Syntax
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.branch_name
	BranchName *string `json:"branchName,omitempty"`

	// Regex matching tags to build.
	//
	//  The syntax of the regular expressions accepted is the syntax accepted by
	//  RE2 and described at https://github.com/google/re2/wiki/Syntax
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.tag_name
	TagName *string `json:"tagName,omitempty"`

	// Explicit commit SHA to build.
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`

	// ID of the project that owns the Cloud Source Repository. If omitted, the
	//  project ID requesting the build is assumed.
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Name of the Cloud Source Repository.
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.repo_name
	RepoName *string `json:"repoName,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	//
	//  This must be a relative path. If a step's `dir` is specified and is an
	//  absolute path, this value is ignored for that step's execution.
	//  eg. helloworld (no leading slash allowed)
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.dir
	Dir *string `json:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision
	//  regex.
	// +kcc:proto:field=google.cloud.functions.v2.RepoSource.invert_regex
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.SecretEnvVar
type SecretEnvVar struct {
	// Name of the environment variable.
	// +kcc:proto:field=google.cloud.functions.v2.SecretEnvVar.key
	Key *string `json:"key,omitempty"`

	// Project identifier (preferably project number but can also be the
	//  project ID) of the project that contains the secret. If not set, it is
	//  assumed that the secret is in the same project as the function.
	// +kcc:proto:field=google.cloud.functions.v2.SecretEnvVar.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	// +kcc:proto:field=google.cloud.functions.v2.SecretEnvVar.secret
	Secret *string `json:"secret,omitempty"`

	// Version of the secret (version number or the string 'latest'). It is
	//  recommended to use a numeric version for secret environment variables as
	//  any updates to the secret value is not reflected until new instances
	//  start.
	// +kcc:proto:field=google.cloud.functions.v2.SecretEnvVar.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.SecretVolume
type SecretVolume struct {
	// The path within the container to mount the secret volume. For example,
	//  setting the mount_path as `/etc/secrets` would mount the secret value files
	//  under the `/etc/secrets` directory. This directory will also be completely
	//  shadowed and unavailable to mount any other secrets.
	//  Recommended mount path: /etc/secrets
	// +kcc:proto:field=google.cloud.functions.v2.SecretVolume.mount_path
	MountPath *string `json:"mountPath,omitempty"`

	// Project identifier (preferably project number but can also be the project
	//  ID) of the project that contains the secret. If not set, it is
	//  assumed that the secret is in the same project as the function.
	// +kcc:proto:field=google.cloud.functions.v2.SecretVolume.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	// +kcc:proto:field=google.cloud.functions.v2.SecretVolume.secret
	Secret *string `json:"secret,omitempty"`

	// List of secret versions to mount for this secret. If empty, the `latest`
	//  version of the secret will be made available in a file named after the
	//  secret under the mount point.
	// +kcc:proto:field=google.cloud.functions.v2.SecretVolume.versions
	Versions []SecretVolume_SecretVersion `json:"versions,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.SecretVolume.SecretVersion
type SecretVolume_SecretVersion struct {
	// Version of the secret (version number or the string 'latest'). It is
	//  preferable to use `latest` version with secret volumes as secret value
	//  changes are reflected immediately.
	// +kcc:proto:field=google.cloud.functions.v2.SecretVolume.SecretVersion.version
	Version *string `json:"version,omitempty"`

	// Relative path of the file under the mount path where the secret value for
	//  this version will be fetched and made available. For example, setting the
	//  mount_path as '/etc/secrets' and path as `secret_foo` would mount the
	//  secret value file at `/etc/secrets/secret_foo`.
	// +kcc:proto:field=google.cloud.functions.v2.SecretVolume.SecretVersion.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.ServiceConfig
type ServiceConfig struct {

	// The function execution timeout. Execution is considered failed and
	//  can be terminated if the function is not completed at the end of the
	//  timeout period. Defaults to 60 seconds.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.timeout_seconds
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`

	// The amount of memory available for a function.
	//  Defaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is
	//  supplied the value is interpreted as bytes.
	//  See
	//  https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//  a full description.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.available_memory
	AvailableMemory *string `json:"availableMemory,omitempty"`

	// The number of CPUs used in a single container instance.
	//  Default value is calculated from available memory.
	//  Supports the same values as Cloud Run, see
	//  https://cloud.google.com/run/docs/reference/rest/v1/Container#resourcerequirements
	//  Example: "1" indicates 1 vCPU
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.available_cpu
	AvailableCpu *string `json:"availableCpu,omitempty"`

	// Environment variables that shall be available during function execution.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.environment_variables
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a
	//  given time.
	//
	//  In some cases, such as rapid traffic surges, Cloud Functions may, for a
	//  short period of time, create more instances than the specified max
	//  instances limit. If your function cannot tolerate this temporary behavior,
	//  you may want to factor in a safety margin and set a lower max instances
	//  value than your function can tolerate.
	//
	//  See the [Max
	//  Instances](https://cloud.google.com/functions/docs/max-instances) Guide for
	//  more details.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.max_instance_count
	MaxInstanceCount *int32 `json:"maxInstanceCount,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a
	//  given time.
	//
	//  Function instances are kept in idle state for a short period after they
	//  finished executing the request to reduce cold start time for subsequent
	//  requests. Setting a minimum instance count will ensure that the given
	//  number of instances are kept running in idle state always. This can help
	//  with cold start times when jump in incoming request count occurs after the
	//  idle instance would have been stopped in the default case.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.min_instance_count
	MinInstanceCount *int32 `json:"minInstanceCount,omitempty"`

	// The Serverless VPC Access connector that this cloud function can connect
	//  to. The format of this field is `projects/*/locations/*/connectors/*`.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.vpc_connector
	VpcConnector *string `json:"vpcConnector,omitempty"`

	// The egress settings for the connector, controlling what traffic is diverted
	//  through it.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.vpc_connector_egress_settings
	VpcConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty"`

	// The ingress settings for the function, controlling what traffic can reach
	//  it.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.ingress_settings
	IngressSettings *string `json:"ingressSettings,omitempty"`

	// The email of the service's service account. If empty, defaults to
	//  `{project_number}-compute@developer.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Whether 100% of traffic is routed to the latest revision.
	//  On CreateFunction and UpdateFunction, when set to true, the revision being
	//  deployed will serve 100% of traffic, ignoring any traffic split settings,
	//  if any. On GetFunction, true will be returned if the latest revision is
	//  serving 100% of traffic.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.all_traffic_on_latest_revision
	AllTrafficOnLatestRevision *bool `json:"allTrafficOnLatestRevision,omitempty"`

	// Secret environment variables configuration.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.secret_environment_variables
	SecretEnvironmentVariables []SecretEnvVar `json:"secretEnvironmentVariables,omitempty"`

	// Secret volumes configuration.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.secret_volumes
	SecretVolumes []SecretVolume `json:"secretVolumes,omitempty"`

	// Sets the maximum number of concurrent requests that each instance
	//  can receive. Defaults to 1.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.max_instance_request_concurrency
	MaxInstanceRequestConcurrency *int32 `json:"maxInstanceRequestConcurrency,omitempty"`

	// Security level configure whether the function only accepts https.
	//  This configuration is only applicable to 1st Gen functions with Http
	//  trigger. By default https is optional for 1st Gen functions; 2nd Gen
	//  functions are https ONLY.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.security_level
	SecurityLevel *string `json:"securityLevel,omitempty"`

	// Optional. The binary authorization policy to be checked when deploying the
	//  Cloud Run service.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.binary_authorization_policy
	BinaryAuthorizationPolicy *string `json:"binaryAuthorizationPolicy,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.Source
type Source struct {
	// If provided, get the source from this location in Google Cloud Storage.
	// +kcc:proto:field=google.cloud.functions.v2.Source.storage_source
	StorageSource *StorageSource `json:"storageSource,omitempty"`

	// If provided, get the source from this location in a Cloud Source
	//  Repository.
	// +kcc:proto:field=google.cloud.functions.v2.Source.repo_source
	RepoSource *RepoSource `json:"repoSource,omitempty"`

	// If provided, get the source from GitHub repository. This option is valid
	//  only for GCF 1st Gen function.
	//  Example: https://github.com/<user>/<repo>/blob/<commit>/<path-to-code>
	// +kcc:proto:field=google.cloud.functions.v2.Source.git_uri
	GitURI *string `json:"gitURI,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.SourceProvenance
type SourceProvenance struct {
	// A copy of the build's `source.storage_source`, if exists, with any
	//  generations resolved.
	// +kcc:proto:field=google.cloud.functions.v2.SourceProvenance.resolved_storage_source
	ResolvedStorageSource *StorageSource `json:"resolvedStorageSource,omitempty"`

	// A copy of the build's `source.repo_source`, if exists, with any
	//  revisions resolved.
	// +kcc:proto:field=google.cloud.functions.v2.SourceProvenance.resolved_repo_source
	ResolvedRepoSource *RepoSource `json:"resolvedRepoSource,omitempty"`

	// A copy of the build's `source.git_uri`, if exists, with any commits
	//  resolved.
	// +kcc:proto:field=google.cloud.functions.v2.SourceProvenance.git_uri
	GitURI *string `json:"gitURI,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.StateMessage
type StateMessage struct {
	// Severity of the state message.
	// +kcc:proto:field=google.cloud.functions.v2.StateMessage.severity
	Severity *string `json:"severity,omitempty"`

	// One-word CamelCase type of the state message.
	// +kcc:proto:field=google.cloud.functions.v2.StateMessage.type
	Type *string `json:"type,omitempty"`

	// The message.
	// +kcc:proto:field=google.cloud.functions.v2.StateMessage.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.StorageSource
type StorageSource struct {
	// Google Cloud Storage bucket containing the source (see
	//  [Bucket Name
	//  Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	// +kcc:proto:field=google.cloud.functions.v2.StorageSource.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Google Cloud Storage object containing the source.
	//
	//  This object must be a gzipped archive file (`.tar.gz`) containing source to
	//  build.
	// +kcc:proto:field=google.cloud.functions.v2.StorageSource.object
	Object *string `json:"object,omitempty"`

	// Google Cloud Storage generation for the object. If the generation is
	//  omitted, the latest generation will be used.
	// +kcc:proto:field=google.cloud.functions.v2.StorageSource.generation
	Generation *int64 `json:"generation,omitempty"`

	// When the specified storage bucket is a 1st gen function uploard url bucket,
	//  this field should be set as the generated upload url for 1st gen
	//  deployment.
	// +kcc:proto:field=google.cloud.functions.v2.StorageSource.source_upload_url
	SourceUploadURL *string `json:"sourceUploadURL,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.BuildConfig
type BuildConfigObservedState struct {
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.on_deploy_update_policy
	OnDeployUpdatePolicy *OnDeployUpdatePolicyObservedState `json:"onDeployUpdatePolicy,omitempty"`

	// Output only. The Cloud Build name of the latest successful deployment of
	//  the function.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.build
	Build *string `json:"build,omitempty"`

	// Output only. A permanent fixed identifier for source.
	// +kcc:proto:field=google.cloud.functions.v2.BuildConfig.source_provenance
	SourceProvenance *SourceProvenance `json:"sourceProvenance,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.EventTrigger
type EventTriggerObservedState struct {
	// Output only. The resource name of the Eventarc trigger. The format of this
	//  field is `projects/{project}/locations/{region}/triggers/{trigger}`.
	// +kcc:proto:field=google.cloud.functions.v2.EventTrigger.trigger
	Trigger *string `json:"trigger,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.Function
type FunctionObservedState struct {
	// Describes the Build step of the function that builds a container from the
	//  given source.
	// +kcc:proto:field=google.cloud.functions.v2.Function.build_config
	BuildConfig *BuildConfigObservedState `json:"buildConfig,omitempty"`

	// Describes the Service being deployed. Currently deploys services to Cloud
	//  Run (fully managed).
	// +kcc:proto:field=google.cloud.functions.v2.Function.service_config
	ServiceConfig *ServiceConfigObservedState `json:"serviceConfig,omitempty"`

	// An Eventarc trigger managed by Google Cloud Functions that fires events in
	//  response to a condition in another service.
	// +kcc:proto:field=google.cloud.functions.v2.Function.event_trigger
	EventTrigger *EventTriggerObservedState `json:"eventTrigger,omitempty"`

	// Output only. State of the function.
	// +kcc:proto:field=google.cloud.functions.v2.Function.state
	State *string `json:"state,omitempty"`

	// Output only. The last update timestamp of a Cloud Function.
	// +kcc:proto:field=google.cloud.functions.v2.Function.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State Messages for this Cloud Function.
	// +kcc:proto:field=google.cloud.functions.v2.Function.state_messages
	StateMessages []StateMessage `json:"stateMessages,omitempty"`

	// Output only. The deployed url for the function.
	// +kcc:proto:field=google.cloud.functions.v2.Function.url
	URL *string `json:"url,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.functions.v2.Function.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. The create timestamp of a Cloud Function. This is only
	//  applicable to 2nd Gen functions.
	// +kcc:proto:field=google.cloud.functions.v2.Function.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.OnDeployUpdatePolicy
type OnDeployUpdatePolicyObservedState struct {
	// Output only. contains the runtime version which was used during latest
	//  function deployment.
	// +kcc:proto:field=google.cloud.functions.v2.OnDeployUpdatePolicy.runtime_version
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
}

// +kcc:proto=google.cloud.functions.v2.ServiceConfig
type ServiceConfigObservedState struct {
	// Output only. Name of the service associated with a Function.
	//  The format of this field is
	//  `projects/{project}/locations/{region}/services/{service}`
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.service
	Service *string `json:"service,omitempty"`

	// Output only. URI of the Service deployed.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.uri
	URI *string `json:"uri,omitempty"`

	// Output only. The name of service revision.
	// +kcc:proto:field=google.cloud.functions.v2.ServiceConfig.revision
	Revision *string `json:"revision,omitempty"`
}
