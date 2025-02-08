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


// +kcc:proto=google.cloud.vision.v1p3beta1.Product
type Product struct {
	// The resource name of the product.
	//
	//  Format is:
	//  `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID`.
	//
	//  This field is ignored when creating a product.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.name
	Name *string `json:"name,omitempty"`

	// The user-provided name for this Product. Must not be empty. Must be at most
	//  4096 characters long.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// User-provided metadata to be stored with this product. Must be at most 4096
	//  characters long.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.description
	Description *string `json:"description,omitempty"`

	// Immutable. The category for the product identified by the reference image. This should
	//  be either "homegoods-v2", "apparel-v2", or "toys-v2". The legacy categories
	//  "homegoods", "apparel", and "toys" are still supported, but these should
	//  not be used for new products.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.product_category
	ProductCategory *string `json:"productCategory,omitempty"`

	// Key-value pairs that can be attached to a product. At query time,
	//  constraints can be specified based on the product_labels.
	//
	//  Note that integer values can be provided as strings, e.g. "1199". Only
	//  strings with integer values can match a range-based restriction which is
	//  to be supported soon.
	//
	//  Multiple values can be assigned to the same key. One product may have up to
	//  100 product_labels.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.product_labels
	ProductLabels []Product_KeyValue `json:"productLabels,omitempty"`
}

// +kcc:proto=google.cloud.vision.v1p3beta1.Product.KeyValue
type Product_KeyValue struct {
	// The key of the label attached to the product. Cannot be empty and cannot
	//  exceed 128 bytes.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.KeyValue.key
	Key *string `json:"key,omitempty"`

	// The value of the label attached to the product. Cannot be empty and
	//  cannot exceed 128 bytes.
	// +kcc:proto:field=google.cloud.vision.v1p3beta1.Product.KeyValue.value
	Value *string `json:"value,omitempty"`
}
