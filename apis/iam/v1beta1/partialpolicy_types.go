// Copyright 2022 Google LLC
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

package v1beta1

import (
	"time"

	bigqueryconnection "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MemberReference represents a resource with an IAM identity
type MemberReference struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
}

// MemberSource represents a source for an IAM identity
type MemberSource struct {
	// The IAMServiceAccount to be bound to the role.
	ServiceAccountRef *MemberReference `json:"serviceAccountRef,omitempty"`

	// The LoggingLogSink whose writer identity (i.e. its
	// 'status.writerIdentity') is to be bound to the role.
	LogSinkRef *MemberReference `json:"logSinkRef,omitempty"`

	// The SQLInstance whose service account (i.e. its
	// 'status.serviceAccountEmailAddress') is to be bound to the role.
	SQLInstanceRef *MemberReference `json:"sqlInstanceRef,omitempty"`

	// The ServiceIdentity whose service account (i.e., its
	// 'status.email') is to be bound to the role.
	ServiceIdentityRef *MemberReference `json:"serviceIdentityRef,omitempty"`

	// BigQueryConnectionConnection whose service account is to be bound to the role.
	// Use the Type field to specifie the connection type.
	// For "spark" connetion, the service account is in `status.observedState.spark.serviceAccountID`.
	// For "cloudSQL" connection, the service account is in `status.observedState.cloudSQL.serviceAccountID`.
	// For "cloudResource" connection, the service account is in `status.observedState.cloudResource.serviceAccountID`.
	BigQueryConnectionConnectionRef *bigqueryconnection.BigQueryConnectionServiceAccountRef `json:"bigQueryConnectionConnectionRef,omitempty"`
}

type IAMPartialPolicyMember struct {
	// The IAM identity to be bound to the role. Exactly one of
	// 'member' or 'memberFrom' must be used.
	Member Member `json:"member,omitempty"`

	// The IAM identity to be bound to the role. Exactly one of
	// 'member' or 'memberFrom' must be used, and only one subfield within
	// 'memberFrom' can be used.
	MemberFrom *MemberSource `json:"memberFrom,omitempty"`
}

// Specifies the members to bind to an IAM role.
type IAMPartialPolicyBinding struct {
	// Optional. The list of IAM users to be bound to the role.
	Members []IAMPartialPolicyMember `json:"members,omitempty"`
	// Required. The role to bind the users to.
	// +kubebuilder:validation:Pattern=^((projects|organizations)/[^/]+/)?roles/[\w_\.]+$
	Role string `json:"role"`
	// Optional. The condition under which the binding applies.
	Condition *IAMCondition `json:"condition,omitempty"`
}

// IAMPartialPolicySpec defines the desired state of IAMPartialPolicy
type IAMPartialPolicySpec struct {
	// Immutable. Required. The GCP resource to set the IAM policy on (e.g.
	// organization, project...)
	ResourceReference ResourceReference `json:"resourceRef"`
	// Optional. The list of IAM bindings managed by Config Connector.
	Bindings []IAMPartialPolicyBinding `json:"bindings,omitempty"`
}

// IAMPartialPolicyStatus defines the observed state of IAMPartialPolicy
type IAMPartialPolicyStatus struct {
	// Conditions represent the latest available observations of the IAM
	// policy's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// LastAppliedBindings is the list of IAM bindings that were most recently applied by Config Connector.
	LastAppliedBindings []IAMPolicyBinding `json:"lastAppliedBindings,omitempty"`
	// AllBindings surfaces all IAM bindings for the referenced resource.
	AllBindings []IAMPolicyBinding `json:"allBindings,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPartialPolicy is the Schema for the iampartialpolicy API
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].status",description="When 'True' the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].reason",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",type="date",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime"
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:resource:categories=gcp,shortName=gcpiampartialpolicy;gcpiampartialpolicies
// +k8s:openapi-gen=true
type IAMPartialPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPartialPolicySpec   `json:"spec,omitempty"`
	Status IAMPartialPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPartialPolicyList contains a list of IAMPartialPolicy
type IAMPartialPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMPartialPolicy `json:"items"`
}

const IAMPartialPolicyReconcileInterval = 10 * time.Minute

func init() {
	SchemeBuilder.Register(&IAMPartialPolicy{}, &IAMPartialPolicyList{})
}
