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

func DCLConnectorSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "VPCAccess/Connector",
			Description: "The VPCAccess Connector resource",
			StructName:  "Connector",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Connector",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "connector",
						Required:    true,
						Description: "A full instance of a Connector",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Connector",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "connector",
						Required:    true,
						Description: "A full instance of a Connector",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Connector",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "connector",
						Required:    true,
						Description: "A full instance of a Connector",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Connector",
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
				Description: "The function used to list information about many Connector",
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
				"Connector": &dcl.Component{
					Title:           "Connector",
					ID:              "projects/{{project}}/locations/{{location}}/connectors/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"connectedProjects": &dcl.Property{
								Type:        "array",
								GoName:      "ConnectedProjects",
								ReadOnly:    true,
								Description: "Output only. List of projects using the connector.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"ipCidrRange": &dcl.Property{
								Type:        "string",
								GoName:      "IPCidrRange",
								Description: "The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"machineType": &dcl.Property{
								Type:          "string",
								GoName:        "MachineType",
								Description:   "Machine type of VM Instance underlying connector. Default is e2-micro",
								Immutable:     true,
								ServerDefault: true,
							},
							"maxInstances": &dcl.Property{
								Type:          "integer",
								Format:        "int64",
								GoName:        "MaxInstances",
								Description:   "Maximum value of instances in autoscaling group underlying the connector.",
								Immutable:     true,
								ServerDefault: true,
							},
							"maxThroughput": &dcl.Property{
								Type:          "integer",
								Format:        "int64",
								GoName:        "MaxThroughput",
								Description:   "Maximum throughput of the connector in Mbps. Default is 200, max is 1000.",
								Immutable:     true,
								ServerDefault: true,
							},
							"minInstances": &dcl.Property{
								Type:          "integer",
								Format:        "int64",
								GoName:        "MinInstances",
								Description:   "Minimum value of instances in autoscaling group underlying the connector.",
								Immutable:     true,
								ServerDefault: true,
							},
							"minThroughput": &dcl.Property{
								Type:          "integer",
								Format:        "int64",
								GoName:        "MinThroughput",
								Description:   "Minimum throughput of the connector in Mbps. Default and min is 200.",
								Immutable:     true,
								ServerDefault: true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The resource name in the format `projects/*/locations/*/connectors/*`.",
								Immutable:   true,
							},
							"network": &dcl.Property{
								Type:        "string",
								GoName:      "Network",
								Description: "Name of a VPC network.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Compute/Network",
										Field:    "name",
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
								GoType:      "ConnectorStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the VPC access connector. Possible values: STATE_UNSPECIFIED, READY, CREATING, DELETING, ERROR, UPDATING",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"READY",
									"CREATING",
									"DELETING",
									"ERROR",
									"UPDATING",
								},
							},
							"subnet": &dcl.Property{
								Type:        "object",
								GoName:      "Subnet",
								GoType:      "ConnectorSubnet",
								Description: "The subnet in which to house the VPC Access Connector.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"name": &dcl.Property{
										Type:        "string",
										GoName:      "Name",
										Description: "Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be: {subnetName}",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Compute/Subnetwork",
												Field:    "name",
												Parent:   true,
											},
										},
									},
									"projectId": &dcl.Property{
										Type:        "string",
										GoName:      "ProjectId",
										Description: "Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Cloudresourcemanager/Project",
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
	}
}
