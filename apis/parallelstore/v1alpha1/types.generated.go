// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.parallelstore.v1beta.Instance
type Instance struct {
	// Identifier. The resource name of the instance, in the format
	//  `projects/{project}/locations/{location}/instances/{instance_id}`.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.name
	Name *string `json:"name,omitempty"`

	// Optional. The description of the instance. 2048 characters or less.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.description
	Description *string `json:"description,omitempty"`

	// Optional. Cloud Labels are a flexible and lightweight mechanism for
	//  organizing cloud resources into groups that reflect a customer's
	//  organizational needs and deployment strategies. See
	//  https://cloud.google.com/resource-manager/docs/labels-overview for details.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Immutable. The instance's storage capacity in Gibibytes (GiB).
	//  Allowed values are between 12000 and 100000, in multiples of 4000; e.g.,
	//  12000, 16000, 20000, ...
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.capacity_gib
	CapacityGib *int64 `json:"capacityGib,omitempty"`

	// Optional. Immutable. The name of the Compute Engine
	//  [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the
	//  instance is connected.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.network
	Network *string `json:"network,omitempty"`

	// Optional. Immutable. The ID of the IP address range being used by the
	//  instance's VPC network. See [Configure a VPC
	//  network](https://cloud.google.com/parallelstore/docs/vpc#create_and_configure_the_vpc).
	//  If no ID is provided, all ranges are considered.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`

	// Optional. Stripe level for files. Allowed values are:
	//
	//  * `FILE_STRIPE_LEVEL_MIN`: offers the best performance for small size
	//    files.
	//  * `FILE_STRIPE_LEVEL_BALANCED`: balances performance for workloads
	//    involving a mix of small and large files.
	//  * `FILE_STRIPE_LEVEL_MAX`: higher throughput performance for larger files.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.file_stripe_level
	FileStripeLevel *string `json:"fileStripeLevel,omitempty"`

	// Optional. Stripe level for directories. Allowed values are:
	//
	//  * `DIRECTORY_STRIPE_LEVEL_MIN`: recommended when directories contain a
	//    small number of files.
	//  * `DIRECTORY_STRIPE_LEVEL_BALANCED`: balances performance for workloads
	//    involving a mix of small and large directories.
	//  * `DIRECTORY_STRIPE_LEVEL_MAX`: recommended for directories with a large
	//    number of files.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.directory_stripe_level
	DirectoryStripeLevel *string `json:"directoryStripeLevel,omitempty"`

	// Optional. The deployment type of the instance. Allowed values are:
	//
	//  * `SCRATCH`: the instance is a scratch instance.
	//  * `PERSISTENT`: the instance is a persistent instance.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.deployment_type
	DeploymentType *string `json:"deploymentType,omitempty"`
}

// +kcc:proto=google.cloud.parallelstore.v1beta.Instance
type InstanceObservedState struct {
	// Output only. The instance state.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The version of DAOS software running in the instance.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.daos_version
	DaosVersion *string `json:"daosVersion,omitempty"`

	// Output only. A list of IPv4 addresses used for client side configuration.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.access_points
	AccessPoints []string `json:"accessPoints,omitempty"`

	// Output only. Immutable. The ID of the IP address range being used by the
	//  instance's VPC network. This field is populated by the service and contains
	//  the value currently used by the service.
	// +kcc:proto:field=google.cloud.parallelstore.v1beta.Instance.effective_reserved_ip_range
	EffectiveReservedIPRange *string `json:"effectiveReservedIPRange,omitempty"`
}
