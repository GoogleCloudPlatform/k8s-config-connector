// Copyright 2025 Google LLC
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

// +generated:types
// krm.group: storage.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.storage.control.v2
// resource: StorageAnywhereCache:AnywhereCache

package v1beta1

// +kcc:proto=google.storage.control.v2.AnywhereCache
type AnywhereCache struct {
	// Immutable. The resource name of this AnywhereCache.
	//  Format:
	//  `projects/{project}/buckets/{bucket}/anywhereCaches/{anywhere_cache}`
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.name
	Name *string `json:"name,omitempty"`

	// Immutable. The zone in which the cache instance is running. For example,
	//  us-central1-a.
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.zone
	Zone *string `json:"zone,omitempty"`

	// Cache entry TTL (ranges between 1h to 7d). This is a cache-level config
	//  that defines how long a cache entry can live. Default ttl value (24hrs)
	//  is applied if not specified in the create request. TTL must be in whole
	//  seconds.
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.ttl
	Ttl *string `json:"ttl,omitempty"`

	// Cache admission policy. Valid policies includes:
	//  `admit-on-first-miss` and `admit-on-second-miss`. Defaults to
	//  `admit-on-first-miss`. Default value is applied if not specified in the
	//  create request.
	// +kcc:proto:field=google.storage.control.v2.AnywhereCache.admission_policy
	AdmissionPolicy *string `json:"admissionPolicy,omitempty"`
}

// +kcc:proto=google.storage.control.v2.AnywhereCache
type AnywhereCacheObservedState struct {
	// Output only. Cache state including RUNNING, CREATING, DISABLED and PAUSED.
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
