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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryDatasetGVK = GroupVersion.WithKind("BigQueryDataset")

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BigQueryDatasetSpec defines the desired state of BigQueryDataset
// +kcc:spec:proto=google.cloud.bigquery.v2.Dataset
type BigQueryDatasetSpec struct {
	// The BigQueryDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An array of objects that define dataset access for one or more entities.
	// +optional
	Access []Access `json:"access,omitempty"`

	// Optional. Defines the default collation specification of future tables
	//  created in the dataset. If a table is created in this dataset without
	//  table-level default collation, then the table inherits the dataset default
	//  collation, which is applied to the string fields that do not have explicit
	//  collation specified. A change to this field affects only tables created
	//  afterwards, and does not alter the existing tables.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case-insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	DefaultCollation *string `json:"defaultCollation,omitempty"`

	// The default encryption key for all tables in the dataset.
	//  After this property is set, the encryption key of all newly-created tables
	//  in the dataset is set to this value unless the table creation request or
	//  query explicitly overrides the key.
	DefaultEncryptionConfiguration *EncryptionConfiguration `json:"defaultEncryptionConfiguration,omitempty"`

	// This default partition expiration, expressed in milliseconds.
	//
	//  When new time-partitioned tables are created in a dataset where this
	//  property is set, the table will inherit this value, propagated as the
	//  `TimePartitioning.expirationMs` property on the new table.  If you set
	//  `TimePartitioning.expirationMs` explicitly when creating a table,
	//  the `defaultPartitionExpirationMs` of the containing dataset is ignored.
	//
	//  When creating a partitioned table, if `defaultPartitionExpirationMs`
	//  is set, the `defaultTableExpirationMs` value is ignored and the table
	//  will not be inherit a table expiration deadline.
	DefaultPartitionExpirationMs *int64 `json:"defaultPartitionExpirationMs,omitempty"`

	// Optional. The default lifetime of all tables in the dataset, in
	//  milliseconds. The minimum lifetime value is 3600000 milliseconds (one
	//  hour). To clear an existing default expiration with a PATCH request, set to
	//  0. Once this property is set, all newly-created tables in the dataset will
	//  have an expirationTime property set to the creation time plus the value in
	//  this property, and changing the value will only affect new tables, not
	//  existing ones. When the expirationTime for a given table is reached, that
	//  table will be deleted automatically.
	//  If a table's expirationTime is modified or removed before the table
	//  expires, or if you provide an explicit expirationTime when creating a
	//  table, that value takes precedence over the default expiration time
	//  indicated by this property.
	DefaultTableExpirationMs *int64 `json:"defaultTableExpirationMs,omitempty"`

	// Optional. A user-friendly description of the dataset.
	Description *string `json:"description,omitempty"`

	// Optional. A descriptive name for the dataset.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Optional. TRUE if the dataset and its table names are case-insensitive,
	//  otherwise FALSE. By default, this is FALSE, which means the dataset and its
	//  table names are case-sensitive. This field does not affect routine
	//  references.
	IsCaseInsensitive *bool `json:"isCaseInsensitive,omitempty"`

	// Optional. The geographic location where the dataset should reside. See
	//  https://cloud.google.com/bigquery/docs/locations for supported
	//  locations.
	Location *string `json:"location,omitempty"`

	// Optional. Defines the time travel window in hours. The value can be from 48
	//  to 168 hours (2 to 7 days). The default value is 168 hours if this is not
	//  set.
	MaxTimeTravelHours *string `json:"maxTimeTravelHours,omitempty"`

	//  Optional. The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Optional. Updates storage_billing_model for the dataset.
	StorageBillingModel *string `json:"storageBillingModel,omitempty"`
}

// BigQueryDatasetStatus defines the config connector machine state of BigQueryDataset
type BigQueryDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Output only. The time when this dataset was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. A hash of the resource.
	Etag *string `json:"etag,omitempty"`

	// A unique specifier for the BigQueryAnalyticsHubDataExchangeListing resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Output only. The date when this dataset was last modified, in milliseconds
	//  since the epoch.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. A URL that can be used to access the resource again. You can
	//  use this URL in Get or Update requests to the resource.
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryDatasetObservedState `json:"observedState,omitempty"`
}

// BigQueryDatasetObservedState defines the desired state of BigQueryDataset
// +kcc:observedstate:proto=google.cloud.bigquery.v2.dataset
type BigQueryDatasetObservedState struct {

	// Optional. If the location is not specified in the spec, the GCP server defaults to a location and will be captured here.
	Location *string `json:"location,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerydataset;gcpbigquerydatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryDataset is the Schema for the BigQueryDataset API
// +k8s:openapi-gen=true
type BigQueryDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDatasetSpec   `json:"spec,omitempty"`
	Status BigQueryDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryDatasetList contains a list of BigQueryDataset
type BigQueryDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryDataset{}, &BigQueryDatasetList{})
}
