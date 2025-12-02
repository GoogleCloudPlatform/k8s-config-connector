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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
)

var FirestoreFieldGVK = GroupVersion.WithKind("FirestoreField")

// FirestoreFieldSpec defines the desired state of FirestoreField
// +kcc:spec:proto=google.firestore.admin.v1.Field
type FirestoreFieldSpec struct {
	// The FirestoreDatabase containing the collection group for this field.
	// +required
	DatabaseRef *v1beta1.FirestoreDatabaseRef `json:"databaseRef"`

	// The FirestoreField name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The collectionGroup of which this field is a part.
	// +required
	CollectionGroup *string `json:"collectionGroup,omitempty"`

	// The index configuration for this field. If unset, field indexing will
	//  revert to the configuration defined by the `ancestor_field`. To
	//  explicitly remove all indexes for this field, specify an index config
	//  with an empty list of indexes.
	// +kcc:proto:field=google.firestore.admin.v1.Field.index_config
	IndexConfig *Field_IndexConfig `json:"indexConfig,omitempty"`

	// The TTL configuration for this `Field`.
	TTLConfig *Field_TTLConfig_Spec `json:"ttlConfig,omitempty"`
}

// The Spec version of the TTL configuration for this `Field`.
// Note that this field is unusual - its presence indicates that TTL should be
// enabled, and its absence indicates that TTL should be disabled.
// We handle it specially.
type Field_TTLConfig_Spec struct {
	// Whether to enable TTL for documents based on this field.
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Field.IndexConfig
type Field_IndexConfig struct {
	// The indexes supported for this field.
	// +kcc:proto:field=google.firestore.admin.v1.Field.IndexConfig.indexes
	Indexes []Index `json:"indexes,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Field.IndexConfig
type Field_IndexConfig_ObservedState struct {
	// The indexes supported for this field.
	// +kcc:proto:field=google.firestore.admin.v1.Field.IndexConfig.indexes
	Indexes []Index_ObservedState `json:"indexes,omitempty"`

	// Output only. When true, the `Field`'s index configuration is set from the
	//  configuration specified by the `ancestor_field`.
	//  When false, the `Field`'s index configuration is defined explicitly.
	// +kcc:proto:field=google.firestore.admin.v1.Field.IndexConfig.uses_ancestor_config
	UsesAncestorConfig *bool `json:"usesAncestorConfig,omitempty"`

	// Output only. Specifies the resource name of the `Field` from which this
	//  field's index configuration is set (when `uses_ancestor_config` is true),
	//  or from which it *would* be set if this field had no index configuration
	//  (when `uses_ancestor_config` is false).
	// +kcc:proto:field=google.firestore.admin.v1.Field.IndexConfig.ancestor_field
	AncestorField *string `json:"ancestorField,omitempty"`

	// Output only
	//  When true, the `Field`'s index configuration is in the process of being
	//  reverted. Once complete, the index config will transition to the same
	//  state as the field specified by `ancestor_field`, at which point
	//  `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	// +kcc:proto:field=google.firestore.admin.v1.Field.IndexConfig.reverting
	Reverting *bool `json:"reverting,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index
type Index struct {
	// Indexes with a collection query scope specified allow queries
	//  against a collection that is the child of a specific document, specified at
	//  query time, and that has the same collection ID.
	//
	//  Indexes with a collection group query scope specified allow queries against
	//  all collections descended from a specific document, specified at query
	//  time, and that have the same collection ID as this index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.query_scope
	QueryScope *string `json:"queryScope,omitempty"`

	// The API scope supported by this index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.api_scope
	APIScope *string `json:"apiScope,omitempty"`

	// The fields supported by this index.
	//
	//  For composite indexes, this requires a minimum of 2 and a maximum of 100
	//  fields. The last field entry is always for the field path `__name__`. If,
	//  on creation, `__name__` was not specified as the last field, it will be
	//  added automatically with the same direction as that of the last field
	//  defined. If the final field in a composite index is not directional, the
	//  `__name__` will be ordered ASCENDING (unless explicitly specified).
	//
	//  For single field indexes, this will always be exactly one entry with a
	//  field path equal to the field path of the associated field.
	// +kcc:proto:field=google.firestore.admin.v1.Index.fields
	Fields []Index_IndexField `json:"fields,omitempty"`

	// Immutable. The density configuration of the index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.density
	Density *string `json:"density,omitempty"`

	// Optional. Whether the index is multikey. By default, the index is not
	//  multikey. For non-multikey indexes, none of the paths in the index
	//  definition reach or traverse an array, except via an explicit array index.
	//  For multikey indexes, at most one of the paths in the index definition
	//  reach or traverse an array, except via an explicit array index. Violations
	//  will result in errors.
	//
	//  Note this field only applies to index with MONGODB_COMPATIBLE_API ApiScope.
	// +kcc:proto:field=google.firestore.admin.v1.Index.multikey
	Multikey *bool `json:"multikey,omitempty"`

	// Optional. The number of shards for the index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.shard_count
	ShardCount *int32 `json:"shardCount,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index
type Index_ObservedState struct {
	// Output only. A server defined name for this index.
	//  The form of this name for composite indexes will be:
	//  `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}`
	//  For single field indexes, this field will be empty.
	// +kcc:proto:field=google.firestore.admin.v1.Index.name
	Name *string `json:"name,omitempty"`

	// Output only. The serving state of the index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.state
	State *string `json:"state,omitempty"`
}

// FirestoreFieldStatus defines the config connector machine state of FirestoreField
type FirestoreFieldStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FirestoreField resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FirestoreFieldObservedState `json:"observedState,omitempty"`
}

// FirestoreFieldObservedState is the state of the FirestoreField resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.firestore.admin.v1.Field
type FirestoreFieldObservedState struct {
	// The index state for this field.
	// +kcc:proto:field=google.firestore.admin.v1.Field.index_config
	IndexConfig *Field_IndexConfig_ObservedState `json:"indexConfig,omitempty"`

	// The TTL state for this `Field`.
	// +kcc:proto:field=google.firestore.admin.v1.Field.ttl_config
	TTLConfig *Field_TTLConfigObservedState `json:"ttlConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfirestorefield;gcpfirestorefields
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FirestoreField is the Schema for the FirestoreField API
// +k8s:openapi-gen=true
type FirestoreField struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FirestoreFieldSpec   `json:"spec,omitempty"`
	Status FirestoreFieldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FirestoreFieldList contains a list of FirestoreField
type FirestoreFieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirestoreField `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirestoreField{}, &FirestoreFieldList{})
}
