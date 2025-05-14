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

package v1beta1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.devtools.cloudbuild.v1.WorkerPool
// CloudBuildWorkerPoolSpec defines the desired state of Instance
type CloudBuildWorkerPoolSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// A user-specified, human-readable name for the `WorkerPool`. If provided,
	//  this value must be 1-63 characters.
	DisplayName string `json:"displayName,omitempty"`

	// The `WorkerPool` name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +required
	Location string `json:"location"`

	// Legacy Private Pool configuration.
	// +required
	PrivatePoolConfig *PrivatePoolV1Config `json:"privatePoolV1Config,omitempty"`

	// TODO: support annotations
}

// +kcc:proto=google.devtools.cloudbuild.v1.PrivatePoolV1Config
type PrivatePoolV1Config struct {
	// Machine configuration for the workers in the pool.
	// +required
	WorkerConfig *PrivatePoolV1Config_WorkerConfig `json:"workerConfig,omitempty"`

	// Network configuration for the pool.
	NetworkConfig *PrivatePoolV1Config_NetworkConfigSpec `json:"networkConfig,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PrivatePoolV1Config.WorkerConfig
type PrivatePoolV1Config_WorkerConfig struct {
	// Machine type of a worker, such as `e2-medium`.
	//  See [Worker pool config
	//  file](https://cloud.google.com/build/docs/private-pools/worker-pool-config-file-schema).
	//  If left blank, Cloud Build will use a sensible default.
	MachineType *string `json:"machineType,omitempty"`

	// Size of the disk attached to the worker, in GB.
	//  See [Worker pool config
	//  file](https://cloud.google.com/build/docs/private-pools/worker-pool-config-file-schema).
	//  Specify a value of up to 2000. If `0` is specified, Cloud Build will use
	//  a standard disk size.
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PrivatePoolV1Config.NetworkConfig
type PrivatePoolV1Config_NetworkConfigSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Immutable. The network definition that the workers are peered
	//  to. If this section is left empty, the workers will be peered to
	//  `WorkerPool.project_id` on the service producer network.
	PeeredNetworkRef computev1beta1.ComputeNetworkRef `json:"peeredNetworkRef,omitempty"`

	// Option to configure network egress for the workers.
	EgressOption *string `json:"egressOption,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Immutable. Subnet IP range within the peered network. This is specified
	//  in CIDR notation with a slash and the subnet prefix size. You can
	//  optionally specify an IP address before the subnet prefix value. e.g.
	//  `192.168.0.0/29` would specify an IP range starting at 192.168.0.0 with a
	//  prefix size of 29 bits.
	//  `/16` would specify a prefix size of 16 bits, with an automatically
	//  determined IP within the peered VPC.
	//  If unspecified, a value of `/24` will be used.
	PeeredNetworkIPRange *string `json:"peeredNetworkIPRange,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.PrivatePoolV1Config.NetworkConfig
type PrivatePoolV1Config_NetworkConfigStatus struct {
	// Immutable. The network definition that the workers are peered
	//  to. If this section is left empty, the workers will be peered to
	//  `WorkerPool.project_id` on the service producer network.
	PeeredNetwork *string `json:"peeredNetwork,omitempty"`

	// Option to configure network egress for the workers.
	EgressOption *string `json:"egressOption,omitempty"`

	// Immutable. Subnet IP range within the peered network. This is specified
	//  in CIDR notation with a slash and the subnet prefix size. You can
	//  optionally specify an IP address before the subnet prefix value. e.g.
	//  `192.168.0.0/29` would specify an IP range starting at 192.168.0.0 with a
	//  prefix size of 29 bits.
	//  `/16` would specify a prefix size of 16 bits, with an automatically
	//  determined IP within the peered VPC.
	//  If unspecified, a value of `/24` will be used.
	PeeredNetworkIPRange *string `json:"peeredNetworkIPRange,omitempty"`
}

// CloudBuildWorkerPoolStatus defines the observed state of Instance
type CloudBuildWorkerPoolStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *CloudBuildWorkerPoolObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.WorkerPool
type CloudBuildWorkerPoolObservedState struct {
	/* The creation timestamp of the workerpool.*/
	// +optional
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* The last update timestamp of the workerpool.*/
	// +optional
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Machine configuration for the workers in the pool.
	// +required
	WorkerConfig *PrivatePoolV1Config_WorkerConfig `json:"workerConfig,omitempty"`

	// Network configuration for the pool.
	NetworkConfig *PrivatePoolV1Config_NetworkConfigStatus `json:"networkConfig,omitempty"`

	/* The Checksum computed by the server, using weak indicator.*/
	// +optional
	ETag *string `json:"etag,omitempty"`
}

// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=beta"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBuildWorkerPool is the Schema for the CloudBuild WorkerPool API
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
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
