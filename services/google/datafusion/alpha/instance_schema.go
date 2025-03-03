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

func DCLInstanceSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "DataFusion/Instance",
			Description: "The DataFusion Instance resource",
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
					Title: "Instance",
					ID:    "projects/{{project}}/locations/{{location}}/instances/{{name}}",
					Locations: []string{
						"zone",
					},
					ParentContainer: "project",
					LabelsField:     "labels",
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
							"apiEndpoint": &dcl.Property{
								Type:        "string",
								GoName:      "ApiEndpoint",
								ReadOnly:    true,
								Description: "Output only. Endpoint on which the REST APIs is accessible.",
								Immutable:   true,
							},
							"availableVersion": &dcl.Property{
								Type:        "array",
								GoName:      "AvailableVersion",
								ReadOnly:    true,
								Description: "Available versions that the instance can be upgraded to.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "InstanceAvailableVersion",
									Properties: map[string]*dcl.Property{
										"availableFeatures": &dcl.Property{
											Type:        "array",
											GoName:      "AvailableFeatures",
											ReadOnly:    true,
											Description: "Represents a list of available feature names for a given version.",
											Immutable:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
											},
										},
										"defaultVersion": &dcl.Property{
											Type:        "boolean",
											GoName:      "DefaultVersion",
											ReadOnly:    true,
											Description: "Whether this is currently the default version for Cloud Data Fusion",
											Immutable:   true,
										},
										"versionNumber": &dcl.Property{
											Type:        "string",
											GoName:      "VersionNumber",
											ReadOnly:    true,
											Description: "The version number of the Data Fusion instance, such as '6.0.1.0'.",
											Immutable:   true,
										},
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time the instance was created.",
								Immutable:   true,
							},
							"dataprocServiceAccount": &dcl.Property{
								Type:        "string",
								GoName:      "DataprocServiceAccount",
								Description: "User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Iam/ServiceAccount",
										Field:    "email",
									},
								},
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "A description of this instance.",
								Immutable:   true,
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Display name for an instance.",
								Immutable:   true,
							},
							"enableStackdriverLogging": &dcl.Property{
								Type:        "boolean",
								GoName:      "EnableStackdriverLogging",
								Description: "Option to enable Stackdriver Logging.",
							},
							"enableStackdriverMonitoring": &dcl.Property{
								Type:        "boolean",
								GoName:      "EnableStackdriverMonitoring",
								Description: "Option to enable Stackdriver Monitoring.",
							},
							"gcsBucket": &dcl.Property{
								Type:        "string",
								GoName:      "GcsBucket",
								ReadOnly:    true,
								Description: "Output only. Cloud Storage bucket generated by Data Fusion in the customer project.",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "The resource labels for instance to use to annotate any related underlying resources such as GCE VMs. The character '=' is not allowed to be used within the labels.",
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
								Description: "Output only. The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.",
								Immutable:   true,
							},
							"networkConfig": &dcl.Property{
								Type:        "object",
								GoName:      "NetworkConfig",
								GoType:      "InstanceNetworkConfig",
								Description: "Network configuration options. These are required when a private Data Fusion instance is to be created.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"ipAllocation": &dcl.Property{
										Type:        "string",
										GoName:      "IPAllocation",
										Description: "The IP range in CIDR notation to use for the managed Data Fusion instance nodes. This range must not overlap with any other ranges used in the customer network.",
										Immutable:   true,
									},
									"network": &dcl.Property{
										Type:        "string",
										GoName:      "Network",
										Description: "Name of the network in the customer project with which the Tenant Project will be peered for executing pipelines. In case of shared VPC where the network resides in another host project the network should specified in the form of projects/{host-project-id}/global/networks/{network}",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Compute/Network",
												Field:    "name",
											},
										},
									},
								},
							},
							"options": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Options",
								Description: "Map of additional options used to configure the behavior of Data Fusion instance.",
								Immutable:   true,
							},
							"p4ServiceAccount": &dcl.Property{
								Type:        "string",
								GoName:      "P4ServiceAccount",
								ReadOnly:    true,
								Description: "Output only. P4 service account for the customer project.",
								Immutable:   true,
							},
							"privateInstance": &dcl.Property{
								Type:        "boolean",
								GoName:      "PrivateInstance",
								Description: "Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.",
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
							"serviceEndpoint": &dcl.Property{
								Type:        "string",
								GoName:      "ServiceEndpoint",
								ReadOnly:    true,
								Description: "Output only. Endpoint on which the Data Fusion UI is accessible.",
								Immutable:   true,
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "InstanceStateEnum",
								ReadOnly:    true,
								Description: "Output only. The current state of this Data Fusion instance. Possible values: STATE_UNSPECIFIED, ENABLED, DISABLED, UNKNOWN",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ENABLED",
									"DISABLED",
									"UNKNOWN",
								},
							},
							"stateMessage": &dcl.Property{
								Type:        "string",
								GoName:      "StateMessage",
								ReadOnly:    true,
								Description: "Output only. Additional information about the current state of this Data Fusion instance if available.",
								Immutable:   true,
							},
							"tenantProjectId": &dcl.Property{
								Type:        "string",
								GoName:      "TenantProjectId",
								ReadOnly:    true,
								Description: "Output only. The name of the tenant project.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
									},
								},
							},
							"type": &dcl.Property{
								Type:        "string",
								GoName:      "Type",
								GoType:      "InstanceTypeEnum",
								Description: "Required. Instance type. Possible values: TYPE_UNSPECIFIED, BASIC, ENTERPRISE, DEVELOPER",
								Immutable:   true,
								Enum: []string{
									"TYPE_UNSPECIFIED",
									"BASIC",
									"ENTERPRISE",
									"DEVELOPER",
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The time the instance was last updated.",
								Immutable:   true,
							},
							"version": &dcl.Property{
								Type:        "string",
								GoName:      "Version",
								Description: "Current version of the Data Fusion.",
							},
							"zone": &dcl.Property{
								Type:        "string",
								GoName:      "Zone",
								Description: "Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
