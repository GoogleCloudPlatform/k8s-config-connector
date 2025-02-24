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

func DCLPrivateCloudSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Vmware/PrivateCloud",
			Description: "The Vmware PrivateCloud resource",
			StructName:  "PrivateCloud",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a PrivateCloud",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "privateCloud",
						Required:    true,
						Description: "A full instance of a PrivateCloud",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a PrivateCloud",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "privateCloud",
						Required:    true,
						Description: "A full instance of a PrivateCloud",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a PrivateCloud",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "privateCloud",
						Required:    true,
						Description: "A full instance of a PrivateCloud",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all PrivateCloud",
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
				Description: "The function used to list information about many PrivateCloud",
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
				"PrivateCloud": &dcl.Component{
					Title:           "PrivateCloud",
					ID:              "projects/{{project}}/locations/{{location}}/privateClouds/{{name}}",
					UsesStateHint:   true,
					ParentContainer: "project",
					HasCreate:       true,
					HasIAM:          true,
					ApplyTimeout:    9600,
					DeleteTimeout:   7200,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"networkConfig",
							"managementCluster",
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
							"deleteTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "DeleteTime",
								ReadOnly:    true,
								Description: "Output only. Time when the resource was scheduled for deletion.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "User-provided description for this private cloud.",
							},
							"expireTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "ExpireTime",
								ReadOnly:    true,
								Description: "Output only. Time when the resource will be irreversibly deleted.",
								Immutable:   true,
							},
							"hcx": &dcl.Property{
								Type:        "object",
								GoName:      "Hcx",
								GoType:      "PrivateCloudHcx",
								ReadOnly:    true,
								Description: "Output only. HCX appliance.",
								Properties: map[string]*dcl.Property{
									"fqdn": &dcl.Property{
										Type:        "string",
										GoName:      "Fqdn",
										Description: "Fully qualified domain name of the appliance.",
									},
									"internalIP": &dcl.Property{
										Type:        "string",
										GoName:      "InternalIP",
										Description: "Internal IP address of the appliance.",
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "PrivateCloudHcxStateEnum",
										ReadOnly:    true,
										Description: "Output only. The state of the appliance. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING",
										Immutable:   true,
										Enum: []string{
											"STATE_UNSPECIFIED",
											"ACTIVE",
											"CREATING",
										},
									},
									"version": &dcl.Property{
										Type:        "string",
										GoName:      "Version",
										Description: "Version of the appliance.",
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Locations/Location",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"managementCluster": &dcl.Property{
								Type:        "object",
								GoName:      "ManagementCluster",
								GoType:      "PrivateCloudManagementCluster",
								Description: "Required. Input only. The management cluster for this private cloud. This field is required during creation of the private cloud to provide details for the default cluster. The following fields can't be changed after private cloud creation: `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.",
								Unreadable:  true,
								Required: []string{
									"clusterId",
								},
								Properties: map[string]*dcl.Property{
									"clusterId": &dcl.Property{
										Type:        "string",
										GoName:      "ClusterId",
										Description: "Required. The user-provided identifier of the new `Cluster`. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)",
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Output only. The resource name of this private cloud. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud`",
								Immutable:   true,
							},
							"networkConfig": &dcl.Property{
								Type:        "object",
								GoName:      "NetworkConfig",
								GoType:      "PrivateCloudNetworkConfig",
								Description: "Required. Network configuration of the private cloud.",
								Required: []string{
									"managementCidr",
								},
								Properties: map[string]*dcl.Property{
									"managementCidr": &dcl.Property{
										Type:        "string",
										GoName:      "ManagementCidr",
										Description: "Required. Management CIDR used by VMware management appliances.",
									},
									"managementIPAddressLayoutVersion": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "ManagementIPAddressLayoutVersion",
										ReadOnly:    true,
										Description: "Output only. The IP address layout version of the management IP address range. Possible versions include: * `managementIpAddressLayoutVersion=1`: Indicates the legacy IP address layout used by some existing private clouds. This is no longer supported for new private clouds as it does not support all features. * `managementIpAddressLayoutVersion=2`: Indicates the latest IP address layout used by all newly created private clouds. This version supports all current features.",
									},
									"vmwareEngineNetwork": &dcl.Property{
										Type:        "string",
										GoName:      "VmwareEngineNetwork",
										Description: "Optional. The relative resource name of the VMware Engine network attached to the private cloud. Specify the name in the following form: `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}` where `{project}` can either be a project number or a project ID.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Vmwareengine/VmwareEngineNetwork",
												Field:    "name",
											},
										},
									},
									"vmwareEngineNetworkCanonical": &dcl.Property{
										Type:        "string",
										GoName:      "VmwareEngineNetworkCanonical",
										ReadOnly:    true,
										Description: "Output only. The canonical name of the VMware Engine network in the form: `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Vmwareengine/VmwareEngineNetwork",
												Field:    "name",
											},
										},
									},
								},
							},
							"nsx": &dcl.Property{
								Type:        "object",
								GoName:      "Nsx",
								GoType:      "PrivateCloudNsx",
								ReadOnly:    true,
								Description: "Output only. NSX appliance.",
								Properties: map[string]*dcl.Property{
									"fqdn": &dcl.Property{
										Type:        "string",
										GoName:      "Fqdn",
										Description: "Fully qualified domain name of the appliance.",
									},
									"internalIP": &dcl.Property{
										Type:        "string",
										GoName:      "InternalIP",
										Description: "Internal IP address of the appliance.",
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "PrivateCloudNsxStateEnum",
										ReadOnly:    true,
										Description: "Output only. The state of the appliance. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING",
										Immutable:   true,
										Enum: []string{
											"STATE_UNSPECIFIED",
											"ACTIVE",
											"CREATING",
										},
									},
									"version": &dcl.Property{
										Type:        "string",
										GoName:      "Version",
										Description: "Version of the appliance.",
									},
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
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "PrivateCloudStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the resource. New values may be added to this enum when appropriate. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, UPDATING, FAILED, DELETED, PURGING",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ACTIVE",
									"CREATING",
									"UPDATING",
									"FAILED",
									"DELETED",
									"PURGING",
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
							"vcenter": &dcl.Property{
								Type:        "object",
								GoName:      "Vcenter",
								GoType:      "PrivateCloudVcenter",
								ReadOnly:    true,
								Description: "Output only. Vcenter appliance.",
								Properties: map[string]*dcl.Property{
									"fqdn": &dcl.Property{
										Type:        "string",
										GoName:      "Fqdn",
										Description: "Fully qualified domain name of the appliance.",
									},
									"internalIP": &dcl.Property{
										Type:        "string",
										GoName:      "InternalIP",
										Description: "Internal IP address of the appliance.",
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "PrivateCloudVcenterStateEnum",
										ReadOnly:    true,
										Description: "Output only. The state of the appliance. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING",
										Immutable:   true,
										Enum: []string{
											"STATE_UNSPECIFIED",
											"ACTIVE",
											"CREATING",
										},
									},
									"version": &dcl.Property{
										Type:        "string",
										GoName:      "Version",
										Description: "Version of the appliance.",
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
