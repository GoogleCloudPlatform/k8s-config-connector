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

func DCLWorkforcePoolSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iam/WorkforcePool",
			Description: "The Iam WorkforcePool resource",
			StructName:  "WorkforcePool",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a WorkforcePool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workforcePool",
						Required:    true,
						Description: "A full instance of a WorkforcePool",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a WorkforcePool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workforcePool",
						Required:    true,
						Description: "A full instance of a WorkforcePool",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a WorkforcePool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workforcePool",
						Required:    true,
						Description: "A full instance of a WorkforcePool",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all WorkforcePool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many WorkforcePool",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
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
				"WorkforcePool": &dcl.Component{
					Title:     "WorkforcePool",
					ID:        "locations/{{location}}/workforcePools/{{name}}",
					HasCreate: true,
					HasIAM:    true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"parent",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "A user-specified description of the pool. Cannot exceed 256 characters.",
							},
							"disabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Disabled",
								Description: "Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.",
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
								Description: "The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen. The prefix `gcp-` is reserved for use by Google, and may not be specified.",
								Immutable:   true,
							},
							"parent": &dcl.Property{
								Type:                "string",
								GoName:              "Parent",
								Description:         "Immutable. The resource name of the parent. Format: `organizations/{org-id}`.",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Organization",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"selfLink": &dcl.Property{
								Type:        "string",
								GoName:      "SelfLink",
								ReadOnly:    true,
								Description: "Output only. The resource name of the pool. Format: `locations/{location}/workforcePools/{workforce_pool_id}`",
								Immutable:   true,
							},
							"sessionDuration": &dcl.Property{
								Type:          "string",
								GoName:        "SessionDuration",
								Description:   "How long the Google Cloud access tokens, console sign-in sessions, and gcloud sign-in sessions from this pool are valid. Must be greater than 15 minutes (900s) and less than 12 hours (43200s). If `session_duration` is not configured, minted credentials will have a default duration of one hour (3600s).",
								ServerDefault: true,
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "WorkforcePoolStateEnum",
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
