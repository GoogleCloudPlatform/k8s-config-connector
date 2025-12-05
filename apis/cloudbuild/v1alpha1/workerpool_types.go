/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudBuildWorkerPoolSpec defines the desired state of Instance
type CloudBuildWorkerPoolSpec struct {
	Name        string  `json:"name,omitempty"`
	DisplayName string  `json:"displayName,omitempty"`
	ResourceID  *string `json:"resourceID,omitempty"`
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +required
	Location string `json:"location"`
	// +required
	PrivatePoolConfig *PrivatePoolV1Config `json:"privatePoolV1Config,omitempty"`
}

type PrivatePoolV1Config struct {
	// +required
	WorkerConfig *WorkerConfig `json:"workerConfig,omitempty"`
	// +optional
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`
}

type WorkerConfig struct {
	// +optional
	MachineType string `json:"machineType,omitempty"`
	// +optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`
}

type NetworkConfig struct {
	// +required
	PeeredNetworkRef computev1beta1.ComputeNetworkRef `json:"peeredNetworkRef,omitempty"`
	// +optional
	EgressOption string `json:"egressOption,omitempty"`
	// +optional
	PeeredNetworkIPRange string `json:"peeredNetworkIPRange,omitempty"`
}

// CloudBuildWorkerPoolStatus defines the observed state of Instance
type CloudBuildWorkerPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *CloudBuildWorkerPoolObservedState `json:"observedState,omitempty"`
}

type CloudBuildWorkerPoolObservedState struct {
	/* The creation timestamp of the workerpool.*/
	// +optional
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* The last update timestamp of the workerpool.*/
	// +optional
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`
	// +optional
	WorkerConfig  *WorkerConfig       `json:"workerConfig,omitempty"`
	NetworkConfig *NetworkConfigState `json:"networkConfig,omitempty"`
}

type NetworkConfigState struct {
	// +optional
	PeeredNetwork string `json:"peeredNetwork,omitempty"`
	// +optional
	EgressOption string `json:"egressOption,omitempty"`
	// +optional
	PeeredNetworkIPRange string `json:"peeredNetworkIPRange,omitempty"`
}

// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=beta"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBuildWorkerPool is the Schema for the CloudBuild WorkerPool API
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildworkerpool;gcpcloudbuildworkerpools
type CloudBuildWorkerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudBuildWorkerPoolSpec   `json:"spec,omitempty"`
	Status CloudBuildWorkerPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBuildWorkerPoolList contains a list of WorkerPool
type CloudBuildWorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildWorkerPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildWorkerPool{}, &CloudBuildWorkerPoolList{})
}
