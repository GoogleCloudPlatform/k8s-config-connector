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


// +kcc:proto=google.cloud.retail.v2.Audience
type Audience struct {
	// The genders of the audience. Strongly encouraged to use the standard
	//  values: "male", "female", "unisex".
	//
	//  At most 5 values are allowed. Each value must be a UTF-8 encoded string
	//  with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error
	//  is returned.
	//
	//  Google Merchant Center property
	//  [gender](https://support.google.com/merchants/answer/6324479). Schema.org
	//  property
	//  [Product.audience.suggestedGender](https://schema.org/suggestedGender).
	// +kcc:proto:field=google.cloud.retail.v2.Audience.genders
	Genders []string `json:"genders,omitempty"`

	// The age groups of the audience. Strongly encouraged to use the standard
	//  values: "newborn" (up to 3 months old), "infant" (3–12 months old),
	//  "toddler" (1–5 years old), "kids" (5–13 years old), "adult" (typically
	//  teens or older).
	//
	//  At most 5 values are allowed. Each value must be a UTF-8 encoded string
	//  with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error
	//  is returned.
	//
	//  Google Merchant Center property
	//  [age_group](https://support.google.com/merchants/answer/6324463).
	//  Schema.org property
	//  [Product.audience.suggestedMinAge](https://schema.org/suggestedMinAge) and
	//  [Product.audience.suggestedMaxAge](https://schema.org/suggestedMaxAge).
	// +kcc:proto:field=google.cloud.retail.v2.Audience.age_groups
	AgeGroups []string `json:"ageGroups,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.ColorInfo
type ColorInfo struct {
	// The standard color families. Strongly recommended to use the following
	//  standard color groups: "Red", "Pink", "Orange", "Yellow", "Purple",
	//  "Green", "Cyan", "Blue", "Brown", "White", "Gray", "Black" and
	//  "Mixed". Normally it is expected to have only 1 color family. May consider
	//  using single "Mixed" instead of multiple values.
	//
	//  A maximum of 5 values are allowed. Each value must be a UTF-8 encoded
	//  string with a length limit of 128 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Google Merchant Center property
	//  [color](https://support.google.com/merchants/answer/6324487). Schema.org
	//  property [Product.color](https://schema.org/color).
	// +kcc:proto:field=google.cloud.retail.v2.ColorInfo.color_families
	ColorFamilies []string `json:"colorFamilies,omitempty"`

	// The color display names, which may be different from standard color family
	//  names, such as the color aliases used in the website frontend. Normally
	//  it is expected to have only 1 color. May consider using single "Mixed"
	//  instead of multiple values.
	//
	//  A maximum of 75 colors are allowed. Each value must be a UTF-8 encoded
	//  string with a length limit of 128 characters. Otherwise, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  Google Merchant Center property
	//  [color](https://support.google.com/merchants/answer/6324487). Schema.org
	//  property [Product.color](https://schema.org/color).
	// +kcc:proto:field=google.cloud.retail.v2.ColorInfo.colors
	Colors []string `json:"colors,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.CustomAttribute
type CustomAttribute struct {
	// The textual values of this custom attribute. For example, `["yellow",
	//  "green"]` when the key is "color".
	//
	//  Empty string is not allowed. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	//
	//  Exactly one of [text][google.cloud.retail.v2.CustomAttribute.text] or
	//  [numbers][google.cloud.retail.v2.CustomAttribute.numbers] should be set.
	//  Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.CustomAttribute.text
	Text []string `json:"text,omitempty"`

	// The numerical values of this custom attribute. For example, `[2.3, 15.4]`
	//  when the key is "lengths_cm".
	//
	//  Exactly one of [text][google.cloud.retail.v2.CustomAttribute.text] or
	//  [numbers][google.cloud.retail.v2.CustomAttribute.numbers] should be set.
	//  Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.CustomAttribute.numbers
	Numbers []float64 `json:"numbers,omitempty"`

	// This field is normally ignored unless
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2.AttributesConfig.attribute_config_level]
	//  of the [Catalog][google.cloud.retail.v2.Catalog] is set to the deprecated
	//  'PRODUCT_LEVEL_ATTRIBUTE_CONFIG' mode. For information about product-level
	//  attribute configuration, see [Configuration
	//  modes](https://cloud.google.com/retail/docs/attribute-config#config-modes).
	//  If true, custom attribute values are searchable by text queries in
	//  [SearchService.Search][google.cloud.retail.v2.SearchService.Search].
	//
	//  This field is ignored in a [UserEvent][google.cloud.retail.v2.UserEvent].
	//
	//  Only set if type [text][google.cloud.retail.v2.CustomAttribute.text] is
	//  set. Otherwise, a INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.CustomAttribute.searchable
	Searchable *bool `json:"searchable,omitempty"`

	// This field is normally ignored unless
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2.AttributesConfig.attribute_config_level]
	//  of the [Catalog][google.cloud.retail.v2.Catalog] is set to the deprecated
	//  'PRODUCT_LEVEL_ATTRIBUTE_CONFIG' mode. For information about product-level
	//  attribute configuration, see [Configuration
	//  modes](https://cloud.google.com/retail/docs/attribute-config#config-modes).
	//  If true, custom attribute values are indexed, so that they can be filtered,
	//  faceted or boosted in
	//  [SearchService.Search][google.cloud.retail.v2.SearchService.Search].
	//
	//  This field is ignored in a [UserEvent][google.cloud.retail.v2.UserEvent].
	//
	//  See [SearchRequest.filter][google.cloud.retail.v2.SearchRequest.filter],
	//  [SearchRequest.facet_specs][google.cloud.retail.v2.SearchRequest.facet_specs]
	//  and
	//  [SearchRequest.boost_spec][google.cloud.retail.v2.SearchRequest.boost_spec]
	//  for more details.
	// +kcc:proto:field=google.cloud.retail.v2.CustomAttribute.indexable
	Indexable *bool `json:"indexable,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.FulfillmentInfo
type FulfillmentInfo struct {
	// The fulfillment type, including commonly used types (such as pickup in
	//  store and same day delivery), and custom types. Customers have to map
	//  custom types to their display names before rendering UI.
	//
	//  Supported values:
	//
	//  * "pickup-in-store"
	//  * "ship-to-store"
	//  * "same-day-delivery"
	//  * "next-day-delivery"
	//  * "custom-type-1"
	//  * "custom-type-2"
	//  * "custom-type-3"
	//  * "custom-type-4"
	//  * "custom-type-5"
	//
	//  If this field is set to an invalid value other than these, an
	//  INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.FulfillmentInfo.type
	Type *string `json:"type,omitempty"`

	// The IDs for this [type][google.cloud.retail.v2.FulfillmentInfo.type], such
	//  as the store IDs for
	//  [FulfillmentInfo.type.pickup-in-store][google.cloud.retail.v2.FulfillmentInfo.type]
	//  or the region IDs for
	//  [FulfillmentInfo.type.same-day-delivery][google.cloud.retail.v2.FulfillmentInfo.type].
	//
	//  A maximum of 3000 values are allowed. Each value must be a string with a
	//  length limit of 30 characters, matching the pattern `[a-zA-Z0-9_-]+`, such
	//  as "store1" or "REGION-2". Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	// +kcc:proto:field=google.cloud.retail.v2.FulfillmentInfo.place_ids
	PlaceIds []string `json:"placeIds,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Image
