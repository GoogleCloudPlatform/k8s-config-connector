// Copyright 2022 Google LLC. All Rights Reserved.
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

func DCLInstanceSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Apigee/Instance",
			Description: "The Apigee Instance resource",
			StructName:  "Instance",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Instance",
						Required:    true,
						Description: "A full instance of a Instance",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Instance",
						Required:    true,
						Description: "A full instance of a Instance",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "Instance",
						Required:    true,
						Description: "A full instance of a Instance",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Instance",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "apigeeorganization",
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
						Name:     "apigeeorganization",
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
					Title:         "Instance",
					ID:            "organizations/{{apigee_organization}}/instances/{{name}}",
					HasCreate:     true,
					ApplyTimeout:  4800,
					DeleteTimeout: 4800,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"location",
							"apigeeOrganization",
						},
						Properties: map[string]*dcl.Property{
							"apigeeOrganization": &dcl.Property{
								Type:        "string",
								GoName:      "ApigeeOrganization",
								Description: "The apigee organization for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Apigee/Organization",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"createdAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "CreatedAt",
								ReadOnly:    true,
								Description: "Output only. Time the instance was created in milliseconds since epoch.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. Description of the instance.",
								Immutable:   true,
							},
							"diskEncryptionKeyName": &dcl.Property{
								Type:        "string",
								GoName:      "DiskEncryptionKeyName",
								Description: "Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudkms/CryptoKey",
										Field:    "name",
									},
								},
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Optional. Display name for the instance.",
								Immutable:   true,
							},
							"host": &dcl.Property{
								Type:        "string",
								GoName:      "Host",
								ReadOnly:    true,
								Description: "Output only. Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.",
								Immutable:   true,
							},
							"lastModifiedAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "LastModifiedAt",
								ReadOnly:    true,
								Description: "Output only. Time the instance was last modified in milliseconds since epoch.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "Required. Compute Engine location where the instance resides.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. Resource ID of the instance. Values must match the regular expression ^[a-z][a-z\\-\\d]{0,30}[a-z\\d]$.",
								Immutable:   true,
							},
							"peeringCidrRange": &dcl.Property{
								Type:        "string",
								GoName:      "PeeringCidrRange",
								GoType:      "InstancePeeringCidrRangeEnum",
								Description: "Optional. Size of the CIDR block range that will be reserved by the instance. PAID apigee_organizations support SLASH_16 to SLASH_20 and defaults to SLASH_16. Evaluation organizations support only SLASH_23. Possible values: CIDR_RANGE_UNSPECIFIED, SLASH_16, SLASH_17, SLASH_18, SLASH_19, SLASH_20, SLASH_23",
								Immutable:   true,
								Enum: []string{
									"CIDR_RANGE_UNSPECIFIED",
									"SLASH_16",
									"SLASH_17",
									"SLASH_18",
									"SLASH_19",
									"SLASH_20",
									"SLASH_23",
								},
							},
							"port": &dcl.Property{
								Type:        "string",
								GoName:      "Port",
								ReadOnly:    true,
								Description: "Output only. Port number of the exposed Apigee endpoint.",
								Immutable:   true,
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "InstanceStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the instance. Values other than `ACTIVE` means the resource is not ready to use. Possible values: SNAPSHOT_STATE_UNSPECIFIED, MISSING, OK_DOCSTORE, OK_SUBMITTED, OK_EXTERNAL, DELETED",
								Immutable:   true,
								Enum: []string{
									"SNAPSHOT_STATE_UNSPECIFIED",
									"MISSING",
									"OK_DOCSTORE",
									"OK_SUBMITTED",
									"OK_EXTERNAL",
									"DELETED",
								},
							},
						},
					},
				},
			},
		},
	}
}
