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

var GKEHubScopeGVK = GroupVersion.WithKind("GKEHubScope")

type GKEHubScopeStateStatus struct {
	/* Output only. Code describes the state of a Scope resource. Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING */
	Code *string `json:"code,omitempty"`
}

// +kcc:spec:proto=google.cloud.gkehub.v1beta.Scope
type GKEHubScopeSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. The location for the resource. Typically "global". */
	Location *string `json:"location"`

	/* Optional. Labels for the GKEHubScope. */
	// +optional
	NamespaceLabels map[string]string `json:"namespaceLabels,omitempty"`

	/* Immutable. Optional. The resourceID of the resource; if not provided, the name of the resource will be used as the resourceID. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type GKEHubScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubScope's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The time at which this scope was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The time at which this scope was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. The time at which this scope was deleted. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* Output only. Google-generated UUID for this resource. This is unique across all scope resources. If a scope resource is deleted and another with the same name is created, it will have a different uid. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. State of the scope resource. */
	// +optional
	State *GKEHubScopeStateStatus `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubscope;gcpgkehubscopes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
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
