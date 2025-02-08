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


// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestamps struct {
	// Creation timestamp of the resource within the given system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Timestamp of the last modification of the resource or its metadata within
	//  a given system.
	//
	//  Note: Depending on the source system, not every modification updates this
	//  timestamp.
	//  For example, BigQuery timestamps every metadata modification but not data
	//  or permission changes.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Taxonomy
type Taxonomy struct {
	// Identifier. Resource name of this taxonomy in URL format.
	//
	//  Note: Policy tag manager generates unique taxonomy IDs.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.name
	Name *string `json:"name,omitempty"`

	// Required. User-defined name of this taxonomy.
	//
	//  The name can't start or end with spaces, must contain only Unicode letters,
	//  numbers, underscores, dashes, and spaces, and be at most 200 bytes long
	//  when encoded in UTF-8.
	//
	//  The taxonomy display name must be unique within an organization.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of this taxonomy. If not set, defaults to empty.
	//
	//  The description must contain only Unicode characters, tabs, newlines,
	//  carriage returns, and page breaks, and be at most 2000 bytes long when
	//  encoded in UTF-8.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.description
	Description *string `json:"description,omitempty"`

	// Optional. A list of policy types that are activated for this taxonomy. If
	//  not set, defaults to an empty list.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.activated_policy_types
	ActivatedPolicyTypes []string `json:"activatedPolicyTypes,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Taxonomy.Service
type Taxonomy_Service struct {
	// The Google Cloud service name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.Service.name
	Name *string `json:"name,omitempty"`

	// The service agent for the service.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.Service.identity
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.SystemTimestamps
type SystemTimestampsObservedState struct {
	// Output only. Expiration timestamp of the resource within the given system.
	//
	//  Currently only applicable to BigQuery resources.
	// +kcc:proto:field=google.cloud.datacatalog.v1.SystemTimestamps.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Taxonomy
type TaxonomyObservedState struct {
	// Output only. Number of policy tags in this taxonomy.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.policy_tag_count
	PolicyTagCount *int32 `json:"policyTagCount,omitempty"`

	// Output only. Creation and modification timestamps of this taxonomy.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.taxonomy_timestamps
	TaxonomyTimestamps *SystemTimestamps `json:"taxonomyTimestamps,omitempty"`

	// Output only. Identity of the service which owns the Taxonomy. This field is
	//  only populated when the taxonomy is created by a Google Cloud service.
	//  Currently only 'DATAPLEX' is supported.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.service
	Service *Taxonomy_Service `json:"service,omitempty"`
}
