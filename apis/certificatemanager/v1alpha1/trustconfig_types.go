// Copyright 2025 Google LLC
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
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CertificateManagerTrustConfigGVK = GroupVersion.WithKind("CertificateManagerTrustConfig")

// CertificateManagerTrustConfigSpec defines the desired state of CertificateManagerTrustConfig
// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig
type TrustConfigSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// +required
	Location string `json:"location,omitempty"`

	// NOT YET
	// // Set of labels associated with a TrustConfig.
	// // +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.description
	Description *string `json:"description,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.etag
	Etag *string `json:"etag,omitempty"`

	// Set of trust stores to perform validation against.
	//
	//  This field is supported when TrustConfig is configured with Load Balancers,
	//  currently not supported for SPIFFE certificate validation.
	//
	//  Only one TrustStore specified is currently allowed.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.trust_stores
	TrustStores []TrustConfig_TrustStore `json:"trustStores,omitempty"`
}

// CertificateManagerTrustConfigStatus defines the config connector machine state of CertificateManagerTrustConfig
type TrustConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CertificateManagerTrustConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TrustConfigObservedState `json:"observedState,omitempty"`
}

// TrustConfigObservedState is the state of the CertificateManagerTrustConfig resource as most recently observed in GCP.
// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig
type TrustConfigObservedState struct {
	// Output only. The creation timestamp of a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a TrustConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.TrustConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcertificatemanagertrustconfig;gcpcertificatemanagertrustconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CertificateManagerTrustConfig is the Schema for the CertificateManagerTrustConfig API
// +k8s:openapi-gen=true
type CertificateManagerTrustConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TrustConfigSpec   `json:"spec,omitempty"`
	Status TrustConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CertificateManagerTrustConfigList contains a list of CertificateManagerTrustConfig
type CertificateManagerTrustConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerTrustConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerTrustConfig{}, &CertificateManagerTrustConfigList{})
}
