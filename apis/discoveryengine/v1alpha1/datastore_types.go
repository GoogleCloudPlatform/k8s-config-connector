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

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineDataStoreGVK = GroupVersion.WithKind("DiscoveryEngineDataStore")

// DiscoveryEngineDataStoreSpec defines the desired state of DiscoveryEngineDataStore
// +kcc:spec:proto=google.cloud.discoveryengine.v1.DataStore
type DiscoveryEngineDataStoreSpec struct {
	// The DiscoveryEngineDataStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The data store display name.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The industry vertical that the data store registers.
	IndustryVertical *string `json:"industryVertical,omitempty"`

	// The solutions that the data store enrolls. Available solutions for each
	// [industry_vertical][google.cloud.discoveryengine.v1.DataStore.industry_vertical]:
	//
	// * `MEDIA`: `SOLUTION_TYPE_RECOMMENDATION` and `SOLUTION_TYPE_SEARCH`.
	// * `SITE_SEARCH`: `SOLUTION_TYPE_SEARCH` is automatically enrolled. Other
	//   solutions cannot be enrolled.
	SolutionTypes []string `json:"solutionTypes,omitempty"`

	// Immutable. The content config of the data store. If this field is unset,
	// the server behavior defaults to
	// [ContentConfig.NO_CONTENT][google.cloud.discoveryengine.v1.DataStore.ContentConfig.NO_CONTENT].
	ContentConfig *string `json:"contentConfig,omitempty"`

	// Config to store data store type configuration for workspace data. This
	// must be set when
	// [DataStore.content_config][google.cloud.discoveryengine.v1.DataStore.content_config]
	// is set as
	// [DataStore.ContentConfig.GOOGLE_WORKSPACE][google.cloud.discoveryengine.v1.DataStore.ContentConfig.GOOGLE_WORKSPACE].
	WorkspaceConfig *WorkspaceConfig `json:"workspaceConfig,omitempty"`

	/* NOTYET: this includes a map[string]object which is difficult to map to KRM
	// Configuration for Document understanding and enrichment.
	DocumentProcessingConfig *DocumentProcessingConfig `json:"documentProcessingConfig,omitempty"`
	*/

	/* The ID of the project in which the resource belongs.*/
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The collection for the DataStore.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Collection field is immutable"
	// +required
	Collection string `json:"collection"`

	// Immutable. The location for the resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// +required
	Location string `json:"location"`
}

// DiscoveryEngineDataStoreStatus defines the config connector machine state of DiscoveryEngineDataStore
type DiscoveryEngineDataStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineDataStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineDataStoreObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineDataStoreObservedState is the state of the DiscoveryEngineDataStore resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.DataStore
type DiscoveryEngineDataStoreObservedState struct {
	// Output only. The id of the default
	// [Schema][google.cloud.discoveryengine.v1.Schema] associated to this data
	// store.
	DefaultSchemaID *string `json:"defaultSchemaID,omitempty"`

	// Output only. Timestamp the
	// [DataStore][google.cloud.discoveryengine.v1.DataStore] was created at.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Data size estimation for billing.
	BillingEstimation *DataStore_BillingEstimation `json:"billingEstimation,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginedatastore;gcpdiscoveryenginedatastores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineDataStore is the Schema for the DiscoveryEngineDataStore API
// +k8s:openapi-gen=true
type DiscoveryEngineDataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineDataStoreSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineDataStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineDataStoreList contains a list of DiscoveryEngineDataStore
type DiscoveryEngineDataStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineDataStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineDataStore{}, &DiscoveryEngineDataStoreList{})
}
