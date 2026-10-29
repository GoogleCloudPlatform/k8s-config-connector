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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeAddonsConfigGVK = GroupVersion.WithKind("ApigeeAddonsConfig")

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AdvancedApiOpsConfig
type GoogleCloudApigeeV1AdvancedApiOpsConfig struct {
	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AdvancedApiOpsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiSecurityConfig
type GoogleCloudApigeeV1ApiSecurityConfig struct {
	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiSecurityConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ApiSecurityConfig.expires_at
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ConnectorsPlatformConfig
type GoogleCloudApigeeV1ConnectorsPlatformConfig struct {
	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ConnectorsPlatformConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1ConnectorsPlatformConfig.expires_at
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1IntegrationConfig
type GoogleCloudApigeeV1IntegrationConfig struct {
	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1IntegrationConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1MonetizationConfig
type GoogleCloudApigeeV1MonetizationConfig struct {
	// Flag that specifies whether the Advanced API Ops add-on is enabled.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1MonetizationConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig
type GoogleCloudApigeeV1AddonsConfig struct {
	// Configuration for the Monetization add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.advanced_api_ops_config
	AdvancedApiOpsConfig *GoogleCloudApigeeV1AdvancedApiOpsConfig `json:"advancedApiOpsConfig,omitempty"`

	// Configuration for the Monetization add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.api_security_config
	ApiSecurityConfig *GoogleCloudApigeeV1ApiSecurityConfig `json:"apiSecurityConfig,omitempty"`

	// Configuration for the Monetization add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.connectors_platform_config
	ConnectorsPlatformConfig *GoogleCloudApigeeV1ConnectorsPlatformConfig `json:"connectorsPlatformConfig,omitempty"`

	// Configuration for the Monetization add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.integration_config
	IntegrationConfig *GoogleCloudApigeeV1IntegrationConfig `json:"integrationConfig,omitempty"`

	// Configuration for the Monetization add-on.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1AddonsConfig.monetization_config
	MonetizationConfig *GoogleCloudApigeeV1MonetizationConfig `json:"monetizationConfig,omitempty"`
}

type ApigeeAddonsConfigSpec struct {
	// Addon configurations of the Apigee organization.
	AddonsConfig *GoogleCloudApigeeV1AddonsConfig `json:"addonsConfig,omitempty"`

	// Immutable. Name of the Apigee organization.
	// +required
	Org string `json:"org"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`
}

type ApigeeAddonsConfigStatus struct {
	/* Conditions represent the latest available observation of the resource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	   If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeaddonsconfig;gcpapigeeaddonsconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeAddonsConfig is the Schema for the apigee API
// +k8s:openapi-gen=true
type ApigeeAddonsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeAddonsConfigSpec   `json:"spec"`
	Status ApigeeAddonsConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApigeeAddonsConfigList contains a list of ApigeeAddonsConfig
type ApigeeAddonsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeAddonsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeAddonsConfig{}, &ApigeeAddonsConfigList{})
}
