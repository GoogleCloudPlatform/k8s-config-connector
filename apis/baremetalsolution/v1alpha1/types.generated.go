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


// +kcc:proto=google.cloud.baremetalsolution.v2.InstanceQuota
type InstanceQuota struct {

	// Instance type.
	//  Deprecated: use gcp_service.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceQuota.instance_type
	InstanceType *string `json:"instanceType,omitempty"`

	// The gcp service of the provisioning quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceQuota.gcp_service
	GcpService *string `json:"gcpService,omitempty"`

	// Location where the quota applies.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceQuota.location
	Location *string `json:"location,omitempty"`

	// Number of machines than can be created for the given location and
	//  instance_type.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceQuota.available_machine_count
	AvailableMachineCount *int32 `json:"availableMachineCount,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.ProvisioningQuota
type ProvisioningQuota struct {

	// The asset type of this provisioning quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.asset_type
	AssetType *string `json:"assetType,omitempty"`

	// The gcp service of the provisioning quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.gcp_service
	GcpService *string `json:"gcpService,omitempty"`

	// The specific location of the provisioining quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.location
	Location *string `json:"location,omitempty"`

	// The available count of the provisioning quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.available_count
	AvailableCount *int32 `json:"availableCount,omitempty"`

	// Instance quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.instance_quota
	InstanceQuota *InstanceQuota `json:"instanceQuota,omitempty"`

	// Server count.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.server_count
	ServerCount *int64 `json:"serverCount,omitempty"`

	// Network bandwidth, Gbps
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.network_bandwidth
	NetworkBandwidth *int64 `json:"networkBandwidth,omitempty"`

	// Storage size (GB).
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.storage_gib
	StorageGib *int64 `json:"storageGib,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.InstanceQuota
type InstanceQuotaObservedState struct {
	// Output only. The name of the instance quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.InstanceQuota.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.ProvisioningQuota
type ProvisioningQuotaObservedState struct {
	// Output only. The name of the provisioning quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.name
	Name *string `json:"name,omitempty"`

	// Instance quota.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.ProvisioningQuota.instance_quota
	InstanceQuota *InstanceQuotaObservedState `json:"instanceQuota,omitempty"`
}
