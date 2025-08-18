// Copyright 2025 Google LLC
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
	billing "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	monitoring "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	pubsub "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BillingBudgetsBudgetGVK = GroupVersion.WithKind("BillingBudgetsBudget")

// BillingBudgetsBudgetSpec defines the desired state of BillingBudgetsBudget
// +kcc:spec:proto=google.cloud.billing.budgets.v1beta1.Budget
type BillingBudgetsBudgetSpec struct {
	// The billing account for which the budget applies.
	BillingAccountRef *billing.BillingAccountRef `json:"billingAccountRef,omitempty"`

	// The BillingBudgetsBudget name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User data for display name in UI.
	//  Validation: <= 60 chars.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Budget.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Filters that define which resources are used to compute the
	//  actual spend against the budget amount, such as projects, services, and the
	//  budget's time period, as well as other filters.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Budget.budget_filter
	BudgetFilter *Filter `json:"budgetFilter,omitempty"`

	// Required. Budgeted amount.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Budget.amount
	Amount *BudgetAmount `json:"amount,omitempty"`

	// Optional. Rules that trigger alerts (notifications of thresholds
	//  being crossed) when spend exceeds the specified percentages of the budget.
	//
	//  Optional for `pubsubTopic` notifications.
	//
	//  Required if using email notifications.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Budget.threshold_rules
	ThresholdRules []*ThresholdRule `json:"thresholdRules,omitempty"`

	// Optional. Rules to apply to notifications sent based on budget spend and
	//  thresholds.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Budget.all_updates_rule
	AllUpdatesRule *AllUpdatesRule `json:"allUpdatesRule,omitempty"`
}

// +kcc:proto=google.cloud.billing.budgets.v1beta1.AllUpdatesRule
type AllUpdatesRule struct {
	// Optional. The name of the Pub/Sub topic where budget related messages will
	//  be published, in the form `projects/{project_id}/topics/{topic_id}`.
	//  Updates are sent at regular intervals to the topic. The topic needs to be
	//  created before the budget is created; see
	//  https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications
	//  for more details.
	//  Caller is expected to have
	//  `pubsub.topics.setIamPolicy` permission on the topic when it's set for a
	//  budget, otherwise, the API call will fail with PERMISSION_DENIED. See
	//  https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task
	//  for more details on Pub/Sub roles and permissions.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.AllUpdatesRule.pubsub_topic
	PubsubTopicRef *pubsub.PubSubTopicRef `json:"pubsubTopicRef,omitempty"`

	// Optional. Required when
	//  [AllUpdatesRule.pubsub_topic][google.cloud.billing.budgets.v1beta1.AllUpdatesRule.pubsub_topic]
	//  is set. The schema version of the notification sent to
	//  [AllUpdatesRule.pubsub_topic][google.cloud.billing.budgets.v1beta1.AllUpdatesRule.pubsub_topic].
	//  Only "1.0" is accepted. It represents the JSON schema as defined in
	//  https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.AllUpdatesRule.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Optional. Targets to send notifications to when a threshold is exceeded.
	//  This is in addition to default recipients who have billing account IAM
	//  roles. The value is the full REST resource name of a monitoring
	//  notification channel with the form
	//  `projects/{project_id}/notificationChannels/{channel_id}`. A maximum of 5
	//  channels are allowed. See
	//  https://cloud.google.com/billing/docs/how-to/budgets-notification-recipients
	//  for more details.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.AllUpdatesRule.monitoring_notification_channels
	MonitoringNotificationChannelRefs []*monitoring.NotificationChannelRef `json:"monitoringNotificationChannels,omitempty"`

	// Optional. When set to true, disables default notifications sent when a
	//  threshold is exceeded. Default notifications are sent to those with Billing
	//  Account Administrator and Billing Account User IAM roles for the target
	//  account.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.AllUpdatesRule.disable_default_iam_recipients
	DisableDefaultIAMRecipients *bool `json:"disableDefaultIamRecipients,omitempty"`

	/* NOTYET: Terraform
	// Optional. When set to true, and when the budget has a single project
	//  configured, notifications will be sent to project level recipients of that
	//  project. This field will be ignored if the budget has multiple or no
	//  project configured.
	//
	//  Currently, project level recipients are the users with `Owner` role on a
	//  cloud project.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.AllUpdatesRule.enable_project_level_recipients
	EnableProjectLevelRecipients *bool `json:"enableProjectLevelRecipients,omitempty"`
	*/
}

