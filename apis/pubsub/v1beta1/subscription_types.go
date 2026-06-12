// Copyright 2026 Google LLC
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
// See the License for the1 specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BigQueryTableRef struct {
	// Allowed value: string of the format `{{project}}.{{dataset_id}}.{{value}}`, where {{value}} is the `name` field of a `BigQueryTable` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

type StorageBucketRef struct {
	// Allowed value: The `name` field of a `StorageBucket` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.pubsub.v1.BigQueryConfig
type BigQueryConfig struct {
	// When true and useTopicSchema is true, any fields that are a part of the topic schema that are not part of the BigQuery table schema are dropped when writing to BigQuery.
	// Otherwise, the schemas must be kept in sync and any messages with extra fields are not written and remain in the subscription's backlog.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.drop_unknown_fields
	DropUnknownFields *bool `json:"dropUnknownFields,omitempty"`

	// The name of the table to which to write data.
	// +required
	TableRef *BigQueryTableRef `json:"tableRef"`

	// When true, use the topic's schema as the columns to write to in BigQuery, if it exists.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.use_topic_schema
	UseTopicSchema *bool `json:"useTopicSchema,omitempty"`

	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey to additional columns in the table.
	// The subscription name, messageId, and publishTime fields are put in their own columns while all other message properties (other than data) are written to a JSON object in the attributes column.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.write_metadata
	WriteMetadata *bool `json:"writeMetadata,omitempty"`
}

// +kcc:proto=google.pubsub.v1.CloudStorageConfig.AvroConfig
type CloudStorageConfigAvroConfig struct {
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey as additional fields in the output.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.AvroConfig.write_metadata
	WriteMetadata *bool `json:"writeMetadata,omitempty"`
}

// +kcc:proto=google.pubsub.v1.CloudStorageConfig
type CloudStorageConfig struct {
	// If set, message data will be written to Cloud Storage in Avro format.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.avro_config
	AvroConfig *CloudStorageConfigAvroConfig `json:"avroConfig,omitempty"`

	// User-provided name for the Cloud Storage bucket. The bucket must be created by the user. The bucket name must be without any prefix like "gs://".
	// +required
	BucketRef *StorageBucketRef `json:"bucketRef"`

	// User-provided prefix for Cloud Storage filename.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.filename_prefix
	FilenamePrefix *string `json:"filenamePrefix,omitempty"`

	// User-provided suffix for Cloud Storage filename. Must not end in "/".
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.filename_suffix
	FilenameSuffix *string `json:"filenameSuffix,omitempty"`

	// The maximum bytes that can be written to a Cloud Storage file before a new file is created. Min 1 KB, max 10 GiB. The maxBytes limit may be exceeded in cases where messages are larger than the limit.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.max_bytes
	MaxBytes *int64 `json:"maxBytes,omitempty"`

	// The maximum duration that can elapse before a new Cloud Storage file is created. Min 1 minute, max 10 minutes, default 5 minutes. May not exceed the subscription's acknowledgement deadline. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.max_duration
	MaxDuration *string `json:"maxDuration,omitempty"`

	// An output-only field that indicates whether or not the subscription can receive messages.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.DeadLetterPolicy
type DeadLetterPolicy struct {
	// The name of the topic to which dead letter messages should be published.
	DeadLetterTopicRef *PubSubTopicRef `json:"deadLetterTopicRef,omitempty"`

	// The maximum number of delivery attempts for any message. The value must be between 5 and 100.
	// +kcc:proto:field=google.pubsub.v1.DeadLetterPolicy.max_delivery_attempts
	MaxDeliveryAttempts *int32 `json:"maxDeliveryAttempts,omitempty"`
}

// +kcc:proto=google.pubsub.v1.ExpirationPolicy
type ExpirationPolicy struct {
	// Specifies the "time-to-live" duration for an associated resource. The resource expires if it is not active for a period of ttl. If ttl is set to "", the associated resource never expires. A duration in seconds with up to nine fractional digits, terminated by 's'. Example - "3.5s".
	// +kcc:proto:field=google.pubsub.v1.ExpirationPolicy.ttl
	// +required
	Ttl *string `json:"ttl"`
}

// +kcc:proto=google.pubsub.v1.PushConfig.NoWrapper
type PushConfigNoWrapper struct {
	// When true, writes the Pub/Sub message metadata to 'x-goog-pubsub-<KEY>:<VAL>' headers of the HTTP request. Writes the Pub/Sub message attributes to '<KEY>:<VAL>' headers of the HTTP request.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.NoWrapper.write_metadata
	// +required
	WriteMetadata *bool `json:"writeMetadata"`
}

// +kcc:proto=google.pubsub.v1.PushConfig.OidcToken
type PushConfigOidcToken struct {
	// Audience to be used when generating OIDC token. The audience claim identifies the recipients that the JWT is intended for. The audience value is a single case-sensitive string. Having multiple values (array) for the audience field is not supported. More info about the OIDC JWT token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3 Note: if not specified, the Push endpoint URL will be used.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.OidcToken.audience
	Audience *string `json:"audience,omitempty"`

	// Service account email to be used for generating the OIDC token. The caller (for subscriptions.create, subscriptions.patch, and subscriptions.modifyPushConfig RPCs) must have the iam.serviceAccounts.actAs permission for the service account.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.OidcToken.service_account_email
	// +required
	ServiceAccountEmail *string `json:"serviceAccountEmail"`
}

// +kcc:proto=google.pubsub.v1.PushConfig
type PushConfig struct {
	// Endpoint configuration attributes.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// When set, the payload to the push endpoint is not wrapped.Sets the 'data' field as the HTTP body for delivery.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.no_wrapper
	NoWrapper *PushConfigNoWrapper `json:"noWrapper,omitempty"`

	// If specified, Pub/Sub will generate and attach an OIDC JWT token as an Authorization header in the HTTP request for every pushed message.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.oidc_token
	OidcToken *PushConfigOidcToken `json:"oidcToken,omitempty"`

	// A URL locating the endpoint to which messages should be pushed. For example, a Webhook endpoint might use "https://example.com/push".
	// +kcc:proto:field=google.pubsub.v1.PushConfig.push_endpoint
	// +required
	PushEndpoint *string `json:"pushEndpoint"`
}

// +kcc:proto=google.pubsub.v1.RetryPolicy
type RetryPolicy struct {
	// The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kcc:proto:field=google.pubsub.v1.RetryPolicy.maximum_backoff
	MaximumBackoff *string `json:"maximumBackoff,omitempty"`

	// The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kcc:proto:field=google.pubsub.v1.RetryPolicy.minimum_backoff
	MinimumBackoff *string `json:"minimumBackoff,omitempty"`
}

// PubSubSubscriptionSpec defines the desired state of PubSubSubscription
// +kcc:spec:proto=google.pubsub.v1.Subscription
type PubSubSubscriptionSpec struct {
	// This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the message.
	// +kcc:proto:field=google.pubsub.v1.Subscription.ack_deadline_seconds
	AckDeadlineSeconds *int32 `json:"ackDeadlineSeconds,omitempty"`

	// If delivery to BigQuery is used with this subscription, this field is used to configure it. Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined. If all three are empty, then the subscriber will pull and ack messages using API methods.
	// +kcc:proto:field=google.pubsub.v1.Subscription.bigquery_config
	BigqueryConfig *BigQueryConfig `json:"bigqueryConfig,omitempty"`

	// If delivery to Cloud Storage is used with this subscription, this field is used to configure it. Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined. If all three are empty, then the subscriber will pull and ack messages using API methods.
	// +kcc:proto:field=google.pubsub.v1.Subscription.cloud_storage_config
	CloudStorageConfig *CloudStorageConfig `json:"cloudStorageConfig,omitempty"`

	// A policy that specifies the conditions for dead lettering messages in this subscription. If dead_letter_policy is not set, dead lettering is disabled.
	// +kcc:proto:field=google.pubsub.v1.Subscription.dead_letter_policy
	DeadLetterPolicy *DeadLetterPolicy `json:"deadLetterPolicy,omitempty"`

	// If 'true', Pub/Sub provides the guarantees for the delivery of a message with a given value of messageId on this Subscriptions'.
	// +kcc:proto:field=google.pubsub.v1.Subscription.enable_exactly_once_delivery
	EnableExactlyOnceDelivery *bool `json:"enableExactlyOnceDelivery,omitempty"`

	// Immutable. If 'true', messages published with the same orderingKey in PubsubMessage will be delivered to the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they may be delivered in any order.
	// +kcc:proto:field=google.pubsub.v1.Subscription.enable_message_ordering
	EnableMessageOrdering *bool `json:"enableMessageOrdering,omitempty"`

	// A policy that specifies the conditions for this subscription's expiration.
	// +kcc:proto:field=google.pubsub.v1.Subscription.expiration_policy
	ExpirationPolicy *ExpirationPolicy `json:"expirationPolicy,omitempty"`

	// Immutable. The subscription only delivers the messages that match the filter.
	// +kcc:proto:field=google.pubsub.v1.Subscription.filter
	Filter *string `json:"filter,omitempty"`

	// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published.
	// +kcc:proto:field=google.pubsub.v1.Subscription.message_retention_duration
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty"`

	// If push delivery is used with this subscription, this field is used to configure it. An empty pushConfig signifies that the subscriber will pull and ack messages using API methods.
	// +kcc:proto:field=google.pubsub.v1.Subscription.push_config
	PushConfig *PushConfig `json:"pushConfig,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Indicates whether to retain acknowledged messages. If 'true', then messages are not expunged from the subscription's backlog, even if they are acknowledged, until they fall out of the messageRetentionDuration window.
	// +kcc:proto:field=google.pubsub.v1.Subscription.retain_acked_messages
	RetainAckedMessages *bool `json:"retainAckedMessages,omitempty"`

	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// +kcc:proto:field=google.pubsub.v1.Subscription.retry_policy
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`

	// Reference to a PubSubTopic.
	// +required
	TopicRef *PubSubTopicRef `json:"topicRef"`
}

// PubSubSubscriptionStatus defines the config connector machine state of PubSubSubscription
type PubSubSubscriptionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsubsubscription;gcppubsubsubscriptions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubSubscription is the Schema for the PubSubSubscription API
// +k8s:openapi-gen=true
type PubSubSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PubSubSubscriptionSpec   `json:"spec,omitempty"`
	Status PubSubSubscriptionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PubSubSubscriptionList contains a list of PubSubSubscription
type PubSubSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubSubscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubSubscription{}, &PubSubSubscriptionList{})
}
