// Copyright 2024 Google LLC
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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageNotificationGVK = GroupVersion.WithKind("StorageNotification")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageNotificationSpec defines the desired state of StorageNotification
// +kcc:proto=google.storage.v1.Notification
type StorageNotificationSpec struct {
	// The StorageNotification name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`

	StorageBucketRef *refv1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// The Cloud PubSub topic to which this subscription publishes. Formatted as:
	//  '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	PubSubTopicRef *refv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`

	// If present, only send notifications about listed event types. If empty,
	//  sent notifications for all event types.
	EventTypes []string `json:"eventTypes,omitempty"`

	// An optional list of additional attributes to attach to each Cloud PubSub
	//  message published for this notification subscription.
	CustomAttributes map[string]string `json:"customAttributes,omitempty"`

	// If present, only apply this notification configuration to object names that
	//  begin with this prefix.
	ObjectNamePrefix *string `json:"objectNamePrefix,omitempty"`

	// The desired content of the Payload.
	PayloadFormat *string `json:"payloadFormat,omitempty"`
}

// StorageNotificationStatus defines the config connector machine state of StorageNotification
type StorageNotificationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageNotification resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *StorageNotificationObservedState `json:"observedState,omitempty"`

	// The SelfLink and NotificatoinID are not from the storage GCP service.

	// Deprecated: use the `.status.externalRef` to identify the resource.
	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty"`
	// Deprecated: use the `.status.externalRef` to identify the resource.
	// The ID of the created notification.
	NotificationId *string `json:"notificationId,omitempty"`
}

// StorageNotificationSpec defines the desired state of StorageNotification
// +kcc:proto=google.storage.v1.Notification
type StorageNotificationObservedState struct {

	// The Cloud PubSub topic to which this subscription publishes. Formatted as:
	//  '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic *string `json:"topic,omitempty"`

	// HTTP 1.1 [https://tools.ietf.org/html/rfc7232#section-2.3][Entity tag]
	//  for this subscription notification.
	Etag *string `json:"etag,omitempty"`

	// The ID of the notification.
	Id *string `json:"id,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageNotification is the Schema for the StorageNotification API
// +k8s:openapi-gen=true
type StorageNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageNotificationSpec   `json:"spec,omitempty"`
	Status StorageNotificationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageNotificationList contains a list of StorageNotification
type StorageNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageNotification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageNotification{}, &StorageNotificationList{})
}
