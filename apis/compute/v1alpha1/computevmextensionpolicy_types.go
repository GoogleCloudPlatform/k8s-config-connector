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

// API sources for ComputeVMExtensionPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/zoneVmExtensionPolicies

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeVMExtensionPolicyGVK = GroupVersion.WithKind("ComputeVMExtensionPolicy")

// ComputeVMExtensionPolicySpec defines the desired state of ComputeVMExtensionPolicy
// +kcc:spec:proto=google.cloud.compute.v1.VmExtensionPolicy
// +kcc:required-from-proto
type ComputeVMExtensionPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Location represents the geographical location of the ComputeVMExtensionPolicy. Specify the zone name.
	Location string `json:"location"`

	// The ComputeVMExtensionPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.description
	Description *string `json:"description,omitempty"`

	// Required. A map of extension names (for example, "ops-agent") to their corresponding
	//  policy configurations.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.extension_policies
	ExtensionPolicies map[string]VMExtensionPolicyExtensionPolicy `json:"extensionPolicies,omitempty"`

	// Optional. Selectors to target VMs for this policy.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.instance_selectors
	InstanceSelectors []VMExtensionPolicyInstanceSelector `json:"instanceSelectors,omitempty"`

	// Optional. Priority of this policy.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.priority
	Priority *int32 `json:"priority,omitempty"`
}

// ComputeVMExtensionPolicyStatus defines the config connector machine state of ComputeVMExtensionPolicy
type ComputeVMExtensionPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeVMExtensionPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeVMExtensionPolicyObservedState `json:"observedState,omitempty"`
}

// ComputeVMExtensionPolicyObservedState is the state of the ComputeVMExtensionPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.VmExtensionPolicy
type ComputeVMExtensionPolicyObservedState struct {
	// Output only. [Output Only] Creation timestamp inRFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Optional. Output only. [Output Only] Link to the global policy that manages this zone policy, if
	//  applicable.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.global_resource_link
	GlobalResourceLink *string `json:"globalResourceLink,omitempty"`

	// Output only. [Output Only] The unique identifier for the resource. This identifier is
	//  defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.id
	ID *uint64 `json:"id,omitempty"`

	// Output only. [Output Only] Type of the resource. Alwayscompute#vmExtensionPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.kind
	Kind *string `json:"kind,omitempty"`

	// Optional. Output only. [Output Only] Indicates if this policy is managed by a global policy.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.managed_by_global
	ManagedByGlobal *bool `json:"managedByGlobal,omitempty"`

	// Output only. [Output Only] Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. [Output Only] Server-defined URL for this resource's resource id.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// Optional. Output only. [Output Only] Current state of the policy: ACTIVE or DELETING.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.state
	State *string `json:"state,omitempty"`

	// Output only. [Output Only] Update timestamp inRFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.VmExtensionPolicy.update_timestamp
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputevmextensionpolicy;gcpcomputevmextensionpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeVMExtensionPolicy is the Schema for the ComputeVMExtensionPolicy API
// +k8s:openapi-gen=true
type ComputeVMExtensionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeVMExtensionPolicySpec   `json:"spec,omitempty"`
	Status ComputeVMExtensionPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeVMExtensionPolicyList contains a list of ComputeVMExtensionPolicy
type ComputeVMExtensionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeVMExtensionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeVMExtensionPolicy{}, &ComputeVMExtensionPolicyList{})
}
