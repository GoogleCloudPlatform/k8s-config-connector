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
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLServiceAccountSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iam/ServiceAccount",
			Description: "The Iam ServiceAccount resource",
			StructName:  "ServiceAccount",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a ServiceAccount",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "serviceAccount",
						Required:    true,
						Description: "A full instance of a ServiceAccount",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a ServiceAccount",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "serviceAccount",
						Required:    true,
						Description: "A full instance of a ServiceAccount",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a ServiceAccount",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "serviceAccount",
						Required:    true,
						Description: "A full instance of a ServiceAccount",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all ServiceAccount",
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
				Description: "The function used to list information about many ServiceAccount",
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
				"ServiceAccount": &dcl.Component{
					Title:           "ServiceAccount",
					ID:              "projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com",
					ParentContainer: "project",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Properties: map[string]*dcl.Property{
							"actasResources": &dcl.Property{
								Type:        "object",
								GoName:      "ActasResources",
								GoType:      "ServiceAccountActasResources",
								Description: "Optional.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"resources": &dcl.Property{
										Type:      "array",
										GoName:    "Resources",
										Immutable: true,
										SendEmpty: true,
										ListType:  "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "ServiceAccountActasResourcesResources",
											Properties: map[string]*dcl.Property{
												"fullResourceName": &dcl.Property{
													Type:      "string",
													GoName:    "FullResourceName",
													Immutable: true,
												},
											},
										},
									},
								},
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.",
							},
							"disabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Disabled",
								ReadOnly:    true,
								Description: "Output only. Whether the service account is disabled.",
								Immutable:   true,
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.",
							},
							"email": &dcl.Property{
								Type:        "string",
								GoName:      "Email",
								ReadOnly:    true,
								Description: "Output only. The email address of the service account.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.",
								Immutable:   true,
							},
							"oauth2ClientId": &dcl.Property{
								Type:        "string",
								GoName:      "OAuth2ClientId",
								ReadOnly:    true,
								Description: "Output only. The OAuth 2.0 client ID for the service account.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The ID of the project that owns the service account.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"uniqueId": &dcl.Property{
								Type:        "string",
								GoName:      "UniqueId",
								ReadOnly:    true,
								Description: "Output only. The unique, stable numeric ID for the service account. Each service account retains its unique ID even if you delete the service account. For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
