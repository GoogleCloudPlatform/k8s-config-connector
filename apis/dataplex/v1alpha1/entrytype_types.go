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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEntryTypeGVK = GroupVersion.WithKind("DataplexEntryType")

type DataplexEntryTypeParent struct {
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	Location string `json:"location"`
}

// DataplexEntryTypeSpec defines the desired state of DataplexEntryType
// +kcc:proto=google.cloud.dataplex.v1.EntryType
type DataplexEntryTypeSpec struct {
	DataplexEntryTypeParent `json:",inline"`
	// The DataplexEntryType name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// AspectInfo contains overriding configuration for aspects.
	// +kcc:proto=required_aspects
	RequiredAspects []EntryType_AspectInfo `json:"requiredAspects,omitempty"`

	// Authorization contains constraints on the visibility of Entries that conform
	//  to the EntryType.
	Authorization *EntryType_Authorization `json:"authorization,omitempty"`
}

// DataplexEntryTypeStatus defines the config connector machine state of DataplexEntryType
type DataplexEntryTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEntryType resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexentrytype;gcpdataplexentrytypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEntryType is the Schema for the DataplexEntryType API
// +k8s:openapi-gen=true
type DataplexEntryType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEntryTypeSpec   `json:"spec,omitempty"`
	Status DataplexEntryTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEntryTypeList contains a list of DataplexEntryType
type DataplexEntryTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEntryType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEntryType{}, &DataplexEntryTypeList{})
}
