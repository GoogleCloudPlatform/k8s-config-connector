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

// API sources for CloudSecurityComplianceFrameworkDeployment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/cloudsecuritycompliance/v1/deployment.proto
// +kcc:source:service-docs=https://cloud.google.com/security-command-center/docs/compliance-manager-overview
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.locations.frameworkDeployments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudSecurityComplianceFrameworkDeploymentGVK = GroupVersion.WithKind("CloudSecurityComplianceFrameworkDeployment")

// CloudSecurityComplianceFrameworkDeploymentSpec defines the desired state of CloudSecurityComplianceFrameworkDeployment
// +kcc:spec:proto=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment
// +kcc:required-from-proto
type CloudSecurityComplianceFrameworkDeploymentSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CloudSecurityComplianceFrameworkDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. target_resource_config referencing either an already existing
	//  target_resource or contains config for a target_resource to be created
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.target_resource_config
	// +required
	TargetResourceConfig *TargetResourceConfig `json:"targetResourceConfig,omitempty"`

	// Required. Framework resource reference
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.framework
	// +required
	Framework *FrameworkReference `json:"framework,omitempty"`

	// Optional. User provided description of the deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.description
	Description *string `json:"description,omitempty"`

	// Required. Deployment mode and parameters for each of the cloud_controls
	//  part of the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.cloud_control_metadata
	// +required
	CloudControlMetadata []CloudControlMetadata `json:"cloudControlMetadata,omitempty"`

	// Optional. To prevent concurrent updates from overwriting each other, always
	//  provide the `etag` when you update a CustomComplianceFramework. You can
	//  also provide the `etag` when you delete a CustomComplianceFramework, to
	//  help ensure that you're deleting the intended version of the
	//  CustomComplianceFramework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.etag
	Etag *string `json:"etag,omitempty"`
}

// CloudSecurityComplianceFrameworkDeploymentStatus defines the config connector machine state of CloudSecurityComplianceFrameworkDeployment
type CloudSecurityComplianceFrameworkDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudSecurityComplianceFrameworkDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudSecurityComplianceFrameworkDeploymentObservedState `json:"observedState,omitempty"`
}

// CloudSecurityComplianceFrameworkDeploymentObservedState is the state of the CloudSecurityComplianceFrameworkDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment
type CloudSecurityComplianceFrameworkDeploymentObservedState struct {
	// Output only. The resource on which the Framework is deployed based on the
	//  provided TargetResourceConfig. In format organizations/{organization},
	//  folders/{folder}, projects/{project} or
	//  projects/{project}/locations/{location}/applications/{application}.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.computed_target_resource
	ComputedTargetResource *string `json:"computedTargetResource,omitempty"`

	// Output only. State of the deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Output only. This field is inlined just for cloudNext because the one
	//  platform apis of CCDeployment does not exist. Beyond cloud
	//  next it will be replaced with the field below which is the
	//  references of cloud control deployment
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.cc_deployments
	CcDeployments []CloudControlDeploymentObservedState `json:"ccDeployments,omitempty"`

	// Output only. The time at which the resource was created.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the resource last updated.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Similarly we'll also have a field for CloudControlGroups
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.cc_group_deployments
	CcGroupDeployments []CloudControlGroupDeploymentObservedState `json:"ccGroupDeployments,omitempty"`

	// Output only. The display name of the target resource.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.target_resource_display_name
	TargetResourceDisplayName *string `json:"targetResourceDisplayName,omitempty"`

	// Output only. The references to the cloud control deployments.
	//  Example: If a framework deployment deploys two cloud controls,
	//  cc-deployment-1 and cc-deployment-2, then the
	//  cloud_control_deployment_references will be:
	//  {
	//   cloud_control_deployment_reference: {
	//     cloud_control_deployment:
	//     "organizations/{organization}/locations/{location}/cloudControlDeployments/cc-deployment-1"
	//   },
	//   cloud_control_deployment_reference: {
	//    cloud_control_deployment:
	//    "organizations/{organization}/locations/{location}/cloudControlDeployments/cc-deployment-2"
	//   }
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment.cloud_control_deployment_references
	CloudControlDeploymentReferences []CloudControlDeploymentReferenceObservedState `json:"cloudControlDeploymentReferences,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudsecuritycomplianceframeworkdeployment;gcpcloudsecuritycomplianceframeworkdeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudSecurityComplianceFrameworkDeployment is the Schema for the CloudSecurityComplianceFrameworkDeployment API
// +k8s:openapi-gen=true
type CloudSecurityComplianceFrameworkDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudSecurityComplianceFrameworkDeploymentSpec   `json:"spec,omitempty"`
	Status CloudSecurityComplianceFrameworkDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudSecurityComplianceFrameworkDeploymentList contains a list of CloudSecurityComplianceFrameworkDeployment
type CloudSecurityComplianceFrameworkDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSecurityComplianceFrameworkDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSecurityComplianceFrameworkDeployment{}, &CloudSecurityComplianceFrameworkDeploymentList{})
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.ParamValue
type ParamValueObservedState struct {
	// Represents a string value.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a repeated string.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.string_list_value
	StringListValue *StringListObservedState `json:"stringListValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents sub-parameter values.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.oneof_value
	OneofValue *ParameterObservedState `json:"oneofValue,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.Parameter
type ParameterObservedState struct {
	// Required. The name of the parameter.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Parameter.name
	Name *string `json:"name,omitempty"`

	// Required. The value of the parameter
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Parameter.parameter_value
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type=object
	ParameterValue *ParamValueObservedState `json:"parameterValue,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.ParameterSpec
type ParameterSpecObservedState struct {
	// Required. The name of the parameter.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.name
	Name *string `json:"name,omitempty"`

	// Optional. The display name of the parameter. The maximum length is 200
	//  characters.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the parameter. The maximum length is 2000
	//  characters.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.description
	Description *string `json:"description,omitempty"`

	// Required. if the parameter is required
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// Required. Parameter value type.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Optional. The default value of the parameter.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.default_value
	DefaultValue *ParamValueObservedState `json:"defaultValue,omitempty"`

	// Optional. List of parameter substitutions.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.substitution_rules
	SubstitutionRules []ParameterSubstitutionRule `json:"substitutionRules,omitempty"`

	// Optional. ParameterSpec for oneof attributes.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.sub_parameters
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	SubParameters []ParameterSpecObservedState `json:"subParameters,omitempty"`

	// Optional. The allowed set of values for the parameter.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParameterSpec.validation
	Validation *ValidationObservedState `json:"validation,omitempty"`
}
