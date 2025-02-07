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

// +kcc:proto=google.cloud.aiplatform.v1.BigQueryDestination
type BigQueryDestination struct {
	// Required. BigQuery URI to a project or table, up to 2000 characters long.
	//
	//  When only the project is specified, the Dataset and Table is created.
	//  When the full table reference is specified, the Dataset must exist and
	//  table must not exist.
	//
	//  Accepted forms:
	//
	//  *  BigQuery path. For example:
	//  `bq://projectId` or `bq://projectId.bqDatasetId` or
	//  `bq://projectId.bqDatasetId.bqTableId`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BigQueryDestination.output_uri
	OutputURI *string `json:"outputURI,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.BlurBaselineConfig
type BlurBaselineConfig struct {
	// The standard deviation of the blur kernel for the blurred baseline. The
	//  same blurring parameter is used for both the height and the width
	//  dimension. If not set, the method defaults to the zero (i.e. black for
	//  images) baseline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.BlurBaselineConfig.max_blur_sigma
	MaxBlurSigma *float32 `json:"maxBlurSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ClientConnectionConfig
type ClientConnectionConfig struct {
	// Customizable online prediction request timeout.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ClientConnectionConfig.inference_timeout
	InferenceTimeout *string `json:"inferenceTimeout,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.DeployedModel
type DeployedModel struct {
	// A description of resources that are dedicated to the DeployedModel, and
	//  that need a higher degree of manual configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.dedicated_resources
	DedicatedResources *DedicatedResources `json:"dedicatedResources,omitempty"`

	// A description of resources that to large degree are decided by Vertex
	//  AI, and require only a modest additional configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.automatic_resources
	AutomaticResources *AutomaticResources `json:"automaticResources,omitempty"`

	// The resource name of the shared DeploymentResourcePool to deploy on.
	//  Format:
	//  `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.shared_resources
	SharedResources *string `json:"sharedResources,omitempty"`

	// Immutable. The ID of the DeployedModel. If not provided upon deployment,
	//  Vertex AI will generate a value for this ID.
	//
	//  This value should be 1-10 characters, and valid characters are `/[0-9]/`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.id
	ID *string `json:"id,omitempty"`

	// Required. The resource name of the Model that this is the deployment of.
	//  Note that the Model may be in a different location than the DeployedModel's
	//  Endpoint.
	//
	//  The resource name may contain version id or version alias to specify the
	//  version.
	//   Example: `projects/{project}/locations/{location}/models/{model}@2`
	//               or
	//             `projects/{project}/locations/{location}/models/{model}@golden`
	//  if no version is specified, the default version will be deployed.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.model
	Model *string `json:"model,omitempty"`

	// The display name of the DeployedModel. If not provided upon creation,
	//  the Model's display_name is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Explanation configuration for this DeployedModel.
	//
	//  When deploying a Model using
	//  [EndpointService.DeployModel][google.cloud.aiplatform.v1.EndpointService.DeployModel],
	//  this value overrides the value of
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec].
	//  All fields of
	//  [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	//  are optional in the request. If a field of
	//  [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	//  is not populated, the value of the same field of
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec]
	//  is inherited. If the corresponding
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec]
	//  is not populated, all fields of the
	//  [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec]
	//  will be used for the explanation configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.explanation_spec
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// If true, deploy the model without explainable feature, regardless the
	//  existence of
	//  [Model.explanation_spec][google.cloud.aiplatform.v1.Model.explanation_spec]
	//  or
	//  [explanation_spec][google.cloud.aiplatform.v1.DeployedModel.explanation_spec].
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.disable_explanations
	DisableExplanations *bool `json:"disableExplanations,omitempty"`

	// The service account that the DeployedModel's container runs as. Specify the
	//  email address of the service account. If this service account is not
	//  specified, the container runs as a service account that doesn't have access
	//  to the resource project.
	//
	//  Users deploying the Model must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// For custom-trained Models and AutoML Tabular Models, the container of the
	//  DeployedModel instances will send `stderr` and `stdout` streams to
	//  Cloud Logging by default. Please note that the logs incur cost,
	//  which are subject to [Cloud Logging
	//  pricing](https://cloud.google.com/logging/pricing).
	//
	//  User can disable container logging by setting this flag to true.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.disable_container_logging
	DisableContainerLogging *bool `json:"disableContainerLogging,omitempty"`

	// If true, online prediction access logs are sent to Cloud
	//  Logging.
	//  These logs are like standard server access logs, containing
	//  information like timestamp and latency for each prediction request.
	//
	//  Note that logs may incur a cost, especially if your project
	//  receives prediction requests at a high queries per second rate (QPS).
	//  Estimate your costs before enabling this option.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.enable_access_logging
	EnableAccessLogging *bool `json:"enableAccessLogging,omitempty"`

	// Configuration for faster model deployment.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.faster_deployment_config
	FasterDeploymentConfig *FasterDeploymentConfig `json:"fasterDeploymentConfig,omitempty"`

	// System labels to apply to Model Garden deployments.
	//  System labels are managed by Google for internal use only.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.system_labels
	SystemLabels map[string]string `json:"systemLabels,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedModel.Status
type DeployedModel_Status struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.Endpoint
type Endpoint struct {

	// Required. The display name of the Endpoint.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.description
	Description *string `json:"description,omitempty"`

	// TODO: unsupported map type with key string and value int32


	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Endpoints.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key spec for an Endpoint. If set, this
	//  Endpoint and all sub-resources of this Endpoint will be secured by
	//  this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. The full name of the Google Compute Engine
	//  [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks)
	//  to which the Endpoint should be peered.
	//
	//  Private services access must already be configured for the network. If left
	//  unspecified, the Endpoint is not peered with any network.
	//
	//  Only one of the fields,
	//  [network][google.cloud.aiplatform.v1.Endpoint.network] or
	//  [enable_private_service_connect][google.cloud.aiplatform.v1.Endpoint.enable_private_service_connect],
	//  can be set.
	//
	//  [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	//  `projects/{project}/global/networks/{network}`.
	//  Where `{project}` is a project number, as in `12345`, and `{network}` is
	//  network name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.network
	Network *string `json:"network,omitempty"`

	// Deprecated: If true, expose the Endpoint via private service connect.
	//
	//  Only one of the fields,
	//  [network][google.cloud.aiplatform.v1.Endpoint.network] or
	//  [enable_private_service_connect][google.cloud.aiplatform.v1.Endpoint.enable_private_service_connect],
	//  can be set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// Optional. Configuration for private service connect.
	//
	//  [network][google.cloud.aiplatform.v1.Endpoint.network] and
	//  [private_service_connect_config][google.cloud.aiplatform.v1.Endpoint.private_service_connect_config]
	//  are mutually exclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.private_service_connect_config
	PrivateServiceConnectConfig *PrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty"`

	// Configures the request-response logging for online prediction.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.predict_request_response_logging_config
	PredictRequestResponseLoggingConfig *PredictRequestResponseLoggingConfig `json:"predictRequestResponseLoggingConfig,omitempty"`

	// If true, the endpoint will be exposed through a dedicated
	//  DNS [Endpoint.dedicated_endpoint_dns]. Your request to the dedicated DNS
	//  will be isolated from other users' traffic and will have better performance
	//  and reliability.
	//  Note: Once you enabled dedicated endpoint, you won't be able to send
	//  request to the shared DNS {region}-aiplatform.googleapis.com. The
	//  limitation will be removed soon.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.dedicated_endpoint_enabled
	DedicatedEndpointEnabled *bool `json:"dedicatedEndpointEnabled,omitempty"`

	// Configurations that are applied to the endpoint for online prediction.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.client_connection_config
	ClientConnectionConfig *ClientConnectionConfig `json:"clientConnectionConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Examples
type Examples struct {
	// The Cloud Storage input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.example_gcs_source
	ExampleGcsSource *Examples_ExampleGcsSource `json:"exampleGcsSource,omitempty"`

	// The full configuration for the generated index, the semantics are the
	//  same as [metadata][google.cloud.aiplatform.v1.Index.metadata] and should
	//  match
	//  [NearestNeighborSearchConfig](https://cloud.google.com/vertex-ai/docs/explainable-ai/configuring-explanations-example-based#nearest-neighbor-search-config).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.nearest_neighbor_search_config
	NearestNeighborSearchConfig *Value `json:"nearestNeighborSearchConfig,omitempty"`

	// Simplified preset configuration, which automatically sets configuration
	//  values based on the desired query speed-precision trade-off and modality.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.presets
	Presets *Presets `json:"presets,omitempty"`

	// The number of neighbors to return when querying for examples.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.neighbor_count
	NeighborCount *int32 `json:"neighborCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Examples.ExampleGcsSource
