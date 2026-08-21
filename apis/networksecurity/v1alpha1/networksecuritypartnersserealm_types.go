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

var NetworkSecurityPartnerSSERealmGVK = GroupVersion.WithKind("NetworkSecurityPartnerSSERealm")

type PartnerSSERealmPanOptions struct {
	// Optional. serial_number is provided by PAN to identify GCP customer on PAN side.
	// +kubebuilder:validation:Optional
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Optional. tenant_id is provided by PAN to identify GCP customer on PAN side.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantID,omitempty"`
}

// NetworkSecurityPartnerSSERealmSpec defines the desired state of NetworkSecurityPartnerSSERealm
// +kcc:spec:proto=google.cloud.networksecurity.v1alpha1.PartnerSSERealm
type NetworkSecurityPartnerSSERealmSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The NetworkSecurityPartnerSSERealm name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels as key value pairs
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Required. value of the key to establish global handshake from SSERealm
	// +kubebuilder:validation:Required
	PairingKey *string `json:"pairingKey"`

	// Optional. VPC owned by the partner to be peered with CDEN sse_vpc in sse_project This field is deprecated. Use partner_network instead.
	// +kubebuilder:validation:Optional
	PartnerVPC *string `json:"partnerVPC,omitempty"`

	// Optional. Partner-owned network to be peered with CDEN's sse_network in sse_project
	// +kubebuilder:validation:Optional
	PartnerNetwork *string `json:"partnerNetwork,omitempty"`

	// Optional. Required only for PAN.
	// +kubebuilder:validation:Optional
	PanOptions *PartnerSSERealmPanOptions `json:"panOptions,omitempty"`
}

// NetworkSecurityPartnerSSERealmStatus defines the config connector machine state of NetworkSecurityPartnerSSERealm
type NetworkSecurityPartnerSSERealmStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityPartnerSSERealm resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityPartnerSSERealmObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityPartnerSSERealmObservedState is the state of the NetworkSecurityPartnerSSERealm resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.PartnerSSERealm
type NetworkSecurityPartnerSSERealmObservedState struct {
	// Output only. Create time stamp
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. CDEN owned VPC to be peered with partner_vpc
	//  This field is deprecated. Use sse_network instead.
	// +kubebuilder:validation:Optional
	SseVPC *string `json:"sseVPC,omitempty"`

	// Output only. CDEN owned project owning sse_vpc. It stores project id in the
	//  TTM flow, but project number in the NCCGW flow. This field will be
	//  deprecated after the partner migrates from using sse_project to using
	//  sse_project_number.
	// +kubebuilder:validation:Optional
	SseProject *string `json:"sseProject,omitempty"`

	// Output only. State of the realm. It can be either CUSTOMER_ATTACHED or
	//  CUSTOMER_DETACHED.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Output only. CDEN-owned network to be peered with partner_network
	// +kubebuilder:validation:Optional
	SseNetwork *string `json:"sseNetwork,omitempty"`

	// Output only. CDEN owned project owning sse_vpc
	// +kubebuilder:validation:Optional
	SseProjectNumber *int64 `json:"sseProjectNumber,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritypartnersserealm;gcpnetworksecuritypartnersserealms
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityPartnerSSERealm is the Schema for the NetworkSecurityPartnerSSERealm API
// +k8s:openapi-gen=true
type NetworkSecurityPartnerSSERealm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityPartnerSSERealmSpec   `json:"spec,omitempty"`
	Status NetworkSecurityPartnerSSERealmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityPartnerSSERealmList contains a list of NetworkSecurityPartnerSSERealm
type NetworkSecurityPartnerSSERealmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityPartnerSSERealm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityPartnerSSERealm{}, &NetworkSecurityPartnerSSERealmList{})
}
