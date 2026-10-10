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

// API sources for ChronicleDataAccessScope, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/chronicle/v1/data_access_control.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.dataAccessScopes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleDataAccessScopeGVK = GroupVersion.WithKind("ChronicleDataAccessScope")

// ChronicleDataAccessScopeSpec defines the desired state of ChronicleDataAccessScope
// +kcc:spec:proto=google.cloud.chronicle.v1.DataAccessScope
// +kcc:required-from-proto
type ChronicleDataAccessScopeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ChronicleDataAccessScope name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The allowed labels for the scope.
	//  Either allow_all or allowed_data_access_labels needs to be provided.
	//  When provided, there has to be at least one label allowed for the scope to
	//  be valid.
	//  The logical operator for evaluation of the allowed labels is OR.
	//  E.g.: A customer with scope with allowed labels A and B will be able
	//  to see data with labeled with A or B or (A and B).
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.allowed_data_access_labels
	AllowedDataAccessLabels []DataAccessLabelReference `json:"allowedDataAccessLabels,omitempty"`

	// Optional. The denied labels for the scope.
	//  The logical operator for evaluation of the denied labels is AND.
	//  E.g.: A customer with scope with denied labels A and B won't be able
	//  to see data labeled with A and data labeled with B
	//  and data with labels A and B.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.denied_data_access_labels
	DeniedDataAccessLabels []DataAccessLabelReference `json:"deniedDataAccessLabels,omitempty"`

	// Optional. A description of the data access scope for a human reader.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.description
	Description *string `json:"description,omitempty"`

	// Optional. Whether or not the scope allows all labels, allow_all and
	//  allowed_data_access_labels are mutually exclusive and one of them must be
	//  present. denied_data_access_labels can still be used along with allow_all.
	//  When combined with denied_data_access_labels, access will be granted to all
	//  data that doesn't have labels mentioned in denied_data_access_labels. E.g.:
	//  A customer with scope with denied labels A and B and allow_all will be able
	//  to see all data except data labeled with A and data labeled with B and data
	//  with labels A and B.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.allow_all
	AllowAll *bool `json:"allowAll,omitempty"`
}

// ChronicleDataAccessScopeStatus defines the config connector machine state of ChronicleDataAccessScope
type ChronicleDataAccessScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleDataAccessScope resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleDataAccessScopeObservedState `json:"observedState,omitempty"`
}

// ChronicleDataAccessScopeObservedState is the state of the ChronicleDataAccessScope resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.DataAccessScope
type ChronicleDataAccessScopeObservedState struct {
	// Optional. The allowed labels for the scope.
	//  Either allow_all or allowed_data_access_labels needs to be provided.
	//  When provided, there has to be at least one label allowed for the scope to
	//  be valid.
	//  The logical operator for evaluation of the allowed labels is OR.
	//  E.g.: A customer with scope with allowed labels A and B will be able
	//  to see data with labeled with A or B or (A and B).
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.allowed_data_access_labels
	AllowedDataAccessLabels []DataAccessLabelReferenceObservedState `json:"allowedDataAccessLabels,omitempty"`

	// Optional. The denied labels for the scope.
	//  The logical operator for evaluation of the denied labels is AND.
	//  E.g.: A customer with scope with denied labels A and B won't be able
	//  to see data labeled with A and data labeled with B
	//  and data with labels A and B.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.denied_data_access_labels
	DeniedDataAccessLabels []DataAccessLabelReferenceObservedState `json:"deniedDataAccessLabels,omitempty"`

	// Output only. The name to be used for display to customers of the data
	//  access scope.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The time at which the data access scope was created.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the data access scope was last updated.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The user who created the data access scope.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.author
	Author *string `json:"author,omitempty"`

	// Output only. The user who last updated the data access scope.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataAccessScope.last_editor
	LastEditor *string `json:"lastEditor,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchronicledataaccessscope;gcpchronicledataaccessscopes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleDataAccessScope is the Schema for the ChronicleDataAccessScope API
// +k8s:openapi-gen=true
type ChronicleDataAccessScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleDataAccessScopeSpec   `json:"spec,omitempty"`
	Status ChronicleDataAccessScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleDataAccessScopeList contains a list of ChronicleDataAccessScope
type ChronicleDataAccessScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleDataAccessScope `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleDataAccessScope{}, &ChronicleDataAccessScopeList{})
}
