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


// +kcc:proto=google.cloud.privatecatalog.v1beta1.Version
type Version struct {
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.Version
type VersionObservedState struct {
	// Output only. The resource name of the version, in the format
	//  `catalogs/{catalog}/products/{product}/versions/[a-z][-a-z0-9]*[a-z0-9]`.
	//
	//  A unique identifier for the version under a product.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Version.name
	Name *string `json:"name,omitempty"`

	// Output only. The user-supplied description of the version. Maximum of 256
	//  characters.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Version.description
	Description *string `json:"description,omitempty"`

	// Output only. The asset which has been validated and is ready to be
	//  provisioned. See
	//  [google.cloud.privatecatalogproducer.v1beta.Version.asset][] for details.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Version.asset
	Asset map[string]string `json:"asset,omitempty"`

	// Output only. The time when the version was created.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Version.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the version was last updated.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Version.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
