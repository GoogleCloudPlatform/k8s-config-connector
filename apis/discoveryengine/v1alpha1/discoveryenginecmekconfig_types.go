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

// API sources for DiscoveryEngineCMEKConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/6a1db32dc52fd9421c1743dc8f5faa7a15e518d1/google/cloud/discoveryengine/v1/cmek_config_service.proto
// +kcc:source:service-docs=https://cloud.google.com/generative-ai-app-builder/docs/
// +kcc:source:resource-docs=https://cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1/projects.locations.cmekConfigs

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineCMEKConfigGVK = GroupVersion.WithKind("DiscoveryEngineCMEKConfig")

// DiscoveryEngineCMEKConfigSpec defines the desired state of DiscoveryEngineCMEKConfig
// +kcc:spec:proto=google.cloud.discoveryengine.v1.CmekConfig
// +kcc:required-from-proto
type DiscoveryEngineCMEKConfigSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The DiscoveryEngineCMEKConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// KMS key resource name which will be used to encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key
	KMSKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.single_region_keys
	SingleRegionKeys []SingleRegionKeyWithRefs `json:"singleRegionKeys,omitempty"`
}

// DiscoveryEngineCMEKConfigStatus defines the config connector machine state of DiscoveryEngineCMEKConfig
type DiscoveryEngineCMEKConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineCMEKConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineCMEKConfigObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineCMEKConfigObservedState is the state of the DiscoveryEngineCMEKConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.CmekConfig
type DiscoveryEngineCMEKConfigObservedState struct {
	// Output only. KMS key version resource name which will be used to encrypt
	//  resources
	//  `<kms_key>/cryptoKeyVersions/{keyVersion}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key_version
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`

	// Output only. The states of the CmekConfig.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.state
	State *string `json:"state,omitempty"`

	// Output only. The default CmekConfig for the Customer.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.is_default
	IsDefault *bool `json:"isDefault,omitempty"`

	// Output only. The timestamp of the last key rotation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.last_rotation_timestamp_micros
	LastRotationTimestampMicros *int64 `json:"lastRotationTimestampMicros,omitempty"`

	// Output only. Whether the NotebookLM Corpus is ready to be used.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.notebooklm_state
	NotebooklmState *string `json:"notebooklmState,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.SingleRegionKey
type SingleRegionKey struct {
	// Required. Single-regional kms key resource name which will be used to
	//  encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SingleRegionKey.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.SingleRegionKey
type SingleRegionKeyWithRequired struct {
	// Required. Single-regional kms key resource name which will be used to
	//  encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SingleRegionKey.kms_key
	// +required
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.SingleRegionKey
type SingleRegionKeyWithRefs struct {
	// Required. Single-regional kms key resource name which will be used to
	//  encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SingleRegionKey.kms_key
	// +required
	KMSKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// CmekConfig is used by DiscoveryEngineIdentityMappingStore.
// +kcc:proto=google.cloud.discoveryengine.v1.CmekConfig
type CmekConfig struct {
	// Required. The name of the CmekConfig of the form
	//  `projects/{project}/locations/{location}/cmekConfig` or
	//  `projects/{project}/locations/{location}/cmekConfigs/{cmek_config}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.name
	Name *string `json:"name,omitempty"`

	// KMS key resource name which will be used to encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// KMS key version resource name which will be used to encrypt resources
	//  `<kms_key>/cryptoKeyVersions/{keyVersion}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key_version
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.single_region_keys
	SingleRegionKeys []SingleRegionKey `json:"singleRegionKeys,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginecmekconfig;gcpdiscoveryenginecmekconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineCMEKConfig is the Schema for the DiscoveryEngineCMEKConfig API
// +k8s:openapi-gen=true
type DiscoveryEngineCMEKConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineCMEKConfigSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineCMEKConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineCMEKConfigList contains a list of DiscoveryEngineCMEKConfig
type DiscoveryEngineCMEKConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineCMEKConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineCMEKConfig{}, &DiscoveryEngineCMEKConfigList{})
}
