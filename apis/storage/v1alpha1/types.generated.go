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


// +kcc:proto=google.storage.v2.Bucket
type Bucket struct {
	// Immutable. The name of the bucket.
	//  Format: `projects/{project}/buckets/{bucket}`
	// +kcc:proto:field=google.storage.v2.Bucket.name
	Name *string `json:"name,omitempty"`

	// The etag of the bucket.
	//  If included in the metadata of an UpdateBucketRequest, the operation will
	//  only be performed if the etag matches that of the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.etag
	Etag *string `json:"etag,omitempty"`

	// Immutable. The project which owns this bucket, in the format of
	//  "projects/{projectIdentifier}".
	//  {projectIdentifier} can be the project ID or project number.
	// +kcc:proto:field=google.storage.v2.Bucket.project
	Project *string `json:"project,omitempty"`

	// Immutable. The location of the bucket. Object data for objects in the
	//  bucket resides in physical storage within this region.  Defaults to `US`.
	//  See the
	//  [https://developers.google.com/storage/docs/concepts-techniques#specifyinglocations"][developer's
	//  guide] for the authoritative list. Attempting to update this field after
	//  the bucket is created will result in an error.
	// +kcc:proto:field=google.storage.v2.Bucket.location
	Location *string `json:"location,omitempty"`

	// The bucket's default storage class, used whenever no storageClass is
	//  specified for a newly-created object. This defines how objects in the
	//  bucket are stored and determines the SLA and the cost of storage.
	//  If this value is not specified when the bucket is created, it will default
	//  to `STANDARD`. For more information, see
	//  https://developers.google.com/storage/docs/storage-classes.
	// +kcc:proto:field=google.storage.v2.Bucket.storage_class
	StorageClass *string `json:"storageClass,omitempty"`

	// The recovery point objective for cross-region replication of the bucket.
	//  Applicable only for dual- and multi-region buckets. "DEFAULT" uses default
	//  replication. "ASYNC_TURBO" enables turbo replication, valid for dual-region
	//  buckets only. If rpo is not specified when the bucket is created, it
	//  defaults to "DEFAULT". For more information, see
	//  https://cloud.google.com/storage/docs/availability-durability#turbo-replication.
	// +kcc:proto:field=google.storage.v2.Bucket.rpo
	Rpo *string `json:"rpo,omitempty"`

	// Access controls on the bucket.
	//  If iam_config.uniform_bucket_level_access is enabled on this bucket,
	//  requests to set, read, or modify acl is an error.
	// +kcc:proto:field=google.storage.v2.Bucket.acl
	Acl []BucketAccessControl `json:"acl,omitempty"`

	// Default access controls to apply to new objects when no ACL is provided.
	//  If iam_config.uniform_bucket_level_access is enabled on this bucket,
	//  requests to set, read, or modify acl is an error.
	// +kcc:proto:field=google.storage.v2.Bucket.default_object_acl
	DefaultObjectAcl []ObjectAccessControl `json:"defaultObjectAcl,omitempty"`

	// The bucket's lifecycle config. See
	//  [https://developers.google.com/storage/docs/lifecycle]Lifecycle Management]
	//  for more information.
	// +kcc:proto:field=google.storage.v2.Bucket.lifecycle
	Lifecycle *Bucket_Lifecycle `json:"lifecycle,omitempty"`

	// The bucket's [https://www.w3.org/TR/cors/][Cross-Origin Resource Sharing]
	//  (CORS) config.
	// +kcc:proto:field=google.storage.v2.Bucket.cors
	Cors []Bucket_Cors `json:"cors,omitempty"`

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
	// +kcc:proto:field=google.storage.v2.Bucket.default_event_based_hold
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty"`

	// User-provided labels, in key/value pairs.
	// +kcc:proto:field=google.storage.v2.Bucket.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The bucket's website config, controlling how the service behaves
	//  when accessing bucket contents as a web site. See the
	//  [https://cloud.google.com/storage/docs/static-website][Static Website
	//  Examples] for more information.
	// +kcc:proto:field=google.storage.v2.Bucket.website
	Website *Bucket_Website `json:"website,omitempty"`

	// The bucket's versioning config.
	// +kcc:proto:field=google.storage.v2.Bucket.versioning
	Versioning *Bucket_Versioning `json:"versioning,omitempty"`

	// The bucket's logging config, which defines the destination bucket
	//  and name prefix (if any) for the current bucket's logs.
	// +kcc:proto:field=google.storage.v2.Bucket.logging
	Logging *Bucket_Logging `json:"logging,omitempty"`

	// Encryption config for a bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.encryption
	Encryption *Bucket_Encryption `json:"encryption,omitempty"`

	// The bucket's billing config.
	// +kcc:proto:field=google.storage.v2.Bucket.billing
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
	// +kcc:proto:field=google.storage.v2.Bucket.retention_policy
	RetentionPolicy *Bucket_RetentionPolicy `json:"retentionPolicy,omitempty"`

	// The bucket's IAM config.
	// +kcc:proto:field=google.storage.v2.Bucket.iam_config
	IamConfig *Bucket_IamConfig `json:"iamConfig,omitempty"`

	// Reserved for future use.
	// +kcc:proto:field=google.storage.v2.Bucket.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Configuration that, if present, specifies the data placement for a
	//  [https://cloud.google.com/storage/docs/locations#location-dr][configurable
	//  dual-region].
	// +kcc:proto:field=google.storage.v2.Bucket.custom_placement_config
	CustomPlacementConfig *Bucket_CustomPlacementConfig `json:"customPlacementConfig,omitempty"`

	// The bucket's Autoclass configuration. If there is no configuration, the
	//  Autoclass feature will be disabled and have no effect on the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.autoclass
	Autoclass *Bucket_Autoclass `json:"autoclass,omitempty"`

	// Optional. The bucket's hierarchical namespace configuration. If there is no
	//  configuration, the hierarchical namespace feature will be disabled and have
	//  no effect on the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.hierarchical_namespace
	HierarchicalNamespace *Bucket_HierarchicalNamespace `json:"hierarchicalNamespace,omitempty"`

	// Optional. The bucket's soft delete policy. The soft delete policy prevents
	//  soft-deleted objects from being permanently deleted.
	// +kcc:proto:field=google.storage.v2.Bucket.soft_delete_policy
	SoftDeletePolicy *Bucket_SoftDeletePolicy `json:"softDeletePolicy,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Autoclass
type Bucket_Autoclass struct {
	// Enables Autoclass.
	// +kcc:proto:field=google.storage.v2.Bucket.Autoclass.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// An object in an Autoclass bucket will eventually cool down to the
	//  terminal storage class if there is no access to the object.
	//  The only valid values are NEARLINE and ARCHIVE.
	// +kcc:proto:field=google.storage.v2.Bucket.Autoclass.terminal_storage_class
	TerminalStorageClass *string `json:"terminalStorageClass,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Billing
type Bucket_Billing struct {
	// When set to true, Requester Pays is enabled for this bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.Billing.requester_pays
	RequesterPays *bool `json:"requesterPays,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Cors
type Bucket_Cors struct {
	// The list of Origins eligible to receive CORS response headers. See
	//  [https://tools.ietf.org/html/rfc6454][RFC 6454] for more on origins.
	//  Note: "*" is permitted in the list of origins, and means "any Origin".
	// +kcc:proto:field=google.storage.v2.Bucket.Cors.origin
	Origin []string `json:"origin,omitempty"`

	// The list of HTTP methods on which to include CORS response headers,
	//  (`GET`, `OPTIONS`, `POST`, etc) Note: "*" is permitted in the list of
	//  methods, and means "any method".
	// +kcc:proto:field=google.storage.v2.Bucket.Cors.method
	Method []string `json:"method,omitempty"`

	// The list of HTTP headers other than the
	//  [https://www.w3.org/TR/cors/#simple-response-header][simple response
	//  headers] to give permission for the user-agent to share across domains.
	// +kcc:proto:field=google.storage.v2.Bucket.Cors.response_header
	ResponseHeader []string `json:"responseHeader,omitempty"`

	// The value, in seconds, to return in the
	//  [https://www.w3.org/TR/cors/#access-control-max-age-response-header][Access-Control-Max-Age
	//  header] used in preflight responses.
	// +kcc:proto:field=google.storage.v2.Bucket.Cors.max_age_seconds
	MaxAgeSeconds *int32 `json:"maxAgeSeconds,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.CustomPlacementConfig
type Bucket_CustomPlacementConfig struct {
	// List of locations to use for data placement.
	// +kcc:proto:field=google.storage.v2.Bucket.CustomPlacementConfig.data_locations
	DataLocations []string `json:"dataLocations,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Encryption
type Bucket_Encryption struct {
	// The name of the Cloud KMS key that will be used to encrypt objects
	//  inserted into this bucket, if no encryption method is specified.
	// +kcc:proto:field=google.storage.v2.Bucket.Encryption.default_kms_key
	DefaultKMSKey *string `json:"defaultKMSKey,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.HierarchicalNamespace
type Bucket_HierarchicalNamespace struct {
	// Optional. Enables the hierarchical namespace feature.
	// +kcc:proto:field=google.storage.v2.Bucket.HierarchicalNamespace.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.IamConfig
type Bucket_IamConfig struct {
	// Bucket restriction options currently enforced on the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.IamConfig.uniform_bucket_level_access
	UniformBucketLevelAccess *Bucket_IamConfig_UniformBucketLevelAccess `json:"uniformBucketLevelAccess,omitempty"`

	// Whether IAM will enforce public access prevention. Valid values are
	//  "enforced" or "inherited".
	// +kcc:proto:field=google.storage.v2.Bucket.IamConfig.public_access_prevention
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.IamConfig.UniformBucketLevelAccess
type Bucket_IamConfig_UniformBucketLevelAccess struct {
	// If set, access checks only use bucket-level IAM policies or above.
	// +kcc:proto:field=google.storage.v2.Bucket.IamConfig.UniformBucketLevelAccess.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The deadline time for changing
	//  `iam_config.uniform_bucket_level_access.enabled` from `true` to
	//  `false`. Mutable until the specified deadline is reached, but not
	//  afterward.
	// +kcc:proto:field=google.storage.v2.Bucket.IamConfig.UniformBucketLevelAccess.lock_time
	LockTime *string `json:"lockTime,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Lifecycle
type Bucket_Lifecycle struct {
	// A lifecycle management rule, which is made of an action to take and the
	//  condition(s) under which the action will be taken.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.rule
	Rule []Bucket_Lifecycle_Rule `json:"rule,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Lifecycle.Rule
type Bucket_Lifecycle_Rule struct {
	// The action to take.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.action
	Action *Bucket_Lifecycle_Rule_Action `json:"action,omitempty"`

	// The condition(s) under which the action will be taken.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.condition
	Condition *Bucket_Lifecycle_Rule_Condition `json:"condition,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Lifecycle.Rule.Action
type Bucket_Lifecycle_Rule_Action struct {
	// Type of the action. Currently, only `Delete`, `SetStorageClass`, and
	//  `AbortIncompleteMultipartUpload` are supported.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Action.type
	Type *string `json:"type,omitempty"`

	// Target storage class. Required iff the type of the action is
	//  SetStorageClass.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Action.storage_class
	StorageClass *string `json:"storageClass,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Lifecycle.Rule.Condition
type Bucket_Lifecycle_Rule_Condition struct {
	// Age of an object (in days). This condition is satisfied when an
	//  object reaches the specified age.
	//  A value of 0 indicates that all objects immediately match this
	//  condition.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.age_days
	AgeDays *int32 `json:"ageDays,omitempty"`

	// This condition is satisfied when an object is created before midnight
	//  of the specified date in UTC.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.created_before
	CreatedBefore *Date `json:"createdBefore,omitempty"`

	// Relevant only for versioned objects. If the value is
	//  `true`, this condition matches live objects; if the value
	//  is `false`, it matches archived objects.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.is_live
	IsLive *bool `json:"isLive,omitempty"`

	// Relevant only for versioned objects. If the value is N, this
	//  condition is satisfied when there are at least N versions (including
	//  the live version) newer than this version of the object.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.num_newer_versions
	NumNewerVersions *int32 `json:"numNewerVersions,omitempty"`

	// Objects having any of the storage classes specified by this condition
	//  will be matched. Values include `MULTI_REGIONAL`, `REGIONAL`,
	//  `NEARLINE`, `COLDLINE`, `STANDARD`, and
	//  `DURABLE_REDUCED_AVAILABILITY`.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.matches_storage_class
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty"`

	// Number of days that have elapsed since the custom timestamp set on an
	//  object.
	//  The value of the field must be a nonnegative integer.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.days_since_custom_time
	DaysSinceCustomTime *int32 `json:"daysSinceCustomTime,omitempty"`

	// An object matches this condition if the custom timestamp set on the
	//  object is before the specified date in UTC.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.custom_time_before
	CustomTimeBefore *Date `json:"customTimeBefore,omitempty"`

	// This condition is relevant only for versioned objects. An object
	//  version satisfies this condition only if these many days have been
	//  passed since it became noncurrent. The value of the field must be a
	//  nonnegative integer. If it's zero, the object version will become
	//  eligible for Lifecycle action as soon as it becomes noncurrent.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.days_since_noncurrent_time
	DaysSinceNoncurrentTime *int32 `json:"daysSinceNoncurrentTime,omitempty"`

	// This condition is relevant only for versioned objects. An object
	//  version satisfies this condition only if it became noncurrent before
	//  the specified date in UTC.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.noncurrent_time_before
	NoncurrentTimeBefore *Date `json:"noncurrentTimeBefore,omitempty"`

	// List of object name prefixes. If any prefix exactly matches the
	//  beginning of the object name, the condition evaluates to true.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.matches_prefix
	MatchesPrefix []string `json:"matchesPrefix,omitempty"`

	// List of object name suffixes. If any suffix exactly matches the
	//  end of the object name, the condition evaluates to true.
	// +kcc:proto:field=google.storage.v2.Bucket.Lifecycle.Rule.Condition.matches_suffix
	MatchesSuffix []string `json:"matchesSuffix,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Logging
type Bucket_Logging struct {
	// The destination bucket where the current bucket's logs should be placed,
	//  using path format (like `projects/123456/buckets/foo`).
	// +kcc:proto:field=google.storage.v2.Bucket.Logging.log_bucket
	LogBucket *string `json:"logBucket,omitempty"`

	// A prefix for log object names.
	// +kcc:proto:field=google.storage.v2.Bucket.Logging.log_object_prefix
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.RetentionPolicy
type Bucket_RetentionPolicy struct {
	// Server-determined value that indicates the time from which policy was
	//  enforced and effective.
	// +kcc:proto:field=google.storage.v2.Bucket.RetentionPolicy.effective_time
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Once locked, an object retention policy cannot be modified.
	// +kcc:proto:field=google.storage.v2.Bucket.RetentionPolicy.is_locked
	IsLocked *bool `json:"isLocked,omitempty"`

	// The duration that objects need to be retained. Retention duration must be
	//  greater than zero and less than 100 years. Note that enforcement of
	//  retention periods less than a day is not guaranteed. Such periods should
	//  only be used for testing purposes. Any `nanos` value specified will be
	//  rounded down to the nearest second.
	// +kcc:proto:field=google.storage.v2.Bucket.RetentionPolicy.retention_duration
	RetentionDuration *string `json:"retentionDuration,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.SoftDeletePolicy
type Bucket_SoftDeletePolicy struct {
	// The period of time that soft-deleted objects in the bucket must be
	//  retained and cannot be permanently deleted. The duration must be greater
	//  than or equal to 7 days and less than 1 year.
	// +kcc:proto:field=google.storage.v2.Bucket.SoftDeletePolicy.retention_duration
	RetentionDuration *string `json:"retentionDuration,omitempty"`

	// Time from which the policy was effective. This is service-provided.
	// +kcc:proto:field=google.storage.v2.Bucket.SoftDeletePolicy.effective_time
	EffectiveTime *string `json:"effectiveTime,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Versioning
type Bucket_Versioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.Versioning.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Website
type Bucket_Website struct {
	// If the requested object path is missing, the service will ensure the path
	//  has a trailing '/', append this suffix, and attempt to retrieve the
	//  resulting object. This allows the creation of `index.html`
	//  objects to represent directory pages.
	// +kcc:proto:field=google.storage.v2.Bucket.Website.main_page_suffix
	MainPageSuffix *string `json:"mainPageSuffix,omitempty"`

	// If the requested object path is missing, and any
	//  `mainPageSuffix` object is missing, if applicable, the service
	//  will return the named object from this bucket as the content for a
	//  [https://tools.ietf.org/html/rfc7231#section-6.5.4][404 Not Found]
	//  result.
	// +kcc:proto:field=google.storage.v2.Bucket.Website.not_found_page
	NotFoundPage *string `json:"notFoundPage,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketAccessControl
type BucketAccessControl struct {
	// The access permission for the entity.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.role
	Role *string `json:"role,omitempty"`

	// The ID of the access-control entry.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.id
	ID *string `json:"id,omitempty"`

	// The entity holding the permission, in one of the following forms:
	//  * `user-{userid}`
	//  * `user-{email}`
	//  * `group-{groupid}`
	//  * `group-{email}`
	//  * `domain-{domain}`
	//  * `project-{team}-{projectnumber}`
	//  * `project-{team}-{projectid}`
	//  * `allUsers`
	//  * `allAuthenticatedUsers`
	//  Examples:
	//  * The user `liz@example.com` would be `user-liz@example.com`.
	//  * The group `example@googlegroups.com` would be
	//  `group-example@googlegroups.com`
	//  * All members of the Google Apps for Business domain `example.com` would be
	//  `domain-example.com`
	//  For project entities, `project-{team}-{projectnumber}` format will be
	//  returned on response.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.entity
	Entity *string `json:"entity,omitempty"`

	// The ID for the entity, if any.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.entity_id
	EntityID *string `json:"entityID,omitempty"`

	// The etag of the BucketAccessControl.
	//  If included in the metadata of an update or delete request message, the
	//  operation operation will only be performed if the etag matches that of the
	//  bucket's BucketAccessControl.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.etag
	Etag *string `json:"etag,omitempty"`

	// The email address associated with the entity, if any.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.email
	Email *string `json:"email,omitempty"`

	// The domain associated with the entity, if any.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.domain
	Domain *string `json:"domain,omitempty"`

	// The project team associated with the entity, if any.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.project_team
	ProjectTeam *ProjectTeam `json:"projectTeam,omitempty"`
}

// +kcc:proto=google.storage.v2.ObjectAccessControl
type ObjectAccessControl struct {
	// The access permission for the entity. One of the following values:
	//  * `READER`
	//  * `WRITER`
	//  * `OWNER`
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.role
	Role *string `json:"role,omitempty"`

	// The ID of the access-control entry.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.id
	ID *string `json:"id,omitempty"`

	// The entity holding the permission, in one of the following forms:
	//  * `user-{userid}`
	//  * `user-{email}`
	//  * `group-{groupid}`
	//  * `group-{email}`
	//  * `domain-{domain}`
	//  * `project-{team}-{projectnumber}`
	//  * `project-{team}-{projectid}`
	//  * `allUsers`
	//  * `allAuthenticatedUsers`
	//  Examples:
	//  * The user `liz@example.com` would be `user-liz@example.com`.
	//  * The group `example@googlegroups.com` would be
	//  `group-example@googlegroups.com`.
	//  * All members of the Google Apps for Business domain `example.com` would be
	//  `domain-example.com`.
	//  For project entities, `project-{team}-{projectnumber}` format will be
	//  returned on response.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.entity
	Entity *string `json:"entity,omitempty"`

	// The ID for the entity, if any.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.entity_id
	EntityID *string `json:"entityID,omitempty"`

	// The etag of the ObjectAccessControl.
	//  If included in the metadata of an update or delete request message, the
	//  operation will only be performed if the etag matches that of the live
	//  object's ObjectAccessControl.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.etag
	Etag *string `json:"etag,omitempty"`

	// The email address associated with the entity, if any.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.email
	Email *string `json:"email,omitempty"`

	// The domain associated with the entity, if any.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.domain
	Domain *string `json:"domain,omitempty"`

	// The project team associated with the entity, if any.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.project_team
	ProjectTeam *ProjectTeam `json:"projectTeam,omitempty"`
}

// +kcc:proto=google.storage.v2.Owner
type Owner struct {
	// The entity, in the form `user-`*userId*.
	// +kcc:proto:field=google.storage.v2.Owner.entity
	Entity *string `json:"entity,omitempty"`

	// The ID for the entity.
	// +kcc:proto:field=google.storage.v2.Owner.entity_id
	EntityID *string `json:"entityID,omitempty"`
}

// +kcc:proto=google.storage.v2.ProjectTeam
type ProjectTeam struct {
	// The project number.
	// +kcc:proto:field=google.storage.v2.ProjectTeam.project_number
	ProjectNumber *string `json:"projectNumber,omitempty"`

	// The team.
	// +kcc:proto:field=google.storage.v2.ProjectTeam.team
	Team *string `json:"team,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket
type BucketObservedState struct {
	// Output only. The user-chosen part of the bucket name. The `{bucket}`
	//  portion of the `name` field. For globally unique buckets, this is equal to
	//  the "bucket name" of other Cloud Storage APIs. Example: "pub".
	// +kcc:proto:field=google.storage.v2.Bucket.bucket_id
	BucketID *string `json:"bucketID,omitempty"`

	// Output only. The metadata generation of this bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.metageneration
	Metageneration *int64 `json:"metageneration,omitempty"`

	// Output only. The location type of the bucket (region, dual-region,
	//  multi-region, etc).
	// +kcc:proto:field=google.storage.v2.Bucket.location_type
	LocationType *string `json:"locationType,omitempty"`

	// Access controls on the bucket.
	//  If iam_config.uniform_bucket_level_access is enabled on this bucket,
	//  requests to set, read, or modify acl is an error.
	// +kcc:proto:field=google.storage.v2.Bucket.acl
	Acl []BucketAccessControlObservedState `json:"acl,omitempty"`

	// Default access controls to apply to new objects when no ACL is provided.
	//  If iam_config.uniform_bucket_level_access is enabled on this bucket,
	//  requests to set, read, or modify acl is an error.
	// +kcc:proto:field=google.storage.v2.Bucket.default_object_acl
	DefaultObjectAcl []ObjectAccessControlObservedState `json:"defaultObjectAcl,omitempty"`

	// Output only. The creation time of the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The modification time of the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The owner of the bucket. This is always the project team's
	//  owner group.
	// +kcc:proto:field=google.storage.v2.Bucket.owner
	Owner *Owner `json:"owner,omitempty"`

	// The bucket's Autoclass configuration. If there is no configuration, the
	//  Autoclass feature will be disabled and have no effect on the bucket.
	// +kcc:proto:field=google.storage.v2.Bucket.autoclass
	Autoclass *Bucket_AutoclassObservedState `json:"autoclass,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket.Autoclass
type Bucket_AutoclassObservedState struct {
	// Output only. Latest instant at which the `enabled` field was set to true
	//  after being disabled/unconfigured or set to false after being enabled. If
	//  Autoclass is enabled when the bucket is created, the toggle_time is set
	//  to the bucket creation time.
	// +kcc:proto:field=google.storage.v2.Bucket.Autoclass.toggle_time
	ToggleTime *string `json:"toggleTime,omitempty"`

	// Output only. Latest instant at which the autoclass terminal storage class
	//  was updated.
	// +kcc:proto:field=google.storage.v2.Bucket.Autoclass.terminal_storage_class_update_time
	TerminalStorageClassUpdateTime *string `json:"terminalStorageClassUpdateTime,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketAccessControl
type BucketAccessControlObservedState struct {
	// Output only. The alternative entity format, if exists. For project
	//  entities, `project-{team}-{projectid}` format will be returned on response.
	// +kcc:proto:field=google.storage.v2.BucketAccessControl.entity_alt
	EntityAlt *string `json:"entityAlt,omitempty"`
}

// +kcc:proto=google.storage.v2.ObjectAccessControl
type ObjectAccessControlObservedState struct {
	// Output only. The alternative entity format, if exists. For project
	//  entities, `project-{team}-{projectid}` format will be returned on response.
	// +kcc:proto:field=google.storage.v2.ObjectAccessControl.entity_alt
	EntityAlt *string `json:"entityAlt,omitempty"`
}
