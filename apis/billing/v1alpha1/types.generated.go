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


// +kcc:proto=google.cloud.billing.budgets.v1.Budget
type Budget struct {

	// User data for display name in UI. The name must be less than or equal to 60
	//  characters.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Filters that define which resources are used to compute the
	//  actual spend against the budget amount, such as projects, services, and the
	//  budget's time period, as well as other filters.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.budget_filter
	BudgetFilter *Filter `json:"budgetFilter,omitempty"`

	// Required. Budgeted amount.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.amount
	Amount *BudgetAmount `json:"amount,omitempty"`

	// Optional. Rules that trigger alerts (notifications of thresholds
	//  being crossed) when spend exceeds the specified percentages of the budget.
	//
	//  Optional for `pubsubTopic` notifications.
	//
	//  Required if using email notifications.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.threshold_rules
	ThresholdRules []ThresholdRule `json:"thresholdRules,omitempty"`

	// Optional. Rules to apply to notifications sent based on budget spend and
	//  thresholds.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.notifications_rule
	NotificationsRule *NotificationsRule `json:"notificationsRule,omitempty"`

	// Optional. Etag to validate that the object is unchanged for a
	//  read-modify-write operation.
	//  An empty etag causes an update to overwrite other changes.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.BudgetAmount
type BudgetAmount struct {
	// A specified amount to use as the budget.
	//  `currency_code` is optional. If specified when creating a budget, it must
	//  match the currency of the billing account. If specified when updating a
	//  budget, it must match the currency_code of the existing budget.
	//  The `currency_code` is provided on output.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.BudgetAmount.specified_amount
	SpecifiedAmount *Money `json:"specifiedAmount,omitempty"`

	// Use the last period's actual spend as the budget for the present period.
	//  LastPeriodAmount can only be set when the budget's time period is a
	//  [Filter.calendar_period][google.cloud.billing.budgets.v1.Filter.calendar_period].
	//  It cannot be set in combination with
	//  [Filter.custom_period][google.cloud.billing.budgets.v1.Filter.custom_period].
	// +kcc:proto:field=google.cloud.billing.budgets.v1.BudgetAmount.last_period_amount
	LastPeriodAmount *LastPeriodAmount `json:"lastPeriodAmount,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.CustomPeriod
type CustomPeriod struct {
	// Required. The start date must be after January 1, 2017.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.CustomPeriod.start_date
	StartDate *Date `json:"startDate,omitempty"`

	// Optional. The end date of the time period. Budgets with elapsed end date
	//  won't be processed. If unset, specifies to track all usage incurred since
	//  the start_date.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.CustomPeriod.end_date
	EndDate *Date `json:"endDate,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.Filter
type Filter struct {
	// Optional. A set of projects of the form `projects/{project}`,
	//  specifying that usage from only this set of projects should be
	//  included in the budget. If omitted, the report includes all usage for
	//  the billing account, regardless of which project the usage occurred on.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.projects
	Projects []string `json:"projects,omitempty"`

	// Optional. A set of folder and organization names of the form
	//  `folders/{folderId}` or `organizations/{organizationId}`, specifying that
	//  usage from only this set of folders and organizations should be included in
	//  the budget. If omitted, the budget includes all usage that the billing
	//  account pays for. If the folder or organization contains projects that are
	//  paid for by a different Cloud Billing account, the budget *doesn't* apply
	//  to those projects.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.resource_ancestors
	ResourceAncestors []string `json:"resourceAncestors,omitempty"`

	// Optional. If
	//  [Filter.credit_types_treatment][google.cloud.billing.budgets.v1.Filter.credit_types_treatment]
	//  is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be
	//  subtracted from gross cost to determine the spend for threshold
	//  calculations. See [a list of acceptable credit type
	//  values](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-tables#credits-type).
	//
	//  If
	//  [Filter.credit_types_treatment][google.cloud.billing.budgets.v1.Filter.credit_types_treatment]
	//  is **not** INCLUDE_SPECIFIED_CREDITS, this field must be empty.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.credit_types
	CreditTypes []string `json:"creditTypes,omitempty"`

	// Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.credit_types_treatment
	CreditTypesTreatment *string `json:"creditTypesTreatment,omitempty"`

	// Optional. A set of services of the form `services/{service_id}`,
	//  specifying that usage from only this set of services should be
	//  included in the budget. If omitted, the report includes usage for
	//  all the services.
	//  The service names are available through the Catalog API:
	//  https://cloud.google.com/billing/v1/how-tos/catalog-api.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.services
	Services []string `json:"services,omitempty"`

	// Optional. A set of subaccounts of the form `billingAccounts/{account_id}`,
	//  specifying that usage from only this set of subaccounts should be included
	//  in the budget. If a subaccount is set to the name of the parent account,
	//  usage from the parent account is included. If the field is omitted,
	//  the report includes usage from the parent account and all subaccounts,
	//  if they exist.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.subaccounts
	Subaccounts []string `json:"subaccounts,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. Specifies to track usage for recurring calendar period.
	//  For example, assume that CalendarPeriod.QUARTER is set. The budget
	//  tracks usage from April 1 to June 30, when the current calendar month is
	//  April, May, June. After that, it tracks usage from July 1 to
	//  September 30 when the current calendar month is July, August, September,
	//  so on.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.calendar_period
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	// Optional. Specifies to track usage from any start date (required) to any
	//  end date (optional). This time period is static, it does not recur.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.custom_period
	CustomPeriod *CustomPeriod `json:"customPeriod,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.LastPeriodAmount
type LastPeriodAmount struct {
}

// +kcc:proto=google.cloud.billing.budgets.v1.NotificationsRule
type NotificationsRule struct {
	// Optional. The name of the Pub/Sub topic where budget-related messages are
	//  published, in the form `projects/{project_id}/topics/{topic_id}`. Updates
	//  are sent to the topic at regular intervals; the timing of the updates is
	//  not dependent on the [threshold rules](#thresholdrule) you've set.
	//
	//  Note that if you want your
	//  [Pub/Sub JSON
	//  object](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format)
	//  to contain data for `alertThresholdExceeded`, you need at least one
	//  [alert threshold rule](#thresholdrule). When you set threshold rules, you
	//  must also enable at least one of the email notification options, either
	//  using the default IAM recipients or Cloud Monitoring email notification
	//  channels.
	//
	//  To use Pub/Sub topics with budgets, you must do the following:
	//
	//  1. Create the Pub/Sub topic
	//  before connecting it to your budget. For guidance, see
	//  [Manage programmatic budget alert
	//  notifications](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications).
	//
	//  2. Grant the API caller the `pubsub.topics.setIamPolicy` permission on
	//  the Pub/Sub topic. If not set, the API call fails with PERMISSION_DENIED.
	//  For additional details on Pub/Sub roles and permissions, see
	//  [Permissions required for this
	//  task](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task).
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.pubsub_topic
	PubsubTopic *string `json:"pubsubTopic,omitempty"`

	// Optional. Required when
	//  [NotificationsRule.pubsub_topic][google.cloud.billing.budgets.v1.NotificationsRule.pubsub_topic]
	//  is set. The schema version of the notification sent to
	//  [NotificationsRule.pubsub_topic][google.cloud.billing.budgets.v1.NotificationsRule.pubsub_topic].
	//  Only "1.0" is accepted. It represents the JSON schema as defined in
	//  https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Optional. Email targets to send notifications to when a threshold is
	//  exceeded. This is in addition to the `DefaultIamRecipients` who receive
	//  alert emails based on their billing account IAM role. The value is the full
	//  REST resource name of a Cloud Monitoring email notification channel with
	//  the form `projects/{project_id}/notificationChannels/{channel_id}`. A
	//  maximum of 5 email notifications are allowed.
	//
	//  To customize budget alert email recipients with monitoring notification
	//  channels, you _must create the monitoring notification channels before
	//  you link them to a budget_. For guidance on setting up notification
	//  channels to use with budgets, see
	//  [Customize budget alert email
	//  recipients](https://cloud.google.com/billing/docs/how-to/budgets-notification-recipients).
	//
	//  For Cloud Billing budget alerts, you _must use email notification
	//  channels_. The other types of notification channels are _not_
	//  supported, such as Slack, SMS, or PagerDuty. If you want to
	//  [send budget notifications to
	//  Slack](https://cloud.google.com/billing/docs/how-to/notify#send_notifications_to_slack),
	//  use a pubsubTopic and configure
	//  [programmatic
	//  notifications](https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications).
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.monitoring_notification_channels
	MonitoringNotificationChannels []string `json:"monitoringNotificationChannels,omitempty"`

	// Optional. When set to true, disables default notifications sent when a
	//  threshold is exceeded. Default notifications are sent to those with Billing
	//  Account Administrator and Billing Account User IAM roles for the target
	//  account.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.disable_default_iam_recipients
	DisableDefaultIamRecipients *bool `json:"disableDefaultIamRecipients,omitempty"`

	// Optional. When set to true, and when the budget has a single project
	//  configured, notifications will be sent to project level recipients of that
	//  project. This field will be ignored if the budget has multiple or no
	//  project configured.
	//
	//  Currently, project level recipients are the users with `Owner` role on a
	//  cloud project.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.enable_project_level_recipients
	EnableProjectLevelRecipients *bool `json:"enableProjectLevelRecipients,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.ThresholdRule
type ThresholdRule struct {
	// Required. Send an alert when this threshold is exceeded.
	//  This is a 1.0-based percentage, so 0.5 = 50%.
	//  Validation: non-negative number.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.ThresholdRule.threshold_percent
	ThresholdPercent *float64 `json:"thresholdPercent,omitempty"`

	// Optional. The type of basis used to determine if spend has passed the
	//  threshold. Behavior defaults to CURRENT_SPEND if not set.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.ThresholdRule.spend_basis
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

// +kcc:proto=google.cloud.billing.budgets.v1.Budget
type BudgetObservedState struct {
	// Output only. Resource name of the budget.
	//  The resource name implies the scope of a budget. Values are of the form
	//  `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.name
	Name *string `json:"name,omitempty"`
}
