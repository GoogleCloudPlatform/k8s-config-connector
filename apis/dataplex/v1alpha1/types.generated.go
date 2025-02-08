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

package v1alpha1


// +kcc:proto=google.cloud.dataplex.v1.EntryType
type EntryType struct {

	// Optional. Description of the EntryType.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the EntryType.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. This checksum is computed by the service, and might be sent on
	//  update and delete requests to ensure the client has an up-to-date value
	//  before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Indicates the classes this Entry Type belongs to, for example,
	//  TABLE, DATABASE, MODEL.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.type_aliases
	TypeAliases []string `json:"typeAliases,omitempty"`

	// Optional. The platform that Entries of this type belongs to.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.platform
	Platform *string `json:"platform,omitempty"`

	// Optional. The system that Entries of this type belongs to. Examples include
	//  CloudSQL, MariaDB etc
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.system
	System *string `json:"system,omitempty"`

	// AspectInfo for the entry type.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.required_aspects
	RequiredAspects []EntryType_AspectInfo `json:"requiredAspects,omitempty"`

	// Immutable. Authorization defined for this type.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.authorization
	Authorization *EntryType_Authorization `json:"authorization,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType.AspectInfo
type EntryType_AspectInfo struct {
	// Required aspect type for the entry type.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.AspectInfo.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType.Authorization
type EntryType_Authorization struct {
	// Immutable. The IAM permission grantable on the Entry Group to allow
	//  access to instantiate Entries of Dataplex owned Entry Types, only
	//  settable for Dataplex owned Types.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.Authorization.alternate_use_permission
	AlternateUsePermission *string `json:"alternateUsePermission,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType
type EntryTypeObservedState struct {
	// Output only. The relative resource name of the EntryType, of the form:
	//  projects/{project_number}/locations/{location_id}/entryTypes/{entry_type_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the EntryType. This ID
	//  will be different if the EntryType is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the EntryType was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the EntryType was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
