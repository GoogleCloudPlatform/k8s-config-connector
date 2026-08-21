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

func DCLTenantSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "IdentityToolkit/Tenant",
			Description: "The IdentityToolkit Tenant resource",
			StructName:  "Tenant",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Tenant",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tenant",
						Required:    true,
						Description: "A full instance of a Tenant",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Tenant",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tenant",
						Required:    true,
						Description: "A full instance of a Tenant",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Tenant",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tenant",
						Required:    true,
						Description: "A full instance of a Tenant",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Tenant",
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
				Description: "The function used to list information about many Tenant",
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
				"Tenant": &dcl.Component{
					Title:           "Tenant",
					ID:              "projects/{{project}}/tenants/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
						},
						Properties: map[string]*dcl.Property{
							"allowPasswordSignup": &dcl.Property{
								Type:        "boolean",
								GoName:      "AllowPasswordSignup",
								Description: "Whether to allow email/password user authentication.",
							},
							"disableAuth": &dcl.Property{
								Type:        "boolean",
								GoName:      "DisableAuth",
								Description: "Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Display name of the tenant.",
							},
							"enableAnonymousUser": &dcl.Property{
								Type:        "boolean",
								GoName:      "EnableAnonymousUser",
								Description: "Whether to enable anonymous user authentication.",
							},
							"enableEmailLinkSignin": &dcl.Property{
								Type:        "boolean",
								GoName:      "EnableEmailLinkSignin",
								Description: "Whether to enable email link user authentication.",
							},
							"mfaConfig": &dcl.Property{
								Type:        "object",
								GoName:      "MfaConfig",
								GoType:      "TenantMfaConfig",
								Description: "The tenant-level configuration of MFA options.",
								Properties: map[string]*dcl.Property{
									"enabledProviders": &dcl.Property{
										Type:        "array",
										GoName:      "EnabledProviders",
										Description: "A list of usable second factors for this project.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "TenantMfaConfigEnabledProvidersEnum",
											Enum: []string{
												"PROVIDER_UNSPECIFIED",
												"PHONE_SMS",
											},
										},
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "TenantMfaConfigStateEnum",
										Description: "Whether MultiFactor Authentication has been enabled for this project. Possible values: STATE_UNSPECIFIED, DISABLED, ENABLED, MANDATORY",
										Enum: []string{
											"STATE_UNSPECIFIED",
											"DISABLED",
											"ENABLED",
											"MANDATORY",
										},
									},
								},
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. Resource name of a tenant. For example: \"projects/{project-id}/tenants/{tenant-id}\"",
								Immutable:                true,
								ServerGeneratedParameter: true,
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
							"testPhoneNumbers": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "TestPhoneNumbers",
								Description: "A map of <test phone number, fake code> pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).",
							},
						},
					},
				},
			},
		},
	}
}
