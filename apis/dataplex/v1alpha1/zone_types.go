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
)

var DataplexZoneGVK = GroupVersion.WithKind("DataplexZone")

// Parent defines the parent resource for the DataplexZone.
type DataplexZoneParent struct {
	// Reference to the parent DataplexLake that owns this Zone.
	// +required
	LakeRef *LakeRef `json:"lakeRef"`
}

// DataplexZoneSpec defines the desired state of DataplexZone
// +kcc:spec:proto=google.cloud.dataplex.v1.Zone
type DataplexZoneSpec struct {
	DataplexZoneParent `json:",inline"`

	// The DataplexZone name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The type of the zone.
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.type
	Type *string `json:"type"`

	// Optional. Specification of the discovery feature applied to data in this
	//  zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.discovery_spec
	DiscoverySpec *Zone_DiscoverySpec `json:"discoverySpec,omitempty"`

	// Required. Specification of the resources that are referenced by the assets
	//  within this zone.
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.resource_spec
	ResourceSpec *Zone_ResourceSpec `json:"resourceSpec"`
}

// DataplexZoneStatus defines the config connector machine state of DataplexZone
type DataplexZoneStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexZone resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexZoneObservedState `json:"observedState,omitempty"`
}

// DataplexZoneObservedState is the state of the DataplexZone resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Zone
type DataplexZoneObservedState struct {
	// Output only. System generated globally unique ID for the zone. This ID will
	//  be different if the zone is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the zone was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the zone was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.state
	State *string `json:"state,omitempty"`

	// Output only. Aggregated status of the underlying assets of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexzone;gcpdataplexzones
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexZone is the Schema for the DataplexZone API
// +k8s:openapi-gen=true
type DataplexZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexZoneSpec   `json:"spec,omitempty"`
	Status DataplexZoneStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexZoneList contains a list of DataplexZone
type DataplexZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexZone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexZone{}, &DataplexZoneList{})
}
