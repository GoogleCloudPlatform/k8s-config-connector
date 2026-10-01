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
	SingleProductOffer *PrivateOffer_SingleProductOffer `json:"singleProductOffer,omitempty"`

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
	PartnerContact *PrivateOffer_PartnerContact `json:"partnerContact,omitempty"`

	// Optional. Information identifying the intended recipient of the offer.
	// Must be provided when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.customer
	// +kubebuilder:validation:Optional
	Customer *PrivateOffer_Customer `json:"customer,omitempty"`

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
	Term *PrivateOffer_Term `json:"term,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer
type PrivateOffer_Customer struct {
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
type PrivateOffer_CustomerObservedState struct {
	// Output only. Legal address of the customer organization.
	// This field can no longer be set, but it is preserved to return the
	// data from existing offers where address is set.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.Customer.address
	Address *string `json:"address,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer
type PrivateOffer_SingleProductOffer struct {
	// Optional. An existing private offer that will be superseded by this offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.amended_private_offer
	AmendedPrivateOfferRef *CommerceProducerPrivateOfferRef `json:"amendedPrivateOfferRef,omitempty"`

	// Optional. An existing standard offer that will be superseded by this offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.amended_standard_offer
	AmendedStandardOfferRef *CommerceProducerStandardOfferRef `json:"amendedStandardOfferRef,omitempty"`

	// Optional. Price configurations for offers with standard intervals.
	// A price must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.standard_interval_price
	StandardIntervalPrice *PrivateOffer_SingleProductOffer_StandardIntervalPrice `json:"standardIntervalPrice,omitempty"`

	// Optional. Price configurations for offers with custom intervals.
	// Custom interval corresponds to "custom billing frequency".
	// A price must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.custom_interval_price
	CustomIntervalPrice *PrivateOffer_SingleProductOffer_CustomIntervalPrice `json:"customIntervalPrice,omitempty"`

	// Optional. The StandardOffer this PrivateOffer is based on.
	// Must be in the same project as the private offer, and must be effective
	// at the time of publishing. Must be present to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.base_standard_offer
	BaseStandardOfferRef *CommerceProducerStandardOfferRef `json:"baseStandardOfferRef,omitempty"`

	// Optional. The custom product features to display for this offer.
	// Feature `display_name` values must be unique to publish the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.features
	Features []PrivateOffer_SingleProductOffer_Feature `json:"features,omitempty"`

	// Optional. Additional contract value that the customer is legally
	// obligated to spend on the product over the duration of the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.additional_contract_value
	AdditionalContractValue *PrivateOffer_SingleProductOffer_AdditionalContractValue `json:"additionalContractValue,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment
type PrivateOffer_SingleProductOffer_Installment struct {
	// Optional. Start time of the installment.
	// Must be set when publishing the offer.
	// Only day boundaries in the America/Los_Angeles time zone are supported.
	// Must not be in the past at publish time.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.start_time
	StartTime *DateTime `json:"startTime,omitempty"`

	// Optional. Price model of the installment.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.price_model
	PriceModel *PrivateOffer_SingleProductOffer_PriceModel `json:"priceModel,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel
type PrivateOffer_SingleProductOffer_PriceModel struct {
	// Optional. Price model for a flat fee (subscription).
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.flat_fee
	FlatFee *PrivateOffer_SingleProductOffer_PriceModel_FlatFee `json:"flatFee,omitempty"`

	// Optional. Price model for usage based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.usage
	Usage *PrivateOffer_SingleProductOffer_PriceModel_Usage `json:"usage,omitempty"`

	// Optional. Price model for commitment based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.commitment
	Commitment *PrivateOffer_SingleProductOffer_PriceModel_Commitment `json:"commitment,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage
type PrivateOffer_SingleProductOffer_PriceModel_Usage struct {
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
	SkuDiscounts []PrivateOffer_SingleProductOffer_PriceModel_SkuDiscount `json:"skuDiscounts,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer
type PrivateOffer_SingleProductOfferObservedState struct {
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
	EffectiveInstallmentTimeline []PrivateOffer_SingleProductOffer_InstallmentObservedState `json:"effectiveInstallmentTimeline,omitempty"`

	// Output only. Contract value of the offer.
	// Not included for `PRIVATE_OFFER_VIEW_BASIC`.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.contract_value
	ContractValue *PrivateOffer_SingleProductOffer_ContractValueObservedState `json:"contractValue,omitempty"`

	// Output only. Revenue share information for this Private Offer.
	// Not included for `PRIVATE_OFFER_VIEW_BASIC`.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.revenue_share
	RevenueShare *PrivateOffer_SingleProductOffer_RevenueShareObservedState `json:"revenueShare,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment
type PrivateOffer_SingleProductOffer_InstallmentObservedState struct {
	// Output only. Start time of the installment.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.start_time
	StartTime *DateTime `json:"startTime,omitempty"`

	// Output only. Price model of the installment.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.Installment.price_model
	PriceModel *PrivateOffer_SingleProductOffer_PriceModelObservedState `json:"priceModel,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel
type PrivateOffer_SingleProductOffer_PriceModelObservedState struct {
	// Output only. Price model for a flat fee (subscription).
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.flat_fee
	FlatFee *PrivateOffer_SingleProductOffer_PriceModel_FlatFee `json:"flatFee,omitempty"`

	// Output only. Price model for usage based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.usage
	Usage *PrivateOffer_SingleProductOffer_PriceModel_UsageObservedState `json:"usage,omitempty"`

	// Output only. Price model for commitment based pricing.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.commitment
	Commitment *PrivateOffer_SingleProductOffer_PriceModel_Commitment `json:"commitment,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage
type PrivateOffer_SingleProductOffer_PriceModel_UsageObservedState struct {
	// Output only. The default discount percentage applied to all SKUs in the
	// product that are not explicitly discounted below.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage.default_discount_percent
	DefaultDiscountPercent *Decimal `json:"defaultDiscountPercent,omitempty"`

	// Output only. Specific discounts for individual SKUs in the product.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.Usage.sku_discounts
	SkuDiscounts []PrivateOffer_SingleProductOffer_PriceModel_SkuDiscountObservedState `json:"skuDiscounts,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount
type PrivateOffer_SingleProductOffer_PriceModel_SkuDiscountObservedState struct {
	// Output only. The SKU to which the discount will be applied.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.sku
	Sku *string `json:"sku,omitempty"`

	// Output only. The discount as a percentage of the list price of the SKU.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.discount_percentage
	DiscountPercentage *Decimal `json:"discountPercentage,omitempty"`

	// Output only. The fixed discounted price of the SKU.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.discounted_price
	DiscountedPrice *Money `json:"discountedPrice,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.ContractValue
type PrivateOffer_SingleProductOffer_ContractValueObservedState struct {
	// Output only. The total contract value of the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.ContractValue.total_contract_value
	TotalContractValue *Money `json:"totalContractValue,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.RevenueShare
type PrivateOffer_SingleProductOffer_RevenueShareObservedState struct {
	// Output only. The revenue share currently in effect.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.RevenueShare.current_term_vendor_net_revenue_percent
	CurrentTermVendorNetRevenuePercent *Decimal `json:"currentTermVendorNetRevenuePercent,omitempty"`

	// Output only. The expected revenue share for the renewal term.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.RevenueShare.renewal_term_vendor_net_revenue_percent
	RenewalTermVendorNetRevenuePercent *Decimal `json:"renewalTermVendorNetRevenuePercent,omitempty"`
}

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.AdditionalContractValue
type PrivateOffer_SingleProductOffer_AdditionalContractValue struct {
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

// +kcc:proto=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount
type PrivateOffer_SingleProductOffer_PriceModel_SkuDiscount struct {
	// Optional. The SKU to which the discount will be applied.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.sku
	SkuRef *CommerceProducerSkuRef `json:"skuRef,omitempty"`

	// Optional. The discount as a percentage of the list price of the SKU.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.discount_percentage
	DiscountPercentage *Decimal `json:"discountPercentage,omitempty"`

	// Optional. The fixed discounted price of the SKU.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.SingleProductOffer.PriceModel.SkuDiscount.discounted_price
	DiscountedPrice *Money `json:"discountedPrice,omitempty"`
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
	SingleProductOffer *PrivateOffer_SingleProductOfferObservedState `json:"singleProductOffer,omitempty"`

	// Output only. The state of the private offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.state
	State *string `json:"state,omitempty"`

	// Output only. Information about the Google review process.
	// Present only when the offer is determined to require review by Google.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.publish_requirement_google_review
	PublishRequirementGoogleReview *PrivateOffer_PublishRequirementGoogleReviewObservedState `json:"publishRequirementGoogleReview,omitempty"`

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
	ResellerContact *PrivateOffer_ResellerContactObservedState `json:"resellerContact,omitempty"`

	// Optional. Information identifying the intended recipient of the offer.
	// Must be provided when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.customer
	Customer *PrivateOffer_CustomerObservedState `json:"customer,omitempty"`

	// Optional. Configuration for the offer term.
	// Must be set when publishing the offer.
	// +kcc:proto:field=google.cloud.commerceproducer.v1beta.PrivateOffer.term
	Term *PrivateOffer_TermObservedState `json:"term,omitempty"`
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
