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

// +kcc:proto=google.cloud.aiplatform.v1.DeploymentResourcePool
type DeploymentResourcePool struct {
	// Immutable. The resource name of the DeploymentResourcePool.
	//  Format:
	//  `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.name
	Name *string `json:"name,omitempty"`

	// Required. The underlying DedicatedResources that the DeploymentResourcePool
	//  uses.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.dedicated_resources
	DedicatedResources *DedicatedResources `json:"dedicatedResources,omitempty"`

	// Customer-managed encryption key spec for a DeploymentResourcePool. If set,
	//  this DeploymentResourcePool will be secured by this key. Endpoints and the
	//  DeploymentResourcePool they deploy in need to have the same EncryptionSpec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The service account that the DeploymentResourcePool's container(s) run as.
	//  Specify the email address of the service account. If this service account
	//  is not specified, the container(s) run as a service account that doesn't
	//  have access to the resource project.
	//
	//  Users deploying the Models to this DeploymentResourcePool must have the
	//  `iam.serviceAccounts.actAs` permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// If the DeploymentResourcePool is deployed with custom-trained Models or
	//  AutoML Tabular Models, the container(s) of the DeploymentResourcePool will
	//  send `stderr` and `stdout` streams to Cloud Logging by default.
	//  Please note that the logs incur cost, which are subject to [Cloud Logging
	//  pricing](https://cloud.google.com/logging/pricing).
	//
	//  User can disable container logging by setting this flag to true.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.disable_container_logging
	DisableContainerLogging *bool `json:"disableContainerLogging,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.DeploymentResourcePool
type DeploymentResourcePoolObservedState struct {
	// Output only. Timestamp when this DeploymentResourcePool was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeploymentResourcePool.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
