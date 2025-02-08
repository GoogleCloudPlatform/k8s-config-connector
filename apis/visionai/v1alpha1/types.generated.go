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


// +kcc:proto=google.cloud.visionai.v1.ApplicationNodeAnnotation
type ApplicationNodeAnnotation struct {
	// The node name of the application graph.
	// +kcc:proto:field=google.cloud.visionai.v1.ApplicationNodeAnnotation.node
	Node *string `json:"node,omitempty"`

	// The node specific stream annotations.
	// +kcc:proto:field=google.cloud.visionai.v1.ApplicationNodeAnnotation.annotations
	Annotations []StreamAnnotation `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Instance
type Instance struct {

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A user friendly display name for the solution.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description for this instance.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.description
	Description *string `json:"description,omitempty"`

	// The instance type for the current instance.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// The input resources for the current application instance.
	//  For example:
	//  input_resources:
	//  visionai.googleapis.com/v1/projects/123/locations/us-central1/clusters/456/streams/stream-a
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.input_resources
	InputResources []Instance_InputResource `json:"inputResources,omitempty"`

	// All the output resources associated to one application instance.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.output_resources
	OutputResources []Instance_OutputResource `json:"outputResources,omitempty"`

	// State of the instance.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Instance.InputResource
type Instance_InputResource struct {
	// The direct input resource name.
	//  If the instance type is STREAMING_PREDICTION, the input resource is in
	//  format of
	//  "projects/123/locations/us-central1/clusters/456/streams/stream-a".
	//  If the instance type is BATCH_PREDICTION from Cloud Storage input
	//  container, the input resource is in format of "gs://bucket-a".
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.InputResource.input_resource
	InputResource *string `json:"inputResource,omitempty"`

	// If the input resource is VisionAI Stream, the associated annotations
	//  can be specified using annotated_stream instead.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.InputResource.annotated_stream
	AnnotatedStream *StreamWithAnnotation `json:"annotatedStream,omitempty"`

	// Data type for the current input resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.InputResource.data_type
	DataType *string `json:"dataType,omitempty"`

	// The name of graph node who receives the input resource.
	//  For example:
	//  input_resource:
	//  visionai.googleapis.com/v1/projects/123/locations/us-central1/clusters/456/streams/input-stream-a
	//  consumer_node: stream-input
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.InputResource.consumer_node
	ConsumerNode *string `json:"consumerNode,omitempty"`

	// The specific input resource binding which will consume the current Input
	//  Resource, can be ignored is there is only 1 input binding.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.InputResource.input_resource_binding
	InputResourceBinding *string `json:"inputResourceBinding,omitempty"`

	// Contains resource annotations.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.InputResource.annotations
	Annotations *ResourceAnnotations `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Instance.OutputResource
type Instance_OutputResource struct {
	// The output resource name for the current application instance.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.OutputResource.output_resource
	OutputResource *string `json:"outputResource,omitempty"`

	// The name of graph node who produces the output resource name.
	//  For example:
	//  output_resource:
	//  /projects/123/locations/us-central1/clusters/456/streams/output-application-789-stream-a-occupancy-counting
	//  producer_node: occupancy-counting
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.OutputResource.producer_node
	ProducerNode *string `json:"producerNode,omitempty"`

	// The specific output resource binding which produces the current
	//  OutputResource.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.OutputResource.output_resource_binding
	OutputResourceBinding *string `json:"outputResourceBinding,omitempty"`
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

// +kcc:proto=google.cloud.visionai.v1.ResourceAnnotations
type ResourceAnnotations struct {
	// Annotations that will be applied to the whole application.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceAnnotations.application_annotations
	ApplicationAnnotations []StreamAnnotation `json:"applicationAnnotations,omitempty"`

	// Annotations that will be applied to the specific node of the application.
	//  If the same type of the annotations is applied to both application and
	//  node, the node annotation will be added in addition to the global
	//  application one.
	//  For example, if there is one active zone annotation for the whole
	//  application and one active zone annotation for the Occupancy Analytic
	//  processor, then the Occupancy Analytic processor will have two active zones
	//  defined.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceAnnotations.node_annotations
	NodeAnnotations []ApplicationNodeAnnotation `json:"nodeAnnotations,omitempty"`
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

// +kcc:proto=google.cloud.visionai.v1.Instance
type InstanceObservedState struct {
	// Output only. name of resource
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.name
	Name *string `json:"name,omitempty"`

	// Output only. [Output only] Create timestamp
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// All the output resources associated to one application instance.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.output_resources
	OutputResources []Instance_OutputResourceObservedState `json:"outputResources,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Instance.OutputResource
type Instance_OutputResourceObservedState struct {
	// Output only. Whether the output resource is temporary which means the
	//  resource is generated during the deployment of the application. Temporary
	//  resource will be deleted during the undeployment of the application.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.OutputResource.is_temporary
	IsTemporary *bool `json:"isTemporary,omitempty"`

	// Output only. Whether the output resource is created automatically by the
	//  Vision AI App Platform.
	// +kcc:proto:field=google.cloud.visionai.v1.Instance.OutputResource.autogen
	Autogen *bool `json:"autogen,omitempty"`
}
