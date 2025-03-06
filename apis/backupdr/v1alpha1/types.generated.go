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

// +kcc:proto=google.cloud.backupdr.v1.ManagementServer
type ManagementServer struct {

	// Optional. The description of the ManagementServer instance (2048 characters
	//  or less).
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.description
	Description *string `json:"description,omitempty"`

	// Optional. Resource labels to represent user provided metadata.
	//  Labels currently defined:
	//  1. migrate_from_go=<false|true>
	//     If set to true, the MS is created in migration ready mode.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The type of the ManagementServer resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.type
	Type *string `json:"type,omitempty"`

	// Optional. VPC networks to which the ManagementServer instance is connected.
	//  For this version, only a single network is supported. This field is
	//  optional if MS is created without PSA
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.networks
	Networks []NetworkConfig `json:"networks,omitempty"`

	// Optional. Server specified ETag for the ManagementServer resource to
	//  prevent simultaneous updates from overwiting each other.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ManagementURI
type ManagementURI struct {
}

// +kcc:proto=google.cloud.backupdr.v1.NetworkConfig
type NetworkConfig struct {
	// Optional. The resource name of the Google Compute Engine VPC network to
	//  which the ManagementServer instance is connected.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkConfig.network
	Network *string `json:"network,omitempty"`

	// Optional. The network connect mode of the ManagementServer instance. For
	//  this version, only PRIVATE_SERVICE_ACCESS is supported.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkConfig.peering_mode
	PeeringMode *string `json:"peeringMode,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURI struct {
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID
type WorkforceIdentityBasedOAuth2ClientID struct {
}

// +kcc:proto=google.cloud.backupdr.v1.ManagementServer
type ManagementServerObservedState struct {
	// Output only. Identifier. The resource name.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The hostname or ip address of the exposed AGM endpoints, used
	//  by clients to connect to AGM/RD graphical user interface and APIs.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.management_uri
	ManagementURI *ManagementURI `json:"managementURI,omitempty"`

	// Output only. The hostnames of the exposed AGM endpoints for both types of
	//  user i.e. 1p and 3p, used to connect AGM/RM UI.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.workforce_identity_based_management_uri
	WorkforceIdentityBasedManagementURI *WorkforceIdentityBasedManagementURI `json:"workforceIdentityBasedManagementURI,omitempty"`

	// Output only. The ManagementServer state.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.state
	State *string `json:"state,omitempty"`

	// Output only. The OAuth 2.0 client id is required to make API calls to the
	//  BackupDR instance API of this ManagementServer. This is the value that
	//  should be provided in the 'aud' field of the OIDC ID Token (see openid
	//  specification
	//  https://openid.net/specs/openid-connect-core-1_0.html#IDToken).
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.oauth2_client_id
	OAUTH2ClientID *string `json:"oauth2ClientID,omitempty"`

	// Output only. The OAuth client IDs for both types of user i.e. 1p and 3p.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.workforce_identity_based_oauth2_client_id
	WorkforceIdentityBasedOAUTH2ClientID *WorkforceIdentityBasedOAuth2ClientID `json:"workforceIdentityBasedOAUTH2ClientID,omitempty"`

	// Output only. The hostname or ip address of the exposed AGM endpoints, used
	//  by BAs to connect to BA proxy.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.ba_proxy_uri
	BaProxyURI []string `json:"baProxyURI,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ManagementURI
type ManagementURIObservedState struct {
	// Output only. The ManagementServer AGM/RD WebUI URL.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementURI.web_ui
	WebUi *string `json:"webUi,omitempty"`

	// Output only. The ManagementServer AGM/RD API URL.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementURI.api
	API *string `json:"api,omitempty"`
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

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID
type WorkforceIdentityBasedOAuth2ClientIDObservedState struct {
	// Output only. First party OAuth Client ID for Google Identities.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID.first_party_oauth2_client_id
	FirstPartyOAUTH2ClientID *string `json:"firstPartyOAUTH2ClientID,omitempty"`

	// Output only. Third party OAuth Client ID for External Identity Providers.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID.third_party_oauth2_client_id
	ThirdPartyOAUTH2ClientID *string `json:"thirdPartyOAUTH2ClientID,omitempty"`
}
