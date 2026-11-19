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

// API sources for DesignCenterApplicationTemplate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/authoring.proto
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

var DesignCenterApplicationTemplateGVK = GroupVersion.WithKind("DesignCenterApplicationTemplate")

// DesignCenterApplicationTemplateSpec defines the desired state of DesignCenterApplicationTemplate
// +kcc:spec:proto=google.cloud.designcenter.v1.ApplicationTemplate
// +kcc:required-from-proto
type DesignCenterApplicationTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/spaces/{space} this resource belongs to.
	// +kcc:guess
	// SpaceRef *DesignCenterSpaceRef `json:"spaceRef,omitempty"`

	// The DesignCenterApplicationTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Application template display name.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Application template description.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Parameters to apply to all components in an application. You can
	//  specify projectID and region.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.application_parameters
	ApplicationParameters []Parameter `json:"applicationParameters,omitempty"`

	// Optional. Immutable. The composition type of the applicationTemplate:
	//  STANDARD OR COMPOSITE. This is a create time only param.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.composition_type
	CompositionType *string `json:"compositionType,omitempty"`

	// Optional. Root level input variables of the application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.root_input_variables
	RootInputVariables []ComponentVariable `json:"rootInputVariables,omitempty"`

	// Optional. Root level output variables of the application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.root_output_variables
	RootOutputVariables []ComponentVariable `json:"rootOutputVariables,omitempty"`

	// Optional. SaaS runtime context for the application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.saas_runtime_context
	SaasRuntimeContext *SaaSRuntimeContext `json:"saasRuntimeContext,omitempty"`

	// Optional. If true, exposes all inputs as root module variables.
	//  This is only supported for STANDARD composition type templates.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.inputs_as_root_module_variables
	InputsAsRootModuleVariables *bool `json:"inputsAsRootModuleVariables,omitempty"`
}

// DesignCenterApplicationTemplateStatus defines the config connector machine state of DesignCenterApplicationTemplate
type DesignCenterApplicationTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterApplicationTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterApplicationTemplateObservedState `json:"observedState,omitempty"`
}

// DesignCenterApplicationTemplateObservedState is the state of the DesignCenterApplicationTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.ApplicationTemplate
type DesignCenterApplicationTemplateObservedState struct {
	// Output only. Application template creation timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Application template update timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Parameters to apply to all components in an application. You can
	//  specify projectID and region.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.application_parameters
	ApplicationParameters []ParameterObservedState `json:"applicationParameters,omitempty"`

	// Output only. The IaC format of the application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.iac_format
	IacFormat *string `json:"iacFormat,omitempty"`

	// Output only. The serialized application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.serialized_application_template
	SerializedApplicationTemplate *SerializedApplicationTemplateObservedState `json:"serializedApplicationTemplate,omitempty"`

	// Output only. The latest application template revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.latest_revision
	LatestRevision *string `json:"latestRevision,omitempty"`

	// Output only. Details of the location where the IaC for this
	//  ApplicationTemplate was last successfully exported.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.artifact_location
	ArtifactLocation *ArtifactLocationObservedState `json:"artifactLocation,omitempty"`

	// Output only. The URI of the most recent System Revision generated for this
	//  ApplicationTemplate. This ID corresponds to the `name` field in the
	//  ApplicationTemplateSystemRevision resource.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.latest_system_revision_uri
	LatestSystemRevisionURI *string `json:"latestSystemRevisionURI,omitempty"`

	// Output only. Indicates the IaC source format of this application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.iac_source
	IacSource *string `json:"iacSource,omitempty"`

	// Output only. The input output variables parsed from the imported terraform
	//  code.
	// +kcc:proto:field=google.cloud.designcenter.v1.ApplicationTemplate.template_metadata
	TemplateMetadata *TfBlueprintMetadataObservedState `json:"templateMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncenterapplicationtemplate;gcpdesigncenterapplicationtemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterApplicationTemplate is the Schema for the DesignCenterApplicationTemplate API
// +k8s:openapi-gen=true
type DesignCenterApplicationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterApplicationTemplateSpec   `json:"spec,omitempty"`
	Status DesignCenterApplicationTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterApplicationTemplateList contains a list of DesignCenterApplicationTemplate
type DesignCenterApplicationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterApplicationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterApplicationTemplate{}, &DesignCenterApplicationTemplateList{})
}
