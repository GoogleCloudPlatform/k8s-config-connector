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

// API sources for CloudchannelChannelPartnerRepricingConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/channel/v1/repricing.proto
// +kcc:source:service-docs=https://cloud.google.com/channel
// +kcc:source:resource-docs=https://docs.cloud.google.com/channel/docs/reference/rest/v1/accounts.channelPartnerLinks.channelPartnerRepricingConfigs

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudchannelChannelPartnerRepricingConfigGVK = GroupVersion.WithKind("CloudchannelChannelPartnerRepricingConfig")

// CloudchannelChannelPartnerRepricingConfigSpec defines the desired state of CloudchannelChannelPartnerRepricingConfig
// +kcc:spec:proto=google.cloud.channel.v1.ChannelPartnerRepricingConfig
// +kcc:required-from-proto
type CloudchannelChannelPartnerRepricingConfigSpec struct {

	// A reference to the accounts/{account}/channelPartnerLinks/{channel_partner} this resource belongs to.
	// +kcc:guess
	// ChannelPartnerLinkRef *CloudchannelChannelPartnerLinkRef `json:"channelPartnerLinkRef,omitempty"`

	// The CloudchannelChannelPartnerRepricingConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The configuration for bill modifications made by a reseller
	//  before sending it to ChannelPartner.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerRepricingConfig.repricing_config
	// +required
	RepricingConfig *RepricingConfig `json:"repricingConfig,omitempty"`
}

// CloudchannelChannelPartnerRepricingConfigStatus defines the config connector machine state of CloudchannelChannelPartnerRepricingConfig
type CloudchannelChannelPartnerRepricingConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudchannelChannelPartnerRepricingConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudchannelChannelPartnerRepricingConfigObservedState `json:"observedState,omitempty"`
}

// CloudchannelChannelPartnerRepricingConfigObservedState is the state of the CloudchannelChannelPartnerRepricingConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.channel.v1.ChannelPartnerRepricingConfig
type CloudchannelChannelPartnerRepricingConfigObservedState struct {
	// Output only. Timestamp of an update to the repricing rule. If `update_time`
	//  is after
	//  [RepricingConfig.effective_invoice_month][google.cloud.channel.v1.RepricingConfig.effective_invoice_month]
	//  then it indicates this was set mid-month.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerRepricingConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudchannelchannelpartnerrepricingconfig;gcpcloudchannelchannelpartnerrepricingconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudchannelChannelPartnerRepricingConfig is the Schema for the CloudchannelChannelPartnerRepricingConfig API
// +k8s:openapi-gen=true
type CloudchannelChannelPartnerRepricingConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudchannelChannelPartnerRepricingConfigSpec   `json:"spec,omitempty"`
	Status CloudchannelChannelPartnerRepricingConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudchannelChannelPartnerRepricingConfigList contains a list of CloudchannelChannelPartnerRepricingConfig
type CloudchannelChannelPartnerRepricingConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudchannelChannelPartnerRepricingConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudchannelChannelPartnerRepricingConfig{}, &CloudchannelChannelPartnerRepricingConfigList{})
}
