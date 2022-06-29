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

func DCLCapacityCommitmentSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "BigqueryReservation/CapacityCommitment",
			Description: "The BigqueryReservation CapacityCommitment resource",
			StructName:  "CapacityCommitment",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a CapacityCommitment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "CapacityCommitment",
						Required:    true,
						Description: "A full instance of a CapacityCommitment",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a CapacityCommitment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "CapacityCommitment",
						Required:    true,
						Description: "A full instance of a CapacityCommitment",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a CapacityCommitment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "CapacityCommitment",
						Required:    true,
						Description: "A full instance of a CapacityCommitment",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all CapacityCommitment",
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
				Description: "The function used to list information about many CapacityCommitment",
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
				"CapacityCommitment": &dcl.Component{
					Title:           "CapacityCommitment",
					ID:              "projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"commitmentEndTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CommitmentEndTime",
								ReadOnly:    true,
								Description: "Output only. The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.",
								Immutable:   true,
							},
							"commitmentStartTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CommitmentStartTime",
								ReadOnly:    true,
								Description: "Output only. The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.",
								Immutable:   true,
							},
							"failureStatus": &dcl.Property{
								Type:        "object",
								GoName:      "FailureStatus",
								GoType:      "CapacityCommitmentFailureStatus",
								ReadOnly:    true,
								Description: "Output only. For FAILED commitment plan, provides the reason of failure.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"code": &dcl.Property{
										Type:        "integer",
										Format:      "int64",
										GoName:      "Code",
										Description: "The status code, which should be an enum value of google.rpc.Code.",
										Immutable:   true,
									},
									"details": &dcl.Property{
										Type:        "array",
										GoName:      "Details",
										Description: "A list of messages that carry the error details. There is a common set of message types for APIs to use.",
										Immutable:   true,
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "CapacityCommitmentFailureStatusDetails",
											Properties: map[string]*dcl.Property{
												"typeUrl": &dcl.Property{
													Type:        "string",
													GoName:      "TypeUrl",
													Description: "A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one \"/\" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading \".\" is not accepted). In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows: * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a google.protobuf.Type value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.) Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com. Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics.",
													Immutable:   true,
												},
												"value": &dcl.Property{
													Type:        "string",
													GoName:      "Value",
													Description: "Must be a valid serialized protocol buffer of the above specified type.",
													Immutable:   true,
												},
											},
										},
									},
									"message": &dcl.Property{
										Type:        "string",
										GoName:      "Message",
										Description: "A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.",
										Immutable:   true,
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123`",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"plan": &dcl.Property{
								Type:        "string",
								GoName:      "Plan",
								GoType:      "CapacityCommitmentPlanEnum",
								Description: "Capacity commitment commitment plan. Possible values: COMMITMENT_PLAN_UNSPECIFIED, FLEX, TRIAL, MONTHLY, ANNUAL",
								Enum: []string{
									"COMMITMENT_PLAN_UNSPECIFIED",
									"FLEX",
									"TRIAL",
									"MONTHLY",
									"ANNUAL",
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
							"renewalPlan": &dcl.Property{
								Type:        "string",
								GoName:      "RenewalPlan",
								GoType:      "CapacityCommitmentRenewalPlanEnum",
								Description: "The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments. Possible values: COMMITMENT_PLAN_UNSPECIFIED, FLEX, TRIAL, MONTHLY, ANNUAL",
								Enum: []string{
									"COMMITMENT_PLAN_UNSPECIFIED",
									"FLEX",
									"TRIAL",
									"MONTHLY",
									"ANNUAL",
								},
							},
							"slotCount": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "SlotCount",
								Description: "Number of slots in this commitment.",
								Immutable:   true,
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "CapacityCommitmentStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the commitment. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, FAILED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"PENDING",
									"ACTIVE",
									"FAILED",
								},
							},
						},
					},
				},
			},
		},
	}
}
