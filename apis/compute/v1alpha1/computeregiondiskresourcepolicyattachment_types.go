// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeRegionDiskResourcePolicyAttachmentGVK = GroupVersion.WithKind("ComputeRegionDiskResourcePolicyAttachment")

type ComputeRegionDiskResourcePolicyAttachmentSpec struct {
	/* The disk to which the resource policy is attached. */
	// +required
	DiskRef *computev1beta1.ComputeDiskRef `json:"diskRef"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. A reference to the region where the disk resides. */
	// +required
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource policy. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ComputeRegionDiskResourcePolicyAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRegionDiskResourcePolicyAttachment's current state. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the ComputeRegionDiskResourcePolicyAttachment resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeregiondiskresourcepolicyattachment;gcpcomputeregiondiskresourcepolicyattachments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRegionDiskResourcePolicyAttachment is the Schema for the ComputeRegionDiskResourcePolicyAttachment API
// +k8s:openapi-gen=true
type ComputeRegionDiskResourcePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRegionDiskResourcePolicyAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeRegionDiskResourcePolicyAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRegionDiskResourcePolicyAttachmentList contains a list of ComputeRegionDiskResourcePolicyAttachment
type ComputeRegionDiskResourcePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRegionDiskResourcePolicyAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRegionDiskResourcePolicyAttachment{}, &ComputeRegionDiskResourcePolicyAttachmentList{})
}
