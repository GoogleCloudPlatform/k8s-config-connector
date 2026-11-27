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

// API sources for DesignCenterComponent, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/authoring.proto
// +kcc:source:service-docs=https://cloud.google.com/application-design-center/docs
// +kcc:source:resource-docs=https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applicationTemplates.components

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DesignCenterComponentGVK = GroupVersion.WithKind("DesignCenterComponent")

// DesignCenterComponentSpec defines the desired state of DesignCenterComponent
// +kcc:spec:proto=google.cloud.designcenter.v1.Component
// +kcc:required-from-proto
type DesignCenterComponentSpec struct {
	// The DesignCenterApplicationTemplate this DesignCenterComponent is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DesignCenterApplicationTemplateRef `json:"parentRef"`

	// The DesignCenterComponent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The shared template used to generate the component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.shared_template_revision_uri
	// +required
	SharedTemplateRevisionURI *string `json:"sharedTemplateRevisionURI,omitempty"`

	// Optional. The component parameters.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.parameters
	Parameters []Parameter `json:"parameters,omitempty"`

	// Optional. The component display name.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The application associated with the component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.application_info
	ApplicationInfo *ComponentApplicationInfo `json:"applicationInfo,omitempty"`

	// Optional. Whether the component is exported as a separate terraform root
	//  module in a composite application template. If this is false, then native
	//  components will be exported as a submodule of a separate terraform root
	//  module.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.use_as_root_module
	UseAsRootModule *bool `json:"useAsRootModule,omitempty"`

	// Optional. Input variable aliases defined on this live component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.input_variable_aliases
	InputVariableAliases []InputVariableAlias `json:"inputVariableAliases,omitempty"`

	// Optional. The application template revision used to generate the component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.application_template_revision_uri
	ApplicationTemplateRevisionURI *string `json:"applicationTemplateRevisionURI,omitempty"`
}

// DesignCenterComponentStatus defines the config connector machine state of DesignCenterComponent
type DesignCenterComponentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterComponent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterComponentObservedState `json:"observedState,omitempty"`
}

// DesignCenterComponentObservedState is the state of the DesignCenterComponent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.Component
type DesignCenterComponentObservedState struct {
	// Optional. The component parameters.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.parameters
	Parameters []ParameterObservedState `json:"parameters,omitempty"`

	// Output only. The component creation timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The component update timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. IAM roles required by the service account to deploy the
	//  component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.roles
	Roles []string `json:"roles,omitempty"`

	// Output only. APIs required to be enabled to deploy the component, in the
	//  form of "*.googleapis.com".
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.apis
	APIs []string `json:"apis,omitempty"`

	// Output only. The connection parameters of the component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.connections_parameters
	ConnectionsParameters []ConnectionParametersObservedState `json:"connectionsParameters,omitempty"`

	// Output only. The component parameter schema, which includes possible
	//  parameter values. values.
	// +kcc:proto:field=google.cloud.designcenter.v1.Component.component_parameter_schema
	ComponentParameterSchema []ComponentParameterSchemaObservedState `json:"componentParameterSchema,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncentercomponent;gcpdesigncentercomponents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterComponent is the Schema for the DesignCenterComponent API
// +k8s:openapi-gen=true
type DesignCenterComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterComponentSpec   `json:"spec,omitempty"`
	Status DesignCenterComponentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterComponentList contains a list of DesignCenterComponent
type DesignCenterComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterComponent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterComponent{}, &DesignCenterComponentList{})
}
