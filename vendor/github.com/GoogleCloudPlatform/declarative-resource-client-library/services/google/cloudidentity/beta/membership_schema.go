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

func DCLMembershipSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Cloudidentity/Membership",
			Description: "The Cloudidentity Membership resource",
			StructName:  "Membership",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "membership",
						Required:    true,
						Description: "A full instance of a Membership",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "membership",
						Required:    true,
						Description: "A full instance of a Membership",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "membership",
						Required:    true,
						Description: "A full instance of a Membership",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "group",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "group",
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
				"Membership": &dcl.Component{
					Title:     "Membership",
					ID:        "groups/{{group}}/memberships/{{name}}",
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"preferredMemberKey",
							"roles",
							"group",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time when the `Membership` was created.",
								Immutable:   true,
							},
							"deliverySetting": &dcl.Property{
								Type:        "string",
								GoName:      "DeliverySetting",
								GoType:      "MembershipDeliverySettingEnum",
								ReadOnly:    true,
								Description: "Output only. Delivery setting associated with the membership. Possible values: DELIVERY_SETTING_UNSPECIFIED, ALL_MAIL, DIGEST, DAILY, NONE, DISABLED",
								Immutable:   true,
								Enum: []string{
									"DELIVERY_SETTING_UNSPECIFIED",
									"ALL_MAIL",
									"DIGEST",
									"DAILY",
									"NONE",
									"DISABLED",
								},
							},
							"displayName": &dcl.Property{
								Type:        "object",
								GoName:      "DisplayName",
								GoType:      "MembershipDisplayName",
								ReadOnly:    true,
								Description: "Output only. The display name of this member, if available",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"familyName": &dcl.Property{
										Type:        "string",
										GoName:      "FamilyName",
										ReadOnly:    true,
										Description: "Output only. Member's family name",
										Immutable:   true,
									},
									"fullName": &dcl.Property{
										Type:        "string",
										GoName:      "FullName",
										ReadOnly:    true,
										Description: "Output only. Localized UTF-16 full name for the member. Localization is done based on the language in the request and the language of the stored display name.",
										Immutable:   true,
									},
									"givenName": &dcl.Property{
										Type:        "string",
										GoName:      "GivenName",
										ReadOnly:    true,
										Description: "Output only. Member's given name",
										Immutable:   true,
									},
								},
							},
							"group": &dcl.Property{
								Type:        "string",
								GoName:      "Group",
								Description: "The group for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudidentity/Group",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"memberKey": &dcl.Property{
								Type:          "object",
								GoName:        "MemberKey",
								GoType:        "MembershipMemberKey",
								Description:   "Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"id": &dcl.Property{
										Type:        "string",
										GoName:      "Id",
										Description: "The ID of the entity. For Google-managed entities, the `id` must be the email address of an existing group or user. For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.",
									},
									"namespace": &dcl.Property{
										Type:        "string",
										GoName:      "Namespace",
										Description: "The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}`.",
									},
								},
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group}/memberships/{membership}`.",
								ServerGeneratedParameter: true,
							},
							"preferredMemberKey": &dcl.Property{
								Type:        "object",
								GoName:      "PreferredMemberKey",
								GoType:      "MembershipPreferredMemberKey",
								Description: "Required. Immutable. The `EntityKey` of the member.",
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
							"roles": &dcl.Property{
								Type:        "array",
								GoName:      "Roles",
								Description: "The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.",
								SendEmpty:   true,
								ListType:    "set",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "MembershipRoles",
									Required: []string{
										"name",
									},
									Properties: map[string]*dcl.Property{
										"expiryDetail": &dcl.Property{
											Type:        "object",
											GoName:      "ExpiryDetail",
											GoType:      "MembershipRolesExpiryDetail",
											Description: "The expiry details of the `MembershipRole`. Expiry details are only supported for `MEMBER` `MembershipRoles`. May be set if `name` is `MEMBER`. Must not be set if `name` is any other value.",
											SendEmpty:   true,
											Properties: map[string]*dcl.Property{
												"expireTime": &dcl.Property{
													Type:        "string",
													Format:      "date-time",
													GoName:      "ExpireTime",
													Description: "The time at which the `MembershipRole` will expire.",
												},
											},
										},
										"name": &dcl.Property{
											Type:   "string",
											GoName: "Name",
										},
										"restrictionEvaluations": &dcl.Property{
											Type:        "object",
											GoName:      "RestrictionEvaluations",
											GoType:      "MembershipRolesRestrictionEvaluations",
											Description: "Evaluations of restrictions applied to parent group on this membership.",
											Properties: map[string]*dcl.Property{
												"memberRestrictionEvaluation": &dcl.Property{
													Type:        "object",
													GoName:      "MemberRestrictionEvaluation",
													GoType:      "MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation",
													Description: "Evaluation of the member restriction applied to this membership. Empty if the user lacks permission to view the restriction evaluation.",
													Properties: map[string]*dcl.Property{
														"state": &dcl.Property{
															Type:        "string",
															GoName:      "State",
															GoType:      "MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum",
															ReadOnly:    true,
															Description: "Output only. The current state of the restriction Possible values: ENCRYPTION_STATE_UNSPECIFIED, UNSUPPORTED_BY_DEVICE, ENCRYPTED, NOT_ENCRYPTED",
															Enum: []string{
																"ENCRYPTION_STATE_UNSPECIFIED",
																"UNSUPPORTED_BY_DEVICE",
																"ENCRYPTED",
																"NOT_ENCRYPTED",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							"type": &dcl.Property{
								Type:        "string",
								GoName:      "Type",
								GoType:      "MembershipTypeEnum",
								ReadOnly:    true,
								Description: "Output only. The type of the membership. Possible values: OWNER_TYPE_UNSPECIFIED, OWNER_TYPE_CUSTOMER, OWNER_TYPE_PARTNER",
								Immutable:   true,
								Enum: []string{
									"OWNER_TYPE_UNSPECIFIED",
									"OWNER_TYPE_CUSTOMER",
									"OWNER_TYPE_PARTNER",
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The time when the `Membership` was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
