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

func DCLHttpRouteSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/HttpRoute",
			Description: "The NetworkServices HttpRoute resource",
			StructName:  "HttpRoute",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a HttpRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "httpRoute",
						Required:    true,
						Description: "A full instance of a HttpRoute",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a HttpRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "httpRoute",
						Required:    true,
						Description: "A full instance of a HttpRoute",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a HttpRoute",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "httpRoute",
						Required:    true,
						Description: "A full instance of a HttpRoute",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all HttpRoute",
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
				Description: "The function used to list information about many HttpRoute",
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
				"HttpRoute": &dcl.Component{
					Title:           "HttpRoute",
					ID:              "projects/{{project}}/locations/{{location}}/httpRoutes/{{name}}",
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
								Description: "Optional. Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`",
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
								Description: "Required. Hostnames define a set of hosts that should match against the HTTP host header to select a HttpRoute to process the request. Hostname is the fully qualified domain name of a network host, as defined by RFC 1123 with the exception that ip addresses are not allowed. Wildcard hosts are supported as \"*\" (no prefix or suffix allowed).",
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
								Description: "Optional. Set of label tags associated with the HttpRoute resource.",
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
								Description: "Optional. Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR",
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
								Description: "Required. Name of the HttpRoute resource. It matches pattern `projects/*/locations/global/httpRoutes/http_route_name>`.",
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
								Description: "Required. Rules that define how traffic is routed and handled.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "HttpRouteRules",
									Properties: map[string]*dcl.Property{
										"action": &dcl.Property{
											Type:        "object",
											GoName:      "Action",
											GoType:      "HttpRouteRulesAction",
											Description: "The detailed rule defining how to route matched traffic.",
											Properties: map[string]*dcl.Property{
												"corsPolicy": &dcl.Property{
													Type:        "object",
													GoName:      "CorsPolicy",
													GoType:      "HttpRouteRulesActionCorsPolicy",
													Description: "The specification for allowing client side cross-origin requests.",
													Properties: map[string]*dcl.Property{
														"allowCredentials": &dcl.Property{
															Type:        "boolean",
															GoName:      "AllowCredentials",
															Description: "In response to a preflight request, setting this to true indicates that the actual request can include user credentials. This translates to the Access-Control-Allow-Credentials header. Default value is false.",
														},
														"allowHeaders": &dcl.Property{
															Type:        "array",
															GoName:      "AllowHeaders",
															Description: "Specifies the content for Access-Control-Allow-Headers header.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"allowMethods": &dcl.Property{
															Type:        "array",
															GoName:      "AllowMethods",
															Description: "Specifies the content for Access-Control-Allow-Methods header.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"allowOriginRegexes": &dcl.Property{
															Type:        "array",
															GoName:      "AllowOriginRegexes",
															Description: "Specifies the regular expression patterns that match allowed origins. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"allowOrigins": &dcl.Property{
															Type:        "array",
															GoName:      "AllowOrigins",
															Description: "Specifies the list of origins that will be allowed to do CORS requests. An origin is allowed if it matches either an item in allow_origins or an item in allow_origin_regexes.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"disabled": &dcl.Property{
															Type:        "boolean",
															GoName:      "Disabled",
															Description: "If true, the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.",
														},
														"exposeHeaders": &dcl.Property{
															Type:        "array",
															GoName:      "ExposeHeaders",
															Description: "Specifies the content for Access-Control-Expose-Headers header.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"maxAge": &dcl.Property{
															Type:        "string",
															GoName:      "MaxAge",
															Description: "Specifies how long result of a preflight request can be cached in seconds. This translates to the Access-Control-Max-Age header.",
														},
													},
												},
												"destinations": &dcl.Property{
													Type:        "array",
													GoName:      "Destinations",
													Description: "The destination to which traffic should be forwarded.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "HttpRouteRulesActionDestinations",
														Properties: map[string]*dcl.Property{
															"serviceName": &dcl.Property{
																Type:        "string",
																GoName:      "ServiceName",
																Description: "The URL of a BackendService to route traffic to.",
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
																Description: "Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weights are specified for any one service name, they need to be specified for all of them. If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.",
															},
														},
													},
												},
												"faultInjectionPolicy": &dcl.Property{
													Type:        "object",
													GoName:      "FaultInjectionPolicy",
													GoType:      "HttpRouteRulesActionFaultInjectionPolicy",
													Description: "The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure. As part of fault injection, when clients send requests to a backend service, delays can be introduced on a percentage of requests before sending those requests to the backend service. Similarly requests from clients can be aborted for a percentage of requests. timeout and retry_policy will be ignored by clients that are configured with a fault_injection_policy",
													Properties: map[string]*dcl.Property{
														"abort": &dcl.Property{
															Type:        "object",
															GoName:      "Abort",
															GoType:      "HttpRouteRulesActionFaultInjectionPolicyAbort",
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
															GoType:      "HttpRouteRulesActionFaultInjectionPolicyDelay",
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
												"redirect": &dcl.Property{
													Type:        "object",
													GoName:      "Redirect",
													GoType:      "HttpRouteRulesActionRedirect",
													Description: "If set, the request is directed as configured by this field.",
													Properties: map[string]*dcl.Property{
														"hostRedirect": &dcl.Property{
															Type:        "string",
															GoName:      "HostRedirect",
															Description: "The host that will be used in the redirect response instead of the one that was supplied in the request.",
														},
														"httpsRedirect": &dcl.Property{
															Type:        "boolean",
															GoName:      "HttpsRedirect",
															Description: "If set to true, the URL scheme in the redirected request is set to https. If set to false, the URL scheme of the redirected request will remain the same as that of the request. The default is set to false.",
														},
														"pathRedirect": &dcl.Property{
															Type:        "string",
															GoName:      "PathRedirect",
															Description: "The path that will be used in the redirect response instead of the one that was supplied in the request. path_redirect can not be supplied together with prefix_redirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.",
														},
														"portRedirect": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "PortRedirect",
															Description: "The port that will be used in the redirected request instead of the one that was supplied in the request.",
														},
														"prefixRewrite": &dcl.Property{
															Type:        "string",
															GoName:      "PrefixRewrite",
															Description: "Indicates that during redirection, the matched prefix (or path) should be swapped with this value. This option allows URLs be dynamically created based on the request.",
														},
														"responseCode": &dcl.Property{
															Type:        "string",
															GoName:      "ResponseCode",
															GoType:      "HttpRouteRulesActionRedirectResponseCodeEnum",
															Description: "The HTTP Status code to use for the redirect. Possible values: MOVED_PERMANENTLY_DEFAULT, FOUND, SEE_OTHER, TEMPORARY_REDIRECT, PERMANENT_REDIRECT",
															Enum: []string{
																"MOVED_PERMANENTLY_DEFAULT",
																"FOUND",
																"SEE_OTHER",
																"TEMPORARY_REDIRECT",
																"PERMANENT_REDIRECT",
															},
														},
														"stripQuery": &dcl.Property{
															Type:        "boolean",
															GoName:      "StripQuery",
															Description: "if set to true, any accompanying query portion of the original URL is removed prior to redirecting the request. If set to false, the query portion of the original URL is retained. The default is set to false.",
														},
													},
												},
												"requestHeaderModifier": &dcl.Property{
													Type:        "object",
													GoName:      "RequestHeaderModifier",
													GoType:      "HttpRouteRulesActionRequestHeaderModifier",
													Description: "The specification for modifying the headers of a matching request prior to delivery of the request to the destination.",
													Properties: map[string]*dcl.Property{
														"add": &dcl.Property{
															Type: "object",
															AdditionalProperties: &dcl.Property{
																Type: "string",
															},
															GoName:      "Add",
															Description: "Add the headers with given map where key is the name of the header, value is the value of the header.",
														},
														"remove": &dcl.Property{
															Type:        "array",
															GoName:      "Remove",
															Description: "Remove headers (matching by header names) specified in the list.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"set": &dcl.Property{
															Type: "object",
															AdditionalProperties: &dcl.Property{
																Type: "string",
															},
															GoName:      "Set",
															Description: "Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.",
														},
													},
												},
												"requestMirrorPolicy": &dcl.Property{
													Type:        "object",
													GoName:      "RequestMirrorPolicy",
													GoType:      "HttpRouteRulesActionRequestMirrorPolicy",
													Description: "Specifies the policy on how requests intended for the routes destination are shadowed to a separate mirrored destination. Proxy will not wait for the shadow destination to respond before returning the response. Prior to sending traffic to the shadow service, the host/authority header is suffixed with -shadow.",
													Properties: map[string]*dcl.Property{
														"destination": &dcl.Property{
															Type:        "object",
															GoName:      "Destination",
															GoType:      "HttpRouteRulesActionRequestMirrorPolicyDestination",
															Description: "The destination the requests will be mirrored to. The weight of the destination will be ignored.",
															Properties: map[string]*dcl.Property{
																"serviceName": &dcl.Property{
																	Type:        "string",
																	GoName:      "ServiceName",
																	Description: "The URL of a BackendService to route traffic to.",
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
																	Description: "Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weights are specified for any one service name, they need to be specified for all of them. If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.",
																},
															},
														},
													},
												},
												"responseHeaderModifier": &dcl.Property{
													Type:        "object",
													GoName:      "ResponseHeaderModifier",
													GoType:      "HttpRouteRulesActionResponseHeaderModifier",
													Description: "The specification for modifying the headers of a response prior to sending the response back to the client.",
													Properties: map[string]*dcl.Property{
														"add": &dcl.Property{
															Type: "object",
															AdditionalProperties: &dcl.Property{
																Type: "string",
															},
															GoName:      "Add",
															Description: "Add the headers with given map where key is the name of the header, value is the value of the header.",
														},
														"remove": &dcl.Property{
															Type:        "array",
															GoName:      "Remove",
															Description: "Remove headers (matching by header names) specified in the list.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"set": &dcl.Property{
															Type: "object",
															AdditionalProperties: &dcl.Property{
																Type: "string",
															},
															GoName:      "Set",
															Description: "Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.",
														},
													},
												},
												"retryPolicy": &dcl.Property{
													Type:        "object",
													GoName:      "RetryPolicy",
													GoType:      "HttpRouteRulesActionRetryPolicy",
													Description: "Specifies the retry policy associated with this route.",
													Properties: map[string]*dcl.Property{
														"numRetries": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "NumRetries",
															Description: "Specifies the allowed number of retries. This number must be > 0. If not specified, default to 1.",
														},
														"perTryTimeout": &dcl.Property{
															Type:        "string",
															GoName:      "PerTryTimeout",
															Description: "Specifies a non-zero timeout per retry attempt.",
														},
														"retryConditions": &dcl.Property{
															Type:        "array",
															GoName:      "RetryConditions",
															Description: "Specifies one or more conditions when this retry policy applies. Valid values are: 5xx: Proxy will attempt a retry if the destination service responds with any 5xx response code, of if the destination service does not respond at all, example: disconnect, reset, read timeout, connection failure and refused streams. gateway-error: Similar to 5xx, but only applies to response codes 502, 503, 504. reset: Proxy will attempt a retry if the destination service does not respond at all (disconnect/reset/read timeout) connect-failure: Proxy will retry on failures connecting to destination for example due to connection timeouts. retriable-4xx: Proxy will retry fro retriable 4xx response codes. Currently the only retriable error supported is 409. refused-stream: Proxy will retry if the destination resets the stream with a REFUSED_STREAM error code. This reset type indicates that it is safe to retry.",
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
													Description: "Specifies the timeout for selected route. Timeout is computed from the time the request has been fully processed (i.e. end of stream) up until the response has been completely processed. Timeout includes all retries.",
												},
												"urlRewrite": &dcl.Property{
													Type:        "object",
													GoName:      "UrlRewrite",
													GoType:      "HttpRouteRulesActionUrlRewrite",
													Description: "The specification for rewrite URL before forwarding requests to the destination.",
													Properties: map[string]*dcl.Property{
														"hostRewrite": &dcl.Property{
															Type:        "string",
															GoName:      "HostRewrite",
															Description: "Prior to forwarding the request to the selected destination, the requests host header is replaced by this value.",
														},
														"pathPrefixRewrite": &dcl.Property{
															Type:        "string",
															GoName:      "PathPrefixRewrite",
															Description: "Prior to forwarding the request to the selected destination, the matching portion of the requests path is replaced by this value.",
														},
													},
												},
											},
										},
										"matches": &dcl.Property{
											Type:        "array",
											GoName:      "Matches",
											Description: "A list of matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if ANY one of the matches is satisfied.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "HttpRouteRulesMatches",
												Properties: map[string]*dcl.Property{
													"fullPathMatch": &dcl.Property{
														Type:        "string",
														GoName:      "FullPathMatch",
														Description: "The HTTP request path value should exactly match this value. Only one of full_path_match, prefix_match, or regex_match should be used.",
														Conflicts: []string{
															"prefixMatch",
															"regexMatch",
														},
													},
													"headers": &dcl.Property{
														Type:        "array",
														GoName:      "Headers",
														Description: "Specifies a list of HTTP request headers to match against. ALL of the supplied headers must be matched.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "HttpRouteRulesMatchesHeaders",
															Properties: map[string]*dcl.Property{
																"exactMatch": &dcl.Property{
																	Type:        "string",
																	GoName:      "ExactMatch",
																	Description: "The value of the header should match exactly the content of exact_match.",
																	Conflicts: []string{
																		"regexMatch",
																		"prefixMatch",
																		"presentMatch",
																		"suffixMatch",
																		"rangeMatch",
																	},
																},
																"header": &dcl.Property{
																	Type:        "string",
																	GoName:      "Header",
																	Description: "The name of the HTTP header to match against.",
																},
																"invertMatch": &dcl.Property{
																	Type:        "boolean",
																	GoName:      "InvertMatch",
																	Description: "If specified, the match result will be inverted before checking. Default value is set to false.",
																},
																"prefixMatch": &dcl.Property{
																	Type:        "string",
																	GoName:      "PrefixMatch",
																	Description: "The value of the header must start with the contents of prefix_match.",
																	Conflicts: []string{
																		"exactMatch",
																		"regexMatch",
																		"presentMatch",
																		"suffixMatch",
																		"rangeMatch",
																	},
																},
																"presentMatch": &dcl.Property{
																	Type:        "boolean",
																	GoName:      "PresentMatch",
																	Description: "A header with header_name must exist. The match takes place whether or not the header has a value.",
																	Conflicts: []string{
																		"exactMatch",
																		"regexMatch",
																		"prefixMatch",
																		"suffixMatch",
																		"rangeMatch",
																	},
																},
																"rangeMatch": &dcl.Property{
																	Type:        "object",
																	GoName:      "RangeMatch",
																	GoType:      "HttpRouteRulesMatchesHeadersRangeMatch",
																	Description: "If specified, the rule will match if the request header value is within the range.",
																	Conflicts: []string{
																		"exactMatch",
																		"regexMatch",
																		"prefixMatch",
																		"presentMatch",
																		"suffixMatch",
																	},
																	Properties: map[string]*dcl.Property{
																		"end": &dcl.Property{
																			Type:        "integer",
																			Format:      "int64",
																			GoName:      "End",
																			Description: "End of the range (exclusive)",
																		},
																		"start": &dcl.Property{
																			Type:        "integer",
																			Format:      "int64",
																			GoName:      "Start",
																			Description: "Start of the range (inclusive)",
																		},
																	},
																},
																"regexMatch": &dcl.Property{
																	Type:        "string",
																	GoName:      "RegexMatch",
																	Description: "The value of the header must match the regular expression specified in regex_match. For regular expression grammar, please see: https://github.com/google/re2/wiki/Syntax",
																	Conflicts: []string{
																		"exactMatch",
																		"prefixMatch",
																		"presentMatch",
																		"suffixMatch",
																		"rangeMatch",
																	},
																},
																"suffixMatch": &dcl.Property{
																	Type:        "string",
																	GoName:      "SuffixMatch",
																	Description: "The value of the header must end with the contents of suffix_match.",
																	Conflicts: []string{
																		"exactMatch",
																		"regexMatch",
																		"prefixMatch",
																		"presentMatch",
																		"rangeMatch",
																	},
																},
															},
														},
													},
													"ignoreCase": &dcl.Property{
														Type:        "boolean",
														GoName:      "IgnoreCase",
														Description: "Specifies if prefix_match and full_path_match matches are case sensitive. The default value is false.",
													},
													"prefixMatch": &dcl.Property{
														Type:        "string",
														GoName:      "PrefixMatch",
														Description: "The HTTP request path value must begin with specified prefix_match. prefix_match must begin with a /. Only one of full_path_match, prefix_match, or regex_match should be used.",
														Conflicts: []string{
															"fullPathMatch",
															"regexMatch",
														},
													},
													"queryParameters": &dcl.Property{
														Type:        "array",
														GoName:      "QueryParameters",
														Description: "Specifies a list of query parameters to match against. ALL of the query parameters must be matched.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "HttpRouteRulesMatchesQueryParameters",
															Properties: map[string]*dcl.Property{
																"exactMatch": &dcl.Property{
																	Type:        "string",
																	GoName:      "ExactMatch",
																	Description: "The value of the query parameter must exactly match the contents of exact_match. Only one of exact_match, regex_match, or present_match must be set.",
																	Conflicts: []string{
																		"regexMatch",
																		"presentMatch",
																	},
																},
																"presentMatch": &dcl.Property{
																	Type:        "boolean",
																	GoName:      "PresentMatch",
																	Description: "Specifies that the QueryParameterMatcher matches if request contains query parameter, irrespective of whether the parameter has a value or not. Only one of exact_match, regex_match, or present_match must be set.",
																	Conflicts: []string{
																		"exactMatch",
																		"regexMatch",
																	},
																},
																"queryParameter": &dcl.Property{
																	Type:        "string",
																	GoName:      "QueryParameter",
																	Description: "The name of the query parameter to match.",
																},
																"regexMatch": &dcl.Property{
																	Type:        "string",
																	GoName:      "RegexMatch",
																	Description: "The value of the query parameter must match the regular expression specified by regex_match. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax Only one of exact_match, regex_match, or present_match must be set.",
																	Conflicts: []string{
																		"exactMatch",
																		"presentMatch",
																	},
																},
															},
														},
													},
													"regexMatch": &dcl.Property{
														Type:        "string",
														GoName:      "RegexMatch",
														Description: "The HTTP request path value must satisfy the regular expression specified by regex_match after removing any query parameters and anchor supplied with the original URL. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax Only one of full_path_match, prefix_match, or regex_match should be used.",
														Conflicts: []string{
															"fullPathMatch",
															"prefixMatch",
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
