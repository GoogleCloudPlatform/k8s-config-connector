// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MultiClusterLeaseSpec defines the desired state of MultiClusterLease
type MultiClusterLeaseSpec struct {
	// LeaseDurationSeconds is the duration that non-leader candidates will
	// wait to force acquire leadership.
	LeaseDurationSeconds *int32 `json:"leaseDurationSeconds,omitempty"`

	// RenewDeadlineSeconds is the duration that a leader candidate will
	// wait to renew its leadership.
	RenewDeadlineSeconds *int32 `json:"renewDeadlineSeconds,omitempty"`

	// RetryPeriodSeconds is the period between attempts to acquire or renew
	// leadership.
	RetryPeriodSeconds *int32 `json:"retryPeriodSeconds,omitempty"`
}

// MultiClusterLeaseStatus defines the observed state of MultiClusterLease
type MultiClusterLeaseStatus struct {
	// Conditions represent the latest available observations of the MultiClusterLease's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the .metadata.generation of the MultiClusterLease CR
	// that was last processed to produce this status.
	// Used by clients to determine if the status reflects their latest spec changes.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// IsLeadingCluster is true if the local cluster is holding the global lease.
	// This is the source of truth for leadership.
	IsLeadingCluster bool `json:"isLeadingCluster,omitempty"`

	// GlobalHolderIdentity is the globally unique identifier of the candidate
	// last observed as holding the global lock for this lease scope.
	GlobalHolderIdentity *string `json:"globalHolderIdentity,omitempty"`

	// GlobalRenewTime is the timestamp when the GlobalHolderIdentity
	// was last observed to renew the lease.
	GlobalRenewTime *string `json:"globalRenewTime,omitempty"`

	// GlobalLeaseDurationSeconds is the lease duration (in seconds) currently
	// being enforced by the global leader.
	GlobalLeaseDurationSeconds *int32 `json:"globalLeaseDurationSeconds,omitempty"`

	// GlobalLeaseTransitions is the total number of times the global leader has changed.
	GlobalLeaseTransitions *int32 `json:"globalLeaseTransitions,omitempty"`
}

// ConditionType is a string alias for defining specific condition types
// for the MultiClusterLease resource.
type ConditionType string

// These are the valid condition types for a MultiClusterLease.
const (
	// ConditionTypeBackendHealthy indicates whether the controller can successfully
	// communicate with and perform operations on the backend store (e.g., GCS).
	ConditionTypeBackendHealthy ConditionType = "BackendHealthy"

	// ConditionTypeLockAcquiredInBackend indicates whether a global lock
	// for this lease scope is currently held in the backend store by anyone.
	// This doesn't necessarily mean it is hold by this candidate.
	ConditionTypeLockAcquiredInBackend ConditionType = "LockAcquiredInBackend"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=multiclusterleases,scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Holder (Spec)",type="string",JSONPath=".spec.holderIdentity"
// +kubebuilder:printcolumn:name="Leader (Status)",type="string",JSONPath=".status.globalHolderIdentity"
// +kubebuilder:printcolumn:name="Renew Time (Status)",type="string",format="date-time",JSONPath=".status.globalRenewTime"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// MultiClusterLease is the Schema for the multiclusterleases API
type MultiClusterLease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultiClusterLeaseSpec   `json:"spec,omitempty"`
	Status MultiClusterLeaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MultiClusterLeaseList contains a list of MultiClusterLease
type MultiClusterLeaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultiClusterLease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MultiClusterLease{}, &MultiClusterLeaseList{})
}
