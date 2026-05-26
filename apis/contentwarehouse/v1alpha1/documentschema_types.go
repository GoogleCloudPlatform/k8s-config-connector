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

var ContentWarehouseSchemaGVK = GroupVersion.WithKind("ContentWarehouseSchema")

// ContentWarehouseSchemaSpec defines the desired state of ContentWarehouseSchema
// +kcc:spec:proto=google.cloud.contentwarehouse.v1.DocumentSchema
type ContentWarehouseSchemaSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// Required. Name of the schema given by the user. Must be unique per project.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.display_name
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Document details.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.property_definitions
	// +kubebuilder:validation:Optional
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`

	// Document Type, true refers the document is a folder, otherwise it is
	//  a typical document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.document_is_folder
	// +kubebuilder:validation:Optional
	DocumentIsFolder *bool `json:"documentIsFolder,omitempty"`

	// Schema description.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentSchema.description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// The ContentWarehouseSchema name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ContentWarehouseSchemaStatus defines the config connector machine state of ContentWarehouseSchema
type ContentWarehouseSchemaStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ContentWarehouseSchema resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ContentWarehouseSchemaObservedState `json:"observedState,omitempty"`
}

// ContentWarehouseSchemaObservedState is the state of the ContentWarehouseSchema resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contentwarehouse.v1.DocumentSchema
type ContentWarehouseSchemaObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontentwarehouseschema;gcpcontentwarehouseschemas
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContentWarehouseSchema is the Schema for the ContentWarehouseSchema API
// +k8s:openapi-gen=true
type ContentWarehouseSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ContentWarehouseSchemaSpec   `json:"spec,omitempty"`
	Status ContentWarehouseSchemaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContentWarehouseSchemaList contains a list of ContentWarehouseSchema
type ContentWarehouseSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentWarehouseSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContentWarehouseSchema{}, &ContentWarehouseSchemaList{})
}

type DateTimeTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.EnumTypeOptions
type EnumTypeOptions struct {
	// Required. List of possible enum values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.EnumTypeOptions.possible_values
	// +kubebuilder:validation:Required
	PossibleValues []string `json:"possibleValues,omitempty"`

	// Make sure the Enum property value provided in the document is in the
	//  possile value list during document creation. The validation check runs by
	//  default.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.EnumTypeOptions.validation_check_disabled
	ValidationCheckDisabled *bool `json:"validationCheckDisabled,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.FloatTypeOptions
type FloatTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.IntegerTypeOptions
type IntegerTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.MapTypeOptions
type MapTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyDefinition
type PropertyDefinition struct {
	// Required. The name of the metadata property.
	//  Must be unique within a document schema and is case insensitive.
	//  Names must be non-blank, start with a letter, and can contain alphanumeric
	//  characters and: /, :, -, _, and .
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.name
	// +kubebuilder:validation:Required
	Name *string `json:"name,omitempty"`
	// The display-name for the property, used for front-end.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Whether the property can have multiple values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_repeatable
	IsRepeatable *bool `json:"isRepeatable,omitempty"`

	// Whether the property can be filtered. If this is a sub-property, all the
	//  parent properties must be marked filterable.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_filterable
	IsFilterable *bool `json:"isFilterable,omitempty"`

	// Indicates that the property should be included in a global search.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_searchable
	IsSearchable *bool `json:"isSearchable,omitempty"`

	// Whether the property is user supplied metadata.
	//  This out-of-the box placeholder setting can be used to tag derived
	//  properties. Its value and interpretation logic should be implemented by API
	//  user.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_metadata
	IsMetadata *bool `json:"isMetadata,omitempty"`

	// Whether the property is mandatory.
	//  Default is 'false', i.e. populating property value can be skipped.
	//  If 'true' then user must populate the value for this property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The retrieval importance of the property during search.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.retrieval_importance
	RetrievalImportance *string `json:"retrievalImportance,omitempty"`

	// Integer property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.integer_type_options
	IntegerTypeOptions *IntegerTypeOptions `json:"integerTypeOptions,omitempty"`

	// Float property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.float_type_options
	FloatTypeOptions *FloatTypeOptions `json:"floatTypeOptions,omitempty"`

	// Text/string property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.text_type_options
	TextTypeOptions *TextTypeOptions `json:"textTypeOptions,omitempty"`

	// Nested structured data property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.property_type_options
	PropertyTypeOptions *PropertyTypeOptions `json:"propertyTypeOptions,omitempty"`

	// Enum/categorical property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.enum_type_options
	EnumTypeOptions *EnumTypeOptions `json:"enumTypeOptions,omitempty"`

	// Date time property.
	//  It is not supported by CMEK compliant deployment.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.date_time_type_options
	DateTimeTypeOptions *DateTimeTypeOptions `json:"dateTimeTypeOptions,omitempty"`

	// Map property.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.map_type_options
	MapTypeOptions *MapTypeOptions `json:"mapTypeOptions,omitempty"`

	// Timestamp property.
	//  It is not supported by CMEK compliant deployment.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.timestamp_type_options
	TimestampTypeOptions *TimestampTypeOptions `json:"timestampTypeOptions,omitempty"`

	// The mapping information between this property to another schema source.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.schema_sources
	SchemaSources []PropertyDefinition_SchemaSource `json:"schemaSources,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyDefinition.SchemaSource
type PropertyDefinition_SchemaSource struct {
	// The schema name in the source.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.SchemaSource.name
	Name *string `json:"name,omitempty"`

	// The Doc AI processor type name.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyDefinition.SchemaSource.processor_type
	ProcessorType *string `json:"processorType,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyTypeOptions
type PropertyTypeOptions struct {
	// Required. List of property definitions.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyTypeOptions.property_definitions
	// +kubebuilder:validation:Required
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TextTypeOptions
type TextTypeOptions struct {
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TimestampTypeOptions
type TimestampTypeOptions struct {
}
