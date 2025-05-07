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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BatchTaskGVK = GroupVersion.WithKind("BatchTask")

// BatchTaskSpec defines the desired state of BatchTask
// +kcc:proto=google.cloud.batch.v1.Task
type BatchTaskSpec struct {
	Parent `json:",inline"`
	// The BatchTask name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// BatchTaskStatus defines the config connector machine state of BatchTask
type BatchTaskStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BatchTask resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Task name.
	//  The name is generated from the parent TaskGroup name and 'id' field.
	//  For example:
	//  "projects/123456/locations/us-west1/jobs/job01/taskGroups/group01/tasks/task01".
	// +kcc:proto:field=google.cloud.batch.v1.Task.name
	Name          *string                 `json:"name,omitempty"`
	ObservedState *BatchTaskObservedState `json:"observedState,omitempty"`
}

// BatchTaskObservedState is the state of the BatchTask resource as most recently observed in GCP.
// +kcc:proto=google.cloud.batch.v1.Task
type BatchTaskObservedState struct {
	// Task Status.
	// +kcc:proto:field=google.cloud.batch.v1.Task.status
	Status *TaskStatus `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpbatchtask;gcpbatchtasks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BatchTask is the Schema for the BatchTask API
// +k8s:openapi-gen=true
type BatchTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BatchTaskSpec   `json:"spec,omitempty"`
	Status BatchTaskStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BatchTaskList contains a list of BatchTask
type BatchTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BatchTask{}, &BatchTaskList{})
}
