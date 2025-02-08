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


// +kcc:proto=google.cloud.retail.v2beta.Catalog
type Catalog struct {
	// Required. Immutable. The fully qualified resource name of the catalog.
	// +kcc:proto:field=google.cloud.retail.v2beta.Catalog.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. The catalog display name.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2beta.Catalog.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The product level configuration.
	// +kcc:proto:field=google.cloud.retail.v2beta.Catalog.product_level_config
	ProductLevelConfig *ProductLevelConfig `json:"productLevelConfig,omitempty"`

	// The Merchant Center linking configuration.
	//  After a link is added, the data stream from Merchant Center to Cloud Retail
	//  will be enabled automatically. The requester must have access to the
	//  Merchant Center account in order to make changes to this field.
	// +kcc:proto:field=google.cloud.retail.v2beta.Catalog.merchant_center_linking_config
	MerchantCenterLinkingConfig *MerchantCenterLinkingConfig `json:"merchantCenterLinkingConfig,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.MerchantCenterFeedFilter
type MerchantCenterFeedFilter struct {
	// Merchant Center primary feed ID.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterFeedFilter.primary_feed_id
	PrimaryFeedID *int64 `json:"primaryFeedID,omitempty"`

	// Merchant Center primary feed name. The name is used for the display
	//  purposes only.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterFeedFilter.primary_feed_name
	PrimaryFeedName *string `json:"primaryFeedName,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.MerchantCenterLink
type MerchantCenterLink struct {
	// Required. The linked [Merchant Center account
	//  ID](https://developers.google.com/shopping-content/guides/accountstatuses).
	//  The account must be a standalone account or a sub-account of a MCA.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLink.merchant_center_account_id
	MerchantCenterAccountID *int64 `json:"merchantCenterAccountID,omitempty"`

	// The branch ID (e.g. 0/1/2) within this catalog that products from
	//  merchant_center_account_id are streamed to. When updating this field, an
	//  empty value will use the currently configured default branch. However,
	//  changing the default branch later on won't change the linked branch here.
	//
	//  A single branch ID can only have one linked Merchant Center account ID.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLink.branch_id
	BranchID *string `json:"branchID,omitempty"`

	// String representing the destination to import for, all if left empty.
	//  List of possible values is given in [Included
	//  destination](https://support.google.com/merchants/answer/7501026).
	//  List of allowed string values:
	//  "Shopping_ads", "Buy_on_google_listings", "Display_ads", "Local_inventory
	//  _ads", "Free_listings", "Free_local_listings"
	//  NOTE: The string values are case sensitive.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLink.destinations
	Destinations []string `json:"destinations,omitempty"`

	// Region code of offers to accept. 2-letter Uppercase ISO 3166-1 alpha-2
	//  code. List of values can be found
	//  [here](https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry)
	//  under the `region` tag. If left blank no region filtering will be
	//  performed.
	//
	//  Example value: `US`.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLink.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// Language of the title/description and other string attributes. Use language
	//  tags defined by [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	//  ISO 639-1.
	//
	//  This specifies the language of offers in Merchant Center that will be
	//  accepted. If  empty no language filtering will be performed.
	//
	//  Example value: `en`.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLink.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Criteria for the Merchant Center feeds to be ingested via the link.
	//  All offers will be ingested if the list is empty.
	//  Otherwise the offers will be ingested from selected feeds.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLink.feeds
	Feeds []MerchantCenterFeedFilter `json:"feeds,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.MerchantCenterLinkingConfig
type MerchantCenterLinkingConfig struct {
	// Links between Merchant Center accounts and branches.
	// +kcc:proto:field=google.cloud.retail.v2beta.MerchantCenterLinkingConfig.links
	Links []MerchantCenterLink `json:"links,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.ProductLevelConfig
type ProductLevelConfig struct {
	// The type of [Product][google.cloud.retail.v2beta.Product]s allowed to be
	//  ingested into the catalog. Acceptable values are:
	//
	//  * `primary` (default): You can ingest
	//  [Product][google.cloud.retail.v2beta.Product]s of all types. When
	//    ingesting a [Product][google.cloud.retail.v2beta.Product], its type will
	//    default to
	//    [Product.Type.PRIMARY][google.cloud.retail.v2beta.Product.Type.PRIMARY]
	//    if unset.
	//  * `variant` (incompatible with Retail Search): You can only
	//    ingest
	//    [Product.Type.VARIANT][google.cloud.retail.v2beta.Product.Type.VARIANT]
	//    [Product][google.cloud.retail.v2beta.Product]s. This means
	//    [Product.primary_product_id][google.cloud.retail.v2beta.Product.primary_product_id]
	//    cannot be empty.
	//
	//  If this field is set to an invalid value other than these, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  If this field is `variant` and
	//  [merchant_center_product_id_field][google.cloud.retail.v2beta.ProductLevelConfig.merchant_center_product_id_field]
	//  is `itemGroupId`, an INVALID_ARGUMENT error is returned.
	//
	//  See [Product
	//  levels](https://cloud.google.com/retail/docs/catalog#product-levels)
	//  for more details.
	// +kcc:proto:field=google.cloud.retail.v2beta.ProductLevelConfig.ingestion_product_type
	IngestionProductType *string `json:"ingestionProductType,omitempty"`

	// Which field of [Merchant Center
	//  Product](/bigquery-transfer/docs/merchant-center-products-schema) should be
	//  imported as [Product.id][google.cloud.retail.v2beta.Product.id]. Acceptable
	//  values are:
	//
	//  * `offerId` (default): Import `offerId` as the product ID.
	//  * `itemGroupId`: Import `itemGroupId` as the product ID. Notice that Retail
	//    API will choose one item from the ones with the same `itemGroupId`, and
	//    use it to represent the item group.
	//
	//  If this field is set to an invalid value other than these, an
	//  INVALID_ARGUMENT error is returned.
	//
	//  If this field is `itemGroupId` and
	//  [ingestion_product_type][google.cloud.retail.v2beta.ProductLevelConfig.ingestion_product_type]
	//  is `variant`, an INVALID_ARGUMENT error is returned.
	//
	//  See [Product
	//  levels](https://cloud.google.com/retail/docs/catalog#product-levels)
	//  for more details.
	// +kcc:proto:field=google.cloud.retail.v2beta.ProductLevelConfig.merchant_center_product_id_field
	MerchantCenterProductIDField *string `json:"merchantCenterProductIDField,omitempty"`
}
