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

var NetworkSecurityMirroringEndpointGroupAssociationGVK = GroupVersion.WithKind("NetworkSecurityMirroringEndpointGroupAssociation")

// NetworkSecurityMirroringEndpointGroupAssociationSpec defines the desired state of NetworkSecurityMirroringEndpointGroupAssociation
// +kcc:spec:proto=google.cloud.networksecurity.v1.MirroringEndpointGroupAssociation
type NetworkSecurityMirroringEndpointGroupAssociationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkSecurityMirroringEndpointGroupAssociation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are key/value pairs that help to organize and filter resources.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Required. The endpoint group that this association is connected to.
	// +kubebuilder:validation:Required
	MirroringEndpointGroupRef *refsv1beta1.NetworkSecurityMirroringEndpointGroupRef `json:"mirroringEndpointGroupRef"`

	// Immutable. Required. The VPC network that is associated.
	// +kubebuilder:validation:Required
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef"`
}

// NetworkSecurityMirroringEndpointGroupAssociationStatus defines the config connector machine state of NetworkSecurityMirroringEndpointGroupAssociation
type NetworkSecurityMirroringEndpointGroupAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	// +kubebuilder:validation:Optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:validation:Optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityMirroringEndpointGroupAssociation resource in GCP.
	// +kubebuilder:validation:Optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +kubebuilder:validation:Optional
	ObservedState *NetworkSecurityMirroringEndpointGroupAssociationObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityMirroringEndpointGroupAssociationObservedState is the state of the NetworkSecurityMirroringEndpointGroupAssociation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.MirroringEndpointGroupAssociation
type NetworkSecurityMirroringEndpointGroupAssociationObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The list of locations where the association is present. This
	//  information is retrieved from the linked endpoint group, and not configured
	//  as part of the association itself.
	// +kubebuilder:validation:Optional
	LocationsDetails []MirroringEndpointGroupAssociation_LocationDetailsObservedState `json:"locationsDetails,omitempty"`

	// Output only. Current state of the endpoint group association.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's
	//  intended state, and the system is working to reconcile them. This part of
	//  the normal operation (e.g. adding a new location to the target deployment
	//  group).
	// +kubebuilder:validation:Optional
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The list of locations where the association is configured.
	//  This information is retrieved from the linked endpoint group.
	// +kubebuilder:validation:Optional
	Locations []MirroringLocationObservedState `json:"locations,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.MirroringEndpointGroupAssociation.LocationDetails
type MirroringEndpointGroupAssociation_LocationDetailsObservedState struct {
	// Output only. The cloud location, e.g. "us-central1-a" or "asia-south1".
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty"`

	// Output only. The current state of the association in this location.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritymirroringendpointgroupassociation;gcpnetworksecuritymirroringendpointgroupassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityMirroringEndpointGroupAssociation is the Schema for the NetworkSecurityMirroringEndpointGroupAssociation API
// +k8s:openapi-gen=true
type NetworkSecurityMirroringEndpointGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityMirroringEndpointGroupAssociationSpec   `json:"spec,omitempty"`
	Status NetworkSecurityMirroringEndpointGroupAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityMirroringEndpointGroupAssociationList contains a list of NetworkSecurityMirroringEndpointGroupAssociation
type NetworkSecurityMirroringEndpointGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityMirroringEndpointGroupAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityMirroringEndpointGroupAssociation{}, &NetworkSecurityMirroringEndpointGroupAssociationList{})
}
