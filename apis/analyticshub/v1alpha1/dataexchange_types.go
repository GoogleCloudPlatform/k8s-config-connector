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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryAnalyticsHubDataExchangeGVK = GroupVersion.WithKind("BigQueryAnalyticsHubDataExchange")

// BigQueryAnalyticsHubDataExchangeSpec defines the desired state of BigQueryAnalyticsHubDataExchange
// +kcc:spec:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type BigQueryAnalyticsHubDataExchangeSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The BigQueryAnalyticsHubDataExchange name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Human-readable display name of the data exchange. The display
	//  name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), ampersands (&) and must not start or end with
	//  spaces. Default value is an empty string. Max length: 63 bytes.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// Optional. Description of the data exchange. The description must not
	//  contain Unicode non-characters as well as C0 and C1 control codes except
	//  tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	//  Default value is an empty string.
	//  Max length: 2000 bytes.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. Email or URL of the primary point of contact of the data
	//  exchange. Max Length: 1000 bytes.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the data exchange.
	// +kubebuilder:validation:Optional
	Documentation *string `json:"documentation,omitempty"`

	// Optional. Base64 encoded image representing the data exchange. Max
	//  Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API
	//  only performs validation on size of the encoded data. Note: For byte
	//  fields, the content of the fields are base64-encoded (which increases the
	//  size of the data by 33-36%) when using JSON on the wire.
	// +kubebuilder:validation:Optional
	Icon []byte `json:"icon,omitempty"`

	// Optional. Configurable data sharing environment option for a data exchange.
	// +kubebuilder:validation:Optional
	// SharingEnvironmentConfig *SharingEnvironmentConfig `json:"sharingEnvironmentConfig,omitempty"`

	// Optional. Type of discovery on the discovery page for all the listings
	//  under this exchange. Updating this field also updates (overwrites) the
	//  discovery_type field for all the listings under this exchange.
	// +kubebuilder:validation:Optional
	DiscoveryType *string `json:"discoveryType,omitempty"`

	// Optional. By default, false.
	//  If true, the DataExchange has an email sharing mandate enabled.
	// +kubebuilder:validation:Optional
	LogLinkedDatasetQueryUserEmail *bool `json:"logLinkedDatasetQueryUserEmail,omitempty"`
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

// BigQueryAnalyticsHubDataExchangeObservedState is the state of the BigQueryAnalyticsHubDataExchange resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.analyticshub.v1.DataExchange
type BigQueryAnalyticsHubDataExchangeObservedState struct {
	// Output only. Number of listings contained in the data exchange.
	ListingCount *int32 `json:"listingCount,omitempty"`

	// Optional. Configurable data sharing environment option for a data exchange.
	// SharingEnvironmentConfig *SharingEnvironmentConfigObservedState `json:"sharingEnvironmentConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryanalyticshubdataexchange;gcpbigqueryanalyticshubdataexchanges
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

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
