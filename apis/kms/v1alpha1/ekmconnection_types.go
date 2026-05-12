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

var KMSEKMConnectionGVK = GroupVersion.WithKind("KMSEKMConnection")

// KMSEKMConnectionSpec defines the desired state of KMSEKMConnection
// +kcc:spec:proto=google.cloud.kms.v1.EkmConnection
type KMSEKMConnectionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The KMSEKMConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.service_resolvers
	ServiceResolvers []KMSEKMConnectionServiceResolver `json:"serviceResolvers,omitempty"`

	// Optional. Etag of the currently stored EkmConnection.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.key_management_mode
	KeyManagementMode *string `json:"keyManagementMode,omitempty"`

	// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if KeyManagementMode is CLOUD_KMS.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.crypto_space_path
	CryptoSpacePath *string `json:"cryptoSpacePath,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection.ServiceResolver
type KMSEKMConnectionServiceResolver struct {
	// Required. The resource name of the Service Directory service pointing to an EKM replica, in the format projects/*/locations/*/namespaces/*/services/*.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.service_directory_service
	ServiceDirectoryService *string `json:"serviceDirectoryService,omitempty"`

	// Optional. The filter applied to the endpoints of the resolved service. If no filter is specified, all endpoints will be considered. An endpoint will be chosen arbitrarily from the filtered list for each request.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.endpoint_filter
	EndpointFilter *string `json:"endpointFilter,omitempty"`

	// Required. The hostname of the EKM replica used at TLS and HTTP layers.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Required. A list of leaf server certificates used to authenticate HTTPS connections to the EKM replica. Currently, a maximum of 10 Certificate is supported.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.server_certificates
	ServerCertificates []KMSEKMConnectionCertificate `json:"serverCertificates,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.Certificate
type KMSEKMConnectionCertificate struct {
	// Required. The raw certificate bytes in DER format.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.raw_der
	RawDer []byte `json:"rawDer,omitempty"`
}

// KMSEKMConnectionStatus defines the config connector machine state of KMSEKMConnection
type KMSEKMConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the KMSEKMConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *KMSEKMConnectionObservedState `json:"observedState,omitempty"`
}

// KMSEKMConnectionObservedState is the state of the KMSEKMConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.kms.v1.EkmConnection
type KMSEKMConnectionObservedState struct {
	// Output only. The time at which the EkmConnection was created.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The parsed server certificates from the service resolvers.
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.service_resolvers
	ServiceResolvers []KMSEKMConnectionObservedServiceResolver `json:"serviceResolvers,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.EkmConnection.ServiceResolver
type KMSEKMConnectionObservedServiceResolver struct {
	// +kcc:proto:field=google.cloud.kms.v1.EkmConnection.ServiceResolver.server_certificates
	ServerCertificates []KMSEKMConnectionObservedCertificate `json:"serverCertificates,omitempty"`
}

// +kcc:proto=google.cloud.kms.v1.Certificate
type KMSEKMConnectionObservedCertificate struct {
	// Output only. True if the certificate was parsed successfully.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.parsed
	Parsed *bool `json:"parsed,omitempty"`

	// Output only. The issuer distinguished name in RFC 2253 format. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.issuer
	Issuer *string `json:"issuer,omitempty"`

	// Output only. The subject distinguished name in RFC 2253 format. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.subject
	Subject *string `json:"subject,omitempty"`

	// Output only. The subject Alternative DNS names. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.subject_alternative_dns_names
	SubjectAlternativeDnsNames []string `json:"subjectAlternativeDnsNames,omitempty"`

	// Output only. The certificate is not valid before this time. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.not_before_time
	NotBeforeTime *string `json:"notBeforeTime,omitempty"`

	// Output only. The certificate is not valid after this time. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.not_after_time
	NotAfterTime *string `json:"notAfterTime,omitempty"`

	// Output only. The certificate serial number as a hex string. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.serial_number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Output only. The SHA-256 certificate fingerprint as a hex string. Only present if parsed is true.
	// +kcc:proto:field=google.cloud.kms.v1.Certificate.sha256_fingerprint
	Sha256Fingerprint *string `json:"sha256Fingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmsekmconnection;gcpkmsekmconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSEKMConnection is the Schema for the KMSEKMConnection API
// +k8s:openapi-gen=true
type KMSEKMConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   KMSEKMConnectionSpec   `json:"spec,omitempty"`
	Status KMSEKMConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KMSEKMConnectionList contains a list of KMSEKMConnection
type KMSEKMConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSEKMConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSEKMConnection{}, &KMSEKMConnectionList{})
}
