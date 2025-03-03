// Copyright 2024 Google LLC. All Rights Reserved.
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
package alpha

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLOAuthIdpConfigSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "IdentityToolkit/OAuthIdpConfig",
			Description: "The IdentityToolkit OAuthIdpConfig resource",
			StructName:  "OAuthIdpConfig",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a OAuthIdpConfig",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "oAuthIdpConfig",
						Required:    true,
						Description: "A full instance of a OAuthIdpConfig",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a OAuthIdpConfig",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "oAuthIdpConfig",
						Required:    true,
						Description: "A full instance of a OAuthIdpConfig",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a OAuthIdpConfig",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "oAuthIdpConfig",
						Required:    true,
						Description: "A full instance of a OAuthIdpConfig",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all OAuthIdpConfig",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many OAuthIdpConfig",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
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
				"OAuthIdpConfig": &dcl.Component{
					Title:           "OAuthIdpConfig",
					ID:              "projects/{{project}}/oauthIdpConfigs/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
						},
						Properties: map[string]*dcl.Property{
							"clientId": &dcl.Property{
								Type:        "string",
								GoName:      "ClientId",
								Description: "The client id of an OAuth client.",
							},
							"clientSecret": &dcl.Property{
								Type:        "string",
								GoName:      "ClientSecret",
								Description: "The client secret of the OAuth client, to enable OIDC code flow.",
								Sensitive:   true,
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "The config's display name set by developers.",
							},
							"enabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Enabled",
								Description: "True if allows the user to sign in with the provider.",
							},
							"issuer": &dcl.Property{
								Type:        "string",
								GoName:      "Issuer",
								Description: "For OIDC Idps, the issuer identifier.",
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The name of the Config resource",
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"responseType": &dcl.Property{
								Type:          "object",
								GoName:        "ResponseType",
								GoType:        "OAuthIdpConfigResponseType",
								Description:   "The multiple response type to request for in the OAuth authorization flow. This can possibly be a combination of set bits (e.g.: {id\\_token, token}).",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"code": &dcl.Property{
										Type:        "boolean",
										GoName:      "Code",
										Description: "If true, authorization code is returned from IdP's authorization endpoint.",
									},
									"idToken": &dcl.Property{
										Type:        "boolean",
										GoName:      "IdToken",
										Description: "If true, ID token is returned from IdP's authorization endpoint.",
									},
									"token": &dcl.Property{
										Type:        "boolean",
										GoName:      "Token",
										Description: "If true, access token is returned from IdP's authorization endpoint.",
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
