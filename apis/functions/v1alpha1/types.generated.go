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


// +kcc:proto=google.cloud.functions.v1.CloudFunction
type CloudFunction struct {
	// A user-defined name of the function. Function names must be unique
	//  globally and match pattern `projects/*/locations/*/functions/*`
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.name
	Name *string `json:"name,omitempty"`

	// User-provided description of a function.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.description
	Description *string `json:"description,omitempty"`

	// The Google Cloud Storage URL, starting with `gs://`, pointing to the zip
	//  archive which contains the function.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_archive_url
	SourceArchiveURL *string `json:"sourceArchiveURL,omitempty"`

	// **Beta Feature**
	//
	//  The source repository where a function is hosted.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_repository
	SourceRepository *SourceRepository `json:"sourceRepository,omitempty"`

	// The Google Cloud Storage signed URL used for source uploading, generated
	//  by calling [google.cloud.functions.v1.GenerateUploadUrl].
	//
	//  The signature is validated on write methods (Create, Update)
	//  The signature is stripped from the Function object on read methods (Get,
	//  List)
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_upload_url
	SourceUploadURL *string `json:"sourceUploadURL,omitempty"`

	// An HTTPS endpoint type of source that can be triggered via URL.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.https_trigger
	HTTPSTrigger *HttpsTrigger `json:"httpsTrigger,omitempty"`

	// A source that fires events in response to a condition in another service.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.event_trigger
	EventTrigger *EventTrigger `json:"eventTrigger,omitempty"`

	// The name of the function (as defined in source code) that will be
	//  executed. Defaults to the resource name suffix (ID of the function), if not
	//  specified.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.entry_point
	EntryPoint *string `json:"entryPoint,omitempty"`

	// The runtime in which to run the function. Required when deploying a new
	//  function, optional when updating an existing function. For a complete
	//  list of possible choices, see the
	//  [`gcloud` command
	//  reference](https://cloud.google.com/sdk/gcloud/reference/functions/deploy#--runtime).
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.runtime
	Runtime *string `json:"runtime,omitempty"`

	// The function execution timeout. Execution is considered failed and
	//  can be terminated if the function is not completed at the end of the
	//  timeout period. Defaults to 60 seconds.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.timeout
	Timeout *string `json:"timeout,omitempty"`

	// The amount of memory in MB available for a function.
	//  Defaults to 256MB.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.available_memory_mb
	AvailableMemoryMb *int32 `json:"availableMemoryMb,omitempty"`

	// The email of the function's service account. If empty, defaults to
	//  `{project_id}@appspot.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Labels associated with this Cloud Function.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Environment variables that shall be available during function execution.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.environment_variables
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	// Build environment variables that shall be available during build time.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.build_environment_variables
	BuildEnvironmentVariables map[string]string `json:"buildEnvironmentVariables,omitempty"`

	// Deprecated: use vpc_connector
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.network
	Network *string `json:"network,omitempty"`

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
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.max_instances
	MaxInstances *int32 `json:"maxInstances,omitempty"`

	// A lower bound for the number function instances that may coexist at a
	//  given time.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.min_instances
	MinInstances *int32 `json:"minInstances,omitempty"`

	// The VPC Network Connector that this cloud function can connect to. It can
	//  be either the fully-qualified URI, or the short name of the network
	//  connector resource. The format of this field is
	//  `projects/*/locations/*/connectors/*`
	//
	//  This field is mutually exclusive with `network` field and will eventually
	//  replace it.
	//
	//  See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for
	//  more information on connecting Cloud projects.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.vpc_connector
	VpcConnector *string `json:"vpcConnector,omitempty"`

	// The egress settings for the connector, controlling what traffic is diverted
	//  through it.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.vpc_connector_egress_settings
	VpcConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty"`

	// The ingress settings for the function, controlling what traffic can reach
	//  it.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.ingress_settings
	IngressSettings *string `json:"ingressSettings,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt function resources.
	//
	//  It must match the pattern
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	//
	//  If specified, you must also provide an artifact registry repository using
	//  the `docker_repository` field that was created with the same KMS crypto
	//  key.
	//
	//  The following service accounts need to be granted the role 'Cloud KMS
	//  CryptoKey Encrypter/Decrypter (roles/cloudkms.cryptoKeyEncrypterDecrypter)'
	//  on the Key/KeyRing/Project/Organization (least access preferred).
	//
	//  1. Google Cloud Functions service account
	//     (service-{project_number}@gcf-admin-robot.iam.gserviceaccount.com) -
	//     Required to protect the function's image.
	//  2. Google Storage service account
	//     (service-{project_number}@gs-project-accounts.iam.gserviceaccount.com) -
	//     Required to protect the function's source code.
	//     If this service account does not exist, deploying a function without a
	//     KMS key or retrieving the service agent name provisions it. For more
	//     information, see
	//     https://cloud.google.com/storage/docs/projects#service-agents and
	//     https://cloud.google.com/storage/docs/getting-service-agent#gsutil.
	//
	//  Google Cloud Functions delegates access to service agents to protect
	//  function resources in internal projects that are not accessible by the
	//  end user.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Name of the Cloud Build Custom Worker Pool that should be used to build the
	//  function. The format of this field is
	//  `projects/{project}/locations/{region}/workerPools/{workerPool}` where
	//  `{project}` and `{region}` are the project id and region respectively where
	//  the worker pool is defined and `{workerPool}` is the short name of the
	//  worker pool.
	//
	//  If the project id is not the same as the function, then the Cloud
	//  Functions Service Agent
	//  (`service-<project_number>@gcf-admin-robot.iam.gserviceaccount.com`) must
	//  be granted the role Cloud Build Custom Workers Builder
	//  (`roles/cloudbuild.customworkers.builder`) in the project.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.build_worker_pool
	BuildWorkerPool *string `json:"buildWorkerPool,omitempty"`

	// Secret environment variables configuration.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.secret_environment_variables
	SecretEnvironmentVariables []SecretEnvVar `json:"secretEnvironmentVariables,omitempty"`

	// Secret volumes configuration.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.secret_volumes
	SecretVolumes []SecretVolume `json:"secretVolumes,omitempty"`

	// Input only. An identifier for Firebase function sources. Disclaimer: This
	//  field is only supported for Firebase function deployments.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_token
	SourceToken *string `json:"sourceToken,omitempty"`

	// User-managed repository created in Artifact Registry to which the
	//  function's Docker image will be pushed after it is built by Cloud Build.
	//  May optionally be encrypted with a customer-managed encryption key (CMEK).
	//  If unspecified and `docker_registry` is not explicitly set to
	//  `CONTAINER_REGISTRY`, GCF will create and use a default Artifact Registry
	//  repository named 'gcf-artifacts' in the region.
	//
	//  It must match the pattern
	//  `projects/{project}/locations/{location}/repositories/{repository}`.
	//
	//  Cross-project repositories are not supported.
	//  Cross-location repositories are not supported.
	//  Repository format must be 'DOCKER'.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.docker_repository
	DockerRepository *string `json:"dockerRepository,omitempty"`

	// Docker Registry to use for this deployment.
	//
	//  If unspecified, it defaults to `ARTIFACT_REGISTRY`.
	//  If `docker_repository` field is specified, this field should either be left
	//  unspecified or set to `ARTIFACT_REGISTRY`.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.docker_registry
	DockerRegistry *string `json:"dockerRegistry,omitempty"`

	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.automatic_update_policy
	AutomaticUpdatePolicy *CloudFunction_AutomaticUpdatePolicy `json:"automaticUpdatePolicy,omitempty"`

	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.on_deploy_update_policy
	OnDeployUpdatePolicy *CloudFunction_OnDeployUpdatePolicy `json:"onDeployUpdatePolicy,omitempty"`

	// A service account the user provides for use with Cloud Build. The format of
	//  this field is
	//  `projects/{projectId}/serviceAccounts/{serviceAccountEmail}`.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.build_service_account
	BuildServiceAccount *string `json:"buildServiceAccount,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.CloudFunction.AutomaticUpdatePolicy
type CloudFunction_AutomaticUpdatePolicy struct {
}

// +kcc:proto=google.cloud.functions.v1.CloudFunction.OnDeployUpdatePolicy
type CloudFunction_OnDeployUpdatePolicy struct {
}

// +kcc:proto=google.cloud.functions.v1.EventTrigger
type EventTrigger struct {
	// Required. The type of event to observe. For example:
	//  `providers/cloud.storage/eventTypes/object.change` and
	//  `providers/cloud.pubsub/eventTypes/topic.publish`.
	//
	//  Event types match pattern `providers/*/eventTypes/*.*`.
	//  The pattern contains:
	//
	//  1. namespace: For example, `cloud.storage` and
	//     `google.firebase.analytics`.
	//  2. resource type: The type of resource on which event occurs. For
	//     example, the Google Cloud Storage API includes the type `object`.
	//  3. action: The action that generates the event. For example, action for
	//     a Google Cloud Storage Object is 'change'.
	//  These parts are lower case.
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.event_type
	EventType *string `json:"eventType,omitempty"`

	// Required. The resource(s) from which to observe events, for example,
	//  `projects/_/buckets/myBucket`.
	//
	//  Not all syntactically correct values are accepted by all services. For
	//  example:
	//
	//  1. The authorization model must support it. Google Cloud Functions
	//     only allows EventTriggers to be deployed that observe resources in the
	//     same project as the `CloudFunction`.
	//  2. The resource type must match the pattern expected for an
	//     `event_type`. For example, an `EventTrigger` that has an
	//     `event_type` of "google.pubsub.topic.publish" should have a resource
	//     that matches Google Cloud Pub/Sub topics.
	//
	//  Additionally, some services may support short names when creating an
	//  `EventTrigger`. These will always be returned in the normalized "long"
	//  format.
	//
	//  See each *service's* documentation for supported formats.
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.resource
	Resource *string `json:"resource,omitempty"`

	// The hostname of the service that should be observed.
	//
	//  If no string is provided, the default service implementing the API will
	//  be used. For example, `storage.googleapis.com` is the default for all
	//  event types in the `google.storage` namespace.
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.service
	Service *string `json:"service,omitempty"`

	// Specifies policy for failed executions.
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.failure_policy
	FailurePolicy *FailurePolicy `json:"failurePolicy,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.FailurePolicy
type FailurePolicy struct {
	// If specified, then the function will be retried in case of a failure.
	// +kcc:proto:field=google.cloud.functions.v1.FailurePolicy.retry
	Retry *FailurePolicy_Retry `json:"retry,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.FailurePolicy.Retry
type FailurePolicy_Retry struct {
}

// +kcc:proto=google.cloud.functions.v1.HttpsTrigger
type HttpsTrigger struct {

	// The security level for the function.
	// +kcc:proto:field=google.cloud.functions.v1.HttpsTrigger.security_level
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SecretEnvVar
type SecretEnvVar struct {
	// Name of the environment variable.
	// +kcc:proto:field=google.cloud.functions.v1.SecretEnvVar.key
	Key *string `json:"key,omitempty"`

	// Project identifier (preferrably project number but can also be the project
	//  ID) of the project that contains the secret. If not set, it will be
	//  populated with the function's project assuming that the secret exists in
	//  the same project as of the function.
	// +kcc:proto:field=google.cloud.functions.v1.SecretEnvVar.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	// +kcc:proto:field=google.cloud.functions.v1.SecretEnvVar.secret
	Secret *string `json:"secret,omitempty"`

	// Version of the secret (version number or the string 'latest'). It is
	//  recommended to use a numeric version for secret environment variables as
	//  any updates to the secret value is not reflected until new instances start.
	// +kcc:proto:field=google.cloud.functions.v1.SecretEnvVar.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SecretVolume
type SecretVolume struct {
	// The path within the container to mount the secret volume. For example,
	//  setting the mount_path as `/etc/secrets` would mount the secret value files
	//  under the `/etc/secrets` directory. This directory will also be completely
	//  shadowed and unavailable to mount any other secrets.
	//
	//  Recommended mount paths: /etc/secrets
	//  Restricted mount paths: /cloudsql, /dev/log, /pod, /proc, /var/log
	// +kcc:proto:field=google.cloud.functions.v1.SecretVolume.mount_path
	MountPath *string `json:"mountPath,omitempty"`

	// Project identifier (preferrably project number but can also be the project
	//  ID) of the project that contains the secret. If not set, it will be
	//  populated with the function's project assuming that the secret exists in
	//  the same project as of the function.
	// +kcc:proto:field=google.cloud.functions.v1.SecretVolume.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Name of the secret in secret manager (not the full resource name).
	// +kcc:proto:field=google.cloud.functions.v1.SecretVolume.secret
	Secret *string `json:"secret,omitempty"`

	// List of secret versions to mount for this secret. If empty, the `latest`
	//  version of the secret will be made available in a file named after the
	//  secret under the mount point.
	// +kcc:proto:field=google.cloud.functions.v1.SecretVolume.versions
	Versions []SecretVolume_SecretVersion `json:"versions,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SecretVolume.SecretVersion
type SecretVolume_SecretVersion struct {
	// Version of the secret (version number or the string 'latest'). It is
	//  preferable to use `latest` version with secret volumes as secret value
	//  changes are reflected immediately.
	// +kcc:proto:field=google.cloud.functions.v1.SecretVolume.SecretVersion.version
	Version *string `json:"version,omitempty"`

	// Relative path of the file under the mount path where the secret value for
	//  this version will be fetched and made available. For example, setting the
	//  mount_path as '/etc/secrets' and path as `/secret_foo` would mount the
	//  secret value file at `/etc/secrets/secret_foo`.
	// +kcc:proto:field=google.cloud.functions.v1.SecretVolume.SecretVersion.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SourceRepository
type SourceRepository struct {
	// The URL pointing to the hosted repository where the function is defined.
	//  There are supported Cloud Source Repository URLs in the following
	//  formats:
	//
	//  To refer to a specific commit:
	//  `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*`
	//  To refer to a moveable alias (branch):
	//  `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*`
	//  In particular, to refer to HEAD use `master` moveable alias.
	//  To refer to a specific fixed alias (tag):
	//  `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*`
	//
	//  You may omit `paths/*` if you want to use the main directory. The function
	//  response may add an empty `/paths/` to the URL.
	// +kcc:proto:field=google.cloud.functions.v1.SourceRepository.url
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.CloudFunction
type CloudFunctionObservedState struct {
	// **Beta Feature**
	//
	//  The source repository where a function is hosted.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_repository
	SourceRepository *SourceRepositoryObservedState `json:"sourceRepository,omitempty"`

	// An HTTPS endpoint type of source that can be triggered via URL.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.https_trigger
	HTTPSTrigger *HttpsTriggerObservedState `json:"httpsTrigger,omitempty"`

	// Output only. Status of the function deployment.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.status
	Status *string `json:"status,omitempty"`

	// Output only. The last update timestamp of a Cloud Function.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The version identifier of the Cloud Function. Each deployment
	//  attempt results in a new version of a function being created.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.version_id
	VersionID *int64 `json:"versionID,omitempty"`

	// Output only. The Cloud Build ID of the latest successful deployment of the
	//  function.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.build_id
	BuildID *string `json:"buildID,omitempty"`

	// Output only. The Cloud Build Name of the function deployment.
	//  `projects/<project-number>/locations/<region>/builds/<build-id>`.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.build_name
	BuildName *string `json:"buildName,omitempty"`

	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.on_deploy_update_policy
	OnDeployUpdatePolicy *CloudFunction_OnDeployUpdatePolicyObservedState `json:"onDeployUpdatePolicy,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.CloudFunction.OnDeployUpdatePolicy
type CloudFunction_OnDeployUpdatePolicyObservedState struct {
	// Output only. Contains the runtime version which was used during latest
	//  function deployment.
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.OnDeployUpdatePolicy.runtime_version
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.HttpsTrigger
type HttpsTriggerObservedState struct {
	// Output only. The deployed url for the function.
	// +kcc:proto:field=google.cloud.functions.v1.HttpsTrigger.url
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SourceRepository
type SourceRepositoryObservedState struct {
	// Output only. The URL pointing to the hosted repository where the function
	//  were defined at the time of deployment. It always points to a specific
	//  commit in the format described above.
	// +kcc:proto:field=google.cloud.functions.v1.SourceRepository.deployed_url
	DeployedURL *string `json:"deployedURL,omitempty"`
}
