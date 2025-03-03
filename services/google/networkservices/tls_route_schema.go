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

func DCLTlsRouteSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/TlsRoute",
			Description: "The NetworkServices TlsRoute resource",
			StructName:  "TlsRoute",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a TlsRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tlsRoute",
						Required:    true,
						Description: "A full instance of a TlsRoute",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a TlsRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tlsRoute",
						Required:    true,
						Description: "A full instance of a TlsRoute",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a TlsRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "tlsRoute",
						Required:    true,
						Description: "A full instance of a TlsRoute",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all TlsRoute",
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
				Description: "The function used to list information about many TlsRoute",
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
				"TlsRoute": &dcl.Component{
					Title:           "TlsRoute",
					ID:              "projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}",
					ParentContainer: "project",
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
								Description: "Optional. Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`",
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
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"meshes": &dcl.Property{
								Type:        "array",
								GoName:      "Meshes",
								Description: "Optional. Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR",
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
								Description: "Required. Name of the TlsRoute resource. It matches pattern `projects/*/locations/global/tlsRoutes/tls_route_name>`.",
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
									GoType: "TlsRouteRules",
									Required: []string{
										"matches",
										"action",
									},
									Properties: map[string]*dcl.Property{
										"action": &dcl.Property{
											Type:        "object",
											GoName:      "Action",
											GoType:      "TlsRouteRulesAction",
											Description: "Required. The detailed rule defining how to route matched traffic.",
											Required: []string{
												"destinations",
											},
											Properties: map[string]*dcl.Property{
												"destinations": &dcl.Property{
													Type:        "array",
													GoName:      "Destinations",
													Description: "Required. The destination services to which traffic should be forwarded. At least one destination service is required.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "TlsRouteRulesActionDestinations",
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
																Description: "Optional. Specifies the proportion of requests forwareded to the backend referenced by the service_name field. This is computed as: weight/Sum(weights in destinations) Weights in all destinations does not need to sum up to 100.",
															},
														},
													},
												},
											},
										},
										"matches": &dcl.Property{
											Type:        "array",
											GoName:      "Matches",
											Description: "Required. RouteMatch defines the predicate used to match requests to a given action. Multiple match types are \"OR\"ed for evaluation.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "TlsRouteRulesMatches",
												Properties: map[string]*dcl.Property{
													"alpn": &dcl.Property{
														Type:        "array",
														GoName:      "Alpn",
														Description: "Optional. ALPN (Application-Layer Protocol Negotiation) to match against. Examples: \"http/1.1\", \"h2\". At least one of sni_host and alpn is required. Up to 5 alpns across all matches can be set.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"sniHost": &dcl.Property{
														Type:        "array",
														GoName:      "SniHost",
														Description: "Optional. SNI (server name indicator) to match against. SNI will be matched against all wildcard domains, i.e. www.example.com will be first matched against www.example.com, then *.example.com, then *.com. Partial wildcards are not supported, and values like *w.example.com are invalid. At least one of sni_host and alpn is required. Up to 5 sni hosts across all matches can be set.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
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
