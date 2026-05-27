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

var NetworkSecurityMirroringEndpointGroupGVK = GroupVersion.WithKind("NetworkSecurityMirroringEndpointGroup")

// NetworkSecurityMirroringEndpointGroupSpec defines the desired state of NetworkSecurityMirroringEndpointGroup
// +kcc:spec:proto=google.cloud.networksecurity.v1alpha1.MirroringEndpointGroup
type NetworkSecurityMirroringEndpointGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkSecurityMirroringEndpointGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are key/value pairs that help to organize and filter
	//  resources.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The deployment group that this DIRECT endpoint group is
	//  connected to, for example:
	//  `projects/123456789/locations/global/mirroringDeploymentGroups/my-dg`.
	//  See https://google.aip.dev/124.
	// +kubebuilder:validation:Required
	MirroringDeploymentGroupRef *refsv1beta1.NetworkSecurityMirroringDeploymentGroupRef `json:"mirroringDeploymentGroupRef"`

	// Immutable. The type of the endpoint group.
	//  If left unspecified, defaults to DIRECT.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty"`

	// Optional. User-provided description of the endpoint group.
	//  Used as additional context for the endpoint group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`
}

// NetworkSecurityMirroringEndpointGroupStatus defines the config connector machine state of NetworkSecurityMirroringEndpointGroup
type NetworkSecurityMirroringEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	// +kubebuilder:validation:Optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:validation:Optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityMirroringEndpointGroup resource in GCP.
	// +kubebuilder:validation:Optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +kubebuilder:validation:Optional
	ObservedState *NetworkSecurityMirroringEndpointGroupObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityMirroringEndpointGroupObservedState is the state of the NetworkSecurityMirroringEndpointGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.MirroringEndpointGroup
type NetworkSecurityMirroringEndpointGroupObservedState struct {
	// Output only. The timestamp when the resource was created.
	//  See https://google.aip.dev/148#timestamps.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	//  See https://google.aip.dev/148#timestamps.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. List of details about the connected deployment groups to this
	//  endpoint group.
	// +kubebuilder:validation:Optional
	ConnectedDeploymentGroups []MirroringEndpointGroup_ConnectedDeploymentGroupObservedState `json:"connectedDeploymentGroups,omitempty"`

	// Output only. The current state of the endpoint group.
	//  See https://google.aip.dev/216.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's
	//  intended state, and the system is working to reconcile them. This is part
	//  of the normal operation (e.g. adding a new association to the group). See
	//  https://google.aip.dev/128.
	// +kubebuilder:validation:Optional
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. List of associations to this endpoint group.
	// +kubebuilder:validation:Optional
	Associations []MirroringEndpointGroup_AssociationDetailsObservedState `json:"associations,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.MirroringEndpointGroup.AssociationDetails
type MirroringEndpointGroup_AssociationDetailsObservedState struct {
	// Output only. The connected association's resource name, for example:
	//  `projects/123456789/locations/global/mirroringEndpointGroupAssociations/my-ega`.
	//  See https://google.aip.dev/124.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty"`

	// Output only. The associated network, for example:
	//  projects/123456789/global/networks/my-network.
	//  See https://google.aip.dev/124.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty"`

	// Output only. Most recent known state of the association.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.MirroringEndpointGroup.ConnectedDeploymentGroup
type MirroringEndpointGroup_ConnectedDeploymentGroupObservedState struct {
	// Output only. The connected deployment group's resource name, for example:
	//  `projects/123456789/locations/global/mirroringDeploymentGroups/my-dg`.
	//  See https://google.aip.dev/124.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty"`

	// Output only. The list of locations where the deployment group is present.
	// +kubebuilder:validation:Optional
	Locations []MirroringLocationObservedState `json:"locations,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.MirroringLocation
type MirroringLocationObservedState struct {
	// Output only. The cloud location, e.g. "us-central1-a" or "asia-south1".
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty"`

	// Output only. The current state of the association in this location.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritymirroringendpointgroup;gcpnetworksecuritymirroringendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityMirroringEndpointGroup is the Schema for the NetworkSecurityMirroringEndpointGroup API
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type NetworkSecurityMirroringEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSecurityMirroringEndpointGroupSpec   `json:"spec,omitempty"`
	Status NetworkSecurityMirroringEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkSecurityMirroringEndpointGroupList contains a list of NetworkSecurityMirroringEndpointGroup
// +kubebuilder:object:root=true
type NetworkSecurityMirroringEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityMirroringEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityMirroringEndpointGroup{}, &NetworkSecurityMirroringEndpointGroupList{})
}
