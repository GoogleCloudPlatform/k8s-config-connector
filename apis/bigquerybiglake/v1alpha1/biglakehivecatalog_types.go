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

// API sources for BigLakeHiveCatalog, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/biglake/hive/v1/hive_metastore.proto
// +kcc:source:service-docs=https://cloud.google.com/products/lakehouse
// +kcc:source:resource-docs=https://cloud.google.com/bigquery/docs/manage-open-source-metadata

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeHiveCatalogGVK = GroupVersion.WithKind("BigLakeHiveCatalog")

// BigLakeHiveCatalogSpec defines the desired state of BigLakeHiveCatalog
// +kcc:spec:proto=google.cloud.biglake.hive.v1.HiveCatalog
// +kcc:required-from-proto
type BigLakeHiveCatalogSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The BigLakeHiveCatalog name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Stores the catalog description.
	//  The maximum length is 4000 characters.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveCatalog.description
	Description *string `json:"description,omitempty"`

	// Required. The Cloud Storage location path where the catalog exists.
	//  Format: gs://bucket/path/to/catalog
	//  The maximum length is 4000 characters.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveCatalog.location_uri
	// +required
	LocationURI *string `json:"locationURI,omitempty"`
}

// BigLakeHiveCatalogStatus defines the config connector machine state of BigLakeHiveCatalog
type BigLakeHiveCatalogStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeHiveCatalog resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeHiveCatalogObservedState `json:"observedState,omitempty"`
}

// BigLakeHiveCatalogObservedState is the state of the BigLakeHiveCatalog resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.biglake.hive.v1.HiveCatalog
type BigLakeHiveCatalogObservedState struct {
	// Output only. The replicas for the catalog metadata.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveCatalog.replicas
	Replicas []HiveCatalog_ReplicaObservedState `json:"replicas,omitempty"`

	// Output only. The creation time of the catalog.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveCatalog.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the catalog.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveCatalog.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakehivecatalog;gcpbiglakehivecatalogs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeHiveCatalog is the Schema for the BigLakeHiveCatalog API
// +k8s:openapi-gen=true
type BigLakeHiveCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeHiveCatalogSpec   `json:"spec,omitempty"`
	Status BigLakeHiveCatalogStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeHiveCatalogList contains a list of BigLakeHiveCatalog
type BigLakeHiveCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeHiveCatalog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeHiveCatalog{}, &BigLakeHiveCatalogList{})
}
