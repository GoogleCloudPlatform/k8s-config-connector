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

var NetworkSecurityInterceptEndpointGroupGVK = GroupVersion.WithKind("NetworkSecurityInterceptEndpointGroup")

// NetworkSecurityInterceptEndpointGroupSpec defines the desired state of NetworkSecurityInterceptEndpointGroup
// +kcc:spec:proto=google.cloud.networksecurity.v1.InterceptEndpointGroup
type NetworkSecurityInterceptEndpointGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkSecurityInterceptEndpointGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are key/value pairs that help to organize and filter
	// resources.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. The deployment group that this endpoint group is
	// connected to, for example:
	// `projects/123456789/locations/global/interceptDeploymentGroups/my-dg`.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.intercept_deployment_group
	// +kubebuilder:validation:Required
	InterceptDeploymentGroupRef *refsv1beta1.NetworkSecurityInterceptDeploymentGroupRef `json:"interceptDeploymentGroupRef"`

	// Optional. User-provided description of the endpoint group.
	// Used as additional context for the endpoint group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.description
	Description *string `json:"description,omitempty"`
}

// NetworkSecurityInterceptEndpointGroupStatus defines the config connector machine state of NetworkSecurityInterceptEndpointGroup
type NetworkSecurityInterceptEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityInterceptEndpointGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityInterceptEndpointGroupObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityInterceptEndpointGroupObservedState is the state of the NetworkSecurityInterceptEndpointGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptEndpointGroup
type NetworkSecurityInterceptEndpointGroupObservedState struct {
	// Output only. The timestamp when the resource was created.
	// See https://google.aip.dev/148#timestamps.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	// See https://google.aip.dev/148#timestamps.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Details about the connected deployment group to this endpoint
	// group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.connected_deployment_group
	ConnectedDeploymentGroup *InterceptEndpointGroup_ConnectedDeploymentGroupObservedState `json:"connectedDeploymentGroup,omitempty"`

	// Output only. The current state of the endpoint group.
	// See https://google.aip.dev/216.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.state
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's
	// intended state, and the system is working to reconcile them. This is part
	// of the normal operation (e.g. adding a new association to the group). See
	// https://google.aip.dev/128.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. List of associations to this endpoint group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.associations
	Associations []InterceptEndpointGroup_AssociationDetailsObservedState `json:"associations,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptEndpointGroup.ConnectedDeploymentGroup
type InterceptEndpointGroup_ConnectedDeploymentGroupObservedState struct {
	// Output only. The connected deployment group's resource name, for example:
	// `projects/123456789/locations/global/interceptDeploymentGroups/my-dg`.
	// See https://google.aip.dev/124.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.ConnectedDeploymentGroup.name
	Name *string `json:"name,omitempty"`

	// Output only. The list of locations where the deployment group is present.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.ConnectedDeploymentGroup.locations
	Locations []InterceptLocationObservedState `json:"locations,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptLocation
type InterceptLocationObservedState struct {
	// Output only. The cloud location, e.g. "us-central1-a" or "asia-south1".
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptLocation.location
	Location *string `json:"location,omitempty"`

	// Output only. The current state of the association in this location.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptLocation.state
	State *string `json:"state,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptEndpointGroup.AssociationDetails
type InterceptEndpointGroup_AssociationDetailsObservedState struct {
	// Output only. The connected association's resource name, for example:
	// `projects/123456789/locations/global/interceptEndpointGroupAssociations/my-ega`.
	// See https://google.aip.dev/124.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.AssociationDetails.name
	Name *string `json:"name,omitempty"`

	// Output only. The associated network, for example:
	// projects/123456789/global/networks/my-network.
	// See https://google.aip.dev/124.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.AssociationDetails.network
	Network *string `json:"network,omitempty"`

	// Output only. Most recent known state of the association.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroup.AssociationDetails.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityinterceptendpointgroup;gcpnetworksecurityinterceptendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityInterceptEndpointGroup is the Schema for the NetworkSecurityInterceptEndpointGroup API
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
type NetworkSecurityInterceptEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityInterceptEndpointGroupSpec   `json:"spec,omitempty"`
	Status NetworkSecurityInterceptEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityInterceptEndpointGroupList contains a list of NetworkSecurityInterceptEndpointGroup
type NetworkSecurityInterceptEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityInterceptEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityInterceptEndpointGroup{}, &NetworkSecurityInterceptEndpointGroupList{})
}
