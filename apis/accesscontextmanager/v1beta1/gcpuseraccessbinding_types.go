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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AccessContextManagerGCPUserAccessBindingGVK = GroupVersion.WithKind("AccessContextManagerGCPUserAccessBinding")

// AccessContextManagerGCPUserAccessBindingSpec defines the desired state of AccessContextManagerGCPUserAccessBinding
// +kcc:spec:proto=google.identity.accesscontextmanager.v1.GcpUserAccessBinding
type AccessContextManagerGCPUserAccessBindingSpec struct {
	// The organization that this resource belongs to.
	// +required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	/* Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted". */
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.GcpUserAccessBinding.access_levels
	AccessLevels []string `json:"accessLevels"`

	/* Immutable. Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht". */
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.GcpUserAccessBinding.group_key
	// +required
	GroupKey *string `json:"groupKey,omitempty"`

	// The AccessContextManagerGCPUserAccessBinding name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// AccessContextManagerGCPUserAccessBindingStatus defines the config connector machine state of AccessContextManagerGCPUserAccessBinding
type AccessContextManagerGCPUserAccessBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AccessContextManagerGCPUserAccessBinding resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AccessContextManagerGCPUserAccessBindingObservedState `json:"observedState,omitempty"`
}

// AccessContextManagerGCPUserAccessBindingObservedState is the state of the AccessContextManagerGCPUserAccessBinding resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.identity.accesscontextmanager.v1.GcpUserAccessBinding
type AccessContextManagerGCPUserAccessBindingObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanagergcpuseraccessbinding;gcpaccesscontextmanagergcpuseraccessbindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AccessContextManagerGCPUserAccessBinding is the Schema for the AccessContextManagerGCPUserAccessBinding API
// +k8s:openapi-gen=true
type AccessContextManagerGCPUserAccessBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AccessContextManagerGCPUserAccessBindingSpec   `json:"spec,omitempty"`
	Status AccessContextManagerGCPUserAccessBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AccessContextManagerGCPUserAccessBindingList contains a list of AccessContextManagerGCPUserAccessBinding
type AccessContextManagerGCPUserAccessBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerGCPUserAccessBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerGCPUserAccessBinding{}, &AccessContextManagerGCPUserAccessBindingList{})
}
