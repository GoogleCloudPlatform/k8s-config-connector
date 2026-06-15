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

// StorageBucketSpec defines the desired state of StorageBucket
// +kcc:spec:proto=google.storage.v1.Bucket
type StorageBucketSpec struct {
	/* The bucket's autoclass configuration. */
	// +kcc:proto:field=google.storage.v1.Bucket.autoclass
	Autoclass *StorageBucketAutoclass `json:"autoclass,omitempty"`

	/* DEPRECATED. Please use the `uniformBucketLevelAccess` field as this field has been renamed by Google. The `uniformBucketLevelAccess` field will supersede this field.
	Enables Bucket PolicyOnly access to a bucket. */
	BucketPolicyOnly *bool `json:"bucketPolicyOnly,omitempty"`

	/* The bucket's Cross-Origin Resource Sharing (CORS) configuration. */
	// +kcc:proto:field=google.storage.v1.Bucket.cors
	Cors []StorageBucketCors `json:"cors,omitempty"`

	/* The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. */
	CustomPlacementConfig *StorageBucketCustomPlacementConfig `json:"customPlacementConfig,omitempty"`

	/* Whether or not to automatically apply an eventBasedHold to new objects added to the bucket. */
	// +kcc:proto:field=google.storage.v1.Bucket.default_event_based_hold
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty"`

	/* The bucket's encryption configuration. */
	// +kcc:proto:field=google.storage.v1.Bucket.encryption
	Encryption *StorageBucketEncryption `json:"encryption,omitempty"`

	/* The bucket IP filtering configuration. */
	// +optional
	IpFilter *StorageBucketIpFilter `json:"ipFilter,omitempty"`

	/* The bucket's Lifecycle Rules configuration. */
	// +kcc:proto:field=google.storage.v1.Bucket.lifecycle.rule
	LifecycleRule []StorageBucketLifecycleRule `json:"lifecycleRule,omitempty"`

	/* The Google Cloud Storage location. */
	// +kcc:proto:field=google.storage.v1.Bucket.location
	// +kubebuilder:default=US
	Location *string `json:"location,omitempty"`

	/* The bucket's Access & Storage Logs configuration. */
	// +kcc:proto:field=google.storage.v1.Bucket.logging
	Logging *StorageBucketLogging `json:"logging,omitempty"`

	/* Prevents public access to a bucket. */
	// +kcc:proto:field=google.storage.v1.Bucket.iam_configuration.public_access_prevention
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty"`

	/* Enables Requester Pays on a storage bucket. */
	// +kcc:proto:field=google.storage.v1.Bucket.billing.requester_pays
	RequesterPays *bool `json:"requesterPays,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	/* Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. */
	// +kcc:proto:field=google.storage.v1.Bucket.retention_policy
	RetentionPolicy *StorageBucketRetentionPolicy `json:"retentionPolicy,omitempty"`

	/* The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy. */
	SoftDeletePolicy *StorageBucketSoftDeletePolicy `json:"softDeletePolicy,omitempty"`

	/* The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. */
	// +kcc:proto:field=google.storage.v1.Bucket.storage_class
	StorageClass *string `json:"storageClass,omitempty"`

	/* Enables uniform bucket-level access on a bucket. */
	// +kcc:proto:field=google.storage.v1.Bucket.iam_configuration.uniform_bucket_level_access.enabled
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty"`

	/* The bucket's Versioning configuration. */
	// +kcc:proto:field=google.storage.v1.Bucket.versioning
	Versioning *StorageBucketVersioning `json:"versioning,omitempty"`

	/* Configuration if the bucket acts as a website. */
	// +kcc:proto:field=google.storage.v1.Bucket.website
	Website *StorageBucketWebsite `json:"website,omitempty"`
}

type StorageBucketAutoclass struct {
	/* While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern. */
	// +required
	Enabled *bool `json:"enabled"`
}

type StorageBucketCors struct {
	/* The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses. */
	MaxAgeSeconds *int `json:"maxAgeSeconds,omitempty"`

	/* The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method". */
	Method []string `json:"method,omitempty"`

	/* The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin". */
	Origin []string `json:"origin,omitempty"`

	/* The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains. */
	ResponseHeader []string `json:"responseHeader,omitempty"`
}

type StorageBucketCustomPlacementConfig struct {
	/* The list of individual regions that comprise a dual-region bucket. See the docs for a list of acceptable regions. */
	// +required
	DataLocations []string `json:"dataLocations"`
}

type StorageBucketEncryption struct {
	/* A reference to the KMS Crypto Key that will be used to encrypt objects inserted into this bucket. */
	// +required
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef"`
}

type StorageBucketLifecycleRule struct {
	/* The Lifecycle Rule's action configuration. A single block of this type is supported. */
	// +required
	Action *StorageBucketLifecycleRuleAction `json:"action"`

	/* The Lifecycle Rule's condition configuration. */
	// +required
	Condition *StorageBucketLifecycleRuleCondition `json:"condition"`
}

type StorageBucketLifecycleRuleAction struct {
	/* The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. */
	StorageClass *string `json:"storageClass,omitempty"`

	/* The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload. */
	// +required
	Type string `json:"type"`
}

