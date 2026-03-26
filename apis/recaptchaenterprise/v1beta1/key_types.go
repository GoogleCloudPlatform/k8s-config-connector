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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RecaptchaEnterpriseKeyGVK = GroupVersion.WithKind("RecaptchaEnterpriseKey")

// RecaptchaEnterpriseKeySpec defines the desired state of RecaptchaEnterpriseKey
// +kcc:spec:proto=google.cloud.recaptchaenterprise.v1.Key
type RecaptchaEnterpriseKeySpec struct {
	// Immutable. The project that this resource belongs to.
	ProjectRef *ProjectRef `json:"projectRef"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// Human-readable display name of this key. Modifiable by user.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.display_name
	// +required
	DisplayName *string `json:"displayName"`

	// Settings for keys that can be used by websites.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.web_settings
	WebSettings *KeyWebSettings `json:"webSettings,omitempty"`

	// Settings for keys that can be used by Android apps.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.android_settings
	AndroidSettings *KeyAndroidSettings `json:"androidSettings,omitempty"`

	// Settings for keys that can be used by iOS apps.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.ios_settings
	IosSettings *KeyIosSettings `json:"iosSettings,omitempty"`

	// Immutable. Options for user acceptance testing.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.testing_options
	TestingOptions *KeyTestingOptions `json:"testingOptions,omitempty"`

	// Immutable. Settings specific to keys that can be used for WAF (Web Application Firewall).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.waf_settings
	WafSettings *KeyWafSettings `json:"wafSettings,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.WebKeySettings
type KeyWebSettings struct {
	// If set to true, it means allowed_domains are not enforced.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.allow_all_domains
	AllowAllDomains *bool `json:"allowAllDomains,omitempty"`

	// Domains or subdomains of websites allowed to use the key. All subdomains of an allowed domain are automatically allowed. A valid domain requires a host and must not include any path, port, query or fragment. Examples: 'example.com' or 'subdomain.example.com'
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.allowed_domains
	AllowedDomains []string `json:"allowedDomains,omitempty"`

	// If set to true, the key can be used on AMP (Accelerated Mobile Pages) websites. This is supported only for the SCORE integration type.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.allow_amp_traffic
	AllowAmpTraffic *bool `json:"allowAmpTraffic,omitempty"`

	// Immutable. Required. Describes how this key is integrated with the website. Possible values: SCORE, CHECKBOX, INVISIBLE
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.integration_type
	// +required
	IntegrationType *string `json:"integrationType"`

	// Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE. Possible values: CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED, USABILITY, BALANCE, SECURITY
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WebKeySettings.challenge_security_preference
	ChallengeSecurityPreference *string `json:"challengeSecurityPreference,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.AndroidKeySettings
type KeyAndroidSettings struct {
	// If set to true, it means allowed_package_names will not be enforced.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AndroidKeySettings.allow_all_package_names
	AllowAllPackageNames *bool `json:"allowAllPackageNames,omitempty"`

	// Android package names of apps allowed to use the key. Example: 'com.companyname.appname'
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.AndroidKeySettings.allowed_package_names
	AllowedPackageNames []string `json:"allowedPackageNames,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.IOSKeySettings
type KeyIosSettings struct {
	// If set to true, it means allowed_bundle_ids will not be enforced.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.IOSKeySettings.allow_all_bundle_ids
	AllowAllBundleIds *bool `json:"allowAllBundleIds,omitempty"`

	// iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname'
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.IOSKeySettings.allowed_bundle_ids
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.TestingOptions
type KeyTestingOptions struct {
	// Immutable. All assessments for this Key will return this score. Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.
	// +kubebuilder:validation:Format=double
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TestingOptions.testing_score
	TestingScore *float64 `json:"testingScore,omitempty"`

	// Immutable. For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if UNSOLVABLE_CHALLENGE. Possible values: TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.TestingOptions.testing_challenge
	TestingChallenge *string `json:"testingChallenge,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.WafSettings
type KeyWafSettings struct {
	// Immutable. The WAF service that uses this key. Possible values: CA, FASTLY
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WafSettings.waf_service
	// +required
	WafService *string `json:"wafService"`

	// Immutable. Supported WAF features. For more information, see https://cloud.google.com/recaptcha-enterprise/docs/usecase#comparison_of_features. Possible values: CHALLENGE_PAGE, SESSION_TOKEN, ACTION_TOKEN, EXPRESS
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.WafSettings.waf_feature
	// +required
	WafFeature *string `json:"wafFeature"`
}

type ProjectRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
}

// RecaptchaEnterpriseKeyStatus defines the config connector machine state of RecaptchaEnterpriseKey
type RecaptchaEnterpriseKeyStatus struct {
	// Conditions represent the latest available observations of the
	// object's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:validation:Type=integer
	// +kubebuilder:validation:Format=""
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RecaptchaEnterpriseKey resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RecaptchaEnterpriseKeyObservedState `json:"observedState,omitempty"`

	// The timestamp corresponding to the creation of this Key.
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`
}

// RecaptchaEnterpriseKeyObservedState is the state of the RecaptchaEnterpriseKey resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.recaptchaenterprise.v1.Key
type RecaptchaEnterpriseKeyObservedState struct {
	// The timestamp corresponding to the creation of this Key.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Key.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcprecaptchaenterprisekey;gcprecaptchaenterprisekeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RecaptchaEnterpriseKey is the Schema for the RecaptchaEnterpriseKey API
// +k8s:openapi-gen=true
type RecaptchaEnterpriseKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RecaptchaEnterpriseKeySpec   `json:"spec,omitempty"`
	Status RecaptchaEnterpriseKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RecaptchaEnterpriseKeyList contains a list of RecaptchaEnterpriseKey
type RecaptchaEnterpriseKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecaptchaEnterpriseKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RecaptchaEnterpriseKey{}, &RecaptchaEnterpriseKeyList{})
}
