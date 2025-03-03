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

func DCLFleetSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "GkeHub/Fleet",
			Description: "The GkeHub Fleet resource",
			StructName:  "Fleet",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Fleet",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "fleet",
						Required:    true,
						Description: "A full instance of a Fleet",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Fleet",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "fleet",
						Required:    true,
						Description: "A full instance of a Fleet",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Fleet",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "fleet",
						Required:    true,
						Description: "A full instance of a Fleet",
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Fleet": &dcl.Component{
					Title:           "Fleet",
					ID:              "projects/{{project}}/locations/{{location}}/fleets/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. When the Fleet was created.",
								Immutable:   true,
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Optional. A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `Production Fleet`",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"managedNamespaces": &dcl.Property{
								Type:        "boolean",
								GoName:      "ManagedNamespaces",
								Description: "Optional. If true, namespaces must be explicitly declared in a `FleetNamespace` object in order to use Fleet Features.",
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. The full, unique resource name of this fleet in the format of `projects/{project}/locations/{location}/fleets/{fleet}`. Each GCP project can have at most one fleet resource, named \"default\".",
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
							"uid": &dcl.Property{
								Type:        "string",
								GoName:      "Uid",
								ReadOnly:    true,
								Description: "Output only. Google-generated UUID for this resource. This is unique across all Fleet resources. If a Fleet resource is deleted and another resource with the same name is created, it gets a different uid.",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. When the Fleet was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
