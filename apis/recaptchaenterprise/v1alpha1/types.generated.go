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
