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

// API sources for MapmanagementMapContextConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/maps/mapmanagement/v2/map_management_service.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MapmanagementMapContextConfigGVK = GroupVersion.WithKind("MapmanagementMapContextConfig")

// MapmanagementMapContextConfigSpec defines the desired state of MapmanagementMapContextConfig
// +kcc:spec:proto=google.maps.mapmanagement.v2.MapContextConfig
// +kcc:required-from-proto
type MapmanagementMapContextConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The MapmanagementMapContextConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The MapConfig resource name that this MapContextConfig is
	//  associated with. Format: projects/{project}/mapConfigs/{map_config}. This
	//  field is required and cannot be omitted.
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.map_config
	// +required
	MapConfig *string `json:"mapConfig,omitempty"`

	// Required. The StyleConfig resource name that is styling this
	//  MapContextConfig. This field is required and cannot be omitted. Format:
	//  projects/{project}/styleConfigs/{style_config}
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.style_config
	// +required
	StyleConfig *string `json:"styleConfig,omitempty"`

	// Optional. The Dataset resource name that is associated with this
	//  MapContextConfig. This field is optional and can be omitted. If omitted, no
	//  datasets will be associated with the MapContextConfig. If a dataset is
	//  specified, it will be applied to the MapContextConfig. Format:
	//  projects/{project}/datasets/{dataset}
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.dataset
	Dataset []string `json:"dataset,omitempty"`

	// Optional. The user defined human readable name for this MapContextConfig.
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.alias
	Alias *string `json:"alias,omitempty"`

	// Required. The map variants that this MapContextConfig can be applied to. If
	//  empty, the MapContextConfig will be default applied to only the ROADMAP map
	//  variant.
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.map_variants
	// +required
	MapVariants []string `json:"mapVariants,omitempty"`
}

// MapmanagementMapContextConfigStatus defines the config connector machine state of MapmanagementMapContextConfig
type MapmanagementMapContextConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MapmanagementMapContextConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MapmanagementMapContextConfigObservedState `json:"observedState,omitempty"`
}

// MapmanagementMapContextConfigObservedState is the state of the MapmanagementMapContextConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.maps.mapmanagement.v2.MapContextConfig
type MapmanagementMapContextConfigObservedState struct {
	// Output only. Denotes the creation time of the MapContextConfig. Output
	//  only.
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Denotes the last update time of the MapContextConfig. Output
	//  only.
	// +kcc:proto:field=google.maps.mapmanagement.v2.MapContextConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmapmanagementmapcontextconfig;gcpmapmanagementmapcontextconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MapmanagementMapContextConfig is the Schema for the MapmanagementMapContextConfig API
// +k8s:openapi-gen=true
type MapmanagementMapContextConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MapmanagementMapContextConfigSpec   `json:"spec,omitempty"`
	Status MapmanagementMapContextConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MapmanagementMapContextConfigList contains a list of MapmanagementMapContextConfig
type MapmanagementMapContextConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MapmanagementMapContextConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MapmanagementMapContextConfig{}, &MapmanagementMapContextConfigList{})
}
