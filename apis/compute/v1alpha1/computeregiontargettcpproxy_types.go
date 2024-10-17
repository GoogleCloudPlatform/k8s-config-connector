// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeRegionTargetTCPProxyGVK = GroupVersion.WithKind("ComputeRegionTargetTCPProxy")

// ComputeRegionTargetTCPProxySpec defines the desired state of ComputeRegionTargetTCPProxy
// +kcc:proto=google.cloud.compute.v1.TargetTcpProxy
type ComputeRegionTargetTCPProxySpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="BackendServiceRef is immutable"
	// Immutable. A reference to the ComputeBackendService resource.
	// +required
	BackendServiceRef *refs.ComputeBackendServiceRef `json:"backendServiceRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Description is immutable"
	// Immutable. An optional description of this resource.
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Region is immutable"
	// Immutable. The geographical location of the ComputeTargetTCPProxy.
	// Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Region *string `json:"region"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ProxyBind is immutable"
	// Immutable. This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `json:"proxyBind,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ProxyHeader is immutable"
	// Immutable. Specifies the type of proxy header to append before sending data to
	// the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The ComputeRegionTargetTCPProxy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeRegionTargetTCPProxyStatus defines the config connector machine state of ComputeRegionTargetTCPProxy
type ComputeRegionTargetTCPProxyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRegionTargetTCPProxy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRegionTargetTCPProxyObservedState `json:"observedState,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
	Kind *string `json:"kind,omitempty"`

	// The unique identifier for the resource.
	ProxyId *int64 `json:"proxyId,omitempty"`

	// The SelfLink for the resource.
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeRegionTargetTCPProxyObservedState is the state of the ComputeRegionTargetTCPProxy resource as most recently observed in GCP.
type ComputeRegionTargetTCPProxyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeregiontargettcpproxy;gcpcomputeregiontargettcpproxys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRegionTargetTCPProxy is the Schema for the ComputeRegionTargetTCPProxy API
// +k8s:openapi-gen=true
type ComputeRegionTargetTCPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRegionTargetTCPProxySpec   `json:"spec,omitempty"`
	Status ComputeRegionTargetTCPProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRegionTargetTCPProxyList contains a list of ComputeRegionTargetTCPProxy
type ComputeRegionTargetTCPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRegionTargetTCPProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRegionTargetTCPProxy{}, &ComputeRegionTargetTCPProxyList{})
}
