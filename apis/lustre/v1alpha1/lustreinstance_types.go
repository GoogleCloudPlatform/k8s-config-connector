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

// API sources for LustreInstance, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/lustre/v1/instance.proto
// +kcc:source:service-docs=https://cloud.google.com/managed-lustre/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/managed-lustre/docs/reference/rest/v1/projects.locations.instances

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LustreInstanceGVK = GroupVersion.WithKind("LustreInstance")

// LustreInstanceSpec defines the desired state of LustreInstance
// +kcc:spec:proto=google.cloud.lustre.v1.Instance
type LustreInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The LustreInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The filesystem name for this instance. This name is
	//  used by client-side tools, including when mounting the instance. Must be
	//  eight characters or less and can only contain letters and numbers.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.filesystem
	// +required
	Filesystem *string `json:"filesystem,omitempty"`

	// Required. The storage capacity of the instance in gibibytes (GiB). Allowed
	//  values are from `18000` to `954000`, in increments of 9000.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.capacity_gib
	// +required
	CapacityGib *int64 `json:"capacityGib,omitempty"`

	// Required. Immutable. The full name of the VPC network to which the instance
	//  is connected. Must be in the format
	//  `projects/{project_id}/global/networks/{network_name}`.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.network
	// +required
	Network *string `json:"network,omitempty"`

	// Optional. A user-readable description of the instance.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The throughput of the instance in MB/s/TiB.
	//  Valid values are 125, 250, 500, 1000.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.per_unit_storage_throughput
	// +required
	PerUnitStorageThroughput *int64 `json:"perUnitStorageThroughput,omitempty"`
}

// LustreInstanceStatus defines the config connector machine state of LustreInstance
type LustreInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LustreInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LustreInstanceObservedState `json:"observedState,omitempty"`
}

// LustreInstanceObservedState is the state of the LustreInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.lustre.v1.Instance
type LustreInstanceObservedState struct {
	// Output only. The state of the instance.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Mount point of the instance in the format
	//  `IP_ADDRESS@tcp:/FILESYSTEM`.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.mount_point
	MountPoint *string `json:"mountPoint,omitempty"`

	// Output only. Timestamp when the instance was created.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the instance was last updated.
	// +kcc:proto:field=google.cloud.lustre.v1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplustreinstance;gcplustreinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LustreInstance is the Schema for the LustreInstance API
// +k8s:openapi-gen=true
type LustreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LustreInstanceSpec   `json:"spec,omitempty"`
	Status LustreInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LustreInstanceList contains a list of LustreInstance
type LustreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LustreInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LustreInstance{}, &LustreInstanceList{})
}
