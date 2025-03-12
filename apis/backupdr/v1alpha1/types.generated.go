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

// +kcc:proto=google.cloud.backupdr.v1.ManagementURI
type ManagementURI struct {
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURI struct {
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID
type WorkforceIdentityBasedOAuth2ClientID struct {
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURIObservedState struct {
	// Output only. First party Management URI for Google Identities.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI.first_party_management_uri
	FirstPartyManagementURI *string `json:"firstPartyManagementURI,omitempty"`

	// Output only. Third party Management URI for External Identity Providers.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI.third_party_management_uri
	ThirdPartyManagementURI *string `json:"thirdPartyManagementURI,omitempty"`
}
