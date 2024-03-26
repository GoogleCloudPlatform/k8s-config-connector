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
package networkservices

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLGatewaySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/Gateway",
			Description: "The NetworkServices Gateway resource",
			StructName:  "Gateway",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Gateway",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "gateway",
						Required:    true,
						Description: "A full instance of a Gateway",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Gateway",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "gateway",
						Required:    true,
						Description: "A full instance of a Gateway",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Gateway",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "gateway",
						Required:    true,
						Description: "A full instance of a Gateway",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Gateway",
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
				Description: "The function used to list information about many Gateway",
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
				"Gateway": &dcl.Component{
					Title:           "Gateway",
					ID:              "projects/{{project}}/locations/{{location}}/gateways/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"ports",
							"scope",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"addresses": &dcl.Property{
								Type:        "array",
								GoName:      "Addresses",
								Description: "One or more addresses with ports in format of \":\" that the Gateway must receive traffic on. The proxy binds to the ports specified. IP address can be anything that is allowed by the underlying infrastructure (auto-allocation, static IP, BYOIP).",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
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
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. Set of label tags associated with the Gateway resource.",
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
								Description: "Required. Name of the Gateway resource. It matches pattern `projects/*/locations/global/gateways/`.",
							},
							"ports": &dcl.Property{
								Type:        "array",
								GoName:      "Ports",
								Description: "Required. One or more ports that the Gateway must receive traffic on. The proxy binds to the ports specified. Gateway listen on 0.0.0.0 on the ports specified below.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "integer",
									Format: "int64",
									GoType: "int64",
								},
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
							"scope": &dcl.Property{
								Type:        "string",
								GoName:      "Scope",
								Description: "Required. Immutable. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.",
								Immutable:   true,
							},
							"selfLink": &dcl.Property{
								Type:        "string",
								GoName:      "SelfLink",
								ReadOnly:    true,
								Description: "Output only. Server-defined URL of this resource",
								Immutable:   true,
							},
							"serverTlsPolicy": &dcl.Property{
								Type:        "string",
								GoName:      "ServerTlsPolicy",
								Description: "Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Networksecurity/ServerTlsPolicy",
										Field:    "name",
										Format:   "projects/{{project}}/locations/global/serverTlsPolicies/{{name}}",
									},
								},
							},
							"type": &dcl.Property{
								Type:        "string",
								GoName:      "Type",
								GoType:      "GatewayTypeEnum",
								Description: "Immutable. The type of the customer managed gateway. Possible values: TYPE_UNSPECIFIED, OPEN_MESH, SECURE_WEB_GATEWAY",
								Immutable:   true,
								Enum: []string{
									"TYPE_UNSPECIFIED",
									"OPEN_MESH",
									"SECURE_WEB_GATEWAY",
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
