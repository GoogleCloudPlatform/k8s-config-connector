// Copyright 2025 Google LLC
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

var NotebooksEnvironmentGVK = GroupVersion.WithKind("NotebooksEnvironment")

// NotebooksEnvironmentSpec defines the desired state of NotebooksEnvironment
// +kcc:proto=google.cloud.notebooks.v1.Environment
type NotebooksEnvironmentSpec struct {
	// The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location for the resource.
	// +required
	Location string `json:"location"`

	// The NotebooksEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Display name of this environment for the UI.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A brief description of this environment.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.description
	Description *string `json:"description,omitempty"`

	// Use a Compute Engine VM image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.vm_image
	VmImage *VmImage `json:"vmImage,omitempty"`

	// Use a container image to start the notebook instance.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.container_image
	ContainerImage *ContainerImage `json:"containerImage,omitempty"`

	// Path to a Bash script that automatically runs after a notebook instance
	//  fully boots up. The path must be a URL or
	//  Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.post_startup_script
	PostStartupScript *string `json:"postStartupScript,omitempty"`
}

// NotebooksEnvironmentStatus defines the config connector machine state of NotebooksEnvironment
type NotebooksEnvironmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NotebooksEnvironment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NotebooksEnvironmentObservedState `json:"observedState,omitempty"`
}

// NotebooksEnvironmentObservedState is the state of the NotebooksEnvironment resource as most recently observed in GCP.
// +kcc:proto=google.cloud.notebooks.v1.Environment
type NotebooksEnvironmentObservedState struct {
	// Output only. Name of this environment.
	//  Format:
	//  `projects/{project_id}/locations/{location}/environments/{environment_id}`
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.name
	// NOTYET: same as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The time at which this environment was created.
	// +kcc:proto:field=google.cloud.notebooks.v1.Environment.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnotebooksenvironment;gcpnotebooksenvironments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NotebooksEnvironment is the Schema for the NotebooksEnvironment API
// +k8s:openapi-gen=true
type NotebooksEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NotebooksEnvironmentSpec   `json:"spec,omitempty"`
	Status NotebooksEnvironmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotebooksEnvironmentList contains a list of NotebooksEnvironment
type NotebooksEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebooksEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebooksEnvironment{}, &NotebooksEnvironmentList{})
}
