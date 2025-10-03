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

func DCLServerTlsPolicySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "NetworkSecurity/ServerTlsPolicy",
			Description: "The NetworkSecurity ServerTlsPolicy resource",
			StructName:  "ServerTlsPolicy",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a ServerTlsPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "serverTlsPolicy",
						Required:    true,
						Description: "A full instance of a ServerTlsPolicy",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a ServerTlsPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "serverTlsPolicy",
						Required:    true,
						Description: "A full instance of a ServerTlsPolicy",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a ServerTlsPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "serverTlsPolicy",
						Required:    true,
						Description: "A full instance of a ServerTlsPolicy",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all ServerTlsPolicy",
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
				Description: "The function used to list information about many ServerTlsPolicy",
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
				"ServerTlsPolicy": &dcl.Component{
					Title:           "ServerTlsPolicy",
					ID:              "projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"allowOpen": &dcl.Property{
								Type:        "boolean",
								GoName:      "AllowOpen",
								Description: "Optional. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allow_open and mtls_policy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.",
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
								Description: "Optional. Set of label tags associated with the resource.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"mtlsPolicy": &dcl.Property{
								Type:        "object",
								GoName:      "MtlsPolicy",
								GoType:      "ServerTlsPolicyMtlsPolicy",
								Description: "Optional. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allow_open and mtls_policy are set, server allows both plain text and mTLS connections.",
								Required: []string{
									"clientValidationCa",
								},
								Properties: map[string]*dcl.Property{
									"clientValidationCa": &dcl.Property{
										Type:        "array",
										GoName:      "ClientValidationCa",
										Description: "Required. Defines the mechanism to obtain the Certificate Authority certificate to validate the client certificate.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "ServerTlsPolicyMtlsPolicyClientValidationCa",
											Properties: map[string]*dcl.Property{
												"caCertPath": &dcl.Property{
													Type:        "string",
													GoName:      "CaCertPath",
													Description: "The path to the file holding the CA certificate to validate the client or server certificate.",
													Conflicts: []string{
														"grpcEndpoint",
														"certificateProviderInstance",
													},
												},
												"certificateProviderInstance": &dcl.Property{
													Type:        "object",
													GoName:      "CertificateProviderInstance",
													GoType:      "ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance",
													Description: "The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.",
													Conflicts: []string{
														"caCertPath",
														"grpcEndpoint",
													},
													Required: []string{
														"pluginInstance",
													},
													Properties: map[string]*dcl.Property{
														"pluginInstance": &dcl.Property{
															Type:        "string",
															GoName:      "PluginInstance",
															Description: "Required. Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\" to use Certificate Authority Service certificate provider instance.",
														},
													},
												},
												"grpcEndpoint": &dcl.Property{
													Type:        "object",
													GoName:      "GrpcEndpoint",
													GoType:      "ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint",
													Description: "gRPC specific configuration to access the gRPC server to obtain the CA certificate.",
													Conflicts: []string{
														"caCertPath",
														"certificateProviderInstance",
													},
													Required: []string{
														"targetUri",
													},
													Properties: map[string]*dcl.Property{
														"targetUri": &dcl.Property{
															Type:        "string",
															GoName:      "TargetUri",
															Description: "Required. The target URI of the gRPC endpoint. Only UDS path is supported, and should start with “unix:”.",
														},
													},
												},
											},
										},
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. Name of the ServerTlsPolicy resource.",
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
							"serverCertificate": &dcl.Property{
								Type:        "object",
								GoName:      "ServerCertificate",
								GoType:      "ServerTlsPolicyServerCertificate",
								Description: "Optional. Defines a mechanism to provision server identity (public and private keys). Cannot be combined with allow_open as a permissive mode that allows both plain text and TLS is not supported.",
								Properties: map[string]*dcl.Property{
									"certificateProviderInstance": &dcl.Property{
										Type:        "object",
										GoName:      "CertificateProviderInstance",
										GoType:      "ServerTlsPolicyServerCertificateCertificateProviderInstance",
										Description: "The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.",
										Conflicts: []string{
											"localFilepath",
											"grpcEndpoint",
										},
										Required: []string{
											"pluginInstance",
										},
										Properties: map[string]*dcl.Property{
											"pluginInstance": &dcl.Property{
												Type:        "string",
												GoName:      "PluginInstance",
												Description: "Required. Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\" to use Certificate Authority Service certificate provider instance.",
											},
										},
									},
									"grpcEndpoint": &dcl.Property{
										Type:        "object",
										GoName:      "GrpcEndpoint",
										GoType:      "ServerTlsPolicyServerCertificateGrpcEndpoint",
										Description: "gRPC specific configuration to access the gRPC server to obtain the cert and private key.",
										Conflicts: []string{
											"localFilepath",
											"certificateProviderInstance",
										},
										Required: []string{
											"targetUri",
										},
										Properties: map[string]*dcl.Property{
											"targetUri": &dcl.Property{
												Type:        "string",
												GoName:      "TargetUri",
												Description: "Required. The target URI of the gRPC endpoint. Only UDS path is supported, and should start with “unix:”.",
											},
										},
									},
									"localFilepath": &dcl.Property{
										Type:        "object",
										GoName:      "LocalFilepath",
										GoType:      "ServerTlsPolicyServerCertificateLocalFilepath",
										Description: "Obtain certificates and private key from a locally mounted filesystem path.",
										Conflicts: []string{
											"grpcEndpoint",
											"certificateProviderInstance",
										},
										Required: []string{
											"certificatePath",
											"privateKeyPath",
										},
										Properties: map[string]*dcl.Property{
											"certificatePath": &dcl.Property{
												Type:        "string",
												GoName:      "CertificatePath",
												Description: "Required. The path to the file that has the certificate containing public key.",
											},
											"privateKeyPath": &dcl.Property{
												Type:        "string",
												GoName:      "PrivateKeyPath",
												Description: "Required. The path to the file that has the private key.",
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
