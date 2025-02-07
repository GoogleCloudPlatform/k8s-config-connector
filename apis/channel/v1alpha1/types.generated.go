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


// +kcc:proto=google.cloud.channel.v1.Constraints
type Constraints struct {
	// Represents constraints required to purchase the Offer for a customer.
	// +kcc:proto:field=google.cloud.channel.v1.Constraints.customer_constraints
	CustomerConstraints *CustomerConstraints `json:"customerConstraints,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.CustomerConstraints
type CustomerConstraints struct {
	// Allowed geographical regions of the customer.
	// +kcc:proto:field=google.cloud.channel.v1.CustomerConstraints.allowed_regions
	AllowedRegions []string `json:"allowedRegions,omitempty"`

	// Allowed Customer Type.
	// +kcc:proto:field=google.cloud.channel.v1.CustomerConstraints.allowed_customer_types
	AllowedCustomerTypes []string `json:"allowedCustomerTypes,omitempty"`

	// Allowed Promotional Order Type. Present for Promotional offers.
	// +kcc:proto:field=google.cloud.channel.v1.CustomerConstraints.promotional_order_types
	PromotionalOrderTypes []string `json:"promotionalOrderTypes,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.MarketingInfo
type MarketingInfo struct {
	// Human readable name.
	// +kcc:proto:field=google.cloud.channel.v1.MarketingInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Human readable description. Description can contain HTML.
	// +kcc:proto:field=google.cloud.channel.v1.MarketingInfo.description
	Description *string `json:"description,omitempty"`

	// Default logo.
	// +kcc:proto:field=google.cloud.channel.v1.MarketingInfo.default_logo
	DefaultLogo *Media `json:"defaultLogo,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Media
type Media struct {
	// Title of the media.
	// +kcc:proto:field=google.cloud.channel.v1.Media.title
	Title *string `json:"title,omitempty"`

	// URL of the media.
	// +kcc:proto:field=google.cloud.channel.v1.Media.content
	Content *string `json:"content,omitempty"`

	// Type of the media.
	// +kcc:proto:field=google.cloud.channel.v1.Media.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Offer
type Offer struct {
	// Resource Name of the Offer.
	//  Format: accounts/{account_id}/offers/{offer_id}
	// +kcc:proto:field=google.cloud.channel.v1.Offer.name
	Name *string `json:"name,omitempty"`

	// Marketing information for the Offer.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.marketing_info
	MarketingInfo *MarketingInfo `json:"marketingInfo,omitempty"`

	// SKU the offer is associated with.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.sku
	Sku *Sku `json:"sku,omitempty"`

	// Describes the payment plan for the Offer.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.plan
	Plan *Plan `json:"plan,omitempty"`

	// Constraints on transacting the Offer.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.constraints
	Constraints *Constraints `json:"constraints,omitempty"`

	// Price for each monetizable resource type.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.price_by_resources
	PriceByResources []PriceByResource `json:"priceByResources,omitempty"`

	// Start of the Offer validity time.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Parameters required to use current Offer to purchase.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.parameter_definitions
	ParameterDefinitions []ParameterDefinition `json:"parameterDefinitions,omitempty"`

	// The deal code of the offer to get a special promotion or discount.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.deal_code
	DealCode *string `json:"dealCode,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ParameterDefinition
type ParameterDefinition struct {
	// Name of the parameter.
	// +kcc:proto:field=google.cloud.channel.v1.ParameterDefinition.name
	Name *string `json:"name,omitempty"`

	// Data type of the parameter. Minimal value, Maximum value and allowed values
	//  will use specified data type here.
	// +kcc:proto:field=google.cloud.channel.v1.ParameterDefinition.parameter_type
	ParameterType *string `json:"parameterType,omitempty"`

	// Minimal value of the parameter, if applicable. Inclusive. For example,
	//  minimal commitment when purchasing Anthos is 0.01.
	//  Applicable to INT64 and DOUBLE parameter types.
	// +kcc:proto:field=google.cloud.channel.v1.ParameterDefinition.min_value
	MinValue *Value `json:"minValue,omitempty"`

	// Maximum value of the parameter, if applicable. Inclusive. For example,
	//  maximum seats when purchasing Google Workspace Business Standard.
	//  Applicable to INT64 and DOUBLE parameter types.
	// +kcc:proto:field=google.cloud.channel.v1.ParameterDefinition.max_value
	MaxValue *Value `json:"maxValue,omitempty"`

	// If not empty, parameter values must be drawn from this list.
	//  For example, [us-west1, us-west2, ...]
	//  Applicable to STRING parameter type.
	// +kcc:proto:field=google.cloud.channel.v1.ParameterDefinition.allowed_values
	AllowedValues []Value `json:"allowedValues,omitempty"`

	// If set to true, parameter is optional to purchase this Offer.
	// +kcc:proto:field=google.cloud.channel.v1.ParameterDefinition.optional
	Optional *bool `json:"optional,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Period
type Period struct {
	// Total duration of Period Type defined.
	// +kcc:proto:field=google.cloud.channel.v1.Period.duration
	Duration *int32 `json:"duration,omitempty"`

	// Period Type.
	// +kcc:proto:field=google.cloud.channel.v1.Period.period_type
	PeriodType *string `json:"periodType,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Plan
type Plan struct {
	// Describes how a reseller will be billed.
	// +kcc:proto:field=google.cloud.channel.v1.Plan.payment_plan
	PaymentPlan *string `json:"paymentPlan,omitempty"`

	// Specifies when the payment needs to happen.
	// +kcc:proto:field=google.cloud.channel.v1.Plan.payment_type
	PaymentType *string `json:"paymentType,omitempty"`

	// Describes how frequently the reseller will be billed, such as
	//  once per month.
	// +kcc:proto:field=google.cloud.channel.v1.Plan.payment_cycle
	PaymentCycle *Period `json:"paymentCycle,omitempty"`

	// Present for Offers with a trial period.
	//  For trial-only Offers, a paid service needs to start before the trial
	//  period ends for continued service.
	//  For Regular Offers with a trial period, the regular pricing goes into
	//  effect when trial period ends, or if paid service is started before the end
	//  of the trial period.
	// +kcc:proto:field=google.cloud.channel.v1.Plan.trial_period
	TrialPeriod *Period `json:"trialPeriod,omitempty"`

	// Reseller Billing account to charge after an offer transaction.
	//  Only present for Google Cloud offers.
	// +kcc:proto:field=google.cloud.channel.v1.Plan.billing_account
	BillingAccount *string `json:"billingAccount,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Price
type Price struct {
	// Base price.
	// +kcc:proto:field=google.cloud.channel.v1.Price.base_price
	BasePrice *Money `json:"basePrice,omitempty"`

	// Discount percentage, represented as decimal.
	//  For example, a 20% discount will be represent as 0.2.
	// +kcc:proto:field=google.cloud.channel.v1.Price.discount
	Discount *float64 `json:"discount,omitempty"`

	// Effective Price after applying the discounts.
	// +kcc:proto:field=google.cloud.channel.v1.Price.effective_price
	EffectivePrice *Money `json:"effectivePrice,omitempty"`

	// Link to external price list, such as link to Google Voice rate card.
	// +kcc:proto:field=google.cloud.channel.v1.Price.external_price_uri
	ExternalPriceURI *string `json:"externalPriceURI,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.PriceByResource
type PriceByResource struct {
	// Resource Type. Example: SEAT
	// +kcc:proto:field=google.cloud.channel.v1.PriceByResource.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Price of the Offer. Present if there are no price phases.
	// +kcc:proto:field=google.cloud.channel.v1.PriceByResource.price
	Price *Price `json:"price,omitempty"`

	// Specifies the price by time range.
	// +kcc:proto:field=google.cloud.channel.v1.PriceByResource.price_phases
	PricePhases []PricePhase `json:"pricePhases,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.PricePhase
type PricePhase struct {
	// Defines the phase period type.
	// +kcc:proto:field=google.cloud.channel.v1.PricePhase.period_type
	PeriodType *string `json:"periodType,omitempty"`

	// Defines first period for the phase.
	// +kcc:proto:field=google.cloud.channel.v1.PricePhase.first_period
	FirstPeriod *int32 `json:"firstPeriod,omitempty"`

	// Defines first period for the phase.
	// +kcc:proto:field=google.cloud.channel.v1.PricePhase.last_period
	LastPeriod *int32 `json:"lastPeriod,omitempty"`

	// Price of the phase. Present if there are no price tiers.
	// +kcc:proto:field=google.cloud.channel.v1.PricePhase.price
	Price *Price `json:"price,omitempty"`

	// Price by the resource tiers.
	// +kcc:proto:field=google.cloud.channel.v1.PricePhase.price_tiers
	PriceTiers []PriceTier `json:"priceTiers,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.PriceTier
type PriceTier struct {
	// First resource for which the tier price applies.
	// +kcc:proto:field=google.cloud.channel.v1.PriceTier.first_resource
	FirstResource *int32 `json:"firstResource,omitempty"`

	// Last resource for which the tier price applies.
	// +kcc:proto:field=google.cloud.channel.v1.PriceTier.last_resource
	LastResource *int32 `json:"lastResource,omitempty"`

	// Price of the tier.
	// +kcc:proto:field=google.cloud.channel.v1.PriceTier.price
	Price *Price `json:"price,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Product
type Product struct {
	// Resource Name of the Product.
	//  Format: products/{product_id}
	// +kcc:proto:field=google.cloud.channel.v1.Product.name
	Name *string `json:"name,omitempty"`

	// Marketing information for the product.
	// +kcc:proto:field=google.cloud.channel.v1.Product.marketing_info
	MarketingInfo *MarketingInfo `json:"marketingInfo,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Sku
type Sku struct {
	// Resource Name of the SKU.
	//  Format: products/{product_id}/skus/{sku_id}
	// +kcc:proto:field=google.cloud.channel.v1.Sku.name
	Name *string `json:"name,omitempty"`

	// Marketing information for the SKU.
	// +kcc:proto:field=google.cloud.channel.v1.Sku.marketing_info
	MarketingInfo *MarketingInfo `json:"marketingInfo,omitempty"`

	// Product the SKU is associated with.
	// +kcc:proto:field=google.cloud.channel.v1.Sku.product
	Product *Product `json:"product,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Value
type Value struct {
	// Represents an int64 value.
	// +kcc:proto:field=google.cloud.channel.v1.Value.int64_value
	Int64Value *int64 `json:"int64Value,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.cloud.channel.v1.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.cloud.channel.v1.Value.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// Represents an 'Any' proto value.
	// +kcc:proto:field=google.cloud.channel.v1.Value.proto_value
	ProtoValue *Any `json:"protoValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.cloud.channel.v1.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
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

// +kcc:proto=google.cloud.channel.v1.Offer
type OfferObservedState struct {
	// Output only. End of the Offer validity time.
	// +kcc:proto:field=google.cloud.channel.v1.Offer.end_time
	EndTime *string `json:"endTime,omitempty"`
}
