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

func DCLEnvironmentGroupSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Apigee/EnvironmentGroup",
			Description: "The Apigee EnvironmentGroup resource",
			StructName:  "EnvironmentGroup",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a EnvironmentGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "environmentGroup",
						Required:    true,
						Description: "A full instance of a EnvironmentGroup",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a EnvironmentGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "environmentGroup",
						Required:    true,
						Description: "A full instance of a EnvironmentGroup",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a EnvironmentGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "environmentGroup",
						Required:    true,
						Description: "A full instance of a EnvironmentGroup",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all EnvironmentGroup",
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
				Description: "The function used to list information about many EnvironmentGroup",
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
				"EnvironmentGroup": &dcl.Component{
					Title:     "EnvironmentGroup",
					ID:        "organizations/{{apigee_organization}}/envgroups/{{name}}",
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"hostnames",
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
								Type:      "integer",
								Format:    "int64",
								GoName:    "CreatedAt",
								ReadOnly:  true,
								Immutable: true,
							},
							"hostnames": &dcl.Property{
								Type:        "array",
								GoName:      "Hostnames",
								Description: "Required. Host names for this environment group.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"lastModifiedAt": &dcl.Property{
								Type:      "integer",
								Format:    "int64",
								GoName:    "LastModifiedAt",
								ReadOnly:  true,
								Immutable: true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "ID of the environment group.",
								Immutable:   true,
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "EnvironmentGroupStateEnum",
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
