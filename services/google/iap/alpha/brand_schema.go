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

func DCLBrandSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iap/Brand",
			Description: "The Iap Brand resource",
			StructName:  "Brand",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Brand",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "brand",
						Required:    true,
						Description: "A full instance of a Brand",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Brand",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "brand",
						Required:    true,
						Description: "A full instance of a Brand",
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Brand",
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
				"Brand": &dcl.Component{
					Title:           "Brand",
					ID:              "projects/{{project}}/brands/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Properties: map[string]*dcl.Property{
							"applicationTitle": &dcl.Property{
								Type:        "string",
								GoName:      "ApplicationTitle",
								Description: "Application name displayed on OAuth consent screen.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. Identifier of the brand. NOTE: GCP project number achieves the same brand identification purpose as only one brand per project can be created.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"orgInternalOnly": &dcl.Property{
								Type:        "boolean",
								GoName:      "OrgInternalOnly",
								ReadOnly:    true,
								Description: "Output only. Whether the brand is only intended for usage inside the G Suite organization only.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "GCP Project id under which the brand is to be created.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"supportEmail": &dcl.Property{
								Type:        "string",
								GoName:      "SupportEmail",
								Description: "Support email displayed on the OAuth consent screen.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
