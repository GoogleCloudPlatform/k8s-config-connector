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


// +kcc:proto=google.cloud.privatecatalog.v1beta1.Catalog
type Catalog struct {
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.Catalog
type CatalogObservedState struct {
	// Output only. The resource name of the target catalog, in the format of
	//  `catalogs/{catalog}`.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Catalog.name
	Name *string `json:"name,omitempty"`

	// Output only. The descriptive name of the catalog as it appears in UIs.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Catalog.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The description of the catalog.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Catalog.description
	Description *string `json:"description,omitempty"`

	// Output only. The time when the catalog was created.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Catalog.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the catalog was last updated.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Catalog.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
