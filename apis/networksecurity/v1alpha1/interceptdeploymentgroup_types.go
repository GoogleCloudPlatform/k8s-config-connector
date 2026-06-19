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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityInterceptDeploymentGroupGVK = GroupVersion.WithKind("NetworkSecurityInterceptDeploymentGroup")

// NetworkSecurityInterceptDeploymentGroupSpec defines the desired state of NetworkSecurityInterceptDeploymentGroup
// +kcc:spec:proto=google.cloud.networksecurity.v1.InterceptDeploymentGroup
type NetworkSecurityInterceptDeploymentGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// Optional. User-provided description of the deployment group.
	// Used as additional context for the deployment group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// The NetworkSecurityInterceptDeploymentGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are key/value pairs that help to organize and filter resources.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. The network that will be used for all child
	// deployments.
	// +kubebuilder:validation:Required
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef"`
}

// NetworkSecurityInterceptDeploymentGroupStatus defines the config connector machine state of NetworkSecurityInterceptDeploymentGroup
type NetworkSecurityInterceptDeploymentGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityInterceptDeploymentGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityInterceptDeploymentGroupObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityInterceptDeploymentGroupObservedState is the state of the NetworkSecurityInterceptDeploymentGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptDeploymentGroup
type NetworkSecurityInterceptDeploymentGroupObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the deployment group.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's intended state, and the system is working to reconcile them.
	// +kubebuilder:validation:Optional
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The list of endpoint groups that are connected to this resource.
	// +kubebuilder:validation:Optional
	ConnectedEndpointGroups []InterceptDeploymentGroup_ConnectedEndpointGroupObservedState `json:"connectedEndpointGroups,omitempty"`

	// Output only. The list of Intercept Deployments that belong to this group.
	// +kubebuilder:validation:Optional
	NestedDeployments []InterceptDeploymentGroup_DeploymentObservedState `json:"nestedDeployments,omitempty"`

	// Output only. The list of locations where the deployment group is present.
	// +kubebuilder:validation:Optional
	Locations []InterceptLocationObservedState `json:"locations,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityinterceptdeploymentgroup;gcpnetworksecurityinterceptdeploymentgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityInterceptDeploymentGroup is the Schema for the NetworkSecurityInterceptDeploymentGroup API
// +k8s:openapi-gen=true
type NetworkSecurityInterceptDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityInterceptDeploymentGroupSpec   `json:"spec,omitempty"`
	Status NetworkSecurityInterceptDeploymentGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityInterceptDeploymentGroupList contains a list of NetworkSecurityInterceptDeploymentGroup
type NetworkSecurityInterceptDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityInterceptDeploymentGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityInterceptDeploymentGroup{}, &NetworkSecurityInterceptDeploymentGroupList{})
}
