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


// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.EkmMetadata
type EkmMetadata struct {
	// The Cloud EKM partner.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.EkmMetadata.ekm_solution
	EkmSolution *string `json:"ekmSolution,omitempty"`

	// Endpoint for sending requests to the EKM for key provisioning during
	//  Assured Workload creation.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.EkmMetadata.ekm_endpoint_uri
	EkmEndpointURI *string `json:"ekmEndpointURI,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.Partner
type Partner struct {
	// Identifier. The resource name of the partner.
	//  Format: `organizations/{organization}/locations/{location}/partner`
	//  Example: "organizations/123456/locations/us-central1/partner"
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.name
	Name *string `json:"name,omitempty"`

	// List of SKUs the partner is offering
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.skus
	Skus []Sku `json:"skus,omitempty"`

	// List of Google Cloud supported EKM partners supported by the partner
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.ekm_solutions
	EkmSolutions []EkmMetadata `json:"ekmSolutions,omitempty"`

	// List of Google Cloud regions that the partner sells services to customers.
	//  Valid Google Cloud regions found here:
	//  https://cloud.google.com/compute/docs/regions-zones
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.operated_cloud_regions
	OperatedCloudRegions []string `json:"operatedCloudRegions,omitempty"`

	// Google Cloud project ID in the partner's Google Cloud organization for
	//  receiving enhanced Logs for Partners.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.partner_project_id
	PartnerProjectID *string `json:"partnerProjectID,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.Sku
type Sku struct {
	// Argentum product SKU, that is associated with the partner offerings to
	//  customers used by Syntro for billing purposes. SKUs can represent resold
	//  Google products or support services.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Sku.id
	ID *string `json:"id,omitempty"`

	// Display name of the product identified by the SKU. A partner may want to
	//  show partner branded names for their offerings such as local sovereign
	//  cloud solutions.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Sku.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.Partner
type PartnerObservedState struct {
	// Output only. Time the resource was created
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the resource was updated
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Partner.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
