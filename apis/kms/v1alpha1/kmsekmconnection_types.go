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

// API sources for KMSEkmConnection, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/kms/v1/ekm_service.proto
// +kcc:source:service-docs=https://cloud.google.com/kms/
// +kcc:source:resource-docs=https://docs.cloud.google.com/kms/docs/reference/rest/v1/projects.locations.ekmConnections

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KMSEkmConnectionGVK = GroupVersion.WithKind("KMSEkmConnection")

// KMSEkmConnectionSpec defines the desired state of KMSEkmConnection
// +kcc:spec:proto=google.cloud.kms.v1.EkmConnection
// +kcc:required-from-proto
type KMSEkmConnectionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The KMSEkmConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A list of
	//  [ServiceResolvers][google.cloud.kms.v1.EkmConnection.ServiceResolver] where
	//  the EKM can be reached. There should be one ServiceResolver per EKM
	//  replica. Currently, only a single
	//  [ServiceResolver][google.cloud.kms.v1.EkmConnection.ServiceResolver] is
	//  supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.service_resolvers
	ServiceResolvers []EkmConnection_ServiceResolver `json:"serviceResolvers,omitempty"`

	// Optional. Etag of the currently stored
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection].
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Describes who can perform control plane operations on the EKM. If
	//  unset, this defaults to
	//  [MANUAL][google.cloud.kms.v1.EkmConnection.KeyManagementMode.MANUAL].
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.key_management_mode
	KeyManagementMode *string `json:"keyManagementMode,omitempty"`

	// Optional. Identifies the EKM Crypto Space that this
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] maps to. Note: This
	//  field is required if
	//  [KeyManagementMode][google.cloud.kms.v1.EkmConnection.KeyManagementMode] is
	//  [CLOUD_KMS][google.cloud.kms.v1.EkmConnection.KeyManagementMode.CLOUD_KMS].
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.crypto_space_path
	CryptoSpacePath *string `json:"cryptoSpacePath,omitempty"`
}

// KMSEkmConnectionStatus defines the config connector machine state of KMSEkmConnection
type KMSEkmConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the KMSEkmConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *KMSEkmConnectionObservedState `json:"observedState,omitempty"`
}

// KMSEkmConnectionObservedState is the state of the KMSEkmConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.kms.v1.EkmConnection
type KMSEkmConnectionObservedState struct {
	// Output only. The time at which the
	//  [EkmConnection][google.cloud.kms.v1.EkmConnection] was created.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. A list of
	//  [ServiceResolvers][google.cloud.kms.v1.EkmConnection.ServiceResolver] where
	//  the EKM can be reached. There should be one ServiceResolver per EKM
	//  replica. Currently, only a single
	//  [ServiceResolver][google.cloud.kms.v1.EkmConnection.ServiceResolver] is
	//  supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.service_resolvers
	ServiceResolvers []EkmConnection_ServiceResolverObservedState `json:"serviceResolvers,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmsekmconnection;gcpkmsekmconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSEkmConnection is the Schema for the KMSEkmConnection API
// +k8s:openapi-gen=true
type KMSEkmConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   KMSEkmConnectionSpec   `json:"spec,omitempty"`
	Status KMSEkmConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KMSEkmConnectionList contains a list of KMSEkmConnection
type KMSEkmConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSEkmConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSEkmConnection{}, &KMSEkmConnectionList{})
}
