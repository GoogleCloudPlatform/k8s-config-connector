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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetAppKMSConfigGVK = GroupVersion.WithKind("NetAppKMSConfig")

// NetAppKMSConfigSpec defines the desired state of NetAppKMSConfig
// +kcc:spec:proto=google.cloud.netapp.v1.KmsConfig
type NetAppKMSConfigSpec struct {
	// The NetAppKMSConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Location of the NetAppKMSConfig.
	// +required
	Location *string `json:"location,omitempty"`

	// Required. Customer managed crypto key resource full name. Format:
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{key}.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.crypto_key_name
	// +kubebuilder:validation:Required
	CryptoKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"cryptoKeyRef,omitempty"`

	// Description of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// NOT YET
	// // Labels as key value pairs
	// // +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.labels
	// Labels map[string]string `json:"labels,omitempty"`
}

// NetAppKMSConfigStatus defines the config connector machine state of NetAppKMSConfig
type NetAppKMSConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetAppKMSConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetAppKMSConfigObservedState `json:"observedState,omitempty"`
}

// NetAppKMSConfigObservedState is the state of the NetAppKMSConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.KmsConfig
type NetAppKMSConfigObservedState struct {
	// Output only. Create time of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instructions to provide the access to the customer provided
	//  encryption key.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.instructions
	Instructions *string `json:"instructions,omitempty"`

	// Output only. The Service account which will have access to the customer
	//  provided encryption key.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. State of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.state_details
	StateDetails *string `json:"stateDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappkmsconfig;gcpnetappkmsconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetAppKMSConfig is the Schema for the NetAppKMSConfig API
// +k8s:openapi-gen=true
type NetAppKMSConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetAppKMSConfigSpec   `json:"spec,omitempty"`
	Status NetAppKMSConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetAppKMSConfigList contains a list of NetAppKMSConfig
type NetAppKMSConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetAppKMSConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetAppKMSConfig{}, &NetAppKMSConfigList{})
}
