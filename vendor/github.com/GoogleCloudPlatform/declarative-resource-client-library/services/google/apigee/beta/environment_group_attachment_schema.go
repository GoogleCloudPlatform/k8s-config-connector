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
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLEnvironmentGroupAttachmentSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Apigee/EnvironmentGroupAttachment",
			Description: "The Apigee EnvironmentGroupAttachment resource",
			StructName:  "EnvironmentGroupAttachment",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a EnvironmentGroupAttachment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "EnvironmentGroupAttachment",
						Required:    true,
						Description: "A full instance of a EnvironmentGroupAttachment",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a EnvironmentGroupAttachment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "EnvironmentGroupAttachment",
						Required:    true,
						Description: "A full instance of a EnvironmentGroupAttachment",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a EnvironmentGroupAttachment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "EnvironmentGroupAttachment",
						Required:    true,
						Description: "A full instance of a EnvironmentGroupAttachment",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all EnvironmentGroupAttachment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "envgroup",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many EnvironmentGroupAttachment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "envgroup",
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
				"EnvironmentGroupAttachment": &dcl.Component{
					Title:     "EnvironmentGroupAttachment",
					ID:        "{{envgroup}}/attachments/{{name}}",
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"environment",
							"envgroup",
						},
						Properties: map[string]*dcl.Property{
							"createdAt": &dcl.Property{
								Type:      "integer",
								Format:    "int64",
								GoName:    "CreatedAt",
								ReadOnly:  true,
								Immutable: true,
							},
							"envgroup": &dcl.Property{
								Type:                "string",
								GoName:              "Envgroup",
								Description:         "The environment group for the resource",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Apigee/EnvironmentGroup",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"environment": &dcl.Property{
								Type:        "string",
								GoName:      "Environment",
								Description: "Required. ID of the attached environment.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Apigee/Environment",
										Field:    "name",
									},
								},
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "ID of the environment group attachment.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
						},
					},
				},
			},
		},
	}
}
