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
package apigee

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLEnvironmentSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Apigee/Environment",
			Description: "The Apigee Environment resource",
			StructName:  "Environment",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Environment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "environment",
						Required:    true,
						Description: "A full instance of a Environment",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Environment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "environment",
						Required:    true,
						Description: "A full instance of a Environment",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Environment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "environment",
						Required:    true,
						Description: "A full instance of a Environment",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Environment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "apigeeOrganization",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Environment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "apigeeOrganization",
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
				"Environment": &dcl.Component{
					Title:     "Environment",
					ID:        "organizations/{{apigee_organization}}/environments/{{name}}",
					HasCreate: true,
					HasIAM:    true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"apigeeOrganization",
						},
						Properties: map[string]*dcl.Property{
							"apigeeOrganization": &dcl.Property{
								Type:        "string",
								GoName:      "ApigeeOrganization",
								Description: "The apigee organization for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Apigee/Organization",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"createdAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "CreatedAt",
								ReadOnly:    true,
								Description: "Output only. Creation time of this environment as milliseconds since epoch.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. Description of the environment.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Optional. Display name for this environment.",
							},
							"lastModifiedAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "LastModifiedAt",
								ReadOnly:    true,
								Description: "Output only. Last modification time of this environment as milliseconds since epoch.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. Name of the environment. Values must match the regular expression ^[.\\p{Alnum}-_]{1,255}$",
								Immutable:   true,
							},
							"properties": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Properties",
								Description: "Optional. Key-value pairs that may be used for customizing the environment.",
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "EnvironmentStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the environment. Values other than ACTIVE means the resource is not ready to use. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"CREATING",
									"ACTIVE",
									"DELETING",
								},
							},
						},
					},
				},
			},
		},
	}
}
