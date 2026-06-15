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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VideoStitcherCdnKeyGVK = GroupVersion.WithKind("VideoStitcherCdnKey")

// VideoStitcherCdnKeySpec defines the desired state of VideoStitcherCdnKey
// +kcc:spec:proto=google.cloud.video.stitcher.v1.CdnKey
type VideoStitcherCdnKeySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The VideoStitcherCdnKey name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The hostname this key applies to.
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// The configuration for a Google Cloud CDN key.
	// +optional
	GoogleCDNKey *GoogleCDNKey `json:"googleCDNKey,omitempty"`

	// The configuration for an Akamai CDN key.
	// +optional
	AkamaiCDNKey *AkamaiCDNKey `json:"akamaiCDNKey,omitempty"`

	// The configuration for a Media CDN key.
	// +optional
	MediaCDNKey *MediaCDNKey `json:"mediaCDNKey,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.GoogleCdnKey
type GoogleCDNKey struct {
	// Input only. Secret for this Google Cloud CDN key.
	// +required
	PrivateKey *GoogleCDNKeyPrivateKey `json:"privateKey,omitempty"`

	// The public name of the Google Cloud CDN key.
	// +required
	KeyName *string `json:"keyName,omitempty"`
}

type GoogleCDNKeyPrivateKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *VideoStitcherCdnKeyValueFrom `json:"valueFrom,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.AkamaiCdnKey
type AkamaiCDNKey struct {
	// Input only. Token key for the Akamai CDN edge configuration.
	// +required
	TokenKey *AkamaiCDNKeyTokenKey `json:"tokenKey,omitempty"`
}

type AkamaiCDNKeyTokenKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *VideoStitcherCdnKeyValueFrom `json:"valueFrom,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.MediaCdnKey
type MediaCDNKey struct {
	// Input only. 64-byte ed25519 private key for this Media CDN key.
	// +required
	PrivateKey *MediaCDNKeyPrivateKey `json:"privateKey,omitempty"`

	// The keyset name of the Media CDN key.
	// +required
	KeyName *string `json:"keyName,omitempty"`

	// Optional. If set, the URL will be signed using the Media CDN token.
	//  Otherwise, the URL would be signed using the standard Media CDN signature.
	// +optional
	TokenConfig *MediaCDNKeyTokenConfig `json:"tokenConfig,omitempty"`
}

type MediaCDNKeyPrivateKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *VideoStitcherCdnKeyValueFrom `json:"valueFrom,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.MediaCdnKey.TokenConfig
type MediaCDNKeyTokenConfig struct {
	// Optional. The query parameter in which to find the token.
	// +optional
	QueryParameter *string `json:"queryParameter,omitempty"`
}

type VideoStitcherCdnKeyValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *k8sv1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
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
