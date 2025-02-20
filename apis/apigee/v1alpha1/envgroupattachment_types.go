// Copyright 2025 Google LLC.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EnvgroupAttachmentGVK = GroupVersion.WithKind("ApigeeEnvgroupAttachment")

// ApigeeEnvgroupAttachmentSpec defines the desired state of EnvgroupAttachment
type ApigeeEnvgroupAttachmentSpec struct {
	// +required
	OrganizationRef *apigeev1beta1.OrganizationRef `json:"organizationRef"`

	// Immutable. The Apigee environment group which will host the environment.
	EnvgroupRef *EnvironmentGroupRef `json:"envgroupRef"`

	// Immutable. The Apigee environment to attach to.
	EnvironmentRef *apigeev1beta1.EnvironmentRef `json:"environmentRef"`

	// The EnvgroupAttachment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ApigeeEnvgroupAttachmentStatus defines the observed state of EnvgroupAttachment
type ApigeeEnvgroupAttachmentStatus struct {
	// Conditions represent the latest available observations of the
	//   EnvgroupAttachment's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EnvgroupAttachment resource.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EnvgroupAttachmentObservedState `json:"observedState,omitempty"`
}

// EnvgroupAttachmentObservedState defines the desired state of ApigeeEnvgroupAttachment
type EnvgroupAttachmentObservedState struct {
	// Output only. The time at which the environment group attachment
	// was created as milliseconds since epoch.
	CreatedAt *string `json:"createdAt,omitempty"`

	// Output only. ID of the environment group.
	EnvironmentGroupID *string `json:"environmentGroupID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApigeeEnvgroupAttachment is the Schema for the EnvgroupAttachments API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeenvgroupattachment;gcpapigeeenvgroupattachments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
type ApigeeEnvgroupAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeEnvgroupAttachmentSpec   `json:"spec,omitempty"`
	Status ApigeeEnvgroupAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApigeeEnvgroupAttachmentList contains a list of EnvgroupAttachment
type ApigeeEnvgroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeEnvgroupAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeEnvgroupAttachment{}, &ApigeeEnvgroupAttachmentList{})
}
