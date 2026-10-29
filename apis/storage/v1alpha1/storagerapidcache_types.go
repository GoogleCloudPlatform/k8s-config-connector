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

// API sources for StorageRapidCache, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/storage/control/v2/storage_control.proto
// +kcc:source:service-docs=https://cloud.google.com/storage/docs/apis
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/storage/docs/apis/reference/rest/v2/projects.buckets.rapidCaches

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageRapidCacheGVK = GroupVersion.WithKind("StorageRapidCache")

// StorageRapidCacheSpec defines the desired state of StorageRapidCache
// +kcc:spec:proto=google.storage.control.v2.RapidCache
// +kcc:required-from-proto
type StorageRapidCacheSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The StorageRapidCache name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. The zone in which the cache instance is running. For example,
	//  us-central1-a.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.zone
	Zone *string `json:"zone,omitempty"`

	// Immutable. The type of Rapid Cache this represents. Valid values include:
	//  'rapid-cache' and 'rapid-cache-ultra'.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.cache_type
	CacheType *string `json:"cacheType,omitempty"`

	// Cache entry TTL (ranges between 1h to 7d). This is a cache-level config
	//  that defines how long a cache entry can live. Default ttl value (24hrs)
	//  is applied if not specified in the create request. TTL must be in whole
	//  seconds.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.ttl
	TTL *string `json:"ttl,omitempty"`

	// Cache admission policy. Valid policies includes:
	//  no_read_admission, admit-on-first-miss and admit-on-second-miss. Defaults
	//  to admit-on-first-miss for both AC and RCU. Default value is applied if not
	//  specified in the create request.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.admission_policy
	AdmissionPolicy *string `json:"admissionPolicy,omitempty"`
}

// StorageRapidCacheStatus defines the config connector machine state of StorageRapidCache
type StorageRapidCacheStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageRapidCache resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageRapidCacheObservedState `json:"observedState,omitempty"`
}

// StorageRapidCacheObservedState is the state of the StorageRapidCache resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.storage.control.v2.RapidCache
type StorageRapidCacheObservedState struct {
	// Output only. Cache state including running, creating, and disabled.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.state
	State *string `json:"state,omitempty"`

	// Output only. Time when Rapid cache instance is allocated.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when Rapid cache instance is last updated, including
	//  creation.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. True if there is an active update operation against this cache
	//  instance. Subsequential update requests will be rejected if this field is
	//  true. Output only.
	// +kcc:proto:field=google.storage.control.v2.RapidCache.pending_update
	PendingUpdate *bool `json:"pendingUpdate,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragerapidcache;gcpstoragerapidcaches
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageRapidCache is the Schema for the StorageRapidCache API
// +k8s:openapi-gen=true
type StorageRapidCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageRapidCacheSpec   `json:"spec,omitempty"`
	Status StorageRapidCacheStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageRapidCacheList contains a list of StorageRapidCache
type StorageRapidCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageRapidCache `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageRapidCache{}, &StorageRapidCacheList{})
}
