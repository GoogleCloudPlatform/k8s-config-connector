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


// +kcc:proto=google.cloud.securitycenter.v2.NotificationConfig
type NotificationConfig struct {
	// Identifier. The relative resource name of this notification config. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  The following list shows some examples:
	//  +
	//  `organizations/{organization_id}/locations/{location_id}/notificationConfigs/notify_public_bucket`
	//  +
	//  `folders/{folder_id}/locations/{location_id}/notificationConfigs/notify_public_bucket`
	//  +
	//  `projects/{project_id}/locations/{location_id}/notificationConfigs/notify_public_bucket`
	// +kcc:proto:field=google.cloud.securitycenter.v2.NotificationConfig.name
	Name *string `json:"name,omitempty"`

	// The description of the notification config (max of 1024 characters).
	// +kcc:proto:field=google.cloud.securitycenter.v2.NotificationConfig.description
	Description *string `json:"description,omitempty"`

	// The Pub/Sub topic to send notifications to. Its format is
	//  "projects/[project_id]/topics/[topic]".
	// +kcc:proto:field=google.cloud.securitycenter.v2.NotificationConfig.pubsub_topic
	PubsubTopic *string `json:"pubsubTopic,omitempty"`

	// The config for triggering streaming-based notifications.
	// +kcc:proto:field=google.cloud.securitycenter.v2.NotificationConfig.streaming_config
	StreamingConfig *NotificationConfig_StreamingConfig `json:"streamingConfig,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.NotificationConfig.StreamingConfig
type NotificationConfig_StreamingConfig struct {
	// Expression that defines the filter to apply across create/update events
	//  of assets or findings as specified by the event type. The expression is a
	//  list of zero or more restrictions combined via logical operators `AND`
	//  and `OR`. Parentheses are supported, and `OR` has higher precedence than
	//  `AND`.
	//
	//  Restrictions have the form `<field> <operator> <value>` and may have a
	//  `-` character in front of them to indicate negation. The fields map to
	//  those defined in the corresponding resource.
	//
	//  The supported operators are:
	//
	//  * `=` for all value types.
	//  * `>`, `<`, `>=`, `<=` for integer values.
	//  * `:`, meaning substring matching, for strings.
	//
	//  The supported value types are:
	//
	//  * string literals in quotes.
	//  * integer literals without quotes.
	//  * boolean literals `true` and `false` without quotes.
	// +kcc:proto:field=google.cloud.securitycenter.v2.NotificationConfig.StreamingConfig.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.NotificationConfig
type NotificationConfigObservedState struct {
	// Output only. The service account that needs "pubsub.topics.publish"
	//  permission to publish to the Pub/Sub topic.
	// +kcc:proto:field=google.cloud.securitycenter.v2.NotificationConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}
