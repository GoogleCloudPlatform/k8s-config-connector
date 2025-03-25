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

var PubSubSnapshotGVK = GroupVersion.WithKind("PubSubSnapshot")

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
}

// PubSubSnapshotSpec defines the desired state of PubSubSnapshot
// +kcc:proto=google.pubsub.v1.Snapshot
type PubSubSnapshotSpec struct {
	Parent `json:",inline"`

	// The PubSubSnapshot name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The name of the topic from which this snapshot is retaining
	//  messages.
	// +kcc:proto:field=google.pubsub.v1.Snapshot.topic
	Topic *string `json:"topic,omitempty"`

	// Optional. The snapshot is guaranteed to exist up until this time.
	//  A newly-created snapshot expires no later than 7 days from the time of its
	//  creation. Its exact lifetime is determined at creation by the existing
	//  backlog in the source subscription. Specifically, the lifetime of the
	//  snapshot is `7 days - (age of oldest unacked message in the subscription)`.
	//  For example, consider a subscription whose oldest unacked message is 3 days
	//  old. If a snapshot is created from this subscription, the snapshot -- which
	//  will always capture this 3-day-old backlog as long as the snapshot
	//  exists -- will expire in 4 days. The service will refuse to create a
	//  snapshot that would expire in less than 1 hour after creation.
	// +kcc:proto:field=google.pubsub.v1.Snapshot.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. See [Creating and managing labels]
	//  (https://cloud.google.com/pubsub/docs/labels).
	// +kcc:proto:field=google.pubsub.v1.Snapshot.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// PubSubSnapshotStatus defines the config connector machine state of PubSubSnapshot
type PubSubSnapshotStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the PubSubSnapshot resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *PubSubSnapshotObservedState `json:"observedState,omitempty"`
}

// PubSubSnapshotObservedState is the state of the PubSubSnapshot resource as most recently observed in GCP.
// +kcc:proto=google.pubsub.v1.Snapshot
type PubSubSnapshotObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsubsnapshot;gcppubsubsnapshots
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubSnapshot is the Schema for the PubSubSnapshot API
// +k8s:openapi-gen=true
type PubSubSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PubSubSnapshotSpec   `json:"spec,omitempty"`
	Status PubSubSnapshotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PubSubSnapshotList contains a list of PubSubSnapshot
type PubSubSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubSnapshot{}, &PubSubSnapshotList{})
}
