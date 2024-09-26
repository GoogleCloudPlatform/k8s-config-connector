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

package v1beta1

// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.DynamicGroupMetadata
type DynamicGroupMetadata struct {
	// Memberships will be the union of all queries. Only one entry with USER resource is currently supported. Customers can create up to 500 dynamic groups.
	Queries []DynamicGroupQuery `json:"queries,omitempty"`

	// Output only. Status of the dynamic group.
	Status *DynamicGroupStatus `json:"status,omitempty"`
}

// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.DynamicGroupQuery
type DynamicGroupQuery struct {
	// Query that determines the memberships of the dynamic group. Examples: All users with at least one `organizations.department` of engineering. `user.organizations.exists(org, org.department=='engineering')` All users with at least one location that has `area` of `foo` and `building_id` of `bar`. `user.locations.exists(loc, loc.area=='foo' && loc.building_id=='bar')` All users with any variation of the name John Doe (case-insensitive queries add `equalsIgnoreCase()` to the value being queried). `user.name.value.equalsIgnoreCase('jOhn DoE')`
	Query *string `json:"query,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.DynamicGroupStatus
type DynamicGroupStatus struct {
	// Status of the dynamic group.
	Status *string `json:"status,omitempty"`

	// The latest time at which the dynamic group is guaranteed to be in the given status. If status is `UP_TO_DATE`, the latest time at which the dynamic group was confirmed to be up-to-date. If status is `UPDATING_MEMBERSHIPS`, the time at which dynamic group was created.
	StatusTime *string `json:"statusTime,omitempty"`
}

// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.EntityKey
type EntityKey struct {
	// The ID of the entity. For Google-managed entities, the `id` must be the email address of an existing group or user. For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.
	ID *string `json:"id,omitempty"`

	// The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}`.
	Namespace *string `json:"namespace,omitempty"`
}

// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.PosixGroup
type PosixGroup struct {
	// GID of the POSIX group.
	Gid *uint64 `json:"gid,omitempty"`

	// Name of the POSIX group.
	Name *string `json:"name,omitempty"`

	// System identifier for which group name and gid apply to. If not specified it will default to empty value.
	SystemID *string `json:"systemID,omitempty"`
}
