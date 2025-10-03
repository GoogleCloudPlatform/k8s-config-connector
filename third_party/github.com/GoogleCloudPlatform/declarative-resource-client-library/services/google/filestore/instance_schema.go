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
package filestore

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLInstanceSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Filestore/Instance",
			Description: "The Filestore Instance resource",
			StructName:  "Instance",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "instance",
						Required:    true,
						Description: "A full instance of a Instance",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "instance",
						Required:    true,
						Description: "A full instance of a Instance",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "instance",
						Required:    true,
						Description: "A full instance of a Instance",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Instance",
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
				Description: "The function used to list information about many Instance",
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
				"Instance": &dcl.Component{
					Title:           "Instance",
					ID:              "projects/{{project}}/locations/{{location}}/instances/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time when the instance was created.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "The description of the instance (2048 characters or less).",
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.",
								Immutable:   true,
							},
							"fileShares": &dcl.Property{
								Type:        "array",
								GoName:      "FileShares",
								Description: "File system shares on the instance. For this version, only a single file share is supported.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "InstanceFileShares",
									Properties: map[string]*dcl.Property{
										"capacityGb": &dcl.Property{
											Type:        "integer",
											Format:      "int64",
											GoName:      "CapacityGb",
											Description: "File share capacity in gigabytes (GB). Cloud Filestore defines 1 GB as 1024^3 bytes.",
										},
										"name": &dcl.Property{
											Type:        "string",
											GoName:      "Name",
											Description: "The name of the file share (must be 16 characters or less).",
										},
										"nfsExportOptions": &dcl.Property{
											Type:        "array",
											GoName:      "NfsExportOptions",
											Description: "Nfs Export Options. There is a limit of 10 export options per file share.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "InstanceFileSharesNfsExportOptions",
												Properties: map[string]*dcl.Property{
													"accessMode": &dcl.Property{
														Type:          "string",
														GoName:        "AccessMode",
														GoType:        "InstanceFileSharesNfsExportOptionsAccessModeEnum",
														Description:   "Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE. Possible values: ACCESS_MODE_UNSPECIFIED, READ_ONLY, READ_WRITE",
														ServerDefault: true,
														Enum: []string{
															"ACCESS_MODE_UNSPECIFIED",
															"READ_ONLY",
															"READ_WRITE",
														},
													},
													"anonGid": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "AnonGid",
														Description: "An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.",
													},
													"anonUid": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "AnonUid",
														Description: "An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.",
													},
													"ipRanges": &dcl.Property{
														Type:        "array",
														GoName:      "IPRanges",
														Description: "List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"squashMode": &dcl.Property{
														Type:          "string",
														GoName:        "SquashMode",
														GoType:        "InstanceFileSharesNfsExportOptionsSquashModeEnum",
														Description:   "Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH. Possible values: SQUASH_MODE_UNSPECIFIED, NO_ROOT_SQUASH, ROOT_SQUASH",
														ServerDefault: true,
														Enum: []string{
															"SQUASH_MODE_UNSPECIFIED",
															"NO_ROOT_SQUASH",
															"ROOT_SQUASH",
														},
													},
												},
											},
										},
										"sourceBackup": &dcl.Property{
											Type:        "string",
											GoName:      "SourceBackup",
											Description: "The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Filestore/Backup",
													Field:    "name",
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
								Description: "Resource labels to represent user provided metadata.",
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
								Description: "Output only. The resource name of the instance, in the format `projects/{project}/locations/{location}/instances/{instance}`.",
								Immutable:   true,
							},
							"networks": &dcl.Property{
								Type:        "array",
								GoName:      "Networks",
								Description: "VPC networks to which the instance is connected. For this version, only a single network is supported.",
								Immutable:   true,
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "InstanceNetworks",
									Properties: map[string]*dcl.Property{
										"ipAddresses": &dcl.Property{
											Type:        "array",
											GoName:      "IPAddresses",
											ReadOnly:    true,
											Description: "Output only. IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or IPv6 addresses in the format `{block1}:{block2}:{block3}:{block4}:{block5}:{block6}:{block7}:{block8}`.",
											Immutable:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
											},
										},
										"modes": &dcl.Property{
											Type:          "array",
											GoName:        "Modes",
											Description:   "Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported.",
											Immutable:     true,
											ServerDefault: true,
											SendEmpty:     true,
											ListType:      "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "InstanceNetworksModesEnum",
												Enum: []string{
													"ADDRESS_MODE_UNSPECIFIED",
													"MODE_IPV4",
												},
											},
										},
										"network": &dcl.Property{
											Type:        "string",
											GoName:      "Network",
											Description: "The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.",
											Immutable:   true,
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Compute/Network",
													Field:    "name",
												},
											},
										},
										"reservedIPRange": &dcl.Property{
											Type:          "string",
											GoName:        "ReservedIPRange",
											Description:   "A /29 CIDR block in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29 or 192.168.0.0/29. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Cloud Filestore instances in the selected VPC network.",
											Immutable:     true,
											ServerDefault: true,
										},
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
								GoType:      "InstanceStateEnum",
								ReadOnly:    true,
								Description: "Output only. The instance state. Possible values: STATE_UNSPECIFIED, CREATING, READY, REPAIRING, DELETING, ERROR",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"CREATING",
									"READY",
									"REPAIRING",
									"DELETING",
									"ERROR",
								},
							},
							"statusMessage": &dcl.Property{
								Type:        "string",
								GoName:      "StatusMessage",
								ReadOnly:    true,
								Description: "Output only. Additional information about the instance state, if available.",
								Immutable:   true,
							},
							"tier": &dcl.Property{
								Type:        "string",
								GoName:      "Tier",
								GoType:      "InstanceTierEnum",
								Description: "The service tier of the instance. Possible values: TIER_UNSPECIFIED, STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ENTERPRISE",
								Immutable:   true,
								Enum: []string{
									"TIER_UNSPECIFIED",
									"STANDARD",
									"PREMIUM",
									"BASIC_HDD",
									"BASIC_SSD",
									"HIGH_SCALE_SSD",
									"ENTERPRISE",
								},
							},
						},
					},
				},
			},
		},
	}
}
