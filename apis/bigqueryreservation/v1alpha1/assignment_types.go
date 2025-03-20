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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryReservationAssignmentGVK = GroupVersion.WithKind("BigQueryReservationAssignment")

// BigQueryReservationAssignmentSpec defines the desired state of BigQueryReservationAssignment
// +kcc:proto=google.cloud.bigquery.reservation.v1.Assignment
type BigQueryReservationAssignmentSpec struct {
	// The name of reservation to create a new assignment in,
	//  or to move the assignment to.
	// +required
	ReservationRef *ReservationRef `json:"reservationRef,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self != null && has(self.projectRef) != has(self.folderRef) != has(self.organizationRef)",message="Exactly one of orojectRef or folderRef or organizationRef must be specified."
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="assignee field is immutable"
	// Immutable.
	// Required. The resource which will use the reservation. E.g.
	//  `projects/myproject`, `folders/123`, or `organizations/456`.
	// +required
	Assignee *Assignee `json:"assignee,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="jobType field is immutable"
	// Immutable.
	// Which type of jobs will use the reservation.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.job_type
	// +required
	JobType *string `json:"jobType,omitempty"`

	// Immutable. Optional.
	// The BigQueryReservationAssignment ID used for resource creation or acquisition.
	// Service-generated.Can be set only if resource acquisition .
	// For acquisition: This field must be provided to identify the Reservation resource to acquire.
	ResourceID *string `json:"resourceID,omitempty"`
}

// BigQueryReservationAssignmentStatus defines the config connector machine state of BigQueryReservationAssignment
type BigQueryReservationAssignmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigqueryReservationAssignment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryReservationAssignmentObservedState `json:"observedState,omitempty"`
}

// BigQueryReservationAssignmentObservedState is the state of the BigQueryReservationAssignment resource as most recently observed in GCP.
// +kcc:proto=google.cloud.bigquery.reservation.v1.Assignment
type BigQueryReservationAssignmentObservedState struct {
	// State of the assignment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryreservationassignment;gcpbigqueryreservationassignments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryReservationAssignment is the Schema for the BigQueryReservationAssignment API
// +k8s:openapi-gen=true
type BigQueryReservationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryReservationAssignmentSpec   `json:"spec,omitempty"`
	Status BigQueryReservationAssignmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryReservationAssignmentList contains a list of BigQueryReservationAssignment
type BigQueryReservationAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryReservationAssignment `json:"items"`
}

// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.assignee
type Assignee struct {
	// Exactly one of ProjectRef or FolderRef or OrganizationRef must be specified.
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// Exactly one of ProjectRef or FolderRef or OrganizationRef must be specified.
	// +optional
	FolderRef *refv1beta1.FolderRef `json:"folderRef"`
	// Exactly one of ProjectRef or FolderRef or OrganizationRef must be specified.
	// +optional
	OrganizationRef *refv1beta1.OrganizationRef `json:"organizationRef"`
}

func init() {
	SchemeBuilder.Register(&BigQueryReservationAssignment{}, &BigQueryReservationAssignmentList{})
}
