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

var DiscoveryEngineUserStoreGVK = GroupVersion.WithKind("DiscoveryEngineUserStore")

// DiscoveryEngineUserStoreSpec defines the desired state of DiscoveryEngineUserStore
// +kcc:spec:proto=google.cloud.discoveryengine.v1beta.UserStore
type DiscoveryEngineUserStoreSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The DiscoveryEngineUserStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The display name of the User Store.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The default subscription LicenseConfig for the UserStore. If
	//  enableLicenseAutoRegister is true, new users will automatically register
	//  under the default subscription.
	DefaultLicenseConfigRef *DiscoveryEngineLicenseConfigRef `json:"defaultLicenseConfigRef,omitempty"`

	// Optional. Whether to enable license auto register for users in this User
	//  Store. If true, new users will automatically register under the default
	//  license config as long as the default license config has seats left.
	EnableLicenseAutoRegister *bool `json:"enableLicenseAutoRegister,omitempty"`

	// Optional. Whether to enable license auto update for users in this User
	//  Store. If true, users with expired licenses will automatically be updated
	//  to use the default license config as long as the default license config has
	//  seats left.
	EnableExpiredLicenseAutoUpdate *bool `json:"enableExpiredLicenseAutoUpdate,omitempty"`
}

// DiscoveryEngineUserStoreStatus defines the config connector machine state of DiscoveryEngineUserStore
type DiscoveryEngineUserStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineUserStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineUserStoreObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineUserStoreObservedState is the state of the DiscoveryEngineUserStore resource as most recently observed in GCP.
// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1beta.UserStore
type DiscoveryEngineUserStoreObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineuserstore;gcpdiscoveryengineuserstores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineUserStore is the Schema for the DiscoveryEngineUserStore API
// +k8s:openapi-gen=true
type DiscoveryEngineUserStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineUserStoreSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineUserStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineUserStoreList contains a list of DiscoveryEngineUserStore
type DiscoveryEngineUserStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineUserStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineUserStore{}, &DiscoveryEngineUserStoreList{})
}
