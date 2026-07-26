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

var ModelArmorFloorSettingGVK = GroupVersion.WithKind("ModelArmorFloorSetting")

// ModelArmorFloorSettingSpec defines the desired state of ModelArmorFloorSetting
// +kcc:spec:proto=google.cloud.modelarmor.v1.FloorSetting
type ModelArmorFloorSettingSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ModelArmorFloorSetting name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. ModelArmor filter configuration.
	// +required
	FilterConfig *FilterConfig `json:"filterConfig"`

	// Optional. Floor Settings enforcement status.
	EnableFloorSettingEnforcement *bool `json:"enableFloorSettingEnforcement,omitempty"`

	// Optional. List of integrated services for which the floor setting is applicable.
	IntegratedServices []string `json:"integratedServices,omitempty"`

	// Optional. AI Platform floor setting.
	AiPlatformFloorSetting *AiPlatformFloorSetting `json:"aiPlatformFloorSetting,omitempty"`

	// Optional. Metadata for FloorSetting
	FloorSettingMetadata *FloorSetting_FloorSettingMetadata `json:"floorSettingMetadata,omitempty"`
}

// ModelArmorFloorSettingStatus defines the config connector machine state of ModelArmorFloorSetting
type ModelArmorFloorSettingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ModelArmorFloorSetting resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ModelArmorFloorSettingObservedState `json:"observedState,omitempty"`
}

// ModelArmorFloorSettingObservedState is the state of the ModelArmorFloorSetting resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.modelarmor.v1.FloorSetting
type ModelArmorFloorSettingObservedState struct {
	// Output only. [Output only] Create timestamp
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmodelarmorfloorsetting;gcpmodelarmorfloorsettings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ModelArmorFloorSetting is the Schema for the ModelArmorFloorSetting API
// +k8s:openapi-gen=true
type ModelArmorFloorSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ModelArmorFloorSettingSpec   `json:"spec,omitempty"`
	Status ModelArmorFloorSettingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ModelArmorFloorSettingList contains a list of ModelArmorFloorSetting
type ModelArmorFloorSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelArmorFloorSetting `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelArmorFloorSetting{}, &ModelArmorFloorSettingList{})
}

// +kcc:proto=google.cloud.modelarmor.v1.AiPlatformFloorSetting
type AiPlatformFloorSetting struct {
	// Optional. If true, Model Armor filters will be run in inspect only mode. No action will be taken on the request.
	InspectOnly *bool `json:"inspectOnly,omitempty"`

	// Optional. If true, Model Armor filters will be run in inspect and block mode. Requests that trip Model Armor filters will be blocked.
	InspectAndBlock *bool `json:"inspectAndBlock,omitempty"`

	// Optional. If true, log Model Armor filter results to Cloud Logging.
	EnableCloudLogging *bool `json:"enableCloudLogging,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.FloorSetting.FloorSettingMetadata
type FloorSetting_FloorSettingMetadata struct {
	// Optional. Metadata for multi language detection.
	MultiLanguageDetection *FloorSetting_FloorSettingMetadata_MultiLanguageDetection `json:"multiLanguageDetection,omitempty"`
}

// +kcc:proto=google.cloud.modelarmor.v1.FloorSetting.FloorSettingMetadata.MultiLanguageDetection
type FloorSetting_FloorSettingMetadata_MultiLanguageDetection struct {
	// Required. If true, multi language detection will be enabled.
	// +required
	EnableMultiLanguageDetection *bool `json:"enableMultiLanguageDetection,omitempty"`
}
