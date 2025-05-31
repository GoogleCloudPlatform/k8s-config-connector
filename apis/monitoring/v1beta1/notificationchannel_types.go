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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringNotificationChannelGVK = GroupVersion.WithKind("MonitoringNotificationChannel")

// MonitoringNotificationChannelSpec defines the desired state of MonitoringNotificationChannel
// +kcc:proto=google.monitoring.v3.NotificationChannel
type MonitoringNotificationChannelSpec struct {
	// Immutable. Optional. The service-generated name of the resource.
	// Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// The type of the notification channel. This field matches the
	//  value of the
	//  [NotificationChannelDescriptor.type][google.monitoring.v3.NotificationChannelDescriptor.type]
	//  field.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.type
	Type *string `json:"type,omitempty"`

	/* IDENTITY
	// Identifier. The full REST resource name for this channel. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
	//
	//  The `[CHANNEL_ID]` is automatically assigned by the server on creation.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.name
	Name *string `json:"name,omitempty"`
	*/

	/* NOTYET - TERRAFORM BACKCOMPAT
	// An optional human-readable name for this notification channel. It is
	//  recommended that you specify a non-empty and unique name in order to
	//  make it easier to identify the channels in your project, though this is
	//  not enforced. The display name is limited to 512 Unicode characters.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.display_name
	DisplayName *string `json:"displayName,omitempty"`
	*/

	// An optional human-readable description of this notification channel. This
	//  description may provide additional details, beyond the display
	//  name, for the channel. This may not exceed 1024 Unicode characters.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.description
	Description *string `json:"description,omitempty"`

	// Whether notifications are forwarded to the described channel. This makes
	//  it possible to disable delivery of notifications to a particular channel
	//  without removing the channel from all alerting policies that reference
	//  the channel. This is a more convenient approach when the change is
	//  temporary and you want to receive notifications from the same set
	//  of alerting policies on the channel at some point in the future.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Configuration fields that define the channel and its behavior. The
	//  permissible and required labels are specified in the
	//  [NotificationChannelDescriptor.labels][google.monitoring.v3.NotificationChannelDescriptor.labels]
	//  of the `NotificationChannelDescriptor` corresponding to the `type` field.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.labels
	Labels map[string]string `json:"labels,omitempty"`

	/* NOTYET - TERRAFORM BACKCOMPAT

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
	*/

	// If true, the notification channel will be deleted regardless
	// of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still
	// referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete *bool `json:"forceDelete,omitempty"`

	// Different notification type behaviors are configured primarily using the the 'labels' field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the 'labels' map in the api request.
	//
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	SensitiveLabels *SensitiveLabels `json:"sensitiveLabels,omitempty"`
}

// SensitiveLabels is for backwards-compatibility with terraform.
type SensitiveLabels struct {
	// An authorization token for a notification channel.
	// Channel types that support this field include: slack
	AuthToken *secret.Legacy `json:"authToken,omitempty"`

	// A password for a notification channel.
	// Channel types that support this field include: webhook_basicauth
	Password *secret.Legacy `json:"password,omitempty"`

	//  A servicekey token for a notification channel.
	// Channel types that support this field include: pagerduty
	ServiceKey *secret.Legacy `json:"serviceKey,omitempty"`
}

// MonitoringNotificationChannelStatus defines the config connector machine state of MonitoringNotificationChannel
// +kcc:proto=google.monitoring.v3.NotificationChannel
type MonitoringNotificationChannelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The full REST resource name for this channel
	// The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	// The [CHANNEL_ID] is automatically assigned by the server on creation.
	Name *string `json:"name,omitempty"`

	/* NOTYET - TERRAFORM BACKCOMPAT
	// A unique specifier for the MonitoringNotificationChannel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/* NOTYET - TERRAFORM BACKCOMPAT
	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MonitoringNotificationChannelObservedState `json:"observedState,omitempty"`
	*/

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
}

// MonitoringNotificationChannelObservedState is the state of the MonitoringNotificationChannel resource as most recently observed in GCP.
// +kcc:proto=google.monitoring.v3.NotificationChannel
type MonitoringNotificationChannelObservedState struct {

	// Record of the creation of this channel.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.creation_record
	CreationRecord *MutationRecord `json:"creationRecord,omitempty"`

	// Records of the modification of this channel.
	// +kcc:proto:field=google.monitoring.v3.NotificationChannel.mutation_records
	MutationRecords []MutationRecord `json:"mutationRecords,omitempty"`

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
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringnotificationchannel;gcpmonitoringnotificationchannels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringNotificationChannel is the Schema for the MonitoringNotificationChannel API
// +k8s:openapi-gen=true
type MonitoringNotificationChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringNotificationChannelSpec   `json:"spec,omitempty"`
	Status MonitoringNotificationChannelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringNotificationChannelList contains a list of MonitoringNotificationChannel
type MonitoringNotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringNotificationChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringNotificationChannel{}, &MonitoringNotificationChannelList{})
}
