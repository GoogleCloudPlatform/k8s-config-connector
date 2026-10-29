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

// API sources for NotebooksRuntime, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/notebooks/v1/runtime.proto
// +kcc:source:service-docs=https://cloud.google.com/notebooks/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/workbench/reference/rest/v1/projects.locations.runtimes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NotebooksRuntimeGVK = GroupVersion.WithKind("NotebooksRuntime")

// NotebooksRuntimeSpec defines the desired state of NotebooksRuntime
// +kcc:spec:proto=google.cloud.notebooks.v1.Runtime
// +kcc:required-from-proto
type NotebooksRuntimeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NotebooksRuntime name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.virtual_machine
	VirtualMachine *VirtualMachine `json:"virtualMachine,omitempty"`

	// The config settings for accessing runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.access_config
	AccessConfig *RuntimeAccessConfig `json:"accessConfig,omitempty"`

	// The config settings for software inside the runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.software_config
	SoftwareConfig *RuntimeSoftwareConfig `json:"softwareConfig,omitempty"`
}

// NotebooksRuntimeStatus defines the config connector machine state of NotebooksRuntime
type NotebooksRuntimeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NotebooksRuntime resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NotebooksRuntimeObservedState `json:"observedState,omitempty"`
}

// NotebooksRuntimeObservedState is the state of the NotebooksRuntime resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.notebooks.v1.Runtime
type NotebooksRuntimeObservedState struct {
	// Use a Compute Engine VM image to start the managed notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.virtual_machine
	VirtualMachine *VirtualMachineObservedState `json:"virtualMachine,omitempty"`

	// Output only. Runtime state.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.state
	State *string `json:"state,omitempty"`

	// Output only. Runtime health_state.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.health_state
	HealthState *string `json:"healthState,omitempty"`

	// The config settings for accessing runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.access_config
	AccessConfig *RuntimeAccessConfigObservedState `json:"accessConfig,omitempty"`

	// The config settings for software inside the runtime.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.software_config
	SoftwareConfig *RuntimeSoftwareConfigObservedState `json:"softwareConfig,omitempty"`

	// Output only. Contains Runtime daemon metrics such as Service status and JupyterLab
	//  stats.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.metrics
	Metrics *RuntimeMetricsObservedState `json:"metrics,omitempty"`

	// Output only. Runtime creation time.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Runtime update time.
	// +kcc:proto:field=google.cloud.notebooks.v1.Runtime.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnotebooksruntime;gcpnotebooksruntimes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NotebooksRuntime is the Schema for the NotebooksRuntime API
// +k8s:openapi-gen=true
type NotebooksRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NotebooksRuntimeSpec   `json:"spec,omitempty"`
	Status NotebooksRuntimeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotebooksRuntimeList contains a list of NotebooksRuntime
type NotebooksRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebooksRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebooksRuntime{}, &NotebooksRuntimeList{})
}
