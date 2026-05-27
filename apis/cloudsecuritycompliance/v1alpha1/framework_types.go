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

var CloudSecurityFrameworkGVK = GroupVersion.WithKind("CloudSecurityFramework")

// CloudSecurityFrameworkSpec defines the desired state of CloudSecurityFramework
// +kcc:spec:proto=google.cloud.cloudsecuritycompliance.v1.Framework
type CloudSecurityFrameworkSpec struct {
	// The organization that this resource belongs to.
	// +required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`
	// The location of this resource.
	Location *string `json:"location"`

	// The CloudSecurityFramework name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The friendly name of the framework. The maximum length is 200
	//  characters.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the framework. The maximum length is 2000
	//  characters.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.description
	Description *string `json:"description,omitempty"`

	// Optional. The cloud control details that are directly added without any
	//  grouping in the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.cloud_control_details
	CloudControlDetails []CloudControlDetails `json:"cloudControlDetails,omitempty"`

	// Optional. The category of the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.category
	Category []string `json:"category,omitempty"`
}

// CloudSecurityFrameworkStatus defines the config connector machine state of CloudSecurityFramework
type CloudSecurityFrameworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudSecurityFramework resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudSecurityFrameworkObservedState `json:"observedState,omitempty"`
}

// CloudSecurityFrameworkObservedState is the state of the CloudSecurityFramework resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.Framework
type CloudSecurityFrameworkObservedState struct {
	// Output only. The major version of the framework, which is incremented in
	//  ascending order.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.major_revision_id
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`

	// Output only. The type of framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.type
	Type *string `json:"type,omitempty"`

	// Output only. The cloud providers that are supported by the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.supported_cloud_providers
	SupportedCloudProviders []string `json:"supportedCloudProviders,omitempty"`

	// Output only. The target resource types that are supported by the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.supported_target_resource_types
	SupportedTargetResourceTypes []string `json:"supportedTargetResourceTypes,omitempty"`

	// Output only. The supported enforcement modes of the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.supported_enforcement_modes
	SupportedEnforcementModes []string `json:"supportedEnforcementModes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudsecurityframework;gcpcloudsecurityframeworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudSecurityFramework is the Schema for the CloudSecurityFramework API
// +k8s:openapi-gen=true
type CloudSecurityFramework struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudSecurityFrameworkSpec   `json:"spec,omitempty"`
	Status CloudSecurityFrameworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudSecurityFrameworkList contains a list of CloudSecurityFramework
type CloudSecurityFrameworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSecurityFramework `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSecurityFramework{}, &CloudSecurityFrameworkList{})
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails
type CloudControlDetails struct {
	// Required. The name of the cloud control, in one of the following formats:
	//  `organizations/{organization}/locations/{location}/cloudControls/{cloud_control}`
	//  or
	//  `projects/{project}/locations/{location}/cloudControls/{cloud_control}`.
	//
	//  The only supported location is `global`.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.name
	Name *string `json:"name,omitempty"`

	// Required. The major version of the cloud control.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.major_revision_id
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`

	// Optional. Parameters are key-value pairs that let you provide your custom
	//  location requirements, environment requirements, or other settings that are
	//  relevant to the cloud control. An example parameter is
	//  `{"name": "location","value": "us-west-1"}`.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.parameters
	Parameters []Parameter `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.ParamValue
type ParamValue struct {
	// Optional. A string value.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Optional. A boolean value.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Optional. A repeated string.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.string_list_value
	StringListValue *StringList `json:"stringListValue,omitempty"`

	// Optional. A double value.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Optional. Sub-parameter values.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.ParamValue.oneof_value
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	OneofValue *Parameter `json:"oneofValue,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.Parameter
type Parameter struct {
	// Required. The name or key of the parameter.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Parameter.name
	Name *string `json:"name,omitempty"`

	// Required. The value of the parameter.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Parameter.parameter_value
	ParameterValue *ParamValue `json:"parameterValue,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.StringList
type StringList struct {
	// Required. The strings in the list.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.StringList.values
	Values []string `json:"values,omitempty"`
}
