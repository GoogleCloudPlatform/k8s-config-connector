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

var CloudSecurityComplianceFrameworkDeploymentGVK = GroupVersion.WithKind("CloudSecurityComplianceFrameworkDeployment")

type CloudControlObservedState CloudSecurityComplianceCloudControlObservedState

// CloudSecurityComplianceFrameworkDeploymentSpec defines the desired state of CloudSecurityComplianceFrameworkDeployment
// +kcc:spec:proto=google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment
type CloudSecurityComplianceFrameworkDeploymentSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CloudSecurityComplianceFrameworkDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. target_resource_config referencing either an already existing
	//  target_resource or contains config for a target_resource to be created
	TargetResourceConfig *TargetResourceConfig `json:"targetResourceConfig,omitempty"`

	// Required. Framework resource reference
	Framework *FrameworkReference `json:"framework,omitempty"`

	// Optional. User provided description of the deployment
	Description *string `json:"description,omitempty"`

	// Required. Deployment mode and parameters for each of the cloud_controls
	//  part of the framework.
	CloudControlMetadata []CloudControlMetadata `json:"cloudControlMetadata,omitempty"`

	// Optional. To prevent concurrent updates from overwriting each other, always
	//  provide the `etag` when you update a CustomComplianceFramework. You can
	//  also provide the `etag` when you delete a CustomComplianceFramework, to
	//  help ensure that you're deleting the intended version of the
	//  CustomComplianceFramework.
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
	ComputedTargetResource *string `json:"computedTargetResource,omitempty"`

	// Output only. State of the deployment
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Output only. This field is inlined just for cloudNext because the one
	//  platform apis of CCDeployment does not exist. Beyond cloud
	//  next it will be replaced with the field below which is the
	//  references of cloud control deployment
	CcDeployments []CloudControlDeploymentObservedState `json:"ccDeployments,omitempty"`

	// Output only. The time at which the resource was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the resource last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Similarly we'll also have a field for CloudControlGroups
	CcGroupDeployments []CloudControlGroupDeploymentObservedState `json:"ccGroupDeployments,omitempty"`

	// Output only. The display name of the target resource.
	TargetResourceDisplayName *string `json:"targetResourceDisplayName,omitempty"`

	// Output only. The references to the cloud control deployments.
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
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
