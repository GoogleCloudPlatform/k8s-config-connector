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

func DCLIdentityAwareProxyClientSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iap/IdentityAwareProxyClient",
			Description: "The Iap IdentityAwareProxyClient resource",
			StructName:  "IdentityAwareProxyClient",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a IdentityAwareProxyClient",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "identityAwareProxyClient",
						Required:    true,
						Description: "A full instance of a IdentityAwareProxyClient",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a IdentityAwareProxyClient",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "identityAwareProxyClient",
						Required:    true,
						Description: "A full instance of a IdentityAwareProxyClient",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a IdentityAwareProxyClient",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "identityAwareProxyClient",
						Required:    true,
						Description: "A full instance of a IdentityAwareProxyClient",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all IdentityAwareProxyClient",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "brand",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many IdentityAwareProxyClient",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "brand",
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
				"IdentityAwareProxyClient": &dcl.Component{
					Title:           "IdentityAwareProxyClient",
					ID:              "projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
							"brand",
						},
						Properties: map[string]*dcl.Property{
							"brand": &dcl.Property{
								Type:        "string",
								GoName:      "Brand",
								Description: "The brand for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Iap/Brand",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Human-friendly name given to the OAuth client.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. Unique identifier of the OAuth client.",
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
							"secret": &dcl.Property{
								Type:        "string",
								GoName:      "Secret",
								ReadOnly:    true,
								Description: "Output only. Client secret of the OAuth client.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