type StorageBucketLifecycleRuleCondition struct {
	/* Minimum age of an object in days to satisfy this condition. */
	Age *int `json:"age,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	CreatedBefore *string `json:"createdBefore,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	CustomTimeBefore *string `json:"customTimeBefore,omitempty"`

	/* Number of days elapsed since the user-specified timestamp set on an object. */
	DaysSinceCustomTime *int `json:"daysSinceCustomTime,omitempty"`

	/* Number of days elapsed since the noncurrent timestamp of an object. This
	condition is relevant only for versioned objects. */
	DaysSinceNoncurrentTime *int `json:"daysSinceNoncurrentTime,omitempty"`

	/* One or more matching name prefixes to satisfy this condition. */
	MatchesPrefix []string `json:"matchesPrefix,omitempty"`

	/* Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY. */
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty"`

	/* One or more matching name suffixes to satisfy this condition. */
	MatchesSuffix []string `json:"matchesSuffix,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty"`

	/* Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition. */
	NumNewerVersions *int `json:"numNewerVersions,omitempty"`

	/* Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY". */
	WithState *string `json:"withState,omitempty"`
}

type StorageBucketLogging struct {
	/* The bucket that will receive log objects. */
	// +required
	LogBucket string `json:"logBucket"`

	/* The object prefix for log objects. If it's not provided, by default Google Cloud Storage sets this to this bucket's name. */
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty"`
}

type StorageBucketRetentionPolicy struct {
	/* If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action. */
	IsLocked *bool `json:"isLocked,omitempty"`

	/* The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds. */
	// +required
	RetentionPeriod int `json:"retentionPeriod"`
}

type StorageBucketSoftDeletePolicy struct {
	/* The duration in seconds that soft-deleted objects in the bucket will be retained and cannot be permanently deleted. Default value is 604800. */
	RetentionDurationSeconds *int `json:"retentionDurationSeconds,omitempty"`
}

type StorageBucketVersioning struct {
	/* While set to true, versioning is fully enabled for this bucket. */
	// +required
	Enabled bool `json:"enabled"`
}

type StorageBucketWebsite struct {
	/* Behaves as the bucket's directory index where missing objects are treated as potential directories. */
	MainPageSuffix *string `json:"mainPageSuffix,omitempty"`

	/* The custom object to return when a requested resource is not found. */
	NotFoundPage *string `json:"notFoundPage,omitempty"`
}

// StorageBucketStatus defines the config connector machine state of StorageBucket
type StorageBucketStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	ObservedState *StorageBucketObservedState `json:"observedState,omitempty"`

	/* The URI of the created resource. */
	SelfLink *string `json:"selfLink,omitempty"`

	/* The base URL of the bucket, in the format gs://<bucket-name>. */
	Url *string `json:"url,omitempty"`
}

// StorageBucketObservedState is the state of the StorageBucket resource as most recently observed in GCP.
type StorageBucketObservedState struct {
	/* The bucket's soft delete policy, which defines the period of time that soft-deleted objects will be retained, and cannot be permanently deleted. If it is not provided, by default Google Cloud Storage sets this to default soft delete policy. */
	SoftDeletePolicy *StorageBucketSoftDeletePolicyObservedState `json:"softDeletePolicy,omitempty"`
}

type StorageBucketSoftDeletePolicyObservedState struct {
	/* Server-determined value that indicates the time from which the policy, or one with a greater retention, was effective. This value is in RFC 3339 format. */
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	/* The duration in seconds that soft-deleted objects in the bucket will be retained and cannot be permanently deleted. Default value is 604800. */
	RetentionDurationSeconds *int64 `json:"retentionDurationSeconds,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragebucket;gcpstoragebuckets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageBucket is the Schema for the StorageBucket API
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

type StorageBucketIpFilter struct {
	/* The mode of the IP filter. Valid values are 'Enabled' and 'Disabled'. */
	Mode string `json:"mode"`

	/* The public network IP address ranges that can access the bucket and its data. */
	// +optional
	PublicNetworkSource *StorageBucketIpFilterPublicNetworkSource `json:"publicNetworkSource,omitempty"`

	/* The list of VPC networks that can access the bucket. */
	// +optional
	VpcNetworkSources []StorageBucketIpFilterVpcNetworkSources `json:"vpcNetworkSources,omitempty"`

	/* Whether to allow cross-org VPCs in the bucket's IP filter configuration. */
	// +optional
	AllowCrossOrgVpcs *bool `json:"allowCrossOrgVpcs,omitempty"`

	/* Whether to allow all service agents to access the bucket regardless of the IP filter configuration. */
	// +optional
	AllowAllServiceAgentAccess *bool `json:"allowAllServiceAgentAccess,omitempty"`
}

type StorageBucketIpFilterPublicNetworkSource struct {
	/* The list of public IPv4, IPv6 cidr ranges that are allowed to access the bucket. */
	AllowedIpCidrRanges []string `json:"allowedIpCidrRanges"`
}

type StorageBucketIpFilterVpcNetworkSources struct {
	/* The list of public or private IPv4 and IPv6 CIDR ranges that can access the bucket. */
	AllowedIpCidrRanges []string `json:"allowedIpCidrRanges"`

	/* The VPC network that can access the bucket. */
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef"`
}
