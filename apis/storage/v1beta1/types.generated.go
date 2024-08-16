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

// +kcc:proto=google.storage.v1.Bucket
type Bucket struct {
	// Access controls on the bucket.
	Acl []BucketAccessControl `json:"acl,omitempty"`

	// Default access controls to apply to new objects when no ACL is provided.
	DefaultObjectAcl []ObjectAccessControl `json:"defaultObjectAcl,omitempty"`

	// The bucket's lifecycle configuration. See
	//  [https://developers.google.com/storage/docs/lifecycle]Lifecycle Management]
	//  for more information.
	Lifecycle *Bucket_Lifecycle `json:"lifecycle,omitempty"`

	// The creation time of the bucket in
	//  [https://tools.ietf.org/html/rfc3339][RFC 3339] format.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	TimeCreated *string `json:"timeCreated,omitempty"`

	// The ID of the bucket. For buckets, the `id` and `name` properties are the
	//  same.
	//  Attempting to update this field after the bucket is created will result in
	//  a [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Id *string `json:"id,omitempty"`

	// The name of the bucket.
	//  Attempting to update this field after the bucket is created will result in
	//  an error.
	Name *string `json:"name,omitempty"`

	// The project number of the project the bucket belongs to.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	ProjectNumber *int64 `json:"projectNumber,omitempty"`

	// The metadata generation of this bucket.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Metageneration *int64 `json:"metageneration,omitempty"`

	// The bucket's [https://www.w3.org/TR/cors/][Cross-Origin Resource Sharing]
	//  (CORS) configuration.
	Cors []Bucket_Cors `json:"cors,omitempty"`

	// The location of the bucket. Object data for objects in the bucket resides
	//  in physical storage within this region.  Defaults to `US`. See the
	//  [https://developers.google.com/storage/docs/concepts-techniques#specifyinglocations"][developer's
	//  guide] for the authoritative list. Attempting to update this field after
	//  the bucket is created will result in an error.
	Location *string `json:"location,omitempty"`

	// The bucket's default storage class, used whenever no storageClass is
	//  specified for a newly-created object. This defines how objects in the
	//  bucket are stored and determines the SLA and the cost of storage.
	//  If this value is not specified when the bucket is created, it will default
	//  to `STANDARD`. For more information, see
	//  https://developers.google.com/storage/docs/storage-classes.
	StorageClass *string `json:"storageClass,omitempty"`

	// HTTP 1.1 [https://tools.ietf.org/html/rfc7232#section-2.3"]Entity tag]
	//  for the bucket.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Etag *string `json:"etag,omitempty"`

	// The modification time of the bucket.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Updated *string `json:"updated,omitempty"`

	// The default value for event-based hold on newly created objects in this
	//  bucket.  Event-based hold is a way to retain objects indefinitely until an
	//  event occurs, signified by the
	//  hold's release. After being released, such objects will be subject to
	//  bucket-level retention (if any).  One sample use case of this flag is for
	//  banks to hold loan documents for at least 3 years after loan is paid in
	//  full. Here, bucket-level retention is 3 years and the event is loan being
	//  paid in full. In this example, these objects will be held intact for any
	//  number of years until the event has occurred (event-based hold on the
	//  object is released) and then 3 more years after that. That means retention
	//  duration of the objects begins from the moment event-based hold
	//  transitioned from true to false.  Objects under event-based hold cannot be
	//  deleted, overwritten or archived until the hold is removed.
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty"`

	// User-provided labels, in key/value pairs.
	Labels map[string]string `json:"labels,omitempty"`

	// The bucket's website configuration, controlling how the service behaves
	//  when accessing bucket contents as a web site. See the
	//  [https://cloud.google.com/storage/docs/static-website][Static Website
	//  Examples] for more information.
	Website *Bucket_Website `json:"website,omitempty"`

	// The bucket's versioning configuration.
	Versioning *Bucket_Versioning `json:"versioning,omitempty"`

	// The bucket's logging configuration, which defines the destination bucket
	//  and optional name prefix for the current bucket's logs.
	Logging *Bucket_Logging `json:"logging,omitempty"`

	// The owner of the bucket. This is always the project team's owner group.
	Owner *Owner `json:"owner,omitempty"`

	// Encryption configuration for a bucket.
	Encryption *Bucket_Encryption `json:"encryption,omitempty"`

	// The bucket's billing configuration.
	Billing *Bucket_Billing `json:"billing,omitempty"`

	// The bucket's retention policy. The retention policy enforces a minimum
	//  retention time for all objects contained in the bucket, based on their
	//  creation time. Any attempt to overwrite or delete objects younger than the
	//  retention period will result in a PERMISSION_DENIED error.  An unlocked
	//  retention policy can be modified or removed from the bucket via a
	//  storage.buckets.update operation. A locked retention policy cannot be
	//  removed or shortened in duration for the lifetime of the bucket.
	//  Attempting to remove or decrease period of a locked retention policy will
	//  result in a PERMISSION_DENIED error.
	RetentionPolicy *Bucket_RetentionPolicy `json:"retentionPolicy,omitempty"`

	// The location type of the bucket (region, dual-region, multi-region, etc).
	LocationType *string `json:"locationType,omitempty"`

	// The bucket's IAM configuration.
	IamConfiguration *Bucket_IamConfiguration `json:"iamConfiguration,omitempty"`

	// The zone or zones from which the bucket is intended to use zonal quota.
	//  Requests for data from outside the specified affinities are still allowed
	//  but won't be able to use zonal quota. The values are case-insensitive.
	//  Attempting to update this field after bucket is created will result in an
	//  error.
	ZoneAffinity []string `json:"zoneAffinity,omitempty"`

	// Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// The bucket's autoclass configuration. If there is no configuration, the
	//  Autoclass feature will be disabled and have no effect on the bucket.
	Autoclass *Bucket_Autoclass `json:"autoclass,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Autoclass
type Bucket_Autoclass struct {
	// Enables Autoclass.
	Enabled *bool `json:"enabled,omitempty"`

	// Latest instant at which the `enabled` bit was flipped.
	ToggleTime *string `json:"toggleTime,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Billing
type Bucket_Billing struct {
	// When set to true, Requester Pays is enabled for this bucket.
	RequesterPays *bool `json:"requesterPays,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Cors
type Bucket_Cors struct {
	// The list of Origins eligible to receive CORS response headers. See
	//  [https://tools.ietf.org/html/rfc6454][RFC 6454] for more on origins.
	//  Note: "*" is permitted in the list of origins, and means "any Origin".
	Origin []string `json:"origin,omitempty"`

	// The list of HTTP methods on which to include CORS response headers,
	//  (`GET`, `OPTIONS`, `POST`, etc) Note: "*" is permitted in the list of
	//  methods, and means "any method".
	Method []string `json:"method,omitempty"`

	// The list of HTTP headers other than the
	//  [https://www.w3.org/TR/cors/#simple-response-header][simple response
	//  headers] to give permission for the user-agent to share across domains.
	ResponseHeader []string `json:"responseHeader,omitempty"`

	// The value, in seconds, to return in the
	//  [https://www.w3.org/TR/cors/#access-control-max-age-response-header][Access-Control-Max-Age
	//  header] used in preflight responses.
	MaxAgeSeconds *int32 `json:"maxAgeSeconds,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Encryption
type Bucket_Encryption struct {
	// A Cloud KMS key that will be used to encrypt objects inserted into this
	//  bucket, if no encryption method is specified.
	DefaultKmsKeyName *string `json:"defaultKmsKeyName,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.IamConfiguration
type Bucket_IamConfiguration struct {
	UniformBucketLevelAccess *Bucket_IamConfiguration_UniformBucketLevelAccess `json:"uniformBucketLevelAccess,omitempty"`

	// Whether IAM will enforce public access prevention.
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.IamConfiguration.UniformBucketLevelAccess
type Bucket_IamConfiguration_UniformBucketLevelAccess struct {
	// If set, access checks only use bucket-level IAM policies or above.
	Enabled *bool `json:"enabled,omitempty"`

	// The deadline time for changing
	//  <code>iamConfiguration.uniformBucketLevelAccess.enabled</code> from
	//  true to false in [https://tools.ietf.org/html/rfc3339][RFC 3339]. After
	//  the deadline is passed the field is immutable.
	LockedTime *string `json:"lockedTime,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Lifecycle
type Bucket_Lifecycle struct {
	// A lifecycle management rule, which is made of an action to take and the
	//  condition(s) under which the action will be taken.
	Rule []Bucket_Lifecycle_Rule `json:"rule,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Lifecycle.Rule
type Bucket_Lifecycle_Rule struct {
	// The action to take.
	Action *Bucket_Lifecycle_Rule_Action `json:"action,omitempty"`

	// The condition(s) under which the action will be taken.
	Condition *Bucket_Lifecycle_Rule_Condition `json:"condition,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Lifecycle.Rule.Action
type Bucket_Lifecycle_Rule_Action struct {
	// Type of the action. Currently, only `Delete`, `SetStorageClass`, and
	//  `AbortIncompleteMultipartUpload` are supported.
	Type *string `json:"type,omitempty"`

	// Target storage class. Required iff the type of the action is
	//  SetStorageClass.
	StorageClass *string `json:"storageClass,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Lifecycle.Rule.Condition
type Bucket_Lifecycle_Rule_Condition struct {
	// Age of an object (in days). This condition is satisfied when an
	//  object reaches the specified age.
	Age *int32 `json:"age,omitempty"`

	// A date in [RFC 3339][1] format with only the date part (for
	//  instance, "2013-01-15"). This condition is satisfied when an
	//  object is created before midnight of the specified date in UTC.
	//  [1]: https://tools.ietf.org/html/rfc3339
	CreatedBefore *string `json:"createdBefore,omitempty"`

	// Relevant only for versioned objects. If the value is
	//  `true`, this condition matches live objects; if the value
	//  is `false`, it matches archived objects.
	// IsLive *google_protobuf_BoolValue `json:"isLive,omitempty"`

	// Relevant only for versioned objects. If the value is N, this
	//  condition is satisfied when there are at least N versions (including
	//  the live version) newer than this version of the object.
	NumNewerVersions *int32 `json:"numNewerVersions,omitempty"`

	// Objects having any of the storage classes specified by this condition
	//  will be matched. Values include `MULTI_REGIONAL`, `REGIONAL`,
	//  `NEARLINE`, `COLDLINE`, `STANDARD`, and
	//  `DURABLE_REDUCED_AVAILABILITY`.
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty"`

	// A regular expression that satisfies the RE2 syntax. This condition is
	//  satisfied when the name of the object matches the RE2 pattern.  Note:
	//  This feature is currently in the "Early Access" launch stage and is
	//  only available to an allowlisted set of users; that means that this
	//  feature may be changed in backward-incompatible ways and that it is
	//  not guaranteed to be released.
	MatchesPattern *string `json:"matchesPattern,omitempty"`

	// Number of days that has elapsed since the custom timestamp set on an
	//  object.
	DaysSinceCustomTime *int32 `json:"daysSinceCustomTime,omitempty"`

	// An object matches this condition if the custom timestamp set on the
	//  object is before this timestamp.
	CustomTimeBefore *string `json:"customTimeBefore,omitempty"`

	// This condition is relevant only for versioned objects. An object
	//  version satisfies this condition only if these many days have been
	//  passed since it became noncurrent. The value of the field must be a
	//  nonnegative integer. If it's zero, the object version will become
	//  eligible for Lifecycle action as soon as it becomes noncurrent.
	DaysSinceNoncurrentTime *int32 `json:"daysSinceNoncurrentTime,omitempty"`

	// This condition is relevant only for versioned objects. An object
	//  version satisfies this condition only if it became noncurrent before
	//  the specified timestamp.
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty"`

	// List of object name prefixes. If any prefix exactly matches the
	//  beginning of the object name, the condition evaluates to true.
	MatchesPrefix []string `json:"matchesPrefix,omitempty"`

	// List of object name suffixes. If any suffix exactly matches the
	//  end of the object name, the condition evaluates to true.
	MatchesSuffix []string `json:"matchesSuffix,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Logging
type Bucket_Logging struct {
	// The destination bucket where the current bucket's logs should be placed.
	LogBucket *string `json:"logBucket,omitempty"`

	// A prefix for log object names.
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.RetentionPolicy
type Bucket_RetentionPolicy struct {
	// Server-determined value that indicates the time from which policy was
	//  enforced and effective. This value is in
	//  [https://tools.ietf.org/html/rfc3339][RFC 3339] format.
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Once locked, an object retention policy cannot be modified.
	IsLocked *bool `json:"isLocked,omitempty"`

	// The duration in seconds that objects need to be retained. Retention
	//  duration must be greater than zero and less than 100 years. Note that
	//  enforcement of retention periods less than a day is not guaranteed. Such
	//  periods should only be used for testing purposes.
	RetentionPeriod *int64 `json:"retentionPeriod,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Versioning
type Bucket_Versioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.storage.v1.Bucket.Website
type Bucket_Website struct {
	// If the requested object path is missing, the service will ensure the path
	//  has a trailing '/', append this suffix, and attempt to retrieve the
	//  resulting object. This allows the creation of `index.html`
	//  objects to represent directory pages.
	MainPageSuffix *string `json:"mainPageSuffix,omitempty"`

	// If the requested object path is missing, and any
	//  `mainPageSuffix` object is missing, if applicable, the service
	//  will return the named object from this bucket as the content for a
	//  [https://tools.ietf.org/html/rfc7231#section-6.5.4][404 Not Found]
	//  result.
	NotFoundPage *string `json:"notFoundPage,omitempty"`
}

// +kcc:proto=google.storage.v1.BucketAccessControl
type BucketAccessControl struct {
	// The access permission for the entity.
	Role *string `json:"role,omitempty"`

	// HTTP 1.1 ["https://tools.ietf.org/html/rfc7232#section-2.3][Entity tag]
	//  for the access-control entry.
	Etag *string `json:"etag,omitempty"`

	// The ID of the access-control entry.
	Id *string `json:"id,omitempty"`

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty"`

	// The entity holding the permission, in one of the following forms:
	//  * `user-{userid}`
	//  * `user-{email}`
	//  * `group-{groupid}`
	//  * `group-{email}`
	//  * `domain-{domain}`
	//  * `project-{team-projectid}`
	//  * `allUsers`
	//  * `allAuthenticatedUsers`
	//  Examples:
	//  * The user `liz@example.com` would be `user-liz@example.com`.
	//  * The group `example@googlegroups.com` would be
	//  `group-example@googlegroups.com`
	//  * All members of the Google Apps for Business domain `example.com` would be
	//  `domain-example.com`
	Entity *string `json:"entity,omitempty"`

	// The ID for the entity, if any.
	EntityID *string `json:"entityID,omitempty"`

	// The email address associated with the entity, if any.
	Email *string `json:"email,omitempty"`

	// The domain associated with the entity, if any.
	Domain *string `json:"domain,omitempty"`

	// The project team associated with the entity, if any.
	ProjectTeam *ProjectTeam `json:"projectTeam,omitempty"`
}

// +kcc:proto=google.storage.v1.Channel
type Channel struct {
	// A UUID or similar unique string that identifies this channel.
	Id *string `json:"id,omitempty"`

	// An opaque ID that identifies the resource being watched on this channel.
	//  Stable across different API versions.
	ResourceID *string `json:"resourceID,omitempty"`

	// A version-specific identifier for the watched resource.
	ResourceUri *string `json:"resourceUri,omitempty"`

	// An arbitrary string delivered to the target address with each notification
	//  delivered over this channel. Optional.
	Token *string `json:"token,omitempty"`

	// Date and time of notification channel expiration. Optional.
	Expiration *string `json:"expiration,omitempty"`

	// The type of delivery mechanism used for this channel.
	Type *string `json:"type,omitempty"`

	// The address where notifications are delivered for this channel.
	Address *string `json:"address,omitempty"`

	// Additional parameters controlling delivery channel behavior. Optional.
	Params map[string]string `json:"params,omitempty"`

	// A Boolean value to indicate whether payload is wanted. Optional.
	Payload *bool `json:"payload,omitempty"`
}

// +kcc:proto=google.storage.v1.ChecksummedData
type ChecksummedData struct {
	// The data.
	Content []byte `json:"content,omitempty"`

	// CRC32C digest of the contents.
	// Crc32c *google_protobuf_UInt32Value `json:"crc32c,omitempty"`
}

// +kcc:proto=google.storage.v1.CommonEnums
type CommonEnums struct {
}

// +kcc:proto=google.storage.v1.CommonObjectRequestParams
type CommonObjectRequestParams struct {
	// Encryption algorithm used with Customer-Supplied Encryption Keys feature.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`

	// Encryption key used with Customer-Supplied Encryption Keys feature.
	EncryptionKey *string `json:"encryptionKey,omitempty"`

	// SHA256 hash of encryption key used with Customer-Supplied Encryption Keys
	//  feature.
	EncryptionKeySha256 *string `json:"encryptionKeySha256,omitempty"`
}

// +kcc:proto=google.storage.v1.CommonRequestParams
type CommonRequestParams struct {
	// Required. Required when using buckets with Requestor Pays feature enabled.
	UserProject *string `json:"userProject,omitempty"`

	// Lets you enforce per-user quotas from a server-side application even in
	//  cases when the user's IP address is unknown. This can occur, for example,
	//  with applications that run cron jobs on App Engine on a user's behalf.
	//  You can choose any arbitrary string that uniquely identifies a user, but it
	//  is limited to 40 characters.
	QuotaUser *string `json:"quotaUser,omitempty"`

	// Subset of fields to include in the response.
	// Fields *google_protobuf_FieldMask `json:"fields,omitempty"`
}

// +kcc:proto=google.storage.v1.ContentRange
type ContentRange struct {
	// The starting offset of the object data.
	Start *int64 `json:"start,omitempty"`

	// The ending offset of the object data.
	End *int64 `json:"end,omitempty"`

	// The complete length of the object data.
	CompleteLength *int64 `json:"completeLength,omitempty"`
}

// +kcc:proto=google.storage.v1.InsertObjectSpec
type InsertObjectSpec struct {
	// Destination object, including its name and its metadata.
	Resource *Object `json:"resource,omitempty"`

	// Apply a predefined set of access controls to this object.
	PredefinedAcl *string `json:"predefinedAcl,omitempty"`

	// Makes the operation conditional on whether the object's current
	//  generation matches the given value. Setting to 0 makes the operation
	//  succeed only if there are no live versions of the object.
	IfGenerationMatch *int64 `json:"ifGenerationMatch,omitempty"`

	// Makes the operation conditional on whether the object's current
	//  generation does not match the given value. If no live object exists, the
	//  precondition fails. Setting to 0 makes the operation succeed only if
	//  there is a live version of the object.
	IfGenerationNotMatch *int64 `json:"ifGenerationNotMatch,omitempty"`

	// Makes the operation conditional on whether the object's current
	//  metageneration matches the given value.
	IfMetagenerationMatch *int64 `json:"ifMetagenerationMatch,omitempty"`

	// Makes the operation conditional on whether the object's current
	//  metageneration does not match the given value.
	IfMetagenerationNotMatch *int64 `json:"ifMetagenerationNotMatch,omitempty"`

	// Set of properties to return. Defaults to `NO_ACL`, unless the
	//  object resource specifies the `acl` property, when it defaults
	//  to `full`.
	Projection *string `json:"projection,omitempty"`
}

// +kcc:proto=google.storage.v1.Notification
type Notification struct {
	// The Cloud PubSub topic to which this subscription publishes. Formatted as:
	//  '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic *string `json:"topic,omitempty"`

	// If present, only send notifications about listed event types. If empty,
	//  sent notifications for all event types.
	EventTypes []string `json:"eventTypes,omitempty"`

	// An optional list of additional attributes to attach to each Cloud PubSub
	//  message published for this notification subscription.
	CustomAttributes map[string]string `json:"customAttributes,omitempty"`

	// HTTP 1.1 [https://tools.ietf.org/html/rfc7232#section-2.3][Entity tag]
	//  for this subscription notification.
	Etag *string `json:"etag,omitempty"`

	// If present, only apply this notification configuration to object names that
	//  begin with this prefix.
	ObjectNamePrefix *string `json:"objectNamePrefix,omitempty"`

	// The desired content of the Payload.
	PayloadFormat *string `json:"payloadFormat,omitempty"`

	// The ID of the notification.
	Id *string `json:"id,omitempty"`
}

// +kcc:proto=google.storage.v1.Object
type Object struct {
	// Content-Encoding of the object data, matching
	//  [https://tools.ietf.org/html/rfc7231#section-3.1.2.2][RFC 7231 §3.1.2.2]
	ContentEncoding *string `json:"contentEncoding,omitempty"`

	// Content-Disposition of the object data, matching
	//  [https://tools.ietf.org/html/rfc6266][RFC 6266].
	ContentDisposition *string `json:"contentDisposition,omitempty"`

	// Cache-Control directive for the object data, matching
	//  [https://tools.ietf.org/html/rfc7234#section-5.2"][RFC 7234 §5.2].
	//  If omitted, and the object is accessible to all anonymous users, the
	//  default will be `public, max-age=3600`.
	CacheControl *string `json:"cacheControl,omitempty"`

	// Access controls on the object.
	Acl []ObjectAccessControl `json:"acl,omitempty"`

	// Content-Language of the object data, matching
	//  [https://tools.ietf.org/html/rfc7231#section-3.1.3.2][RFC 7231 §3.1.3.2].
	ContentLanguage *string `json:"contentLanguage,omitempty"`

	// The version of the metadata for this object at this generation. Used for
	//  preconditions and for detecting changes in metadata. A metageneration
	//  number is only meaningful in the context of a particular generation of a
	//  particular object.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Metageneration *int64 `json:"metageneration,omitempty"`

	// The deletion time of the object. Will be returned if and only if this
	//  version of the object has been deleted.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	TimeDeleted *string `json:"timeDeleted,omitempty"`

	// Content-Type of the object data, matching
	//  [https://tools.ietf.org/html/rfc7231#section-3.1.1.5][RFC 7231 §3.1.1.5].
	//  If an object is stored without a Content-Type, it is served as
	//  `application/octet-stream`.
	ContentType *string `json:"contentType,omitempty"`

	// Content-Length of the object data in bytes, matching
	//  [https://tools.ietf.org/html/rfc7230#section-3.3.2][RFC 7230 §3.3.2].
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Size *int64 `json:"size,omitempty"`

	// The creation time of the object.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	TimeCreated *string `json:"timeCreated,omitempty"`

	// CRC32c checksum. For more information about using the CRC32c
	//  checksum, see
	//  [https://cloud.google.com/storage/docs/hashes-etags#json-api][Hashes and
	//  ETags: Best Practices]. This is a server determined value and should not be
	//  supplied by the user when sending an Object. The server will ignore any
	//  value provided. Users should instead use the object_checksums field on the
	//  InsertObjectRequest when uploading an object.
	// Crc32c *google_protobuf_UInt32Value `json:"crc32c,omitempty"`

	// Number of underlying components that make up this object. Components are
	//  accumulated by compose operations.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	ComponentCount *int32 `json:"componentCount,omitempty"`

	// MD5 hash of the data; encoded using base64 as per
	//  [https://tools.ietf.org/html/rfc4648#section-4][RFC 4648 §4]. For more
	//  information about using the MD5 hash, see
	//  [https://cloud.google.com/storage/docs/hashes-etags#json-api][Hashes and
	//  ETags: Best Practices]. This is a server determined value and should not be
	//  supplied by the user when sending an Object. The server will ignore any
	//  value provided. Users should instead use the object_checksums field on the
	//  InsertObjectRequest when uploading an object.
	Md5Hash *string `json:"md5Hash,omitempty"`

	// HTTP 1.1 Entity tag for the object. See
	//  [https://tools.ietf.org/html/rfc7232#section-2.3][RFC 7232 §2.3].
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Etag *string `json:"etag,omitempty"`

	// The modification time of the object metadata.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Updated *string `json:"updated,omitempty"`

	// Storage class of the object.
	StorageClass *string `json:"storageClass,omitempty"`

	// Cloud KMS Key used to encrypt this object, if the object is encrypted by
	//  such a key.
	KmsKeyName *string `json:"kmsKeyName,omitempty"`

	// The time at which the object's storage class was last changed. When the
	//  object is initially created, it will be set to time_created.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	TimeStorageClassUpdated *string `json:"timeStorageClassUpdated,omitempty"`

	// Whether an object is under temporary hold. While this flag is set to true,
	//  the object is protected against deletion and overwrites.  A common use case
	//  of this flag is regulatory investigations where objects need to be retained
	//  while the investigation is ongoing. Note that unlike event-based hold,
	//  temporary hold does not impact retention expiration time of an object.
	TemporaryHold *bool `json:"temporaryHold,omitempty"`

	// A server-determined value that specifies the earliest time that the
	//  object's retention period expires. This value is in
	//  [https://tools.ietf.org/html/rfc3339][RFC 3339] format.
	//  Note 1: This field is not provided for objects with an active event-based
	//  hold, since retention expiration is unknown until the hold is removed.
	//  Note 2: This value can be provided even when temporary hold is set (so that
	//  the user can reason about policy without having to first unset the
	//  temporary hold).
	RetentionExpirationTime *string `json:"retentionExpirationTime,omitempty"`

	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Whether an object is under event-based hold. Event-based hold is a way to
	//  retain objects until an event occurs, which is signified by the
	//  hold's release (i.e. this value is set to false). After being released (set
	//  to false), such objects will be subject to bucket-level retention (if any).
	//  One sample use case of this flag is for banks to hold loan documents for at
	//  least 3 years after loan is paid in full. Here, bucket-level retention is 3
	//  years and the event is the loan being paid in full. In this example, these
	//  objects will be held intact for any number of years until the event has
	//  occurred (event-based hold on the object is released) and then 3 more years
	//  after that. That means retention duration of the objects begins from the
	//  moment event-based hold transitioned from true to false.
	// EventBasedHold *google_protobuf_BoolValue `json:"eventBasedHold,omitempty"`

	// The name of the object.
	//  Attempting to update this field after the object is created will result in
	//  an error.
	Name *string `json:"name,omitempty"`

	// The ID of the object, including the bucket name, object name, and
	//  generation number.
	//  Attempting to update this field after the object is created will result in
	//  an error.
	Id *string `json:"id,omitempty"`

	// The name of the bucket containing this object.
	//  Attempting to update this field after the object is created will result in
	//  an error.
	Bucket *string `json:"bucket,omitempty"`

	// The content generation of this object. Used for object versioning.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Generation *int64 `json:"generation,omitempty"`

	// The owner of the object. This will always be the uploader of the object.
	//  Attempting to set or update this field will result in a
	//  [FieldViolation][google.rpc.BadRequest.FieldViolation].
	Owner *Owner `json:"owner,omitempty"`

	// Metadata of customer-supplied encryption key, if the object is encrypted by
	//  such a key.
	CustomerEncryption *Object_CustomerEncryption `json:"customerEncryption,omitempty"`

	// A user-specified timestamp set on an object.
	CustomTime *string `json:"customTime,omitempty"`
}

// +kcc:proto=google.storage.v1.Object.CustomerEncryption
type Object_CustomerEncryption struct {
	// The encryption algorithm.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`

	// SHA256 hash value of the encryption key.
	KeySha256 *string `json:"keySha256,omitempty"`
}

// +kcc:proto=google.storage.v1.ObjectAccessControl
type ObjectAccessControl struct {
	// The access permission for the entity.
	Role *string `json:"role,omitempty"`

	// HTTP 1.1 Entity tag for the access-control entry.
	//  See [https://tools.ietf.org/html/rfc7232#section-2.3][RFC 7232 §2.3].
	Etag *string `json:"etag,omitempty"`

	// The ID of the access-control entry.
	Id *string `json:"id,omitempty"`

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty"`

	// The name of the object, if applied to an object.
	Object *string `json:"object,omitempty"`

	// The content generation of the object, if applied to an object.
	Generation *int64 `json:"generation,omitempty"`

	// The entity holding the permission, in one of the following forms:
	//  * `user-{userid}`
	//  * `user-{email}`
	//  * `group-{groupid}`
	//  * `group-{email}`
	//  * `domain-{domain}`
	//  * `project-{team-projectid}`
	//  * `allUsers`
	//  * `allAuthenticatedUsers`
	//  Examples:
	//  * The user `liz@example.com` would be `user-liz@example.com`.
	//  * The group `example@googlegroups.com` would be
	//  `group-example@googlegroups.com`.
	//  * All members of the Google Apps for Business domain `example.com` would be
	//  `domain-example.com`.
	Entity *string `json:"entity,omitempty"`

	// The ID for the entity, if any.
	EntityID *string `json:"entityID,omitempty"`

	// The email address associated with the entity, if any.
	Email *string `json:"email,omitempty"`

	// The domain associated with the entity, if any.
	Domain *string `json:"domain,omitempty"`

	// The project team associated with the entity, if any.
	ProjectTeam *ProjectTeam `json:"projectTeam,omitempty"`
}

// +kcc:proto=google.storage.v1.ObjectChecksums
type ObjectChecksums struct {
	// CRC32C digest of the object data. Computed by the GCS service for
	//  all written objects, and validated by the GCS service against
	//  client-supplied values if present in an InsertObjectRequest.
	// Crc32c *google_protobuf_UInt32Value `json:"crc32c,omitempty"`

	// Hex-encoded MD5 hash of the object data (hexdigest). Whether/how this
	//  checksum is provided and validated is service-dependent.
	Md5Hash *string `json:"md5Hash,omitempty"`
}

// +kcc:proto=google.storage.v1.Owner
type Owner struct {
	// The entity, in the form `user-`*userId*.
	Entity *string `json:"entity,omitempty"`

	// The ID for the entity.
	EntityID *string `json:"entityID,omitempty"`
}

// +kcc:proto=google.storage.v1.ProjectTeam
type ProjectTeam struct {
	// The project number.
	ProjectNumber *string `json:"projectNumber,omitempty"`

	// The team.
	Team *string `json:"team,omitempty"`
}

// +kcc:proto=google.storage.v1.ServiceAccount
type ServiceAccount struct {
	// The ID of the notification.
	EmailAddress *string `json:"emailAddress,omitempty"`
}
