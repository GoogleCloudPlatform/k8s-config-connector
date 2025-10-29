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

// +generated:types
// krm.group: billingbudgets.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.billing.budgets.v1beta1
// resource: BillingBudgetsBudget:Budget

package v1beta1

// +kcc:proto=google.cloud.billing.budgets.v1beta1.BudgetAmount
type BudgetAmount struct {
	// A specified amount to use as the budget.
	//  `currency_code` is optional. If specified when creating a budget, it must
	//  match the currency of the billing account. If specified when updating a
	//  budget, it must match the currency_code of the existing budget.
	//  The `currency_code` is provided on output.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.BudgetAmount.specified_amount
	SpecifiedAmount *Money `json:"specifiedAmount,omitempty"`

	// Use the last period's actual spend as the budget for the present period.
	//  LastPeriodAmount can only be set when the budget's time period is a
	//  [Filter.calendar_period][google.cloud.billing.budgets.v1beta1.Filter.calendar_period].
	//  It cannot be set in combination with
	//  [Filter.custom_period][google.cloud.billing.budgets.v1beta1.Filter.custom_period].
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.BudgetAmount.last_period_amount
	LastPeriodAmount *LastPeriodAmount `json:"lastPeriodAmount,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1beta1.CustomPeriod
type CustomPeriod struct {
	// Required. The start date must be after January 1, 2017.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.CustomPeriod.start_date
	StartDate *Date `json:"startDate,omitempty"`

	// Optional. The end date of the time period. Budgets with elapsed end date
	//  won't be processed. If unset, specifies to track all usage incurred since
	//  the start_date.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.CustomPeriod.end_date
	EndDate *Date `json:"endDate,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1beta1.LastPeriodAmount
type LastPeriodAmount struct {
}

// +kcc:proto=google.cloud.billing.budgets.v1beta1.ThresholdRule
type ThresholdRule struct {
	// Required. Send an alert when this threshold is exceeded.
	//  This is a 1.0-based percentage, so 0.5 = 50%.
	//  Validation: non-negative number.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.ThresholdRule.threshold_percent
	ThresholdPercent *float64 `json:"thresholdPercent,omitempty"`

	// Optional. The type of basis used to determine if spend has passed the
	//  threshold. Behavior defaults to CURRENT_SPEND if not set.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.ThresholdRule.spend_basis
	SpendBasis *string `json:"spendBasis,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
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
