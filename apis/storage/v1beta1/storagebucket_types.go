// Copyright 2026 Google LLC
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

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageBucketGVK = GroupVersion.WithKind("StorageBucket")

// +kcc:proto=google.storage.v2.Action
type BucketAction struct {
	/* The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Action.storage_class
	StorageClass *string `json:"storageClass,omitempty"`

	/* The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload. */
	// +kcc:proto:field=google.storage.v2.Action.type
	Type string `json:"type"`
}

// +kcc:proto=google.storage.v2.BucketAutoclass
type BucketAutoclass struct {
	/* While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern. */
	// +kcc:proto:field=google.storage.v2.BucketAutoclass.enabled
	Enabled bool `json:"enabled"`
}

// +kcc:proto=google.storage.v2.Condition
type BucketCondition struct {
	/* Minimum age of an object in days to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.age
	Age *int64 `json:"age,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.created_before
	CreatedBefore *string `json:"createdBefore,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.custom_time_before
	CustomTimeBefore *string `json:"customTimeBefore,omitempty"`

	/* Number of days elapsed since the user-specified timestamp set on an object. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.days_since_custom_time
	DaysSinceCustomTime *int64 `json:"daysSinceCustomTime,omitempty"`

	/* Number of days elapsed since the noncurrent timestamp of an object. This
	condition is relevant only for versioned objects. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.days_since_noncurrent_time
	DaysSinceNoncurrentTime *int64 `json:"daysSinceNoncurrentTime,omitempty"`

	/* One or more matching name prefixes to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.matches_prefix
	MatchesPrefix []string `json:"matchesPrefix,omitempty"`

	/* Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.matches_storage_class
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty"`

	/* One or more matching name suffixes to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.matches_suffix
	MatchesSuffix []string `json:"matchesSuffix,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.noncurrent_time_before
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty"`

	/* Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.num_newer_versions
	NumNewerVersions *int64 `json:"numNewerVersions,omitempty"`

	/* Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY". */
	// +optional
	// +kcc:proto:field=google.storage.v2.Condition.is_live
	WithState *string `json:"withState,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketCors
type BucketCors struct {
	/* The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketCors.max_age_seconds
	MaxAgeSeconds *int64 `json:"maxAgeSeconds,omitempty"`

	/* The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method". */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketCors.method
	Method []string `json:"method,omitempty"`

	/* The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin". */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketCors.origin
	Origin []string `json:"origin,omitempty"`

	/* The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketCors.response_header
	ResponseHeader []string `json:"responseHeader,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketCustomPlacementConfig
type BucketCustomPlacementConfig struct {
	/* The list of individual regions that comprise a dual-region bucket. See the docs for a list of acceptable regions. */
	// +kcc:proto:field=google.storage.v2.BucketCustomPlacementConfig.data_locations
	DataLocations []string `json:"dataLocations"`
}

// +kcc:proto=google.storage.v2.BucketEncryption
type BucketEncryption struct {
	// +kcc:proto:field=google.storage.v2.BucketEncryption.default_kms_key_name
	KmsKeyRef refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef"`
}

// +kcc:proto=google.storage.v2.Rule
type BucketLifecycleRule struct {
	/* The Lifecycle Rule's action configuration. A single block of this type is supported. */
	// +kcc:proto:field=google.storage.v2.Rule.action
	Action BucketAction `json:"action"`

	/* The Lifecycle Rule's condition configuration. */
	// +kcc:proto:field=google.storage.v2.Rule.condition
	Condition BucketCondition `json:"condition"`
}

// +kcc:proto=google.storage.v2.BucketLogging
type BucketLogging struct {
	/* The bucket that will receive log objects. */
	// +kcc:proto:field=google.storage.v2.BucketLogging.log_bucket
	LogBucket string `json:"logBucket"`

	/* The object prefix for log objects. If it's not provided, by default Google Cloud Storage sets this to this bucket's name. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketLogging.log_object_prefix
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketRetentionPolicy
type BucketRetentionPolicy struct {
	/* If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketRetentionPolicy.is_locked
	IsLocked *bool `json:"isLocked,omitempty"`

	/* The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds. */
	// +kcc:proto:field=google.storage.v2.BucketRetentionPolicy.retention_period
	RetentionPeriod int64 `json:"retentionPeriod"`
}

// +kcc:proto=google.storage.v2.BucketSoftDeletePolicy
type BucketSoftDeletePolicy struct {
	/* The duration in seconds that soft-deleted objects in the bucket will be retained and cannot be permanently deleted. Default value is 604800. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketSoftDeletePolicy.retention_duration_seconds
	RetentionDurationSeconds *int64 `json:"retentionDurationSeconds,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket_IpFilter
type BucketIPFilter struct {
	/* The mode of the IP filter. Possible values: ENABLED, DISABLED. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket_IpFilter.mode
	Mode *string `json:"mode,omitempty"`

	/* The public network source IPs to allow access from. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket_IpFilter.public_network_source
	PublicNetworkSource *BucketIPFilterPublicNetworkSource `json:"publicNetworkSource,omitempty"`

	/* The VPC network sources IPs to allow access from. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket_IpFilter.vpc_network_sources
	VpcNetworkSources []BucketIPFilterVpcNetworkSource `json:"vpcNetworkSources,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket_IpFilter_VpcNetworkSource
type BucketIPFilterVpcNetworkSource struct {
	/* The network to allow access from. */
	// +kcc:proto:field=google.storage.v2.Bucket_IpFilter_VpcNetworkSource.network
	NetworkRef computev1beta1.ComputeNetworkRef `json:"networkRef"`

	/* The list of VPC network source IPs to allow access from. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket_IpFilter_VpcNetworkSource.allowed_ip_cidr_ranges
	AllowedIpRanges []string `json:"allowedIpRanges,omitempty"`
}

// +kcc:proto=google.storage.v2.Bucket_IpFilter_PublicNetworkSource
type BucketIPFilterPublicNetworkSource struct {
	/* The list of public network source IPs to allow access from. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket_IpFilter_PublicNetworkSource.allowed_ip_cidr_ranges
	AllowedIpRanges []string `json:"allowedIpRanges,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketVersioning
type BucketVersioning struct {
	/* While set to true, versioning is fully enabled for this bucket. */
	// +kcc:proto:field=google.storage.v2.BucketVersioning.enabled
	Enabled bool `json:"enabled"`
}

// +kcc:proto=google.storage.v2.BucketWebsite
type BucketWebsite struct {
	/* Behaves as the bucket's directory index where missing objects are treated as potential directories. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketWebsite.main_page_suffix
	MainPageSuffix *string `json:"mainPageSuffix,omitempty"`

	/* The custom object to return when a requested resource is not found. */
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketWebsite.not_found_page
	NotFoundPage *string `json:"notFoundPage,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketIamConfiguration
type BucketIAMConfiguration struct {
	// +optional
	// +kcc:proto:field=google.storage.v2.BucketIamConfiguration.bucket_policy_only
	BucketPolicyOnly *BucketPolicyOnly `json:"bucketPolicyOnly,omitempty"`

	// +optional
	// +kcc:proto:field=google.storage.v2.BucketIamConfiguration.public_access_prevention
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty"`

	// +optional
	// +kcc:proto:field=google.storage.v2.BucketIamConfiguration.uniform_bucket_level_access
	UniformBucketLevelAccess *UniformBucketLevelAccess `json:"uniformBucketLevelAccess,omitempty"`
}

// +kcc:proto=google.storage.v2.BucketPolicyOnly
type BucketPolicyOnly struct {
	// +kcc:proto:field=google.storage.v2.BucketPolicyOnly.enabled
	Enabled bool `json:"enabled"`
}

// +kcc:proto=google.storage.v2.UniformBucketLevelAccess
type UniformBucketLevelAccess struct {
	// +kcc:proto:field=google.storage.v2.UniformBucketLevelAccess.enabled
	Enabled bool `json:"enabled"`
}

// StorageBucketSpec defines the desired state of StorageBucket
// +kcc:spec:proto=google.storage.v2.Bucket
type StorageBucketSpec struct {
	// The project that this resource belongs to.
	// +optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	/* The bucket's autoclass configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.autoclass
	Autoclass *BucketAutoclass `json:"autoclass,omitempty"`

	/* The bucket's Cross-Origin Resource Sharing (CORS) configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.cors
	Cors []BucketCors `json:"cors,omitempty"`

	/* The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.custom_placement_config
	CustomPlacementConfig *BucketCustomPlacementConfig `json:"customPlacementConfig,omitempty"`

	/* Whether or not to automatically apply an eventBasedHold to new objects added to the bucket. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.default_event_based_hold
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty"`

	/* The bucket's encryption configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.encryption
	Encryption *BucketEncryption `json:"encryption,omitempty"`

	/* The bucket's Lifecycle Rules configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.lifecycle
	LifecycleRule []BucketLifecycleRule `json:"lifecycleRule,omitempty"`

	/* The Google Cloud Storage location. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.location
	Location *string `json:"location,omitempty"`

	/* DEPRECATED. Please use the `uniformBucketLevelAccess` field as this field has been renamed by Google. The `uniformBucketLevelAccess` field will supersede this field. Enables Bucket PolicyOnly access to a bucket. */
	// +optional
	BucketPolicyOnly *bool `json:"bucketPolicyOnly,omitempty"`

	/* Enables uniform bucket-level access on a bucket. */
	// +optional
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty"`

	/* The bucket's Access & Storage Logs configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.logging
	Logging *BucketLogging `json:"logging,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.name
	ResourceID *string `json:"resourceID,omitempty"`

	/* Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.retention_policy
	RetentionPolicy *BucketRetentionPolicy `json:"retentionPolicy,omitempty"`

	/* The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.soft_delete_policy
	SoftDeletePolicy *BucketSoftDeletePolicy `json:"softDeletePolicy,omitempty"`

	/* The bucket's IP filter configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.ip_filter
	IpFilter *BucketIPFilter `json:"ipFilter,omitempty"`

	/* The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.storage_class
	StorageClass *string `json:"storageClass,omitempty"`

	/* The bucket's IAM configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.iam_configuration
	IAMConfiguration *BucketIAMConfiguration `json:"iamConfiguration,omitempty"`

	/* The bucket's Versioning configuration. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.versioning
	Versioning *BucketVersioning `json:"versioning,omitempty"`

	/* Configuration if the bucket acts as a website. */
	// +optional
	// +kcc:proto:field=google.storage.v2.Bucket.website
	Website *BucketWebsite `json:"website,omitempty"`
}

type BucketObservedStateStatus struct {
	/* The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy. */
	// +optional
	SoftDeletePolicy *BucketSoftDeletePolicyStatus `json:"softDeletePolicy,omitempty"`
}

type BucketSoftDeletePolicyStatus struct {
	/* Server-determined value that indicates the time from which the policy, or one with a greater retention, was effective. This value is in RFC 3339 format. */
	// +optional
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	/* The duration in seconds that soft-deleted objects in the bucket will be retained and cannot be permanently deleted. Default value is 604800. */
	// +optional
	RetentionDurationSeconds *int64 `json:"retentionDurationSeconds,omitempty"`
}

type StorageBucketStatus struct {
	/* Conditions represent the latest available observations of the
	   StorageBucket's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	// +optional
	ObservedState *BucketObservedStateStatus `json:"observedState,omitempty"`

	/* The URI of the created resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* The base URL of the bucket, in the format gs://<bucket-name>. */
	// +optional
	Url *string `json:"url,omitempty"`

	/* A unique specifier for the StorageBucket resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragebucket;gcpstoragebuckets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageBucket is the Schema for the storage API
// +k8s:openapi-gen=true
type StorageBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageBucketSpec   `json:"spec,omitempty"`
	Status StorageBucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageBucketList contains a list of StorageBucket
type StorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageBucket{}, &StorageBucketList{})
}
