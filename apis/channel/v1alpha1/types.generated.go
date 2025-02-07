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


// +kcc:proto=google.cloud.channel.v1.ChannelPartnerRepricingConfig
type ChannelPartnerRepricingConfig struct {

	// Required. The configuration for bill modifications made by a reseller
	//  before sending it to ChannelPartner.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerRepricingConfig.repricing_config
	RepricingConfig *RepricingConfig `json:"repricingConfig,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ConditionalOverride
type ConditionalOverride struct {
	// Required. Information about the applied override's adjustment.
	// +kcc:proto:field=google.cloud.channel.v1.ConditionalOverride.adjustment
	Adjustment *RepricingAdjustment `json:"adjustment,omitempty"`

	// Required. The [RebillingBasis][google.cloud.channel.v1.RebillingBasis] to
	//  use for the applied override. Shows the relative cost based on your
	//  repricing costs.
	// +kcc:proto:field=google.cloud.channel.v1.ConditionalOverride.rebilling_basis
	RebillingBasis *string `json:"rebillingBasis,omitempty"`

	// Required. Specifies the condition which, if met, will apply the override.
	// +kcc:proto:field=google.cloud.channel.v1.ConditionalOverride.repricing_condition
	RepricingCondition *RepricingCondition `json:"repricingCondition,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.PercentageAdjustment
type PercentageAdjustment struct {
	// The percentage of the bill to adjust.
	//  For example:
	//  Mark down by 1% => "-1.00"
	//  Mark up by 1%   => "1.00"
	//  Pass-Through    => "0.00"
	// +kcc:proto:field=google.cloud.channel.v1.PercentageAdjustment.percentage
	Percentage *Decimal `json:"percentage,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.RepricingAdjustment
type RepricingAdjustment struct {
	// Flat markup or markdown on an entire bill.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingAdjustment.percentage_adjustment
	PercentageAdjustment *PercentageAdjustment `json:"percentageAdjustment,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.RepricingCondition
type RepricingCondition struct {
	// SKU Group condition for override.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingCondition.sku_group_condition
	SkuGroupCondition *SkuGroupCondition `json:"skuGroupCondition,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.RepricingConfig
type RepricingConfig struct {
	// Applies the repricing configuration at the entitlement level.
	//
	//  Note: If a
	//  [ChannelPartnerRepricingConfig][google.cloud.channel.v1.ChannelPartnerRepricingConfig]
	//  using
	//  [RepricingConfig.EntitlementGranularity][google.cloud.channel.v1.RepricingConfig.EntitlementGranularity]
	//  becomes effective, then no existing or future
	//  [RepricingConfig.ChannelPartnerGranularity][google.cloud.channel.v1.RepricingConfig.ChannelPartnerGranularity]
	//  will apply to the
	//  [RepricingConfig.EntitlementGranularity.entitlement][google.cloud.channel.v1.RepricingConfig.EntitlementGranularity.entitlement].
	//  This is the recommended value for both
	//  [CustomerRepricingConfig][google.cloud.channel.v1.CustomerRepricingConfig]
	//  and
	//  [ChannelPartnerRepricingConfig][google.cloud.channel.v1.ChannelPartnerRepricingConfig].
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.entitlement_granularity
	EntitlementGranularity *RepricingConfig_EntitlementGranularity `json:"entitlementGranularity,omitempty"`

	// Applies the repricing configuration at the channel partner level.
	//  Only
	//  [ChannelPartnerRepricingConfig][google.cloud.channel.v1.ChannelPartnerRepricingConfig]
	//  supports this value. Deprecated: This is no longer supported. Use
	//  [RepricingConfig.entitlement_granularity][google.cloud.channel.v1.RepricingConfig.entitlement_granularity]
	//  instead.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.channel_partner_granularity
	ChannelPartnerGranularity *RepricingConfig_ChannelPartnerGranularity `json:"channelPartnerGranularity,omitempty"`

	// Required. The YearMonth when these adjustments activate. The Day field
	//  needs to be "0" since we only accept YearMonth repricing boundaries.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.effective_invoice_month
	EffectiveInvoiceMonth *Date `json:"effectiveInvoiceMonth,omitempty"`

	// Required. Information about the adjustment.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.adjustment
	Adjustment *RepricingAdjustment `json:"adjustment,omitempty"`

	// Required. The [RebillingBasis][google.cloud.channel.v1.RebillingBasis] to
	//  use for this bill. Specifies the relative cost based on repricing costs you
	//  will apply.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.rebilling_basis
	RebillingBasis *string `json:"rebillingBasis,omitempty"`

	// The conditional overrides to apply for this configuration. If you list
	//  multiple overrides, only the first valid override is used.  If you don't
	//  list any overrides, the API uses the normal adjustment and rebilling basis.
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.conditional_overrides
	ConditionalOverrides []ConditionalOverride `json:"conditionalOverrides,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.RepricingConfig.ChannelPartnerGranularity
type RepricingConfig_ChannelPartnerGranularity struct {
}

// +kcc:proto=google.cloud.channel.v1.RepricingConfig.EntitlementGranularity
type RepricingConfig_EntitlementGranularity struct {
	// Resource name of the entitlement.
	//  Format:
	//  accounts/{account_id}/customers/{customer_id}/entitlements/{entitlement_id}
	// +kcc:proto:field=google.cloud.channel.v1.RepricingConfig.EntitlementGranularity.entitlement
	Entitlement *string `json:"entitlement,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.SkuGroupCondition
type SkuGroupCondition struct {
	// Specifies a SKU group (https://cloud.google.com/skus/sku-groups).
	//  Resource name of SKU group. Format:
	//  accounts/{account}/skuGroups/{sku_group}.
	//  Example:
	//  "accounts/C01234/skuGroups/3d50fd57-3157-4577-a5a9-a219b8490041".
	// +kcc:proto:field=google.cloud.channel.v1.SkuGroupCondition.sku_group
	SkuGroup *string `json:"skuGroup,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.type.Decimal
type Decimal struct {
	// The decimal value, as a string.
	//
	//  The string representation consists of an optional sign, `+` (`U+002B`)
	//  or `-` (`U+002D`), followed by a sequence of zero or more decimal digits
	//  ("the integer"), optionally followed by a fraction, optionally followed
	//  by an exponent.
	//
	//  The fraction consists of a decimal point followed by zero or more decimal
	//  digits. The string must contain at least one digit in either the integer
	//  or the fraction. The number formed by the sign, the integer and the
	//  fraction is referred to as the significand.
	//
	//  The exponent consists of the character `e` (`U+0065`) or `E` (`U+0045`)
	//  followed by one or more decimal digits.
	//
	//  Services **should** normalize decimal values before storing them by:
	//
	//    - Removing an explicitly-provided `+` sign (`+2.5` -> `2.5`).
	//    - Replacing a zero-length integer value with `0` (`.5` -> `0.5`).
	//    - Coercing the exponent character to lower-case (`2.5E8` -> `2.5e8`).
	//    - Removing an explicitly-provided zero exponent (`2.5e0` -> `2.5`).
	//
	//  Services **may** perform additional normalization based on its own needs
	//  and the internal decimal implementation selected, such as shifting the
	//  decimal point and exponent value together (example: `2.5e-1` <-> `0.25`).
	//  Additionally, services **may** preserve trailing zeroes in the fraction
	//  to indicate increased precision, but are not required to do so.
	//
	//  Note that only the `.` character is supported to divide the integer
	//  and the fraction; `,` **should not** be supported regardless of locale.
	//  Additionally, thousand separators **should not** be supported. If a
	//  service does support them, values **must** be normalized.
	//
	//  The ENBF grammar is:
	//
	//      DecimalString =
	//        [Sign] Significand [Exponent];
	//
	//      Sign = '+' | '-';
	//
	//      Significand =
	//        Digits ['.'] [Digits] | [Digits] '.' Digits;
	//
	//      Exponent = ('e' | 'E') [Sign] Digits;
	//
	//      Digits = { '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' };
	//
	//  Services **should** clearly document the range of supported values, the
	//  maximum supported precision (total number of digits), and, if applicable,
	//  the scale (number of digits after the decimal point), as well as how it
	//  behaves when receiving out-of-bounds values.
	//
	//  Services **may** choose to accept values passed as input even when the
	//  value has a higher precision or scale than the service supports, and
	//  **should** round the value to fit the supported scale. Alternatively, the
	//  service **may** error with `400 Bad Request` (`INVALID_ARGUMENT` in gRPC)
	//  if precision would be lost.
	//
	//  Services **should** error with `400 Bad Request` (`INVALID_ARGUMENT` in
	//  gRPC) if the service receives a value outside of the supported range.
	// +kcc:proto:field=google.type.Decimal.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ChannelPartnerRepricingConfig
type ChannelPartnerRepricingConfigObservedState struct {
	// Output only. Resource name of the ChannelPartnerRepricingConfig.
	//  Format:
	//  accounts/{account_id}/channelPartnerLinks/{channel_partner_id}/channelPartnerRepricingConfigs/{id}.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerRepricingConfig.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp of an update to the repricing rule. If `update_time`
	//  is after
	//  [RepricingConfig.effective_invoice_month][google.cloud.channel.v1.RepricingConfig.effective_invoice_month]
	//  then it indicates this was set mid-month.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerRepricingConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
