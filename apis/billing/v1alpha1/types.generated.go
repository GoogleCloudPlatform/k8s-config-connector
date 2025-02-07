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


// +kcc:proto=google.cloud.billing.v1.AggregationInfo
type AggregationInfo struct {
	// +kcc:proto:field=google.cloud.billing.v1.AggregationInfo.aggregation_level
	AggregationLevel *string `json:"aggregationLevel,omitempty"`

	// +kcc:proto:field=google.cloud.billing.v1.AggregationInfo.aggregation_interval
	AggregationInterval *string `json:"aggregationInterval,omitempty"`

	// The number of intervals to aggregate over.
	//  Example: If aggregation_level is "DAILY" and aggregation_count is 14,
	//  aggregation will be over 14 days.
	// +kcc:proto:field=google.cloud.billing.v1.AggregationInfo.aggregation_count
	AggregationCount *int32 `json:"aggregationCount,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.Category
type Category struct {
	// The display name of the service this SKU belongs to.
	// +kcc:proto:field=google.cloud.billing.v1.Category.service_display_name
	ServiceDisplayName *string `json:"serviceDisplayName,omitempty"`

	// The type of product the SKU refers to.
	//  Example: "Compute", "Storage", "Network", "ApplicationServices" etc.
	// +kcc:proto:field=google.cloud.billing.v1.Category.resource_family
	ResourceFamily *string `json:"resourceFamily,omitempty"`

	// A group classification for related SKUs.
	//  Example: "RAM", "GPU", "Prediction", "Ops", "GoogleEgress" etc.
	// +kcc:proto:field=google.cloud.billing.v1.Category.resource_group
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// Represents how the SKU is consumed.
	//  Example: "OnDemand", "Preemptible", "Commit1Mo", "Commit1Yr" etc.
	// +kcc:proto:field=google.cloud.billing.v1.Category.usage_type
	UsageType *string `json:"usageType,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.GeoTaxonomy
type GeoTaxonomy struct {
	// The type of Geo Taxonomy: GLOBAL, REGIONAL, or MULTI_REGIONAL.
	// +kcc:proto:field=google.cloud.billing.v1.GeoTaxonomy.type
	Type *string `json:"type,omitempty"`

	// The list of regions associated with a sku. Empty for Global skus, which are
	//  associated with all Google Cloud regions.
	// +kcc:proto:field=google.cloud.billing.v1.GeoTaxonomy.regions
	Regions []string `json:"regions,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.PricingExpression
type PricingExpression struct {
	// The short hand for unit of usage this pricing is specified in.
	//  Example: usage_unit of "GiBy" means that usage is specified in "Gibi Byte".
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.usage_unit
	UsageUnit *string `json:"usageUnit,omitempty"`

	// The recommended quantity of units for displaying pricing info. When
	//  displaying pricing info it is recommended to display:
	//  (unit_price * display_quantity) per display_quantity usage_unit.
	//  This field does not affect the pricing formula and is for display purposes
	//  only.
	//  Example: If the unit_price is "0.0001 USD", the usage_unit is "GB" and
	//  the display_quantity is "1000" then the recommended way of displaying the
	//  pricing info is "0.10 USD per 1000 GB"
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.display_quantity
	DisplayQuantity *float64 `json:"displayQuantity,omitempty"`

	// The list of tiered rates for this pricing. The total cost is computed by
	//  applying each of the tiered rates on usage. This repeated list is sorted
	//  by ascending order of start_usage_amount.
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.tiered_rates
	TieredRates []PricingExpression_TierRate `json:"tieredRates,omitempty"`

	// The unit of usage in human readable form.
	//  Example: "gibi byte".
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.usage_unit_description
	UsageUnitDescription *string `json:"usageUnitDescription,omitempty"`

	// The base unit for the SKU which is the unit used in usage exports.
	//  Example: "By"
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.base_unit
	BaseUnit *string `json:"baseUnit,omitempty"`

	// The base unit in human readable form.
	//  Example: "byte".
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.base_unit_description
	BaseUnitDescription *string `json:"baseUnitDescription,omitempty"`

	// Conversion factor for converting from price per usage_unit to price per
	//  base_unit, and start_usage_amount to start_usage_amount in base_unit.
	//  unit_price / base_unit_conversion_factor = price per base_unit.
	//  start_usage_amount * base_unit_conversion_factor = start_usage_amount in
	//  base_unit.
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.base_unit_conversion_factor
	BaseUnitConversionFactor *float64 `json:"baseUnitConversionFactor,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.PricingExpression.TierRate
type PricingExpression_TierRate struct {
	// Usage is priced at this rate only after this amount.
	//  Example: start_usage_amount of 10 indicates that the usage will be priced
	//  at the unit_price after the first 10 usage_units.
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.TierRate.start_usage_amount
	StartUsageAmount *float64 `json:"startUsageAmount,omitempty"`

	// The price per unit of usage.
	//  Example: unit_price of amount $10 indicates that each unit will cost $10.
	// +kcc:proto:field=google.cloud.billing.v1.PricingExpression.TierRate.unit_price
	UnitPrice *Money `json:"unitPrice,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.PricingInfo
type PricingInfo struct {
	// The timestamp from which this pricing was effective within the requested
	//  time range. This is guaranteed to be greater than or equal to the
	//  start_time field in the request and less than the end_time field in the
	//  request. If a time range was not specified in the request this field will
	//  be equivalent to a time within the last 12 hours, indicating the latest
	//  pricing info.
	// +kcc:proto:field=google.cloud.billing.v1.PricingInfo.effective_time
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// An optional human readable summary of the pricing information, has a
	//  maximum length of 256 characters.
	// +kcc:proto:field=google.cloud.billing.v1.PricingInfo.summary
	Summary *string `json:"summary,omitempty"`

	// Expresses the pricing formula. See `PricingExpression` for an example.
	// +kcc:proto:field=google.cloud.billing.v1.PricingInfo.pricing_expression
	PricingExpression *PricingExpression `json:"pricingExpression,omitempty"`

	// Aggregation Info. This can be left unspecified if the pricing expression
	//  doesn't require aggregation.
	// +kcc:proto:field=google.cloud.billing.v1.PricingInfo.aggregation_info
	AggregationInfo *AggregationInfo `json:"aggregationInfo,omitempty"`

	// Conversion rate used for currency conversion, from USD to the currency
	//  specified in the request. This includes any surcharge collected for billing
	//  in non USD currency. If a currency is not specified in the request this
	//  defaults to 1.0.
	//  Example: USD * currency_conversion_rate = JPY
	// +kcc:proto:field=google.cloud.billing.v1.PricingInfo.currency_conversion_rate
	CurrencyConversionRate *float64 `json:"currencyConversionRate,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.Sku
type Sku struct {
	// The resource name for the SKU.
	//  Example: "services/6F81-5844-456A/skus/D041-B8A1-6E0B"
	// +kcc:proto:field=google.cloud.billing.v1.Sku.name
	Name *string `json:"name,omitempty"`

	// The identifier for the SKU.
	//  Example: "D041-B8A1-6E0B"
	// +kcc:proto:field=google.cloud.billing.v1.Sku.sku_id
	SkuID *string `json:"skuID,omitempty"`

	// A human readable description of the SKU, has a maximum length of 256
	//  characters.
	// +kcc:proto:field=google.cloud.billing.v1.Sku.description
	Description *string `json:"description,omitempty"`

	// The category hierarchy of this SKU, purely for organizational purpose.
	// +kcc:proto:field=google.cloud.billing.v1.Sku.category
	Category *Category `json:"category,omitempty"`

	// List of service regions this SKU is offered at.
	//  Example: "asia-east1"
	//  Service regions can be found at https://cloud.google.com/about/locations/
	// +kcc:proto:field=google.cloud.billing.v1.Sku.service_regions
	ServiceRegions []string `json:"serviceRegions,omitempty"`

	// A timeline of pricing info for this SKU in chronological order.
	// +kcc:proto:field=google.cloud.billing.v1.Sku.pricing_info
	PricingInfo []PricingInfo `json:"pricingInfo,omitempty"`

	// Identifies the service provider.
	//  This is 'Google' for first party services in Google Cloud Platform.
	// +kcc:proto:field=google.cloud.billing.v1.Sku.service_provider_name
	ServiceProviderName *string `json:"serviceProviderName,omitempty"`

	// The geographic taxonomy for this sku.
	// +kcc:proto:field=google.cloud.billing.v1.Sku.geo_taxonomy
	GeoTaxonomy *GeoTaxonomy `json:"geoTaxonomy,omitempty"`
}

// +kcc:proto=google.type.Money
type Money struct {
	// The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`

	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}
