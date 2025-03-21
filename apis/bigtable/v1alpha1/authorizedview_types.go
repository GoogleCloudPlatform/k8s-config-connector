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
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableAuthorizedViewGVK = GroupVersion.WithKind("BigtableAuthorizedView")

type BigtableAuthorizedViewParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +required
	InstanceRef bigtablev1beta1.InstanceRef `json:"instanceRef"`

	// +required
	TableRef bigtablev1beta1.TableRef `json:"tableRef"`
}

// BigtableAuthorizedViewSpec defines the desired state of BigtableAuthorizedView
// +kcc:proto=google.bigtable.admin.v2.AuthorizedView
type BigtableAuthorizedViewSpec struct {
	// The BigtableAuthorizedView name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +required
	BigtableAuthorizedViewParent `json:",inline"`

	// An AuthorizedView permitting access to an explicit subset of a Table.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.subset_view
	SubsetView *AuthorizedView_SubsetView `json:"subsetView,omitempty"`

	// The etag for this AuthorizedView.
	//  If this is provided on update, it must match the server's etag. The server
	//  returns ABORTED error on a mismatched etag.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.etag
	Etag *string `json:"etag,omitempty"`

	// Set to true to make the AuthorizedView protected against deletion.
	//  The parent Table and containing Instance cannot be deleted if an
	//  AuthorizedView has this bit set.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}

// BigtableAuthorizedViewStatus defines the config connector machine state of BigtableAuthorizedView
type BigtableAuthorizedViewStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableAuthorizedView resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigtableAuthorizedViewObservedState `json:"observedState,omitempty"`
}

// BigtableAuthorizedViewObservedState is the state of the BigtableAuthorizedView resource as most recently observed in GCP.
// +kcc:proto=google.bigtable.admin.v2.AuthorizedView
type BigtableAuthorizedViewObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableauthorizedview;gcpbigtableauthorizedviews
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableAuthorizedView is the Schema for the BigtableAuthorizedView API
// +k8s:openapi-gen=true
type BigtableAuthorizedView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableAuthorizedViewSpec   `json:"spec,omitempty"`
	Status BigtableAuthorizedViewStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableAuthorizedViewList contains a list of BigtableAuthorizedView
type BigtableAuthorizedViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableAuthorizedView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableAuthorizedView{}, &BigtableAuthorizedViewList{})
}
