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
	apphubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DevConnectInsightsConfigGVK = GroupVersion.WithKind("DevConnectInsightsConfig")

// DevConnectInsightsConfigSpec defines the desired state of DevConnectInsightsConfig
// +kcc:spec:proto=google.cloud.developerconnect.v1.InsightsConfig
type DevConnectInsightsConfigSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	Location *string `json:"location"`

	// The DevConnectInsightsConfig name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The name of the App Hub Application.
	// +optional
	AppHubApplicationRef *apphubv1beta1.ApplicationRef `json:"appHubApplicationRef,omitempty"`

	// Optional. The artifact configurations of the artifacts that are deployed.
	// +optional
	ArtifactConfigs []ArtifactConfig `json:"artifactConfigs,omitempty"`
}

// DevConnectInsightsConfigStatus defines the config connector machine state of DevConnectInsightsConfig
type DevConnectInsightsConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DevConnectInsightsConfig resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *DevConnectInsightsConfigObservedState `json:"observedState,omitempty"`
}

// DevConnectInsightsConfigObservedState is the state of the DevConnectInsightsConfig resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.developerconnect.v1.InsightsConfig
type DevConnectInsightsConfigObservedState struct {
	// Output only. [Output only] Create timestamp
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The runtime configurations where the application is deployed.
	// +optional
	RuntimeConfigs []RuntimeConfigObservedState `json:"runtimeConfigs,omitempty"`

	// Output only. The state of the InsightsConfig.
	// +optional
	State *string `json:"state,omitempty"`

	// Output only. Reconciling (https://google.aip.dev/128#reconciliation).
	//  Set to true if the current state of InsightsConfig does not match the
	//  user's intended state, and the service is actively updating the resource to
	//  reconcile them. This can happen due to user-triggered updates or
	//  system actions like failover or maintenance.
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Any errors that occurred while setting up the InsightsConfig.
	//  Each error will be in the format: `field_name: error_message`, e.g.
	//  GetAppHubApplication: Permission denied while getting App Hub
	//  application. Please grant permissions to the P4SA.
	// +optional
	Errors []common.Status `json:"errors,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdevconnectinsightsconfig;gcpdevconnectinsightsconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DevConnectInsightsConfig is the Schema for the DevConnectInsightsConfig API
// +k8s:openapi-gen=true
type DevConnectInsightsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DevConnectInsightsConfigSpec   `json:"spec,omitempty"`
	Status DevConnectInsightsConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DevConnectInsightsConfigList contains a list of DevConnectInsightsConfig
type DevConnectInsightsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevConnectInsightsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevConnectInsightsConfig{}, &DevConnectInsightsConfigList{})
}
