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
	bigquery "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryAnalyticsHubListingGVK = GroupVersion.WithKind("BigQueryAnalyticsHubListing")

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.SelectedResource
type SelectedResource struct {
	// Optional. A reference to a BigQueryTable.
	// Format:
	//  `projects/{projectId}/datasets/{datasetId}/tables/{tableId}`
	//  Example:"projects/test_project/datasets/test_dataset/tables/test_table"
	TableRef *refv1beta1.BigQueryTableRef `json:"tableRef,omitempty"`
}

type BigQueryDatasetSource struct {
	// +required
	// Resource name of the dataset source for this listing.
	//  e.g. `projects/myproject/datasets/123`
	DatasetRef *bigquery.DatasetRef `json:"datasetRef,omitempty"`

	// Optional. Resources in this dataset that are selectively shared.
	//  If this field is empty, then the entire dataset (all resources) are
	//  shared. This field is only valid for data clean room exchanges.
	SelectedResources []SelectedResource `json:"selectedResources,omitempty"`

	// Optional. If set, restricted export policy will be propagated and
	//  enforced on the linked dataset.
	RestrictedExportPolicy *RestrictedExportPolicy `json:"restrictedExportPolicy,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.analyticshub.v1.Listing.BigQueryDatasetSource.RestrictedExportPolicy
type RestrictedExportPolicy struct {
	// Optional. If true, enable restricted export.
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. If true, restrict direct table access (read
	//  api/tabledata.list) on linked table.
	RestrictDirectTableAccess *bool `json:"restrictDirectTableAccess,omitempty"`

	// Optional. If true, restrict export of query result derived from
	//  restricted linked dataset table.
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty"`
}

type Source struct {
	// One of the following fields must be set.
	BigQueryDatasetSource *BigQueryDatasetSource `json:"bigQueryDatasetSource,omitempty"`

	// NOT YET
	// PubsubTopicSource *PubsubTopicSource `json:"pubsubTopicSource,omitempty"`
}

// BigQueryAnalyticsHubListingSpec defines the desired state of BigQueryAnalyticsHubDataExchangeListing
// +kcc:spec:proto=google.cloud.bigquery.analyticshub.v1.Listing
type BigQueryAnalyticsHubListingSpec struct {
	// +required
	Source *Source `json:"source,omitempty"`

	// +required
	// Required. Human-readable display name of the listing. The display name must
	//  contain only Unicode letters, numbers (0-9), underscores (_), dashes (-),
	//  spaces ( ), ampersands (&) and can't start or end with spaces. Default
	//  value is an empty string. Max length: 63 bytes.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Short description of the listing. The description must contain only
	//  Unicode characters or tabs  (HT), new lines (LF), carriage returns (CR), and
	//  page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the listing.
	//  Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the listing.
	Documentation *string `json:"documentation,omitempty"`

	// Optional. Details of the data provider who owns the source data.
	DataProvider *DataProvider `json:"dataProvider,omitempty"`

	// Optional. Categories of the listing. Up to two categories are allowed.
	Categories []string `json:"categories,omitempty"`

	// Optional. Details of the publisher who owns the listing and who can share
	//  the source data.
	Publisher *Publisher `json:"publisher,omitempty"`

	// Optional. Email or URL of the request access of the listing.
	//  Subscribers can use this reference to request access.
	//  Max Length: 1000 bytes.
	RequestAccess *string `json:"requestAccess,omitempty"`

	// Not yet
	// // Optional. If set, restricted export configuration will be propagated and
	// //  enforced on the linked dataset.
	// RestrictedExportConfig *Listing_RestrictedExportConfig `json:"restrictedExportConfig,omitempty"`

	// Optional. Type of discovery of the listing on the discovery page.
	DiscoveryType *string `json:"discoveryType,omitempty"`

	// Not yet
	// // Optional. Base64 encoded image representing the listing. Max Size: 3.0MiB
	// //  Expected image dimensions are 512x512 pixels, however the API only
	// //  performs validation on size of the encoded data.
	// //  Note: For byte fields, the contents of the field are base64-encoded (which
	// //  increases the size of the data by 33-36%) when using JSON on the wire.
	// Icon []byte `json:"icon,omitempty"`

	/* Immutable. The name of the location this data exchange. */
	// +required
	Location string `json:"location"`

	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +required
	DataExchangeRef *BigQueryAnalyticsHubDataExchangeRef `json:"dataExchangeRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The BigQueryAnalyticsHubDataExchangeListing name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// BigQueryAnalyticsHubListingStatus defines the config connector machine state of BigQueryAnalyticsHubDataExchangeListing
type BigQueryAnalyticsHubListingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryAnalyticsHubDataExchangeListing resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryAnalyticsHubListingObservedState `json:"observedState,omitempty"`
}

// BigQueryAnalyticsHubDataExchangeListingSpec defines the desired state of BigQueryAnalyticsHubDataExchangeListing
// +kcc:observedstate:proto=google.cloud.bigquery.analyticshub.v1.Listing
type BigQueryAnalyticsHubListingObservedState struct {
	// This field is in the same format as our externalRef! So it's redundant.
	// // Output only. The resource name of the data exchange.
	// //  e.g. `projects/myproject/locations/US/dataExchanges/123/listing/456`.
	// Name *string `json:"name,omitempty"`

	// Output only. Current state of the listing.
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryanalyticshublisting;gcpbigqueryanalyticshublistings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryAnalyticsHubListing is the Schema for the BigQueryAnalyticsHubListing API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type BigQueryAnalyticsHubListing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryAnalyticsHubListingSpec   `json:"spec,omitempty"`
	Status BigQueryAnalyticsHubListingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryAnalyticsHubListingList contains a list of BigQueryAnalyticsHubDataExchangeListing
type BigQueryAnalyticsHubListingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryAnalyticsHubListing `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryAnalyticsHubListing{}, &BigQueryAnalyticsHubListingList{})
}
