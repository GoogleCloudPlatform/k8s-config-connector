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


// +kcc:proto=google.cloud.aiplatform.v1.DiskSpec
type DiskSpec struct {
	// Type of the boot disk (default is "pd-ssd").
	//  Valid values: "pd-ssd" (Persistent Disk Solid State Drive) or
	//  "pd-standard" (Persistent Disk Hard Disk Drive).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DiskSpec.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Size in GB of the boot disk (default is 100GB).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DiskSpec.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.PersistentResource
type PersistentResource struct {
	// Immutable. Resource name of a PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.name
	Name *string `json:"name,omitempty"`

	// Optional. The display name of the PersistentResource.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The spec of the pools of different resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	ResourcePools []ResourcePool `json:"resourcePools,omitempty"`

	// Optional. The labels with user-defined metadata to organize
	//  PersistentResource.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The full name of the Compute Engine
	//  [network](/compute/docs/networks-and-firewalls#networks) to peered with
	//  Vertex AI to host the persistent resources.
	//  For example, `projects/12345/global/networks/myVPC`.
	//  [Format](/compute/docs/reference/rest/v1/networks/insert)
	//  is of the form `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in `12345`, and {network} is a
	//  network name.
	//
	//  To specify this field, you must have already [configured VPC Network
	//  Peering for Vertex
	//  AI](https://cloud.google.com/vertex-ai/docs/general/vpc-peering).
	//
	//  If this field is left unspecified, the resources aren't peered with any
	//  network.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.network
	Network *string `json:"network,omitempty"`

	// Optional. Customer-managed encryption key spec for a PersistentResource.
	//  If set, this PersistentResource and all sub-resources of this
	//  PersistentResource will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Persistent Resource runtime spec.
	//  For example, used for Ray cluster configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime_spec
	ResourceRuntimeSpec *ResourceRuntimeSpec `json:"resourceRuntimeSpec,omitempty"`

	// Optional. A list of names for the reserved IP ranges under the VPC network
	//  that can be used for this persistent resource.
	//
	//  If set, we will deploy the persistent resource within the provided IP
	//  ranges. Otherwise, the persistent resource is deployed to any IP
	//  ranges under the provided VPC network.
	//
	//  Example: ['vertex-ai-ip-range'].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RayLogsSpec
type RayLogsSpec struct {
	// Optional. Flag to disable the export of Ray OSS logs to Cloud Logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RayLogsSpec.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RayMetricSpec
type RayMetricSpec struct {
	// Optional. Flag to disable the Ray metrics collection.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RayMetricSpec.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.RaySpec
type RaySpec struct {
	// Optional. Default image for user to choose a preferred ML framework
	//  (for example, TensorFlow or Pytorch) by choosing from [Vertex prebuilt
	//  images](https://cloud.google.com/vertex-ai/docs/training/pre-built-containers).
	//  Either this or the resource_pool_images is required. Use this field if
	//  you need all the resource pools to have the same Ray image. Otherwise, use
	//  the {@code resource_pool_images} field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RaySpec.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Optional. Required if image_uri isn't set. A map of resource_pool_id to
	//  prebuild Ray image if user need to use different images for different
	//  head/worker pools. This map needs to cover all the resource pool ids.
	//  Example:
	//  {
	//    "ray_head_node_pool": "head image"
	//    "ray_worker_node_pool1": "worker image"
	//    "ray_worker_node_pool2": "another worker image"
	//  }
	// +kcc:proto:field=google.cloud.aiplatform.v1.RaySpec.resource_pool_images
	ResourcePoolImages map[string]string `json:"resourcePoolImages,omitempty"`

	// Optional. This will be used to indicate which resource pool will serve as
	//  the Ray head node(the first node within that pool). Will use the machine
	//  from the first workerpool as the head node by default if this field isn't
	//  set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RaySpec.head_node_resource_pool_id
	HeadNodeResourcePoolID *string `json:"headNodeResourcePoolID,omitempty"`

	// Optional. Ray metrics configurations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RaySpec.ray_metric_spec
	RayMetricSpec *RayMetricSpec `json:"rayMetricSpec,omitempty"`

	// Optional. OSS Ray logging configurations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RaySpec.ray_logs_spec
	RayLogsSpec *RayLogsSpec `json:"rayLogsSpec,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.ResourcePool
type ResourcePool struct {
	// Immutable. The unique ID in a PersistentResource for referring to this
	//  resource pool. User can specify it if necessary. Otherwise, it's generated
	//  automatically.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.id
	ID *string `json:"id,omitempty"`

	// Required. Immutable. The specification of a single machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Optional. The total number of machines to use for this resource pool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.replica_count
	ReplicaCount *int64 `json:"replicaCount,omitempty"`

	// Optional. Disk spec for the machine in this node pool.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.disk_spec
	DiskSpec *DiskSpec `json:"diskSpec,omitempty"`

	// Optional. Optional spec to configure GKE or Ray-on-Vertex autoscaling
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.autoscaling_spec
	AutoscalingSpec *ResourcePool_AutoscalingSpec `json:"autoscalingSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourcePool.AutoscalingSpec
type ResourcePool_AutoscalingSpec struct {
	// Optional. min replicas in the node pool,
	//  must be ≤ replica_count and < max_replica_count or will throw error.
	//  For autoscaling enabled Ray-on-Vertex, we allow min_replica_count of a
	//  resource_pool to be 0 to match the OSS Ray
	//  behavior(https://docs.ray.io/en/latest/cluster/vms/user-guides/configuring-autoscaling.html#cluster-config-parameters).
	//  As for Persistent Resource, the min_replica_count must be > 0, we added
	//  a corresponding validation inside
	//  CreatePersistentResourceRequestValidator.java.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.AutoscalingSpec.min_replica_count
	MinReplicaCount *int64 `json:"minReplicaCount,omitempty"`

	// Optional. max replicas in the node pool,
	//  must be ≥ replica_count and > min_replica_count or will throw error
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.AutoscalingSpec.max_replica_count
	MaxReplicaCount *int64 `json:"maxReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourceRuntime
type ResourceRuntime struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourceRuntimeSpec
type ResourceRuntimeSpec struct {
	// Optional. Configure the use of workload identity on the PersistentResource
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourceRuntimeSpec.service_account_spec
	ServiceAccountSpec *ServiceAccountSpec `json:"serviceAccountSpec,omitempty"`

	// Optional. Ray cluster configuration.
	//  Required when creating a dedicated RayCluster on the PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourceRuntimeSpec.ray_spec
	RaySpec *RaySpec `json:"raySpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ServiceAccountSpec
type ServiceAccountSpec struct {
	// Required. If true, custom user-managed service account is enforced to run
	//  any workloads (for example, Vertex Jobs) on the resource. Otherwise, uses
	//  the [Vertex AI Custom Code Service
	//  Agent](https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ServiceAccountSpec.enable_custom_service_account
	EnableCustomServiceAccount *bool `json:"enableCustomServiceAccount,omitempty"`

	// Optional. Required when all below conditions are met
	//   * `enable_custom_service_account` is true;
	//   * any runtime is specified via `ResourceRuntimeSpec` on creation time,
	//     for example, Ray
	//
	//  The users must have `iam.serviceAccounts.actAs` permission on this service
	//  account and then the specified runtime containers will run as it.
	//
	//  Do not set this field if you want to submit jobs using custom service
	//  account to this PersistentResource after creation, but only specify the
	//  `service_account` inside the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ServiceAccountSpec.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PersistentResource
type PersistentResourceObservedState struct {
	// Required. The spec of the pools of different resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	ResourcePools []ResourcePoolObservedState `json:"resourcePools,omitempty"`

	// Output only. The detailed state of a Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when persistent resource's state is `STOPPING`
	//  or `ERROR`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.error
	Error *Status `json:"error,omitempty"`

	// Output only. Time when the PersistentResource was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the PersistentResource for the first time entered
	//  the `RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the PersistentResource was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Runtime information of the Persistent Resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime
	ResourceRuntime *ResourceRuntime `json:"resourceRuntime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourcePool
type ResourcePoolObservedState struct {
	// Output only. The number of machines currently in use by training jobs for
	//  this resource pool. Will replace idle_replica_count.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourcePool.used_replica_count
	UsedReplicaCount *int64 `json:"usedReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ResourceRuntime
type ResourceRuntimeObservedState struct {
	// Output only. URIs for user to connect to the Cluster.
	//  Example:
	//  {
	//    "RAY_HEAD_NODE_INTERNAL_IP": "head-node-IP:10001"
	//    "RAY_DASHBOARD_URI": "ray-dashboard-address:8888"
	//  }
	// +kcc:proto:field=google.cloud.aiplatform.v1.ResourceRuntime.access_uris
	AccessUris map[string]string `json:"accessUris,omitempty"`
}
