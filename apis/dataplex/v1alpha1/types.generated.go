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


// +kcc:proto=google.cloud.dataplex.v1.DataAttributeBinding
type DataAttributeBinding struct {

	// Optional. Description of the DataAttributeBinding.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the DataAttributeBinding.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.labels
	Labels map[string]string `json:"labels,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	//  Etags must be used when calling the DeleteDataAttributeBinding and the
	//  UpdateDataAttributeBinding method.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Immutable. The resource name of the resource that is associated
	//  to attributes. Presently, only entity resource is supported in the form:
	//  projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id}
	//  Must belong in the same project and region as the attribute binding, and
	//  there can only exist one active binding for a resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.resource
	Resource *string `json:"resource,omitempty"`

	// Optional. List of attributes to be associated with the resource, provided
	//  in the form:
	//  projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.attributes
	Attributes []string `json:"attributes,omitempty"`

	// Optional. The list of paths for items within the associated resource (eg.
	//  columns and partitions within a table) along with attribute bindings.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.paths
	Paths []DataAttributeBinding_Path `json:"paths,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataAttributeBinding.Path
type DataAttributeBinding_Path struct {
	// Required. The name identifier of the path.
	//  Nested columns should be of the form: 'address.city'.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.Path.name
	Name *string `json:"name,omitempty"`

	// Optional. List of attributes to be associated with the path of the
	//  resource, provided in the form:
	//  projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.Path.attributes
	Attributes []string `json:"attributes,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataAttributeBinding
type DataAttributeBindingObservedState struct {
	// Output only. The relative resource name of the Data Attribute Binding, of
	//  the form:
	//  projects/{project_number}/locations/{location}/dataAttributeBindings/{data_attribute_binding_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the
	//  DataAttributeBinding. This ID will be different if the DataAttributeBinding
	//  is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the DataAttributeBinding was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the DataAttributeBinding was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttributeBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
