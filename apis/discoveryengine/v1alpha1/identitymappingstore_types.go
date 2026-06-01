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

var DiscoveryEngineIdentityMappingStoreGVK = GroupVersion.WithKind("DiscoveryEngineIdentityMappingStore")

// DiscoveryEngineIdentityMappingStoreSpec defines the desired state of DiscoveryEngineIdentityMappingStore
// +kcc:spec:proto=google.cloud.discoveryengine.v1.IdentityMappingStore
type DiscoveryEngineIdentityMappingStoreSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The DiscoveryEngineIdentityMappingStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Input only. The KMS key to be used to protect this Identity Mapping Store
	// at creation time.
	//
	// Must be set for requests that need to comply with CMEK Org Policy
	// protections.
	//
	// If this field is set and processed successfully, the Identity Mapping Store
	// will be protected by the KMS key, as indicated in the cmek_config field.
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// DiscoveryEngineIdentityMappingStoreStatus defines the config connector machine state of DiscoveryEngineIdentityMappingStore
type DiscoveryEngineIdentityMappingStoreStatus struct { /* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineIdentityMappingStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineIdentityMappingStoreObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineIdentityMappingStoreObservedState is the state of the DiscoveryEngineIdentityMappingStore resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.IdentityMappingStore
type DiscoveryEngineIdentityMappingStoreObservedState struct {
	// Output only. CMEK-related information for the Identity Mapping Store.
	CmekConfig *CmekConfig `json:"cmekConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineidentitymappingstore;gcpdiscoveryengineidentitymappingstores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineIdentityMappingStore is the Schema for the DiscoveryEngineIdentityMappingStore API
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
type DiscoveryEngineIdentityMappingStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineIdentityMappingStoreSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineIdentityMappingStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineIdentityMappingStoreList contains a list of DiscoveryEngineIdentityMappingStore
type DiscoveryEngineIdentityMappingStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineIdentityMappingStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineIdentityMappingStore{}, &DiscoveryEngineIdentityMappingStoreList{})
}
