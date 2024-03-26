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

func DCLTcpRouteSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/TcpRoute",
			Description: "The NetworkServices TcpRoute resource",
			StructName:  "TcpRoute",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a TcpRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tcpRoute",
						Required:    true,
						Description: "A full instance of a TcpRoute",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a TcpRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tcpRoute",
						Required:    true,
						Description: "A full instance of a TcpRoute",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a TcpRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tcpRoute",
						Required:    true,
						Description: "A full instance of a TcpRoute",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all TcpRoute",
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
				Description: "The function used to list information about many TcpRoute",
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
				"TcpRoute": &dcl.Component{
					Title:           "TcpRoute",
					ID:              "projects/{{project}}/locations/{{location}}/tcpRoutes/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"rules",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
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
							"gateways": &dcl.Property{
								Type:        "array",
								GoName:      "Gateways",
								Description: "Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
									ResourceReferences: []*dcl.PropertyResourceReference{
										&dcl.PropertyResourceReference{
											Resource: "Networkservices/Gateway",
											Field:    "selfLink",
										},
									},
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. Set of label tags associated with the TcpRoute resource.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"meshes": &dcl.Property{
								Type:        "array",
								GoName:      "Meshes",
								Description: "Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
									ResourceReferences: []*dcl.PropertyResourceReference{
										&dcl.PropertyResourceReference{
											Resource: "Networkservices/Mesh",
											Field:    "selfLink",
										},
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.",
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
							"rules": &dcl.Property{
								Type:        "array",
								GoName:      "Rules",
								Description: "Required. Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "TcpRouteRules",
									Required: []string{
										"action",
									},
									Properties: map[string]*dcl.Property{
										"action": &dcl.Property{
											Type:        "object",
											GoName:      "Action",
											GoType:      "TcpRouteRulesAction",
											Description: "Required. The detailed rule defining how to route matched traffic.",
											Properties: map[string]*dcl.Property{
												"destinations": &dcl.Property{
													Type:        "array",
													GoName:      "Destinations",
													Description: "Optional. The destination services to which traffic should be forwarded. At least one destination service is required.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "TcpRouteRulesActionDestinations",
														Required: []string{
															"serviceName",
														},
														Properties: map[string]*dcl.Property{
															"serviceName": &dcl.Property{
																Type:        "string",
																GoName:      "ServiceName",
																Description: "Required. The URL of a BackendService to route traffic to.",
																ResourceReferences: []*dcl.PropertyResourceReference{
																	&dcl.PropertyResourceReference{
																		Resource: "Compute/BackendService",
																		Field:    "name",
																		Format:   "projects/{{project}}/locations/global/backendServices/{{name}}",
																	},
																},
															},
															"weight": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "Weight",
																Description: "Optional. Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weights are specified for any one service name, they need to be specified for all of them. If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.",
															},
														},
													},
												},
												"originalDestination": &dcl.Property{
													Type:        "boolean",
													GoName:      "OriginalDestination",
													Description: "Optional. If true, Router will use the destination IP and port of the original connection as the destination of the request. Default is false.",
												},
											},
										},
										"matches": &dcl.Property{
											Type:        "array",
											GoName:      "Matches",
											Description: "Optional. RouteMatch defines the predicate used to match requests to a given action. Multiple match types are “OR”ed for evaluation. If no routeMatch field is specified, this rule will unconditionally match traffic.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "TcpRouteRulesMatches",
												Required: []string{
													"address",
													"port",
												},
												Properties: map[string]*dcl.Property{
													"address": &dcl.Property{
														Type:        "string",
														GoName:      "Address",
														Description: "Required. Must be specified in the CIDR range format. A CIDR range consists of an IP Address and a prefix length to construct the subnet mask. By default, the prefix length is 32 (i.e. matches a single IP address). Only IPV4 addresses are supported. Examples: “10.0.0.1” - matches against this exact IP address. “10.0.0.0/8\" - matches against any IP address within the 10.0.0.0 subnet and 255.255.255.0 mask. \"0.0.0.0/0\" - matches against any IP address'.",
													},
													"port": &dcl.Property{
														Type:        "string",
														GoName:      "Port",
														Description: "Required. Specifies the destination port to match against.",
													},
												},
											},
										},
									},
								},
							},
							"selfLink": &dcl.Property{
								Type:        "string",
								GoName:      "SelfLink",
								ReadOnly:    true,
								Description: "Output only. Server-defined URL of this resource",
								Immutable:   true,
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
