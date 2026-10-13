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

package v1alpha1

import (
	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CommerceProducerPrivateOfferGVK = GroupVersion.WithKind("CommerceProducerPrivateOffer")

// CommerceProducerPrivateOfferSpec defines the desired state of CommerceProducerPrivateOffer
// +kcc:spec:proto=google.cloud.commerceproducer.v1beta.PrivateOffer
type CommerceProducerPrivateOfferSpec struct {
	// The project that this resource belongs to.
	// +required
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The CommerceProducerPrivateOffer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Configurations for the offer that is associated with a single
	// product.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.single_product_offer
	// +kubebuilder:validation:Optional
	SingleProductOffer *PrivateOfferSingleProductOffer `json:"singleProductOffer,omitempty"`

	// Optional. Unstructured text content that is not visible to the customer.
	// Intended to be used by partners for storing notes about the private offer.
	// Maximum length: 1500 characters.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.internal_note
	// +kubebuilder:validation:Optional
	InternalNote *string `json:"internalNote,omitempty"`

	// Optional. The type of the deal transacted with the offer.
	// The deal type is not visible to customers.
	// Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.offer_deal_type
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=CHANNEL_SHIFT;MIGRATION;NATIVE_RENEWAL;NEW
	OfferDealType *string `json:"offerDealType,omitempty"`

	// Optional. A title that describes the offer and helps your customers
	// identify it. This title will be visible to the customer. Maximum length:
	// 256 characters. Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.title
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty"`

	// Optional. Unstructured text content that is visible to the customer.
	// Maximum length: 120 characters.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.customer_note
	// +kubebuilder:validation:Optional
	CustomerNote *string `json:"customerNote,omitempty"`

	// Optional. Information about the partner contact.
	// Must be provided when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.partner_contact
	// +kubebuilder:validation:Optional
	PartnerContact *PrivateOfferPartnerContact `json:"partnerContact,omitempty"`

	// Optional. Information identifying the intended recipient of the offer.
	// Must be provided when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.customer
	// +kubebuilder:validation:Optional
	Customer *PrivateOfferCustomer `json:"customer,omitempty"`

	// Optional. Deadline for acceptance of published offers.
	// A published offer not accepted by this time will expire.
	// Only day boundaries in the America/Los_Angeles time zone are supported.
	// Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.accept_deadline_time
	// +kubebuilder:validation:Optional
	AcceptDeadlineTime *DateTime `json:"acceptDeadlineTime,omitempty"`

	// Optional. Configuration for the offer term.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.term
	// +kubebuilder:validation:Optional
	Term *PrivateOfferTerm `json:"term,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer
type PrivateOfferCustomer struct {
	// Optional. A string identifying the customer's entity (for example, the
	// customer's organization or company name). Must be provided when
	// publishing the offer. Maximum length: 256 characters.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer.entity_title
	EntityTitle *string `json:"entityTitle,omitempty"`

	// Optional. A string identifying the customer contact.
	// Must be provided when publishing the offer.
	// Maximum length: 256 characters.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer.contact
	Contact *string `json:"contact,omitempty"`

	// Optional. Email of customer contact.
	// If provided, it must be a well-formed email address.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer.email
	Email *string `json:"email,omitempty"`

	// Optional. The customer's billing account targeted by the offer.
	// The private offer once published can be accepted by a billing
	// administrator of the target billing account. If the customer accepts the
	// offer and later moves the resulting order to a new billing account, this
	// field will continue to reflect the original billing account to which the
	// private offer was extended. Must be provided when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer.target_billing_account
	TargetBillingAccountRef *billingv1alpha1.BillingAccountRef `json:"targetBillingAccountRef,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer
type PrivateOfferCustomerObservedState struct {
	// Output only. Legal address of the customer organization.
	// This field can no longer be set, but it is preserved to return the
	// data from existing offers where address is set.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer.address
	Address *string `json:"address,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.PartnerContact
type PrivateOfferPartnerContact struct {
	// Optional. Display text identifying the partner sales contact.
	// Displayed to the customer.
	// Must be provided when publishing the offer.
	// Maximum length: 256 characters.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.PartnerContact.contact
	Contact *string `json:"contact,omitempty"`

	// Optional. The email address of the partner sales contact.
	// Displayed to the customer.
	// The format of the provided email address is validated when publishing the
	// offer, but no verification is performed that the email address actually
	// exists, accepts email, or is otherwise functional.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.PartnerContact.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer
type PrivateOfferSingleProductOffer struct {
	// Optional. An existing private offer that will be superseded by this offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.amended_private_offer
	AmendedPrivateOfferRef *CommerceProducerPrivateOfferRef `json:"amendedPrivateOfferRef,omitempty"`

	// Optional. An existing standard offer that will be superseded by this offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.amended_standard_offer
	AmendedStandardOfferRef *CommerceProducerStandardOfferRef `json:"amendedStandardOfferRef,omitempty"`

	// Optional. Price configurations for offers with standard intervals.
	// A price must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.standard_interval_price
	StandardIntervalPrice *PrivateOfferSingleProductOfferStandardIntervalPrice `json:"standardIntervalPrice,omitempty"`

	// Optional. Price configurations for offers with custom intervals.
	// Custom interval corresponds to "custom billing frequency".
	// A price must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.custom_interval_price
	CustomIntervalPrice *PrivateOfferSingleProductOfferCustomIntervalPrice `json:"customIntervalPrice,omitempty"`

	// Optional. The StandardOffer this PrivateOffer is based on.
	// Must be in the same project as the private offer, and must be effective
	// at the time of publishing. Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.base_standard_offer
	BaseStandardOfferRef *CommerceProducerStandardOfferRef `json:"baseStandardOfferRef,omitempty"`

	// Optional. The custom product features to display for this offer.
	// Feature `display_name` values must be unique to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.features
	Features []PrivateOfferSingleProductOfferFeature `json:"features,omitempty"`

	// Optional. Additional contract value that the customer is legally
	// obligated to spend on the product over the duration of the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.additional_contract_value
	AdditionalContractValue *PrivateOfferSingleProductOfferAdditionalContractValue `json:"additionalContractValue,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.StandardIntervalPrice
type PrivateOfferSingleProductOfferStandardIntervalPrice struct {
	// Optional. The standard interval of the offer.
	// Must be present to publish unless the price model is usage-only.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.StandardIntervalPrice.standard_interval
	StandardInterval *string `json:"standardInterval,omitempty"`

	// Optional. The price model of the offer.
	// Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.StandardIntervalPrice.price_model
	PriceModel *PrivateOfferSingleProductOfferPriceModel `json:"priceModel,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.CustomIntervalPrice
type PrivateOfferSingleProductOfferCustomIntervalPrice struct {
	// Optional. The installments that make up the installment timeline.
	// All installments must have the same form of price model (e.g.
	// all commitment, or all flat fee). A subscription must be present
	// under the price model for every installment.
	//
	// Must contain at least one installment to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.CustomIntervalPrice.installments
	Installments []PrivateOfferSingleProductOfferInstallment `json:"installments,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Feature
type PrivateOfferSingleProductOfferFeature struct {
	// Optional. Human readable display text characterizing the feature.
	// Should be sufficiently detailed to identify the feature, allowing
	// features to be correlated across separate offers and service levels.
	// Must be non-empty to publish the offer.
	// The maximum allowed length is 128 characters.
	// Allows characters from the following Unicode Property Classes:
	// Letters, Numbers, Punctuation, Symbols, and Separators
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Feature.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Human readable display text reflecting the value of the
	// feature. Used for variable features not captured by the display_name
	// alone. The maximum allowed length is 3000 characters. Allows characters
	// from the following Unicode Property Classes: Letters, Numbers,
	// Punctuation, Symbols, and Separators
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Feature.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment
type PrivateOfferSingleProductOfferInstallment struct {
	// Optional. Start time of the installment.
	// Must be set when publishing the offer.
	// Only day boundaries in the America/Los_Angeles time zone are supported.
	// Must not be in the past at publish time.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.start_time
	StartTime *DateTime `json:"startTime,omitempty"`

	// Optional. Price model of the installment.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.price_model
	PriceModel *PrivateOfferSingleProductOfferPriceModel `json:"priceModel,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel
type PrivateOfferSingleProductOfferPriceModel struct {
	// Optional. Price model for a flat fee (subscription).
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.flat_fee
	FlatFee *PrivateOfferSingleProductOfferPriceModelFlatFee `json:"flatFee,omitempty"`

	// Optional. Price model for usage based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.usage
	Usage *PrivateOfferSingleProductOfferPriceModelUsage `json:"usage,omitempty"`

	// Optional. Price model for commitment based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.commitment
	Commitment *PrivateOfferSingleProductOfferPriceModelCommitment `json:"commitment,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Commitment
type PrivateOfferSingleProductOfferPriceModelCommitment struct {
	// Optional. The commitment balance that the customer will receive.
	// Value cannot be negative.
	// The currency must be "USD" and precision is limited to cents.
	// Must be present to publish the offer, if the parent message is set.
	// The maximum allowed value is 1,000,000,000 USD.
	//
	// If the current offer amends a private offer where the term has not
	// ended, and if the field
	// `single_product_offer.standard_interval_price.price_model.commitment.commitment_amount`
	// is set on both the current offer and the amended private offer, then
	// the commitment amount must be equal or higher on the current offer,
	// compared to the amended private offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Commitment.commitment_amount
	CommitmentAmount *Money `json:"commitmentAmount,omitempty"`

	// Optional. The discount percent on `commitment_amount`.
	//
	// For example, `10` means a discount of 10%. If the original
	// `commitment_amount` is $100 then the discounted amount will be $90.
	// The customer is charged $90 and receives $100 in credits.
	//
	// All reported usage will be charged at the standard price.
	// If this field is set, the `price_model.usage` must be unset. A
	// private offer can apply a discount to all usage or to a usage
	// commitment, but not both.
	//
	// Must be between 0 and 100 inclusive, with precision up to 2 decimal
	// places. If this field is set, then the allowed pattern for 'value'
	// is `^([0-9]{1,2}(\\.\\d{1,2})?|100(\\.0{1,2})?)$`.
	// Examples of valid values: "0", "100.0", "12.34", "40.0".
	// Example of invalid values: ".3", "-1", "100.1", "12.345", "+12", "".
	//
	// Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Commitment.discount_percent
	DiscountPercent *Decimal `json:"discountPercent,omitempty"`

	// Optional. Additional credit granted to the customer.
	// Additional credits are only supported for custom interval offers.
	// Equivalent behavior can be achieved for standard interval offers by
	// varying the `commitment_amount` and `discount_percent`.
	//
	// Value cannot be negative.
	// The currency must be "USD" and precision is limited to cents.
	// The maximum allowed value is 1,000,000 USD.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Commitment.additional_credit
	AdditionalCredit *Money `json:"additionalCredit,omitempty"`

	// Optional. Whether to discard the previous credit balance when the
	// associated installment starts. If not set, the previous credit
	// balance will be rolled over to the current installment.
	//
	// If there are no previous installments, then the value of this field
	// will not matter - since it has no effect when no credits exist.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Commitment.discard_previous_credit_balance
	DiscardPreviousCreditBalance *bool `json:"discardPreviousCreditBalance,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.FlatFee
type PrivateOfferSingleProductOfferPriceModelFlatFee struct {
	// Optional. Flat fee overriding the default flat fee in the base
	// standard offer. Value cannot be negative. The currency must be "USD"
	// and precision is limited to cents. Must be present to publish the
	// offer, if the parent message is set. The maximum allowed value is
	// 1,000,000,000 USD.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.FlatFee.flat_fee_override
	FlatFeeOverride *Money `json:"flatFeeOverride,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage
type PrivateOfferSingleProductOfferPriceModelUsage struct {
	// Optional. The default discount percentage applied to all SKUs in the
	// product that are not explicitly discounted below.
	// If unset, default discount percentage is 0.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage.default_discount_percent
	DefaultDiscountPercent *Decimal `json:"defaultDiscountPercent,omitempty"`

	// Optional. Specific discounts for individual SKUs in the product.
	// If a discount is not specified for a SKU, `default_discount_percent`
	// will be applied.
	// At most 100 discounts may be specified.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage.sku_discounts
	SkuDiscounts []PrivateOfferSingleProductOfferPriceModelSkuDiscount `json:"skuDiscounts,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount
type PrivateOfferSingleProductOfferPriceModelSkuDiscount struct {
	// Optional. The SKU to which the discount will be applied.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.sku
	SkuRef *CommerceProducerSkuRef `json:"skuRef,omitempty"`

	// Optional. The name of a Google SKU from the Cloud Billing API.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.cloud_billing_sku
	CloudBillingSku *string `json:"cloudBillingSku,omitempty"`

	// Optional. The discount percent for the SKU.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.discount_percent
	DiscountPercent *Decimal `json:"discountPercent,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.AdditionalContractValue
type PrivateOfferSingleProductOfferAdditionalContractValue struct {
	// Optional. The absolute, cumulative contract value of the customer's
	// spend obligation that is added on top of the automatically billed fees
	// from Google.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.AdditionalContractValue.contract_value
	ContractValue *Money `json:"contractValue,omitempty"`

	// Optional. The resource names of the SKUs whose tracked usage is
	// eligible to contribute toward satisfying this additional contract value
	// obligation.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.AdditionalContractValue.eligible_skus
	EligibleSkuRefs []CommerceProducerSkuRef `json:"eligibleSkuRefs,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.Term
type PrivateOfferTerm struct {
	// Optional. Duration in months of the offer term.
	// The offer will end on the same day of the month after this duration. If
	// that date is not available, the offer will end on the last day of the
	// month. For example, if the offer starts on 01/31/2024 and the duration
	// is 13 months, the offer will end on 02/28/2025.
	// Must be set to publish the offer if the end policy is `AFTER_DURATION`
	// or `MATCH_AMENDED_OFFER`, and must be a positive value.
	//
	// The maximum possible offer duration, including all renewal terms,
	// cannot exceed 7 years. For example, if `duration_months` is 9 and
	// `term.max_renewal_count` is 9, then the maximum possible offer duration
	// is 90 months and exceeds the 7 year limit.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.duration_months
	DurationMonths *int32 `json:"durationMonths,omitempty"`

	// Optional. Specified end time of the offer.
	// Must be set to publish the offer if the end policy is
	// `SCHEDULED_END_TIME`.
	// When set, the time must be in the future and at a day boundary in the
	// America/Los_Angeles time zone, and must be after
	// `term.scheduled_start_time` if applicable.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.scheduled_end_time
	ScheduledEndTime *DateTime `json:"scheduledEndTime,omitempty"`

	// Optional. The max number of renewals allowed, after the initial term
	// ends. This field is only relevant to standard interval offers with
	// interval of 'MONTHLY_PRORATED', 'MONTHLY_NOT_PRORATED',
	// 'QUARTERLY_NOT_PRORATED' or 'YEARLY_NOT_PRORATED'. Other types of
	// offers cannot be renewed.
	//
	// For example, if the initial term is 1 year and the max renewals is 3,
	// the customer can renew the offer for up to 3 additional years after the
	// initial term ends. Customers control whether they renew the offer.
	//
	// Allowing a customer to renew an offer does not constitute a guarantee
	// to the customer that the offered product will still be available to be
	// renewed at the renewal date.
	//
	// An offer with a max renewal count of 0 cannot be renewed.
	// Must be non-negative.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.max_renewal_count
	MaxRenewalCount *int32 `json:"maxRenewalCount,omitempty"`

	// Optional. Defines when the resulting order should start.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.start_policy
	StartPolicy *string `json:"startPolicy,omitempty"`

	// Optional. The scheduled start time of the offer. Cannot exceed one year
	// from the time of publish.
	//
	// If the start policy is `SCHEDULED_START_TIME`, a future time at a day
	// boundary in the America/Los_Angeles time zone must be provided when
	// publishing the offer.
	//
	// If the offer amends another private or standard offer, then the
	// scheduled start time must be:
	//
	// * Before or equal to the amended offer's end time if the offer's
	//   end policy is not MATCH_AMENDED_OFFER. Otherwise, it has to be
	//   strictly before the amended offer's end time.
	// * Before or equal to the start time of the next pending installment of
	//   the amended offer, if amending a private offer with custom
	//   installments. For example, if the amended offer has installments on
	//   01/01, 02/01, 03/01, 04/01, and today is 02/15, then the new offer's
	//   `scheduled_start_time` must be at or before 03/01.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.scheduled_start_time
	ScheduledStartTime *DateTime `json:"scheduledStartTime,omitempty"`

	// Optional. Defines when an offer should end.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.end_policy
	EndPolicy *string `json:"endPolicy,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer
type PrivateOfferSingleProductOfferObservedState struct {
	// Output only. The service level (also known as the 'plan') of the base
	// standard offer. The value is populated at publish time from the base
	// standard offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.service_level
	ServiceLevel *string `json:"serviceLevel,omitempty"`

	// Output only. Present for offers created by a reseller from a reseller
	// private offer plan (RPOP). When set, contains the ID of the originating
	// RPOP. Not included for `PRIVATE_OFFER_VIEW_BASIC`.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.reseller_private_offer_plan_id
	ResellerPrivateOfferPlanID *string `json:"resellerPrivateOfferPlanID,omitempty"`

	// Output only. The effective installment timeline of the offer.
	// Not included for `PRIVATE_OFFER_VIEW_BASIC`.
	// Included for `PRIVATE_OFFER_VIEW_FULL` if all necessary information is
	// available to generate the timeline, and if the offer has
	// 'standard_interval_price' of 'MONTHLY_PRORATED', 'MONTHLY_NOT_PRORATED',
	// 'QUARTERLY_NOT_PRORATED', or 'YEARLY_NOT_PRORATED'.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.effective_installment_timeline
	EffectiveInstallmentTimeline []PrivateOfferSingleProductOfferInstallmentObservedState `json:"effectiveInstallmentTimeline,omitempty"`

	// Output only. Contract value of the offer.
	// Not included for `PRIVATE_OFFER_VIEW_BASIC`.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.contract_value
	ContractValue *PrivateOfferSingleProductOfferContractValueObservedState `json:"contractValue,omitempty"`

	// Output only. Revenue share information for this Private Offer.
	// Not included for `PRIVATE_OFFER_VIEW_BASIC`.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.revenue_share
	RevenueShare *PrivateOfferSingleProductOfferRevenueShareObservedState `json:"revenueShare,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment
type PrivateOfferSingleProductOfferInstallmentObservedState struct {
	// Output only. Start time of the installment.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.start_time
	StartTime *DateTime `json:"startTime,omitempty"`

	// Output only. Price model of the installment.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.price_model
	PriceModel *PrivateOfferSingleProductOfferPriceModelObservedState `json:"priceModel,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel
type PrivateOfferSingleProductOfferPriceModelObservedState struct {
	// Output only. Price model for a flat fee (subscription).
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.flat_fee
	FlatFee *PrivateOfferSingleProductOfferPriceModelFlatFee `json:"flatFee,omitempty"`

	// Output only. Price model for usage based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.usage
	Usage *PrivateOfferSingleProductOfferPriceModelUsageObservedState `json:"usage,omitempty"`

	// Output only. Price model for commitment based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.commitment
	Commitment *PrivateOfferSingleProductOfferPriceModelCommitment `json:"commitment,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage
type PrivateOfferSingleProductOfferPriceModelUsageObservedState struct {
	// Output only. The default discount percentage applied to all SKUs in the
	// product that are not explicitly discounted below.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage.default_discount_percent
	DefaultDiscountPercent *Decimal `json:"defaultDiscountPercent,omitempty"`

	// Output only. Specific discounts for individual SKUs in the product.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage.sku_discounts
	SkuDiscounts []PrivateOfferSingleProductOfferPriceModelSkuDiscountObservedState `json:"skuDiscounts,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount
type PrivateOfferSingleProductOfferPriceModelSkuDiscountObservedState struct {
	// Output only. The name of a partner SKU from this API.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.sku
	Sku *string `json:"sku,omitempty"`

	// Output only. The name of a Google SKU from the Cloud Billing API.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.cloud_billing_sku
	CloudBillingSku *string `json:"cloudBillingSku,omitempty"`

	// Output only. The discount percent for the SKU.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.discount_percent
	DiscountPercent *Decimal `json:"discountPercent,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.ContractValue
type PrivateOfferSingleProductOfferContractValueObservedState struct {
	// Output only. The total contract value of the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.ContractValue.total_contract_value
	TotalContractValue *Money `json:"totalContractValue,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.RevenueShare
type PrivateOfferSingleProductOfferRevenueShareObservedState struct {
	// Output only. The revenue share currently in effect.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.RevenueShare.current_term_vendor_net_revenue_percent
	CurrentTermVendorNetRevenuePercent *Decimal `json:"currentTermVendorNetRevenuePercent,omitempty"`

	// Output only. The expected revenue share for the renewal term.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.RevenueShare.renewal_term_vendor_net_revenue_percent
	RenewalTermVendorNetRevenuePercent *Decimal `json:"renewalTermVendorNetRevenuePercent,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.PublishRequirementGoogleReview
type PrivateOfferPublishRequirementGoogleReviewObservedState struct {
	// Output only. The approval time of the Google review process.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.PublishRequirementGoogleReview.review_approve_time
	ReviewApproveTime *string `json:"reviewApproveTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.ResellerContact
type PrivateOfferResellerContactObservedState struct {
	// Output only. The name of the reseller contact for this offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.ResellerContact.contact
	Contact *string `json:"contact,omitempty"`

	// Output only. The email of the reseller contact for this offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.ResellerContact.email
	Email *string `json:"email,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.Term
type PrivateOfferTermObservedState struct {
	// Output only. Deprecated: This field can only be true on existing offers
	// with standard interval of 'MONTHLY_POSTPAY'. As of May 2024, we no
	// longer allow unlimited renewals.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.unlimited_renewal
	UnlimitedRenewal *bool `json:"unlimitedRenewal,omitempty"`

	// Output only. The expected end time of the current offer term.
	//
	// At the end of each offer term, an offer associated with an active
	// order will either renew or end. When an offer renews, a new term begins
	// and this value changes to reflect the end time of the new term.
	// When an offer ends this value is no longer set and instead `end_time` is
	// set.
	//
	// When the term of an offer ends and the offer does not renew, the
	// associated order also ends if the base standard offer has a
	// subscription price model. Otherwise, the associated order does not end
	// and remains active.
	//
	// Not included for PRIVATE_OFFER_VIEW_BASIC.
	// Included for PRIVATE_OFFER_VIEW_FULL when the offer has not ended and
	// either the offer has started or the value can be derived from other
	// fields.
	//
	// For offers that have not started, this field is set when one of the
	// following conditions is true.
	//
	// * The offer sets `term.scheduled_end_time`.
	// * The offer sets `term.scheduled_start_time` and `term.duration_months`,
	//   the offer's `term.end_policy` is not `MATCH_AMENDED_OFFER`, and the
	//   offer does not have standard_interval set to `MONTHLY_POSTPAY`.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Term.effective_term_end_time
	EffectiveTermEndTime *DateTime `json:"effectiveTermEndTime,omitempty"`
}

// CommerceProducerPrivateOfferStatus defines the config connector machine state of CommerceProducerPrivateOffer
type CommerceProducerPrivateOfferStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CommerceProducerPrivateOffer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CommerceProducerPrivateOfferObservedState `json:"observedState,omitempty"`
}

// CommerceProducerPrivateOfferObservedState is the state of the CommerceProducerPrivateOffer resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer
type CommerceProducerPrivateOfferObservedState struct {
	// Optional. Configurations for the offer that is associated with a single
	// product.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.single_product_offer
	SingleProductOffer *PrivateOfferSingleProductOfferObservedState `json:"singleProductOffer,omitempty"`

	// Output only. The state of the private offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.state
	State *string `json:"state,omitempty"`

	// Output only. Information about the Google review process.
	// Present only when the offer is determined to require review by Google.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.publish_requirement_google_review
	PublishRequirementGoogleReview *PrivateOfferPublishRequirementGoogleReviewObservedState `json:"publishRequirementGoogleReview,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time the offer transitioned to PUBLISHED state.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.publish_time
	PublishTime *string `json:"publishTime,omitempty"`

	// Output only. The time the offer transitioned to ACCEPTED state.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.accept_time
	AcceptTime *string `json:"acceptTime,omitempty"`

	// Output only. The time the offer transited to CANCELLED state.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.cancel_time
	CancelTime *string `json:"cancelTime,omitempty"`

	// Output only. The time when the offer ended. This can only be set for offers
	// with `ENDED` state.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Internal note supplied when the offer was cancelled.
	// Present only for cancelled offers and only if a note was supplied.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.cancellation_note
	CancellationNote *string `json:"cancellationNote,omitempty"`

	// Output only. Information about the reseller contact.
	// Present only for offers created by a reseller.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.reseller_contact
	ResellerContact *PrivateOfferResellerContactObservedState `json:"resellerContact,omitempty"`

	// Optional. Information identifying the intended recipient of the offer.
	// Must be provided when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.customer
	Customer *PrivateOfferCustomerObservedState `json:"customer,omitempty"`

	// Optional. Configuration for the offer term.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.term
	Term *PrivateOfferTermObservedState `json:"term,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcommerceproducerprivateoffer;gcpcommerceproducerprivateoffers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CommerceProducerPrivateOffer is the Schema for the CommerceProducerPrivateOffer API
// +k8s:openapi-gen=true
type CommerceProducerPrivateOffer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CommerceProducerPrivateOfferSpec   `json:"spec,omitempty"`
	Status CommerceProducerPrivateOfferStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CommerceProducerPrivateOfferList contains a list of CommerceProducerPrivateOffer
type CommerceProducerPrivateOfferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CommerceProducerPrivateOffer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CommerceProducerPrivateOffer{}, &CommerceProducerPrivateOfferList{})
}
