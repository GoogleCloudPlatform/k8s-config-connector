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

func DCLEkmConnectionSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Cloudkms/EkmConnection",
			Description: "The Cloudkms EkmConnection resource",
			StructName:  "EkmConnection",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a EkmConnection",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "ekmConnection",
						Required:    true,
						Description: "A full instance of a EkmConnection",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a EkmConnection",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "ekmConnection",
						Required:    true,
						Description: "A full instance of a EkmConnection",
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many EkmConnection",
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
				"EkmConnection": &dcl.Component{
					Title:           "EkmConnection",
					ID:              "projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"serviceResolvers",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time at which the EkmConnection was created.",
								Immutable:   true,
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update requests to ensure the client has an up-to-date value before proceeding.",
								Immutable:   true,
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
								Description: "The resource name for the EkmConnection in the format `projects/*/locations/*/ekmConnections/*`.",
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
							"serviceResolvers": &dcl.Property{
								Type:        "array",
								GoName:      "ServiceResolvers",
								Description: "A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "EkmConnectionServiceResolvers",
									Required: []string{
										"serviceDirectoryService",
										"hostname",
										"serverCertificates",
									},
									Properties: map[string]*dcl.Property{
										"endpointFilter": &dcl.Property{
											Type:        "string",
											GoName:      "EndpointFilter",
											Description: "Optional. The filter applied to the endpoints of the resolved service. If no filter is specified, all endpoints will be considered. An endpoint will be chosen arbitrarily from the filtered list for each request. For endpoint filter syntax and examples, see https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.",
										},
										"hostname": &dcl.Property{
											Type:        "string",
											GoName:      "Hostname",
											Description: "Required. The hostname of the EKM replica used at TLS and HTTP layers.",
										},
										"serverCertificates": &dcl.Property{
											Type:        "array",
											GoName:      "ServerCertificates",
											Description: "Required. A list of leaf server certificates used to authenticate HTTPS connections to the EKM replica. Currently, a maximum of 10 Certificate is supported.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "EkmConnectionServiceResolversServerCertificates",
												Required: []string{
													"rawDer",
												},
												Properties: map[string]*dcl.Property{
													"issuer": &dcl.Property{
														Type:        "string",
														GoName:      "Issuer",
														ReadOnly:    true,
														Description: "Output only. The issuer distinguished name in RFC 2253 format. Only present if parsed is true.",
													},
													"notAfterTime": &dcl.Property{
														Type:        "string",
														Format:      "date-time",
														GoName:      "NotAfterTime",
														ReadOnly:    true,
														Description: "Output only. The certificate is not valid after this time. Only present if parsed is true.",
													},
													"notBeforeTime": &dcl.Property{
														Type:        "string",
														Format:      "date-time",
														GoName:      "NotBeforeTime",
														ReadOnly:    true,
														Description: "Output only. The certificate is not valid before this time. Only present if parsed is true.",
													},
													"parsed": &dcl.Property{
														Type:        "boolean",
														GoName:      "Parsed",
														ReadOnly:    true,
														Description: "Output only. True if the certificate was parsed successfully.",
													},
													"rawDer": &dcl.Property{
														Type:        "string",
														GoName:      "RawDer",
														Description: "Required. The raw certificate bytes in DER format.",
													},
													"serialNumber": &dcl.Property{
														Type:        "string",
														GoName:      "SerialNumber",
														ReadOnly:    true,
														Description: "Output only. The certificate serial number as a hex string. Only present if parsed is true.",
													},
													"sha256Fingerprint": &dcl.Property{
														Type:        "string",
														GoName:      "Sha256Fingerprint",
														ReadOnly:    true,
														Description: "Output only. The SHA-256 certificate fingerprint as a hex string. Only present if parsed is true.",
													},
													"subject": &dcl.Property{
														Type:        "string",
														GoName:      "Subject",
														ReadOnly:    true,
														Description: "Output only. The subject distinguished name in RFC 2253 format. Only present if parsed is true.",
													},
													"subjectAlternativeDnsNames": &dcl.Property{
														Type:        "array",
														GoName:      "SubjectAlternativeDnsNames",
														ReadOnly:    true,
														Description: "Output only. The subject Alternative DNS names. Only present if parsed is true.",
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
												},
											},
										},
										"serviceDirectoryService": &dcl.Property{
											Type:        "string",
											GoName:      "ServiceDirectoryService",
											Description: "Required. The resource name of the Service Directory service pointing to an EKM replica, in the format `projects/*/locations/*/namespaces/*/services/*`.",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Servicedirectory/Service",
													Field:    "name",
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
		},
	}
}
