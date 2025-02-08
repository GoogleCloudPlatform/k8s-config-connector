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


// +kcc:proto=google.cloud.kms.inventory.v1.ProtectedResource
type ProtectedResource struct {
	// The full resource name of the resource.
	//  Example:
	//  `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.name
	Name *string `json:"name,omitempty"`

	// Format: `projects/{PROJECT_NUMBER}`.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.project
	Project *string `json:"project,omitempty"`

	// The ID of the project that owns the resource.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// The Cloud product that owns the resource.
	//  Example: `compute`
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.cloud_product
	CloudProduct *string `json:"cloudProduct,omitempty"`

	// Example: `compute.googleapis.com/Disk`
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Location can be `global`, regional like `us-east1`, or zonal like
	//  `us-west1-b`.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.location
	Location *string `json:"location,omitempty"`

	// A key-value pair of the resource's labels (v1) to their values.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The name of the Cloud KMS
	//  [CryptoKeyVersion](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions?hl=en)
	//  used to protect this resource via CMEK. This field is empty if the
	//  Google Cloud product owning the resource does not provide key version data
	//  to Asset Inventory. If there are multiple key versions protecting the
	//  resource, then this is same value as the first element of
	//  crypto_key_versions.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.crypto_key_version
	CryptoKeyVersion *string `json:"cryptoKeyVersion,omitempty"`

	// The names of the Cloud KMS
	//  [CryptoKeyVersion](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions?hl=en)
	//  used to protect this resource via CMEK. This field is empty if the
	//  Google Cloud product owning the resource does not provide key versions data
	//  to Asset Inventory. The first element of this field is stored in
	//  crypto_key_version.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.crypto_key_versions
	CryptoKeyVersions []string `json:"cryptoKeyVersions,omitempty"`
}

// +kcc:proto=google.cloud.kms.inventory.v1.ProtectedResource
type ProtectedResourceObservedState struct {
	// Output only. The time at which this resource was created. The granularity
	//  is in seconds. Timestamp.nanos will always be 0.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResource.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
