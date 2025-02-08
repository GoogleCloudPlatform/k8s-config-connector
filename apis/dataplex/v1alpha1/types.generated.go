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


// +kcc:proto=google.cloud.dataplex.v1.DataTaxonomy
type DataTaxonomy struct {

	// Optional. Description of the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataTaxonomy
type DataTaxonomyObservedState struct {
	// Output only. The relative resource name of the DataTaxonomy, of the form:
	//  projects/{project_number}/locations/{location_id}/dataTaxonomies/{data_taxonomy_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the dataTaxonomy. This
	//  ID will be different if the DataTaxonomy is deleted and re-created with the
	//  same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the DataTaxonomy was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the DataTaxonomy was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The number of attributes in the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.attribute_count
	AttributeCount *int32 `json:"attributeCount,omitempty"`

	// Output only. The number of classes in the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.class_count
	ClassCount *int32 `json:"classCount,omitempty"`
}
