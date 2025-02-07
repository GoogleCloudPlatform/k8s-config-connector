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


// +kcc:proto=google.cloud.channel.v1.AssociationInfo
type AssociationInfo struct {
	// The name of the base entitlement, for which this entitlement is an add-on.
	// +kcc:proto:field=google.cloud.channel.v1.AssociationInfo.base_entitlement
	BaseEntitlement *string `json:"baseEntitlement,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.CommitmentSettings
type CommitmentSettings struct {

	// Optional. Renewal settings applicable for a commitment-based Offer.
	// +kcc:proto:field=google.cloud.channel.v1.CommitmentSettings.renewal_settings
	RenewalSettings *RenewalSettings `json:"renewalSettings,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Entitlement
type Entitlement struct {

	// Required. The offer resource name for which the entitlement is to be
	//  created. Takes the form: accounts/{account_id}/offers/{offer_id}.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.offer
	Offer *string `json:"offer,omitempty"`

	// Commitment settings for a commitment-based Offer.
	//  Required for commitment based offers.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.commitment_settings
	CommitmentSettings *CommitmentSettings `json:"commitmentSettings,omitempty"`

	// Optional. This purchase order (PO) information is for resellers to use for
	//  their company tracking usage. If a purchaseOrderId value is given, it
	//  appears in the API responses and shows up in the invoice. The property
	//  accepts up to 80 plain text characters. This is only supported for Google
	//  Workspace entitlements.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.purchase_order_id
	PurchaseOrderID *string `json:"purchaseOrderID,omitempty"`

	// Association information to other entitlements.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.association_info
	AssociationInfo *AssociationInfo `json:"associationInfo,omitempty"`

	// Extended entitlement parameters. When creating an entitlement, valid
	//  parameter names and values are defined in the
	//  [Offer.parameter_definitions][google.cloud.channel.v1.Offer.parameter_definitions].
	//
	//  For Google Workspace, the following Parameters may be accepted as input:
	//
	//  - max_units: The maximum assignable units for a flexible offer
	//
	//  OR
	//
	//  - num_units: The total commitment for commitment-based offers
	//
	//  The response may additionally include the following output-only Parameters:
	//
	//  - assigned_units: The number of licenses assigned to users.
	//
	//  For Google Cloud billing subaccounts, the following Parameter may be
	//  accepted as input:
	//
	//  - display_name: The display name of the billing subaccount.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.parameters
	Parameters []Parameter `json:"parameters,omitempty"`

	// Optional. The billing account resource name that is used to pay for this
	//  entitlement.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.billing_account
	BillingAccount *string `json:"billingAccount,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Parameter
type Parameter struct {
	// Name of the parameter.
	// +kcc:proto:field=google.cloud.channel.v1.Parameter.name
	Name *string `json:"name,omitempty"`

	// Value of the parameter.
	// +kcc:proto:field=google.cloud.channel.v1.Parameter.value
	Value *Value `json:"value,omitempty"`
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

// +kcc:proto=google.cloud.channel.v1.ProvisionedService
type ProvisionedService struct {
}

// +kcc:proto=google.cloud.channel.v1.RenewalSettings
type RenewalSettings struct {
	// If false, the plan will be completed at the end date.
	// +kcc:proto:field=google.cloud.channel.v1.RenewalSettings.enable_renewal
	EnableRenewal *bool `json:"enableRenewal,omitempty"`

	// If true and enable_renewal = true, the unit (for example seats or licenses)
	//  will be set to the number of active units at renewal time.
	// +kcc:proto:field=google.cloud.channel.v1.RenewalSettings.resize_unit_count
	ResizeUnitCount *bool `json:"resizeUnitCount,omitempty"`

	// Describes how a reseller will be billed.
	// +kcc:proto:field=google.cloud.channel.v1.RenewalSettings.payment_plan
	PaymentPlan *string `json:"paymentPlan,omitempty"`

	// Describes how frequently the reseller will be billed, such as
	//  once per month.
	// +kcc:proto:field=google.cloud.channel.v1.RenewalSettings.payment_cycle
	PaymentCycle *Period `json:"paymentCycle,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.TrialSettings
type TrialSettings struct {
	// Determines if the entitlement is in a trial or not:
	//
	//  * `true` - The entitlement is in trial.
	//  * `false` - The entitlement is not in trial.
	// +kcc:proto:field=google.cloud.channel.v1.TrialSettings.trial
	Trial *bool `json:"trial,omitempty"`

	// Date when the trial ends. The value is in milliseconds
	//  using the UNIX Epoch format. See an example [Epoch
	//  converter](https://www.epochconverter.com).
	// +kcc:proto:field=google.cloud.channel.v1.TrialSettings.end_time
	EndTime *string `json:"endTime,omitempty"`
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

// +kcc:proto=google.cloud.channel.v1.CommitmentSettings
type CommitmentSettingsObservedState struct {
	// Output only. Commitment start timestamp.
	// +kcc:proto:field=google.cloud.channel.v1.CommitmentSettings.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Commitment end timestamp.
	// +kcc:proto:field=google.cloud.channel.v1.CommitmentSettings.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Entitlement
type EntitlementObservedState struct {
	// Output only. Resource name of an entitlement in the form:
	//  accounts/{account_id}/customers/{customer_id}/entitlements/{entitlement_id}.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which the entitlement is created.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the entitlement is updated.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Commitment settings for a commitment-based Offer.
	//  Required for commitment based offers.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.commitment_settings
	CommitmentSettings *CommitmentSettingsObservedState `json:"commitmentSettings,omitempty"`

	// Output only. Current provisioning state of the entitlement.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.provisioning_state
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Output only. Service provisioning details for the entitlement.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.provisioned_service
	ProvisionedService *ProvisionedService `json:"provisionedService,omitempty"`

	// Output only. Enumerable of all current suspension reasons for an
	//  entitlement.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.suspension_reasons
	SuspensionReasons []string `json:"suspensionReasons,omitempty"`

	// Output only. Settings for trial offers.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.trial_settings
	TrialSettings *TrialSettings `json:"trialSettings,omitempty"`

	// Extended entitlement parameters. When creating an entitlement, valid
	//  parameter names and values are defined in the
	//  [Offer.parameter_definitions][google.cloud.channel.v1.Offer.parameter_definitions].
	//
	//  For Google Workspace, the following Parameters may be accepted as input:
	//
	//  - max_units: The maximum assignable units for a flexible offer
	//
	//  OR
	//
	//  - num_units: The total commitment for commitment-based offers
	//
	//  The response may additionally include the following output-only Parameters:
	//
	//  - assigned_units: The number of licenses assigned to users.
	//
	//  For Google Cloud billing subaccounts, the following Parameter may be
	//  accepted as input:
	//
	//  - display_name: The display name of the billing subaccount.
	// +kcc:proto:field=google.cloud.channel.v1.Entitlement.parameters
	Parameters []ParameterObservedState `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Parameter
type ParameterObservedState struct {
	// Output only. Specifies whether this parameter is allowed to be changed. For
	//  example, for a Google Workspace Business Starter entitlement in commitment
	//  plan, num_units is editable when entitlement is active.
	// +kcc:proto:field=google.cloud.channel.v1.Parameter.editable
	Editable *bool `json:"editable,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ProvisionedService
type ProvisionedServiceObservedState struct {
	// Output only. Provisioning ID of the entitlement. For Google Workspace, this
	//  is the underlying Subscription ID. For Google Cloud, this is the Billing
	//  Account ID of the billing subaccount.
	// +kcc:proto:field=google.cloud.channel.v1.ProvisionedService.provisioning_id
	ProvisioningID *string `json:"provisioningID,omitempty"`

	// Output only. The product pertaining to the provisioning resource as
	//  specified in the Offer.
	// +kcc:proto:field=google.cloud.channel.v1.ProvisionedService.product_id
	ProductID *string `json:"productID,omitempty"`

	// Output only. The SKU pertaining to the provisioning resource as specified
	//  in the Offer.
	// +kcc:proto:field=google.cloud.channel.v1.ProvisionedService.sku_id
	SkuID *string `json:"skuID,omitempty"`
}
