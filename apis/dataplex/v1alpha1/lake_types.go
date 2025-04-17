// Copyright 2025 Google LLC
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
	dataprocv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexLakeGVK = GroupVersion.WithKind("DataplexLake")

type Parent struct {
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	Location string `json:"location"`
}

// DataplexLakeSpec defines the desired state of DataplexLake
// +kcc:spec:proto=google.cloud.dataplex.v1.Lake
type DataplexLakeSpec struct {
	// The DataplexLake name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.description
	Description *string `json:"description,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance
	//  association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore
	Metastore *Lake_Metastore `json:"metastore,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake.Metastore
type Lake_Metastore struct {
	// Optional. A relative reference to the Dataproc Metastore
	//  (https://cloud.google.com/dataproc-metastore/docs) service associated
	//  with the lake:
	//  `projects/{project_id}/locations/{location_id}/services/{service_id}`
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.Metastore.service
	ServiceRef *dataprocv1alpha1.ServiceRef `json:"serviceRef,omitempty"`
}

// DataplexLakeStatus defines the config connector machine state of DataplexLake
type DataplexLakeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexLake resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexLakeObservedState `json:"observedState,omitempty"`
}

// DataplexLakeObservedState is the state of the DataplexLake resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Lake
type DataplexLakeObservedState struct {
	// Output only. System generated globally unique ID for the lake. This ID will
	//  be different if the lake is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the lake was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the lake was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.state
	State *string `json:"state,omitempty"`

	// Output only. Service account associated with this lake. This service
	//  account must be authorized to access or operate on resources managed by the
	//  lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Aggregated status of the underlying assets of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`

	// Output only. Metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore_status
	MetastoreStatus *Lake_MetastoreStatus `json:"metastoreStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexlake;gcpdataplexlakes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexLake is the Schema for the DataplexLake API
// +k8s:openapi-gen=true
type DataplexLake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexLakeSpec   `json:"spec,omitempty"`
	Status DataplexLakeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexLakeList contains a list of DataplexLake
type DataplexLakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexLake `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexLake{}, &DataplexLakeList{})
}
