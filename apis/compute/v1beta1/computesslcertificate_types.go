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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSSLCertificateSpec defines the desired state of ComputeSSLCertificate
// +kcc:spec:proto=google.cloud.compute.v1.SslCertificate
type ComputeSSLCertificateSpec struct {
	/* Immutable. The certificate in PEM format.
	The certificate chain must be no greater than 5 certs long.
	The chain must include at least one intermediate cert. */
	Certificate secret.Legacy `json:"certificate"`

	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Location represents the geographical location of the ComputeSSLCertificate. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. The write-only private key in PEM format. */
	PrivateKey secret.Legacy `json:"privateKey"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeSSLCertificateStatus defines the config connector machine state of ComputeSSLCertificate
type ComputeSSLCertificateStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeSSLCertificate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The unique identifier for the resource. */
	// +optional
	CertificateId *int64 `json:"certificateId,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Expire time of the certificate in RFC3339 text format. */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesslcertificate;gcpcomputesslcertificates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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

	Spec   ComputeSSLCertificateSpec   `json:"spec"`
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
