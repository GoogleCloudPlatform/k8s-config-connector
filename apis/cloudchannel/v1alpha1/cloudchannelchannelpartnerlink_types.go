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

// API sources for CloudchannelChannelPartnerLink, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/channel/v1/channel_partner_links.proto
// +kcc:source:service-docs=https://cloud.google.com/channel
// +kcc:source:resource-docs=https://docs.cloud.google.com/channel/docs/reference/rest/v1/accounts.channelPartnerLinks

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudchannelChannelPartnerLinkGVK = GroupVersion.WithKind("CloudchannelChannelPartnerLink")

// CloudchannelChannelPartnerLinkSpec defines the desired state of CloudchannelChannelPartnerLink
// +kcc:spec:proto=google.cloud.channel.v1.ChannelPartnerLink
// +kcc:required-from-proto
type CloudchannelChannelPartnerLinkSpec struct {

	// The CloudchannelChannelPartnerLink name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Cloud Identity ID of the linked reseller.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.reseller_cloud_identity_id
	// +required
	ResellerCloudIdentityID *string `json:"resellerCloudIdentityID,omitempty"`

	// Required. State of the channel partner link.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.link_state
	// +required
	LinkState *string `json:"linkState,omitempty"`
}

// CloudchannelChannelPartnerLinkStatus defines the config connector machine state of CloudchannelChannelPartnerLink
type CloudchannelChannelPartnerLinkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudchannelChannelPartnerLink resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudchannelChannelPartnerLinkObservedState `json:"observedState,omitempty"`
}

// CloudchannelChannelPartnerLinkObservedState is the state of the CloudchannelChannelPartnerLink resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.channel.v1.ChannelPartnerLink
type CloudchannelChannelPartnerLinkObservedState struct {
	// Output only. URI of the web page where partner accepts the link invitation.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.invite_link_uri
	InviteLinkURI *string `json:"inviteLinkURI,omitempty"`

	// Output only. Timestamp of when the channel partner link is created.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp of when the channel partner link is updated.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Public identifier that a customer must use to generate a
	//  transfer token to move to this distributor-reseller combination.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.public_id
	PublicID *string `json:"publicID,omitempty"`

	// Output only. Cloud Identity info of the channel partner (IR).
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.channel_partner_cloud_identity_info
	ChannelPartnerCloudIdentityInfo *CloudIdentityInfoObservedState `json:"channelPartnerCloudIdentityInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudchannelchannelpartnerlink;gcpcloudchannelchannelpartnerlinks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudchannelChannelPartnerLink is the Schema for the CloudchannelChannelPartnerLink API
// +k8s:openapi-gen=true
type CloudchannelChannelPartnerLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudchannelChannelPartnerLinkSpec   `json:"spec,omitempty"`
	Status CloudchannelChannelPartnerLinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudchannelChannelPartnerLinkList contains a list of CloudchannelChannelPartnerLink
type CloudchannelChannelPartnerLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudchannelChannelPartnerLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudchannelChannelPartnerLink{}, &CloudchannelChannelPartnerLinkList{})
}