type Examples_ExampleGcsSource struct {
	// The format in which instances are given, if not specified, assume it's
	//  JSONL format. Currently only JSONL format is supported.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.ExampleGcsSource.data_format
	DataFormat *string `json:"dataFormat,omitempty"`

	// The Cloud Storage location for the input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Examples.ExampleGcsSource.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata
type ExplanationMetadata struct {

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// Points to a YAML file stored on Google Cloud Storage describing the format
	//  of the [feature
	//  attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML tabular Models always have this field populated by Vertex AI.
	//  Note: The URI given on output may be different, including the URI scheme,
	//  than the one given on input. The output URI will point to a location where
	//  the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.feature_attributions_schema_uri
	FeatureAttributionsSchemaURI *string `json:"featureAttributionsSchemaURI,omitempty"`

	// Name of the source to generate embeddings for example based explanations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.latent_space_source
	LatentSpaceSource *string `json:"latentSpaceSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata
type ExplanationMetadata_InputMetadata struct {
	// Baseline inputs for this feature.
	//
	//  If no baseline is specified, Vertex AI chooses the baseline for this
	//  feature. If multiple baselines are specified, Vertex AI returns the
	//  average attributions across them in
	//  [Attribution.feature_attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions].
	//
	//  For Vertex AI-provided Tensorflow images (both 1.x and 2.x), the shape
	//  of each baseline must match the shape of the input tensor. If a scalar is
	//  provided, we broadcast to the same shape as the input tensor.
	//
	//  For custom images, the element of the baselines must be in the same
	//  format as the feature's input in the
	//  [instance][google.cloud.aiplatform.v1.ExplainRequest.instances][]. The
	//  schema of any single instance may be specified via Endpoint's
	//  DeployedModels' [Model's][google.cloud.aiplatform.v1.DeployedModel.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1.PredictSchemata.instance_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.input_baselines
	InputBaselines []Value `json:"inputBaselines,omitempty"`

	// Name of the input tensor for this feature. Required and is only
	//  applicable to Vertex AI-provided images for Tensorflow.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.input_tensor_name
	InputTensorName *string `json:"inputTensorName,omitempty"`

	// Defines how the feature is encoded into the input tensor. Defaults to
	//  IDENTITY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Modality of the feature. Valid values are: numeric, image. Defaults to
	//  numeric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.modality
	Modality *string `json:"modality,omitempty"`

	// The domain details of the input feature value. Like min/max, original
	//  mean or standard deviation if normalized.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.feature_value_domain
	FeatureValueDomain *ExplanationMetadata_InputMetadata_FeatureValueDomain `json:"featureValueDomain,omitempty"`

	// Specifies the index of the values of the input tensor.
	//  Required when the input tensor is a sparse representation. Refer to
	//  Tensorflow documentation for more details:
	//  https://www.tensorflow.org/api_docs/python/tf/sparse/SparseTensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.indices_tensor_name
	IndicesTensorName *string `json:"indicesTensorName,omitempty"`

	// Specifies the shape of the values of the input if the input is a sparse
	//  representation. Refer to Tensorflow documentation for more details:
	//  https://www.tensorflow.org/api_docs/python/tf/sparse/SparseTensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.dense_shape_tensor_name
	DenseShapeTensorName *string `json:"denseShapeTensorName,omitempty"`

	// A list of feature names for each index in the input tensor.
	//  Required when the input
	//  [InputMetadata.encoding][google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoding]
	//  is BAG_OF_FEATURES, BAG_OF_FEATURES_SPARSE, INDICATOR.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.index_feature_mapping
	IndexFeatureMapping []string `json:"indexFeatureMapping,omitempty"`

	// Encoded tensor is a transformation of the input tensor. Must be provided
	//  if choosing
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution]
	//  or [XRAI
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution]
	//  and the input tensor is not differentiable.
	//
	//  An encoded tensor is generated if the input tensor is encoded by a lookup
	//  table.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoded_tensor_name
	EncodedTensorName *string `json:"encodedTensorName,omitempty"`

	// A list of baselines for the encoded tensor.
	//
	//  The shape of each baseline should match the shape of the encoded tensor.
	//  If a scalar is provided, Vertex AI broadcasts to the same shape as the
	//  encoded tensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.encoded_baselines
	EncodedBaselines []Value `json:"encodedBaselines,omitempty"`

	// Visualization configurations for image explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.visualization
	Visualization *ExplanationMetadata_InputMetadata_Visualization `json:"visualization,omitempty"`

	// Name of the group that the input belongs to. Features with the same group
	//  name will be treated as one feature when computing attributions. Features
	//  grouped together can have different shapes in value. If provided, there
	//  will be one single attribution generated in
	//  [Attribution.feature_attributions][google.cloud.aiplatform.v1.Attribution.feature_attributions],
	//  keyed by the group name.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.group_name
	GroupName *string `json:"groupName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain
type ExplanationMetadata_InputMetadata_FeatureValueDomain struct {
	// The minimum permissible value for this feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.min_value
	MinValue *float32 `json:"minValue,omitempty"`

	// The maximum permissible value for this feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.max_value
	MaxValue *float32 `json:"maxValue,omitempty"`

	// If this input feature has been normalized to a mean value of 0,
	//  the original_mean specifies the mean value of the domain prior to
	//  normalization.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.original_mean
	OriginalMean *float32 `json:"originalMean,omitempty"`

	// If this input feature has been normalized to a standard deviation of
	//  1.0, the original_stddev specifies the standard deviation of the domain
	//  prior to normalization.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.FeatureValueDomain.original_stddev
	OriginalStddev *float32 `json:"originalStddev,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization
type ExplanationMetadata_InputMetadata_Visualization struct {
	// Type of the image visualization. Only applicable to
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution].
	//  OUTLINES shows regions of attribution, while PIXELS shows per-pixel
	//  attribution. Defaults to OUTLINES.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.type
	Type *string `json:"type,omitempty"`

	// Whether to only highlight pixels with positive contributions, negative
	//  or both. Defaults to POSITIVE.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.polarity
	Polarity *string `json:"polarity,omitempty"`

	// The color scheme used for the highlighted areas.
	//
	//  Defaults to PINK_GREEN for
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution],
	//  which shows positive attributions in green and negative in pink.
	//
	//  Defaults to VIRIDIS for
	//  [XRAI
	//  attribution][google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution],
	//  which highlights the most influential regions in yellow and the least
	//  influential in blue.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.color_map
	ColorMap *string `json:"colorMap,omitempty"`

	// Excludes attributions above the specified percentile from the
	//  highlighted areas. Using the clip_percent_upperbound and
	//  clip_percent_lowerbound together can be useful for filtering out noise
	//  and making it easier to see areas of strong attribution. Defaults to
	//  99.9.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.clip_percent_upperbound
	ClipPercentUpperbound *float32 `json:"clipPercentUpperbound,omitempty"`

	// Excludes attributions below the specified percentile, from the
	//  highlighted areas. Defaults to 62.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.clip_percent_lowerbound
	ClipPercentLowerbound *float32 `json:"clipPercentLowerbound,omitempty"`

	// How the original image is displayed in the visualization.
	//  Adjusting the overlay can help increase visual clarity if the original
	//  image makes it difficult to view the visualization. Defaults to NONE.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.InputMetadata.Visualization.overlay_type
	OverlayType *string `json:"overlayType,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata
type ExplanationMetadata_OutputMetadata struct {
	// Static mapping between the index and display name.
	//
	//  Use this if the outputs are a deterministic n-dimensional array, e.g. a
	//  list of scores of all the classes in a pre-defined order for a
	//  multi-classification Model. It's not feasible if the outputs are
	//  non-deterministic, e.g. the Model produces top-k classes or sort the
	//  outputs by their values.
	//
	//  The shape of the value must be an n-dimensional array of strings. The
	//  number of dimensions must match that of the outputs to be explained.
	//  The
	//  [Attribution.output_display_name][google.cloud.aiplatform.v1.Attribution.output_display_name]
	//  is populated by locating in the mapping with
	//  [Attribution.output_index][google.cloud.aiplatform.v1.Attribution.output_index].
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata.index_display_name_mapping
	IndexDisplayNameMapping *Value `json:"indexDisplayNameMapping,omitempty"`

	// Specify a field name in the prediction to look for the display name.
	//
	//  Use this if the prediction contains the display names for the outputs.
	//
	//  The display names in the prediction must have the same shape of the
	//  outputs, so that it can be located by
	//  [Attribution.output_index][google.cloud.aiplatform.v1.Attribution.output_index]
	//  for a specific output.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata.display_name_mapping_key
	DisplayNameMappingKey *string `json:"displayNameMappingKey,omitempty"`

	// Name of the output tensor. Required and is only applicable to Vertex
	//  AI provided images for Tensorflow.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationMetadata.OutputMetadata.output_tensor_name
	OutputTensorName *string `json:"outputTensorName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationParameters
type ExplanationParameters struct {
	// An attribution method that approximates Shapley values for features that
	//  contribute to the label being predicted. A sampling strategy is used to
	//  approximate the value rather than considering all subsets of features.
	//  Refer to this paper for model details: https://arxiv.org/abs/1306.4265.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.sampled_shapley_attribution
	SampledShapleyAttribution *SampledShapleyAttribution `json:"sampledShapleyAttribution,omitempty"`

	// An attribution method that computes Aumann-Shapley values taking
	//  advantage of the model's fully differentiable structure. Refer to this
	//  paper for more details: https://arxiv.org/abs/1703.01365
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution
	IntegratedGradientsAttribution *IntegratedGradientsAttribution `json:"integratedGradientsAttribution,omitempty"`

	// An attribution method that redistributes Integrated Gradients
	//  attribution to segmented regions, taking advantage of the model's fully
	//  differentiable structure. Refer to this paper for
	//  more details: https://arxiv.org/abs/1906.02825
	//
	//  XRAI currently performs better on natural images, like a picture of a
	//  house or an animal. If the images are taken in artificial environments,
	//  like a lab or manufacturing line, or from diagnostic equipment, like
	//  x-rays or quality-control cameras, use Integrated Gradients instead.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution
	XraiAttribution *XraiAttribution `json:"xraiAttribution,omitempty"`

	// Example-based explanations that returns the nearest neighbors from the
	//  provided dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.examples
	Examples *Examples `json:"examples,omitempty"`

	// If populated, returns attributions for top K indices of outputs
	//  (defaults to 1). Only applies to Models that predicts more than one outputs
	//  (e,g, multi-class Models). When set to -1, returns explanations for all
	//  outputs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.top_k
	TopK *int32 `json:"topK,omitempty"`

	// If populated, only returns attributions that have
	//  [output_index][google.cloud.aiplatform.v1.Attribution.output_index]
	//  contained in output_indices. It must be an ndarray of integers, with the
	//  same shape of the output it's explaining.
	//
	//  If not populated, returns attributions for
	//  [top_k][google.cloud.aiplatform.v1.ExplanationParameters.top_k] indices of
	//  outputs. If neither top_k nor output_indices is populated, returns the
	//  argmax index of the outputs.
	//
	//  Only applicable to Models that predict multiple outputs (e,g, multi-class
	//  Models that predict multiple classes).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.output_indices
	OutputIndices *ListValue `json:"outputIndices,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationSpec
type ExplanationSpec struct {
	// Required. Parameters that configure explaining of the Model's predictions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationSpec.parameters
	Parameters *ExplanationParameters `json:"parameters,omitempty"`

	// Optional. Metadata describing the Model's input and output for explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationSpec.metadata
	Metadata *ExplanationMetadata `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FasterDeploymentConfig
type FasterDeploymentConfig struct {
	// If true, enable fast tryout feature for this deployed model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FasterDeploymentConfig.fast_tryout_enabled
	FastTryoutEnabled *bool `json:"fastTryoutEnabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureNoiseSigma
type FeatureNoiseSigma struct {
	// Noise sigma per feature. No noise is added to features that are not set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureNoiseSigma.noise_sigma
	NoiseSigma []FeatureNoiseSigma_NoiseSigmaForFeature `json:"noiseSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureNoiseSigma.NoiseSigmaForFeature
type FeatureNoiseSigma_NoiseSigmaForFeature struct {
	// The name of the input feature for which noise sigma is provided. The
	//  features are defined in
	//  [explanation metadata
	//  inputs][google.cloud.aiplatform.v1.ExplanationMetadata.inputs].
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureNoiseSigma.NoiseSigmaForFeature.name
	Name *string `json:"name,omitempty"`

	// This represents the standard deviation of the Gaussian kernel that will
	//  be used to add noise to the feature prior to computing gradients. Similar
	//  to [noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma]
	//  but represents the noise added to the current feature. Defaults to 0.1.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureNoiseSigma.NoiseSigmaForFeature.sigma
	Sigma *float32 `json:"sigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.GcsSource
type GcsSource struct {
	// Required. Google Cloud Storage URI(-s) to the input file(s). May contain
	//  wildcards. For more information on wildcards, see
	//  https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames.
	// +kcc:proto:field=google.cloud.aiplatform.v1.GcsSource.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.IntegratedGradientsAttribution
type IntegratedGradientsAttribution struct {
	// Required. The number of steps for approximating the path integral.
	//  A good value to start is 50 and gradually increase until the
	//  sum to diff property is within the desired error range.
	//
	//  Valid range of its value is [1, 100], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.IntegratedGradientsAttribution.step_count
	StepCount *int32 `json:"stepCount,omitempty"`

	// Config for SmoothGrad approximation of gradients.
	//
	//  When enabled, the gradients are approximated by averaging the gradients
	//  from noisy samples in the vicinity of the inputs. Adding
	//  noise can help improve the computed gradients. Refer to this paper for more
	//  details: https://arxiv.org/pdf/1706.03825.pdf
	// +kcc:proto:field=google.cloud.aiplatform.v1.IntegratedGradientsAttribution.smooth_grad_config
	SmoothGradConfig *SmoothGradConfig `json:"smoothGradConfig,omitempty"`

	// Config for IG with blur baseline.
	//
	//  When enabled, a linear path from the maximally blurred image to the input
	//  image is created. Using a blurred baseline instead of zero (black image) is
	//  motivated by the BlurIG approach explained here:
	//  https://arxiv.org/abs/2004.03383
	// +kcc:proto:field=google.cloud.aiplatform.v1.IntegratedGradientsAttribution.blur_baseline_config
	BlurBaselineConfig *BlurBaselineConfig `json:"blurBaselineConfig,omitempty"`
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

// +kcc:proto=google.cloud.aiplatform.v1.PredictRequestResponseLoggingConfig
type PredictRequestResponseLoggingConfig struct {
	// If logging is enabled or not.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictRequestResponseLoggingConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Percentage of requests to be logged, expressed as a fraction in
	//  range(0,1].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictRequestResponseLoggingConfig.sampling_rate
	SamplingRate *float64 `json:"samplingRate,omitempty"`

	// BigQuery table for logging.
	//  If only given a project, a new dataset will be created with name
	//  `logging_<endpoint-display-name>_<endpoint-id>` where
	//  <endpoint-display-name> will be made BigQuery-dataset-name compatible (e.g.
	//  most special characters will become underscores). If no table name is
	//  given, a new table will be created with name `request_response_logging`
	// +kcc:proto:field=google.cloud.aiplatform.v1.PredictRequestResponseLoggingConfig.bigquery_destination
	BigqueryDestination *BigQueryDestination `json:"bigqueryDestination,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Presets
type Presets struct {
	// Preset option controlling parameters for speed-precision trade-off when
	//  querying for examples. If omitted, defaults to `PRECISE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Presets.query
	Query *string `json:"query,omitempty"`

	// The modality of the uploaded model, which automatically configures the
	//  distance measurement and feature normalization for the underlying example
	//  index and queries. If your model does not precisely fit one of these types,
	//  it is okay to choose the closest type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Presets.modality
	Modality *string `json:"modality,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateEndpoints
type PrivateEndpoints struct {
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

// +kcc:proto=google.cloud.aiplatform.v1.SampledShapleyAttribution
type SampledShapleyAttribution struct {
	// Required. The number of feature permutations to consider when approximating
	//  the Shapley values.
	//
	//  Valid range of its value is [1, 50], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SampledShapleyAttribution.path_count
	PathCount *int32 `json:"pathCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.SmoothGradConfig
type SmoothGradConfig struct {
	// This is a single float value and will be used to add noise to all the
	//  features. Use this field when all features are normalized to have the
	//  same distribution: scale to range [0, 1], [-1, 1] or z-scoring, where
	//  features are normalized to have 0-mean and 1-variance. Learn more about
	//  [normalization](https://developers.google.com/machine-learning/data-prep/transform/normalization).
	//
	//  For best results the recommended value is about 10% - 20% of the standard
	//  deviation of the input feature. Refer to section 3.2 of the SmoothGrad
	//  paper: https://arxiv.org/pdf/1706.03825.pdf. Defaults to 0.1.
	//
	//  If the distribution is different per feature, set
	//  [feature_noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.feature_noise_sigma]
	//  instead for each feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma
	NoiseSigma *float32 `json:"noiseSigma,omitempty"`

	// This is similar to
	//  [noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma],
	//  but provides additional flexibility. A separate noise sigma can be
	//  provided for each feature, which is useful if their distributions are
	//  different. No noise is added to features that are not set. If this field
	//  is unset,
	//  [noise_sigma][google.cloud.aiplatform.v1.SmoothGradConfig.noise_sigma]
	//  will be used for all features.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SmoothGradConfig.feature_noise_sigma
	FeatureNoiseSigma *FeatureNoiseSigma `json:"featureNoiseSigma,omitempty"`

	// The number of gradient samples to use for
	//  approximation. The higher this number, the more accurate the gradient
	//  is, but the runtime complexity increases by this factor as well.
	//  Valid range of its value is [1, 50]. Defaults to 3.
	// +kcc:proto:field=google.cloud.aiplatform.v1.SmoothGradConfig.noisy_sample_count
	NoisySampleCount *int32 `json:"noisySampleCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.XraiAttribution
type XraiAttribution struct {
	// Required. The number of steps for approximating the path integral.
	//  A good value to start is 50 and gradually increase until the
	//  sum to diff property is met within the desired error range.
	//
	//  Valid range of its value is [1, 100], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.XraiAttribution.step_count
	StepCount *int32 `json:"stepCount,omitempty"`

	// Config for SmoothGrad approximation of gradients.
	//
	//  When enabled, the gradients are approximated by averaging the gradients
	//  from noisy samples in the vicinity of the inputs. Adding
	//  noise can help improve the computed gradients. Refer to this paper for more
	//  details: https://arxiv.org/pdf/1706.03825.pdf
	// +kcc:proto:field=google.cloud.aiplatform.v1.XraiAttribution.smooth_grad_config
	SmoothGradConfig *SmoothGradConfig `json:"smoothGradConfig,omitempty"`

	// Config for XRAI with blur baseline.
	//
	//  When enabled, a linear path from the maximally blurred image to the input
	//  image is created. Using a blurred baseline instead of zero (black image) is
	//  motivated by the BlurIG approach explained here:
	//  https://arxiv.org/abs/2004.03383
	// +kcc:proto:field=google.cloud.aiplatform.v1.XraiAttribution.blur_baseline_config
	BlurBaselineConfig *BlurBaselineConfig `json:"blurBaselineConfig,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedModel
type DeployedModelObservedState struct {
	// Output only. The version ID of the model that is deployed.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.model_version_id
	ModelVersionID *string `json:"modelVersionID,omitempty"`

	// Output only. Timestamp when the DeployedModel was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Provide paths for users to send predict/explain/health
	//  requests directly to the deployed model services running on Cloud via
	//  private services access. This field is populated if
	//  [network][google.cloud.aiplatform.v1.Endpoint.network] is configured.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.private_endpoints
	PrivateEndpoints *PrivateEndpoints `json:"privateEndpoints,omitempty"`

	// Output only. Runtime status of the deployed model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.status
	Status *DeployedModel_Status `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DeployedModel.Status
type DeployedModel_StatusObservedState struct {
	// Output only. The latest deployed model's status message (if any).
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.Status.message
	Message *string `json:"message,omitempty"`

	// Output only. The time at which the status was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.Status.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// Output only. The number of available replicas of the deployed model.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModel.Status.available_replica_count
	AvailableReplicaCount *int32 `json:"availableReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Endpoint
type EndpointObservedState struct {
	// Output only. The resource name of the Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.name
	Name *string `json:"name,omitempty"`

	// Output only. The models deployed in this Endpoint.
	//  To add or remove DeployedModels use
	//  [EndpointService.DeployModel][google.cloud.aiplatform.v1.EndpointService.DeployModel]
	//  and
	//  [EndpointService.UndeployModel][google.cloud.aiplatform.v1.EndpointService.UndeployModel]
	//  respectively.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.deployed_models
	DeployedModels []DeployedModel `json:"deployedModels,omitempty"`

	// Output only. Timestamp when this Endpoint was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Endpoint was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Configuration for private service connect.
	//
	//  [network][google.cloud.aiplatform.v1.Endpoint.network] and
	//  [private_service_connect_config][google.cloud.aiplatform.v1.Endpoint.private_service_connect_config]
	//  are mutually exclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.private_service_connect_config
	PrivateServiceConnectConfig *PrivateServiceConnectConfigObservedState `json:"privateServiceConnectConfig,omitempty"`

	// Output only. Resource name of the Model Monitoring job associated with this
	//  Endpoint if monitoring is enabled by
	//  [JobService.CreateModelDeploymentMonitoringJob][google.cloud.aiplatform.v1.JobService.CreateModelDeploymentMonitoringJob].
	//  Format:
	//  `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.model_deployment_monitoring_job
	ModelDeploymentMonitoringJob *string `json:"modelDeploymentMonitoringJob,omitempty"`

	// Output only. DNS of the dedicated endpoint. Will only be populated if
	//  dedicated_endpoint_enabled is true.
	//  Format:
	//  `https://{endpoint_id}.{region}-{project_number}.prediction.vertexai.goog`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.dedicated_endpoint_dns
	DedicatedEndpointDns *string `json:"dedicatedEndpointDns,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Endpoint.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateEndpoints
type PrivateEndpointsObservedState struct {
	// Output only. Http(s) path to send prediction requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateEndpoints.predict_http_uri
	PredictHTTPURI *string `json:"predictHTTPURI,omitempty"`

	// Output only. Http(s) path to send explain requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateEndpoints.explain_http_uri
	ExplainHTTPURI *string `json:"explainHTTPURI,omitempty"`

	// Output only. Http(s) path to send health check requests.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateEndpoints.health_http_uri
	HealthHTTPURI *string `json:"healthHTTPURI,omitempty"`

	// Output only. The name of the service attachment resource. Populated if
	//  private service connect is enabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateEndpoints.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PrivateServiceConnectConfig
type PrivateServiceConnectConfigObservedState struct {
	// Output only. The name of the generated service attachment resource.
	//  This is only populated if the endpoint is deployed with
	//  PrivateServiceConnect.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PrivateServiceConnectConfig.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}
