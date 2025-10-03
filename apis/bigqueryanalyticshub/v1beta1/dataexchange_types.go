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

var BigQueryAnalyticsHubDataExchangeGVK = GroupVersion.WithKind("BigQueryAnalyticsHubDataExchange")

// BigQueryAnalyticsHubDataExchangeSpec defines the desired state of BigQueryAnalyticsHubDataExchange
// +kcc:spec:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type BigQueryAnalyticsHubDataExchangeSpec struct {
	// Required. Human-readable display name of the data exchange. The display
	//  name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), ampersands (&) and must not start or end with
	//  spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the data exchange. The description must not
	//  contain Unicode non-characters as well as C0 and C1 control codes except
	//  tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	//  Default value is an empty string.
	//  Max length: 2000 bytes.
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the data
	//  exchange. Max Length: 1000 bytes.
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the data exchange.
	Documentation *string `json:"documentation,omitempty"`

	// TODO(KCC): NOT YET
	// // Optional. Base64 encoded image representing the data exchange. Max
	// //  Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API
	// //  only performs validation on size of the encoded data. Note: For byte
	// //  fields, the content of the fields are base64-encoded (which increases the
	// //  size of the data by 33-36%) when using JSON on the wire.
	// Icon *string `json:"icon,omitempty"`

	// As of now SharingEnvironmentConfig is empty or output only so let's not include it as
	// part of the spec yet.
	// // Optional. Configurable data sharing environment option for a data exchange.
	// SharingEnvironmentConfig *SharingEnvironmentConfig `json:"sharingEnvironmentConfig,omitempty"`

	// Optional. Type of discovery on the discovery page for all the listings
	//  under this exchange. Updating this field also updates (overwrites) the
	//  discovery_type field for all the listings under this exchange.
	DiscoveryType *string `json:"discoveryType,omitempty"`

	/* Immutable. The name of the location this data exchange. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The BigQueryAnalyticsHubDataExchange name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// BigQueryAnalyticsHubDataExchangeStatus defines the config connector machine state of BigQueryAnalyticsHubDataExchange
type BigQueryAnalyticsHubDataExchangeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryAnalyticsHubDataExchange resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryAnalyticsHubDataExchangeObservedState `json:"observedState,omitempty"`
}

// BigQueryAnalyticsHubDataExchangeSpec defines the desired state of BigQueryAnalyticsHubDataExchange
// +kcc:observedstate:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type BigQueryAnalyticsHubDataExchangeObservedState struct {
	// This field is in the same format as our externalRef! So it's redundant.
	// // Output only. The resource name of the data exchange.
	// //  e.g. `projects/myproject/locations/US/dataExchanges/123`.
	// Name *string `json:"name,omitempty"`

	/* Number of listings contained in the data exchange. */
	// +optional
	ListingCount *int64 `json:"listingCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryanalyticshubdataexchange;gcpbigqueryanalyticshubdataexchanges
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion
// BigQueryAnalyticsHubDataExchange is the Schema for the BigQueryAnalyticsHubDataExchange API
// +k8s:openapi-gen=true
type BigQueryAnalyticsHubDataExchange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryAnalyticsHubDataExchangeSpec   `json:"spec,omitempty"`
	Status BigQueryAnalyticsHubDataExchangeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryAnalyticsHubDataExchangeList contains a list of BigQueryAnalyticsHubDataExchange
type BigQueryAnalyticsHubDataExchangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryAnalyticsHubDataExchange `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryAnalyticsHubDataExchange{}, &BigQueryAnalyticsHubDataExchangeList{})
}
