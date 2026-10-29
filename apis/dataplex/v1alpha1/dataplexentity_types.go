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

// API sources for DataplexEntity, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/metadata.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.lakes.zones.entities

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEntityGVK = GroupVersion.WithKind("DataplexEntity")

// DataplexEntitySpec defines the desired state of DataplexEntity
// +kcc:spec:proto=google.cloud.dataplex.v1.Entity
// +kcc:required-from-proto
type DataplexEntitySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexEntity name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Display name must be shorter than or equal to 256 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User friendly longer description text. Must be shorter than or
	//  equal to 1024 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.description
	Description *string `json:"description,omitempty"`

	// Required. A user-provided entity ID. It is mutable, and will be used as the
	//  published table name. Specifying a new ID in an update entity
	//  request will override the existing value.
	//  The ID must contain only letters (a-z, A-Z), numbers (0-9), and
	//  underscores, and consist of 256 or fewer characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.id
	// +required
	ID *string `json:"id,omitempty"`

	// Optional. The etag associated with the entity, which can be retrieved with
	//  a [GetEntity][] request. Required for update and delete requests.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.etag
	Etag *string `json:"etag,omitempty"`

	// Required. Immutable. The type of entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.type
	// +required
	Type *string `json:"type,omitempty"`

	// Required. Immutable. The ID of the asset associated with the storage
	//  location containing the entity data. The entity must be with in the same
	//  zone with the asset.
	// +kcc:guess=possible-reference target=DataplexAsset
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.asset
	// +required
	Asset *string `json:"asset,omitempty"`

	// Required. Immutable. The storage path of the entity data.
	//  For Cloud Storage data, this is the fully-qualified path to the entity,
	//  such as `gs://bucket/path/to/data`. For BigQuery data, this is the name of
	//  the table resource, such as
	//  `projects/project_id/datasets/dataset_id/tables/table_id`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.data_path
	// +required
	DataPath *string `json:"dataPath,omitempty"`

	// Optional. The set of items within the data path constituting the data in
	//  the entity, represented as a glob path. Example:
	//  `gs://bucket/path/to/data/**/*.csv`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.data_path_pattern
	DataPathPattern *string `json:"dataPathPattern,omitempty"`

	// Required. Immutable. Identifies the storage system of the entity data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.system
	// +required
	System *string `json:"system,omitempty"`

	// Required. Identifies the storage format of the entity data.
	//  It does not apply to entities with data stored in BigQuery.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.format
	// +required
	Format *StorageFormat `json:"format,omitempty"`

	// Required. The description of the data structure and layout.
	//  The schema is not included in list responses. It is only included in
	//  `SCHEMA` and `FULL` entity views of a `GetEntity` response.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.schema
	// +required
	Schema *Schema `json:"schema,omitempty"`
}

// DataplexEntityStatus defines the config connector machine state of DataplexEntity
type DataplexEntityStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEntity resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexEntityObservedState `json:"observedState,omitempty"`
}

// DataplexEntityObservedState is the state of the DataplexEntity resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Entity
type DataplexEntityObservedState struct {
	// Output only. The time when the entity was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the entity was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The name of the associated Data Catalog entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.catalog_entry
	CatalogEntry *string `json:"catalogEntry,omitempty"`

	// Required. Identifies the storage format of the entity data.
	//  It does not apply to entities with data stored in BigQuery.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.format
	Format *StorageFormatObservedState `json:"format,omitempty"`

	// Output only. Metadata stores that the entity is compatible with.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.compatibility
	Compatibility *Entity_CompatibilityStatusObservedState `json:"compatibility,omitempty"`

	// Output only. Identifies the access mechanism to the entity. Not user
	//  settable.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.access
	Access *StorageAccessObservedState `json:"access,omitempty"`

	// Output only. System generated unique ID for the Entity. This ID will be
	//  different if the Entity is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entity.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexentity;gcpdataplexentities
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEntity is the Schema for the DataplexEntity API
// +k8s:openapi-gen=true
type DataplexEntity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEntitySpec   `json:"spec,omitempty"`
	Status DataplexEntityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEntityList contains a list of DataplexEntity
type DataplexEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEntity `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEntity{}, &DataplexEntityList{})
}
