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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexAspectTypeGVK = GroupVersion.WithKind("DataplexAspectType")

// DataplexAspectTypeSpec defines the desired state of DataplexAspectType
// +kcc:spec:proto=google.cloud.dataplex.v1.AspectType
type DataplexAspectTypeSpec struct {
	ParentRef *parent.ProjectAndLocationRef `json:",inline"`

	// The DataplexAspectType name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Description of the AspectType.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the AspectType.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Defines the Authorization for this type.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.authorization
	Authorization *AspectType_Authorization `json:"authorization,omitempty"`

	// Required. MetadataTemplate of the aspect.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.metadata_template
	MetadataTemplate *AspectType_MetadataTemplate `json:"metadataTemplate,omitempty"`
}

// DataplexAspectTypeStatus defines the config connector machine state of DataplexAspectType
type DataplexAspectTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexAspectType resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexAspectTypeObservedState `json:"observedState,omitempty"`
}

// DataplexAspectTypeObservedState is the state of the DataplexAspectType resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.AspectType
type DataplexAspectTypeObservedState struct {
	// Output only. System generated globally unique ID for the AspectType.
	//  If you delete and recreate the AspectType with the same name, then this ID
	//  will be different.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the AspectType was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the AspectType was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Denotes the transfer status of the Aspect Type. It is
	//  unspecified for Aspect Types created from Dataplex API.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.transfer_status
	TransferStatus *string `json:"transferStatus,omitempty"`

	// Optional. This checksum is computed by the service, and might be sent on
	//  update and delete requests to ensure the client has an up-to-date value
	//  before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexaspecttype;gcpdataplexaspecttypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexAspectType is the Schema for the DataplexAspectType API
// +k8s:openapi-gen=true
type DataplexAspectType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexAspectTypeSpec   `json:"spec,omitempty"`
	Status DataplexAspectTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexAspectTypeList contains a list of DataplexAspectType
type DataplexAspectTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexAspectType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexAspectType{}, &DataplexAspectTypeList{})
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType.MetadataTemplate
type AspectType_MetadataTemplate struct {
	// Optional. Index is used to encode Template messages. The value of index
	//  can range between 1 and 2,147,483,647. Index must be unique within all
	//  fields in a Template. (Nested Templates can reuse indexes). Once a
	//  Template is defined, the index cannot be changed, because it identifies
	//  the field in the actual storage format. Index is a mandatory field, but
	//  it is optional for top level fields, and map/array "values" definitions.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.index
	Index *int32 `json:"index,omitempty"`

	// Required. The name of the field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.name
	Name *string `json:"name,omitempty"`

	// Required. The datatype of this field. The following values are supported:
	//
	//  Primitive types:
	//
	//  * string
	//  * integer
	//  * boolean
	//  * double
	//  * datetime. Must be of the format RFC3339 UTC "Zulu" (Examples:
	//  "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z").
	//
	//  Complex types:
	//
	//  * enum
	//  * array
	//  * map
	//  * record
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.type
	Type *string `json:"type,omitempty"`

	// Optional. Field definition. You must specify it if the type is record. It
	//  defines the nested fields.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.record_fields
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	RecordFields []AspectType_MetadataTemplate `json:"recordFields,omitempty"`

	// Optional. The list of values for an enum type. You must define it if the
	//  type is enum.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.enum_values
	EnumValues []AspectType_MetadataTemplate_EnumValue `json:"enumValues,omitempty"`

	// Optional. If the type is map, set map_items. map_items can refer to a
	//  primitive field or a complex (record only) field. To specify a primitive
	//  field, you only need to set name and type in the nested
	//  MetadataTemplate. The recommended value for the name field is item, as
	//  this isn't used in the actual payload.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.map_items
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	MapItems *AspectType_MetadataTemplate `json:"mapItems,omitempty"`

	// Optional. If the type is array, set array_items. array_items can refer
	//  to a primitive field or a complex (record only) field. To specify a
	//  primitive field, you only need to set name and type in the nested
	//  MetadataTemplate. The recommended value for the name field is item, as
	//  this isn't used in the actual payload.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.array_items
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	ArrayItems *AspectType_MetadataTemplate `json:"arrayItems,omitempty"`

	// Optional. You can use type id if this definition of the field needs to be
	//  reused later. The type id must be unique across the entire template. You
	//  can only specify it if the field type is record.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.type_id
	TypeID *string `json:"typeID,omitempty"`

	// Optional. A reference to another field definition (not an inline
	//  definition). The value must be equal to the value of an id field defined
	//  elsewhere in the MetadataTemplate. Only fields with record type can
	//  refer to other fields.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.type_ref
	TypeRef *string `json:"typeRef,omitempty"`

	// Optional. Specifies the constraints on this field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.constraints
	Constraints *AspectType_MetadataTemplate_Constraints `json:"constraints,omitempty"`

	// Optional. Specifies annotations on this field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.annotations
	Annotations *AspectType_MetadataTemplate_Annotations `json:"annotations,omitempty"`
}
