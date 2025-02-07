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

// +kcc:proto=google.cloud.aiplatform.v1.EnvVar
type EnvVar struct {
	// Required. Name of the environment variable. Must be a valid C identifier.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EnvVar.name
	Name *string `json:"name,omitempty"`

	// Required. Variables that reference a $(VAR_NAME) are expanded
	//  using the previous defined environment variables in the container and
	//  any service environment variables. If a variable cannot be resolved,
	//  the reference in the input string will be unchanged. The $(VAR_NAME)
	//  syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	//  references will never be expanded, regardless of whether the variable
	//  exists or not.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EnvVar.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.LargeModelReference
type LargeModelReference struct {
	// Required. The unique name of the large Foundation or pre-built model. Like
	//  "chat-bison", "text-bison". Or model name with version ID, like
	//  "chat-bison@001", "text-bison@005", etc.
	// +kcc:proto:field=google.cloud.aiplatform.v1.LargeModelReference.name
	Name *string `json:"name,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.ModelContainerSpec
type ModelContainerSpec struct {
	// Required. Immutable. URI of the Docker image to be used as the custom
	//  container for serving predictions. This URI must identify an image in
	//  Artifact Registry or Container Registry. Learn more about the [container
	//  publishing
	//  requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing),
	//  including permissions requirements for the Vertex AI Service Agent.
	//
	//  The container image is ingested upon
	//  [ModelService.UploadModel][google.cloud.aiplatform.v1.ModelService.UploadModel],
	//  stored internally, and this original path is afterwards not used.
	//
	//  To learn about the requirements for the Docker image itself, see
	//  [Custom container
	//  requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#).
	//
	//  You can use the URI to one of Vertex AI's [pre-built container images for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers)
	//  in this field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Immutable. Specifies the command that runs when the container starts. This
	//  overrides the container's
	//  [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint).
	//  Specify this field as an array of executable and arguments, similar to a
	//  Docker `ENTRYPOINT`'s "exec" form, not its "shell" form.
	//
	//  If you do not specify this field, then the container's `ENTRYPOINT` runs,
	//  in conjunction with the
	//  [args][google.cloud.aiplatform.v1.ModelContainerSpec.args] field or the
	//  container's [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd),
	//  if either exists. If this field is not specified and the container does not
	//  have an `ENTRYPOINT`, then refer to the Docker documentation about [how
	//  `CMD` and `ENTRYPOINT`
	//  interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	//
	//  If you specify this field, then you can also specify the `args` field to
	//  provide additional arguments for this command. However, if you specify this
	//  field, then the container's `CMD` is ignored. See the
	//  [Kubernetes documentation about how the
	//  `command` and `args` fields interact with a container's `ENTRYPOINT` and
	//  `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).
	//
	//  In this field, you can reference [environment variables set by Vertex
	//  AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)
	//  and environment variables set in the
	//  [env][google.cloud.aiplatform.v1.ModelContainerSpec.env] field. You cannot
	//  reference environment variables set in the Docker image. In order for
	//  environment variables to be expanded, reference them by using the following
	//  syntax: <code>$(<var>VARIABLE_NAME</var>)</code> Note that this differs
	//  from Bash variable expansion, which does not use parentheses. If a variable
	//  cannot be resolved, the reference in the input string is used unchanged. To
	//  avoid variable expansion, you can escape this syntax with `$$`; for
	//  example: <code>$$(<var>VARIABLE_NAME</var>)</code> This field corresponds
	//  to the `command` field of the Kubernetes Containers [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.command
	Command []string `json:"command,omitempty"`

	// Immutable. Specifies arguments for the command that runs when the container
	//  starts. This overrides the container's
	//  [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd). Specify
	//  this field as an array of executable and arguments, similar to a Docker
	//  `CMD`'s "default parameters" form.
	//
	//  If you don't specify this field but do specify the
	//  [command][google.cloud.aiplatform.v1.ModelContainerSpec.command] field,
	//  then the command from the `command` field runs without any additional
	//  arguments. See the [Kubernetes documentation about how the `command` and
	//  `args` fields interact with a container's `ENTRYPOINT` and
	//  `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).
	//
	//  If you don't specify this field and don't specify the `command` field,
	//  then the container's
	//  [`ENTRYPOINT`](https://docs.docker.com/engine/reference/builder/#cmd) and
	//  `CMD` determine what runs based on their default behavior. See the Docker
	//  documentation about [how `CMD` and `ENTRYPOINT`
	//  interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	//
	//  In this field, you can reference [environment variables
	//  set by Vertex
	//  AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)
	//  and environment variables set in the
	//  [env][google.cloud.aiplatform.v1.ModelContainerSpec.env] field. You cannot
	//  reference environment variables set in the Docker image. In order for
	//  environment variables to be expanded, reference them by using the following
	//  syntax: <code>$(<var>VARIABLE_NAME</var>)</code> Note that this differs
	//  from Bash variable expansion, which does not use parentheses. If a variable
	//  cannot be resolved, the reference in the input string is used unchanged. To
	//  avoid variable expansion, you can escape this syntax with `$$`; for
	//  example: <code>$$(<var>VARIABLE_NAME</var>)</code> This field corresponds
	//  to the `args` field of the Kubernetes Containers [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.args
	Args []string `json:"args,omitempty"`

	// Immutable. List of environment variables to set in the container. After the
	//  container starts running, code running in the container can read these
	//  environment variables.
	//
	//  Additionally, the
	//  [command][google.cloud.aiplatform.v1.ModelContainerSpec.command] and
	//  [args][google.cloud.aiplatform.v1.ModelContainerSpec.args] fields can
	//  reference these variables. Later entries in this list can also reference
	//  earlier entries. For example, the following example sets the variable
	//  `VAR_2` to have the value `foo bar`:
	//
	//  ```json
	//  [
	//    {
	//      "name": "VAR_1",
	//      "value": "foo"
	//    },
	//    {
	//      "name": "VAR_2",
	//      "value": "$(VAR_1) bar"
	//    }
	//  ]
	//  ```
	//
	//  If you switch the order of the variables in the example, then the expansion
	//  does not occur.
	//
	//  This field corresponds to the `env` field of the Kubernetes Containers
	//  [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.env
	Env []EnvVar `json:"env,omitempty"`

	// Immutable. List of ports to expose from the container. Vertex AI sends any
	//  prediction requests that it receives to the first port on this list. Vertex
	//  AI also sends
	//  [liveness and health
	//  checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#liveness)
	//  to this port.
	//
	//  If you do not specify this field, it defaults to following value:
	//
	//  ```json
	//  [
	//    {
	//      "containerPort": 8080
	//    }
	//  ]
	//  ```
	//
	//  Vertex AI does not use ports other than the first one listed. This field
	//  corresponds to the `ports` field of the Kubernetes Containers
	//  [v1 core
	//  API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.ports
	Ports []Port `json:"ports,omitempty"`

	// Immutable. HTTP path on the container to send prediction requests to.
	//  Vertex AI forwards requests sent using
	//  [projects.locations.endpoints.predict][google.cloud.aiplatform.v1.PredictionService.Predict]
	//  to this path on the container's IP address and port. Vertex AI then returns
	//  the container's response in the API response.
	//
	//  For example, if you set this field to `/foo`, then when Vertex AI
	//  receives a prediction request, it forwards the request body in a POST
	//  request to the `/foo` path on the port of your container specified by the
	//  first value of this `ModelContainerSpec`'s
	//  [ports][google.cloud.aiplatform.v1.ModelContainerSpec.ports] field.
	//
	//  If you don't specify this field, it defaults to the following value when
	//  you [deploy this Model to an
	//  Endpoint][google.cloud.aiplatform.v1.EndpointService.DeployModel]:
	//  <code>/v1/endpoints/<var>ENDPOINT</var>/deployedModels/<var>DEPLOYED_MODEL</var>:predict</code>
	//  The placeholders in this value are replaced as follows:
	//
	//  * <var>ENDPOINT</var>: The last segment (following `endpoints/`)of the
	//    Endpoint.name][] field of the Endpoint where this Model has been
	//    deployed. (Vertex AI makes this value available to your container code
	//    as the [`AIP_ENDPOINT_ID` environment
	//   variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	//  * <var>DEPLOYED_MODEL</var>:
	//  [DeployedModel.id][google.cloud.aiplatform.v1.DeployedModel.id] of the
	//  `DeployedModel`.
	//    (Vertex AI makes this value available to your container code
	//    as the [`AIP_DEPLOYED_MODEL_ID` environment
	//    variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.predict_route
	PredictRoute *string `json:"predictRoute,omitempty"`

	// Immutable. HTTP path on the container to send health checks to. Vertex AI
	//  intermittently sends GET requests to this path on the container's IP
	//  address and port to check that the container is healthy. Read more about
	//  [health
	//  checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health).
	//
	//  For example, if you set this field to `/bar`, then Vertex AI
	//  intermittently sends a GET request to the `/bar` path on the port of your
	//  container specified by the first value of this `ModelContainerSpec`'s
	//  [ports][google.cloud.aiplatform.v1.ModelContainerSpec.ports] field.
	//
	//  If you don't specify this field, it defaults to the following value when
	//  you [deploy this Model to an
	//  Endpoint][google.cloud.aiplatform.v1.EndpointService.DeployModel]:
	//  <code>/v1/endpoints/<var>ENDPOINT</var>/deployedModels/<var>DEPLOYED_MODEL</var>:predict</code>
	//  The placeholders in this value are replaced as follows:
	//
	//  * <var>ENDPOINT</var>: The last segment (following `endpoints/`)of the
	//    Endpoint.name][] field of the Endpoint where this Model has been
	//    deployed. (Vertex AI makes this value available to your container code
	//    as the [`AIP_ENDPOINT_ID` environment
	//    variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	//  * <var>DEPLOYED_MODEL</var>:
	//  [DeployedModel.id][google.cloud.aiplatform.v1.DeployedModel.id] of the
	//  `DeployedModel`.
	//    (Vertex AI makes this value available to your container code as the
	//    [`AIP_DEPLOYED_MODEL_ID` environment
	//    variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.health_route
	HealthRoute *string `json:"healthRoute,omitempty"`

	// Immutable. List of ports to expose from the container. Vertex AI sends gRPC
	//  prediction requests that it receives to the first port on this list. Vertex
	//  AI also sends liveness and health checks to this port.
	//
	//  If you do not specify this field, gRPC requests to the container will be
	//  disabled.
	//
	//  Vertex AI does not use ports other than the first one listed. This field
	//  corresponds to the `ports` field of the Kubernetes Containers v1 core API.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.grpc_ports
	GrpcPorts []Port `json:"grpcPorts,omitempty"`

	// Immutable. Deployment timeout.
	//  Limit for deployment timeout is 2 hours.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.deployment_timeout
	DeploymentTimeout *string `json:"deploymentTimeout,omitempty"`

	// Immutable. The amount of the VM memory to reserve as the shared memory for
	//  the model in megabytes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.shared_memory_size_mb
	SharedMemorySizeMb *int64 `json:"sharedMemorySizeMb,omitempty"`

	// Immutable. Specification for Kubernetes startup probe.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.startup_probe
	StartupProbe *Probe `json:"startupProbe,omitempty"`

	// Immutable. Specification for Kubernetes readiness probe.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ModelContainerSpec.health_probe
	HealthProbe *Probe `json:"healthProbe,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Port
type Port struct {
	// The number of the port to expose on the pod's IP address.
	//  Must be a valid port number, between 1 and 65535 inclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Port.container_port
	ContainerPort *int32 `json:"containerPort,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PredictSchemata
type PredictSchemata struct {
	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  the format of a single instance, which are used in
	//  [PredictRequest.instances][google.cloud.aiplatform.v1.PredictRequest.instances],
	//  [ExplainRequest.instances][google.cloud.aiplatform.v1.ExplainRequest.instances]
	//  and
	//  [BatchPredictionJob.input_config][google.cloud.aiplatform.v1.BatchPredictionJob.input_config].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri
	InstanceSchemaURI *string `json:"instanceSchemaURI,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  the parameters of prediction and explanation via
	//  [PredictRequest.parameters][google.cloud.aiplatform.v1.PredictRequest.parameters],
	//  [ExplainRequest.parameters][google.cloud.aiplatform.v1.ExplainRequest.parameters]
	//  and
	//  [BatchPredictionJob.model_parameters][google.cloud.aiplatform.v1.BatchPredictionJob.model_parameters].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI, if no
	//  parameters are supported, then it is set to an empty string.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictSchemata.parameters_schema_uri
	ParametersSchemaURI *string `json:"parametersSchemaURI,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing
	//  the format of a single prediction produced by this Model, which are
	//  returned via
	//  [PredictResponse.predictions][google.cloud.aiplatform.v1.PredictResponse.predictions],
	//  [ExplainResponse.explanations][google.cloud.aiplatform.v1.ExplainResponse.explanations],
	//  and
	//  [BatchPredictionJob.output_config][google.cloud.aiplatform.v1.BatchPredictionJob.output_config].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML Models always have this field populated by Vertex AI.
	//  Note: The URI given on output will be immutable and probably different,
	//  including the URI scheme, than the one given on input. The output URI will
	//  point to a location where the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictSchemata.prediction_schema_uri
	PredictionSchemaURI *string `json:"predictionSchemaURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Probe
type Probe struct {
	// ExecAction probes the health of a container by executing a command.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.exec
	Exec *Probe_ExecAction `json:"exec,omitempty"`

	// How often (in seconds) to perform the probe. Default to 10 seconds.
	//  Minimum value is 1. Must be less than timeout_seconds.
	//
	//  Maps to Kubernetes probe argument 'periodSeconds'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.period_seconds
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`

	// Number of seconds after which the probe times out. Defaults to 1 second.
	//  Minimum value is 1. Must be greater or equal to period_seconds.
	//
	//  Maps to Kubernetes probe argument 'timeoutSeconds'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.timeout_seconds
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Probe.ExecAction
type Probe_ExecAction struct {
	// Command is the command line to execute inside the container, the working
	//  directory for the command is root ('/') in the container's filesystem.
	//  The command is simply exec'd, it is not run inside a shell, so
	//  traditional shell instructions ('|', etc) won't work. To use a shell, you
	//  need to explicitly call out to that shell. Exit status of 0 is treated as
	//  live/healthy and non-zero is unhealthy.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Probe.ExecAction.command
	Command []string `json:"command,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel
type PublisherModel struct {

	// Required. Indicates the open source category of the publisher model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.open_source_category
	OpenSourceCategory *string `json:"openSourceCategory,omitempty"`

	// Optional. Supported call-to-action options.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.supported_actions
	SupportedActions *PublisherModel_CallToAction `json:"supportedActions,omitempty"`

	// Optional. Additional information about the model's Frameworks.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.frameworks
	Frameworks []string `json:"frameworks,omitempty"`

	// Optional. Indicates the launch stage of the model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// Optional. Indicates the state of the model version.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.version_state
	VersionState *string `json:"versionState,omitempty"`

	// Optional. The schemata that describes formats of the PublisherModel's
	//  predictions and explanations as given and returned via
	//  [PredictionService.Predict][google.cloud.aiplatform.v1.PredictionService.Predict].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.predict_schemata
	PredictSchemata *PredictSchemata `json:"predictSchemata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction
type PublisherModel_CallToAction struct {
	// Optional. To view Rest API docs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.view_rest_api
	ViewRestApi *PublisherModel_CallToAction_ViewRestApi `json:"viewRestApi,omitempty"`

	// Optional. Open notebook of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_notebook
	OpenNotebook *PublisherModel_CallToAction_RegionalResourceReferences `json:"openNotebook,omitempty"`

	// Optional. Open notebooks of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_notebooks
	OpenNotebooks *PublisherModel_CallToAction_OpenNotebooks `json:"openNotebooks,omitempty"`

	// Optional. Create application using the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.create_application
	CreateApplication *PublisherModel_CallToAction_RegionalResourceReferences `json:"createApplication,omitempty"`

	// Optional. Open fine-tuning pipeline of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_fine_tuning_pipeline
	OpenFineTuningPipeline *PublisherModel_CallToAction_RegionalResourceReferences `json:"openFineTuningPipeline,omitempty"`

	// Optional. Open fine-tuning pipelines of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_fine_tuning_pipelines
	OpenFineTuningPipelines *PublisherModel_CallToAction_OpenFineTuningPipelines `json:"openFineTuningPipelines,omitempty"`

	// Optional. Open prompt-tuning pipeline of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_prompt_tuning_pipeline
	OpenPromptTuningPipeline *PublisherModel_CallToAction_RegionalResourceReferences `json:"openPromptTuningPipeline,omitempty"`

	// Optional. Open Genie / Playground.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_genie
	OpenGenie *PublisherModel_CallToAction_RegionalResourceReferences `json:"openGenie,omitempty"`

	// Optional. Deploy the PublisherModel to Vertex Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.deploy
	Deploy *PublisherModel_CallToAction_Deploy `json:"deploy,omitempty"`

	// Optional. Deploy PublisherModel to Google Kubernetes Engine.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.deploy_gke
	DeployGke *PublisherModel_CallToAction_DeployGke `json:"deployGke,omitempty"`

	// Optional. Open in Generation AI Studio.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_generation_ai_studio
	OpenGenerationAiStudio *PublisherModel_CallToAction_RegionalResourceReferences `json:"openGenerationAiStudio,omitempty"`

	// Optional. Request for access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.request_access
	RequestAccess *PublisherModel_CallToAction_RegionalResourceReferences `json:"requestAccess,omitempty"`

	// Optional. Open evaluation pipeline of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.open_evaluation_pipeline
	OpenEvaluationPipeline *PublisherModel_CallToAction_RegionalResourceReferences `json:"openEvaluationPipeline,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy
type PublisherModel_CallToAction_Deploy struct {
	// A description of resources that are dedicated to the DeployedModel,
	//  and that need a higher degree of manual configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.dedicated_resources
	DedicatedResources *DedicatedResources `json:"dedicatedResources,omitempty"`

	// A description of resources that to large degree are decided by Vertex
	//  AI, and require only a modest additional configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.automatic_resources
	AutomaticResources *AutomaticResources `json:"automaticResources,omitempty"`

	// The resource name of the shared DeploymentResourcePool to deploy on.
	//  Format:
	//  `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.shared_resources
	SharedResources *string `json:"sharedResources,omitempty"`

	// Optional. Default model display name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.model_display_name
	ModelDisplayName *string `json:"modelDisplayName,omitempty"`

	// Optional. Large model reference. When this is set, model_artifact_spec
	//  is not needed.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.large_model_reference
	LargeModelReference *LargeModelReference `json:"largeModelReference,omitempty"`

	// Optional. The specification of the container that is to be used when
	//  deploying this Model in Vertex AI. Not present for Large Models.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.container_spec
	ContainerSpec *ModelContainerSpec `json:"containerSpec,omitempty"`

	// Optional. The path to the directory containing the Model artifact and
	//  any of its supporting files.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.artifact_uri
	ArtifactURI *string `json:"artifactURI,omitempty"`

	// Optional. The name of the deploy task (e.g., "text to image
	//  generation").
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.deploy_task_name
	DeployTaskName *string `json:"deployTaskName,omitempty"`

	// Optional. Metadata information about this deployment config.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.deploy_metadata
	DeployMetadata *PublisherModel_CallToAction_Deploy_DeployMetadata `json:"deployMetadata,omitempty"`

	// Required. The title of the regional resource reference.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.title
	Title *string `json:"title,omitempty"`

	// Optional. The signed URI for ephemeral Cloud Storage access to model
	//  artifact.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.public_artifact_uri
	PublicArtifactURI *string `json:"publicArtifactURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.DeployMetadata
type PublisherModel_CallToAction_Deploy_DeployMetadata struct {
	// Optional. Labels for the deployment. For managing deployment config
	//  like verifying, source of deployment config, etc.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.DeployMetadata.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Sample request for deployed endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.Deploy.DeployMetadata.sample_request
	SampleRequest *string `json:"sampleRequest,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.DeployGke
type PublisherModel_CallToAction_DeployGke struct {
	// Optional. GKE deployment configuration in yaml format.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.DeployGke.gke_yaml_configs
	GkeYamlConfigs []string `json:"gkeYamlConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.OpenFineTuningPipelines
type PublisherModel_CallToAction_OpenFineTuningPipelines struct {
	// Required. Regional resource references to fine tuning pipelines.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.OpenFineTuningPipelines.fine_tuning_pipelines
	FineTuningPipelines []PublisherModel_CallToAction_RegionalResourceReferences `json:"fineTuningPipelines,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.OpenNotebooks
type PublisherModel_CallToAction_OpenNotebooks struct {
	// Required. Regional resource references to notebooks.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.OpenNotebooks.notebooks
	Notebooks []PublisherModel_CallToAction_RegionalResourceReferences `json:"notebooks,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.RegionalResourceReferences
type PublisherModel_CallToAction_RegionalResourceReferences struct {

	// TODO: unsupported map type with key string and value message


	// Required.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.RegionalResourceReferences.title
	Title *string `json:"title,omitempty"`

	// Optional. Title of the resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.RegionalResourceReferences.resource_title
	ResourceTitle *string `json:"resourceTitle,omitempty"`

	// Optional. Use case (CUJ) of the resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.RegionalResourceReferences.resource_use_case
	ResourceUseCase *string `json:"resourceUseCase,omitempty"`

	// Optional. Description of the resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.RegionalResourceReferences.resource_description
	ResourceDescription *string `json:"resourceDescription,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.CallToAction.ViewRestApi
type PublisherModel_CallToAction_ViewRestApi struct {
	// Required.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.ViewRestApi.documentations
	Documentations []PublisherModel_Documentation `json:"documentations,omitempty"`

	// Required. The title of the view rest API.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.CallToAction.ViewRestApi.title
	Title *string `json:"title,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.Documentation
type PublisherModel_Documentation struct {
	// Required. E.g., OVERVIEW, USE CASES, DOCUMENTATION, SDK & SAMPLES, JAVA,
	//  NODE.JS, etc..
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.Documentation.title
	Title *string `json:"title,omitempty"`

	// Required. Content of this piece of document (in Markdown format).
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.Documentation.content
	Content *string `json:"content,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel.ResourceReference
type PublisherModel_ResourceReference struct {
	// The URI of the resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.ResourceReference.uri
	URI *string `json:"uri,omitempty"`

	// The resource name of the Google Cloud resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.ResourceReference.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// Use case (CUJ) of the resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.ResourceReference.use_case
	UseCase *string `json:"useCase,omitempty"`

	// Description of the resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.ResourceReference.description
	Description *string `json:"description,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.PublisherModel
type PublisherModelObservedState struct {
	// Output only. The resource name of the PublisherModel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.name
	Name *string `json:"name,omitempty"`

	// Output only. Immutable. The version ID of the PublisherModel.
	//  A new version is committed when a new model version is uploaded under an
	//  existing model id. It is an auto-incrementing decimal number in string
	//  representation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.version_id
	VersionID *string `json:"versionID,omitempty"`

	// Optional. Output only. Immutable. Used to indicate this model has a
	//  publisher model and provide the template of the publisher model resource
	//  name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PublisherModel.publisher_model_template
	PublisherModelTemplate *string `json:"publisherModelTemplate,omitempty"`
}
