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

var VideoStitcherCdnKeyGVK = GroupVersion.WithKind("VideoStitcherCdnKey")

// VideoStitcherCdnKeySpec defines the desired state of VideoStitcherCdnKey
// +kcc:spec:proto=google.cloud.video.stitcher.v1.CdnKey
type VideoStitcherCdnKeySpec struct {
	/* Immutable. The project that this resource belongs to. */
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. The location of this resource. */
	Location string `json:"location"`

	/* Immutable. Optional. The VideoStitcherCdnKey name. If not given, the metadata.name will be used. */
	ResourceID *string `json:"resourceID,omitempty"`

	// The hostname this key applies to.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.hostname
	Hostname *string `json:"hostname,omitempty"`

	// The configuration for a Google Cloud CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.google_cdn_key
	GoogleCDNKey *GoogleCDNKey `json:"googleCDNKey,omitempty"`

	// The configuration for an Akamai CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.akamai_cdn_key
	AkamaiCDNKey *AkamaiCDNKey `json:"akamaiCDNKey,omitempty"`

	// The configuration for a Media CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.media_cdn_key
	MediaCDNKey *MediaCDNKey `json:"mediaCDNKey,omitempty"`
}

// VideoStitcherCdnKeyStatus defines the config connector machine state of VideoStitcherCdnKey
type VideoStitcherCdnKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VideoStitcherCdnKey resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VideoStitcherCdnKeyObservedState `json:"observedState,omitempty"`
}

// VideoStitcherCdnKeyObservedState is the state of the VideoStitcherCdnKey resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.stitcher.v1.CdnKey
type VideoStitcherCdnKeyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvideostitchercdnkey;gcpvideostitchercdnkeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VideoStitcherCdnKey is the Schema for the VideoStitcherCdnKey API
// +k8s:openapi-gen=true
type VideoStitcherCdnKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VideoStitcherCdnKeySpec   `json:"spec,omitempty"`
	Status VideoStitcherCdnKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VideoStitcherCdnKeyList contains a list of VideoStitcherCdnKey
type VideoStitcherCdnKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VideoStitcherCdnKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VideoStitcherCdnKey{}, &VideoStitcherCdnKeyList{})
}
