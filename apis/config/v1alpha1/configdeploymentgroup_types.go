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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConfigDeploymentGroupGVK = GroupVersion.WithKind("ConfigDeploymentGroup")

// ConfigDeploymentGroupSpec defines the desired state of ConfigDeploymentGroup
// +kcc:spec:proto=google.cloud.config.v1.DeploymentGroup
type ConfigDeploymentGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ConfigDeploymentGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-defined metadata for the deployment group.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Arbitrary key-value metadata storage e.g. to help client tools
	//  identify deployment group during automation. See
	//  https://google.aip.dev/148#annotations for details on format and size
	//  limitations.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// The deployment units of the deployment group in a DAG like structure.
	//  When a deployment group is being provisioned, the deployment units are
	//  deployed in a DAG order.
	//  The provided units must be in a DAG order, otherwise an error will be
	//  returned.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.deployment_units
	DeploymentUnits []DeploymentUnit `json:"deploymentUnits,omitempty"`
}

// ConfigDeploymentGroupStatus defines the config connector machine state of ConfigDeploymentGroup
type ConfigDeploymentGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConfigDeploymentGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ConfigDeploymentGroupObservedState `json:"observedState,omitempty"`
}

// ConfigDeploymentGroupObservedState is the state of the ConfigDeploymentGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.config.v1.DeploymentGroup
type ConfigDeploymentGroupObservedState struct {
	// Output only. Time when the deployment group was created.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the deployment group was last updated.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the deployment group.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information regarding the current state.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.state_description
	StateDescription *string `json:"stateDescription,omitempty"`

	// Output only. The provisioning state of the deployment group.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.provisioning_state
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Output only. Additional information regarding the current provisioning
	//  state.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.provisioning_state_description
	ProvisioningStateDescription *string `json:"provisioningStateDescription,omitempty"`

	// Output only. The error status of the deployment group provisioning or
	//  deprovisioning.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentGroup.provisioning_error
	ProvisioningError *common.Status `json:"provisioningError,omitempty"`
}

// +kcc:proto=google.cloud.config.v1.DeploymentUnit
type DeploymentUnit struct {
	// The id of the deployment unit. Must be unique within the deployment group.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentUnit.id
	ID *string `json:"id,omitempty"`

	// Optional. The name of the deployment to be provisioned.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentUnit.deployment
	DeploymentRef *ConfigDeploymentRef `json:"deploymentRef,omitempty"`

	// Required. The IDs of the deployment units within the deployment group that
	//  this unit depends on.
	// +kcc:proto:field=google.cloud.config.v1.DeploymentUnit.dependencies
	Dependencies []string `json:"dependencies,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconfigdeploymentgroup;gcpconfigdeploymentgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConfigDeploymentGroup is the Schema for the ConfigDeploymentGroup API
// +k8s:openapi-gen=true
type ConfigDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConfigDeploymentGroupSpec   `json:"spec,omitempty"`
	Status ConfigDeploymentGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConfigDeploymentGroupList contains a list of ConfigDeploymentGroup
type ConfigDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDeploymentGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigDeploymentGroup{}, &ConfigDeploymentGroupList{})
}
