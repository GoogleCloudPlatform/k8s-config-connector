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
package cloudkms

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLKeyRingSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Cloudkms/KeyRing",
			Description: "The Cloudkms KeyRing resource",
			StructName:  "KeyRing",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a KeyRing",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "keyRing",
						Required:    true,
						Description: "A full instance of a KeyRing",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a KeyRing",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "keyRing",
						Required:    true,
						Description: "A full instance of a KeyRing",
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many KeyRing",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
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
				"KeyRing": &dcl.Component{
					Title:           "KeyRing",
					ID:              "projects/{{project}}/locations/{{location}}/keyRings/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time at which this KeyRing was created.",
								Immutable:   true,
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
								Description: "The resource name for the KeyRing in the format `projects/*/locations/*/keyRings/*`.",
								Immutable:   true,
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
						},
					},
				},
			},
		},
	}
}
