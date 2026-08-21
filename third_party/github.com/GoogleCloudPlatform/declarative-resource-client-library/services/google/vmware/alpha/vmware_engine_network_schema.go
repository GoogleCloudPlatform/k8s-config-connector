// Copyright 2023 Google LLC. All Rights Reserved.
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

func DCLVmwareEngineNetworkSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Vmware/VmwareEngineNetwork",
			Description: "The Vmware VmwareEngineNetwork resource",
			StructName:  "VmwareEngineNetwork",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a VmwareEngineNetwork",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "vmwareEngineNetwork",
						Required:    true,
						Description: "A full instance of a VmwareEngineNetwork",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a VmwareEngineNetwork",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "vmwareEngineNetwork",
						Required:    true,
						Description: "A full instance of a VmwareEngineNetwork",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a VmwareEngineNetwork",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "vmwareEngineNetwork",
						Required:    true,
						Description: "A full instance of a VmwareEngineNetwork",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all VmwareEngineNetwork",
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
				Description: "The function used to list information about many VmwareEngineNetwork",
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
				"VmwareEngineNetwork": &dcl.Component{
					Title:           "VmwareEngineNetwork",
					ID:              "projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"type",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Creation time of this resource.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "User-provided description for this VMware Engine network.",
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Checksum that may be sent on update and delete requests to ensure that the user-provided value is up to date before the server processes a request. The server computes checksums based on the value of other fields in the request.",
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
								Description: "Output only. The resource name of the VMware Engine network. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/global/vmwareEngineNetworks/my-network`",
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
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "VmwareEngineNetworkStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the VMware Engine network. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, UPDATING, DELETING",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"CREATING",
									"ACTIVE",
									"UPDATING",
									"DELETING",
								},
							},
							"type": &dcl.Property{
								Type:        "string",
								GoName:      "Type",
								GoType:      "VmwareEngineNetworkTypeEnum",
								Description: "Required. VMware Engine network type. Possible values: TYPE_UNSPECIFIED, LEGACY",
								Enum: []string{
									"TYPE_UNSPECIFIED",
									"LEGACY",
								},
							},
							"uid": &dcl.Property{
								Type:        "string",
								GoName:      "Uid",
								ReadOnly:    true,
								Description: "Output only. System-generated unique identifier for the resource.",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Last update time of this resource.",
								Immutable:   true,
							},
							"vpcNetworks": &dcl.Property{
								Type:        "array",
								GoName:      "VPCNetworks",
								ReadOnly:    true,
								Description: "Output only. VMware Engine service VPC networks that provide connectivity from a private cloud to customer projects, the internet, and other Google Cloud services.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "VmwareEngineNetworkVPCNetworks",
									Properties: map[string]*dcl.Property{
										"network": &dcl.Property{
											Type:        "string",
											GoName:      "Network",
											ReadOnly:    true,
											Description: "Output only. The relative resource name of the service VPC network this VMware Engine network is attached to. For example: `projects/123123/global/networks/my-network`",
											Immutable:   true,
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Compute/Network",
													Field:    "selfLink",
												},
											},
										},
										"type": &dcl.Property{
											Type:        "string",
											GoName:      "Type",
											GoType:      "VmwareEngineNetworkVPCNetworksTypeEnum",
											ReadOnly:    true,
											Description: "Output only. Type of VPC network (INTRANET, INTERNET, or GOOGLE_CLOUD) Possible values: TYPE_UNSPECIFIED, INTRANET, INTERNET, GOOGLE_CLOUD",
											Immutable:   true,
											Enum: []string{
												"TYPE_UNSPECIFIED",
												"INTRANET",
												"INTERNET",
												"GOOGLE_CLOUD",
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
