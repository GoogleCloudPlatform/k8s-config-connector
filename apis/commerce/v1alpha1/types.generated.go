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


// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.AssignmentProtocol
type AssignmentProtocol struct {
	// Allow manual assignments triggered by administrative operations only.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.AssignmentProtocol.manual_assignment_type
	ManualAssignmentType *AssignmentProtocol_ManualAssignmentType `json:"manualAssignmentType,omitempty"`

	// Allow automatic assignments triggered by data plane operations.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.AssignmentProtocol.auto_assignment_type
	AutoAssignmentType *AssignmentProtocol_AutoAssignmentType `json:"autoAssignmentType,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.AssignmentProtocol.AutoAssignmentType
type AssignmentProtocol_AutoAssignmentType struct {
	// Optional. The time to live for an inactive license. After this time has
	//  passed, the license will be automatically unassigned from the user. Must
	//  be at least 7 days, if set. If unset, the license will never expire.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.AssignmentProtocol.AutoAssignmentType.inactive_license_ttl
	InactiveLicenseTtl *string `json:"inactiveLicenseTtl,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.AssignmentProtocol.ManualAssignmentType
type AssignmentProtocol_ManualAssignmentType struct {
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LicensePool
type LicensePool struct {
	// Identifier. Format:
	//  `billingAccounts/{billing_account}/orders/{order}/licensePool`
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LicensePool.name
	Name *string `json:"name,omitempty"`

	// Required. Assignment protocol for the license pool.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LicensePool.license_assignment_protocol
	LicenseAssignmentProtocol *AssignmentProtocol `json:"licenseAssignmentProtocol,omitempty"`
}

// +kcc:proto=google.cloud.commerce.consumer.procurement.v1.LicensePool
type LicensePoolObservedState struct {
	// Output only. Licenses count that are available to be assigned.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LicensePool.available_license_count
	AvailableLicenseCount *int32 `json:"availableLicenseCount,omitempty"`

	// Output only. Total number of licenses in the pool.
	// +kcc:proto:field=google.cloud.commerce.consumer.procurement.v1.LicensePool.total_license_count
	TotalLicenseCount *int32 `json:"totalLicenseCount,omitempty"`
}
