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


// +kcc:proto=google.cloud.visionai.v1.AIEnabledDevicesInputConfig
type AIEnabledDevicesInputConfig struct {
}

// +kcc:proto=google.cloud.visionai.v1.Application
type Application struct {
	// name of resource
	// +kcc:proto:field=google.cloud.visionai.v1.Application.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.visionai.v1.Application.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A user friendly display name for the solution.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description for this application.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.description
	Description *string `json:"description,omitempty"`

	// Application graph configuration.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.application_configs
	ApplicationConfigs *ApplicationConfigs `json:"applicationConfigs,omitempty"`

	// Billing mode of the application.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.billing_mode
	BillingMode *string `json:"billingMode,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo
type Application_ApplicationRuntimeInfo struct {
	// Timestamp when the engine be deployed
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.deploy_time
	DeployTime *string `json:"deployTime,omitempty"`

	// Globally created resources like warehouse dataschemas.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.global_output_resources
	GlobalOutputResources []Application_ApplicationRuntimeInfo_GlobalOutputResource `json:"globalOutputResources,omitempty"`

	// Monitoring-related configuration for this application.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.monitoring_config
	MonitoringConfig *Application_ApplicationRuntimeInfo_MonitoringConfig `json:"monitoringConfig,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.GlobalOutputResource
type Application_ApplicationRuntimeInfo_GlobalOutputResource struct {
	// The full resource name of the outputted resources.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.GlobalOutputResource.output_resource
	OutputResource *string `json:"outputResource,omitempty"`

	// The name of graph node who produces the output resource name.
	//  For example:
	//  output_resource:
	//  /projects/123/locations/us-central1/corpora/my-corpus/dataSchemas/my-schema
	//  producer_node: occupancy-count
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.GlobalOutputResource.producer_node
	ProducerNode *string `json:"producerNode,omitempty"`

	// The key of the output resource, it has to be unique within the same
	//  producer node. One producer node can output several output resources,
	//  the key can be used to match corresponding output resources.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.GlobalOutputResource.key
	Key *string `json:"key,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.MonitoringConfig
type Application_ApplicationRuntimeInfo_MonitoringConfig struct {
	// Whether this application has monitoring enabled.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.ApplicationRuntimeInfo.MonitoringConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ApplicationConfigs
type ApplicationConfigs struct {
	// A list of nodes  in the application graph.
	// +kcc:proto:field=google.cloud.visionai.v1.ApplicationConfigs.nodes
	Nodes []Node `json:"nodes,omitempty"`

	// Event-related configuration for this application.
	// +kcc:proto:field=google.cloud.visionai.v1.ApplicationConfigs.event_delivery_config
	EventDeliveryConfig *ApplicationConfigs_EventDeliveryConfig `json:"eventDeliveryConfig,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ApplicationConfigs.EventDeliveryConfig
type ApplicationConfigs_EventDeliveryConfig struct {
	// The delivery channel for the event notification, only pub/sub topic is
	//  supported now.
	//  Example channel:
	//  [//pubsub.googleapis.com/projects/visionai-testing-stable/topics/test-topic]
	// +kcc:proto:field=google.cloud.visionai.v1.ApplicationConfigs.EventDeliveryConfig.channel
	Channel *string `json:"channel,omitempty"`

	// The expected delivery interval for the same event. The same event won't
	//  be notified multiple times during this internal event that it is
	//  happening multiple times during the period of time.The same event is
	//  identified by <event_id, app_platform_metadata>.
	// +kcc:proto:field=google.cloud.visionai.v1.ApplicationConfigs.EventDeliveryConfig.minimal_delivery_interval
	MinimalDeliveryInterval *string `json:"minimalDeliveryInterval,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AutoscalingMetricSpec
type AutoscalingMetricSpec struct {
	// Required. The resource metric name.
	//  Supported metrics:
	//
	//  * For Online Prediction:
	//  * `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle`
	//  * `aiplatform.googleapis.com/prediction/online/cpu/utilization`
	// +kcc:proto:field=google.cloud.visionai.v1.AutoscalingMetricSpec.metric_name
	MetricName *string `json:"metricName,omitempty"`

	// The target resource utilization in percentage (1% - 100%) for the given
	//  metric; once the real usage deviates from the target by a certain
	//  percentage, the machine replicas change. The default value is 60
	//  (representing 60%) if not provided.
	// +kcc:proto:field=google.cloud.visionai.v1.AutoscalingMetricSpec.target
	Target *int32 `json:"target,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.BigQueryConfig
type BigQueryConfig struct {
	// BigQuery table resource for Vision AI Platform to ingest annotations to.
	// +kcc:proto:field=google.cloud.visionai.v1.BigQueryConfig.table
	Table *string `json:"table,omitempty"`

	// Data Schema
	//  By default, Vision AI Application will try to write annotations to the
	//  target BigQuery table using the following schema:
	//
	//  ingestion_time: TIMESTAMP, the ingestion time of the original data.
	//
	//  application: STRING, name of the application which produces the annotation.
	//
	//  instance: STRING, Id of the instance which produces the annotation.
	//
	//  node: STRING, name of the application graph node which produces the
	//  annotation.
	//
	//  annotation: STRING or JSON, the actual annotation protobuf will be
	//  converted to json string with bytes field as 64 encoded string. It can be
	//  written to both String or Json type column.
	//
	//  To forward annotation data to an existing BigQuery table, customer needs to
	//  make sure the compatibility of the schema.
	//  The map maps application node name to its corresponding cloud function
	//  endpoint to transform the annotations directly to the
	//  google.cloud.bigquery.storage.v1.AppendRowsRequest (only avro_rows or
	//  proto_rows should be set). If configured, annotations produced by
	//  corresponding application node will sent to the Cloud Function at first
	//  before be forwarded to BigQuery.
	//
	//  If the default table schema doesn't fit, customer is able to transform the
	//  annotation output from Vision AI Application to arbitrary BigQuery table
	//  schema with CloudFunction.
	//  * The cloud function will receive AppPlatformCloudFunctionRequest where
	//  the annotations field will be the json format of Vision AI annotation.
	//  * The cloud function should return AppPlatformCloudFunctionResponse with
	//  AppendRowsRequest stored in the annotations field.
	//  * To drop the annotation, simply clear the annotations field in the
	//  returned AppPlatformCloudFunctionResponse.
	// +kcc:proto:field=google.cloud.visionai.v1.BigQueryConfig.cloud_function_mapping
	CloudFunctionMapping map[string]string `json:"cloudFunctionMapping,omitempty"`

	// If true, App Platform will create the BigQuery DataSet and the
	//  BigQuery Table with default schema if the specified table doesn't exist.
	//  This doesn't work if any cloud function customized schema is specified
	//  since the system doesn't know your desired schema.
	//  JSON column will be used in the default table created by App Platform.
	// +kcc:proto:field=google.cloud.visionai.v1.BigQueryConfig.create_default_table_if_not_exists
	CreateDefaultTableIfNotExists *bool `json:"createDefaultTableIfNotExists,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.DedicatedResources
type DedicatedResources struct {
	// Required. Immutable. The specification of a single machine used by the
	//  prediction.
	// +kcc:proto:field=google.cloud.visionai.v1.DedicatedResources.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Required. Immutable. The minimum number of machine replicas this
	//  DeployedModel will be always deployed on. This value must be greater than
	//  or equal to 1.
	//
	//  If traffic against the DeployedModel increases, it may dynamically be
	//  deployed onto more replicas, and as traffic decreases, some of these extra
	//  replicas may be freed.
	// +kcc:proto:field=google.cloud.visionai.v1.DedicatedResources.min_replica_count
	MinReplicaCount *int32 `json:"minReplicaCount,omitempty"`

	// Immutable. The maximum number of replicas this DeployedModel may be
	//  deployed on when the traffic against it increases. If the requested value
	//  is too large, the deployment will error, but if deployment succeeds then
	//  the ability to scale the model to that many replicas is guaranteed (barring
	//  service outages). If traffic against the DeployedModel increases beyond
	//  what its replicas at maximum may handle, a portion of the traffic will be
	//  dropped. If this value is not provided, will use
	//  [min_replica_count][google.cloud.visionai.v1.DedicatedResources.min_replica_count]
	//  as the default value.
	//
	//  The value of this field impacts the charge against Vertex CPU and GPU
	//  quotas. Specifically, you will be charged for max_replica_count *
	//  number of cores in the selected machine type) and (max_replica_count *
	//  number of GPUs per replica in the selected machine type).
	// +kcc:proto:field=google.cloud.visionai.v1.DedicatedResources.max_replica_count
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`

	// Immutable. The metric specifications that overrides a resource
	//  utilization metric (CPU utilization, accelerator's duty cycle, and so on)
	//  target value (default to 60 if not set). At most one entry is allowed per
	//  metric.
	//
	//  If
	//  [machine_spec.accelerator_count][google.cloud.visionai.v1.MachineSpec.accelerator_count]
	//  is above 0, the autoscaling will be based on both CPU utilization and
	//  accelerator's duty cycle metrics and scale up when either metrics exceeds
	//  its target value while scale down if both metrics are under their target
	//  value. The default target value is 60 for both metrics.
	//
	//  If
	//  [machine_spec.accelerator_count][google.cloud.visionai.v1.MachineSpec.accelerator_count]
	//  is 0, the autoscaling will be based on CPU utilization metric only with
	//  default target value 60 if not explicitly set.
	//
	//  For example, in the case of Online Prediction, if you want to override
	//  target CPU utilization to 80, you should set
	//  [autoscaling_metric_specs.metric_name][google.cloud.visionai.v1.AutoscalingMetricSpec.metric_name]
	//  to `aiplatform.googleapis.com/prediction/online/cpu/utilization` and
	//  [autoscaling_metric_specs.target][google.cloud.visionai.v1.AutoscalingMetricSpec.target]
	//  to `80`.
	// +kcc:proto:field=google.cloud.visionai.v1.DedicatedResources.autoscaling_metric_specs
	AutoscalingMetricSpecs []AutoscalingMetricSpec `json:"autoscalingMetricSpecs,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.GcsOutputConfig
type GcsOutputConfig struct {
	// The Cloud Storage path for Vision AI Platform to ingest annotations to.
	// +kcc:proto:field=google.cloud.visionai.v1.GcsOutputConfig.gcs_path
	GcsPath *string `json:"gcsPath,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.GeneralObjectDetectionConfig
type GeneralObjectDetectionConfig struct {
}

// +kcc:proto=google.cloud.visionai.v1.MachineSpec
type MachineSpec struct {
	// Immutable. The type of the machine.
	//
	//  See the [list of machine types supported for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	//  See the [list of machine types supported for custom
	//  training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	//  For [DeployedModel][] this field is optional, and the default
	//  value is `n1-standard-2`. For [BatchPredictionJob][] or as part of
	//  [WorkerPoolSpec][] this field is required.
	// +kcc:proto:field=google.cloud.visionai.v1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine
	//  as per
	//  [accelerator_count][google.cloud.visionai.v1.MachineSpec.accelerator_count].
	// +kcc:proto:field=google.cloud.visionai.v1.MachineSpec.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.visionai.v1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.MediaWarehouseConfig
type MediaWarehouseConfig struct {
	// Resource name of the Media Warehouse corpus.
	//  Format:
	//  projects/${project_id}/locations/${location_id}/corpora/${corpus_id}
	// +kcc:proto:field=google.cloud.visionai.v1.MediaWarehouseConfig.corpus
	Corpus *string `json:"corpus,omitempty"`

	// Deprecated.
	// +kcc:proto:field=google.cloud.visionai.v1.MediaWarehouseConfig.region
	Region *string `json:"region,omitempty"`

	// The duration for which all media assets, associated metadata, and search
	//  documents can exist.
	// +kcc:proto:field=google.cloud.visionai.v1.MediaWarehouseConfig.ttl
	Ttl *string `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Node
type Node struct {
	// By default, the output of the node will only be available to downstream
	//  nodes. To consume the direct output from the application node, the output
	//  must be sent to Vision AI Streams at first.
	//
	//  By setting output_all_output_channels_to_stream to true, App Platform
	//  will automatically send all the outputs of the current node to Vision AI
	//  Stream resources (one stream per output channel). The output stream
	//  resource will be created by App Platform automatically during deployment
	//  and deleted after application un-deployment.
	//  Note that this config applies to all the Application Instances.
	//
	//  The output stream can be override at instance level by
	//  configuring the `output_resources` section of Instance resource.
	//  `producer_node` should be current node, `output_resource_binding` should
	//  be the output channel name (or leave it blank if there is only 1 output
	//  channel of the processor) and `output_resource` should be the target
	//  output stream.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.output_all_output_channels_to_stream
	OutputAllOutputChannelsToStream *bool `json:"outputAllOutputChannelsToStream,omitempty"`

	// Required. A unique name for the node.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.name
	Name *string `json:"name,omitempty"`

	// A user friendly display name for the node.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Node config.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.node_config
	NodeConfig *ProcessorConfig `json:"nodeConfig,omitempty"`

	// Processor name refer to the chosen processor resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.processor
	Processor *string `json:"processor,omitempty"`

	// Parent node. Input node should not have parent node. For V1 Alpha1/Beta
	//  only media warehouse node can have multiple parents, other types of nodes
	//  will only have one parent.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.parents
	Parents []Node_InputEdge `json:"parents,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Node.InputEdge
type Node_InputEdge struct {
	// The name of the parent node.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.InputEdge.parent_node
	ParentNode *string `json:"parentNode,omitempty"`

	// The connected output artifact of the parent node.
	//  It can be omitted if target processor only has 1 output artifact.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.InputEdge.parent_output_channel
	ParentOutputChannel *string `json:"parentOutputChannel,omitempty"`

	// The connected input channel of the current node's processor.
	//  It can be omitted if target processor only has 1 input channel.
	// +kcc:proto:field=google.cloud.visionai.v1.Node.InputEdge.connected_input_channel
	ConnectedInputChannel *string `json:"connectedInputChannel,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.NormalizedPolygon
type NormalizedPolygon struct {
	// The bounding polygon normalized vertices. Top left corner of the image
	//  will be [0, 0].
	// +kcc:proto:field=google.cloud.visionai.v1.NormalizedPolygon.normalized_vertices
	NormalizedVertices []NormalizedVertex `json:"normalizedVertices,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.NormalizedPolyline
type NormalizedPolyline struct {
	// A sequence of vertices connected by straight lines.
	// +kcc:proto:field=google.cloud.visionai.v1.NormalizedPolyline.normalized_vertices
	NormalizedVertices []NormalizedVertex `json:"normalizedVertices,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.NormalizedVertex
type NormalizedVertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.visionai.v1.NormalizedVertex.x
	X *float32 `json:"x,omitempty"`

	// Y coordinate.
	// +kcc:proto:field=google.cloud.visionai.v1.NormalizedVertex.y
	Y *float32 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.OccupancyCountConfig
type OccupancyCountConfig struct {
	// Whether to count the appearances of people, output counts have 'people' as
	//  the key.
	// +kcc:proto:field=google.cloud.visionai.v1.OccupancyCountConfig.enable_people_counting
	EnablePeopleCounting *bool `json:"enablePeopleCounting,omitempty"`

	// Whether to count the appearances of vehicles, output counts will have
	//  'vehicle' as the key.
	// +kcc:proto:field=google.cloud.visionai.v1.OccupancyCountConfig.enable_vehicle_counting
	EnableVehicleCounting *bool `json:"enableVehicleCounting,omitempty"`

	// Whether to track each invidual object's loitering time inside the scene or
	//  specific zone.
	// +kcc:proto:field=google.cloud.visionai.v1.OccupancyCountConfig.enable_dwelling_time_tracking
	EnableDwellingTimeTracking *bool `json:"enableDwellingTimeTracking,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.PersonBlurConfig
type PersonBlurConfig struct {
	// Person blur type.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonBlurConfig.person_blur_type
	PersonBlurType *string `json:"personBlurType,omitempty"`

	// Whether only blur faces other than the whole object in the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonBlurConfig.faces_only
	FacesOnly *bool `json:"facesOnly,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.PersonVehicleDetectionConfig
type PersonVehicleDetectionConfig struct {
	// At least one of enable_people_counting and enable_vehicle_counting fields
	//  must be set to true.
	//  Whether to count the appearances of people, output counts have 'people' as
	//  the key.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonVehicleDetectionConfig.enable_people_counting
	EnablePeopleCounting *bool `json:"enablePeopleCounting,omitempty"`

	// Whether to count the appearances of vehicles, output counts will have
	//  'vehicle' as the key.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonVehicleDetectionConfig.enable_vehicle_counting
	EnableVehicleCounting *bool `json:"enableVehicleCounting,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.PersonalProtectiveEquipmentDetectionConfig
type PersonalProtectiveEquipmentDetectionConfig struct {
	// Whether to enable face coverage detection.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonalProtectiveEquipmentDetectionConfig.enable_face_coverage_detection
	EnableFaceCoverageDetection *bool `json:"enableFaceCoverageDetection,omitempty"`

	// Whether to enable head coverage detection.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonalProtectiveEquipmentDetectionConfig.enable_head_coverage_detection
	EnableHeadCoverageDetection *bool `json:"enableHeadCoverageDetection,omitempty"`

	// Whether to enable hands coverage detection.
	// +kcc:proto:field=google.cloud.visionai.v1.PersonalProtectiveEquipmentDetectionConfig.enable_hands_coverage_detection
	EnableHandsCoverageDetection *bool `json:"enableHandsCoverageDetection,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ProcessorConfig
type ProcessorConfig struct {
	// Configs of stream input processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.video_stream_input_config
	VideoStreamInputConfig *VideoStreamInputConfig `json:"videoStreamInputConfig,omitempty"`

	// Config of AI-enabled input devices.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.ai_enabled_devices_input_config
	AiEnabledDevicesInputConfig *AIEnabledDevicesInputConfig `json:"aiEnabledDevicesInputConfig,omitempty"`

	// Configs of media warehouse processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.media_warehouse_config
	MediaWarehouseConfig *MediaWarehouseConfig `json:"mediaWarehouseConfig,omitempty"`

	// Configs of person blur processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.person_blur_config
	PersonBlurConfig *PersonBlurConfig `json:"personBlurConfig,omitempty"`

	// Configs of occupancy count processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.occupancy_count_config
	OccupancyCountConfig *OccupancyCountConfig `json:"occupancyCountConfig,omitempty"`

	// Configs of Person Vehicle Detection processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.person_vehicle_detection_config
	PersonVehicleDetectionConfig *PersonVehicleDetectionConfig `json:"personVehicleDetectionConfig,omitempty"`

	// Configs of Vertex AutoML vision processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.vertex_automl_vision_config
	VertexAutomlVisionConfig *VertexAutoMLVisionConfig `json:"vertexAutomlVisionConfig,omitempty"`

	// Configs of Vertex AutoML video processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.vertex_automl_video_config
	VertexAutomlVideoConfig *VertexAutoMLVideoConfig `json:"vertexAutomlVideoConfig,omitempty"`

	// Configs of Vertex Custom processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.vertex_custom_config
	VertexCustomConfig *VertexCustomConfig `json:"vertexCustomConfig,omitempty"`

	// Configs of General Object Detection processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.general_object_detection_config
	GeneralObjectDetectionConfig *GeneralObjectDetectionConfig `json:"generalObjectDetectionConfig,omitempty"`

	// Configs of BigQuery processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.big_query_config
	BigQueryConfig *BigQueryConfig `json:"bigQueryConfig,omitempty"`

	// Configs of Cloud Storage output processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.gcs_output_config
	GcsOutputConfig *GcsOutputConfig `json:"gcsOutputConfig,omitempty"`

	// Runtime configs of Product Recognizer processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.product_recognizer_config
	ProductRecognizerConfig *ProductRecognizerConfig `json:"productRecognizerConfig,omitempty"`

	// Configs of personal_protective_equipment_detection_config
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.personal_protective_equipment_detection_config
	PersonalProtectiveEquipmentDetectionConfig *PersonalProtectiveEquipmentDetectionConfig `json:"personalProtectiveEquipmentDetectionConfig,omitempty"`

	// Runtime configs of Tag Recognizer processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.tag_recognizer_config
	TagRecognizerConfig *TagRecognizerConfig `json:"tagRecognizerConfig,omitempty"`

	// Runtime configs of UniversalInput processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.universal_input_config
	UniversalInputConfig *UniversalInputConfig `json:"universalInputConfig,omitempty"`

	// Experimental configurations. Structured object containing not-yet-stable
	//  processor parameters.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorConfig.experimental_config
	ExperimentalConfig map[string]string `json:"experimentalConfig,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ProductRecognizerConfig
type ProductRecognizerConfig struct {
	// The resource name of retail endpoint to use.
	// +kcc:proto:field=google.cloud.visionai.v1.ProductRecognizerConfig.retail_endpoint
	RetailEndpoint *string `json:"retailEndpoint,omitempty"`

	// Confidence threshold to filter detection results. If not set, a system
	//  default value will be used.
	// +kcc:proto:field=google.cloud.visionai.v1.ProductRecognizerConfig.recognition_confidence_threshold
	RecognitionConfidenceThreshold *float32 `json:"recognitionConfidenceThreshold,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.StreamAnnotation
type StreamAnnotation struct {
	// Annotation for type ACTIVE_ZONE
	// +kcc:proto:field=google.cloud.visionai.v1.StreamAnnotation.active_zone
	ActiveZone *NormalizedPolygon `json:"activeZone,omitempty"`

	// Annotation for type CROSSING_LINE
	// +kcc:proto:field=google.cloud.visionai.v1.StreamAnnotation.crossing_line
	CrossingLine *NormalizedPolyline `json:"crossingLine,omitempty"`

	// ID of the annotation. It must be unique when used in the certain context.
	//  For example, all the annotations to one input streams of a Vision AI
	//  application.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamAnnotation.id
	ID *string `json:"id,omitempty"`

	// User-friendly name for the annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamAnnotation.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The Vision AI stream resource name.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamAnnotation.source_stream
	SourceStream *string `json:"sourceStream,omitempty"`

	// The actual type of Annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamAnnotation.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.StreamWithAnnotation
type StreamWithAnnotation struct {
	// Vision AI Stream resource name.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamWithAnnotation.stream
	Stream *string `json:"stream,omitempty"`

	// Annotations that will be applied to the whole application.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamWithAnnotation.application_annotations
	ApplicationAnnotations []StreamAnnotation `json:"applicationAnnotations,omitempty"`

	// Annotations that will be applied to the specific node of the application.
	//  If the same type of the annotations is applied to both application and
	//  node, the node annotation will be added in addition to the global
	//  application one.
	//  For example, if there is one active zone annotation for the whole
	//  application and one active zone annotation for the Occupancy Analytic
	//  processor, then the Occupancy Analytic processor will have two active zones
	//  defined.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamWithAnnotation.node_annotations
	NodeAnnotations []StreamWithAnnotation_NodeAnnotation `json:"nodeAnnotations,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.StreamWithAnnotation.NodeAnnotation
type StreamWithAnnotation_NodeAnnotation struct {
	// The node name of the application graph.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamWithAnnotation.NodeAnnotation.node
	Node *string `json:"node,omitempty"`

	// The node specific stream annotations.
	// +kcc:proto:field=google.cloud.visionai.v1.StreamWithAnnotation.NodeAnnotation.annotations
	Annotations []StreamAnnotation `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.TagParsingConfig
type TagParsingConfig struct {
	// Each tag entity class may have an optional EntityParsingConfig which is
	//  used to help parse the entities of the class.
	// +kcc:proto:field=google.cloud.visionai.v1.TagParsingConfig.entity_parsing_configs
	EntityParsingConfigs []TagParsingConfig_EntityParsingConfig `json:"entityParsingConfigs,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.TagParsingConfig.EntityParsingConfig
type TagParsingConfig_EntityParsingConfig struct {
	// Required. The tag entity class name. This should match the class name
	//  produced by the tag entity detection model.
	// +kcc:proto:field=google.cloud.visionai.v1.TagParsingConfig.EntityParsingConfig.entity_class
	EntityClass *string `json:"entityClass,omitempty"`

	// Optional. An regular expression hint.
	// +kcc:proto:field=google.cloud.visionai.v1.TagParsingConfig.EntityParsingConfig.regex
	Regex *string `json:"regex,omitempty"`

	// Optional. Entity matching strategy.
	// +kcc:proto:field=google.cloud.visionai.v1.TagParsingConfig.EntityParsingConfig.entity_matching_strategy
	EntityMatchingStrategy *string `json:"entityMatchingStrategy,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.TagRecognizerConfig
type TagRecognizerConfig struct {
	// Confidence threshold to filter detection results. If not set, a system
	//  default value will be used.
	// +kcc:proto:field=google.cloud.visionai.v1.TagRecognizerConfig.entity_detection_confidence_threshold
	EntityDetectionConfidenceThreshold *float32 `json:"entityDetectionConfidenceThreshold,omitempty"`

	// Configuration to customize how tags are parsed.
	// +kcc:proto:field=google.cloud.visionai.v1.TagRecognizerConfig.tag_parsing_config
	TagParsingConfig *TagParsingConfig `json:"tagParsingConfig,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.UniversalInputConfig
type UniversalInputConfig struct {
}

// +kcc:proto=google.cloud.visionai.v1.VertexAutoMLVideoConfig
type VertexAutoMLVideoConfig struct {
	// Only entities with higher score than the threshold will be returned.
	//  Value 0.0 means returns all the detected entities.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexAutoMLVideoConfig.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Labels specified in this field won't be returned.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexAutoMLVideoConfig.blocked_labels
	BlockedLabels []string `json:"blockedLabels,omitempty"`

	// At most this many predictions will be returned per output frame.
	//  Value 0 means to return all the detected entities.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexAutoMLVideoConfig.max_predictions
	MaxPredictions *int32 `json:"maxPredictions,omitempty"`

	// Only Bounding Box whose size is larger than this limit will be returned.
	//  Object Tracking only.
	//  Value 0.0 means to return all the detected entities.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexAutoMLVideoConfig.bounding_box_size_limit
	BoundingBoxSizeLimit *float32 `json:"boundingBoxSizeLimit,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.VertexAutoMLVisionConfig
type VertexAutoMLVisionConfig struct {
	// Only entities with higher score than the threshold will be returned.
	//  Value 0.0 means to return all the detected entities.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexAutoMLVisionConfig.confidence_threshold
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// At most this many predictions will be returned per output frame.
	//  Value 0 means to return all the detected entities.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexAutoMLVisionConfig.max_predictions
	MaxPredictions *int32 `json:"maxPredictions,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.VertexCustomConfig
type VertexCustomConfig struct {
	// The max prediction frame per second. This attribute sets how fast the
	//  operator sends prediction requests to Vertex AI endpoint. Default value is
	//  0, which means there is no max prediction fps limit. The operator sends
	//  prediction requests at input fps.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexCustomConfig.max_prediction_fps
	MaxPredictionFps *int32 `json:"maxPredictionFps,omitempty"`

	// A description of resources that are dedicated to the DeployedModel, and
	//  that need a higher degree of manual configuration.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexCustomConfig.dedicated_resources
	DedicatedResources *DedicatedResources `json:"dedicatedResources,omitempty"`

	// If not empty, the prediction result will be sent to the specified cloud
	//  function for post processing.
	//  * The cloud function will receive AppPlatformCloudFunctionRequest where
	//  the annotations field will be the json format of proto PredictResponse.
	//  * The cloud function should return AppPlatformCloudFunctionResponse with
	//  PredictResponse stored in the annotations field.
	//  * To drop the prediction output, simply clear the payload field in the
	//  returned AppPlatformCloudFunctionResponse.
	// +kcc:proto:field=google.cloud.visionai.v1.VertexCustomConfig.post_processing_cloud_function
	PostProcessingCloudFunction *string `json:"postProcessingCloudFunction,omitempty"`

	// If true, the prediction request received by custom model will also contain
	//  metadata with the following schema:
	//  'appPlatformMetadata': {
	//        'ingestionTime': DOUBLE; (UNIX timestamp)
	//        'application': STRING;
	//        'instanceId': STRING;
	//        'node': STRING;
	//        'processor': STRING;
	//   }
	// +kcc:proto:field=google.cloud.visionai.v1.VertexCustomConfig.attach_application_metadata
	AttachApplicationMetadata *bool `json:"attachApplicationMetadata,omitempty"`

	// Optional. By setting the configuration_input_topic, processor will
	//  subscribe to given topic, only pub/sub topic is supported now. Example
	//  channel:
	//  //pubsub.googleapis.com/projects/visionai-testing-stable/topics/test-topic
	//  message schema should be:
	//  message Message {
	//  // The ID of the stream that associates with the application instance.
	//  string stream_id = 1;
	//  // The target fps. By default, the custom processor will *not* send any
	//  data to the Vertex Prediction container. Note that once the
	//  dynamic_config_input_topic is set, max_prediction_fps will not work and be
	//  preceded by the fps set inside the topic.
	//  int32 fps = 2;
	//  }
	// +kcc:proto:field=google.cloud.visionai.v1.VertexCustomConfig.dynamic_config_input_topic
	DynamicConfigInputTopic *string `json:"dynamicConfigInputTopic,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.VideoStreamInputConfig
type VideoStreamInputConfig struct {
	// +kcc:proto:field=google.cloud.visionai.v1.VideoStreamInputConfig.streams
	Streams []string `json:"streams,omitempty"`

	// +kcc:proto:field=google.cloud.visionai.v1.VideoStreamInputConfig.streams_with_annotation
	StreamsWithAnnotation []StreamWithAnnotation `json:"streamsWithAnnotation,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Application
type ApplicationObservedState struct {
	// Output only. [Output only] Create timestamp
	// +kcc:proto:field=google.cloud.visionai.v1.Application.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	// +kcc:proto:field=google.cloud.visionai.v1.Application.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Application graph runtime info. Only exists when application
	//  state equals to DEPLOYED.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.runtime_info
	RuntimeInfo *Application_ApplicationRuntimeInfo `json:"runtimeInfo,omitempty"`

	// Output only. State of the application.
	// +kcc:proto:field=google.cloud.visionai.v1.Application.state
	State *string `json:"state,omitempty"`
}
