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


// +kcc:proto=google.cloud.advisorynotifications.v1.Attachment
type Attachment struct {
	// A CSV file attachment. Max size is 10 MB.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Attachment.csv
	Csv *Csv `json:"csv,omitempty"`

	// The title of the attachment.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Attachment.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Csv
type Csv struct {
	// The list of headers for data columns in a CSV file.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Csv.headers
	Headers []string `json:"headers,omitempty"`

	// The list of data rows in a CSV file, as string arrays rather than as a
	//  single comma-separated string.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Csv.data_rows
	DataRows []Csv_CsvRow `json:"dataRows,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Csv.CsvRow
type Csv_CsvRow struct {
	// The data entries in a CSV file row, as a string array rather than a
	//  single comma-separated string.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Csv.CsvRow.entries
	Entries []string `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Message
type Message struct {
	// The message content.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Message.body
	Body *Message_Body `json:"body,omitempty"`

	// The attachments to download.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Message.attachments
	Attachments []Attachment `json:"attachments,omitempty"`

	// The Message creation timestamp.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Message.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Time when Message was localized
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Message.localization_time
	LocalizationTime *string `json:"localizationTime,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Message.Body
type Message_Body struct {
	// The text content of the message body.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Message.Body.text
	Text *Text `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Notification
type Notification struct {
	// The resource name of the notification.
	//  Format:
	//  organizations/{organization}/locations/{location}/notifications/{notification}
	//  or projects/{project}/locations/{location}/notifications/{notification}.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Notification.name
	Name *string `json:"name,omitempty"`

	// The subject line of the notification.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Notification.subject
	Subject *Subject `json:"subject,omitempty"`

	// A list of messages in the notification.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Notification.messages
	Messages []Message `json:"messages,omitempty"`

	// Type of notification
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Notification.notification_type
	NotificationType *string `json:"notificationType,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Subject
type Subject struct {
	// The text content.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Subject.text
	Text *Text `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Text
type Text struct {
	// The English copy.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Text.en_text
	EnText *string `json:"enText,omitempty"`

	// The requested localized copy (if applicable).
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Text.localized_text
	LocalizedText *string `json:"localizedText,omitempty"`

	// Status of the localization.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Text.localization_state
	LocalizationState *string `json:"localizationState,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Notification
type NotificationObservedState struct {
	// Output only. Time the notification was created.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Notification.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
