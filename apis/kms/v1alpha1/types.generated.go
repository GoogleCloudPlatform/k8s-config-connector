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


// +kcc:proto=google.cloud.kms.inventory.v1.ProtectedResourcesSummary
type ProtectedResourcesSummary struct {
	// The full name of the ProtectedResourcesSummary resource.
	//  Example:
	//  projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/test-key/protectedResourcesSummary
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResourcesSummary.name
	Name *string `json:"name,omitempty"`

	// The total number of protected resources in the same Cloud organization as
	//  the key.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResourcesSummary.resource_count
	ResourceCount *int64 `json:"resourceCount,omitempty"`

	// The number of distinct Cloud projects in the same Cloud organization as the
	//  key that have resources protected by the key.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResourcesSummary.project_count
	ProjectCount *int32 `json:"projectCount,omitempty"`

	// The number of resources protected by the key grouped by resource type.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResourcesSummary.resource_types
	ResourceTypes map[string]int64 `json:"resourceTypes,omitempty"`

	// The number of resources protected by the key grouped by Cloud product.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResourcesSummary.cloud_products
	CloudProducts map[string]int64 `json:"cloudProducts,omitempty"`

	// The number of resources protected by the key grouped by region.
	// +kcc:proto:field=google.cloud.kms.inventory.v1.ProtectedResourcesSummary.locations
	Locations map[string]int64 `json:"locations,omitempty"`
}
