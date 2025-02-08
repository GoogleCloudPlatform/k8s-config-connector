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


// +kcc:proto=google.pubsub.v1.BigQueryConfig
type BigQueryConfig struct {
	// Optional. The name of the table to which to write data, of the form
	//  {projectId}.{datasetId}.{tableId}
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.table
	Table *string `json:"table,omitempty"`

	// Optional. When true, use the topic's schema as the columns to write to in
	//  BigQuery, if it exists. `use_topic_schema` and `use_table_schema` cannot be
	//  enabled at the same time.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.use_topic_schema
	UseTopicSchema *bool `json:"useTopicSchema,omitempty"`

	// Optional. When true, write the subscription name, message_id, publish_time,
	//  attributes, and ordering_key to additional columns in the table. The
	//  subscription name, message_id, and publish_time fields are put in their own
	//  columns while all other message properties (other than data) are written to
	//  a JSON object in the attributes column.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.write_metadata
	WriteMetadata *bool `json:"writeMetadata,omitempty"`

	// Optional. When true and use_topic_schema is true, any fields that are a
	//  part of the topic schema that are not part of the BigQuery table schema are
	//  dropped when writing to BigQuery. Otherwise, the schemas must be kept in
	//  sync and any messages with extra fields are not written and remain in the
	//  subscription's backlog.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.drop_unknown_fields
	DropUnknownFields *bool `json:"dropUnknownFields,omitempty"`

	// Optional. When true, use the BigQuery table's schema as the columns to
	//  write to in BigQuery. `use_table_schema` and `use_topic_schema` cannot be
	//  enabled at the same time.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.use_table_schema
	UseTableSchema *bool `json:"useTableSchema,omitempty"`

	// Optional. The service account to use to write to BigQuery. The subscription
	//  creator or updater that specifies this field must have
	//  `iam.serviceAccounts.actAs` permission on the service account. If not
	//  specified, the Pub/Sub [service
	//  agent](https://cloud.google.com/iam/docs/service-agents),
	//  service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`
}

