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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesecuritypolicy;gcpcomputesecuritypolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:object:root=true
// ComputeSecurityPolicy is the Schema for the ComputeSecurityPolicy API
// +k8s:openapi-gen=true
type ComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status ComputeSecurityPolicyStatus `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicy
type ComputeSecurityPolicySpec struct {
	// The ComputeSecurityPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kcc:proto:field=adaptive_protection_config
	AdaptiveProtectionConfig *SecurityPolicyAdaptiveProtectionConfig `json:"adaptiveProtectionConfig,omitempty"`

	// +kcc:proto:field=advanced_options_config
	AdvancedOptionsConfig *SecurityPolicyAdvancedOptionsConfig `json:"advancedOptionsConfig,omitempty"`

	// +kcc:proto:field=description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=recaptcha_options_config
	RecaptchaOptionsConfig *SecurityPolicyRecaptchaOptionsConfig `json:"recaptchaOptionsConfig,omitempty"`

	// +kcc:proto:field=rules
	Rules []SecurityPolicyRule `json:"rule,omitempty"`

	// +kcc:proto:field=type
	Type *string `json:"type,omitempty"`
}

type SecurityPolicyAdaptiveProtectionConfig struct {
	// Layer 7 DDoS Defense Config of this security policy
	Layer7DdosDefenseConfig *SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig `json:"layer7DdosDefenseConfig,omitempty"`

	// Auto Deploy Config of this security policy
	AutoDeployConfig *SecurityPolicyAdaptiveProtectionConfigAutoDeployConfig `json:"autoDeployConfig,omitempty"`
}

type SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig struct {
	// If set to true, enables CAAP for L7 DDoS detection.
	Enable *bool `json:"enable,omitempty"`

	// Rule visibility. Supported values include: "STANDARD", "PREMIUM".
	RuleVisibility *string `json:"ruleVisibility,omitempty"`
}

type SecurityPolicyAdaptiveProtectionConfigAutoDeployConfig struct {
	// Identifies new attackers only when the load to the backend service that is under attack exceeds this threshold.
	LoadThreshold *float32 `json:"loadThreshold,omitempty"`

	// Rules are only automatically deployed for alerts on potential attacks with confidence scores greater than this threshold.
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Rules are only automatically deployed when the estimated impact to baseline traffic from the suggested mitigation is below this threshold.
	ImpactedBaselineThreshold *float32 `json:"impactedBaselineThreshold,omitempty"`

	// Google Cloud Armor stops applying the action in the automatically deployed rule to an identified attacker after this duration. The rule continues to operate against new requests.
	ExpirationSec *int32 `json:"expirationSec,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfig
