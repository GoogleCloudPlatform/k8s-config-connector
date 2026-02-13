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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEHubScopeRBACRoleBindingGVK = GroupVersion.WithKind("GKEHubScopeRBACRoleBinding")

type GKEHubScopeRBACRoleBindingSpec struct {
	/* The project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. The GKEHubScopeRBACRoleBinding name. If not given, the metadata.name will be used. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. */
	ScopeRef GKEHubScopeRef `json:"scopeRef"`

	/* Role to bind to the principal. */
	Role ScopeRBACRole `json:"role"`

	/* Principal that is be authorized in the cluster (at least of one the oneof is required). group is the group, as seen by the kubernetes cluster. */
	// +optional
	Group *string `json:"group,omitempty"`

	/* Principal that is be authorized in the cluster (at least of one the oneof is required). user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld" */
	// +optional
	User *string `json:"user,omitempty"`

	/* Labels for this ScopeRBACRoleBinding. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type ScopeRBACRole struct {
	/* PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: ["UNKNOWN", "ADMIN", "EDIT", "VIEW"] */
	// +optional
	PredefinedRole *string `json:"predefinedRole,omitempty"`
}

type GKEHubScopeRBACRoleBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubScopeRBACRoleBinding's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. Google-generated UUID for this resource. This is unique across all ScopeRBACRoleBinding resources, even those that have been deleted. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. Time the ScopeRBACRoleBinding was created in UTC. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. Time the ScopeRBACRoleBinding was updated in UTC. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. Time the ScopeRBACRoleBinding was deleted in UTC. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* Output only. State of the scope rbac role binding resource. */
	// +optional
	State *ScopeRBACRoleBindingStateStatus `json:"state,omitempty"`

	/* Output only. The unique identifier of the scope rbac role binding */
	// +optional
	Name *string `json:"name,omitempty"`
}

type ScopeRBACRoleBindingStateStatus struct {
	/* Output only. Code describes the state of a ScopeRBACRoleBinding resource. */
	// +optional
	Code *string `json:"code,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubscoperbacrolebinding;gcpgkehubscoperbacrolebindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
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
