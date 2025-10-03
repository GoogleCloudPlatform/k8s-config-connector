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

func DCLWorkloadIdentityPoolSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iam/WorkloadIdentityPool",
			Description: "The Iam WorkloadIdentityPool resource",
			StructName:  "WorkloadIdentityPool",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a WorkloadIdentityPool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workloadIdentityPool",
						Required:    true,
						Description: "A full instance of a WorkloadIdentityPool",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a WorkloadIdentityPool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workloadIdentityPool",
						Required:    true,
						Description: "A full instance of a WorkloadIdentityPool",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a WorkloadIdentityPool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workloadIdentityPool",
						Required:    true,
						Description: "A full instance of a WorkloadIdentityPool",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all WorkloadIdentityPool",
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
				Description: "The function used to list information about many WorkloadIdentityPool",
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
				"WorkloadIdentityPool": &dcl.Component{
					Title:           "WorkloadIdentityPool",
					ID:              "projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{name}}",
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
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "A description of the pool. Cannot exceed 256 characters.",
							},
							"disabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Disabled",
								Description: "Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "A display name for the pool. Cannot exceed 32 characters.",
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
								Description: "Output only. The resource name of the pool.",
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
								GoType:      "WorkloadIdentityPoolStateEnum",
								ReadOnly:    true,
								Description: "Output only. The state of the pool. Possible values: STATE_UNSPECIFIED, ACTIVE, DELETED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ACTIVE",
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
