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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityClientTLSPolicyGVK = GroupVersion.WithKind("NetworkSecurityClientTLSPolicy")

// +kcc:proto=google.cloud.networksecurity.v1beta1.CertificateProvider
// +kubebuilder:validation:XValidation:rule="has(self.grpcEndpoint) != has(self.certificateProviderInstance)",message="Exactly one of grpcEndpoint or certificateProviderInstance must be specified"
type CertificateProvider struct {
	// gRPC specific configuration to access the gRPC server to
	//  obtain the cert and private key.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.CertificateProvider.grpc_endpoint
	GrpcEndpoint *GrpcEndpoint `json:"grpcEndpoint,omitempty"`

	// The certificate provider instance specification that will be passed to
	//  the data plane, which will be used to load necessary credential
	//  information.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.CertificateProvider.certificate_provider_instance
	CertificateProviderInstance *CertificateProviderInstance `json:"certificateProviderInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.CertificateProviderInstance
type CertificateProviderInstance struct {
	// Required. Plugin instance name, used to locate and load CertificateProvider
	//  instance configuration. Set to "google_cloud_private_spiffe" to use
	//  Certificate Authority Service certificate provider instance.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.CertificateProviderInstance.plugin_instance
	// +required
	PluginInstance *string `json:"pluginInstance,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.GrpcEndpoint
type GrpcEndpoint struct {
	// Required. The target URI of the gRPC endpoint. Only UDS path is supported,
	//  and should start with "unix:".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.GrpcEndpoint.target_uri
	// +required
	TargetURI *string `json:"targetUri,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.ValidationCA
// +kubebuilder:validation:XValidation:rule="has(self.grpcEndpoint) != has(self.certificateProviderInstance)",message="Exactly one of grpcEndpoint or certificateProviderInstance must be specified"
type ValidationCA struct {
	// gRPC specific configuration to access the gRPC server to
	//  obtain the CA certificate.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ValidationCA.grpc_endpoint
	GrpcEndpoint *GrpcEndpoint `json:"grpcEndpoint,omitempty"`

	// The certificate provider instance specification that will be passed to
	//  the data plane, which will be used to load necessary credential
	//  information.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ValidationCA.certificate_provider_instance
	CertificateProviderInstance *CertificateProviderInstance `json:"certificateProviderInstance,omitempty"`
}

// NetworkSecurityClientTLSPolicySpec defines the desired state of NetworkSecurityClientTLSPolicy
// +kcc:spec:proto=google.cloud.networksecurity.v1beta1.ClientTlsPolicy
type NetworkSecurityClientTLSPolicySpec struct {
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The Project that this resource belongs to.
	// +optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location for the resource
	// +required
	Location string `json:"location"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of label tags associated with the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.labels
	// NOTEYET: this field is not in the DCL-based beta resource
	// Labels map[string]string `json:"labels,omitempty"`

	// Optional. Server Name Indication string to present to the server during TLS
	//  handshake. E.g: "secure.example.com".
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.sni
	Sni *string `json:"sni,omitempty"`

	// Optional. Defines a mechanism to provision client identity (public and private keys)
	//  for peer to peer authentication. The presence of this dictates mTLS.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.client_certificate
	ClientCertificate *CertificateProvider `json:"clientCertificate,omitempty"`

	// Required. Defines the mechanism to obtain the Certificate Authority certificate to
	//  validate the server certificate.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.server_validation_ca
	ServerValidationCA []ValidationCA `json:"serverValidationCa,omitempty"`
}

// NetworkSecurityClientTLSPolicyStatus defines the config connector machine state of NetworkSecurityClientTLSPolicy
// +kcc:status:proto=google.cloud.networksecurity.v1beta1.ClientTlsPolicy
type NetworkSecurityClientTLSPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityClientTLSPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// NOTEYET: output only fields are moved directly under "status" due to backward compatibility
	// ObservedState *NetworkSecurityClientTLSPolicyObservedState `json:"observedState,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.ClientTlsPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// NetworkSecurityClientTLSPolicyObservedState is the state of the NetworkSecurityClientTLSPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1beta1.ClientTlsPolicy
type NetworkSecurityClientTLSPolicyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityclienttlspolicy;gcpnetworksecurityclienttlspolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityClientTLSPolicy is the Schema for the NetworkSecurityClientTLSPolicy API
// +k8s:openapi-gen=true
type NetworkSecurityClientTLSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityClientTLSPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityClientTLSPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityClientTLSPolicyList contains a list of NetworkSecurityClientTLSPolicy
type NetworkSecurityClientTLSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityClientTLSPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityClientTLSPolicy{}, &NetworkSecurityClientTLSPolicyList{})
}
