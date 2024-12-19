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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ComputeTargetTCPProxyGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    "ComputeTargetTCPProxy",
	}
)

// ComputeTargetTCPProxySpec defines the desired state of ComputeTargetTCPProxy
// +kcc:proto=google.cloud.compute.v1.TargetTcpProxy
type ComputeTargetTCPProxySpec struct {
	// A reference to the ComputeBackendService resource.
	// +required
	BackendServiceRef *refs.ComputeBackendServiceRef `json:"backendServiceRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Description is immutable"
	// Immutable. An optional description of this resource.
	Description *string `json:"description,omitempty"`

	// The geographical location of the ComputeTargetTCPProxy.
	// Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ProxyBind is immutable"
	// Immutable. This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `json:"proxyBind,omitempty"`

	// Specifies the type of proxy header to append before sending data to
	// the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// The ComputeTargetTCPProxy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeTargetTCPProxyStatus defines the config connector machine state of ComputeTargetTCPProxy
type ComputeTargetTCPProxyStatus struct {
	// Conditions represent the latest available observations of the object's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetTCPProxy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The unique identifier for the resource.
	ProxyId *int64 `json:"proxyId,omitempty"`

	// The SelfLink for the resource.
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeTargetTCPProxyObservedState is the state of the ComputeTargetTCPProxy resource as most recently observed in GCP.
type ComputeTargetTCPProxyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargettcpproxy;gcpcomputetargettcpproxies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetTCPProxy is the Schema for the ComputeTargetTCPProxy API
// +k8s:openapi-gen=true
type ComputeTargetTCPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetTCPProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetTCPProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetTCPProxyList contains a list of ComputeTargetTCPProxy
type ComputeTargetTCPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetTCPProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetTCPProxy{}, &ComputeTargetTCPProxyList{})
}
