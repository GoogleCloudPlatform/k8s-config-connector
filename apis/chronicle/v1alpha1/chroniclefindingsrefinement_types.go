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

// API sources for ChronicleFindingsRefinement, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/chronicle/v1/findings_refinement.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.findingsRefinements

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleFindingsRefinementGVK = GroupVersion.WithKind("ChronicleFindingsRefinement")

// ChronicleFindingsRefinementSpec defines the desired state of ChronicleFindingsRefinement
// +kcc:spec:proto=google.cloud.chronicle.v1.FindingsRefinement
// +kcc:required-from-proto
type ChronicleFindingsRefinementSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ChronicleFindingsRefinement name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Display name of the findings refinement.
	// +kcc:proto:field=google.cloud.chronicle.v1.FindingsRefinement.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The type of findings refinement.
	// +kcc:proto:field=google.cloud.chronicle.v1.FindingsRefinement.type
	Type *string `json:"type,omitempty"`

	// The query for the findings refinement. Works in conjunction with the type
	//  field to determine the findings refinement behavior. The syntax of this
	//  query is the same as a UDM search string. See the following for more
	//  information:
	//  https://cloud.google.com/chronicle/docs/investigation/udm-search
	// +kcc:proto:field=google.cloud.chronicle.v1.FindingsRefinement.query
	Query *string `json:"query,omitempty"`

	// Optional. The outcome filters for the findings refinement. These allow you
	//  to specify filters that are applied to the outcome variables in the
	//  detection. All filters must be true for a detection to match the findings
	//  refinement.
	// +kcc:proto:field=google.cloud.chronicle.v1.FindingsRefinement.outcome_filters
	OutcomeFilters []OutcomeFilter `json:"outcomeFilters,omitempty"`
}

// ChronicleFindingsRefinementStatus defines the config connector machine state of ChronicleFindingsRefinement
type ChronicleFindingsRefinementStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleFindingsRefinement resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleFindingsRefinementObservedState `json:"observedState,omitempty"`
}

// ChronicleFindingsRefinementObservedState is the state of the ChronicleFindingsRefinement resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.FindingsRefinement
type ChronicleFindingsRefinementObservedState struct {
	// Output only. The timestamp of when the findings refinement was created.
	// +kcc:proto:field=google.cloud.chronicle.v1.FindingsRefinement.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of when the findings refinement was last
	//  updated.
	// +kcc:proto:field=google.cloud.chronicle.v1.FindingsRefinement.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchroniclefindingsrefinement;gcpchroniclefindingsrefinements
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleFindingsRefinement is the Schema for the ChronicleFindingsRefinement API
// +k8s:openapi-gen=true
type ChronicleFindingsRefinement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleFindingsRefinementSpec   `json:"spec,omitempty"`
	Status ChronicleFindingsRefinementStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleFindingsRefinementList contains a list of ChronicleFindingsRefinement
type ChronicleFindingsRefinementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleFindingsRefinement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleFindingsRefinement{}, &ChronicleFindingsRefinementList{})
}
