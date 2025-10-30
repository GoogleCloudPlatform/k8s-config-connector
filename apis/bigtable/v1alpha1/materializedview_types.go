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
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableMaterializedViewGVK = GroupVersion.WithKind("BigtableMaterializedView")

type BigtableMaterializedViewParent struct {
	// +reqired
	InstanceRef *v1beta1.InstanceRef `json:"instanceRef,omitempty"`
}

// BigtableMaterializedViewSpec defines the desired state of BigtableMaterializedView
// +kcc:spec:proto=google.bigtable.admin.v2.MaterializedView
type BigtableMaterializedViewSpec struct {
	// The BigtableMaterializedView name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	BigtableMaterializedViewParent `json:",inline"`

	// Immutable. MaterializedView's select query.
	Query *string `json:"query,omitempty"`

	// Optional. Set to true to make the MaterializedView protected against deletion.
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}

// BigtableMaterializedViewStatus defines the config connector machine state of BigtableMaterializedView
type BigtableMaterializedViewStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableMaterializedView resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtablematerializedview;gcpbigtablematerializedviews
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableMaterializedView is the Schema for the BigtableMaterializedView API
// +k8s:openapi-gen=true
type BigtableMaterializedView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableMaterializedViewSpec   `json:"spec,omitempty"`
	Status BigtableMaterializedViewStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableMaterializedViewList contains a list of BigtableMaterializedView
type BigtableMaterializedViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableMaterializedView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableMaterializedView{}, &BigtableMaterializedViewList{})
}
