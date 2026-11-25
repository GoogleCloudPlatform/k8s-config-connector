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

// API sources for BigLakeHiveDatabase, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/biglake/hive/v1/hive_metastore.proto
// +kcc:source:service-docs=https://cloud.google.com/products/lakehouse
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/products/lakehouse/docs/reference/rest/v1/projects.catalogs.databases

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeHiveDatabaseGVK = GroupVersion.WithKind("BigLakeHiveDatabase")

// BigLakeHiveDatabaseSpec defines the desired state of BigLakeHiveDatabase
// +kcc:spec:proto=google.cloud.biglake.hive.v1.HiveDatabase
// +kcc:required-from-proto
type BigLakeHiveDatabaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// A reference to the projects/{project}/catalogs/{catalog} this resource belongs to.
	// +kcc:guess
	// CatalogRef *BigQueryBigLakeCatalogRef `json:"catalogRef,omitempty"`

	// The BigLakeHiveDatabase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Stores the database description.
	//  The maximum length is 4000 characters.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveDatabase.description
	Description *string `json:"description,omitempty"`

	// Optional. The Cloud Storage location path where the database exists.
	//  Format: `gs://bucket/path/to/database`
	//  If unspecified, the database will be stored in the catalog location.
	//  The maximum length is 4000 characters.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveDatabase.location_uri
	LocationURI *string `json:"locationURI,omitempty"`

	// Optional. Stores the properties associated with the database.
	//  The maximum size is 2 MiB.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveDatabase.parameters
	Parameters map[string]string `json:"parameters,omitempty"`
}

// BigLakeHiveDatabaseStatus defines the config connector machine state of BigLakeHiveDatabase
type BigLakeHiveDatabaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeHiveDatabase resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeHiveDatabaseObservedState `json:"observedState,omitempty"`
}

// BigLakeHiveDatabaseObservedState is the state of the BigLakeHiveDatabase resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.biglake.hive.v1.HiveDatabase
type BigLakeHiveDatabaseObservedState struct {
	// Output only. The creation time of the database.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveDatabase.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the database.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveDatabase.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakehivedatabase;gcpbiglakehivedatabases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeHiveDatabase is the Schema for the BigLakeHiveDatabase API
// +k8s:openapi-gen=true
type BigLakeHiveDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeHiveDatabaseSpec   `json:"spec,omitempty"`
	Status BigLakeHiveDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeHiveDatabaseList contains a list of BigLakeHiveDatabase
type BigLakeHiveDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeHiveDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeHiveDatabase{}, &BigLakeHiveDatabaseList{})
}
