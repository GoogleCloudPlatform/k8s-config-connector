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

var NetAppHostGroupGVK = GroupVersion.WithKind("NetAppHostGroup")

// NetAppHostGroupSpec defines the desired state of NetAppHostGroup
// +kcc:spec:proto=google.cloud.netapp.v1.HostGroup
type NetAppHostGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NetAppHostGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Type of the host group.
	// +kubebuilder:validation:Enum=TYPE_UNSPECIFIED;ISCSI_INITIATOR
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.type
	Type *string `json:"type,omitempty"`

	// Required. The list of hosts associated with the host group.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.hosts
	Hosts []string `json:"hosts,omitempty"`

	// Required. The OS type of the host group. It indicates the type of operating
	// system used by all of the hosts in the HostGroup. All hosts in a HostGroup
	// must be of the same OS type. This can be set only when creating a
	// HostGroup.
	// +kubebuilder:validation:Enum=OS_TYPE_UNSPECIFIED;LINUX;WINDOWS;ESXI
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.os_type
	OSType *string `json:"osType,omitempty"`

	// Optional. Description of the host group.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels of the host group.
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// NetAppHostGroupStatus defines the config connector machine state of NetAppHostGroup
type NetAppHostGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetAppHostGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetAppHostGroupObservedState `json:"observedState,omitempty"`
}

// NetAppHostGroupObservedState is the state of the NetAppHostGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.HostGroup
type NetAppHostGroupObservedState struct {
	// Output only. State of the host group.
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.state
	State *string `json:"state,omitempty"`

	// Output only. Create time of the host group.
	// +kcc:proto:field=google.cloud.netapp.v1.HostGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetapphostgroup;gcpnetapphostgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetAppHostGroup is the Schema for the NetAppHostGroup API
// +k8s:openapi-gen=true
type NetAppHostGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetAppHostGroupSpec   `json:"spec,omitempty"`
	Status NetAppHostGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetAppHostGroupList contains a list of NetAppHostGroup
type NetAppHostGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetAppHostGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetAppHostGroup{}, &NetAppHostGroupList{})
}
