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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

var DataCatalogTagGVK = GroupVersion.WithKind("DataCatalogTag")

// Parent defines the parent resource for this DataCatalogTag
type DataCatalogTagParent struct {
	// Required. Reference to the DataCatalogEntry that owns this Tag.
	// The entry must be in the same project and location as the tag.
	// +required
	EntryRef *EntryRef `json:"entryRef,omitempty"`
}

// DataCatalogTagSpec defines the desired state of DataCatalogTag
// +kcc:proto=google.cloud.datacatalog.v1.Tag
type DataCatalogTagSpec struct {
	// Specifies the parent resource where this DataCatalogTag resides.
	// +required
	DataCatalogTagParent `json:",inline"`

	// The DataCatalogTag name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The resource name of the tag template this tag uses.
	//
	//  This field cannot be modified after creation.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.template
	// +required
	TemplateRef *TagTemplateRef `json:"templateRef,omitempty"`

	// Resources like entry can have schemas associated with them. This scope
	//  allows you to attach tags to an individual column based on that schema.
	//
	//  To attach a tag to a nested column, separate column names with a dot
	//  (`.`). Example: `column.nested_column`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.column
	Column *string `json:"column,omitempty"`

	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.fields
	Fields map[string]TagField `json:"fields,omitempty"`
}

// DataCatalogTagStatus defines the config connector machine state of DataCatalogTag
type DataCatalogTagStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogTag resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataCatalogTagObservedState `json:"observedState,omitempty"`
}

// DataCatalogTagObservedState is the state of the DataCatalogTag resource as most recently observed in GCP.
// +kcc:proto=google.cloud.datacatalog.v1.Tag
type DataCatalogTagObservedState struct {
	// Output only. The display name of the tag template.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.template_display_name
	TemplateDisplayName *string `json:"templateDisplayName,omitempty"`

	// Output only. Denotes the transfer status of the Tag Template.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Tag.dataplex_transfer_status
	DataplexTransferStatus *string `json:"dataplexTransferStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogtag;gcpdatacatalogtags
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogTag is the Schema for the DataCatalogTag API
// +k8s:openapi-gen=true
type DataCatalogTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogTagSpec   `json:"spec,omitempty"`
	Status DataCatalogTagStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogTagList contains a list of DataCatalogTag
type DataCatalogTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogTag `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogTag{}, &DataCatalogTagList{})
}
