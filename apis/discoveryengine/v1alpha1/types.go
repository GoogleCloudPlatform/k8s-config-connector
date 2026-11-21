// Copyright 2026 Google LLC
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

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// WorkspaceConfig is kept here to preserve backward compatibility for
// DiscoveryEngineDataStore, as upstream proto marked dasher_customer_id OUTPUT_ONLY.
// +kcc:proto=google.cloud.discoveryengine.v1.WorkspaceConfig
type WorkspaceConfig struct {
	// The type of workspace.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.type
	Type *string `json:"type,omitempty"`

	// Output only. Obfuscated Dasher customer ID.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.dasher_customer_id
	DasherCustomerID *string `json:"dasherCustomerID,omitempty"`

	// Optional. The super admin service account for the workspace.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.super_admin_service_account
	SuperAdminServiceAccount *string `json:"superAdminServiceAccount,omitempty"`

	// Optional. The super admin email address for the workspace.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.super_admin_email_address
	SuperAdminEmailAddress *string `json:"superAdminEmailAddress,omitempty"`
}

// CmekConfig is kept here to preserve backward compatibility for
// DiscoveryEngineIdentityMappingStore, as upstream proto marked kms_key_version OUTPUT_ONLY.
// +kcc:proto=google.cloud.discoveryengine.v1.CmekConfig
type CmekConfig struct {
	// Required. The name of the CmekConfig.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.name
	Name *string `json:"name,omitempty"`

	// KMS key resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// KMS key version resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key_version
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.single_region_keys
	SingleRegionKeys []SingleRegionKey `json:"singleRegionKeys,omitempty"`
}
