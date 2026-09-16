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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContentWarehouseDocumentSchemaGVK = GroupVersion.WithKind("ContentWarehouseDocumentSchema")

// ContentWarehouseDocumentSchemaSpec defines the desired state of ContentWarehouseDocumentSchema
// +kcc:spec:proto=google.cloud.contentwarehouse.v1.DocumentSchema
type ContentWarehouseDocumentSchemaSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// Required. Name of the schema given by the user. Must be unique per project.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.display_name
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Document details.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.property_definitions
	// +kubebuilder:validation:Optional
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`

	// Document Type, true refers the document is a folder, otherwise it is a typical document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.document_is_folder
	// +kubebuilder:validation:Optional
	DocumentIsFolder *bool `json:"documentIsFolder,omitempty"`

	// Schema description.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// The ContentWarehouseDocumentSchema name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
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
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontentwarehousedocumentschema;gcpcontentwarehousedocumentschemas
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
