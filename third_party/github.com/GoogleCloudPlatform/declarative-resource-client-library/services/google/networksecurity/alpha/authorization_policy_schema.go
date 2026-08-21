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

func DCLAuthorizationPolicySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkSecurity/AuthorizationPolicy",
			Description: "The NetworkSecurity AuthorizationPolicy resource",
			StructName:  "AuthorizationPolicy",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a AuthorizationPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "authorizationPolicy",
						Required:    true,
						Description: "A full instance of a AuthorizationPolicy",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a AuthorizationPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "authorizationPolicy",
						Required:    true,
						Description: "A full instance of a AuthorizationPolicy",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a AuthorizationPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "authorizationPolicy",
						Required:    true,
						Description: "A full instance of a AuthorizationPolicy",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all AuthorizationPolicy",
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
				Description: "The function used to list information about many AuthorizationPolicy",
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
				"AuthorizationPolicy": &dcl.Component{
					Title:           "AuthorizationPolicy",
					ID:              "projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"action",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"action": &dcl.Property{
								Type:        "string",
								GoName:      "Action",
								GoType:      "AuthorizationPolicyActionEnum",
								Description: "Required. The action to take when a rule match is found. Possible values are \"ALLOW\" or \"DENY\". Possible values: ACTION_UNSPECIFIED, ALLOW, DENY",
								Enum: []string{
									"ACTION_UNSPECIFIED",
									"ALLOW",
									"DENY",
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
								Description: "Optional. Free-text description of the resource.",
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. Set of label tags associated with the AuthorizationPolicy resource.",
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
								Description: "Required. Name of the AuthorizationPolicy resource.",
								Immutable:   true,
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
								Description: "Optional. List of rules to match. If not set, the action specified in the ‘action’ field will be applied without any additional rule checks.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "AuthorizationPolicyRules",
									Properties: map[string]*dcl.Property{
										"destinations": &dcl.Property{
											Type:        "array",
											GoName:      "Destinations",
											Description: "Optional. List of attributes for the traffic destination. If not set, the action specified in the ‘action’ field will be applied without any rule checks for the destination.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "AuthorizationPolicyRulesDestinations",
												Required: []string{
													"hosts",
													"ports",
												},
												Properties: map[string]*dcl.Property{
													"hosts": &dcl.Property{
														Type:        "array",
														GoName:      "Hosts",
														Description: "Required. List of host names to match. Matched against HOST header in http requests. Each host can be an exact match, or a prefix match (example, “mydomain.*”) or a suffix match (example, *.myorg.com”) or a presence(any) match “*”.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"httpHeaderMatch": &dcl.Property{
														Type:        "object",
														GoName:      "HttpHeaderMatch",
														GoType:      "AuthorizationPolicyRulesDestinationsHttpHeaderMatch",
														Description: "Optional. Match against key:value pair in http header. Provides a flexible match based on HTTP headers, for potentially advanced use cases.",
														Required: []string{
															"headerName",
															"regexMatch",
														},
														Properties: map[string]*dcl.Property{
															"headerName": &dcl.Property{
																Type:        "string",
																GoName:      "HeaderName",
																Description: "Required. The name of the HTTP header to match. For matching against the HTTP request's authority, use a headerMatch with the header name \":authority\". For matching a request's method, use the headerName \":method\".",
															},
															"regexMatch": &dcl.Property{
																Type:        "string",
																GoName:      "RegexMatch",
																Description: "Required. The value of the header must match the regular expression specified in regexMatch. For regular expression grammar, please see: en.cppreference.com/w/cpp/regex/ecmascript For matching against a port specified in the HTTP request, use a headerMatch with headerName set to Host and a regular expression that satisfies the RFC2616 Host header's port specifier.",
															},
														},
													},
													"methods": &dcl.Property{
														Type:        "array",
														GoName:      "Methods",
														Description: "Optional. A list of HTTP methods to match. Should not be set for gRPC services.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"ports": &dcl.Property{
														Type:        "array",
														GoName:      "Ports",
														Description: "Required. List of destination ports to match.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "integer",
															Format: "int64",
															GoType: "int64",
														},
													},
												},
											},
										},
										"sources": &dcl.Property{
											Type:        "array",
											GoName:      "Sources",
											Description: "Optional. List of attributes for the traffic source. If not set, the action specified in the ‘action’ field will be applied without any rule checks for the source.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "AuthorizationPolicyRulesSources",
												Properties: map[string]*dcl.Property{
													"ipBlocks": &dcl.Property{
														Type:        "array",
														GoName:      "IPBlocks",
														Description: "Optional. List of CIDR ranges to match based on source IP address. Single IP (e.g., \"1.2.3.4\") and CIDR (e.g., \"1.2.3.0/24\") are supported.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"principals": &dcl.Property{
														Type:        "array",
														GoName:      "Principals",
														Description: "Optional. List of peer identities to match for authorization. Each peer can be an exact match, or a prefix match (example, “namespace/*”) or a suffix match (example, */service-account”) or a presence match “*”.",
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
