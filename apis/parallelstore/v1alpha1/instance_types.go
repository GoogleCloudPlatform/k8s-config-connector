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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ParallelstoreInstanceGVK = GroupVersion.WithKind("ParallelstoreInstance")

// ParallelstoreInstanceSpec defines the desired state of ParallelstoreInstance
// +kcc:spec:proto=google.cloud.parallelstore.v1.Instance
type ParallelstoreInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ParallelstoreInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The description of the instance. 2048 characters or less.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.description
	Description *string `json:"description,omitempty"`

	// Optional. Cloud Labels are a flexible and lightweight mechanism for
	//  organizing cloud resources into groups that reflect a customer's
	//  organizational needs and deployment strategies. See
	//  https://cloud.google.com/resource-manager/docs/labels-overview for details.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. The instance's storage capacity in Gibibytes (GiB).
	//  Allowed values are between 12000 and 100000, in multiples of 4000; e.g.,
	//  12000, 16000, 20000, ...
	// +required
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.capacity_gib
	CapacityGib *int64 `json:"capacityGib"`

	// Optional. Immutable. The name of the Compute Engine
	//  [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the
	//  instance is connected.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Immutable. The ID of the IP address range being used by the
	//  instance's VPC network. See [Configure a VPC
	//  network](https://cloud.google.com/parallelstore/docs/vpc#create_and_configure_the_vpc).
	//  If no ID is provided, all ranges are considered.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`

	// Optional. Immutable. Stripe level for files. Allowed values are:
	//
	//  * `FILE_STRIPE_LEVEL_MIN`: offers the best performance for small size
	//    files.
	//  * `FILE_STRIPE_LEVEL_BALANCED`: balances performance for workloads
	//    involving a mix of small and large files.
	//  * `FILE_STRIPE_LEVEL_MAX`: higher throughput performance for larger files.
	// +kubebuilder:validation:Enum=FILE_STRIPE_LEVEL_MIN;FILE_STRIPE_LEVEL_BALANCED;FILE_STRIPE_LEVEL_MAX
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.file_stripe_level
	FileStripeLevel *string `json:"fileStripeLevel,omitempty"`

	// Optional. Immutable. Stripe level for directories. Allowed values are:
	//
	//  * `DIRECTORY_STRIPE_LEVEL_MIN`: recommended when directories contain a
	//    small number of files.
	//  * `DIRECTORY_STRIPE_LEVEL_BALANCED`: balances performance for workloads
	//    involving a mix of small and large directories.
	//  * `DIRECTORY_STRIPE_LEVEL_MAX`: recommended for directories with a large
	//    number of files.
	// +kubebuilder:validation:Enum=DIRECTORY_STRIPE_LEVEL_MIN;DIRECTORY_STRIPE_LEVEL_BALANCED;DIRECTORY_STRIPE_LEVEL_MAX
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.directory_stripe_level
	DirectoryStripeLevel *string `json:"directoryStripeLevel,omitempty"`

	// Optional. Immutable. The deployment type of the instance. Allowed values
	//  are:
	//
	//  * `SCRATCH`: the instance is a scratch instance.
	//  * `PERSISTENT`: the instance is a persistent instance.
	// +kubebuilder:validation:Enum=SCRATCH;PERSISTENT
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.deployment_type
	DeploymentType *string `json:"deploymentType,omitempty"`
}

// ParallelstoreInstanceStatus defines the config connector machine state of ParallelstoreInstance
type ParallelstoreInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ParallelstoreInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ParallelstoreInstanceObservedState `json:"observedState,omitempty"`
}

// ParallelstoreInstanceObservedState is the state of the ParallelstoreInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.parallelstore.v1.Instance
type ParallelstoreInstanceObservedState struct {
	// Output only. The instance state.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Deprecated 'daos_version' field.
	//  Output only. The version of DAOS software running in the instance.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.daos_version
	DaosVersion *string `json:"daosVersion,omitempty"`

	// Output only. A list of IPv4 addresses used for client side configuration.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.access_points
	AccessPoints []string `json:"accessPoints,omitempty"`

	// Output only. Immutable. The ID of the IP address range being used by the
	//  instance's VPC network. This field is populated by the service and contains
	//  the value currently used by the service.
	// +kcc:proto:field=google.cloud.parallelstore.v1.Instance.effective_reserved_ip_range
	EffectiveReservedIPRange *string `json:"effectiveReservedIPRange,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpparallelstoreinstance;gcpparallelstoreinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ParallelstoreInstance is the Schema for the ParallelstoreInstance API
// +k8s:openapi-gen=true
type ParallelstoreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ParallelstoreInstanceSpec   `json:"spec,omitempty"`
	Status ParallelstoreInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ParallelstoreInstanceList contains a list of ParallelstoreInstance
type ParallelstoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParallelstoreInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ParallelstoreInstance{}, &ParallelstoreInstanceList{})
}
