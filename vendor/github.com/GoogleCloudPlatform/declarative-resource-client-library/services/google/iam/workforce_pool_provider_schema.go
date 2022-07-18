// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package iam

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLWorkforcePoolProviderSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iam/WorkforcePoolProvider",
			Description: "The Iam WorkforcePoolProvider resource",
			StructName:  "WorkforcePoolProvider",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a WorkforcePoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "WorkforcePoolProvider",
						Required:    true,
						Description: "A full instance of a WorkforcePoolProvider",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a WorkforcePoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "WorkforcePoolProvider",
						Required:    true,
						Description: "A full instance of a WorkforcePoolProvider",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a WorkforcePoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "WorkforcePoolProvider",
						Required:    true,
						Description: "A full instance of a WorkforcePoolProvider",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all WorkforcePoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "workforcepool",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many WorkforcePoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "workforcepool",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"WorkforcePoolProvider": &dcl.Component{
					Title:     "WorkforcePoolProvider",
					ID:        "locations/{{location}}/workforcePools/{{workforce_pool}}/providers/{{name}}",
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"attributeMapping",
							"location",
							"workforcePool",
						},
						Properties: map[string]*dcl.Property{
							"attributeCondition": &dcl.Property{
								Type:        "string",
								GoName:      "AttributeCondition",
								Description: "A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` \"'admins' in google.groups\" ```",
							},
							"attributeMapping": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "AttributeMapping",
								Description: "Required. Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example:",
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "A user-specified description of the provider. Cannot exceed 256 characters.",
							},
							"disabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Disabled",
								Description: "Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "A user-specified display name for the provider. Cannot exceed 32 characters.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Output only. The resource name of the provider. Format: `locations/{location}/workforcePools/{workforce_pool_id}/providers/{provider_id}`",
								Immutable:   true,
							},
							"oidc": &dcl.Property{
								Type:        "object",
								GoName:      "Oidc",
								GoType:      "WorkforcePoolProviderOidc",
								Description: "An OpenId Connect 1.0 identity provider configuration.",
								Conflicts: []string{
									"saml",
								},
								Required: []string{
									"issuerUri",
									"clientId",
								},
								Properties: map[string]*dcl.Property{
									"clientId": &dcl.Property{
										Type:        "string",
										GoName:      "ClientId",
										Description: "Required. The client ID. Must match the audience claim of the JWT issued by the identity provider.",
									},
									"issuerUri": &dcl.Property{
										Type:        "string",
										GoName:      "IssuerUri",
										Description: "Required. The OIDC issuer URI. Must be a valid URI using the 'https' scheme.",
									},
								},
							},
							"saml": &dcl.Property{
								Type:        "object",
								GoName:      "Saml",
								GoType:      "WorkforcePoolProviderSaml",
								Description: "A SAML identity provider configuration.",
								Conflicts: []string{
									"oidc",
								},
								Required: []string{
									"idpMetadataXml",
								},
								Properties: map[string]*dcl.Property{
									"idpMetadataXml": &dcl.Property{
										Type:        "string",
										GoName:      "IdpMetadataXml",
										Description: "Required. SAML Identity provider configuration metadata xml doc. The xml document should comply with [SAML 2.0 specification](https://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf). The max size of the acceptable xml document will be bounded to 128k characters. The metadata xml document should satisfy the following constraints: 1) Must contain an Identity Provider Entity ID. 2) Must contain at least one non-expired signing key certificate. 3) For each signing key: a) Valid from should be no more than 7 days from now. b) Valid to should be no more than 10 years in the future. 4) Up to 3 IdP signing keys are allowed in the metadata xml. When updating the provider's metadata xml, at least one non-expired signing key must overlap with the existing metadata. This requirement is skipped if there are no non-expired signing keys present in the existing metadata.",
									},
								},
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "WorkforcePoolProviderStateEnum",
								ReadOnly:    true,
								Description: "Output only. The state of the provider. Possible values: STATE_UNSPECIFIED, ACTIVE, DELETED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ACTIVE",
									"DELETED",
								},
							},
							"workforcePool": &dcl.Property{
								Type:        "string",
								GoName:      "WorkforcePool",
								Description: "The workforce_pool for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Iam/WorkforcePool",
										Field:    "name",
										Parent:   true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
