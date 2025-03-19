// Copyright 2025 Google LLC
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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataCatalogTagTemplateGVK = GroupVersion.WithKind("DataCatalogTagTemplate")

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +required
	Location string `json:"location"`
}

// +kcc:proto=google.cloud.datacatalog.v1.FieldType
type FieldType struct {
	// Primitive types, such as string, boolean, etc.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FieldType.primitive_type
	// +kubebuilder:validation:Enum=STRING;BOOL;DOUBLE;TIMESTAMP;RICHTEXT
	PrimitiveType *string `json:"primitiveType,omitempty"`

	// An enum type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FieldType.enum_type
	EnumType *FieldType_EnumType `json:"enumType,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.TagTemplateField
type TagTemplateField struct {
	// Identifier. The resource name of the tag template field in URL format.
	//  Example:
	//
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/tagTemplates/{TAG_TEMPLATE}/fields/{FIELD}`
	//
	//  Note: The tag template field itself might not be stored in the location
	//  specified in its name.
	//
	//  The name must contain only letters (a-z, A-Z), numbers (0-9),
	//  or underscores (_), and must start with a letter or underscore.
	//  The maximum length is 64 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.name
	// +required
	Name *string `json:"name,omitempty"`

	// The display name for this field. Defaults to an empty string.
	//
	//  The name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), and can't start or end with spaces.
	//  The maximum length is 200 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of value this tag field can contain.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.type
	// +required
	Type *FieldType `json:"type,omitempty"`

	// If true, this field is required. Defaults to false.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The description for this field. Defaults to an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.description
	Description *string `json:"description,omitempty"`

	// The order of this field with respect to other fields in this tag
	//  template.
	//
	//  For example, a higher value can indicate a more important field.
	//  The value can be negative. Multiple fields can have the same order and
	//  field orders within a tag don't have to be sequential.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.order
	Order *int32 `json:"order,omitempty"`
}

// DataCatalogTagTemplateSpec defines the desired state of DataCatalogTagTemplate
// +kcc:proto=google.cloud.datacatalog.v1.TagTemplate
type DataCatalogTagTemplateSpec struct {
	Parent `json:",inline"`

	// The DataCatalogTagTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Display name for this template. Defaults to an empty string.
	//
	//  The name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), and can't start or end with spaces.
	//  The maximum length is 200 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether tags created with this template are public. Public tags
	//  do not require tag template access to appear in
	//  [ListTags][google.cloud.datacatalog.v1.DataCatalog.ListTags] API response.
	//
	//  Additionally, you can search for a public tag by value with a
	//  simple search query in addition to using a ``tag:`` predicate.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.is_publicly_readable
	IsPubliclyReadable *bool `json:"isPubliclyReadable,omitempty"`

	// Fields used to create a Tag
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.fields
	Fields map[string]TagTemplateField `json:"fields,omitempty"`

	// Optional. Transfer status of the TagTemplate
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.dataplex_transfer_status
	// +kubebuilder:validation:Enum=TRANSFERRED
	DataplexTransferStatus *string `json:"dataplexTransferStatus,omitempty"`
}

// DataCatalogTagTemplateStatus defines the config connector machine state of DataCatalogTagTemplate
type DataCatalogTagTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogTagTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogtagtemplate;gcpdatacatalogtagtemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogTagTemplate is the Schema for the DataCatalogTagTemplate API
// +k8s:openapi-gen=true
type DataCatalogTagTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogTagTemplateSpec   `json:"spec,omitempty"`
	Status DataCatalogTagTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogTagTemplateList contains a list of DataCatalogTagTemplate
type DataCatalogTagTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogTagTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogTagTemplate{}, &DataCatalogTagTemplateList{})
}
