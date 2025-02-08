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


// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.BillingSettings
type BillingSettings struct {
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.ComponentSettings
type ComponentSettings struct {
	// The relative resource name of the component settings.
	//  Formats:
	//   * `organizations/{organization}/components/{component}/settings`
	//   * `folders/{folder}/components/{component}/settings`
	//   * `projects/{project}/components/{component}/settings`
	//   * `projects/{project}/locations/{location}/clusters/{cluster}/components/{component}/settings`
	//   * `projects/{project}/regions/{region}/clusters/{cluster}/components/{component}/settings`
	//   * `projects/{project}/zones/{zone}/clusters/{cluster}/components/{component}/settings`
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.name
	Name *string `json:"name,omitempty"`

	// ENABLE to enable component, DISABLE to disable and INHERIT to inherit
	//  setting from ancestors.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.state
	State *string `json:"state,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Container Threate Detection specific settings
	//  For component, expect CONTAINER_THREAT_DETECTION
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.container_threat_detection_settings
	ContainerThreatDetectionSettings *ContainerThreatDetectionSettings `json:"containerThreatDetectionSettings,omitempty"`

	// Event Threat Detection specific settings
	//  For component, expect EVENT_THREAT_DETECTION
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.event_threat_detection_settings
	EventThreatDetectionSettings *EventThreatDetectionSettings `json:"eventThreatDetectionSettings,omitempty"`

	// Security Health Analytics specific settings
	//  For component, expect SECURITY_HEALTH_ANALYTICS
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.security_health_analytics_settings
	SecurityHealthAnalyticsSettings *SecurityHealthAnalyticsSettings `json:"securityHealthAnalyticsSettings,omitempty"`

	// Web Security Scanner specific settings
	//  For component, expect WEB_SECURITY_SCANNER
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.web_security_scanner_settings
	WebSecurityScannerSettings *WebSecurityScanner `json:"webSecurityScannerSettings,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.DetectorSettings
type ComponentSettings_DetectorSettings struct {
	// ENABLE to enable component, DISABLE to disable and INHERIT to inherit
	//  setting from ancestors.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ComponentSettings.DetectorSettings.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.ContainerThreatDetectionSettings
type ContainerThreatDetectionSettings struct {
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.EventThreatDetectionSettings
type EventThreatDetectionSettings struct {
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings
type SecurityHealthAnalyticsSettings struct {
	// Settings for "NON_ORG_IAM_MEMBER" scanner.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.non_org_iam_member_settings
	NonOrgIamMemberSettings *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings `json:"nonOrgIamMemberSettings,omitempty"`

	// Settings for "ADMIN_SERVICE_ACCOUNT" scanner.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.admin_service_account_settings
	AdminServiceAccountSettings *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings `json:"adminServiceAccountSettings,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.AdminServiceAccountSettings
type SecurityHealthAnalyticsSettings_AdminServiceAccountSettings struct {
	// User-created service accounts ending in the provided identities are
	//  allowed to have Admin, Owner or Editor roles granted to them. Otherwise
	//  a finding will be created.
	//  A valid identity can be:
	//    *  a partilly specified service account that starts with "@", e.g.
	//    "@myproject.iam.gserviceaccount.com". This approves all the service
	//    accounts suffixed with the specified identity.
	//    *  a fully specified service account that does not start with "@", e.g.
	//    "myadmin@myproject.iam.gserviceaccount.com".
	//  Google-created service accounts are all approved.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.AdminServiceAccountSettings.approved_identities
	ApprovedIdentities []string `json:"approvedIdentities,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.NonOrgIamMemberSettings
type SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings struct {
	// User emails ending in the provided identities are allowed to have IAM
	//  permissions on a project or the organization. Otherwise a finding will
	//  be created.
	//  A valid identity can be:
	//    *  a domain that starts with "@", e.g. "@yourdomain.com".
	//    *  a fully specified email address that does not start with "@", e.g.
	//    "abc@gmail.com"
	//  Regular expressions are not supported.
	//  Service accounts are not examined by the scanner and will be omitted if
	//  added to the list.
	//  If not specified, only Gmail accounts will be considered as non-approved.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.NonOrgIamMemberSettings.approved_identities
	ApprovedIdentities []string `json:"approvedIdentities,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.Settings
type Settings struct {
	// The relative resource name of the settings resource.
	//  Formats:
	//   * `organizations/{organization}/settings`
	//   * `folders/{folder}/settings`
	//   * `projects/{project}/settings`
	//   * `projects/{project}/locations/{location}/clusters/{cluster}/settings`
	//   * `projects/{project}/regions/{region}/clusters/{cluster}/settings`
	//   * `projects/{project}/zones/{zone}/clusters/{cluster}/settings`
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.name
	Name *string `json:"name,omitempty"`

	// Billing settings
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.billing_settings
	BillingSettings *BillingSettings `json:"billingSettings,omitempty"`

	// An enum representing the current on boarding state of SCC.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.state
	State *string `json:"state,omitempty"`

	// Sink settings.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.sink_settings
	SinkSettings *SinkSettings `json:"sinkSettings,omitempty"`

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message


	// A fingerprint used for optimistic concurrency. If none is provided
	//  on updates then the existing metadata will be blindly overwritten.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettings
type Settings_DetectorGroupSettings struct {
	// The state determines if the group is enabled or not.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettings.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.SinkSettings
type SinkSettings struct {
	// The resource name of the project to send logs to. This project must be
	//  part of the same organization where the Security Center API is
	//  enabled. The format is `projects/{project}`. If it is empty, we do
	//  not output logs. If a project ID is provided it will be normalized to a
	//  project number.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.SinkSettings.logging_sink_project
	LoggingSinkProject *string `json:"loggingSinkProject,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.WebSecurityScanner
type WebSecurityScanner struct {
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.BillingSettings
type BillingSettingsObservedState struct {
	// Output only. Billing tier selected by customer
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.BillingSettings.billing_tier
	BillingTier *string `json:"billingTier,omitempty"`

	// Output only. Type of billing method
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.BillingSettings.billing_type
	BillingType *string `json:"billingType,omitempty"`

	// Output only. The absolute point in time when the subscription became effective.
	//  Can be compared to expire_time value to determine full contract duration
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.BillingSettings.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The absolute point in time when the subscription expires.
	//
	//  If this field is populated and billing_tier is STANDARD, this is
	//  indication of a point in the _past_ when a PREMIUM access ended.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.BillingSettings.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.Settings
type SettingsObservedState struct {
	// Billing settings
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.billing_settings
	BillingSettings *BillingSettingsObservedState `json:"billingSettings,omitempty"`

	// Output only. The organization-level service account to be used for security center
	//  components. The component must have permission to "act as" the service
	//  account.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.org_service_account
	OrgServiceAccount *string `json:"orgServiceAccount,omitempty"`

	// Output only. The time these settings were last updated.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.Settings.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
