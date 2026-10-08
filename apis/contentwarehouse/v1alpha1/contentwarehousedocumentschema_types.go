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

// API sources for ContentWarehouseDocumentSchema, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/contentwarehouse/v1/document_schema.proto
// +kcc:source:service-docs=https://cloud.google.com/document-warehouse
// +kcc:source:resource-docs=https://docs.cloud.google.com/document-warehouse/docs/reference/rest/v1/projects.locations.documentSchemas

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContentWarehouseDocumentSchemaGVK = GroupVersion.WithKind("ContentWarehouseDocumentSchema")

// ContentWarehouseDocumentSchemaSpec defines the desired state of ContentWarehouseDocumentSchema
// +kcc:spec:proto=google.cloud.contentwarehouse.v1.DocumentSchema
// +kcc:required-from-proto
type ContentWarehouseDocumentSchemaSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ContentWarehouseDocumentSchema name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Name of the schema given by the user. Must be unique per project.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Document details.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.property_definitions
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`

	// Document Type, true refers the document is a folder, otherwise it is
	//  a typical document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.document_is_folder
	DocumentIsFolder *bool `json:"documentIsFolder,omitempty"`

	// Schema description.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.description
	Description *string `json:"description,omitempty"`
}

// ContentWarehouseDocumentSchemaStatus defines the config connector machine state of ContentWarehouseDocumentSchema
type ContentWarehouseDocumentSchemaStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ContentWarehouseDocumentSchema resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ContentWarehouseDocumentSchemaObservedState `json:"observedState,omitempty"`
}

// ContentWarehouseDocumentSchemaObservedState is the state of the ContentWarehouseDocumentSchema resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contentwarehouse.v1.DocumentSchema
type ContentWarehouseDocumentSchemaObservedState struct {
	// Output only. The time when the document schema is last updated.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the document schema is created.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontentwarehousedocumentschema;gcpcontentwarehousedocumentschemas
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContentWarehouseDocumentSchema is the Schema for the ContentWarehouseDocumentSchema API
// +k8s:openapi-gen=true
type ContentWarehouseDocumentSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ContentWarehouseDocumentSchemaSpec   `json:"spec,omitempty"`
	Status ContentWarehouseDocumentSchemaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContentWarehouseDocumentSchemaList contains a list of ContentWarehouseDocumentSchema
type ContentWarehouseDocumentSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentWarehouseDocumentSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContentWarehouseDocumentSchema{}, &ContentWarehouseDocumentSchemaList{})
}

// PropertyTypeOptions is written by hand to give the list items a type.
// PropertyDefinition contains PropertyTypeOptions, so the type is recursive.
// Without these markers controller-gen leaves the items schema empty, and the
// API server rejects the CRD.
// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyTypeOptions
type PropertyTypeOptions struct {
	// Required. List of property definitions.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyTypeOptions.property_definitions
	// +required
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`
}
