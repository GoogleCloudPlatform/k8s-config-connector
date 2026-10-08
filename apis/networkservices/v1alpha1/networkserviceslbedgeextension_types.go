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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesLBEdgeExtensionGVK = GroupVersion.WithKind("NetworkServicesLBEdgeExtension")

// NetworkServicesLBEdgeExtensionSpec defines the desired state of NetworkServicesLBEdgeExtension
// +kcc:spec:proto=google.cloud.networkservices.v1.LbEdgeExtension
type NetworkServicesLBEdgeExtensionSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The NetworkServicesLBEdgeExtension name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of labels associated with the LbEdgeExtension resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// Required. A list of references to the forwarding rules to which this service extension is attached. At least one forwarding rule is required. Only one LbEdgeExtension resource can be associated with a forwarding rule.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.forwarding_rules
	// +required
	ForwardingRuleRefs []*computev1beta1.ForwardingRuleRef `json:"forwardingRuleRefs,omitempty"`

	// Required. A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.extension_chains
	// +required
	ExtensionChains []ExtensionChain `json:"extensionChains,omitempty"`

	// Required. All forwarding rules referenced by this extension must share the same load balancing scheme. Supported values: EXTERNAL_MANAGED.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.load_balancing_scheme
	// +required
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`
}

// NetworkServicesLBEdgeExtensionStatus defines the config connector machine state of NetworkServicesLBEdgeExtension
type NetworkServicesLBEdgeExtensionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesLBEdgeExtension resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesLBEdgeExtensionObservedState `json:"observedState,omitempty"`
}

// NetworkServicesLBEdgeExtensionObservedState is the state of the NetworkServicesLBEdgeExtension resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.LbEdgeExtension
type NetworkServicesLBEdgeExtensionObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbEdgeExtension.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkserviceslbedgeextension;gcpnetworkserviceslbedgeextensions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesLBEdgeExtension is the Schema for the NetworkServicesLBEdgeExtension API
// +k8s:openapi-gen=true
type NetworkServicesLBEdgeExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesLBEdgeExtensionSpec   `json:"spec,omitempty"`
	Status NetworkServicesLBEdgeExtensionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesLBEdgeExtensionList contains a list of NetworkServicesLBEdgeExtension
type NetworkServicesLBEdgeExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesLBEdgeExtension `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesLBEdgeExtension{}, &NetworkServicesLBEdgeExtensionList{})
}
