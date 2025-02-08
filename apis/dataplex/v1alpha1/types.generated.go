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


// +kcc:proto=google.cloud.dataplex.v1.DataAccessSpec
type DataAccessSpec struct {
	// Optional. The format of strings follows the pattern followed by IAM in the
	//  bindings. user:{email}, serviceAccount:{email} group:{email}.
	//  The set of principals to be granted reader role on data
	//  stored within resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAccessSpec.readers
	Readers []string `json:"readers,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataAttribute
type DataAttribute struct {

	// Optional. Description of the DataAttribute.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the DataAttribute.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The ID of the parent DataAttribute resource, should belong to the
	//  same data taxonomy. Circular dependency in parent chain is not valid.
	//  Maximum depth of the hierarchy allowed is 4.
	//  [a -> b -> c -> d -> e, depth = 4]
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.parent_id
	ParentID *string `json:"parentID,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Specified when applied to a resource (eg: Cloud Storage bucket,
	//  BigQuery dataset, BigQuery table).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.resource_access_spec
	ResourceAccessSpec *ResourceAccessSpec `json:"resourceAccessSpec,omitempty"`

	// Optional. Specified when applied to data stored on the resource (eg: rows,
	//  columns in BigQuery Tables).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.data_access_spec
	DataAccessSpec *DataAccessSpec `json:"dataAccessSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.ResourceAccessSpec
type ResourceAccessSpec struct {
	// Optional. The format of strings follows the pattern followed by IAM in the
	//  bindings. user:{email}, serviceAccount:{email} group:{email}.
	//  The set of principals to be granted reader role on the resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.ResourceAccessSpec.readers
	Readers []string `json:"readers,omitempty"`

	// Optional. The set of principals to be granted writer role on the resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.ResourceAccessSpec.writers
	Writers []string `json:"writers,omitempty"`

	// Optional. The set of principals to be granted owner role on the resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.ResourceAccessSpec.owners
	Owners []string `json:"owners,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataAttribute
type DataAttributeObservedState struct {
	// Output only. The relative resource name of the dataAttribute, of the form:
	//  projects/{project_number}/locations/{location_id}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the DataAttribute.
	//  This ID will be different if the DataAttribute is deleted and re-created
	//  with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the DataAttribute was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the DataAttribute was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The number of child attributes present for this attribute.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.attribute_count
	AttributeCount *int32 `json:"attributeCount,omitempty"`
}
