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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudNumberRegistryRegistryBookGVK = GroupVersion.WithKind("CloudNumberRegistryRegistryBook")

// CloudNumberRegistryRegistryBookSpec defines the desired state of CloudNumberRegistryRegistryBook
// +kcc:spec:proto=google.cloud.numberregistry.v1alpha.RegistryBook
type CloudNumberRegistryRegistryBookSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CloudNumberRegistryRegistryBook name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-defined labels.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. List of scopes claimed by the RegistryBook. In Preview, Only
	//  project scope is supported. Each scope is in the format of
	//  projects/{project}. Each scope can only be claimed once.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.claimed_scopes
	ClaimedScopeRefs []refsv1beta1.ProjectRef `json:"claimedScopeRefs,omitempty"`
}

// CloudNumberRegistryRegistryBookStatus defines the config connector machine state of CloudNumberRegistryRegistryBook
type CloudNumberRegistryRegistryBookStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudNumberRegistryRegistryBook resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudNumberRegistryRegistryBookObservedState `json:"observedState,omitempty"`
}

// CloudNumberRegistryRegistryBookObservedState is the state of the CloudNumberRegistryRegistryBook resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.numberregistry.v1alpha.RegistryBook
type CloudNumberRegistryRegistryBookObservedState struct {
	// Output only. The time at which the RegistryBook was created.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the RegistryBook was last updated.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Whether the RegistryBook is the default one.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.is_default
	IsDefault *bool `json:"isDefault,omitempty"`

	// Output only. Aggregated data for the RegistryBook. Populated only when the
	//  view is AGGREGATE.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.aggregated_data
	AggregatedData *RegistryBook_AggregatedDataObservedState `json:"aggregatedData,omitempty"`
}

// +kcc:proto=google.cloud.numberregistry.v1alpha.RegistryBook.AggregatedData
type RegistryBook_AggregatedDataObservedState struct {
	// Output only. Number of scopes unique to the RegistryBook.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.AggregatedData.unique_scopes_count
	UniqueScopesCount *int32 `json:"uniqueScopesCount,omitempty"`

	// Output only. Number of discovered Realms in the RegistryBook.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.AggregatedData.discovered_realms_count
	DiscoveredRealmsCount *int32 `json:"discoveredRealmsCount,omitempty"`

	// Output only. Number of DiscoveredRanges in the RegistryBook.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.AggregatedData.discovered_ranges_count
	DiscoveredRangesCount *int32 `json:"discoveredRangesCount,omitempty"`

	// Output only. Number of custom Realms in the RegistryBook.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.AggregatedData.custom_realms_count
	CustomRealmsCount *int32 `json:"customRealmsCount,omitempty"`

	// Output only. Number of CustomRanges in the RegistryBook.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.RegistryBook.AggregatedData.custom_ranges_count
	CustomRangesCount *int32 `json:"customRangesCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudnumberregistryregistrybook;gcpcloudnumberregistryregistrybooks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudNumberRegistryRegistryBook is the Schema for the CloudNumberRegistryRegistryBook API
// +k8s:openapi-gen=true
type CloudNumberRegistryRegistryBook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudNumberRegistryRegistryBookSpec   `json:"spec,omitempty"`
	Status CloudNumberRegistryRegistryBookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudNumberRegistryRegistryBookList contains a list of CloudNumberRegistryRegistryBook
type CloudNumberRegistryRegistryBookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudNumberRegistryRegistryBook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudNumberRegistryRegistryBook{}, &CloudNumberRegistryRegistryBookList{})
}
