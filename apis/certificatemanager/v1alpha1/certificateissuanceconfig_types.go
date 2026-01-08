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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CertificateManagerCertificateIssuanceConfigGVK = GroupVersion.WithKind("CertificateManagerCertificateIssuanceConfig")

// CertificateManagerCertificateIssuanceConfigSpec defines the desired state of CertificateManagerCertificateIssuanceConfig
// +kcc:spec:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig
type CertificateManagerCertificateIssuanceConfigSpec struct {
	// Required. Defines the parent path of the resource.
	ParentRef *parent.ProjectAndLocationRef `json:",inline"`

	// Set of labels associated with a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.description
	Description *string `json:"description,omitempty"`

	// Required. The CA that issues the workload certificate. It includes the CA
	//  address, type, authentication to CA service, etc.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.certificate_authority_config
	CertificateAuthorityConfig *CertificateIssuanceConfig_CertificateAuthorityConfig `json:"certificateAuthorityConfigRef,omitempty"`

	// Required. Workload certificate lifetime requested.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.lifetime
	Lifetime *string `json:"lifetime,omitempty"`

	// Required. Specifies the percentage of elapsed time of the certificate
	//  lifetime to wait before renewing the certificate. Must be a number between
	//  1-99, inclusive.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.rotation_window_percentage
	RotationWindowPercentage *int32 `json:"rotationWindowPercentage,omitempty"`

	// Required. The key algorithm to use when generating the private key.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.key_algorithm
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty"`

	// The CertificateManagerCertificateIssuanceConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// CertificateManagerCertificateIssuanceConfigStatus defines the config connector machine state of CertificateManagerCertificateIssuanceConfig
type CertificateManagerCertificateIssuanceConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CertificateManagerCertificateIssuanceConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CertificateManagerCertificateIssuanceConfigObservedState `json:"observedState,omitempty"`
}

// CertificateManagerCertificateIssuanceConfigObservedState is the state of the CertificateManagerCertificateIssuanceConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig
type CertificateManagerCertificateIssuanceConfigObservedState struct {
	// Output only. The creation timestamp of a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a CertificateIssuanceConfig.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcertificatemanagercertificateissuanceconfig;gcpcertificatemanagercertificateissuanceconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CertificateManagerCertificateIssuanceConfig is the Schema for the CertificateManagerCertificateIssuanceConfig API
// +k8s:openapi-gen=true
type CertificateManagerCertificateIssuanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CertificateManagerCertificateIssuanceConfigSpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateIssuanceConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CertificateManagerCertificateIssuanceConfigList contains a list of CertificateManagerCertificateIssuanceConfig
type CertificateManagerCertificateIssuanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerCertificateIssuanceConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerCertificateIssuanceConfig{}, &CertificateManagerCertificateIssuanceConfigList{})
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.CertificateAuthorityServiceConfig
type CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig struct {
	// Required. A CA pool resource used to issue a certificate.
	//  The CA pool string has a relative resource path following the form
	//  "projects/{project}/locations/{location}/caPools/{ca_pool}".
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.CertificateAuthorityServiceConfig.ca_pool
	CAPoolRef *refsv1beta1.PrivateCACAPoolRef `json:"caPoolRef,omitempty"`
}
