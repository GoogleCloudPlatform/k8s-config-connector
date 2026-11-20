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

// API sources for DesignCenterCatalogTemplateRevision, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/catalog_management.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DesignCenterCatalogTemplateRevisionGVK = GroupVersion.WithKind("DesignCenterCatalogTemplateRevision")

// DesignCenterCatalogTemplateRevisionSpec defines the desired state of DesignCenterCatalogTemplateRevision
// +kcc:spec:proto=google.cloud.designcenter.v1.CatalogTemplateRevision
// +kcc:required-from-proto
type DesignCenterCatalogTemplateRevisionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/spaces/{space}/catalogs/{catalog}/templates/{catalog_template} this resource belongs to.
	// +kcc:guess
	// TemplateRef *DesignCenterCatalogTemplateRef `json:"templateRef,omitempty"`

	// The DesignCenterCatalogTemplateRevision name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The git source.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.git_source
	GitSource *GitSource `json:"gitSource,omitempty"`

	// Optional. The application template revision source.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.application_template_revision_source
	ApplicationTemplateRevisionSource *string `json:"applicationTemplateRevisionSource,omitempty"`

	// Optional. Configuration for fetching content from source code repository
	//  such as GitHub or Bitbucket through Developer Connect.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.developer_connect_source_config
	DeveloperConnectSourceConfig *DeveloperConnectSourceConfig `json:"developerConnectSourceConfig,omitempty"`

	// Optional. The Open Container Initiative (OCI) repo source that contains
	//  helm charts.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.oci_repo
	OciRepo *OciRepo `json:"ociRepo,omitempty"`

	// Optional. The catalog template revision description.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.description
	Description *string `json:"description,omitempty"`

	// Optional. The Application Design Center application template type.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.type
	Type *string `json:"type,omitempty"`

	// Optional. Metadata input.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.metadata_input
	MetadataInput *MetadataInput `json:"metadataInput,omitempty"`

	// Optional. The annotations of the template revision.
	//  Key is the annotation name. Value is the annotation value.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. The values provided for the input variables.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.application_parameters
	ApplicationParameters []Parameter `json:"applicationParameters,omitempty"`
}

// DesignCenterCatalogTemplateRevisionStatus defines the config connector machine state of DesignCenterCatalogTemplateRevision
type DesignCenterCatalogTemplateRevisionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterCatalogTemplateRevision resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterCatalogTemplateRevisionObservedState `json:"observedState,omitempty"`
}

// DesignCenterCatalogTemplateRevisionObservedState is the state of the DesignCenterCatalogTemplateRevision resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.CatalogTemplateRevision
type DesignCenterCatalogTemplateRevisionObservedState struct {
	// Optional. The git source.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.git_source
	GitSource *GitSourceObservedState `json:"gitSource,omitempty"`

	// Optional. Configuration for fetching content from source code repository
	//  such as GitHub or Bitbucket through Developer Connect.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.developer_connect_source_config
	DeveloperConnectSourceConfig *DeveloperConnectSourceConfigObservedState `json:"developerConnectSourceConfig,omitempty"`

	// Output only. The catalog template creation timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The catalog template update timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Template metadata related to Terraform input and output.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.template_metadata
	TemplateMetadata *TfBlueprintMetadataObservedState `json:"templateMetadata,omitempty"`

	// Output only. The template state (validating/ready/invalid).
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.state
	State *string `json:"state,omitempty"`

	// Output only. The application template revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.application_template_revision
	ApplicationTemplateRevision *SerializedApplicationTemplateObservedState `json:"applicationTemplateRevision,omitempty"`

	// Output only. The helm chart metadata.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.helm_chart_metadata
	HelmChartMetadata *HelmChartMetadataObservedState `json:"helmChartMetadata,omitempty"`

	// Output only. UUID of the template revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Optional. Metadata input.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.metadata_input
	MetadataInput *MetadataInputObservedState `json:"metadataInput,omitempty"`

	// Output only. The category of the template.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.template_category
	TemplateCategory *string `json:"templateCategory,omitempty"`

	// Output only. The Product Main logical product type information.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.logical_products
	LogicalProducts []LogicalProductObservedState `json:"logicalProducts,omitempty"`

	// Output only. Metadata that was automatically inferred from the template
	//  content. This field can be updated by the system as it gets new
	//  information.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.inferred_metadata
	InferredMetadata *InferredMetadataObservedState `json:"inferredMetadata,omitempty"`

	// Output only. The resource types present in the template revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.resource_types
	ResourceTypes []ResourceTypeObservedState `json:"resourceTypes,omitempty"`

	// Output only. Indicates the IaC source format of this catalog template
	//  revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.iac_source
	IacSource *string `json:"iacSource,omitempty"`

	// Optional. The values provided for the input variables.
	// +kcc:proto:field=google.cloud.designcenter.v1.CatalogTemplateRevision.application_parameters
	ApplicationParameters []ParameterObservedState `json:"applicationParameters,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncentercatalogtemplaterevision;gcpdesigncentercatalogtemplaterevisions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterCatalogTemplateRevision is the Schema for the DesignCenterCatalogTemplateRevision API
// +k8s:openapi-gen=true
type DesignCenterCatalogTemplateRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterCatalogTemplateRevisionSpec   `json:"spec,omitempty"`
	Status DesignCenterCatalogTemplateRevisionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterCatalogTemplateRevisionList contains a list of DesignCenterCatalogTemplateRevision
type DesignCenterCatalogTemplateRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterCatalogTemplateRevision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterCatalogTemplateRevision{}, &DesignCenterCatalogTemplateRevisionList{})
}

// +kcc:proto=google.cloud.designcenter.v1.DisplayVariable
type DisplayVariable struct {
	// Required. The variable name from the corresponding standard Template file.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.name
	// +required
	Name *string `json:"name,omitempty"`

	// Required. Visible title for the variable on the UI.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.title
	// +required
	Title *string `json:"title,omitempty"`

	// Optional. Alternate defaults for the input.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.alternate_defaults
	AlternateDefaults []AlternateDefaultInput `json:"alternateDefaults,omitempty"`

	// Optional. Regex based validation rules for the variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.regex_validation
	RegexValidation *string `json:"regexValidation,omitempty"`

	// Optional. Minimum value for numeric types.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.min
	Min *float32 `json:"min,omitempty"`

	// Optional. Maximum value for numeric types.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.max
	Max *float32 `json:"max,omitempty"`

	// Optional. Text describing the validation rules for the property. Typically
	//  shown after an invalid input. Optional. UTF-8 text. No markup. At most 128
	//  characters.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.validation
	Validation *string `json:"validation,omitempty"`

	// Optional. Labels for enum values.
	//  Values must be UTF-8 text with no markup, and at most 64 characters.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.enum_value_labels
	EnumValueLabels []ValueLabelInput `json:"enumValueLabels,omitempty"`

	// Optional. Indicates the "advanced" level of the input property. Level 0
	//  (default) will always be shown. Level 1 corresponds to one expansion (user
	//  clicks "show advanced options" or "more options"). Higher levels correspond
	//  to further expansions, or they may be collapsed to level 1 by the UI
	//  implementation.
	//  Optional.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.level
	Level *int32 `json:"level,omitempty"`

	// Optional. Property subtext, displayed below the title.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.subtext
	Subtext *string `json:"subtext,omitempty"`

	// Optional. properties is a map defining all the fields of the input
	//  variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariable.properties
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	Properties map[string]DisplayVariable `json:"properties,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	Values []Value `json:"values,omitempty"`
}
