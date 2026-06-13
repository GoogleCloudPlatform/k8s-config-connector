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

// ComputeSSLPolicySpec defines the desired state of ComputeSSLPolicy
// +kcc:spec:proto=google.cloud.compute.v1.SslPolicy
type ComputeSSLPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeSSLPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.custom_features
	CustomFeatures []string `json:"customFeatures,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.description
	Description *string `json:"description,omitempty"`

	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	// Check the MinTlsVersion enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.min_tls_version
	// +kubebuilder:validation:Enum=TLS_1_0;TLS_1_1;TLS_1_2
	MinTlsVersion *string `json:"minTlsVersion,omitempty"`

	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	// Check the Profile enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.profile
	// +kubebuilder:validation:Enum=COMPATIBLE;MODERN;RESTRICTED;CUSTOM
	Profile *string `json:"profile,omitempty"`
}

// ComputeSSLPolicyStatus defines the config connector machine state of ComputeSSLPolicy
type ComputeSSLPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeSSLPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The list of features enabled in the SSL policy.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.enabled_features
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeSSLPolicyObservedState `json:"observedState,omitempty"`
}

// ComputeSSLPolicyObservedState is the state of the ComputeSSLPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.SslPolicy
type ComputeSSLPolicyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputesslpolicy;gcpcomputesslpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeSSLPolicy is the Schema for the ComputeSSLPolicy API
// +k8s:openapi-gen=true
type ComputeSSLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeSSLPolicySpec   `json:"spec,omitempty"`
	Status ComputeSSLPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeSSLPolicyList contains a list of ComputeSSLPolicy
type ComputeSSLPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeSSLPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSSLPolicy{}, &ComputeSSLPolicyList{})
}
