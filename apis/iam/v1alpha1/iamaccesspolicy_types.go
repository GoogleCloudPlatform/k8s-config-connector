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

var IamAccessPolicyGVK = GroupVersion.WithKind("IamAccessPolicy")

// IamAccessPolicySpec defines the desired state of IamAccessPolicy
// +kcc:spec:proto=google.identity.accesscontextmanager.v1.AccessPolicy
type IamAccessPolicySpec struct {
	// The organization that this resource belongs to.
	// +required
	// +kubebuilder:validation:Required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	// +required
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The IamAccessPolicy name. If not given, the metadata.name will be used.
	// +optional
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Human readable title. Does not affect behavior.
	// +required
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.title
	Title *string `json:"title"`

	// Optional. The scopes of a policy define which resources an ACM policy can restrict,
	// and where ACM resources can be referenced.
	// +optional
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// IamAccessPolicyStatus defines the config connector machine state of IamAccessPolicy
type IamAccessPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IamAccessPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *IamAccessPolicyObservedState `json:"observedState,omitempty"`
}

// IamAccessPolicyObservedState is the state of the IamAccessPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.identity.accesscontextmanager.v1.AccessPolicy
type IamAccessPolicyObservedState struct {
	// Output only. Time the `AccessPolicy` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `AccessPolicy` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. An opaque identifier for the current version of the `AccessPolicy`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamaccesspolicy;gcpiamaccesspolicys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IamAccessPolicy is the Schema for the IamAccessPolicy API
// +k8s:openapi-gen=true
type IamAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IamAccessPolicySpec   `json:"spec,omitempty"`
	Status IamAccessPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IamAccessPolicyList contains a list of IamAccessPolicy
type IamAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IamAccessPolicy{}, &IamAccessPolicyList{})
}