// +kcc:proto=google.cloud.billing.budgets.v1beta1.Filter
type Filter struct {
	// Optional. A set of projects of the form `projects/{project}`,
	//  specifying that usage from only this set of projects should be
	//  included in the budget. If omitted, the report will include all usage for
	//  the billing account, regardless of which project the usage occurred on.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.projects
	ProjectRefs []*refs.ProjectRef `json:"projects,omitempty"`

	/* NOTYET: Terraform
	// Optional. A set of folder and organization names of the form
	//  `folders/{folderId}` or `organizations/{organizationId}`, specifying that
	//  usage from only this set of folders and organizations should be included in
	//  the budget. If omitted, the budget includes all usage that the billing
	//  account pays for. If the folder or organization contains projects that are
	//  paid for by a different Cloud Billing account, the budget *doesn't* apply
	//  to those projects.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.resource_ancestors
	ResourceAncestors []string `json:"resourceAncestors,omitempty"`
	*/

	// Optional. If
	//  [Filter.credit_types_treatment][google.cloud.billing.budgets.v1beta1.Filter.credit_types_treatment]
	//  is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be
	//  subtracted from gross cost to determine the spend for threshold
	//  calculations. See [a list of acceptable credit type
	//  values](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-tables#credits-type).
	//
	//  If
	//  [Filter.credit_types_treatment][google.cloud.billing.budgets.v1beta1.Filter.credit_types_treatment]
	//  is **not** INCLUDE_SPECIFIED_CREDITS, this field must be empty.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.credit_types
	CreditTypes []string `json:"creditTypes,omitempty"`

	// Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.credit_types_treatment
	CreditTypesTreatment *string `json:"creditTypesTreatment,omitempty"`

	// Optional. A set of services of the form `services/{service_id}`,
	//  specifying that usage from only this set of services should be
	//  included in the budget. If omitted, the report will include usage for
	//  all the services.
	//  The service names are available through the Catalog API:
	//  https://cloud.google.com/billing/v1/how-tos/catalog-api.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.services
	Services []string `json:"services,omitempty"`

	// Optional. A set of subaccounts of the form `billingAccounts/{account_id}`,
	//  specifying that usage from only this set of subaccounts should be included
	//  in the budget. If a subaccount is set to the name of the parent account,
	//  usage from the parent account will be included. If omitted, the
	//  report will include usage from the parent account and all
	//  subaccounts, if they exist.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.subaccounts
	Subaccounts []string `json:"subaccounts,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. Specifies to track usage for recurring calendar period.
	//  For example, assume that CalendarPeriod.QUARTER is set. The budget will
	//  track usage from April 1 to June 30, when the current calendar month is
	//  April, May, June. After that, it will track usage from July 1 to
	//  September 30 when the current calendar month is July, August, September,
	//  so on.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.calendar_period
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	// Optional. Specifies to track usage from any start date (required) to any
	//  end date (optional). This time period is static, it does not recur.
	// +kcc:proto:field=google.cloud.billing.budgets.v1beta1.Filter.custom_period
	CustomPeriod *CustomPeriod `json:"customPeriod,omitempty"`
}

// BillingBudgetsBudgetStatus defines the config connector machine state of BillingBudgetsBudget
type BillingBudgetsBudgetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// This field is deprecated and will not be set in future versions
	Etag *string `json:"etag,omitempty"`

	// A unique specifier for the BillingBudgetsBudget resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	/* NOTYET: Terraform compat
	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BillingBudgetsBudgetObservedState `json:"observedState,omitempty"`
	*/
}

// BillingBudgetsBudgetObservedState is the state of the BillingBudgetsBudget resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.billing.budgets.v1beta1.Budget
type BillingBudgetsBudgetObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbillingbudgetsbudget;gcpbillingbudgetsbudgets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
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
