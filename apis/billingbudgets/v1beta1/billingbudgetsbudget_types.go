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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BillingBudgetsBudgetGVK = GroupVersion.WithKind("BillingBudgetsBudget")

// BillingBudgetsBudgetSpec defines the desired state of BillingBudgetsBudget
// +kcc:spec:proto=google.cloud.billing.budgets.v1.Budget
type BillingBudgetsBudgetSpec struct {
	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.notifications_rule
	AllUpdatesRule *BudgetNotificationsRule `json:"allUpdatesRule,omitempty"`

	// Required. Budgeted amount.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.amount
	Amount *BudgetAmount `json:"amount"`

	// Immutable.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="BillingAccountRef field is immutable"
	BillingAccountRef BillingAccountRef `json:"billingAccountRef"`

	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.budget_filter
	BudgetFilter *BudgetFilter `json:"budgetFilter,omitempty"`

	// User data for display name in UI. The name must be less than or equal to 60 characters.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.threshold_rules
	ThresholdRules []BudgetThresholdRule `json:"thresholdRules,omitempty"`
}

type BillingAccountRef struct {
	/* The billing account of the resource

	Allowed value: The Google Cloud resource name of a Google Cloud Billing Account (format: `billingAccounts/{{name}}`). */
	External string `json:"external,omitempty"`

	/* [WARNING] BillingAccount not yet supported in Config Connector, use 'external' field to reference existing resources.
	Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type PubSubTopicRef struct {
	/* Optional. The name of the Pub/Sub topic where budget related messages will be published, in the form `projects/{project_id}/topics/{topic_id}`. Updates are sent at regular intervals to the topic. The topic needs to be created before the budget is created; see https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications for more details. Caller is expected to have `pubsub.topics.setIamPolicy` permission on the topic when it's set for a budget, otherwise, the API call will fail with PERMISSION_DENIED. See https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task for more details on Pub/Sub roles and permissions.

	Allowed value: The Google Cloud resource name of a `PubSubTopic` resource (format: `projects/{{project}}/topics/{{name}}`). */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type MonitoringNotificationChannelRef struct {
	/* Allowed value: The Google Cloud resource name of a `MonitoringNotificationChannel` resource (format: `projects/{{project}}/notificationChannels/{{name}}`). */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type ProjectRef struct {
	/* Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`). */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type CloudBillingBillingAccountRef struct {
	External string `json:"external,omitempty"`

	/* [WARNING] CloudBillingBillingAccount not yet supported in Config Connector, use 'external' field to reference existing resources.
	Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.NotificationsRule
type BudgetNotificationsRule struct {
	// Optional. When set to true, disables default notifications sent when a threshold is exceeded. Default notifications are sent to those with Billing Account Administrator and Billing Account User IAM roles for the target account.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.disable_default_iam_recipients
	DisableDefaultIAMRecipients *bool `json:"disableDefaultIamRecipients,omitempty"`

	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.monitoring_notification_channels
	MonitoringNotificationChannels []MonitoringNotificationChannelRef `json:"monitoringNotificationChannels,omitempty"`

	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.pubsub_topic
	PubsubTopicRef *PubSubTopicRef `json:"pubsubTopicRef,omitempty"`

	// Optional. Required when NotificationsRule.pubsub_topic is set. The schema version of the notification sent to NotificationsRule.pubsub_topic. Only "1.0" is accepted. It represents the JSON schema as defined in https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.NotificationsRule.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.BudgetAmount
// +kubebuilder:validation:XValidation:rule="has(self.specifiedAmount) != has(self.lastPeriodAmount)",message="Only one of specifiedAmount or lastPeriodAmount can be set"
type BudgetAmount struct {
	// Use the last period's actual spend as the budget for the present period. LastPeriodAmount can only be set when the budget's time period is a Filter.calendar_period.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.BudgetAmount.last_period_amount
	LastPeriodAmount *apiextensionsv1.JSON `json:"lastPeriodAmount,omitempty"`

	// A specified amount to use as the budget. `currency_code` is optional. If specified when creating a budget, it must match the currency of the billing account. If specified when updating a budget, it must match the currency_code of the existing budget. The `currency_code` is provided on output.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.BudgetAmount.specified_amount
	SpecifiedAmount *Money `json:"specifiedAmount,omitempty"`
}

// +kcc:proto=google.type.Money
type Money struct {
	// Immutable. The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int64 `json:"nanos,omitempty"`

	// The whole units of the amount. For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.Filter
// +kubebuilder:validation:XValidation:rule="has(self.calendarPeriod) != has(self.customPeriod)",message="Only one of calendarPeriod or customPeriod can be set"
type BudgetFilter struct {
	// Optional. Specifies to track usage for recurring calendar period. For example, assume that CalendarPeriod.QUARTER is set. The budget will track usage from April 1 to June 30, when the current calendar month is April, May, June. After that, it will track usage from July 1 to September 30 when the current calendar month is July, August, September, so on. Possible values: CALENDAR_PERIOD_UNSPECIFIED, MONTH, QUARTER, YEAR
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.calendar_period
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	// Optional. If Filter.credit_types_treatment is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See a list of acceptable credit type values. If Filter.credit_types_treatment is not INCLUDE_SPECIFIED_CREDITS, this field must be empty.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.credit_types
	CreditTypes []string `json:"creditTypes,omitempty"`

	// Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.credit_types_treatment
	CreditTypesTreatment *string `json:"creditTypesTreatment,omitempty"`

	// Optional. Specifies to track usage from any start date (required) to any end date (optional). This time period is static, it does not recur.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.custom_period
	CustomPeriod *BudgetCustomPeriod `json:"customPeriod,omitempty"`

	// Optional. A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget. Currently, multiple entries or multiple values per entry are not allowed. If omitted, the report will include all labeled and unlabeled usage.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.labels
	Labels map[string]BudgetLabels `json:"labels,omitempty"`

	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.projects
	Projects []ProjectRef `json:"projects,omitempty"`

	// Optional. A set of services of the form `services/{service_id}`, specifying that usage from only this set of services should be included in the budget. If omitted, the report will include usage for all the services. The service names are available through the Catalog API: https://cloud.google.com/billing/v1/how-tos/catalog-api.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.services
	Services []string `json:"services,omitempty"`

	// +kcc:proto:field=google.cloud.billing.budgets.v1.Filter.subaccounts
	Subaccounts []CloudBillingBillingAccountRef `json:"subaccounts,omitempty"`
}

type BudgetLabels struct {
	// The values of the label
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.CustomPeriod
type BudgetCustomPeriod struct {
	// Optional. The end date of the time period. Budgets with elapsed end date won't be processed. If unset, specifies to track all usage incurred since the start_date.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.CustomPeriod.end_date
	EndDate *Date `json:"endDate,omitempty"`

	// Required. The start date must be after January 1, 2017.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.CustomPeriod.start_date
	StartDate *Date `json:"startDate"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int64 `json:"day,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int64 `json:"month,omitempty"`

	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int64 `json:"year,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1.ThresholdRule
type BudgetThresholdRule struct {
	// Optional. The type of basis used to determine if spend has passed the threshold. Behavior defaults to CURRENT_SPEND if not set. Possible values: BASIS_UNSPECIFIED, CURRENT_SPEND, FORECASTED_SPEND
	// +kcc:proto:field=google.cloud.billing.budgets.v1.ThresholdRule.spend_basis
	SpendBasis *string `json:"spendBasis,omitempty"`

	// Required. Send an alert when this threshold is exceeded. This is a 1.0-based percentage, so 0.5 = 50%. Validation: non-negative number.
	// +kubebuilder:validation:Format=double
	// +kcc:proto:field=google.cloud.billing.budgets.v1.ThresholdRule.threshold_percent
	ThresholdPercent *float64 `json:"thresholdPercent"`
}

// BillingBudgetsBudgetStatus defines the config connector machine state of BillingBudgetsBudget
type BillingBudgetsBudgetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.Budget.etag
	Etag *string `json:"etag,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// BillingBudgetsBudgetObservedState is the state of the BillingBudgetsBudget resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.billing.budgets.v1.Budget
type BillingBudgetsBudgetObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbillingbudgetsbudget;gcpbillingbudgetsbudgets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BillingBudgetsBudget is the Schema for the BillingBudgetsBudget API
// +k8s:openapi-gen=true
type BillingBudgetsBudget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BillingBudgetsBudgetSpec   `json:"spec,omitempty"`
	Status BillingBudgetsBudgetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BillingBudgetsBudgetList contains a list of BillingBudgetsBudget
type BillingBudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BillingBudgetsBudget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BillingBudgetsBudget{}, &BillingBudgetsBudgetList{})
}
