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


// +kcc:proto=google.cloud.support.v2.Actor
type Actor struct {
	// The name to display for the actor. If not provided, it is inferred from
	//  credentials supplied during case creation. When an email is provided, a
	//  display name must also be provided. This will be obfuscated if the user
	//  is a Google Support agent.
	// +kcc:proto:field=google.cloud.support.v2.Actor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The email address of the actor. If not provided, it is inferred from
	//  credentials supplied during case creation. If the authenticated principal
	//  does not have an email address, one must be provided. When a name is
	//  provided, an email must also be provided. This will be obfuscated if the
	//  user is a Google Support agent.
	// +kcc:proto:field=google.cloud.support.v2.Actor.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.Case
type Case struct {
	// The resource name for the case.
	// +kcc:proto:field=google.cloud.support.v2.Case.name
	Name *string `json:"name,omitempty"`

	// The short summary of the issue reported in this case.
	// +kcc:proto:field=google.cloud.support.v2.Case.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A broad description of the issue.
	// +kcc:proto:field=google.cloud.support.v2.Case.description
	Description *string `json:"description,omitempty"`

	// The issue classification applicable to this case.
	// +kcc:proto:field=google.cloud.support.v2.Case.classification
	Classification *CaseClassification `json:"classification,omitempty"`

	// The timezone of the user who created the support case.
	//  It should be in a format IANA recognizes: https://www.iana.org/time-zones.
	//  There is no additional validation done by the API.
	// +kcc:proto:field=google.cloud.support.v2.Case.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// The email addresses to receive updates on this case.
	// +kcc:proto:field=google.cloud.support.v2.Case.subscriber_email_addresses
	SubscriberEmailAddresses []string `json:"subscriberEmailAddresses,omitempty"`

	// The user who created the case.
	//
	//  Note: The name and email will be obfuscated if the case was created by
	//  Google Support.
	// +kcc:proto:field=google.cloud.support.v2.Case.creator
	Creator *Actor `json:"creator,omitempty"`

	// A user-supplied email address to send case update notifications for. This
	//  should only be used in BYOID flows, where we cannot infer the user's email
	//  address directly from their EUCs.
	// +kcc:proto:field=google.cloud.support.v2.Case.contact_email
	ContactEmail *string `json:"contactEmail,omitempty"`

	// Whether the case is currently escalated.
	// +kcc:proto:field=google.cloud.support.v2.Case.escalated
	Escalated *bool `json:"escalated,omitempty"`

	// Whether this case was created for internal API testing and should not be
	//  acted on by the support team.
	// +kcc:proto:field=google.cloud.support.v2.Case.test_case
	TestCase *bool `json:"testCase,omitempty"`

	// The language the user has requested to receive support in. This should be a
	//  BCP 47 language code (e.g., `"en"`, `"zh-CN"`, `"zh-TW"`, `"ja"`, `"ko"`).
	//  If no language or an unsupported language is specified, this field defaults
	//  to English (en).
	//
	//  Language selection during case creation may affect your available support
	//  options. For a list of supported languages and their support working hours,
	//  see: https://cloud.google.com/support/docs/language-working-hours
	// +kcc:proto:field=google.cloud.support.v2.Case.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The priority of this case.
	// +kcc:proto:field=google.cloud.support.v2.Case.priority
	Priority *string `json:"priority,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.CaseClassification
type CaseClassification struct {
	// The unique ID for a classification. Must be specified for case creation.
	//
	//  To retrieve valid classification IDs for case creation, use
	//  `caseClassifications.search`.
	// +kcc:proto:field=google.cloud.support.v2.CaseClassification.id
	ID *string `json:"id,omitempty"`

	// The display name of the classification.
	// +kcc:proto:field=google.cloud.support.v2.CaseClassification.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.Actor
type ActorObservedState struct {
	// Output only. Whether the actor is a Google support actor.
	// +kcc:proto:field=google.cloud.support.v2.Actor.google_support
	GoogleSupport *bool `json:"googleSupport,omitempty"`
}

// +kcc:proto=google.cloud.support.v2.Case
type CaseObservedState struct {
	// Output only. The current status of the support case.
	// +kcc:proto:field=google.cloud.support.v2.Case.state
	State *string `json:"state,omitempty"`

	// Output only. The time this case was created.
	// +kcc:proto:field=google.cloud.support.v2.Case.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this case was last updated.
	// +kcc:proto:field=google.cloud.support.v2.Case.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The user who created the case.
	//
	//  Note: The name and email will be obfuscated if the case was created by
	//  Google Support.
	// +kcc:proto:field=google.cloud.support.v2.Case.creator
	Creator *ActorObservedState `json:"creator,omitempty"`
}
