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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FirebaseHostingChannelGVK = GroupVersion.WithKind("FirebaseHostingChannel")

type FirebaseHostingChannelSpec struct {
	/* The time at which the channel will be automatically deleted. If null, the channel
	will not be automatically deleted. This field is present in the output whether it's
	set directly or via the 'ttl' field. */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* Immutable. Optional. The channelId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The number of previous releases to retain on the channel for rollback or other
	purposes. Must be a number between 1-100. Defaults to 10 for new channels. */
	// +optional
	RetainedReleaseCount *int `json:"retainedReleaseCount,omitempty"`

	/* Immutable. Required. The ID of the site in which to create this channel. */
	// +required
	SiteID string `json:"siteId"`

	/* Immutable. Input only. A time-to-live for this channel. Sets 'expire_time' to the provided
	duration past the time of the request. A duration in seconds with up to nine fractional
	digits, terminated by 's'. Example: "86400s" (one day). */
	// +optional
	TTL *string `json:"ttl,omitempty"`
}

type FirebaseHostingChannelStatus struct {
	/* Conditions represent the latest available observation of the
	   resource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* A unique specifier for the FirebaseHostingChannel resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* The fully-qualified resource name for the channel, in the format:
	sites/SITE_ID/channels/CHANNEL_ID. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfirebasehostingchannel;gcpfirebasehostingchannels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FirebaseHostingChannel is the Schema for the firebasehosting API
// +k8s:openapi-gen=true
type FirebaseHostingChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FirebaseHostingChannelSpec   `json:"spec"`
	Status FirebaseHostingChannelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirebaseHostingChannelList contains a list of FirebaseHostingChannel
type FirebaseHostingChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirebaseHostingChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirebaseHostingChannel{}, &FirebaseHostingChannelList{})
}
