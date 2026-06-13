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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstanceTemplateGVK = GroupVersion.WithKind("ComputeInstanceTemplate")

// +kcc:proto=google.cloud.compute.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// The number of the guest accelerator cards exposed to this instance.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// The full or partial URL of the accelerator type to attach to this instance.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.AcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// ComputeInstanceTemplateSpec defines the desired state of ComputeInstanceTemplate
// +kcc:spec:proto=google.cloud.compute.v1.InstanceTemplate
type ComputeInstanceTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeInstanceTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	/* An optional description of this resource. Provide this property when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The instance properties for this instance template. */
	// +optional
	Properties *InstanceProperties `json:"properties,omitempty"`

	/* The source instance used to create the template. */
	// +optional
	SourceInstanceRef *InstanceRef `json:"sourceInstanceRef,omitempty"`

	/* The source instance params to use to create this instance template. */
	// +optional
	SourceInstanceParams *SourceInstanceParams `json:"sourceInstanceParams,omitempty"`
}

// ComputeInstanceTemplateStatus defines the config connector machine state of ComputeInstanceTemplate
type ComputeInstanceTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstanceTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstanceTemplateObservedState `json:"observedState,omitempty"`
}

// ComputeInstanceTemplateObservedState is the state of the ComputeInstanceTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.InstanceTemplate
type ComputeInstanceTemplateObservedState struct {
	// [Output Only] The creation timestamp for this instance template in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] A unique identifier for this instance template. The server defines this identifier.
	// +optional
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] The resource type, which is always compute#instanceTemplate for instance templates.
	// +optional
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the instance template resides. Only applicable for regional resources.
	// +optional
	Region *string `json:"region,omitempty"`

	// [Output Only] The URL for this instance template. The server defines this URL.
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancetemplate;gcpcomputeinstancetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceTemplate is the Schema for the ComputeInstanceTemplate API
// +k8s:openapi-gen=true
type ComputeInstanceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceTemplateSpec   `json:"spec,omitempty"`
	Status ComputeInstanceTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceTemplateList contains a list of ComputeInstanceTemplate
type ComputeInstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceTemplate{}, &ComputeInstanceTemplateList{})
}
