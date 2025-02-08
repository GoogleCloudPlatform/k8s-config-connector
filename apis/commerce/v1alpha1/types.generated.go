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


// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LineItem
type LineItem struct {
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LineItemChange
type LineItemChange struct {

	// Required. Type of the change to make.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.change_type
	ChangeType *string `json:"changeType,omitempty"`

	// Line item info after the change.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.new_line_item_info
	NewLineItemInfo *LineItemInfo `json:"newLineItemInfo,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LineItemInfo
type LineItemInfo struct {
	// Optional. The name of the offer can have either of these formats:
	//  'billingAccounts/{billing_account}/offers/{offer}',
	//  or 'services/{service}/standardOffers/{offer}'.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemInfo.offer
	Offer *string `json:"offer,omitempty"`

	// Optional. User-provided parameters.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemInfo.parameters
	Parameters []Parameter `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.Order
type Order struct {

	// Required. The user-specified name of the order.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The weak etag of the order.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.Parameter
type Parameter struct {
	// Name of the parameter.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Parameter.name
	Name *string `json:"name,omitempty"`

	// Value of parameter.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Parameter.value
	Value *Parameter_Value `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.Parameter.Value
type Parameter_Value struct {
	// Represents an int64 value.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Parameter.Value.int64_value
	Int64Value *int64 `json:"int64Value,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Parameter.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Parameter.Value.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.Subscription
type Subscription struct {
	// The timestamp when the subscription begins, if applicable.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Subscription.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The timestamp when the subscription ends, if applicable.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Subscription.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Whether auto renewal is enabled by user choice on current subscription.
	//  This field indicates order/subscription status after pending plan change is
	//  cancelled or rejected.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Subscription.auto_renewal_enabled
	AutoRenewalEnabled *bool `json:"autoRenewalEnabled,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LineItem
type LineItemObservedState struct {
	// Output only. Line item ID.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItem.line_item_id
	LineItemID *string `json:"lineItemID,omitempty"`

	// Output only. Current state and information of this item. It tells what,
	//  e.g. which offer, is currently effective.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItem.line_item_info
	LineItemInfo *LineItemInfo `json:"lineItemInfo,omitempty"`

	// Output only. A change made on the item which is pending and not yet
	//  effective. Absence of this field indicates the line item is not undergoing
	//  a change.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItem.pending_change
	PendingChange *LineItemChange `json:"pendingChange,omitempty"`

	// Output only. Changes made on the item that are not pending anymore which
	//  might be because they already took effect, were reverted by the customer,
	//  or were rejected by the partner. No more operations are allowed on these
	//  changes.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItem.change_history
	ChangeHistory []LineItemChange `json:"changeHistory,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LineItemChange
type LineItemChangeObservedState struct {
	// Output only. Change ID.
	//  All changes made within one order update operation have the same change_id.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.change_id
	ChangeID *string `json:"changeID,omitempty"`

	// Output only. Line item info before the change.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.old_line_item_info
	OldLineItemInfo *LineItemInfo `json:"oldLineItemInfo,omitempty"`

	// Output only. State of the change.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.change_state
	ChangeState *string `json:"changeState,omitempty"`

	// Output only. Provider-supplied message explaining the LineItemChange's
	//  state. Mainly used to communicate progress and ETA for provisioning in the
	//  case of `PENDING_APPROVAL`, and to explain why the change request was
	//  denied or canceled in the case of `REJECTED` and `CANCELED` states.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.state_reason
	StateReason *string `json:"stateReason,omitempty"`

	// Output only. Predefined enum types for why this line item change is in
	//  current state. For example, a line item change's state could be
	//  `LINE_ITEM_CHANGE_STATE_COMPLETED` because of end-of-term expiration,
	//  immediate cancellation initiated by the user, or system-initiated
	//  cancellation.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.change_state_reason_type
	ChangeStateReasonType *string `json:"changeStateReasonType,omitempty"`

	// Output only. A time at which the change became or will become (in case of
	//  pending change) effective.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.change_effective_time
	ChangeEffectiveTime *string `json:"changeEffectiveTime,omitempty"`

	// Output only. The time when change was initiated.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when change was updated, e.g. approved/rejected by
	//  partners or cancelled by the user.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemChange.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LineItemInfo
type LineItemInfoObservedState struct {
	// Output only. Information about the subscription created, if applicable.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LineItemInfo.subscription
	Subscription *Subscription `json:"subscription,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.Order
type OrderObservedState struct {
	// Output only. The resource name of the order.
	//  Has the form
	//  `billingAccounts/{billing_account}/orders/{order}`.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.name
	Name *string `json:"name,omitempty"`

	// Output only. The items being purchased.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.line_items
	LineItems []LineItem `json:"lineItems,omitempty"`

	// Output only. Line items that were cancelled.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.cancelled_line_items
	CancelledLineItems []LineItem `json:"cancelledLineItems,omitempty"`

	// Output only. The creation timestamp.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.Order.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
