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

var NetworkSecurityInterceptEndpointGroupAssociationGVK = GroupVersion.WithKind("NetworkSecurityInterceptEndpointGroupAssociation")

// NetworkSecurityInterceptEndpointGroupAssociationSpec defines the desired state of NetworkSecurityInterceptEndpointGroupAssociation
// +kcc:spec:proto=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation
type NetworkSecurityInterceptEndpointGroupAssociationSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The NetworkSecurityInterceptEndpointGroupAssociation name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are key/value pairs that help to organize and filter
	//  resources.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. The endpoint group that this association is connected
	//  to, for example:
	//  `projects/123456789/locations/global/interceptEndpointGroups/my-eg`.
	//  See https://google.aip.dev/124.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.intercept_endpoint_group
	InterceptEndpointGroupRef *refsv1beta1.NetworkSecurityInterceptEndpointGroupRef `json:"interceptEndpointGroupRef"`

	// Required. Immutable. The VPC network that is associated. for example:
	//  `projects/123456789/global/networks/my-network`.
	//  See https://google.aip.dev/124.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef"`
}

// NetworkSecurityInterceptEndpointGroupAssociationStatus defines the config connector machine state of NetworkSecurityInterceptEndpointGroupAssociation
type NetworkSecurityInterceptEndpointGroupAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityInterceptEndpointGroupAssociation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityInterceptEndpointGroupAssociationObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityInterceptEndpointGroupAssociationObservedState is the state of the NetworkSecurityInterceptEndpointGroupAssociation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation
type NetworkSecurityInterceptEndpointGroupAssociationObservedState struct {
	// Output only. The timestamp when the resource was created.
	//  See https://google.aip.dev/148#timestamps.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	//  See https://google.aip.dev/148#timestamps.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The list of locations where the association is present. This
	//  information is retrieved from the linked endpoint group, and not configured
	//  as part of the association itself.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.locations_details
	LocationsDetails []InterceptEndpointGroupAssociation_LocationDetailsObservedState `json:"locationsDetails,omitempty"`

	// Output only. Current state of the endpoint group association.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.state
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's
	//  intended state, and the system is working to reconcile them. This part of
	//  the normal operation (e.g. adding a new location to the target deployment
	//  group). See https://google.aip.dev/128.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The list of locations where the association is configured.
	//  This information is retrieved from the linked endpoint group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.locations
	Locations []InterceptLocationObservedState `json:"locations,omitempty"`

	// Output only. Identifier used by the data-path.
	//  See the NSI GENEVE format for more details:
	//  https://docs.cloud.google.com/network-security-integration/docs/understand-geneve#network_id
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.network_cookie
	NetworkCookie *int64 `json:"networkCookie,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.LocationDetails
type InterceptEndpointGroupAssociation_LocationDetailsObservedState struct {
	// Output only. The cloud location, e.g. "us-central1-a" or "asia-south1".
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.LocationDetails.location
	Location *string `json:"location,omitempty"`

	// Output only. The current state of the association in this location.
	// +kcc:proto:field=google.cloud.networksecurity.v1.InterceptEndpointGroupAssociation.LocationDetails.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityinterceptendpointgroupassociation;gcpnetworksecurityinterceptendpointgroupassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityInterceptEndpointGroupAssociation is the Schema for the NetworkSecurityInterceptEndpointGroupAssociation API
// +k8s:openapi-gen=true
type NetworkSecurityInterceptEndpointGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityInterceptEndpointGroupAssociationSpec   `json:"spec,omitempty"`
	Status NetworkSecurityInterceptEndpointGroupAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityInterceptEndpointGroupAssociationList contains a list of NetworkSecurityInterceptEndpointGroupAssociation
type NetworkSecurityInterceptEndpointGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityInterceptEndpointGroupAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityInterceptEndpointGroupAssociation{}, &NetworkSecurityInterceptEndpointGroupAssociationList{})
}
