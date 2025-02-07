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


// +kcc:proto=google.analytics.admin.v1beta.DataSharingSettings
type DataSharingSettings struct {

	// Allows Google support to access the data in order to help troubleshoot
	//  issues.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataSharingSettings.sharing_with_google_support_enabled
	SharingWithGoogleSupportEnabled *bool `json:"sharingWithGoogleSupportEnabled,omitempty"`

	// Allows Google sales teams that are assigned to the customer to access the
	//  data in order to suggest configuration changes to improve results.
	//  Sales team restrictions still apply when enabled.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataSharingSettings.sharing_with_google_assigned_sales_enabled
	SharingWithGoogleAssignedSalesEnabled *bool `json:"sharingWithGoogleAssignedSalesEnabled,omitempty"`

	// Allows any of Google sales to access the data in order to suggest
	//  configuration changes to improve results.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataSharingSettings.sharing_with_google_any_sales_enabled
	SharingWithGoogleAnySalesEnabled *bool `json:"sharingWithGoogleAnySalesEnabled,omitempty"`

	// Allows Google to use the data to improve other Google products or services.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataSharingSettings.sharing_with_google_products_enabled
	SharingWithGoogleProductsEnabled *bool `json:"sharingWithGoogleProductsEnabled,omitempty"`

	// Allows Google to share the data anonymously in aggregate form with others.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataSharingSettings.sharing_with_others_enabled
	SharingWithOthersEnabled *bool `json:"sharingWithOthersEnabled,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataSharingSettings
type DataSharingSettingsObservedState struct {
	// Output only. Resource name.
	//  Format: accounts/{account}/dataSharingSettings
	//  Example: "accounts/1000/dataSharingSettings"
	// +kcc:proto:field=google.analytics.admin.v1beta.DataSharingSettings.name
	Name *string `json:"name,omitempty"`
}
