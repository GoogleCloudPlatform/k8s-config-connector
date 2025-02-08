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


// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.AccountDefenderAssessment
type AccountDefenderAssessment struct {
	// Labels for this request.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.AccountDefenderAssessment.labels
	Labels []string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.Assessment
type Assessment struct {

	// The event being assessed.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.event
	Event *Event `json:"event,omitempty"`

	// Information about the user's credentials used to check for leaks.
	//  This feature is part of the Early Access Program (EAP). Exercise caution,
	//  and do not deploy integrations based on this feature in a production
	//  environment.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.password_leak_verification
	PasswordLeakVerification *PasswordLeakVerification `json:"passwordLeakVerification,omitempty"`

	// Assessment returned by account defender when a hashed_account_id is
	//  provided.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.account_defender_assessment
	AccountDefenderAssessment *AccountDefenderAssessment `json:"accountDefenderAssessment,omitempty"`

	// Assessment returned by Fraud Prevention when TransactionData is provided.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.fraud_prevention_assessment
	FraudPreventionAssessment *FraudPreventionAssessment `json:"fraudPreventionAssessment,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.Event
type Event struct {
	// Optional. The user response token provided by the reCAPTCHA client-side
	//  integration on your site.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.token
	Token *string `json:"token,omitempty"`

	// Optional. The site key that was used to invoke reCAPTCHA on your site and
	//  generate the token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.site_key
	SiteKey *string `json:"siteKey,omitempty"`

	// Optional. The user agent present in the request from the user's device
	//  related to this event.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.user_agent
	UserAgent *string `json:"userAgent,omitempty"`

	// Optional. The IP address in the request from the user's device related to
	//  this event.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.user_ip_address
	UserIPAddress *string `json:"userIPAddress,omitempty"`

	// Optional. The expected action for this type of event. This should be the
	//  same action provided at token generation time on client-side platforms
	//  already integrated with reCAPTCHA.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.expected_action
	ExpectedAction *string `json:"expectedAction,omitempty"`

	// Optional. Unique stable hashed user identifier for the request. The
	//  identifier must be hashed using hmac-sha256 with stable secret.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.hashed_account_id
	HashedAccountID []byte `json:"hashedAccountID,omitempty"`

	// Optional. Data describing a payment transaction to be assessed. Sending
	//  this data enables reCAPTCHA Fraud Prevention and the
	//  FraudPreventionAssessment component in the response.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.transaction_data
	TransactionData *TransactionData `json:"transactionData,omitempty"`

	// Optional. The Fraud Prevention setting for this Assessment.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Event.fraud_prevention
	FraudPrevention *string `json:"fraudPrevention,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment
type FraudPreventionAssessment struct {
	// Probability (0-1) of this transaction being fraudulent. Summarizes the
	//  combined risk of attack vectors below.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.transaction_risk
	TransactionRisk *float32 `json:"transactionRisk,omitempty"`

	// Assessment of this transaction for risk of a stolen instrument.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.stolen_instrument_verdict
	StolenInstrumentVerdict *FraudPreventionAssessment_StolenInstrumentVerdict `json:"stolenInstrumentVerdict,omitempty"`

	// Assessment of this transaction for risk of being part of a card testing
	//  attack.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.card_testing_verdict
	CardTestingVerdict *FraudPreventionAssessment_CardTestingVerdict `json:"cardTestingVerdict,omitempty"`

	// Assessment of this transaction for behavioral trust.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.behavioral_trust_verdict
	BehavioralTrustVerdict *FraudPreventionAssessment_BehavioralTrustVerdict `json:"behavioralTrustVerdict,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.BehavioralTrustVerdict
type FraudPreventionAssessment_BehavioralTrustVerdict struct {
	// Probability (0-1) of this transaction attempt being executed in a
	//  behaviorally trustworthy way.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.BehavioralTrustVerdict.trust
	Trust *float32 `json:"trust,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.CardTestingVerdict
type FraudPreventionAssessment_CardTestingVerdict struct {
	// Probability (0-1) of this transaction attempt being part of a card
	//  testing attack.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.CardTestingVerdict.risk
	Risk *float32 `json:"risk,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.StolenInstrumentVerdict
type FraudPreventionAssessment_StolenInstrumentVerdict struct {
	// Probability (0-1) of this transaction being executed with a stolen
	//  instrument.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.FraudPreventionAssessment.StolenInstrumentVerdict.risk
	Risk *float32 `json:"risk,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.PasswordLeakVerification
type PasswordLeakVerification struct {
	// Optional. Scrypt hash of the username+password that the customer wants to
	//  verify against a known password leak.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.PasswordLeakVerification.hashed_user_credentials
	HashedUserCredentials []byte `json:"hashedUserCredentials,omitempty"`

	// Optional. The username part of the user credentials for which we want to
	//  trigger a leak check in canonicalized form. This is the same data used to
	//  create the hashed_user_credentials on the customer side.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.PasswordLeakVerification.canonicalized_username
	CanonicalizedUsername *string `json:"canonicalizedUsername,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.TokenProperties
type TokenProperties struct {
	// Whether the provided user response token is valid. When valid = false, the
	//  reason could be specified in invalid_reason or it could also be due to
	//  a user failing to solve a challenge or a sitekey mismatch (i.e the sitekey
	//  used to generate the token was different than the one specified in the
	//  assessment).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TokenProperties.valid
	Valid *bool `json:"valid,omitempty"`

	// Reason associated with the response when valid = false.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TokenProperties.invalid_reason
	InvalidReason *string `json:"invalidReason,omitempty"`

	// The timestamp corresponding to the generation of the token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TokenProperties.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The hostname of the page on which the token was generated.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TokenProperties.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Action name provided at token generation.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TokenProperties.action
	Action *string `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.TransactionData
type TransactionData struct {
	// Unique identifier for the transaction. This custom identifier can be used
	//  to reference this transaction in the future, for example, labeling a refund
	//  or chargeback event. Two attempts at the same transaction should use the
	//  same transaction id.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.transaction_id
	TransactionID *string `json:"transactionID,omitempty"`

	// The payment method for the transaction. The allowed values are:
	//
	//  * credit-card
	//  * debit-card
	//  * gift-card
	//  * processor-{name} (If a third-party is used, for example,
	//  processor-paypal)
	//  * custom-{name} (If an alternative method is used, for example,
	//  custom-crypto)
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.payment_method
	PaymentMethod *string `json:"paymentMethod,omitempty"`

	// The Bank Identification Number - generally the first 6 or 8 digits of the
	//  card.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.card_bin
	CardBin *string `json:"cardBin,omitempty"`

	// The last four digits of the card.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.card_last_four
	CardLastFour *string `json:"cardLastFour,omitempty"`

	// The currency code in ISO-4217 format.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The decimal value of the transaction in the specified currency.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.value
	Value *float64 `json:"value,omitempty"`

	// The value of shipping in the specified currency. 0 for free or no shipping.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.shipping_value
	ShippingValue *float64 `json:"shippingValue,omitempty"`

	// Destination address if this transaction involves shipping a physical item.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.shipping_address
	ShippingAddress *TransactionData_Address `json:"shippingAddress,omitempty"`

	// Address associated with the payment method when applicable.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.billing_address
	BillingAddress *TransactionData_Address `json:"billingAddress,omitempty"`

	// Information about the user paying/initiating the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.user
	User *TransactionData_User `json:"user,omitempty"`

	// Information about the user or users fulfilling the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.merchants
	Merchants []TransactionData_User `json:"merchants,omitempty"`

	// Items purchased in this transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.items
	Items []TransactionData_Item `json:"items,omitempty"`

	// Information about the payment gateway's response to the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.gateway_info
	GatewayInfo *TransactionData_GatewayInfo `json:"gatewayInfo,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address
type TransactionData_Address struct {
	// The recipient name, potentially including information such as "care of".
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address.recipient
	Recipient *string `json:"recipient,omitempty"`

	// The first lines of the address. The first line generally contains the
	//  street name and number, and further lines may include information such as
	//  an apartment number.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address.address
	Address []string `json:"address,omitempty"`

	// The town/city of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address.locality
	Locality *string `json:"locality,omitempty"`

	// The state, province, or otherwise administrative area of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address.administrative_area
	AdministrativeArea *string `json:"administrativeArea,omitempty"`

	// The CLDR country/region of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// The postal or ZIP code of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Address.postal_code
	PostalCode *string `json:"postalCode,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.TransactionData.GatewayInfo
type TransactionData_GatewayInfo struct {
	// Name of the gateway service (for example, stripe, square, paypal).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.GatewayInfo.name
	Name *string `json:"name,omitempty"`

	// Gateway response code describing the state of the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.GatewayInfo.gateway_response_code
	GatewayResponseCode *string `json:"gatewayResponseCode,omitempty"`

	// AVS response code from the gateway
	//  (available only when reCAPTCHA is called after authorization).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.GatewayInfo.avs_response_code
	AvsResponseCode *string `json:"avsResponseCode,omitempty"`

	// CVV response code from the gateway
	//  (available only when reCAPTCHA is called after authorization).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.GatewayInfo.cvv_response_code
	CvvResponseCode *string `json:"cvvResponseCode,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Item
type TransactionData_Item struct {
	// The full name of the item.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Item.name
	Name *string `json:"name,omitempty"`

	// The value per item that the user is paying, in the transaction currency,
	//  after discounts.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Item.value
	Value *float64 `json:"value,omitempty"`

	// The quantity of this item that is being purchased.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Item.quantity
	Quantity *int64 `json:"quantity,omitempty"`

	// When a merchant is specified, its corresponding account_id. Necessary to
	//  populate marketplace-style transactions.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.Item.merchant_account_id
	MerchantAccountID *string `json:"merchantAccountID,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User
type TransactionData_User struct {
	// Unique account identifier for this user. If using account defender,
	//  this should match the hashed_account_id field. Otherwise, a unique and
	//  persistent identifier for this account.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User.account_id
	AccountID *string `json:"accountID,omitempty"`

	// The epoch milliseconds of the user's account creation.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User.creation_ms
	CreationMs *int64 `json:"creationMs,omitempty"`

	// The email address of the user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User.email
	Email *string `json:"email,omitempty"`

	// Whether the email has been verified to be accessible by the user (OTP or
	//  similar).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User.email_verified
	EmailVerified *bool `json:"emailVerified,omitempty"`

	// The phone number of the user, with country code.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Whether the phone number has been verified to be accessible by the user
	//  (OTP or similar).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.TransactionData.User.phone_verified
	PhoneVerified *bool `json:"phoneVerified,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.Assessment
type AssessmentObservedState struct {
	// Output only. The resource name for the Assessment in the format
	//  `projects/{project_number}/assessments/{assessment_id}`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.name
	Name *string `json:"name,omitempty"`

	// Output only. Legitimate event score from 0.0 to 1.0.
	//  (1.0 means very likely legitimate traffic while 0.0 means very likely
	//  non-legitimate traffic).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.score
	Score *float32 `json:"score,omitempty"`

	// Output only. Properties of the provided event token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.token_properties
	TokenProperties *TokenProperties `json:"tokenProperties,omitempty"`

	// Output only. Reasons contributing to the risk analysis verdict.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.reasons
	Reasons []string `json:"reasons,omitempty"`

	// Information about the user's credentials used to check for leaks.
	//  This feature is part of the Early Access Program (EAP). Exercise caution,
	//  and do not deploy integrations based on this feature in a production
	//  environment.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.Assessment.password_leak_verification
	PasswordLeakVerification *PasswordLeakVerificationObservedState `json:"passwordLeakVerification,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1beta1.PasswordLeakVerification
type PasswordLeakVerificationObservedState struct {
	// Output only. Whether or not the user's credentials are present in a known
	//  leak.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1beta1.PasswordLeakVerification.credentials_leaked
	CredentialsLeaked *bool `json:"credentialsLeaked,omitempty"`
}