type SecurityPolicyAdvancedOptionsConfig struct {
	// Custom configuration to apply the JSON parsing. Only applicable when json_parsing is set to STANDARD.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfig.json_custom_config
	JsonCustomConfig *SecurityPolicyAdvancedOptionsConfigJsonCustomConfig `json:"jsonCustomConfig,omitempty"`

	// Check the JsonParsing enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfig.json_parsing
	JsonParsing *string `json:"jsonParsing,omitempty"`

	// Check the LogLevel enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfig.log_level
	LogLevel *string `json:"logLevel,omitempty"`

	// An optional list of case-insensitive request header names to use for resolving the callers client IP address.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfig.user_ip_request_headers
	UserIPRequestHeaders []string `json:"userIpRequestHeaders,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig
type SecurityPolicyAdvancedOptionsConfigJsonCustomConfig struct {
	// +required
	// A list of custom Content-Type header values to apply the JSON parsing. As per RFC 1341, a Content-Type header value has the following format: Content-Type := type "/" subtype *[";" parameter] When configuring a custom Content-Type header value, only the type/subtype needs to be specified, and the parameters should be excluded.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig.content_types
	ContentTypes []string `json:"contentTypes,omitempty"`
}

type SecurityPolicyRecaptchaOptionsConfig struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRecaptchaOptionsConfig.redirect_site_key
	RedirectSiteKeyRef *RecaptchaEnterpriseKeyRef `json:"redirectSiteKeyRef,omitempty"`
}

type RecaptchaEnterpriseKeyRef struct {
	/* The value of an externally managed RecaptchaEnterpriseKey resource. */
	External string `json:"external,omitempty"`

	/* The name of a RecaptchaEnterpriseKey resource. */
	Name string `json:"name,omitempty"`

	/* The namespace of a RecaptchaEnterpriseKey resource. */
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRule
type SecurityPolicyRule struct {
	// +required
	// The Action to perform when the rule is matched. The following are the valid actions: - allow: allow access to target. - deny(STATUS): deny access to target, returns the HTTP response code specified. Valid values for `STATUS` are 403, 404, and 502. - rate_based_ban: limit client traffic to the configured threshold and ban the client if the traffic exceeds the threshold. Configure parameters for this action in RateLimitOptions. Requires rate_limit_options to be set. - redirect: redirect to a different target. This can either be an internal reCAPTCHA redirect, or an external URL-based redirect via a 302 response. Parameters for this action can be configured via redirectOptions. This action is only supported in Global Security Policies of type CLOUD_ARMOR. - throttle: limit client traffic to the configured threshold. Configure parameters for this action in rateLimitOptions. Requires rate_limit_options to be set for this.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.action
	Action *string `json:"action,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.description
	Description *string `json:"description,omitempty"`

	// Optional, additional actions that are performed on headers. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.header_action
	HeaderAction *SecurityPolicyRuleHTTPHeaderAction `json:"headerAction,omitempty"`

	// +required
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.match
	Match *SecurityPolicyRuleMatcher `json:"match,omitempty"`

	// Preconfigured WAF configuration to be applied for the rule. If the rule does not evaluate preconfigured WAF rules, i.e., if evaluatePreconfiguredWaf() is not used, this field will have no effect.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.preconfigured_waf_config
	PreconfiguredWafConfig *SecurityPolicyRulePreconfiguredWafConfig `json:"preconfiguredWafConfig,omitempty"`

	// If set to true, the specified action is not enforced.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.preview
	Preview *bool `json:"preview,omitempty"`

	// +required
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.priority
	Priority *int32 `json:"priority,omitempty"`

	// Must be specified if the action is "rate_based_ban" or "throttle". Cannot be specified for any other actions.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.rate_limit_options
	RateLimitOptions *SecurityPolicyRuleRateLimitOptions `json:"rateLimitOptions,omitempty"`

	// Parameters defining the redirect action. Cannot be specified for any other actions. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRule.redirect_options
	RedirectOptions *SecurityPolicyRuleRedirectOptions `json:"redirectOptions,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleMatcher
type SecurityPolicyRuleMatcher struct {
	// The configuration options available when specifying versioned_expr. This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcher.config
	Config *SecurityPolicyRuleMatcherConfig `json:"config,omitempty"`

	// User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header. Expressions containing `evaluateThreatIntelligence` require Cloud Armor Managed Protection Plus tier and are not supported in Edge Policies nor in Regional Policies. Expressions containing `evaluatePreconfiguredExpr('sourceiplist-*')` require Cloud Armor Managed Protection Plus tier and are only supported in Global Security Policies.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcher.expr
	Expr *SecurityPolicyRuleMatcherExpr `json:"expr,omitempty"`

	// Preconfigured versioned expression. If this field is specified, config must also be specified. Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding src_ip_range field in config.
	//  Check the VersionedExpr enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcher.versioned_expr
	VersionedExpr *string `json:"versionedExpr,omitempty"`
}

type SecurityPolicyRuleMatcherExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfig
type SecurityPolicyRuleMatcherConfig struct {
	// CIDR IP address range. Maximum number of src_ip_ranges allowed is 10.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleMatcherConfig.src_ip_ranges
	SrcIPRanges []string `json:"srcIpRanges,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions
type SecurityPolicyRuleRateLimitOptions struct {
	// Can only be specified if the action for the rule is "rate_based_ban". If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.ban_duration_sec
	BanDurationSec *int32 `json:"banDurationSec,omitempty"`

	// Can only be specified if the action for the rule is "rate_based_ban". If specified, the key will be banned for the configured 'ban_duration_sec' when the number of requests that exceed the 'rate_limit_threshold' also exceed this 'ban_threshold'.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.ban_threshold
	BanThreshold *SecurityPolicyRuleRateLimitOptionsThreshold `json:"banThreshold,omitempty"`

	// Action to take for requests that are under the configured rate limit threshold. Valid option is "allow" only.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.conform_action
	ConformAction *string `json:"conformAction,omitempty"`

	// Determines the key to enforce the rate_limit_threshold on. Possible values are: - ALL: A single rate limit threshold is applied to all the requests matching this rule. This is the default value if "enforceOnKey" is not configured. - IP: The source IP address of the request is the key. Each IP has this limit enforced separately. - HTTP_HEADER: The value of the HTTP header whose name is configured under "enforceOnKeyName". The key value is truncated to the first 128 bytes of the header value. If no such header is present in the request, the key type defaults to ALL. - XFF_IP: The first IP address (i.e. the originating client IP address) specified in the list of IPs under X-Forwarded-For HTTP header. If no such header is present or the value is not a valid IP, the key defaults to the source IP address of the request i.e. key type IP. - HTTP_COOKIE: The value of the HTTP cookie whose name is configured under "enforceOnKeyName". The key value is truncated to the first 128 bytes of the cookie value. If no such cookie is present in the request, the key type defaults to ALL. - HTTP_PATH: The URL path of the HTTP request. The key value is truncated to the first 128 bytes. - SNI: Server name indication in the TLS session of the HTTPS request. The key value is truncated to the first 128 bytes. The key type defaults to ALL on a HTTP session. - REGION_CODE: The country/region from which the request originates. - TLS_JA3_FINGERPRINT: JA3 TLS/SSL fingerprint if the client connects using HTTPS, HTTP/2 or HTTP/3. If not available, the key type defaults to ALL. - USER_IP: The IP address of the originating client, which is resolved based on "userIpRequestHeaders" configured with the security policy. If there is no "userIpRequestHeaders" configuration or an IP address cannot be resolved from it, the key type defaults to IP. - TLS_JA4_FINGERPRINT: JA4 TLS/SSL fingerprint if the client connects using HTTPS, HTTP/2 or HTTP/3. If not available, the key type defaults to ALL.
	//  Check the EnforceOnKey enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.enforce_on_key
	EnforceOnKey *string `json:"enforceOnKey,omitempty"`

	// If specified, any combination of values of enforce_on_key_type/enforce_on_key_name is treated as the key on which ratelimit threshold/action is enforced. You can specify up to 3 enforce_on_key_configs. If enforce_on_key_configs is specified, enforce_on_key must not be specified.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.enforce_on_key_configs
	EnforceOnKeyConfigs []SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig `json:"enforceOnKeyConfigs,omitempty"`

	// Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.enforce_on_key_name
	EnforceOnKeyName *string `json:"enforceOnKeyName,omitempty"`

	// Action to take for requests that are above the configured rate limit threshold, to either deny with a specified HTTP response code, or redirect to a different endpoint. Valid options are `deny(STATUS)`, where valid values for `STATUS` are 403, 404, 429, and 502, and `redirect`, where the redirect parameters come from `exceedRedirectOptions` below. The `redirect` action is only supported in Global Security Policies of type CLOUD_ARMOR.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.exceed_action
	ExceedAction *string `json:"exceedAction,omitempty"`

	// Parameters defining the redirect action that is used as the exceed action. Cannot be specified if the exceed action is not redirect. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.exceed_redirect_options
	ExceedRedirectOptions *SecurityPolicyRuleRedirectOptions `json:"exceedRedirectOptions,omitempty"`

	// Threshold at which to begin ratelimiting.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptions.rate_limit_threshold
	RateLimitThreshold *SecurityPolicyRuleRateLimitOptionsThreshold `json:"rateLimitThreshold,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptionsThreshold
type SecurityPolicyRuleRateLimitOptionsThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptionsThreshold.count
	Count *int32 `json:"count,omitempty"`

	// Interval over which the threshold is computed.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRateLimitOptionsThreshold.interval_sec
	IntervalSec *int32 `json:"intervalSec,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecurityPolicyRuleRedirectOptions
type SecurityPolicyRuleRedirectOptions struct {
	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRedirectOptions.target
	Target *string `json:"target,omitempty"`

	// +required
	// Type of the redirect action. Possible values are: - GOOGLE_RECAPTCHA: redirect to reCAPTCHA for manual challenge assessment. - EXTERNAL_302: redirect to a different URL via a 302 response.
	//  Check the Type enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SecurityPolicyRuleRedirectOptions.type
	Type *string `json:"type,omitempty"`
}

type SecurityPolicyRulePreconfiguredWafConfig struct {
	// A list of exclusions to apply during preconfigured WAF evaluation.
	Exclusion []SecurityPolicyRulePreconfiguredWafConfigExclusion `json:"exclusion,omitempty"`
}

type SecurityPolicyRulePreconfiguredWafConfigExclusion struct {
	// A list of request cookie names whose value will be excluded from inspection during preconfigured WAF evaluation.
	RequestCookie []SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams `json:"requestCookie,omitempty"`

	// A list of request header names whose value will be excluded from inspection during preconfigured WAF evaluation.
	RequestHeader []SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams `json:"requestHeader,omitempty"`

	// A list of request query parameter names whose value will be excluded from inspection during preconfigured WAF evaluation. Note that the parameter can be in the query string or in the POST body.
	RequestQueryParam []SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams `json:"requestQueryParam,omitempty"`

	// A list of request URIs from the request line to be excluded from inspection during preconfigured WAF evaluation. When specifying this field, the query or fragment part should be excluded.
	RequestUri []SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams `json:"requestUri,omitempty"`

	// A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion. If omitted, it refers to all the rule IDs under the WAF rule set.
	TargetRuleIds []string `json:"targetRuleIds,omitempty"`

	// Target WAF rule set to apply the preconfigured WAF exclusion.
	TargetRuleSet *string `json:"targetRuleSet,omitempty"`
}

type SecurityPolicyRulePreconfiguredWafConfigExclusionFieldParams struct {
	// The match operator for the field.
	//  Check the Op enum for the list of possible values.
	Operator *string `json:"operator,omitempty"`

	// The value of the field.
	Value *string `json:"value,omitempty"`
}

// ComputeSecurityPolicyStatus defines the config connector machine state of ComputeSecurityPolicy
type ComputeSecurityPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +kcc:proto:field=fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// +kcc:proto:field=self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeSecurityPolicyObservedState is the state of the ComputeSecurityPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.SecurityPolicy
type ComputeSecurityPolicyObservedState struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// ComputeSecurityPolicyList contains a list of ComputeSecurityPolicy
type ComputeSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeSecurityPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSecurityPolicy{}, &ComputeSecurityPolicyList{})
}
