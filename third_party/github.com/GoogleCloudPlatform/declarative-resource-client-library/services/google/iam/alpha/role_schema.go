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

func DCLRoleSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iam/Role",
			Description: "The Iam Role resource",
			StructName:  "Role",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Role",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "role",
						Required:    true,
						Description: "A full instance of a Role",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Role",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "role",
						Required:    true,
						Description: "A full instance of a Role",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Role",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "role",
						Required:    true,
						Description: "A full instance of a Role",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Role",
				Parameters: []dcl.PathParameters{
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
				Description: "The function used to list information about many Role",
				Parameters: []dcl.PathParameters{
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
				"Role": &dcl.Component{
					Title:     "Role",
					ID:        "{{parent}}/roles/{{name}}",
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Properties: map[string]*dcl.Property{
							"deleted": &dcl.Property{
								Type:        "boolean",
								GoName:      "Deleted",
								Description: "The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. A human-readable description for the role.",
								Immutable:   true,
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								Description: "Used to perform a consistent read-modify-write.",
								Immutable:   true,
							},
							"groupName": &dcl.Property{
								Type:      "string",
								GoName:    "GroupName",
								Immutable: true,
							},
							"groupTitle": &dcl.Property{
								Type:      "string",
								GoName:    "GroupTitle",
								Immutable: true,
							},
							"includedPermissions": &dcl.Property{
								Type:        "array",
								GoName:      "IncludedPermissions",
								Description: "The names of the permissions this role grants when bound in an IAM policy.",
								Immutable:   true,
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"includedRoles": &dcl.Property{
								Type:      "array",
								GoName:    "IncludedRoles",
								Immutable: true,
								SendEmpty: true,
								ListType:  "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"lifecyclePhase": &dcl.Property{
								Type:      "string",
								GoName:    "LifecyclePhase",
								Immutable: true,
							},
							"localizedValues": &dcl.Property{
								Type:      "object",
								GoName:    "LocalizedValues",
								GoType:    "RoleLocalizedValues",
								Immutable: true,
								Properties: map[string]*dcl.Property{
									"localizedDescription": &dcl.Property{
										Type:        "string",
										GoName:      "LocalizedDescription",
										Description: "Will be English by default or if an error occurred during translation.",
										Immutable:   true,
									},
									"localizedTitle": &dcl.Property{
										Type:        "string",
										GoName:      "LocalizedTitle",
										Description: "Will be English by default or if an error occurred during translation.",
										Immutable:   true,
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.",
								Immutable:   true,
							},
							"parent": &dcl.Property{
								Type:                "string",
								GoName:              "Parent",
								Description:         "The parent parameter's value depends on the target resource for the request, namely projects or organizations. Each resource type's parent value format is described below: projects.roles.create(): projects/{PROJECT_ID}. This method creates project-level custom roles. Example request URL: https://iam.googleapis.com/v1/projects/{PROJECT_ID}/roles organizations.roles.create(): organizations/{ORGANIZATION_ID}. This method creates organization-level custom roles. Example request URL: https://iam.googleapis.com/v1/organizations/{ORGANIZATION_ID}/roles Note: Wildcard (*) values are invalid; you must specify a complete project ID or organization ID. Authorization requires the following IAM permission on the specified resource parent: iam.roles.create",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Organization",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"stage": &dcl.Property{
								Type:        "string",
								GoName:      "Stage",
								GoType:      "RoleStageEnum",
								Description: "The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.",
								Immutable:   true,
								Enum: []string{
									"ALPHA",
									"BETA",
									"GA",
									"DEPRECATED",
									"DISABLED",
									"EAP",
								},
							},
							"title": &dcl.Property{
								Type:        "string",
								GoName:      "Title",
								Description: "Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
