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

// API sources for DataplexDataAttribute, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/data_taxonomy.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.dataTaxonomies.attributes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexDataAttributeGVK = GroupVersion.WithKind("DataplexDataAttribute")

// DataplexDataAttributeSpec defines the desired state of DataplexDataAttribute
// +kcc:spec:proto=google.cloud.dataplex.v1.DataAttribute
// +kcc:required-from-proto
type DataplexDataAttributeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexDataAttribute name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Description of the DataAttribute.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the DataAttribute.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The ID of the parent DataAttribute resource, should belong to the
	//  same data taxonomy. Circular dependency in parent chain is not valid.
	//  Maximum depth of the hierarchy allowed is 4.
	//  [a -> b -> c -> d -> e, depth = 4]
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.parent_id
	ParentID *string `json:"parentID,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Specified when applied to a resource (eg: Cloud Storage bucket,
	//  BigQuery dataset, BigQuery table).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.resource_access_spec
	ResourceAccessSpec *ResourceAccessSpec `json:"resourceAccessSpec,omitempty"`

	// Optional. Specified when applied to data stored on the resource (eg: rows,
	//  columns in BigQuery Tables).
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.data_access_spec
	DataAccessSpec *DataAccessSpec `json:"dataAccessSpec,omitempty"`
}

// DataplexDataAttributeStatus defines the config connector machine state of DataplexDataAttribute
type DataplexDataAttributeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexDataAttribute resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexDataAttributeObservedState `json:"observedState,omitempty"`
}

// DataplexDataAttributeObservedState is the state of the DataplexDataAttribute resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataAttribute
type DataplexDataAttributeObservedState struct {
	// Output only. System generated globally unique ID for the DataAttribute.
	//  This ID will be different if the DataAttribute is deleted and re-created
	//  with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the DataAttribute was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the DataAttribute was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The number of child attributes present for this attribute.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataAttribute.attribute_count
	AttributeCount *int32 `json:"attributeCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexdataattribute;gcpdataplexdataattributes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexDataAttribute is the Schema for the DataplexDataAttribute API
// +k8s:openapi-gen=true
type DataplexDataAttribute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexDataAttributeSpec   `json:"spec,omitempty"`
	Status DataplexDataAttributeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexDataAttributeList contains a list of DataplexDataAttribute
type DataplexDataAttributeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexDataAttribute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexDataAttribute{}, &DataplexDataAttributeList{})
}
