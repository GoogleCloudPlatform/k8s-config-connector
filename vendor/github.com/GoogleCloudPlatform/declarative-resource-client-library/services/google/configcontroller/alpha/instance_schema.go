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
			Title:       "ConfigController/Instance",
			Description: "The ConfigController Instance resource",
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
					ID:              "projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"managementConfig",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"gkeResourceLink": &dcl.Property{
								Type:        "string",
								GoName:      "GkeResourceLink",
								ReadOnly:    true,
								Description: "Output only. KrmApiHost GCP self link used for identifying the underlying endpoint (GKE cluster currently).",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Labels are used for additional information for a KrmApiHost.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"managementConfig": &dcl.Property{
								Type:        "object",
								GoName:      "ManagementConfig",
								GoType:      "InstanceManagementConfig",
								Description: "Configuration of the cluster management",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"fullManagementConfig": &dcl.Property{
										Type:        "object",
										GoName:      "FullManagementConfig",
										GoType:      "InstanceManagementConfigFullManagementConfig",
										Description: "Configuration of the full (Autopilot) cluster management",
										Immutable:   true,
										Conflicts: []string{
											"standardManagementConfig",
										},
										ServerDefault: true,
										SendEmpty:     true,
										Properties: map[string]*dcl.Property{
											"clusterCidrBlock": &dcl.Property{
												Type:          "string",
												GoName:        "ClusterCidrBlock",
												Description:   "The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.",
												Immutable:     true,
												ServerDefault: true,
											},
											"clusterNamedRange": &dcl.Property{
												Type:        "string",
												GoName:      "ClusterNamedRange",
												Description: "The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_cidr_block can be used to automatically create a GKE-managed one.",
												Immutable:   true,
											},
											"manBlock": &dcl.Property{
												Type:          "string",
												GoName:        "ManBlock",
												Description:   "Master Authorized Network. Allows access to the k8s master from this block.",
												Immutable:     true,
												ServerDefault: true,
											},
											"masterIPv4CidrBlock": &dcl.Property{
												Type:        "string",
												GoName:      "MasterIPv4CidrBlock",
												Description: "The /28 network that the masters will use.",
												Immutable:   true,
											},
											"network": &dcl.Property{
												Type:          "string",
												GoName:        "Network",
												Description:   "Existing VPC Network to put the GKE cluster and nodes in.",
												Immutable:     true,
												ServerDefault: true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Network",
														Field:    "name",
													},
												},
											},
											"servicesCidrBlock": &dcl.Property{
												Type:          "string",
												GoName:        "ServicesCidrBlock",
												Description:   "The IP address range for the cluster service IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.",
												Immutable:     true,
												ServerDefault: true,
											},
											"servicesNamedRange": &dcl.Property{
												Type:        "string",
												GoName:      "ServicesNamedRange",
												Description: "The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_cidr_block can be used to automatically create a GKE-managed one.",
												Immutable:   true,
											},
										},
									},
									"standardManagementConfig": &dcl.Property{
										Type:        "object",
										GoName:      "StandardManagementConfig",
										GoType:      "InstanceManagementConfigStandardManagementConfig",
										Description: "Configuration of the standard (GKE) cluster management",
										Immutable:   true,
										Conflicts: []string{
											"fullManagementConfig",
										},
										Required: []string{
											"masterIPv4CidrBlock",
										},
										Properties: map[string]*dcl.Property{
											"clusterCidrBlock": &dcl.Property{
												Type:        "string",
												GoName:      "ClusterCidrBlock",
												Description: "The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.",
												Immutable:   true,
											},
											"clusterNamedRange": &dcl.Property{
												Type:        "string",
												GoName:      "ClusterNamedRange",
												Description: "The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_cidr_block can be used to automatically create a GKE-managed one.",
												Immutable:   true,
											},
											"manBlock": &dcl.Property{
												Type:          "string",
												GoName:        "ManBlock",
												Description:   "Master Authorized Network. Allows access to the k8s master from this block.",
												Immutable:     true,
												ServerDefault: true,
											},
											"masterIPv4CidrBlock": &dcl.Property{
												Type:        "string",
												GoName:      "MasterIPv4CidrBlock",
												Description: "The /28 network that the masters will use.",
												Immutable:   true,
											},
											"network": &dcl.Property{
												Type:        "string",
												GoName:      "Network",
												Description: "Existing VPC Network to put the GKE cluster and nodes in.",
												Immutable:   true,
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Compute/Network",
														Field:    "name",
													},
												},
											},
											"servicesCidrBlock": &dcl.Property{
												Type:        "string",
												GoName:      "ServicesCidrBlock",
												Description: "The IP address range for the cluster service IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.",
												Immutable:   true,
											},
											"servicesNamedRange": &dcl.Property{
												Type:        "string",
												GoName:      "ServicesNamedRange",
												Description: "The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_cidr_block can be used to automatically create a GKE-managed one.",
												Immutable:   true,
											},
										},
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The name of this KrmApiHost resource in the format: 'projects/{project_id}/locations/{location}/krmApiHosts/{krm_api_host_id}'.",
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
								GoType:      "InstanceStateEnum",
								ReadOnly:    true,
								Description: "Output only. The current state of the internal state machine for the KrmApiHost. Possible values: STATE_UNSPECIFIED, CREATING, RUNNING, DELETING, SUSPENDED, READ_ONLY",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"CREATING",
									"RUNNING",
									"DELETING",
									"SUSPENDED",
									"READ_ONLY",
								},
							},
							"usePrivateEndpoint": &dcl.Property{
								Type:        "boolean",
								GoName:      "UsePrivateEndpoint",
								Description: "Only allow access to the master's private endpoint IP.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
