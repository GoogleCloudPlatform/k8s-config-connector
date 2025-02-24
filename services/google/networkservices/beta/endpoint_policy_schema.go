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
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLEndpointPolicySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkServices/EndpointPolicy",
			Description: "The NetworkServices EndpointPolicy resource",
			StructName:  "EndpointPolicy",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a EndpointPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "endpointPolicy",
						Required:    true,
						Description: "A full instance of a EndpointPolicy",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a EndpointPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "endpointPolicy",
						Required:    true,
						Description: "A full instance of a EndpointPolicy",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a EndpointPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "endpointPolicy",
						Required:    true,
						Description: "A full instance of a EndpointPolicy",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all EndpointPolicy",
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
				Description: "The function used to list information about many EndpointPolicy",
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
				"EndpointPolicy": &dcl.Component{
					Title:           "EndpointPolicy",
					ID:              "projects/{{project}}/locations/{{location}}/endpointPolicies/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"type",
							"endpointMatcher",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"authorizationPolicy": &dcl.Property{
								Type:        "string",
								GoName:      "AuthorizationPolicy",
								Description: "Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Networksecurity/AuthorizationPolicy",
										Field:    "name",
									},
								},
							},
							"clientTlsPolicy": &dcl.Property{
								Type:        "string",
								GoName:      "ClientTlsPolicy",
								Description: "Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Networksecurity/ClientTlsPolicy",
										Field:    "name",
									},
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
							"endpointMatcher": &dcl.Property{
								Type:        "object",
								GoName:      "EndpointMatcher",
								GoType:      "EndpointPolicyEndpointMatcher",
								Description: "Required. A matcher that selects endpoints to which the policies should be applied.",
								Properties: map[string]*dcl.Property{
									"metadataLabelMatcher": &dcl.Property{
										Type:        "object",
										GoName:      "MetadataLabelMatcher",
										GoType:      "EndpointPolicyEndpointMatcherMetadataLabelMatcher",
										Description: "The matcher is based on node metadata presented by xDS clients.",
										Properties: map[string]*dcl.Property{
											"metadataLabelMatchCriteria": &dcl.Property{
												Type:        "string",
												GoName:      "MetadataLabelMatchCriteria",
												GoType:      "EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum",
												Description: "Specifies how matching should be done. Supported values are: MATCH_ANY: At least one of the Labels specified in the matcher should match the metadata presented by xDS client. MATCH_ALL: The metadata presented by the xDS client should contain all of the labels specified here. The selection is determined based on the best match. For example, suppose there are three EndpointPolicy resources P1, P2 and P3 and if P1 has a the matcher as MATCH_ANY , P2 has MATCH_ALL , and P3 has MATCH_ALL . If a client with label connects, the config from P1 will be selected. If a client with label connects, the config from P2 will be selected. If a client with label connects, the config from P3 will be selected. If there is more than one best match, (for example, if a config P4 with selector exists and if a client with label connects), an error will be thrown. Possible values: METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED, MATCH_ANY, MATCH_ALL",
												Enum: []string{
													"METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED",
													"MATCH_ANY",
													"MATCH_ALL",
												},
											},
											"metadataLabels": &dcl.Property{
												Type:        "array",
												GoName:      "MetadataLabels",
												Description: "The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria This list can have at most 64 entries. The list can be empty if the match criteria is MATCH_ANY, to specify a wildcard match (i.e this matches any client).",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels",
													Required: []string{
														"labelName",
														"labelValue",
													},
													Properties: map[string]*dcl.Property{
														"labelName": &dcl.Property{
															Type:        "string",
															GoName:      "LabelName",
															Description: "Required. Label name presented as key in xDS Node Metadata.",
														},
														"labelValue": &dcl.Property{
															Type:        "string",
															GoName:      "LabelValue",
															Description: "Required. Label value presented as value corresponding to the above key, in xDS Node Metadata.",
														},
													},
												},
											},
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
								Description: "Optional. Set of label tags associated with the EndpointPolicy resource.",
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
								Description: "Required. Name of the EndpointPolicy resource.",
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
							"serverTlsPolicy": &dcl.Property{
								Type:        "string",
								GoName:      "ServerTlsPolicy",
								Description: "Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Networksecurity/ServerTlsPolicy",
										Field:    "name",
									},
								},
							},
							"trafficPortSelector": &dcl.Property{
								Type:        "object",
								GoName:      "TrafficPortSelector",
								GoType:      "EndpointPolicyTrafficPortSelector",
								Description: "Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.",
								Properties: map[string]*dcl.Property{
									"ports": &dcl.Property{
										Type:        "array",
										GoName:      "Ports",
										Description: "Optional. A list of ports. Can be port numbers or port range (example, specifies all ports from 80 to 90, including 80 and 90) or named ports or * to specify all ports. If the list is empty, all ports are selected.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
								},
							},
							"type": &dcl.Property{
								Type:        "string",
								GoName:      "Type",
								GoType:      "EndpointPolicyTypeEnum",
								Description: "Required. The type of endpoint config. This is primarily used to validate the configuration. Possible values: ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED, SIDECAR_PROXY, GRPC_SERVER",
								Enum: []string{
									"ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED",
									"SIDECAR_PROXY",
									"GRPC_SERVER",
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
