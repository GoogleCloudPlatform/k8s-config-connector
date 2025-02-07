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


// +kcc:proto=google.cloud.billing.v1.ProjectBillingInfo
type ProjectBillingInfo struct {

	// The resource name of the billing account associated with the project, if
	//  any. For example, `billingAccounts/012345-567890-ABCDEF`.
	// +kcc:proto:field=google.cloud.billing.v1.ProjectBillingInfo.billing_account_name
	BillingAccountName *string `json:"billingAccountName,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.ProjectBillingInfo
type ProjectBillingInfoObservedState struct {
	// Output only. The resource name for the `ProjectBillingInfo`; has the form
	//  `projects/{project_id}/billingInfo`. For example, the resource name for the
	//  billing information for project `tokyo-rain-123` would be
	//  `projects/tokyo-rain-123/billingInfo`.
	// +kcc:proto:field=google.cloud.billing.v1.ProjectBillingInfo.name
	Name *string `json:"name,omitempty"`

	// Output only. The ID of the project that this `ProjectBillingInfo`
	//  represents, such as `tokyo-rain-123`. This is a convenience field so that
	//  you don't need to parse the `name` field to obtain a project ID.
	// +kcc:proto:field=google.cloud.billing.v1.ProjectBillingInfo.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Output only. True if the project is associated with an open billing
	//  account, to which usage on the project is charged. False if the project is
	//  associated with a closed billing account, or no billing account at all, and
	//  therefore cannot use paid services.
	// +kcc:proto:field=google.cloud.billing.v1.ProjectBillingInfo.billing_enabled
	BillingEnabled *bool `json:"billingEnabled,omitempty"`
}
