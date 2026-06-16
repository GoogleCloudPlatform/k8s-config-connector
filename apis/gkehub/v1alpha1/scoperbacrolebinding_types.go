// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEHubScopeRBACRoleBindingGVK = GroupVersion.WithKind("GKEHubScopeRBACRoleBinding")

// +kcc:spec:proto=google.cloud.gkehub.v1beta.RBACRoleBinding
// +kubebuilder:validation:XValidation:rule="has(self.user) != has(self.group)",message="exactly one of user or group must be specified"
type GKEHubScopeRBACRoleBindingSpec struct {
	/* Immutable. The scope that this rbac role binding belongs to. */
	// +required
	ScopeRef *GKEHubScopeRef `json:"scopeRef"`

	/* Required. Role to bind to the principal. */
	// +required
	Role *RBACRoleBindingRole `json:"role"`

	/* Immutable. user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld".
	   One of 'user' or 'group' must be specified. */
	// +optional
	User *string `json:"user,omitempty"`

	/* Immutable. group is the group, as seen by the kubernetes cluster.
	   One of 'user' or 'group' must be specified. */
	// +optional
	Group *string `json:"group,omitempty"`

	/* Immutable. Optional. The rbacRoleBindingID of the resource; if not provided, the name of the resource will be used as the rbacRoleBindingID. */
	// +optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="rbacRoleBindingID is immutable"
	RBACRoleBindingID *string `json:"rbacRoleBindingID,omitempty"`

	/* Optional. Labels for this RBACRoleBinding. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type GKEHubScopeRBACRoleBindingObservedState struct {
	/* Output only. The time at which this rbac role binding was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The time at which this rbac role binding was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. The time at which this rbac role binding was deleted. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* Output only. Google-generated UUID for this resource. This is unique across all rbac role binding resources. If a rbac role binding resource is deleted and another with the same name is created, it will have a different uid. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. State of the rbac role binding resource. */
	// +optional
	State *RBACRoleBindingStateStatus `json:"state,omitempty"`
}

type GKEHubScopeRBACRoleBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubScopeRBACRoleBinding's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the GKEHubScopeRBACRoleBinding resource in Google Cloud. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in Google Cloud. */
	// +optional
	ObservedState *GKEHubScopeRBACRoleBindingObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubscoperbacrolebinding;gcpgkehubscoperbacrolebindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEHubScopeRBACRoleBinding is the Schema for the gkehub API
// +k8s:openapi-gen=true
type GKEHubScopeRBACRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   GKEHubScopeRBACRoleBindingSpec   `json:"spec,omitempty"`
	Status GKEHubScopeRBACRoleBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubScopeRBACRoleBindingList contains a list of GKEHubScopeRBACRoleBinding
type GKEHubScopeRBACRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubScopeRBACRoleBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubScopeRBACRoleBinding{}, &GKEHubScopeRBACRoleBindingList{})
}
