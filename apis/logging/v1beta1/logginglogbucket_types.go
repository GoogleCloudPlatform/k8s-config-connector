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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLogBucketSpec defines the desired state of LoggingLogBucket
// +kcc:spec:proto=google.logging.v2.LogBucket
type LoggingLogBucketSpec struct {
	/* Immutable. The BillingAccount that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	BillingAccountRef *LoggingLogBucketBillingAccountRef `json:"billingAccountRef,omitempty"`

	/* Immutable. The Folder that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	FolderRef *LoggingLogBucketFolderRef `json:"folderRef,omitempty"`

	/* Immutable. The Organization that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	OrganizationRef *LoggingLogBucketOrganizationRef `json:"organizationRef,omitempty"`

	/* Immutable. The Project that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	ProjectRef *LoggingLogBucketProjectRef `json:"projectRef,omitempty"`

	/* Immutable. The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1. */
	Location string `json:"location"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	/* Describes this bucket. */
	// +kcc:proto:field=google.logging.v2.LogBucket.description
	Description *string `json:"description,omitempty"`

	/* Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the Log Analytics page using SQL queries. Cannot be disabled once enabled. */
	// +kcc:proto:field=google.logging.v2.LogBucket.analytics_enabled
	AnalyticsEnabled *bool `json:"enableAnalytics,omitempty"`

	/* Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty. */
	// +kcc:proto:field=google.logging.v2.LogBucket.locked
	Locked *bool `json:"locked,omitempty"`

	/* Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. */
	// +kcc:proto:field=google.logging.v2.LogBucket.retention_days
	RetentionDays *int32 `json:"retentionDays,omitempty"`
}

type LoggingLogBucketBillingAccountRef struct {
	/* Allowed value: The Google Cloud resource name of a Google Cloud Billing Account (format: `billingAccounts/{{name}}`). */
	External string `json:"external,omitempty"`

	/* [WARNING] BillingAccount not yet supported in Config Connector, use 'external' field to reference existing resources.
	   Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type LoggingLogBucketFolderRef struct {
	/* Allowed value: The Google Cloud resource name of a `Folder` resource (format: `folders/{{name}}`). */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type LoggingLogBucketOrganizationRef struct {
	/* Allowed value: The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`). */
	External string `json:"external,omitempty"`

	/* [WARNING] Organization not yet supported in Config Connector, use 'external' field to reference existing resources.
	   Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type LoggingLogBucketProjectRef struct {
	/* Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`). */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// LoggingLogBucketStatus defines the config connector machine state of LoggingLogBucket
type LoggingLogBucketStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Output only. The creation timestamp of the bucket. This is not set for any of the default buckets. */
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.logging.v2.LogBucket.create_time
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The bucket lifecycle state. Possible values: LIFECYCLE_STATE_UNSPECIFIED, ACTIVE, DELETE_REQUESTED */
	// +kcc:proto:field=google.logging.v2.LogBucket.lifecycle_state
	LifecycleState *string `json:"lifecycleState,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The last update timestamp of the bucket. */
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.logging.v2.LogBucket.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplogginglogbucket;gcplogginglogbuckets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LoggingLogBucket is the Schema for the LoggingLogBucket API
// +k8s:openapi-gen=true
type LoggingLogBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LoggingLogBucketSpec   `json:"spec,omitempty"`
	Status LoggingLogBucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LoggingLogBucketList contains a list of LoggingLogBucket
type LoggingLogBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingLogBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingLogBucket{}, &LoggingLogBucketList{})
}
