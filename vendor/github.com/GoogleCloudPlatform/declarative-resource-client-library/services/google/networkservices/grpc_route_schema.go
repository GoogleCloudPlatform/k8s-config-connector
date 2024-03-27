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

func DCLGrpcRouteSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/GrpcRoute",
			Description: "The NetworkServices GrpcRoute resource",
			StructName:  "GrpcRoute",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a GrpcRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "grpcRoute",
						Required:    true,
						Description: "A full instance of a GrpcRoute",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a GrpcRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "grpcRoute",
						Required:    true,
						Description: "A full instance of a GrpcRoute",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a GrpcRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "grpcRoute",
						Required:    true,
						Description: "A full instance of a GrpcRoute",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all GrpcRoute",
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
				Description: "The function used to list information about many GrpcRoute",
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
				"GrpcRoute": &dcl.Component{
					Title:           "GrpcRoute",
					ID:              "projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"hostnames",
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
								Description: "Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`",
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
							"hostnames": &dcl.Property{
								Type:        "array",
								GoName:      "Hostnames",
								Description: "Required. Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be “precise” which is a domain name without the terminating dot of a network host (e.g. “foo.example.com”) or “wildcard”, which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or ‘-’, and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Router must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames \"*.foo.bar.com\" and \"*.bar.com\" to be associated with the same route, it is not possible to associate two routes both with \"*.bar.com\" or both with \"bar.com\". In the case that multiple routes match the hostname, the most specific match will be selected. For example, \"foo.bar.baz.com\" will take precedence over \"*.bar.baz.com\" and \"*.bar.baz.com\" will take precedence over \"*.baz.com\". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. \"xds:///service:123\"), otherwise they must supply the URI without a port (i.e. \"xds:///service\").",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. Set of label tags associated with the GrpcRoute resource.",
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
								Description: "Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`",
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
								Description: "Required. Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`",
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
								Description: "Required. A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "GrpcRouteRules",
									Required: []string{
										"action",
									},
									Properties: map[string]*dcl.Property{
										"action": &dcl.Property{
											Type:        "object",
											GoName:      "Action",
											GoType:      "GrpcRouteRulesAction",
											Description: "Required. A detailed rule defining how to route traffic. This field is required.",
											Properties: map[string]*dcl.Property{
												"destinations": &dcl.Property{
													Type:        "array",
													GoName:      "Destinations",
													Description: "Optional. The destination services to which traffic should be forwarded. If multiple destinations are specified, traffic will be split between Backend Service(s) according to the weight field of these destinations.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "GrpcRouteRulesActionDestinations",
														Required: []string{
															"serviceName",
														},
														Properties: map[string]*dcl.Property{
															"serviceName": &dcl.Property{
																Type:        "string",
																GoName:      "ServiceName",
																Description: "Required. The URL of a destination service to which to route traffic. Must refer to either a BackendService or ServiceDirectoryService.",
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
												"faultInjectionPolicy": &dcl.Property{
													Type:        "object",
													GoName:      "FaultInjectionPolicy",
													GoType:      "GrpcRouteRulesActionFaultInjectionPolicy",
													Description: "Optional. The specification for fault injection introduced into traffic to test the resiliency of clients to destination service failure. As part of fault injection, when clients send requests to a destination, delays can be introduced on a percentage of requests before sending those requests to the destination service. Similarly requests from clients can be aborted by for a percentage of requests. timeout and retry_policy will be ignored by clients that are configured with a fault_injection_policy",
													Properties: map[string]*dcl.Property{
														"abort": &dcl.Property{
															Type:        "object",
															GoName:      "Abort",
															GoType:      "GrpcRouteRulesActionFaultInjectionPolicyAbort",
															Description: "The specification for aborting to client requests.",
															Properties: map[string]*dcl.Property{
																"httpStatus": &dcl.Property{
																	Type:        "integer",
																	Format:      "int64",
																	GoName:      "HttpStatus",
																	Description: "The HTTP status code used to abort the request. The value must be between 200 and 599 inclusive.",
																},
																"percentage": &dcl.Property{
																	Type:        "integer",
																	Format:      "int64",
																	GoName:      "Percentage",
																	Description: "The percentage of traffic which will be aborted. The value must be between [0, 100]",
																},
															},
														},
														"delay": &dcl.Property{
															Type:        "object",
															GoName:      "Delay",
															GoType:      "GrpcRouteRulesActionFaultInjectionPolicyDelay",
															Description: "The specification for injecting delay to client requests.",
															Properties: map[string]*dcl.Property{
																"fixedDelay": &dcl.Property{
																	Type:        "string",
																	GoName:      "FixedDelay",
																	Description: "Specify a fixed delay before forwarding the request.",
																},
																"percentage": &dcl.Property{
																	Type:        "integer",
																	Format:      "int64",
																	GoName:      "Percentage",
																	Description: "The percentage of traffic on which delay will be injected. The value must be between [0, 100]",
																},
															},
														},
													},
												},
												"retryPolicy": &dcl.Property{
													Type:        "object",
													GoName:      "RetryPolicy",
													GoType:      "GrpcRouteRulesActionRetryPolicy",
													Description: "Optional. Specifies the retry policy associated with this route.",
													Properties: map[string]*dcl.Property{
														"numRetries": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "NumRetries",
															Description: "Specifies the allowed number of retries. This number must be > 0. If not specpfied, default to 1.",
														},
														"retryConditions": &dcl.Property{
															Type:        "array",
															GoName:      "RetryConditions",
															Description: "- connect-failure: Router will retry on failures connecting to Backend Services, for example due to connection timeouts. - refused-stream: Router will retry if the backend service resets the stream with a REFUSED_STREAM error code. This reset type indicates that it is safe to retry. - cancelled: Router will retry if the gRPC status code in the response header is set to cancelled - deadline-exceeded: Router will retry if the gRPC status code in the response header is set to deadline-exceeded - resource-exhausted: Router will retry if the gRPC status code in the response header is set to resource-exhausted - unavailable: Router will retry if the gRPC status code in the response header is set to unavailable",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
													},
												},
												"timeout": &dcl.Property{
													Type:        "string",
													GoName:      "Timeout",
													Description: "Optional. Specifies the timeout for selected route. Timeout is computed from the time the request has been fully processed (i.e. end of stream) up until the response has been completely processed. Timeout includes all retries.",
												},
											},
										},
										"matches": &dcl.Property{
											Type:        "array",
											GoName:      "Matches",
											Description: "Optional. Matches define conditions used for matching the rule against incoming gRPC requests. Each match is independent, i.e. this rule will be matched if ANY one of the matches is satisfied. If no matches field is specified, this rule will unconditionally match traffic.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "GrpcRouteRulesMatches",
												Properties: map[string]*dcl.Property{
													"headers": &dcl.Property{
														Type:        "array",
														GoName:      "Headers",
														Description: "Optional. Specifies a collection of headers to match.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "GrpcRouteRulesMatchesHeaders",
															Required: []string{
																"key",
																"value",
															},
															Properties: map[string]*dcl.Property{
																"key": &dcl.Property{
																	Type:        "string",
																	GoName:      "Key",
																	Description: "Required. The key of the header.",
																},
																"type": &dcl.Property{
																	Type:        "string",
																	GoName:      "Type",
																	GoType:      "GrpcRouteRulesMatchesHeadersTypeEnum",
																	Description: "Optional. Specifies how to match against the value of the header. If not specified, a default value of EXACT is used. Possible values: MATCH_TYPE_UNSPECIFIED, MATCH_ANY, MATCH_ALL",
																	Enum: []string{
																		"MATCH_TYPE_UNSPECIFIED",
																		"MATCH_ANY",
																		"MATCH_ALL",
																	},
																},
																"value": &dcl.Property{
																	Type:        "string",
																	GoName:      "Value",
																	Description: "Required. The value of the header.",
																},
															},
														},
													},
													"method": &dcl.Property{
														Type:        "object",
														GoName:      "Method",
														GoType:      "GrpcRouteRulesMatchesMethod",
														Description: "Optional. A gRPC method to match against. If this field is empty or omitted, will match all methods.",
														Required: []string{
															"grpcService",
															"grpcMethod",
														},
														Properties: map[string]*dcl.Property{
															"caseSensitive": &dcl.Property{
																Type:        "boolean",
																GoName:      "CaseSensitive",
																Description: "Optional. Specifies that matches are case sensitive. The default value is true. case_sensitive must not be used with a type of REGULAR_EXPRESSION.",
															},
															"grpcMethod": &dcl.Property{
																Type:        "string",
																GoName:      "GrpcMethod",
																Description: "Required. Name of the method to match against. If unspecified, will match all methods.",
															},
															"grpcService": &dcl.Property{
																Type:        "string",
																GoName:      "GrpcService",
																Description: "Required. Name of the service to match against. If unspecified, will match all services.",
															},
															"type": &dcl.Property{
																Type:        "string",
																GoName:      "Type",
																GoType:      "GrpcRouteRulesMatchesMethodTypeEnum",
																Description: "Optional. Specifies how to match against the name. If not specified, a default value of \"EXACT\" is used. Possible values: TYPE_UNSPECIFIED, EXACT, REGULAR_EXPRESSION",
																Enum: []string{
																	"TYPE_UNSPECIFIED",
																	"EXACT",
																	"REGULAR_EXPRESSION",
																},
															},
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
