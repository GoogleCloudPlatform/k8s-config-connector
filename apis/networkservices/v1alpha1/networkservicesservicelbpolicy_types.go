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

var NetworkServicesServiceLBPolicyGVK = GroupVersion.WithKind("NetworkServicesServiceLBPolicy")

// NetworkServicesServiceLBPolicySpec defines the desired state of NetworkServicesServiceLBPolicy
// +kcc:spec:proto=google.cloud.networkservices.v1.ServiceLbPolicy
type NetworkServicesServiceLBPolicySpec struct {
	// Immutable. The Project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	Location *string `json:"location"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Set of label tags associated with the ServiceLbPolicy resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024 characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. The type of load balancing algorithm to be used. The default behavior is WATERFALL_BY_REGION.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.load_balancing_algorithm
	LoadBalancingAlgorithm *string `json:"loadBalancingAlgorithm,omitempty"`

	// Optional. Configuration to automatically move traffic away for unhealthy IG/NEG for the associated Backend Service.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.auto_capacity_drain
	AutoCapacityDrain *ServiceLBPolicyAutoCapacityDrain `json:"autoCapacityDrain,omitempty"`

	// Optional. Configuration related to health based failover.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.failover_config
	FailoverConfig *ServiceLBPolicyFailoverConfig `json:"failoverConfig,omitempty"`

	// Optional. Configuration to provide isolation support for the associated Backend Service.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.isolation_config
	IsolationConfig *ServiceLBPolicyIsolationConfig `json:"isolationConfig,omitempty"`
}

// NetworkServicesServiceLBPolicyStatus defines the config connector machine state of NetworkServicesServiceLBPolicy
type NetworkServicesServiceLBPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesServiceLBPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesServiceLBPolicyObservedState `json:"observedState,omitempty"`
}

// NetworkServicesServiceLBPolicyObservedState is the state of the NetworkServicesServiceLBPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.ServiceLbPolicy
type NetworkServicesServiceLBPolicyObservedState struct {
	// Output only. The timestamp when this resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.create_time
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this resource was last updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceLbPolicy.update_time
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesservicelbpolicy;gcpnetworkservicesservicelbpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesServiceLBPolicy is the Schema for the NetworkServicesServiceLBPolicy API
// +k8s:openapi-gen=true
type NetworkServicesServiceLBPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesServiceLBPolicySpec   `json:"spec,omitempty"`
	Status NetworkServicesServiceLBPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesServiceLBPolicyList contains a list of NetworkServicesServiceLBPolicy
type NetworkServicesServiceLBPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesServiceLBPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesServiceLBPolicy{}, &NetworkServicesServiceLBPolicyList{})
}
