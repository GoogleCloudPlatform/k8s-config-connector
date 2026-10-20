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

// API sources for VideoStitcherSlate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/video/stitcher/v1/slates.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VideoStitcherSlateGVK = GroupVersion.WithKind("VideoStitcherSlate")

// VideoStitcherSlateSpec defines the desired state of VideoStitcherSlate
// +kcc:spec:proto=google.cloud.video.stitcher.v1.Slate
// +kcc:required-from-proto
type VideoStitcherSlateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VideoStitcherSlate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The URI to fetch the source content for the slate. This URI must return an
	//  MP4 video with at least one audio track.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.uri
	URI *string `json:"uri,omitempty"`

	// gam_slate has all the GAM-related attributes of slates.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.gam_slate
	GamSlate *Slate_GamSlate `json:"gamSlate,omitempty"`
}

// VideoStitcherSlateStatus defines the config connector machine state of VideoStitcherSlate
type VideoStitcherSlateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VideoStitcherSlate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VideoStitcherSlateObservedState `json:"observedState,omitempty"`
}

// VideoStitcherSlateObservedState is the state of the VideoStitcherSlate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.stitcher.v1.Slate
type VideoStitcherSlateObservedState struct {
	// gam_slate has all the GAM-related attributes of slates.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.gam_slate
	GamSlate *Slate_GamSlateObservedState `json:"gamSlate,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvideostitcherslate;gcpvideostitcherslates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VideoStitcherSlate is the Schema for the VideoStitcherSlate API
// +k8s:openapi-gen=true
type VideoStitcherSlate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VideoStitcherSlateSpec   `json:"spec,omitempty"`
	Status VideoStitcherSlateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VideoStitcherSlateList contains a list of VideoStitcherSlate
type VideoStitcherSlateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VideoStitcherSlate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VideoStitcherSlate{}, &VideoStitcherSlateList{})
}
