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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerSessionGVK = GroupVersion.WithKind("SpannerSession")

// SpannerSessionSpec defines the desired state of SpannerSession
// +kcc:proto=google.spanner.v1.Session
type SpannerSessionSpec struct {
	// Immutable. The Project that this resource belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The location of the cluster.
	Location string `json:"location,omitempty"`

	// The SpannerSession name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The labels for the session.
	//
	//   * Label keys must be between 1 and 63 characters long and must conform to
	//     the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`.
	//   * Label values must be between 0 and 63 characters long and must conform
	//     to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	//   * No more than 64 labels can be associated with a given session.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	// +kcc:proto:field=google.spanner.v1.Session.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The reference to the database role which created this session.
	// +kcc:proto:field=google.spanner.v1.Session.creator_role
	CreatorRoleRef *DatabaseRoleRef `json:"creatorRoleRef,omitempty"`

	// Optional. If true, specifies a multiplexed session. A multiplexed session
	//  may be used for multiple, concurrent read-only operations but can not be
	//  used for read-write transactions, partitioned reads, or partitioned
	//  queries. Multiplexed sessions can be created via
	//  [CreateSession][google.spanner.v1.Spanner.CreateSession] but not via
	//  [BatchCreateSessions][google.spanner.v1.Spanner.BatchCreateSessions].
	//  Multiplexed sessions may not be deleted nor listed.
	// +kcc:proto:field=google.spanner.v1.Session.multiplexed
	Multiplexed *bool `json:"multiplexed,omitempty"`
}

// SpannerSessionStatus defines the config connector machine state of SpannerSession
type SpannerSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerSessionObservedState `json:"observedState,omitempty"`
}

// SpannerSessionSpec defines the desired state of SpannerSession
// +kcc:proto=google.spanner.v1.Session
// SpannerSessionObservedState is the state of the SpannerSession resource as most recently observed in GCP.
type SpannerSessionObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpspannersession;gcpspannersessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerSession is the Schema for the SpannerSession API
// +k8s:openapi-gen=true
type SpannerSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerSessionSpec   `json:"spec,omitempty"`
	Status SpannerSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerSessionList contains a list of SpannerSession
type SpannerSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerSession{}, &SpannerSessionList{})
}
