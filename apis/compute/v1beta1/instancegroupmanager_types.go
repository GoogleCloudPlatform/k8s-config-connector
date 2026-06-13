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

var ComputeInstanceGroupManagerGVK = GroupVersion.WithKind("ComputeInstanceGroupManager")

// ComputeInstanceGroupManagerSpec defines the desired state of ComputeInstanceGroupManager
// +kcc:spec:proto=google.cloud.compute.v1.InstanceGroupManager
type ComputeInstanceGroupManagerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeInstanceGroupManager name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Specifies configuration that overrides the instance template configuration for the group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.all_instances_config
	AllInstancesConfig *InstanceGroupManagerAllInstancesConfig `json:"allInstancesConfig,omitempty"`

	// The autohealing policy for this managed instance group. You can specify only one value.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.auto_healing_policies
	AutoHealingPolicies []InstanceGroupManagerAutoHealingPolicy `json:"autoHealingPolicies,omitempty"`

	// The base instance name to use for instances in this group. The value must be 1-58 characters long.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.base_instance_name
	BaseInstanceName *string `json:"baseInstanceName,omitempty"`

	// An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.description
	Description *string `json:"description,omitempty"`

	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.distribution_policy
	DistributionPolicy *DistributionPolicy `json:"distributionPolicy,omitempty"`

	// Legacy field, not supported in proto.
	// +optional
	FailoverAction *string `json:"failoverAction,omitempty"`

	// The URL of the instance template that is specified for this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.instance_template
	InstanceTemplateRef *InstanceResourceRef `json:"instanceTemplateRef,omitempty"`

	// The repair policy for this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.instance_lifecycle_policy
	InstanceLifecyclePolicy *InstanceGroupManagerInstanceLifecyclePolicy `json:"instanceLifecyclePolicy,omitempty"`

	// Legacy field, not supported in proto.
	// +optional
	ServiceAccountRef *InstanceResourceRef `json:"serviceAccountRef,omitempty"`

	// Stateful configuration for this Instanced Group Manager
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.stateful_policy
	StatefulPolicy *StatefulPolicy `json:"statefulPolicy,omitempty"`

	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.target_pools
	TargetPools []InstanceResourceRef `json:"targetPools,omitempty"`

	// The target number of running instances for this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.target_size
	TargetSize *int32 `json:"targetSize,omitempty"`

	// The update policy for this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.update_policy
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `json:"updatePolicy,omitempty"`

	// Specifies the instance templates used by this managed instance group to create instances.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.versions
	Versions []InstanceGroupManagerVersion `json:"versions,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InstanceGroupManagerAutoHealingPolicy
type InstanceGroupManagerAutoHealingPolicy struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerAutoHealingPolicy.health_check
	HealthCheckRef *InstanceResourceRef `json:"healthCheckRef,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerAutoHealingPolicy.initial_delay_sec
	InitialDelaySec *int32 `json:"initialDelaySec,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InstanceGroupManagerVersion
type InstanceGroupManagerVersion struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerVersion.instance_template
	InstanceTemplateRef *InstanceResourceRef `json:"instanceTemplateRef,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerVersion.name
	Name *string `json:"name,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerVersion.target_size
	TargetSize *FixedOrPercent `json:"targetSize,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.StatefulPolicy
type StatefulPolicy struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.StatefulPolicy.preserved_state
	PreservedState *StatefulPolicyPreservedState `json:"preservedState,omitempty"`
}

type StatefulPolicyPreservedState struct {
	// +optional
	Disks map[string]StatefulPolicyPreservedStateDiskDevice `json:"disks,omitempty"`

	// +optional
	ExternalIPs map[string]StatefulPolicyPreservedStateNetworkIP `json:"externalIps,omitempty"`

	// +optional
	InternalIPs map[string]StatefulPolicyPreservedStateNetworkIP `json:"internalIps,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.StatefulPolicyPreservedStateDiskDevice
type StatefulPolicyPreservedStateDiskDevice struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.StatefulPolicyPreservedStateDiskDevice.auto_delete
	AutoDelete *string `json:"autoDelete,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.StatefulPolicyPreservedStateNetworkIp
type StatefulPolicyPreservedStateNetworkIP struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.StatefulPolicyPreservedStateNetworkIp.auto_delete
	AutoDelete *string `json:"autoDelete,omitempty"`
}

// ComputeInstanceGroupManagerStatus defines the config connector machine state of ComputeInstanceGroupManager
// +kcc:status:proto=google.cloud.compute.v1.InstanceGroupManager
type ComputeInstanceGroupManagerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstanceGroupManager resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The creation timestamp for this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The list of instance actions and the number of instances in this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.current_actions
	CurrentActions *InstanceGroupManagerActionsSummary `json:"currentActions,omitempty"`

	// Fingerprint of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// A unique identifier for this resource type. The server generates this identifier.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.id
	Id *int64 `json:"id,omitempty"`

	// The URL of the Instance Group resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.instance_group
	InstanceGroup *string `json:"instanceGroup,omitempty"`

	// The URL of the region where the managed instance group resides.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.region
	Region *string `json:"region,omitempty"`

	// The URL for this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// The status of this managed instance group.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.status
	Status *InstanceGroupManagerStatus `json:"status,omitempty"`

	// The URL of a zone where the managed instance group is located.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManager.zone
	Zone *string `json:"zone,omitempty"`
}

// ComputeInstanceGroupManagerObservedState is the state of the ComputeInstanceGroupManager resource as most recently observed in GCP.
type ComputeInstanceGroupManagerObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancegroupmanager;gcpcomputeinstancegroupmanagers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceGroupManager is the Schema for the ComputeInstanceGroupManager API
// +k8s:openapi-gen=true
type ComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceGroupManagerList contains a list of ComputeInstanceGroupManager
type ComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceGroupManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceGroupManager{}, &ComputeInstanceGroupManagerList{})
}
