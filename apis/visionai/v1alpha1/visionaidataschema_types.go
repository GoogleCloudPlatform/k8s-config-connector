// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for VisionAIDataSchema, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/warehouse.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.corpora.dataSchemas

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIDataSchemaGVK = GroupVersion.WithKind("VisionAIDataSchema")

// VisionAIDataSchemaSpec defines the desired state of VisionAIDataSchema
// +kcc:spec:proto=google.cloud.visionai.v1.DataSchema
// +kcc:required-from-proto
type VisionAIDataSchemaSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project_number}/locations/{location}/corpora/{corpus} this resource belongs to.
	// +kcc:guess
	// CorporaRef *VisionAICorpusRef `json:"corporaRef,omitempty"`

	// The VisionAIDataSchema name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The key of this data schema. This key should be matching the key
	//  of user specified annotation and unique inside corpus. This value can be up
	//  to 63 characters, and valid characters are /[a-z][0-9]-/. The first
	//  character must be a letter, the last could be a letter or a number.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchema.key
	// +required
	Key *string `json:"key,omitempty"`

	// The schema details mapping to the key.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchema.schema_details
	SchemaDetails *DataSchemaDetails `json:"schemaDetails,omitempty"`
}

// VisionAIDataSchemaStatus defines the config connector machine state of VisionAIDataSchema
type VisionAIDataSchemaStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIDataSchema resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *VisionAIDataSchemaObservedState `json:"observedState,omitempty"`
}

// VisionAIDataSchemaObservedState is the state of the VisionAIDataSchema resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.DataSchema
// type VisionAIDataSchemaObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaidataschema;gcpvisionaidataschemas
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIDataSchema is the Schema for the VisionAIDataSchema API
// +k8s:openapi-gen=true
type VisionAIDataSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIDataSchemaSpec   `json:"spec,omitempty"`
	Status VisionAIDataSchemaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIDataSchemaList contains a list of VisionAIDataSchema
type VisionAIDataSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIDataSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAIDataSchema{}, &VisionAIDataSchemaList{})
}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails
type DataSchemaDetails struct {
	// Type of the annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.type
	Type *string `json:"type,omitempty"`

	// Config for protobuf any type.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.proto_any_config
	ProtoAnyConfig *DataSchemaDetails_ProtoAnyConfig `json:"protoAnyConfig,omitempty"`

	// Config for List data type.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.list_config
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	ListConfig *DataSchemaDetails_ListConfig `json:"listConfig,omitempty"`

	// Config for CustomizedStruct data type.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.customized_struct_config
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	CustomizedStructConfig *DataSchemaDetails_CustomizedStructConfig `json:"customizedStructConfig,omitempty"`

	// The granularity associated with this DataSchema.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.granularity
	Granularity *string `json:"granularity,omitempty"`

	// The search strategy to be applied on the `key` above.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.search_strategy
	SearchStrategy *DataSchemaDetails_SearchStrategy `json:"searchStrategy,omitempty"`
}
