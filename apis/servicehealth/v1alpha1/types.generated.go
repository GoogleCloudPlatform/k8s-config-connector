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


// +kcc:proto=google.cloud.servicehealth.v1.Asset
type Asset struct {
}

// +kcc:proto=google.cloud.servicehealth.v1.OrganizationImpact
type OrganizationImpact struct {
}

// +kcc:proto=google.cloud.servicehealth.v1.Asset
type AssetObservedState struct {
	// Output only. Full name of the resource as defined in
	//  [Resource
	//  Names](https://cloud.google.com/apis/design/resource_names#full_resource_name).
	// +kcc:proto:field=google.cloud.servicehealth.v1.Asset.asset_name
	AssetName *string `json:"assetName,omitempty"`

	// Output only. Type of the asset. Example:
	//  `"cloudresourcemanager.googleapis.com/Project"`
	// +kcc:proto:field=google.cloud.servicehealth.v1.Asset.asset_type
	AssetType *string `json:"assetType,omitempty"`
}

// +kcc:proto=google.cloud.servicehealth.v1.OrganizationImpact
type OrganizationImpactObservedState struct {
	// Output only. Identifier. Unique name of the organization impact in this
	//  scope including organization and location using the form
	//  `organizations/{organization_id}/locations/{location}/organizationImpacts/{organization_impact_id}`.
	//
	//  `organization_id` - ID (number) of the organization that contains the
	//  event. To get your `organization_id`, see
	//  [Getting your organization resource
	//  ID](https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id).<br>
	//  `organization_impact_id` - ID of the [OrganizationImpact
	//  resource](/service-health/docs/reference/rest/v1beta/organizations.locations.organizationImpacts#OrganizationImpact).
	// +kcc:proto:field=google.cloud.servicehealth.v1.OrganizationImpact.name
	Name *string `json:"name,omitempty"`

	// Output only. A list of event names impacting the asset.
	// +kcc:proto:field=google.cloud.servicehealth.v1.OrganizationImpact.events
	Events []string `json:"events,omitempty"`

	// Output only. Google Cloud asset possibly impacted by the specified events.
	// +kcc:proto:field=google.cloud.servicehealth.v1.OrganizationImpact.asset
	Asset *Asset `json:"asset,omitempty"`

	// Output only. The time when the affected project was last modified.
	// +kcc:proto:field=google.cloud.servicehealth.v1.OrganizationImpact.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
