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

// API sources for ParameterManagerTemplate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/4fabd6e4b356307d319b55a86a3f200e9c52b594/google/cloud/parametermanager/v1/service.proto
// +kcc:source:service-docs=https://cloud.google.com/secret-manager/parameter-manager/docs/overview
// +kcc:source:resource-docs=https://cloud.google.com/secret-manager/parameter-manager/docs/reference/rest/v1/projects.locations.templates

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ParameterManagerTemplateGVK = GroupVersion.WithKind("ParameterManagerTemplate")

// ParameterManagerTemplateSpec defines the desired state of ParameterManagerTemplate
// +kcc:spec:proto=google.cloud.parametermanager.v1.Template
// +kcc:required-from-proto
type ParameterManagerTemplateSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The ParameterManagerTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Specifies the format of a Template.
	// +kcc:proto:field=google.cloud.parametermanager.v1.Template.format
	Format *string `json:"format,omitempty"`
}

// ParameterManagerTemplateStatus defines the config connector machine state of ParameterManagerTemplate
type ParameterManagerTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ParameterManagerTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ParameterManagerTemplateObservedState `json:"observedState,omitempty"`
}

// ParameterManagerTemplateObservedState is the state of the ParameterManagerTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.parametermanager.v1.Template
type ParameterManagerTemplateObservedState struct {
	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.Template.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.parametermanager.v1.Template.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpparametermanagertemplate;gcpparametermanagertemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ParameterManagerTemplate is the Schema for the ParameterManagerTemplate API
// +k8s:openapi-gen=true
type ParameterManagerTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ParameterManagerTemplateSpec   `json:"spec,omitempty"`
	Status ParameterManagerTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ParameterManagerTemplateList contains a list of ParameterManagerTemplate
type ParameterManagerTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterManagerTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ParameterManagerTemplate{}, &ParameterManagerTemplateList{})
}
