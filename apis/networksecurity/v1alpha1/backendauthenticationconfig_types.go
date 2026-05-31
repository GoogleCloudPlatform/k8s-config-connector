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

var NetworkSecurityBackendAuthenticationConfigGVK = GroupVersion.WithKind("NetworkSecurityBackendAuthenticationConfig")

// NetworkSecurityBackendAuthenticationConfigSpec defines the desired state of NetworkSecurityBackendAuthenticationConfig
// +kcc:spec:proto=google.cloud.networksecurity.v1.BackendAuthenticationConfig
type NetworkSecurityBackendAuthenticationConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location,omitempty"`

	// The NetworkSecurityBackendAuthenticationConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.description
	Description *string `json:"description,omitempty"`

	// Set of label tags associated with the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A reference to a certificatemanager.googleapis.com.Certificate
	//  resource. This is a relative resource path following the form
	//  "projects/{project}/locations/{location}/certificates/{certificate}".
	//
	//  Used by a BackendService to negotiate mTLS when the backend connection uses
	//  TLS and the backend requests a client certificate. Must have a CLIENT_AUTH
	//  scope.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.client_certificate
	ClientCertificateRef *refsv1beta1.CertificateManagerCertificateRef `json:"clientCertificateRef,omitempty"`

	// Optional. A reference to a TrustConfig resource from the
	//  certificatemanager.googleapis.com namespace. This is a relative resource
	//  path following the form
	//  "projects/{project}/locations/{location}/trustConfigs/{trust_config}".
	//
	//  A BackendService uses the chain of trust represented by this TrustConfig,
	//  if specified, to validate the server certificates presented by the backend.
	//  Required unless wellKnownRoots is set to PUBLIC_ROOTS.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.trust_config
	TrustConfigRef *refsv1beta1.CertificateManagerTrustConfigRef `json:"trustConfigRef,omitempty"`

	// Well known roots to use for server certificate validation.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.well_known_roots
	// +kubebuilder:validation:Enum=WELL_KNOWN_ROOTS_UNSPECIFIED;NONE;PUBLIC_ROOTS
	WellKnownRoots *string `json:"wellKnownRoots,omitempty"`
}

// NetworkSecurityBackendAuthenticationConfigStatus defines the config connector machine state of NetworkSecurityBackendAuthenticationConfig
type NetworkSecurityBackendAuthenticationConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityBackendAuthenticationConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityBackendAuthenticationConfigObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityBackendAuthenticationConfigObservedState is the state of the NetworkSecurityBackendAuthenticationConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.BackendAuthenticationConfig
type NetworkSecurityBackendAuthenticationConfigObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Etag of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.BackendAuthenticationConfig.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritybackendauthenticationconfig;gcpnetworksecuritybackendauthenticationconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityBackendAuthenticationConfig is the Schema for the NetworkSecurityBackendAuthenticationConfig API
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
type NetworkSecurityBackendAuthenticationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityBackendAuthenticationConfigSpec   `json:"spec,omitempty"`
	Status NetworkSecurityBackendAuthenticationConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityBackendAuthenticationConfigList contains a list of NetworkSecurityBackendAuthenticationConfig
type NetworkSecurityBackendAuthenticationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityBackendAuthenticationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityBackendAuthenticationConfig{}, &NetworkSecurityBackendAuthenticationConfigList{})
}
