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

var CloudSecurityComplianceFrameworkGVK = GroupVersion.WithKind("CloudSecurityComplianceFramework")

// CloudSecurityComplianceFrameworkSpec defines the desired state of CloudSecurityComplianceFramework
// +kcc:spec:proto=google.cloud.cloudsecuritycompliance.v1.Framework
type CloudSecurityComplianceFrameworkSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CloudSecurityComplianceFramework name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Display name of the framework. The maximum length is 200 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the framework. The maximum length is 2000 characters.
	Description *string `json:"description,omitempty"`

	// Optional. The details of the cloud control groups included in the framework.
	CloudControlGroupDetails []Framework_CloudControlGroupDetails `json:"cloudControlGroupDetails,omitempty"`

	// Optional. The details of the cloud controls directly added without any grouping in the framework.
	CloudControlDetails []CloudControlDetails `json:"cloudControlDetails,omitempty"`

	// Optional. The category of the framework.
	Category []string `json:"category,omitempty"`
}

// CloudSecurityComplianceFrameworkStatus defines the config connector machine state of CloudSecurityComplianceFramework
type CloudSecurityComplianceFrameworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudSecurityComplianceFramework resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudSecurityComplianceFrameworkObservedState `json:"observedState,omitempty"`
}

// CloudSecurityComplianceFrameworkObservedState is the state of the CloudSecurityComplianceFramework resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.Framework
type CloudSecurityComplianceFrameworkObservedState struct {
	// Output only. Major revision of the framework incremented in ascending order.
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`

	// Output only. The type of the framework. The default is TYPE_CUSTOM.
	Type *string `json:"type,omitempty"`

	// Optional. The details of the cloud control groups included in the framework.
	CloudControlGroupDetails []Framework_CloudControlGroupDetailsObservedState `json:"cloudControlGroupDetails,omitempty"`

	// Output only. cloud providers supported.
	SupportedCloudProviders []string `json:"supportedCloudProviders,omitempty"`

	// Output only. target resource types supported by the Framework.
	SupportedTargetResourceTypes []string `json:"supportedTargetResourceTypes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudsecuritycomplianceframework;gcpcloudsecuritycomplianceframeworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudSecurityComplianceFramework is the Schema for the CloudSecurityComplianceFramework API
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
type CloudSecurityComplianceFramework struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudSecurityComplianceFrameworkSpec   `json:"spec,omitempty"`
	Status CloudSecurityComplianceFrameworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudSecurityComplianceFrameworkList contains a list of CloudSecurityComplianceFramework
type CloudSecurityComplianceFrameworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSecurityComplianceFramework `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSecurityComplianceFramework{}, &CloudSecurityComplianceFrameworkList{})
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails
type CloudControlDetails struct {
	// Required. The name of the CloudControl in the format:
	//  “organizations/{organization}/locations/{location}/
	//  cloudControls/{cloud-control}”
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.name
	CloudControlRef *CloudSecurityComplianceCloudControlRef `json:"cloudControlRef,omitempty"`

	// Required. Major revision of cloudcontrol
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.major_revision_id
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`

	// Optional. Parameters is a key-value pair that is required by the
	//  CloudControl. The specification of these parameters will be present in
	//  cloudcontrol.Eg: { "name": "location","value": "us-west-1"}.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlDetails.parameters
	Parameters []Parameter `json:"parameters,omitempty"`
}

// kcc:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup
type CloudControlGroup struct {
	// Required. The name of the cloud control group in the format:
	//  “organizations/{organization}/locations/{location}/
	//  cloudControlGroups/{cloud-control-group}”
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.name
	CloudControlGroupRef *CloudSecurityComplianceCloudControlGroupRef `json:"cloudControlGroupRef,omitempty"`

	// Optional. The description of the cloud control group.The maximum length is
	//  2000 characters.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. The control identifier used to fetch the findings. This is same
	//  as the control report name.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.control_id
	ControlID *string `json:"controlID,omitempty"`

	// Required. The details of the cloud controls to be referred to in the
	//  framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.cloud_control_details
	CloudControlDetails []CloudControlDetails `json:"cloudControlDetails,omitempty"`

	// Optional. Major revision of the cloud control group.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.major_revision_id
	MajorRevisionID *int64 `json:"majorRevisionID,omitempty"`

	// Optional. The industry-defined Control assciated with the cloud controls in
	//  this group.
	//  organizations/{organization}/locations/{location}/controls/{control_id}
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.control
	ControlRef *CloudSecurityComplianceControlRef `json:"controlRef,omitempty"`
}

// kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup
type CloudControlGroupObservedState struct {
	// Optional. Output only. The type of the cloud control group. Default is
	//  TYPE_CUSTOM.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.CloudControlGroup.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.cloudsecuritycompliance.v1.Framework.CloudControlGroupDetails
type Framework_CloudControlGroupDetails struct {
	// The cloud control group included in the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.CloudControlGroupDetails.cloud_control_group
	CloudControlGroup *CloudControlGroup `json:"cloudControlGroup,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.cloudsecuritycompliance.v1.Framework.CloudControlGroupDetails
type Framework_CloudControlGroupDetailsObservedState struct {
	// The cloud control group included in the framework.
	// +kcc:proto:field=google.cloud.cloudsecuritycompliance.v1.Framework.CloudControlGroupDetails.cloud_control_group
	CloudControlGroup *CloudControlGroupObservedState `json:"cloudControlGroup,omitempty"`
}
