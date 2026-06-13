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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeTargetGRPCProxyGVK = GroupVersion.WithKind("ComputeTargetGRPCProxy")

// ComputeTargetGRPCProxySpec defines the desired state of ComputeTargetGRPCProxy
// +kcc:spec:proto=google.cloud.compute.v1.TargetGrpcProxy
type ComputeTargetGRPCProxySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputeTargetGRPCProxy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.description
	Description *string `json:"description,omitempty"`

	// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.url_map
	URLMapRef *computev1beta1.ComputeURLMapRef `json:"urlMapRef"`

	// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.validate_for_proxyless
	ValidateForProxyless *bool `json:"validateForProxyless,omitempty"`
}

// ComputeTargetGRPCProxyStatus defines the config connector machine state of ComputeTargetGRPCProxy
type ComputeTargetGRPCProxyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetGRPCProxy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeTargetGRPCProxyObservedState `json:"observedState,omitempty"`
}

// ComputeTargetGRPCProxyObservedState is the state of the ComputeTargetGRPCProxy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetGrpcProxy
type ComputeTargetGRPCProxyObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#targetGrpcProxy for target grpc proxies.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL with id for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object.
	// +kcc:proto:field=google.cloud.compute.v1.TargetGrpcProxy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetgrpcproxy;gcpcomputetargetgrpcproxys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetGRPCProxy is the Schema for the ComputeTargetGRPCProxy API
// +k8s:openapi-gen=true
type ComputeTargetGRPCProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetGRPCProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetGRPCProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetGRPCProxyList contains a list of ComputeTargetGRPCProxy
type ComputeTargetGRPCProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetGRPCProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetGRPCProxy{}, &ComputeTargetGRPCProxyList{})
}
