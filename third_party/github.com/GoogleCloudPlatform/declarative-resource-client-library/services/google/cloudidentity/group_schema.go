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
package cloudidentity

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLGroupSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Cloudidentity/Group",
			Description: "The Cloudidentity Group resource",
			StructName:  "Group",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Group",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "group",
						Required:    true,
						Description: "A full instance of a Group",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Group",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "group",
						Required:    true,
						Description: "A full instance of a Group",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Group",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "group",
						Required:    true,
						Description: "A full instance of a Group",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Group",
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
				Description: "The function used to list information about many Group",
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
				"Group": &dcl.Component{
					Title:         "Group",
					ID:            "groups/{{name}}",
					UsesStateHint: true,
					HasCreate:     true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"groupKey",
							"parent",
							"labels",
						},
						Properties: map[string]*dcl.Property{
							"additionalGroupKeys": &dcl.Property{
								Type:        "array",
								GoName:      "AdditionalGroupKeys",
								Description: "Optional. Additional entity key aliases for a Group.",
								Immutable:   true,
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "GroupAdditionalGroupKeys",
									Required: []string{
										"id",
									},
									Properties: map[string]*dcl.Property{
										"id": &dcl.Property{
											Type:        "string",
											GoName:      "Id",
											Description: "The ID of the entity. For Google-managed entities, the `id` must be the email address of a group or user. For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.",
											Immutable:   true,
										},
										"namespace": &dcl.Property{
											Type:        "string",
											GoName:      "Namespace",
											Description: "The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}`.",
											Immutable:   true,
										},
									},
								},
								Unreadable: true,
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time when the `Group` was created.",
								Immutable:   true,
							},
							"derivedAliases": &dcl.Property{
								Type:        "array",
								GoName:      "DerivedAliases",
								ReadOnly:    true,
								Description: "Output only. Aliases for groups derived from domain aliases.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "GroupDerivedAliases",
									Required: []string{
										"id",
									},
									Properties: map[string]*dcl.Property{
										"id": &dcl.Property{
											Type:        "string",
											GoName:      "Id",
											Description: "The ID of the entity. For Google-managed entities, the `id` must be the email address of a group or user. For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.",
											Immutable:   true,
										},
										"namespace": &dcl.Property{
											Type:        "string",
											GoName:      "Namespace",
											Description: "The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}`.",
											Immutable:   true,
										},
									},
								},
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.",
							},
							"directMemberCount": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "DirectMemberCount",
								ReadOnly:    true,
								Description: "Output only. The number of all direct members. Including groups and users, The special member: all-user-in-domain will be counted as one member. Output only.",
								Immutable:   true,
							},
							"directMemberCountPerType": &dcl.Property{
								Type:        "object",
								GoName:      "DirectMemberCountPerType",
								GoType:      "GroupDirectMemberCountPerType",
								ReadOnly:    true,
								Description: "Output only. Direct membership counts grouped by user/group type",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"groupCount": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "GroupCount",
										ReadOnly:    true,
										Description: "Output only. Direct group type membership count",
										Immutable:   true,
									},
									"userCount": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "UserCount",
										ReadOnly:    true,
										Description: "Output only. Direct user type membership count",
										Immutable:   true,
									},
								},
							},
							"displayName": &dcl.Property{
								Type:          "string",
								GoName:        "DisplayName",
								Description:   "The display name of the `Group`.",
								ServerDefault: true,
							},
							"dynamicGroupMetadata": &dcl.Property{
								Type:        "object",
								GoName:      "DynamicGroupMetadata",
								GoType:      "GroupDynamicGroupMetadata",
								Description: "Optional. Dynamic group metadata like queries and status.",
								Properties: map[string]*dcl.Property{
									"queries": &dcl.Property{
										Type:        "array",
										GoName:      "Queries",
										Description: "Only one entry is supported for now. Memberships will be the union of all queries. Customers can create up to 100 dynamic groups.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "GroupDynamicGroupMetadataQueries",
											Properties: map[string]*dcl.Property{
												"query": &dcl.Property{
													Type:        "string",
													GoName:      "Query",
													Description: "Query that determines the memberships of the dynamic group.",
												},
												"resourceType": &dcl.Property{
													Type:        "string",
													GoName:      "ResourceType",
													GoType:      "GroupDynamicGroupMetadataQueriesResourceTypeEnum",
													Description: " Possible values: RESOURCE_TYPE_UNSPECIFIED, USER",
													Enum: []string{
														"RESOURCE_TYPE_UNSPECIFIED",
														"USER",
													},
												},
											},
										},
									},
									"status": &dcl.Property{
										Type:        "object",
										GoName:      "Status",
										GoType:      "GroupDynamicGroupMetadataStatus",
										ReadOnly:    true,
										Description: "Status of the dynamic group. Output only.",
										Properties: map[string]*dcl.Property{
											"status": &dcl.Property{
												Type:        "string",
												GoName:      "Status",
												GoType:      "GroupDynamicGroupMetadataStatusStatusEnum",
												Description: "Status of the dynamic group. Possible values: STATUS_UNSPECIFIED, UP_TO_DATE, UPDATING_MEMBERSHIPS, INVALID_QUERY",
												Enum: []string{
													"STATUS_UNSPECIFIED",
													"UP_TO_DATE",
													"UPDATING_MEMBERSHIPS",
													"INVALID_QUERY",
												},
											},
											"statusTime": &dcl.Property{
												Type:        "string",
												Format:      "date-time",
												GoName:      "StatusTime",
												Description: "The latest time at which the dynamic group is guaranteed to be in the given status. For example, if status is: UP_TO_DATE - The latest time at which this dynamic group was confirmed to be up to date. UPDATING_MEMBERSHIPS - The time at which dynamic group was created.",
											},
										},
									},
								},
							},
							"groupKey": &dcl.Property{
								Type:        "object",
								GoName:      "GroupKey",
								GoType:      "GroupGroupKey",
								Description: "Required. Immutable. The `EntityKey` of the `Group`.",
								Immutable:   true,
								Required: []string{
									"id",
								},
								Properties: map[string]*dcl.Property{
									"id": &dcl.Property{
										Type:        "string",
										GoName:      "Id",
										Description: "The ID of the entity. For Google-managed entities, the `id` must be the email address of a group or user. For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.",
										Immutable:   true,
									},
									"namespace": &dcl.Property{
										Type:        "string",
										GoName:      "Namespace",
										Description: "The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}`.",
										Immutable:   true,
									},
								},
							},
							"initialGroupConfig": &dcl.Property{
								Type:        "string",
								GoName:      "InitialGroupConfig",
								GoType:      "GroupInitialGroupConfigEnum",
								Description: "The initial configuration option for the `Group`. Possible values: INITIAL_GROUP_CONFIG_UNSPECIFIED, WITH_INITIAL_OWNER, EMPTY",
								Immutable:   true,
								Enum: []string{
									"INITIAL_GROUP_CONFIG_UNSPECIFIED",
									"WITH_INITIAL_OWNER",
									"EMPTY",
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Required. One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value.",
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group}`.",
								ServerGeneratedParameter: true,
							},
							"parent": &dcl.Property{
								Type:        "string",
								GoName:      "Parent",
								Description: "Required. Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source}` for external- identity-mapped groups or `customers/{customer}` for Google Groups. The `customer` must begin with \"C\" (for example, 'C046psxkn').",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The time when the `Group` was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
