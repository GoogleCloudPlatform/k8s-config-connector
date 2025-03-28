// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNetworkEdgeSecurityServiceGVK = GroupVersion.WithKind("ComputeNetworkEdgeSecurityService")

type ComputeNetworkEdgeSecurityServiceParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +required
	Location string `json:"location"`
}

// ComputeNetworkEdgeSecurityServiceSpec defines the desired state of ComputeNetworkEdgeSecurityService
// +kcc:proto=google.cloud.compute.v1.NetworkEdgeSecurityService
type ComputeNetworkEdgeSecurityServiceSpec struct {
	ComputeNetworkEdgeSecurityServiceParent `json:"inline"`
	// The ComputeNetworkEdgeSecurityService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.description
	Description *string `json:"description,omitempty"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService. An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a NetworkEdgeSecurityService.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
	// The resource URL for the network edge security service associated with this network edge security service.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.security_policy
	SecurityPolicy *string `json:"securityPolicy,omitempty"`
}

// ComputeNetworkEdgeSecurityServiceStatus defines the config connector machine state of ComputeNetworkEdgeSecurityService
type ComputeNetworkEdgeSecurityServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNetworkEdgeSecurityService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNetworkEdgeSecurityServiceObservedState `json:"observedState,omitempty"`
}

// ComputeNetworkEdgeSecurityServiceObservedState is the state of the ComputeNetworkEdgeSecurityService resource as most recently observed in GCP.
// +kcc:proto=google.cloud.compute.v1.NetworkEdgeSecurityService
type ComputeNetworkEdgeSecurityServiceObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.id
	ID *uint64 `json:"id,omitempty"`

	// [Output only] Type of the resource. Always compute#networkEdgeSecurityService for NetworkEdgeSecurityServices
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the resource resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource with the resource id.
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEdgeSecurityService.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenetworkedgesecurityservice;gcpcomputenetworkedgesecurityservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNetworkEdgeSecurityService is the Schema for the ComputeNetworkEdgeSecurityService API
// +k8s:openapi-gen=true
type ComputeNetworkEdgeSecurityService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNetworkEdgeSecurityServiceSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEdgeSecurityServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNetworkEdgeSecurityServiceList contains a list of ComputeNetworkEdgeSecurityService
type ComputeNetworkEdgeSecurityServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNetworkEdgeSecurityService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNetworkEdgeSecurityService{}, &ComputeNetworkEdgeSecurityServiceList{})
}
