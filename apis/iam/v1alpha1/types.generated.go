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


// +kcc:proto=google.iam.v1beta.WorkloadIdentityPoolProvider
type WorkloadIdentityPoolProvider struct {

	// A display name for the provider. Cannot exceed 32 characters.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description for the provider. Cannot exceed 256 characters.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.description
	Description *string `json:"description,omitempty"`

	// Whether the provider is disabled. You cannot use a disabled provider to
	//  exchange tokens. However, existing tokens still grant access.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Maps attributes from authentication credentials issued by an external
	//  identity provider to Google Cloud attributes, such as `subject` and
	//  `segment`.
	//
	//  Each key must be a string specifying the Google Cloud IAM attribute to
	//  map to.
	//
	//  The following keys are supported:
	//
	//  * `google.subject`: The principal IAM is authenticating. You can reference
	//                      this value in IAM bindings. This is also the
	//                      subject that appears in Cloud Logging logs.
	//                      Cannot exceed 127 characters.
	//
	//  * `google.groups`: Groups the external identity belongs to. You can grant
	//                     groups access to resources using an IAM `principalSet`
	//                     binding; access applies to all members of the group.
	//
	//  You can also provide custom attributes by specifying
	//  `attribute.{custom_attribute}`, where `{custom_attribute}` is the name of
	//  the custom attribute to be mapped. You can define a maximum of 50 custom
	//  attributes. The maximum length of a mapped attribute key is
	//  100 characters, and the key may only contain the characters [a-z0-9_].
	//
	//  You can reference these attributes in IAM policies to define fine-grained
	//  access for a workload to Google Cloud resources. For example:
	//
	//  * `google.subject`:
	//  `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	//
	//  * `google.groups`:
	//  `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	//
	//  * `attribute.{custom_attribute}`:
	//  `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	//
	//  Each value must be a [Common Expression Language]
	//  (https://opensource.google/projects/cel) function that maps an
	//  identity provider credential to the normalized attribute specified by the
	//  corresponding map key.
	//
	//  You can use the `assertion` keyword in the expression to access a JSON
	//  representation of the authentication credential issued by the provider.
	//
	//  The maximum length of an attribute mapping expression is 2048 characters.
	//  When evaluated, the total size of all mapped attributes must not exceed
	//  8KB.
	//
	//  For AWS providers, the following rules apply:
	//
	//  - If no attribute mapping is defined, the following default mapping
	//    applies:
	//
	//    ```
	//    {
	//      "google.subject":"assertion.arn",
	//      "attribute.aws_role":
	//          "assertion.arn.contains('assumed-role')"
	//          " ? assertion.arn.extract('{account_arn}assumed-role/')"
	//          "   + 'assumed-role/'"
	//          "   + assertion.arn.extract('assumed-role/{role_name}/')"
	//          " : assertion.arn",
	//    }
	//    ```
	//
	//  - If any custom attribute mappings are defined, they must include a mapping
	//    to the `google.subject` attribute.
	//
	//
	//  For OIDC providers, the following rules apply:
	//
	//  - Custom attribute mappings must be defined, and must include a mapping to
	//    the `google.subject` attribute. For example, the following maps the
	//    `sub` claim of the incoming credential to the `subject` attribute on
	//    a Google token.
	//
	//    ```
	//    {"google.subject": "assertion.sub"}
	//    ```
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.attribute_mapping
	AttributeMapping map[string]string `json:"attributeMapping,omitempty"`

	// [A Common Expression Language](https://opensource.google/projects/cel)
	//  expression, in plain text, to restrict what otherwise valid authentication
	//  credentials issued by the provider should not be accepted.
	//
	//  The expression must output a boolean representing whether to allow the
	//  federation.
	//
	//  The following keywords may be referenced in the expressions:
	//
	//  * `assertion`: JSON representing the authentication credential issued by
	//                 the provider.
	//  * `google`: The Google attributes mapped from the assertion in the
	//              `attribute_mappings`.
	//  * `attribute`: The custom attributes mapped from the assertion in the
	//                 `attribute_mappings`.
	//
	//  The maximum length of the attribute condition expression is 4096
	//  characters. If unspecified, all valid authentication credential are
	//  accepted.
	//
	//  The following example shows how to only allow credentials with a mapped
	//  `google.groups` value of `admins`:
	//
	//  ```
	//  "'admins' in google.groups"
	//  ```
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.attribute_condition
	AttributeCondition *string `json:"attributeCondition,omitempty"`

	// An Amazon Web Services identity provider.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.aws
	Aws *WorkloadIdentityPoolProvider_Aws `json:"aws,omitempty"`

	// An OpenId Connect 1.0 identity provider.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.oidc
	Oidc *WorkloadIdentityPoolProvider_Oidc `json:"oidc,omitempty"`
}

// +kcc:proto=google.iam.v1beta.WorkloadIdentityPoolProvider.Aws
type WorkloadIdentityPoolProvider_Aws struct {
	// Required. The AWS account ID.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.Aws.account_id
	AccountID *string `json:"accountID,omitempty"`
}

// +kcc:proto=google.iam.v1beta.WorkloadIdentityPoolProvider.Oidc
type WorkloadIdentityPoolProvider_Oidc struct {
	// Required. The OIDC issuer URL.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.Oidc.issuer_uri
	IssuerURI *string `json:"issuerURI,omitempty"`

	// Acceptable values for the `aud` field (audience) in the OIDC token. Token
	//  exchange requests are rejected if the token audience does not match one
	//  of the configured values. Each audience may be at most 256 characters. A
	//  maximum of 10 audiences may be configured.
	//
	//  If this list is empty, the OIDC token audience must be equal to
	//  the full canonical resource name of the WorkloadIdentityPoolProvider,
	//  with or without the HTTPS prefix. For example:
	//
	//  ```
	//  //iam.googleapis.com/projects/<project-number>/locations/<location>/workloadIdentityPools/<pool-id>/providers/<provider-id>
	//  https://iam.googleapis.com/projects/<project-number>/locations/<location>/workloadIdentityPools/<pool-id>/providers/<provider-id>
	//  ```
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.Oidc.allowed_audiences
	AllowedAudiences []string `json:"allowedAudiences,omitempty"`
}

// +kcc:proto=google.iam.v1beta.WorkloadIdentityPoolProvider
type WorkloadIdentityPoolProviderObservedState struct {
	// Output only. The resource name of the provider.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of the provider.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPoolProvider.state
	State *string `json:"state,omitempty"`
}
