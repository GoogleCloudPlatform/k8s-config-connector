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


// +kcc:proto=google.cloud.dataplex.v1.Aspect
type Aspect struct {

	// Required. The content of the aspect, according to its aspect type schema.
	//  The maximum size of the field is 120KB (encoded as UTF-8).
	// +kcc:proto:field=google.cloud.dataplex.v1.Aspect.data
	Data map[string]string `json:"data,omitempty"`

	// Optional. Information related to the source system of the aspect.
	// +kcc:proto:field=google.cloud.dataplex.v1.Aspect.aspect_source
	AspectSource *AspectSource `json:"aspectSource,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectSource
type AspectSource struct {
	// The time the aspect was created in the source system.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectSource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The time the aspect was last updated in the source system.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectSource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The version of the data format used to produce this data. This field is
	//  used to indicated when the underlying data format changes (e.g., schema
	//  modifications, changes to the source URL format definition, etc).
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectSource.data_version
	DataVersion *string `json:"dataVersion,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Entry
type Entry struct {
	// Identifier. The relative resource name of the entry, in the format
	//  `projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}/entries/{entry_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. The relative resource name of the entry type that was
	//  used to create this entry, in the format
	//  `projects/{project_id_or_number}/locations/{location_id}/entryTypes/{entry_type_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.entry_type
	EntryType *string `json:"entryType,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. Immutable. The resource name of the parent entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.parent_entry
	ParentEntry *string `json:"parentEntry,omitempty"`

	// Optional. A name for the entry that can be referenced by an external
	//  system. For more information, see [Fully qualified
	//  names](https://cloud.google.com/data-catalog/docs/fully-qualified-names).
	//  The maximum size of the field is 4000 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.fully_qualified_name
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`

	// Optional. Information related to the source system of the data resource
	//  that is represented by the entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.entry_source
	EntrySource *EntrySource `json:"entrySource,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntrySource
type EntrySource struct {
	// The name of the resource in the source system.
	//  Maximum length is 4,000 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.resource
	Resource *string `json:"resource,omitempty"`

	// The name of the source system.
	//  Maximum length is 64 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.system
	System *string `json:"system,omitempty"`

	// The platform containing the source system.
	//  Maximum length is 64 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.platform
	Platform *string `json:"platform,omitempty"`

	// A user-friendly display name.
	//  Maximum length is 500 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description of the data resource.
	//  Maximum length is 2,000 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.description
	Description *string `json:"description,omitempty"`

	// User-defined labels.
	//  The maximum size of keys and values is 128 characters each.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The entries representing the ancestors of the data resource in
	//  the source system.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.ancestors
	Ancestors []EntrySource_Ancestor `json:"ancestors,omitempty"`

	// The time when the resource was created in the source system.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The time when the resource was last updated in the source system. If the
	//  entry exists in the system and its `EntrySource` has `update_time`
	//  populated, further updates to the `EntrySource` of the entry must provide
	//  incremental updates to its `update_time`.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntrySource.Ancestor
type EntrySource_Ancestor struct {
	// Optional. The name of the ancestor resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.Ancestor.name
	Name *string `json:"name,omitempty"`

	// Optional. The type of the ancestor resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.Ancestor.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Entry
type EntryObservedState struct {
	// Output only. The time when the entry was created in Dataplex.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the entry was last updated in Dataplex.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Information related to the source system of the data resource
	//  that is represented by the entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.entry_source
	EntrySource *EntrySourceObservedState `json:"entrySource,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntrySource
type EntrySourceObservedState struct {
	// Output only. Location of the resource in the source system. You can search
	//  the entry by this location. By default, this should match the location of
	//  the entry group containing this entry. A different value allows capturing
	//  the source location for data external to Google Cloud.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntrySource.location
	Location *string `json:"location,omitempty"`
}
