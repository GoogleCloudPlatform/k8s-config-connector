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

var DiscoveryEngineCMEKConfigGVK = GroupVersion.WithKind("DiscoveryEngineCMEKConfig")

// +kcc:proto=google.cloud.discoveryengine.v1.CmekConfig
type CMEKConfig struct {
	// Required. The name of the CmekConfig of the form
	//  `projects/{project}/locations/{location}/cmekConfig` or
	//  `projects/{project}/locations/{location}/cmekConfigs/{cmek_config}`.
	Name *string `json:"name,omitempty"`

	// KMS key resource name which will be used to encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	KMSKey *string `json:"kmsKey,omitempty"`

	// KMS key version resource name which will be used to encrypt resources
	//  `<kms_key>/cryptoKeyVersions/{keyVersion}`.
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	SingleRegionKeys []SingleRegionKeyObservedState `json:"singleRegionKeys,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.SingleRegionKey
type SingleRegionKey struct {
	// Optional. Single-regional kms key resource name which will be used to encrypt resources.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SingleRegionKey.kms_key
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.SingleRegionKey
type SingleRegionKeyObservedState struct {
	// Optional. Single-regional kms key resource name which will be used to encrypt resources.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SingleRegionKey.kms_key
	KmsKey *string `json:"kmsKey,omitempty"`
}

// DiscoveryEngineCMEKConfigSpec defines the desired state of DiscoveryEngineCMEKConfig
// +kcc:spec:proto=google.cloud.discoveryengine.v1.CmekConfig
type DiscoveryEngineCMEKConfigSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// KMS key resource name which will be used to encrypt resources.
	// +required
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	SingleRegionKeys []SingleRegionKey `json:"singleRegionKeys,omitempty"`
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
	// Output only. The states of the CmekConfig.
	State *string `json:"state,omitempty"`

	// Output only. The default CmekConfig for the Customer.
	IsDefault *bool `json:"isDefault,omitempty"`

	// Output only. The timestamp of the last key rotation.
	LastRotationTimestampMicros *int64 `json:"lastRotationTimestampMicros,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	SingleRegionKeys []SingleRegionKeyObservedState `json:"singleRegionKeys,omitempty"`

	// Output only. Whether the NotebookLM Corpus is ready to be used.
	NotebooklmState *string `json:"notebooklmState,omitempty"`

	// KMS key version resource name which will be used to encrypt resources.
	KmsKeyVersion *string `json:"kmsKeyVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginecmekconfig;gcpdiscoveryenginecmekconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