// +kcc:proto=google.pubsub.v1.CloudStorageConfig
type CloudStorageConfig struct {
	// Required. User-provided name for the Cloud Storage bucket.
	//  The bucket must be created by the user. The bucket name must be without
	//  any prefix like "gs://". See the [bucket naming
	//  requirements] (https://cloud.google.com/storage/docs/buckets#naming).
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Optional. User-provided prefix for Cloud Storage filename. See the [object
	//  naming requirements](https://cloud.google.com/storage/docs/objects#naming).
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.filename_prefix
	FilenamePrefix *string `json:"filenamePrefix,omitempty"`

	// Optional. User-provided suffix for Cloud Storage filename. See the [object
	//  naming requirements](https://cloud.google.com/storage/docs/objects#naming).
	//  Must not end in "/".
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.filename_suffix
	FilenameSuffix *string `json:"filenameSuffix,omitempty"`

	// Optional. User-provided format string specifying how to represent datetimes
	//  in Cloud Storage filenames. See the [datetime format
	//  guidance](https://cloud.google.com/pubsub/docs/create-cloudstorage-subscription#file_names).
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.filename_datetime_format
	FilenameDatetimeFormat *string `json:"filenameDatetimeFormat,omitempty"`

	// Optional. If set, message data will be written to Cloud Storage in text
	//  format.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.text_config
	TextConfig *CloudStorageConfig_TextConfig `json:"textConfig,omitempty"`

	// Optional. If set, message data will be written to Cloud Storage in Avro
	//  format.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.avro_config
	AvroConfig *CloudStorageConfig_AvroConfig `json:"avroConfig,omitempty"`

	// Optional. The maximum duration that can elapse before a new Cloud Storage
	//  file is created. Min 1 minute, max 10 minutes, default 5 minutes. May not
	//  exceed the subscription's acknowledgement deadline.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.max_duration
	MaxDuration *string `json:"maxDuration,omitempty"`

	// Optional. The maximum bytes that can be written to a Cloud Storage file
	//  before a new file is created. Min 1 KB, max 10 GiB. The max_bytes limit may
	//  be exceeded in cases where messages are larger than the limit.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.max_bytes
	MaxBytes *int64 `json:"maxBytes,omitempty"`

	// Optional. The maximum number of messages that can be written to a Cloud
	//  Storage file before a new file is created. Min 1000 messages.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.max_messages
	MaxMessages *int64 `json:"maxMessages,omitempty"`

	// Optional. The service account to use to write to Cloud Storage. The
	//  subscription creator or updater that specifies this field must have
	//  `iam.serviceAccounts.actAs` permission on the service account. If not
	//  specified, the Pub/Sub
	//  [service agent](https://cloud.google.com/iam/docs/service-agents),
	//  service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`
}

// +kcc:proto=google.pubsub.v1.CloudStorageConfig.AvroConfig
type CloudStorageConfig_AvroConfig struct {
	// Optional. When true, write the subscription name, message_id,
	//  publish_time, attributes, and ordering_key as additional fields in the
	//  output. The subscription name, message_id, and publish_time fields are
	//  put in their own fields while all other message properties other than
	//  data (for example, an ordering_key, if present) are added as entries in
	//  the attributes map.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.AvroConfig.write_metadata
	WriteMetadata *bool `json:"writeMetadata,omitempty"`

	// Optional. When true, the output Cloud Storage file will be serialized
	//  using the topic schema, if it exists.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.AvroConfig.use_topic_schema
	UseTopicSchema *bool `json:"useTopicSchema,omitempty"`
}

// +kcc:proto=google.pubsub.v1.CloudStorageConfig.TextConfig
type CloudStorageConfig_TextConfig struct {
}

// +kcc:proto=google.pubsub.v1.DeadLetterPolicy
type DeadLetterPolicy struct {
	// Optional. The name of the topic to which dead letter messages should be
	//  published. Format is `projects/{project}/topics/{topic}`.The Pub/Sub
	//  service account associated with the enclosing subscription's parent project
	//  (i.e., service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must
	//  have permission to Publish() to this topic.
	//
	//  The operation will fail if the topic does not exist.
	//  Users should ensure that there is a subscription attached to this topic
	//  since messages published to a topic with no subscriptions are lost.
	// +kcc:proto:field=google.pubsub.v1.DeadLetterPolicy.dead_letter_topic
	DeadLetterTopic *string `json:"deadLetterTopic,omitempty"`

	// Optional. The maximum number of delivery attempts for any message. The
	//  value must be between 5 and 100.
	//
	//  The number of delivery attempts is defined as 1 + (the sum of number of
	//  NACKs and number of times the acknowledgement deadline has been exceeded
	//  for the message).
	//
	//  A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
	//  client libraries may automatically extend ack_deadlines.
	//
	//  This field will be honored on a best effort basis.
	//
	//  If this parameter is 0, a default value of 5 is used.
	// +kcc:proto:field=google.pubsub.v1.DeadLetterPolicy.max_delivery_attempts
	MaxDeliveryAttempts *int32 `json:"maxDeliveryAttempts,omitempty"`
}

// +kcc:proto=google.pubsub.v1.ExpirationPolicy
type ExpirationPolicy struct {
	// Optional. Specifies the "time-to-live" duration for an associated resource.
	//  The resource expires if it is not active for a period of `ttl`. The
	//  definition of "activity" depends on the type of the associated resource.
	//  The minimum and maximum allowed values for `ttl` depend on the type of the
	//  associated resource, as well. If `ttl` is not set, the associated resource
	//  never expires.
	// +kcc:proto:field=google.pubsub.v1.ExpirationPolicy.ttl
	Ttl *string `json:"ttl,omitempty"`
}

// +kcc:proto=google.pubsub.v1.PushConfig
type PushConfig struct {
	// Optional. A URL locating the endpoint to which messages should be pushed.
	//  For example, a Webhook endpoint might use `https://example.com/push`.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.push_endpoint
	PushEndpoint *string `json:"pushEndpoint,omitempty"`

	// Optional. Endpoint configuration attributes that can be used to control
	//  different aspects of the message delivery.
	//
	//  The only currently supported attribute is `x-goog-version`, which you can
	//  use to change the format of the pushed message. This attribute
	//  indicates the version of the data expected by the endpoint. This
	//  controls the shape of the pushed message (i.e., its fields and metadata).
	//
	//  If not present during the `CreateSubscription` call, it will default to
	//  the version of the Pub/Sub API used to make such call. If not present in a
	//  `ModifyPushConfig` call, its value will not be changed. `GetSubscription`
	//  calls will always return a valid version, even if the subscription was
	//  created without this attribute.
	//
	//  The only supported values for the `x-goog-version` attribute are:
	//
	//  * `v1beta1`: uses the push format defined in the v1beta1 Pub/Sub API.
	//  * `v1` or `v1beta2`: uses the push format defined in the v1 Pub/Sub API.
	//
	//  For example:
	//  `attributes { "x-goog-version": "v1" }`
	// +kcc:proto:field=google.pubsub.v1.PushConfig.attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// Optional. If specified, Pub/Sub will generate and attach an OIDC JWT
	//  token as an `Authorization` header in the HTTP request for every pushed
	//  message.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.oidc_token
	OidcToken *PushConfig_OidcToken `json:"oidcToken,omitempty"`

	// Optional. When set, the payload to the push endpoint is in the form of
	//  the JSON representation of a PubsubMessage
	//  (https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#pubsubmessage).
	// +kcc:proto:field=google.pubsub.v1.PushConfig.pubsub_wrapper
	PubsubWrapper *PushConfig_PubsubWrapper `json:"pubsubWrapper,omitempty"`

	// Optional. When set, the payload to the push endpoint is not wrapped.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.no_wrapper
	NoWrapper *PushConfig_NoWrapper `json:"noWrapper,omitempty"`
}

// +kcc:proto=google.pubsub.v1.PushConfig.NoWrapper
type PushConfig_NoWrapper struct {
	// Optional. When true, writes the Pub/Sub message metadata to
	//  `x-goog-pubsub-<KEY>:<VAL>` headers of the HTTP request. Writes the
	//  Pub/Sub message attributes to `<KEY>:<VAL>` headers of the HTTP request.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.NoWrapper.write_metadata
	WriteMetadata *bool `json:"writeMetadata,omitempty"`
}

// +kcc:proto=google.pubsub.v1.PushConfig.OidcToken
type PushConfig_OidcToken struct {
	// Optional. [Service account
	//  email](https://cloud.google.com/iam/docs/service-accounts)
	//  used for generating the OIDC token. For more information
	//  on setting up authentication, see
	//  [Push subscriptions](https://cloud.google.com/pubsub/docs/push).
	// +kcc:proto:field=google.pubsub.v1.PushConfig.OidcToken.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// Optional. Audience to be used when generating OIDC token. The audience
	//  claim identifies the recipients that the JWT is intended for. The
	//  audience value is a single case-sensitive string. Having multiple values
	//  (array) for the audience field is not supported. More info about the OIDC
	//  JWT token audience here:
	//  https://tools.ietf.org/html/rfc7519#section-4.1.3 Note: if not specified,
	//  the Push endpoint URL will be used.
	// +kcc:proto:field=google.pubsub.v1.PushConfig.OidcToken.audience
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.pubsub.v1.PushConfig.PubsubWrapper
type PushConfig_PubsubWrapper struct {
}

// +kcc:proto=google.pubsub.v1.RetryPolicy
type RetryPolicy struct {
	// Optional. The minimum delay between consecutive deliveries of a given
	//  message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
	// +kcc:proto:field=google.pubsub.v1.RetryPolicy.minimum_backoff
	MinimumBackoff *string `json:"minimumBackoff,omitempty"`

	// Optional. The maximum delay between consecutive deliveries of a given
	//  message. Value should be between 0 and 600 seconds. Defaults to 600
	//  seconds.
	// +kcc:proto:field=google.pubsub.v1.RetryPolicy.maximum_backoff
	MaximumBackoff *string `json:"maximumBackoff,omitempty"`
}

// +kcc:proto=google.pubsub.v1.Subscription
type Subscription struct {
	// Required. The name of the subscription. It must have the format
	//  `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must
	//  start with a letter, and contain only letters (`[A-Za-z]`), numbers
	//  (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`),
	//  plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters
	//  in length, and it must not start with `"goog"`.
	// +kcc:proto:field=google.pubsub.v1.Subscription.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the topic from which this subscription is receiving
	//  messages. Format is `projects/{project}/topics/{topic}`. The value of this
	//  field will be `_deleted-topic_` if the topic has been deleted.
	// +kcc:proto:field=google.pubsub.v1.Subscription.topic
	Topic *string `json:"topic,omitempty"`

	// Optional. If push delivery is used with this subscription, this field is
	//  used to configure it.
	// +kcc:proto:field=google.pubsub.v1.Subscription.push_config
	PushConfig *PushConfig `json:"pushConfig,omitempty"`

	// Optional. If delivery to BigQuery is used with this subscription, this
	//  field is used to configure it.
	// +kcc:proto:field=google.pubsub.v1.Subscription.bigquery_config
	BigqueryConfig *BigQueryConfig `json:"bigqueryConfig,omitempty"`

	// Optional. If delivery to Google Cloud Storage is used with this
	//  subscription, this field is used to configure it.
	// +kcc:proto:field=google.pubsub.v1.Subscription.cloud_storage_config
	CloudStorageConfig *CloudStorageConfig `json:"cloudStorageConfig,omitempty"`

	// Optional. The approximate amount of time (on a best-effort basis) Pub/Sub
	//  waits for the subscriber to acknowledge receipt before resending the
	//  message. In the interval after the message is delivered and before it is
	//  acknowledged, it is considered to be _outstanding_. During that time
	//  period, the message will not be redelivered (on a best-effort basis).
	//
	//  For pull subscriptions, this value is used as the initial value for the ack
	//  deadline. To override this value for a given message, call
	//  `ModifyAckDeadline` with the corresponding `ack_id` if using
	//  non-streaming pull or send the `ack_id` in a
	//  `StreamingModifyAckDeadlineRequest` if using streaming pull.
	//  The minimum custom deadline you can specify is 10 seconds.
	//  The maximum custom deadline you can specify is 600 seconds (10 minutes).
	//  If this parameter is 0, a default value of 10 seconds is used.
	//
	//  For push delivery, this value is also used to set the request timeout for
	//  the call to the push endpoint.
	//
	//  If the subscriber never acknowledges the message, the Pub/Sub
	//  system will eventually redeliver the message.
	// +kcc:proto:field=google.pubsub.v1.Subscription.ack_deadline_seconds
	AckDeadlineSeconds *int32 `json:"ackDeadlineSeconds,omitempty"`

	// Optional. Indicates whether to retain acknowledged messages. If true, then
	//  messages are not expunged from the subscription's backlog, even if they are
	//  acknowledged, until they fall out of the `message_retention_duration`
	//  window. This must be true if you would like to [`Seek` to a timestamp]
	//  (https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) in
	//  the past to replay previously-acknowledged messages.
	// +kcc:proto:field=google.pubsub.v1.Subscription.retain_acked_messages
	RetainAckedMessages *bool `json:"retainAckedMessages,omitempty"`

	// Optional. How long to retain unacknowledged messages in the subscription's
	//  backlog, from the moment a message is published. If `retain_acked_messages`
	//  is true, then this also configures the retention of acknowledged messages,
	//  and thus configures how far back in time a `Seek` can be done. Defaults to
	//  7 days. Cannot be more than 31 days or less than 10 minutes.
	// +kcc:proto:field=google.pubsub.v1.Subscription.message_retention_duration
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty"`

	// Optional. See [Creating and managing
	//  labels](https://cloud.google.com/pubsub/docs/labels).
	// +kcc:proto:field=google.pubsub.v1.Subscription.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. If true, messages published with the same `ordering_key` in
	//  `PubsubMessage` will be delivered to the subscribers in the order in which
	//  they are received by the Pub/Sub system. Otherwise, they may be delivered
	//  in any order.
	// +kcc:proto:field=google.pubsub.v1.Subscription.enable_message_ordering
	EnableMessageOrdering *bool `json:"enableMessageOrdering,omitempty"`

	// Optional. A policy that specifies the conditions for this subscription's
	//  expiration. A subscription is considered active as long as any connected
	//  subscriber is successfully consuming messages from the subscription or is
	//  issuing operations on the subscription. If `expiration_policy` is not set,
	//  a *default policy* with `ttl` of 31 days will be used. The minimum allowed
	//  value for `expiration_policy.ttl` is 1 day. If `expiration_policy` is set,
	//  but `expiration_policy.ttl` is not set, the subscription never expires.
	// +kcc:proto:field=google.pubsub.v1.Subscription.expiration_policy
	ExpirationPolicy *ExpirationPolicy `json:"expirationPolicy,omitempty"`

	// Optional. An expression written in the Pub/Sub [filter
	//  language](https://cloud.google.com/pubsub/docs/filtering). If non-empty,
	//  then only `PubsubMessage`s whose `attributes` field matches the filter are
	//  delivered on this subscription. If empty, then no messages are filtered
	//  out.
	// +kcc:proto:field=google.pubsub.v1.Subscription.filter
	Filter *string `json:"filter,omitempty"`

	// Optional. A policy that specifies the conditions for dead lettering
	//  messages in this subscription. If dead_letter_policy is not set, dead
	//  lettering is disabled.
	//
	//  The Pub/Sub service account associated with this subscriptions's
	//  parent project (i.e.,
	//  service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	//  permission to Acknowledge() messages on this subscription.
	// +kcc:proto:field=google.pubsub.v1.Subscription.dead_letter_policy
	DeadLetterPolicy *DeadLetterPolicy `json:"deadLetterPolicy,omitempty"`

	// Optional. A policy that specifies how Pub/Sub retries message delivery for
	//  this subscription.
	//
	//  If not set, the default retry policy is applied. This generally implies
	//  that messages will be retried as soon as possible for healthy subscribers.
	//  RetryPolicy will be triggered on NACKs or acknowledgement deadline
	//  exceeded events for a given message.
	// +kcc:proto:field=google.pubsub.v1.Subscription.retry_policy
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`

	// Optional. Indicates whether the subscription is detached from its topic.
	//  Detached subscriptions don't receive messages from their topic and don't
	//  retain any backlog. `Pull` and `StreamingPull` requests will return
	//  FAILED_PRECONDITION. If the subscription is a push subscription, pushes to
	//  the endpoint will not be made.
	// +kcc:proto:field=google.pubsub.v1.Subscription.detached
	Detached *bool `json:"detached,omitempty"`

	// Optional. If true, Pub/Sub provides the following guarantees for the
	//  delivery of a message with a given value of `message_id` on this
	//  subscription:
	//
	//  * The message sent to a subscriber is guaranteed not to be resent
	//  before the message's acknowledgement deadline expires.
	//  * An acknowledged message will not be resent to a subscriber.
	//
	//  Note that subscribers may still receive multiple copies of a message
	//  when `enable_exactly_once_delivery` is true if the message was published
	//  multiple times by a publisher client. These copies are  considered distinct
	//  by Pub/Sub and have distinct `message_id` values.
	// +kcc:proto:field=google.pubsub.v1.Subscription.enable_exactly_once_delivery
	EnableExactlyOnceDelivery *bool `json:"enableExactlyOnceDelivery,omitempty"`
}

// +kcc:proto=google.pubsub.v1.Subscription.AnalyticsHubSubscriptionInfo
type Subscription_AnalyticsHubSubscriptionInfo struct {
	// Optional. The name of the associated Analytics Hub listing resource.
	//  Pattern:
	//  "projects/{project}/locations/{location}/dataExchanges/{data_exchange}/listings/{listing}"
	// +kcc:proto:field=google.pubsub.v1.Subscription.AnalyticsHubSubscriptionInfo.listing
	Listing *string `json:"listing,omitempty"`

	// Optional. The name of the associated Analytics Hub subscription resource.
	//  Pattern:
	//  "projects/{project}/locations/{location}/subscriptions/{subscription}"
	// +kcc:proto:field=google.pubsub.v1.Subscription.AnalyticsHubSubscriptionInfo.subscription
	Subscription *string `json:"subscription,omitempty"`
}

// +kcc:proto=google.pubsub.v1.BigQueryConfig
type BigQueryConfigObservedState struct {
	// Output only. An output-only field that indicates whether or not the
	//  subscription can receive messages.
	// +kcc:proto:field=google.pubsub.v1.BigQueryConfig.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.CloudStorageConfig
type CloudStorageConfigObservedState struct {
	// Output only. An output-only field that indicates whether or not the
	//  subscription can receive messages.
	// +kcc:proto:field=google.pubsub.v1.CloudStorageConfig.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.Subscription
type SubscriptionObservedState struct {
	// Optional. If delivery to BigQuery is used with this subscription, this
	//  field is used to configure it.
	// +kcc:proto:field=google.pubsub.v1.Subscription.bigquery_config
	BigqueryConfig *BigQueryConfigObservedState `json:"bigqueryConfig,omitempty"`

	// Optional. If delivery to Google Cloud Storage is used with this
	//  subscription, this field is used to configure it.
	// +kcc:proto:field=google.pubsub.v1.Subscription.cloud_storage_config
	CloudStorageConfig *CloudStorageConfigObservedState `json:"cloudStorageConfig,omitempty"`

	// Output only. Indicates the minimum duration for which a message is retained
	//  after it is published to the subscription's topic. If this field is set,
	//  messages published to the subscription's topic in the last
	//  `topic_message_retention_duration` are always available to subscribers. See
	//  the `message_retention_duration` field in `Topic`. This field is set only
	//  in responses from the server; it is ignored if it is set in any requests.
	// +kcc:proto:field=google.pubsub.v1.Subscription.topic_message_retention_duration
	TopicMessageRetentionDuration *string `json:"topicMessageRetentionDuration,omitempty"`

	// Output only. An output-only field indicating whether or not the
	//  subscription can receive messages.
	// +kcc:proto:field=google.pubsub.v1.Subscription.state
	State *string `json:"state,omitempty"`

	// Output only. Information about the associated Analytics Hub subscription.
	//  Only set if the subscritpion is created by Analytics Hub.
	// +kcc:proto:field=google.pubsub.v1.Subscription.analytics_hub_subscription_info
	AnalyticsHubSubscriptionInfo *Subscription_AnalyticsHubSubscriptionInfo `json:"analyticsHubSubscriptionInfo,omitempty"`
}
