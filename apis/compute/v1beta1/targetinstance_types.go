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

var ComputeTargetInstanceGVK = GroupVersion.WithKind("ComputeTargetInstance")

// ComputeTargetInstanceSpec defines the desired state of ComputeTargetInstance
// +kcc:spec:proto=google.cloud.compute.v1.TargetInstance
type ComputeTargetInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.description
	Description *string `json:"description,omitempty"`

	// The ComputeInstance handling traffic for this target instance.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.instance
	InstanceRef *InstanceRef `json:"instanceRef"`

	// Immutable. NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported. Default value: "NO_NAT" Possible values: ["NO_NAT"].
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.nat_policy
	NatPolicy *string `json:"natPolicy,omitempty"`

	// The network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.network
	NetworkRef *ComputeNetworkRef `json:"networkRef,omitempty"`

	// The ComputeTargetInstance name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The resource URL for the security policy associated with this target instance.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.security_policy
	SecurityPolicyRef *ComputeSecurityPolicyRef `json:"securityPolicyRef,omitempty"`

	// Immutable. URL of the zone where the target instance resides.
	// +required
	Zone string `json:"zone"`
}

// ComputeTargetInstanceStatus defines the config connector machine state of ComputeTargetInstance
// +kcc:status:proto=google.cloud.compute.v1.TargetInstance
type ComputeTargetInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Server-defined URL for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.TargetInstance.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeTargetInstanceObservedState `json:"observedState,omitempty"`
}

// ComputeTargetInstanceObservedState is the state of the ComputeTargetInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetInstance
type ComputeTargetInstanceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetinstance;gcpcomputetargetinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetInstance is the Schema for the ComputeTargetInstance API
// +k8s:openapi-gen=true
type ComputeTargetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetInstanceSpec   `json:"spec,omitempty"`
	Status ComputeTargetInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetInstanceList contains a list of ComputeTargetInstance
type ComputeTargetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetInstance{}, &ComputeTargetInstanceList{})
}
