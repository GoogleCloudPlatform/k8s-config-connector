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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SslcertificateCertificate struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *SslcertificateValueFrom `json:"valueFrom,omitempty"`
}

type SslcertificatePrivateKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *SslcertificateValueFrom `json:"valueFrom,omitempty"`
}

type SslcertificateValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *k8sv1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

// ComputeSSLCertificateSpec defines the desired state of ComputeSSLCertificate
// +kcc:spec:proto=google.cloud.compute.v1.SslCertificate
type ComputeSSLCertificateSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	/* Immutable. The certificate in PEM format.
	The certificate chain must be no greater than 5 certs long.
	The chain must include at least one intermediate cert. */
	// +required
	Certificate SslcertificateCertificate `json:"certificate"`

	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The write-only private key in PEM format. */
	// +required
	PrivateKey SslcertificatePrivateKey `json:"privateKey"`

	// The ComputeSSLCertificate name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeSSLCertificateStatus defines the config connector machine state of ComputeSSLCertificate
type ComputeSSLCertificateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeSSLCertificate resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *ComputeSSLCertificateObservedState `json:"observedState,omitempty"`

	/* The unique identifier for the resource. */
	// +optional
	CertificateId *int64 `json:"certificateId,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Expire time of the certificate in RFC3339 text format. */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeSSLCertificateObservedState is the state of the ComputeSSLCertificate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.SslCertificate
type ComputeSSLCertificateObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesslcertificate;gcpcomputesslcertificates
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeSSLCertificate is the Schema for the ComputeSSLCertificate API
// +k8s:openapi-gen=true
type ComputeSSLCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeSSLCertificateSpec   `json:"spec,omitempty"`
	Status ComputeSSLCertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeSSLCertificateList contains a list of ComputeSSLCertificate
type ComputeSSLCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeSSLCertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSSLCertificate{}, &ComputeSSLCertificateList{})
}
