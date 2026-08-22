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

var VertexAISecurityPolicyGVK = GroupVersion.WithKind("VertexAISecurityPolicy")

// VertexAISecurityPolicySpec defines the desired state of VertexAISecurityPolicy
// +kcc:spec:proto=google.cloud.modelarmor.v1.Template
type VertexAISecurityPolicySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAISecurityPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. ModelArmor filter configuration.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.filter_config
	// +required
	FilterConfig *FilterConfig `json:"filterConfig"`

	// Optional. Floor Settings enforcement status.
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.enable_floor_setting_enforcement
	EnableFloorSettingEnforcement *bool `json:"enableFloorSettingEnforcement,omitempty"`

	// Optional. Metadata for Template
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.template_metadata
	TemplateMetadata *Template_TemplateMetadata `json:"templateMetadata,omitempty"`
}

// VertexAISecurityPolicyStatus defines the config connector machine state of VertexAISecurityPolicy
type VertexAISecurityPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAISecurityPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAISecurityPolicyObservedState `json:"observedState,omitempty"`
}

// VertexAISecurityPolicyObservedState is the state of the VertexAISecurityPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.modelarmor.v1.Template
type VertexAISecurityPolicyObservedState struct {
	// Output only. Create timestamp
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp
	// +kcc:proto:field=google.cloud.modelarmor.v1.Template.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaisecuritypolicy;gcpvertexaisecuritypolicys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAISecurityPolicy is the Schema for the VertexAISecurityPolicy API
// +k8s:openapi-gen=true
type VertexAISecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAISecurityPolicySpec   `json:"spec,omitempty"`
	Status VertexAISecurityPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAISecurityPolicyList contains a list of VertexAISecurityPolicy
type VertexAISecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAISecurityPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAISecurityPolicy{}, &VertexAISecurityPolicyList{})
}
