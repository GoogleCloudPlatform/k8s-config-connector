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


// +kcc:proto=google.cloud.securitycenter.settings.v1beta1.ServiceAccount
type ServiceAccount struct {
	// The relative resource name of the service account resource.
	//  Format:
	//   * `organizations/{organization}/serviceAccount`
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ServiceAccount.name
	Name *string `json:"name,omitempty"`

	// Security Center managed service account for the organization
	//  example service-org-1234@scc.iam.gserviceaccount.com
	//  This service_account will be stored in the ComponentSettings field for the
	//  SCC, SHA, and Infra Automation components.
	// +kcc:proto:field=google.cloud.securitycenter.settings.v1beta1.ServiceAccount.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}
