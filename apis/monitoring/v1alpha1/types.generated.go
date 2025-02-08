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


// +kcc:proto=google.api.LabelDescriptor
type LabelDescriptor struct {
	// The label key.
	// +kcc:proto:field=google.api.LabelDescriptor.key
	Key *string `json:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// +kcc:proto:field=google.api.LabelDescriptor.value_type
	ValueType *string `json:"valueType,omitempty"`

	// A human-readable description for the label.
	// +kcc:proto:field=google.api.LabelDescriptor.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.monitoring.v3.NotificationChannelDescriptor
type NotificationChannelDescriptor struct {
	// The full REST resource name for this descriptor. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/notificationChannelDescriptors/[TYPE]
	//
	//  In the above, `[TYPE]` is the value of the `type` field.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.name
	Name *string `json:"name,omitempty"`

	// The type of notification channel, such as "email" and "sms". To view the
	//  full list of channels, see
	//  [Channel
	//  descriptors](https://cloud.google.com/monitoring/alerts/using-channels-api#ncd).
	//  Notification channel types are globally unique.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.type
	Type *string `json:"type,omitempty"`

	// A human-readable name for the notification channel type.  This
	//  form of the name is suitable for a user interface.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A human-readable description of the notification channel
	//  type. The description may include a description of the properties
	//  of the channel and pointers to external documentation.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.description
	Description *string `json:"description,omitempty"`

	// The set of labels that must be defined to identify a particular
	//  channel of the corresponding type. Each label includes a
	//  description for how that field should be populated.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.labels
	Labels []LabelDescriptor `json:"labels,omitempty"`

	// The tiers that support this notification channel; the project service tier
	//  must be one of the supported_tiers.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.supported_tiers
	SupportedTiers []string `json:"supportedTiers,omitempty"`

	// The product launch stage for channels of this type.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannelDescriptor.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`
}
