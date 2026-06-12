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
	privatecav1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityTLSInspectionPolicyGVK = GroupVersion.WithKind("NetworkSecurityTLSInspectionPolicy")

// NetworkSecurityTLSInspectionPolicySpec defines the desired state of NetworkSecurityTLSInspectionPolicy
// +kcc:spec:proto=google.cloud.networksecurity.v1.TlsInspectionPolicy
type NetworkSecurityTLSInspectionPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`
	// The NetworkSecurityTLSInspectionPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Free-text description of the resource.
	Description *string `json:"description,omitempty"`

	// Required. A CA pool resource used to issue interception certificates.
	// The CA pool string has a relative resource path following the form
	// "projects/{project}/locations/{location}/caPools/{ca_pool}".
	// +kubebuilder:validation:Required
	CAPoolRef *privatecav1beta1.PrivateCACAPoolRef `json:"caPoolRef"`

	// Optional. A TrustConfig resource used when making a connection to the TLS
	// server. This is a relative resource path following the form
	// "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This
	// is necessary to intercept TLS connections to servers with certificates
	// signed by a private CA or self-signed certificates.
	// Note that Secure Web Proxy does not yet honor this field.
	TrustConfigRef *refsv1beta1.CertificateManagerTrustConfigRef `json:"trustConfigRef,omitempty"`

	// Optional. If  FALSE (the default), use our default set of public CAs in
	// addition to any CAs specified in trust_config. These public CAs are
	// currently based on the Mozilla Root Program and are subject to change over
	// time. If TRUE, do not accept our default set of public CAs. Only CAs
	// specified in trust_config will be accepted. This defaults to FALSE (use
	// public CAs in addition to trust_config) for backwards compatibility, but
	// trusting public root CAs is *not recommended* unless the traffic in
	// question is outbound to public web servers. When possible, prefer setting
	// this to "false" and explicitly specifying trusted CAs and certificates in a
	// TrustConfig. Note that Secure Web Proxy does not yet honor this field.
	ExcludePublicCASet *bool `json:"excludePublicCASet,omitempty"`

	// Optional. Minimum TLS version that the firewall should use when negotiating
	// connections with both clients and servers. If this is not set, then the
	// default value is to allow the broadest set of clients and servers (TLS 1.0
	// or higher). Setting this to more restrictive values may improve security,
	// but may also prevent the firewall from connecting to some clients or
	// servers.
	// Note that Secure Web Proxy does not yet honor this field.
	MinTLSVersion *string `json:"minTLSVersion,omitempty"`

	// Optional. The selected Profile. If this is not set, then the default value
	// is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE").
	// Setting this to more restrictive values may improve security, but may also
	// prevent the TLS inspection proxy from connecting to some clients or
	// servers. Note that Secure Web Proxy does not yet honor this field.
	TLSFeatureProfile *string `json:"tlsFeatureProfile,omitempty"`

	// Optional. List of custom TLS cipher suites selected.
	// This field is valid only if the selected tls_feature_profile is CUSTOM.
	// The [compute.SslPoliciesService.ListAvailableFeatures][] method returns the
	// set of features that can be specified in this list.
	// Note that Secure Web Proxy does not yet honor this field.
	CustomTLSFeatures []string `json:"customTLSFeatures,omitempty"`
}

// NetworkSecurityTLSInspectionPolicyStatus defines the config connector machine state of NetworkSecurityTLSInspectionPolicy
type NetworkSecurityTLSInspectionPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityTLSInspectionPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityTLSInspectionPolicyObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityTLSInspectionPolicyObservedState is the state of the NetworkSecurityTLSInspectionPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.TlsInspectionPolicy
type NetworkSecurityTLSInspectionPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritytlsinspectionpolicy;gcpnetworksecuritytlsinspectionpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityTLSInspectionPolicy is the Schema for the NetworkSecurityTLSInspectionPolicy API
// +k8s:openapi-gen=true
type NetworkSecurityTLSInspectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityTLSInspectionPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityTLSInspectionPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityTLSInspectionPolicyList contains a list of NetworkSecurityTLSInspectionPolicy
type NetworkSecurityTLSInspectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityTLSInspectionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityTLSInspectionPolicy{}, &NetworkSecurityTLSInspectionPolicyList{})
}
