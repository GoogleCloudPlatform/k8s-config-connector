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
// +kcc:source:service-docs=https://cloud.google.com/application-design-center/docs
// +kcc:source:resource-docs=https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.catalogs.templates.revisions

package v1alpha1

import (
	developerconnectv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DesignCenterCatalogTemplateRevisionGVK = GroupVersion.WithKind("DesignCenterCatalogTemplateRevision")

// +kcc:proto=google.cloud.designcenter.v1.DeveloperConnectSourceConfig
type DeveloperConnectSourceConfig struct {
	// Required. The Developer Connect Git repository link, formatted as
	//  `projects/*/locations/*/connections/*/gitRepositoryLinks/*`.
	// +kcc:proto:field=google.cloud.designcenter.v1.DeveloperConnectSourceConfig.developer_connect_repo_uri
	// +required
	DeveloperConnectRepoRef *developerconnectv1alpha1.DeveloperConnectGitRepositoryLinkRef `json:"developerConnectRepoRef,omitempty"`

	// Required. The reference (for example, a branch, tag, or commit SHA) from
	//  which the content should be read.
	// +kcc:proto:field=google.cloud.designcenter.v1.DeveloperConnectSourceConfig.reference
	// +required
	Reference *GitReference `json:"reference,omitempty"`

	// Required. The sub-directory within the repository from which to read
	//  content. The path must be relative to the repository's root such as
	//  dir1/dir2. To read content from the root dir, provide "/" as the value of
	//  the field.
	// +kcc:proto:field=google.cloud.designcenter.v1.DeveloperConnectSourceConfig.dir
	// +required
	Dir *string `json:"dir,omitempty"`
}

// +kcc:proto=google.cloud.designcenter.v1.GitSource
type GitSource struct {
	// Required. Location of the git repo to build.
	// +kcc:proto:field=google.cloud.designcenter.v1.GitSource.repo
	// +required
	Repo *string `json:"repo,omitempty"`

	// Optional. Directory, relative to the source root, in which to run the
	//  build. This must be a relative path. If a step's `dir` is specified and is
	//  an absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.cloud.designcenter.v1.GitSource.dir
	Dir *string `json:"dir,omitempty"`

	// Required. The reference tag of the git repo.
	// +kcc:proto:field=google.cloud.designcenter.v1.GitSource.ref_tag
	// +required
	RefTag *string `json:"refTag,omitempty"`
}

// DesignCenterCatalogTemplateRevisionSpec defines the desired state of DesignCenterCatalogTemplateRevision
// +kcc:spec:proto=google.cloud.designcenter.v1.CatalogTemplateRevision
// +kcc:required-from-proto
type DesignCenterCatalogTemplateRevisionSpec struct {
	// The DesignCenterCatalogTemplate this DesignCenterCatalogTemplateRevision is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DesignCenterCatalogTemplateRef `json:"parentRef"`

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

// +kcc:proto=google.cloud.designcenter.v1.AlternateDefault
type AlternateDefault struct {
	// Output only. Type of alternate default.
	// +kcc:proto:field=google.cloud.designcenter.v1.AlternateDefault.type
	Type *string `json:"type,omitempty"`

	// Output only. Value of the alternate default.
	// +kcc:proto:field=google.cloud.designcenter.v1.AlternateDefault.value
	Value *Value `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.designcenter.v1.ValueLabel
type ValueLabel struct {
	// Output only. Value of the enum.
	// +kcc:proto:field=google.cloud.designcenter.v1.ValueLabel.value
	Value *string `json:"value,omitempty"`

	// Output only. Label of the enum.
	// +kcc:proto:field=google.cloud.designcenter.v1.ValueLabel.label
	Label *string `json:"label,omitempty"`
}

// +kcc:proto=google.cloud.designcenter.v1.DisplayVariableToggle
type DisplayVariableToggle struct {
	// Output only. The name of the variable used to toggle the display of another
	//  variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariableToggle.variable
	Variable *string `json:"variable,omitempty"`

	// Output only. The value of the variable used to toggle the display of
	//  another variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariableToggle.variable_values
	VariableValues []string `json:"variableValues,omitempty"`

	// Output only. The type of the variable used to toggle the display of another
	//  variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.DisplayVariableToggle.variable_type
	VariableType *string `json:"variableType,omitempty"`
}

// +kcc:proto=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable
type TerraformUiInputDisplayVariable struct {
	// Output only. Display name of the input.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Title of the input.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.title
	Title *string `json:"title,omitempty"`

	// Output only. Alternate defaults for the input.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.alternate_defaults
	AlternateDefaults []AlternateDefault `json:"alternateDefaults,omitempty"`

	// Output only. Regex based validation rules for the variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.regex_validation
	RegexValidation *string `json:"regexValidation,omitempty"`

	// Output only. Minimum value for numeric types.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.min
	Min *float32 `json:"min,omitempty"`

	// Output only. Maximum value for numeric types.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.max
	Max *float32 `json:"max,omitempty"`

	// Output only. Text describing the validation rules for the property.
	//  Typically shown after an invalid input. Optional. UTF-8 text. No markup. At
	//  most 128 characters.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.validation
	Validation *string `json:"validation,omitempty"`

	// Output only. Labels for enum values.
	//  Values must be UTF-8 text with no markup, and at most 64 characters.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.enum_value_labels
	EnumValueLabels []ValueLabel `json:"enumValueLabels,omitempty"`

	// Output only. Indicates the "advanced" level of the input property. Level 0
	//  (default) will always be shown. Level 1 corresponds to one expansion (user
	//  clicks "show advanced options" or "more options"). Higher levels correspond
	//  to further expansions, or they may be collapsed to level 1 by the UI
	//  implementation.
	//  Optional.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.level
	Level *int32 `json:"level,omitempty"`

	// Output only. Property subtext, displayed below the title.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.subtext
	Subtext *string `json:"subtext,omitempty"`

	// Output only. Variables used to toggle the display of another variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.toggle_using_variables
	ToggleUsingVariables []DisplayVariableToggle `json:"toggleUsingVariables,omitempty"`

	// Output only. A map that defines all fields of the input variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiInputDisplayVariable.properties
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	Properties map[string]TerraformUiInputDisplayVariable `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.designcenter.v1.TerraformUiOutputDisplay
type TerraformUiOutputDisplay struct {
	// Output only. Visibility of the output.
	// +kcc:proto:field=google.cloud.designcenter.v1.TerraformUiOutputDisplay.visibility
	Visibility *string `json:"visibility,omitempty"`
}

// +kcc:proto=google.cloud.designcenter.v1.HelmChartInput
type HelmChartInput struct {
	// Output only. Input variable name present in Helm Chart values.yaml
	// +kcc:proto:field=google.cloud.designcenter.v1.HelmChartInput.helm_input_variable
	HelmInputVariable *string `json:"helmInputVariable,omitempty"`

	// Output only. Input variable description.
	// +kcc:proto:field=google.cloud.designcenter.v1.HelmChartInput.description
	Description *string `json:"description,omitempty"`

	// Output only. Input data type.
	// +kcc:proto:field=google.cloud.designcenter.v1.HelmChartInput.type
	Type *string `json:"type,omitempty"`

	// Output only. The default value of the input variable.
	// +kcc:proto:field=google.cloud.designcenter.v1.HelmChartInput.default_value
	DefaultValue *Value `json:"defaultValue,omitempty"`

	// Output only. Indicates if input is required.
	// +kcc:proto:field=google.cloud.designcenter.v1.HelmChartInput.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// Output only. Contains details of nested inputs.
	// +kcc:proto:field=google.cloud.designcenter.v1.HelmChartInput.nested_inputs
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	NestedInputs map[string]HelmChartInput `json:"nestedInputs,omitempty"`
}
