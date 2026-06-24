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

var NetworkSecuritySACRealmGVK = GroupVersion.WithKind("NetworkSecuritySACRealm")

// NetworkSecuritySACRealmSpec defines the desired state of NetworkSecuritySACRealm
// +kcc:spec:proto=google.cloud.networksecurity.v1.SACRealm
type NetworkSecuritySACRealmSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NetworkSecuritySACRealm name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Optional list of labels applied to the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. SSE service provider associated with the realm.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.security_service
	// +kubebuilder:validation:Enum=SECURITY_SERVICE_UNSPECIFIED;PALO_ALTO_PRISMA_ACCESS
	SecurityService *string `json:"securityService,omitempty"`
}

// NetworkSecuritySACRealmStatus defines the config connector machine state of NetworkSecuritySACRealm
type NetworkSecuritySACRealmStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecuritySACRealm resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecuritySACRealmObservedState `json:"observedState,omitempty"`
}

// NetworkSecuritySACRealmObservedState is the state of the NetworkSecuritySACRealm resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.SACRealm
type NetworkSecuritySACRealmObservedState struct {
	// Output only. Timestamp when the realm was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the realm was last updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Key to be shared with SSE service provider during pairing.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.pairing_key
	PairingKey *SACRealm_PairingKeyObservedState `json:"pairingKey,omitempty"`

	// Output only. State of the realm.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.state
	State *string `json:"state,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1.SACRealm.PairingKey
type SACRealm_PairingKeyObservedState struct {
	// Output only. Key value.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.PairingKey.key
	Key *string `json:"key,omitempty"`

	// Output only. Timestamp in UTC of when this resource is considered
	//  expired. It expires 7 days after creation.
	// +kcc:proto:field=google.cloud.networksecurity.v1.SACRealm.PairingKey.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritysacrealm;gcpnetworksecuritysacrealms
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecuritySACRealm is the Schema for the NetworkSecuritySACRealm API
// +k8s:openapi-gen=true
type NetworkSecuritySACRealm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecuritySACRealmSpec   `json:"spec,omitempty"`
	Status NetworkSecuritySACRealmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecuritySACRealmList contains a list of NetworkSecuritySACRealm
type NetworkSecuritySACRealmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecuritySACRealm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecuritySACRealm{}, &NetworkSecuritySACRealmList{})
}
