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

// +kcc:proto=google.monitoring.v3.MutationRecord
type MutationRecord struct {
	// When the change occurred.
	// +kcc:proto:field=google.monitoring.v3.MutationRecord.mutate_time
	MutateTime *string `json:"mutateTime,omitempty"`

	// The email address of the user making the change.
	// +kcc:proto:field=google.monitoring.v3.MutationRecord.mutated_by
	MutatedBy *string `json:"mutatedBy,omitempty"`
}

// +kcc:proto=google.monitoring.v3.NotificationChannel
type NotificationChannel struct {
	// The type of the notification channel. This field matches the
	//  value of the
	//  [NotificationChannelDescriptor.type][google.monitoring.v3.NotificationChannelDescriptor.type]
	//  field.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.type
	Type *string `json:"type,omitempty"`

	// Identifier. The full REST resource name for this channel. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
	//
	//  The `[CHANNEL_ID]` is automatically assigned by the server on creation.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.name
	Name *string `json:"name,omitempty"`

	// An optional human-readable name for this notification channel. It is
	//  recommended that you specify a non-empty and unique name in order to
	//  make it easier to identify the channels in your project, though this is
	//  not enforced. The display name is limited to 512 Unicode characters.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// An optional human-readable description of this notification channel. This
	//  description may provide additional details, beyond the display
	//  name, for the channel. This may not exceed 1024 Unicode characters.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.description
	Description *string `json:"description,omitempty"`

	// Configuration fields that define the channel and its behavior. The
	//  permissible and required labels are specified in the
	//  [NotificationChannelDescriptor.labels][google.monitoring.v3.NotificationChannelDescriptor.labels]
	//  of the `NotificationChannelDescriptor` corresponding to the `type` field.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.labels
	Labels map[string]string `json:"labels,omitempty"`

	// User-supplied key/value data that does not need to conform to
	//  the corresponding `NotificationChannelDescriptor`'s schema, unlike
	//  the `labels` field. This field is intended to be used for organizing
	//  and identifying the `NotificationChannel` objects.
	//
	//  The field can contain up to 64 entries. Each key and value is limited to
	//  63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	//  values can contain only lowercase letters, numerals, underscores, and
	//  dashes. Keys must begin with a letter.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// Indicates whether this channel has been verified or not. On a
	//  [`ListNotificationChannels`][google.monitoring.v3.NotificationChannelService.ListNotificationChannels]
	//  or
	//  [`GetNotificationChannel`][google.monitoring.v3.NotificationChannelService.GetNotificationChannel]
	//  operation, this field is expected to be populated.
	//
	//  If the value is `UNVERIFIED`, then it indicates that the channel is
	//  non-functioning (it both requires verification and lacks verification);
	//  otherwise, it is assumed that the channel works.
	//
	//  If the channel is neither `VERIFIED` nor `UNVERIFIED`, it implies that
	//  the channel is of a type that does not require verification or that
	//  this specific channel has been exempted from verification because it was
	//  created prior to verification being required for channels of this type.
	//
	//  This field cannot be modified using a standard
	//  [`UpdateNotificationChannel`][google.monitoring.v3.NotificationChannelService.UpdateNotificationChannel]
	//  operation. To change the value of this field, you must call
	//  [`VerifyNotificationChannel`][google.monitoring.v3.NotificationChannelService.VerifyNotificationChannel].
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.verification_status
	VerificationStatus *string `json:"verificationStatus,omitempty"`

	// Whether notifications are forwarded to the described channel. This makes
	//  it possible to disable delivery of notifications to a particular channel
	//  without removing the channel from all alerting policies that reference
	//  the channel. This is a more convenient approach when the change is
	//  temporary and you want to receive notifications from the same set
	//  of alerting policies on the channel at some point in the future.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Record of the creation of this channel.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.creation_record
	CreationRecord *MutationRecord `json:"creationRecord,omitempty"`

	// Records of the modification of this channel.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.mutation_records
	MutationRecords []MutationRecord `json:"mutationRecords,omitempty"`
}
