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
package alpha

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLHttpFilterSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/HttpFilter",
			Description: "The NetworkServices HttpFilter resource",
			StructName:  "HttpFilter",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a HttpFilter",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "httpFilter",
						Required:    true,
						Description: "A full instance of a HttpFilter",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a HttpFilter",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "httpFilter",
						Required:    true,
						Description: "A full instance of a HttpFilter",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a HttpFilter",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "httpFilter",
						Required:    true,
						Description: "A full instance of a HttpFilter",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all HttpFilter",
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
			List: &dcl.Path{
				Description: "The function used to list information about many HttpFilter",
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
				"HttpFilter": &dcl.Component{
					Title:           "HttpFilter",
					ID:              "projects/{{project}}/locations/{{location}}/httpFilters/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"filterName",
							"configTypeUrl",
							"config",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"config": &dcl.Property{
								Type:        "string",
								GoName:      "Config",
								Description: "Required. The configuration needed to enable the HTTP filter. The configuration must be JSON formatted and only contain fields defined in the protobuf identified in config_type_url.",
							},
							"configTypeUrl": &dcl.Property{
								Type:        "string",
								GoName:      "ConfigTypeUrl",
								Description: "Required. The fully qualified versioned proto3 type url that the filter expects for its configuration. For example: 'type.googleapis.com/envoy.config.wasm.v2.WasmService'.",
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The timestamp when the resource was created.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. A free-text description of the resource. Max length 1024 characters.",
							},
							"filterName": &dcl.Property{
								Type:        "string",
								GoName:      "FilterName",
								Description: "Required. Name of the HTTP filter defined in the `config` field. It is used by the xDS API client to identify specific filter implementation the `config` must be applied to. It is different from the name of the HttpFilter resource and does not have to be unique. Example: 'envoy.wasm'.",
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. Set of label tags associated with the HttpFilter resource.",
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
								Description: "Required. Name of the HttpFilter resource.",
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
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The timestamp when the resource was updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
