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

var StorageAnywhereCacheGVK = GroupVersion.WithKind("StorageAnywhereCache")

// StorageAnywhereCacheSpec defines the desired state of StorageAnywhereCache
// +kcc:spec:proto=google.storage.control.v2.AnywhereCache
type StorageAnywhereCacheSpec struct {
	// Immutable. The reference to bucket where cache needs to be created.
	// +required
	BucketRef *refs.StorageBucketRef `json:"bucketRef"`

	// Immutable. The zone in which the cache instance needs to be created. For example, us-central1-a.
	// +required
	Zone *string `json:"zone"`

	// The AnywhereCacheID generated via backend, It can be used by users to manage an existing cache.
	ResourceID *string `json:"resourceID,omitempty"`

	// The desired state of the cache. Possible values include "running", "disabled", and "paused".
	// If not specified, the default value is "running". This field controls the runtime behavior of the cache.
	// Please note that changes to the `desiredState` are prioritized over any other updates.
	// For instance, if both the `desiredState` and `ttl` are updated simultaneously, the state would be
	// updated first, followed by `ttl`.
	// +kubebuilder:default=running
	// +kubebuilder:validation:Enum=running;paused;disabled
	DesiredState *string `json:"desiredState,omitempty"`

	// Cache entry TTL (ranges between 1h to 7d). This is a cache-level config
	//  that defines how long a cache entry can live. Defaults to "86400s"
	//  TTL must be in whole seconds.
	// +kubebuilder:default="86400s"
	Ttl *string `json:"ttl,omitempty"`

	// Cache admission policy. Valid values includes:
	//  `admit-on-first-miss` and `admit-on-second-miss`. Defaults to
	//  `admit-on-first-miss`.
	// +kubebuilder:default=admit-on-first-miss
	// +kubebuilder:validation:Enum=admit-on-first-miss;admit-on-second-miss
	AdmissionPolicy *string `json:"admissionPolicy,omitempty"`
}

// StorageAnywhereCacheStatus defines the config connector machine state of StorageAnywhereCache
type StorageAnywhereCacheStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageAnywhereCache resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageAnywhereCacheObservedState `json:"observedState,omitempty"`
}

// StorageAnywhereCacheObservedState is the state of the StorageAnywhereCache resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.storage.control.v2.AnywhereCache
type StorageAnywhereCacheObservedState struct {
	// Output only. Cache state including "running", "creating", "disabled" and "paused".
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.state
	State *string `json:"state,omitempty"`

	// Output only. Time when Anywhere cache instance is allocated.
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when Anywhere cache instance is last updated, including
	//  creation.
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. True if there is an active update operation against this cache
	//  instance. Subsequential update requests will be rejected if this field is
	//  true. Output only.
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.pending_update
	PendingUpdate *bool `json:"pendingUpdate,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstorageanywherecache;gcpstorageanywherecaches
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageAnywhereCache is the Schema for the StorageAnywhereCache API
// +k8s:openapi-gen=true
type StorageAnywhereCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageAnywhereCacheSpec   `json:"spec,omitempty"`
	Status StorageAnywhereCacheStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageAnywhereCacheList contains a list of StorageAnywhereCache
type StorageAnywhereCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAnywhereCache `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageAnywhereCache{}, &StorageAnywhereCacheList{})
}
