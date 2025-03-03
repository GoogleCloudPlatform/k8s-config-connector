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

func DCLAddressGroupSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkSecurity/AddressGroup",
			Description: "The NetworkSecurity AddressGroup resource",
			StructName:  "AddressGroup",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a AddressGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "addressGroup",
						Required:    true,
						Description: "A full instance of a AddressGroup",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a AddressGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "addressGroup",
						Required:    true,
						Description: "A full instance of a AddressGroup",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a AddressGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "addressGroup",
						Required:    true,
						Description: "A full instance of a AddressGroup",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all AddressGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many AddressGroup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
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
				"AddressGroup": &dcl.Component{
					Title:       "AddressGroup",
					ID:          "{{parent}}/locations/{{location}}/addressGroups/{{name}}",
					LabelsField: "labels",
					HasCreate:   true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"type",
							"capacity",
							"parent",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"capacity": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "Capacity",
								Description: "Required. Capacity of the Address Group.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. Free-text description of the resource.",
							},
							"items": &dcl.Property{
								Type:        "array",
								GoName:      "Items",
								Description: "Optional. List of items.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
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
								Description: "Required. Name of the AddressGroup resource.",
								Immutable:   true,
							},
							"parent": &dcl.Property{
								Type:                "string",
								GoName:              "Parent",
								Description:         "The parent of the resource.",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Organization",
										Field:    "name",
										Parent:   true,
									},
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"type": &dcl.Property{
								Type:        "string",
								GoName:      "Type",
								GoType:      "AddressGroupTypeEnum",
								Description: "Required. The type of the Address Group. Possible values are \"IPv4\" or \"IPV6\". Possible values: TYPE_UNSPECIFIED, IPV4, IPV6",
								Immutable:   true,
								Enum: []string{
									"TYPE_UNSPECIFIED",
									"IPV4",
									"IPV6",
								},
							},
						},
					},
				},
			},
		},
	}
}
