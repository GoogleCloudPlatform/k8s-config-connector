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

// API sources for DataplexDataAsset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/data_products.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.dataProducts.dataAssets

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexDataAssetGVK = GroupVersion.WithKind("DataplexDataAsset")

// DataplexDataAssetSpec defines the desired state of DataplexDataAsset
// +kcc:spec:proto=google.cloud.dataplex.v1.DataAsset
// +kcc:required-from-proto
type DataplexDataAssetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexDataAsset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. User-defined labels for the data asset.
	//
	//  Example:
	//   ```
	//   {
	//     "environment": "production",
	//     "billing": "marketing-department"
	//   }
	//   ```
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. Full resource name of the cloud resource represented
	//  by the data asset. This must follow
	//  https://cloud.google.com/iam/docs/full-resource-names. Example:
	//  `//bigquery.googleapis.com/projects/my_project_123/datasets/dataset_456/tables/table_789`
	//  Only BigQuery tables and datasets are currently supported.
	//  Data asset creator must have getIamPolicy and setIamPolicy permissions on
	//  the resource. Data asset creator must also have resource specific get
	//  permission, for instance, bigquery.tables.get for BigQuery tables.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.resource
	// +required
	Resource *string `json:"resource,omitempty"`

	// Optional. Access groups configurations for this data asset.
	//
	//  The key is `DataProduct.AccessGroup.id` and the value is
	//  `AccessGroupConfig`.
	//
	//  Example:
	//  ```
	//   {
	//     "analyst": {
	//       "iamRoles": ["roles/bigquery.dataViewer"]
	//     }
	//   }
	//  ```
	//
	//  Currently, at most one IAM role is allowed per access group. For providing
	//  multiple predefined IAM roles, wrap them in a custom IAM role as per
	//  https://cloud.google.com/iam/docs/creating-custom-roles.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.access_group_configs
	AccessGroupConfigs map[string]DataAsset_AccessGroupConfig `json:"accessGroupConfigs,omitempty"`
}

// DataplexDataAssetStatus defines the config connector machine state of DataplexDataAsset
type DataplexDataAssetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexDataAsset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexDataAssetObservedState `json:"observedState,omitempty"`
}

// DataplexDataAssetObservedState is the state of the DataplexDataAsset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataAsset
type DataplexDataAssetObservedState struct {
	// Output only. System generated globally unique ID for the data asset.
	//  This ID will be different if the data asset is deleted and re-created
	//  with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which the data asset was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the data asset was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAsset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexdataasset;gcpdataplexdataassets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexDataAsset is the Schema for the DataplexDataAsset API
// +k8s:openapi-gen=true
type DataplexDataAsset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexDataAssetSpec   `json:"spec,omitempty"`
	Status DataplexDataAssetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexDataAssetList contains a list of DataplexDataAsset
type DataplexDataAssetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexDataAsset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexDataAsset{}, &DataplexDataAssetList{})
}
