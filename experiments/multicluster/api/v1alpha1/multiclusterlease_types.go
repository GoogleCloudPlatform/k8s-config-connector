// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
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
	// MultiClusterUID is a unique string identifying a multicluster election.
	MultiClusterUID string `json:"multiClusterUID,omitempty"`

	// Identity is the unique string identifying a lease holder across
	// all cluters in a multicluster election.
	Identity string `json:"identity,omitempty"`

	// The Google Cloud storage bucket to hold the lease.
	Bucket string `json:"bucket,omitempty"`

	// LeaseDurationSeconds is the duration that non-leader candidates will
	// wait to force acquire leadership. This is measured against time of
	// last observed ack.
	LeaseDurationSeconds *int32 `json:"leaseDurationSeconds,omitempty"`

	// RenewDeadlineSeconds is the duration that the acting master will retry
	// refreshing leadership before giving up.
	RenewDeadlineSeconds *int32 `json:"renewDeadlineSeconds,omitempty"`

	// RetryPeriodSeconds is the duration the LeaderElector clients should wait
	// between tries of actions.
	RetryPeriodSeconds *int32 `json:"retryPeriodSeconds,omitempty"`
}

// MultiClusterLeaseStatus defines the observed state of MultiClusterLease
type MultiClusterLeaseStatus struct {
	// TODO: add "IsLeader bool". Make sure the value is NOT stale when the cluster is disconnected.
	// Potentially KCC resource controllers could fetch this field to determine reconciliation mode.
	IsLeader         bool   `json:"isLeader"`
	LastObservedTime string `json:"lastObservedTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=multiclusterleases,scope=Namespaced
// +kubebuilder:subresource:status

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
