// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for RetailProduct, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/retail/v2/product.proto
// +kcc:source:service-docs=https://cloud.google.com/recommendations
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/recommendations/docs/reference/rest/v2/projects.locations.catalogs.branches.products

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RetailProductGVK = GroupVersion.WithKind("RetailProduct")

// RetailProductSpec defines the desired state of RetailProduct
// +kcc:spec:proto=google.cloud.retail.v2.Product
// +kcc:required-from-proto
type RetailProductSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The RetailProduct name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Note that this field is applied in the following ways:
	//
	//  * If the [Product][google.cloud.retail.v2.Product] is already expired
	//  when it is uploaded, this product
	//    is not indexed for search.
	//
	//  * If the [Product][google.cloud.retail.v2.Product] is not expired when it
	//  is uploaded, only the
	//    [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]'s and
	//    [Type.COLLECTION][google.cloud.retail.v2.Product.Type.COLLECTION]'s
	//    expireTime is respected, and
	//    [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]'s
	//    expireTime is not used.
	//
	//  In general, we suggest the users to delete the stale
	//  products explicitly, instead of using this field to determine staleness.
	//
	//  [expire_time][google.cloud.retail.v2.Product.expire_time] must be later
	//  than [available_time][google.cloud.retail.v2.Product.available_time] and
	//  [publish_time][google.cloud.retail.v2.Product.publish_time], otherwise an
	//  INVALID_ARGUMENT error is thrown.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [expiration_date](https://support.google.com/merchants/answer/6324499).
	// +kcc:proto:field=google.cloud.retail.v2.Product.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL (time to live) of the product. Note that this is only
	//  applicable to [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  and [Type.COLLECTION][google.cloud.retail.v2.Product.Type.COLLECTION],
	//  and ignored for
	//  [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]. In general,
	//  we suggest the users to delete the stale products explicitly, instead of
	//  using this field to determine staleness.
	//
	//  If it is set, it must be a non-negative value, and
	//  [expire_time][google.cloud.retail.v2.Product.expire_time] is set as
	//  current timestamp plus [ttl][google.cloud.retail.v2.Product.ttl]. The
	//  derived [expire_time][google.cloud.retail.v2.Product.expire_time] is
	//  returned in the output and [ttl][google.cloud.retail.v2.Product.ttl] is
	//  left blank when retrieving the [Product][google.cloud.retail.v2.Product].
	//
	//  If it is set, the product is not available for
	//  [SearchService.Search][google.cloud.retail.v2.SearchService.Search] after
	//  current timestamp plus [ttl][google.cloud.retail.v2.Product.ttl].
	//  However, the product can still be retrieved by
	//  [ProductService.GetProduct][google.cloud.retail.v2.ProductService.GetProduct]
	//  and
	//  [ProductService.ListProducts][google.cloud.retail.v2.ProductService.ListProducts].
	// +kcc:proto:field=google.cloud.retail.v2.Product.ttl
	TTL *string `json:"ttl,omitempty"`

	// Immutable. [Product][google.cloud.retail.v2.Product] identifier, which is
	//  the final component of [name][google.cloud.retail.v2.Product.name]. For
	//  example, this field is "id_1", if
	//  [name][google.cloud.retail.v2.Product.name] is
	//  `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/id_1`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [id](https://support.google.com/merchants/answer/6324405). Schema.org
	//  property [Product.sku](https://schema.org/sku).
	// +kcc:proto:field=google.cloud.retail.v2.Product.id
	ID *string `json:"id,omitempty"`

	// Immutable. The type of the product. Default to
	//  [Catalog.product_level_config.ingestion_product_type][google.cloud.retail.v2.ProductLevelConfig.ingestion_product_type]
	//  if unset.
	// +kcc:proto:field=google.cloud.retail.v2.Product.type
	Type *string `json:"type,omitempty"`

	// Variant group identifier. Must be an
	//  [id][google.cloud.retail.v2.Product.id], with the same parent branch with
	//  this product. Otherwise, an error is thrown.
	//
	//  For [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  [Product][google.cloud.retail.v2.Product]s, this field can only be empty or
	//  set to the same value as [id][google.cloud.retail.v2.Product.id].
	//
	//  For VARIANT [Product][google.cloud.retail.v2.Product]s, this field cannot
	//  be empty. A maximum of 2,000 products are allowed to share the same
	//  [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  [Product][google.cloud.retail.v2.Product]. Otherwise, an INVALID_ARGUMENT
	//  error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [item_group_id](https://support.google.com/merchants/answer/6324507).
	//  Schema.org property
	//  [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID).
	// +kcc:proto:field=google.cloud.retail.v2.Product.primary_product_id
	PrimaryProductID *string `json:"primaryProductID,omitempty"`

	// The [id][google.cloud.retail.v2.Product.id] of the collection members when
	//  [type][google.cloud.retail.v2.Product.type] is
	//  [Type.COLLECTION][google.cloud.retail.v2.Product.Type.COLLECTION].
	//
	//  Non-existent product ids are allowed.
	//  The [type][google.cloud.retail.v2.Product.type] of the members must be
	//  either [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY] or
	//  [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT] otherwise an
	//  INVALID_ARGUMENT error is thrown. Should not set it for other types. A
	//  maximum of 1000 values are allowed. Otherwise, an INVALID_ARGUMENT error is
	//  return.
	// +kcc:proto:field=google.cloud.retail.v2.Product.collection_member_ids
	CollectionMemberIDs []string `json:"collectionMemberIDs,omitempty"`

	// The Global Trade Item Number (GTIN) of the product.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  This field must be a Unigram. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [gtin](https://support.google.com/merchants/answer/6324461).
	//  Schema.org property
	//  [Product.isbn](https://schema.org/isbn),
	//  [Product.gtin8](https://schema.org/gtin8),
	//  [Product.gtin12](https://schema.org/gtin12),
	//  [Product.gtin13](https://schema.org/gtin13), or
	//  [Product.gtin14](https://schema.org/gtin14).
	//
	//  If the value is not a valid GTIN, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.Product.gtin
	Gtin *string `json:"gtin,omitempty"`

	// Product categories. This field is repeated for supporting one product
	//  belonging to several parallel categories. Strongly recommended using the
	//  full path for better search / recommendation quality.
	//
	//
	//  To represent full path of category, use '>' sign to separate different
	//  hierarchies. If '>' is part of the category name, replace it with
	//  other character(s).
	//
	//  For example, if a shoes product belongs to both
	//  ["Shoes & Accessories" -> "Shoes"] and
	//  ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be
	//  represented as:
	//
	//       "categories": [
	//         "Shoes & Accessories > Shoes",
	//         "Sports & Fitness > Athletic Clothing > Shoes"
	//       ]
	//
	//  Must be set for [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  [Product][google.cloud.retail.v2.Product] otherwise an INVALID_ARGUMENT
	//  error is returned.
	//
	//  At most 250 values are allowed per
	//  [Product][google.cloud.retail.v2.Product] unless overridden through the
	//  Google Cloud console. Empty values are not allowed. Each value must be a
	//  UTF-8 encoded string with a length limit of 5,000 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [google_product_category][mc_google_product_category]. Schema.org property
	//  [Product.category] (https://schema.org/category).
	//
	//  [mc_google_product_category]:
	//  https://support.google.com/merchants/answer/6324436
	// +kcc:proto:field=google.cloud.retail.v2.Product.categories
	Categories []string `json:"categories,omitempty"`

	// Required. Product title.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1,000
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [title](https://support.google.com/merchants/answer/6324415). Schema.org
	//  property [Product.name](https://schema.org/name).
	// +kcc:proto:field=google.cloud.retail.v2.Product.title
	// +required
	Title *string `json:"title,omitempty"`

	// The brands of the product.
	//
	//  A maximum of 30 brands are allowed unless overridden through the Google
	//  Cloud console. Each
	//  brand must be a UTF-8 encoded string with a length limit of 1,000
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [brand](https://support.google.com/merchants/answer/6324351). Schema.org
	//  property [Product.brand](https://schema.org/brand).
	// +kcc:proto:field=google.cloud.retail.v2.Product.brands
	Brands []string `json:"brands,omitempty"`

	// Product description.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 5,000
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [description](https://support.google.com/merchants/answer/6324468).
	//  Schema.org property [Product.description](https://schema.org/description).
	// +kcc:proto:field=google.cloud.retail.v2.Product.description
	Description *string `json:"description,omitempty"`

	// Language of the title/description and other string attributes. Use language
	//  tags defined by [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	//
	//  For product prediction, this field is ignored and the model automatically
	//  detects the text language. The [Product][google.cloud.retail.v2.Product]
	//  can include text in different languages, but duplicating
	//  [Product][google.cloud.retail.v2.Product]s to provide text in multiple
	//  languages can result in degraded model performance.
	//
	//  For product search this field is in use. It defaults to "en-US" if unset.
	// +kcc:proto:field=google.cloud.retail.v2.Product.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Highly encouraged. Extra product attributes to be included. For example,
	//  for products, this could include the store name, vendor, style, color, etc.
	//  These are very strong signals for recommendation model, thus we highly
	//  recommend providing the attributes here.
	//
	//  Features that can take on one of a limited number of possible values. Two
	//  types of features can be set are:
	//
	//  Textual features. some examples would be the brand/maker of a product, or
	//  country of a customer. Numerical features. Some examples would be the
	//  height/weight of a product, or age of a customer.
	//
	//  For example: `{ "vendor": {"text": ["vendor123", "vendor456"]},
	//  "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]}
	//  }`.
	//
	//  This field needs to pass all below criteria, otherwise an INVALID_ARGUMENT
	//  error is returned:
	//
	//  * Max entries count: 200.
	//  * The key must be a UTF-8 encoded string with a length limit of 128
	//    characters.
	//  * For indexable attribute, the key must match the pattern:
	//    `[a-zA-Z0-9][a-zA-Z0-9_]*`. For example, `key0LikeThis` or
	//    `KEY_1_LIKE_THIS`.
	//  * For text attributes, at most 400 values are allowed. Empty values are not
	//    allowed. Each value must be a non-empty UTF-8 encoded string with a
	//    length limit of 256 characters.
	//  * For number attributes, at most 400 values are allowed.
	// +kcc:proto:field=google.cloud.retail.v2.Product.attributes
	Attributes map[string]CustomAttribute `json:"attributes,omitempty"`

	// Custom tags associated with the product.
	//
	//  At most 250 values are allowed per
	//  [Product][google.cloud.retail.v2.Product]. This value must be a UTF-8
	//  encoded string with a length limit of 1,000 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  This tag can be used for filtering recommendation results by passing the
	//  tag as part of the
	//  [PredictRequest.filter][google.cloud.retail.v2.PredictRequest.filter].
	//
	//  Corresponding properties: Google Merchant Center property
	//  [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	// +kcc:proto:field=google.cloud.retail.v2.Product.tags
	Tags []string `json:"tags,omitempty"`

	// Product price and cost information.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [price](https://support.google.com/merchants/answer/6324371).
	// +kcc:proto:field=google.cloud.retail.v2.Product.price_info
	PriceInfo *PriceInfo `json:"priceInfo,omitempty"`

	// The rating of this product.
	// +kcc:proto:field=google.cloud.retail.v2.Product.rating
	Rating *Rating `json:"rating,omitempty"`

	// The timestamp when this [Product][google.cloud.retail.v2.Product] becomes
	//  available for
	//  [SearchService.Search][google.cloud.retail.v2.SearchService.Search]. Note
	//  that this is only applicable to
	//  [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY] and
	//  [Type.COLLECTION][google.cloud.retail.v2.Product.Type.COLLECTION], and
	//  ignored for [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT].
	// +kcc:proto:field=google.cloud.retail.v2.Product.available_time
	AvailableTime *string `json:"availableTime,omitempty"`

	// The online availability of the [Product][google.cloud.retail.v2.Product].
	//  Default to
	//  [Availability.IN_STOCK][google.cloud.retail.v2.Product.Availability.IN_STOCK].
	//
	//  For primary products with variants set the availability of the primary as
	//  [Availability.OUT_OF_STOCK][google.cloud.retail.v2.Product.Availability.OUT_OF_STOCK]
	//  and set the true availability at the variant level. This way the primary
	//  product will be considered "in stock" as long as it has at least one
	//  variant in stock.
	//
	//  For primary products with no variants set the true availability at the
	//  primary level.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [availability](https://support.google.com/merchants/answer/6324448).
	//  Schema.org property [Offer.availability](https://schema.org/availability).
	// +kcc:proto:field=google.cloud.retail.v2.Product.availability
	Availability *string `json:"availability,omitempty"`

	// The available quantity of the item.
	// +kcc:proto:field=google.cloud.retail.v2.Product.available_quantity
	AvailableQuantity *Int32Value `json:"availableQuantity,omitempty"`

	// Fulfillment information, such as the store IDs for in-store pickup or
	//  region IDs for different shipping methods.
	//
	//  All the elements must have distinct
	//  [FulfillmentInfo.type][google.cloud.retail.v2.FulfillmentInfo.type].
	//  Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.Product.fulfillment_info
	FulfillmentInfo []FulfillmentInfo `json:"fulfillmentInfo,omitempty"`

	// Canonical URL directly linking to the product detail page.
	//
	//  It is strongly recommended to provide a valid uri for the product,
	//  otherwise the service performance could be significantly degraded.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 5,000
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [link](https://support.google.com/merchants/answer/6324416). Schema.org
	//  property [Offer.url](https://schema.org/url).
	// +kcc:proto:field=google.cloud.retail.v2.Product.uri
	URI *string `json:"uri,omitempty"`

	// Product images for the product. We highly recommend putting the main
	//  image first.
	//
	//  A maximum of 300 images are allowed.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [image_link](https://support.google.com/merchants/answer/6324350).
	//  Schema.org property [Product.image](https://schema.org/image).
	// +kcc:proto:field=google.cloud.retail.v2.Product.images
	Images []Image `json:"images,omitempty"`

	// The target group associated with a given audience (e.g. male, veterans,
	//  car owners, musicians, etc.) of the product.
	// +kcc:proto:field=google.cloud.retail.v2.Product.audience
	Audience *Audience `json:"audience,omitempty"`

	// The color of the product.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [color](https://support.google.com/merchants/answer/6324487). Schema.org
	//  property [Product.color](https://schema.org/color).
	// +kcc:proto:field=google.cloud.retail.v2.Product.color_info
	ColorInfo *ColorInfo `json:"colorInfo,omitempty"`

	// The size of the product. To represent different size systems or size types,
	//  consider using this format: [[[size_system:]size_type:]size_value].
	//
	//  For example, in "US:MENS:M", "US" represents size system; "MENS" represents
	//  size type; "M" represents size value. In "GIRLS:27", size system is empty;
	//  "GIRLS" represents size type; "27" represents size value. In "32 inches",
	//  both size system and size type are empty, while size value is "32 inches".
	//
	//  A maximum of 20 values are allowed per
	//  [Product][google.cloud.retail.v2.Product]. Each value must be a UTF-8
	//  encoded string with a length limit of 128 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [size](https://support.google.com/merchants/answer/6324492),
	//  [size_type](https://support.google.com/merchants/answer/6324497), and
	//  [size_system](https://support.google.com/merchants/answer/6324502).
	//  Schema.org property [Product.size](https://schema.org/size).
	// +kcc:proto:field=google.cloud.retail.v2.Product.sizes
	Sizes []string `json:"sizes,omitempty"`

	// The material of the product. For example, "leather", "wooden".
	//
	//  A maximum of 20 values are allowed. Each value must be a UTF-8 encoded
	//  string with a length limit of 200 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [material](https://support.google.com/merchants/answer/6324410). Schema.org
	//  property [Product.material](https://schema.org/material).
	// +kcc:proto:field=google.cloud.retail.v2.Product.materials
	Materials []string `json:"materials,omitempty"`

	// The pattern or graphic print of the product. For example, "striped", "polka
	//  dot", "paisley".
	//
	//  A maximum of 20 values are allowed per
	//  [Product][google.cloud.retail.v2.Product]. Each value must be a UTF-8
	//  encoded string with a length limit of 128 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [pattern](https://support.google.com/merchants/answer/6324483). Schema.org
	//  property [Product.pattern](https://schema.org/pattern).
	// +kcc:proto:field=google.cloud.retail.v2.Product.patterns
	Patterns []string `json:"patterns,omitempty"`

	// The condition of the product. Strongly encouraged to use the standard
	//  values: "new", "refurbished", "used".
	//
	//  A maximum of 1 value is allowed per
	//  [Product][google.cloud.retail.v2.Product]. Each value must be a UTF-8
	//  encoded string with a length limit of 128 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Corresponding properties: Google Merchant Center property
	//  [condition](https://support.google.com/merchants/answer/6324469).
	//  Schema.org property
	//  [Offer.itemCondition](https://schema.org/itemCondition).
	// +kcc:proto:field=google.cloud.retail.v2.Product.conditions
	Conditions []string `json:"conditions,omitempty"`

	// The promotions applied to the product. A maximum of 10 values are allowed
	//  per [Product][google.cloud.retail.v2.Product]. Only
	//  [Promotion.promotion_id][google.cloud.retail.v2.Promotion.promotion_id]
	//  will be used, other fields will be ignored if set.
	// +kcc:proto:field=google.cloud.retail.v2.Product.promotions
	Promotions []Promotion `json:"promotions,omitempty"`

	// The timestamp when the product is published by the retailer for the first
	//  time, which indicates the freshness of the products. Note that this field
	//  is different from
	//  [available_time][google.cloud.retail.v2.Product.available_time], given it
	//  purely describes product freshness regardless of when it is available on
	//  search and recommendation.
	// +kcc:proto:field=google.cloud.retail.v2.Product.publish_time
	PublishTime *string `json:"publishTime,omitempty"`
}

// RetailProductStatus defines the config connector machine state of RetailProduct
type RetailProductStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RetailProduct resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RetailProductObservedState `json:"observedState,omitempty"`
}

// RetailProductObservedState is the state of the RetailProduct resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.retail.v2.Product
type RetailProductObservedState struct {
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpretailproduct;gcpretailproducts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RetailProduct is the Schema for the RetailProduct API
// +k8s:openapi-gen=true
type RetailProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RetailProductSpec   `json:"spec,omitempty"`
	Status RetailProductStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RetailProductList contains a list of RetailProduct
type RetailProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RetailProduct `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RetailProduct{}, &RetailProductList{})
}
