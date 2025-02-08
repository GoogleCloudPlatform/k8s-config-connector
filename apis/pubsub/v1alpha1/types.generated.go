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


// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings
type IngestionDataSourceSettings struct {
	// Optional. Amazon Kinesis Data Streams.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.aws_kinesis
	AwsKinesis *IngestionDataSourceSettings_AwsKinesis `json:"awsKinesis,omitempty"`

	// Optional. Cloud Storage.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.cloud_storage
	CloudStorage *IngestionDataSourceSettings_CloudStorage `json:"cloudStorage,omitempty"`

	// Optional. Azure Event Hubs.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.azure_event_hubs
	AzureEventHubs *IngestionDataSourceSettings_AzureEventHubs `json:"azureEventHubs,omitempty"`

	// Optional. Amazon MSK.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.aws_msk
	AwsMsk *IngestionDataSourceSettings_AwsMsk `json:"awsMsk,omitempty"`

	// Optional. Confluent Cloud.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.confluent_cloud
	ConfluentCloud *IngestionDataSourceSettings_ConfluentCloud `json:"confluentCloud,omitempty"`

	// Optional. Platform Logs settings. If unset, no Platform Logs will be
	//  generated.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.platform_logs_settings
	PlatformLogsSettings *PlatformLogsSettings `json:"platformLogsSettings,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis
type IngestionDataSourceSettings_AwsKinesis struct {

	// Required. The Kinesis stream ARN to ingest data from.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis.stream_arn
	StreamArn *string `json:"streamArn,omitempty"`

	// Required. The Kinesis consumer ARN to used for ingestion in Enhanced
	//  Fan-Out mode. The consumer must be already created and ready to be used.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis.consumer_arn
	ConsumerArn *string `json:"consumerArn,omitempty"`

	// Required. AWS role ARN to be used for Federated Identity authentication
	//  with Kinesis. Check the Pub/Sub docs for how to set up this role and the
	//  required permissions that need to be attached to it.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis.aws_role_arn
	AwsRoleArn *string `json:"awsRoleArn,omitempty"`

	// Required. The GCP service account to be used for Federated Identity
	//  authentication with Kinesis (via a `AssumeRoleWithWebIdentity` call for
	//  the provided role). The `aws_role_arn` must be set up with
	//  `accounts.google.com:sub` equals to this service account number.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis.gcp_service_account
	GcpServiceAccount *string `json:"gcpServiceAccount,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk
type IngestionDataSourceSettings_AwsMsk struct {

	// Required. The Amazon Resource Name (ARN) that uniquely identifies the
	//  cluster.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk.cluster_arn
	ClusterArn *string `json:"clusterArn,omitempty"`

	// Required. The name of the topic in the Amazon MSK cluster that Pub/Sub
	//  will import from.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk.topic
	Topic *string `json:"topic,omitempty"`

	// Required. AWS role ARN to be used for Federated Identity authentication
	//  with Amazon MSK. Check the Pub/Sub docs for how to set up this role and
	//  the required permissions that need to be attached to it.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk.aws_role_arn
	AwsRoleArn *string `json:"awsRoleArn,omitempty"`

	// Required. The GCP service account to be used for Federated Identity
	//  authentication with Amazon MSK (via a `AssumeRoleWithWebIdentity` call
	//  for the provided role). The `aws_role_arn` must be set up with
	//  `accounts.google.com:sub` equals to this service account number.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk.gcp_service_account
	GcpServiceAccount *string `json:"gcpServiceAccount,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs
type IngestionDataSourceSettings_AzureEventHubs struct {

	// Optional. Name of the resource group within the azure subscription.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.resource_group
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// Optional. The name of the Event Hubs namespace.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.namespace
	Namespace *string `json:"namespace,omitempty"`

	// Optional. The name of the Event Hub.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.event_hub
	EventHub *string `json:"eventHub,omitempty"`

	// Optional. The client id of the Azure application that is being used to
	//  authenticate Pub/Sub.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Optional. The tenant id of the Azure application that is being used to
	//  authenticate Pub/Sub.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.tenant_id
	TenantID *string `json:"tenantID,omitempty"`

	// Optional. The Azure subscription id.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.subscription_id
	SubscriptionID *string `json:"subscriptionID,omitempty"`

	// Optional. The GCP service account to be used for Federated Identity
	//  authentication.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.gcp_service_account
	GcpServiceAccount *string `json:"gcpServiceAccount,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage
type IngestionDataSourceSettings_CloudStorage struct {

	// Optional. Cloud Storage bucket. The bucket name must be without any
	//  prefix like "gs://". See the [bucket naming requirements]
	//  (https://cloud.google.com/storage/docs/buckets#naming).
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Optional. Data from Cloud Storage will be interpreted as text.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.text_format
	TextFormat *IngestionDataSourceSettings_CloudStorage_TextFormat `json:"textFormat,omitempty"`

	// Optional. Data from Cloud Storage will be interpreted in Avro format.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.avro_format
	AvroFormat *IngestionDataSourceSettings_CloudStorage_AvroFormat `json:"avroFormat,omitempty"`

	// Optional. It will be assumed data from Cloud Storage was written via
	//  [Cloud Storage
	//  subscriptions](https://cloud.google.com/pubsub/docs/cloudstorage).
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.pubsub_avro_format
	PubsubAvroFormat *IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat `json:"pubsubAvroFormat,omitempty"`

	// Optional. Only objects with a larger or equal creation timestamp will be
	//  ingested.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.minimum_object_create_time
	MinimumObjectCreateTime *string `json:"minimumObjectCreateTime,omitempty"`

	// Optional. Glob pattern used to match objects that will be ingested. If
	//  unset, all objects will be ingested. See the [supported
	//  patterns](https://cloud.google.com/storage/docs/json_api/v1/objects/list#list-objects-and-prefixes-using-glob).
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.match_glob
	MatchGlob *string `json:"matchGlob,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.AvroFormat
type IngestionDataSourceSettings_CloudStorage_AvroFormat struct {
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.PubSubAvroFormat
type IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat struct {
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.TextFormat
type IngestionDataSourceSettings_CloudStorage_TextFormat struct {
	// Optional. When unset, '\n' is used.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.TextFormat.delimiter
	Delimiter *string `json:"delimiter,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud
type IngestionDataSourceSettings_ConfluentCloud struct {

	// Required. The address of the bootstrap server. The format is url:port.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud.bootstrap_server
	BootstrapServer *string `json:"bootstrapServer,omitempty"`

	// Required. The id of the cluster.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud.cluster_id
	ClusterID *string `json:"clusterID,omitempty"`

	// Required. The name of the topic in the Confluent Cloud cluster that
	//  Pub/Sub will import from.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud.topic
	Topic *string `json:"topic,omitempty"`

	// Required. The id of the identity pool to be used for Federated Identity
	//  authentication with Confluent Cloud. See
	//  https://docs.confluent.io/cloud/current/security/authenticate/workload-identities/identity-providers/oauth/identity-pools.html#add-oauth-identity-pools.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud.identity_pool_id
	IdentityPoolID *string `json:"identityPoolID,omitempty"`

	// Required. The GCP service account to be used for Federated Identity
	//  authentication with `identity_pool_id`.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud.gcp_service_account
	GcpServiceAccount *string `json:"gcpServiceAccount,omitempty"`
}

// +kcc:proto=google.pubsub.v1.MessageStoragePolicy
type MessageStoragePolicy struct {
	// Optional. A list of IDs of Google Cloud regions where messages that are
	//  published to the topic may be persisted in storage. Messages published by
	//  publishers running in non-allowed Google Cloud regions (or running outside
	//  of Google Cloud altogether) are routed for storage in one of the allowed
	//  regions. An empty list means that no regions are allowed, and is not a
	//  valid configuration.
	// +kcc:proto:field=google.pubsub.v1.MessageStoragePolicy.allowed_persistence_regions
	AllowedPersistenceRegions []string `json:"allowedPersistenceRegions,omitempty"`

	// Optional. If true, `allowed_persistence_regions` is also used to enforce
	//  in-transit guarantees for messages. That is, Pub/Sub will fail
	//  Publish operations on this topic and subscribe operations
	//  on any subscription attached to this topic in any region that is
	//  not in `allowed_persistence_regions`.
	// +kcc:proto:field=google.pubsub.v1.MessageStoragePolicy.enforce_in_transit
	EnforceInTransit *bool `json:"enforceInTransit,omitempty"`
}

// +kcc:proto=google.pubsub.v1.PlatformLogsSettings
type PlatformLogsSettings struct {
	// Optional. The minimum severity level of Platform Logs that will be written.
	// +kcc:proto:field=google.pubsub.v1.PlatformLogsSettings.severity
	Severity *string `json:"severity,omitempty"`
}

// +kcc:proto=google.pubsub.v1.SchemaSettings
type SchemaSettings struct {
	// Required. The name of the schema that messages published should be
	//  validated against. Format is `projects/{project}/schemas/{schema}`. The
	//  value of this field will be `_deleted-schema_` if the schema has been
	//  deleted.
	// +kcc:proto:field=google.pubsub.v1.SchemaSettings.schema
	Schema *string `json:"schema,omitempty"`

	// Optional. The encoding of messages validated against `schema`.
	// +kcc:proto:field=google.pubsub.v1.SchemaSettings.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. The minimum (inclusive) revision allowed for validating messages.
	//  If empty or not present, allow any revision to be validated against
	//  last_revision or any revision created before.
	// +kcc:proto:field=google.pubsub.v1.SchemaSettings.first_revision_id
	FirstRevisionID *string `json:"firstRevisionID,omitempty"`

	// Optional. The maximum (inclusive) revision allowed for validating messages.
	//  If empty or not present, allow any revision to be validated against
	//  first_revision or any revision created after.
	// +kcc:proto:field=google.pubsub.v1.SchemaSettings.last_revision_id
	LastRevisionID *string `json:"lastRevisionID,omitempty"`
}

// +kcc:proto=google.pubsub.v1.Topic
type Topic struct {
	// Required. The name of the topic. It must have the format
	//  `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter,
	//  and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`),
	//  underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent
	//  signs (`%`). It must be between 3 and 255 characters in length, and it
	//  must not start with `"goog"`.
	// +kcc:proto:field=google.pubsub.v1.Topic.name
	Name *string `json:"name,omitempty"`

	// Optional. See [Creating and managing labels]
	//  (https://cloud.google.com/pubsub/docs/labels).
	// +kcc:proto:field=google.pubsub.v1.Topic.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Policy constraining the set of Google Cloud Platform regions
	//  where messages published to the topic may be stored. If not present, then
	//  no constraints are in effect.
	// +kcc:proto:field=google.pubsub.v1.Topic.message_storage_policy
	MessageStoragePolicy *MessageStoragePolicy `json:"messageStoragePolicy,omitempty"`

	// Optional. The resource name of the Cloud KMS CryptoKey to be used to
	//  protect access to messages published on this topic.
	//
	//  The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.pubsub.v1.Topic.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Optional. Settings for validating messages published against a schema.
	// +kcc:proto:field=google.pubsub.v1.Topic.schema_settings
	SchemaSettings *SchemaSettings `json:"schemaSettings,omitempty"`

	// Optional. Reserved for future use. This field is set only in responses from
	//  the server; it is ignored if it is set in any requests.
	// +kcc:proto:field=google.pubsub.v1.Topic.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Optional. Indicates the minimum duration to retain a message after it is
	//  published to the topic. If this field is set, messages published to the
	//  topic in the last `message_retention_duration` are always available to
	//  subscribers. For instance, it allows any attached subscription to [seek to
	//  a
	//  timestamp](https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time)
	//  that is up to `message_retention_duration` in the past. If this field is
	//  not set, message retention is controlled by settings on individual
	//  subscriptions. Cannot be more than 31 days or less than 10 minutes.
	// +kcc:proto:field=google.pubsub.v1.Topic.message_retention_duration
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty"`

	// Optional. Settings for ingestion from a data source into this topic.
	// +kcc:proto:field=google.pubsub.v1.Topic.ingestion_data_source_settings
	IngestionDataSourceSettings *IngestionDataSourceSettings `json:"ingestionDataSourceSettings,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings
type IngestionDataSourceSettingsObservedState struct {
	// Optional. Amazon Kinesis Data Streams.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.aws_kinesis
	AwsKinesis *IngestionDataSourceSettings_AwsKinesisObservedState `json:"awsKinesis,omitempty"`

	// Optional. Cloud Storage.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.cloud_storage
	CloudStorage *IngestionDataSourceSettings_CloudStorageObservedState `json:"cloudStorage,omitempty"`

	// Optional. Azure Event Hubs.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.azure_event_hubs
	AzureEventHubs *IngestionDataSourceSettings_AzureEventHubsObservedState `json:"azureEventHubs,omitempty"`

	// Optional. Amazon MSK.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.aws_msk
	AwsMsk *IngestionDataSourceSettings_AwsMskObservedState `json:"awsMsk,omitempty"`

	// Optional. Confluent Cloud.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.confluent_cloud
	ConfluentCloud *IngestionDataSourceSettings_ConfluentCloudObservedState `json:"confluentCloud,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis
type IngestionDataSourceSettings_AwsKinesisObservedState struct {
	// Output only. An output-only field that indicates the state of the Kinesis
	//  ingestion source.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsKinesis.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk
type IngestionDataSourceSettings_AwsMskObservedState struct {
	// Output only. An output-only field that indicates the state of the Amazon
	//  MSK ingestion source.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AwsMsk.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs
type IngestionDataSourceSettings_AzureEventHubsObservedState struct {
	// Output only. An output-only field that indicates the state of the Event
	//  Hubs ingestion source.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.AzureEventHubs.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage
type IngestionDataSourceSettings_CloudStorageObservedState struct {
	// Output only. An output-only field that indicates the state of the Cloud
	//  Storage ingestion source.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.CloudStorage.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud
type IngestionDataSourceSettings_ConfluentCloudObservedState struct {
	// Output only. An output-only field that indicates the state of the
	//  Confluent Cloud ingestion source.
	// +kcc:proto:field=google.pubsub.v1.IngestionDataSourceSettings.ConfluentCloud.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.pubsub.v1.Topic
type TopicObservedState struct {
	// Output only. An output-only field indicating the state of the topic.
	// +kcc:proto:field=google.pubsub.v1.Topic.state
	State *string `json:"state,omitempty"`

	// Optional. Settings for ingestion from a data source into this topic.
	// +kcc:proto:field=google.pubsub.v1.Topic.ingestion_data_source_settings
	IngestionDataSourceSettings *IngestionDataSourceSettingsObservedState `json:"ingestionDataSourceSettings,omitempty"`
}
