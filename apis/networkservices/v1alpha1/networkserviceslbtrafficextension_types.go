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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesLBTrafficExtensionGVK = GroupVersion.WithKind("NetworkServicesLBTrafficExtension")

// NetworkServicesLBTrafficExtensionSpec defines the desired state of NetworkServicesLBTrafficExtension
// +kcc:spec:proto=google.cloud.networkservices.v1.LbTrafficExtension
type NetworkServicesLBTrafficExtensionSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The NetworkServicesLBTrafficExtension name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of labels associated with the LbTrafficExtension resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A list of references to the forwarding rules to which this service extension is attached. At least one forwarding rule is required. Only one LBTrafficExtension resource can be associated with a forwarding rule.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.forwarding_rules
	// +required
	ForwardingRuleRefs []*computev1beta1.ForwardingRuleRef `json:"forwardingRuleRefs,omitempty"`

	// Required. A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.extension_chains
	// +required
	ExtensionChains []ExtensionChain `json:"extensionChains,omitempty"`

	// Required. All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. Supported values: INTERNAL_MANAGED, EXTERNAL_MANAGED.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.load_balancing_scheme
	// +required
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// Optional. The metadata provided here is included as part of the metadata_context (of type google.protobuf.Struct) in the ProcessingRequest message sent to the extension server.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.metadata
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`
}

// NetworkServicesLBTrafficExtensionStatus defines the config connector machine state of NetworkServicesLBTrafficExtension
type NetworkServicesLBTrafficExtensionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesLBTrafficExtension resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesLBTrafficExtensionObservedState `json:"observedState,omitempty"`
}

// NetworkServicesLBTrafficExtensionObservedState is the state of the NetworkServicesLBTrafficExtension resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.LbTrafficExtension
type NetworkServicesLBTrafficExtensionObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkserviceslbtrafficextension;gcpnetworkserviceslbtrafficextensions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesLBTrafficExtension is the Schema for the NetworkServicesLBTrafficExtension API
// +k8s:openapi-gen=true
type NetworkServicesLBTrafficExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesLBTrafficExtensionSpec   `json:"spec,omitempty"`
	Status NetworkServicesLBTrafficExtensionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesLBTrafficExtensionList contains a list of NetworkServicesLBTrafficExtension
type NetworkServicesLBTrafficExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesLBTrafficExtension `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesLBTrafficExtension{}, &NetworkServicesLBTrafficExtensionList{})
}
