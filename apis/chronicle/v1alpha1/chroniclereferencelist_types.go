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

// API sources for ChronicleReferenceList, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/chronicle/v1/reference_list.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.referenceLists

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleReferenceListGVK = GroupVersion.WithKind("ChronicleReferenceList")

// ChronicleReferenceListSpec defines the desired state of ChronicleReferenceList
// +kcc:spec:proto=google.cloud.chronicle.v1.ReferenceList
// +kcc:required-from-proto
type ChronicleReferenceListSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ChronicleReferenceList name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. A user-provided description of the reference list.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.description
	// +required
	Description *string `json:"description,omitempty"`

	// Required. The entries of the reference list.
	//  When listed, they are returned in the order that was specified at creation
	//  or update. The combined size of the values of the reference list may not
	//  exceed 6MB.
	//  This is returned only when the view is REFERENCE_LIST_VIEW_FULL.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.entries
	// +required
	Entries []ReferenceListEntry `json:"entries,omitempty"`

	// Required. The syntax type indicating how list entries should be validated.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.syntax_type
	// +required
	SyntaxType *string `json:"syntaxType,omitempty"`

	// The scope info of the reference list.
	//  During reference list creation, if this field is not set, the reference
	//  list without scopes (an unscoped list) will be created for an unscoped
	//  user. For a scoped user, this field must be set. During reference list
	//  update, if scope_info is requested to be updated, this field must be set.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.scope_info
	ScopeInfo *ScopeInfo `json:"scopeInfo,omitempty"`
}

// ChronicleReferenceListStatus defines the config connector machine state of ChronicleReferenceList
type ChronicleReferenceListStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleReferenceList resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleReferenceListObservedState `json:"observedState,omitempty"`
}

// ChronicleReferenceListObservedState is the state of the ChronicleReferenceList resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.ReferenceList
type ChronicleReferenceListObservedState struct {
	// Output only. The unique display name of the reference list.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The timestamp when the reference list was last updated.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. The resource names for the associated self-authored Rules that
	//  use this reference list. This is returned only when the view is
	//  REFERENCE_LIST_VIEW_FULL.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.rules
	Rules []string `json:"rules,omitempty"`

	// Output only. The count of self-authored rules using the reference list.
	// +kcc:proto:field=google.cloud.chronicle.v1.ReferenceList.rule_associations_count
	RuleAssociationsCount *int32 `json:"ruleAssociationsCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchroniclereferencelist;gcpchroniclereferencelists
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleReferenceList is the Schema for the ChronicleReferenceList API
// +k8s:openapi-gen=true
type ChronicleReferenceList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleReferenceListSpec   `json:"spec,omitempty"`
	Status ChronicleReferenceListStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleReferenceListList contains a list of ChronicleReferenceList
type ChronicleReferenceListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleReferenceList `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleReferenceList{}, &ChronicleReferenceListList{})
}
