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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableAuthorizedViewGVK = GroupVersion.WithKind("BigtableAuthorizedView")

type BigtableAuthorizedViewParent struct {
	// Immutable. The instance to create the authorized view in.
	// +required
	InstanceRef InstanceRef `json:"instanceRef"`

	// Immutable. The table to create the authorized view in.
	// +required
	TableRef TableRef `json:"tableRef"`
}

// BigtableAuthorizedViewSpec defines the desired state of BigtableAuthorizedView
// +kcc:spec:proto=google.bigtable.admin.v2.AuthorizedView
type BigtableAuthorizedViewSpec struct {
	// The BigtableAuthorizedView name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	BigtableAuthorizedViewParent `json:",inline"`

	// An AuthorizedView permitting access to an explicit subset of a Table.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.subset_view
	SubsetView *AuthorizedView_SubsetView `json:"subsetView,omitempty"`

	// Set to true to make the AuthorizedView protected against deletion.
	//  The parent Table and containing Instance cannot be deleted if an
	//  AuthorizedView has this bit set.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.SubsetView
type AuthorizedView_SubsetView struct {
	// Row prefixes to be included in the AuthorizedView.
	//  To provide access to all rows, include the empty string as a prefix ("").
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.subset_view.row_prefixes
	RowPrefixes [][]byte `json:"rowPrefixes,omitempty"`

	// Map from family ID to its subsets to be included in the AuthorizedView.
	FamilySubsets []AuthorizedView_FamilySubsets `json:"familySubsets,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.FamilySubsets
type AuthorizedView_FamilySubsets struct {
	// Name of the family.
	Name string `json:"name"`

	// Individual exact column qualifiers to be included in the AuthorizedView.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.FamilySubsets.qualifiers
	Qualifiers [][]byte `json:"qualifiers,omitempty"`

	// Prefixes for qualifiers to be included in the AuthorizedView. Every
	//  qualifier starting with one of these prefixes is included in the
	//  AuthorizedView. To provide access to all qualifiers, include the empty
	//  string as a prefix ("").
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.FamilySubsets.qualifier_prefixes
	QualifierPrefixes [][]byte `json:"qualifierPrefixes,omitempty"`
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
// +kcc:observedstate:proto=google.bigtable.admin.v2.AuthorizedView
type BigtableAuthorizedViewObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableauthorizedview;gcpbigtableauthorizedviews
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
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
