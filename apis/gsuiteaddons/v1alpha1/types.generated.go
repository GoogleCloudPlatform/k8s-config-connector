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


// +kcc:proto=google.cloud.gsuiteaddons.v1.Authorization
type Authorization struct {
	// The canonical full name of this resource.
	//  Example:  `projects/123/authorization`
	// +kcc:proto:field=google.cloud.gsuiteaddons.v1.Authorization.name
	Name *string `json:"name,omitempty"`

	// The email address of the service account used to authenticate requests to
	//  add-on callback endpoints.
	// +kcc:proto:field=google.cloud.gsuiteaddons.v1.Authorization.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`

	// The OAuth client ID used to obtain OAuth access tokens for a user on the
	//  add-on's behalf.
	// +kcc:proto:field=google.cloud.gsuiteaddons.v1.Authorization.oauth_client_id
	OauthClientID *string `json:"oauthClientID,omitempty"`
}
