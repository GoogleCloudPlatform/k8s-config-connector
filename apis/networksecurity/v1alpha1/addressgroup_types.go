// Copyright 2025 Google LLC
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

var NetworkSecurityAddressGroupGVK = GroupVersion.WithKind("NetworkSecurityAddressGroup")

type Parent struct {
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef,omitempty"`

	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// NetworkSecurityAddressGroupSpec defines the desired state of NetworkSecurityAddressGroup
// +kcc:spec:proto=google.cloud.networksecurity.v1.AddressGroup
type NetworkSecurityAddressGroupSpec struct {
	// The NetworkSecurity name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Required. Capacity of the Address Group
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.capacity
	Capacity *int64 `json:"capacity"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. List of items.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.items
	Items []string `json:"items,omitempty"`

	// Optional. Set of label tags associated with the AddressGroup resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. List of supported purposes of the Address Group.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.purpose
	Purpose []string `json:"purpose,omitempty"`

	// Required. The type of the Address Group. Possible values are "IPv4" or "IPV6".
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.type
	Type *string `json:"type"`
}

// NetworkSecurityAddressGroupStatus defines the config connector machine state of NetworkSecurityAddressGroup
type NetworkSecurityAddressGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityAddressGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityAddressGroupObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityAddressGroupObservedState is the state of the NetworkSecurityAddressGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.AddressGroup
type NetworkSecurityAddressGroupObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AddressGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityaddressgroup;gcpnetworksecurityaddressgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityAddressGroup is the Schema for the NetworkSecurityAddressGroup API
// +k8s:openapi-gen=true
type NetworkSecurityAddressGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityAddressGroupSpec   `json:"spec,omitempty"`
	Status NetworkSecurityAddressGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityAddressGroupList contains a list of NetworkSecurityAddressGroup
type NetworkSecurityAddressGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityAddressGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityAddressGroup{}, &NetworkSecurityAddressGroupList{})
}
