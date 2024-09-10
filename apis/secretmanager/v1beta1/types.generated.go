// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.cloud.secretmanager.v1.CustomerManagedEncryption
type CustomerManagedEncryption struct {
	// Required. The resource name of the Cloud KMS CryptoKey used to encrypt
	//  secret payloads.
	//
	//  For secrets using the
	//  [UserManaged][google.cloud.secretmanager.v1.Replication.UserManaged]
	//  replication policy type, Cloud KMS CryptoKeys must reside in the same
	//  location as the [replica location][Secret.UserManaged.Replica.location].
	//
	//  For secrets using the
	//  [Automatic][google.cloud.secretmanager.v1.Replication.Automatic]
	//  replication policy type, Cloud KMS CryptoKeys must reside in `global`.
	//
	//  The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	KmsKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication
type Replication struct {
	// The [Secret][google.cloud.secretmanager.v1.Secret] will automatically be
	//  replicated without any restrictions.
	Automatic *Replication_Automatic `json:"automatic,omitempty"`

	// The [Secret][google.cloud.secretmanager.v1.Secret] will only be
	//  replicated into the locations specified.
	UserManaged *Replication_UserManaged `json:"userManaged,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication.Automatic
type Replication_Automatic struct {
	// Optional. The customer-managed encryption configuration of the
	//  [Secret][google.cloud.secretmanager.v1.Secret]. If no configuration is
	//  provided, Google-managed default encryption is used.
	//
	//  Updates to the [Secret][google.cloud.secretmanager.v1.Secret] encryption
	//  configuration only apply to
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion] added
	//  afterwards. They do not apply retroactively to existing
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
	CustomerManagedEncryption *CustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication.UserManaged
type Replication_UserManaged struct {
	// Required. The list of Replicas for this
	//  [Secret][google.cloud.secretmanager.v1.Secret].
	//
	//  Cannot be empty.
	Replicas []Replication_UserManaged_Replica `json:"replicas,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Replication.UserManaged.Replica
type Replication_UserManaged_Replica struct {
	// The canonical IDs of the location to replicate data.
	//  For example: `"us-east1"`.
	Location *string `json:"location,omitempty"`

	// Optional. The customer-managed encryption configuration of the
	//  [User-Managed Replica][Replication.UserManaged.Replica]. If no
	//  configuration is provided, Google-managed default encryption is used.
	//
	//  Updates to the [Secret][google.cloud.secretmanager.v1.Secret]
	//  encryption configuration only apply to
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion] added
	//  afterwards. They do not apply retroactively to existing
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
	CustomerManagedEncryption *CustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Rotation
type Rotation struct {
	// Optional. Timestamp in UTC at which the
	//  [Secret][google.cloud.secretmanager.v1.Secret] is scheduled to rotate.
	//  Cannot be set to less than 300s (5 min) in the future and at most
	//  3153600000s (100 years).
	//
	//  [next_rotation_time][google.cloud.secretmanager.v1.Rotation.next_rotation_time]
	//  MUST  be set if
	//  [rotation_period][google.cloud.secretmanager.v1.Rotation.rotation_period]
	//  is set.
	NextRotationTime *string `json:"nextRotationTime,omitempty"`

	// Input only. The Duration between rotation notifications. Must be in seconds
	//  and at least 3600s (1h) and at most 3153600000s (100 years).
	//
	//  If
	//  [rotation_period][google.cloud.secretmanager.v1.Rotation.rotation_period]
	//  is set,
	//  [next_rotation_time][google.cloud.secretmanager.v1.Rotation.next_rotation_time]
	//  must be set.
	//  [next_rotation_time][google.cloud.secretmanager.v1.Rotation.next_rotation_time]
	//  will be advanced by this period when the service automatically sends
	//  rotation notifications.
	RotationPeriod *string `json:"rotationPeriod,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Secret
type Secret struct {
	// Output only. The resource name of the
	//  [Secret][google.cloud.secretmanager.v1.Secret] in the format
	//  `projects/*/secrets/*`.
	Name *string `json:"name,omitempty"`

	// Optional. Immutable. The replication policy of the secret data attached to
	//  the [Secret][google.cloud.secretmanager.v1.Secret].
	//
	//  The replication policy cannot be changed after the Secret has been created.
	Replication *Replication `json:"replication,omitempty"`

	// Output only. The time at which the
	//  [Secret][google.cloud.secretmanager.v1.Secret] was created.
	CreateTime *string `json:"createTime,omitempty"`

	// The labels assigned to this Secret.
	//
	//  Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	//  of maximum 128 bytes, and must conform to the following PCRE regular
	//  expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
	//
	//  Label values must be between 0 and 63 characters long, have a UTF-8
	//  encoding of maximum 128 bytes, and must conform to the following PCRE
	//  regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
	//
	//  No more than 64 labels can be assigned to a given resource.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A list of up to 10 Pub/Sub topics to which messages are published
	//  when control plane operations are called on the secret or its versions.
	Topics []Topic `json:"topics,omitempty"`

	// Optional. Timestamp in UTC when the
	//  [Secret][google.cloud.secretmanager.v1.Secret] is scheduled to expire.
	//  This is always provided on output, regardless of what was sent on input.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL for the
	//  [Secret][google.cloud.secretmanager.v1.Secret].
	Ttl *string `json:"ttl,omitempty"`

	// Optional. Etag of the currently stored
	//  [Secret][google.cloud.secretmanager.v1.Secret].
	Etag *string `json:"etag,omitempty"`

	// Optional. Rotation policy attached to the
	//  [Secret][google.cloud.secretmanager.v1.Secret]. May be excluded if there is
	//  no rotation policy.
	Rotation *Rotation `json:"rotation,omitempty"`

	// Optional. Mapping from version alias to version name.
	//
	//  A version alias is a string with a maximum length of 63 characters and can
	//  contain uppercase and lowercase letters, numerals, and the hyphen (`-`)
	//  and underscore ('_') characters. An alias string must start with a
	//  letter and cannot be the string 'latest' or 'NEW'.
	//  No more than 50 aliases can be assigned to a given secret.
	//
	//  Version-Alias pairs will be viewable via GetSecret and modifiable via
	//  UpdateSecret. Access by alias is only be supported on
	//  GetSecretVersion and AccessSecretVersion.
	VersionAliases map[string]int64 `json:"versionAliases,omitempty"`

	// Optional. Custom metadata about the secret.
	//
	//  Annotations are distinct from various forms of labels.
	//  Annotations exist to allow client tools to store their own state
	//  information without requiring a database.
	//
	//  Annotation keys must be between 1 and 63 characters long, have a UTF-8
	//  encoding of maximum 128 bytes, begin and end with an alphanumeric character
	//  ([a-z0-9A-Z]), and may have dashes (-), underscores (_), dots (.), and
	//  alphanumerics in between these symbols.
	//
	//  The total size of annotation keys and values must be less than 16KiB.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Secret Version TTL after destruction request
	//
	//  This is a part of the Delayed secret version destroy feature.
	//  For secret with TTL>0, version destruction doesn't happen immediately
	//  on calling destroy instead the version goes to a disabled state and
	//  destruction happens after the TTL expires.
	VersionDestroyTtl *string `json:"versionDestroyTtl,omitempty"`

	// Optional. The customer-managed encryption configuration of the Regionalised
	//  Secrets. If no configuration is provided, Google-managed default encryption
	//  is used.
	//
	//  Updates to the [Secret][google.cloud.secretmanager.v1.Secret] encryption
	//  configuration only apply to
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion] added
	//  afterwards. They do not apply retroactively to existing
	//  [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
	CustomerManagedEncryption *CustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
}

// +kcc:proto=google.cloud.secretmanager.v1.Topic
type Topic struct {
	// Required. The resource name of the Pub/Sub topic that will be published to,
	//  in the following format: `projects/*/topics/*`. For publication to succeed,
	//  the Secret Manager service agent must have the `pubsub.topic.publish`
	//  permission on the topic. The Pub/Sub Publisher role
	//  (`roles/pubsub.publisher`) includes this permission.
	Name *string `json:"name,omitempty"`
}
