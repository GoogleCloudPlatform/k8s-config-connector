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

var NetworkSecurityMirroringDeploymentGroupGVK = GroupVersion.WithKind("NetworkSecurityMirroringDeploymentGroup")

// NetworkSecurityMirroringDeploymentGroupSpec defines the desired state of NetworkSecurityMirroringDeploymentGroup
// +kcc:spec:proto=google.cloud.networksecurity.v1.MirroringDeploymentGroup
type NetworkSecurityMirroringDeploymentGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NetworkSecurityMirroringDeploymentGroup name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are key/value pairs that help to organize and filter
	// resources.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. The network that will be used for all child deployments.
	// +required
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef"`

	// Optional. User-provided description of the deployment group.
	// Used as additional context for the deployment group.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.description
	Description *string `json:"description,omitempty"`
}

// NetworkSecurityMirroringDeploymentGroupStatus defines the config connector machine state of NetworkSecurityMirroringDeploymentGroup
type NetworkSecurityMirroringDeploymentGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityMirroringDeploymentGroup resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *NetworkSecurityMirroringDeploymentGroupObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityMirroringDeploymentGroupObservedState is the state of the NetworkSecurityMirroringDeploymentGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.MirroringDeploymentGroup
type NetworkSecurityMirroringDeploymentGroupObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The list of endpoint groups that are connected to this resource.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.connected_endpoint_groups
	ConnectedEndpointGroups []MirroringDeploymentGroup_ConnectedEndpointGroupObservedState `json:"connectedEndpointGroups,omitempty"`

	// Output only. The current state of the deployment group.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.state
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's
	// intended state, and the system is working to reconcile them.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The list of locations where the deployment group is present.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.locations
	Locations []MirroringLocationObservedState `json:"locations,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.MirroringDeploymentGroup.ConnectedEndpointGroup
type MirroringDeploymentGroup_ConnectedEndpointGroupObservedState struct {
	// Output only. The connected endpoint group's resource name.
	// +optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.MirroringDeploymentGroup.ConnectedEndpointGroup.name
	Name *string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritymirroringdeploymentgroup;gcpnetworksecuritymirroringdeploymentgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityMirroringDeploymentGroup is the Schema for the NetworkSecurityMirroringDeploymentGroup API
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type NetworkSecurityMirroringDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityMirroringDeploymentGroupSpec   `json:"spec,omitempty"`
	Status NetworkSecurityMirroringDeploymentGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityMirroringDeploymentGroupList contains a list of NetworkSecurityMirroringDeploymentGroup
// +kubebuilder:object:root=true
type NetworkSecurityMirroringDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityMirroringDeploymentGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityMirroringDeploymentGroup{}, &NetworkSecurityMirroringDeploymentGroupList{})
}
