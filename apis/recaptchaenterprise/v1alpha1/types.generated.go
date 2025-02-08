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


// +kcc:proto=google.cloud.recaptchaenterprise.v1.AndroidKeySettings
type AndroidKeySettings struct {
	// Optional. If set to true, allowed_package_names are not enforced.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AndroidKeySettings.allow_all_package_names
	AllowAllPackageNames *bool `json:"allowAllPackageNames,omitempty"`

	// Optional. Android package names of apps allowed to use the key.
	//  Example: 'com.companyname.appname'
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AndroidKeySettings.allowed_package_names
	AllowedPackageNames []string `json:"allowedPackageNames,omitempty"`

	// Optional. Set to true for keys that are used in an Android application that
	//  is available for download in app stores in addition to the Google Play
	//  Store.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AndroidKeySettings.support_non_google_app_store_distribution
	SupportNonGoogleAppStoreDistribution *bool `json:"supportNonGoogleAppStoreDistribution,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.AppleDeveloperId
type AppleDeveloperId struct {
	// Required. Input only. A private key (downloaded as a text file with a .p8
	//  file extension) generated for your Apple Developer account. Ensure that
	//  Apple DeviceCheck is enabled for the private key.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AppleDeveloperId.private_key
	PrivateKey *string `json:"privateKey,omitempty"`

	// Required. The Apple developer key ID (10-character string).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AppleDeveloperId.key_id
	KeyID *string `json:"keyID,omitempty"`

	// Required. The Apple team ID (10-character string) owning the provisioning
	//  profile used to build your application.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AppleDeveloperId.team_id
	TeamID *string `json:"teamID,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.ExpressKeySettings
type ExpressKeySettings struct {
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.IOSKeySettings
type IOSKeySettings struct {
	// Optional. If set to true, allowed_bundle_ids are not enforced.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.IOSKeySettings.allow_all_bundle_ids
	AllowAllBundleIds *bool `json:"allowAllBundleIds,omitempty"`

	// Optional. iOS bundle ids of apps allowed to use the key.
	//  Example: 'com.companyname.productname.appname'
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.IOSKeySettings.allowed_bundle_ids
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty"`

	// Optional. Apple Developer account details for the app that is protected by
	//  the reCAPTCHA Key. reCAPTCHA leverages platform-specific checks like Apple
	//  App Attest and Apple DeviceCheck to protect your app from abuse. Providing
	//  these fields allows reCAPTCHA to get a better assessment of the integrity
	//  of your app.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.IOSKeySettings.apple_developer_id
	AppleDeveloperID *AppleDeveloperId `json:"appleDeveloperID,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Key
type Key struct {
	// Identifier. The resource name for the Key in the format
	//  `projects/{project}/keys/{key}`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.name
	Name *string `json:"name,omitempty"`

	// Required. Human-readable display name of this key. Modifiable by user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Settings for keys that can be used by websites.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.web_settings
	WebSettings *WebKeySettings `json:"webSettings,omitempty"`

	// Settings for keys that can be used by Android apps.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.android_settings
	AndroidSettings *AndroidKeySettings `json:"androidSettings,omitempty"`

	// Settings for keys that can be used by iOS apps.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.ios_settings
	IosSettings *IOSKeySettings `json:"iosSettings,omitempty"`

	// Settings for keys that can be used by reCAPTCHA Express.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.express_settings
	ExpressSettings *ExpressKeySettings `json:"expressSettings,omitempty"`

	// Optional. See [Creating and managing labels]
	//  (https://cloud.google.com/recaptcha/docs/labels).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Options for user acceptance testing.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.testing_options
	TestingOptions *TestingOptions `json:"testingOptions,omitempty"`

	// Optional. Settings for WAF
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.waf_settings
	WafSettings *WafSettings `json:"wafSettings,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TestingOptions
type TestingOptions struct {
	// Optional. All assessments for this Key return this score. Must be between 0
	//  (likely not legitimate) and 1 (likely legitimate) inclusive.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TestingOptions.testing_score
	TestingScore *float32 `json:"testingScore,omitempty"`

	// Optional. For challenge-based keys only (CHECKBOX, INVISIBLE), all
	//  challenge requests for this site return nocaptcha if NOCAPTCHA, or an
	//  unsolvable challenge if CHALLENGE.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TestingOptions.testing_challenge
	TestingChallenge *string `json:"testingChallenge,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.WafSettings
type WafSettings struct {
	// Required. The WAF service that uses this key.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WafSettings.waf_service
	WafService *string `json:"wafService,omitempty"`

	// Required. The WAF feature for which this key is enabled.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WafSettings.waf_feature
	WafFeature *string `json:"wafFeature,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.WebKeySettings
type WebKeySettings struct {
	// Optional. If set to true, it means allowed_domains are not enforced.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.allow_all_domains
	AllowAllDomains *bool `json:"allowAllDomains,omitempty"`

	// Optional. Domains or subdomains of websites allowed to use the key. All
	//  subdomains of an allowed domain are automatically allowed. A valid domain
	//  requires a host and must not include any path, port, query or fragment.
	//  Examples: 'example.com' or 'subdomain.example.com'
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.allowed_domains
	AllowedDomains []string `json:"allowedDomains,omitempty"`

	// Optional. If set to true, the key can be used on AMP (Accelerated Mobile
	//  Pages) websites. This is supported only for the SCORE integration type.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.allow_amp_traffic
	AllowAmpTraffic *bool `json:"allowAmpTraffic,omitempty"`

	// Required. Describes how this key is integrated with the website.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.integration_type
	IntegrationType *string `json:"integrationType,omitempty"`

	// Optional. Settings for the frequency and difficulty at which this key
	//  triggers captcha challenges. This should only be specified for
	//  IntegrationTypes CHECKBOX and INVISIBLE and SCORE_AND_CHALLENGE.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.challenge_security_preference
	ChallengeSecurityPreference *string `json:"challengeSecurityPreference,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Key
type KeyObservedState struct {
	// Output only. The timestamp corresponding to the creation of this key.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
