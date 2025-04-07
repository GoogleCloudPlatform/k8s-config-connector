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
	// GCP Resource Reference type.
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEntryGroupGVK = GroupVersion.WithKind("DataplexEntryGroup")

// The Parent resource that the DataplexEntryGroup resource resides in.
type DataplexEntryGroupParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +required
	Location string `json:"location"`
}

// DataplexEntryGroupSpec defines the desired state of DataplexEntryGroup
// +kcc:proto=google.cloud.dataplex.v1.EntryGroup
type DataplexEntryGroupSpec struct {
	DataplexEntryGroupParent `json:",inline"`
	// The DataplexEntryGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Description of the EntryGroup.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the EntryGroup.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// This checksum is computed by the service, and might be sent on update and
	//  delete requests to ensure the client has an up-to-date value before
	//  proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.etag
	Etag *string `json:"etag,omitempty"`
}

// DataplexEntryGroupStatus defines the config connector machine state of DataplexEntryGroup
type DataplexEntryGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEntryGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexEntryGroupObservedState `json:"observedState,omitempty"`
}

// DataplexEntryGroupObservedState is the state of the DataplexEntryGroup resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataplex.v1.EntryGroup
type DataplexEntryGroupObservedState struct {
	// Output only. The relative resource name of the EntryGroup, in the format
	//  projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the EntryGroup. If you
	//  delete and recreate the EntryGroup with the same name, this ID will be
	//  different.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the EntryGroup was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the EntryGroup was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Denotes the transfer status of the Entry Group. It is
	//  unspecified for Entry Group created from Dataplex API.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.transfer_status
	TransferStatus *string `json:"transferStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexentrygroup;gcpdataplexentrygroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEntryGroup is the Schema for the DataplexEntryGroup API
// +k8s:openapi-gen=true
type DataplexEntryGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEntryGroupSpec   `json:"spec,omitempty"`
	Status DataplexEntryGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEntryGroupList contains a list of DataplexEntryGroup
type DataplexEntryGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEntryGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEntryGroup{}, &DataplexEntryGroupList{})
}
