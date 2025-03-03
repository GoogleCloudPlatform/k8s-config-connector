// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package billingbudgets

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLBudgetSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "BillingBudgets/Budget",
			Description: "The BillingBudgets Budget resource",
			StructName:  "Budget",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Budget",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "budget",
						Required:    true,
						Description: "A full instance of a Budget",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Budget",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "budget",
						Required:    true,
						Description: "A full instance of a Budget",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Budget",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "budget",
						Required:    true,
						Description: "A full instance of a Budget",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Budget",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "billingAccount",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Budget",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "billingAccount",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Budget": &dcl.Component{
					Title:     "Budget",
					ID:        "billingAccounts/{{billing_account}}/budgets/{{name}}",
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"amount",
							"billingAccount",
						},
						Properties: map[string]*dcl.Property{
							"allUpdatesRule": &dcl.Property{
								Type:        "object",
								GoName:      "AllUpdatesRule",
								GoType:      "BudgetAllUpdatesRule",
								Description: "Optional. Rules to apply to notifications sent based on budget spend and thresholds.",
								Properties: map[string]*dcl.Property{
									"disableDefaultIamRecipients": &dcl.Property{
										Type:        "boolean",
										GoName:      "DisableDefaultIamRecipients",
										Description: "Optional. When set to true, disables default notifications sent when a threshold is exceeded. Default notifications are sent to those with Billing Account Administrator and Billing Account User IAM roles for the target account.",
									},
									"monitoringNotificationChannels": &dcl.Property{
										Type:        "array",
										GoName:      "MonitoringNotificationChannels",
										Description: "Optional. Targets to send notifications to when a threshold is exceeded. This is in addition to default recipients who have billing account IAM roles. The value is the full REST resource name of a monitoring notification channel with the form `projects/{project_id}/notificationChannels/{channel_id}`. A maximum of 5 channels are allowed. See https://cloud.google.com/billing/docs/how-to/budgets-notification-recipients for more details.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Monitoring/NotificationChannel",
													Field:    "name",
												},
											},
										},
									},
									"pubsubTopic": &dcl.Property{
										Type:        "string",
										GoName:      "PubsubTopic",
										Description: "Optional. The name of the Pub/Sub topic where budget related messages will be published, in the form `projects/{project_id}/topics/{topic_id}`. Updates are sent at regular intervals to the topic. The topic needs to be created before the budget is created; see https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications for more details. Caller is expected to have `pubsub.topics.setIamPolicy` permission on the topic when it's set for a budget, otherwise, the API call will fail with PERMISSION_DENIED. See https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task for more details on Pub/Sub roles and permissions.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Pubsub/Topic",
												Field:    "name",
											},
										},
									},
									"schemaVersion": &dcl.Property{
										Type:        "string",
										GoName:      "SchemaVersion",
										Description: "Optional. Required when NotificationsRule.pubsub_topic is set. The schema version of the notification sent to NotificationsRule.pubsub_topic. Only \"1.0\" is accepted. It represents the JSON schema as defined in https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.",
									},
								},
							},
							"amount": &dcl.Property{
								Type:        "object",
								GoName:      "Amount",
								GoType:      "BudgetAmount",
								Description: "Required. Budgeted amount.",
								Properties: map[string]*dcl.Property{
									"lastPeriodAmount": &dcl.Property{
										Type:        "object",
										GoName:      "LastPeriodAmount",
										GoType:      "BudgetAmountLastPeriodAmount",
										Description: "Use the last period's actual spend as the budget for the present period. LastPeriodAmount can only be set when the budget's time period is a .",
										Conflicts: []string{
											"specifiedAmount",
										},
										Properties: map[string]*dcl.Property{},
									},
									"specifiedAmount": &dcl.Property{
										Type:        "object",
										GoName:      "SpecifiedAmount",
										GoType:      "BudgetAmountSpecifiedAmount",
										Description: "A specified amount to use as the budget. `currency_code` is optional. If specified when creating a budget, it must match the currency of the billing account. If specified when updating a budget, it must match the currency_code of the existing budget. The `currency_code` is provided on output.",
										Conflicts: []string{
											"lastPeriodAmount",
										},
										Properties: map[string]*dcl.Property{
											"currencyCode": &dcl.Property{
												Type:        "string",
												GoName:      "CurrencyCode",
												Description: "The three-letter currency code defined in ISO 4217.",
												Immutable:   true,
											},
											"nanos": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Nanos",
												Description: "Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.",
												SendEmpty:   true,
											},
											"units": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Units",
												Description: "The whole units of the amount. For example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.",
												SendEmpty:   true,
											},
										},
									},
								},
							},
							"billingAccount": &dcl.Property{
								Type:        "string",
								GoName:      "BillingAccount",
								Description: "The billing account of the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/BillingAccount",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"budgetFilter": &dcl.Property{
								Type:          "object",
								GoName:        "BudgetFilter",
								GoType:        "BudgetBudgetFilter",
								Description:   "Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"calendarPeriod": &dcl.Property{
										Type:        "string",
										GoName:      "CalendarPeriod",
										GoType:      "BudgetBudgetFilterCalendarPeriodEnum",
										Description: "Optional. Specifies to track usage for recurring calendar period. For example, assume that CalendarPeriod.QUARTER is set. The budget will track usage from April 1 to June 30, when the current calendar month is April, May, June. After that, it will track usage from July 1 to September 30 when the current calendar month is July, August, September, so on. Possible values: CALENDAR_PERIOD_UNSPECIFIED, MONTH, QUARTER, YEAR",
										Conflicts: []string{
											"customPeriod",
										},
										ServerDefault: true,
										Enum: []string{
											"CALENDAR_PERIOD_UNSPECIFIED",
											"MONTH",
											"QUARTER",
											"YEAR",
										},
									},
									"creditTypes": &dcl.Property{
										Type:        "array",
										GoName:      "CreditTypes",
										Description: "Optional. If Filter.credit_types_treatment is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See a list of acceptable credit type values. If Filter.credit_types_treatment is not INCLUDE_SPECIFIED_CREDITS, this field must be empty.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"creditTypesTreatment": &dcl.Property{
										Type:          "string",
										GoName:        "CreditTypesTreatment",
										GoType:        "BudgetBudgetFilterCreditTypesTreatmentEnum",
										Description:   "Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.",
										ServerDefault: true,
										Enum: []string{
											"INCLUDE_ALL_CREDITS",
											"EXCLUDE_ALL_CREDITS",
											"INCLUDE_SPECIFIED_CREDITS",
										},
									},
									"customPeriod": &dcl.Property{
										Type:        "object",
										GoName:      "CustomPeriod",
										GoType:      "BudgetBudgetFilterCustomPeriod",
										Description: "Optional. Specifies to track usage from any start date (required) to any end date (optional). This time period is static, it does not recur.",
										Conflicts: []string{
											"calendarPeriod",
										},
										Required: []string{
											"startDate",
										},
										Properties: map[string]*dcl.Property{
											"endDate": &dcl.Property{
												Type:        "object",
												GoName:      "EndDate",
												GoType:      "BudgetBudgetFilterCustomPeriodEndDate",
												Description: "Optional. The end date of the time period. Budgets with elapsed end date won't be processed. If unset, specifies to track all usage incurred since the start_date.",
												Immutable:   true,
												Properties: map[string]*dcl.Property{
													"day": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Day",
														Description: "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
														Immutable:   true,
													},
													"month": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Month",
														Description: "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
														Immutable:   true,
													},
													"year": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Year",
														Description: "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
														Immutable:   true,
													},
												},
											},
											"startDate": &dcl.Property{
												Type:        "object",
												GoName:      "StartDate",
												GoType:      "BudgetBudgetFilterCustomPeriodStartDate",
												Description: "Required. The start date must be after January 1, 2017.",
												Immutable:   true,
												Properties: map[string]*dcl.Property{
													"day": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Day",
														Description: "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
														Immutable:   true,
													},
													"month": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Month",
														Description: "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
														Immutable:   true,
													},
													"year": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Year",
														Description: "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
														Immutable:   true,
													},
												},
											},
										},
									},
									"labels": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type:   "object",
											GoType: "BudgetBudgetFilterLabels",
											Properties: map[string]*dcl.Property{
												"values": &dcl.Property{
													Type:        "array",
													GoName:      "Values",
													Description: "The values of the label",
													Immutable:   true,
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "string",
														GoType: "string",
													},
												},
											},
										},
										GoName:      "Labels",
										Description: "Optional. A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget. Currently, multiple entries or multiple values per entry are not allowed. If omitted, the report will include all labeled and unlabeled usage.",
									},
									"projects": &dcl.Property{
										Type:        "array",
										GoName:      "Projects",
										Description: "Optional. A set of projects of the form `projects/{project}`, specifying that usage from only this set of projects should be included in the budget. If omitted, the report will include all usage for the billing account, regardless of which project the usage occurred on. Only zero or one project can be specified currently.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Cloudresourcemanager/Project",
													Field:    "name",
												},
											},
										},
									},
									"services": &dcl.Property{
										Type:        "array",
										GoName:      "Services",
										Description: "Optional. A set of services of the form `services/{service_id}`, specifying that usage from only this set of services should be included in the budget. If omitted, the report will include usage for all the services. The service names are available through the Catalog API: https://cloud.google.com/billing/v1/how-tos/catalog-api.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"subaccounts": &dcl.Property{
										Type:        "array",
										GoName:      "Subaccounts",
										Description: "Optional. A set of subaccounts of the form `billingAccounts/{account_id}`, specifying that usage from only this set of subaccounts should be included in the budget. If a subaccount is set to the name of the parent account, usage from the parent account will be included. If the field is omitted, the report will include usage from the parent account and all subaccounts, if they exist.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Cloudbilling/BillingAccount",
													Field:    "name",
												},
											},
										},
									},
								},
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "User data for display name in UI. The name must be less than or equal to 60 characters.",
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. Resource name of the budget.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"thresholdRules": &dcl.Property{
								Type:        "array",
								GoName:      "ThresholdRules",
								Description: "Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "BudgetThresholdRules",
									Required: []string{
										"thresholdPercent",
									},
									Properties: map[string]*dcl.Property{
										"spendBasis": &dcl.Property{
											Type:        "string",
											GoName:      "SpendBasis",
											GoType:      "BudgetThresholdRulesSpendBasisEnum",
											Description: "Optional. The type of basis used to determine if spend has passed the threshold. Behavior defaults to CURRENT_SPEND if not set. Possible values: BASIS_UNSPECIFIED, CURRENT_SPEND, FORECASTED_SPEND",
											Enum: []string{
												"BASIS_UNSPECIFIED",
												"CURRENT_SPEND",
												"FORECASTED_SPEND",
											},
										},
										"thresholdPercent": &dcl.Property{
											Type:        "number",
											Format:      "double",
											GoName:      "ThresholdPercent",
											Description: "Required. Send an alert when this threshold is exceeded. This is a 1.0-based percentage, so 0.5 = 50%. Validation: non-negative number.",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
