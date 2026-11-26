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

// API sources for DesignCenterApplication, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/deployment_resources.proto
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

var DesignCenterApplicationGVK = GroupVersion.WithKind("DesignCenterApplication")

// DesignCenterApplicationSpec defines the desired state of DesignCenterApplication
// +kcc:spec:proto=google.cloud.designcenter.v1.Application
// +kcc:required-from-proto
type DesignCenterApplicationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/spaces/{space} this resource belongs to.
	// +kcc:guess
	// SpaceRef *DesignCenterSpaceRef `json:"spaceRef,omitempty"`

	// The DesignCenterApplication name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The application deployment source.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.source
	// +required
	Source *DeploymentSource `json:"source,omitempty"`

	// Optional. A list of parameters to attach to the deployment source object,
	//  which is a catalog entry or application template snapshot.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.app_parameters
	AppParameters []Parameter `json:"appParameters,omitempty"`

	// Optional. A list of component parameters to associate with the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.component_parameters
	ComponentParameters []ComponentParameters `json:"componentParameters,omitempty"`

	// Optional. Your own service account that you use to deploy an application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Display name of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.description
	Description *string `json:"description,omitempty"`

	// Optional. Attributes of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.attributes
	Attributes *Attributes `json:"attributes,omitempty"`

	// Required. Scope of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.scope
	// +required
	Scope *Scope `json:"scope,omitempty"`

	// Optional. Import existing resources into the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.import_existing_resources
	ImportExistingResources *bool `json:"importExistingResources,omitempty"`

	// Optional. The region where the application is deployed.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.deployment_region
	DeploymentRegion *string `json:"deploymentRegion,omitempty"`

	// Optional. Deployment project of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.deployment_project
	DeploymentProject *string `json:"deploymentProject,omitempty"`

	// Optional. Connection configuration for the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.connection_configs
	ConnectionConfigs []ConnectionConfig `json:"connectionConfigs,omitempty"`

	// Optional. The type of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.type
	Type *string `json:"type,omitempty"`

	// Optional. The deployment target of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.deployment_target
	DeploymentTarget *DeploymentTarget `json:"deploymentTarget,omitempty"`

	// Optional. Parameters that are required for composite applications.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.composite_application_parameters
	CompositeApplicationParameters *CompositeApplicationParameters `json:"compositeApplicationParameters,omitempty"`

	// Optional. Specifies the strategy to use when updating the application
	//  parameters while updating the application template revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.params_update_strategy
	ParamsUpdateStrategy *string `json:"paramsUpdateStrategy,omitempty"`

	// Optional. Project-specific configurations required for application
	//  deployment, including IAM roles and APIs. The list can contain up to 100
	//  items.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.roles_apis_config
	RolesAPIsConfig []RolesAndApisConfig `json:"rolesAPIsConfig,omitempty"`
}

// DesignCenterApplicationStatus defines the config connector machine state of DesignCenterApplication
type DesignCenterApplicationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterApplication resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterApplicationObservedState `json:"observedState,omitempty"`
}

// DesignCenterApplicationObservedState is the state of the DesignCenterApplication resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.Application
type DesignCenterApplicationObservedState struct {
	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. A list of parameters to attach to the deployment source object,
	//  which is a catalog entry or application template snapshot.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.app_parameters
	AppParameters []ParameterObservedState `json:"appParameters,omitempty"`

	// Optional. A list of component parameters to associate with the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.component_parameters
	ComponentParameters []ComponentParametersObservedState `json:"componentParameters,omitempty"`

	// Output only. Preview reference for the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.preview_reference
	PreviewReference *string `json:"previewReference,omitempty"`

	// Output only. [Output only] Optional Infra Manager deployment Id with
	//  revision
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.deployment_revision
	DeploymentRevision *string `json:"deploymentRevision,omitempty"`

	// Output only. The App Hub App associated with the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.apphub_application
	ApphubApplication *string `json:"apphubApplication,omitempty"`

	// Output only. [Output only] Serialized application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.serialized_application_template
	SerializedApplicationTemplate *SerializedApplicationTemplateObservedState `json:"serializedApplicationTemplate,omitempty"`

	// Output only. Deployment state of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.state
	State *string `json:"state,omitempty"`

	// Output only. Deployment metadata of the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.deployment_metadata
	DeploymentMetadata *DeploymentMetadataObservedState `json:"deploymentMetadata,omitempty"`

	// Output only. List of project parameters for the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.project_parameters
	ProjectParameters []ProjectParametersObservedState `json:"projectParameters,omitempty"`

	// Output only. The updated template revision because of which the application
	//  is outdated.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.updated_template_revision
	UpdatedTemplateRevision *UpdatedTemplateRevisionObservedState `json:"updatedTemplateRevision,omitempty"`

	// Optional. Connection configuration for the application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.connection_configs
	ConnectionConfigs []ConnectionConfigObservedState `json:"connectionConfigs,omitempty"`

	// Output only. Details of the location where the IaC for this application was
	//  last successfully exported.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.artifact_location
	ArtifactLocation *ArtifactLocationObservedState `json:"artifactLocation,omitempty"`

	// Output only. The composition type of the application: STANDARD OR
	//  COMPOSITE.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.composition_type
	CompositionType *string `json:"compositionType,omitempty"`

	// Optional. Parameters that are required for composite applications.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.composite_application_parameters
	CompositeApplicationParameters *CompositeApplicationParametersObservedState `json:"compositeApplicationParameters,omitempty"`

	// Output only. The deployment group metadata of the application. This is
	//  applicable only for composite applications.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.deployment_group_metadata
	DeploymentGroupMetadata *DeploymentGroupMetadataObservedState `json:"deploymentGroupMetadata,omitempty"`

	// Optional. Project-specific configurations required for application
	//  deployment, including IAM roles and APIs. The list can contain up to 100
	//  items.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.roles_apis_config
	RolesAPIsConfig []RolesAndApisConfigObservedState `json:"rolesAPIsConfig,omitempty"`

	// Output only. The input output variables parsed from the imported terraform
	//  code.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.template_metadata
	TemplateMetadata *TfBlueprintMetadataObservedState `json:"templateMetadata,omitempty"`

	// Output only. Indicates the IaC source format of this application.
	// +kcc:proto:field=google.cloud.designcenter.v1.Application.iac_source
	IacSource *string `json:"iacSource,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncenterapplication;gcpdesigncenterapplications
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterApplication is the Schema for the DesignCenterApplication API
// +k8s:openapi-gen=true
type DesignCenterApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterApplicationSpec   `json:"spec,omitempty"`
	Status DesignCenterApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterApplicationList contains a list of DesignCenterApplication
type DesignCenterApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterApplication{}, &DesignCenterApplicationList{})
}
