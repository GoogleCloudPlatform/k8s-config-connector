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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEHubNamespaceGVK = GroupVersion.WithKind("GKEHubNamespace")

type GKEHubNamespaceStateStatus struct {
	/* Output only. Code describes the state of a Namespace resource. Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING */
	// +kubebuilder:validation:Enum=CODE_UNSPECIFIED;CREATING;READY;DELETING;UPDATING
	Code *string `json:"code,omitempty"`
}

// +kcc:spec:proto=google.cloud.gkehub.v1beta.Namespace
type GKEHubNamespaceSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	/* Immutable. The location for the resource */
	Location *string `json:"location,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The scope that this namespace belongs to. */
	ScopeRef *GKEHubScopeRef `json:"scopeRef"`

	/* Immutable. The namespaceID of the resource. */
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="namespaceID is immutable"
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern="^[a-z0-9]([-a-z0-9]*[a-z0-9])?$"
	NamespaceID *string `json:"namespaceID"`

	/* Optional. Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound to the parent Scope. Scope-level labels (namespace_labels in the Fleet Scope resource) take precedence over Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant. */
	// +optional
	NamespaceLabels map[string]string `json:"namespaceLabels,omitempty"`

	/* Optional. Labels for this Namespace. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type GKEHubNamespaceObservedState struct {
	/* Output only. The time at which this namespace was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The time at which this namespace was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. The time at which this namespace was deleted. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* Output only. Google-generated UUID for this resource. This is unique across all namespace resources. If a namespace resource is deleted and another with the same name is created, it will have a different uid. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. State of the namespace resource. */
	// +optional
	State *string `json:"state,omitempty"`
}

type GKEHubNamespaceStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubNamespace's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the GKEHubNamespace resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *GKEHubNamespaceObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubnamespace;gcpgkehubnamespaces
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEHubNamespace is the Schema for the gkehub API
// +k8s:openapi-gen=true
type GKEHubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   GKEHubNamespaceSpec   `json:"spec,omitempty"`
	Status GKEHubNamespaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubNamespaceList contains a list of GKEHubNamespace
type GKEHubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubNamespace{}, &GKEHubNamespaceList{})
}
