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

// +kcc:proto=google.cloud.essentialcontacts.v1.Contact
type Contact struct {

	// Required. The email address to send notifications to. The email address
	//  does not need to be a Google Account.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.email
	Email *string `json:"email,omitempty"`

	// Required. The categories of notifications that the contact will receive
	//  communications for.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.notification_category_subscriptions
	NotificationCategorySubscriptions []string `json:"notificationCategorySubscriptions,omitempty"`

	// Required. The preferred language for notifications, as a ISO 639-1 language
	//  code. See [Supported
	//  languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages)
	//  for a list of supported languages.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.language_tag
	LanguageTag *string `json:"languageTag,omitempty"`

	// The last time the validation_state was updated, either manually or
	//  automatically. A contact is considered stale if its validation state was
	//  updated more than 1 year ago.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.validate_time
	ValidateTime *string `json:"validateTime,omitempty"`
}

// +kcc:proto=google.cloud.essentialcontacts.v1.Contact
type ContactObservedState struct {
	// Output only. The identifier for the contact.
	//  Format: {resource_type}/{resource_id}/contacts/{contact_id}
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.name
	Name *string `json:"name,omitempty"`

	// Output only. The validity of the contact. A contact is considered valid if
	//  it is the correct recipient for notifications for a particular resource.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.validation_state
	ValidationState *string `json:"validationState,omitempty"`
}
