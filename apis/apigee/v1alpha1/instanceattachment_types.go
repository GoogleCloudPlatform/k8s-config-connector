// Copyright 2025 Google LLC
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
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeInstanceAttachmentGVK = GroupVersion.WithKind("ApigeeInstanceAttachment")

// ApigeeInstanceAttachmentSpec defines the desired state of ApigeeInstanceAttachment
// +kcc:spec:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1InstanceAttachment
type ApigeeInstanceAttachmentSpec struct {
	// Reference to parent Apigee Instance.
	// +required
	InstanceRef *ApigeeInstanceRef `json:"instanceRef"`

	// ID of the attached environment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1InstanceAttachment.environment
	// +required
	EnvironmentRef *apigeev1beta1.ApigeeEnvironmentRef `json:"environmentRef,omitempty"`

	// The ApigeeInstanceAttachment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ApigeeInstanceAttachmentStatus defines the config connector machine state of ApigeeInstanceAttachment
type ApigeeInstanceAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeInstanceAttachment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeInstanceAttachmentObservedState `json:"observedState,omitempty"`
}

// ApigeeInstanceAttachmentSpec defines the desired state of ApigeeInstanceAttachment
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1InstanceAttachment
// ApigeeInstanceAttachmentObservedState is the state of the ApigeeInstanceAttachment resource as most recently observed in GCP.
type ApigeeInstanceAttachmentObservedState struct {
	// Output only. Time the attachment was created in milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1InstanceAttachment.created_at
	CreatedAt *string `json:"createdAt,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeinstanceattachment;gcpapigeeinstanceattachments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeInstanceAttachment is the Schema for the ApigeeInstanceAttachment API
// +k8s:openapi-gen=true
type ApigeeInstanceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeInstanceAttachmentSpec   `json:"spec,omitempty"`
	Status ApigeeInstanceAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeInstanceAttachmentList contains a list of ApigeeInstanceAttachment
type ApigeeInstanceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeInstanceAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeInstanceAttachment{}, &ApigeeInstanceAttachmentList{})
}
