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

// API sources for CCInsightsQAScorecardRevision, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/contactcenterinsights/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/contact-center/insights/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/contact-center/insights/docs/reference/rest/v1/projects.locations.qaScorecards.revisions

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CCInsightsQAScorecardRevisionGVK = GroupVersion.WithKind("CCInsightsQAScorecardRevision")

// CCInsightsQAScorecardRevisionSpec defines the desired state of CCInsightsQAScorecardRevision
// +kcc:spec:proto=google.cloud.contactcenterinsights.v1.QaScorecardRevision
// +kcc:required-from-proto
type CCInsightsQAScorecardRevisionSpec struct {
	// The CCInsightsQAScorecard this CCInsightsQAScorecardRevision is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *CCInsightsQAScorecardRef `json:"parentRef"`

	// The CCInsightsQAScorecardRevision name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The snapshot of the scorecard at the time of this revision's creation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardRevision.snapshot
	Snapshot *QaScorecard `json:"snapshot,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecard
type QaScorecard struct {
	// The user-specified display name of the scorecard.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A text description explaining the intent of the scorecard.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.description
	Description *string `json:"description,omitempty"`
}

// CCInsightsQAScorecardRevisionStatus defines the config connector machine state of CCInsightsQAScorecardRevision
type CCInsightsQAScorecardRevisionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CCInsightsQAScorecardRevision resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CCInsightsQAScorecardRevisionObservedState `json:"observedState,omitempty"`
}

// CCInsightsQAScorecardRevisionObservedState is the state of the CCInsightsQAScorecardRevision resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.QaScorecardRevision
type CCInsightsQAScorecardRevisionObservedState struct {
	// The snapshot of the scorecard at the time of this revision's creation.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardRevision.snapshot
	Snapshot *QaScorecardObservedState `json:"snapshot,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardRevision.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Alternative IDs for this revision of the scorecard, e.g.,
	//  `latest`.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardRevision.alternate_ids
	AlternateIDs []string `json:"alternateIDs,omitempty"`

	// Output only. State of the scorecard revision, indicating whether it's ready
	//  to be used in analysis.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecardRevision.state
	State *string `json:"state,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.contactcenterinsights.v1.QaScorecard
type QaScorecardObservedState struct {
	// Identifier. The scorecard name.
	// Format:
	// projects/{project}/locations/{location}/qaScorecards/{qa_scorecard}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which this scorecard was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the scorecard was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpccinsightsqascorecardrevision;gcpccinsightsqascorecardrevisions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CCInsightsQAScorecardRevision is the Schema for the CCInsightsQAScorecardRevision API
// +k8s:openapi-gen=true
type CCInsightsQAScorecardRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CCInsightsQAScorecardRevisionSpec   `json:"spec,omitempty"`
	Status CCInsightsQAScorecardRevisionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CCInsightsQAScorecardRevisionList contains a list of CCInsightsQAScorecardRevision
type CCInsightsQAScorecardRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CCInsightsQAScorecardRevision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CCInsightsQAScorecardRevision{}, &CCInsightsQAScorecardRevisionList{})
}
