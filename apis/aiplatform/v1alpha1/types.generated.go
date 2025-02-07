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


// +kcc:proto=google.cloud.aiplatform.v1.AutomaticResources
type AutomaticResources struct {
	// Immutable. The minimum number of replicas this DeployedModel will be always
	//  deployed on. If traffic against it increases, it may dynamically be
	//  deployed onto more replicas up to
	//  [max_replica_count][google.cloud.aiplatform.v1.AutomaticResources.max_replica_count],
	//  and as traffic decreases, some of these extra replicas may be freed. If the
	//  requested value is too large, the deployment will error.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AutomaticResources.min_replica_count
	MinReplicaCount *int32 `json:"minReplicaCount,omitempty"`

	// Immutable. The maximum number of replicas this DeployedModel may be
	//  deployed on when the traffic against it increases. If the requested value
	//  is too large, the deployment will error, but if deployment succeeds then
	//  the ability to scale the model to that many replicas is guaranteed (barring
	//  service outages). If traffic against the DeployedModel increases beyond
	//  what its replicas at maximum may handle, a portion of the traffic will be
	//  dropped. If this value is not provided, a no upper bound for scaling under
	//  heavy traffic will be assume, though Vertex AI may be unable to scale
	//  beyond certain replica number.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AutomaticResources.max_replica_count
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.AutoscalingMetricSpec
type AutoscalingMetricSpec struct {
	// Required. The resource metric name.
	//  Supported metrics:
	//
	//  * For Online Prediction:
	//  * `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle`
	//  * `aiplatform.googleapis.com/prediction/online/cpu/utilization`
	// +kcc:proto:field=google.cloud.aiplatform.v1.AutoscalingMetricSpec.metric_name
	MetricName *string `json:"metricName,omitempty"`

	// The target resource utilization in percentage (1% - 100%) for the given
	//  metric; once the real usage deviates from the target by a certain
	//  percentage, the machine replicas change. The default value is 60
	//  (representing 60%) if not provided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AutoscalingMetricSpec.target
	Target *int32 `json:"target,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DedicatedResources
type DedicatedResources struct {
	// Required. Immutable. The specification of a single machine used by the
	//  prediction.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DedicatedResources.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Required. Immutable. The minimum number of machine replicas this
	//  DeployedModel will be always deployed on. This value must be greater than
	//  or equal to 1.
	//
	//  If traffic against the DeployedModel increases, it may dynamically be
	//  deployed onto more replicas, and as traffic decreases, some of these extra
	//  replicas may be freed.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DedicatedResources.min_replica_count
	MinReplicaCount *int32 `json:"minReplicaCount,omitempty"`

	// Immutable. The maximum number of replicas this DeployedModel may be
	//  deployed on when the traffic against it increases. If the requested value
	//  is too large, the deployment will error, but if deployment succeeds then
	//  the ability to scale the model to that many replicas is guaranteed (barring
	//  service outages). If traffic against the DeployedModel increases beyond
	//  what its replicas at maximum may handle, a portion of the traffic will be
	//  dropped. If this value is not provided, will use
	//  [min_replica_count][google.cloud.aiplatform.v1.DedicatedResources.min_replica_count]
	//  as the default value.
	//
	//  The value of this field impacts the charge against Vertex CPU and GPU
	//  quotas. Specifically, you will be charged for (max_replica_count *
	//  number of cores in the selected machine type) and (max_replica_count *
	//  number of GPUs per replica in the selected machine type).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DedicatedResources.max_replica_count
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`

	// Optional. Number of required available replicas for the deployment to
	//  succeed. This field is only needed when partial model deployment/mutation
	//  is desired. If set, the model deploy/mutate operation will succeed once
	//  available_replica_count reaches required_replica_count, and the rest of
	//  the replicas will be retried. If not set, the default
	//  required_replica_count will be min_replica_count.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DedicatedResources.required_replica_count
	RequiredReplicaCount *int32 `json:"requiredReplicaCount,omitempty"`

	// Immutable. The metric specifications that overrides a resource
	//  utilization metric (CPU utilization, accelerator's duty cycle, and so on)
	//  target value (default to 60 if not set). At most one entry is allowed per
	//  metric.
	//
	//  If
	//  [machine_spec.accelerator_count][google.cloud.aiplatform.v1.MachineSpec.accelerator_count]
	//  is above 0, the autoscaling will be based on both CPU utilization and
	//  accelerator's duty cycle metrics and scale up when either metrics exceeds
	//  its target value while scale down if both metrics are under their target
	//  value. The default target value is 60 for both metrics.
	//
	//  If
	//  [machine_spec.accelerator_count][google.cloud.aiplatform.v1.MachineSpec.accelerator_count]
	//  is 0, the autoscaling will be based on CPU utilization metric only with
	//  default target value 60 if not explicitly set.
	//
	//  For example, in the case of Online Prediction, if you want to override
	//  target CPU utilization to 80, you should set
	//  [autoscaling_metric_specs.metric_name][google.cloud.aiplatform.v1.AutoscalingMetricSpec.metric_name]
	//  to `aiplatform.googleapis.com/prediction/online/cpu/utilization` and
	//  [autoscaling_metric_specs.target][google.cloud.aiplatform.v1.AutoscalingMetricSpec.target]
	//  to `80`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DedicatedResources.autoscaling_metric_specs
	AutoscalingMetricSpecs []AutoscalingMetricSpec `json:"autoscalingMetricSpecs,omitempty"`

	// Optional. If true, schedule the deployment workload on [spot
	//  VMs](https://cloud.google.com/kubernetes-engine/docs/concepts/spot-vms).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DedicatedResources.spot
	Spot *bool `json:"spot,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedIndex
type DeployedIndex struct {
	// Required. The user specified ID of the DeployedIndex.
	//  The ID can be up to 128 characters long and must start with a letter and
	//  only contain letters, numbers, and underscores.
	//  The ID must be unique within the project it is created in.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.id
	ID *string `json:"id,omitempty"`

	// Required. The name of the Index this is the deployment of.
	//  We may refer to this Index as the DeployedIndex's "original" Index.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.index
	Index *string `json:"index,omitempty"`

	// The display name of the DeployedIndex. If not provided upon creation,
	//  the Index's display_name is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A description of resources that the DeployedIndex uses, which to
	//  large degree are decided by Vertex AI, and optionally allows only a modest
	//  additional configuration.
	//  If min_replica_count is not set, the default value is 2 (we don't provide
	//  SLA when min_replica_count=1). If max_replica_count is not set, the
	//  default value is min_replica_count. The max allowed replica count is
	//  1000.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.automatic_resources
	AutomaticResources *AutomaticResources `json:"automaticResources,omitempty"`

	// Optional. A description of resources that are dedicated to the
	//  DeployedIndex, and that need a higher degree of manual configuration. The
	//  field min_replica_count must be set to a value strictly greater than 0, or
	//  else validation will fail. We don't provide SLA when min_replica_count=1.
	//  If max_replica_count is not set, the default value is min_replica_count.
	//  The max allowed replica count is 1000.
	//
	//  Available machine types for SMALL shard:
	//  e2-standard-2 and all machine types available for MEDIUM and LARGE shard.
	//
	//  Available machine types for MEDIUM shard:
	//  e2-standard-16 and all machine types available for LARGE shard.
	//
	//  Available machine types for LARGE shard:
	//  e2-highmem-16, n2d-standard-32.
	//
	//  n1-standard-16 and n1-standard-32 are still available, but we recommend
	//  e2-standard-16 and e2-highmem-16 for cost efficiency.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.dedicated_resources
	DedicatedResources *DedicatedResources `json:"dedicatedResources,omitempty"`

	// Optional. If true, private endpoint's access logs are sent to Cloud
	//  Logging.
	//
	//  These logs are like standard server access logs, containing
	//  information like timestamp and latency for each MatchRequest.
	//
	//  Note that logs may incur a cost, especially if the deployed
	//  index receives a high queries per second rate (QPS).
	//  Estimate your costs before enabling this option.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.enable_access_logging
	EnableAccessLogging *bool `json:"enableAccessLogging,omitempty"`

	// Optional. If set, the authentication is enabled for the private endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.deployed_index_auth_config
	DeployedIndexAuthConfig *DeployedIndexAuthConfig `json:"deployedIndexAuthConfig,omitempty"`

	// Optional. A list of reserved ip ranges under the VPC network that can be
	//  used for this DeployedIndex.
	//
	//  If set, we will deploy the index within the provided ip ranges. Otherwise,
	//  the index might be deployed to any ip ranges under the provided VPC
	//  network.
	//
	//  The value should be the name of the address
	//  (https://cloud.google.com/compute/docs/reference/rest/v1/addresses)
	//  Example: ['vertex-ai-ip-range'].
	//
	//  For more information about subnets and network IP ranges, please see
	//  https://cloud.google.com/vpc/docs/subnets#manually_created_subnet_ip_ranges.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`

	// Optional. The deployment group can be no longer than 64 characters (eg:
	//  'test', 'prod'). If not set, we will use the 'default' deployment group.
	//
	//  Creating `deployment_groups` with `reserved_ip_ranges` is a recommended
	//  practice when the peered network has multiple peering ranges. This creates
	//  your deployments from predictable IP spaces for easier traffic
	//  administration. Also, one deployment_group (except 'default') can only be
	//  used with the same reserved_ip_ranges which means if the deployment_group
	//  has been used with reserved_ip_ranges: [a, b, c], using it with [a, b] or
	//  [d, e] is disallowed.
	//
	//  Note: we only support up to 5 deployment groups(not including 'default').
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.deployment_group
	DeploymentGroup *string `json:"deploymentGroup,omitempty"`

	// Optional. If set for PSC deployed index, PSC connection will be
	//  automatically created after deployment is done and the endpoint information
	//  is populated in private_endpoints.psc_automated_endpoints.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.psc_automation_configs
	PscAutomationConfigs []PSCAutomationConfig `json:"pscAutomationConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedIndexAuthConfig
type DeployedIndexAuthConfig struct {
	// Defines the authentication provider that the DeployedIndex uses.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexAuthConfig.auth_provider
	AuthProvider *DeployedIndexAuthConfig_AuthProvider `json:"authProvider,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedIndexAuthConfig.AuthProvider
type DeployedIndexAuthConfig_AuthProvider struct {
	// The list of JWT
	//  [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	//  that are allowed to access. A JWT containing any of these audiences will
	//  be accepted.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexAuthConfig.AuthProvider.audiences
	Audiences []string `json:"audiences,omitempty"`

	// A list of allowed JWT issuers. Each entry must be a valid Google
	//  service account, in the following format:
	//
	//  `service-account-name@project-id.iam.gserviceaccount.com`
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndexAuthConfig.AuthProvider.allowed_issuers
	AllowedIssuers []string `json:"allowedIssuers,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IndexEndpoint
type IndexEndpoint struct {

	// Required. The display name of the IndexEndpoint.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the IndexEndpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.description
	Description *string `json:"description,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your IndexEndpoints.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks)
	//  to which the IndexEndpoint should be peered.
	//
	//  Private services access must already be configured for the network. If left
	//  unspecified, the Endpoint is not peered with any network.
	//
	//  [network][google.cloud.aiplatform.v1.IndexEndpoint.network] and
	//  [private_service_connect_config][google.cloud.aiplatform.v1.IndexEndpoint.private_service_connect_config]
	//  are mutually exclusive.
	//
	//  [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	//  `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in '12345', and {network} is
	//  network name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.network
	Network *string `json:"network,omitempty"`

	// Optional. Deprecated: If true, expose the IndexEndpoint via private service
	//  connect.
	//
	//  Only one of the fields,
	//  [network][google.cloud.aiplatform.v1.IndexEndpoint.network] or
	//  [enable_private_service_connect][google.cloud.aiplatform.v1.IndexEndpoint.enable_private_service_connect],
	//  can be set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// Optional. Configuration for private service connect.
	//
	//  [network][google.cloud.aiplatform.v1.IndexEndpoint.network] and
	//  [private_service_connect_config][google.cloud.aiplatform.v1.IndexEndpoint.private_service_connect_config]
	//  are mutually exclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.private_service_connect_config
	PrivateServiceConnectConfig *PrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty"`

	// Optional. If true, the deployed index will be accessible through public
	//  endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.public_endpoint_enabled
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty"`

	// Immutable. Customer-managed encryption key spec for an IndexEndpoint. If
	//  set, this IndexEndpoint and all sub-resources of this IndexEndpoint will be
	//  secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IndexPrivateEndpoints
type IndexPrivateEndpoints struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.MachineSpec
type MachineSpec struct {
	// Immutable. The type of the machine.
	//
	//  See the [list of machine types supported for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	//  See the [list of machine types supported for custom
	//  training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	//  For [DeployedModel][google.cloud.aiplatform.v1.DeployedModel] this field is
	//  optional, and the default value is `n1-standard-2`. For
	//  [BatchPredictionJob][google.cloud.aiplatform.v1.BatchPredictionJob] or as
	//  part of [WorkerPoolSpec][google.cloud.aiplatform.v1.WorkerPoolSpec] this
	//  field is required.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine
	//  as per
	//  [accelerator_count][google.cloud.aiplatform.v1.MachineSpec.accelerator_count].
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Immutable. The topology of the TPUs. Corresponds to the TPU topologies
	//  available from GKE. (Example: tpu_topology: "2x2x1").
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	// Optional. Immutable. Configuration controlling how this resource pool
	//  consumes reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.MachineSpec.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PSCAutomationConfig
type PSCAutomationConfig struct {
	// Required. Project id used to create forwarding rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PSCAutomationConfig.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks).
	//  [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	//  `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in '12345', and {network} is
	//  network name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PSCAutomationConfig.network
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateServiceConnectConfig
type PrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// A list of Projects from which the forwarding rule will target the service
	//  attachment.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.project_allowlist
	ProjectAllowlist []string `json:"projectAllowlist,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PscAutomatedEndpoints
type PscAutomatedEndpoints struct {
	// Corresponding project_id in pscAutomationConfigs
	// +kcc:proto:field=google.cloud.aiplatform.v1.PscAutomatedEndpoints.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Corresponding network in pscAutomationConfigs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PscAutomatedEndpoints.network
	Network *string `json:"network,omitempty"`

	// Ip Address created by the automated forwarding rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PscAutomatedEndpoints.match_address
	MatchAddress *string `json:"matchAddress,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ReservationAffinity
type ReservationAffinity struct {
	// Required. Specifies the reservation affinity type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.reservation_affinity_type
	ReservationAffinityType *string `json:"reservationAffinityType,omitempty"`

	// Optional. Corresponds to the label key of a reservation resource. To target
	//  a SPECIFIC_RESERVATION by name, use
	//  `compute.googleapis.com/reservation-name` as the key and specify the name
	//  of your reservation as its value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of a reservation resource. This
	//  must be the full resource name of the reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedIndex
type DeployedIndexObservedState struct {
	// Output only. Timestamp when the DeployedIndex was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Provides paths for users to send requests directly to the
	//  deployed index services running on Cloud via private services access. This
	//  field is populated if
	//  [network][google.cloud.aiplatform.v1.IndexEndpoint.network] is configured.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.private_endpoints
	PrivateEndpoints *IndexPrivateEndpoints `json:"privateEndpoints,omitempty"`

	// Output only. The DeployedIndex may depend on various data on its original
	//  Index. Additionally when certain changes to the original Index are being
	//  done (e.g. when what the Index contains is being changed) the DeployedIndex
	//  may be asynchronously updated in the background to reflect these changes.
	//  If this timestamp's value is at least the
	//  [Index.update_time][google.cloud.aiplatform.v1.Index.update_time] of the
	//  original Index, it means that this DeployedIndex and the original Index are
	//  in sync. If this timestamp is older, then to see which updates this
	//  DeployedIndex already contains (and which it does not), one must
	//  [list][google.longrunning.Operations.ListOperations] the operations that
	//  are running on the original Index. Only the successfully completed
	//  Operations with
	//  [update_time][google.cloud.aiplatform.v1.GenericOperationMetadata.update_time]
	//  equal or before this sync time are contained in this DeployedIndex.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedIndex.index_sync_time
	IndexSyncTime *string `json:"indexSyncTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IndexEndpoint
type IndexEndpointObservedState struct {
	// Output only. The resource name of the IndexEndpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.name
	Name *string `json:"name,omitempty"`

	// Output only. The indexes deployed in this endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.deployed_indexes
	DeployedIndexes []DeployedIndex `json:"deployedIndexes,omitempty"`

	// Output only. Timestamp when this IndexEndpoint was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this IndexEndpoint was last updated.
	//  This timestamp is not updated when the endpoint's DeployedIndexes are
	//  updated, e.g. due to updates of the original Indexes they are the
	//  deployments of.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Configuration for private service connect.
	//
	//  [network][google.cloud.aiplatform.v1.IndexEndpoint.network] and
	//  [private_service_connect_config][google.cloud.aiplatform.v1.IndexEndpoint.private_service_connect_config]
	//  are mutually exclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.private_service_connect_config
	PrivateServiceConnectConfig *PrivateServiceConnectConfigObservedState `json:"privateServiceConnectConfig,omitempty"`

	// Output only. If
	//  [public_endpoint_enabled][google.cloud.aiplatform.v1.IndexEndpoint.public_endpoint_enabled]
	//  is true, this field will be populated with the domain name to use for this
	//  index endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.public_endpoint_domain_name
	PublicEndpointDomainName *string `json:"publicEndpointDomainName,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexEndpoint.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IndexPrivateEndpoints
type IndexPrivateEndpointsObservedState struct {
	// Output only. The ip address used to send match gRPC requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexPrivateEndpoints.match_grpc_address
	MatchGrpcAddress *string `json:"matchGrpcAddress,omitempty"`

	// Output only. The name of the service attachment resource. Populated if
	//  private service connect is enabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexPrivateEndpoints.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`

	// Output only. PscAutomatedEndpoints is populated if private service connect
	//  is enabled if PscAutomatedConfig is set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IndexPrivateEndpoints.psc_automated_endpoints
	PscAutomatedEndpoints []PscAutomatedEndpoints `json:"pscAutomatedEndpoints,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateServiceConnectConfig
type PrivateServiceConnectConfigObservedState struct {
	// Output only. The name of the generated service attachment resource.
	//  This is only populated if the endpoint is deployed with
	//  PrivateServiceConnect.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}
