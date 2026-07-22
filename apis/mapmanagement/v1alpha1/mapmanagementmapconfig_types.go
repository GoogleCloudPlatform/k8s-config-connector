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

var MapManagementMapConfigGVK = GroupVersion.WithKind("MapManagementMapConfig")

type MapFeatures struct {
	// Optional. The visual feature to use for this map.
	// +kubebuilder:validation:Optional
	SimpleFeatures []string `json:"simpleFeatures,omitempty"`

	// Optional. POI Boost level, where 0 denotes no boostings and negative values
	// denotes de-boosting. Boosted POIs are shown at lower zoom than default and
	// vice versa de-boosted. Currently supports 2 levels of boosting, so the
	// level is clamped to [-2, 2]. If not specified, the POI density defined in
	// the style sheet will be used if it exists. Otherwise, no POI density will
	// be applied.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=-2
	// +kubebuilder:validation:Maximum=2
	PoiBoostLevel *int32 `json:"poiBoostLevel,omitempty"`
}

// MapManagementMapConfigSpec defines the desired state of MapManagementMapConfig
// +kcc:spec:proto=google.maps.mapmanagement.v2beta.MapConfig
type MapManagementMapConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The MapManagementMapConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The display name of this MapConfig, as specified by the user.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of this MapConfig, as specified by the user.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. The Map Features that apply to this Map Config.
	// +kubebuilder:validation:Optional
	MapFeatures *MapFeatures `json:"mapFeatures,omitempty"`

	// Optional. Represents the Map Type of the MapConfig. If this is unset, the
	// default behavior is to use the raster map type.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=RASTER;VECTOR
	MapType *string `json:"mapType,omitempty"`
}

// MapManagementMapConfigStatus defines the config connector machine state of MapManagementMapConfig
type MapManagementMapConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by Config Connector. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MapManagementMapConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MapManagementMapConfigObservedState `json:"observedState,omitempty"`
}

// MapManagementMapConfigObservedState is the state of the MapManagementMapConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.maps.mapmanagement.v2beta.MapConfig
type MapManagementMapConfigObservedState struct {
	// Output only. The Map ID of this MapConfig, used to identify the map in
	// client applications. This read-only field is generated when the MapConfig
	// is created. Output only.
	MapID *string `json:"mapID,omitempty"`

	// Output only. Denotes the creation time of the Map Config. Output only.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Denotes the last update time of the Map Config. Output only.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmapmanagementmapconfig;gcpmapmanagementmapconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MapManagementMapConfig is the Schema for the MapManagementMapConfig API
// +k8s:openapi-gen=true
type MapManagementMapConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MapManagementMapConfigSpec   `json:"spec,omitempty"`
	Status MapManagementMapConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MapManagementMapConfigList contains a list of MapManagementMapConfig
type MapManagementMapConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MapManagementMapConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MapManagementMapConfig{}, &MapManagementMapConfigList{})
}