type Image struct {
	// Required. URI of the image.
	//
	//  This field must be a valid UTF-8 encoded URI with a length limit of 5,000
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Google Merchant Center property
	//  [image_link](https://support.google.com/merchants/answer/6324350).
	//  Schema.org property [Product.image](https://schema.org/image).
	// +kcc:proto:field=google.cloud.retail.v2.Image.uri
	URI *string `json:"uri,omitempty"`

	// Height of the image in number of pixels.
	//
	//  This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	// +kcc:proto:field=google.cloud.retail.v2.Image.height
	Height *int32 `json:"height,omitempty"`

	// Width of the image in number of pixels.
	//
	//  This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	// +kcc:proto:field=google.cloud.retail.v2.Image.width
	Width *int32 `json:"width,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Interval
type Interval struct {
	// Inclusive lower bound.
	// +kcc:proto:field=google.cloud.retail.v2.Interval.minimum
	Minimum *float64 `json:"minimum,omitempty"`

	// Exclusive lower bound.
	// +kcc:proto:field=google.cloud.retail.v2.Interval.exclusive_minimum
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty"`

	// Inclusive upper bound.
	// +kcc:proto:field=google.cloud.retail.v2.Interval.maximum
	Maximum *float64 `json:"maximum,omitempty"`

	// Exclusive upper bound.
	// +kcc:proto:field=google.cloud.retail.v2.Interval.exclusive_maximum
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.LocalInventory
type LocalInventory struct {
	// The place ID for the current set of inventory information.
	// +kcc:proto:field=google.cloud.retail.v2.LocalInventory.place_id
	PlaceID *string `json:"placeID,omitempty"`

	// Product price and cost information.
	//
	//  Google Merchant Center property
	//  [price](https://support.google.com/merchants/answer/6324371).
	// +kcc:proto:field=google.cloud.retail.v2.LocalInventory.price_info
	PriceInfo *PriceInfo `json:"priceInfo,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Input only. Supported fulfillment types. Valid fulfillment type values
	//  include commonly used types (such as pickup in store and same day
	//  delivery), and custom types. Customers have to map custom types to their
	//  display names before rendering UI.
	//
	//  Supported values:
	//
	//  * "pickup-in-store"
	//  * "ship-to-store"
	//  * "same-day-delivery"
	//  * "next-day-delivery"
	//  * "custom-type-1"
	//  * "custom-type-2"
	//  * "custom-type-3"
	//  * "custom-type-4"
	//  * "custom-type-5"
	//
	//  If this field is set to an invalid value other than these, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  All the elements must be distinct. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	// +kcc:proto:field=google.cloud.retail.v2.LocalInventory.fulfillment_types
	FulfillmentTypes []string `json:"fulfillmentTypes,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.PriceInfo
type PriceInfo struct {
	// The 3-letter currency code defined in [ISO
	//  4217](https://www.iso.org/iso-4217-currency-codes.html).
	//
	//  If this field is an unrecognizable currency code, an INVALID_ARGUMENT
	//  error is returned.
	//
	//  The [Product.Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]
	//  [Product][google.cloud.retail.v2.Product]s with the same
	//  [Product.primary_product_id][google.cloud.retail.v2.Product.primary_product_id]
	//  must share the same
	//  [currency_code][google.cloud.retail.v2.PriceInfo.currency_code]. Otherwise,
	//  a FAILED_PRECONDITION error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// Price of the product.
	//
	//  Google Merchant Center property
	//  [price](https://support.google.com/merchants/answer/6324371). Schema.org
	//  property [Offer.price](https://schema.org/price).
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.price
	Price *float32 `json:"price,omitempty"`

	// Price of the product without any discount. If zero, by default set to be
	//  the [price][google.cloud.retail.v2.PriceInfo.price]. If set,
	//  [original_price][google.cloud.retail.v2.PriceInfo.original_price] should be
	//  greater than or equal to [price][google.cloud.retail.v2.PriceInfo.price],
	//  otherwise an INVALID_ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.original_price
	OriginalPrice *float32 `json:"originalPrice,omitempty"`

	// The costs associated with the sale of a particular product. Used for gross
	//  profit reporting.
	//
	//  * Profit = [price][google.cloud.retail.v2.PriceInfo.price] -
	//  [cost][google.cloud.retail.v2.PriceInfo.cost]
	//
	//  Google Merchant Center property
	//  [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.cost
	Cost *float32 `json:"cost,omitempty"`

	// The timestamp when the [price][google.cloud.retail.v2.PriceInfo.price]
	//  starts to be effective. This can be set as a future timestamp, and the
	//  [price][google.cloud.retail.v2.PriceInfo.price] is only used for search
	//  after
	//  [price_effective_time][google.cloud.retail.v2.PriceInfo.price_effective_time].
	//  If so, the
	//  [original_price][google.cloud.retail.v2.PriceInfo.original_price] must be
	//  set and [original_price][google.cloud.retail.v2.PriceInfo.original_price]
	//  is used before
	//  [price_effective_time][google.cloud.retail.v2.PriceInfo.price_effective_time].
	//
	//  Do not set if [price][google.cloud.retail.v2.PriceInfo.price] is always
	//  effective because it will cause additional latency during search.
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.price_effective_time
	PriceEffectiveTime *string `json:"priceEffectiveTime,omitempty"`

	// The timestamp when the [price][google.cloud.retail.v2.PriceInfo.price]
	//  stops to be effective. The [price][google.cloud.retail.v2.PriceInfo.price]
	//  is used for search before
	//  [price_expire_time][google.cloud.retail.v2.PriceInfo.price_expire_time]. If
	//  this field is set, the
	//  [original_price][google.cloud.retail.v2.PriceInfo.original_price] must be
	//  set and [original_price][google.cloud.retail.v2.PriceInfo.original_price]
	//  is used after
	//  [price_expire_time][google.cloud.retail.v2.PriceInfo.price_expire_time].
	//
	//  Do not set if [price][google.cloud.retail.v2.PriceInfo.price] is always
	//  effective because it will cause additional latency during search.
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.price_expire_time
	PriceExpireTime *string `json:"priceExpireTime,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.PriceInfo.PriceRange
type PriceInfo_PriceRange struct {
	// The inclusive
	//  [Product.pricing_info.price][google.cloud.retail.v2.PriceInfo.price]
	//  interval of all [variant][google.cloud.retail.v2.Product.Type.VARIANT]
	//  [Product][google.cloud.retail.v2.Product] having the same
	//  [Product.primary_product_id][google.cloud.retail.v2.Product.primary_product_id].
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.PriceRange.price
	Price *Interval `json:"price,omitempty"`

	// The inclusive
	//  [Product.pricing_info.original_price][google.cloud.retail.v2.PriceInfo.original_price]
	//  internal of all [variant][google.cloud.retail.v2.Product.Type.VARIANT]
	//  [Product][google.cloud.retail.v2.Product] having the same
	//  [Product.primary_product_id][google.cloud.retail.v2.Product.primary_product_id].
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.PriceRange.original_price
	OriginalPrice *Interval `json:"originalPrice,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Product
type Product struct {
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
	Ttl *string `json:"ttl,omitempty"`

	// Immutable. Full resource name of the product, such as
	//  `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`.
	// +kcc:proto:field=google.cloud.retail.v2.Product.name
	Name *string `json:"name,omitempty"`

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
	CollectionMemberIds []string `json:"collectionMemberIds,omitempty"`

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

	// TODO: unsupported map type with key string and value message


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

	// Indicates which fields in the [Product][google.cloud.retail.v2.Product]s
	//  are returned in [SearchResponse][google.cloud.retail.v2.SearchResponse].
	//
	//  Supported fields for all [type][google.cloud.retail.v2.Product.type]s:
	//
	//  * [audience][google.cloud.retail.v2.Product.audience]
	//  * [availability][google.cloud.retail.v2.Product.availability]
	//  * [brands][google.cloud.retail.v2.Product.brands]
	//  * [color_info][google.cloud.retail.v2.Product.color_info]
	//  * [conditions][google.cloud.retail.v2.Product.conditions]
	//  * [gtin][google.cloud.retail.v2.Product.gtin]
	//  * [materials][google.cloud.retail.v2.Product.materials]
	//  * [name][google.cloud.retail.v2.Product.name]
	//  * [patterns][google.cloud.retail.v2.Product.patterns]
	//  * [price_info][google.cloud.retail.v2.Product.price_info]
	//  * [rating][google.cloud.retail.v2.Product.rating]
	//  * [sizes][google.cloud.retail.v2.Product.sizes]
	//  * [title][google.cloud.retail.v2.Product.title]
	//  * [uri][google.cloud.retail.v2.Product.uri]
	//
	//  Supported fields only for
	//  [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY] and
	//  [Type.COLLECTION][google.cloud.retail.v2.Product.Type.COLLECTION]:
	//
	//  * [categories][google.cloud.retail.v2.Product.categories]
	//  * [description][google.cloud.retail.v2.Product.description]
	//  * [images][google.cloud.retail.v2.Product.images]
	//
	//  Supported fields only for
	//  [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]:
	//
	//  * Only the first image in [images][google.cloud.retail.v2.Product.images]
	//
	//  To mark [attributes][google.cloud.retail.v2.Product.attributes] as
	//  retrievable, include paths of the form "attributes.key" where "key" is the
	//  key of a custom attribute, as specified in
	//  [attributes][google.cloud.retail.v2.Product.attributes].
	//
	//  For [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY] and
	//  [Type.COLLECTION][google.cloud.retail.v2.Product.Type.COLLECTION], the
	//  following fields are always returned in
	//  [SearchResponse][google.cloud.retail.v2.SearchResponse] by default:
	//
	//  * [name][google.cloud.retail.v2.Product.name]
	//
	//  For [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT], the
	//  following fields are always returned in by default:
	//
	//  * [name][google.cloud.retail.v2.Product.name]
	//  * [color_info][google.cloud.retail.v2.Product.color_info]
	//
	//
	//  Note: Returning more fields in
	//  [SearchResponse][google.cloud.retail.v2.SearchResponse] can increase
	//  response payload size and serving latency.
	//
	//  This field is deprecated. Use the retrievable site-wide control instead.
	// +kcc:proto:field=google.cloud.retail.v2.Product.retrievable_fields
	RetrievableFields *FieldMask `json:"retrievableFields,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Promotion
type Promotion struct {
	// ID of the promotion. For example, "free gift".
	//
	//  The value must be a UTF-8 encoded string with a length limit of 128
	//  characters, and match the pattern: `[a-zA-Z][a-zA-Z0-9_]*`. For example,
	//  id0LikeThis or ID_1_LIKE_THIS. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	//
	//  Corresponds to Google Merchant Center property
	//  [promotion_id](https://support.google.com/merchants/answer/7050148).
	// +kcc:proto:field=google.cloud.retail.v2.Promotion.promotion_id
	PromotionID *string `json:"promotionID,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rating
type Rating struct {
	// The total number of ratings. This value is independent of the value of
	//  [rating_histogram][google.cloud.retail.v2.Rating.rating_histogram].
	//
	//  This value must be nonnegative. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	// +kcc:proto:field=google.cloud.retail.v2.Rating.rating_count
	RatingCount *int32 `json:"ratingCount,omitempty"`

	// The average rating of the [Product][google.cloud.retail.v2.Product].
	//
	//  The rating is scaled at 1-5. Otherwise, an INVALID_ARGUMENT error is
	//  returned.
	// +kcc:proto:field=google.cloud.retail.v2.Rating.average_rating
	AverageRating *float32 `json:"averageRating,omitempty"`

	// List of rating counts per rating value (index = rating - 1). The list is
	//  empty if there is no rating. If the list is non-empty, its size is
	//  always 5. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  For example, [41, 14, 13, 47, 303]. It means that the
	//  [Product][google.cloud.retail.v2.Product] got 41 ratings with 1 star, 14
	//  ratings with 2 star, and so on.
	// +kcc:proto:field=google.cloud.retail.v2.Rating.rating_histogram
	RatingHistogram []int32 `json:"ratingHistogram,omitempty"`
}

// +kcc:proto=google.protobuf.FieldMask
type FieldMask struct {
	// The set of field mask paths.
	// +kcc:proto:field=google.protobuf.FieldMask.paths
	Paths []string `json:"paths,omitempty"`
}

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.PriceInfo
type PriceInfoObservedState struct {
	// Output only. The price range of all the child
	//  [Product.Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]
	//  [Product][google.cloud.retail.v2.Product]s grouped together on the
	//  [Product.Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  [Product][google.cloud.retail.v2.Product]. Only populated for
	//  [Product.Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	//  [Product][google.cloud.retail.v2.Product]s.
	//
	//  Note: This field is OUTPUT_ONLY for
	//  [ProductService.GetProduct][google.cloud.retail.v2.ProductService.GetProduct].
	//  Do not set this field in API requests.
	// +kcc:proto:field=google.cloud.retail.v2.PriceInfo.price_range
	PriceRange *PriceInfo_PriceRange `json:"priceRange,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Product
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
	Variants []Product `json:"variants,omitempty"`

	// Output only. A list of local inventories specific to different places.
	//
	//  This field can be managed by
	//  [ProductService.AddLocalInventories][google.cloud.retail.v2.ProductService.AddLocalInventories]
	//  and
	//  [ProductService.RemoveLocalInventories][google.cloud.retail.v2.ProductService.RemoveLocalInventories]
	//  APIs if fine-grained, high-volume updates are necessary.
	// +kcc:proto:field=google.cloud.retail.v2.Product.local_inventories
	LocalInventories []LocalInventory `json:"localInventories,omitempty"`

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
	Variants []Product `json:"variants,omitempty"`

	// Output only. A list of local inventories specific to different places.
	//
	//  This field can be managed by
	//  [ProductService.AddLocalInventories][google.cloud.retail.v2.ProductService.AddLocalInventories]
	//  and
	//  [ProductService.RemoveLocalInventories][google.cloud.retail.v2.ProductService.RemoveLocalInventories]
	//  APIs if fine-grained, high-volume updates are necessary.
	// +kcc:proto:field=google.cloud.retail.v2.Product.local_inventories
	LocalInventories []LocalInventory `json:"localInventories,omitempty"`
}
