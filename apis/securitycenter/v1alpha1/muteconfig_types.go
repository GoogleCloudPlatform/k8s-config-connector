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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecurityCenterMuteConfigGVK = GroupVersion.WithKind("SecurityCenterMuteConfig")

// SecurityCenterMuteConfigSpec defines the desired state of SecurityCenterMuteConfig
// +kcc:spec:proto=google.cloud.securitycenter.v1.MuteConfig
type SecurityCenterMuteConfigSpec struct {
	// The organization, folder, or project that this resource belongs to.
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The SecurityCenterMuteConfig name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The human-readable name to be displayed for the mute config.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// A description of this mute config (max 4000 characters).
	// +optional
	Description *string `json:"description,omitempty"`

	// Required. An expression that defines the filter to apply across create/update events of findings.
	// +required
	Filter *string `json:"filter,omitempty"`

	// Required. The type of the mute config. Possible values: ["STATIC", "DYNAMIC"]
	// +required
	Type *string `json:"type,omitempty"`

	// Optional. The expiry of the mute config. Only applicable for dynamic configs.
	// +optional
	ExpiryTime *string `json:"expiryTime,omitempty"`
}

// SecurityCenterMuteConfigStatus defines the observed state of SecurityCenterMuteConfig
type SecurityCenterMuteConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecurityCenterMuteConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecurityCenterMuteConfigObservedState `json:"observedState,omitempty"`
}

// SecurityCenterMuteConfigObservedState defines the observed state of SecurityCenterMuteConfig
// +kcc:observedstate:proto=google.cloud.securitycenter.v1.MuteConfig
type SecurityCenterMuteConfigObservedState struct {
	// Output only. The time at which the mute config was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the mute config was last updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The email address of the person who last edited this mute config.
	// +optional
	MostRecentEditor *string `json:"mostRecentEditor,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycentermuteconfig;gcpsecuritycentermuteconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityCenterMuteConfig is the Schema for the SecurityCenterMuteConfig API
// +k8s:openapi-gen=true
type SecurityCenterMuteConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityCenterMuteConfigSpec   `json:"spec,omitempty"`
	Status SecurityCenterMuteConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecurityCenterMuteConfigList contains a list of SecurityCenterMuteConfig
type SecurityCenterMuteConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterMuteConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityCenterMuteConfig{}, &SecurityCenterMuteConfigList{})
}
