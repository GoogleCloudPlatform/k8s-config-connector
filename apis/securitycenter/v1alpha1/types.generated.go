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


// +kcc:proto=google.cloud.securitycenter.v1.OrgPolicy
type OrgPolicy struct {
	// The resource name of the org policy.
	//  Example:
	//  "organizations/{organization_id}/policies/{constraint_name}"
	// +kcc:proto:field=google.cloud.securitycenter.v1.OrgPolicy.name
	Name *string `json:"name,omitempty"`
}
