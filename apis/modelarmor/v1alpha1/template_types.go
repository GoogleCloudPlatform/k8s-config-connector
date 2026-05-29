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

var ModelArmorTemplateGVK = GroupVersion.WithKind("ModelArmorTemplate")

// ModelArmorTemplateSpec defines the desired state of ModelArmorTemplate
// +kcc:spec:proto=google.cloud.modelarmor.v1.Template
type ModelArmorTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ModelArmorTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels as key value pairs
	Labels map[string]string `json:"labels,omitempty"`

	// Required. filter configuration for this template
	// +required
	FilterConfig *FilterConfig `json:"filterConfig"`

	// Optional. metadata for this template
	TemplateMetadata *Template_TemplateMetadata `json:"templateMetadata,omitempty"`
}

// ModelArmorTemplateStatus defines the config connector machine state of ModelArmorTemplate
type ModelArmorTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ModelArmorTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ModelArmorTemplateObservedState `json:"observedState,omitempty"`
}

// ModelArmorTemplateObservedState is the state of the ModelArmorTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.modelarmor.v1.Template
type ModelArmorTemplateObservedState struct {
	// Output only. [Output only] Create time stamp
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmodelarmortemplate;gcpmodelarmortemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ModelArmorTemplate is the Schema for the ModelArmorTemplate API
// +k8s:openapi-gen=true
type ModelArmorTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ModelArmorTemplateSpec   `json:"spec,omitempty"`
	Status ModelArmorTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ModelArmorTemplateList contains a list of ModelArmorTemplate
type ModelArmorTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelArmorTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelArmorTemplate{}, &ModelArmorTemplateList{})
}

// +kcc:proto=google.cloud.modelarmor.v1.SdpAdvancedConfig
type SdpAdvancedConfig struct {
	// Optional. Sensitive Data Protection inspect template resource name
	//
	//  If only inspect template is provided (de-identify template not provided),
	//  then Sensitive Data Protection InspectContent action is performed during
	//  Sanitization. All Sensitive Data Protection findings identified during
	//  inspection will be returned as SdpFinding in SdpInsepctionResult.
	//
	//  e.g.
	//  `projects/{project}/locations/{location}/inspectTemplates/{inspect_template}`
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpAdvancedConfig.inspect_template
	InspectTemplateRef *refsv1beta1.DLPInspectTemplateRef `json:"inspectTemplateRef,omitempty"`

	// Optional. Optional Sensitive Data Protection Deidentify template resource
	//  name.
	//
	//  If provided then DeidentifyContent action is performed during Sanitization
	//  using this template and inspect template. The De-identified data will
	//  be returned in SdpDeidentifyResult.
	//  Note that all info-types present in the deidentify template must be present
	//  in inspect template.
	//
	//  e.g.
	//  `projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}`
	// +kcc:proto:field=google.cloud.modelarmor.v1.SdpAdvancedConfig.deidentify_template
	DeidentifyTemplateRef *refsv1beta1.DLPDeidentifyTemplateRef `json:"deidentifyTemplateRef,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.RaiFilterSettings
type RaiFilterSettings struct {
	// Required. List of Responsible AI filters enabled for template.
	// +required
	// +kcc:proto:field=google.cloud.modelarmor.v1.RaiFilterSettings.rai_filters
	RaiFilters []RaiFilterSettings_RaiFilter `json:"raiFilters,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.RaiFilterSettings.RaiFilter
type RaiFilterSettings_RaiFilter struct {
	// Required. Type of responsible AI filter.
	// +required
	// +kcc:proto:field=google.cloud.modelarmor.v1.RaiFilterSettings.RaiFilter.filter_type
	FilterType *string `json:"filterType,omitempty"`

	// Optional. Confidence level for this RAI filter.
	//  During data sanitization, if data is classified under this filter with a
	//  confidence level equal to or greater than the specified level, a positive
	//  match is reported. If the confidence level is unspecified (i.e., 0), the
	//  system will use a reasonable default level based on the `filter_type`.
	// +kcc:proto:field=google.cloud.modelarmor.v1.RaiFilterSettings.RaiFilter.confidence_level
	ConfidenceLevel *string `json:"confidenceLevel,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.Template.TemplateMetadata.MultiLanguageDetection
type Template_TemplateMetadata_MultiLanguageDetection struct {
	// Required. If true, multi language detection will be enabled.
	// +required
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.TemplateMetadata.MultiLanguageDetection.enable_multi_language_detection
	EnableMultiLanguageDetection *bool `json:"enableMultiLanguageDetection,omitempty"`
}
