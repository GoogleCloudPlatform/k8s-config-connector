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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
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
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.description
	Description *string `json:"description,omitempty"`

	// Required. A CA pool resource used to issue interception certificates.
	// The CA pool string has a relative resource path following the form
	// "projects/{project}/locations/{location}/caPools/{ca_pool}".
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.ca_pool
	CaPoolRef *privatecarefs.PrivateCACAPoolRef `json:"caPoolRef"`

	// Optional. A TrustConfig resource that contains the trust store for validation of client credentials.
	// The TrustConfig resource must be in the format "projects/{project}/locations/{location}/trustConfigs/{trustConfig}".
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.trust_config
	TrustConfigRef *refsv1beta1.CertificateManagerTrustConfigRef `json:"trustConfigRef,omitempty"`

	// Optional. If True, then the TrustConfig will exclude public CA certificates.
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.exclude_public_ca_set
	ExcludePublicCASet *bool `json:"excludePublicCASet,omitempty"`

	// Optional. Minimum TLS version that is allowed.
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.min_tls_version
	// +kubebuilder:validation:Enum=TLS_VERSION_UNSPECIFIED;TLS_1_0;TLS_1_1;TLS_1_2;TLS_1_3
	MinTLSVersion *string `json:"minTLSVersion,omitempty"`

	// Optional. TLS feature profile that is allowed.
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.tls_feature_profile
	// +kubebuilder:validation:Enum=PROFILE_UNSPECIFIED;PROFILE_COMPATIBLE;PROFILE_MODERN;PROFILE_RESTRICTED;PROFILE_CUSTOM
	TLSFeatureProfile *string `json:"tlsFeatureProfile,omitempty"`

	// Optional. Custom TLS features that are allowed.
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.custom_tls_features
	CustomTLSFeatures []string `json:"customTLSFeatures,omitempty"`
}

// NetworkSecurityTLSInspectionPolicyObservedState defines the observed state of NetworkSecurityTLSInspectionPolicy
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.TlsInspectionPolicy
type NetworkSecurityTLSInspectionPolicyObservedState struct {
	// Output only. The time when the TLS inspection policy was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the TLS inspection policy was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.TlsInspectionPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// NetworkSecurityTLSInspectionPolicyStatus defines the observed state of NetworkSecurityTLSInspectionPolicy
type NetworkSecurityTLSInspectionPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkSecurityTLSInspectionPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the NetworkSecurityTLSInspectionPolicy resource in GCP. */
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState contains the state of the resource as most recently observed in GCP. */
	ObservedState *NetworkSecurityTLSInspectionPolicyObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritytlsinspectionpolicy;gcpnetworksecuritytlsinspectionpolicies
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityTLSInspectionPolicy is the Schema for the NetworkSecurityTLSInspectionPolicy API
type NetworkSecurityTLSInspectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSecurityTLSInspectionPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityTLSInspectionPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkSecurityTLSInspectionPolicyList contains a list of NetworkSecurityTLSInspectionPolicy
type NetworkSecurityTLSInspectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityTLSInspectionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityTLSInspectionPolicy{}, &NetworkSecurityTLSInspectionPolicyList{})
}
