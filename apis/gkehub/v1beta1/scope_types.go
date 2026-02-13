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

var GKEHubScopeGVK = GroupVersion.WithKind("GKEHubScope")

type GKEHubScopeSpec struct {
	/* The project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. The GKEHubScope name. If not given, the metadata.name will be used. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Labels for this Scope. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type GKEHubScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubScope's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. Google-generated UUID for this resource. This is unique across all Scope resources, even those that have been deleted. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. Time the Scope was created in UTC. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. Time the Scope was updated in UTC. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. Time the Scope was deleted in UTC. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* Output only. State of the scope resource. */
	// +optional
	State *ScopeStateStatus `json:"state,omitempty"`

	/* Output only. The unique identifier of the scope */
	// +optional
	Name *string `json:"name,omitempty"`
}

type ScopeStateStatus struct {
	/* Output only. Code describes the state of a Scope resource. */
	// +optional
	Code *string `json:"code,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubscope;gcpgkehubscopes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEHubScope is the Schema for the gkehub API
// +k8s:openapi-gen=true
type GKEHubScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEHubScopeSpec   `json:"spec,omitempty"`
	Status GKEHubScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubScopeList contains a list of GKEHubScope
type GKEHubScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubScope `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubScope{}, &GKEHubScopeList{})
}
