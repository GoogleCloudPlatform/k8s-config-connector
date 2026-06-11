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

var SecurityCenterMuteConfigGVK = GroupVersion.WithKind("SecurityCenterMuteConfig")

// SecurityCenterMuteConfigSpec defines the desired state of SecurityCenterMuteConfig
// +kcc:spec:proto=google.cloud.securitycenter.v1.MuteConfig
type SecurityCenterMuteConfigSpec struct {
	// The organization that this resource belongs to.
	// +kubebuilder:validation:Required
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The SecurityCenterMuteConfig name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// A description of the mute config.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Required. An expression that defines the filter to apply across create/update events of findings.
	// While creating a filter string, be mindful of the scope in which the mute configuration is being created.
	// E.g., If a filter contains project = X but is created under the project = Y scope, it might not match any findings.
	// +kubebuilder:validation:Required
	Filter *string `json:"filter"`

	// Optional. The type of the mute config, which determines what type of mute state the config affects.
	// The static mute state takes precedence over the dynamic mute state.
	// Immutable after creation. STATIC by default if not set during creation.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=MUTE_CONFIG_TYPE_UNSPECIFIED;STATIC;DYNAMIC
	Type *string `json:"type,omitempty"`

	// Optional. The expiry of the mute config. Only applicable for dynamic configs.
	// If the expiry is set, when the config expires, it is removed from all findings.
	// +kubebuilder:validation:Optional
	ExpiryTime *string `json:"expiryTime,omitempty"`
}

// SecurityCenterMuteConfigStatus defines the config connector machine state of SecurityCenterMuteConfig
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

// SecurityCenterMuteConfigObservedState is the state of the SecurityCenterMuteConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securitycenter.v1.MuteConfig
type SecurityCenterMuteConfigObservedState struct {
	// Output only. The time at which the mute config was created.
	// This field is set by the server and will be ignored if provided on config creation.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the mute config was updated.
	// This field is set by the server and will be ignored if provided on config creation or update.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Email address of the user who last edited the mute config.
	// This field is set by the server and will be ignored if provided on config creation or update.
	// +kubebuilder:validation:Optional
	MostRecentEditor *string `json:"mostRecentEditor,omitempty"`
}

// +kubebuilder:metadata:labels="cnrm.cloud.google.com/unverified-greenfield=true"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycentermuteconfig;gcpsecuritycentermuteconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityCenterMuteConfig is the Schema for the SecurityCenterMuteConfig API
// +k8s:openapi-gen=true
type SecurityCenterMuteConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
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
