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

// API sources for DataMigrationMappingRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/clouddms/v1/conversionworkspace_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/database-migration/
// +kcc:source:resource-docs=https://docs.cloud.google.com/database-migration/docs/reference/rest/v1/projects.locations.conversionWorkspaces.mappingRules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataMigrationMappingRuleGVK = GroupVersion.WithKind("DataMigrationMappingRule")

// DataMigrationMappingRuleSpec defines the desired state of DataMigrationMappingRule
// +kcc:spec:proto=google.cloud.clouddms.v1.MappingRule
// +kcc:required-from-proto
type DataMigrationMappingRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataMigrationMappingRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A human readable name
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The mapping rule state
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.state
	State *string `json:"state,omitempty"`

	// Required. The rule scope
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.rule_scope
	// +required
	RuleScope *string `json:"ruleScope,omitempty"`

	// Required. The rule filter
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.filter
	// +required
	Filter *MappingRuleFilter `json:"filter,omitempty"`

	// Required. The order in which the rule is applied. Lower order rules are
	//  applied before higher value rules so they may end up being overridden.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.rule_order
	// +required
	RuleOrder *int64 `json:"ruleOrder,omitempty"`

	// Optional. Rule to specify how a single entity should be renamed.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.single_entity_rename
	SingleEntityRename *SingleEntityRename `json:"singleEntityRename,omitempty"`

	// Optional. Rule to specify how multiple entities should be renamed.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.multi_entity_rename
	MultiEntityRename *MultiEntityRename `json:"multiEntityRename,omitempty"`

	// Optional. Rule to specify how multiple entities should be relocated into
	//  a different schema.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.entity_move
	EntityMove *EntityMove `json:"entityMove,omitempty"`

	// Optional. Rule to specify how a single column is converted.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.single_column_change
	SingleColumnChange *SingleColumnChange `json:"singleColumnChange,omitempty"`

	// Optional. Rule to specify how multiple columns should be converted to a
	//  different data type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.multi_column_data_type_change
	MultiColumnDataTypeChange *MultiColumnDatatypeChange `json:"multiColumnDataTypeChange,omitempty"`

	// Optional. Rule to specify how the data contained in a column should be
	//  transformed (such as trimmed, rounded, etc) provided that the data meets
	//  certain criteria.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.conditional_column_set_value
	ConditionalColumnSetValue *ConditionalColumnSetValue `json:"conditionalColumnSetValue,omitempty"`

	// Optional. Rule to specify how multiple tables should be converted with an
	//  additional rowid column.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.convert_rowid_column
	ConvertRowidColumn *ConvertRowIDToColumn `json:"convertRowidColumn,omitempty"`

	// Optional. Rule to specify the primary key for a table
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.set_table_primary_key
	SetTablePrimaryKey *SetTablePrimaryKey `json:"setTablePrimaryKey,omitempty"`

	// Optional. Rule to specify how a single package is converted.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.single_package_change
	SinglePackageChange *SinglePackageChange `json:"singlePackageChange,omitempty"`

	// Optional. Rule to change the sql code for an entity, for example,
	//  function, procedure.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.source_sql_change
	SourceSQLChange *SourceSQLChange `json:"sourceSQLChange,omitempty"`

	// Optional. Rule to specify the list of columns to include or exclude from
	//  a table.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.filter_table_columns
	FilterTableColumns *FilterTableColumns `json:"filterTableColumns,omitempty"`
}

// DataMigrationMappingRuleStatus defines the config connector machine state of DataMigrationMappingRule
type DataMigrationMappingRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataMigrationMappingRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataMigrationMappingRuleObservedState `json:"observedState,omitempty"`
}

// DataMigrationMappingRuleObservedState is the state of the DataMigrationMappingRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.MappingRule
type DataMigrationMappingRuleObservedState struct {
	// Output only. The revision ID of the mapping rule.
	//  A new revision is committed whenever the mapping rule is changed in any
	//  way. The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.clouddms.v1.MappingRule.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatamigrationmappingrule;gcpdatamigrationmappingrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataMigrationMappingRule is the Schema for the DataMigrationMappingRule API
// +k8s:openapi-gen=true
type DataMigrationMappingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataMigrationMappingRuleSpec   `json:"spec,omitempty"`
	Status DataMigrationMappingRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataMigrationMappingRuleList contains a list of DataMigrationMappingRule
type DataMigrationMappingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataMigrationMappingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataMigrationMappingRule{}, &DataMigrationMappingRuleList{})
}
