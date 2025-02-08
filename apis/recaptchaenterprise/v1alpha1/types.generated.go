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


// +kcc:proto=google.cloud.recaptchaenterprise.v1.AccountDefenderAssessment
type AccountDefenderAssessment struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo
type AccountVerificationInfo struct {
	// Optional. Endpoints that can be used for identity verification.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo.endpoints
	Endpoints []EndpointVerificationInfo `json:"endpoints,omitempty"`

	// Optional. Language code preference for the verification message, set as a
	//  IETF BCP 47 language code.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Username of the account that is being verified. Deprecated. Customers
	//  should now provide the `account_id` field in `event.user_info`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Assessment
type Assessment struct {

	// Optional. The event being assessed.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.event
	Event *Event `json:"event,omitempty"`

	// Optional. Account verification information for identity verification. The
	//  assessment event must include a token and site key to use this feature.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.account_verification
	AccountVerification *AccountVerificationInfo `json:"accountVerification,omitempty"`

	// Optional. The private password leak verification field contains the
	//  parameters that are used to to check for leaks privately without sharing
	//  user credentials.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.private_password_leak_verification
	PrivatePasswordLeakVerification *PrivatePasswordLeakVerification `json:"privatePasswordLeakVerification,omitempty"`

	// Optional. The environment creating the assessment. This describes your
	//  environment (the system invoking CreateAssessment), NOT the environment of
	//  your user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.assessment_environment
	AssessmentEnvironment *AssessmentEnvironment `json:"assessmentEnvironment,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.AssessmentEnvironment
type AssessmentEnvironment struct {
	// Optional. Identifies the client module initiating the CreateAssessment
	//  request. This can be the link to the client module's project. Examples
	//  include:
	//  - "github.com/GoogleCloudPlatform/recaptcha-enterprise-google-tag-manager"
	//  - "cloud.google.com/recaptcha/docs/implement-waf-akamai"
	//  - "cloud.google.com/recaptcha/docs/implement-waf-cloudflare"
	//  - "wordpress.org/plugins/recaptcha-something"
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AssessmentEnvironment.client
	Client *string `json:"client,omitempty"`

	// Optional. The version of the client module. For example, "1.0.0".
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AssessmentEnvironment.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.EndpointVerificationInfo
type EndpointVerificationInfo struct {
	// Email address for which to trigger a verification request.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.EndpointVerificationInfo.email_address
	EmailAddress *string `json:"emailAddress,omitempty"`

	// Phone number for which to trigger a verification request. Should be given
	//  in E.164 format.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.EndpointVerificationInfo.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Event
type Event struct {
	// Optional. The user response token provided by the reCAPTCHA Enterprise
	//  client-side integration on your site.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.token
	Token *string `json:"token,omitempty"`

	// Optional. The site key that was used to invoke reCAPTCHA Enterprise on your
	//  site and generate the token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.site_key
	SiteKey *string `json:"siteKey,omitempty"`

	// Optional. The user agent present in the request from the user's device
	//  related to this event.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.user_agent
	UserAgent *string `json:"userAgent,omitempty"`

	// Optional. The IP address in the request from the user's device related to
	//  this event.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.user_ip_address
	UserIPAddress *string `json:"userIPAddress,omitempty"`

	// Optional. The expected action for this type of event. This should be the
	//  same action provided at token generation time on client-side platforms
	//  already integrated with recaptcha enterprise.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.expected_action
	ExpectedAction *string `json:"expectedAction,omitempty"`

	// Optional. Deprecated: use `user_info.account_id` instead.
	//  Unique stable hashed user identifier for the request. The identifier must
	//  be hashed using hmac-sha256 with stable secret.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.hashed_account_id
	HashedAccountID []byte `json:"hashedAccountID,omitempty"`

	// Optional. Flag for a reCAPTCHA express request for an assessment without a
	//  token. If enabled, `site_key` must reference an Express site key.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.express
	Express *bool `json:"express,omitempty"`

	// Optional. The URI resource the user requested that triggered an assessment.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.requested_uri
	RequestedURI *string `json:"requestedURI,omitempty"`

	// Optional. Flag for running WAF token assessment.
	//  If enabled, the token must be specified, and have been created by a
	//  WAF-enabled key.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.waf_token_assessment
	WafTokenAssessment *bool `json:"wafTokenAssessment,omitempty"`

	// Optional. JA3 fingerprint for SSL clients.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.ja3
	Ja3 *string `json:"ja3,omitempty"`

	// Optional. HTTP header information about the request.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.headers
	Headers []string `json:"headers,omitempty"`

	// Optional. Flag for enabling firewall policy config assessment.
	//  If this flag is enabled, the firewall policy is evaluated and a
	//  suggested firewall action is returned in the response.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.firewall_policy_evaluation
	FirewallPolicyEvaluation *bool `json:"firewallPolicyEvaluation,omitempty"`

	// Optional. Data describing a payment transaction to be assessed. Sending
	//  this data enables reCAPTCHA Enterprise Fraud Prevention and the
	//  FraudPreventionAssessment component in the response.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.transaction_data
	TransactionData *TransactionData `json:"transactionData,omitempty"`

	// Optional. Information about the user that generates this event, when they
	//  can be identified. They are often identified through the use of an account
	//  for logged-in requests or login/registration requests, or by providing user
	//  identifiers for guest actions like checkout.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.user_info
	UserInfo *UserInfo `json:"userInfo,omitempty"`

	// Optional. The Fraud Prevention setting for this assessment.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Event.fraud_prevention
	FraudPrevention *string `json:"fraudPrevention,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction
type FirewallAction struct {
	// The user request did not match any policy and should be allowed
	//  access to the requested resource.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.allow
	Allow *FirewallAction_AllowAction `json:"allow,omitempty"`

	// This action denies access to a given page. The user gets an HTTP
	//  error code.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.block
	Block *FirewallAction_BlockAction `json:"block,omitempty"`

	// This action injects reCAPTCHA JavaScript code into the HTML page
	//  returned by the site backend.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.include_recaptcha_script
	IncludeRecaptchaScript *FirewallAction_IncludeRecaptchaScriptAction `json:"includeRecaptchaScript,omitempty"`

	// This action redirects the request to a reCAPTCHA interstitial to
	//  attach a token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.redirect
	Redirect *FirewallAction_RedirectAction `json:"redirect,omitempty"`

	// This action transparently serves a different page to an offending
	//  user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.substitute
	Substitute *FirewallAction_SubstituteAction `json:"substitute,omitempty"`

	// This action sets a custom header but allow the request to continue
	//  to the customer backend.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.set_header
	SetHeader *FirewallAction_SetHeaderAction `json:"setHeader,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction.AllowAction
type FirewallAction_AllowAction struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction.BlockAction
type FirewallAction_BlockAction struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction.IncludeRecaptchaScriptAction
type FirewallAction_IncludeRecaptchaScriptAction struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction.RedirectAction
type FirewallAction_RedirectAction struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction.SetHeaderAction
type FirewallAction_SetHeaderAction struct {
	// Optional. The header key to set in the request to the backend server.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.SetHeaderAction.key
	Key *string `json:"key,omitempty"`

	// Optional. The header value to set in the request to the backend server.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.SetHeaderAction.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallAction.SubstituteAction
type FirewallAction_SubstituteAction struct {
	// Optional. The address to redirect to. The target is a relative path in
	//  the current host. Example: "/blog/404.html".
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallAction.SubstituteAction.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallPolicy
type FirewallPolicy struct {
	// Identifier. The resource name for the FirewallPolicy in the format
	//  `projects/{project}/firewallpolicies/{firewallpolicy}`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.name
	Name *string `json:"name,omitempty"`

	// Optional. A description of what this policy aims to achieve, for
	//  convenience purposes. The description can at most include 256 UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. The path for which this policy applies, specified as a glob
	//  pattern. For more information on glob, see the [manual
	//  page](https://man7.org/linux/man-pages/man7/glob.7.html).
	//  A path has a max length of 200 characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.path
	Path *string `json:"path,omitempty"`

	// Optional. A CEL (Common Expression Language) conditional expression that
	//  specifies if this policy applies to an incoming user request. If this
	//  condition evaluates to true and the requested path matched the path
	//  pattern, the associated actions should be executed by the caller. The
	//  condition string is checked for CEL syntax correctness on creation. For
	//  more information, see the [CEL spec](https://github.com/google/cel-spec)
	//  and its [language
	//  definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md).
	//  A condition has a max length of 500 characters.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.condition
	Condition *string `json:"condition,omitempty"`

	// Optional. The actions that the caller should take regarding user access.
	//  There should be at most one terminal action. A terminal action is any
	//  action that forces a response, such as `AllowAction`,
	//  `BlockAction` or `SubstituteAction`.
	//  Zero or more non-terminal actions such as `SetHeader` might be
	//  specified. A single policy can contain up to 16 actions.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicy.actions
	Actions []FirewallAction `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallPolicyAssessment
type FirewallPolicyAssessment struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment
type FraudPreventionAssessment struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.BehavioralTrustVerdict
type FraudPreventionAssessment_BehavioralTrustVerdict struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.CardTestingVerdict
type FraudPreventionAssessment_CardTestingVerdict struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.StolenInstrumentVerdict
type FraudPreventionAssessment_StolenInstrumentVerdict struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudSignals
type FraudSignals struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudSignals.CardSignals
type FraudSignals_CardSignals struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudSignals.UserSignals
type FraudSignals_UserSignals struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.PhoneFraudAssessment
type PhoneFraudAssessment struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.PrivatePasswordLeakVerification
type PrivatePasswordLeakVerification struct {
	// Required. Exactly 26-bit prefix of the SHA-256 hash of the canonicalized
	//  username. It is used to look up password leaks associated with that hash
	//  prefix.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.PrivatePasswordLeakVerification.lookup_hash_prefix
	LookupHashPrefix []byte `json:"lookupHashPrefix,omitempty"`

	// Optional. Encrypted Scrypt hash of the canonicalized username+password. It
	//  is re-encrypted by the server and returned through
	//  `reencrypted_user_credentials_hash`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.PrivatePasswordLeakVerification.encrypted_user_credentials_hash
	EncryptedUserCredentialsHash []byte `json:"encryptedUserCredentialsHash,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.RiskAnalysis
type RiskAnalysis struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.SmsTollFraudVerdict
type SmsTollFraudVerdict struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TokenProperties
type TokenProperties struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TransactionData
type TransactionData struct {
	// Unique identifier for the transaction. This custom identifier can be used
	//  to reference this transaction in the future, for example, labeling a refund
	//  or chargeback event. Two attempts at the same transaction should use the
	//  same transaction id.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.transaction_id
	TransactionID *string `json:"transactionID,omitempty"`

	// Optional. The payment method for the transaction. The allowed values are:
	//
	//  * credit-card
	//  * debit-card
	//  * gift-card
	//  * processor-{name} (If a third-party is used, for example,
	//  processor-paypal)
	//  * custom-{name} (If an alternative method is used, for example,
	//  custom-crypto)
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.payment_method
	PaymentMethod *string `json:"paymentMethod,omitempty"`

	// Optional. The Bank Identification Number - generally the first 6 or 8
	//  digits of the card.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.card_bin
	CardBin *string `json:"cardBin,omitempty"`

	// Optional. The last four digits of the card.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.card_last_four
	CardLastFour *string `json:"cardLastFour,omitempty"`

	// Optional. The currency code in ISO-4217 format.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// Optional. The decimal value of the transaction in the specified currency.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.value
	Value *float64 `json:"value,omitempty"`

	// Optional. The value of shipping in the specified currency. 0 for free or no
	//  shipping.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.shipping_value
	ShippingValue *float64 `json:"shippingValue,omitempty"`

	// Optional. Destination address if this transaction involves shipping a
	//  physical item.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.shipping_address
	ShippingAddress *TransactionData_Address `json:"shippingAddress,omitempty"`

	// Optional. Address associated with the payment method when applicable.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.billing_address
	BillingAddress *TransactionData_Address `json:"billingAddress,omitempty"`

	// Optional. Information about the user paying/initiating the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.user
	User *TransactionData_User `json:"user,omitempty"`

	// Optional. Information about the user or users fulfilling the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.merchants
	Merchants []TransactionData_User `json:"merchants,omitempty"`

	// Optional. Items purchased in this transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.items
	Items []TransactionData_Item `json:"items,omitempty"`

	// Optional. Information about the payment gateway's response to the
	//  transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.gateway_info
	GatewayInfo *TransactionData_GatewayInfo `json:"gatewayInfo,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TransactionData.Address
type TransactionData_Address struct {
	// Optional. The recipient name, potentially including information such as
	//  "care of".
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Address.recipient
	Recipient *string `json:"recipient,omitempty"`

	// Optional. The first lines of the address. The first line generally
	//  contains the street name and number, and further lines may include
	//  information such as an apartment number.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Address.address
	Address []string `json:"address,omitempty"`

	// Optional. The town/city of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Address.locality
	Locality *string `json:"locality,omitempty"`

	// Optional. The state, province, or otherwise administrative area of the
	//  address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Address.administrative_area
	AdministrativeArea *string `json:"administrativeArea,omitempty"`

	// Optional. The CLDR country/region of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Address.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// Optional. The postal or ZIP code of the address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Address.postal_code
	PostalCode *string `json:"postalCode,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TransactionData.GatewayInfo
type TransactionData_GatewayInfo struct {
	// Optional. Name of the gateway service (for example, stripe, square,
	//  paypal).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.GatewayInfo.name
	Name *string `json:"name,omitempty"`

	// Optional. Gateway response code describing the state of the transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.GatewayInfo.gateway_response_code
	GatewayResponseCode *string `json:"gatewayResponseCode,omitempty"`

	// Optional. AVS response code from the gateway
	//  (available only when reCAPTCHA Enterprise is called after authorization).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.GatewayInfo.avs_response_code
	AvsResponseCode *string `json:"avsResponseCode,omitempty"`

	// Optional. CVV response code from the gateway
	//  (available only when reCAPTCHA Enterprise is called after authorization).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.GatewayInfo.cvv_response_code
	CvvResponseCode *string `json:"cvvResponseCode,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TransactionData.Item
type TransactionData_Item struct {
	// Optional. The full name of the item.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Item.name
	Name *string `json:"name,omitempty"`

	// Optional. The value per item that the user is paying, in the transaction
	//  currency, after discounts.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Item.value
	Value *float64 `json:"value,omitempty"`

	// Optional. The quantity of this item that is being purchased.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Item.quantity
	Quantity *int64 `json:"quantity,omitempty"`

	// Optional. When a merchant is specified, its corresponding account_id.
	//  Necessary to populate marketplace-style transactions.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.Item.merchant_account_id
	MerchantAccountID *string `json:"merchantAccountID,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TransactionData.User
type TransactionData_User struct {
	// Optional. Unique account identifier for this user. If using account
	//  defender, this should match the hashed_account_id field. Otherwise, a
	//  unique and persistent identifier for this account.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.User.account_id
	AccountID *string `json:"accountID,omitempty"`

	// Optional. The epoch milliseconds of the user's account creation.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.User.creation_ms
	CreationMs *int64 `json:"creationMs,omitempty"`

	// Optional. The email address of the user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.User.email
	Email *string `json:"email,omitempty"`

	// Optional. Whether the email has been verified to be accessible by the
	//  user (OTP or similar).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.User.email_verified
	EmailVerified *bool `json:"emailVerified,omitempty"`

	// Optional. The phone number of the user, with country code.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.User.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Optional. Whether the phone number has been verified to be accessible by
	//  the user (OTP or similar).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TransactionData.User.phone_verified
	PhoneVerified *bool `json:"phoneVerified,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.UserId
type UserId struct {
	// Optional. An email address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.UserId.email
	Email *string `json:"email,omitempty"`

	// Optional. A phone number. Should use the E.164 format.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.UserId.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Optional. A unique username, if different from all the other identifiers
	//  and `account_id` that are provided. Can be a unique login handle or
	//  display name for a user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.UserId.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.UserInfo
type UserInfo struct {
	// Optional. Creation time for this account associated with this user. Leave
	//  blank for non logged-in actions, guest checkout, or when there is no
	//  account associated with the current user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.UserInfo.create_account_time
	CreateAccountTime *string `json:"createAccountTime,omitempty"`

	// Optional. For logged-in requests or login/registration requests, the unique
	//  account identifier associated with this user. You can use the username if
	//  it is stable (meaning it is the same for every request associated with the
	//  same user), or any stable user ID of your choice. Leave blank for non
	//  logged-in actions or guest checkout.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.UserInfo.account_id
	AccountID *string `json:"accountID,omitempty"`

	// Optional. Identifiers associated with this user or request.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.UserInfo.user_ids
	UserIds []UserId `json:"userIds,omitempty"`
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

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.AccountDefenderAssessment
type AccountDefenderAssessmentObservedState struct {
	// Output only. Labels for this request.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AccountDefenderAssessment.labels
	Labels []string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo
type AccountVerificationInfoObservedState struct {
	// Optional. Endpoints that can be used for identity verification.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo.endpoints
	Endpoints []EndpointVerificationInfoObservedState `json:"endpoints,omitempty"`

	// Output only. Result of the latest account verification challenge.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AccountVerificationInfo.latest_verification_result
	LatestVerificationResult *string `json:"latestVerificationResult,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Assessment
type AssessmentObservedState struct {
	// Output only. Identifier. The resource name for the Assessment in the format
	//  `projects/{project}/assessments/{assessment}`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.name
	Name *string `json:"name,omitempty"`

	// Output only. The risk analysis result for the event being assessed.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.risk_analysis
	RiskAnalysis *RiskAnalysis `json:"riskAnalysis,omitempty"`

	// Output only. Properties of the provided event token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.token_properties
	TokenProperties *TokenProperties `json:"tokenProperties,omitempty"`

	// Optional. Account verification information for identity verification. The
	//  assessment event must include a token and site key to use this feature.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.account_verification
	AccountVerification *AccountVerificationInfoObservedState `json:"accountVerification,omitempty"`

	// Output only. Assessment returned by account defender when an account
	//  identifier is provided.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.account_defender_assessment
	AccountDefenderAssessment *AccountDefenderAssessment `json:"accountDefenderAssessment,omitempty"`

	// Optional. The private password leak verification field contains the
	//  parameters that are used to to check for leaks privately without sharing
	//  user credentials.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.private_password_leak_verification
	PrivatePasswordLeakVerification *PrivatePasswordLeakVerificationObservedState `json:"privatePasswordLeakVerification,omitempty"`

	// Output only. Assessment returned when firewall policies belonging to the
	//  project are evaluated using the field firewall_policy_evaluation.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.firewall_policy_assessment
	FirewallPolicyAssessment *FirewallPolicyAssessment `json:"firewallPolicyAssessment,omitempty"`

	// Output only. Assessment returned by Fraud Prevention when TransactionData
	//  is provided.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.fraud_prevention_assessment
	FraudPreventionAssessment *FraudPreventionAssessment `json:"fraudPreventionAssessment,omitempty"`

	// Output only. Fraud Signals specific to the users involved in a payment
	//  transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.fraud_signals
	FraudSignals *FraudSignals `json:"fraudSignals,omitempty"`

	// Output only. Assessment returned when a site key, a token, and a phone
	//  number as `user_id` are provided. Account defender and SMS toll fraud
	//  protection need to be enabled.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Assessment.phone_fraud_assessment
	PhoneFraudAssessment *PhoneFraudAssessment `json:"phoneFraudAssessment,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.EndpointVerificationInfo
type EndpointVerificationInfoObservedState struct {
	// Output only. Token to provide to the client to trigger endpoint
	//  verification. It must be used within 15 minutes.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.EndpointVerificationInfo.request_token
	RequestToken *string `json:"requestToken,omitempty"`

	// Output only. Timestamp of the last successful verification for the
	//  endpoint, if any.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.EndpointVerificationInfo.last_verification_time
	LastVerificationTime *string `json:"lastVerificationTime,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FirewallPolicyAssessment
type FirewallPolicyAssessmentObservedState struct {
	// Output only. If the processing of a policy config fails, an error is
	//  populated and the firewall_policy is left empty.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicyAssessment.error
	Error *Status `json:"error,omitempty"`

	// Output only. The policy that matched the request. If more than one policy
	//  may match, this is the first match. If no policy matches the incoming
	//  request, the policy field is left empty.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FirewallPolicyAssessment.firewall_policy
	FirewallPolicy *FirewallPolicy `json:"firewallPolicy,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment
type FraudPreventionAssessmentObservedState struct {
	// Output only. Probability of this transaction being fraudulent. Summarizes
	//  the combined risk of attack vectors below. Values are from 0.0 (lowest)
	//  to 1.0 (highest).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.transaction_risk
	TransactionRisk *float32 `json:"transactionRisk,omitempty"`

	// Output only. Assessment of this transaction for risk of a stolen
	//  instrument.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.stolen_instrument_verdict
	StolenInstrumentVerdict *FraudPreventionAssessment_StolenInstrumentVerdict `json:"stolenInstrumentVerdict,omitempty"`

	// Output only. Assessment of this transaction for risk of being part of a
	//  card testing attack.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.card_testing_verdict
	CardTestingVerdict *FraudPreventionAssessment_CardTestingVerdict `json:"cardTestingVerdict,omitempty"`

	// Output only. Assessment of this transaction for behavioral trust.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.behavioral_trust_verdict
	BehavioralTrustVerdict *FraudPreventionAssessment_BehavioralTrustVerdict `json:"behavioralTrustVerdict,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.BehavioralTrustVerdict
type FraudPreventionAssessment_BehavioralTrustVerdictObservedState struct {
	// Output only. Probability of this transaction attempt being executed in a
	//  behaviorally trustworthy way. Values are from 0.0 (lowest) to 1.0
	//  (highest).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.BehavioralTrustVerdict.trust
	Trust *float32 `json:"trust,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.CardTestingVerdict
type FraudPreventionAssessment_CardTestingVerdictObservedState struct {
	// Output only. Probability of this transaction attempt being part of a card
	//  testing attack. Values are from 0.0 (lowest) to 1.0 (highest).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.CardTestingVerdict.risk
	Risk *float32 `json:"risk,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.StolenInstrumentVerdict
type FraudPreventionAssessment_StolenInstrumentVerdictObservedState struct {
	// Output only. Probability of this transaction being executed with a stolen
	//  instrument. Values are from 0.0 (lowest) to 1.0 (highest).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudPreventionAssessment.StolenInstrumentVerdict.risk
	Risk *float32 `json:"risk,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudSignals
type FraudSignalsObservedState struct {
	// Output only. Signals describing the end user in this transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudSignals.user_signals
	UserSignals *FraudSignals_UserSignals `json:"userSignals,omitempty"`

	// Output only. Signals describing the payment card or cards used in this
	//  transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudSignals.card_signals
	CardSignals *FraudSignals_CardSignals `json:"cardSignals,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudSignals.CardSignals
type FraudSignals_CardSignalsObservedState struct {
	// Output only. The labels for the payment card in this transaction.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudSignals.CardSignals.card_labels
	CardLabels []string `json:"cardLabels,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.FraudSignals.UserSignals
type FraudSignals_UserSignalsObservedState struct {
	// Output only. This user (based on email, phone, and other identifiers) has
	//  been seen on the internet for at least this number of days.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudSignals.UserSignals.active_days_lower_bound
	ActiveDaysLowerBound *int32 `json:"activeDaysLowerBound,omitempty"`

	// Output only. Likelihood (from 0.0 to 1.0) this user includes synthetic
	//  components in their identity, such as a randomly generated email address,
	//  temporary phone number, or fake shipping address.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.FraudSignals.UserSignals.synthetic_risk
	SyntheticRisk *float32 `json:"syntheticRisk,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.PhoneFraudAssessment
type PhoneFraudAssessmentObservedState struct {
	// Output only. Assessment of this phone event for risk of SMS toll fraud.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.PhoneFraudAssessment.sms_toll_fraud_verdict
	SmsTollFraudVerdict *SmsTollFraudVerdict `json:"smsTollFraudVerdict,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.PrivatePasswordLeakVerification
type PrivatePasswordLeakVerificationObservedState struct {
	// Output only. List of prefixes of the encrypted potential password leaks
	//  that matched the given parameters. They must be compared with the
	//  client-side decryption prefix of `reencrypted_user_credentials_hash`
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.PrivatePasswordLeakVerification.encrypted_leak_match_prefixes
	EncryptedLeakMatchPrefixes [][]byte `json:"encryptedLeakMatchPrefixes,omitempty"`

	// Output only. Corresponds to the re-encryption of the
	//  `encrypted_user_credentials_hash` field. It is used to match potential
	//  password leaks within `encrypted_leak_match_prefixes`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.PrivatePasswordLeakVerification.reencrypted_user_credentials_hash
	ReencryptedUserCredentialsHash []byte `json:"reencryptedUserCredentialsHash,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.RiskAnalysis
type RiskAnalysisObservedState struct {
	// Output only. Legitimate event score from 0.0 to 1.0.
	//  (1.0 means very likely legitimate traffic while 0.0 means very likely
	//  non-legitimate traffic).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RiskAnalysis.score
	Score *float32 `json:"score,omitempty"`

	// Output only. Reasons contributing to the risk analysis verdict.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RiskAnalysis.reasons
	Reasons []string `json:"reasons,omitempty"`

	// Output only. Extended verdict reasons to be used for experimentation only.
	//  The set of possible reasons is subject to change.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RiskAnalysis.extended_verdict_reasons
	ExtendedVerdictReasons []string `json:"extendedVerdictReasons,omitempty"`

	// Output only. Challenge information for SCORE_AND_CHALLENGE and INVISIBLE
	//  keys
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RiskAnalysis.challenge
	Challenge *string `json:"challenge,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.SmsTollFraudVerdict
type SmsTollFraudVerdictObservedState struct {
	// Output only. Probability of an SMS event being fraudulent.
	//  Values are from 0.0 (lowest) to 1.0 (highest).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.SmsTollFraudVerdict.risk
	Risk *float32 `json:"risk,omitempty"`

	// Output only. Reasons contributing to the SMS toll fraud verdict.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.SmsTollFraudVerdict.reasons
	Reasons []string `json:"reasons,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TokenProperties
type TokenPropertiesObservedState struct {
	// Output only. Whether the provided user response token is valid. When valid
	//  = false, the reason could be specified in invalid_reason or it could also
	//  be due to a user failing to solve a challenge or a sitekey mismatch (i.e
	//  the sitekey used to generate the token was different than the one specified
	//  in the assessment).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.valid
	Valid *bool `json:"valid,omitempty"`

	// Output only. Reason associated with the response when valid = false.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.invalid_reason
	InvalidReason *string `json:"invalidReason,omitempty"`

	// Output only. The timestamp corresponding to the generation of the token.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The hostname of the page on which the token was generated (Web
	//  keys only).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Output only. The name of the Android package with which the token was
	//  generated (Android keys only).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.android_package_name
	AndroidPackageName *string `json:"androidPackageName,omitempty"`

	// Output only. The ID of the iOS bundle with which the token was generated
	//  (iOS keys only).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.ios_bundle_id
	IosBundleID *string `json:"iosBundleID,omitempty"`

	// Output only. Action name provided at token generation.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TokenProperties.action
	Action *string `json:"action,omitempty"`
}
