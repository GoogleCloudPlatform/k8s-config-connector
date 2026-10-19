// Copyright 2026 Google LLC
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

// +kcc:observedstate:proto=google.cloud.retail.v2.Product
type ProductObservedState struct {
	// Product price and cost information.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [price](https://support.google.com/merchants/answer/6324371).
	// +kcc:proto:field=google.cloud.retail.v2.Product.price_info
	PriceInfo *PriceInfoObservedState `json:"priceInfo,omitempty"`

	// Output only. Product variants grouped together on primary product which
	//  share similar product attributes. It's automatically grouped by
	//  [primary_product_id][google.cloud.retail.v2.Product.primary_product_id] for
	//  all the product variants. Only populated for
	//  [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  [Product][google.cloud.retail.v2.Product]s.
	//
	//  Note: This field is OUTPUT_ONLY for
	//  [ProductService.GetProduct][google.cloud.retail.v2.ProductService.GetProduct].
	//  Do not set this field in API requests.
	// +kcc:proto:field=google.cloud.retail.v2.Product.variants
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	Variants []ProductObservedState `json:"variants,omitempty"`

	// Output only. A list of local inventories specific to different places.
	//
	//  This field can be managed by
	//  [ProductService.AddLocalInventories][google.cloud.retail.v2.ProductService.AddLocalInventories]
	//  and
	//  [ProductService.RemoveLocalInventories][google.cloud.retail.v2.ProductService.RemoveLocalInventories]
	//  APIs if fine-grained, high-volume updates are necessary.
	// +kcc:proto:field=google.cloud.retail.v2.Product.local_inventories
	LocalInventories []LocalInventoryObservedState `json:"localInventories,omitempty"`
}
