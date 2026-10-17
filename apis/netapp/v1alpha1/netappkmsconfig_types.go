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

// API sources for NetappKmsConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/netapp/v1/kms.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.kmsConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappKmsConfigGVK = GroupVersion.WithKind("NetappKmsConfig")

// NetappKmsConfigSpec defines the desired state of NetappKmsConfig
// +kcc:spec:proto=google.cloud.netapp.v1.KmsConfig
// +kcc:required-from-proto
type NetappKmsConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetappKmsConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Customer managed crypto key resource full name. Format:
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{key}.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.crypto_key_name
	// +required
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`

	// Description of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.description
	Description *string `json:"description,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// NetappKmsConfigStatus defines the config connector machine state of NetappKmsConfig
type NetappKmsConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappKmsConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappKmsConfigObservedState `json:"observedState,omitempty"`
}

// NetappKmsConfigObservedState is the state of the NetappKmsConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.KmsConfig
type NetappKmsConfigObservedState struct {
	// Output only. State of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the KmsConfig.
	// +kcc:proto:field=google.cloud.netapp.v1.KmsConfig.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

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
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappkmsconfig;gcpnetappkmsconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappKmsConfig is the Schema for the NetappKmsConfig API
// +k8s:openapi-gen=true
type NetappKmsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappKmsConfigSpec   `json:"spec,omitempty"`
	Status NetappKmsConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappKmsConfigList contains a list of NetappKmsConfig
type NetappKmsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappKmsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappKmsConfig{}, &NetappKmsConfigList{})
}
