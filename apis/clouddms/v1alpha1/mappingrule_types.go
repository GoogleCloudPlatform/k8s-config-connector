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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSMappingRuleGVK = GroupVersion.WithKind("CloudDMSMappingRule")

// CloudDMSMappingRuleSpec defines the desired state of CloudDMSMappingRule
// +kcc:spec:proto=google.cloud.clouddms.v1.MappingRule
type CloudDMSMappingRuleSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The CloudDMSMappingRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The rule display name.
	// +kcc:proto=display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The rule scope
	// +kcc:proto=rule_scope
	RuleScope *string `json:"ruleScope,omitempty"`

	// The rule filter
	// +kcc:proto=filter
	Filter *MappingRuleFilter `json:"filter,omitempty"`

	ConnectionProfileRef CloudDMSConnectionProfileRef `json:"connectionProfileRef,omitempty"`
}

// CloudDMSMappingRuleStatus defines the config connector machine state of CloudDMSMappingRule
type CloudDMSMappingRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDMSMappingRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDMSMappingRuleObservedState `json:"observedState,omitempty"`
}

// CloudDMSMappingRuleObservedState is the state of the CloudDMSMappingRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.MappingRule
type CloudDMSMappingRuleObservedState struct {
	// The current state of the mapping rule.
	// +kcc:proto=state
	State *string `json:"state,omitempty"`

	// The revision ID of the mapping rule.
	// +kcc:proto=revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// The timestamp that the revision was created.
	// +kcc:proto=revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsmappingrule;gcpclouddmsmappingrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDMSMappingRule is the Schema for the CloudDMSMappingRule API
// +k8s:openapi-gen=true
type CloudDMSMappingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDMSMappingRuleSpec   `json:"spec,omitempty"`
	Status CloudDMSMappingRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDMSMappingRuleList contains a list of CloudDMSMappingRule
type CloudDMSMappingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDMSMappingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDMSMappingRule{}, &CloudDMSMappingRuleList{})
}

// +kcc:proto=google.cloud.clouddms.v1.ConditionalColumnSetValue
type ConditionalColumnSetValue struct {
	// Optional. Optional filter on source column length. Used for text based
	//  data types like varchar.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.source_text_filter
	SourceTextFilter *SourceTextFilter `json:"sourceTextFilter,omitempty"`

	// Optional. Optional filter on source column precision and scale. Used for
	//  fixed point numbers such as NUMERIC/NUMBER data types.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.source_numeric_filter
	SourceNumericFilter *SourceNumericFilter `json:"sourceNumericFilter,omitempty"`

	// Required. Description of data transformation during migration.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.value_transformation
	ValueTransformation *ValueTransformation `json:"valueTransformation,omitempty"`

	// NOTYET
	// Optional. Custom engine specific features.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConditionalColumnSetValue.custom_features
	// CustomFeatures apiextensionsv1.JSON `json:"customFeatures,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.MultiColumnDatatypeChange
type MultiColumnDatatypeChange struct {
	// Required. Filter on source data type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.source_data_type_filter
	SourceDataTypeFilter *string `json:"sourceDataTypeFilter,omitempty"`

	// Optional. Filter for text-based data types like varchar.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.source_text_filter
	SourceTextFilter *SourceTextFilter `json:"sourceTextFilter,omitempty"`

	// Optional. Filter for fixed point number data types such as
	//  NUMERIC/NUMBER.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.source_numeric_filter
	SourceNumericFilter *SourceNumericFilter `json:"sourceNumericFilter,omitempty"`

	// Required. New data type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.new_data_type
	NewDataType *string `json:"newDataType,omitempty"`

	// Optional. Column length - e.g. varchar (50) - if not specified and relevant
	//  uses the source column length.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_length
	OverrideLength *int64 `json:"overrideLength,omitempty"`

	// Optional. Column scale - when relevant - if not specified and relevant
	//  uses the source column scale.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_scale
	OverrideScale *int32 `json:"overrideScale,omitempty"`

	// Optional. Column precision - when relevant - if not specified and relevant
	//  uses the source column precision.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_precision
	OverridePrecision *int32 `json:"overridePrecision,omitempty"`

	// Optional. Column fractional seconds precision - used only for timestamp
	//  based datatypes - if not specified and relevant uses the source column
	//  fractional seconds precision.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.override_fractional_seconds_precision
	OverrideFractionalSecondsPrecision *int32 `json:"overrideFractionalSecondsPrecision,omitempty"`

	// NOTYET
	// Optional. Custom engine specific features.
	// +kcc:proto:field=google.cloud.clouddms.v1.MultiColumnDatatypeChange.custom_features
	// CustomFeatures apiextensionsv1.JSON `json:"customFeatures,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.SingleColumnChange
type SingleColumnChange struct {
	// Optional. Column data type name.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.data_type
	DataType *string `json:"dataType,omitempty"`

	// Optional. Charset override - instead of table level charset.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.charset
	Charset *string `json:"charset,omitempty"`

	// Optional. Collation override - instead of table level collation.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.collation
	Collation *string `json:"collation,omitempty"`

	// Optional. Column length - e.g. 50 as in varchar (50) - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.length
	Length *int64 `json:"length,omitempty"`

	// Optional. Column precision - e.g. 8 as in double (8,2) - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.precision
	Precision *int32 `json:"precision,omitempty"`

	// Optional. Column scale - e.g. 2 as in double (8,2) - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.scale
	Scale *int32 `json:"scale,omitempty"`

	// Optional. Column fractional seconds precision - e.g. 2 as in timestamp (2)
	//  - when relevant.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.fractional_seconds_precision
	FractionalSecondsPrecision *int32 `json:"fractionalSecondsPrecision,omitempty"`

	// Optional. Is the column of array type.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.array
	Array *bool `json:"array,omitempty"`

	// Optional. The length of the array, only relevant if the column type is an
	//  array.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.array_length
	ArrayLength *int32 `json:"arrayLength,omitempty"`

	// Optional. Is the column nullable.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// Optional. Is the column auto-generated/identity.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.auto_generated
	AutoGenerated *bool `json:"autoGenerated,omitempty"`

	// Optional. Is the column a UDT (User-defined Type).
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.udt
	Udt *bool `json:"udt,omitempty"`

	// NOTYET
	// Optional. Custom engine specific features.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.custom_features
	// CustomFeatures apiextensionsv1.JSON `json:"customFeatures,omitempty"`

	// Optional. Specifies the list of values allowed in the column.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.set_values
	SetValues []string `json:"setValues,omitempty"`

	// Optional. Comment associated with the column.
	// +kcc:proto:field=google.cloud.clouddms.v1.SingleColumnChange.comment
	Comment *string `json:"comment,omitempty"`
}
